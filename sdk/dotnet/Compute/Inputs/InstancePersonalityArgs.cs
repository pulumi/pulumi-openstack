// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class InstancePersonalityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The contents of the file. Limited to 255 bytes.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The absolute path of the destination file.
        /// </summary>
        [Input("file", required: true)]
        public Input<string> File { get; set; } = null!;

        public InstancePersonalityArgs()
        {
        }
        public static new InstancePersonalityArgs Empty => new InstancePersonalityArgs();
    }
}
