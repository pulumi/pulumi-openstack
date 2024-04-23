// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this resource to configure a share.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     name: "network_1",
 *     adminStateUp: true,
 * });
 * const subnet1 = new openstack.networking.Subnet("subnet_1", {
 *     name: "subnet_1",
 *     cidr: "192.168.199.0/24",
 *     ipVersion: 4,
 *     networkId: network1.id,
 * });
 * const sharenetwork1 = new openstack.sharedfilesystem.ShareNetwork("sharenetwork_1", {
 *     name: "test_sharenetwork",
 *     description: "test share network with security services",
 *     neutronNetId: network1.id,
 *     neutronSubnetId: subnet1.id,
 * });
 * const share1 = new openstack.sharedfilesystem.Share("share_1", {
 *     name: "nfs_share",
 *     description: "test share description",
 *     shareProto: "NFS",
 *     size: 1,
 *     shareNetworkId: sharenetwork1.id,
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported by specifying the ID of the share:
 *
 * ```sh
 * $ pulumi import openstack:sharedfilesystem/share:Share share_1 id
 * ```
 */
export class Share extends pulumi.CustomResource {
    /**
     * Get an existing Share resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShareState, opts?: pulumi.CustomResourceOptions): Share {
        return new Share(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:sharedfilesystem/share:Share';

    /**
     * Returns true if the given object is an instance of Share.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Share {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Share.__pulumiType;
    }

    /**
     * The map of metadata, assigned on the share, which has been
     * explicitly and implicitly added.
     */
    public /*out*/ readonly allMetadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * The share availability zone. Changing this creates a
     * new share.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The human-readable description for the share.
     * Changing this updates the description of the existing share.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A list of export locations. For example, when a share server
     * has more than one network interface, it can have multiple export locations.
     */
    public /*out*/ readonly exportLocations!: pulumi.Output<outputs.sharedfilesystem.ShareExportLocation[]>;
    /**
     * Indicates whether a share has replicas or not.
     */
    public /*out*/ readonly hasReplicas!: pulumi.Output<boolean>;
    /**
     * The share host name.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * The level of visibility for the share. Set to true to make
     * share public. Set to false to make it private. Default value is false. Changing this
     * updates the existing share.
     */
    public readonly isPublic!: pulumi.Output<boolean | undefined>;
    /**
     * One or more metadata key and value pairs as a dictionary of
     * strings.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The name of the share. Changing this updates the name
     * of the existing share.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The owner of the Share.
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share. Changing this
     * creates a new share.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The share replication type.
     */
    public /*out*/ readonly replicationType!: pulumi.Output<string>;
    /**
     * The UUID of a share network where the share server exists
     * or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
     * the shareNetworkId value from the snapshot is used. Changing this creates a new share.
     */
    public readonly shareNetworkId!: pulumi.Output<string>;
    /**
     * The share protocol - can either be NFS, CIFS,
     * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
     */
    public readonly shareProto!: pulumi.Output<string>;
    /**
     * The UUID of the share server.
     */
    public /*out*/ readonly shareServerId!: pulumi.Output<string>;
    /**
     * The share type name. If you omit this parameter, the default
     * share type is used.
     */
    public readonly shareType!: pulumi.Output<string>;
    /**
     * The share size, in GBs. The requested share size cannot be greater
     * than the allowed GB quota. Changing this resizes the existing share.
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * The UUID of the share's base snapshot. Changing this creates
     * a new share.
     */
    public readonly snapshotId!: pulumi.Output<string | undefined>;

    /**
     * Create a Share resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShareArgs | ShareState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShareState | undefined;
            resourceInputs["allMetadata"] = state ? state.allMetadata : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["exportLocations"] = state ? state.exportLocations : undefined;
            resourceInputs["hasReplicas"] = state ? state.hasReplicas : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["isPublic"] = state ? state.isPublic : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["replicationType"] = state ? state.replicationType : undefined;
            resourceInputs["shareNetworkId"] = state ? state.shareNetworkId : undefined;
            resourceInputs["shareProto"] = state ? state.shareProto : undefined;
            resourceInputs["shareServerId"] = state ? state.shareServerId : undefined;
            resourceInputs["shareType"] = state ? state.shareType : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["snapshotId"] = state ? state.snapshotId : undefined;
        } else {
            const args = argsOrState as ShareArgs | undefined;
            if ((!args || args.shareProto === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shareProto'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["isPublic"] = args ? args.isPublic : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["shareNetworkId"] = args ? args.shareNetworkId : undefined;
            resourceInputs["shareProto"] = args ? args.shareProto : undefined;
            resourceInputs["shareType"] = args ? args.shareType : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["allMetadata"] = undefined /*out*/;
            resourceInputs["exportLocations"] = undefined /*out*/;
            resourceInputs["hasReplicas"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["replicationType"] = undefined /*out*/;
            resourceInputs["shareServerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Share.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Share resources.
 */
export interface ShareState {
    /**
     * The map of metadata, assigned on the share, which has been
     * explicitly and implicitly added.
     */
    allMetadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The share availability zone. Changing this creates a
     * new share.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The human-readable description for the share.
     * Changing this updates the description of the existing share.
     */
    description?: pulumi.Input<string>;
    /**
     * A list of export locations. For example, when a share server
     * has more than one network interface, it can have multiple export locations.
     */
    exportLocations?: pulumi.Input<pulumi.Input<inputs.sharedfilesystem.ShareExportLocation>[]>;
    /**
     * Indicates whether a share has replicas or not.
     */
    hasReplicas?: pulumi.Input<boolean>;
    /**
     * The share host name.
     */
    host?: pulumi.Input<string>;
    /**
     * The level of visibility for the share. Set to true to make
     * share public. Set to false to make it private. Default value is false. Changing this
     * updates the existing share.
     */
    isPublic?: pulumi.Input<boolean>;
    /**
     * One or more metadata key and value pairs as a dictionary of
     * strings.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the share. Changing this updates the name
     * of the existing share.
     */
    name?: pulumi.Input<string>;
    /**
     * The owner of the Share.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share. Changing this
     * creates a new share.
     */
    region?: pulumi.Input<string>;
    /**
     * The share replication type.
     */
    replicationType?: pulumi.Input<string>;
    /**
     * The UUID of a share network where the share server exists
     * or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
     * the shareNetworkId value from the snapshot is used. Changing this creates a new share.
     */
    shareNetworkId?: pulumi.Input<string>;
    /**
     * The share protocol - can either be NFS, CIFS,
     * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
     */
    shareProto?: pulumi.Input<string>;
    /**
     * The UUID of the share server.
     */
    shareServerId?: pulumi.Input<string>;
    /**
     * The share type name. If you omit this parameter, the default
     * share type is used.
     */
    shareType?: pulumi.Input<string>;
    /**
     * The share size, in GBs. The requested share size cannot be greater
     * than the allowed GB quota. Changing this resizes the existing share.
     */
    size?: pulumi.Input<number>;
    /**
     * The UUID of the share's base snapshot. Changing this creates
     * a new share.
     */
    snapshotId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Share resource.
 */
export interface ShareArgs {
    /**
     * The share availability zone. Changing this creates a
     * new share.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The human-readable description for the share.
     * Changing this updates the description of the existing share.
     */
    description?: pulumi.Input<string>;
    /**
     * The level of visibility for the share. Set to true to make
     * share public. Set to false to make it private. Default value is false. Changing this
     * updates the existing share.
     */
    isPublic?: pulumi.Input<boolean>;
    /**
     * One or more metadata key and value pairs as a dictionary of
     * strings.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the share. Changing this updates the name
     * of the existing share.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share. Changing this
     * creates a new share.
     */
    region?: pulumi.Input<string>;
    /**
     * The UUID of a share network where the share server exists
     * or will be created. If `shareNetworkId` is not set and you provide a `snapshotId`,
     * the shareNetworkId value from the snapshot is used. Changing this creates a new share.
     */
    shareNetworkId?: pulumi.Input<string>;
    /**
     * The share protocol - can either be NFS, CIFS,
     * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
     */
    shareProto: pulumi.Input<string>;
    /**
     * The share type name. If you omit this parameter, the default
     * share type is used.
     */
    shareType?: pulumi.Input<string>;
    /**
     * The share size, in GBs. The requested share size cannot be greater
     * than the allowed GB quota. Changing this resizes the existing share.
     */
    size: pulumi.Input<number>;
    /**
     * The UUID of the share's base snapshot. Changing this creates
     * a new share.
     */
    snapshotId?: pulumi.Input<string>;
}
