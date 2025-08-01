// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of a running server
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := compute.GetInstanceV2(ctx, &compute.GetInstanceV2Args{
//				Id: "2ba26dc6-a12d-4889-8f25-794ea5bf4453",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstanceV2(ctx *pulumi.Context, args *GetInstanceV2Args, opts ...pulumi.InvokeOption) (*GetInstanceV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstanceV2Result
	err := ctx.Invoke("openstack:compute/getInstanceV2:getInstanceV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceV2.
type GetInstanceV2Args struct {
	// The UUID of the instance
	Id string `pulumi:"id"`
	// An array of maps, detailed below.
	Networks []GetInstanceV2Network `pulumi:"networks"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The user data added when the server was created.
	UserData *string `pulumi:"userData"`
}

// A collection of values returned by getInstanceV2.
type GetInstanceV2Result struct {
	// The first IPv4 address assigned to this server.
	AccessIpV4 string `pulumi:"accessIpV4"`
	// The first IPv6 address assigned to this server.
	AccessIpV6 string `pulumi:"accessIpV6"`
	// The availability zone of this server.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The creation time of the instance.
	Created string `pulumi:"created"`
	// The flavor ID used to create the server.
	FlavorId string `pulumi:"flavorId"`
	// The flavor name used to create the server.
	FlavorName string `pulumi:"flavorName"`
	Id         string `pulumi:"id"`
	// The image ID used to create the server.
	ImageId string `pulumi:"imageId"`
	// The image name used to create the server.
	ImageName string `pulumi:"imageName"`
	// The name of the key pair assigned to this server.
	KeyPair string `pulumi:"keyPair"`
	// A set of key/value pairs made available to the server.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the network
	Name string `pulumi:"name"`
	// An array of maps, detailed below.
	Networks   []GetInstanceV2Network `pulumi:"networks"`
	PowerState string                 `pulumi:"powerState"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// An array of security group names associated with this server.
	SecurityGroups []string `pulumi:"securityGroups"`
	// A set of string tags assigned to this server.
	Tags []string `pulumi:"tags"`
	// The time when the instance was last updated.
	Updated string `pulumi:"updated"`
	// The user data added when the server was created.
	UserData string `pulumi:"userData"`
}

func GetInstanceV2Output(ctx *pulumi.Context, args GetInstanceV2OutputArgs, opts ...pulumi.InvokeOption) GetInstanceV2ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetInstanceV2ResultOutput, error) {
			args := v.(GetInstanceV2Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:compute/getInstanceV2:getInstanceV2", args, GetInstanceV2ResultOutput{}, options).(GetInstanceV2ResultOutput), nil
		}).(GetInstanceV2ResultOutput)
}

// A collection of arguments for invoking getInstanceV2.
type GetInstanceV2OutputArgs struct {
	// The UUID of the instance
	Id pulumi.StringInput `pulumi:"id"`
	// An array of maps, detailed below.
	Networks GetInstanceV2NetworkArrayInput `pulumi:"networks"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The user data added when the server was created.
	UserData pulumi.StringPtrInput `pulumi:"userData"`
}

func (GetInstanceV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceV2Args)(nil)).Elem()
}

// A collection of values returned by getInstanceV2.
type GetInstanceV2ResultOutput struct{ *pulumi.OutputState }

func (GetInstanceV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceV2Result)(nil)).Elem()
}

func (o GetInstanceV2ResultOutput) ToGetInstanceV2ResultOutput() GetInstanceV2ResultOutput {
	return o
}

func (o GetInstanceV2ResultOutput) ToGetInstanceV2ResultOutputWithContext(ctx context.Context) GetInstanceV2ResultOutput {
	return o
}

// The first IPv4 address assigned to this server.
func (o GetInstanceV2ResultOutput) AccessIpV4() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.AccessIpV4 }).(pulumi.StringOutput)
}

// The first IPv6 address assigned to this server.
func (o GetInstanceV2ResultOutput) AccessIpV6() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.AccessIpV6 }).(pulumi.StringOutput)
}

// The availability zone of this server.
func (o GetInstanceV2ResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The creation time of the instance.
func (o GetInstanceV2ResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.Created }).(pulumi.StringOutput)
}

// The flavor ID used to create the server.
func (o GetInstanceV2ResultOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.FlavorId }).(pulumi.StringOutput)
}

// The flavor name used to create the server.
func (o GetInstanceV2ResultOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.FlavorName }).(pulumi.StringOutput)
}

func (o GetInstanceV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.Id }).(pulumi.StringOutput)
}

// The image ID used to create the server.
func (o GetInstanceV2ResultOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.ImageId }).(pulumi.StringOutput)
}

// The image name used to create the server.
func (o GetInstanceV2ResultOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.ImageName }).(pulumi.StringOutput)
}

// The name of the key pair assigned to this server.
func (o GetInstanceV2ResultOutput) KeyPair() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.KeyPair }).(pulumi.StringOutput)
}

// A set of key/value pairs made available to the server.
func (o GetInstanceV2ResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetInstanceV2Result) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the network
func (o GetInstanceV2ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.Name }).(pulumi.StringOutput)
}

// An array of maps, detailed below.
func (o GetInstanceV2ResultOutput) Networks() GetInstanceV2NetworkArrayOutput {
	return o.ApplyT(func(v GetInstanceV2Result) []GetInstanceV2Network { return v.Networks }).(GetInstanceV2NetworkArrayOutput)
}

func (o GetInstanceV2ResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.PowerState }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetInstanceV2ResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.Region }).(pulumi.StringOutput)
}

// An array of security group names associated with this server.
func (o GetInstanceV2ResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceV2Result) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// A set of string tags assigned to this server.
func (o GetInstanceV2ResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstanceV2Result) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The time when the instance was last updated.
func (o GetInstanceV2ResultOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.Updated }).(pulumi.StringOutput)
}

// The user data added when the server was created.
func (o GetInstanceV2ResultOutput) UserData() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceV2Result) string { return v.UserData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceV2ResultOutput{})
}
