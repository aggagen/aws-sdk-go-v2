package changes

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/internal/tools/changes/git"
	"github.com/aws/aws-sdk-go-v2/internal/tools/changes/golist"
	"github.com/aws/aws-sdk-go-v2/internal/tools/changes/util"
	"log"
	"path/filepath"
)

// Repository is a representation of a git repository containing multiple Go modules.
type Repository struct {
	RootPath string    // RootPath is the path to the root of the repository.
	Metadata *Metadata // Metadata is the repository's .changes metadata.

	Logf func(string, ...interface{})

	modules []string // modules contains the shortened module path of all modules in the repository. modules is lazily loaded by Modules()

	git    git.VcsClient
	golist golist.ModuleClient
}

// NewRepository loads the repository at the given path.
func NewRepository(path string) (*Repository, error) {
	metadata, err := LoadMetadata(filepath.Join(path, metadataDir))
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		RootPath: path,
		Metadata: metadata,
		git: git.Client{
			RepoPath: path,
		},
		golist: golist.Client{
			RootPath:       path,
			ShortenModPath: shortenModPath,
		},
		Logf: log.Printf,
	}

	return repo, nil
}

func (r *Repository) Modules() ([]string, error) {
	if r.modules != nil {
		return r.modules, nil
	}

	mods, _, err := discoverModules(r.golist, r.RootPath)
	if err != nil {
		return nil, err
	}

	r.modules = mods
	return mods, err
}

func (r *Repository) DoRelease(releaseID string) error {
	//if err := r.Metadata.CurrentVersions.isValid(r.RootPath); err != nil {
	//	return fmt.Errorf("couldn't create a release: %v", err)
	//}

	enc, bumps, err := r.DiscoverVersions(ReleaseVersionSelector)
	if err != nil {
		return fmt.Errorf("couldn't do release: %v", err)
	}

	err = r.resolveDependencies(&enc, bumps) // update all dependencies on SDK modules to their latest versions.
	if err != nil {
		return err
	}

	rel, err := r.Metadata.CreateRelease(releaseID, bumps, false)
	if err != nil {
		return err
	}

	err = r.Metadata.SaveEnclosure(enc)
	if err != nil {
		return err
	}

	err = r.UpdateAllChangelogs(rel, false)
	if err != nil {
		return err
	}

	return r.tagAndPush(releaseID, bumps)
}

func (r *Repository) tagAndPush(releaseID string, bumps map[string]VersionBump) error {
	err := r.git.Commit([]string{"."})
	if err != nil {
		return err
	}

	for mod, bump := range bumps {
		r.logf("tagging module %s: %s -> %s\n", mod, bump.From, bump.To)
		err = tagRepo(r.git, releaseID, mod, bump.To)
		if err != nil {
			return err
		}
	}

	return r.git.Push()
}

func (r *Repository) resolveDependencies(enc *VersionEnclosure, bumps map[string]VersionBump) error {
	updates, err := r.updateEnclosure(enc, bumps)
	if err != nil {
		return err
	}

	for m, deps := range updates {
		err = updateDependencies(r.RootPath, m, deps, enc)
		if err != nil {
			return err
		}
	}

	return nil
}

// updateEnclosure finds all necessary dependency updates, updating the given VersionEnclosure and VersionBumps. updateEnclosure
// returns a mapping between modules and a list of that module's dependencies that must be updated to the latest version.
// updateEnclosure also creates change metadata for the dependency updates.
func (r *Repository) updateEnclosure(enc *VersionEnclosure, bumps map[string]VersionBump) (map[string][]string, error) {
	mods, err := r.Modules()
	if err != nil {
		return nil, fmt.Errorf("couldn't resolve dependencies: %v", err)
	}

	depGraph, err := moduleGraph(r.golist, mods)
	if err != nil {
		return nil, err
	}

	updates := depGraph.dependencyUpdates(AffectedModules(r.Metadata.Changes))
	if len(updates) == 0 {
		return nil, nil
	}

	var depMods []string
	for m, _ := range updates {
		if _, ok := bumps[m]; !ok {
			depMods = append(depMods, m)

			bumps[m], err = enc.bump(m, PatchBump)
			if err != nil {
				return nil, err
			}
		}
	}

	err = r.Metadata.addDependencyUpdateChange(depMods)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func (r *Repository) UpdateAllChangelogs(release *Release, pending bool) error {
	err := r.UpdateChangelog(release, pending)
	if err != nil {
		return err
	}

	for _, m := range release.AffectedModules() {
		if m == rootModule {
			// skip updating the root module, since this would overwrite the top level consolidated CHANGELOG
			continue
		}

		err = r.UpdateModuleChangelog(release, m, pending)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateChangelog generates a new CHANGELOG entry for the given release. If pending is true, the contents of
// CHANGELOG_PENDING.md will be replaced with the new entry. Otherwise, the entry will be prepended to CHANGELOG.md.
func (r *Repository) UpdateChangelog(release *Release, pending bool) error {
	entry, err := release.RenderChangelog()
	if err != nil {
		return err
	}

	return util.WriteFile([]byte(entry), filepath.Join(r.RootPath, "CHANGELOG.md"), !pending)
}

func (r *Repository) UpdateModuleChangelog(release *Release, module string, pending bool) error {
	entry, err := release.RenderChangelogForModule(module, false)
	if err != nil {
		return err
	}

	return util.WriteFile([]byte(entry), filepath.Join(r.RootPath, module, "CHANGELOG.md"), !pending)
}

// InitializeVersions creates an initial versions.json enclosure in the Repository's .changes directory. The VersionEnclosure
// created by InitializeVersions will include all modules that have a tagged version in git and their corresponding versions.
func (r *Repository) InitializeVersions() error {
	enc, _, err := r.DiscoverVersions(TaggedVersionSelector)
	if err != nil {
		return fmt.Errorf("failed to initialize versions: %v", err)
	}

	return r.Metadata.SaveEnclosure(enc)
}

// DiscoverVersions creates a VersionEnclosure containing all Go modules in the Repository. The version of each module
// is determined by the provided VersionSelector.
func (r *Repository) DiscoverVersions(selector VersionSelector) (VersionEnclosure, map[string]VersionBump, error) {
	modules, packages, err := discoverModules(r.golist, r.RootPath)
	if err != nil {
		return VersionEnclosure{}, nil, fmt.Errorf("failed to discover versions: %v", err)
	}

	enc, bumps, err := r.discoverVersions(modules, selector)
	if err != nil {
		return VersionEnclosure{}, nil, fmt.Errorf("failed to discover versions: %v", err)
	}

	enc.Packages = packages

	return enc, bumps, nil
}

func (r *Repository) discoverVersions(modules []string, selector VersionSelector) (VersionEnclosure, map[string]VersionBump, error) {
	enclosure := VersionEnclosure{
		SchemaVersion:  SchemaVersion,
		ModuleVersions: map[string]Version{},
	}

	bumps := map[string]VersionBump{}

	for _, m := range modules {
		v, incr, err := selector(r, m)
		if err != nil {
			return VersionEnclosure{}, nil, fmt.Errorf("failed to discover version of module %s: %v", m, err)
		}

		if v != "" {
			//modHash, err := goChecksum(r.RootPath, m, v)
			//if err != nil {
			//	return VersionEnclosure{}, nil, fmt.Errorf("failed to discover version of module %s: %v", m, err)
			//}

			enclosure.ModuleVersions[m] = Version{
				Module:     m,
				ImportPath: lengthenModPath(m),
				Version:    v,
				//ModuleHash: modHash,
			}
		}

		if incr != NoBump {
			bumps[m] = VersionBump{
				From: r.Metadata.CurrentVersions.ModuleVersions[m].Version,
				To:   v,
			}
		}
	}

	return enclosure, bumps, nil
}

func (r *Repository) logf(format string, a ...interface{}) {
	if r.Logf != nil {
		r.Logf(format, a...)
	}
}
