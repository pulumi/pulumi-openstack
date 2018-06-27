import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 port resource within OpenStack.
 */
export declare class Port extends pulumi.CustomResource {
    /**
     * Get an existing Port resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortState): Port;
    /**
     * Administrative up/down status for the port
     * (must be "true" or "false" if provided). Changing this updates the
     * `admin_state_up` of an existing port.
     */
    readonly adminStateUp: pulumi.Output<boolean>;
    /**
     * The collection of Fixed IP addresses on the port in the
     * order returned by the Network v2 API.
     */
    readonly allFixedIps: pulumi.Output<string[]>;
    /**
     * The collection of Security Group IDs on the port
     * which have been explicitly and implicitly added.
     */
    readonly allSecurityGroupIds: pulumi.Output<string[]>;
    /**
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     */
    readonly allowedAddressPairs: pulumi.Output<{
        ipAddress: string;
        macAddress?: string;
    }[] | undefined>;
    /**
     * The ID of the device attached to the port. Changing this
     * creates a new port.
     */
    readonly deviceId: pulumi.Output<string>;
    /**
     * The device owner of the Port. Changing this creates
     * a new port.
     */
    readonly deviceOwner: pulumi.Output<string>;
    /**
     * An array of desired IPs for this port. The structure is
     * described below.
     */
    readonly fixedIps: pulumi.Output<{
        ipAddress?: string;
        subnetId: string;
    }[] | undefined>;
    /**
     * The additional MAC address.
     */
    readonly macAddress: pulumi.Output<string>;
    /**
     * A unique name for the port. Changing this
     * updates the `name` of an existing port.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The ID of the network to attach the port to. Changing
     * this creates a new port.
     */
    readonly networkId: pulumi.Output<string>;
    /**
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the Port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    readonly noSecurityGroups: pulumi.Output<boolean | undefined>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port.
     */
    readonly region: pulumi.Output<string>;
    /**
     * A list
     * of security group IDs to apply to the port. The security groups must be
     * specified by ID and not name (as opposed to how they are configured with
     * the Compute Instance).
     */
    readonly securityGroupIds: pulumi.Output<string[] | undefined>;
    /**
     * The owner of the Port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a Port resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortArgs, opts?: pulumi.ResourceOptions);
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
     * An IP/MAC Address pair of additional IP
     * addresses that can be active on this port. The structure is described
     * below.
     */
    readonly allowedAddressPairs?: pulumi.Input<{
        ipAddress: pulumi.Input<string>;
        macAddress?: pulumi.Input<string>;
    }[]>;
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
     * An array of desired IPs for this port. The structure is
     * described below.
     */
    readonly fixedIps?: pulumi.Input<{
        ipAddress?: pulumi.Input<string>;
        subnetId: pulumi.Input<string>;
    }[]>;
    /**
     * The additional MAC address.
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
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the Port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    readonly noSecurityGroups?: pulumi.Input<boolean>;
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
     * The owner of the Port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    readonly allowedAddressPairs?: pulumi.Input<{
        ipAddress: pulumi.Input<string>;
        macAddress?: pulumi.Input<string>;
    }[]>;
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
     * An array of desired IPs for this port. The structure is
     * described below.
     */
    readonly fixedIps?: pulumi.Input<{
        ipAddress?: pulumi.Input<string>;
        subnetId: pulumi.Input<string>;
    }[]>;
    /**
     * The additional MAC address.
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
     * If set to
     * `true`, then no security groups are applied to the port. If set to `false` and
     * no `security_group_ids` are specified, then the Port will yield to the default
     * behavior of the Networking service, which is to usually apply the "default"
     * security group.
     */
    readonly noSecurityGroups?: pulumi.Input<boolean>;
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
     * The owner of the Port. Required if admin wants
     * to create a port for another tenant. Changing this creates a new port.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
