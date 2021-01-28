// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Identity
{
    /// <summary>
    /// Manages a V3 Service resource within OpenStack Keystone.
    /// 
    /// &gt; **Note:** This usually requires admin privileges.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var service1 = new OpenStack.Identity.ServiceV3("service1", new OpenStack.Identity.ServiceV3Args
    ///         {
    ///             Type = "custom",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Services can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:identity/serviceV3:ServiceV3 service_1 6688e967-158a-496f-a224-cae3414e6b61
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/serviceV3:ServiceV3")]
    public partial class ServiceV3 : Pulumi.CustomResource
    {
        /// <summary>
        /// The service description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The service status. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The service name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The service type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceV3(string name, ServiceV3Args args, CustomResourceOptions? options = null)
            : base("openstack:identity/serviceV3:ServiceV3", name, args ?? new ServiceV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceV3(string name, Input<string> id, ServiceV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/serviceV3:ServiceV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceV3 Get(string name, Input<string> id, ServiceV3State? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceV3(name, id, state, options);
        }
    }

    public sealed class ServiceV3Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The service status. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The service type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ServiceV3Args()
        {
        }
    }

    public sealed class ServiceV3State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The service status. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to obtain the V3 Keystone client.
        /// If omitted, the `region` argument of the provider is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The service type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ServiceV3State()
        {
        }
    }
}
