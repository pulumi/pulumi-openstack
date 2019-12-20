// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Identity
{
    /// <summary>
    /// Manages a V3 User resource within OpenStack Keystone.
    /// 
    /// Note: You _must_ have admin privileges in your OpenStack cloud to use
    /// this resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_user_v3.html.markdown.
    /// </summary>
    public partial class User : Pulumi.CustomResource
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
        public Output<ImmutableArray<Outputs.UserMultiFactorAuthRules>> MultiFactorAuthRules { get; private set; } = null!;

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
            : base("openstack:identity/user:User", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class UserArgs : Pulumi.ResourceArgs
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
        private InputList<Inputs.UserMultiFactorAuthRulesArgs>? _multiFactorAuthRules;

        /// <summary>
        /// A multi-factor authentication rule.
        /// The structure is documented below. Please see the
        /// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        /// for more information on how to use mulit-factor rules.
        /// </summary>
        public InputList<Inputs.UserMultiFactorAuthRulesArgs> MultiFactorAuthRules
        {
            get => _multiFactorAuthRules ?? (_multiFactorAuthRules = new InputList<Inputs.UserMultiFactorAuthRulesArgs>());
            set => _multiFactorAuthRules = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password for the user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

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
    }

    public sealed class UserState : Pulumi.ResourceArgs
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
        private InputList<Inputs.UserMultiFactorAuthRulesGetArgs>? _multiFactorAuthRules;

        /// <summary>
        /// A multi-factor authentication rule.
        /// The structure is documented below. Please see the
        /// [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        /// for more information on how to use mulit-factor rules.
        /// </summary>
        public InputList<Inputs.UserMultiFactorAuthRulesGetArgs> MultiFactorAuthRules
        {
            get => _multiFactorAuthRules ?? (_multiFactorAuthRules = new InputList<Inputs.UserMultiFactorAuthRulesGetArgs>());
            set => _multiFactorAuthRules = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password for the user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

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
    }

    namespace Inputs
    {

    public sealed class UserMultiFactorAuthRulesArgs : Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<string>? _rules;
        public InputList<string> Rules
        {
            get => _rules ?? (_rules = new InputList<string>());
            set => _rules = value;
        }

        public UserMultiFactorAuthRulesArgs()
        {
        }
    }

    public sealed class UserMultiFactorAuthRulesGetArgs : Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<string>? _rules;
        public InputList<string> Rules
        {
            get => _rules ?? (_rules = new InputList<string>());
            set => _rules = value;
        }

        public UserMultiFactorAuthRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class UserMultiFactorAuthRules
    {
        public readonly ImmutableArray<string> Rules;

        [OutputConstructor]
        private UserMultiFactorAuthRules(ImmutableArray<string> rules)
        {
            Rules = rules;
        }
    }
    }
}
