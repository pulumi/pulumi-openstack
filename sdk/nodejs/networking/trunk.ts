// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a networking V2 trunk resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const network1 = new openstack.networking.Network("network_1", {
 *     adminStateUp: true,
 * });
 * const subnet1 = new openstack.networking.Subnet("subnet_1", {
 *     cidr: "192.168.1.0/24",
 *     enableDhcp: true,
 *     ipVersion: 4,
 *     networkId: network1.id,
 *     noGateway: true,
 * });
 * const parentPort1 = new openstack.networking.Port("parent_port_1", {
 *     adminStateUp: true,
 *     networkId: network1.id,
 * }, {dependsOn: [subnet1]});
 * const subport1 = new openstack.networking.Port("subport_1", {
 *     adminStateUp: true,
 *     networkId: network1.id,
 * }, {dependsOn: [subnet1]});
 * const trunk1 = new openstack.networking.Trunk("trunk_1", {
 *     adminStateUp: true,
 *     portId: parentPort1.id,
 *     subPorts: [{
 *         portId: subport1.id,
 *         segmentationId: 1,
 *         segmentationType: "vlan",
 *     }],
 * });
 * const instance1 = new openstack.compute.Instance("instance_1", {
 *     networks: [{
 *         port: trunk1.portId,
 *     }],
 *     securityGroups: ["default"],
 * });
 * ```
 */
export class Trunk extends pulumi.CustomResource {
    /**
     * Get an existing Trunk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrunkState, opts?: pulumi.CustomResourceOptions): Trunk {
        return new Trunk(name, <any>state, { ...opts, id: id });
    }

    /**
     * Administrative up/down status for the trunk
     * (must be "true" or "false" if provided). Changing this updates the
     * `admin_state_up` of an existing trunk.
     */
    public readonly adminStateUp: pulumi.Output<boolean | undefined>;
    /**
     * The collection of tags assigned on the trunk, which have been
     * explicitly and implicitly added.
     */
    public /*out*/ readonly allTags: pulumi.Output<string[]>;
    /**
     * Human-readable description of the trunk. Changing this
     * updates the name of the existing trunk.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * A unique name for the trunk. Changing this
     * updates the `name` of an existing trunk.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the port to be used as the parent port of the
     * trunk. This is the port that should be used as the compute instance network
     * port. Changing this creates a new trunk.
     */
    public readonly portId: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a trunk. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * trunk.
     */
    public readonly region: pulumi.Output<string>;
    /**
     * The set of ports that will be made subports of the trunk.
     * The structure of each subport is described below.
     */
    public readonly subPorts: pulumi.Output<{ portId: string, segmentationId: number, segmentationType: string }[] | undefined>;
    /**
     * A set of string tags for the port.
     */
    public readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * The owner of the Trunk. Required if admin wants
     * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
     */
    public readonly tenantId: pulumi.Output<string>;

    /**
     * Create a Trunk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrunkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrunkArgs | TrunkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TrunkState = argsOrState as TrunkState | undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["allTags"] = state ? state.allTags : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["portId"] = state ? state.portId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["subPorts"] = state ? state.subPorts : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as TrunkArgs | undefined;
            if (!args || args.portId === undefined) {
                throw new Error("Missing required property 'portId'");
            }
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["portId"] = args ? args.portId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["subPorts"] = args ? args.subPorts : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["allTags"] = undefined /*out*/;
        }
        super("openstack:networking/trunk:Trunk", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trunk resources.
 */
export interface TrunkState {
    /**
     * Administrative up/down status for the trunk
     * (must be "true" or "false" if provided). Changing this updates the
     * `admin_state_up` of an existing trunk.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The collection of tags assigned on the trunk, which have been
     * explicitly and implicitly added.
     */
    readonly allTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable description of the trunk. Changing this
     * updates the name of the existing trunk.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique name for the trunk. Changing this
     * updates the `name` of an existing trunk.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the port to be used as the parent port of the
     * trunk. This is the port that should be used as the compute instance network
     * port. Changing this creates a new trunk.
     */
    readonly portId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a trunk. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * trunk.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The set of ports that will be made subports of the trunk.
     * The structure of each subport is described below.
     */
    readonly subPorts?: pulumi.Input<pulumi.Input<{ portId: pulumi.Input<string>, segmentationId: pulumi.Input<number>, segmentationType: pulumi.Input<string> }>[]>;
    /**
     * A set of string tags for the port.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the Trunk. Required if admin wants
     * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trunk resource.
 */
export interface TrunkArgs {
    /**
     * Administrative up/down status for the trunk
     * (must be "true" or "false" if provided). Changing this updates the
     * `admin_state_up` of an existing trunk.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description of the trunk. Changing this
     * updates the name of the existing trunk.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique name for the trunk. Changing this
     * updates the `name` of an existing trunk.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the port to be used as the parent port of the
     * trunk. This is the port that should be used as the compute instance network
     * port. Changing this creates a new trunk.
     */
    readonly portId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a trunk. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * trunk.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The set of ports that will be made subports of the trunk.
     * The structure of each subport is described below.
     */
    readonly subPorts?: pulumi.Input<pulumi.Input<{ portId: pulumi.Input<string>, segmentationId: pulumi.Input<number>, segmentationType: pulumi.Input<string> }>[]>;
    /**
     * A set of string tags for the port.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the Trunk. Required if admin wants
     * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
     */
    readonly tenantId?: pulumi.Input<string>;
}
