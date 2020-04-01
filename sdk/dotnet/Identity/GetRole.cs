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
        /// Use this data source to get the ID of an OpenStack role.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_role_v3.html.markdown.
        /// </summary>
        [Obsolete("Use GetRole.InvokeAsync() instead")]
        public static Task<GetRoleResult> GetRole(GetRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleResult>("openstack:identity/getRole:getRole", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetRole
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack role.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_role_v3.html.markdown.
        /// </summary>
        public static Task<GetRoleResult> InvokeAsync(GetRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoleResult>("openstack:identity/getRole:getRole", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRoleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain the role belongs to.
        /// </summary>
        [Input("domainId")]
        public string? DomainId { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetRoleArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRoleResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string DomainId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRoleResult(
            string domainId,
            string name,
            string region,
            string id)
        {
            DomainId = domainId;
            Name = name;
            Region = region;
            Id = id;
        }
    }
}
