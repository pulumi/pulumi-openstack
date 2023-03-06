// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.VPNaaS.Inputs
{

    public sealed class IkePolicyLifetimeArgs : global::Pulumi.ResourceArgs
    {
        [Input("units")]
        public Input<string>? Units { get; set; }

        /// <summary>
        /// The value for the lifetime of the security association. Must be a positive integer.
        /// Default is 3600.
        /// </summary>
        [Input("value")]
        public Input<int>? Value { get; set; }

        public IkePolicyLifetimeArgs()
        {
        }
        public static new IkePolicyLifetimeArgs Empty => new IkePolicyLifetimeArgs();
    }
}
