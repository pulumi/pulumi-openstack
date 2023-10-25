// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetAddressScope
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack address-scope.
        /// </summary>
        public static Task<GetAddressScopeResult> InvokeAsync(GetAddressScopeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAddressScopeResult>("openstack:networking/getAddressScope:getAddressScope", args ?? new GetAddressScopeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack address-scope.
        /// </summary>
        public static Output<GetAddressScopeResult> Invoke(GetAddressScopeInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddressScopeResult>("openstack:networking/getAddressScope:getAddressScope", args ?? new GetAddressScopeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddressScopeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// IP version.
        /// </summary>
        [Input("ipVersion")]
        public int? IpVersion { get; set; }

        /// <summary>
        /// Name of the address-scope.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The owner of the address-scope.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve address-scopes. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Indicates whether this address-scope is shared across
        /// all projects.
        /// </summary>
        [Input("shared")]
        public bool? Shared { get; set; }

        public GetAddressScopeArgs()
        {
        }
        public static new GetAddressScopeArgs Empty => new GetAddressScopeArgs();
    }

    public sealed class GetAddressScopeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// IP version.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// Name of the address-scope.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the address-scope.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve address-scopes. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates whether this address-scope is shared across
        /// all projects.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        public GetAddressScopeInvokeArgs()
        {
        }
        public static new GetAddressScopeInvokeArgs Empty => new GetAddressScopeInvokeArgs();
    }


    [OutputType]
    public sealed class GetAddressScopeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int? IpVersion;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? ProjectId;
        public readonly string? Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? Shared;

        [OutputConstructor]
        private GetAddressScopeResult(
            string id,

            int? ipVersion,

            string? name,

            string? projectId,

            string? region,

            bool? shared)
        {
            Id = id;
            IpVersion = ipVersion;
            Name = name;
            ProjectId = projectId;
            Region = region;
            Shared = shared;
        }
    }
}
