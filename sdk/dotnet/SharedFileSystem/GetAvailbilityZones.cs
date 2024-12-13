// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.SharedFileSystem
{
    public static class GetAvailbilityZones
    {
        /// <summary>
        /// Use this data source to get a list of Shared File System availability zones
        /// from OpenStack
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
        ///     var zones = OpenStack.SharedFileSystem.GetAvailbilityZones.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAvailbilityZonesResult> InvokeAsync(GetAvailbilityZonesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAvailbilityZonesResult>("openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones", args ?? new GetAvailbilityZonesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a list of Shared File System availability zones
        /// from OpenStack
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
        ///     var zones = OpenStack.SharedFileSystem.GetAvailbilityZones.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAvailbilityZonesResult> Invoke(GetAvailbilityZonesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailbilityZonesResult>("openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones", args ?? new GetAvailbilityZonesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a list of Shared File System availability zones
        /// from OpenStack
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
        ///     var zones = OpenStack.SharedFileSystem.GetAvailbilityZones.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAvailbilityZonesResult> Invoke(GetAvailbilityZonesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAvailbilityZonesResult>("openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones", args ?? new GetAvailbilityZonesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAvailbilityZonesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region in which to obtain the V2 Shared File System
        /// client. If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetAvailbilityZonesArgs()
        {
        }
        public static new GetAvailbilityZonesArgs Empty => new GetAvailbilityZonesArgs();
    }

    public sealed class GetAvailbilityZonesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region in which to obtain the V2 Shared File System
        /// client. If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetAvailbilityZonesInvokeArgs()
        {
        }
        public static new GetAvailbilityZonesInvokeArgs Empty => new GetAvailbilityZonesInvokeArgs();
    }


    [OutputType]
    public sealed class GetAvailbilityZonesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The names of the availability zones, ordered alphanumerically.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetAvailbilityZonesResult(
            string id,

            ImmutableArray<string> names,

            string region)
        {
            Id = id;
            Names = names;
            Region = region;
        }
    }
}
