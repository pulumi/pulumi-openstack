// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    public static class GetHypervisorV2
    {
        /// <summary>
        /// Use this data source to get information about hypervisors
        /// by hostname.
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
        ///         var host01 = Output.Create(OpenStack.Compute.GetHypervisorV2.InvokeAsync(new OpenStack.Compute.GetHypervisorV2Args
        ///         {
        ///             Hostname = "host01",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHypervisorV2Result> InvokeAsync(GetHypervisorV2Args args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHypervisorV2Result>("openstack:compute/getHypervisorV2:getHypervisorV2", args ?? new GetHypervisorV2Args(), options.WithVersion());
    }


    public sealed class GetHypervisorV2Args : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hostname of the hypervisor
        /// </summary>
        [Input("hostname", required: true)]
        public string Hostname { get; set; } = null!;

        public GetHypervisorV2Args()
        {
        }
    }


    [OutputType]
    public sealed class GetHypervisorV2Result
    {
        /// <summary>
        /// The amount in GigaBytes of local storage the hypervisor can provide
        /// </summary>
        public readonly int Disk;
        /// <summary>
        /// The IP address of the Hypervisor
        /// </summary>
        public readonly string HostIp;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The number in MegaBytes of memory the hypervisor can provide
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// The state of the hypervisor (`up` or `down`)
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The status of the hypervisor (`enabled` or `disabled`)
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of the hypervisor (example: `QEMU`)
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The number of virtual CPUs the hypervisor can provide
        /// </summary>
        public readonly int Vcpus;

        [OutputConstructor]
        private GetHypervisorV2Result(
            int disk,

            string hostIp,

            string hostname,

            string id,

            int memory,

            string state,

            string status,

            string type,

            int vcpus)
        {
            Disk = disk;
            HostIp = hostIp;
            Hostname = hostname;
            Id = id;
            Memory = memory;
            State = state;
            Status = status;
            Type = type;
            Vcpus = vcpus;
        }
    }
}