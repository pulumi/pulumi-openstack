// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 router resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const router1 = new openstack.networking.Router("router_1", {
 *     name: "my_router",
 *     adminStateUp: true,
 *     externalNetworkId: "f67f0d72-0ddf-11e4-9d95-e1f29f417e2f",
 * });
 * ```
 *
 * ## Import
 *
 * Routers can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:networking/router:Router router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2
 * ```
 */
export class Router extends pulumi.CustomResource {
    /**
     * Get an existing Router resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterState, opts?: pulumi.CustomResourceOptions): Router {
        return new Router(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/router:Router';

    /**
     * Returns true if the given object is an instance of Router.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Router {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Router.__pulumiType;
    }

    /**
     * Administrative up/down status for the router
     * (must be "true" or "false" if provided). Changing this updates the
     * `adminStateUp` of an existing router.
     */
    public readonly adminStateUp!: pulumi.Output<boolean>;
    /**
     * The collection of tags assigned on the router, which have been
     * explicitly and implicitly added.
     */
    public /*out*/ readonly allTags!: pulumi.Output<string[]>;
    /**
     * An availability zone is used to make 
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new router.
     */
    public readonly availabilityZoneHints!: pulumi.Output<string[]>;
    /**
     * Human-readable description for the router.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether or not to create a
     * distributed router. The default policy setting in Neutron restricts
     * usage of this property to administrative users only.
     */
    public readonly distributed!: pulumi.Output<boolean>;
    /**
     * Enable Source NAT for the router. Valid values are
     * "true" or "false". An `externalNetworkId` has to be set in order to
     * set this property. Changing this updates the `enableSnat` of the router.
     * Setting this value **requires** an **ext-gw-mode** extension to be enabled
     * in OpenStack Neutron.
     */
    public readonly enableSnat!: pulumi.Output<boolean>;
    /**
     * An external fixed IP for the router. This
     * can be repeated. The structure is described below. An `externalNetworkId`
     * has to be set in order to set this property. Changing this updates the
     * external fixed IPs of the router.
     */
    public readonly externalFixedIps!: pulumi.Output<outputs.networking.RouterExternalFixedIp[]>;
    /**
     * The
     * network UUID of an external gateway for the router. A router with an
     * external gateway is required if any compute instances or load balancers
     * will be using floating IPs. Changing this updates the external gateway
     * of an existing router.
     *
     * @deprecated use externalNetworkId instead
     */
    public readonly externalGateway!: pulumi.Output<string>;
    /**
     * The network UUID of an external gateway
     * for the router. A router with an external gateway is required if any
     * compute instances or load balancers will be using floating IPs. Changing
     * this updates the external gateway of the router.
     */
    public readonly externalNetworkId!: pulumi.Output<string>;
    /**
     * A list of external subnet IDs to try over
     * each to obtain a fixed IP for the router. If a subnet ID in a list has
     * exhausted floating IP pool, the next subnet ID will be tried. This argument is
     * used only during the router creation and allows to set only one external fixed
     * IP. Conflicts with an `externalFixedIp` argument.
     */
    public readonly externalSubnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * A unique name for the router. Changing this
     * updates the `name` of an existing router.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A set of string tags for the router.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a router for another tenant. Changing this creates a new router.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Map of additional driver-specific options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    public readonly vendorOptions!: pulumi.Output<outputs.networking.RouterVendorOptions | undefined>;

    /**
     * Create a Router resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RouterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterArgs | RouterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouterState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["allTags"] = state ? state.allTags : undefined;
            resourceInputs["availabilityZoneHints"] = state ? state.availabilityZoneHints : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["distributed"] = state ? state.distributed : undefined;
            resourceInputs["enableSnat"] = state ? state.enableSnat : undefined;
            resourceInputs["externalFixedIps"] = state ? state.externalFixedIps : undefined;
            resourceInputs["externalGateway"] = state ? state.externalGateway : undefined;
            resourceInputs["externalNetworkId"] = state ? state.externalNetworkId : undefined;
            resourceInputs["externalSubnetIds"] = state ? state.externalSubnetIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
            resourceInputs["vendorOptions"] = state ? state.vendorOptions : undefined;
        } else {
            const args = argsOrState as RouterArgs | undefined;
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["availabilityZoneHints"] = args ? args.availabilityZoneHints : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["distributed"] = args ? args.distributed : undefined;
            resourceInputs["enableSnat"] = args ? args.enableSnat : undefined;
            resourceInputs["externalFixedIps"] = args ? args.externalFixedIps : undefined;
            resourceInputs["externalGateway"] = args ? args.externalGateway : undefined;
            resourceInputs["externalNetworkId"] = args ? args.externalNetworkId : undefined;
            resourceInputs["externalSubnetIds"] = args ? args.externalSubnetIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["vendorOptions"] = args ? args.vendorOptions : undefined;
            resourceInputs["allTags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Router.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Router resources.
 */
export interface RouterState {
    /**
     * Administrative up/down status for the router
     * (must be "true" or "false" if provided). Changing this updates the
     * `adminStateUp` of an existing router.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The collection of tags assigned on the router, which have been
     * explicitly and implicitly added.
     */
    allTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An availability zone is used to make 
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new router.
     */
    availabilityZoneHints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable description for the router.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates whether or not to create a
     * distributed router. The default policy setting in Neutron restricts
     * usage of this property to administrative users only.
     */
    distributed?: pulumi.Input<boolean>;
    /**
     * Enable Source NAT for the router. Valid values are
     * "true" or "false". An `externalNetworkId` has to be set in order to
     * set this property. Changing this updates the `enableSnat` of the router.
     * Setting this value **requires** an **ext-gw-mode** extension to be enabled
     * in OpenStack Neutron.
     */
    enableSnat?: pulumi.Input<boolean>;
    /**
     * An external fixed IP for the router. This
     * can be repeated. The structure is described below. An `externalNetworkId`
     * has to be set in order to set this property. Changing this updates the
     * external fixed IPs of the router.
     */
    externalFixedIps?: pulumi.Input<pulumi.Input<inputs.networking.RouterExternalFixedIp>[]>;
    /**
     * The
     * network UUID of an external gateway for the router. A router with an
     * external gateway is required if any compute instances or load balancers
     * will be using floating IPs. Changing this updates the external gateway
     * of an existing router.
     *
     * @deprecated use externalNetworkId instead
     */
    externalGateway?: pulumi.Input<string>;
    /**
     * The network UUID of an external gateway
     * for the router. A router with an external gateway is required if any
     * compute instances or load balancers will be using floating IPs. Changing
     * this updates the external gateway of the router.
     */
    externalNetworkId?: pulumi.Input<string>;
    /**
     * A list of external subnet IDs to try over
     * each to obtain a fixed IP for the router. If a subnet ID in a list has
     * exhausted floating IP pool, the next subnet ID will be tried. This argument is
     * used only during the router creation and allows to set only one external fixed
     * IP. Conflicts with an `externalFixedIp` argument.
     */
    externalSubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique name for the router. Changing this
     * updates the `name` of an existing router.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router.
     */
    region?: pulumi.Input<string>;
    /**
     * A set of string tags for the router.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a router for another tenant. Changing this creates a new router.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional driver-specific options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    vendorOptions?: pulumi.Input<inputs.networking.RouterVendorOptions>;
}

/**
 * The set of arguments for constructing a Router resource.
 */
export interface RouterArgs {
    /**
     * Administrative up/down status for the router
     * (must be "true" or "false" if provided). Changing this updates the
     * `adminStateUp` of an existing router.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * An availability zone is used to make 
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new router.
     */
    availabilityZoneHints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable description for the router.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates whether or not to create a
     * distributed router. The default policy setting in Neutron restricts
     * usage of this property to administrative users only.
     */
    distributed?: pulumi.Input<boolean>;
    /**
     * Enable Source NAT for the router. Valid values are
     * "true" or "false". An `externalNetworkId` has to be set in order to
     * set this property. Changing this updates the `enableSnat` of the router.
     * Setting this value **requires** an **ext-gw-mode** extension to be enabled
     * in OpenStack Neutron.
     */
    enableSnat?: pulumi.Input<boolean>;
    /**
     * An external fixed IP for the router. This
     * can be repeated. The structure is described below. An `externalNetworkId`
     * has to be set in order to set this property. Changing this updates the
     * external fixed IPs of the router.
     */
    externalFixedIps?: pulumi.Input<pulumi.Input<inputs.networking.RouterExternalFixedIp>[]>;
    /**
     * The
     * network UUID of an external gateway for the router. A router with an
     * external gateway is required if any compute instances or load balancers
     * will be using floating IPs. Changing this updates the external gateway
     * of an existing router.
     *
     * @deprecated use externalNetworkId instead
     */
    externalGateway?: pulumi.Input<string>;
    /**
     * The network UUID of an external gateway
     * for the router. A router with an external gateway is required if any
     * compute instances or load balancers will be using floating IPs. Changing
     * this updates the external gateway of the router.
     */
    externalNetworkId?: pulumi.Input<string>;
    /**
     * A list of external subnet IDs to try over
     * each to obtain a fixed IP for the router. If a subnet ID in a list has
     * exhausted floating IP pool, the next subnet ID will be tried. This argument is
     * used only during the router creation and allows to set only one external fixed
     * IP. Conflicts with an `externalFixedIp` argument.
     */
    externalSubnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique name for the router. Changing this
     * updates the `name` of an existing router.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router.
     */
    region?: pulumi.Input<string>;
    /**
     * A set of string tags for the router.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a router for another tenant. Changing this creates a new router.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional driver-specific options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    vendorOptions?: pulumi.Input<inputs.networking.RouterVendorOptions>;
}
