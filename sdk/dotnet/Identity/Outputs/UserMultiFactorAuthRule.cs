// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity.Outputs
{

    [OutputType]
    public sealed class UserMultiFactorAuthRule
    {
        /// <summary>
        /// A list of authentication plugins that the user must
        /// authenticate with.
        /// </summary>
        public readonly ImmutableArray<string> Rules;

        [OutputConstructor]
        private UserMultiFactorAuthRule(ImmutableArray<string> rules)
        {
            Rules = rules;
        }
    }
}