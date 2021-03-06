// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for the get artifact request.
type GetArtifactUrlInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for a artifact.
	//
	// ArtifactId is a required field
	ArtifactId *string `location:"uri" locationName:"artifactId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetArtifactUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetArtifactUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetArtifactUrlInput"}

	if s.ArtifactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ArtifactId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetArtifactUrlInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ArtifactId != nil {
		v := *s.ArtifactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "artifactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the get artifact request.
type GetArtifactUrlOutput struct {
	_ struct{} `type:"structure"`

	// Unique Id for a artifact.
	//
	// ArtifactId is a required field
	ArtifactId *string `locationName:"artifactId" type:"string" required:"true"`

	// Presigned url for the artifact.
	//
	// ArtifactUrl is a required field
	ArtifactUrl *string `locationName:"artifactUrl" type:"string" required:"true"`
}

// String returns the string representation
func (s GetArtifactUrlOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetArtifactUrlOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ArtifactId != nil {
		v := *s.ArtifactId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "artifactId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ArtifactUrl != nil {
		v := *s.ArtifactUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "artifactUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetArtifactUrl = "GetArtifactUrl"

// GetArtifactUrlRequest returns a request value for making API operation for
// AWS Amplify.
//
// Retrieves artifact info that corresponds to a artifactId.
//
//    // Example sending a request using GetArtifactUrlRequest.
//    req := client.GetArtifactUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetArtifactUrl
func (c *Client) GetArtifactUrlRequest(input *GetArtifactUrlInput) GetArtifactUrlRequest {
	op := &aws.Operation{
		Name:       opGetArtifactUrl,
		HTTPMethod: "GET",
		HTTPPath:   "/artifacts/{artifactId}",
	}

	if input == nil {
		input = &GetArtifactUrlInput{}
	}

	req := c.newRequest(op, input, &GetArtifactUrlOutput{})

	return GetArtifactUrlRequest{Request: req, Input: input, Copy: c.GetArtifactUrlRequest}
}

// GetArtifactUrlRequest is the request type for the
// GetArtifactUrl API operation.
type GetArtifactUrlRequest struct {
	*aws.Request
	Input *GetArtifactUrlInput
	Copy  func(*GetArtifactUrlInput) GetArtifactUrlRequest
}

// Send marshals and sends the GetArtifactUrl API request.
func (r GetArtifactUrlRequest) Send(ctx context.Context) (*GetArtifactUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetArtifactUrlResponse{
		GetArtifactUrlOutput: r.Request.Data.(*GetArtifactUrlOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetArtifactUrlResponse is the response type for the
// GetArtifactUrl API operation.
type GetArtifactUrlResponse struct {
	*GetArtifactUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetArtifactUrl request.
func (r *GetArtifactUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
