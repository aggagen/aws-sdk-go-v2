// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateDashboardPublishedVersionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the dashboard that you're updating.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" min:"1" type:"string" required:"true"`

	// The version number of the dashboard.
	//
	// VersionNumber is a required field
	VersionNumber *int64 `location:"uri" locationName:"VersionNumber" min:"1" type:"long" required:"true"`
}

// String returns the string representation
func (s UpdateDashboardPublishedVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDashboardPublishedVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDashboardPublishedVersionInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if s.DashboardId != nil && len(*s.DashboardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardId", 1))
	}

	if s.VersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionNumber"))
	}
	if s.VersionNumber != nil && *s.VersionNumber < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("VersionNumber", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDashboardPublishedVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

type UpdateDashboardPublishedVersionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string `type:"string"`

	// The ID for the dashboard.
	DashboardId *string `min:"1" type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s UpdateDashboardPublishedVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDashboardPublishedVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DashboardArn != nil {
		v := *s.DashboardArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DashboardArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateDashboardPublishedVersion = "UpdateDashboardPublishedVersion"

// UpdateDashboardPublishedVersionRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates the published version of a dashboard.
//
//    // Example sending a request using UpdateDashboardPublishedVersionRequest.
//    req := client.UpdateDashboardPublishedVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateDashboardPublishedVersion
func (c *Client) UpdateDashboardPublishedVersionRequest(input *UpdateDashboardPublishedVersionInput) UpdateDashboardPublishedVersionRequest {
	op := &aws.Operation{
		Name:       opUpdateDashboardPublishedVersion,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}/versions/{VersionNumber}",
	}

	if input == nil {
		input = &UpdateDashboardPublishedVersionInput{}
	}

	req := c.newRequest(op, input, &UpdateDashboardPublishedVersionOutput{})

	return UpdateDashboardPublishedVersionRequest{Request: req, Input: input, Copy: c.UpdateDashboardPublishedVersionRequest}
}

// UpdateDashboardPublishedVersionRequest is the request type for the
// UpdateDashboardPublishedVersion API operation.
type UpdateDashboardPublishedVersionRequest struct {
	*aws.Request
	Input *UpdateDashboardPublishedVersionInput
	Copy  func(*UpdateDashboardPublishedVersionInput) UpdateDashboardPublishedVersionRequest
}

// Send marshals and sends the UpdateDashboardPublishedVersion API request.
func (r UpdateDashboardPublishedVersionRequest) Send(ctx context.Context) (*UpdateDashboardPublishedVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDashboardPublishedVersionResponse{
		UpdateDashboardPublishedVersionOutput: r.Request.Data.(*UpdateDashboardPublishedVersionOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDashboardPublishedVersionResponse is the response type for the
// UpdateDashboardPublishedVersion API operation.
type UpdateDashboardPublishedVersionResponse struct {
	*UpdateDashboardPublishedVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDashboardPublishedVersion request.
func (r *UpdateDashboardPublishedVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
