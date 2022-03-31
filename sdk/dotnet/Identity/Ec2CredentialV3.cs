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
    /// ## Example Usage
    /// ### EC2 credential in current project scope
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ec2Key1 = new OpenStack.Identity.Ec2CredentialV3("ec2Key1", new OpenStack.Identity.Ec2CredentialV3Args
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### EC2 credential in pre-defined project scope
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ec2Key1 = new OpenStack.Identity.Ec2CredentialV3("ec2Key1", new OpenStack.Identity.Ec2CredentialV3Args
    ///         {
    ///             ProjectId = "f7ac731cc11f40efbc03a9f9e1d1d21f",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// EC2 Credentials can be imported using the `access`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:identity/ec2CredentialV3:Ec2CredentialV3 ec2_cred_1 2d0ac4a2f81b4b0f9513ee49e780647d
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/ec2CredentialV3:Ec2CredentialV3")]
    public partial class Ec2CredentialV3 : Pulumi.CustomResource
    {
        /// <summary>
        /// contains an EC2 credential access UUID
        /// </summary>
        [Output("access")]
        public Output<string> Access { get; private set; } = null!;

        /// <summary>
        /// The ID of the project the EC2 credential is created
        /// for and that authentication requests using this EC2 credential will
        /// be scoped to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new EC2 credential.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// contains an EC2 credential secret UUID
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// contains an EC2 credential trust ID scope
        /// </summary>
        [Output("trustId")]
        public Output<string> TrustId { get; private set; } = null!;

        /// <summary>
        /// The ID of the user the EC2 credential is created for.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a Ec2CredentialV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ec2CredentialV3(string name, Ec2CredentialV3Args? args = null, CustomResourceOptions? options = null)
            : base("openstack:identity/ec2CredentialV3:Ec2CredentialV3", name, args ?? new Ec2CredentialV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private Ec2CredentialV3(string name, Input<string> id, Ec2CredentialV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/ec2CredentialV3:Ec2CredentialV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ec2CredentialV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ec2CredentialV3 Get(string name, Input<string> id, Ec2CredentialV3State? state = null, CustomResourceOptions? options = null)
        {
            return new Ec2CredentialV3(name, id, state, options);
        }
    }

    public sealed class Ec2CredentialV3Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the project the EC2 credential is created
        /// for and that authentication requests using this EC2 credential will
        /// be scoped to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new EC2 credential.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the user the EC2 credential is created for.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public Ec2CredentialV3Args()
        {
        }
    }

    public sealed class Ec2CredentialV3State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// contains an EC2 credential access UUID
        /// </summary>
        [Input("access")]
        public Input<string>? Access { get; set; }

        /// <summary>
        /// The ID of the project the EC2 credential is created
        /// for and that authentication requests using this EC2 credential will
        /// be scoped to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used. Changing this
        /// creates a new EC2 credential.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// contains an EC2 credential secret UUID
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// contains an EC2 credential trust ID scope
        /// </summary>
        [Input("trustId")]
        public Input<string>? TrustId { get; set; }

        /// <summary>
        /// The ID of the user the EC2 credential is created for.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public Ec2CredentialV3State()
        {
        }
    }
}
