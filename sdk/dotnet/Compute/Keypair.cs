// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ### Import an Existing Public Key
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_keypair = new OpenStack.Compute.Keypair("test-keypair", new()
    ///     {
    ///         PublicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLotBCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAnOfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZqd9LvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TaIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIF61p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Generate a Public/Private Key Pair
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_keypair = new OpenStack.Compute.Keypair("test-keypair");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Keypairs can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:compute/keypair:Keypair my-keypair test-keypair
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:compute/keypair:Keypair")]
    public partial class Keypair : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The fingerprint of the public key.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// A unique name for the keypair. Changing this creates a new
        /// keypair.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The generated private key when no public key is specified.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// A pregenerated OpenSSH-formatted public key.
        /// Changing this creates a new keypair. If a public key is not specified, then
        /// a public/private key pair will be automatically generated. If a pair is
        /// created, then destroying this resource means you will lose access to that
        /// keypair forever.
        /// </summary>
        [Output("publicKey")]
        public Output<string> PublicKey { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new keypair.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// This allows administrative users to operate key-pairs
        /// of specified user ID. For this feature your need to have openstack microversion
        /// 2.10 (Liberty) or later.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a Keypair resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Keypair(string name, KeypairArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:compute/keypair:Keypair", name, args ?? new KeypairArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Keypair(string name, Input<string> id, KeypairState? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/keypair:Keypair", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "privateKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Keypair resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Keypair Get(string name, Input<string> id, KeypairState? state = null, CustomResourceOptions? options = null)
        {
            return new Keypair(name, id, state, options);
        }
    }

    public sealed class KeypairArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique name for the keypair. Changing this creates a new
        /// keypair.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A pregenerated OpenSSH-formatted public key.
        /// Changing this creates a new keypair. If a public key is not specified, then
        /// a public/private key pair will be automatically generated. If a pair is
        /// created, then destroying this resource means you will lose access to that
        /// keypair forever.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new keypair.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// This allows administrative users to operate key-pairs
        /// of specified user ID. For this feature your need to have openstack microversion
        /// 2.10 (Liberty) or later.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public KeypairArgs()
        {
        }
        public static new KeypairArgs Empty => new KeypairArgs();
    }

    public sealed class KeypairState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The fingerprint of the public key.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// A unique name for the keypair. Changing this creates a new
        /// keypair.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// The generated private key when no public key is specified.
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// A pregenerated OpenSSH-formatted public key.
        /// Changing this creates a new keypair. If a public key is not specified, then
        /// a public/private key pair will be automatically generated. If a pair is
        /// created, then destroying this resource means you will lose access to that
        /// keypair forever.
        /// </summary>
        [Input("publicKey")]
        public Input<string>? PublicKey { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new keypair.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// This allows administrative users to operate key-pairs
        /// of specified user ID. For this feature your need to have openstack microversion
        /// 2.10 (Liberty) or later.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public KeypairState()
        {
        }
        public static new KeypairState Empty => new KeypairState();
    }
}
