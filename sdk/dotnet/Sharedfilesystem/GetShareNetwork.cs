// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.SharedFileSystem
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to get the ID of an available Shared File System share network.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/sharedfilesystem_sharenetwork_v2.html.markdown.
        /// </summary>
        public static Task<GetShareNetworkResult> GetShareNetwork(GetShareNetworkArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetShareNetworkResult>("openstack:sharedfilesystem/getShareNetwork:getShareNetwork", args ?? ResourceArgs.Empty, options.WithVersion());
    }

    public sealed class GetShareNetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description of the share network.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IP version of the share network. Can either be 4 or 6.
        /// </summary>
        [Input("ipVersion")]
        public Input<int>? IpVersion { get; set; }

        /// <summary>
        /// The name of the share network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The share network type. Can either be VLAN, VXLAN,
        /// GRE, or flat.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The neutron network UUID of the share network.
        /// </summary>
        [Input("neutronNetId")]
        public Input<string>? NeutronNetId { get; set; }

        /// <summary>
        /// The neutron subnet UUID of the share network.
        /// </summary>
        [Input("neutronSubnetId")]
        public Input<string>? NeutronSubnetId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to read a share network. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The security service IDs associated with
        /// the share network.
        /// </summary>
        [Input("securityServiceId")]
        public Input<string>? SecurityServiceId { get; set; }

        /// <summary>
        /// The share network segmentation ID.
        /// </summary>
        [Input("segmentationId")]
        public Input<int>? SegmentationId { get; set; }

        public GetShareNetworkArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetShareNetworkResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Cidr;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int IpVersion;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string NeutronNetId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string NeutronSubnetId;
        /// <summary>
        /// The owner of the Share Network.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? SecurityServiceId;
        /// <summary>
        /// The list of security service IDs associated with
        /// the share network.
        /// </summary>
        public readonly ImmutableArray<string> SecurityServiceIds;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly int SegmentationId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetShareNetworkResult(
            string cidr,
            string description,
            int ipVersion,
            string name,
            string networkType,
            string neutronNetId,
            string neutronSubnetId,
            string projectId,
            string region,
            string? securityServiceId,
            ImmutableArray<string> securityServiceIds,
            int segmentationId,
            string id)
        {
            Cidr = cidr;
            Description = description;
            IpVersion = ipVersion;
            Name = name;
            NetworkType = networkType;
            NeutronNetId = neutronNetId;
            NeutronSubnetId = neutronSubnetId;
            ProjectId = projectId;
            Region = region;
            SecurityServiceId = securityServiceId;
            SecurityServiceIds = securityServiceIds;
            SegmentationId = segmentationId;
            Id = id;
        }
    }
}
