// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 load balancer quota resource within OpenStack.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack
 *    API in case of delete call.
 *
 * > **Note:** This resource has attributes that depend on octavia minor versions.
 * Please ensure your Openstack cloud supports the required minor version.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const project1 = new openstack.identity.Project("project_1", {name: "project_1"});
 * const quota1 = new openstack.loadbalancer.Quota("quota_1", {
 *     projectId: project1.id,
 *     loadbalancer: 6,
 *     listener: 7,
 *     member: 8,
 *     pool: 9,
 *     healthMonitor: 10,
 *     l7Policy: 11,
 *     l7Rule: 12,
 * });
 * ```
 *
 * ## Import
 *
 * Quotas can be imported using the `project_id/region_name`, where region_name is the
 * one defined is the Openstack credentials that are in use. E.g.
 *
 * ```sh
 * $ pulumi import openstack:loadbalancer/quota:Quota quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
 * ```
 */
export class Quota extends pulumi.CustomResource {
    /**
     * Get an existing Quota resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuotaState, opts?: pulumi.CustomResourceOptions): Quota {
        return new Quota(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/quota:Quota';

    /**
     * Returns true if the given object is an instance of Quota.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Quota {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Quota.__pulumiType;
    }

    /**
     * Quota value for health_monitors. Changing
     * this updates the existing quota. Omitting it sets it to 0.
     */
    public readonly healthMonitor!: pulumi.Output<number>;
    /**
     * Quota value for l7_policies. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     */
    public readonly l7Policy!: pulumi.Output<number>;
    /**
     * Quota value for l7_rules. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     */
    public readonly l7Rule!: pulumi.Output<number>;
    /**
     * Quota value for listeners. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     */
    public readonly listener!: pulumi.Output<number>;
    /**
     * Quota value for loadbalancers. Changing this
     * updates the existing quota. Omitting it sets it to 0.
     */
    public readonly loadbalancer!: pulumi.Output<number>;
    /**
     * Quota value for members. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     */
    public readonly member!: pulumi.Output<number>;
    /**
     * Quota value for pools. Changing this updates the
     * the existing quota. Omitting it sets it to 0.
     */
    public readonly pool!: pulumi.Output<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quota.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Region in which to manage quotas. Changing this
     * creates a new quota. If ommited, the region of the credentials is used.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a Quota resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QuotaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuotaArgs | QuotaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QuotaState | undefined;
            resourceInputs["healthMonitor"] = state ? state.healthMonitor : undefined;
            resourceInputs["l7Policy"] = state ? state.l7Policy : undefined;
            resourceInputs["l7Rule"] = state ? state.l7Rule : undefined;
            resourceInputs["listener"] = state ? state.listener : undefined;
            resourceInputs["loadbalancer"] = state ? state.loadbalancer : undefined;
            resourceInputs["member"] = state ? state.member : undefined;
            resourceInputs["pool"] = state ? state.pool : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as QuotaArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["healthMonitor"] = args ? args.healthMonitor : undefined;
            resourceInputs["l7Policy"] = args ? args.l7Policy : undefined;
            resourceInputs["l7Rule"] = args ? args.l7Rule : undefined;
            resourceInputs["listener"] = args ? args.listener : undefined;
            resourceInputs["loadbalancer"] = args ? args.loadbalancer : undefined;
            resourceInputs["member"] = args ? args.member : undefined;
            resourceInputs["pool"] = args ? args.pool : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Quota.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Quota resources.
 */
export interface QuotaState {
    /**
     * Quota value for health_monitors. Changing
     * this updates the existing quota. Omitting it sets it to 0.
     */
    healthMonitor?: pulumi.Input<number>;
    /**
     * Quota value for l7_policies. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     */
    l7Policy?: pulumi.Input<number>;
    /**
     * Quota value for l7_rules. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     */
    l7Rule?: pulumi.Input<number>;
    /**
     * Quota value for listeners. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     */
    listener?: pulumi.Input<number>;
    /**
     * Quota value for loadbalancers. Changing this
     * updates the existing quota. Omitting it sets it to 0.
     */
    loadbalancer?: pulumi.Input<number>;
    /**
     * Quota value for members. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     */
    member?: pulumi.Input<number>;
    /**
     * Quota value for pools. Changing this updates the
     * the existing quota. Omitting it sets it to 0.
     */
    pool?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quota.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Region in which to manage quotas. Changing this
     * creates a new quota. If ommited, the region of the credentials is used.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Quota resource.
 */
export interface QuotaArgs {
    /**
     * Quota value for health_monitors. Changing
     * this updates the existing quota. Omitting it sets it to 0.
     */
    healthMonitor?: pulumi.Input<number>;
    /**
     * Quota value for l7_policies. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     */
    l7Policy?: pulumi.Input<number>;
    /**
     * Quota value for l7_rules. Changing this
     * updates the existing quota. Omitting it sets it to 0. Available in
     * **Octavia minor version 2.19**.
     */
    l7Rule?: pulumi.Input<number>;
    /**
     * Quota value for listeners. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     */
    listener?: pulumi.Input<number>;
    /**
     * Quota value for loadbalancers. Changing this
     * updates the existing quota. Omitting it sets it to 0.
     */
    loadbalancer?: pulumi.Input<number>;
    /**
     * Quota value for members. Changing this updates
     * the existing quota. Omitting it sets it to 0.
     */
    member?: pulumi.Input<number>;
    /**
     * Quota value for pools. Changing this updates the
     * the existing quota. Omitting it sets it to 0.
     */
    pool?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quota.
     */
    projectId: pulumi.Input<string>;
    /**
     * Region in which to manage quotas. Changing this
     * creates a new quota. If ommited, the region of the credentials is used.
     */
    region?: pulumi.Input<string>;
}
