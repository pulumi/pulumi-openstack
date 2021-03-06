// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking.Outputs
{

    [OutputType]
    public sealed class SubnetAllocationPoolsCollection
    {
        /// <summary>
        /// The ending address.
        /// </summary>
        public readonly string End;
        /// <summary>
        /// The starting address.
        /// </summary>
        public readonly string Start;

        [OutputConstructor]
        private SubnetAllocationPoolsCollection(
            string end,

            string start)
        {
            End = end;
            Start = start;
        }
    }
}
