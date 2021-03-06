// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotdataplane

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DeleteThingShadow operation.
type DeleteThingShadowInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteThingShadowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteThingShadowInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteThingShadowInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the DeleteThingShadow operation.
type DeleteThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	//
	// Payload is a required field
	Payload []byte `locationName:"payload" type:"blob" required:"true"`
}

// String returns the string representation
func (s DeleteThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteThingShadowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Payload != nil {
		v := s.Payload

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "payload", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opDeleteThingShadow = "DeleteThingShadow"

// DeleteThingShadowRequest returns a request value for making API operation for
// AWS IoT Data Plane.
//
// Deletes the thing shadow for the specified thing.
//
// For more information, see DeleteThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_DeleteThingShadow.html)
// in the AWS IoT Developer Guide.
//
//    // Example sending a request using DeleteThingShadowRequest.
//    req := client.DeleteThingShadowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteThingShadowRequest(input *DeleteThingShadowInput) DeleteThingShadowRequest {
	op := &aws.Operation{
		Name:       opDeleteThingShadow,
		HTTPMethod: "DELETE",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &DeleteThingShadowInput{}
	}

	req := c.newRequest(op, input, &DeleteThingShadowOutput{})

	return DeleteThingShadowRequest{Request: req, Input: input, Copy: c.DeleteThingShadowRequest}
}

// DeleteThingShadowRequest is the request type for the
// DeleteThingShadow API operation.
type DeleteThingShadowRequest struct {
	*aws.Request
	Input *DeleteThingShadowInput
	Copy  func(*DeleteThingShadowInput) DeleteThingShadowRequest
}

// Send marshals and sends the DeleteThingShadow API request.
func (r DeleteThingShadowRequest) Send(ctx context.Context) (*DeleteThingShadowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteThingShadowResponse{
		DeleteThingShadowOutput: r.Request.Data.(*DeleteThingShadowOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteThingShadowResponse is the response type for the
// DeleteThingShadow API operation.
type DeleteThingShadowResponse struct {
	*DeleteThingShadowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteThingShadow request.
func (r *DeleteThingShadowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
