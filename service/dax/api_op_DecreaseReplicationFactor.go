// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dax

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DecreaseReplicationFactorInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone(s) from which to remove nodes.
	AvailabilityZones []string `type:"list"`

	// The name of the DAX cluster from which you want to remove nodes.
	//
	// ClusterName is a required field
	ClusterName *string `type:"string" required:"true"`

	// The new number of nodes for the DAX cluster.
	//
	// NewReplicationFactor is a required field
	NewReplicationFactor *int64 `type:"integer" required:"true"`

	// The unique identifiers of the nodes to be removed from the cluster.
	NodeIdsToRemove []string `type:"list"`
}

// String returns the string representation
func (s DecreaseReplicationFactorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecreaseReplicationFactorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DecreaseReplicationFactorInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NewReplicationFactor == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewReplicationFactor"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DecreaseReplicationFactorOutput struct {
	_ struct{} `type:"structure"`

	// A description of the DAX cluster, after you have decreased its replication
	// factor.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s DecreaseReplicationFactorOutput) String() string {
	return awsutil.Prettify(s)
}

const opDecreaseReplicationFactor = "DecreaseReplicationFactor"

// DecreaseReplicationFactorRequest returns a request value for making API operation for
// Amazon DynamoDB Accelerator (DAX).
//
// Removes one or more nodes from a DAX cluster.
//
// You cannot use DecreaseReplicationFactor to remove the last node in a DAX
// cluster. If you need to do this, use DeleteCluster instead.
//
//    // Example sending a request using DecreaseReplicationFactorRequest.
//    req := client.DecreaseReplicationFactorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dax-2017-04-19/DecreaseReplicationFactor
func (c *Client) DecreaseReplicationFactorRequest(input *DecreaseReplicationFactorInput) DecreaseReplicationFactorRequest {
	op := &aws.Operation{
		Name:       opDecreaseReplicationFactor,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecreaseReplicationFactorInput{}
	}

	req := c.newRequest(op, input, &DecreaseReplicationFactorOutput{})

	return DecreaseReplicationFactorRequest{Request: req, Input: input, Copy: c.DecreaseReplicationFactorRequest}
}

// DecreaseReplicationFactorRequest is the request type for the
// DecreaseReplicationFactor API operation.
type DecreaseReplicationFactorRequest struct {
	*aws.Request
	Input *DecreaseReplicationFactorInput
	Copy  func(*DecreaseReplicationFactorInput) DecreaseReplicationFactorRequest
}

// Send marshals and sends the DecreaseReplicationFactor API request.
func (r DecreaseReplicationFactorRequest) Send(ctx context.Context) (*DecreaseReplicationFactorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DecreaseReplicationFactorResponse{
		DecreaseReplicationFactorOutput: r.Request.Data.(*DecreaseReplicationFactorOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DecreaseReplicationFactorResponse is the response type for the
// DecreaseReplicationFactor API operation.
type DecreaseReplicationFactorResponse struct {
	*DecreaseReplicationFactorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DecreaseReplicationFactor request.
func (r *DecreaseReplicationFactorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
