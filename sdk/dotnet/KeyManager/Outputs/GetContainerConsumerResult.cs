// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Outputs
{

    [OutputType]
    public sealed class GetContainerConsumerResult
    {
        /// <summary>
        /// The Container name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The consumer URL.
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private GetContainerConsumerResult(
            string? name,

            string? url)
        {
            Name = name;
            Url = url;
        }
    }
}
