// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a project access for flavor V2 resource within OpenStack.
 *
 * > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 *
 * ***
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const project1 = new openstack.identity.Project("project1", {});
 * const flavor1 = new openstack.compute.Flavor("flavor1", {
 *     ram: 8096,
 *     vcpus: 2,
 *     disk: 20,
 *     isPublic: false,
 * });
 * const access1 = new openstack.compute.FlavorAccess("access1", {
 *     tenantId: project1.id,
 *     flavorId: flavor1.id,
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported by specifying all two arguments, separated
 * by a forward slash:
 *
 * ```sh
 * $ pulumi import openstack:compute/flavorAccess:FlavorAccess access_1 flavor_id/tenant_id
 * ```
 */
export class FlavorAccess extends pulumi.CustomResource {
    /**
     * Get an existing FlavorAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FlavorAccessState, opts?: pulumi.CustomResourceOptions): FlavorAccess {
        return new FlavorAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/flavorAccess:FlavorAccess';

    /**
     * Returns true if the given object is an instance of FlavorAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlavorAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlavorAccess.__pulumiType;
    }

    /**
     * The UUID of flavor to use. Changing this creates a new flavor access.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor access.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The UUID of tenant which is allowed to use the flavor.
     * Changing this creates a new flavor access.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a FlavorAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlavorAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FlavorAccessArgs | FlavorAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FlavorAccessState | undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as FlavorAccessArgs | undefined;
            if ((!args || args.flavorId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavorId'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FlavorAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FlavorAccess resources.
 */
export interface FlavorAccessState {
    /**
     * The UUID of flavor to use. Changing this creates a new flavor access.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor access.
     */
    region?: pulumi.Input<string>;
    /**
     * The UUID of tenant which is allowed to use the flavor.
     * Changing this creates a new flavor access.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FlavorAccess resource.
 */
export interface FlavorAccessArgs {
    /**
     * The UUID of flavor to use. Changing this creates a new flavor access.
     */
    flavorId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new flavor access.
     */
    region?: pulumi.Input<string>;
    /**
     * The UUID of tenant which is allowed to use the flavor.
     * Changing this creates a new flavor access.
     */
    tenantId: pulumi.Input<string>;
}
