// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute.Inputs
{

    public sealed class InstanceSchedulerHintGetArgs : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<object>? _additionalProperties;

        /// <summary>
        /// Arbitrary key/value pairs of additional
        /// properties to pass to the scheduler.
        /// </summary>
        public InputMap<object> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<object>());
            set => _additionalProperties = value;
        }

        /// <summary>
        /// An IP Address in CIDR form. The instance
        /// will be placed on a compute node that is in the same subnet.
        /// </summary>
        [Input("buildNearHostIp")]
        public Input<string>? BuildNearHostIp { get; set; }

        [Input("differentCells")]
        private InputList<string>? _differentCells;

        /// <summary>
        /// The names of cells where not to build the instance.
        /// </summary>
        public InputList<string> DifferentCells
        {
            get => _differentCells ?? (_differentCells = new InputList<string>());
            set => _differentCells = value;
        }

        [Input("differentHosts")]
        private InputList<string>? _differentHosts;

        /// <summary>
        /// A list of instance UUIDs. The instance will
        /// be scheduled on a different host than all other instances.
        /// </summary>
        public InputList<string> DifferentHosts
        {
            get => _differentHosts ?? (_differentHosts = new InputList<string>());
            set => _differentHosts = value;
        }

        /// <summary>
        /// A UUID of a Server Group. The instance will be placed
        /// into that group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("queries")]
        private InputList<string>? _queries;

        /// <summary>
        /// A conditional query that a compute node must pass in
        /// order to host an instance. The query must use the `JsonFilter` syntax
        /// which is described
        /// [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter).
        /// At this time, only simple queries are supported. Compound queries using
        /// `and`, `or`, or `not` are not supported. An example of a simple query is:
        /// </summary>
        public InputList<string> Queries
        {
            get => _queries ?? (_queries = new InputList<string>());
            set => _queries = value;
        }

        [Input("sameHosts")]
        private InputList<string>? _sameHosts;

        /// <summary>
        /// A list of instance UUIDs. The instance will be
        /// scheduled on the same host of those specified.
        /// </summary>
        public InputList<string> SameHosts
        {
            get => _sameHosts ?? (_sameHosts = new InputList<string>());
            set => _sameHosts = value;
        }

        /// <summary>
        /// The name of a cell to host the instance.
        /// </summary>
        [Input("targetCell")]
        public Input<string>? TargetCell { get; set; }

        public InstanceSchedulerHintGetArgs()
        {
        }
    }
}
