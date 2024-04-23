// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity
{
    public static class GetProject
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack project.
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
        ///     var project1 = OpenStack.Identity.GetProject.Invoke(new()
        ///     {
        ///         Name = "demo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("openstack:identity/getProject:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an OpenStack project.
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
        ///     var project1 = OpenStack.Identity.GetProject.Invoke(new()
        ///     {
        ///         Name = "demo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectResult>("openstack:identity/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain this project belongs to.
        /// </summary>
        [Input("domainId")]
        public string? DomainId { get; set; }

        /// <summary>
        /// Whether the project is enabled or disabled. Valid
        /// values are `true` and `false`.
        /// </summary>
        [Input("enabled")]
        public bool? Enabled { get; set; }

        /// <summary>
        /// Whether this project is a domain. Valid values
        /// are `true` and `false`.
        /// </summary>
        [Input("isDomain")]
        public bool? IsDomain { get; set; }

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The parent of this project.
        /// </summary>
        [Input("parentId")]
        public string? ParentId { get; set; }

        /// <summary>
        /// The id of the project. Conflicts with any of the
        /// above arguments.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The region the project is located in.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetProjectArgs()
        {
        }
        public static new GetProjectArgs Empty => new GetProjectArgs();
    }

    public sealed class GetProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain this project belongs to.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// Whether the project is enabled or disabled. Valid
        /// values are `true` and `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Whether this project is a domain. Valid values
        /// are `true` and `false`.
        /// </summary>
        [Input("isDomain")]
        public Input<bool>? IsDomain { get; set; }

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent of this project.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// The id of the project. Conflicts with any of the
        /// above arguments.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region the project is located in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetProjectInvokeArgs()
        {
        }
        public static new GetProjectInvokeArgs Empty => new GetProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// The description of the project.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string DomainId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly bool? IsDomain;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? ParentId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// The region the project is located in.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly ImmutableArray<string> Tags;

        [OutputConstructor]
        private GetProjectResult(
            string description,

            string domainId,

            bool? enabled,

            string id,

            bool? isDomain,

            string? name,

            string? parentId,

            string? projectId,

            string region,

            ImmutableArray<string> tags)
        {
            Description = description;
            DomainId = domainId;
            Enabled = enabled;
            Id = id;
            IsDomain = isDomain;
            Name = name;
            ParentId = parentId;
            ProjectId = projectId;
            Region = region;
            Tags = tags;
        }
    }
}
