import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 Neutron IPSec site connection resource within OpenStack.
 */
export declare class SiteConnection extends pulumi.CustomResource {
    /**
     * Get an existing SiteConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SiteConnectionState): SiteConnection;
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     */
    readonly adminStateUp: pulumi.Output<boolean | undefined>;
    /**
     * The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * A dictionary with dead peer detection (DPD) protocol controls.
     * - `action` - (Optional) The dead peer detection (DPD) action.
     * A valid value is clear, hold, restart, disabled, or restart-by-peer.
     * Default value is hold.
     */
    readonly dpds: pulumi.Output<{
        action: string;
        interval: number;
        timeout: number;
    }[]>;
    /**
     * The ID of the IKE policy. Changing this creates a new connection.
     */
    readonly ikepolicyId: pulumi.Output<string>;
    /**
     * A valid value is response-only or bi-directional. Default is bi-directional.
     */
    readonly initiator: pulumi.Output<string>;
    /**
     * The ID of the IPsec policy. Changing this creates a new connection.
     */
    readonly ipsecpolicyId: pulumi.Output<string>;
    /**
     * The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peer_ep_group_id parameter unless
     * in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
     * Changing this updates the existing connection.
     */
    readonly localEpGroupId: pulumi.Output<string | undefined>;
    /**
     * An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     */
    readonly localId: pulumi.Output<string | undefined>;
    /**
     * The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     */
    readonly mtu: pulumi.Output<number>;
    /**
     * The name of the connection. Changing this updates the name of
     * the existing connection.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The peer gateway public IPv4 or IPv6 address or FQDN.
     */
    readonly peerAddress: pulumi.Output<string>;
    /**
     * Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
     */
    readonly peerCidrs: pulumi.Output<string[] | undefined>;
    /**
     * The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
     * You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
     * where peer_cidrs is provided with a subnet_id for the VPN service.
     */
    readonly peerEpGroupId: pulumi.Output<string>;
    /**
     * The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peer_address value.
     * Changing this updates the existing policy.
     */
    readonly peerId: pulumi.Output<string>;
    /**
     * The pre-shared key. A valid value is any string.
     */
    readonly psk: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The ID of the VPN service. Changing this creates a new connection.
     */
    readonly vpnserviceId: pulumi.Output<string>;
    /**
     * Create a SiteConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteConnectionArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering SiteConnection resources.
 */
export interface SiteConnectionState {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A dictionary with dead peer detection (DPD) protocol controls.
     * - `action` - (Optional) The dead peer detection (DPD) action.
     * A valid value is clear, hold, restart, disabled, or restart-by-peer.
     * Default value is hold.
     */
    readonly dpds?: pulumi.Input<{
        action?: pulumi.Input<string>;
        interval?: pulumi.Input<number>;
        timeout?: pulumi.Input<number>;
    }[]>;
    /**
     * The ID of the IKE policy. Changing this creates a new connection.
     */
    readonly ikepolicyId?: pulumi.Input<string>;
    /**
     * A valid value is response-only or bi-directional. Default is bi-directional.
     */
    readonly initiator?: pulumi.Input<string>;
    /**
     * The ID of the IPsec policy. Changing this creates a new connection.
     */
    readonly ipsecpolicyId?: pulumi.Input<string>;
    /**
     * The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peer_ep_group_id parameter unless
     * in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
     * Changing this updates the existing connection.
     */
    readonly localEpGroupId?: pulumi.Input<string>;
    /**
     * An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     */
    readonly localId?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name of the connection. Changing this updates the name of
     * the existing connection.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The peer gateway public IPv4 or IPv6 address or FQDN.
     */
    readonly peerAddress?: pulumi.Input<string>;
    /**
     * Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
     */
    readonly peerCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
     * You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
     * where peer_cidrs is provided with a subnet_id for the VPN service.
     */
    readonly peerEpGroupId?: pulumi.Input<string>;
    /**
     * The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peer_address value.
     * Changing this updates the existing policy.
     */
    readonly peerId?: pulumi.Input<string>;
    /**
     * The pre-shared key. A valid value is any string.
     */
    readonly psk?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the VPN service. Changing this creates a new connection.
     */
    readonly vpnserviceId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a SiteConnection resource.
 */
export interface SiteConnectionArgs {
    /**
     * The administrative state of the resource. Can either be up(true) or down(false).
     * Changing this updates the administrative state of the existing connection.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The human-readable description for the connection.
     * Changing this updates the description of the existing connection.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A dictionary with dead peer detection (DPD) protocol controls.
     * - `action` - (Optional) The dead peer detection (DPD) action.
     * A valid value is clear, hold, restart, disabled, or restart-by-peer.
     * Default value is hold.
     */
    readonly dpds?: pulumi.Input<{
        action?: pulumi.Input<string>;
        interval?: pulumi.Input<number>;
        timeout?: pulumi.Input<number>;
    }[]>;
    /**
     * The ID of the IKE policy. Changing this creates a new connection.
     */
    readonly ikepolicyId: pulumi.Input<string>;
    /**
     * A valid value is response-only or bi-directional. Default is bi-directional.
     */
    readonly initiator?: pulumi.Input<string>;
    /**
     * The ID of the IPsec policy. Changing this creates a new connection.
     */
    readonly ipsecpolicyId: pulumi.Input<string>;
    /**
     * The ID for the endpoint group that contains private subnets for the local side of the connection.
     * You must specify this parameter with the peer_ep_group_id parameter unless
     * in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
     * Changing this updates the existing connection.
     */
    readonly localEpGroupId?: pulumi.Input<string>;
    /**
     * An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
     * Most often, local ID would be domain name, email address, etc.
     * If this is not configured then the external IP address will be used as the ID.
     */
    readonly localId?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) value to address fragmentation.
     * Minimum value is 68 for IPv4, and 1280 for IPv6.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name of the connection. Changing this updates the name of
     * the existing connection.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The peer gateway public IPv4 or IPv6 address or FQDN.
     */
    readonly peerAddress: pulumi.Input<string>;
    /**
     * Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
     */
    readonly peerCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
     * You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
     * where peer_cidrs is provided with a subnet_id for the VPN service.
     */
    readonly peerEpGroupId: pulumi.Input<string>;
    /**
     * The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
     * Typically, this value matches the peer_address value.
     * Changing this updates the existing policy.
     */
    readonly peerId: pulumi.Input<string>;
    /**
     * The pre-shared key. A valid value is any string.
     */
    readonly psk: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec site connection. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * site connection.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the connection. Required if admin wants to
     * create a connection for another project. Changing this creates a new connection.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the VPN service. Changing this creates a new connection.
     */
    readonly vpnserviceId: pulumi.Input<string>;
}
