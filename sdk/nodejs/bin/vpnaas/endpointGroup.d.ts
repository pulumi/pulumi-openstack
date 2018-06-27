import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 Neutron Endpoint Group resource within OpenStack.
 */
export declare class EndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing EndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointGroupState): EndpointGroup;
    /**
     * The human-readable description for the group.
     * Changing this updates the description of the existing group.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     */
    readonly endpoints: pulumi.Output<string[] | undefined>;
    /**
     * The name of the group. Changing this updates the name of
     * the existing group.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     */
    readonly type: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a EndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointGroupArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering EndpointGroup resources.
 */
export interface EndpointGroupState {
    /**
     * The human-readable description for the group.
     * Changing this updates the description of the existing group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     */
    readonly endpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the group. Changing this updates the name of
     * the existing group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a EndpointGroup resource.
 */
export interface EndpointGroupArgs {
    /**
     * The human-readable description for the group.
     * Changing this updates the description of the existing group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     */
    readonly endpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the group. Changing this updates the name of
     * the existing group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
