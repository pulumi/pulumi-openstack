// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Inputs
{

    public sealed class ContainerV1ConsumerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human-readable name for the Container. Does not have
        /// to be unique.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The consumer URL.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ContainerV1ConsumerArgs()
        {
        }
        public static new ContainerV1ConsumerArgs Empty => new ContainerV1ConsumerArgs();
    }
}