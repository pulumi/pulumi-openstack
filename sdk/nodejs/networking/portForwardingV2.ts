// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 portforwarding resource within OpenStack.
 *
 * ## Example Usage
 *
 * ### Simple portforwarding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const pf1 = new openstack.networking.PortForwardingV2("pf1", {
 *     externalPort: 7233,
 *     floatingipId: "7a52eb59-7d47-415d-a884-046666a6fbae",
 *     internalPort: 25,
 *     internalPortId: "b930d7f6-ceb7-40a0-8b81-a425dd994ccf",
 *     protocol: "tcp",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class PortForwardingV2 extends pulumi.CustomResource {
    /**
     * Get an existing PortForwardingV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortForwardingV2State, opts?: pulumi.CustomResourceOptions): PortForwardingV2 {
        return new PortForwardingV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/portForwardingV2:PortForwardingV2';

    /**
     * Returns true if the given object is an instance of PortForwardingV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PortForwardingV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PortForwardingV2.__pulumiType;
    }

    /**
     * A text describing the port forwarding. Changing this
     * updates the `description` of an existing port forwarding.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The TCP/UDP/other protocol port number of the port forwarding. Changing this
     * updates the `externalPort` of an existing port forwarding.
     */
    public readonly externalPort!: pulumi.Output<number>;
    /**
     * The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
     */
    public readonly floatingipId!: pulumi.Output<string>;
    /**
     * The fixed IPv4 address of the Neutron port associated with the port forwarding.
     * Changing this updates the `internalIpAddress` of an existing port forwarding.
     */
    public readonly internalIpAddress!: pulumi.Output<string>;
    /**
     * The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
     * port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
     */
    public readonly internalPort!: pulumi.Output<number>;
    /**
     * The ID of the Neutron port associated with the port forwarding. Changing
     * this updates the `internalPortId` of an existing port forwarding.
     */
    public readonly internalPortId!: pulumi.Output<string>;
    /**
     * The IP protocol used in the port forwarding. Changing this updates the `protocol`
     * of an existing port forwarding.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port forwarding. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port forwarding.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a PortForwardingV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortForwardingV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortForwardingV2Args | PortForwardingV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortForwardingV2State | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalPort"] = state ? state.externalPort : undefined;
            resourceInputs["floatingipId"] = state ? state.floatingipId : undefined;
            resourceInputs["internalIpAddress"] = state ? state.internalIpAddress : undefined;
            resourceInputs["internalPort"] = state ? state.internalPort : undefined;
            resourceInputs["internalPortId"] = state ? state.internalPortId : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as PortForwardingV2Args | undefined;
            if ((!args || args.externalPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalPort'");
            }
            if ((!args || args.floatingipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'floatingipId'");
            }
            if ((!args || args.internalIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internalIpAddress'");
            }
            if ((!args || args.internalPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internalPort'");
            }
            if ((!args || args.internalPortId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internalPortId'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["externalPort"] = args ? args.externalPort : undefined;
            resourceInputs["floatingipId"] = args ? args.floatingipId : undefined;
            resourceInputs["internalIpAddress"] = args ? args.internalIpAddress : undefined;
            resourceInputs["internalPort"] = args ? args.internalPort : undefined;
            resourceInputs["internalPortId"] = args ? args.internalPortId : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PortForwardingV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PortForwardingV2 resources.
 */
export interface PortForwardingV2State {
    /**
     * A text describing the port forwarding. Changing this
     * updates the `description` of an existing port forwarding.
     */
    description?: pulumi.Input<string>;
    /**
     * The TCP/UDP/other protocol port number of the port forwarding. Changing this
     * updates the `externalPort` of an existing port forwarding.
     */
    externalPort?: pulumi.Input<number>;
    /**
     * The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
     */
    floatingipId?: pulumi.Input<string>;
    /**
     * The fixed IPv4 address of the Neutron port associated with the port forwarding.
     * Changing this updates the `internalIpAddress` of an existing port forwarding.
     */
    internalIpAddress?: pulumi.Input<string>;
    /**
     * The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
     * port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
     */
    internalPort?: pulumi.Input<number>;
    /**
     * The ID of the Neutron port associated with the port forwarding. Changing
     * this updates the `internalPortId` of an existing port forwarding.
     */
    internalPortId?: pulumi.Input<string>;
    /**
     * The IP protocol used in the port forwarding. Changing this updates the `protocol`
     * of an existing port forwarding.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port forwarding. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port forwarding.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PortForwardingV2 resource.
 */
export interface PortForwardingV2Args {
    /**
     * A text describing the port forwarding. Changing this
     * updates the `description` of an existing port forwarding.
     */
    description?: pulumi.Input<string>;
    /**
     * The TCP/UDP/other protocol port number of the port forwarding. Changing this
     * updates the `externalPort` of an existing port forwarding.
     */
    externalPort: pulumi.Input<number>;
    /**
     * The ID of the Neutron floating IP address. Changing this creates a new port forwarding.
     */
    floatingipId: pulumi.Input<string>;
    /**
     * The fixed IPv4 address of the Neutron port associated with the port forwarding.
     * Changing this updates the `internalIpAddress` of an existing port forwarding.
     */
    internalIpAddress: pulumi.Input<string>;
    /**
     * The TCP/UDP/other protocol port number of the Neutron port fixed IP address associated to the
     * port forwarding. Changing this updates the `internalPort` of an existing port forwarding.
     */
    internalPort: pulumi.Input<number>;
    /**
     * The ID of the Neutron port associated with the port forwarding. Changing
     * this updates the `internalPortId` of an existing port forwarding.
     */
    internalPortId: pulumi.Input<string>;
    /**
     * The IP protocol used in the port forwarding. Changing this updates the `protocol`
     * of an existing port forwarding.
     */
    protocol: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port forwarding. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * port forwarding.
     */
    region?: pulumi.Input<string>;
}
