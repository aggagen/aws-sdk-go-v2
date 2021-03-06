// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sso

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListAccountRolesInput struct {
	_ struct{} `type:"structure"`

	// The token issued by the CreateToken API call. For more information, see CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the AWS SSO OIDC API Reference Guide.
	//
	// AccessToken is a required field
	AccessToken *string `location:"header" locationName:"x-amz-sso_bearer_token" type:"string" required:"true" sensitive:"true"`

	// The identifier for the AWS account that is assigned to the user.
	//
	// AccountId is a required field
	AccountId *string `location:"querystring" locationName:"account_id" type:"string" required:"true"`

	// The number of items that clients can request per page.
	MaxResults *int64 `location:"querystring" locationName:"max_result" min:"1" type:"integer"`

	// The page token from the previous response output when you request subsequent
	// pages.
	NextToken *string `location:"querystring" locationName:"next_token" type:"string"`
}

// String returns the string representation
func (s ListAccountRolesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAccountRolesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAccountRolesInput"}

	if s.AccessToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessToken"))
	}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAccountRolesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AccessToken != nil {
		v := *s.AccessToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-sso_bearer_token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "account_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max_result", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next_token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListAccountRolesOutput struct {
	_ struct{} `type:"structure"`

	// The page token client that is used to retrieve the list of accounts.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A paginated response with the list of roles and the next token if more results
	// are available.
	RoleList []RoleInfo `locationName:"roleList" type:"list"`
}

// String returns the string representation
func (s ListAccountRolesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAccountRolesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoleList != nil {
		v := s.RoleList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "roleList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListAccountRoles = "ListAccountRoles"

// ListAccountRolesRequest returns a request value for making API operation for
// AWS Single Sign-On.
//
// Lists all roles that are assigned to the user for a given AWS account.
//
//    // Example sending a request using ListAccountRolesRequest.
//    req := client.ListAccountRolesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/ListAccountRoles
func (c *Client) ListAccountRolesRequest(input *ListAccountRolesInput) ListAccountRolesRequest {
	op := &aws.Operation{
		Name:       opListAccountRoles,
		HTTPMethod: "GET",
		HTTPPath:   "/assignment/roles",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAccountRolesInput{}
	}

	req := c.newRequest(op, input, &ListAccountRolesOutput{})
	req.Config.Credentials = aws.AnonymousCredentials

	return ListAccountRolesRequest{Request: req, Input: input, Copy: c.ListAccountRolesRequest}
}

// ListAccountRolesRequest is the request type for the
// ListAccountRoles API operation.
type ListAccountRolesRequest struct {
	*aws.Request
	Input *ListAccountRolesInput
	Copy  func(*ListAccountRolesInput) ListAccountRolesRequest
}

// Send marshals and sends the ListAccountRoles API request.
func (r ListAccountRolesRequest) Send(ctx context.Context) (*ListAccountRolesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAccountRolesResponse{
		ListAccountRolesOutput: r.Request.Data.(*ListAccountRolesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAccountRolesRequestPaginator returns a paginator for ListAccountRoles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAccountRolesRequest(input)
//   p := sso.NewListAccountRolesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAccountRolesPaginator(req ListAccountRolesRequest) ListAccountRolesPaginator {
	return ListAccountRolesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAccountRolesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListAccountRolesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAccountRolesPaginator struct {
	aws.Pager
}

func (p *ListAccountRolesPaginator) CurrentPage() *ListAccountRolesOutput {
	return p.Pager.CurrentPage().(*ListAccountRolesOutput)
}

// ListAccountRolesResponse is the response type for the
// ListAccountRoles API operation.
type ListAccountRolesResponse struct {
	*ListAccountRolesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAccountRoles request.
func (r *ListAccountRolesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
