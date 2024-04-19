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
    /// Manages a V1 Barbican order resource within OpenStack.
    /// 
    /// ## Example Usage
    /// 
    /// ### Symmetric key order
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var order1 = new OpenStack.KeyManager.OrderV1("order1", new()
    ///     {
    ///         Meta = new OpenStack.KeyManager.Inputs.OrderV1MetaArgs
    ///         {
    ///             Algorithm = "aes",
    ///             BitLength = 256,
    ///             Mode = "cbc",
    ///             Name = "mysecret",
    ///         },
    ///         Type = "key",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Asymmetric key pair order
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var order1 = new OpenStack.KeyManager.OrderV1("order1", new()
    ///     {
    ///         Meta = new OpenStack.KeyManager.Inputs.OrderV1MetaArgs
    ///         {
    ///             Algorithm = "rsa",
    ///             BitLength = 4096,
    ///             Name = "mysecret",
    ///         },
    ///         Type = "asymmetric",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Orders can be imported using the order id (the last part of the order reference), e.g.:
    /// 
    /// ```sh
    /// $ pulumi import openstack:keymanager/orderV1:OrderV1 order_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:keymanager/orderV1:OrderV1")]
    public partial class OrderV1 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The container reference / where to find the container.
        /// </summary>
        [Output("containerRef")]
        public Output<string> ContainerRef { get; private set; } = null!;

        /// <summary>
        /// The date the order was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// The creator of the order.
        /// </summary>
        [Output("creatorId")]
        public Output<string> CreatorId { get; private set; } = null!;

        /// <summary>
        /// Dictionary containing the order metadata used to generate the order. The structure is described below.
        /// </summary>
        [Output("meta")]
        public Output<Outputs.OrderV1Meta> Meta { get; private set; } = null!;

        /// <summary>
        /// The order reference / where to find the order.
        /// </summary>
        [Output("orderRef")]
        public Output<string> OrderRef { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a order. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 order.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The secret reference / where to find the secret.
        /// </summary>
        [Output("secretRef")]
        public Output<string> SecretRef { get; private set; } = null!;

        /// <summary>
        /// The status of the order.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The sub status of the order.
        /// </summary>
        [Output("subStatus")]
        public Output<string> SubStatus { get; private set; } = null!;

        /// <summary>
        /// The sub status message of the order.
        /// </summary>
        [Output("subStatusMessage")]
        public Output<string> SubStatusMessage { get; private set; } = null!;

        /// <summary>
        /// The type of key to be generated. Must be one of `asymmetric`, `key`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The date the order was last updated.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;


        /// <summary>
        /// Create a OrderV1 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrderV1(string name, OrderV1Args args, CustomResourceOptions? options = null)
            : base("openstack:keymanager/orderV1:OrderV1", name, args ?? new OrderV1Args(), MakeResourceOptions(options, ""))
        {
        }

        private OrderV1(string name, Input<string> id, OrderV1State? state = null, CustomResourceOptions? options = null)
            : base("openstack:keymanager/orderV1:OrderV1", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrderV1 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrderV1 Get(string name, Input<string> id, OrderV1State? state = null, CustomResourceOptions? options = null)
        {
            return new OrderV1(name, id, state, options);
        }
    }

    public sealed class OrderV1Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dictionary containing the order metadata used to generate the order. The structure is described below.
        /// </summary>
        [Input("meta", required: true)]
        public Input<Inputs.OrderV1MetaArgs> Meta { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a order. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 order.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The type of key to be generated. Must be one of `asymmetric`, `key`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public OrderV1Args()
        {
        }
        public static new OrderV1Args Empty => new OrderV1Args();
    }

    public sealed class OrderV1State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The container reference / where to find the container.
        /// </summary>
        [Input("containerRef")]
        public Input<string>? ContainerRef { get; set; }

        /// <summary>
        /// The date the order was created.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// The creator of the order.
        /// </summary>
        [Input("creatorId")]
        public Input<string>? CreatorId { get; set; }

        /// <summary>
        /// Dictionary containing the order metadata used to generate the order. The structure is described below.
        /// </summary>
        [Input("meta")]
        public Input<Inputs.OrderV1MetaGetArgs>? Meta { get; set; }

        /// <summary>
        /// The order reference / where to find the order.
        /// </summary>
        [Input("orderRef")]
        public Input<string>? OrderRef { get; set; }

        /// <summary>
        /// The region in which to obtain the V1 KeyManager client.
        /// A KeyManager client is needed to create a order. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// V1 order.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The secret reference / where to find the secret.
        /// </summary>
        [Input("secretRef")]
        public Input<string>? SecretRef { get; set; }

        /// <summary>
        /// The status of the order.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The sub status of the order.
        /// </summary>
        [Input("subStatus")]
        public Input<string>? SubStatus { get; set; }

        /// <summary>
        /// The sub status message of the order.
        /// </summary>
        [Input("subStatusMessage")]
        public Input<string>? SubStatusMessage { get; set; }

        /// <summary>
        /// The type of key to be generated. Must be one of `asymmetric`, `key`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The date the order was last updated.
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        public OrderV1State()
        {
        }
        public static new OrderV1State Empty => new OrderV1State();
    }
}
