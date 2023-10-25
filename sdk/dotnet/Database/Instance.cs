// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Database
{
    /// <summary>
    /// Manages a V1 DB instance resource within OpenStack.
    /// 
    /// &gt; **Note:** All arguments including the instance user password will be stored
    /// in the raw state as plain-text. Read more about sensitive data in
    /// state.
    /// 
    /// ## Example Usage
    /// </summary>
    [OpenStackResourceType("openstack:database/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of IP addresses assigned to the instance.
        /// </summary>
        [Output("addresses")]
        public Output<ImmutableArray<string>> Addresses { get; private set; } = null!;

        /// <summary>
        /// Configuration ID to be attached to the instance. Database instance
        /// will be rebooted when configuration is detached.
        /// </summary>
        [Output("configurationId")]
        public Output<string?> ConfigurationId { get; private set; } = null!;

        /// <summary>
        /// An array of database name, charset and collate. The database
        /// object structure is documented below.
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<Outputs.InstanceDatabase>> Databases { get; private set; } = null!;

        /// <summary>
        /// An array of database engine type and version. The datastore
        /// object structure is documented below. Changing this creates a new instance.
        /// </summary>
        [Output("datastore")]
        public Output<Outputs.InstanceDatastore> Datastore { get; private set; } = null!;

        /// <summary>
        /// The flavor ID of the desired flavor for the instance.
        /// Changing this creates new instance.
        /// </summary>
        [Output("flavorId")]
        public Output<string> FlavorId { get; private set; } = null!;

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An array of one or more networks to attach to the
        /// instance. The network object structure is documented below. Changing this
        /// creates a new instance.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<Outputs.InstanceNetwork>> Networks { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the db instance. Changing this
        /// creates a new instance.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Specifies the volume size in GB. Changing this creates new instance.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// An array of username, password, host and databases. The user
        /// object structure is documented below.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<Outputs.InstanceUser>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("openstack:database/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("openstack:database/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration ID to be attached to the instance. Database instance
        /// will be rebooted when configuration is detached.
        /// </summary>
        [Input("configurationId")]
        public Input<string>? ConfigurationId { get; set; }

        [Input("databases")]
        private InputList<Inputs.InstanceDatabaseArgs>? _databases;

        /// <summary>
        /// An array of database name, charset and collate. The database
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceDatabaseArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.InstanceDatabaseArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// An array of database engine type and version. The datastore
        /// object structure is documented below. Changing this creates a new instance.
        /// </summary>
        [Input("datastore", required: true)]
        public Input<Inputs.InstanceDatastoreArgs> Datastore { get; set; } = null!;

        /// <summary>
        /// The flavor ID of the desired flavor for the instance.
        /// Changing this creates new instance.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<Inputs.InstanceNetworkArgs>? _networks;

        /// <summary>
        /// An array of one or more networks to attach to the
        /// instance. The network object structure is documented below. Changing this
        /// creates a new instance.
        /// </summary>
        public InputList<Inputs.InstanceNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.InstanceNetworkArgs>());
            set => _networks = value;
        }

        /// <summary>
        /// The region in which to create the db instance. Changing this
        /// creates a new instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the volume size in GB. Changing this creates new instance.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("users")]
        private InputList<Inputs.InstanceUserArgs>? _users;

        /// <summary>
        /// An array of username, password, host and databases. The user
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceUserArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.InstanceUserArgs>());
            set => _users = value;
        }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<string>? _addresses;

        /// <summary>
        /// A list of IP addresses assigned to the instance.
        /// </summary>
        public InputList<string> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<string>());
            set => _addresses = value;
        }

        /// <summary>
        /// Configuration ID to be attached to the instance. Database instance
        /// will be rebooted when configuration is detached.
        /// </summary>
        [Input("configurationId")]
        public Input<string>? ConfigurationId { get; set; }

        [Input("databases")]
        private InputList<Inputs.InstanceDatabaseGetArgs>? _databases;

        /// <summary>
        /// An array of database name, charset and collate. The database
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceDatabaseGetArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.InstanceDatabaseGetArgs>());
            set => _databases = value;
        }

        /// <summary>
        /// An array of database engine type and version. The datastore
        /// object structure is documented below. Changing this creates a new instance.
        /// </summary>
        [Input("datastore")]
        public Input<Inputs.InstanceDatastoreGetArgs>? Datastore { get; set; }

        /// <summary>
        /// The flavor ID of the desired flavor for the instance.
        /// Changing this creates new instance.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("networks")]
        private InputList<Inputs.InstanceNetworkGetArgs>? _networks;

        /// <summary>
        /// An array of one or more networks to attach to the
        /// instance. The network object structure is documented below. Changing this
        /// creates a new instance.
        /// </summary>
        public InputList<Inputs.InstanceNetworkGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.InstanceNetworkGetArgs>());
            set => _networks = value;
        }

        /// <summary>
        /// The region in which to create the db instance. Changing this
        /// creates a new instance.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Specifies the volume size in GB. Changing this creates new instance.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        [Input("users")]
        private InputList<Inputs.InstanceUserGetArgs>? _users;

        /// <summary>
        /// An array of username, password, host and databases. The user
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.InstanceUserGetArgs>());
            set => _users = value;
        }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
