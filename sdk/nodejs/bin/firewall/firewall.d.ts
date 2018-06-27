import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a v1 firewall resource within OpenStack.
 */
export declare class Firewall extends pulumi.CustomResource {
    /**
     * Get an existing Firewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallState): Firewall;
    /**
     * Administrative up/down status for the firewall
     * (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `admin_state_up` of an existing firewall.
     */
    readonly adminStateUp: pulumi.Output<boolean | undefined>;
    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     */
    readonly associatedRouters: pulumi.Output<string[]>;
    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Should this firewall not be associated with any routers
     * (must be "true" or "false" if provide - defaults to "false").
     * Conflicts with `associated_routers`.
     */
    readonly noRouters: pulumi.Output<boolean | undefined>;
    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     */
    readonly policyId: pulumi.Output<string>;
    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a Firewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Firewall resources.
 */
export interface FirewallState {
    /**
     * Administrative up/down status for the firewall
     * (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `admin_state_up` of an existing firewall.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     */
    readonly associatedRouters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Should this firewall not be associated with any routers
     * (must be "true" or "false" if provide - defaults to "false").
     * Conflicts with `associated_routers`.
     */
    readonly noRouters?: pulumi.Input<boolean>;
    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     */
    readonly policyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a Firewall resource.
 */
export interface FirewallArgs {
    /**
     * Administrative up/down status for the firewall
     * (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `admin_state_up` of an existing firewall.
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Router(s) to associate this firewall instance
     * with. Must be a list of strings. Changing this updates the associated routers
     * of an existing firewall. Conflicts with `no_routers`.
     */
    readonly associatedRouters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description for the firewall. Changing this
     * updates the `description` of an existing firewall.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A name for the firewall. Changing this
     * updates the `name` of an existing firewall.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Should this firewall not be associated with any routers
     * (must be "true" or "false" if provide - defaults to "false").
     * Conflicts with `associated_routers`.
     */
    readonly noRouters?: pulumi.Input<boolean>;
    /**
     * The policy resource id for the firewall. Changing
     * this updates the `policy_id` of an existing firewall.
     */
    readonly policyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the v1 networking client.
     * A networking client is needed to create a firewall. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The owner of the floating IP. Required if admin wants
     * to create a firewall for another tenant. Changing this creates a new
     * firewall.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
