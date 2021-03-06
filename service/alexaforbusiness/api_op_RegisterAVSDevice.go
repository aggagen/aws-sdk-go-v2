// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterAVSDeviceInput struct {
	_ struct{} `type:"structure"`

	// The device type ID for your AVS device generated by Amazon when the OEM creates
	// a new product on Amazon's Developer Console.
	//
	// AmazonId is a required field
	AmazonId *string `type:"string" required:"true"`

	// The client ID of the OEM used for code-based linking authorization on an
	// AVS device.
	//
	// ClientId is a required field
	ClientId *string `type:"string" required:"true"`

	// The key generated by the OEM that uniquely identifies a specified instance
	// of your AVS device.
	//
	// DeviceSerialNumber is a required field
	DeviceSerialNumber *string `type:"string" required:"true"`

	// The product ID used to identify your AVS device during authorization.
	//
	// ProductId is a required field
	ProductId *string `type:"string" required:"true"`

	// The code that is obtained after your AVS device has made a POST request to
	// LWA as a part of the Device Authorization Request component of the OAuth
	// code-based linking specification.
	//
	// UserCode is a required field
	UserCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterAVSDeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterAVSDeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterAVSDeviceInput"}

	if s.AmazonId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AmazonId"))
	}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}

	if s.DeviceSerialNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceSerialNumber"))
	}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}

	if s.UserCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserCode"))
	}
	if s.UserCode != nil && len(*s.UserCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterAVSDeviceOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the device.
	DeviceArn *string `type:"string"`
}

// String returns the string representation
func (s RegisterAVSDeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opRegisterAVSDevice = "RegisterAVSDevice"

// RegisterAVSDeviceRequest returns a request value for making API operation for
// Alexa For Business.
//
// Registers an Alexa-enabled device built by an Original Equipment Manufacturer
// (OEM) using Alexa Voice Service (AVS).
//
//    // Example sending a request using RegisterAVSDeviceRequest.
//    req := client.RegisterAVSDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/RegisterAVSDevice
func (c *Client) RegisterAVSDeviceRequest(input *RegisterAVSDeviceInput) RegisterAVSDeviceRequest {
	op := &aws.Operation{
		Name:       opRegisterAVSDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterAVSDeviceInput{}
	}

	req := c.newRequest(op, input, &RegisterAVSDeviceOutput{})

	return RegisterAVSDeviceRequest{Request: req, Input: input, Copy: c.RegisterAVSDeviceRequest}
}

// RegisterAVSDeviceRequest is the request type for the
// RegisterAVSDevice API operation.
type RegisterAVSDeviceRequest struct {
	*aws.Request
	Input *RegisterAVSDeviceInput
	Copy  func(*RegisterAVSDeviceInput) RegisterAVSDeviceRequest
}

// Send marshals and sends the RegisterAVSDevice API request.
func (r RegisterAVSDeviceRequest) Send(ctx context.Context) (*RegisterAVSDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterAVSDeviceResponse{
		RegisterAVSDeviceOutput: r.Request.Data.(*RegisterAVSDeviceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterAVSDeviceResponse is the response type for the
// RegisterAVSDevice API operation.
type RegisterAVSDeviceResponse struct {
	*RegisterAVSDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterAVSDevice request.
func (r *RegisterAVSDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
