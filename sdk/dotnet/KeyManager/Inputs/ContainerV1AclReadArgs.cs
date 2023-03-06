// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.KeyManager.Inputs
{

    public sealed class ContainerV1AclReadArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date the container ACL was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Whether the container is accessible project wide.
        /// Defaults to `true`.
        /// </summary>
        [Input("projectAccess")]
        public Input<bool>? ProjectAccess { get; set; }

        /// <summary>
        /// The date the container ACL was last updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// The list of user IDs, which are allowed to access the
        /// container, when `project_access` is set to `false`.
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public ContainerV1AclReadArgs()
        {
        }
        public static new ContainerV1AclReadArgs Empty => new ContainerV1AclReadArgs();
    }
}
