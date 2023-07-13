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
    /// Manages a V1 DB user resource within OpenStack.
    /// 
    /// &gt; **Note:** All arguments including the database password will be stored in the
    /// raw state as plain-text. Read more about sensitive data in
    /// state.
    /// 
    /// ## Example Usage
    /// ### User
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basic = new OpenStack.Database.User("basic", new()
    ///     {
    ///         InstanceId = openstack_db_instance_v1.Basic.Id,
    ///         Password = "password",
    ///         Databases = new[]
    ///         {
    ///             "testdb",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:database/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of database user should have access to.
        /// </summary>
        [Output("databases")]
        public Output<ImmutableArray<string>> Databases { get; private set; } = null!;

        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// The ID for the database instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// User's password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Openstack region resource is created in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("openstack:database/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("openstack:database/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<string>? _databases;

        /// <summary>
        /// A list of database user should have access to.
        /// </summary>
        public InputList<string> Databases
        {
            get => _databases ?? (_databases = new InputList<string>());
            set => _databases = value;
        }

        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The ID for the database instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// User's password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Openstack region resource is created in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<string>? _databases;

        /// <summary>
        /// A list of database user should have access to.
        /// </summary>
        public InputList<string> Databases
        {
            get => _databases ?? (_databases = new InputList<string>());
            set => _databases = value;
        }

        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The ID for the database instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// A unique name for the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// User's password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Openstack region resource is created in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
