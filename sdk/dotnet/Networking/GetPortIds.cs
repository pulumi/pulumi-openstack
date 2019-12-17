// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get a list of Openstack Port IDs matching the
        /// specified criteria.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_port_ids_v2.html.markdown.
        /// </summary>
        public static Task<GetPortIdsResult> GetPortIds(GetPortIdsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPortIdsResult>("openstack:networking/getPortIds:getPortIds", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetPortIdsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The administrative state of the port.
        /// </summary>
        [Input("adminStateUp")]
        public Input<bool>? AdminStateUp { get; set; }

        /// <summary>
        /// Human-readable description of the port.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the device the port belongs to.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// The device owner of the port.
        /// </summary>
        [Input("deviceOwner")]
        public Input<string>? DeviceOwner { get; set; }

        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// The port IP address filter.
        /// </summary>
        [Input("fixedIp")]
        public Input<string>? FixedIp { get; set; }

        /// <summary>
        /// The MAC address of the port.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The name of the port.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the network the port belongs to.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The owner of the port.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve port ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The list of port security group IDs to filter.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Order the results in either `asc` or `desc`.
        /// Defaults to none.
        /// </summary>
        [Input("sortDirection")]
        public Input<string>? SortDirection { get; set; }

        /// <summary>
        /// Sort ports based on a certain key. Defaults to none.
        /// </summary>
        [Input("sortKey")]
        public Input<string>? SortKey { get; set; }

        /// <summary>
        /// The status of the port.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The list of port tags to filter.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public GetPortIdsArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetPortIdsResult
    {
        public readonly bool? AdminStateUp;
        public readonly string? Description;
        public readonly string? DeviceId;
        public readonly string? DeviceOwner;
        public readonly string? DnsName;
        public readonly string? FixedIp;
        public readonly ImmutableArray<string> Ids;
        public readonly string? MacAddress;
        public readonly string? Name;
        public readonly string? NetworkId;
        public readonly string? ProjectId;
        public readonly string? Region;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly string? SortDirection;
        public readonly string? SortKey;
        public readonly string? Status;
        public readonly ImmutableArray<string> Tags;
        public readonly string? TenantId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetPortIdsResult(
            bool? adminStateUp,
            string? description,
            string? deviceId,
            string? deviceOwner,
            string? dnsName,
            string? fixedIp,
            ImmutableArray<string> ids,
            string? macAddress,
            string? name,
            string? networkId,
            string? projectId,
            string? region,
            ImmutableArray<string> securityGroupIds,
            string? sortDirection,
            string? sortKey,
            string? status,
            ImmutableArray<string> tags,
            string? tenantId,
            string id)
        {
            AdminStateUp = adminStateUp;
            Description = description;
            DeviceId = deviceId;
            DeviceOwner = deviceOwner;
            DnsName = dnsName;
            FixedIp = fixedIp;
            Ids = ids;
            MacAddress = macAddress;
            Name = name;
            NetworkId = networkId;
            ProjectId = projectId;
            Region = region;
            SecurityGroupIds = securityGroupIds;
            SortDirection = sortDirection;
            SortKey = sortKey;
            Status = status;
            Tags = tags;
            TenantId = tenantId;
            Id = id;
        }
    }
}