// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V1 DB instance resource within OpenStack.
 *
 * > **Note:** All arguments including the instance user password will be stored
 * in the raw state as plain-text. Read more about sensitive data in
 * state.
 *
 * ## Example Usage
 *
 * ### Instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const test = new openstack.database.Instance("test", {
 *     region: "region-test",
 *     name: "test",
 *     flavorId: "31792d21-c355-4587-9290-56c1ed0ca376",
 *     size: 8,
 *     networks: [{
 *         uuid: "c0612505-caf2-4fb0-b7cb-56a0240a2b12",
 *     }],
 *     datastore: {
 *         version: "mysql-5.7",
 *         type: "mysql",
 *     },
 * });
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:database/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * A list of IP addresses assigned to the instance.
     */
    public /*out*/ readonly addresses!: pulumi.Output<string[]>;
    /**
     * Configuration ID to be attached to the instance. Database instance
     * will be rebooted when configuration is detached.
     */
    public readonly configurationId!: pulumi.Output<string | undefined>;
    /**
     * An array of database name, charset and collate. The database
     * object structure is documented below.
     */
    public readonly databases!: pulumi.Output<outputs.database.InstanceDatabase[] | undefined>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates a new instance.
     */
    public readonly datastore!: pulumi.Output<outputs.database.InstanceDatastore>;
    /**
     * The flavor ID of the desired flavor for the instance.
     * Changing this creates new instance.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * A unique name for the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new instance.
     */
    public readonly networks!: pulumi.Output<outputs.database.InstanceNetwork[] | undefined>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the volume size in GB. Changing this creates new instance.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * An array of username, password, host and databases. The user
     * object structure is documented below.
     */
    public readonly users!: pulumi.Output<outputs.database.InstanceUser[] | undefined>;
    /**
     * Specifies the volume type to use. If you want to
     * specify a volume type, you must also specify a volume size. Changing this
     * creates new instance.
     */
    public readonly volumeType!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["addresses"] = state ? state.addresses : undefined;
            resourceInputs["configurationId"] = state ? state.configurationId : undefined;
            resourceInputs["databases"] = state ? state.databases : undefined;
            resourceInputs["datastore"] = state ? state.datastore : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["users"] = state ? state.users : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.datastore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datastore'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["configurationId"] = args ? args.configurationId : undefined;
            resourceInputs["databases"] = args ? args.databases : undefined;
            resourceInputs["datastore"] = args ? args.datastore : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["users"] = args ? args.users : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["addresses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * A list of IP addresses assigned to the instance.
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configuration ID to be attached to the instance. Database instance
     * will be rebooted when configuration is detached.
     */
    configurationId?: pulumi.Input<string>;
    /**
     * An array of database name, charset and collate. The database
     * object structure is documented below.
     */
    databases?: pulumi.Input<pulumi.Input<inputs.database.InstanceDatabase>[]>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates a new instance.
     */
    datastore?: pulumi.Input<inputs.database.InstanceDatastore>;
    /**
     * The flavor ID of the desired flavor for the instance.
     * Changing this creates new instance.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new instance.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.database.InstanceNetwork>[]>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the volume size in GB. Changing this creates new instance.
     */
    size?: pulumi.Input<number>;
    /**
     * An array of username, password, host and databases. The user
     * object structure is documented below.
     */
    users?: pulumi.Input<pulumi.Input<inputs.database.InstanceUser>[]>;
    /**
     * Specifies the volume type to use. If you want to
     * specify a volume type, you must also specify a volume size. Changing this
     * creates new instance.
     */
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Configuration ID to be attached to the instance. Database instance
     * will be rebooted when configuration is detached.
     */
    configurationId?: pulumi.Input<string>;
    /**
     * An array of database name, charset and collate. The database
     * object structure is documented below.
     */
    databases?: pulumi.Input<pulumi.Input<inputs.database.InstanceDatabase>[]>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates a new instance.
     */
    datastore: pulumi.Input<inputs.database.InstanceDatastore>;
    /**
     * The flavor ID of the desired flavor for the instance.
     * Changing this creates new instance.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new instance.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.database.InstanceNetwork>[]>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the volume size in GB. Changing this creates new instance.
     */
    size: pulumi.Input<number>;
    /**
     * An array of username, password, host and databases. The user
     * object structure is documented below.
     */
    users?: pulumi.Input<pulumi.Input<inputs.database.InstanceUser>[]>;
    /**
     * Specifies the volume type to use. If you want to
     * specify a volume type, you must also specify a volume size. Changing this
     * creates new instance.
     */
    volumeType?: pulumi.Input<string>;
}
