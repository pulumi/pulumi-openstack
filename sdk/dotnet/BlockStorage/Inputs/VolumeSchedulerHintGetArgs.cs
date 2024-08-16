// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage.Inputs
{

    public sealed class VolumeSchedulerHintGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        private InputMap<string>? _additionalProperties;

        /// <summary>
        /// Arbitrary key/value pairs of additional
        /// properties to pass to the scheduler.
        /// </summary>
        public InputMap<string> AdditionalProperties
        {
            get => _additionalProperties ?? (_additionalProperties = new InputMap<string>());
            set => _additionalProperties = value;
        }

        [Input("differentHosts")]
        private InputList<string>? _differentHosts;

        /// <summary>
        /// The volume should be scheduled on a 
        /// different host from the set of volumes specified in the list provided.
        /// </summary>
        public InputList<string> DifferentHosts
        {
            get => _differentHosts ?? (_differentHosts = new InputList<string>());
            set => _differentHosts = value;
        }

        /// <summary>
        /// An instance UUID. The volume should be 
        /// scheduled on the same host as the instance.
        /// </summary>
        [Input("localToInstance")]
        public Input<string>? LocalToInstance { get; set; }

        /// <summary>
        /// A conditional query that a back-end must pass in
        /// order to host a volume. The query must use the `JsonFilter` syntax
        /// which is described
        /// [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
        /// At this time, only simple queries are supported. Compound queries using
        /// `and`, `or`, or `not` are not supported. An example of a simple query is:
        /// 
        /// ```
        /// [“=”, “$backend_id”, “rbd:vol@ceph#cloud”]
        /// ```
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        [Input("sameHosts")]
        private InputList<string>? _sameHosts;

        /// <summary>
        /// A list of volume UUIDs. The volume should be
        /// scheduled on the same host as another volume specified in the list provided.
        /// </summary>
        public InputList<string> SameHosts
        {
            get => _sameHosts ?? (_sameHosts = new InputList<string>());
            set => _sameHosts = value;
        }

        public VolumeSchedulerHintGetArgs()
        {
        }
        public static new VolumeSchedulerHintGetArgs Empty => new VolumeSchedulerHintGetArgs();
    }
}
