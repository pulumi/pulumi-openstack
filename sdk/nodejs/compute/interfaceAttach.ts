// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches a Network Interface (a Port) to an Instance using the OpenStack
 * Compute (Nova) v2 API.
 *
 * ## Example Usage
 * ### Basic Attachment
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     adminStateUp: true,
 * });
 * const instance1 = new openstack.compute.Instance("instance_1", {
 *     securityGroups: ["default"],
 * });
 * const ai1 = new openstack.compute.InterfaceAttach("ai_1", {
 *     instanceId: instance1.id,
 *     networkId: openstack_networking_port_v2_network_1.id,
 * });
 * ```
 * ### Attachment Specifying a Fixed IP
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     adminStateUp: true,
 * });
 * const instance1 = new openstack.compute.Instance("instance_1", {
 *     securityGroups: ["default"],
 * });
 * const ai1 = new openstack.compute.InterfaceAttach("ai_1", {
 *     fixedIp: "10.0.10.10",
 *     instanceId: instance1.id,
 *     networkId: openstack_networking_port_v2_network_1.id,
 * });
 * ```
 * ### Attachment Using an Existing Port
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
 * const instance1 = new openstack.compute.Instance("instance_1", {
 *     securityGroups: ["default"],
 * });
 * const ai1 = new openstack.compute.InterfaceAttach("ai_1", {
 *     instanceId: instance1.id,
 *     portId: port1.id,
 * });
 * ```
 * ### Attaching Multiple Interfaces
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     adminStateUp: true,
 * });
 * const ports: openstack.networking.Port[] = [];
 * for (let i = 0; i < 2; i++) {
 *     ports.push(new openstack.networking.Port(`ports-${i}`, {
 *         adminStateUp: true,
 *         networkId: network1.id,
 *     }));
 * }
 * const instance1 = new openstack.compute.Instance("instance_1", {
 *     securityGroups: ["default"],
 * });
 * const attachments: openstack.compute.InterfaceAttach[] = [];
 * for (let i = 0; i < 2; i++) {
 *     attachments.push(new openstack.compute.InterfaceAttach(`attachments-${i}`, {
 *         instanceId: instance1.id,
 *         portId: pulumi.all(ports.map(v => v.id)).apply(id => id.map(v => v)[i]),
 *     }));
 * }
 * ```
 *
 * Note that the above example will not guarantee that the ports are attached in
 * a deterministic manner. The ports will be attached in a seemingly random
 * order.
 *
 * If you want to ensure that the ports are attached in a given order, create
 * explicit dependencies between the ports, such as:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     adminStateUp: true,
 * });
 * const ports: openstack.networking.Port[] = [];
 * for (let i = 0; i < 2; i++) {
 *     ports.push(new openstack.networking.Port(`ports-${i}`, {
 *         adminStateUp: true,
 *         networkId: network1.id,
 *     }));
 * }
 * const instance1 = new openstack.compute.Instance("instance_1", {
 *     securityGroups: ["default"],
 * });
 * const ai1 = new openstack.compute.InterfaceAttach("ai_1", {
 *     instanceId: instance1.id,
 *     portId: pulumi.all(ports.map(v => v.id)).apply(id => id.map(v => v)[0]),
 * });
 * const ai2 = new openstack.compute.InterfaceAttach("ai_2", {
 *     instanceId: instance1.id,
 *     portId: pulumi.all(ports.map(v => v.id)).apply(id => id.map(v => v)[1]),
 * });
 * ```
 */
export class InterfaceAttach extends pulumi.CustomResource {
    /**
     * Get an existing InterfaceAttach resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InterfaceAttachState, opts?: pulumi.CustomResourceOptions): InterfaceAttach {
        return new InterfaceAttach(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/interfaceAttach:InterfaceAttach';

    /**
     * Returns true if the given object is an instance of InterfaceAttach.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InterfaceAttach {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InterfaceAttach.__pulumiType;
    }

    /**
     * An IP address to assosciate with the port.
     * _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
     */
    public readonly fixedIp!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Instance to attach the Port or Network to.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The ID of the Network to attach to an Instance. A port will be created automatically.
     * _NOTE_: This option and `portId` are mutually exclusive.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The ID of the Port to attach to an Instance.
     * _NOTE_: This option and `networkId` are mutually exclusive.
     */
    public readonly portId!: pulumi.Output<string>;
    /**
     * The region in which to create the interface attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new attachment.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a InterfaceAttach resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InterfaceAttachArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InterfaceAttachArgs | InterfaceAttachState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InterfaceAttachState | undefined;
            inputs["fixedIp"] = state ? state.fixedIp : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["networkId"] = state ? state.networkId : undefined;
            inputs["portId"] = state ? state.portId : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as InterfaceAttachArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["fixedIp"] = args ? args.fixedIp : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["networkId"] = args ? args.networkId : undefined;
            inputs["portId"] = args ? args.portId : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InterfaceAttach.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InterfaceAttach resources.
 */
export interface InterfaceAttachState {
    /**
     * An IP address to assosciate with the port.
     * _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
     */
    readonly fixedIp?: pulumi.Input<string>;
    /**
     * The ID of the Instance to attach the Port or Network to.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The ID of the Network to attach to an Instance. A port will be created automatically.
     * _NOTE_: This option and `portId` are mutually exclusive.
     */
    readonly networkId?: pulumi.Input<string>;
    /**
     * The ID of the Port to attach to an Instance.
     * _NOTE_: This option and `networkId` are mutually exclusive.
     */
    readonly portId?: pulumi.Input<string>;
    /**
     * The region in which to create the interface attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new attachment.
     */
    readonly region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InterfaceAttach resource.
 */
export interface InterfaceAttachArgs {
    /**
     * An IP address to assosciate with the port.
     * _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
     */
    readonly fixedIp?: pulumi.Input<string>;
    /**
     * The ID of the Instance to attach the Port or Network to.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * The ID of the Network to attach to an Instance. A port will be created automatically.
     * _NOTE_: This option and `portId` are mutually exclusive.
     */
    readonly networkId?: pulumi.Input<string>;
    /**
     * The ID of the Port to attach to an Instance.
     * _NOTE_: This option and `networkId` are mutually exclusive.
     */
    readonly portId?: pulumi.Input<string>;
    /**
     * The region in which to create the interface attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new attachment.
     */
    readonly region?: pulumi.Input<string>;
}
