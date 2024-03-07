// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Database.Inputs
{

    public sealed class InstanceUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("databases")]
        private InputList<string>? _databases;

        /// <summary>
        /// A list of databases that user will have access to. If not specified,
        /// user has access to all databases on th einstance. Changing this creates a new instance.
        /// </summary>
        public InputList<string> Databases
        {
            get => _databases ?? (_databases = new InputList<string>());
            set => _databases = value;
        }

        /// <summary>
        /// An ip address or %!s(MISSING)ign indicating what ip addresses can connect with
        /// this user credentials. Changing this creates a new instance.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Username to be created on new instance. Changing this creates a
        /// new instance.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// User's password. Changing this creates a
        /// new instance.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public InstanceUserArgs()
        {
        }
        public static new InstanceUserArgs Empty => new InstanceUserArgs();
    }
}
