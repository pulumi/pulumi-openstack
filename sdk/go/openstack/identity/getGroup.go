// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an OpenStack group.
// 
// Note: This usually requires admin privileges.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_group_v3.html.markdown.
func LookupGroup(ctx *pulumi.Context, args *GetGroupArgs) (*GetGroupResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["domainId"] = args.DomainId
		inputs["name"] = args.Name
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("openstack:identity/getGroup:getGroup", inputs)
	if err != nil {
		return nil, err
	}
	return &GetGroupResult{
		DomainId: outputs["domainId"],
		Name: outputs["name"],
		Region: outputs["region"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getGroup.
type GetGroupArgs struct {
	// The domain the group belongs to.
	DomainId interface{}
	// The name of the group.
	Name interface{}
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region interface{}
}

// A collection of values returned by getGroup.
type GetGroupResult struct {
	// See Argument Reference above.
	DomainId interface{}
	// See Argument Reference above.
	Name interface{}
	// See Argument Reference above.
	Region interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
