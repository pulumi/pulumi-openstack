// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Networking
{
    public static class GetSecGroup
    {
        /// <summary>
        /// Use this data source to get the ID of an available OpenStack security group.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var secgroup = OpenStack.Networking.GetSecGroup.Invoke(new()
        ///     {
        ///         Name = "tf_test_secgroup",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSecGroupResult> InvokeAsync(GetSecGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecGroupResult>("openstack:networking/getSecGroup:getSecGroup", args ?? new GetSecGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID of an available OpenStack security group.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using OpenStack = Pulumi.OpenStack;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var secgroup = OpenStack.Networking.GetSecGroup.Invoke(new()
        ///     {
        ///         Name = "tf_test_secgroup",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSecGroupResult> Invoke(GetSecGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecGroupResult>("openstack:networking/getSecGroup:getSecGroup", args ?? new GetSecGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Human-readable description the the subnet.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// The name of the security group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve security groups ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("secgroupId")]
        public string? SecgroupId { get; set; }

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The list of security group tags to filter.
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the security group.
        /// </summary>
        [Input("tenantId")]
        public string? TenantId { get; set; }

        public GetSecGroupArgs()
        {
        }
        public static new GetSecGroupArgs Empty => new GetSecGroupArgs();
    }

    public sealed class GetSecGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Human-readable description the the subnet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the security group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Neutron client.
        /// A Neutron client is needed to retrieve security groups ids. If omitted, the
        /// `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The ID of the security group.
        /// </summary>
        [Input("secgroupId")]
        public Input<string>? SecgroupId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The list of security group tags to filter.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The owner of the security group.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public GetSecGroupInvokeArgs()
        {
        }
        public static new GetSecGroupInvokeArgs Empty => new GetSecGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecGroupResult
    {
        /// <summary>
        /// The set of string tags applied on the security group.
        /// </summary>
        public readonly ImmutableArray<string> AllTags;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;
        public readonly string? SecgroupId;
        public readonly ImmutableArray<string> Tags;
        public readonly string TenantId;

        [OutputConstructor]
        private GetSecGroupResult(
            ImmutableArray<string> allTags,

            string? description,

            string id,

            string? name,

            string region,

            string? secgroupId,

            ImmutableArray<string> tags,

            string tenantId)
        {
            AllTags = allTags;
            Description = description;
            Id = id;
            Name = name;
            Region = region;
            SecgroupId = secgroupId;
            Tags = tags;
            TenantId = tenantId;
        }
    }
}
