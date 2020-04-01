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
    /// will be stored in the raw state as plain-text. [Read more about sensitive data
    /// in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; **Note:** An Application Credential is created within the authenticated user
    /// project scope and is not visible by an admin or other accounts.
    /// The Application Credential visibility is similar to
    /// `openstack.compute.Keypair`.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_application_credential_v3.html.markdown.
    /// </summary>
    public partial class ApplicationCredential : Pulumi.CustomResource
    {
        /// <summary>
        /// A collection of one or more access rules, which
        /// this application credential allows to follow. The structure is described
        /// below. Changing this creates a new application credential.
        /// </summary>
        [Output("accessRules")]
        public Output<ImmutableArray<Outputs.ApplicationCredentialAccessRules>> AccessRules { get; private set; } = null!;

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
            : base("openstack:identity/applicationCredential:ApplicationCredential", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class ApplicationCredentialArgs : Pulumi.ResourceArgs
    {
        [Input("accessRules")]
        private InputList<Inputs.ApplicationCredentialAccessRulesArgs>? _accessRules;

        /// <summary>
        /// A collection of one or more access rules, which
        /// this application credential allows to follow. The structure is described
        /// below. Changing this creates a new application credential.
        /// </summary>
        public InputList<Inputs.ApplicationCredentialAccessRulesArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.ApplicationCredentialAccessRulesArgs>());
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

        /// <summary>
        /// The secret for the application credential. If omitted,
        /// it will be generated by the server. Changing this creates a new application
        /// credential.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

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
    }

    public sealed class ApplicationCredentialState : Pulumi.ResourceArgs
    {
        [Input("accessRules")]
        private InputList<Inputs.ApplicationCredentialAccessRulesGetArgs>? _accessRules;

        /// <summary>
        /// A collection of one or more access rules, which
        /// this application credential allows to follow. The structure is described
        /// below. Changing this creates a new application credential.
        /// </summary>
        public InputList<Inputs.ApplicationCredentialAccessRulesGetArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.ApplicationCredentialAccessRulesGetArgs>());
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

        /// <summary>
        /// The secret for the application credential. If omitted,
        /// it will be generated by the server. Changing this creates a new application
        /// credential.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

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
    }

    namespace Inputs
    {

    public sealed class ApplicationCredentialAccessRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the existing access rule. The access rule ID of
        /// another application credential can be provided.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The request method that the application credential is
        /// permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
        /// `HEAD`, `PATCH`, `PUT` and `DELETE`.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// The API path that the application credential is permitted
        /// to access. May use named wildcards such as **{tag}** or the unnamed wildcard
        /// **\*** to match against any string in the path up to a **/**, or the recursive
        /// wildcard **\*\*** to include **/** in the matched path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The service type identifier for the service that the
        /// application credential is granted to access. Must be a service type that is
        /// listed in the service catalog and not a code name for a service. E.g.
        /// **identity**, **compute**, **volumev3**, **image**, **network**,
        /// **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ApplicationCredentialAccessRulesArgs()
        {
        }
    }

    public sealed class ApplicationCredentialAccessRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the existing access rule. The access rule ID of
        /// another application credential can be provided.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The request method that the application credential is
        /// permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
        /// `HEAD`, `PATCH`, `PUT` and `DELETE`.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// The API path that the application credential is permitted
        /// to access. May use named wildcards such as **{tag}** or the unnamed wildcard
        /// **\*** to match against any string in the path up to a **/**, or the recursive
        /// wildcard **\*\*** to include **/** in the matched path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The service type identifier for the service that the
        /// application credential is granted to access. Must be a service type that is
        /// listed in the service catalog and not a code name for a service. E.g.
        /// **identity**, **compute**, **volumev3**, **image**, **network**,
        /// **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ApplicationCredentialAccessRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ApplicationCredentialAccessRules
    {
        /// <summary>
        /// The ID of the existing access rule. The access rule ID of
        /// another application credential can be provided.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The request method that the application credential is
        /// permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
        /// `HEAD`, `PATCH`, `PUT` and `DELETE`.
        /// </summary>
        public readonly string Method;
        /// <summary>
        /// The API path that the application credential is permitted
        /// to access. May use named wildcards such as **{tag}** or the unnamed wildcard
        /// **\*** to match against any string in the path up to a **/**, or the recursive
        /// wildcard **\*\*** to include **/** in the matched path.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The service type identifier for the service that the
        /// application credential is granted to access. Must be a service type that is
        /// listed in the service catalog and not a code name for a service. E.g.
        /// **identity**, **compute**, **volumev3**, **image**, **network**,
        /// **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private ApplicationCredentialAccessRules(
            string id,
            string method,
            string path,
            string service)
        {
            Id = id;
            Method = method;
            Path = path;
            Service = service;
        }
    }
    }
}
