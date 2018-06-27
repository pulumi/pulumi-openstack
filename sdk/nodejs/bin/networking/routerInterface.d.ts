import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 router interface resource within OpenStack.
 */
export declare class RouterInterface extends pulumi.CustomResource {
    /**
     * Get an existing RouterInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterInterfaceState): RouterInterface;
    /**
     * ID of the port this interface connects to. Changing
     * this creates a new router interface.
     */
    readonly portId: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router interface.
     */
    readonly region: pulumi.Output<string>;
    /**
     * ID of the router this interface belongs to. Changing
     * this creates a new router interface.
     */
    readonly routerId: pulumi.Output<string>;
    /**
     * ID of the subnet this interface connects to. Changing
     * this creates a new router interface.
     */
    readonly subnetId: pulumi.Output<string>;
    /**
     * Create a RouterInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterInterfaceArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering RouterInterface resources.
 */
export interface RouterInterfaceState {
    /**
     * ID of the port this interface connects to. Changing
     * this creates a new router interface.
     */
    readonly portId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router interface.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * ID of the router this interface belongs to. Changing
     * this creates a new router interface.
     */
    readonly routerId?: pulumi.Input<string>;
    /**
     * ID of the subnet this interface connects to. Changing
     * this creates a new router interface.
     */
    readonly subnetId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a RouterInterface resource.
 */
export interface RouterInterfaceArgs {
    /**
     * ID of the port this interface connects to. Changing
     * this creates a new router interface.
     */
    readonly portId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a router. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * router interface.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * ID of the router this interface belongs to. Changing
     * this creates a new router interface.
     */
    readonly routerId: pulumi.Input<string>;
    /**
     * ID of the subnet this interface connects to. Changing
     * this creates a new router interface.
     */
    readonly subnetId?: pulumi.Input<string>;
}
