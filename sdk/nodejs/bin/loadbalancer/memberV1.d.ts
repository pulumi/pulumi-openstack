import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 load balancer member resource within OpenStack.
 */
export declare class MemberV1 extends pulumi.CustomResource {
    /**
     * Get an existing MemberV1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MemberV1State): MemberV1;
    /**
     * The IP address of the member. Changing this creates a
     * new member.
     */
    readonly address: pulumi.Output<string>;
    /**
     * The administrative state of the member.
     * Acceptable values are 'true' and 'false'. Changing this value updates the
     * state of the existing member.
     */
    readonly adminStateUp: pulumi.Output<boolean>;
    /**
     * The ID of the LB pool. Changing this creates a new
     * member.
     */
    readonly poolId: pulumi.Output<string>;
    /**
     * An integer representing the port on which the member is
     * hosted. Changing this creates a new member.
     */
    readonly port: pulumi.Output<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The owner of the member. Required if admin wants to
     * create a member for another tenant. Changing this creates a new member.
     */
    readonly tenantId: pulumi.Output<string | undefined>;
    /**
     * The load balancing weight of the member. This is currently unable
     * to be set through Terraform.
     */
    readonly weight: pulumi.Output<number>;
    /**
     * Create a MemberV1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MemberV1Args, opts?: pulumi.ResourceOptions);
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
    /**
     * The load balancing weight of the member. This is currently unable
     * to be set through Terraform.
     */
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
    /**
     * The load balancing weight of the member. This is currently unable
     * to be set through Terraform.
     */
    readonly weight?: pulumi.Input<number>;
}
