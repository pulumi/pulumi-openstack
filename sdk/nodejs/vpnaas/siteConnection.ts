// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron IPSec site connection resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const conn1 = new openstack.vpnaas.SiteConnection("conn_1", {
 *     ikepolicyId: openstack_vpnaas_ike_policy_v2_policy_2.id,
 *     ipsecpolicyId: openstack_vpnaas_ipsec_policy_v2_policy_1.id,
 *     localEpGroupId: openstack_vpnaas_endpoint_group_v2_group_2.id,
 *     peerAddress: "192.168.10.1",
 *     peerEpGroupId: openstack_vpnaas_endpoint_group_v2_group_1.id,
 *     psk: "secret",
 *     vpnserviceId: openstack_vpnaas_service_v2_service_1.id,
 * });
 * ```
 *
 * ## Import
 *
 * Site Connections can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:vpnaas/siteConnection:SiteConnection conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 */
export class SiteConnection extends pulumi.CustomResource {
    /**
     * Get an existing SiteConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SiteConnectionState, opts?: pulumi.CustomResourceOptions): SiteConnection {
        return new SiteConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:vpnaas/siteConnection:SiteConnection';

    /**
     * Returns true if the given object is an instance of SiteConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SiteConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SiteConnection.__pulumiType;
    }

    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A dictionary with dead peer detection (DPD) protocol controls.
     */
    public readonly dpds!: pulumi.Output<outputs.vpnaas.SiteConnectionDpd[]>;
    /**
     * The ID of the IKE policy. Changing this creates a new connection.
     */
    public readonly ikepolicyId!: pulumi.Output<string>;
    /**
     * A valid value is response-only or bi-directional. Default is bi-directional.
     */
    public readonly initiator!: pulumi.Output<string>;
    /**
     * The ID of the IPsec policy. Changing this creates a new connection.
     */
    public readonly ipsecpolicyId!: pulumi.Output<string>;
    /**
     * The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peerEpGroupId parameter unless
     * in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
     * Changing this updates the existing connection.
     */
    public readonly localEpGroupId!: pulumi.Output<string | undefined>;
    /**
     * An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     */
    public readonly localId!: pulumi.Output<string | undefined>;
    /**
     * The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * The name of the connection. Changing this updates the name of
     * the existing connection.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The peer gateway public IPv4 or IPv6 address or FQDN.
     */
    public readonly peerAddress!: pulumi.Output<string>;
    /**
     * Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
     */
    public readonly peerCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
     * You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
     * where peerCidrs is provided with a subnetId for the VPN service.
     */
    public readonly peerEpGroupId!: pulumi.Output<string | undefined>;
    /**
     * The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peerAddress value.
     * Changing this updates the existing policy.
     */
    public readonly peerId!: pulumi.Output<string>;
    /**
     * The pre-shared key. A valid value is any string.
     */
    public readonly psk!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the VPN service. Changing this creates a new connection.
     */
    public readonly vpnserviceId!: pulumi.Output<string>;

    /**
     * Create a SiteConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SiteConnectionArgs | SiteConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SiteConnectionState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dpds"] = state ? state.dpds : undefined;
            resourceInputs["ikepolicyId"] = state ? state.ikepolicyId : undefined;
            resourceInputs["initiator"] = state ? state.initiator : undefined;
            resourceInputs["ipsecpolicyId"] = state ? state.ipsecpolicyId : undefined;
            resourceInputs["localEpGroupId"] = state ? state.localEpGroupId : undefined;
            resourceInputs["localId"] = state ? state.localId : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["peerAddress"] = state ? state.peerAddress : undefined;
            resourceInputs["peerCidrs"] = state ? state.peerCidrs : undefined;
            resourceInputs["peerEpGroupId"] = state ? state.peerEpGroupId : undefined;
            resourceInputs["peerId"] = state ? state.peerId : undefined;
            resourceInputs["psk"] = state ? state.psk : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
            resourceInputs["vpnserviceId"] = state ? state.vpnserviceId : undefined;
        } else {
            const args = argsOrState as SiteConnectionArgs | undefined;
            if ((!args || args.ikepolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ikepolicyId'");
            }
            if ((!args || args.ipsecpolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipsecpolicyId'");
            }
            if ((!args || args.peerAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerAddress'");
            }
            if ((!args || args.peerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'peerId'");
            }
            if ((!args || args.psk === undefined) && !opts.urn) {
                throw new Error("Missing required property 'psk'");
            }
            if ((!args || args.vpnserviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpnserviceId'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dpds"] = args ? args.dpds : undefined;
            resourceInputs["ikepolicyId"] = args ? args.ikepolicyId : undefined;
            resourceInputs["initiator"] = args ? args.initiator : undefined;
            resourceInputs["ipsecpolicyId"] = args ? args.ipsecpolicyId : undefined;
            resourceInputs["localEpGroupId"] = args ? args.localEpGroupId : undefined;
            resourceInputs["localId"] = args ? args.localId : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["peerAddress"] = args ? args.peerAddress : undefined;
            resourceInputs["peerCidrs"] = args ? args.peerCidrs : undefined;
            resourceInputs["peerEpGroupId"] = args ? args.peerEpGroupId : undefined;
            resourceInputs["peerId"] = args ? args.peerId : undefined;
            resourceInputs["psk"] = args ? args.psk : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["vpnserviceId"] = args ? args.vpnserviceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SiteConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SiteConnection resources.
 */
export interface SiteConnectionState {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     */
    description?: pulumi.Input<string>;
    /**
     * A dictionary with dead peer detection (DPD) protocol controls.
     */
    dpds?: pulumi.Input<pulumi.Input<inputs.vpnaas.SiteConnectionDpd>[]>;
    /**
     * The ID of the IKE policy. Changing this creates a new connection.
     */
    ikepolicyId?: pulumi.Input<string>;
    /**
     * A valid value is response-only or bi-directional. Default is bi-directional.
     */
    initiator?: pulumi.Input<string>;
    /**
     * The ID of the IPsec policy. Changing this creates a new connection.
     */
    ipsecpolicyId?: pulumi.Input<string>;
    /**
     * The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peerEpGroupId parameter unless
     * in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
     * Changing this updates the existing connection.
     */
    localEpGroupId?: pulumi.Input<string>;
    /**
     * An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     */
    localId?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name of the connection. Changing this updates the name of
     * the existing connection.
     */
    name?: pulumi.Input<string>;
    /**
     * The peer gateway public IPv4 or IPv6 address or FQDN.
     */
    peerAddress?: pulumi.Input<string>;
    /**
     * Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
     */
    peerCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
     * You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
     * where peerCidrs is provided with a subnetId for the VPN service.
     */
    peerEpGroupId?: pulumi.Input<string>;
    /**
     * The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peerAddress value.
     * Changing this updates the existing policy.
     */
    peerId?: pulumi.Input<string>;
    /**
     * The pre-shared key. A valid value is any string.
     */
    psk?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     */
    region?: pulumi.Input<string>;
    /**
     * The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the VPN service. Changing this creates a new connection.
     */
    vpnserviceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SiteConnection resource.
 */
export interface SiteConnectionArgs {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     */
    description?: pulumi.Input<string>;
    /**
     * A dictionary with dead peer detection (DPD) protocol controls.
     */
    dpds?: pulumi.Input<pulumi.Input<inputs.vpnaas.SiteConnectionDpd>[]>;
    /**
     * The ID of the IKE policy. Changing this creates a new connection.
     */
    ikepolicyId: pulumi.Input<string>;
    /**
     * A valid value is response-only or bi-directional. Default is bi-directional.
     */
    initiator?: pulumi.Input<string>;
    /**
     * The ID of the IPsec policy. Changing this creates a new connection.
     */
    ipsecpolicyId: pulumi.Input<string>;
    /**
     * The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peerEpGroupId parameter unless
     * in backward- compatible mode where peerCidrs is provided with a subnetId for the VPN service.
     * Changing this updates the existing connection.
     */
    localEpGroupId?: pulumi.Input<string>;
    /**
     * An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     */
    localId?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name of the connection. Changing this updates the name of
     * the existing connection.
     */
    name?: pulumi.Input<string>;
    /**
     * The peer gateway public IPv4 or IPv6 address or FQDN.
     */
    peerAddress: pulumi.Input<string>;
    /**
     * Unique list of valid peer private CIDRs in the form < netAddress > / < prefix > .
     */
    peerCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID for the endpoint group that contains private CIDRs in the form < netAddress > / < prefix > for the peer side of the connection.
     * You must specify this parameter with the localEpGroupId parameter unless in backward-compatible mode
     * where peerCidrs is provided with a subnetId for the VPN service.
     */
    peerEpGroupId?: pulumi.Input<string>;
    /**
     * The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peerAddress value.
     * Changing this updates the existing policy.
     */
    peerId: pulumi.Input<string>;
    /**
     * The pre-shared key. A valid value is any string.
     */
    psk: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     */
    region?: pulumi.Input<string>;
    /**
     * The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the VPN service. Changing this creates a new connection.
     */
    vpnserviceId: pulumi.Input<string>;
}
