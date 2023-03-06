// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 Endpoint resource within OpenStack Keystone.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const service1 = new openstack.identity.ServiceV3("service1", {type: "my-service-type"});
 * const endpoint1 = new openstack.identity.EndpointV3("endpoint1", {
 *     endpointRegion: service1.region,
 *     serviceId: service1.id,
 *     url: "http://my-endpoint",
 * });
 * ```
 *
 * ## Import
 *
 * Endpoints can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:identity/endpointV3:EndpointV3 endpoint_1 5392472b-106a-4845-90c6-7c8445f18770
 * ```
 */
export class EndpointV3 extends pulumi.CustomResource {
    /**
     * Get an existing EndpointV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointV3State, opts?: pulumi.CustomResourceOptions): EndpointV3 {
        return new EndpointV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:identity/endpointV3:EndpointV3';

    /**
     * Returns true if the given object is an instance of EndpointV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointV3.__pulumiType;
    }

    /**
     * The endpoint region. The `region` and
     * `endpointRegion` can be different.
     */
    public readonly endpointRegion!: pulumi.Output<string>;
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     */
    public readonly interface!: pulumi.Output<string | undefined>;
    /**
     * The endpoint name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The endpoint service ID.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The service name of the endpoint.
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * The service type of the endpoint.
     */
    public /*out*/ readonly serviceType!: pulumi.Output<string>;
    /**
     * The endpoint url.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a EndpointV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointV3Args | EndpointV3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointV3State | undefined;
            resourceInputs["endpointRegion"] = state ? state.endpointRegion : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["serviceType"] = state ? state.serviceType : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as EndpointV3Args | undefined;
            if ((!args || args.endpointRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointRegion'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["endpointRegion"] = args ? args.endpointRegion : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["serviceType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndpointV3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointV3 resources.
 */
export interface EndpointV3State {
    /**
     * The endpoint region. The `region` and
     * `endpointRegion` can be different.
     */
    endpointRegion?: pulumi.Input<string>;
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     */
    interface?: pulumi.Input<string>;
    /**
     * The endpoint name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * The service name of the endpoint.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The service type of the endpoint.
     */
    serviceType?: pulumi.Input<string>;
    /**
     * The endpoint url.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointV3 resource.
 */
export interface EndpointV3Args {
    /**
     * The endpoint region. The `region` and
     * `endpointRegion` can be different.
     */
    endpointRegion: pulumi.Input<string>;
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     */
    interface?: pulumi.Input<string>;
    /**
     * The endpoint name.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    serviceId: pulumi.Input<string>;
    /**
     * The endpoint url.
     */
    url: pulumi.Input<string>;
}
