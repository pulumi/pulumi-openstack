// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package identity

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID of an OpenStack endpoint.
// 
// > **Note:** This usually requires admin privileges.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_endpoint_v3.html.markdown.
func GetEndpoint(ctx *pulumi.Context, args *GetEndpointArgs, opts ...pulumi.InvokeOption) (*GetEndpointResult, error) {
	var rv GetEndpointResult
	err := ctx.Invoke("openstack:identity/getEndpoint:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpoint.
type GetEndpointArgs struct {
	// The region the endpoint is assigned to. The
	// `region` and `endpointRegion` can be different.
	EndpointRegion *string `pulumi:"endpointRegion"`
	// The endpoint interface. Valid values are `public`,
	// `internal`, and `admin`. Default value is `public`
	Interface *string `pulumi:"interface"`
	// The name of the endpoint.
	Name *string `pulumi:"name"`
	// The region in which to obtain the V3 Keystone client.
	// If omitted, the `region` argument of the provider is used.
	Region *string `pulumi:"region"`
	// The service id this endpoint belongs to.
	ServiceId *string `pulumi:"serviceId"`
	// The service name of the endpoint.
	ServiceName *string `pulumi:"serviceName"`
	// The service type of the endpoint.
	ServiceType *string `pulumi:"serviceType"`
}


// A collection of values returned by getEndpoint.
type GetEndpointResult struct {
	// See Argument Reference above.
	EndpointRegion *string `pulumi:"endpointRegion"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Interface *string `pulumi:"interface"`
	// See Argument Reference above.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	Region string `pulumi:"region"`
	// See Argument Reference above.
	ServiceId *string `pulumi:"serviceId"`
	// See Argument Reference above.
	ServiceName *string `pulumi:"serviceName"`
	// See Argument Reference above.
	ServiceType *string `pulumi:"serviceType"`
	// The endpoint URL.
	Url string `pulumi:"url"`
}

