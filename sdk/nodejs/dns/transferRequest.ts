// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DNS zone transfer request in the OpenStack DNS Service.
 *
 * ## Example Usage
 * ### Automatically detect the correct network
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const exampleZone = new openstack.dns.Zone("exampleZone", {
 *     email: "jdoe@example.com",
 *     description: "An example zone",
 *     ttl: 3000,
 *     type: "PRIMARY",
 * });
 * const request1 = new openstack.dns.TransferRequest("request1", {
 *     zoneId: exampleZone.id,
 *     description: "a transfer request",
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported by specifying the transferRequest ID
 *
 * ```sh
 *  $ pulumi import openstack:dns/transferRequest:TransferRequest request_1 request_id
 * ```
 */
export class TransferRequest extends pulumi.CustomResource {
    /**
     * Get an existing TransferRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransferRequestState, opts?: pulumi.CustomResourceOptions): TransferRequest {
        return new TransferRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:dns/transferRequest:TransferRequest';

    /**
     * Returns true if the given object is an instance of TransferRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransferRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransferRequest.__pulumiType;
    }

    /**
     * A description of the zone tranfer request.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     */
    public readonly disableStatusCheck!: pulumi.Output<boolean | undefined>;
    public readonly key!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The target Project ID to transfer to.
     */
    public readonly targetProjectId!: pulumi.Output<string>;
    /**
     * Map of additional options. Changing this creates a
     * new transfer request.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the zone for which to create the transfer
     * request.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a TransferRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransferRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransferRequestArgs | TransferRequestState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransferRequestState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableStatusCheck"] = state ? state.disableStatusCheck : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["targetProjectId"] = state ? state.targetProjectId : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as TransferRequestArgs | undefined;
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableStatusCheck"] = args ? args.disableStatusCheck : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["targetProjectId"] = args ? args.targetProjectId : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransferRequest.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransferRequest resources.
 */
export interface TransferRequestState {
    /**
     * A description of the zone tranfer request.
     */
    description?: pulumi.Input<string>;
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     */
    disableStatusCheck?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    region?: pulumi.Input<string>;
    /**
     * The target Project ID to transfer to.
     */
    targetProjectId?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new transfer request.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the zone for which to create the transfer
     * request.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TransferRequest resource.
 */
export interface TransferRequestArgs {
    /**
     * A description of the zone tranfer request.
     */
    description?: pulumi.Input<string>;
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack request returned success.
     */
    disableStatusCheck?: pulumi.Input<boolean>;
    key?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone.
     */
    region?: pulumi.Input<string>;
    /**
     * The target Project ID to transfer to.
     */
    targetProjectId?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new transfer request.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the zone for which to create the transfer
     * request.
     */
    zoneId: pulumi.Input<string>;
}
