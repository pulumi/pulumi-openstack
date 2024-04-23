// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity
{
    /// <summary>
    /// Manages a V3 User resource within OpenStack Keystone.
    /// 
    /// &gt; **Note:** All arguments including the user password will be stored in the
    /// raw state as plain-text. Read more about sensitive data in
    /// state.
    /// 
    /// &gt; **Note:** You _must_ have admin privileges in your OpenStack cloud to use
    /// this resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project1 = new OpenStack.Identity.Project("project_1", new()
    ///     {
    ///         Name = "project_1",
    ///     });
    /// 
    ///     var user1 = new OpenStack.Identity.User("user_1", new()
    ///     {
    ///         DefaultProjectId = project1.Id,
    ///         Name = "user_1",
    ///         Description = "A user",
    ///         Password = "password123",
    ///         IgnoreChangePasswordUponFirstUse = true,
    ///         MultiFactorAuthEnabled = true,
    ///         MultiFactorAuthRules = new[]
    ///         {
    ///             new OpenStack.Identity.Inputs.UserMultiFactorAuthRuleArgs
    ///             {
    ///                 Rules = new[]
    ///                 {
    ///                     "password",
    ///                     "totp",
    ///                 },
    ///             },
    ///             new OpenStack.Identity.Inputs.UserMultiFactorAuthRuleArgs
    ///             {
    ///                 Rules = new[]
    ///                 {
    ///                     "password",
    ///                 },
    ///             },
    ///         },
    ///         Extra = 
    ///         {
    ///             { "email", "user_1@foobar.com" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Users can be imported using the `id`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:identity/user:User user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The default project this user belongs to.
        /// </summary>
        [Output("defaultProjectId")]
        public Output<string> DefaultProjectId { get; private set; } = null!;

        /// <summary>
        /// A description of the user.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The domain this user belongs to.
        /// </summary>
        [Output("domainId")]
        public Output<string> DomainId { get; private set; } = null!;

        /// <summary>
        /// Whether the user is enabled or disabled. Valid
        /// values are `true` and `false`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Free-form key/value pairs of extra information.
        /// </summary>
        [Output("extra")]
        public Output<ImmutableDictionary<string, object>?> Extra { get; private set; } = null!;

        /// <summary>
        /// User will not have to
        /// change their password upon first use. Valid values are `true` and `false`.
        /// </summary>
        [Output("ignoreChangePasswordUponFirstUse")]
        public Output<bool?> IgnoreChangePasswordUponFirstUse { get; private set; } = null!;

        /// <summary>
        /// User will not have a failure
        /// lockout placed on their account. Valid values are `true` and `false`.
        /// </summary>
        [Output("ignoreLockoutFailureAttempts")]
        public Output<bool?> IgnoreLockoutFailureAttempts { get; private set; } = null!;

        /// <summary>
        /// User's password will not expire.
        /// Valid values are `true` and `false`.
        /// </summary>
        [Output("ignorePasswordExpiry")]
        public Output<bool?> IgnorePasswordExpiry { get; private set; } = null!;

        /// <summary>
        /// Whether to enable multi-factor
        /// authentication. Valid values are `true` and `false`.
        /// </summary>
        [Output("multiFactorAuthEnabled")]
        public Output<bool?> MultiFactorAuthEnabled { get; private set; } = null!;

        /// <summary>
        /// A multi-factor authentication rule.
        /// The structure is documented below. Please see the
        /// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        /// for more information on how to use mulit-factor rules.
        /// </summary>
        [Output("multiFactorAuthRules")]
        public Output<ImmutableArray<Outputs.UserMultiFactorAuthRule>> MultiFactorAuthRules { get; private set; } = null!;

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The password for the user.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new User.
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
        public User(string name, UserArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:identity/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/user:User", name, state, MakeResourceOptions(options, id))
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
        /// <summary>
        /// The default project this user belongs to.
        /// </summary>
        [Input("defaultProjectId")]
        public Input<string>? DefaultProjectId { get; set; }

        /// <summary>
        /// A description of the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain this user belongs to.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Whether the user is enabled or disabled. Valid
        /// values are `true` and `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("extra")]
        private InputMap<object>? _extra;

        /// <summary>
        /// Free-form key/value pairs of extra information.
        /// </summary>
        public InputMap<object> Extra
        {
            get => _extra ?? (_extra = new InputMap<object>());
            set => _extra = value;
        }

        /// <summary>
        /// User will not have to
        /// change their password upon first use. Valid values are `true` and `false`.
        /// </summary>
        [Input("ignoreChangePasswordUponFirstUse")]
        public Input<bool>? IgnoreChangePasswordUponFirstUse { get; set; }

        /// <summary>
        /// User will not have a failure
        /// lockout placed on their account. Valid values are `true` and `false`.
        /// </summary>
        [Input("ignoreLockoutFailureAttempts")]
        public Input<bool>? IgnoreLockoutFailureAttempts { get; set; }

        /// <summary>
        /// User's password will not expire.
        /// Valid values are `true` and `false`.
        /// </summary>
        [Input("ignorePasswordExpiry")]
        public Input<bool>? IgnorePasswordExpiry { get; set; }

        /// <summary>
        /// Whether to enable multi-factor
        /// authentication. Valid values are `true` and `false`.
        /// </summary>
        [Input("multiFactorAuthEnabled")]
        public Input<bool>? MultiFactorAuthEnabled { get; set; }

        [Input("multiFactorAuthRules")]
        private InputList<Inputs.UserMultiFactorAuthRuleArgs>? _multiFactorAuthRules;

        /// <summary>
        /// A multi-factor authentication rule.
        /// The structure is documented below. Please see the
        /// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        /// for more information on how to use mulit-factor rules.
        /// </summary>
        public InputList<Inputs.UserMultiFactorAuthRuleArgs> MultiFactorAuthRules
        {
            get => _multiFactorAuthRules ?? (_multiFactorAuthRules = new InputList<Inputs.UserMultiFactorAuthRuleArgs>());
            set => _multiFactorAuthRules = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password for the user.
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
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new User.
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
        /// <summary>
        /// The default project this user belongs to.
        /// </summary>
        [Input("defaultProjectId")]
        public Input<string>? DefaultProjectId { get; set; }

        /// <summary>
        /// A description of the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain this user belongs to.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Whether the user is enabled or disabled. Valid
        /// values are `true` and `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("extra")]
        private InputMap<object>? _extra;

        /// <summary>
        /// Free-form key/value pairs of extra information.
        /// </summary>
        public InputMap<object> Extra
        {
            get => _extra ?? (_extra = new InputMap<object>());
            set => _extra = value;
        }

        /// <summary>
        /// User will not have to
        /// change their password upon first use. Valid values are `true` and `false`.
        /// </summary>
        [Input("ignoreChangePasswordUponFirstUse")]
        public Input<bool>? IgnoreChangePasswordUponFirstUse { get; set; }

        /// <summary>
        /// User will not have a failure
        /// lockout placed on their account. Valid values are `true` and `false`.
        /// </summary>
        [Input("ignoreLockoutFailureAttempts")]
        public Input<bool>? IgnoreLockoutFailureAttempts { get; set; }

        /// <summary>
        /// User's password will not expire.
        /// Valid values are `true` and `false`.
        /// </summary>
        [Input("ignorePasswordExpiry")]
        public Input<bool>? IgnorePasswordExpiry { get; set; }

        /// <summary>
        /// Whether to enable multi-factor
        /// authentication. Valid values are `true` and `false`.
        /// </summary>
        [Input("multiFactorAuthEnabled")]
        public Input<bool>? MultiFactorAuthEnabled { get; set; }

        [Input("multiFactorAuthRules")]
        private InputList<Inputs.UserMultiFactorAuthRuleGetArgs>? _multiFactorAuthRules;

        /// <summary>
        /// A multi-factor authentication rule.
        /// The structure is documented below. Please see the
        /// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        /// for more information on how to use mulit-factor rules.
        /// </summary>
        public InputList<Inputs.UserMultiFactorAuthRuleGetArgs> MultiFactorAuthRules
        {
            get => _multiFactorAuthRules ?? (_multiFactorAuthRules = new InputList<Inputs.UserMultiFactorAuthRuleGetArgs>());
            set => _multiFactorAuthRules = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password for the user.
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
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new User.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
