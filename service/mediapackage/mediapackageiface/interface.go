// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mediapackageiface provides an interface to enable mocking the AWS Elemental MediaPackage service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mediapackageiface

import (
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
)

// ClientAPI provides an interface to enable mocking the
// mediapackage.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // MediaPackage.
//    func myFunc(svc mediapackageiface.ClientAPI) bool {
//        // Make svc.CreateChannel request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := mediapackage.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        mediapackageiface.ClientPI
//    }
//    func (m *mockClientClient) CreateChannel(input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateChannelRequest(*mediapackage.CreateChannelInput) mediapackage.CreateChannelRequest

	CreateHarvestJobRequest(*mediapackage.CreateHarvestJobInput) mediapackage.CreateHarvestJobRequest

	CreateOriginEndpointRequest(*mediapackage.CreateOriginEndpointInput) mediapackage.CreateOriginEndpointRequest

	DeleteChannelRequest(*mediapackage.DeleteChannelInput) mediapackage.DeleteChannelRequest

	DeleteOriginEndpointRequest(*mediapackage.DeleteOriginEndpointInput) mediapackage.DeleteOriginEndpointRequest

	DescribeChannelRequest(*mediapackage.DescribeChannelInput) mediapackage.DescribeChannelRequest

	DescribeHarvestJobRequest(*mediapackage.DescribeHarvestJobInput) mediapackage.DescribeHarvestJobRequest

	DescribeOriginEndpointRequest(*mediapackage.DescribeOriginEndpointInput) mediapackage.DescribeOriginEndpointRequest

	ListChannelsRequest(*mediapackage.ListChannelsInput) mediapackage.ListChannelsRequest

	ListHarvestJobsRequest(*mediapackage.ListHarvestJobsInput) mediapackage.ListHarvestJobsRequest

	ListOriginEndpointsRequest(*mediapackage.ListOriginEndpointsInput) mediapackage.ListOriginEndpointsRequest

	ListTagsForResourceRequest(*mediapackage.ListTagsForResourceInput) mediapackage.ListTagsForResourceRequest

	RotateChannelCredentialsRequest(*mediapackage.RotateChannelCredentialsInput) mediapackage.RotateChannelCredentialsRequest

	RotateIngestEndpointCredentialsRequest(*mediapackage.RotateIngestEndpointCredentialsInput) mediapackage.RotateIngestEndpointCredentialsRequest

	TagResourceRequest(*mediapackage.TagResourceInput) mediapackage.TagResourceRequest

	UntagResourceRequest(*mediapackage.UntagResourceInput) mediapackage.UntagResourceRequest

	UpdateChannelRequest(*mediapackage.UpdateChannelInput) mediapackage.UpdateChannelRequest

	UpdateOriginEndpointRequest(*mediapackage.UpdateOriginEndpointInput) mediapackage.UpdateOriginEndpointRequest
}

var _ ClientAPI = (*mediapackage.Client)(nil)
