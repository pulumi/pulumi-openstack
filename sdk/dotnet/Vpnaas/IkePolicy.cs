// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Vpnaas
{
    /// <summary>
    /// Manages a V2 Neutron IKE policy resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/vpnaas_ike_policy_v2.html.markdown.
    /// </summary>
    public partial class IkePolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
        /// Default is sha1. Changing this updates the algorithm of the existing policy.
        /// </summary>
        [Output("authAlgorithm")]
        public Output<string?> AuthAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The human-readable description for the policy.
        /// Changing this updates the description of the existing policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        /// The default value is aes-128. Changing this updates the existing policy.
        /// </summary>
        [Output("encryptionAlgorithm")]
        public Output<string?> EncryptionAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The IKE mode. A valid value is v1 or v2. Default is v1.
        /// Changing this updates the existing policy.
        /// </summary>
        [Output("ikeVersion")]
        public Output<string?> IkeVersion { get; private set; } = null!;

        /// <summary>
        /// The lifetime of the security association. Consists of Unit and Value.
        /// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
        /// Default is seconds.
        /// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
        /// Default is 3600.
        /// </summary>
        [Output("lifetimes")]
        public Output<ImmutableArray<Outputs.IkePolicyLifetimes>> Lifetimes { get; private set; } = null!;

        /// <summary>
        /// The name of the policy. Changing this updates the name of
        /// the existing policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
        /// Changing this updates the existing policy.
        /// </summary>
        [Output("pfs")]
        public Output<string?> Pfs { get; private set; } = null!;

        /// <summary>
        /// The IKE mode. A valid value is main, which is the default.
        /// Changing this updates the existing policy.
        /// </summary>
        [Output("phase1NegotiationMode")]
        public Output<string?> Phase1NegotiationMode { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a VPN service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// service.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The owner of the policy. Required if admin wants to
        /// create a service for another policy. Changing this creates a new policy.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a IkePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IkePolicy(string name, IkePolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/ikePolicy:IkePolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private IkePolicy(string name, Input<string> id, IkePolicyState? state = null, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/ikePolicy:IkePolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IkePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IkePolicy Get(string name, Input<string> id, IkePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IkePolicy(name, id, state, options);
        }
    }

    public sealed class IkePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
        /// Default is sha1. Changing this updates the algorithm of the existing policy.
        /// </summary>
        [Input("authAlgorithm")]
        public Input<string>? AuthAlgorithm { get; set; }

        /// <summary>
        /// The human-readable description for the policy.
        /// Changing this updates the description of the existing policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        /// The default value is aes-128. Changing this updates the existing policy.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        /// <summary>
        /// The IKE mode. A valid value is v1 or v2. Default is v1.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("ikeVersion")]
        public Input<string>? IkeVersion { get; set; }

        [Input("lifetimes")]
        private InputList<Inputs.IkePolicyLifetimesArgs>? _lifetimes;

        /// <summary>
        /// The lifetime of the security association. Consists of Unit and Value.
        /// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
        /// Default is seconds.
        /// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
        /// Default is 3600.
        /// </summary>
        public InputList<Inputs.IkePolicyLifetimesArgs> Lifetimes
        {
            get => _lifetimes ?? (_lifetimes = new InputList<Inputs.IkePolicyLifetimesArgs>());
            set => _lifetimes = value;
        }

        /// <summary>
        /// The name of the policy. Changing this updates the name of
        /// the existing policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("pfs")]
        public Input<string>? Pfs { get; set; }

        /// <summary>
        /// The IKE mode. A valid value is main, which is the default.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("phase1NegotiationMode")]
        public Input<string>? Phase1NegotiationMode { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a VPN service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// service.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the policy. Required if admin wants to
        /// create a service for another policy. Changing this creates a new policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public IkePolicyArgs()
        {
        }
    }

    public sealed class IkePolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
        /// Default is sha1. Changing this updates the algorithm of the existing policy.
        /// </summary>
        [Input("authAlgorithm")]
        public Input<string>? AuthAlgorithm { get; set; }

        /// <summary>
        /// The human-readable description for the policy.
        /// Changing this updates the description of the existing policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        /// The default value is aes-128. Changing this updates the existing policy.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        /// <summary>
        /// The IKE mode. A valid value is v1 or v2. Default is v1.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("ikeVersion")]
        public Input<string>? IkeVersion { get; set; }

        [Input("lifetimes")]
        private InputList<Inputs.IkePolicyLifetimesGetArgs>? _lifetimes;

        /// <summary>
        /// The lifetime of the security association. Consists of Unit and Value.
        /// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
        /// Default is seconds.
        /// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
        /// Default is 3600.
        /// </summary>
        public InputList<Inputs.IkePolicyLifetimesGetArgs> Lifetimes
        {
            get => _lifetimes ?? (_lifetimes = new InputList<Inputs.IkePolicyLifetimesGetArgs>());
            set => _lifetimes = value;
        }

        /// <summary>
        /// The name of the policy. Changing this updates the name of
        /// the existing policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("pfs")]
        public Input<string>? Pfs { get; set; }

        /// <summary>
        /// The IKE mode. A valid value is main, which is the default.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("phase1NegotiationMode")]
        public Input<string>? Phase1NegotiationMode { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a VPN service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// service.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the policy. Required if admin wants to
        /// create a service for another policy. Changing this creates a new policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

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

        public IkePolicyState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class IkePolicyLifetimesArgs : Pulumi.ResourceArgs
    {
        [Input("units")]
        public Input<string>? Units { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public IkePolicyLifetimesArgs()
        {
        }
    }

    public sealed class IkePolicyLifetimesGetArgs : Pulumi.ResourceArgs
    {
        [Input("units")]
        public Input<string>? Units { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public IkePolicyLifetimesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class IkePolicyLifetimes
    {
        public readonly string Units;
        public readonly int Value;

        [OutputConstructor]
        private IkePolicyLifetimes(
            string units,
            int value)
        {
            Units = units;
            Value = value;
        }
    }
    }
}
