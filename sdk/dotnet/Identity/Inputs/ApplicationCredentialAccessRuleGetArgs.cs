// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity.Inputs
{

    public sealed class ApplicationCredentialAccessRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the existing access rule. The access rule ID of
        /// another application credential can be provided.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The request method that the application credential is
        /// permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
        /// `HEAD`, `PATCH`, `PUT` and `DELETE`.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        /// <summary>
        /// The API path that the application credential is permitted
        /// to access. May use named wildcards such as **{tag}** or the unnamed wildcard
        /// **\*** to match against any string in the path up to a **/**, or the recursive
        /// wildcard **\*\*** to include **/** in the matched path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The service type identifier for the service that the
        /// application credential is granted to access. Must be a service type that is
        /// listed in the service catalog and not a code name for a service. E.g.
        /// **identity**, **compute**, **volumev3**, **image**, **network**,
        /// **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public ApplicationCredentialAccessRuleGetArgs()
        {
        }
        public static new ApplicationCredentialAccessRuleGetArgs Empty => new ApplicationCredentialAccessRuleGetArgs();
    }
}
