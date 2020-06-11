// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V1 Barbican order resource within OpenStack.
 *
 * ## Example Usage
 *
 * ### Symmetric key order
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const order1 = new openstack.keymanager.OrderV1("order1", {
 *     meta: {
 *         algorithm: "aes",
 *         bitLength: 256,
 *         mode: "cbc",
 *         name: "mysecret",
 *     },
 *     type: "key",
 * });
 * ```
 *
 * ### Asymmetric key pair order
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const order1 = new openstack.keymanager.OrderV1("order1", {
 *     meta: {
 *         algorithm: "rsa",
 *         bitLength: 4096,
 *         name: "mysecret",
 *     },
 *     type: "asymmetric",
 * });
 * ```
 */
export class OrderV1 extends pulumi.CustomResource {
    /**
     * Get an existing OrderV1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrderV1State, opts?: pulumi.CustomResourceOptions): OrderV1 {
        return new OrderV1(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:keymanager/orderV1:OrderV1';

    /**
     * Returns true if the given object is an instance of OrderV1.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrderV1 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrderV1.__pulumiType;
    }

    /**
     * The container reference / where to find the container.
     */
    public /*out*/ readonly containerRef!: pulumi.Output<string>;
    /**
     * The date the order was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * The creator of the order.
     */
    public /*out*/ readonly creatorId!: pulumi.Output<string>;
    /**
     * Dictionary containing the order metadata used to generate the order. The structure is described below.
     */
    public readonly meta!: pulumi.Output<outputs.keymanager.OrderV1Meta>;
    /**
     * The order reference / where to find the order.
     */
    public /*out*/ readonly orderRef!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a order. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 order.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret reference / where to find the secret.
     */
    public /*out*/ readonly secretRef!: pulumi.Output<string>;
    /**
     * The status of the order.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The sub status of the order.
     */
    public /*out*/ readonly subStatus!: pulumi.Output<string>;
    /**
     * The sub status message of the order.
     */
    public /*out*/ readonly subStatusMessage!: pulumi.Output<string>;
    /**
     * The type of key to be generated. Must be one of `asymmetric`, `key`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The date the order was last updated.
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;

    /**
     * Create a OrderV1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrderV1Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrderV1Args | OrderV1State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as OrderV1State | undefined;
            inputs["containerRef"] = state ? state.containerRef : undefined;
            inputs["created"] = state ? state.created : undefined;
            inputs["creatorId"] = state ? state.creatorId : undefined;
            inputs["meta"] = state ? state.meta : undefined;
            inputs["orderRef"] = state ? state.orderRef : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["secretRef"] = state ? state.secretRef : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subStatus"] = state ? state.subStatus : undefined;
            inputs["subStatusMessage"] = state ? state.subStatusMessage : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["updated"] = state ? state.updated : undefined;
        } else {
            const args = argsOrState as OrderV1Args | undefined;
            if (!args || args.meta === undefined) {
                throw new Error("Missing required property 'meta'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["meta"] = args ? args.meta : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["containerRef"] = undefined /*out*/;
            inputs["created"] = undefined /*out*/;
            inputs["creatorId"] = undefined /*out*/;
            inputs["orderRef"] = undefined /*out*/;
            inputs["secretRef"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["subStatus"] = undefined /*out*/;
            inputs["subStatusMessage"] = undefined /*out*/;
            inputs["updated"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(OrderV1.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrderV1 resources.
 */
export interface OrderV1State {
    /**
     * The container reference / where to find the container.
     */
    readonly containerRef?: pulumi.Input<string>;
    /**
     * The date the order was created.
     */
    readonly created?: pulumi.Input<string>;
    /**
     * The creator of the order.
     */
    readonly creatorId?: pulumi.Input<string>;
    /**
     * Dictionary containing the order metadata used to generate the order. The structure is described below.
     */
    readonly meta?: pulumi.Input<inputs.keymanager.OrderV1Meta>;
    /**
     * The order reference / where to find the order.
     */
    readonly orderRef?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a order. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 order.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The secret reference / where to find the secret.
     */
    readonly secretRef?: pulumi.Input<string>;
    /**
     * The status of the order.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The sub status of the order.
     */
    readonly subStatus?: pulumi.Input<string>;
    /**
     * The sub status message of the order.
     */
    readonly subStatusMessage?: pulumi.Input<string>;
    /**
     * The type of key to be generated. Must be one of `asymmetric`, `key`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The date the order was last updated.
     */
    readonly updated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrderV1 resource.
 */
export interface OrderV1Args {
    /**
     * Dictionary containing the order metadata used to generate the order. The structure is described below.
     */
    readonly meta: pulumi.Input<inputs.keymanager.OrderV1Meta>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a order. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 order.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The type of key to be generated. Must be one of `asymmetric`, `key`.
     */
    readonly type: pulumi.Input<string>;
}