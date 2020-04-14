// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Database.Outputs
{

    [OutputType]
    public sealed class InstanceUser
    {
        /// <summary>
        /// A list of databases that user will have access to. If not specified, 
        /// user has access to all databases on th einstance. Changing this creates a new instance.
        /// </summary>
        public readonly ImmutableArray<string> Databases;
        /// <summary>
        /// An ip address or % sign indicating what ip addresses can connect with
        /// this user credentials. Changing this creates a new instance.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// Database to be created on new instance. Changing this creates a
        /// new instance.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User's password. Changing this creates a
        /// new instance.
        /// </summary>
        public readonly string? Password;

        [OutputConstructor]
        private InstanceUser(
            ImmutableArray<string> databases,

            string? host,

            string name,

            string? password)
        {
            Databases = databases;
            Host = host;
            Name = name;
            Password = password;
        }
    }
}
