// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.OpenStack.Compute
{
    public static class GetInstanceV2
    {
        /// <summary>
        /// Use this data source to get the details of a running server
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
        ///         var instance = Output.Create(OpenStack.Compute.GetInstanceV2.InvokeAsync(new OpenStack.Compute.GetInstanceV2Args
        ///         {
        ///             Id = "2ba26dc6-a12d-4889-8f25-794ea5bf4453",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceV2Result> InvokeAsync(GetInstanceV2Args args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceV2Result>("openstack:compute/getInstanceV2:getInstanceV2", args ?? new GetInstanceV2Args(), options.WithVersion());

        /// <summary>
        /// Use this data source to get the details of a running server
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
        ///         var instance = Output.Create(OpenStack.Compute.GetInstanceV2.InvokeAsync(new OpenStack.Compute.GetInstanceV2Args
        ///         {
        ///             Id = "2ba26dc6-a12d-4889-8f25-794ea5bf4453",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceV2Result> Invoke(GetInstanceV2InvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceV2Result>("openstack:compute/getInstanceV2:getInstanceV2", args ?? new GetInstanceV2InvokeArgs(), options.WithVersion());
    }


    public sealed class GetInstanceV2Args : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The UUID of the instance
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("networks")]
        private List<Inputs.GetInstanceV2NetworkArgs>? _networks;

        /// <summary>
        /// An array of maps, detailed below.
        /// </summary>
        public List<Inputs.GetInstanceV2NetworkArgs> Networks
        {
            get => _networks ?? (_networks = new List<Inputs.GetInstanceV2NetworkArgs>());
            set => _networks = value;
        }

        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The user data added when the server was created.
        /// </summary>
        [Input("userData")]
        public string? UserData { get; set; }

        public GetInstanceV2Args()
        {
        }
    }

    public sealed class GetInstanceV2InvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The UUID of the instance
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("networks")]
        private InputList<Inputs.GetInstanceV2NetworkInputArgs>? _networks;

        /// <summary>
        /// An array of maps, detailed below.
        /// </summary>
        public InputList<Inputs.GetInstanceV2NetworkInputArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.GetInstanceV2NetworkInputArgs>());
            set => _networks = value;
        }

        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The user data added when the server was created.
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public GetInstanceV2InvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceV2Result
    {
        /// <summary>
        /// The first IPv4 address assigned to this server.
        /// </summary>
        public readonly string AccessIpV4;
        /// <summary>
        /// The first IPv6 address assigned to this server.
        /// </summary>
        public readonly string AccessIpV6;
        /// <summary>
        /// The availability zone of this server.
        /// </summary>
        public readonly string AvailabilityZone;
        /// <summary>
        /// The flavor ID used to create the server.
        /// </summary>
        public readonly string FlavorId;
        /// <summary>
        /// The flavor name used to create the server.
        /// </summary>
        public readonly string FlavorName;
        public readonly string Id;
        /// <summary>
        /// The image ID used to create the server.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The image name used to create the server.
        /// </summary>
        public readonly string ImageName;
        /// <summary>
        /// The name of the key pair assigned to this server.
        /// </summary>
        public readonly string KeyPair;
        /// <summary>
        /// A set of key/value pairs made available to the server.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Metadata;
        /// <summary>
        /// The name of the network
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An array of maps, detailed below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceV2NetworkResult> Networks;
        public readonly string PowerState;
        public readonly string Region;
        /// <summary>
        /// An array of security group names associated with this server.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroups;
        /// <summary>
        /// A set of string tags assigned to this server.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The user data added when the server was created.
        /// </summary>
        public readonly string UserData;

        [OutputConstructor]
        private GetInstanceV2Result(
            string accessIpV4,

            string accessIpV6,

            string availabilityZone,

            string flavorId,

            string flavorName,

            string id,

            string imageId,

            string imageName,

            string keyPair,

            ImmutableDictionary<string, object> metadata,

            string name,

            ImmutableArray<Outputs.GetInstanceV2NetworkResult> networks,

            string powerState,

            string region,

            ImmutableArray<string> securityGroups,

            ImmutableArray<string> tags,

            string userData)
        {
            AccessIpV4 = accessIpV4;
            AccessIpV6 = accessIpV6;
            AvailabilityZone = availabilityZone;
            FlavorId = flavorId;
            FlavorName = flavorName;
            Id = id;
            ImageId = imageId;
            ImageName = imageName;
            KeyPair = keyPair;
            Metadata = metadata;
            Name = name;
            Networks = networks;
            PowerState = powerState;
            Region = region;
            SecurityGroups = securityGroups;
            Tags = tags;
            UserData = userData;
        }
    }
}
