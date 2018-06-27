import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 Server Group resource within OpenStack.
 */
export declare class ServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerGroupState): ServerGroup;
    /**
     * The instances that are part of this server group.
     */
    readonly members: pulumi.Output<string[]>;
    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The set of policies for the server group. Only two
     * two policies are available right now, and both are mutually exclusive. See
     * the Policies section for more information. Changing this creates a new
     * server group.
     */
    readonly policies: pulumi.Output<string[] | undefined>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     */
    readonly region: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a ServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServerGroupArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering ServerGroup resources.
 */
export interface ServerGroupState {
    /**
     * The instances that are part of this server group.
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The set of policies for the server group. Only two
     * two policies are available right now, and both are mutually exclusive. See
     * the Policies section for more information. Changing this creates a new
     * server group.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a ServerGroup resource.
 */
export interface ServerGroupArgs {
    /**
     * A unique name for the server group. Changing this creates
     * a new server group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The set of policies for the server group. Only two
     * two policies are available right now, and both are mutually exclusive. See
     * the Policies section for more information. Changing this creates a new
     * server group.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new server group.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
