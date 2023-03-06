// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Outputs
{

    [OutputType]
    public sealed class ContainerV1Consumer
    {
        /// <summary>
        /// Human-readable name for the Container. Does not have
        /// to be unique.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The consumer URL.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private ContainerV1Consumer(
            string? name,

            string? url)
        {
            Name = name;
            Url = url;
        }
    }
}
