// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 load balancer member resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const member1 = new openstack.loadbalancer.MemberV1("member1", {
 *     address: "192.168.0.10",
 *     poolId: "d9415786-5f1a-428b-b35f-2f1523e146d2",
 *     port: 80,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_member_v1.html.markdown.
 */
export class MemberV1 extends pulumi.CustomResource {
    /**
     * Get an existing MemberV1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberV1State, opts?: pulumi.CustomResourceOptions): MemberV1 {
        return new MemberV1(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/memberV1:MemberV1';

    /**
     * Returns true if the given object is an instance of MemberV1.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MemberV1 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MemberV1.__pulumiType;
    }

    /**
     * The IP address of the member. Changing this creates a
     * new member.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The administrative state of the member.
     * Acceptable values are 'true' and 'false'. Changing this value updates the
     * state of the existing member.
     */
    public readonly adminStateUp!: pulumi.Output<boolean>;
    /**
     * The ID of the LB pool. Changing this creates a new
     * member.
     */
    public readonly poolId!: pulumi.Output<string>;
    /**
     * An integer representing the port on which the member is
     * hosted. Changing this creates a new member.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The owner of the member. Required if admin wants to
     * create a member for another tenant. Changing this creates a new member.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    public readonly weight!: pulumi.Output<number>;

    /**
     * Create a MemberV1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberV1Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MemberV1Args | MemberV1State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MemberV1State | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["poolId"] = state ? state.poolId : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["weight"] = state ? state.weight : undefined;
        } else {
            const args = argsOrState as MemberV1Args | undefined;
            if (!args || args.address === undefined) {
                throw new Error("Missing required property 'address'");
            }
            if (!args || args.poolId === undefined) {
                throw new Error("Missing required property 'poolId'");
            }
            if (!args || args.port === undefined) {
                throw new Error("Missing required property 'port'");
            }
            inputs["address"] = args ? args.address : undefined;
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["poolId"] = args ? args.poolId : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["weight"] = args ? args.weight : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MemberV1.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MemberV1 resources.
 */
export interface MemberV1State {
    /**
     * The IP address of the member. Changing this creates a
     * new member.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The administrative state of the member.
     * Acceptable values are 'true' and 'false'. Changing this value updates the
     * state of the existing member.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The ID of the LB pool. Changing this creates a new
     * member.
     */
    readonly poolId?: pulumi.Input<string>;
    /**
     * An integer representing the port on which the member is
     * hosted. Changing this creates a new member.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the member. Required if admin wants to
     * create a member for another tenant. Changing this creates a new member.
     */
    readonly tenantId?: pulumi.Input<string>;
    readonly weight?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a MemberV1 resource.
 */
export interface MemberV1Args {
    /**
     * The IP address of the member. Changing this creates a
     * new member.
     */
    readonly address: pulumi.Input<string>;
    /**
     * The administrative state of the member.
     * Acceptable values are 'true' and 'false'. Changing this value updates the
     * state of the existing member.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The ID of the LB pool. Changing this creates a new
     * member.
     */
    readonly poolId: pulumi.Input<string>;
    /**
     * An integer representing the port on which the member is
     * hosted. Changing this creates a new member.
     */
    readonly port: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the member. Required if admin wants to
     * create a member for another tenant. Changing this creates a new member.
     */
    readonly tenantId?: pulumi.Input<string>;
    readonly weight?: pulumi.Input<number>;
}
