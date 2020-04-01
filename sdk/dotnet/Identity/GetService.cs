// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack service.
        /// 
        /// &gt; **Note:** This usually requires admin privileges.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_service_v3.html.markdown.
        /// </summary>
        [Obsolete("Use GetService.InvokeAsync() instead")]
        public static Task<GetServiceResult> GetService(GetServiceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("openstack:identity/getService:getService", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetService
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack service.
        /// 
        /// &gt; **Note:** This usually requires admin privileges.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_service_v3.html.markdown.
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("openstack:identity/getService:getService", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service status.
        /// </summary>
        [Input("enabled")]
        public bool? Enabled { get; set; }

        /// <summary>
        /// The service name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The service type.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetServiceArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetServiceResult
    {
        /// <summary>
        /// The service description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetServiceResult(
            string description,
            bool? enabled,
            string? name,
            string region,
            string? type,
            string id)
        {
            Description = description;
            Enabled = enabled;
            Name = name;
            Region = region;
            Type = type;
            Id = id;
        }
    }
}
