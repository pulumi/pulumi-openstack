// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Networking
{
    /// <summary>
    /// Manages a V2 Neutron addressscope resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_addressscope_v2.html.markdown.
    /// </summary>
    public partial class AddressScope : Pulumi.CustomResource
    {
        /// <summary>
        /// IP version, either 4 (default) or 6. Changing this
        /// creates a new address-scope.
        /// </summary>
        [Output("ipVersion")]
        public Output<int?> IpVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the address-scope. Changing this updates the
        /// name of the existing address-scope.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner of the address-scope. Required if admin
        /// wants to create a address-scope for another project. Changing this creates a
        /// new address-scope.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron address-scope. If omitted,
        /// the `region` argument of the provider is used. Changing this creates a new
        /// address-scope.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this address-scope is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// address-scope.
        /// </summary>
        [Output("shared")]
        public Output<bool> Shared { get; private set; } = null!;


        /// <summary>
        /// Create a AddressScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AddressScope(string name, AddressScopeArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:networking/addressScope:AddressScope", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AddressScope(string name, Input<string> id, AddressScopeState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/addressScope:AddressScope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AddressScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AddressScope Get(string name, Input<string> id, AddressScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new AddressScope(name, id, state, options);
        }
    }

    public sealed class AddressScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP version, either 4 (default) or 6. Changing this
        /// creates a new address-scope.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// The name of the address-scope. Changing this updates the
        /// name of the existing address-scope.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the address-scope. Required if admin
        /// wants to create a address-scope for another project. Changing this creates a
        /// new address-scope.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron address-scope. If omitted,
        /// the `region` argument of the provider is used. Changing this creates a new
        /// address-scope.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates whether this address-scope is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// address-scope.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        public AddressScopeArgs()
        {
        }
    }

    public sealed class AddressScopeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP version, either 4 (default) or 6. Changing this
        /// creates a new address-scope.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// The name of the address-scope. Changing this updates the
        /// name of the existing address-scope.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the address-scope. Required if admin
        /// wants to create a address-scope for another project. Changing this creates a
        /// new address-scope.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Networking client.
        /// A Networking client is needed to create a Neutron address-scope. If omitted,
        /// the `region` argument of the provider is used. Changing this creates a new
        /// address-scope.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates whether this address-scope is shared across
        /// all projects. Changing this updates the shared status of the existing
        /// address-scope.
        /// </summary>
        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        public AddressScopeState()
        {
        }
    }
}
