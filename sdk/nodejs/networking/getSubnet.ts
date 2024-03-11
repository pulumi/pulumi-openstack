// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an available OpenStack subnet.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const subnet1 = openstack.networking.getSubnet({
 *     name: "subnet_1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSubnet(args?: GetSubnetArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:networking/getSubnet:getSubnet", {
        "cidr": args.cidr,
        "description": args.description,
        "dhcpDisabled": args.dhcpDisabled,
        "dhcpEnabled": args.dhcpEnabled,
        "gatewayIp": args.gatewayIp,
        "ipVersion": args.ipVersion,
        "ipv6AddressMode": args.ipv6AddressMode,
        "ipv6RaMode": args.ipv6RaMode,
        "name": args.name,
        "networkId": args.networkId,
        "region": args.region,
        "subnetId": args.subnetId,
        "subnetpoolId": args.subnetpoolId,
        "tags": args.tags,
        "tenantId": args.tenantId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnet.
 */
export interface GetSubnetArgs {
    /**
     * The CIDR of the subnet.
     */
    cidr?: string;
    /**
     * Human-readable description of the subnet.
     */
    description?: string;
    /**
     * @deprecated use dhcp_enabled instead
     */
    dhcpDisabled?: boolean;
    /**
     * If the subnet has DHCP enabled.
     */
    dhcpEnabled?: boolean;
    /**
     * The IP of the subnet's gateway.
     */
    gatewayIp?: string;
    /**
     * The IP version of the subnet (either 4 or 6).
     */
    ipVersion?: number;
    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    ipv6AddressMode?: string;
    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    ipv6RaMode?: string;
    /**
     * The name of the subnet.
     */
    name?: string;
    /**
     * The ID of the network the subnet belongs to.
     */
    networkId?: string;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve subnet ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * The ID of the subnet.
     */
    subnetId?: string;
    /**
     * The ID of the subnetpool associated with the subnet.
     */
    subnetpoolId?: string;
    /**
     * The list of subnet tags to filter.
     */
    tags?: string[];
    /**
     * The owner of the subnet.
     */
    tenantId?: string;
}

/**
 * A collection of values returned by getSubnet.
 */
export interface GetSubnetResult {
    /**
     * A set of string tags applied on the subnet.
     */
    readonly allTags: string[];
    /**
     * Allocation pools of the subnet.
     */
    readonly allocationPools: outputs.networking.GetSubnetAllocationPool[];
    readonly cidr: string;
    readonly description: string;
    /**
     * @deprecated use dhcp_enabled instead
     */
    readonly dhcpDisabled?: boolean;
    readonly dhcpEnabled?: boolean;
    /**
     * DNS Nameservers of the subnet.
     */
    readonly dnsNameservers: string[];
    /**
     * Whether the subnet has DHCP enabled or not.
     */
    readonly enableDhcp: boolean;
    readonly gatewayIp: string;
    /**
     * Host Routes of the subnet.
     */
    readonly hostRoutes: outputs.networking.GetSubnetHostRoute[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipVersion: number;
    readonly ipv6AddressMode: string;
    readonly ipv6RaMode: string;
    readonly name: string;
    readonly networkId: string;
    /**
     * See Argument Reference above.
     */
    readonly region: string;
    /**
     * Service types of the subnet.
     */
    readonly serviceTypes: string[];
    readonly subnetId: string;
    readonly subnetpoolId: string;
    readonly tags?: string[];
    readonly tenantId: string;
}
/**
 * Use this data source to get the ID of an available OpenStack subnet.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const subnet1 = openstack.networking.getSubnet({
 *     name: "subnet_1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getSubnetOutput(args?: GetSubnetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubnetResult> {
    return pulumi.output(args).apply((a: any) => getSubnet(a, opts))
}

/**
 * A collection of arguments for invoking getSubnet.
 */
export interface GetSubnetOutputArgs {
    /**
     * The CIDR of the subnet.
     */
    cidr?: pulumi.Input<string>;
    /**
     * Human-readable description of the subnet.
     */
    description?: pulumi.Input<string>;
    /**
     * @deprecated use dhcp_enabled instead
     */
    dhcpDisabled?: pulumi.Input<boolean>;
    /**
     * If the subnet has DHCP enabled.
     */
    dhcpEnabled?: pulumi.Input<boolean>;
    /**
     * The IP of the subnet's gateway.
     */
    gatewayIp?: pulumi.Input<string>;
    /**
     * The IP version of the subnet (either 4 or 6).
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * The IPv6 address mode. Valid values are
     * `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    ipv6AddressMode?: pulumi.Input<string>;
    /**
     * The IPv6 Router Advertisement mode. Valid values
     * are `dhcpv6-stateful`, `dhcpv6-stateless`, or `slaac`.
     */
    ipv6RaMode?: pulumi.Input<string>;
    /**
     * The name of the subnet.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the network the subnet belongs to.
     */
    networkId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve subnet ids. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The ID of the subnet.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The ID of the subnetpool associated with the subnet.
     */
    subnetpoolId?: pulumi.Input<string>;
    /**
     * The list of subnet tags to filter.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the subnet.
     */
    tenantId?: pulumi.Input<string>;
}
