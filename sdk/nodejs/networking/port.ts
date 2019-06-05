// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 port resource within OpenStack.
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
 *     adminStateUp: true,
 * });
 * const port1 = new openstack.networking.Port("port_1", {
 *     adminStateUp: true,
 *     networkId: network1.id,
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
 *     adminStateUp: true,
 * });
 * const port1 = new openstack.networking.Port("port_1", {
 *     adminStateUp: true,
 *     binding: {
 *         hostId: "b080b9cf-46e0-4ce8-ad47-0fd4accc872b",
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
 *         vnicType: "baremetal",
 *     },
 *     deviceId: "cdf70fcf-c161-4f24-9c70-96b3f5a54b71",
 *     deviceOwner: "baremetal:none",
 *     networkId: network1.id,
 * });
 * ```
 * 
 * ## Notes
 * 
 * ### Ports and Instances
 * 
 * There are some notes to consider when connecting Instances to networks using
 * Ports. Please see the `openstack_compute_instance_v2` documentation for further
 * documentation.
 */
export class Port extends pulumi.CustomResource {
    /**
     * Get an existing Port resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * (must be "true" or "false" if provided). Changing this updates the
     * `admin_state_up` of an existing port.
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
    public readonly allowedAddressPairs!: pulumi.Output<{ ipAddress: string, macAddress?: string }[] | undefined>;
    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     */
    public readonly binding!: pulumi.Output<{ hostId?: string, profile?: string, vifDetails: {[key: string]: any}, vifType: string, vnicType?: string }>;
    /**
     * Human-readable description of the floating IP. Changing
     * this updates the `description` of an existing port.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * The device owner of the Port. Changing this creates
     * a new port.
     */
    public readonly deviceOwner!: pulumi.Output<string>;
    /**
     * The list of maps representing port DNS assignments.
     */
    public /*out*/ readonly dnsAssignments!: pulumi.Output<{[key: string]: any}[]>;
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
    public readonly extraDhcpOptions!: pulumi.Output<{ ipVersion?: number, name: string, value: string }[] | undefined>;
    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     */
    public readonly fixedIps!: pulumi.Output<{ ipAddress?: string, subnetId: string }[] | undefined>;
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
     * no `security_group_ids` are specified, then the Port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    public readonly noSecurityGroups!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of "true". Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     */
    public readonly portSecurityEnabled!: pulumi.Output<boolean>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
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
     * The owner of the Port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Port resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortArgs | PortState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PortState | undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["allFixedIps"] = state ? state.allFixedIps : undefined;
            inputs["allSecurityGroupIds"] = state ? state.allSecurityGroupIds : undefined;
            inputs["allTags"] = state ? state.allTags : undefined;
            inputs["allowedAddressPairs"] = state ? state.allowedAddressPairs : undefined;
            inputs["binding"] = state ? state.binding : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["deviceId"] = state ? state.deviceId : undefined;
            inputs["deviceOwner"] = state ? state.deviceOwner : undefined;
            inputs["dnsAssignments"] = state ? state.dnsAssignments : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["extraDhcpOptions"] = state ? state.extraDhcpOptions : undefined;
            inputs["fixedIps"] = state ? state.fixedIps : undefined;
            inputs["macAddress"] = state ? state.macAddress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkId"] = state ? state.networkId : undefined;
            inputs["noFixedIp"] = state ? state.noFixedIp : undefined;
            inputs["noSecurityGroups"] = state ? state.noSecurityGroups : undefined;
            inputs["portSecurityEnabled"] = state ? state.portSecurityEnabled : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as PortArgs | undefined;
            if (!args || args.networkId === undefined) {
                throw new Error("Missing required property 'networkId'");
            }
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["allowedAddressPairs"] = args ? args.allowedAddressPairs : undefined;
            inputs["binding"] = args ? args.binding : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["deviceId"] = args ? args.deviceId : undefined;
            inputs["deviceOwner"] = args ? args.deviceOwner : undefined;
            inputs["dnsName"] = args ? args.dnsName : undefined;
            inputs["extraDhcpOptions"] = args ? args.extraDhcpOptions : undefined;
            inputs["fixedIps"] = args ? args.fixedIps : undefined;
            inputs["macAddress"] = args ? args.macAddress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkId"] = args ? args.networkId : undefined;
            inputs["noFixedIp"] = args ? args.noFixedIp : undefined;
            inputs["noSecurityGroups"] = args ? args.noSecurityGroups : undefined;
            inputs["portSecurityEnabled"] = args ? args.portSecurityEnabled : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            inputs["allFixedIps"] = undefined /*out*/;
            inputs["allSecurityGroupIds"] = undefined /*out*/;
            inputs["allTags"] = undefined /*out*/;
            inputs["dnsAssignments"] = undefined /*out*/;
        }
        super(Port.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Port resources.
 */
export interface PortState {
    /**
     * Administrative up/down status for the port
     * (must be "true" or "false" if provided). Changing this updates the
     * `admin_state_up` of an existing port.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The collection of Fixed IP addresses on the port in the
     * order returned by the Network v2 API.
     */
    readonly allFixedIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The collection of Security Group IDs on the port
     * which have been explicitly and implicitly added.
     */
    readonly allSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The collection of tags assigned on the port, which have been
     * explicitly and implicitly added.
     */
    readonly allTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     */
    readonly allowedAddressPairs?: pulumi.Input<pulumi.Input<{ ipAddress: pulumi.Input<string>, macAddress?: pulumi.Input<string> }>[]>;
    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     */
    readonly binding?: pulumi.Input<{ hostId?: pulumi.Input<string>, profile?: pulumi.Input<string>, vifDetails?: pulumi.Input<{[key: string]: any}>, vifType?: pulumi.Input<string>, vnicType?: pulumi.Input<string> }>;
    /**
     * Human-readable description of the floating IP. Changing
     * this updates the `description` of an existing port.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     */
    readonly deviceId?: pulumi.Input<string>;
    /**
     * The device owner of the Port. Changing this creates
     * a new port.
     */
    readonly deviceOwner?: pulumi.Input<string>;
    /**
     * The list of maps representing port DNS assignments.
     */
    readonly dnsAssignments?: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
    /**
     * The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     */
    readonly extraDhcpOptions?: pulumi.Input<pulumi.Input<{ ipVersion?: pulumi.Input<number>, name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     */
    readonly fixedIps?: pulumi.Input<pulumi.Input<{ ipAddress?: pulumi.Input<string>, subnetId: pulumi.Input<string> }>[]>;
    /**
     * Specify a specific MAC address for the port. Changing
     * this creates a new port.
     */
    readonly macAddress?: pulumi.Input<string>;
    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     */
    readonly networkId?: pulumi.Input<string>;
    /**
     * Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     */
    readonly noFixedIp?: pulumi.Input<boolean>;
    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the Port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    readonly noSecurityGroups?: pulumi.Input<boolean>;
    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of "true". Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     */
    readonly portSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of string tags for the port.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the Port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Port resource.
 */
export interface PortArgs {
    /**
     * Administrative up/down status for the port
     * (must be "true" or "false" if provided). Changing this updates the
     * `admin_state_up` of an existing port.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     */
    readonly allowedAddressPairs?: pulumi.Input<pulumi.Input<{ ipAddress: pulumi.Input<string>, macAddress?: pulumi.Input<string> }>[]>;
    /**
     * The port binding allows to specify binding information
     * for the port. The structure is described below.
     */
    readonly binding?: pulumi.Input<{ hostId?: pulumi.Input<string>, profile?: pulumi.Input<string>, vifDetails?: pulumi.Input<{[key: string]: any}>, vifType?: pulumi.Input<string>, vnicType?: pulumi.Input<string> }>;
    /**
     * Human-readable description of the floating IP. Changing
     * this updates the `description` of an existing port.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     */
    readonly deviceId?: pulumi.Input<string>;
    /**
     * The device owner of the Port. Changing this creates
     * a new port.
     */
    readonly deviceOwner?: pulumi.Input<string>;
    /**
     * The port DNS name. Available, when Neutron DNS extension
     * is enabled.
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * An extra DHCP option that needs to be configured
     * on the port. The structure is described below. Can be specified multiple
     * times.
     */
    readonly extraDhcpOptions?: pulumi.Input<pulumi.Input<{ ipVersion?: pulumi.Input<number>, name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
    /**
     * An array of desired IPs for
     * this port. The structure is described below.
     */
    readonly fixedIps?: pulumi.Input<pulumi.Input<{ ipAddress?: pulumi.Input<string>, subnetId: pulumi.Input<string> }>[]>;
    /**
     * Specify a specific MAC address for the port. Changing
     * this creates a new port.
     */
    readonly macAddress?: pulumi.Input<string>;
    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     */
    readonly networkId: pulumi.Input<string>;
    /**
     * Create a port with no fixed
     * IP address. This will also remove any fixed IPs previously set on a port. `true`
     * is the only valid value for this argument.
     */
    readonly noFixedIp?: pulumi.Input<boolean>;
    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the Port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    readonly noSecurityGroups?: pulumi.Input<boolean>;
    /**
     * Whether to explicitly enable or disable
     * port security on the port. Port Security is usually enabled by default, so
     * omitting argument will usually result in a value of "true". Setting this
     * explicitly to `false` will disable port security. In order to disable port
     * security, the port must not have any security groups. Valid values are `true`
     * and `false`.
     */
    readonly portSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A set of string tags for the port.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the Port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
