// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Inputs
{

    public sealed class OrderV1MetaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Algorithm to use for key generation.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// - Bit lenght of key to be generated.
        /// </summary>
        [Input("bitLength", required: true)]
        public Input<int> BitLength { get; set; } = null!;

        /// <summary>
        /// This is a UTC timestamp in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ. If set, the secret will not be available after this time.
        /// </summary>
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        /// <summary>
        /// The mode to use for key generation.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name of the secret set by the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The media type for the content of the secrets payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        /// </summary>
        [Input("payloadContentType")]
        public Input<string>? PayloadContentType { get; set; }

        public OrderV1MetaArgs()
        {
        }
    }
}