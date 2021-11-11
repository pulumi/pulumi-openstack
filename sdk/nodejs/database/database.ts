// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 DB database resource within OpenStack.
 *
 * ## Example Usage
 * ### Database
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const mydb = new openstack.database.Database("mydb", {
 *     instanceId: openstack_db_instance_v1_basic.id,
 * });
 * ```
 *
 * ## Import
 *
 * Databases can be imported by using `instance-id/db-name`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:database/database:Database mydb 7b9e3cd3-00d9-449c-b074-8439f8e274fa/mydb
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:database/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * The ID for the database instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * A unique name for the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Openstack region resource is created in.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * The ID for the database instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The ID for the database instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Openstack region resource is created in.
     */
    region?: pulumi.Input<string>;
}
