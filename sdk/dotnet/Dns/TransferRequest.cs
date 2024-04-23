// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Dns
{
    /// <summary>
    /// Manages a DNS zone transfer request in the OpenStack DNS Service.
    /// 
    /// ## Example Usage
    /// 
    /// ### Automatically detect the correct network
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleZone = new OpenStack.Dns.Zone("example_zone", new()
    ///     {
    ///         Name = "example.com.",
    ///         Email = "jdoe@example.com",
    ///         Description = "An example zone",
    ///         Ttl = 3000,
    ///         Type = "PRIMARY",
    ///     });
    /// 
    ///     var request1 = new OpenStack.Dns.TransferRequest("request_1", new()
    ///     {
    ///         ZoneId = exampleZone.Id,
    ///         Description = "a transfer request",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying the transferRequest ID:
    /// 
    /// ```sh
    /// $ pulumi import openstack:dns/transferRequest:TransferRequest request_1 request_id
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:dns/transferRequest:TransferRequest")]
    public partial class TransferRequest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description of the zone tranfer request.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Disable wait for zone to reach ACTIVE
        /// status. The check is enabled by default. If this argument is true, zone
        /// will be considered as created/updated if OpenStack request returned success.
        /// </summary>
        [Output("disableStatusCheck")]
        public Output<bool?> DisableStatusCheck { get; private set; } = null!;

        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The target Project ID to transfer to.
        /// </summary>
        [Output("targetProjectId")]
        public Output<string> TargetProjectId { get; private set; } = null!;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new transfer request.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;

        /// <summary>
        /// The ID of the zone for which to create the transfer
        /// request.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a TransferRequest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransferRequest(string name, TransferRequestArgs args, CustomResourceOptions? options = null)
            : base("openstack:dns/transferRequest:TransferRequest", name, args ?? new TransferRequestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransferRequest(string name, Input<string> id, TransferRequestState? state = null, CustomResourceOptions? options = null)
            : base("openstack:dns/transferRequest:TransferRequest", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransferRequest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransferRequest Get(string name, Input<string> id, TransferRequestState? state = null, CustomResourceOptions? options = null)
        {
            return new TransferRequest(name, id, state, options);
        }
    }

    public sealed class TransferRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the zone tranfer request.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disable wait for zone to reach ACTIVE
        /// status. The check is enabled by default. If this argument is true, zone
        /// will be considered as created/updated if OpenStack request returned success.
        /// </summary>
        [Input("disableStatusCheck")]
        public Input<bool>? DisableStatusCheck { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The target Project ID to transfer to.
        /// </summary>
        [Input("targetProjectId")]
        public Input<string>? TargetProjectId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new transfer request.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the zone for which to create the transfer
        /// request.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public TransferRequestArgs()
        {
        }
        public static new TransferRequestArgs Empty => new TransferRequestArgs();
    }

    public sealed class TransferRequestState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the zone tranfer request.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disable wait for zone to reach ACTIVE
        /// status. The check is enabled by default. If this argument is true, zone
        /// will be considered as created/updated if OpenStack request returned success.
        /// </summary>
        [Input("disableStatusCheck")]
        public Input<bool>? DisableStatusCheck { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The target Project ID to transfer to.
        /// </summary>
        [Input("targetProjectId")]
        public Input<string>? TargetProjectId { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new transfer request.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the zone for which to create the transfer
        /// request.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public TransferRequestState()
        {
        }
        public static new TransferRequestState Empty => new TransferRequestState();
    }
}
