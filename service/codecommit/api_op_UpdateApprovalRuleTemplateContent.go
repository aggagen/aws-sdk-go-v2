// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateApprovalRuleTemplateContentInput struct {
	_ struct{} `type:"structure"`

	// The name of the approval rule template where you want to update the content
	// of the rule.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`

	// The SHA-256 hash signature for the content of the approval rule. You can
	// retrieve this information by using GetPullRequest.
	ExistingRuleContentSha256 *string `locationName:"existingRuleContentSha256" type:"string"`

	// The content that replaces the existing content of the rule. Content statements
	// must be complete. You cannot provide only the changes.
	//
	// NewRuleContent is a required field
	NewRuleContent *string `locationName:"newRuleContent" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateContentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApprovalRuleTemplateContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApprovalRuleTemplateContentInput"}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if s.NewRuleContent == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewRuleContent"))
	}
	if s.NewRuleContent != nil && len(*s.NewRuleContent) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NewRuleContent", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateApprovalRuleTemplateContentOutput struct {
	_ struct{} `type:"structure"`

	// Returns information about an approval rule template.
	//
	// ApprovalRuleTemplate is a required field
	ApprovalRuleTemplate *ApprovalRuleTemplate `locationName:"approvalRuleTemplate" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApprovalRuleTemplateContentOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateApprovalRuleTemplateContent = "UpdateApprovalRuleTemplateContent"

// UpdateApprovalRuleTemplateContentRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Updates the content of an approval rule template. You can change the number
// of required approvals, the membership of the approval rule, and whether an
// approval pool is defined.
//
//    // Example sending a request using UpdateApprovalRuleTemplateContentRequest.
//    req := client.UpdateApprovalRuleTemplateContentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/UpdateApprovalRuleTemplateContent
func (c *Client) UpdateApprovalRuleTemplateContentRequest(input *UpdateApprovalRuleTemplateContentInput) UpdateApprovalRuleTemplateContentRequest {
	op := &aws.Operation{
		Name:       opUpdateApprovalRuleTemplateContent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateApprovalRuleTemplateContentInput{}
	}

	req := c.newRequest(op, input, &UpdateApprovalRuleTemplateContentOutput{})

	return UpdateApprovalRuleTemplateContentRequest{Request: req, Input: input, Copy: c.UpdateApprovalRuleTemplateContentRequest}
}

// UpdateApprovalRuleTemplateContentRequest is the request type for the
// UpdateApprovalRuleTemplateContent API operation.
type UpdateApprovalRuleTemplateContentRequest struct {
	*aws.Request
	Input *UpdateApprovalRuleTemplateContentInput
	Copy  func(*UpdateApprovalRuleTemplateContentInput) UpdateApprovalRuleTemplateContentRequest
}

// Send marshals and sends the UpdateApprovalRuleTemplateContent API request.
func (r UpdateApprovalRuleTemplateContentRequest) Send(ctx context.Context) (*UpdateApprovalRuleTemplateContentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApprovalRuleTemplateContentResponse{
		UpdateApprovalRuleTemplateContentOutput: r.Request.Data.(*UpdateApprovalRuleTemplateContentOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApprovalRuleTemplateContentResponse is the response type for the
// UpdateApprovalRuleTemplateContent API operation.
type UpdateApprovalRuleTemplateContentResponse struct {
	*UpdateApprovalRuleTemplateContentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApprovalRuleTemplateContent request.
func (r *UpdateApprovalRuleTemplateContentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
