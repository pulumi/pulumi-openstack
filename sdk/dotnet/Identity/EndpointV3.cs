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
    /// Manages a V3 Endpoint resource within OpenStack Keystone.
    /// 
    /// &gt; **Note:** This usually requires admin privileges.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var service1 = new OpenStack.Identity.ServiceV3("service1", new()
    ///     {
    ///         Type = "my-service-type",
    ///     });
    /// 
    ///     var endpoint1 = new OpenStack.Identity.EndpointV3("endpoint1", new()
    ///     {
    ///         EndpointRegion = service1.Region,
    ///         ServiceId = service1.Id,
    ///         Url = "http://my-endpoint",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Endpoints can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import openstack:identity/endpointV3:EndpointV3 endpoint_1 5392472b-106a-4845-90c6-7c8445f18770
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:identity/endpointV3:EndpointV3")]
    public partial class EndpointV3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The endpoint region. The `region` and
        /// `endpoint_region` can be different.
        /// </summary>
        [Output("endpointRegion")]
        public Output<string> EndpointRegion { get; private set; } = null!;

        /// <summary>
        /// The endpoint interface. Valid values are `public`,
        /// `internal` and `admin`. Default value is `public`
        /// </summary>
        [Output("interface")]
        public Output<string?> Interface { get; private set; } = null!;

        /// <summary>
        /// The endpoint name.
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
        /// The endpoint service ID.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// The service name of the endpoint.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// The service type of the endpoint.
        /// </summary>
        [Output("serviceType")]
        public Output<string> ServiceType { get; private set; } = null!;

        /// <summary>
        /// The endpoint url.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointV3(string name, EndpointV3Args args, CustomResourceOptions? options = null)
            : base("openstack:identity/endpointV3:EndpointV3", name, args ?? new EndpointV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointV3(string name, Input<string> id, EndpointV3State? state = null, CustomResourceOptions? options = null)
            : base("openstack:identity/endpointV3:EndpointV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointV3 Get(string name, Input<string> id, EndpointV3State? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointV3(name, id, state, options);
        }
    }

    public sealed class EndpointV3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint region. The `region` and
        /// `endpoint_region` can be different.
        /// </summary>
        [Input("endpointRegion", required: true)]
        public Input<string> EndpointRegion { get; set; } = null!;

        /// <summary>
        /// The endpoint interface. Valid values are `public`,
        /// `internal` and `admin`. Default value is `public`
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// The endpoint name.
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
        /// The endpoint service ID.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// The endpoint url.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public EndpointV3Args()
        {
        }
        public static new EndpointV3Args Empty => new EndpointV3Args();
    }

    public sealed class EndpointV3State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint region. The `region` and
        /// `endpoint_region` can be different.
        /// </summary>
        [Input("endpointRegion")]
        public Input<string>? EndpointRegion { get; set; }

        /// <summary>
        /// The endpoint interface. Valid values are `public`,
        /// `internal` and `admin`. Default value is `public`
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// The endpoint name.
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
        /// The endpoint service ID.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The service name of the endpoint.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The service type of the endpoint.
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        /// <summary>
        /// The endpoint url.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public EndpointV3State()
        {
        }
        public static new EndpointV3State Empty => new EndpointV3State();
    }
}
