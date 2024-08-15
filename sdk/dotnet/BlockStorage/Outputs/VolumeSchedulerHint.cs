// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.BlockStorage.Outputs
{

    [OutputType]
    public sealed class VolumeSchedulerHint
    {
        /// <summary>
        /// Arbitrary key/value pairs of additional
        /// properties to pass to the scheduler.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AdditionalProperties;
        /// <summary>
        /// The volume should be scheduled on a 
        /// different host from the set of volumes specified in the list provided.
        /// </summary>
        public readonly ImmutableArray<string> DifferentHosts;
        /// <summary>
        /// An instance UUID. The volume should be 
        /// scheduled on the same host as the instance.
        /// </summary>
        public readonly string? LocalToInstance;
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
        public readonly string? Query;
        /// <summary>
        /// A list of volume UUIDs. The volume should be
        /// scheduled on the same host as another volume specified in the list provided.
        /// </summary>
        public readonly ImmutableArray<string> SameHosts;

        [OutputConstructor]
        private VolumeSchedulerHint(
            ImmutableDictionary<string, string>? additionalProperties,

            ImmutableArray<string> differentHosts,

            string? localToInstance,

            string? query,

            ImmutableArray<string> sameHosts)
        {
            AdditionalProperties = additionalProperties;
            DifferentHosts = differentHosts;
            LocalToInstance = localToInstance;
            Query = query;
            SameHosts = sameHosts;
        }
    }
}
