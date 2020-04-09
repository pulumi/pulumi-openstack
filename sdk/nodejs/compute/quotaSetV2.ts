// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 compute quotaset resource within OpenStack.
 * 
 * > **Note:** This usually requires admin privileges.
 * 
 * > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API 
 *     in case of delete call.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_quotaset_v2.html.markdown.
 */
export class QuotaSetV2 extends pulumi.CustomResource {
    /**
     * Get an existing QuotaSetV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuotaSetV2State, opts?: pulumi.CustomResourceOptions): QuotaSetV2 {
        return new QuotaSetV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/quotaSetV2:QuotaSetV2';

    /**
     * Returns true if the given object is an instance of QuotaSetV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QuotaSetV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QuotaSetV2.__pulumiType;
    }

    /**
     * Quota value for cores.
     * Changing this updates the existing quotaset.
     */
    public readonly cores!: pulumi.Output<number>;
    /**
     * Quota value for fixed IPs.
     * Changing this updates the existing quotaset.
     */
    public readonly fixedIps!: pulumi.Output<number>;
    /**
     * Quota value for floating IPs.
     * Changing this updates the existing quotaset.
     */
    public readonly floatingIps!: pulumi.Output<number>;
    /**
     * Quota value for content bytes
     * of injected files. Changing this updates the existing quotaset.
     */
    public readonly injectedFileContentBytes!: pulumi.Output<number>;
    /**
     * Quota value for path bytes of
     * injected files. Changing this updates the existing quotaset.
     */
    public readonly injectedFilePathBytes!: pulumi.Output<number>;
    /**
     * Quota value for injected files.
     * Changing this updates the existing quotaset.
     */
    public readonly injectedFiles!: pulumi.Output<number>;
    /**
     * Quota value for instances.
     * Changing this updates the existing quotaset.
     */
    public readonly instances!: pulumi.Output<number>;
    /**
     * Quota value for key pairs.
     * Changing this updates the existing quotaset.
     */
    public readonly keyPairs!: pulumi.Output<number>;
    /**
     * Quota value for metadata items.
     * Changing this updates the existing quotaset.
     */
    public readonly metadataItems!: pulumi.Output<number>;
    /**
     * ID of the project to manage quotas.
     * Changing this creates a new quotaset.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Quota value for RAM.
     * Changing this updates the existing quotaset.
     */
    public readonly ram!: pulumi.Output<number>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Quota value for security group rules.
     * Changing this updates the existing quotaset.
     */
    public readonly securityGroupRules!: pulumi.Output<number>;
    /**
     * Quota value for security groups.
     * Changing this updates the existing quotaset.
     */
    public readonly securityGroups!: pulumi.Output<number>;
    /**
     * Quota value for server groups members.
     * Changing this updates the existing quotaset.
     */
    public readonly serverGroupMembers!: pulumi.Output<number>;
    /**
     * Quota value for server groups.
     * Changing this updates the existing quotaset.
     */
    public readonly serverGroups!: pulumi.Output<number>;

    /**
     * Create a QuotaSetV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QuotaSetV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuotaSetV2Args | QuotaSetV2State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QuotaSetV2State | undefined;
            inputs["cores"] = state ? state.cores : undefined;
            inputs["fixedIps"] = state ? state.fixedIps : undefined;
            inputs["floatingIps"] = state ? state.floatingIps : undefined;
            inputs["injectedFileContentBytes"] = state ? state.injectedFileContentBytes : undefined;
            inputs["injectedFilePathBytes"] = state ? state.injectedFilePathBytes : undefined;
            inputs["injectedFiles"] = state ? state.injectedFiles : undefined;
            inputs["instances"] = state ? state.instances : undefined;
            inputs["keyPairs"] = state ? state.keyPairs : undefined;
            inputs["metadataItems"] = state ? state.metadataItems : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["ram"] = state ? state.ram : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["securityGroupRules"] = state ? state.securityGroupRules : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["serverGroupMembers"] = state ? state.serverGroupMembers : undefined;
            inputs["serverGroups"] = state ? state.serverGroups : undefined;
        } else {
            const args = argsOrState as QuotaSetV2Args | undefined;
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["cores"] = args ? args.cores : undefined;
            inputs["fixedIps"] = args ? args.fixedIps : undefined;
            inputs["floatingIps"] = args ? args.floatingIps : undefined;
            inputs["injectedFileContentBytes"] = args ? args.injectedFileContentBytes : undefined;
            inputs["injectedFilePathBytes"] = args ? args.injectedFilePathBytes : undefined;
            inputs["injectedFiles"] = args ? args.injectedFiles : undefined;
            inputs["instances"] = args ? args.instances : undefined;
            inputs["keyPairs"] = args ? args.keyPairs : undefined;
            inputs["metadataItems"] = args ? args.metadataItems : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["ram"] = args ? args.ram : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["securityGroupRules"] = args ? args.securityGroupRules : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["serverGroupMembers"] = args ? args.serverGroupMembers : undefined;
            inputs["serverGroups"] = args ? args.serverGroups : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(QuotaSetV2.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QuotaSetV2 resources.
 */
export interface QuotaSetV2State {
    /**
     * Quota value for cores.
     * Changing this updates the existing quotaset.
     */
    readonly cores?: pulumi.Input<number>;
    /**
     * Quota value for fixed IPs.
     * Changing this updates the existing quotaset.
     */
    readonly fixedIps?: pulumi.Input<number>;
    /**
     * Quota value for floating IPs.
     * Changing this updates the existing quotaset.
     */
    readonly floatingIps?: pulumi.Input<number>;
    /**
     * Quota value for content bytes
     * of injected files. Changing this updates the existing quotaset.
     */
    readonly injectedFileContentBytes?: pulumi.Input<number>;
    /**
     * Quota value for path bytes of
     * injected files. Changing this updates the existing quotaset.
     */
    readonly injectedFilePathBytes?: pulumi.Input<number>;
    /**
     * Quota value for injected files.
     * Changing this updates the existing quotaset.
     */
    readonly injectedFiles?: pulumi.Input<number>;
    /**
     * Quota value for instances.
     * Changing this updates the existing quotaset.
     */
    readonly instances?: pulumi.Input<number>;
    /**
     * Quota value for key pairs.
     * Changing this updates the existing quotaset.
     */
    readonly keyPairs?: pulumi.Input<number>;
    /**
     * Quota value for metadata items.
     * Changing this updates the existing quotaset.
     */
    readonly metadataItems?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas.
     * Changing this creates a new quotaset.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * Quota value for RAM.
     * Changing this updates the existing quotaset.
     */
    readonly ram?: pulumi.Input<number>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Quota value for security group rules.
     * Changing this updates the existing quotaset.
     */
    readonly securityGroupRules?: pulumi.Input<number>;
    /**
     * Quota value for security groups.
     * Changing this updates the existing quotaset.
     */
    readonly securityGroups?: pulumi.Input<number>;
    /**
     * Quota value for server groups members.
     * Changing this updates the existing quotaset.
     */
    readonly serverGroupMembers?: pulumi.Input<number>;
    /**
     * Quota value for server groups.
     * Changing this updates the existing quotaset.
     */
    readonly serverGroups?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a QuotaSetV2 resource.
 */
export interface QuotaSetV2Args {
    /**
     * Quota value for cores.
     * Changing this updates the existing quotaset.
     */
    readonly cores?: pulumi.Input<number>;
    /**
     * Quota value for fixed IPs.
     * Changing this updates the existing quotaset.
     */
    readonly fixedIps?: pulumi.Input<number>;
    /**
     * Quota value for floating IPs.
     * Changing this updates the existing quotaset.
     */
    readonly floatingIps?: pulumi.Input<number>;
    /**
     * Quota value for content bytes
     * of injected files. Changing this updates the existing quotaset.
     */
    readonly injectedFileContentBytes?: pulumi.Input<number>;
    /**
     * Quota value for path bytes of
     * injected files. Changing this updates the existing quotaset.
     */
    readonly injectedFilePathBytes?: pulumi.Input<number>;
    /**
     * Quota value for injected files.
     * Changing this updates the existing quotaset.
     */
    readonly injectedFiles?: pulumi.Input<number>;
    /**
     * Quota value for instances.
     * Changing this updates the existing quotaset.
     */
    readonly instances?: pulumi.Input<number>;
    /**
     * Quota value for key pairs.
     * Changing this updates the existing quotaset.
     */
    readonly keyPairs?: pulumi.Input<number>;
    /**
     * Quota value for metadata items.
     * Changing this updates the existing quotaset.
     */
    readonly metadataItems?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas.
     * Changing this creates a new quotaset.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * Quota value for RAM.
     * Changing this updates the existing quotaset.
     */
    readonly ram?: pulumi.Input<number>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Quota value for security group rules.
     * Changing this updates the existing quotaset.
     */
    readonly securityGroupRules?: pulumi.Input<number>;
    /**
     * Quota value for security groups.
     * Changing this updates the existing quotaset.
     */
    readonly securityGroups?: pulumi.Input<number>;
    /**
     * Quota value for server groups members.
     * Changing this updates the existing quotaset.
     */
    readonly serverGroupMembers?: pulumi.Input<number>;
    /**
     * Quota value for server groups.
     * Changing this updates the existing quotaset.
     */
    readonly serverGroups?: pulumi.Input<number>;
}
