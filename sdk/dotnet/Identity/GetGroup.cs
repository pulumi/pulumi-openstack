// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.OpenStack.Identity
{
    public static class GetGroup
    {
        /// <summary>
        /// Use this data source to get the ID of an OpenStack group.
        /// 
        /// &gt; **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        /// this resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var admins = Output.Create(OpenStack.Identity.GetGroup.InvokeAsync(new OpenStack.Identity.GetGroupArgs
        ///         {
        ///             Name = "admins",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("openstack:identity/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithVersion());

        /// <summary>
        /// Use this data source to get the ID of an OpenStack group.
        /// 
        /// &gt; **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        /// this resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var admins = Output.Create(OpenStack.Identity.GetGroup.InvokeAsync(new OpenStack.Identity.GetGroupArgs
        ///         {
        ///             Name = "admins",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupResult>("openstack:identity/getGroup:getGroup", args ?? new GetGroupInvokeArgs(), options.WithVersion());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain the group belongs to.
        /// </summary>
        [Input("domainId")]
        public string? DomainId { get; set; }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetGroupArgs()
        {
        }
    }

    public sealed class GetGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The domain the group belongs to.
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// A description of the group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string DomainId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetGroupResult(
            string description,

            string domainId,

            string id,

            string name,

            string region)
        {
            Description = description;
            DomainId = domainId;
            Id = id;
            Name = name;
            Region = region;
        }
    }
}
