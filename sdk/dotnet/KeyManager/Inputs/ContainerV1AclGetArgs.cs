// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Inputs
{

    public sealed class ContainerV1AclGetArgs : Pulumi.ResourceArgs
    {
        [Input("read")]
        public Input<Inputs.ContainerV1AclReadGetArgs>? Read { get; set; }

        public ContainerV1AclGetArgs()
        {
        }
    }
}