// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Dns
{
    /// <summary>
    /// Manages DNS quota in OpenStack DNS Service.
    /// 
    /// &gt; **Note:** This usually requires admin privileges.
    /// 
    /// &gt; **Note:** This resource has a no-op deletion so no actual actions will be
    /// done against the OpenStack API in case of delete call.
    /// 
    /// ## Import
    /// 
    /// Quotas can be imported using the `project_id/region_name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import openstack:dns/quotaV2:QuotaV2 quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:dns/quotaV2:QuotaV2")]
    public partial class QuotaV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The maximum number of zones that can be
        /// exported via the API.
        /// </summary>
        [Output("apiExportSize")]
        public Output<int> ApiExportSize { get; private set; } = null!;

        /// <summary>
        /// ID of the project to manage quota. Changing this
        /// creates new quota.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The maximum number of records in a
        /// recordset.
        /// </summary>
        [Output("recordsetRecords")]
        public Output<int> RecordsetRecords { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 DNS client. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new DNS quota.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The maximum number of records in a zone.
        /// </summary>
        [Output("zoneRecords")]
        public Output<int> ZoneRecords { get; private set; } = null!;

        /// <summary>
        /// The maximum number of recordsets in a zone.
        /// </summary>
        [Output("zoneRecordsets")]
        public Output<int> ZoneRecordsets { get; private set; } = null!;

        /// <summary>
        /// The maximum number of zones that can be created.
        /// </summary>
        [Output("zones")]
        public Output<int> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a QuotaV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QuotaV2(string name, QuotaV2Args args, CustomResourceOptions? options = null)
            : base("openstack:dns/quotaV2:QuotaV2", name, args ?? new QuotaV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private QuotaV2(string name, Input<string> id, QuotaV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:dns/quotaV2:QuotaV2", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing QuotaV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QuotaV2 Get(string name, Input<string> id, QuotaV2State? state = null, CustomResourceOptions? options = null)
        {
            return new QuotaV2(name, id, state, options);
        }
    }

    public sealed class QuotaV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of zones that can be
        /// exported via the API.
        /// </summary>
        [Input("apiExportSize")]
        public Input<int>? ApiExportSize { get; set; }

        /// <summary>
        /// ID of the project to manage quota. Changing this
        /// creates new quota.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The maximum number of records in a
        /// recordset.
        /// </summary>
        [Input("recordsetRecords")]
        public Input<int>? RecordsetRecords { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 DNS client. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new DNS quota.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The maximum number of records in a zone.
        /// </summary>
        [Input("zoneRecords")]
        public Input<int>? ZoneRecords { get; set; }

        /// <summary>
        /// The maximum number of recordsets in a zone.
        /// </summary>
        [Input("zoneRecordsets")]
        public Input<int>? ZoneRecordsets { get; set; }

        /// <summary>
        /// The maximum number of zones that can be created.
        /// </summary>
        [Input("zones")]
        public Input<int>? Zones { get; set; }

        public QuotaV2Args()
        {
        }
        public static new QuotaV2Args Empty => new QuotaV2Args();
    }

    public sealed class QuotaV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of zones that can be
        /// exported via the API.
        /// </summary>
        [Input("apiExportSize")]
        public Input<int>? ApiExportSize { get; set; }

        /// <summary>
        /// ID of the project to manage quota. Changing this
        /// creates new quota.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The maximum number of records in a
        /// recordset.
        /// </summary>
        [Input("recordsetRecords")]
        public Input<int>? RecordsetRecords { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 DNS client. If
        /// omitted, the `region` argument of the provider is used. Changing this creates
        /// a new DNS quota.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The maximum number of records in a zone.
        /// </summary>
        [Input("zoneRecords")]
        public Input<int>? ZoneRecords { get; set; }

        /// <summary>
        /// The maximum number of recordsets in a zone.
        /// </summary>
        [Input("zoneRecordsets")]
        public Input<int>? ZoneRecordsets { get; set; }

        /// <summary>
        /// The maximum number of zones that can be created.
        /// </summary>
        [Input("zones")]
        public Input<int>? Zones { get; set; }

        public QuotaV2State()
        {
        }
        public static new QuotaV2State Empty => new QuotaV2State();
    }
}
