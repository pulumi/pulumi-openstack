import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 DB configuration resource within OpenStack.
 */
export declare class Configuration extends pulumi.CustomResource {
    /**
     * Get an existing Configuration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationState): Configuration;
    /**
     * An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     */
    readonly configurations: pulumi.Output<{
        name: string;
        value: string;
    }[] | undefined>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     */
    readonly datastore: pulumi.Output<{
        type: string;
        version: string;
    }>;
    /**
     * Description of the resource.
     */
    readonly description: pulumi.Output<string>;
    /**
     * Configuration parameter name. Changing this creates a new resource.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    readonly region: pulumi.Output<string>;
    /**
     * Create a Configuration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Configuration resources.
 */
export interface ConfigurationState {
    /**
     * An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     */
    readonly configurations?: pulumi.Input<{
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }[]>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     */
    readonly datastore?: pulumi.Input<{
        type: pulumi.Input<string>;
        version: pulumi.Input<string>;
    }>;
    /**
     * Description of the resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Configuration parameter name. Changing this creates a new resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Configuration resource.
 */
export interface ConfigurationArgs {
    /**
     * An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     */
    readonly configurations?: pulumi.Input<{
        name: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }[]>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     */
    readonly datastore: pulumi.Input<{
        type: pulumi.Input<string>;
        version: pulumi.Input<string>;
    }>;
    /**
     * Description of the resource.
     */
    readonly description: pulumi.Input<string>;
    /**
     * Configuration parameter name. Changing this creates a new resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    readonly region: pulumi.Input<string>;
}
