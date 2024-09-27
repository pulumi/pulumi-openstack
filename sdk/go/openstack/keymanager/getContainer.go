// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available Barbican container.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/keymanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := keymanager.GetContainer(ctx, &keymanager.GetContainerArgs{
//				Name: pulumi.StringRef("my_container"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetContainer(ctx *pulumi.Context, args *GetContainerArgs, opts ...pulumi.InvokeOption) (*GetContainerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetContainerResult
	err := ctx.Invoke("openstack:keymanager/getContainer:getContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainer.
type GetContainerArgs struct {
	// The Container name.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to fetch a container. If omitted, the `region`
	// argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getContainer.
type GetContainerResult struct {
	// The list of ACLs assigned to a container. The `read` structure is
	// described below.
	Acls []GetContainerAcl `pulumi:"acls"`
	// The list of the container consumers. The structure is described
	// below.
	Consumers []GetContainerConsumer `pulumi:"consumers"`
	// The container reference / where to find the container.
	ContainerRef string `pulumi:"containerRef"`
	// The date the container ACL was created.
	CreatedAt string `pulumi:"createdAt"`
	// The creator of the container.
	CreatorId string `pulumi:"creatorId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the consumer.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	Region *string `pulumi:"region"`
	// A set of dictionaries containing references to secrets. The
	// structure is described below.
	SecretRefs []GetContainerSecretRef `pulumi:"secretRefs"`
	// The status of the container.
	Status string `pulumi:"status"`
	// The container type.
	Type string `pulumi:"type"`
	// The date the container ACL was last updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetContainerOutput(ctx *pulumi.Context, args GetContainerOutputArgs, opts ...pulumi.InvokeOption) GetContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContainerResultOutput, error) {
			args := v.(GetContainerArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetContainerResult
			secret, err := ctx.InvokePackageRaw("openstack:keymanager/getContainer:getContainer", args, &rv, "", opts...)
			if err != nil {
				return GetContainerResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetContainerResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetContainerResultOutput), nil
			}
			return output, nil
		}).(GetContainerResultOutput)
}

// A collection of arguments for invoking getContainer.
type GetContainerOutputArgs struct {
	// The Container name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to fetch a container. If omitted, the `region`
	// argument of the provider is used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerArgs)(nil)).Elem()
}

// A collection of values returned by getContainer.
type GetContainerResultOutput struct{ *pulumi.OutputState }

func (GetContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerResult)(nil)).Elem()
}

func (o GetContainerResultOutput) ToGetContainerResultOutput() GetContainerResultOutput {
	return o
}

func (o GetContainerResultOutput) ToGetContainerResultOutputWithContext(ctx context.Context) GetContainerResultOutput {
	return o
}

// The list of ACLs assigned to a container. The `read` structure is
// described below.
func (o GetContainerResultOutput) Acls() GetContainerAclArrayOutput {
	return o.ApplyT(func(v GetContainerResult) []GetContainerAcl { return v.Acls }).(GetContainerAclArrayOutput)
}

// The list of the container consumers. The structure is described
// below.
func (o GetContainerResultOutput) Consumers() GetContainerConsumerArrayOutput {
	return o.ApplyT(func(v GetContainerResult) []GetContainerConsumer { return v.Consumers }).(GetContainerConsumerArrayOutput)
}

// The container reference / where to find the container.
func (o GetContainerResultOutput) ContainerRef() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerResult) string { return v.ContainerRef }).(pulumi.StringOutput)
}

// The date the container ACL was created.
func (o GetContainerResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The creator of the container.
func (o GetContainerResultOutput) CreatorId() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerResult) string { return v.CreatorId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the consumer.
func (o GetContainerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o GetContainerResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// A set of dictionaries containing references to secrets. The
// structure is described below.
func (o GetContainerResultOutput) SecretRefs() GetContainerSecretRefArrayOutput {
	return o.ApplyT(func(v GetContainerResult) []GetContainerSecretRef { return v.SecretRefs }).(GetContainerSecretRefArrayOutput)
}

// The status of the container.
func (o GetContainerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerResult) string { return v.Status }).(pulumi.StringOutput)
}

// The container type.
func (o GetContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

// The date the container ACL was last updated.
func (o GetContainerResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerResultOutput{})
}
