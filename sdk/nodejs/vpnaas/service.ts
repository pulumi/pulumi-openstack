// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron VPN service resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const service1 = new openstack.vpnaas.Service("service_1", {
 *     name: "my_service",
 *     routerId: "14a75700-fc03-4602-9294-26ee44f366b3",
 *     adminStateUp: true,
 * });
 * ```
 *
 * ## Import
 *
 * Services can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:vpnaas/service:Service service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:vpnaas/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing service.
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * The human-readable description for the service.
     * Changing this updates the description of the existing service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The read-only external (public) IPv4 address that is used for the VPN service.
     */
    public /*out*/ readonly externalV4Ip!: pulumi.Output<string>;
    /**
     * The read-only external (public) IPv6 address that is used for the VPN service.
     */
    public /*out*/ readonly externalV6Ip!: pulumi.Output<string>;
    /**
     * The name of the service. Changing this updates the name of
     * the existing service.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The ID of the router. Changing this creates a new service.
     */
    public readonly routerId!: pulumi.Output<string>;
    /**
     * Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * SubnetID is the ID of the subnet. Default is null.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * The owner of the service. Required if admin wants to
     * create a service for another project. Changing this creates a new service.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalV4Ip"] = state ? state.externalV4Ip : undefined;
            resourceInputs["externalV6Ip"] = state ? state.externalV6Ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["routerId"] = state ? state.routerId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.routerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routerId'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["routerId"] = args ? args.routerId : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["externalV4Ip"] = undefined /*out*/;
            resourceInputs["externalV6Ip"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing service.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The human-readable description for the service.
     * Changing this updates the description of the existing service.
     */
    description?: pulumi.Input<string>;
    /**
     * The read-only external (public) IPv4 address that is used for the VPN service.
     */
    externalV4Ip?: pulumi.Input<string>;
    /**
     * The read-only external (public) IPv6 address that is used for the VPN service.
     */
    externalV6Ip?: pulumi.Input<string>;
    /**
     * The name of the service. Changing this updates the name of
     * the existing service.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the router. Changing this creates a new service.
     */
    routerId?: pulumi.Input<string>;
    /**
     * Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.
     */
    status?: pulumi.Input<string>;
    /**
     * SubnetID is the ID of the subnet. Default is null.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The owner of the service. Required if admin wants to
     * create a service for another project. Changing this creates a new service.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing service.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The human-readable description for the service.
     * Changing this updates the description of the existing service.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the service. Changing this updates the name of
     * the existing service.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * service.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the router. Changing this creates a new service.
     */
    routerId: pulumi.Input<string>;
    /**
     * SubnetID is the ID of the subnet. Default is null.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The owner of the service. Required if admin wants to
     * create a service for another project. Changing this creates a new service.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
