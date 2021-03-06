package v4

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/internal/sdk"
)

func buildRequest(serviceName, region, body string) (*http.Request, io.ReadSeeker) {

	reader := strings.NewReader(body)
	return buildRequestWithBodyReader(serviceName, region, reader)
}

func buildRequestWithBodyReader(serviceName, region string, body io.Reader) (*http.Request, io.ReadSeeker) {
	var bodyLen int

	type lenner interface {
		Len() int
	}
	if lr, ok := body.(lenner); ok {
		bodyLen = lr.Len()
	}

	endpoint := "https://" + serviceName + "." + region + ".amazonaws.com"
	req, _ := http.NewRequest("POST", endpoint, body)
	req.URL.Opaque = "//example.org/bucket/key-._~,!@#$%^&*()"
	req.Header.Set("X-Amz-Target", "prefix.Operation")
	req.Header.Set("Content-Type", "application/x-amz-json-1.0")

	if bodyLen > 0 {
		req.Header.Set("Content-Length", strconv.Itoa(bodyLen))
	}

	req.Header.Set("X-Amz-Meta-Other-Header", "some-value=!@#$%^&* (+)")
	req.Header.Add("X-Amz-Meta-Other-Header_With_Underscore", "some-value=!@#$%^&* (+)")
	req.Header.Add("X-amz-Meta-Other-Header_With_Underscore", "some-value=!@#$%^&* (+)")

	var seeker io.ReadSeeker
	if sr, ok := body.(io.ReadSeeker); ok {
		seeker = sr
	} else {
		seeker = aws.ReadSeekCloser(body)
	}

	return req, seeker
}

func buildSigner() Signer {
	return Signer{
		Credentials: aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION"),
	}
}

func removeWS(text string) string {
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", "", -1)
	return text
}

func assertEqual(t *testing.T, expected, given string) {
	if removeWS(expected) != removeWS(given) {
		t.Errorf("\nExpected: %s\nGiven:    %s", expected, given)
	}
}

func TestPresignRequest(t *testing.T) {
	req, body := buildRequest("dynamodb", "us-east-1", "{}")

	signer := buildSigner()
	signer.Presign(context.Background(), req, body, "dynamodb", "us-east-1", 300*time.Second, time.Unix(0, 0))

	expectedDate := "19700101T000000Z"
	expectedHeaders := "content-length;content-type;host;x-amz-meta-other-header;x-amz-meta-other-header_with_underscore"
	expectedSig := "122f0b9e091e4ba84286097e2b3404a1f1f4c4aad479adda95b7dff0ccbe5581"
	expectedCred := "AKID/19700101/us-east-1/dynamodb/aws4_request"
	expectedTarget := "prefix.Operation"

	q := req.URL.Query()
	if e, a := expectedSig, q.Get("X-Amz-Signature"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectedCred, q.Get("X-Amz-Credential"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectedHeaders, q.Get("X-Amz-SignedHeaders"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectedDate, q.Get("X-Amz-Date"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if a := q.Get("X-Amz-Meta-Other-Header"); len(a) != 0 {
		t.Errorf("expect %v to be empty", a)
	}
	if e, a := expectedTarget, q.Get("X-Amz-Target"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestPresignBodyWithArrayRequest(t *testing.T) {
	req, body := buildRequest("dynamodb", "us-east-1", "{}")
	req.URL.RawQuery = "Foo=z&Foo=o&Foo=m&Foo=a"

	signer := buildSigner()
	signer.Presign(context.Background(), req, body, "dynamodb", "us-east-1", 300*time.Second, time.Unix(0, 0))

	expectedDate := "19700101T000000Z"
	expectedHeaders := "content-length;content-type;host;x-amz-meta-other-header;x-amz-meta-other-header_with_underscore"
	expectedSig := "e3ac55addee8711b76c6d608d762cff285fe8b627a057f8b5ec9268cf82c08b1"
	expectedCred := "AKID/19700101/us-east-1/dynamodb/aws4_request"
	expectedTarget := "prefix.Operation"

	q := req.URL.Query()
	if e, a := expectedSig, q.Get("X-Amz-Signature"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectedCred, q.Get("X-Amz-Credential"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectedHeaders, q.Get("X-Amz-SignedHeaders"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectedDate, q.Get("X-Amz-Date"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if a := q.Get("X-Amz-Meta-Other-Header"); len(a) != 0 {
		t.Errorf("expect %v to be empty, was not", a)
	}
	if e, a := expectedTarget, q.Get("X-Amz-Target"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestSignRequest(t *testing.T) {
	req, body := buildRequest("dynamodb", "us-east-1", "{}")
	signer := buildSigner()
	signer.Sign(context.Background(), req, body, "dynamodb", "us-east-1", time.Unix(0, 0))

	expectedDate := "19700101T000000Z"
	expectedSig := "AWS4-HMAC-SHA256 Credential=AKID/19700101/us-east-1/dynamodb/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date;x-amz-meta-other-header;x-amz-meta-other-header_with_underscore;x-amz-security-token;x-amz-target, Signature=a518299330494908a70222cec6899f6f32f297f8595f6df1776d998936652ad9"

	q := req.Header
	if e, a := expectedSig, q.Get("Authorization"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := expectedDate, q.Get("X-Amz-Date"); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestSignUnseekableBody(t *testing.T) {
	req, body := buildRequestWithBodyReader("mock-service", "mock-region", bytes.NewBuffer([]byte("hello")))
	signer := buildSigner()
	_, err := signer.Sign(context.Background(), req, body, "mock-service", "mock-region", time.Now())
	if err == nil {
		t.Fatalf("expect error signing request")
	}

	if e, a := "unseekable request body", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %q to be in %q", e, a)
	}
}

func TestSignUnsignedPayloadUnseekableBody(t *testing.T) {
	req, body := buildRequestWithBodyReader("mock-service", "mock-region", bytes.NewBuffer([]byte("hello")))

	signer := buildSigner()
	signer.UnsignedPayload = true

	_, err := signer.Sign(context.Background(), req, body, "mock-service", "mock-region", time.Now())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	hash := req.Header.Get("X-Amz-Content-Sha256")
	if e, a := "UNSIGNED-PAYLOAD", hash; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestSignPreComputedHashUnseekableBody(t *testing.T) {
	req, body := buildRequestWithBodyReader("mock-service", "mock-region", bytes.NewBuffer([]byte("hello")))

	signer := buildSigner()

	req.Header.Set("X-Amz-Content-Sha256", "some-content-sha256")
	_, err := signer.Sign(context.Background(), req, body, "mock-service", "mock-region", time.Now())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	hash := req.Header.Get("X-Amz-Content-Sha256")
	if e, a := "some-content-sha256", hash; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestSignBodyS3(t *testing.T) {
	req, body := buildRequest("s3", "us-east-1", "hello")
	signer := buildSigner()
	signer.Sign(context.Background(), req, body, "s3", "us-east-1", time.Now())
	hash := req.Header.Get("X-Amz-Content-Sha256")
	if e, a := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", hash; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestSignBodyGlacier(t *testing.T) {
	req, body := buildRequest("glacier", "us-east-1", "hello")
	signer := buildSigner()
	signer.Sign(context.Background(), req, body, "glacier", "us-east-1", time.Now())
	hash := req.Header.Get("X-Amz-Content-Sha256")
	if e, a := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", hash; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestPresign_SignedPayload(t *testing.T) {
	req, body := buildRequest("glacier", "us-east-1", "hello")
	signer := buildSigner()
	signer.Presign(context.Background(), req, body, "glacier", "us-east-1", 5*time.Minute, time.Now())
	hash := req.Header.Get("X-Amz-Content-Sha256")
	if e, a := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", hash; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestPresign_UnsignedPayload(t *testing.T) {
	req, body := buildRequest("service-name", "us-east-1", "hello")
	signer := buildSigner()
	signer.UnsignedPayload = true
	signer.Presign(context.Background(), req, body, "service-name", "us-east-1", 5*time.Minute, time.Now())
	hash := req.Header.Get("X-Amz-Content-Sha256")
	if e, a := "UNSIGNED-PAYLOAD", hash; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestPresign_UnsignedPayload_S3(t *testing.T) {
	req, body := buildRequest("s3", "us-east-1", "hello")
	signer := buildSigner()
	signer.Presign(context.Background(), req, body, "s3", "us-east-1", 5*time.Minute, time.Now())
	if a := req.Header.Get("X-Amz-Content-Sha256"); len(a) != 0 {
		t.Errorf("expect no content sha256 got %v", a)
	}
}

func TestSignPrecomputedBodyChecksum(t *testing.T) {
	req, body := buildRequest("dynamodb", "us-east-1", "hello")
	req.Header.Set("X-Amz-Content-Sha256", "PRECOMPUTED")
	signer := buildSigner()
	signer.Sign(context.Background(), req, body, "dynamodb", "us-east-1", time.Now())
	hash := req.Header.Get("X-Amz-Content-Sha256")
	if e, a := "PRECOMPUTED", hash; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestAnonymousCredentials(t *testing.T) {
	cfg := unit.Config()
	cfg.Credentials = aws.AnonymousCredentials

	svc := awstesting.NewClient(cfg)
	r := svc.NewRequest(
		&aws.Operation{
			Name:       "BatchGetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)
	SignSDKRequest(r)

	urlQ := r.HTTPRequest.URL.Query()
	if a := urlQ.Get("X-Amz-Signature"); len(a) != 0 {
		t.Errorf("expect %v to be empty, was not", a)
	}
	if a := urlQ.Get("X-Amz-Credential"); len(a) != 0 {
		t.Errorf("expect %v to be empty, was not", a)
	}
	if a := urlQ.Get("X-Amz-SignedHeaders"); len(a) != 0 {
		t.Errorf("expect %v to be empty, was not", a)
	}
	if a := urlQ.Get("X-Amz-Date"); len(a) != 0 {
		t.Errorf("expect %v to be empty, was not", a)
	}

	hQ := r.HTTPRequest.Header
	if a := hQ.Get("Authorization"); len(a) != 0 {
		t.Errorf("expect %v to be empty, was not", a)
	}
	if a := hQ.Get("X-Amz-Date"); len(a) != 0 {
		t.Errorf("expect %v to be empty, was not", a)
	}
}

func TestIgnoreResignRequestWithValidCreds(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()
	mockTime := time.Time{}
	sdk.NowTime = func() time.Time { return mockTime }

	cfg := unit.Config()
	cfg.Credentials = aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION")
	cfg.Region = "us-west-2"

	svc := awstesting.NewClient(cfg)
	r := svc.NewRequest(
		&aws.Operation{
			Name:       "BatchGetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)

	SignSDKRequest(r)
	sig := r.HTTPRequest.Header.Get("Authorization")

	mockTime = mockTime.Add(1 * time.Second)
	SignSDKRequest(r)
	if e, a := sig, r.HTTPRequest.Header.Get("Authorization"); e == a {
		t.Errorf("expect %v to be %v, but was not", e, a)
	}
}

func TestIgnorePreResignRequestWithValidCreds(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()
	mockTime := time.Time{}
	sdk.NowTime = func() time.Time { return mockTime }

	cfg := unit.Config()
	cfg.Credentials = aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION")
	cfg.Region = "us-west-2"

	svc := awstesting.NewClient(cfg)
	r := svc.NewRequest(
		&aws.Operation{
			Name:       "BatchGetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)
	r.ExpireTime = time.Minute * 10

	SignSDKRequest(r)
	sig := r.HTTPRequest.URL.Query().Get("X-Amz-Signature")

	mockTime = mockTime.Add(1 * time.Second)
	SignSDKRequest(r)
	if e, a := sig, r.HTTPRequest.URL.Query().Get("X-Amz-Signature"); e == a {
		t.Errorf("expect %v to be %v, but was not", e, a)
	}
}

func TestResignRequestExpiredCreds(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()
	mockTime := time.Time{}
	sdk.NowTime = func() time.Time { return mockTime }

	creds := func() awstesting.MockCredentialsProvider {
		creds := aws.Credentials{
			AccessKeyID:     "expiredKey",
			SecretAccessKey: "expiredSecret",
		}
		return awstesting.MockCredentialsProvider{
			RetrieveFn: func(ctx context.Context) (aws.Credentials, error) {
				return creds, nil
			},
			InvalidateFn: func() {
				creds = aws.Credentials{
					AccessKeyID:     "AKID",
					SecretAccessKey: "SECRET",
				}
			},
		}
	}()

	cfg := unit.Config()
	cfg.Credentials = creds

	svc := awstesting.NewClient(cfg)
	r := svc.NewRequest(
		&aws.Operation{
			Name:       "BatchGetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)
	SignSDKRequest(r)
	querySig := r.HTTPRequest.Header.Get("Authorization")
	var origSignedHeaders string
	for _, p := range strings.Split(querySig, ", ") {
		if strings.HasPrefix(p, "SignedHeaders=") {
			origSignedHeaders = p[len("SignedHeaders="):]
			break
		}
	}
	if a := origSignedHeaders; len(a) == 0 {
		t.Errorf("expect not to be empty, but was")
	}
	if e, a := origSignedHeaders, "authorization"; strings.Contains(a, e) {
		t.Errorf("expect %v to not be in %v, but was", e, a)
	}
	origSignedAt := r.LastSignedAt

	creds.Invalidate()

	mockTime = mockTime.Add(1 * time.Second)
	SignSDKRequest(r)
	updatedQuerySig := r.HTTPRequest.Header.Get("Authorization")
	if e, a := querySig, updatedQuerySig; e == a {
		t.Errorf("expect %v to be %v, was not", e, a)
	}

	var updatedSignedHeaders string
	for _, p := range strings.Split(updatedQuerySig, ", ") {
		if strings.HasPrefix(p, "SignedHeaders=") {
			updatedSignedHeaders = p[len("SignedHeaders="):]
			break
		}
	}
	if a := updatedSignedHeaders; len(a) == 0 {
		t.Errorf("expect not to be empty, but was")
	}
	if e, a := updatedQuerySig, "authorization"; strings.Contains(a, e) {
		t.Errorf("expect %v to not be in %v, but was", e, a)
	}
	if e, a := origSignedAt, r.LastSignedAt; e == a {
		t.Errorf("expect %v to be %v, was not", e, a)
	}
}

func TestPreResignRequestExpiredCreds(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()
	mockTime := time.Time{}
	sdk.NowTime = func() time.Time { return mockTime }

	creds := func() awstesting.MockCredentialsProvider {
		creds := aws.Credentials{
			AccessKeyID:     "expiredKey",
			SecretAccessKey: "expiredSecret",
		}
		return awstesting.MockCredentialsProvider{
			RetrieveFn: func(ctx context.Context) (aws.Credentials, error) {
				return creds, nil
			},
			InvalidateFn: func() {
				creds = aws.Credentials{
					AccessKeyID:     "AKID",
					SecretAccessKey: "SECRET",
				}
			},
		}
	}()

	cfg := unit.Config()
	cfg.Credentials = creds

	svc := awstesting.NewClient(cfg)
	r := svc.NewRequest(
		&aws.Operation{
			Name:       "BatchGetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)
	r.ExpireTime = time.Minute * 10

	SignSDKRequest(r)
	querySig := r.HTTPRequest.URL.Query().Get("X-Amz-Signature")
	signedHeaders := r.HTTPRequest.URL.Query().Get("X-Amz-SignedHeaders")
	if a := signedHeaders; len(a) == 0 {
		t.Errorf("expect not to be empty, but was")
	}
	origSignedAt := r.LastSignedAt

	creds.Invalidate()

	// Simulate the request occurred 48 hours in the future
	mockTime = mockTime.Add(-48 * time.Hour)
	SignSDKRequest(r)
	if e, a := querySig, r.HTTPRequest.URL.Query().Get("X-Amz-Signature"); e == a {
		t.Errorf("expect %v to be %v, was not", e, a)
	}
	resignedHeaders := r.HTTPRequest.URL.Query().Get("X-Amz-SignedHeaders")
	if e, a := signedHeaders, resignedHeaders; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := signedHeaders, "x-amz-signedHeaders"; strings.Contains(a, e) {
		t.Errorf("expect %v to not be in %v, but was", e, a)
	}
	if e, a := origSignedAt, r.LastSignedAt; e == a {
		t.Errorf("expect %v to be %v, was not", e, a)
	}
}

func TestResignRequestExpiredRequest(t *testing.T) {
	orig := sdk.NowTime
	defer func() { sdk.NowTime = orig }()
	mockTime := time.Time{}
	sdk.NowTime = func() time.Time { return mockTime }

	creds := aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION")

	cfg := unit.Config()
	cfg.Credentials = creds

	svc := awstesting.NewClient(cfg)
	r := svc.NewRequest(
		&aws.Operation{
			Name:       "BatchGetItem",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		},
		nil,
		nil,
	)

	SignSDKRequest(r)
	querySig := r.HTTPRequest.Header.Get("Authorization")
	origSignedAt := r.LastSignedAt

	// Simulate the request occurred 15 minutes in the past
	mockTime = mockTime.Add(15 * time.Minute)
	SignSDKRequest(r)
	if e, a := querySig, r.HTTPRequest.Header.Get("Authorization"); e == a {
		t.Errorf("expected %v, got %v", e, a)
	}
	if e, a := origSignedAt, r.LastSignedAt; e == a {
		t.Errorf("expect %v to be %v, was not", e, a)
	}
}

func TestSignWithRequestBody(t *testing.T) {
	creds := aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION")
	signer := NewSigner(creds)

	expectBody := []byte("abc123")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
		if e, a := expectBody, b; !reflect.DeepEqual(e, a) {
			t.Errorf("expect %v, got %v", e, a)
		}
		w.WriteHeader(http.StatusOK)
	}))

	req, err := http.NewRequest("POST", server.URL, nil)

	_, err = signer.Sign(context.Background(), req, bytes.NewReader(expectBody), "service", "region", time.Now())
	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}
	if e, a := http.StatusOK, resp.StatusCode; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestSignWithRequestBody_Overwrite(t *testing.T) {
	creds := aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION")
	signer := NewSigner(creds)

	var expectBody []byte

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			t.Errorf("expect not no error, got %v", err)
		}
		if e, a := len(expectBody), len(b); e != a {
			t.Errorf("expect %v, got %v", e, a)
		}
		w.WriteHeader(http.StatusOK)
	}))

	req, err := http.NewRequest("GET", server.URL, strings.NewReader("invalid body"))

	_, err = signer.Sign(context.Background(), req, nil, "service", "region", time.Now())
	req.ContentLength = 0

	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expect not no error, got %v", err)
	}
	if e, a := http.StatusOK, resp.StatusCode; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestBuildCanonicalRequest(t *testing.T) {
	req, _ := buildRequest("dynamodb", "us-east-1", "{}")
	req.URL.RawQuery = "Foo=z&Foo=o&Foo=m&Foo=a"
	ctx := &httpSigner{
		ServiceName: "dynamodb",
		Region:      "us-east-1",
		Request:     req,
		Time:        time.Now(),
		ExpireTime:  5 * time.Second,
	}

	build, err := ctx.Build()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := "https://example.org/bucket/key-._~,!@#$%^&*()?Foo=a&Foo=m&Foo=o&Foo=z"
	if e, a := expected, build.Request.URL.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestSignWithBody_ReplaceRequestBody(t *testing.T) {
	creds := aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION")
	req, seekerBody := buildRequest("dynamodb", "us-east-1", "{}")
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{}))

	s := NewSigner(creds)
	origBody := req.Body

	_, err := s.Sign(context.Background(), req, seekerBody, "dynamodb", "us-east-1", time.Now())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if req.Body == origBody {
		t.Errorf("expeect request body to not be origBody")
	}

	if req.Body == nil {
		t.Errorf("expect request body to be changed but was nil")
	}
}

func TestSignWithBody_NoReplaceRequestBody(t *testing.T) {
	creds := aws.NewStaticCredentialsProvider("AKID", "SECRET", "SESSION")
	req, seekerBody := buildRequest("dynamodb", "us-east-1", "{}")
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte{}))

	s := NewSigner(creds, func(signer *Signer) {
		signer.DisableRequestBodyOverwrite = true
	})

	origBody := req.Body

	_, err := s.Sign(context.Background(), req, seekerBody, "dynamodb", "us-east-1", time.Now())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if req.Body != origBody {
		t.Errorf("expect request body to not be chagned")
	}
}

func TestRequestHost(t *testing.T) {
	req, _ := buildRequest("dynamodb", "us-east-1", "{}")
	req.URL.RawQuery = "Foo=z&Foo=o&Foo=m&Foo=a"
	req.Host = "myhost"
	ctx := &httpSigner{
		ServiceName: "dynamodb",
		Region:      "us-east-1",
		Request:     req,
		Time:        time.Now(),
		ExpireTime:  5 * time.Second,
	}

	build, err := ctx.Build()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !strings.Contains(build.CanonicalString, "host:"+req.Host) {
		t.Errorf("canonical host header invalid")
	}
}

func BenchmarkPresignRequest(b *testing.B) {
	signer := buildSigner()
	req, body := buildRequest("dynamodb", "us-east-1", "{}")
	for i := 0; i < b.N; i++ {
		signer.Presign(context.Background(), req, body, "dynamodb", "us-east-1", 300*time.Second, time.Now())
	}
}

func BenchmarkSignRequest(b *testing.B) {
	signer := buildSigner()
	req, body := buildRequest("dynamodb", "us-east-1", "{}")
	for i := 0; i < b.N; i++ {
		signer.Sign(context.Background(), req, body, "dynamodb", "us-east-1", time.Now())
	}
}
