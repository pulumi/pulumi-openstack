// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    /// <summary>
    /// ## Import
    /// 
    /// Security Groups can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:networking/secGroup:SecGroup secgroup_1 38809219-5e8a-4852-9139-6f461c90e8bc
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:networking/secGroup:SecGroup")]
    public partial class SecGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The collection of tags assigned on the security group, which have
        /// been explicitly and implicitly added.
        /// </summary>
        [Output("allTags")]
        public Output<ImmutableArray<string>> AllTags { get; private set; } = null!;

        /// <summary>
        /// Whether or not to delete the default
        /// egress security rules. This is `false` by default. See the below note
        /// for more information.
        /// </summary>
        [Output("deleteDefaultRules")]
        public Output<bool?> DeleteDefaultRules { get; private set; } = null!;

        /// <summary>
        /// A unique name for the security group.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A unique name for the security group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a port. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security group.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Indicates if the security group is stateful or
        /// stateless. Update of the stateful argument is allowed when there is no port
        /// associated with the security group. Available only in OpenStack environments
        /// with the `stateful-security-group` extension. Defaults to true.
        /// </summary>
        [Output("stateful")]
        public Output<bool> Stateful { get; private set; } = null!;

        /// <summary>
        /// A set of string tags for the security group.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The owner of the security group. Required if admin
        /// wants to create a port for another tenant. Changing this creates a new
        /// security group.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a SecGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecGroup(string name, SecGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:networking/secGroup:SecGroup", name, args ?? new SecGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecGroup(string name, Input<string> id, SecGroupState? state = null, CustomResourceOptions? options = null)
            : base("openstack:networking/secGroup:SecGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecGroup Get(string name, Input<string> id, SecGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SecGroup(name, id, state, options);
        }
    }

    public sealed class SecGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not to delete the default
        /// egress security rules. This is `false` by default. See the below note
        /// for more information.
        /// </summary>
        [Input("deleteDefaultRules")]
        public Input<bool>? DeleteDefaultRules { get; set; }

        /// <summary>
        /// A unique name for the security group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for the security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a port. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates if the security group is stateful or
        /// stateless. Update of the stateful argument is allowed when there is no port
        /// associated with the security group. Available only in OpenStack environments
        /// with the `stateful-security-group` extension. Defaults to true.
        /// </summary>
        [Input("stateful")]
        public Input<bool>? Stateful { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the security group.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the security group. Required if admin
        /// wants to create a port for another tenant. Changing this creates a new
        /// security group.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public SecGroupArgs()
        {
        }
        public static new SecGroupArgs Empty => new SecGroupArgs();
    }

    public sealed class SecGroupState : global::Pulumi.ResourceArgs
    {
        [Input("allTags")]
        private InputList<string>? _allTags;

        /// <summary>
        /// The collection of tags assigned on the security group, which have
        /// been explicitly and implicitly added.
        /// </summary>
        public InputList<string> AllTags
        {
            get => _allTags ?? (_allTags = new InputList<string>());
            set => _allTags = value;
        }

        /// <summary>
        /// Whether or not to delete the default
        /// egress security rules. This is `false` by default. See the below note
        /// for more information.
        /// </summary>
        [Input("deleteDefaultRules")]
        public Input<bool>? DeleteDefaultRules { get; set; }

        /// <summary>
        /// A unique name for the security group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for the security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 networking client.
        /// A networking client is needed to create a port. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security group.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Indicates if the security group is stateful or
        /// stateless. Update of the stateful argument is allowed when there is no port
        /// associated with the security group. Available only in OpenStack environments
        /// with the `stateful-security-group` extension. Defaults to true.
        /// </summary>
        [Input("stateful")]
        public Input<bool>? Stateful { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of string tags for the security group.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the security group. Required if admin
        /// wants to create a port for another tenant. Changing this creates a new
        /// security group.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public SecGroupState()
        {
        }
        public static new SecGroupState Empty => new SecGroupState();
    }
}
