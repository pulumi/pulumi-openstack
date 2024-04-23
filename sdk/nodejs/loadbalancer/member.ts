// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 member resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const member1 = new openstack.loadbalancer.Member("member_1", {
 *     poolId: "935685fb-a896-40f9-9ff4-ae531a3a00fe",
 *     address: "192.168.199.23",
 *     protocolPort: 8080,
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancer Pool Member can be imported using the Pool ID and Member ID
 * separated by a slash, e.g.:
 *
 * ```sh
 * $ pulumi import openstack:loadbalancer/member:Member member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f
 * ```
 */
export class Member extends pulumi.CustomResource {
    /**
     * Get an existing Member resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberState, opts?: pulumi.CustomResourceOptions): Member {
        return new Member(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/member:Member';

    /**
     * Returns true if the given object is an instance of Member.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Member {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Member.__pulumiType;
    }

    /**
     * The IP address of the member to receive traffic from
     * the load balancer. Changing this creates a new member.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean that indicates whether that member works as a backup or not. Available 
     * only for Octavia >= 2.1.
     */
    public readonly backup!: pulumi.Output<boolean | undefined>;
    /**
     * An alternate IP address used for health monitoring a backend member.
     * Available only for Octavia
     */
    public readonly monitorAddress!: pulumi.Output<string | undefined>;
    /**
     * An alternate protocol port used for health monitoring a backend member.
     * Available only for Octavia
     */
    public readonly monitorPort!: pulumi.Output<number | undefined>;
    /**
     * Human-readable name for the member.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the pool that this member will be assigned
     * to. Changing this creates a new member.
     */
    public readonly poolId!: pulumi.Output<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new member.
     */
    public readonly protocolPort!: pulumi.Output<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a member. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new member.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The subnet in which to access the member. Changing
     * this creates a new member.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the member.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new member.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * A positive integer value that indicates the relative
     * portion of traffic that this member should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     */
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a Member resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MemberArgs | MemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MemberState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["backup"] = state ? state.backup : undefined;
            resourceInputs["monitorAddress"] = state ? state.monitorAddress : undefined;
            resourceInputs["monitorPort"] = state ? state.monitorPort : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["poolId"] = state ? state.poolId : undefined;
            resourceInputs["protocolPort"] = state ? state.protocolPort : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as MemberArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.poolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolId'");
            }
            if ((!args || args.protocolPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocolPort'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["backup"] = args ? args.backup : undefined;
            resourceInputs["monitorAddress"] = args ? args.monitorAddress : undefined;
            resourceInputs["monitorPort"] = args ? args.monitorPort : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["poolId"] = args ? args.poolId : undefined;
            resourceInputs["protocolPort"] = args ? args.protocolPort : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["weight"] = args ? args.weight : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Member.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Member resources.
 */
export interface MemberState {
    /**
     * The IP address of the member to receive traffic from
     * the load balancer. Changing this creates a new member.
     */
    address?: pulumi.Input<string>;
    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Boolean that indicates whether that member works as a backup or not. Available 
     * only for Octavia >= 2.1.
     */
    backup?: pulumi.Input<boolean>;
    /**
     * An alternate IP address used for health monitoring a backend member.
     * Available only for Octavia
     */
    monitorAddress?: pulumi.Input<string>;
    /**
     * An alternate protocol port used for health monitoring a backend member.
     * Available only for Octavia
     */
    monitorPort?: pulumi.Input<number>;
    /**
     * Human-readable name for the member.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the pool that this member will be assigned
     * to. Changing this creates a new member.
     */
    poolId?: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new member.
     */
    protocolPort?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a member. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new member.
     */
    region?: pulumi.Input<string>;
    /**
     * The subnet in which to access the member. Changing
     * this creates a new member.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the member.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new member.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * A positive integer value that indicates the relative
     * portion of traffic that this member should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     */
    weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Member resource.
 */
export interface MemberArgs {
    /**
     * The IP address of the member to receive traffic from
     * the load balancer. Changing this creates a new member.
     */
    address: pulumi.Input<string>;
    /**
     * The administrative state of the member.
     * A valid value is true (UP) or false (DOWN). Defaults to true.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Boolean that indicates whether that member works as a backup or not. Available 
     * only for Octavia >= 2.1.
     */
    backup?: pulumi.Input<boolean>;
    /**
     * An alternate IP address used for health monitoring a backend member.
     * Available only for Octavia
     */
    monitorAddress?: pulumi.Input<string>;
    /**
     * An alternate protocol port used for health monitoring a backend member.
     * Available only for Octavia
     */
    monitorPort?: pulumi.Input<number>;
    /**
     * Human-readable name for the member.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the pool that this member will be assigned
     * to. Changing this creates a new member.
     */
    poolId: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new member.
     */
    protocolPort: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a member. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new member.
     */
    region?: pulumi.Input<string>;
    /**
     * The subnet in which to access the member. Changing
     * this creates a new member.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the member.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new member.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * A positive integer value that indicates the relative
     * portion of traffic that this member should receive from the pool. For
     * example, a member with a weight of 10 receives five times as much traffic
     * as a member with a weight of 2. Defaults to 1.
     */
    weight?: pulumi.Input<number>;
}
