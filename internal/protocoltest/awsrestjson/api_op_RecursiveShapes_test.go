// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClient_RecursiveShapes_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *RecursiveShapesInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes recursive structures
		"RestJsonRecursiveShapes": {
			Params: &RecursiveShapesInput{
				Nested: &types.RecursiveShapesInputOutputNested1{
					Foo: ptr.String("Foo1"),
					Nested: &types.RecursiveShapesInputOutputNested2{
						Bar: ptr.String("Bar1"),
						RecursiveMember: &types.RecursiveShapesInputOutputNested1{
							Foo: ptr.String("Foo2"),
							Nested: &types.RecursiveShapesInputOutputNested2{
								Bar: ptr.String("Bar2"),
							},
						},
					},
				},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/RecursiveShapes",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "nested": {
			        "foo": "Foo1",
			        "nested": {
			            "bar": "Bar1",
			            "recursiveMember": {
			                "foo": "Foo2",
			                "nested": {
			                    "bar": "Bar2"
			                }
			            }
			        }
			    }
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					actualReq = r
					return &http.Response{
						StatusCode: 200,
						Header:     http.Header{},
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}, nil
				}),
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = "https://127.0.0.1"
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2"})
			result, err := client.RecursiveShapes(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_RecursiveShapes_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *RecursiveShapesOutput
	}{
		// Serializes recursive structures
		"RestJsonRecursiveShapes": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "nested": {
			        "foo": "Foo1",
			        "nested": {
			            "bar": "Bar1",
			            "recursiveMember": {
			                "foo": "Foo2",
			                "nested": {
			                    "bar": "Bar2"
			                }
			            }
			        }
			    }
			}`),
			ExpectResult: &RecursiveShapesOutput{
				Nested: &types.RecursiveShapesInputOutputNested1{
					Foo: ptr.String("Foo1"),
					Nested: &types.RecursiveShapesInputOutputNested2{
						Bar: ptr.String("Bar1"),
						RecursiveMember: &types.RecursiveShapesInputOutputNested1{
							Foo: ptr.String("Foo2"),
							Nested: &types.RecursiveShapesInputOutputNested2{
								Bar: ptr.String("Bar2"),
							},
						},
					},
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: c.StatusCode,
						Header:     c.Header.Clone(),
						Body:       ioutil.NopCloser(bytes.NewReader(c.Body)),
					}, nil
				}),
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Build.Clear()
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = "https://127.0.0.1"
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				Region: "us-west-2"})
			var params RecursiveShapesInput
			result, err := client.RecursiveShapes(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if diff := cmp.Diff(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); len(diff) != 0 {
				t.Errorf("expect c.ExpectResult value match:\n%s", diff)
			}
		})
	}
}
