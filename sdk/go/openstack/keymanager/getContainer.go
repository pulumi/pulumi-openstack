// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ID of an available Barbican container.
func GetContainer(ctx *pulumi.Context, args *GetContainerArgs, opts ...pulumi.InvokeOption) (*GetContainerResult, error) {
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
	// id is the provider-assigned unique ID for this managed resource.
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
