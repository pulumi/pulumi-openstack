// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * This resource can be imported by specifying the ID of the share and the ID of the share access, separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:sharedfilesystem/shareAccess:ShareAccess share_access_1 <share id>/<share access id>
 * ```
 */
export class ShareAccess extends pulumi.CustomResource {
    /**
     * Get an existing ShareAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShareAccessState, opts?: pulumi.CustomResourceOptions): ShareAccess {
        return new ShareAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:sharedfilesystem/shareAccess:ShareAccess';

    /**
     * Returns true if the given object is an instance of ShareAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShareAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShareAccess.__pulumiType;
    }

    /**
     * The access credential of the entity granted access.
     */
    public /*out*/ readonly accessKey!: pulumi.Output<string>;
    /**
     * The access level to the share. Can either be `rw` or `ro`.
     */
    public readonly accessLevel!: pulumi.Output<string>;
    /**
     * The value that defines the access. Can either be an IP
     * address or a username verified by configured Security Service of the Share Network.
     */
    public readonly accessTo!: pulumi.Output<string>;
    /**
     * The access rule type. Can either be an ip, user,
     * cert, or cephx. cephx support requires an OpenStack environment that supports
     * Shared Filesystem microversion 2.13 (Mitaka) or later.
     */
    public readonly accessType!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share access. Changing this
     * creates a new share access.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The UUID of the share to which you are granted access.
     */
    public readonly shareId!: pulumi.Output<string>;

    /**
     * Create a ShareAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShareAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShareAccessArgs | ShareAccessState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShareAccessState | undefined;
            inputs["accessKey"] = state ? state.accessKey : undefined;
            inputs["accessLevel"] = state ? state.accessLevel : undefined;
            inputs["accessTo"] = state ? state.accessTo : undefined;
            inputs["accessType"] = state ? state.accessType : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["shareId"] = state ? state.shareId : undefined;
        } else {
            const args = argsOrState as ShareAccessArgs | undefined;
            if ((!args || args.accessLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessLevel'");
            }
            if ((!args || args.accessTo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessTo'");
            }
            if ((!args || args.accessType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessType'");
            }
            if ((!args || args.shareId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shareId'");
            }
            inputs["accessLevel"] = args ? args.accessLevel : undefined;
            inputs["accessTo"] = args ? args.accessTo : undefined;
            inputs["accessType"] = args ? args.accessType : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["shareId"] = args ? args.shareId : undefined;
            inputs["accessKey"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ShareAccess.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShareAccess resources.
 */
export interface ShareAccessState {
    /**
     * The access credential of the entity granted access.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The access level to the share. Can either be `rw` or `ro`.
     */
    accessLevel?: pulumi.Input<string>;
    /**
     * The value that defines the access. Can either be an IP
     * address or a username verified by configured Security Service of the Share Network.
     */
    accessTo?: pulumi.Input<string>;
    /**
     * The access rule type. Can either be an ip, user,
     * cert, or cephx. cephx support requires an OpenStack environment that supports
     * Shared Filesystem microversion 2.13 (Mitaka) or later.
     */
    accessType?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share access. Changing this
     * creates a new share access.
     */
    region?: pulumi.Input<string>;
    /**
     * The UUID of the share to which you are granted access.
     */
    shareId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ShareAccess resource.
 */
export interface ShareAccessArgs {
    /**
     * The access level to the share. Can either be `rw` or `ro`.
     */
    accessLevel: pulumi.Input<string>;
    /**
     * The value that defines the access. Can either be an IP
     * address or a username verified by configured Security Service of the Share Network.
     */
    accessTo: pulumi.Input<string>;
    /**
     * The access rule type. Can either be an ip, user,
     * cert, or cephx. cephx support requires an OpenStack environment that supports
     * Shared Filesystem microversion 2.13 (Mitaka) or later.
     */
    accessType: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share access. Changing this
     * creates a new share access.
     */
    region?: pulumi.Input<string>;
    /**
     * The UUID of the share to which you are granted access.
     */
    shareId: pulumi.Input<string>;
}
