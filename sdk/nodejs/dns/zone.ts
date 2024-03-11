// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DNS zone in the OpenStack DNS Service.
 *
 * ## Example Usage
 *
 * ### Automatically detect the correct network
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const exampleCom = new openstack.dns.Zone("exampleCom", {
 *     description: "An example zone",
 *     email: "jdoe@example.com",
 *     ttl: 3000,
 *     type: "PRIMARY",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * This resource can be imported by specifying the zone ID with optional project ID:
 *
 * ```sh
 * $ pulumi import openstack:dns/zone:Zone zone_1 zone_id
 * ```
 *
 * ```sh
 * $ pulumi import openstack:dns/zone:Zone zone_1 zone_id/project_id
 * ```
 */
export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:dns/zone:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     */
    public readonly attributes!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A description of the zone.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     */
    public readonly disableStatusCheck!: pulumi.Output<boolean | undefined>;
    /**
     * The email contact for the zone record.
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     */
    public readonly masters!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project DNS zone is created
     * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
     * user role in target project)
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The time to live (TTL) of the zone.
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Map of additional options. Changing this creates a
     * new zone.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneArgs | ZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableStatusCheck"] = state ? state.disableStatusCheck : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["masters"] = state ? state.masters : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableStatusCheck"] = args ? args.disableStatusCheck : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["masters"] = args ? args.masters : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Zone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     */
    attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * A description of the zone.
     */
    description?: pulumi.Input<string>;
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     */
    disableStatusCheck?: pulumi.Input<boolean>;
    /**
     * The email contact for the zone record.
     */
    email?: pulumi.Input<string>;
    /**
     * An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     */
    masters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project DNS zone is created
     * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
     * user role in target project)
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    region?: pulumi.Input<string>;
    /**
     * The time to live (TTL) of the zone.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     */
    type?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new zone.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     */
    attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * A description of the zone.
     */
    description?: pulumi.Input<string>;
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     */
    disableStatusCheck?: pulumi.Input<boolean>;
    /**
     * The email contact for the zone record.
     */
    email?: pulumi.Input<string>;
    /**
     * An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     */
    masters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project DNS zone is created
     * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
     * user role in target project)
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    region?: pulumi.Input<string>;
    /**
     * The time to live (TTL) of the zone.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     */
    type?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new zone.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
