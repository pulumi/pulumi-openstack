import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a DNS record set in the OpenStack DNS Service.
 */
export declare class RecordSet extends pulumi.CustomResource {
    /**
     * Get an existing RecordSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecordSetState): RecordSet;
    /**
     * A description of the  record set.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * The name of the record set. Note the `.` at the end of the name.
     * Changing this creates a new DNS  record set.
     */
    readonly name: pulumi.Output<string>;
    /**
     * An array of DNS records.
     */
    readonly records: pulumi.Output<string[] | undefined>;
    /**
     * The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS  record set.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The time to live (TTL) of the record set.
     */
    readonly ttl: pulumi.Output<number>;
    /**
     * The type of record set. Examples: "A", "MX".
     * Changing this creates a new DNS  record set.
     */
    readonly type: pulumi.Output<string>;
    /**
     * Map of additional options. Changing this creates a
     * new record set.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The ID of the zone in which to create the record set.
     * Changing this creates a new DNS  record set.
     */
    readonly zoneId: pulumi.Output<string>;
    /**
     * Create a RecordSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordSetArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering RecordSet resources.
 */
export interface RecordSetState {
    /**
     * A description of the  record set.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the record set. Note the `.` at the end of the name.
     * Changing this creates a new DNS  record set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An array of DNS records.
     */
    readonly records?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS  record set.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The time to live (TTL) of the record set.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of record set. Examples: "A", "MX".
     * Changing this creates a new DNS  record set.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new record set.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the zone in which to create the record set.
     * Changing this creates a new DNS  record set.
     */
    readonly zoneId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a RecordSet resource.
 */
export interface RecordSetArgs {
    /**
     * A description of the  record set.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the record set. Note the `.` at the end of the name.
     * Changing this creates a new DNS  record set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An array of DNS records.
     */
    readonly records?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS  record set.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The time to live (TTL) of the record set.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The type of record set. Examples: "A", "MX".
     * Changing this creates a new DNS  record set.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new record set.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the zone in which to create the record set.
     * Changing this creates a new DNS  record set.
     */
    readonly zoneId: pulumi.Input<string>;
}
