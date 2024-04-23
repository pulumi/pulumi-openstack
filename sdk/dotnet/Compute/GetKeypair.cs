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
        ///     var kp = OpenStack.Compute.GetKeypair.Invoke(new()
        ///     {
        ///         Name = "sand",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetKeypairResult> InvokeAsync(GetKeypairArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKeypairResult>("openstack:compute/getKeypair:getKeypair", args ?? new GetKeypairArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the ID and public key of an OpenStack keypair.
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
        ///     var kp = OpenStack.Compute.GetKeypair.Invoke(new()
        ///     {
        ///         Name = "sand",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetKeypairResult> Invoke(GetKeypairInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKeypairResult>("openstack:compute/getKeypair:getKeypair", args ?? new GetKeypairInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeypairArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// The user id of the owner of the key pair.
        /// This parameter can be specified only if the provider is configured to use
        /// the credentials of an OpenStack administrator.
        /// </summary>
        [Input("userId")]
        public string? UserId { get; set; }

        public GetKeypairArgs()
        {
        }
        public static new GetKeypairArgs Empty => new GetKeypairArgs();
    }

    public sealed class GetKeypairInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the keypair.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The user id of the owner of the key pair.
        /// This parameter can be specified only if the provider is configured to use
        /// the credentials of an OpenStack administrator.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public GetKeypairInvokeArgs()
        {
        }
        public static new GetKeypairInvokeArgs Empty => new GetKeypairInvokeArgs();
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
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string UserId;

        [OutputConstructor]
        private GetKeypairResult(
            string fingerprint,

            string id,

            string name,

            string publicKey,

            string region,

            string userId)
        {
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            PublicKey = publicKey;
            Region = region;
            UserId = userId;
        }
    }
}
