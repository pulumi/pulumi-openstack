import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 security group resource within OpenStack.
 *
 * Please note that managing security groups through the OpenStack Compute API
 * has been deprecated. Unless you are using an older OpenStack environment, it is
 * recommended to use the [`openstack_networking_secgroup_v2`](networking_secgroup_v2.html)
 * and [`openstack_networking_secgroup_rule_v2`](networking_secgroup_rule_v2.html)
 * resources instead, which uses the OpenStack Networking API.
 */
export declare class SecGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecGroupState): SecGroup;
    /**
     * A description for the security group. Changing this
     * updates the `description` of an existing security group.
     */
    readonly description: pulumi.Output<string>;
    /**
     * A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    readonly region: pulumi.Output<string>;
    /**
     * A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     */
    readonly rules: pulumi.Output<{
        cidr?: string;
        fromGroupId?: string;
        fromPort: number;
        id: string;
        ipProtocol: string;
        self?: boolean;
        toPort: number;
    }[]>;
    /**
     * Create a SecGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecGroupArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering SecGroup resources.
 */
export interface SecGroupState {
    /**
     * A description for the security group. Changing this
     * updates the `description` of an existing security group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     */
    readonly rules?: pulumi.Input<{
        cidr?: pulumi.Input<string>;
        fromGroupId?: pulumi.Input<string>;
        fromPort: pulumi.Input<number>;
        id?: pulumi.Input<string>;
        ipProtocol: pulumi.Input<string>;
        self?: pulumi.Input<boolean>;
        toPort: pulumi.Input<number>;
    }[]>;
}
/**
 * The set of arguments for constructing a SecGroup resource.
 */
export interface SecGroupArgs {
    /**
     * A description for the security group. Changing this
     * updates the `description` of an existing security group.
     */
    readonly description: pulumi.Input<string>;
    /**
     * A unique name for the security group. Changing this
     * updates the `name` of an existing security group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * A Compute client is needed to create a security group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A rule describing how the security group operates. The
     * rule object structure is documented below. Changing this updates the
     * security group rules. As shown in the example above, multiple rule blocks
     * may be used.
     */
    readonly rules?: pulumi.Input<{
        cidr?: pulumi.Input<string>;
        fromGroupId?: pulumi.Input<string>;
        fromPort: pulumi.Input<number>;
        id?: pulumi.Input<string>;
        ipProtocol: pulumi.Input<string>;
        self?: pulumi.Input<boolean>;
        toPort: pulumi.Input<number>;
    }[]>;
}
