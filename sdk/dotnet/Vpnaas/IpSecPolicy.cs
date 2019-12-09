// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.VPNaaS
{
    /// <summary>
    /// Manages a V2 Neutron IPSec policy resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/vpnaas_ipsec_policy_v2.html.markdown.
    /// </summary>
    public partial class IpSecPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
        /// Default is sha1. Changing this updates the algorithm of the existing policy.
        /// </summary>
        [Output("authAlgorithm")]
        public Output<string> AuthAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The human-readable description for the policy.
        /// Changing this updates the description of the existing policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
        /// Changing this updates the existing policy.
        /// </summary>
        [Output("encapsulationMode")]
        public Output<string> EncapsulationMode { get; private set; } = null!;

        /// <summary>
        /// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        /// The default value is aes-128. Changing this updates the existing policy.
        /// </summary>
        [Output("encryptionAlgorithm")]
        public Output<string> EncryptionAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The lifetime of the security association. Consists of Unit and Value.
        /// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
        /// Default is seconds.
        /// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
        /// Default is 3600.
        /// </summary>
        [Output("lifetimes")]
        public Output<ImmutableArray<Outputs.IpSecPolicyLifetimes>> Lifetimes { get; private set; } = null!;

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
        public Output<string> Pfs { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an IPSec policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// policy.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The owner of the policy. Required if admin wants to
        /// create a policy for another project. Changing this creates a new policy.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// The transform protocol. Valid values are ESP, AH and AH-ESP.
        /// Changing this updates the existing policy. Default is ESP.
        /// </summary>
        [Output("transformProtocol")]
        public Output<string> TransformProtocol { get; private set; } = null!;

        /// <summary>
        /// Map of additional options.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a IpSecPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpSecPolicy(string name, IpSecPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/ipSecPolicy:IpSecPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private IpSecPolicy(string name, Input<string> id, IpSecPolicyState? state = null, CustomResourceOptions? options = null)
            : base("openstack:vpnaas/ipSecPolicy:IpSecPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpSecPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpSecPolicy Get(string name, Input<string> id, IpSecPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IpSecPolicy(name, id, state, options);
        }
    }

    public sealed class IpSecPolicyArgs : Pulumi.ResourceArgs
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
        /// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("encapsulationMode")]
        public Input<string>? EncapsulationMode { get; set; }

        /// <summary>
        /// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        /// The default value is aes-128. Changing this updates the existing policy.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        [Input("lifetimes")]
        private InputList<Inputs.IpSecPolicyLifetimesArgs>? _lifetimes;

        /// <summary>
        /// The lifetime of the security association. Consists of Unit and Value.
        /// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
        /// Default is seconds.
        /// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
        /// Default is 3600.
        /// </summary>
        public InputList<Inputs.IpSecPolicyLifetimesArgs> Lifetimes
        {
            get => _lifetimes ?? (_lifetimes = new InputList<Inputs.IpSecPolicyLifetimesArgs>());
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
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an IPSec policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the policy. Required if admin wants to
        /// create a policy for another project. Changing this creates a new policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The transform protocol. Valid values are ESP, AH and AH-ESP.
        /// Changing this updates the existing policy. Default is ESP.
        /// </summary>
        [Input("transformProtocol")]
        public Input<string>? TransformProtocol { get; set; }

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

        public IpSecPolicyArgs()
        {
        }
    }

    public sealed class IpSecPolicyState : Pulumi.ResourceArgs
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
        /// The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
        /// Changing this updates the existing policy.
        /// </summary>
        [Input("encapsulationMode")]
        public Input<string>? EncapsulationMode { get; set; }

        /// <summary>
        /// The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        /// The default value is aes-128. Changing this updates the existing policy.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        [Input("lifetimes")]
        private InputList<Inputs.IpSecPolicyLifetimesGetArgs>? _lifetimes;

        /// <summary>
        /// The lifetime of the security association. Consists of Unit and Value.
        /// - `unit` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes.
        /// Default is seconds.
        /// - `value` - (Optional) The value for the lifetime of the security association. Must be a positive integer.
        /// Default is 3600.
        /// </summary>
        public InputList<Inputs.IpSecPolicyLifetimesGetArgs> Lifetimes
        {
            get => _lifetimes ?? (_lifetimes = new InputList<Inputs.IpSecPolicyLifetimesGetArgs>());
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
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create an IPSec policy. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// policy.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The owner of the policy. Required if admin wants to
        /// create a policy for another project. Changing this creates a new policy.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The transform protocol. Valid values are ESP, AH and AH-ESP.
        /// Changing this updates the existing policy. Default is ESP.
        /// </summary>
        [Input("transformProtocol")]
        public Input<string>? TransformProtocol { get; set; }

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

        public IpSecPolicyState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class IpSecPolicyLifetimesArgs : Pulumi.ResourceArgs
    {
        [Input("units")]
        public Input<string>? Units { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public IpSecPolicyLifetimesArgs()
        {
        }
    }

    public sealed class IpSecPolicyLifetimesGetArgs : Pulumi.ResourceArgs
    {
        [Input("units")]
        public Input<string>? Units { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public IpSecPolicyLifetimesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class IpSecPolicyLifetimes
    {
        public readonly string Units;
        public readonly int Value;

        [OutputConstructor]
        private IpSecPolicyLifetimes(
            string units,
            int value)
        {
            Units = units;
            Value = value;
        }
    }
    }
}
