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
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const example_com = new openstack.dns.Zone("example.com", {
 *     description: "An example zone",
 *     email: "jdoe@example.com",
 *     ttl: 3000,
 *     type: "PRIMARY",
 * });
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ZoneState | undefined;
            inputs["attributes"] = state ? state.attributes : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["email"] = state ? state.email : undefined;
            inputs["masters"] = state ? state.masters : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            inputs["attributes"] = args ? args.attributes : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["email"] = args ? args.email : undefined;
            inputs["masters"] = args ? args.masters : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        super(Zone.__pulumiType, name, inputs, opts);
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
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * A description of the zone.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The email contact for the zone record.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     */
    readonly masters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The time to live (TTL) of the zone.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new zone.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     */
    readonly attributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * A description of the zone.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The email contact for the zone record.
     */
    readonly email?: pulumi.Input<string>;
    /**
     * An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     */
    readonly masters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The time to live (TTL) of the zone.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new zone.
     */
    readonly valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
