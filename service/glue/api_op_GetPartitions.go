// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetPartitionsRequest
type GetPartitionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the partitions in question reside. If none
	// is provided, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database where the partitions reside.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// An expression that filters the partitions to be returned.
	//
	// The expression uses SQL syntax similar to the SQL WHERE filter clause. The
	// SQL statement parser JSQLParser (http://jsqlparser.sourceforge.net/home.php)
	// parses the expression.
	//
	// Operators: The following are the operators that you can use in the Expression
	// API call:
	//
	// =
	//
	// Checks whether the values of the two operands are equal; if yes, then the
	// condition becomes true.
	//
	// Example: Assume 'variable a' holds 10 and 'variable b' holds 20.
	//
	// (a = b) is not true.
	//
	// < >
	//
	// Checks whether the values of two operands are equal; if the values are not
	// equal, then the condition becomes true.
	//
	// Example: (a < > b) is true.
	//
	// >
	//
	// Checks whether the value of the left operand is greater than the value of
	// the right operand; if yes, then the condition becomes true.
	//
	// Example: (a > b) is not true.
	//
	// <
	//
	// Checks whether the value of the left operand is less than the value of the
	// right operand; if yes, then the condition becomes true.
	//
	// Example: (a < b) is true.
	//
	// >=
	//
	// Checks whether the value of the left operand is greater than or equal to
	// the value of the right operand; if yes, then the condition becomes true.
	//
	// Example: (a >= b) is not true.
	//
	// <=
	//
	// Checks whether the value of the left operand is less than or equal to the
	// value of the right operand; if yes, then the condition becomes true.
	//
	// Example: (a <= b) is true.
	//
	// AND, OR, IN, BETWEEN, LIKE, NOT, IS NULL
	//
	// Logical operators.
	//
	// Supported Partition Key Types: The following are the supported partition
	// keys.
	//
	//    * string
	//
	//    * date
	//
	//    * timestamp
	//
	//    * int
	//
	//    * bigint
	//
	//    * long
	//
	//    * tinyint
	//
	//    * smallint
	//
	//    * decimal
	//
	// If an invalid type is encountered, an exception is thrown.
	//
	// The following list shows the valid operators on each type. When you define
	// a crawler, the partitionKey type is created as a STRING, to be compatible
	// with the catalog partitions.
	//
	// Sample API Call:
	Expression *string `type:"string"`

	// The maximum number of partitions to return in a single response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is not the first call to retrieve these partitions.
	NextToken *string `type:"string"`

	// The segment of the table's partitions to scan in this request.
	Segment *Segment `type:"structure"`

	// The name of the partitions' table.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPartitionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPartitionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetPartitionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
			invalidParams.AddNested("Segment", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetPartitionsResponse
type GetPartitionsOutput struct {
	_ struct{} `type:"structure"`

	// A continuation token, if the returned list of partitions does not include
	// the last one.
	NextToken *string `type:"string"`

	// A list of requested partitions.
	Partitions []Partition `type:"list"`
}

// String returns the string representation
func (s GetPartitionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetPartitions = "GetPartitions"

// GetPartitionsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves information about the partitions in a table.
//
//    // Example sending a request using GetPartitionsRequest.
//    req := client.GetPartitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetPartitions
func (c *Client) GetPartitionsRequest(input *GetPartitionsInput) GetPartitionsRequest {
	op := &aws.Operation{
		Name:       opGetPartitions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetPartitionsInput{}
	}

	req := c.newRequest(op, input, &GetPartitionsOutput{})
	return GetPartitionsRequest{Request: req, Input: input, Copy: c.GetPartitionsRequest}
}

// GetPartitionsRequest is the request type for the
// GetPartitions API operation.
type GetPartitionsRequest struct {
	*aws.Request
	Input *GetPartitionsInput
	Copy  func(*GetPartitionsInput) GetPartitionsRequest
}

// Send marshals and sends the GetPartitions API request.
func (r GetPartitionsRequest) Send(ctx context.Context) (*GetPartitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPartitionsResponse{
		GetPartitionsOutput: r.Request.Data.(*GetPartitionsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetPartitionsRequestPaginator returns a paginator for GetPartitions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetPartitionsRequest(input)
//   p := glue.NewGetPartitionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetPartitionsPaginator(req GetPartitionsRequest) GetPartitionsPaginator {
	return GetPartitionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetPartitionsInput
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

// GetPartitionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetPartitionsPaginator struct {
	aws.Pager
}

func (p *GetPartitionsPaginator) CurrentPage() *GetPartitionsOutput {
	return p.Pager.CurrentPage().(*GetPartitionsOutput)
}

// GetPartitionsResponse is the response type for the
// GetPartitions API operation.
type GetPartitionsResponse struct {
	*GetPartitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPartitions request.
func (r *GetPartitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}