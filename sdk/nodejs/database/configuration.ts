// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class Configuration extends pulumi.CustomResource {
    /**
     * Get an existing Configuration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationState, opts?: pulumi.CustomResourceOptions): Configuration {
        return new Configuration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:database/configuration:Configuration';

    /**
     * Returns true if the given object is an instance of Configuration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Configuration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Configuration.__pulumiType;
    }

    /**
     * An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     */
    public readonly configurations!: pulumi.Output<outputs.database.ConfigurationConfiguration[] | undefined>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     */
    public readonly datastore!: pulumi.Output<outputs.database.ConfigurationDatastore>;
    /**
     * Description of the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Configuration parameter name. Changing this creates a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Configuration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigurationArgs | ConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigurationState | undefined;
            resourceInputs["configurations"] = state ? state.configurations : undefined;
            resourceInputs["datastore"] = state ? state.datastore : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as ConfigurationArgs | undefined;
            if ((!args || args.datastore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datastore'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            resourceInputs["configurations"] = args ? args.configurations : undefined;
            resourceInputs["datastore"] = args ? args.datastore : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Configuration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Configuration resources.
 */
export interface ConfigurationState {
    /**
     * An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     */
    configurations?: pulumi.Input<pulumi.Input<inputs.database.ConfigurationConfiguration>[]>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     */
    datastore?: pulumi.Input<inputs.database.ConfigurationDatastore>;
    /**
     * Description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Configuration parameter name. Changing this creates a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Configuration resource.
 */
export interface ConfigurationArgs {
    /**
     * An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below.
     */
    configurations?: pulumi.Input<pulumi.Input<inputs.database.ConfigurationConfiguration>[]>;
    /**
     * An array of database engine type and version. The datastore
     * object structure is documented below. Changing this creates resource.
     */
    datastore: pulumi.Input<inputs.database.ConfigurationDatastore>;
    /**
     * Description of the resource.
     */
    description: pulumi.Input<string>;
    /**
     * Configuration parameter name. Changing this creates a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the db instance. Changing this
     * creates a new instance.
     */
    region?: pulumi.Input<string>;
}
