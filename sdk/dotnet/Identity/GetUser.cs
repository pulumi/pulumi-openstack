// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity
{
    public static class GetUser
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack user.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var user1 = OpenStack.Identity.GetUser.Invoke(new()
        ///     {
        ///         Name = "user_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("openstack:identity/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an OpenStack user.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var user1 = OpenStack.Identity.GetUser.Invoke(new()
        ///     {
        ///         Name = "user_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("openstack:identity/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// The region the user is located in.
        /// </summary>
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
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain this user belongs to.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Whether the user is enabled or disabled. Valid
        /// values are `true` and `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The identity provider ID of the user.
        /// </summary>
        [Input("idpId")]
        public Input<string>? IdpId { get; set; }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
        /// </summary>
        [Input("passwordExpiresAt")]
        public Input<string>? PasswordExpiresAt { get; set; }

        /// <summary>
        /// The protocol ID of the user.
        /// </summary>
        [Input("protocolId")]
        public Input<string>? ProtocolId { get; set; }

        /// <summary>
        /// The region the user is located in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The unique ID of the user.
        /// </summary>
        [Input("uniqueId")]
        public Input<string>? UniqueId { get; set; }

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
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
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
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

        [OutputConstructor]
        private GetUserResult(
            string defaultProjectId,

            string description,

            string domainId,

            bool? enabled,

            string id,

            string? idpId,

            string? name,

            string? passwordExpiresAt,

            string? protocolId,

            string region,

            string? uniqueId)
        {
            DefaultProjectId = defaultProjectId;
            Description = description;
            DomainId = domainId;
            Enabled = enabled;
            Id = id;
            IdpId = idpId;
            Name = name;
            PasswordExpiresAt = passwordExpiresAt;
            ProtocolId = protocolId;
            Region = region;
            UniqueId = uniqueId;
        }
    }
}
