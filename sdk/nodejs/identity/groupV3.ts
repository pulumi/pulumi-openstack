// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 group resource within OpenStack Keystone.
 *
 * > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const group1 = new openstack.identity.GroupV3("group_1", {
 *     name: "group_1",
 *     description: "group 1",
 * });
 * ```
 *
 * ## Import
 *
 * groups can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:identity/groupV3:GroupV3 group_1 89c60255-9bd6-460c-822a-e2b959ede9d2
 * ```
 */
export class GroupV3 extends pulumi.CustomResource {
    /**
     * Get an existing GroupV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupV3State, opts?: pulumi.CustomResourceOptions): GroupV3 {
        return new GroupV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:identity/groupV3:GroupV3';

    /**
     * Returns true if the given object is an instance of GroupV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupV3.__pulumiType;
    }

    /**
     * A description of the group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain the group belongs to.
     */
    public readonly domainId!: pulumi.Output<string>;
    /**
     * The name of the group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new group.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a GroupV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupV3Args | GroupV3State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupV3State | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as GroupV3Args | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupV3.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupV3 resources.
 */
export interface GroupV3State {
    /**
     * A description of the group.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain the group belongs to.
     */
    domainId?: pulumi.Input<string>;
    /**
     * The name of the group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new group.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupV3 resource.
 */
export interface GroupV3Args {
    /**
     * A description of the group.
     */
    description?: pulumi.Input<string>;
    /**
     * The domain the group belongs to.
     */
    domainId?: pulumi.Input<string>;
    /**
     * The name of the group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new group.
     */
    region?: pulumi.Input<string>;
}
