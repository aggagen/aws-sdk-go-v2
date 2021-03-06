// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the GetSnapshotLimits operation.
type GetSnapshotLimitsInput struct {
	_ struct{} `type:"structure"`

	// Contains the identifier of the directory to obtain the limits for.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetSnapshotLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSnapshotLimitsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSnapshotLimitsInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the results of the GetSnapshotLimits operation.
type GetSnapshotLimitsOutput struct {
	_ struct{} `type:"structure"`

	// A SnapshotLimits object that contains the manual snapshot limits for the
	// specified directory.
	SnapshotLimits *SnapshotLimits `type:"structure"`
}

// String returns the string representation
func (s GetSnapshotLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSnapshotLimits = "GetSnapshotLimits"

// GetSnapshotLimitsRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Obtains the manual snapshot limits for a directory.
//
//    // Example sending a request using GetSnapshotLimitsRequest.
//    req := client.GetSnapshotLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/GetSnapshotLimits
func (c *Client) GetSnapshotLimitsRequest(input *GetSnapshotLimitsInput) GetSnapshotLimitsRequest {
	op := &aws.Operation{
		Name:       opGetSnapshotLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSnapshotLimitsInput{}
	}

	req := c.newRequest(op, input, &GetSnapshotLimitsOutput{})

	return GetSnapshotLimitsRequest{Request: req, Input: input, Copy: c.GetSnapshotLimitsRequest}
}

// GetSnapshotLimitsRequest is the request type for the
// GetSnapshotLimits API operation.
type GetSnapshotLimitsRequest struct {
	*aws.Request
	Input *GetSnapshotLimitsInput
	Copy  func(*GetSnapshotLimitsInput) GetSnapshotLimitsRequest
}

// Send marshals and sends the GetSnapshotLimits API request.
func (r GetSnapshotLimitsRequest) Send(ctx context.Context) (*GetSnapshotLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSnapshotLimitsResponse{
		GetSnapshotLimitsOutput: r.Request.Data.(*GetSnapshotLimitsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSnapshotLimitsResponse is the response type for the
// GetSnapshotLimits API operation.
type GetSnapshotLimitsResponse struct {
	*GetSnapshotLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSnapshotLimits request.
func (r *GetSnapshotLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
