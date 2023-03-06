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
    /// Manages a DNS record set in the OpenStack DNS Service.
    /// 
    /// ## Example Usage
    /// ### Automatically detect the correct network
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleZone = new OpenStack.Dns.Zone("exampleZone", new()
    ///     {
    ///         Description = "a zone",
    ///         Email = "email2@example.com",
    ///         Ttl = 6000,
    ///         Type = "PRIMARY",
    ///     });
    /// 
    ///     var rsExampleCom = new OpenStack.Dns.RecordSet("rsExampleCom", new()
    ///     {
    ///         Description = "An example record set",
    ///         Records = new[]
    ///         {
    ///             "10.0.0.1",
    ///         },
    ///         Ttl = 3000,
    ///         Type = "A",
    ///         ZoneId = exampleZone.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:dns/recordSet:RecordSet recordset_1 &lt;zone_id&gt;/&lt;recordset_id&gt;
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:dns/recordSet:RecordSet")]
    public partial class RecordSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description of the  record set.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Disable wait for recordset to reach ACTIVE
        /// status. This argumen is disabled by default. If it is set to true, the recordset
        /// will be considered as created/updated/deleted if OpenStack request returned success.
        /// </summary>
        [Output("disableStatusCheck")]
        public Output<bool?> DisableStatusCheck { get; private set; } = null!;

        /// <summary>
        /// The name of the record set. Note the `.` at the end of the name.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project DNS zone is created
        /// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
        /// user role in target project)
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// An array of DNS records. _Note:_ if an IPv6 address
        /// contains brackets (`[ ]`), the brackets will be stripped and the modified
        /// address will be recorded in the state.
        /// </summary>
        [Output("records")]
        public Output<ImmutableArray<string>> Records { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The time to live (TTL) of the record set.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of record set. Examples: "A", "MX".
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new record set.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;

        /// <summary>
        /// The ID of the zone in which to create the record set.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a RecordSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecordSet(string name, RecordSetArgs args, CustomResourceOptions? options = null)
            : base("openstack:dns/recordSet:RecordSet", name, args ?? new RecordSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecordSet(string name, Input<string> id, RecordSetState? state = null, CustomResourceOptions? options = null)
            : base("openstack:dns/recordSet:RecordSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RecordSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecordSet Get(string name, Input<string> id, RecordSetState? state = null, CustomResourceOptions? options = null)
        {
            return new RecordSet(name, id, state, options);
        }
    }

    public sealed class RecordSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the  record set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disable wait for recordset to reach ACTIVE
        /// status. This argumen is disabled by default. If it is set to true, the recordset
        /// will be considered as created/updated/deleted if OpenStack request returned success.
        /// </summary>
        [Input("disableStatusCheck")]
        public Input<bool>? DisableStatusCheck { get; set; }

        /// <summary>
        /// The name of the record set. Note the `.` at the end of the name.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project DNS zone is created
        /// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
        /// user role in target project)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("records")]
        private InputList<string>? _records;

        /// <summary>
        /// An array of DNS records. _Note:_ if an IPv6 address
        /// contains brackets (`[ ]`), the brackets will be stripped and the modified
        /// address will be recorded in the state.
        /// </summary>
        public InputList<string> Records
        {
            get => _records ?? (_records = new InputList<string>());
            set => _records = value;
        }

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The time to live (TTL) of the record set.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of record set. Examples: "A", "MX".
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new record set.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the zone in which to create the record set.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public RecordSetArgs()
        {
        }
        public static new RecordSetArgs Empty => new RecordSetArgs();
    }

    public sealed class RecordSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the  record set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Disable wait for recordset to reach ACTIVE
        /// status. This argumen is disabled by default. If it is set to true, the recordset
        /// will be considered as created/updated/deleted if OpenStack request returned success.
        /// </summary>
        [Input("disableStatusCheck")]
        public Input<bool>? DisableStatusCheck { get; set; }

        /// <summary>
        /// The name of the record set. Note the `.` at the end of the name.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project DNS zone is created
        /// for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
        /// user role in target project)
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("records")]
        private InputList<string>? _records;

        /// <summary>
        /// An array of DNS records. _Note:_ if an IPv6 address
        /// contains brackets (`[ ]`), the brackets will be stripped and the modified
        /// address will be recorded in the state.
        /// </summary>
        public InputList<string> Records
        {
            get => _records ?? (_records = new InputList<string>());
            set => _records = value;
        }

        /// <summary>
        /// The region in which to obtain the V2 DNS client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The time to live (TTL) of the record set.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of record set. Examples: "A", "MX".
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new record set.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        /// <summary>
        /// The ID of the zone in which to create the record set.
        /// Changing this creates a new DNS  record set.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public RecordSetState()
        {
        }
        public static new RecordSetState Empty => new RecordSetState();
    }
}
