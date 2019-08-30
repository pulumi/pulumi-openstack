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
 * const service1 = new openstack.identity.ServiceV3("service1", {
 *     type: "my-service-type",
 * });
 * const endpoint1 = new openstack.identity.EndpointV3("endpoint1", {
 *     endpointRegion: service1.region,
 *     serviceId: service1.id,
 *     url: "http://my-endpoint",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_endpoint_v3.html.markdown.
 */
export class EndpointV3 extends pulumi.CustomResource {
    /**
     * Get an existing EndpointV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointV3State | undefined;
            inputs["endpointRegion"] = state ? state.endpointRegion : undefined;
            inputs["interface"] = state ? state.interface : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["serviceId"] = state ? state.serviceId : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["serviceType"] = state ? state.serviceType : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as EndpointV3Args | undefined;
            if (!args || args.endpointRegion === undefined) {
                throw new Error("Missing required property 'endpointRegion'");
            }
            if (!args || args.serviceId === undefined) {
                throw new Error("Missing required property 'serviceId'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["endpointRegion"] = args ? args.endpointRegion : undefined;
            inputs["interface"] = args ? args.interface : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["serviceId"] = args ? args.serviceId : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["serviceName"] = undefined /*out*/;
            inputs["serviceType"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EndpointV3.__pulumiType, name, inputs, opts);
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
    readonly endpointRegion?: pulumi.Input<string>;
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     */
    readonly interface?: pulumi.Input<string>;
    /**
     * The endpoint name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    readonly serviceId?: pulumi.Input<string>;
    /**
     * The service name of the endpoint.
     */
    readonly serviceName?: pulumi.Input<string>;
    /**
     * The service type of the endpoint.
     */
    readonly serviceType?: pulumi.Input<string>;
    /**
     * The endpoint url.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointV3 resource.
 */
export interface EndpointV3Args {
    /**
     * The endpoint region. The `region` and
     * `endpointRegion` can be different.
     */
    readonly endpointRegion: pulumi.Input<string>;
    /**
     * The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     */
    readonly interface?: pulumi.Input<string>;
    /**
     * The endpoint name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The endpoint service ID.
     */
    readonly serviceId: pulumi.Input<string>;
    /**
     * The endpoint url.
     */
    readonly url: pulumi.Input<string>;
}
