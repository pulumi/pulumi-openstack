// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Dns
{
    /// <summary>
    /// Manages a DNS zone transfer accept in the OpenStack DNS Service.
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
    ///         Description = "a transfer accept",
    ///     });
    /// 
    ///     var accept1 = new OpenStack.Dns.TransferAccept("accept_1", new()
    ///     {
    ///         ZoneTransferRequestId = request1.Id,
    ///         Key = request1.Key,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying the transferAccept ID:
    /// 
    /// ```sh
    /// $ pulumi import openstack:dns/transferAccept:TransferAccept accept_1 accept_id
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:dns/transferAccept:TransferAccept")]
    public partial class TransferAccept : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Disable wait for zone to reach ACTIVE
        /// status. The check is enabled by default. If this argument is true, zone
        /// will be considered as created/updated if OpenStack accept returned success.
        /// </summary>
        [Output("disableStatusCheck")]
        public Output<bool?> DisableStatusCheck { get; private set; } = null!;

        /// <summary>
        /// The transfer key.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone zone transfer accept.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new transfer accept.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, string>?> ValueSpecs { get; private set; } = null!;

        /// <summary>
        /// The ID of the zone transfer request.
        /// </summary>
        [Output("zoneTransferRequestId")]
        public Output<string> ZoneTransferRequestId { get; private set; } = null!;


        /// <summary>
        /// Create a TransferAccept resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TransferAccept(string name, TransferAcceptArgs args, CustomResourceOptions? options = null)
            : base("openstack:dns/transferAccept:TransferAccept", name, args ?? new TransferAcceptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TransferAccept(string name, Input<string> id, TransferAcceptState? state = null, CustomResourceOptions? options = null)
            : base("openstack:dns/transferAccept:TransferAccept", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TransferAccept resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TransferAccept Get(string name, Input<string> id, TransferAcceptState? state = null, CustomResourceOptions? options = null)
        {
            return new TransferAccept(name, id, state, options);
        }
    }

    public sealed class TransferAcceptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable wait for zone to reach ACTIVE
        /// status. The check is enabled by default. If this argument is true, zone
        /// will be considered as created/updated if OpenStack accept returned success.
        /// </summary>
        [Input("disableStatusCheck")]
        public Input<bool>? DisableStatusCheck { get; set; }

        /// <summary>
        /// The transfer key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone zone transfer accept.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("valueSpecs")]
        private InputMap<string>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new transfer accept.
        /// </summary>
        public InputMap<string> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<string>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the zone transfer request.
        /// </summary>
        [Input("zoneTransferRequestId", required: true)]
        public Input<string> ZoneTransferRequestId { get; set; } = null!;

        public TransferAcceptArgs()
        {
        }
        public static new TransferAcceptArgs Empty => new TransferAcceptArgs();
    }

    public sealed class TransferAcceptState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable wait for zone to reach ACTIVE
        /// status. The check is enabled by default. If this argument is true, zone
        /// will be considered as created/updated if OpenStack accept returned success.
        /// </summary>
        [Input("disableStatusCheck")]
        public Input<bool>? DisableStatusCheck { get; set; }

        /// <summary>
        /// The transfer key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone zone transfer accept.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("valueSpecs")]
        private InputMap<string>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new transfer accept.
        /// </summary>
        public InputMap<string> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<string>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the zone transfer request.
        /// </summary>
        [Input("zoneTransferRequestId")]
        public Input<string>? ZoneTransferRequestId { get; set; }

        public TransferAcceptState()
        {
        }
        public static new TransferAcceptState Empty => new TransferAcceptState();
    }
}
