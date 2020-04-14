// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the ID of an OpenStack group.
//
// Note: This usually requires admin privileges.
func GetGroup(ctx *pulumi.Context, args *GetGroupArgs, opts ...pulumi.InvokeOption) (*GetGroupResult, error) {
	var rv GetGroupResult
	err := ctx.Invoke("openstack:identity/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type GetGroupArgs struct {
	// The domain the group belongs to.
	DomainId *string `pulumi:"domainId"`
	// The name of the group.
	Name string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getGroup.
type GetGroupResult struct {
	// A description of the group.
	Description string `pulumi:"description"`
	// See Argument Reference above.
	DomainId string `pulumi:"domainId"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
}
