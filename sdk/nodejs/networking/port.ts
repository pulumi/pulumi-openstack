// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 port resource within OpenStack.
 *
 * > **Note:** Ports do not get an IP if the network they are attached
 * to does not have a subnet. If you create the subnet resource in the
 * same run as the port, make sure to use `fixed_ip.subnet_id` or
 * `dependsOn` to enforce the subnet resource creation before the port
 * creation is triggered. More info here
 *
 * ## Example Usage
 *
 * ### Simple port
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     name: "network_1",
 *     adminStateUp: true,
 * });
 * const port1 = new openstack.networking.Port("port_1", {
 *     name: "port_1",
 *     networkId: network1.id,
 *     adminStateUp: true,
 * });
 * ```
 *
 * ### Port defining fixed_ip.subnet_id
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     name: "network_1",
 *     adminStateUp: true,
 * });
 * const subnet1 = new openstack.networking.Subnet("subnet_1", {
 *     name: "subnet_1",
 *     networkId: network1.id,
 *     cidr: "192.168.199.0/24",
 * });
 * const port1 = new openstack.networking.Port("port_1", {
 *     name: "port_1",
 *     networkId: network1.id,
 *     adminStateUp: true,
 *     fixedIps: [{
 *         subnetId: subnet1.id,
 *     }],
 * });
 * ```
 *
 * ### Port with physical binding information
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     name: "network_1",
 *     adminStateUp: true,
 * });
 * const port1 = new openstack.networking.Port("port_1", {
 *     name: "port_1",
 *     networkId: network1.id,
 *     deviceId: "cdf70fcf-c161-4f24-9c70-96b3f5a54b71",
 *     deviceOwner: "baremetal:none",
 *     adminStateUp: true,
 *     binding: {
 *         hostId: "b080b9cf-46e0-4ce8-ad47-0fd4accc872b",
 *         vnicType: "baremetal",
 *         profile: `{
 *   "local_link_information": [
 *     {
 *       "switch_info": "info1",
 *       "port_id": "Ethernet3/4",
 *       "switch_id": "12:34:56:78:9A:BC"
 *     },
 *     {
 *       "switch_info": "info2",
 *       "port_id": "Ethernet3/4",
 *       "switch_id": "12:34:56:78:9A:BD"
 *     }
 *   ],
 *   "vlan_type": "allowed"
 * }
 * `,
 *     },
 * });
 * ```
 *
 * ## Notes
 *
 * ### Ports and Instances
 *
 * There are some notes to consider when connecting Instances to networks using
 * Ports. Please see the `openstack.compute.Instance` documentation for further
 * documentation.
 *
 * ## Import
 *
 * Ports can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:networking/port:Port port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
 * ```
 */
export class Port extends pulumi.CustomResource {
    /**
     * Get an existing Port resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortState, opts?: pulumi.CustomResourceOptions): Port {
        return new Port(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/port:Port';

    /**
     * Returns true if the given object is an instance of Port.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Port {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Port.__pulumiType;
    }

    /**
     * Administrative up/down status for the port
     * (must be `true` or `false` if provided). Changing this updates the
     * `adminStateUp` of an existing port.
     */
    public readonly adminStateUp!: pulumi.Output<boolean>;
    /**
     * The collection of Fixed IP addresses on the port in the
     * order returned by the Network v2 API.
     */
    public /*out*/ readonly allFixedIps!: pulumi.Output<string[]>;
    /**
     * The collection of Security Group IDs on the port
     * which have been explicitly and implicitly added.
     */
    public /*out*/ readonly allSecurityGroupIds!: pulumi.Output<string[]>;
    /**
     * The collection of tags assigned on the port, which have been
     * explicitly and implicitly added.
     */
    public /*out*/ readonly allTags!: pulumi.Output<string[]>;
    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     */
    public readonly allowedAddressPairs!: pulumi.Output<outputs.networking.PortAllowedAddressPair[] | undefined>;
    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     */
    public readonly binding!: pulumi.Output<outputs.networking.PortBinding>;
    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing port.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * The device owner of the port. Changing this creates
     * a new port.
     */
    public readonly deviceOwner!: pulumi.Output<string>;
    /**
     * The list of maps representing port DNS assignments.
     */
    public /*out*/ readonly dnsAssignments!: pulumi.Output<{[key: string]: string}[]>;
    /**
     * The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     */
    public readonly dnsName!: pulumi.Output<string>;
    /**
     * An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     */
    public readonly extraDhcpOptions!: pulumi.Output<outputs.networking.PortExtraDhcpOption[] | undefined>;
    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     */
    public readonly fixedIps!: pulumi.Output<outputs.networking.PortFixedIp[] | undefined>;
    /**
     * Specify a specific MAC address for the port. Changing
     * this creates a new port.
     */
    public readonly macAddress!: pulumi.Output<string>;
    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     */
    public readonly noFixedIp!: pulumi.Output<boolean | undefined>;
    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `securityGroupIds` are specified, then the port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    public readonly noSecurityGroups!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of `true`. Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     */
    public readonly portSecurityEnabled!: pulumi.Output<boolean>;
    /**
     * Reference to the associated QoS policy.
     */
    public readonly qosPolicyId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * A set of string tags for the port.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The owner of the port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Port resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortArgs | PortState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["allFixedIps"] = state ? state.allFixedIps : undefined;
            resourceInputs["allSecurityGroupIds"] = state ? state.allSecurityGroupIds : undefined;
            resourceInputs["allTags"] = state ? state.allTags : undefined;
            resourceInputs["allowedAddressPairs"] = state ? state.allowedAddressPairs : undefined;
            resourceInputs["binding"] = state ? state.binding : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["deviceOwner"] = state ? state.deviceOwner : undefined;
            resourceInputs["dnsAssignments"] = state ? state.dnsAssignments : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["extraDhcpOptions"] = state ? state.extraDhcpOptions : undefined;
            resourceInputs["fixedIps"] = state ? state.fixedIps : undefined;
            resourceInputs["macAddress"] = state ? state.macAddress : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["noFixedIp"] = state ? state.noFixedIp : undefined;
            resourceInputs["noSecurityGroups"] = state ? state.noSecurityGroups : undefined;
            resourceInputs["portSecurityEnabled"] = state ? state.portSecurityEnabled : undefined;
            resourceInputs["qosPolicyId"] = state ? state.qosPolicyId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as PortArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["allowedAddressPairs"] = args ? args.allowedAddressPairs : undefined;
            resourceInputs["binding"] = args ? args.binding : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["deviceOwner"] = args ? args.deviceOwner : undefined;
            resourceInputs["dnsName"] = args ? args.dnsName : undefined;
            resourceInputs["extraDhcpOptions"] = args ? args.extraDhcpOptions : undefined;
            resourceInputs["fixedIps"] = args ? args.fixedIps : undefined;
            resourceInputs["macAddress"] = args ? args.macAddress : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["noFixedIp"] = args ? args.noFixedIp : undefined;
            resourceInputs["noSecurityGroups"] = args ? args.noSecurityGroups : undefined;
            resourceInputs["portSecurityEnabled"] = args ? args.portSecurityEnabled : undefined;
            resourceInputs["qosPolicyId"] = args ? args.qosPolicyId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["allFixedIps"] = undefined /*out*/;
            resourceInputs["allSecurityGroupIds"] = undefined /*out*/;
            resourceInputs["allTags"] = undefined /*out*/;
            resourceInputs["dnsAssignments"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Port.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Port resources.
 */
export interface PortState {
    /**
     * Administrative up/down status for the port
     * (must be `true` or `false` if provided). Changing this updates the
     * `adminStateUp` of an existing port.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The collection of Fixed IP addresses on the port in the
     * order returned by the Network v2 API.
     */
    allFixedIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The collection of Security Group IDs on the port
     * which have been explicitly and implicitly added.
     */
    allSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The collection of tags assigned on the port, which have been
     * explicitly and implicitly added.
     */
    allTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     */
    allowedAddressPairs?: pulumi.Input<pulumi.Input<inputs.networking.PortAllowedAddressPair>[]>;
    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     */
    binding?: pulumi.Input<inputs.networking.PortBinding>;
    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing port.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * The device owner of the port. Changing this creates
     * a new port.
     */
    deviceOwner?: pulumi.Input<string>;
    /**
     * The list of maps representing port DNS assignments.
     */
    dnsAssignments?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    /**
     * The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     */
    extraDhcpOptions?: pulumi.Input<pulumi.Input<inputs.networking.PortExtraDhcpOption>[]>;
    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     */
    fixedIps?: pulumi.Input<pulumi.Input<inputs.networking.PortFixedIp>[]>;
    /**
     * Specify a specific MAC address for the port. Changing
     * this creates a new port.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     */
    networkId?: pulumi.Input<string>;
    /**
     * Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     */
    noFixedIp?: pulumi.Input<boolean>;
    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `securityGroupIds` are specified, then the port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    noSecurityGroups?: pulumi.Input<boolean>;
    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of `true`. Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     */
    portSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * Reference to the associated QoS policy.
     */
    qosPolicyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     */
    region?: pulumi.Input<string>;
    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of string tags for the port.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Port resource.
 */
export interface PortArgs {
    /**
     * Administrative up/down status for the port
     * (must be `true` or `false` if provided). Changing this updates the
     * `adminStateUp` of an existing port.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     */
    allowedAddressPairs?: pulumi.Input<pulumi.Input<inputs.networking.PortAllowedAddressPair>[]>;
    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     */
    binding?: pulumi.Input<inputs.networking.PortBinding>;
    /**
     * Human-readable description of the port. Changing
     * this updates the `description` of an existing port.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * The device owner of the port. Changing this creates
     * a new port.
     */
    deviceOwner?: pulumi.Input<string>;
    /**
     * The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     */
    extraDhcpOptions?: pulumi.Input<pulumi.Input<inputs.networking.PortExtraDhcpOption>[]>;
    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     */
    fixedIps?: pulumi.Input<pulumi.Input<inputs.networking.PortFixedIp>[]>;
    /**
     * Specify a specific MAC address for the port. Changing
     * this creates a new port.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     */
    networkId: pulumi.Input<string>;
    /**
     * Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     */
    noFixedIp?: pulumi.Input<boolean>;
    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `securityGroupIds` are specified, then the port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    noSecurityGroups?: pulumi.Input<boolean>;
    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of `true`. Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     */
    portSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * Reference to the associated QoS policy.
     */
    qosPolicyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     */
    region?: pulumi.Input<string>;
    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of string tags for the port.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
