// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.LoadBalancer.Inputs
{

    public sealed class PoolPersistenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cookie if persistence mode is set
        /// appropriately. Required if `type = APP_COOKIE`.
        /// </summary>
        [Input("cookieName")]
        public Input<string>? CookieName { get; set; }

        /// <summary>
        /// The type of persistence mode. The current specification
        /// supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public PoolPersistenceArgs()
        {
        }
        public static new PoolPersistenceArgs Empty => new PoolPersistenceArgs();
    }
}
