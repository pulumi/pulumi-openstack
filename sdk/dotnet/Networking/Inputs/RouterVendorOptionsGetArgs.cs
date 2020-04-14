// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Inputs
{

    public sealed class RouterVendorOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean to control whether
        /// the Router gateway is assigned during creation or updated after creation.
        /// </summary>
        [Input("setRouterGatewayAfterCreate")]
        public Input<bool>? SetRouterGatewayAfterCreate { get; set; }

        public RouterVendorOptionsGetArgs()
        {
        }
    }
}
