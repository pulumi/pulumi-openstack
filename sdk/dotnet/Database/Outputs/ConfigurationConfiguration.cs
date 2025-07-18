// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Database.Outputs
{

    [OutputType]
    public sealed class ConfigurationConfiguration
    {
        /// <summary>
        /// Configuration parameter name. Changing this creates a new resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        /// </summary>
        public readonly bool? StringType;
        /// <summary>
        /// Configuration parameter value. Changing this creates a new resource.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ConfigurationConfiguration(
            string name,

            bool? stringType,

            string value)
        {
            Name = name;
            StringType = stringType;
            Value = value;
        }
    }
}
