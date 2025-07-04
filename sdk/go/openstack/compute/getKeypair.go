// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID and public key of an OpenStack keypair.
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
//			_, err := compute.LookupKeypair(ctx, &compute.LookupKeypairArgs{
//				Name: "sand",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupKeypair(ctx *pulumi.Context, args *LookupKeypairArgs, opts ...pulumi.InvokeOption) (*LookupKeypairResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKeypairResult
	err := ctx.Invoke("openstack:compute/getKeypair:getKeypair", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeypair.
type LookupKeypairArgs struct {
	// The unique name of the keypair.
	Name string `pulumi:"name"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The user id of the owner of the key pair.
	// This parameter can be specified only if the provider is configured to use
	// the credentials of an OpenStack administrator.
	UserId *string `pulumi:"userId"`
}

// A collection of values returned by getKeypair.
type LookupKeypairResult struct {
	// The fingerprint of the OpenSSH key.
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// The OpenSSH-formatted public key of the keypair.
	PublicKey string `pulumi:"publicKey"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	UserId string `pulumi:"userId"`
}

func LookupKeypairOutput(ctx *pulumi.Context, args LookupKeypairOutputArgs, opts ...pulumi.InvokeOption) LookupKeypairResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKeypairResultOutput, error) {
			args := v.(LookupKeypairArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("openstack:compute/getKeypair:getKeypair", args, LookupKeypairResultOutput{}, options).(LookupKeypairResultOutput), nil
		}).(LookupKeypairResultOutput)
}

// A collection of arguments for invoking getKeypair.
type LookupKeypairOutputArgs struct {
	// The unique name of the keypair.
	Name pulumi.StringInput `pulumi:"name"`
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The user id of the owner of the key pair.
	// This parameter can be specified only if the provider is configured to use
	// the credentials of an OpenStack administrator.
	UserId pulumi.StringPtrInput `pulumi:"userId"`
}

func (LookupKeypairOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeypairArgs)(nil)).Elem()
}

// A collection of values returned by getKeypair.
type LookupKeypairResultOutput struct{ *pulumi.OutputState }

func (LookupKeypairResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeypairResult)(nil)).Elem()
}

func (o LookupKeypairResultOutput) ToLookupKeypairResultOutput() LookupKeypairResultOutput {
	return o
}

func (o LookupKeypairResultOutput) ToLookupKeypairResultOutputWithContext(ctx context.Context) LookupKeypairResultOutput {
	return o
}

// The fingerprint of the OpenSSH key.
func (o LookupKeypairResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKeypairResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupKeypairResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.Name }).(pulumi.StringOutput)
}

// The OpenSSH-formatted public key of the keypair.
func (o LookupKeypairResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupKeypairResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.Region }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupKeypairResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeypairResultOutput{})
}
