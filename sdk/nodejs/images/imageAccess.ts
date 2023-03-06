// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages members for the shared OpenStack Glance V2 Image within the source
 * project, which owns the Image.
 *
 * ## Example Usage
 * ### Unprivileged user
 *
 * Create a shared image and propose a membership to the
 * `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const rancheros = new openstack.images.Image("rancheros", {
 *     containerFormat: "bare",
 *     diskFormat: "qcow2",
 *     imageSourceUrl: "https://releases.rancher.com/os/latest/rancheros-openstack.img",
 *     properties: {
 *         key: "value",
 *     },
 *     visibility: "shared",
 * });
 * const rancherosMember = new openstack.images.ImageAccess("rancherosMember", {
 *     imageId: rancheros.id,
 *     memberId: "bed6b6cbb86a4e2d8dc2735c2f1000e4",
 * });
 * ```
 * ### Privileged user
 *
 * Create a shared image and set a membership to the
 * `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const rancheros = new openstack.images.Image("rancheros", {
 *     containerFormat: "bare",
 *     diskFormat: "qcow2",
 *     imageSourceUrl: "https://releases.rancher.com/os/latest/rancheros-openstack.img",
 *     properties: {
 *         key: "value",
 *     },
 *     visibility: "shared",
 * });
 * const rancherosMember = new openstack.images.ImageAccess("rancherosMember", {
 *     imageId: rancheros.id,
 *     memberId: "bed6b6cbb86a4e2d8dc2735c2f1000e4",
 *     status: "accepted",
 * });
 * ```
 *
 * ## Import
 *
 * Image access can be imported using the `image_id` and the `member_id`, separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:images/imageAccess:ImageAccess openstack_images_image_access_v2 89c60255-9bd6-460c-822a-e2b959ede9d2/bed6b6cbb86a4e2d8dc2735c2f1000e4
 * ```
 */
export class ImageAccess extends pulumi.CustomResource {
    /**
     * Get an existing ImageAccess resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageAccessState, opts?: pulumi.CustomResourceOptions): ImageAccess {
        return new ImageAccess(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:images/imageAccess:ImageAccess';

    /**
     * Returns true if the given object is an instance of ImageAccess.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ImageAccess {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ImageAccess.__pulumiType;
    }

    /**
     * The date the image access was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The image ID.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The member ID, e.g. the target project ID.
     */
    public readonly memberId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to manage Image members. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The member schema.
     */
    public /*out*/ readonly schema!: pulumi.Output<string>;
    /**
     * The member proposal status. Optional if admin wants to
     * force the member proposal acceptance. Can either be `accepted`, `rejected` or
     * `pending`. Defaults to `pending`. Foridden for non-admin users.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The date the image access was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ImageAccess resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageAccessArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageAccessArgs | ImageAccessState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ImageAccessState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["memberId"] = state ? state.memberId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["schema"] = state ? state.schema : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ImageAccessArgs | undefined;
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.memberId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memberId'");
            }
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["memberId"] = args ? args.memberId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["schema"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ImageAccess.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ImageAccess resources.
 */
export interface ImageAccessState {
    /**
     * The date the image access was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The image ID.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The member ID, e.g. the target project ID.
     */
    memberId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to manage Image members. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The member schema.
     */
    schema?: pulumi.Input<string>;
    /**
     * The member proposal status. Optional if admin wants to
     * force the member proposal acceptance. Can either be `accepted`, `rejected` or
     * `pending`. Defaults to `pending`. Foridden for non-admin users.
     */
    status?: pulumi.Input<string>;
    /**
     * The date the image access was last updated.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ImageAccess resource.
 */
export interface ImageAccessArgs {
    /**
     * The image ID.
     */
    imageId: pulumi.Input<string>;
    /**
     * The member ID, e.g. the target project ID.
     */
    memberId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to manage Image members. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The member proposal status. Optional if admin wants to
     * force the member proposal acceptance. Can either be `accepted`, `rejected` or
     * `pending`. Defaults to `pending`. Foridden for non-admin users.
     */
    status?: pulumi.Input<string>;
}
