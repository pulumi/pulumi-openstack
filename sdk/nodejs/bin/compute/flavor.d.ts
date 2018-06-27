import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 flavor resource within OpenStack.
 */
export declare class Flavor extends pulumi.CustomResource {
    /**
     * Get an existing Flavor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlavorState): Flavor;
    /**
     * The amount of disk space in gigabytes to use for the root
     * (/) partition. Changing this creates a new flavor.
     */
    readonly disk: pulumi.Output<number>;
    readonly ephemeral: pulumi.Output<number | undefined>;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    readonly extraSpecs: pulumi.Output<{
        [key: string]: any;
    }>;
    /**
     * Whether the flavor is public. Changing this creates
     * a new flavor.
     */
    readonly isPublic: pulumi.Output<boolean | undefined>;
    /**
     * A unique name for the flavor. Changing this creates a new
     * flavor.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     */
    readonly ram: pulumi.Output<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     */
    readonly region: pulumi.Output<string>;
    /**
     * RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     */
    readonly rxTxFactor: pulumi.Output<number | undefined>;
    /**
     * The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     */
    readonly swap: pulumi.Output<number | undefined>;
    /**
     * The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     */
    readonly vcpus: pulumi.Output<number>;
    /**
     * Create a Flavor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlavorArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Flavor resources.
 */
export interface FlavorState {
    /**
     * The amount of disk space in gigabytes to use for the root
     * (/) partition. Changing this creates a new flavor.
     */
    readonly disk?: pulumi.Input<number>;
    readonly ephemeral?: pulumi.Input<number>;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    readonly extraSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * Whether the flavor is public. Changing this creates
     * a new flavor.
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * A unique name for the flavor. Changing this creates a new
     * flavor.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     */
    readonly ram?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     */
    readonly rxTxFactor?: pulumi.Input<number>;
    /**
     * The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     */
    readonly swap?: pulumi.Input<number>;
    /**
     * The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     */
    readonly vcpus?: pulumi.Input<number>;
}
/**
 * The set of arguments for constructing a Flavor resource.
 */
export interface FlavorArgs {
    /**
     * The amount of disk space in gigabytes to use for the root
     * (/) partition. Changing this creates a new flavor.
     */
    readonly disk: pulumi.Input<number>;
    readonly ephemeral?: pulumi.Input<number>;
    /**
     * Key/Value pairs of metadata for the flavor.
     */
    readonly extraSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * Whether the flavor is public. Changing this creates
     * a new flavor.
     */
    readonly isPublic?: pulumi.Input<boolean>;
    /**
     * A unique name for the flavor. Changing this creates a new
     * flavor.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The amount of RAM to use, in megabytes. Changing this
     * creates a new flavor.
     */
    readonly ram: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Flavors are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * RX/TX bandwith factor. The default is 1. Changing
     * this creates a new flavor.
     */
    readonly rxTxFactor?: pulumi.Input<number>;
    /**
     * The amount of disk space in megabytes to use. If
     * unspecified, the default is 0. Changing this creates a new flavor.
     */
    readonly swap?: pulumi.Input<number>;
    /**
     * The number of virtual CPUs to use. Changing this creates
     * a new flavor.
     */
    readonly vcpus: pulumi.Input<number>;
}
