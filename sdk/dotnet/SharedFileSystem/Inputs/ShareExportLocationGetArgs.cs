// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.SharedFileSystem.Inputs
{

    public sealed class ShareExportLocationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("preferred")]
        public Input<string>? Preferred { get; set; }

        public ShareExportLocationGetArgs()
        {
        }
        public static new ShareExportLocationGetArgs Empty => new ShareExportLocationGetArgs();
    }
}
