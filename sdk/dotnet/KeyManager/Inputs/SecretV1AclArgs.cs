// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Inputs
{

    public sealed class SecretV1AclArgs : Pulumi.ResourceArgs
    {
        [Input("read")]
        public Input<Inputs.SecretV1AclReadArgs>? Read { get; set; }

        public SecretV1AclArgs()
        {
        }
    }
}
