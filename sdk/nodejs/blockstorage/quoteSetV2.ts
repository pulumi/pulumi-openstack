// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 block storage quotaset resource within OpenStack.
 *
 * > **Note:** This usually requires admin privileges.
 *
 * > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
 * in case of delete call.
 *
 * > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
 * created with zero value. This excludes volume type quota.
 *
 * ## Import
 *
 * Quotasets can be imported using the `project_id/region`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:blockstorage/quoteSetV2:QuoteSetV2 quotaset_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
 * ```
 */
export class QuoteSetV2 extends pulumi.CustomResource {
    /**
     * Get an existing QuoteSetV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuoteSetV2State, opts?: pulumi.CustomResourceOptions): QuoteSetV2 {
        return new QuoteSetV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:blockstorage/quoteSetV2:QuoteSetV2';

    /**
     * Returns true if the given object is an instance of QuoteSetV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QuoteSetV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QuoteSetV2.__pulumiType;
    }

    /**
     * Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     */
    public readonly backupGigabytes!: pulumi.Output<number>;
    /**
     * Quota value for backups. Changing this updates the
     * existing quotaset.
     */
    public readonly backups!: pulumi.Output<number>;
    /**
     * Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     */
    public readonly gigabytes!: pulumi.Output<number>;
    /**
     * Quota value for groups. Changing this updates the
     * existing quotaset.
     */
    public readonly groups!: pulumi.Output<number>;
    /**
     * Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     */
    public readonly perVolumeGigabytes!: pulumi.Output<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Quota value for snapshots. Changing this updates the
     * existing quotaset.
     */
    public readonly snapshots!: pulumi.Output<number>;
    /**
     * Key/Value pairs for setting quota for
     * volumes types. Possible keys are `snapshots_<volume_type_name>`,
     * `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
     */
    public readonly volumeTypeQuota!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Quota value for volumes. Changing this updates the
     * existing quotaset.
     */
    public readonly volumes!: pulumi.Output<number>;

    /**
     * Create a QuoteSetV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QuoteSetV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuoteSetV2Args | QuoteSetV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QuoteSetV2State | undefined;
            resourceInputs["backupGigabytes"] = state ? state.backupGigabytes : undefined;
            resourceInputs["backups"] = state ? state.backups : undefined;
            resourceInputs["gigabytes"] = state ? state.gigabytes : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["perVolumeGigabytes"] = state ? state.perVolumeGigabytes : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["snapshots"] = state ? state.snapshots : undefined;
            resourceInputs["volumeTypeQuota"] = state ? state.volumeTypeQuota : undefined;
            resourceInputs["volumes"] = state ? state.volumes : undefined;
        } else {
            const args = argsOrState as QuoteSetV2Args | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["backupGigabytes"] = args ? args.backupGigabytes : undefined;
            resourceInputs["backups"] = args ? args.backups : undefined;
            resourceInputs["gigabytes"] = args ? args.gigabytes : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["perVolumeGigabytes"] = args ? args.perVolumeGigabytes : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["snapshots"] = args ? args.snapshots : undefined;
            resourceInputs["volumeTypeQuota"] = args ? args.volumeTypeQuota : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(QuoteSetV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QuoteSetV2 resources.
 */
export interface QuoteSetV2State {
    /**
     * Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     */
    backupGigabytes?: pulumi.Input<number>;
    /**
     * Quota value for backups. Changing this updates the
     * existing quotaset.
     */
    backups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     */
    gigabytes?: pulumi.Input<number>;
    /**
     * Quota value for groups. Changing this updates the
     * existing quotaset.
     */
    groups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     */
    perVolumeGigabytes?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    region?: pulumi.Input<string>;
    /**
     * Quota value for snapshots. Changing this updates the
     * existing quotaset.
     */
    snapshots?: pulumi.Input<number>;
    /**
     * Key/Value pairs for setting quota for
     * volumes types. Possible keys are `snapshots_<volume_type_name>`,
     * `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
     */
    volumeTypeQuota?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Quota value for volumes. Changing this updates the
     * existing quotaset.
     */
    volumes?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a QuoteSetV2 resource.
 */
export interface QuoteSetV2Args {
    /**
     * Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     */
    backupGigabytes?: pulumi.Input<number>;
    /**
     * Quota value for backups. Changing this updates the
     * existing quotaset.
     */
    backups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     */
    gigabytes?: pulumi.Input<number>;
    /**
     * Quota value for groups. Changing this updates the
     * existing quotaset.
     */
    groups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     */
    perVolumeGigabytes?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     */
    projectId: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    region?: pulumi.Input<string>;
    /**
     * Quota value for snapshots. Changing this updates the
     * existing quotaset.
     */
    snapshots?: pulumi.Input<number>;
    /**
     * Key/Value pairs for setting quota for
     * volumes types. Possible keys are `snapshots_<volume_type_name>`,
     * `volumes_<volume_type_name>` and `gigabytes_<volume_type_name>`.
     */
    volumeTypeQuota?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Quota value for volumes. Changing this updates the
     * existing quotaset.
     */
    volumes?: pulumi.Input<number>;
}
