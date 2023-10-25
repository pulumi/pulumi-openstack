// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    /// <summary>
    /// Manages a V2 compute quotaset resource within OpenStack.
    /// 
    /// &gt; **Note:** This usually requires admin privileges.
    /// 
    /// &gt; **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
    ///     in case of delete call.
    /// 
    /// &gt; **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
    ///     created with zero value.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project1 = new OpenStack.Identity.Project("project1");
    /// 
    ///     var quotaset1 = new OpenStack.Compute.QuotaSetV2("quotaset1", new()
    ///     {
    ///         ProjectId = project1.Id,
    ///         KeyPairs = 10,
    ///         Ram = 40960,
    ///         Cores = 32,
    ///         Instances = 20,
    ///         ServerGroups = 4,
    ///         ServerGroupMembers = 8,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Quotasets can be imported using the `project_id/region_name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:compute/quotaSetV2:QuotaSetV2 quotaset_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:compute/quotaSetV2:QuotaSetV2")]
    public partial class QuotaSetV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Quota value for cores.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("cores")]
        public Output<int> Cores { get; private set; } = null!;

        /// <summary>
        /// Quota value for fixed IPs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("fixedIps")]
        public Output<int> FixedIps { get; private set; } = null!;

        /// <summary>
        /// Quota value for floating IPs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("floatingIps")]
        public Output<int> FloatingIps { get; private set; } = null!;

        /// <summary>
        /// Quota value for content bytes
        /// of injected files. Changing this updates the existing quotaset.
        /// </summary>
        [Output("injectedFileContentBytes")]
        public Output<int> InjectedFileContentBytes { get; private set; } = null!;

        /// <summary>
        /// Quota value for path bytes of
        /// injected files. Changing this updates the existing quotaset.
        /// </summary>
        [Output("injectedFilePathBytes")]
        public Output<int> InjectedFilePathBytes { get; private set; } = null!;

        /// <summary>
        /// Quota value for injected files.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("injectedFiles")]
        public Output<int> InjectedFiles { get; private set; } = null!;

        /// <summary>
        /// Quota value for instances.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("instances")]
        public Output<int> Instances { get; private set; } = null!;

        /// <summary>
        /// Quota value for key pairs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("keyPairs")]
        public Output<int> KeyPairs { get; private set; } = null!;

        /// <summary>
        /// Quota value for metadata items.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("metadataItems")]
        public Output<int> MetadataItems { get; private set; } = null!;

        /// <summary>
        /// ID of the project to manage quotas.
        /// Changing this creates a new quotaset.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Quota value for RAM.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("ram")]
        public Output<int> Ram { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Quota value for security group rules.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("securityGroupRules")]
        public Output<int> SecurityGroupRules { get; private set; } = null!;

        /// <summary>
        /// Quota value for security groups.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("securityGroups")]
        public Output<int> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// Quota value for server groups members.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("serverGroupMembers")]
        public Output<int> ServerGroupMembers { get; private set; } = null!;

        /// <summary>
        /// Quota value for server groups.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Output("serverGroups")]
        public Output<int> ServerGroups { get; private set; } = null!;


        /// <summary>
        /// Create a QuotaSetV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QuotaSetV2(string name, QuotaSetV2Args args, CustomResourceOptions? options = null)
            : base("openstack:compute/quotaSetV2:QuotaSetV2", name, args ?? new QuotaSetV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private QuotaSetV2(string name, Input<string> id, QuotaSetV2State? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/quotaSetV2:QuotaSetV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QuotaSetV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QuotaSetV2 Get(string name, Input<string> id, QuotaSetV2State? state = null, CustomResourceOptions? options = null)
        {
            return new QuotaSetV2(name, id, state, options);
        }
    }

    public sealed class QuotaSetV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Quota value for cores.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("cores")]
        public Input<int>? Cores { get; set; }

        /// <summary>
        /// Quota value for fixed IPs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("fixedIps")]
        public Input<int>? FixedIps { get; set; }

        /// <summary>
        /// Quota value for floating IPs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("floatingIps")]
        public Input<int>? FloatingIps { get; set; }

        /// <summary>
        /// Quota value for content bytes
        /// of injected files. Changing this updates the existing quotaset.
        /// </summary>
        [Input("injectedFileContentBytes")]
        public Input<int>? InjectedFileContentBytes { get; set; }

        /// <summary>
        /// Quota value for path bytes of
        /// injected files. Changing this updates the existing quotaset.
        /// </summary>
        [Input("injectedFilePathBytes")]
        public Input<int>? InjectedFilePathBytes { get; set; }

        /// <summary>
        /// Quota value for injected files.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("injectedFiles")]
        public Input<int>? InjectedFiles { get; set; }

        /// <summary>
        /// Quota value for instances.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("instances")]
        public Input<int>? Instances { get; set; }

        /// <summary>
        /// Quota value for key pairs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("keyPairs")]
        public Input<int>? KeyPairs { get; set; }

        /// <summary>
        /// Quota value for metadata items.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("metadataItems")]
        public Input<int>? MetadataItems { get; set; }

        /// <summary>
        /// ID of the project to manage quotas.
        /// Changing this creates a new quotaset.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Quota value for RAM.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("ram")]
        public Input<int>? Ram { get; set; }

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Quota value for security group rules.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("securityGroupRules")]
        public Input<int>? SecurityGroupRules { get; set; }

        /// <summary>
        /// Quota value for security groups.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("securityGroups")]
        public Input<int>? SecurityGroups { get; set; }

        /// <summary>
        /// Quota value for server groups members.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("serverGroupMembers")]
        public Input<int>? ServerGroupMembers { get; set; }

        /// <summary>
        /// Quota value for server groups.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("serverGroups")]
        public Input<int>? ServerGroups { get; set; }

        public QuotaSetV2Args()
        {
        }
        public static new QuotaSetV2Args Empty => new QuotaSetV2Args();
    }

    public sealed class QuotaSetV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Quota value for cores.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("cores")]
        public Input<int>? Cores { get; set; }

        /// <summary>
        /// Quota value for fixed IPs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("fixedIps")]
        public Input<int>? FixedIps { get; set; }

        /// <summary>
        /// Quota value for floating IPs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("floatingIps")]
        public Input<int>? FloatingIps { get; set; }

        /// <summary>
        /// Quota value for content bytes
        /// of injected files. Changing this updates the existing quotaset.
        /// </summary>
        [Input("injectedFileContentBytes")]
        public Input<int>? InjectedFileContentBytes { get; set; }

        /// <summary>
        /// Quota value for path bytes of
        /// injected files. Changing this updates the existing quotaset.
        /// </summary>
        [Input("injectedFilePathBytes")]
        public Input<int>? InjectedFilePathBytes { get; set; }

        /// <summary>
        /// Quota value for injected files.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("injectedFiles")]
        public Input<int>? InjectedFiles { get; set; }

        /// <summary>
        /// Quota value for instances.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("instances")]
        public Input<int>? Instances { get; set; }

        /// <summary>
        /// Quota value for key pairs.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("keyPairs")]
        public Input<int>? KeyPairs { get; set; }

        /// <summary>
        /// Quota value for metadata items.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("metadataItems")]
        public Input<int>? MetadataItems { get; set; }

        /// <summary>
        /// ID of the project to manage quotas.
        /// Changing this creates a new quotaset.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Quota value for RAM.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("ram")]
        public Input<int>? Ram { get; set; }

        /// <summary>
        /// The region in which to create the volume. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new quotaset.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Quota value for security group rules.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("securityGroupRules")]
        public Input<int>? SecurityGroupRules { get; set; }

        /// <summary>
        /// Quota value for security groups.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("securityGroups")]
        public Input<int>? SecurityGroups { get; set; }

        /// <summary>
        /// Quota value for server groups members.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("serverGroupMembers")]
        public Input<int>? ServerGroupMembers { get; set; }

        /// <summary>
        /// Quota value for server groups.
        /// Changing this updates the existing quotaset.
        /// </summary>
        [Input("serverGroups")]
        public Input<int>? ServerGroups { get; set; }

        public QuotaSetV2State()
        {
        }
        public static new QuotaSetV2State Empty => new QuotaSetV2State();
    }
}
