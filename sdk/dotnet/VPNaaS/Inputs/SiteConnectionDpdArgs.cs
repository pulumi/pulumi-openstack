// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.VPNaaS.Inputs
{

    public sealed class SiteConnectionDpdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dead peer detection (DPD) action.
        /// A valid value is clear, hold, restart, disabled, or restart-by-peer.
        /// Default value is hold.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The dead peer detection (DPD) interval, in seconds.
        /// A valid value is a positive integer.
        /// Default is 30.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The dead peer detection (DPD) timeout in seconds.
        /// A valid value is a positive integer that is greater than the DPD interval value.
        /// Default is 120.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public SiteConnectionDpdArgs()
        {
        }
        public static new SiteConnectionDpdArgs Empty => new SiteConnectionDpdArgs();
    }
}
