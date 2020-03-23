// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keymanager

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an available Barbican container.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/keymanager_container_v1.html.markdown.
func LookupContainer(ctx *pulumi.Context, args *GetContainerArgs) (*GetContainerResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("openstack:keymanager/getContainer:getContainer", inputs)
	if err != nil {
		return nil, err
	}
	return &GetContainerResult{
		Acls: outputs["acls"],
		Consumers: outputs["consumers"],
		ContainerRef: outputs["containerRef"],
		CreatedAt: outputs["createdAt"],
		CreatorId: outputs["creatorId"],
		Name: outputs["name"],
		Region: outputs["region"],
		SecretRefs: outputs["secretRefs"],
		Status: outputs["status"],
		Type: outputs["type"],
		UpdatedAt: outputs["updatedAt"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getContainer.
type GetContainerArgs struct {
	// The Container name.
	Name interface{}
	// The region in which to obtain the V1 KeyManager client.
	// A KeyManager client is needed to fetch a container. If omitted, the `region`
	// argument of the provider is used.
	Region interface{}
}

// A collection of values returned by getContainer.
type GetContainerResult struct {
	// The list of ACLs assigned to a container. The `read` structure is
	// described below.
	Acls interface{}
	// The list of the container consumers. The structure is described
	// below.
	Consumers interface{}
	// The container reference / where to find the container.
	ContainerRef interface{}
	// The date the container ACL was created.
	CreatedAt interface{}
	// The creator of the container.
	CreatorId interface{}
	// The name of the consumer.
	Name interface{}
	// See Argument Reference above.
	Region interface{}
	// A set of dictionaries containing references to secrets. The
	// structure is described below.
	SecretRefs interface{}
	// The status of the container.
	Status interface{}
	// The container type.
	Type interface{}
	// The date the container ACL was last updated.
	UpdatedAt interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
