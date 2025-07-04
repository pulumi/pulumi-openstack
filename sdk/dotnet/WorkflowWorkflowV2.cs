// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    public static class WorkflowWorkflowV2
    {
        /// <summary>
        /// Use this data source to get the ID of an available workflow.
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
        ///     var workflow1 = OpenStack.WorkflowWorkflowV2.Invoke(new()
        ///     {
        ///         Name = "workflow_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<WorkflowWorkflowV2Result> InvokeAsync(WorkflowWorkflowV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<WorkflowWorkflowV2Result>("openstack:index/workflowWorkflowV2:WorkflowWorkflowV2", args ?? new WorkflowWorkflowV2Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available workflow.
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
        ///     var workflow1 = OpenStack.WorkflowWorkflowV2.Invoke(new()
        ///     {
        ///         Name = "workflow_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<WorkflowWorkflowV2Result> Invoke(WorkflowWorkflowV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<WorkflowWorkflowV2Result>("openstack:index/workflowWorkflowV2:WorkflowWorkflowV2", args ?? new WorkflowWorkflowV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available workflow.
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
        ///     var workflow1 = OpenStack.WorkflowWorkflowV2.Invoke(new()
        ///     {
        ///         Name = "workflow_1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<WorkflowWorkflowV2Result> Invoke(WorkflowWorkflowV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<WorkflowWorkflowV2Result>("openstack:index/workflowWorkflowV2:WorkflowWorkflowV2", args ?? new WorkflowWorkflowV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class WorkflowWorkflowV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the workflow.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The namespace of the workflow.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// The id of the project to retrieve the workflow.
        /// Requires admin privileges.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Workflow client.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public WorkflowWorkflowV2Args()
        {
        }
        public static new WorkflowWorkflowV2Args Empty => new WorkflowWorkflowV2Args();
    }

    public sealed class WorkflowWorkflowV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the workflow.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace of the workflow.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The id of the project to retrieve the workflow.
        /// Requires admin privileges.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Workflow client.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public WorkflowWorkflowV2InvokeArgs()
        {
        }
        public static new WorkflowWorkflowV2InvokeArgs Empty => new WorkflowWorkflowV2InvokeArgs();
    }


    [OutputType]
    public sealed class WorkflowWorkflowV2Result
    {
        /// <summary>
        /// The date the workflow was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The workflow definition in Mistral v2 DSL.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of input parameters required for workflow execution.
        /// </summary>
        public readonly string Input;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Scope (private or public).
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// A set of string tags for the workflow.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private WorkflowWorkflowV2Result(
            string createdAt,

            string definition,

            string id,

            string input,

            string name,

            string @namespace,

            string projectId,

            string region,

            string scope,

            ImmutableArray<string> tags)
        {
            CreatedAt = createdAt;
            Definition = definition;
            Id = id;
            Input = input;
            Name = name;
            Namespace = @namespace;
            ProjectId = projectId;
            Region = region;
            Scope = scope;
            Tags = tags;
        }
    }
}
