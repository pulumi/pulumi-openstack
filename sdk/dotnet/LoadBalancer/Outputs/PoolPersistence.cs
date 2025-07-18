// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer.Outputs
{

    [OutputType]
    public sealed class PoolPersistence
    {
        /// <summary>
        /// The name of the cookie if persistence mode is set
        /// appropriately. Required if `type = APP_COOKIE`.
        /// </summary>
        public readonly string? CookieName;
        /// <summary>
        /// The type of persistence mode. The current specification
        /// supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PoolPersistence(
            string? cookieName,

            string type)
        {
            CookieName = cookieName;
            Type = type;
        }
    }
}
