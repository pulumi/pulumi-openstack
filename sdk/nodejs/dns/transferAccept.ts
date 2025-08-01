// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a DNS zone transfer accept in the OpenStack DNS Service.
 *
 * ## Example Usage
 *
 * ### Automatically detect the correct network
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const exampleZone = new openstack.dns.Zone("example_zone", {
 *     name: "example.com.",
 *     email: "jdoe@example.com",
 *     description: "An example zone",
 *     ttl: 3000,
 *     type: "PRIMARY",
 * });
 * const request1 = new openstack.dns.TransferRequest("request_1", {
 *     zoneId: exampleZone.id,
 *     description: "a transfer accept",
 * });
 * const accept1 = new openstack.dns.TransferAccept("accept_1", {
 *     zoneTransferRequestId: request1.id,
 *     key: request1.key,
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported by specifying the transferAccept ID:
 *
 * ```sh
 * $ pulumi import openstack:dns/transferAccept:TransferAccept accept_1 accept_id
 * ```
 */
export class TransferAccept extends pulumi.CustomResource {
    /**
     * Get an existing TransferAccept resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransferAcceptState, opts?: pulumi.CustomResourceOptions): TransferAccept {
        return new TransferAccept(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:dns/transferAccept:TransferAccept';

    /**
     * Returns true if the given object is an instance of TransferAccept.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransferAccept {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransferAccept.__pulumiType;
    }

    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack accept returned success.
     */
    public readonly disableStatusCheck!: pulumi.Output<boolean | undefined>;
    /**
     * The transfer key.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone zone transfer accept.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Map of additional options. Changing this creates a
     * new transfer accept.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the zone transfer request.
     */
    public readonly zoneTransferRequestId!: pulumi.Output<string>;

    /**
     * Create a TransferAccept resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransferAcceptArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransferAcceptArgs | TransferAcceptState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TransferAcceptState | undefined;
            resourceInputs["disableStatusCheck"] = state ? state.disableStatusCheck : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
            resourceInputs["zoneTransferRequestId"] = state ? state.zoneTransferRequestId : undefined;
        } else {
            const args = argsOrState as TransferAcceptArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.zoneTransferRequestId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneTransferRequestId'");
            }
            resourceInputs["disableStatusCheck"] = args ? args.disableStatusCheck : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["zoneTransferRequestId"] = args ? args.zoneTransferRequestId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TransferAccept.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransferAccept resources.
 */
export interface TransferAcceptState {
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack accept returned success.
     */
    disableStatusCheck?: pulumi.Input<boolean>;
    /**
     * The transfer key.
     */
    key?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone zone transfer accept.
     */
    region?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new transfer accept.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the zone transfer request.
     */
    zoneTransferRequestId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TransferAccept resource.
 */
export interface TransferAcceptArgs {
    /**
     * Disable wait for zone to reach ACTIVE
     * status. The check is enabled by default. If this argument is true, zone
     * will be considered as created/updated if OpenStack accept returned success.
     */
    disableStatusCheck?: pulumi.Input<boolean>;
    /**
     * The transfer key.
     */
    key: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS zone zone transfer accept.
     */
    region?: pulumi.Input<string>;
    /**
     * Map of additional options. Changing this creates a
     * new transfer accept.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the zone transfer request.
     */
    zoneTransferRequestId: pulumi.Input<string>;
}
