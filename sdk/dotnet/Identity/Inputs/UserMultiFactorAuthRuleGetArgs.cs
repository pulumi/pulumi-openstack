// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity.Inputs
{

    public sealed class UserMultiFactorAuthRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<string>? _rules;

        /// <summary>
        /// A list of authentication plugins that the user must
        /// authenticate with.
        /// </summary>
        public InputList<string> Rules
        {
            get => _rules ?? (_rules = new InputList<string>());
            set => _rules = value;
        }

        public UserMultiFactorAuthRuleGetArgs()
        {
        }
    }
}
