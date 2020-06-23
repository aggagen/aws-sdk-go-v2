// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of table names associated with the current account and
// endpoint. The output from ListTables is paginated, with each page returning a
// maximum of 100 table names.
func (c *Client) ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) {
	stack := middleware.NewStack("ListTables", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTables(options.Region), middleware.Before)
	addawsAwsjson10_serdeOpListTablesMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "ListTables",
			Err:           err,
		}
	}
	out := result.(*ListTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListTables operation.
type ListTablesInput struct {
	// The first table name that this operation will evaluate. Use the value that was
	// returned for LastEvaluatedTableName in a previous operation, so that you can
	// obtain the next page of results.
	ExclusiveStartTableName *string
	// A maximum number of table names to return. If this parameter is not specified,
	// the limit is 100.
	Limit *int32
}

// Represents the output of a ListTables operation.
type ListTablesOutput struct {
	// The name of the last table in the current page of results. Use this value as the
	// ExclusiveStartTableName in a new request to obtain the next page of results,
	// until all the table names are returned. If you do not receive a
	// LastEvaluatedTableName value in the response, this means that there are no more
	// table names to be retrieved.
	LastEvaluatedTableName *string
	// The names of the tables associated with the current account at the current
	// endpoint. The maximum size of this array is 100. If LastEvaluatedTableName also
	// appears in the output, you can use this value as the ExclusiveStartTableName
	// parameter in a subsequent ListTables request and obtain the next page of
	// results.
	TableNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListTablesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListTables{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListTables{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "ListTables",
	}
}