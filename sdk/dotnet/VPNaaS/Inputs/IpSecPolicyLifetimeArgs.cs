// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.VPNaaS.Inputs
{

    public sealed class IpSecPolicyLifetimeArgs : Pulumi.ResourceArgs
    {
        [Input("units")]
        public Input<string>? Units { get; set; }

        [Input("value")]
        public Input<int>? Value { get; set; }

        public IpSecPolicyLifetimeArgs()
        {
        }
    }
}
