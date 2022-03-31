// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager
{
    public static class GetSecret
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(OpenStack.KeyManager.GetSecret.InvokeAsync(new OpenStack.KeyManager.GetSecretArgs
        ///         {
        ///             Mode = "cbc",
        ///             SecretType = "passphrase",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Date Filters
        /// 
        /// The values for the `expiration_filter`, `created_at_filter`, and
        /// `updated_at_filter` parameters are comma-separated lists of time stamps in
        /// RFC3339 format. The time stamps can be prefixed with any of these comparison
        /// operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
        /// (less-than), *lte:* (less-than-or-equal).
        /// 
        /// For example, to get a passphrase a Secret with CBC moda, that will expire in
        /// January of 2020:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var dateFilterExample = Output.Create(OpenStack.KeyManager.GetSecret.InvokeAsync(new OpenStack.KeyManager.GetSecretArgs
        ///         {
        ///             ExpirationFilter = "gt:2020-01-01T00:00:00Z",
        ///             Mode = "cbc",
        ///             SecretType = "passphrase",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public static Task<GetSecretResult> InvokeAsync(GetSecretArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretResult>("openstack:keymanager/getSecret:getSecret", args ?? new GetSecretArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(OpenStack.KeyManager.GetSecret.InvokeAsync(new OpenStack.KeyManager.GetSecretArgs
        ///         {
        ///             Mode = "cbc",
        ///             SecretType = "passphrase",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Date Filters
        /// 
        /// The values for the `expiration_filter`, `created_at_filter`, and
        /// `updated_at_filter` parameters are comma-separated lists of time stamps in
        /// RFC3339 format. The time stamps can be prefixed with any of these comparison
        /// operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
        /// (less-than), *lte:* (less-than-or-equal).
        /// 
        /// For example, to get a passphrase a Secret with CBC moda, that will expire in
        /// January of 2020:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var dateFilterExample = Output.Create(OpenStack.KeyManager.GetSecret.InvokeAsync(new OpenStack.KeyManager.GetSecretArgs
        ///         {
        ///             ExpirationFilter = "gt:2020-01-01T00:00:00Z",
        ///             Mode = "cbc",
        ///             SecretType = "passphrase",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public static Output<GetSecretResult> Invoke(GetSecretInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSecretResult>("openstack:keymanager/getSecret:getSecret", args ?? new GetSecretInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Select the Secret with an ACL that contains the user.
        /// Project scope is ignored. Defaults to `false`.
        /// </summary>
        [Input("aclOnly")]
        public bool? AclOnly { get; set; }

        /// <summary>
        /// The Secret algorithm.
        /// </summary>
        [Input("algorithm")]
        public string? Algorithm { get; set; }

        /// <summary>
        /// The Secret bit length.
        /// </summary>
        [Input("bitLength")]
        public int? BitLength { get; set; }

        /// <summary>
        /// Date filter to select the Secret with
        /// created matching the specified criteria. See Date Filters below for more
        /// detail.
        /// </summary>
        [Input("createdAtFilter")]
        public string? CreatedAtFilter { get; set; }

        /// <summary>
        /// Date filter to select the Secret with
        /// expiration matching the specified criteria. See Date Filters below for more
        /// detail.
        /// </summary>
        [Input("expirationFilter")]
        public string? ExpirationFilter { get; set; }

        /// <summary>
        /// The Secret mode.
        /// </summary>
        [Input("mode")]
        public string? Mode { get; set; }

        /// <summary>
        /// The Secret name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to fetch a secret. If omitted, the `region`
        /// argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The Secret type. For more information see
        /// [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
        /// </summary>
        [Input("secretType")]
        public string? SecretType { get; set; }

        /// <summary>
        /// Date filter to select the Secret with
        /// updated matching the specified criteria. See Date Filters below for more
        /// detail.
        /// </summary>
        [Input("updatedAtFilter")]
        public string? UpdatedAtFilter { get; set; }

        public GetSecretArgs()
        {
        }
    }

    public sealed class GetSecretInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Select the Secret with an ACL that contains the user.
        /// Project scope is ignored. Defaults to `false`.
        /// </summary>
        [Input("aclOnly")]
        public Input<bool>? AclOnly { get; set; }

        /// <summary>
        /// The Secret algorithm.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The Secret bit length.
        /// </summary>
        [Input("bitLength")]
        public Input<int>? BitLength { get; set; }

        /// <summary>
        /// Date filter to select the Secret with
        /// created matching the specified criteria. See Date Filters below for more
        /// detail.
        /// </summary>
        [Input("createdAtFilter")]
        public Input<string>? CreatedAtFilter { get; set; }

        /// <summary>
        /// Date filter to select the Secret with
        /// expiration matching the specified criteria. See Date Filters below for more
        /// detail.
        /// </summary>
        [Input("expirationFilter")]
        public Input<string>? ExpirationFilter { get; set; }

        /// <summary>
        /// The Secret mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The Secret name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to fetch a secret. If omitted, the `region`
        /// argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The Secret type. For more information see
        /// [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
        /// </summary>
        [Input("secretType")]
        public Input<string>? SecretType { get; set; }

        /// <summary>
        /// Date filter to select the Secret with
        /// updated matching the specified criteria. See Date Filters below for more
        /// detail.
        /// </summary>
        [Input("updatedAtFilter")]
        public Input<string>? UpdatedAtFilter { get; set; }

        public GetSecretInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecretResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? AclOnly;
        /// <summary>
        /// The list of ACLs assigned to a secret. The `read` structure is described below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecretAclResult> Acls;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Algorithm;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int? BitLength;
        /// <summary>
        /// The map of the content types, assigned on the secret.
        /// </summary>
        public readonly ImmutableDictionary<string, object> ContentTypes;
        /// <summary>
        /// The date the secret ACL was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? CreatedAtFilter;
        /// <summary>
        /// The creator of the secret.
        /// </summary>
        public readonly string CreatorId;
        /// <summary>
        /// The date the secret will expire.
        /// </summary>
        public readonly string Expiration;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? ExpirationFilter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The map of metadata, assigned on the secret, which has been
        /// explicitly and implicitly added.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The secret payload.
        /// </summary>
        public readonly string Payload;
        /// <summary>
        /// The Secret encoding.
        /// </summary>
        public readonly string PayloadContentEncoding;
        /// <summary>
        /// The Secret content type.
        /// </summary>
        public readonly string PayloadContentType;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The secret reference / where to find the secret.
        /// </summary>
        public readonly string SecretRef;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? SecretType;
        /// <summary>
        /// The status of the secret.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The date the secret ACL was last updated.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? UpdatedAtFilter;

        [OutputConstructor]
        private GetSecretResult(
            bool? aclOnly,

            ImmutableArray<Outputs.GetSecretAclResult> acls,

            string? algorithm,

            int? bitLength,

            ImmutableDictionary<string, object> contentTypes,

            string createdAt,

            string? createdAtFilter,

            string creatorId,

            string expiration,

            string? expirationFilter,

            string id,

            ImmutableDictionary<string, object> metadata,

            string? mode,

            string? name,

            string payload,

            string payloadContentEncoding,

            string payloadContentType,

            string? region,

            string secretRef,

            string? secretType,

            string status,

            string updatedAt,

            string? updatedAtFilter)
        {
            AclOnly = aclOnly;
            Acls = acls;
            Algorithm = algorithm;
            BitLength = bitLength;
            ContentTypes = contentTypes;
            CreatedAt = createdAt;
            CreatedAtFilter = createdAtFilter;
            CreatorId = creatorId;
            Expiration = expiration;
            ExpirationFilter = expirationFilter;
            Id = id;
            Metadata = metadata;
            Mode = mode;
            Name = name;
            Payload = payload;
            PayloadContentEncoding = payloadContentEncoding;
            PayloadContentType = payloadContentType;
            Region = region;
            SecretRef = secretRef;
            SecretType = secretType;
            Status = status;
            UpdatedAt = updatedAt;
            UpdatedAtFilter = updatedAtFilter;
        }
    }
}
