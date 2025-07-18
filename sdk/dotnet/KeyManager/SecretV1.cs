// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ### Simple secret
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var secret1 = new OpenStack.KeyManager.SecretV1("secret_1", new()
    ///     {
    ///         Algorithm = "aes",
    ///         BitLength = 256,
    ///         Mode = "cbc",
    ///         Name = "mysecret",
    ///         Payload = "foobar",
    ///         PayloadContentType = "text/plain",
    ///         SecretType = "passphrase",
    ///         Metadata = 
    ///         {
    ///             { "key", "foo" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Secret with whitespaces
    /// 
    /// &gt; **Note** If you want to store payload with leading or trailing whitespaces,
    /// it's recommended to store it in a base64 encoding. Plain text payload can also
    /// work, but further addind or removing of the leading or trailing whitespaces
    /// won't be detected as a state change, e.g. changing plain text payload from
    /// ` password  ` to `password` won't recreate the secret.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// using Std = Pulumi.Std;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var secret1 = new OpenStack.KeyManager.SecretV1("secret_1", new()
    ///     {
    ///         Name = "password",
    ///         Payload = Std.Base64encode.Invoke(new()
    ///         {
    ///             Input = "password with the whitespace at the end ",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         SecretType = "passphrase",
    ///         PayloadContentType = "application/octet-stream",
    ///         PayloadContentEncoding = "base64",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Secret with the ACL
    /// 
    /// &gt; **Note** Only read ACLs are supported
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// using Std = Pulumi.Std;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var secret1 = new OpenStack.KeyManager.SecretV1("secret_1", new()
    ///     {
    ///         Name = "certificate",
    ///         Payload = Std.File.Invoke(new()
    ///         {
    ///             Input = "certificate.pem",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///         SecretType = "certificate",
    ///         PayloadContentType = "text/plain",
    ///         Acl = new OpenStack.KeyManager.Inputs.SecretV1AclArgs
    ///         {
    ///             Read = new OpenStack.KeyManager.Inputs.SecretV1AclReadArgs
    ///             {
    ///                 ProjectAccess = false,
    ///                 Users = new[]
    ///                 {
    ///                     "userid1",
    ///                     "userid2",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Secrets can be imported using the secret id (the last part of the secret reference), e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:keymanager/secretV1:SecretV1 secret_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:keymanager/secretV1:SecretV1")]
    public partial class SecretV1 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allows to control an access to a secret. Currently only the
        /// `read` operation is supported. If not specified, the secret is accessible
        /// project wide.
        /// </summary>
        [Output("acl")]
        public Output<Outputs.SecretV1Acl> Acl { get; private set; } = null!;

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// The map of metadata, assigned on the secret, which has been
        /// explicitly and implicitly added.
        /// </summary>
        [Output("allMetadata")]
        public Output<ImmutableDictionary<string, string>> AllMetadata { get; private set; } = null!;

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Output("bitLength")]
        public Output<int> BitLength { get; private set; } = null!;

        /// <summary>
        /// The map of the content types, assigned on the secret.
        /// </summary>
        [Output("contentTypes")]
        public Output<ImmutableDictionary<string, string>> ContentTypes { get; private set; } = null!;

        /// <summary>
        /// The date the secret was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The creator of the secret.
        /// </summary>
        [Output("creatorId")]
        public Output<string> CreatorId { get; private set; } = null!;

        /// <summary>
        /// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
        /// </summary>
        [Output("expiration")]
        public Output<string?> Expiration { get; private set; } = null!;

        /// <summary>
        /// Additional Metadata for the secret.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Human-readable name for the Secret. Does not have
        /// to be unique.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
        /// </summary>
        [Output("payload")]
        public Output<string?> Payload { get; private set; } = null!;

        /// <summary>
        /// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
        /// </summary>
        [Output("payloadContentEncoding")]
        public Output<string?> PayloadContentEncoding { get; private set; } = null!;

        /// <summary>
        /// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        /// </summary>
        [Output("payloadContentType")]
        public Output<string?> PayloadContentType { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a secret. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 secret.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The secret reference / where to find the secret.
        /// </summary>
        [Output("secretRef")]
        public Output<string> SecretRef { get; private set; } = null!;

        /// <summary>
        /// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
        /// </summary>
        [Output("secretType")]
        public Output<string> SecretType { get; private set; } = null!;

        /// <summary>
        /// The status of the secret.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The date the secret was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a SecretV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretV1(string name, SecretV1Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:keymanager/secretV1:SecretV1", name, args ?? new SecretV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private SecretV1(string name, Input<string> id, SecretV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:keymanager/secretV1:SecretV1", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "payload",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretV1 Get(string name, Input<string> id, SecretV1State? state = null, CustomResourceOptions? options = null)
        {
            return new SecretV1(name, id, state, options);
        }
    }

    public sealed class SecretV1Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows to control an access to a secret. Currently only the
        /// `read` operation is supported. If not specified, the secret is accessible
        /// project wide.
        /// </summary>
        [Input("acl")]
        public Input<Inputs.SecretV1AclArgs>? Acl { get; set; }

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Input("bitLength")]
        public Input<int>? BitLength { get; set; }

        /// <summary>
        /// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
        /// </summary>
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Additional Metadata for the secret.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Human-readable name for the Secret. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("payload")]
        private Input<string>? _payload;

        /// <summary>
        /// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
        /// </summary>
        public Input<string>? Payload
        {
            get => _payload;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _payload = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
        /// </summary>
        [Input("payloadContentEncoding")]
        public Input<string>? PayloadContentEncoding { get; set; }

        /// <summary>
        /// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        /// </summary>
        [Input("payloadContentType")]
        public Input<string>? PayloadContentType { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a secret. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 secret.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
        /// </summary>
        [Input("secretType")]
        public Input<string>? SecretType { get; set; }

        public SecretV1Args()
        {
        }
        public static new SecretV1Args Empty => new SecretV1Args();
    }

    public sealed class SecretV1State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows to control an access to a secret. Currently only the
        /// `read` operation is supported. If not specified, the secret is accessible
        /// project wide.
        /// </summary>
        [Input("acl")]
        public Input<Inputs.SecretV1AclGetArgs>? Acl { get; set; }

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("allMetadata")]
        private InputMap<string>? _allMetadata;

        /// <summary>
        /// The map of metadata, assigned on the secret, which has been
        /// explicitly and implicitly added.
        /// </summary>
        public InputMap<string> AllMetadata
        {
            get => _allMetadata ?? (_allMetadata = new InputMap<string>());
            set => _allMetadata = value;
        }

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Input("bitLength")]
        public Input<int>? BitLength { get; set; }

        [Input("contentTypes")]
        private InputMap<string>? _contentTypes;

        /// <summary>
        /// The map of the content types, assigned on the secret.
        /// </summary>
        public InputMap<string> ContentTypes
        {
            get => _contentTypes ?? (_contentTypes = new InputMap<string>());
            set => _contentTypes = value;
        }

        /// <summary>
        /// The date the secret was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The creator of the secret.
        /// </summary>
        [Input("creatorId")]
        public Input<string>? CreatorId { get; set; }

        /// <summary>
        /// The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
        /// </summary>
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Additional Metadata for the secret.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Metadata provided by a user or system for informational purposes.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Human-readable name for the Secret. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("payload")]
        private Input<string>? _payload;

        /// <summary>
        /// The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
        /// </summary>
        public Input<string>? Payload
        {
            get => _payload;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _payload = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
        /// </summary>
        [Input("payloadContentEncoding")]
        public Input<string>? PayloadContentEncoding { get; set; }

        /// <summary>
        /// (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        /// </summary>
        [Input("payloadContentType")]
        public Input<string>? PayloadContentType { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a secret. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 secret.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The secret reference / where to find the secret.
        /// </summary>
        [Input("secretRef")]
        public Input<string>? SecretRef { get; set; }

        /// <summary>
        /// Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
        /// </summary>
        [Input("secretType")]
        public Input<string>? SecretType { get; set; }

        /// <summary>
        /// The status of the secret.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The date the secret was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public SecretV1State()
        {
        }
        public static new SecretV1State Empty => new SecretV1State();
    }
}
