// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Identity
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack user.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_user_v3.html.markdown.
        /// </summary>
        public static Task<GetUserResult> GetUser(GetUserArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("openstack:identity/getUser:getUser", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetUserArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain this user belongs to.
        /// </summary>
        [Input("domainId")]
        public string? DomainId { get; set; }

        /// <summary>
        /// Whether the user is enabled or disabled. Valid
        /// values are `true` and `false`.
        /// </summary>
        [Input("enabled")]
        public bool? Enabled { get; set; }

        /// <summary>
        /// The identity provider ID of the user.
        /// </summary>
        [Input("idpId")]
        public string? IdpId { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
        /// </summary>
        [Input("passwordExpiresAt")]
        public string? PasswordExpiresAt { get; set; }

        /// <summary>
        /// The protocol ID of the user.
        /// </summary>
        [Input("protocolId")]
        public string? ProtocolId { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The unique ID of the user.
        /// </summary>
        [Input("uniqueId")]
        public string? UniqueId { get; set; }

        public GetUserArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string DefaultProjectId;
        /// <summary>
        /// A description of the user.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string DomainId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? IdpId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? PasswordExpiresAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? ProtocolId;
        /// <summary>
        /// The region the user is located in.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? UniqueId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetUserResult(
            string defaultProjectId,
            string description,
            string domainId,
            bool? enabled,
            string? idpId,
            string? name,
            string? passwordExpiresAt,
            string? protocolId,
            string region,
            string? uniqueId,
            string id)
        {
            DefaultProjectId = defaultProjectId;
            Description = description;
            DomainId = domainId;
            Enabled = enabled;
            IdpId = idpId;
            Name = name;
            PasswordExpiresAt = passwordExpiresAt;
            ProtocolId = protocolId;
            Region = region;
            UniqueId = uniqueId;
            Id = id;
        }
    }
}
