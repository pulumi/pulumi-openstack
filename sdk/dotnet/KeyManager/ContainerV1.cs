// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager
{
    /// <summary>
    /// Manages a V1 Barbican container resource within OpenStack.
    /// 
    /// ## Example Usage
    /// 
    /// ### Simple secret
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var certificate1 = new OpenStack.KeyManager.SecretV1("certificate1", new OpenStack.KeyManager.SecretV1Args
    ///         {
    ///             Payload = File.ReadAllText("cert.pem"),
    ///             PayloadContentType = "text/plain",
    ///             SecretType = "certificate",
    ///         });
    ///         var privateKey1 = new OpenStack.KeyManager.SecretV1("privateKey1", new OpenStack.KeyManager.SecretV1Args
    ///         {
    ///             Payload = File.ReadAllText("cert-key.pem"),
    ///             PayloadContentType = "text/plain",
    ///             SecretType = "private",
    ///         });
    ///         var intermediate1 = new OpenStack.KeyManager.SecretV1("intermediate1", new OpenStack.KeyManager.SecretV1Args
    ///         {
    ///             Payload = File.ReadAllText("intermediate-ca.pem"),
    ///             PayloadContentType = "text/plain",
    ///             SecretType = "certificate",
    ///         });
    ///         var tls1 = new OpenStack.KeyManager.ContainerV1("tls1", new OpenStack.KeyManager.ContainerV1Args
    ///         {
    ///             SecretRefs = 
    ///             {
    ///                 new OpenStack.KeyManager.Inputs.ContainerV1SecretRefArgs
    ///                 {
    ///                     Name = "certificate",
    ///                     SecretRef = certificate1.SecretRef,
    ///                 },
    ///                 new OpenStack.KeyManager.Inputs.ContainerV1SecretRefArgs
    ///                 {
    ///                     Name = "private_key",
    ///                     SecretRef = privateKey1.SecretRef,
    ///                 },
    ///                 new OpenStack.KeyManager.Inputs.ContainerV1SecretRefArgs
    ///                 {
    ///                     Name = "intermediates",
    ///                     SecretRef = intermediate1.SecretRef,
    ///                 },
    ///             },
    ///             Type = "certificate",
    ///         });
    ///         var subnet1 = Output.Create(OpenStack.Networking.GetSubnet.InvokeAsync(new OpenStack.Networking.GetSubnetArgs
    ///         {
    ///             Name = "my-subnet",
    ///         }));
    ///         var lb1 = new OpenStack.LoadBalancer.LoadBalancer("lb1", new OpenStack.LoadBalancer.LoadBalancerArgs
    ///         {
    ///             VipSubnetId = subnet1.Apply(subnet1 =&gt; subnet1.Id),
    ///         });
    ///         var listener1 = new OpenStack.LoadBalancer.Listener("listener1", new OpenStack.LoadBalancer.ListenerArgs
    ///         {
    ///             DefaultTlsContainerRef = tls1.ContainerRef,
    ///             LoadbalancerId = lb1.Id,
    ///             Protocol = "TERMINATED_HTTPS",
    ///             ProtocolPort = 443,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Container with the ACL
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var tls1 = new OpenStack.KeyManager.ContainerV1("tls1", new OpenStack.KeyManager.ContainerV1Args
    ///         {
    ///             Acl = new OpenStack.KeyManager.Inputs.ContainerV1AclArgs
    ///             {
    ///                 Read = new OpenStack.KeyManager.Inputs.ContainerV1AclReadArgs
    ///                 {
    ///                     ProjectAccess = false,
    ///                     Users = 
    ///                     {
    ///                         "userid1",
    ///                         "userid2",
    ///                     },
    ///                 },
    ///             },
    ///             SecretRefs = 
    ///             {
    ///                 new OpenStack.KeyManager.Inputs.ContainerV1SecretRefArgs
    ///                 {
    ///                     Name = "certificate",
    ///                     SecretRef = openstack_keymanager_secret_v1.Certificate_1.Secret_ref,
    ///                 },
    ///                 new OpenStack.KeyManager.Inputs.ContainerV1SecretRefArgs
    ///                 {
    ///                     Name = "private_key",
    ///                     SecretRef = openstack_keymanager_secret_v1.Private_key_1.Secret_ref,
    ///                 },
    ///                 new OpenStack.KeyManager.Inputs.ContainerV1SecretRefArgs
    ///                 {
    ///                     Name = "intermediates",
    ///                     SecretRef = openstack_keymanager_secret_v1.Intermediate_1.Secret_ref,
    ///                 },
    ///             },
    ///             Type = "certificate",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ContainerV1 : Pulumi.CustomResource
    {
        /// <summary>
        /// Allows to control an access to a container. Currently only
        /// the `read` operation is supported. If not specified, the container is
        /// accessible project wide. The `read` structure is described below.
        /// </summary>
        [Output("acl")]
        public Output<Outputs.ContainerV1Acl> Acl { get; private set; } = null!;

        /// <summary>
        /// The list of the container consumers. The structure is described below.
        /// </summary>
        [Output("consumers")]
        public Output<ImmutableArray<Outputs.ContainerV1Consumer>> Consumers { get; private set; } = null!;

        /// <summary>
        /// The container reference / where to find the container.
        /// </summary>
        [Output("containerRef")]
        public Output<string> ContainerRef { get; private set; } = null!;

        /// <summary>
        /// The date the container ACL was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The creator of the container.
        /// </summary>
        [Output("creatorId")]
        public Output<string> CreatorId { get; private set; } = null!;

        /// <summary>
        /// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a container. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 container.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A set of dictionaries containing references to secrets. The structure is described
        /// below.
        /// </summary>
        [Output("secretRefs")]
        public Output<ImmutableArray<Outputs.ContainerV1SecretRef>> SecretRefs { get; private set; } = null!;

        /// <summary>
        /// The status of the container.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The date the container ACL was last updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerV1(string name, ContainerV1Args args, CustomResourceOptions? options = null)
            : base("openstack:keymanager/containerV1:ContainerV1", name, args ?? new ContainerV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerV1(string name, Input<string> id, ContainerV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:keymanager/containerV1:ContainerV1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerV1 Get(string name, Input<string> id, ContainerV1State? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerV1(name, id, state, options);
        }
    }

    public sealed class ContainerV1Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows to control an access to a container. Currently only
        /// the `read` operation is supported. If not specified, the container is
        /// accessible project wide. The `read` structure is described below.
        /// </summary>
        [Input("acl")]
        public Input<Inputs.ContainerV1AclArgs>? Acl { get; set; }

        /// <summary>
        /// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a container. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 container.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretRefs")]
        private InputList<Inputs.ContainerV1SecretRefArgs>? _secretRefs;

        /// <summary>
        /// A set of dictionaries containing references to secrets. The structure is described
        /// below.
        /// </summary>
        public InputList<Inputs.ContainerV1SecretRefArgs> SecretRefs
        {
            get => _secretRefs ?? (_secretRefs = new InputList<Inputs.ContainerV1SecretRefArgs>());
            set => _secretRefs = value;
        }

        /// <summary>
        /// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ContainerV1Args()
        {
        }
    }

    public sealed class ContainerV1State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows to control an access to a container. Currently only
        /// the `read` operation is supported. If not specified, the container is
        /// accessible project wide. The `read` structure is described below.
        /// </summary>
        [Input("acl")]
        public Input<Inputs.ContainerV1AclGetArgs>? Acl { get; set; }

        [Input("consumers")]
        private InputList<Inputs.ContainerV1ConsumerGetArgs>? _consumers;

        /// <summary>
        /// The list of the container consumers. The structure is described below.
        /// </summary>
        public InputList<Inputs.ContainerV1ConsumerGetArgs> Consumers
        {
            get => _consumers ?? (_consumers = new InputList<Inputs.ContainerV1ConsumerGetArgs>());
            set => _consumers = value;
        }

        /// <summary>
        /// The container reference / where to find the container.
        /// </summary>
        [Input("containerRef")]
        public Input<string>? ContainerRef { get; set; }

        /// <summary>
        /// The date the container ACL was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The creator of the container.
        /// </summary>
        [Input("creatorId")]
        public Input<string>? CreatorId { get; set; }

        /// <summary>
        /// The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a container. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 container.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretRefs")]
        private InputList<Inputs.ContainerV1SecretRefGetArgs>? _secretRefs;

        /// <summary>
        /// A set of dictionaries containing references to secrets. The structure is described
        /// below.
        /// </summary>
        public InputList<Inputs.ContainerV1SecretRefGetArgs> SecretRefs
        {
            get => _secretRefs ?? (_secretRefs = new InputList<Inputs.ContainerV1SecretRefGetArgs>());
            set => _secretRefs = value;
        }

        /// <summary>
        /// The status of the container.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The date the container ACL was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ContainerV1State()
        {
        }
    }
}
