// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a V2 members resource within OpenStack (batch members update).
 *
 * > **Note:** This resource has attributes that depend on octavia minor versions.
 * Please ensure your Openstack cloud supports the required minor version.
 *
 * > **Note:** This resource works only within Octavia API. For
 * legacy Neutron LBaaS v2 extension please use
 * openstack.loadbalancer.Member resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const members1 = new openstack.loadbalancer.Members("members_1", {
 *     members: [
 *         {
 *             address: "192.168.199.23",
 *             protocolPort: 8080,
 *         },
 *         {
 *             address: "192.168.199.24",
 *             protocolPort: 8080,
 *         },
 *     ],
 *     poolId: "935685fb-a896-40f9-9ff4-ae531a3a00fe",
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancer Pool Members can be imported using the Pool ID, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:loadbalancer/members:Members members_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5
 * ```
 */
export class Members extends pulumi.CustomResource {
    /**
     * Get an existing Members resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MembersState, opts?: pulumi.CustomResourceOptions): Members {
        return new Members(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/members:Members';

    /**
     * Returns true if the given object is an instance of Members.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Members {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Members.__pulumiType;
    }

    /**
     * A set of dictionaries containing member parameters. The
     * structure is described below.
     */
    public readonly members!: pulumi.Output<outputs.loadbalancer.MembersMember[] | undefined>;
    /**
     * The id of the pool that members will be assigned to.
     * Changing this creates a new members resource.
     */
    public readonly poolId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create pool members. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * members resource.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Members resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MembersArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MembersArgs | MembersState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MembersState | undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["poolId"] = state ? state.poolId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as MembersArgs | undefined;
            if ((!args || args.poolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolId'");
            }
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["poolId"] = args ? args.poolId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Members.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Members resources.
 */
export interface MembersState {
    /**
     * A set of dictionaries containing member parameters. The
     * structure is described below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.loadbalancer.MembersMember>[]>;
    /**
     * The id of the pool that members will be assigned to.
     * Changing this creates a new members resource.
     */
    poolId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create pool members. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * members resource.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Members resource.
 */
export interface MembersArgs {
    /**
     * A set of dictionaries containing member parameters. The
     * structure is described below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.loadbalancer.MembersMember>[]>;
    /**
     * The id of the pool that members will be assigned to.
     * Changing this creates a new members resource.
     */
    poolId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create pool members. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * members resource.
     */
    region?: pulumi.Input<string>;
}
