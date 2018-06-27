import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a DNS zone in the OpenStack DNS Service.
 */
export declare class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState): Zone;
    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     */
    readonly attributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * A description of the zone.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * The email contact for the zone record.
     */
    readonly email: pulumi.Output<string | undefined>;
    /**
     * An array of master DNS servers. For when `type` is
     * `SECONDARY`.
     */
    readonly masters: pulumi.Output<string[] | undefined>;
    /**
     * The name of the zone. Note the `.` at the end of the name.
     * Changing this creates a new DNS zone.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    readonly region: pulumi.Output<string>;
    /**
     * The time to live (TTL) of the zone.
     */
    readonly ttl: pulumi.Output<number>;
    /**
     * The type of zone. Can either be `PRIMARY` or `SECONDARY`.
     * Changing this creates a new zone.
     */
    readonly type: pulumi.Output<string>;
    /**
     * Map of additional options. Changing this creates a
     * new zone.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ZoneArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     */
    readonly attributes?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * Attributes for the DNS Service scheduler.
     * Changing this creates a new zone.
     */
    readonly attributes?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
