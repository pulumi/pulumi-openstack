// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Inputs
{

    public sealed class ContainerV1AclArgs : global::Pulumi.ResourceArgs
    {
        [Input("read")]
        public Input<Inputs.ContainerV1AclReadArgs>? Read { get; set; }

        public ContainerV1AclArgs()
        {
        }
        public static new ContainerV1AclArgs Empty => new ContainerV1AclArgs();
    }
}
