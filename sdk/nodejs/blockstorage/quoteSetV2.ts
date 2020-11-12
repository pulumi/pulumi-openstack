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
 * const quotaset1 = new openstack.blockstorage.QuoteSetV2("quotaset1", {
 *     projectId: project1.id,
 *     volumes: 10,
 *     snapshots: 4,
 *     gigabytes: 100,
 *     perVolumeGigabytes: 10,
 *     backups: 4,
 *     backupGigabytes: 10,
 *     groups: 100,
 * });
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QuoteSetV2State | undefined;
            inputs["backupGigabytes"] = state ? state.backupGigabytes : undefined;
            inputs["backups"] = state ? state.backups : undefined;
            inputs["gigabytes"] = state ? state.gigabytes : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["perVolumeGigabytes"] = state ? state.perVolumeGigabytes : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["snapshots"] = state ? state.snapshots : undefined;
            inputs["volumes"] = state ? state.volumes : undefined;
        } else {
            const args = argsOrState as QuoteSetV2Args | undefined;
            if (!args || args.projectId === undefined) {
                throw new Error("Missing required property 'projectId'");
            }
            inputs["backupGigabytes"] = args ? args.backupGigabytes : undefined;
            inputs["backups"] = args ? args.backups : undefined;
            inputs["gigabytes"] = args ? args.gigabytes : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["perVolumeGigabytes"] = args ? args.perVolumeGigabytes : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["snapshots"] = args ? args.snapshots : undefined;
            inputs["volumes"] = args ? args.volumes : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(QuoteSetV2.__pulumiType, name, inputs, opts);
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
    readonly backupGigabytes?: pulumi.Input<number>;
    /**
     * Quota value for backups. Changing this updates the
     * existing quotaset.
     */
    readonly backups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     */
    readonly gigabytes?: pulumi.Input<number>;
    /**
     * Quota value for groups. Changing this updates the
     * existing quotaset.
     */
    readonly groups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     */
    readonly perVolumeGigabytes?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Quota value for snapshots. Changing this updates the
     * existing quotaset.
     */
    readonly snapshots?: pulumi.Input<number>;
    /**
     * Quota value for volumes. Changing this updates the
     * existing quotaset.
     */
    readonly volumes?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a QuoteSetV2 resource.
 */
export interface QuoteSetV2Args {
    /**
     * Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     */
    readonly backupGigabytes?: pulumi.Input<number>;
    /**
     * Quota value for backups. Changing this updates the
     * existing quotaset.
     */
    readonly backups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     */
    readonly gigabytes?: pulumi.Input<number>;
    /**
     * Quota value for groups. Changing this updates the
     * existing quotaset.
     */
    readonly groups?: pulumi.Input<number>;
    /**
     * Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     */
    readonly perVolumeGigabytes?: pulumi.Input<number>;
    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     */
    readonly projectId: pulumi.Input<string>;
    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Quota value for snapshots. Changing this updates the
     * existing quotaset.
     */
    readonly snapshots?: pulumi.Input<number>;
    /**
     * Quota value for volumes. Changing this updates the
     * existing quotaset.
     */
    readonly volumes?: pulumi.Input<number>;
}
