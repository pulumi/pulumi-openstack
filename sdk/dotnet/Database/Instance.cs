// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Database
{
    /// <summary>
    /// Manages a V1 DB instance resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/db_instance_v1.html.markdown.
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
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
        public Output<ImmutableArray<Outputs.InstanceDatabases>> Databases { get; private set; } = null!;

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
        public Output<ImmutableArray<Outputs.InstanceNetworks>> Networks { get; private set; } = null!;

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
        public Output<ImmutableArray<Outputs.InstanceUsers>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("openstack:database/instance:Instance", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration ID to be attached to the instance. Database instance
        /// will be rebooted when configuration is detached.
        /// </summary>
        [Input("configurationId")]
        public Input<string>? ConfigurationId { get; set; }

        [Input("databases")]
        private InputList<Inputs.InstanceDatabasesArgs>? _databases;

        /// <summary>
        /// An array of database name, charset and collate. The database
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceDatabasesArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.InstanceDatabasesArgs>());
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
        private InputList<Inputs.InstanceNetworksArgs>? _networks;

        /// <summary>
        /// An array of one or more networks to attach to the
        /// instance. The network object structure is documented below. Changing this
        /// creates a new instance.
        /// </summary>
        public InputList<Inputs.InstanceNetworksArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.InstanceNetworksArgs>());
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
        private InputList<Inputs.InstanceUsersArgs>? _users;

        /// <summary>
        /// An array of username, password, host and databases. The user
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceUsersArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.InstanceUsersArgs>());
            set => _users = value;
        }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration ID to be attached to the instance. Database instance
        /// will be rebooted when configuration is detached.
        /// </summary>
        [Input("configurationId")]
        public Input<string>? ConfigurationId { get; set; }

        [Input("databases")]
        private InputList<Inputs.InstanceDatabasesGetArgs>? _databases;

        /// <summary>
        /// An array of database name, charset and collate. The database
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceDatabasesGetArgs> Databases
        {
            get => _databases ?? (_databases = new InputList<Inputs.InstanceDatabasesGetArgs>());
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
        private InputList<Inputs.InstanceNetworksGetArgs>? _networks;

        /// <summary>
        /// An array of one or more networks to attach to the
        /// instance. The network object structure is documented below. Changing this
        /// creates a new instance.
        /// </summary>
        public InputList<Inputs.InstanceNetworksGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.InstanceNetworksGetArgs>());
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
        private InputList<Inputs.InstanceUsersGetArgs>? _users;

        /// <summary>
        /// An array of username, password, host and databases. The user
        /// object structure is documented below.
        /// </summary>
        public InputList<Inputs.InstanceUsersGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.InstanceUsersGetArgs>());
            set => _users = value;
        }

        public InstanceState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class InstanceDatabasesArgs : Pulumi.ResourceArgs
    {
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        [Input("collate")]
        public Input<string>? Collate { get; set; }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public InstanceDatabasesArgs()
        {
        }
    }

    public sealed class InstanceDatabasesGetArgs : Pulumi.ResourceArgs
    {
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        [Input("collate")]
        public Input<string>? Collate { get; set; }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public InstanceDatabasesGetArgs()
        {
        }
    }

    public sealed class InstanceDatastoreArgs : Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public InstanceDatastoreArgs()
        {
        }
    }

    public sealed class InstanceDatastoreGetArgs : Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public InstanceDatastoreGetArgs()
        {
        }
    }

    public sealed class InstanceNetworksArgs : Pulumi.ResourceArgs
    {
        [Input("fixedIpV4")]
        public Input<string>? FixedIpV4 { get; set; }

        [Input("fixedIpV6")]
        public Input<string>? FixedIpV6 { get; set; }

        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public InstanceNetworksArgs()
        {
        }
    }

    public sealed class InstanceNetworksGetArgs : Pulumi.ResourceArgs
    {
        [Input("fixedIpV4")]
        public Input<string>? FixedIpV4 { get; set; }

        [Input("fixedIpV6")]
        public Input<string>? FixedIpV6 { get; set; }

        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public InstanceNetworksGetArgs()
        {
        }
    }

    public sealed class InstanceUsersArgs : Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<string>? _databases;
        public InputList<string> Databases
        {
            get => _databases ?? (_databases = new InputList<string>());
            set => _databases = value;
        }

        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password")]
        public Input<string>? Password { get; set; }

        public InstanceUsersArgs()
        {
        }
    }

    public sealed class InstanceUsersGetArgs : Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<string>? _databases;
        public InputList<string> Databases
        {
            get => _databases ?? (_databases = new InputList<string>());
            set => _databases = value;
        }

        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password")]
        public Input<string>? Password { get; set; }

        public InstanceUsersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class InstanceDatabases
    {
        public readonly string? Charset;
        public readonly string? Collate;
        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private InstanceDatabases(
            string? charset,
            string? collate,
            string name)
        {
            Charset = charset;
            Collate = collate;
            Name = name;
        }
    }

    [OutputType]
    public sealed class InstanceDatastore
    {
        public readonly string Type;
        public readonly string Version;

        [OutputConstructor]
        private InstanceDatastore(
            string type,
            string version)
        {
            Type = type;
            Version = version;
        }
    }

    [OutputType]
    public sealed class InstanceNetworks
    {
        public readonly string? FixedIpV4;
        public readonly string? FixedIpV6;
        public readonly string? Port;
        public readonly string? Uuid;

        [OutputConstructor]
        private InstanceNetworks(
            string? fixedIpV4,
            string? fixedIpV6,
            string? port,
            string? uuid)
        {
            FixedIpV4 = fixedIpV4;
            FixedIpV6 = fixedIpV6;
            Port = port;
            Uuid = uuid;
        }
    }

    [OutputType]
    public sealed class InstanceUsers
    {
        public readonly ImmutableArray<string> Databases;
        public readonly string? Host;
        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        public readonly string Name;
        public readonly string? Password;

        [OutputConstructor]
        private InstanceUsers(
            ImmutableArray<string> databases,
            string? host,
            string name,
            string? password)
        {
            Databases = databases;
            Host = host;
            Name = name;
            Password = password;
        }
    }
    }
}
