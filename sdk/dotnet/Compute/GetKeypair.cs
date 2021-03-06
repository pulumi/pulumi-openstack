// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    public static class GetKeypair
    {
        /// <summary>
        /// Use this data source to get the ID and public key of an OpenStack keypair.
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
        ///         var kp = Output.Create(OpenStack.Compute.GetKeypair.InvokeAsync(new OpenStack.Compute.GetKeypairArgs
        ///         {
        ///             Name = "sand",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKeypairResult> InvokeAsync(GetKeypairArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetKeypairResult>("openstack:compute/getKeypair:getKeypair", args ?? new GetKeypairArgs(), options.WithVersion());
    }


    public sealed class GetKeypairArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the keypair.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetKeypairArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetKeypairResult
    {
        /// <summary>
        /// The fingerprint of the OpenSSH key.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The OpenSSH-formatted public key of the keypair.
        /// </summary>
        public readonly string PublicKey;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetKeypairResult(
            string fingerprint,

            string id,

            string name,

            string publicKey,

            string region)
        {
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            PublicKey = publicKey;
            Region = region;
        }
    }
}
