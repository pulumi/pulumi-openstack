// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Neutron network segment resource within OpenStack.
 *
 * > **Note:** This resource is only available if the Neutron service is
 * configured with the `segments` extension.
 *
 * > **Note:** This ussually requires admin privileges to create or manage
 * segments.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const net1 = new openstack.networking.Network("net_1", {name: "demo-net"});
 * const segment1 = new openstack.networking.SegmentV2("segment_1", {
 *     name: "flat-segment",
 *     description: "Example flat segment",
 *     networkId: net1.id,
 *     networkType: "flat",
 *     physicalNetwork: "public",
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported by specifying the segment ID:
 *
 * ```sh
 * $ pulumi import openstack:networking/segmentV2:SegmentV2 segment1 a5e3a494-26ee-4fde-ad26-2d846c47072e
 * ```
 */
export class SegmentV2 extends pulumi.CustomResource {
    /**
     * Get an existing SegmentV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SegmentV2State, opts?: pulumi.CustomResourceOptions): SegmentV2 {
        return new SegmentV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/segmentV2:SegmentV2';

    /**
     * Returns true if the given object is an instance of SegmentV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SegmentV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SegmentV2.__pulumiType;
    }

    /**
     * Creation timestamp (RFC3339 format).
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A description for the segment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A name for the segment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The UUID of the network this segment belongs to.
     * Changing this will create a new segment.
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The network type. Valid values depend on the
     * backend (e.g., `vlan`, `vxlan`, `flat`, `gre`, `geneve`, `local`). Changing
     * this will create a new segment.
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * The name of the physical network. Changing this
     * will create a new segment.
     */
    public readonly physicalNetwork!: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * segment.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The revision number of the segment.
     */
    public /*out*/ readonly revisionNumber!: pulumi.Output<number>;
    /**
     * A segmentation identifier. Changing is allowed
     * only for `vlan`.
     */
    public readonly segmentationId!: pulumi.Output<number>;
    /**
     * Last update timestamp (RFC3339 format).
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a SegmentV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SegmentV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SegmentV2Args | SegmentV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SegmentV2State | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["physicalNetwork"] = state ? state.physicalNetwork : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["revisionNumber"] = state ? state.revisionNumber : undefined;
            resourceInputs["segmentationId"] = state ? state.segmentationId : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as SegmentV2Args | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.networkType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["physicalNetwork"] = args ? args.physicalNetwork : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["segmentationId"] = args ? args.segmentationId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["revisionNumber"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SegmentV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SegmentV2 resources.
 */
export interface SegmentV2State {
    /**
     * Creation timestamp (RFC3339 format).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * A description for the segment.
     */
    description?: pulumi.Input<string>;
    /**
     * A name for the segment.
     */
    name?: pulumi.Input<string>;
    /**
     * The UUID of the network this segment belongs to.
     * Changing this will create a new segment.
     */
    networkId?: pulumi.Input<string>;
    /**
     * The network type. Valid values depend on the
     * backend (e.g., `vlan`, `vxlan`, `flat`, `gre`, `geneve`, `local`). Changing
     * this will create a new segment.
     */
    networkType?: pulumi.Input<string>;
    /**
     * The name of the physical network. Changing this
     * will create a new segment.
     */
    physicalNetwork?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * segment.
     */
    region?: pulumi.Input<string>;
    /**
     * The revision number of the segment.
     */
    revisionNumber?: pulumi.Input<number>;
    /**
     * A segmentation identifier. Changing is allowed
     * only for `vlan`.
     */
    segmentationId?: pulumi.Input<number>;
    /**
     * Last update timestamp (RFC3339 format).
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SegmentV2 resource.
 */
export interface SegmentV2Args {
    /**
     * A description for the segment.
     */
    description?: pulumi.Input<string>;
    /**
     * A name for the segment.
     */
    name?: pulumi.Input<string>;
    /**
     * The UUID of the network this segment belongs to.
     * Changing this will create a new segment.
     */
    networkId: pulumi.Input<string>;
    /**
     * The network type. Valid values depend on the
     * backend (e.g., `vlan`, `vxlan`, `flat`, `gre`, `geneve`, `local`). Changing
     * this will create a new segment.
     */
    networkType: pulumi.Input<string>;
    /**
     * The name of the physical network. Changing this
     * will create a new segment.
     */
    physicalNetwork?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * segment.
     */
    region?: pulumi.Input<string>;
    /**
     * A segmentation identifier. Changing is allowed
     * only for `vlan`.
     */
    segmentationId?: pulumi.Input<number>;
}
