// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Inputs
{

    public sealed class SubnetAllocationPoolsCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ending address.
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// The starting address.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public SubnetAllocationPoolsCollectionArgs()
        {
        }
        public static new SubnetAllocationPoolsCollectionArgs Empty => new SubnetAllocationPoolsCollectionArgs();
    }
}
