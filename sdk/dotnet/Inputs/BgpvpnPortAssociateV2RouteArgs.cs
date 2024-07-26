// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Inputs
{

    public sealed class BgpvpnPortAssociateV2RouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the BGP VPN to be advertised. Required
        /// if `type` is `bgpvpn`. Conflicts with `prefix`.
        /// </summary>
        [Input("bgpvpnId")]
        public Input<string>? BgpvpnId { get; set; }

        /// <summary>
        /// The BGP LOCAL\_PREF value of the routes that will
        /// be advertised.
        /// </summary>
        [Input("localPref")]
        public Input<int>? LocalPref { get; set; }

        /// <summary>
        /// The CIDR prefix (v4 or v6) to be advertised. Required
        /// if `type` is `prefix`. Conflicts with `bgpvpn_id`.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Can be `prefix` or `bgpvpn`. For the `prefix` type, the
        /// CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
        /// `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BgpvpnPortAssociateV2RouteArgs()
        {
        }
        public static new BgpvpnPortAssociateV2RouteArgs Empty => new BgpvpnPortAssociateV2RouteArgs();
    }
}