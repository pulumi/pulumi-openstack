// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 networking quota resource within OpenStack.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
 *     in case of delete call.
 *
 * > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
 *     created with zero value.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const project1 = new openstack.identity.Project("project1", {});
 * const quota1 = new openstack.networking.QuotaV2("quota1", {
 *     projectId: project1.id,
 *     floatingip: 10,
 *     network: 4,
 *     port: 100,
 *     rbacPolicy: 10,
 *     router: 4,
 *     securityGroup: 10,
 *     securityGroupRule: 100,
 *     subnet: 8,
 *     subnetpool: 2,
 * });
 * ```
 *
 * ## Import
 *
 * Quotas can be imported using the `project_id/region_name`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:networking/quotaV2:QuotaV2 quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
 * ```
 */
export class QuotaV2 extends pulumi.CustomResource {
    /**
     * Get an existing QuotaV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuotaV2State, opts?: pulumi.CustomResourceOptions): QuotaV2 {
        return new QuotaV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/quotaV2:QuotaV2';

    /**
     * Returns true if the given object is an instance of QuotaV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QuotaV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QuotaV2.__pulumiType;
    }

    /**
     * Quota value for floating IPs. Changing this updates the
     * existing quota.
     */
    public readonly floatingip!: pulumi.Output<number>;
    /**
     * Quota value for networks. Changing this updates the
     * existing quota.
     */
    public readonly network!: pulumi.Output<number>;
    /**
     * Quota value for ports. Changing this updates the
     * existing quota.
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * ID of the project to manage quota. Changing this
     * creates new quota.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Quota value for RBAC policies.
     * Changing this updates the existing quota.
     */
    public readonly rbacPolicy!: pulumi.Output<number>;
    /**
     * The region in which to create the quota. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates new quota.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Quota value for routers. Changing this updates the
     * existing quota.
     */
    public readonly router!: pulumi.Output<number>;
    /**
     * Quota value for security groups. Changing
     * this updates the existing quota.
     */
    public readonly securityGroup!: pulumi.Output<number>;
    /**
     * Quota value for security group rules.
     * Changing this updates the existing quota.
     */
    public readonly securityGroupRule!: pulumi.Output<number>;
    /**
     * Quota value for subnets. Changing
     * this updates the existing quota.
     */
    public readonly subnet!: pulumi.Output<number>;
    /**
     * Quota value for subnetpools.
     * Changing this updates the existing quota.
     */
    public readonly subnetpool!: pulumi.Output<number>;

    /**
     * Create a QuotaV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QuotaV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuotaV2Args | QuotaV2State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QuotaV2State | undefined;
            inputs["floatingip"] = state ? state.floatingip : undefined;
            inputs["network"] = state ? state.network : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["rbacPolicy"] = state ? state.rbacPolicy : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["securityGroup"] = state ? state.securityGroup : undefined;
            inputs["securityGroupRule"] = state ? state.securityGroupRule : undefined;
            inputs["subnet"] = state ? state.subnet : undefined;
            inputs["subnetpool"] = state ? state.subnetpool : undefined;
        } else {
            const args = argsOrState as QuotaV2Args | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["floatingip"] = args ? args.floatingip : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["port"] = args ? args.port : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["rbacPolicy"] = args ? args.rbacPolicy : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["securityGroup"] = args ? args.securityGroup : undefined;
            inputs["securityGroupRule"] = args ? args.securityGroupRule : undefined;
            inputs["subnet"] = args ? args.subnet : undefined;
            inputs["subnetpool"] = args ? args.subnetpool : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(QuotaV2.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QuotaV2 resources.
 */
export interface QuotaV2State {
    /**
     * Quota value for floating IPs. Changing this updates the
     * existing quota.
     */
    floatingip?: pulumi.Input<number>;
    /**
     * Quota value for networks. Changing this updates the
     * existing quota.
     */
    network?: pulumi.Input<number>;
    /**
     * Quota value for ports. Changing this updates the
     * existing quota.
     */
    port?: pulumi.Input<number>;
    /**
     * ID of the project to manage quota. Changing this
     * creates new quota.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Quota value for RBAC policies.
     * Changing this updates the existing quota.
     */
    rbacPolicy?: pulumi.Input<number>;
    /**
     * The region in which to create the quota. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates new quota.
     */
    region?: pulumi.Input<string>;
    /**
     * Quota value for routers. Changing this updates the
     * existing quota.
     */
    router?: pulumi.Input<number>;
    /**
     * Quota value for security groups. Changing
     * this updates the existing quota.
     */
    securityGroup?: pulumi.Input<number>;
    /**
     * Quota value for security group rules.
     * Changing this updates the existing quota.
     */
    securityGroupRule?: pulumi.Input<number>;
    /**
     * Quota value for subnets. Changing
     * this updates the existing quota.
     */
    subnet?: pulumi.Input<number>;
    /**
     * Quota value for subnetpools.
     * Changing this updates the existing quota.
     */
    subnetpool?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a QuotaV2 resource.
 */
export interface QuotaV2Args {
    /**
     * Quota value for floating IPs. Changing this updates the
     * existing quota.
     */
    floatingip?: pulumi.Input<number>;
    /**
     * Quota value for networks. Changing this updates the
     * existing quota.
     */
    network?: pulumi.Input<number>;
    /**
     * Quota value for ports. Changing this updates the
     * existing quota.
     */
    port?: pulumi.Input<number>;
    /**
     * ID of the project to manage quota. Changing this
     * creates new quota.
     */
    projectId: pulumi.Input<string>;
    /**
     * Quota value for RBAC policies.
     * Changing this updates the existing quota.
     */
    rbacPolicy?: pulumi.Input<number>;
    /**
     * The region in which to create the quota. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates new quota.
     */
    region?: pulumi.Input<string>;
    /**
     * Quota value for routers. Changing this updates the
     * existing quota.
     */
    router?: pulumi.Input<number>;
    /**
     * Quota value for security groups. Changing
     * this updates the existing quota.
     */
    securityGroup?: pulumi.Input<number>;
    /**
     * Quota value for security group rules.
     * Changing this updates the existing quota.
     */
    securityGroupRule?: pulumi.Input<number>;
    /**
     * Quota value for subnets. Changing
     * this updates the existing quota.
     */
    subnet?: pulumi.Input<number>;
    /**
     * Quota value for subnetpools.
     * Changing this updates the existing quota.
     */
    subnetpool?: pulumi.Input<number>;
}
