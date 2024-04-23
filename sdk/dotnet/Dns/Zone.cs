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
    /// Manages a DNS zone in the OpenStack DNS Service.
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
    ///     var exampleCom = new OpenStack.Dns.Zone("example_com", new()
    ///     {
    ///         Name = "example.com.",
    ///         Email = "jdoe@example.com",
    ///         Description = "An example zone",
    ///         Ttl = 3000,
    ///         Type = "PRIMARY",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying the zone ID with optional project ID:
    /// 
    /// ```sh
    /// $ pulumi import openstack:dns/zone:Zone zone_1 zone_id
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import openstack:dns/zone:Zone zone_1 zone_id/project_id
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:dns/zone:Zone")]
    public partial class Zone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Attributes for the DNS Service scheduler.
        /// Changing this creates a new zone.
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableDictionary<string, object>?> Attributes { get; private set; } = null!;

        /// <summary>
        /// A description of the zone.
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

        /// <summary>
        /// The email contact for the zone record.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// An array of master DNS servers. For when `type` is
        /// `SECONDARY`.
        /// </summary>
        [Output("masters")]
        public Output<ImmutableArray<string>> Masters { get; private set; } = null!;

        /// <summary>
        /// The name of the zone. Note the `.` at the end of the name.
        /// Changing this creates a new DNS zone.
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
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The time to live (TTL) of the zone.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
        /// Changing this creates a new zone.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new zone.
        /// </summary>
        [Output("valueSpecs")]
        public Output<ImmutableDictionary<string, object>?> ValueSpecs { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:dns/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("openstack:dns/zone:Zone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;

        /// <summary>
        /// Attributes for the DNS Service scheduler.
        /// Changing this creates a new zone.
        /// </summary>
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        /// <summary>
        /// A description of the zone.
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

        /// <summary>
        /// The email contact for the zone record.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("masters")]
        private InputList<string>? _masters;

        /// <summary>
        /// An array of master DNS servers. For when `type` is
        /// `SECONDARY`.
        /// </summary>
        public InputList<string> Masters
        {
            get => _masters ?? (_masters = new InputList<string>());
            set => _masters = value;
        }

        /// <summary>
        /// The name of the zone. Note the `.` at the end of the name.
        /// Changing this creates a new DNS zone.
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

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The time to live (TTL) of the zone.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
        /// Changing this creates a new zone.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new zone.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public ZoneArgs()
        {
        }
        public static new ZoneArgs Empty => new ZoneArgs();
    }

    public sealed class ZoneState : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;

        /// <summary>
        /// Attributes for the DNS Service scheduler.
        /// Changing this creates a new zone.
        /// </summary>
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        /// <summary>
        /// A description of the zone.
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

        /// <summary>
        /// The email contact for the zone record.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("masters")]
        private InputList<string>? _masters;

        /// <summary>
        /// An array of master DNS servers. For when `type` is
        /// `SECONDARY`.
        /// </summary>
        public InputList<string> Masters
        {
            get => _masters ?? (_masters = new InputList<string>());
            set => _masters = value;
        }

        /// <summary>
        /// The name of the zone. Note the `.` at the end of the name.
        /// Changing this creates a new DNS zone.
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

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// Keypairs are associated with accounts, but a Compute client is needed to
        /// create one. If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new DNS zone.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The time to live (TTL) of the zone.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The type of zone. Can either be `PRIMARY` or `SECONDARY`.
        /// Changing this creates a new zone.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("valueSpecs")]
        private InputMap<object>? _valueSpecs;

        /// <summary>
        /// Map of additional options. Changing this creates a
        /// new zone.
        /// </summary>
        public InputMap<object> ValueSpecs
        {
            get => _valueSpecs ?? (_valueSpecs = new InputMap<object>());
            set => _valueSpecs = value;
        }

        public ZoneState()
        {
        }
        public static new ZoneState Empty => new ZoneState();
    }
}
