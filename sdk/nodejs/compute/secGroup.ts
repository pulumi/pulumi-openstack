// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 security group resource within OpenStack.
 *
 * Please note that managing security groups through the OpenStack Compute API
 * has been deprecated. Unless you are using an older OpenStack environment, it is
 * recommended to use the `openstack.networking.SecGroup`
 * and `openstack.networking.SecGroupRule`
 * resources instead, which uses the OpenStack Networking API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const secgroup1 = new openstack.compute.SecGroup("secgroup1", {
 *     description: "my security group",
 *     rules: [
 *         {
 *             cidr: "0.0.0.0/0",
 *             fromPort: 22,
 *             ipProtocol: "tcp",
 *             toPort: 22,
 *         },
 *         {
 *             cidr: "0.0.0.0/0",
 *             fromPort: 80,
 *             ipProtocol: "tcp",
 *             toPort: 80,
 *         },
 *     ],
 * });
 * ```
 *
 * ## Notes
 *
 * ### ICMP Rules
 *
 * When using ICMP as the `ipProtocol`, the `fromPort` sets the ICMP _type_ and the `toPort` sets the ICMP _code_. To allow all ICMP types, set each value to `-1`, like so:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages).
 *
 * ### Referencing Security Groups
 *
 * When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const test_server = new openstack.compute.Instance("test-server", {
 *     imageId: "ad091b52-742f-469e-8f3c-fd81cadf0743",
 *     flavorId: "3",
 *     keyPair: "my_key_pair_name",
 *     securityGroups: [openstack_compute_secgroup_v2.secgroup_1.name],
 * });
 * ```
 *
 * ## Import
 *
 * Security Groups can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:compute/secGroup:SecGroup my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
 * ```
 */
export class SecGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecGroupState, opts?: pulumi.CustomResourceOptions): SecGroup {
        return new SecGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/secGroup:SecGroup';

    /**
     * Returns true if the given object is an instance of SecGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecGroup.__pulumiType;
    }

    /**
     * A description for the security group. Changing this
     * updates the `description` of an existing security group.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     */
    public readonly rules!: pulumi.Output<outputs.compute.SecGroupRule[]>;

    /**
     * Create a SecGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecGroupArgs | SecGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as SecGroupArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecGroup resources.
 */
export interface SecGroupState {
    /**
     * A description for the security group. Changing this
     * updates the `description` of an existing security group.
     */
    description?: pulumi.Input<string>;
    /**
     * A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    region?: pulumi.Input<string>;
    /**
     * A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.SecGroupRule>[]>;
}

/**
 * The set of arguments for constructing a SecGroup resource.
 */
export interface SecGroupArgs {
    /**
     * A description for the security group. Changing this
     * updates the `description` of an existing security group.
     */
    description: pulumi.Input<string>;
    /**
     * A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    region?: pulumi.Input<string>;
    /**
     * A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.SecGroupRule>[]>;
}
