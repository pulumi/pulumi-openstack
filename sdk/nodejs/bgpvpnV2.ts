// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a V2 BGP VPN service resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const bgpvpn1 = new openstack.bgpvpn.V2("bgpvpn_1", {
 *     name: "bgpvpn1",
 *     routeDistinguishers: ["64512:1"],
 *     routeTargets: ["64512:1"],
 *     importTargets: ["64512:2"],
 *     exportTargets: ["64512:3"],
 * });
 * ```
 *
 * ## Import
 *
 * BGP VPNs can be imported using the `id`, e.g.
 *
 * hcl
 *
 * ```sh
 * $ pulumi import openstack:index/bgpvpnV2:BgpvpnV2 bgpvpn_1 1eec2c66-6be2-4305-af3f-354c9b81f18c
 * ```
 *
 * @deprecated openstack.index/bgpvpnv2.BgpvpnV2 has been deprecated in favor of openstack.bgpvpn/v2.V2
 */
export class BgpvpnV2 extends pulumi.CustomResource {
    /**
     * Get an existing BgpvpnV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpvpnV2State, opts?: pulumi.CustomResourceOptions): BgpvpnV2 {
        pulumi.log.warn("BgpvpnV2 is deprecated: openstack.index/bgpvpnv2.BgpvpnV2 has been deprecated in favor of openstack.bgpvpn/v2.V2")
        return new BgpvpnV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:index/bgpvpnV2:BgpvpnV2';

    /**
     * Returns true if the given object is an instance of BgpvpnV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BgpvpnV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BgpvpnV2.__pulumiType;
    }

    /**
     * A list of additional Route Targets that will be
     * used for export.
     */
    public readonly exportTargets!: pulumi.Output<string[]>;
    /**
     * A list of additional Route Targets that will be
     * imported.
     */
    public readonly importTargets!: pulumi.Output<string[]>;
    /**
     * The default BGP LOCAL\_PREF of routes that will be
     * advertised to the BGP VPN, unless overridden per-route.
     */
    public readonly localPref!: pulumi.Output<number | undefined>;
    /**
     * The name of the BGP VPN. Changing this updates the name of
     * the existing BGP VPN.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of network IDs that are associated with the BGP VPN.
     */
    public /*out*/ readonly networks!: pulumi.Output<string[]>;
    /**
     * A list of port IDs that are associated with the BGP VPN.
     */
    public /*out*/ readonly ports!: pulumi.Output<string[]>;
    /**
     * The ID of the project that owns the BGPVPN. Only
     * administrative and users with `advsvc` role can specify a project ID other
     * than their own. Changing this creates a new BGP VPN.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a BGP VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * BGP VPN.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A list of route distinguisher strings. If
     * specified, one of these RDs will be used to advertise VPN routes.
     */
    public readonly routeDistinguishers!: pulumi.Output<string[]>;
    /**
     * A list of Route Targets that will be both
     * imported and used for export.
     */
    public readonly routeTargets!: pulumi.Output<string[]>;
    /**
     * A list of router IDs that are associated with the BGP VPN.
     */
    public /*out*/ readonly routers!: pulumi.Output<string[]>;
    /**
     * Indicates whether the BGP VPN is shared across projects.
     */
    public /*out*/ readonly shared!: pulumi.Output<boolean>;
    /**
     * The type of the BGP VPN (either `l2` or `l3`). Changing this
     * creates a new BGP VPN. Defaults to `l3`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The globally-assigned VXLAN VNI for the BGP VPN. Changing
     * this creates a new BGP VPN.
     */
    public readonly vni!: pulumi.Output<number | undefined>;

    /**
     * Create a BgpvpnV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated openstack.index/bgpvpnv2.BgpvpnV2 has been deprecated in favor of openstack.bgpvpn/v2.V2 */
    constructor(name: string, args?: BgpvpnV2Args, opts?: pulumi.CustomResourceOptions)
    /** @deprecated openstack.index/bgpvpnv2.BgpvpnV2 has been deprecated in favor of openstack.bgpvpn/v2.V2 */
    constructor(name: string, argsOrState?: BgpvpnV2Args | BgpvpnV2State, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("BgpvpnV2 is deprecated: openstack.index/bgpvpnv2.BgpvpnV2 has been deprecated in favor of openstack.bgpvpn/v2.V2")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BgpvpnV2State | undefined;
            resourceInputs["exportTargets"] = state ? state.exportTargets : undefined;
            resourceInputs["importTargets"] = state ? state.importTargets : undefined;
            resourceInputs["localPref"] = state ? state.localPref : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["routeDistinguishers"] = state ? state.routeDistinguishers : undefined;
            resourceInputs["routeTargets"] = state ? state.routeTargets : undefined;
            resourceInputs["routers"] = state ? state.routers : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vni"] = state ? state.vni : undefined;
        } else {
            const args = argsOrState as BgpvpnV2Args | undefined;
            resourceInputs["exportTargets"] = args ? args.exportTargets : undefined;
            resourceInputs["importTargets"] = args ? args.importTargets : undefined;
            resourceInputs["localPref"] = args ? args.localPref : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["routeDistinguishers"] = args ? args.routeDistinguishers : undefined;
            resourceInputs["routeTargets"] = args ? args.routeTargets : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vni"] = args ? args.vni : undefined;
            resourceInputs["networks"] = undefined /*out*/;
            resourceInputs["ports"] = undefined /*out*/;
            resourceInputs["routers"] = undefined /*out*/;
            resourceInputs["shared"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BgpvpnV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpvpnV2 resources.
 */
export interface BgpvpnV2State {
    /**
     * A list of additional Route Targets that will be
     * used for export.
     */
    exportTargets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of additional Route Targets that will be
     * imported.
     */
    importTargets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default BGP LOCAL\_PREF of routes that will be
     * advertised to the BGP VPN, unless overridden per-route.
     */
    localPref?: pulumi.Input<number>;
    /**
     * The name of the BGP VPN. Changing this updates the name of
     * the existing BGP VPN.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of network IDs that are associated with the BGP VPN.
     */
    networks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of port IDs that are associated with the BGP VPN.
     */
    ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project that owns the BGPVPN. Only
     * administrative and users with `advsvc` role can specify a project ID other
     * than their own. Changing this creates a new BGP VPN.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a BGP VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * BGP VPN.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of route distinguisher strings. If
     * specified, one of these RDs will be used to advertise VPN routes.
     */
    routeDistinguishers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Route Targets that will be both
     * imported and used for export.
     */
    routeTargets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of router IDs that are associated with the BGP VPN.
     */
    routers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether the BGP VPN is shared across projects.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * The type of the BGP VPN (either `l2` or `l3`). Changing this
     * creates a new BGP VPN. Defaults to `l3`.
     */
    type?: pulumi.Input<string>;
    /**
     * The globally-assigned VXLAN VNI for the BGP VPN. Changing
     * this creates a new BGP VPN.
     */
    vni?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BgpvpnV2 resource.
 */
export interface BgpvpnV2Args {
    /**
     * A list of additional Route Targets that will be
     * used for export.
     */
    exportTargets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of additional Route Targets that will be
     * imported.
     */
    importTargets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default BGP LOCAL\_PREF of routes that will be
     * advertised to the BGP VPN, unless overridden per-route.
     */
    localPref?: pulumi.Input<number>;
    /**
     * The name of the BGP VPN. Changing this updates the name of
     * the existing BGP VPN.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project that owns the BGPVPN. Only
     * administrative and users with `advsvc` role can specify a project ID other
     * than their own. Changing this creates a new BGP VPN.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a BGP VPN service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * BGP VPN.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of route distinguisher strings. If
     * specified, one of these RDs will be used to advertise VPN routes.
     */
    routeDistinguishers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of Route Targets that will be both
     * imported and used for export.
     */
    routeTargets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of the BGP VPN (either `l2` or `l3`). Changing this
     * creates a new BGP VPN. Defaults to `l3`.
     */
    type?: pulumi.Input<string>;
    /**
     * The globally-assigned VXLAN VNI for the BGP VPN. Changing
     * this creates a new BGP VPN.
     */
    vni?: pulumi.Input<number>;
}
