package changes

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"sort"
	"strings"
	"time"
)

// ChangeType describes the type of change made to a Go module.
type ChangeType string

const (
	// FeatureChangeType is a constant change type for a new feature.
	FeatureChangeType ChangeType = "feature"
	// BugFixChangeType is a constant change type for a bug fix.
	BugFixChangeType ChangeType = "bugfix"
	// DependencyChangeType is a constant change type for a dependency update.
	DependencyChangeType ChangeType = "dependency"
	// AnnouncementChangeType is a constant change type for an SDK announcement.
	AnnouncementChangeType ChangeType = "announcement"
)

const dependencyUpdateMessage = "Updated SDK dependencies to their latest versions."

// ParseChangeType attempts to parse the given string v into a ChangeType, returning an error if the string is invalid.
func ParseChangeType(v string) (ChangeType, error) {
	switch strings.ToLower(v) {
	case string(FeatureChangeType):
		return FeatureChangeType, nil
	case string(BugFixChangeType):
		return BugFixChangeType, nil
	case string(DependencyChangeType):
		return DependencyChangeType, nil
	case string(AnnouncementChangeType):
		return AnnouncementChangeType, nil
	default:
		return "", fmt.Errorf("unknown change type: %s", v)
	}
}

// ChangelogPrefix returns the CHANGELOG header the ChangeType should be grouped under.
func (c ChangeType) ChangelogPrefix() string {
	switch c {
	case FeatureChangeType:
		return "Feature: "
	case BugFixChangeType:
		return "Bug Fix: "
	case DependencyChangeType:
		return "Dependency Update: "
	case AnnouncementChangeType:
		return "" // Announcements do not have a Changelog prefix.
	default:
		panic("unknown change type: " + string(c))
	}
}

// VersionIncrement returns the VersionIncrement corresponding to the given ChangeType.
func (c ChangeType) VersionIncrement() VersionIncrement {
	switch c {
	case FeatureChangeType:
		return MinorBump
	case BugFixChangeType:
		return PatchBump
	case DependencyChangeType:
		return PatchBump
	case AnnouncementChangeType:
		return NoBump
	default:
		panic("unknown change type: " + string(c))
	}
}

// String returns a string representation of the ChangeType
func (c ChangeType) String() string {
	return string(c)
}

// Set parses the given string and correspondingly sets the ChangeType, returning an error if the string could not be parsed.
func (c *ChangeType) Set(s string) error {
	p, err := ParseChangeType(s)
	if err != nil {
		return err
	}

	*c = p
	return nil
}

// UnmarshalJSON implements the encoding/json package's Unmarshaler interface, additionally providing change type
// validation.
func (c *ChangeType) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*c, err = ParseChangeType(s)
	return err
}

// UnmarshalYAML implements yaml.v2's Unmarshaler interface, additionally providing change type validation.
func (c *ChangeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	err := unmarshal(&s)
	if err != nil {
		return err
	}

	*c, err = ParseChangeType(s)
	return err
}

const changeTemplateSuffix = `
# type may be one of "feature" or "bugfix".
# multiple modules may be listed. A change metadata file will be created for each module.

# affected_modules should not be provided unless you are creating a wildcard change (by passing
# the wildcard and module flag to the add command).`

type changeTemplate struct {
	Modules         []string
	Type            ChangeType
	Description     string
	AffectedModules []string `yaml:"affected_modules"`
}

// Change represents a change to a single Go module.
type Change struct {
	ID            string     // ID is a unique identifier for this Change.
	SchemaVersion int        // SchemaVersion is the version of the library's types used to create this Change.
	Module        string     // Module is a shortened Go module path for the module affected by this Change. Module is the path from the root of the repository to the module.
	Type          ChangeType // Type indicates what category of Change was made. Type may be either "feature" or "bugfix".
	Description   string     // Description is a human readable description of this Change meant to be included in a CHANGELOG.

	// AffectedModules is a list of modules affected by this Change. AffectedModules is only non-nil when the Change's
	// Module has a wildcard, in which case AffectedModules contains all modules matching the wildcard Module that the
	// Change affects.
	AffectedModules []string
}

// NewChanges returns a Change slice containing a Change with the given type and description for each of the specified
// modules.
func NewChanges(modules []string, changeType ChangeType, description string) ([]Change, error) {
	if len(modules) == 0 || changeType == "" || description == "" {
		return nil, errors.New("missing module, type, or description")
	}

	changes := make([]Change, 0, len(modules))

	for _, module := range modules {
		module = shortenModPath(module)

		changes = append(changes, Change{
			ID:            generateID(module, changeType),
			SchemaVersion: SchemaVersion,
			Module:        module,
			Type:          changeType,
			Description:   description,
		})
	}

	for _, change := range changes {
		if change.isWildcard() {
			return nil, fmt.Errorf("module %s provided to NewChanges is a wildcard", change.Module)
		}
	}

	return changes, nil
}

// NewWildcardChange creates a wildcard Change.
func NewWildcardChange(module string, changeType ChangeType, description string, affectedModules []string) (Change, error) {
	module = shortenModPath(module)

	change := Change{
		ID:              generateID(module, changeType),
		SchemaVersion:   SchemaVersion,
		Module:          module,
		Type:            changeType,
		Description:     description,
		AffectedModules: affectedModules,
	}

	if !change.isWildcard() {
		return Change{}, fmt.Errorf("module %s is not a wildcard", module)
	}

	return change, nil
}

func (c Change) isWildcard() bool {
	return modIsWildcard(c.Module)
}

func modIsWildcard(mod string) bool {
	return strings.Contains(mod, "...")
}

// matches returns whether the Change c affects the given module.
func (c Change) matches(module string) bool {
	if !c.isWildcard() {
		return module == c.Module
	}

	prefix := strings.TrimSuffix(c.Module, "/...")
	return strings.HasPrefix(module, prefix)
}

// String returns a string representation of a Change suitable to be included in a Changelog.
func (c Change) String() string {
	return fmt.Sprintf("* %s%s", c.Type.ChangelogPrefix(), c.Description)
}

// MatchWildcardModules filters modules, returning only the modules that match the given wildcard.
func MatchWildcardModules(modules []string, wildcard string) ([]string, error) {
	var matches []string

	prefix := strings.TrimSuffix(wildcard, "/...")

	for _, m := range modules {
		if strings.HasPrefix(m, prefix) {
			matches = append(matches, m)
		}
	}

	return matches, nil
}

// TemplateToChanges parses the provided filledTemplate into the provided Change. If Change has no ID, TemplateToChange
// will set the ID.
func TemplateToChanges(filledTemplate []byte) ([]Change, error) {
	var template changeTemplate

	err := yaml.UnmarshalStrict(filledTemplate, &template)
	if err != nil {
		return nil, err
	}

	if len(template.AffectedModules) != 0 {
		if len(template.Modules) != 1 {
			return nil, fmt.Errorf("expected wildcard template to have only one module, got %d", len(template.Modules))
		}

		change, err := NewWildcardChange(template.Modules[0], template.Type, template.Description, template.AffectedModules)
		if err != nil {
			return nil, err
		}

		return []Change{change}, nil
	}

	return NewChanges(template.Modules, template.Type, template.Description)
}

// ChangeToTemplate returns a Change template populated with the given Change's data.
func ChangeToTemplate(change Change) ([]byte, error) {
	templateBytes, err := yaml.Marshal(changeTemplate{
		Modules:         []string{change.Module},
		Type:            change.Type,
		Description:     change.Description,
		AffectedModules: change.AffectedModules,
	})
	if err != nil {
		return nil, err
	}

	return append(templateBytes, []byte(changeTemplateSuffix)...), nil
}

// AffectedModules returns a sorted list of all modules affected by the given Changes. A module is considered affected if
// it is the Module of one or more Changes that will result in a version increment.
func AffectedModules(changes []Change) []string {
	var modules []string
	seen := make(map[string]struct{})

	for _, c := range changes {
		if c.Type.VersionIncrement() == NoBump {
			continue
		}

		if c.isWildcard() {
			for _, affectedModule := range c.AffectedModules {
				if _, ok := seen[affectedModule]; !ok {
					seen[affectedModule] = struct{}{}
					modules = append(modules, affectedModule)
				}
			}
		} else {
			// todo remove need for this by populating affectedmodules on non-wildcards
			if _, ok := seen[c.Module]; !ok {
				seen[c.Module] = struct{}{}
				modules = append(modules, c.Module)
			}
		}
	}

	sort.Strings(modules)
	return modules
}

func generateID(module string, changeType ChangeType) string {
	module = strings.ReplaceAll(module, "...", "wildcard")
	module = strings.ReplaceAll(module, "/", ".")

	return fmt.Sprintf("%s-%s-%v", module, changeType, time.Now().UnixNano())
}
