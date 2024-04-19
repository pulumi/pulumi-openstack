// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 Server Group resource within OpenStack.
 *
 * ## Example Usage
 *
 * ### Compute service API version 2.63 or below:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const test_sg = new openstack.compute.ServerGroup("test-sg", {policies: ["anti-affinity"]});
 * ```
 *
 * ### Compute service API version 2.64 or above:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const test_sg = new openstack.compute.ServerGroup("test-sg", {
 *     policies: ["anti-affinity"],
 *     rules: {
 *         maxServerPerHost: 3,
 *     },
 * });
 * ```
 *
 * ## Policies
 *
 * * `affinity` - All instances/servers launched in this group will be hosted on
 *   the same compute node.
 *
 * * `anti-affinity` - All instances/servers launched in this group will be
 *   hosted on different compute nodes.
 *
 * * `soft-affinity` - All instances/servers launched in this group will be hosted
 *   on the same compute node if possible, but if not possible they
 *   still will be scheduled instead of failure. To use this policy your
 *   OpenStack environment should support Compute service API 2.15 or above.
 *
 * * `soft-anti-affinity` - All instances/servers launched in this group will be
 *   hosted on different compute nodes if possible, but if not possible they
 *   still will be scheduled instead of failure. To use this policy your
 *   OpenStack environment should support Compute service API 2.15 or above.
 *
 * ## Import
 *
 * Server Groups can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:compute/serverGroup:ServerGroup test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
 * ```
 */
export class ServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerGroupState, opts?: pulumi.CustomResourceOptions): ServerGroup {
        return new ServerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/serverGroup:ServerGroup';

    /**
     * Returns true if the given object is an instance of ServerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerGroup.__pulumiType;
    }

    /**
     * The instances that are part of this server group.
     */
    public /*out*/ readonly members!: pulumi.Output<string[]>;
    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of exactly one policy name to associate with
     * the server group. See the Policies section for more information. Changing this
     * creates a new server group.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The rules which are applied to specified `policy`. Currently,
     * only the `maxServerPerHost` rule is supported for the `anti-affinity` policy.
     */
    public readonly rules!: pulumi.Output<outputs.compute.ServerGroupRules>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a ServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerGroupArgs | ServerGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerGroupState | undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as ServerGroupArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["members"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerGroup resources.
 */
export interface ServerGroupState {
    /**
     * The instances that are part of this server group.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of exactly one policy name to associate with
     * the server group. See the Policies section for more information. Changing this
     * creates a new server group.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     */
    region?: pulumi.Input<string>;
    /**
     * The rules which are applied to specified `policy`. Currently,
     * only the `maxServerPerHost` rule is supported for the `anti-affinity` policy.
     */
    rules?: pulumi.Input<inputs.compute.ServerGroupRules>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ServerGroup resource.
 */
export interface ServerGroupArgs {
    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of exactly one policy name to associate with
     * the server group. See the Policies section for more information. Changing this
     * creates a new server group.
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     */
    region?: pulumi.Input<string>;
    /**
     * The rules which are applied to specified `policy`. Currently,
     * only the `maxServerPerHost` rule is supported for the `anti-affinity` policy.
     */
    rules?: pulumi.Input<inputs.compute.ServerGroupRules>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
