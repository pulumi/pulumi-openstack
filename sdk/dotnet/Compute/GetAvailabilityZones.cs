// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get a list of availability zones from OpenStack
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/compute_availability_zones_v2.html.markdown.
        /// </summary>
        [Obsolete("Use GetAvailabilityZones.InvokeAsync() instead")]
        public static Task<GetAvailabilityZonesResult> GetAvailabilityZones(GetAvailabilityZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityZonesResult>("openstack:compute/getAvailabilityZones:getAvailabilityZones", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetAvailabilityZones
    {
        /// <summary>
        /// Use this data source to get a list of availability zones from OpenStack
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/compute_availability_zones_v2.html.markdown.
        /// </summary>
        public static Task<GetAvailabilityZonesResult> InvokeAsync(GetAvailabilityZonesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAvailabilityZonesResult>("openstack:compute/getAvailabilityZones:getAvailabilityZones", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAvailabilityZonesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The `region` to fetch availability zones from, defaults to the provider's `region`
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The `state` of the availability zones to match, default ("available").
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetAvailabilityZonesArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAvailabilityZonesResult
    {
        /// <summary>
        /// The names of the availability zones, ordered alphanumerically, that match the queried `state`
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string Region;
        public readonly string? State;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAvailabilityZonesResult(
            ImmutableArray<string> names,
            string region,
            string? state,
            string id)
        {
            Names = names;
            Region = region;
            State = state;
            Id = id;
        }
    }
}
