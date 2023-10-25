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
    /// Manages a V3 Application Credential resource within OpenStack Keystone.
    /// 
    /// &gt; **Note:** All arguments including the application credential name and secret
    /// will be stored in the raw state as plain-text. Read more about sensitive data
    /// in state.
    /// 
    /// &gt; **Note:** An Application Credential is created within the authenticated user
    /// project scope and is not visible by an admin or other accounts.
    /// The Application Credential visibility is similar to
    /// `openstack.compute.Keypair`.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Application Credentials can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:identity/applicationCredential:ApplicationCredential application_credential_1 c17304b7-0953-4738-abb0-67005882b0a0
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/applicationCredential:ApplicationCredential")]
    public partial class ApplicationCredential : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A collection of one or more access rules, which
        /// this application credential allows to follow. The structure is described
        /// below. Changing this creates a new application credential.
        /// </summary>
        [Output("accessRules")]
        public Output<ImmutableArray<Outputs.ApplicationCredentialAccessRule>> AccessRules { get; private set; } = null!;

        /// <summary>
        /// A description of the application credential.
        /// Changing this creates a new application credential.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The expiration time of the application credential
        /// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
        /// an application credential will never expire. Changing this creates a new
        /// application credential.
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// A name of the application credential. Changing this
        /// creates a new application credential.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project the application credential was created
        /// for and that authentication requests using this application credential will
        /// be scoped to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new application credential.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A collection of one or more role names, which this
        /// application credential has to be associated with its project. If omitted,
        /// all the current user's roles within the scoped project will be inherited by
        /// a new application credential. Changing this creates a new application
        /// credential.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// The secret for the application credential. If omitted,
        /// it will be generated by the server. Changing this creates a new application
        /// credential.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// A flag indicating whether the application
        /// credential may be used for creation or destruction of other application
        /// credentials or trusts. Changing this creates a new application credential.
        /// </summary>
        [Output("unrestricted")]
        public Output<bool?> Unrestricted { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationCredential resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationCredential(string name, ApplicationCredentialArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:identity/applicationCredential:ApplicationCredential", name, args ?? new ApplicationCredentialArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationCredential(string name, Input<string> id, ApplicationCredentialState? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/applicationCredential:ApplicationCredential", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationCredential resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationCredential Get(string name, Input<string> id, ApplicationCredentialState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationCredential(name, id, state, options);
        }
    }

    public sealed class ApplicationCredentialArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessRules")]
        private InputList<Inputs.ApplicationCredentialAccessRuleArgs>? _accessRules;

        /// <summary>
        /// A collection of one or more access rules, which
        /// this application credential allows to follow. The structure is described
        /// below. Changing this creates a new application credential.
        /// </summary>
        public InputList<Inputs.ApplicationCredentialAccessRuleArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.ApplicationCredentialAccessRuleArgs>());
            set => _accessRules = value;
        }

        /// <summary>
        /// A description of the application credential.
        /// Changing this creates a new application credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration time of the application credential
        /// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
        /// an application credential will never expire. Changing this creates a new
        /// application credential.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// A name of the application credential. Changing this
        /// creates a new application credential.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new application credential.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// A collection of one or more role names, which this
        /// application credential has to be associated with its project. If omitted,
        /// all the current user's roles within the scoped project will be inherited by
        /// a new application credential. Changing this creates a new application
        /// credential.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// The secret for the application credential. If omitted,
        /// it will be generated by the server. Changing this creates a new application
        /// credential.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// A flag indicating whether the application
        /// credential may be used for creation or destruction of other application
        /// credentials or trusts. Changing this creates a new application credential.
        /// </summary>
        [Input("unrestricted")]
        public Input<bool>? Unrestricted { get; set; }

        public ApplicationCredentialArgs()
        {
        }
        public static new ApplicationCredentialArgs Empty => new ApplicationCredentialArgs();
    }

    public sealed class ApplicationCredentialState : global::Pulumi.ResourceArgs
    {
        [Input("accessRules")]
        private InputList<Inputs.ApplicationCredentialAccessRuleGetArgs>? _accessRules;

        /// <summary>
        /// A collection of one or more access rules, which
        /// this application credential allows to follow. The structure is described
        /// below. Changing this creates a new application credential.
        /// </summary>
        public InputList<Inputs.ApplicationCredentialAccessRuleGetArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.ApplicationCredentialAccessRuleGetArgs>());
            set => _accessRules = value;
        }

        /// <summary>
        /// A description of the application credential.
        /// Changing this creates a new application credential.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration time of the application credential
        /// in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted,
        /// an application credential will never expire. Changing this creates a new
        /// application credential.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// A name of the application credential. Changing this
        /// creates a new application credential.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the application credential was created
        /// for and that authentication requests using this application credential will
        /// be scoped to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new application credential.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// A collection of one or more role names, which this
        /// application credential has to be associated with its project. If omitted,
        /// all the current user's roles within the scoped project will be inherited by
        /// a new application credential. Changing this creates a new application
        /// credential.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// The secret for the application credential. If omitted,
        /// it will be generated by the server. Changing this creates a new application
        /// credential.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// A flag indicating whether the application
        /// credential may be used for creation or destruction of other application
        /// credentials or trusts. Changing this creates a new application credential.
        /// </summary>
        [Input("unrestricted")]
        public Input<bool>? Unrestricted { get; set; }

        public ApplicationCredentialState()
        {
        }
        public static new ApplicationCredentialState Empty => new ApplicationCredentialState();
    }
}
