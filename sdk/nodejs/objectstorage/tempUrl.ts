// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to generate an OpenStack Object Storage temporary URL.
 *
 * The temporary URL will be valid for as long as TTL is set to (in seconds).
 * Once the URL has expired, it will no longer be valid, but the resource
 * will remain in place. If you wish to automatically regenerate a URL, set
 * the `regenerate` argument to `true`. This will create a new resource with
 * a new ID and URL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const container1 = new openstack.objectstorage.Container("container_1", {
 *     name: "test",
 *     metadata: {
 *         "Temp-URL-Key": "testkey",
 *     },
 * });
 * const object1 = new openstack.objectstorage.ContainerObject("object_1", {
 *     containerName: container1.name,
 *     name: "test",
 *     content: "Hello, world!",
 * });
 * const objTempurl = new openstack.objectstorage.TempUrl("obj_tempurl", {
 *     container: container1.name,
 *     object: object1.name,
 *     method: "post",
 *     ttl: 20,
 * });
 * ```
 */
export class TempUrl extends pulumi.CustomResource {
    /**
     * Get an existing TempUrl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TempUrlState, opts?: pulumi.CustomResourceOptions): TempUrl {
        return new TempUrl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:objectstorage/tempUrl:TempUrl';

    /**
     * Returns true if the given object is an instance of TempUrl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TempUrl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TempUrl.__pulumiType;
    }

    /**
     * The container name the object belongs to.
     */
    public readonly container!: pulumi.Output<string>;
    /**
     * The digest to use when generating the tempurl.
     * Supported values are `sha1`, `sha256` and `sha512`. Default is `sha1`.
     */
    public readonly digest!: pulumi.Output<string | undefined>;
    /**
     * The key to use when generating the tempurl. If not
     * provided, the key will be read from the container or account metadata.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * The method allowed when accessing this URL.
     * Valid values are `GET`, and `POST`. Default is `GET`.
     */
    public readonly method!: pulumi.Output<string | undefined>;
    /**
     * The object name the tempurl is for.
     */
    public readonly object!: pulumi.Output<string>;
    /**
     * Whether to automatically regenerate the URL when
     * it has expired. If set to true, this will create a new resource with a new
     * ID and new URL. Defaults to false.
     */
    public readonly regenerate!: pulumi.Output<boolean | undefined>;
    /**
     * The region the tempurl is located in.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Split is the string on which to split the object URL.
     * Default is `/v1/`.
     */
    public readonly split!: pulumi.Output<string | undefined>;
    /**
     * The TTL, in seconds, for the URL. For how long it should
     * be valid.
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * The URL
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a TempUrl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TempUrlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TempUrlArgs | TempUrlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TempUrlState | undefined;
            resourceInputs["container"] = state ? state.container : undefined;
            resourceInputs["digest"] = state ? state.digest : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["method"] = state ? state.method : undefined;
            resourceInputs["object"] = state ? state.object : undefined;
            resourceInputs["regenerate"] = state ? state.regenerate : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["split"] = state ? state.split : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as TempUrlArgs | undefined;
            if ((!args || args.container === undefined) && !opts.urn) {
                throw new Error("Missing required property 'container'");
            }
            if ((!args || args.object === undefined) && !opts.urn) {
                throw new Error("Missing required property 'object'");
            }
            if ((!args || args.ttl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ttl'");
            }
            resourceInputs["container"] = args ? args.container : undefined;
            resourceInputs["digest"] = args ? args.digest : undefined;
            resourceInputs["key"] = args?.key ? pulumi.secret(args.key) : undefined;
            resourceInputs["method"] = args ? args.method : undefined;
            resourceInputs["object"] = args ? args.object : undefined;
            resourceInputs["regenerate"] = args ? args.regenerate : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["split"] = args ? args.split : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["key", "url"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(TempUrl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TempUrl resources.
 */
export interface TempUrlState {
    /**
     * The container name the object belongs to.
     */
    container?: pulumi.Input<string>;
    /**
     * The digest to use when generating the tempurl.
     * Supported values are `sha1`, `sha256` and `sha512`. Default is `sha1`.
     */
    digest?: pulumi.Input<string>;
    /**
     * The key to use when generating the tempurl. If not
     * provided, the key will be read from the container or account metadata.
     */
    key?: pulumi.Input<string>;
    /**
     * The method allowed when accessing this URL.
     * Valid values are `GET`, and `POST`. Default is `GET`.
     */
    method?: pulumi.Input<string>;
    /**
     * The object name the tempurl is for.
     */
    object?: pulumi.Input<string>;
    /**
     * Whether to automatically regenerate the URL when
     * it has expired. If set to true, this will create a new resource with a new
     * ID and new URL. Defaults to false.
     */
    regenerate?: pulumi.Input<boolean>;
    /**
     * The region the tempurl is located in.
     */
    region?: pulumi.Input<string>;
    /**
     * Split is the string on which to split the object URL.
     * Default is `/v1/`.
     */
    split?: pulumi.Input<string>;
    /**
     * The TTL, in seconds, for the URL. For how long it should
     * be valid.
     */
    ttl?: pulumi.Input<number>;
    /**
     * The URL
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TempUrl resource.
 */
export interface TempUrlArgs {
    /**
     * The container name the object belongs to.
     */
    container: pulumi.Input<string>;
    /**
     * The digest to use when generating the tempurl.
     * Supported values are `sha1`, `sha256` and `sha512`. Default is `sha1`.
     */
    digest?: pulumi.Input<string>;
    /**
     * The key to use when generating the tempurl. If not
     * provided, the key will be read from the container or account metadata.
     */
    key?: pulumi.Input<string>;
    /**
     * The method allowed when accessing this URL.
     * Valid values are `GET`, and `POST`. Default is `GET`.
     */
    method?: pulumi.Input<string>;
    /**
     * The object name the tempurl is for.
     */
    object: pulumi.Input<string>;
    /**
     * Whether to automatically regenerate the URL when
     * it has expired. If set to true, this will create a new resource with a new
     * ID and new URL. Defaults to false.
     */
    regenerate?: pulumi.Input<boolean>;
    /**
     * The region the tempurl is located in.
     */
    region?: pulumi.Input<string>;
    /**
     * Split is the string on which to split the object URL.
     * Default is `/v1/`.
     */
    split?: pulumi.Input<string>;
    /**
     * The TTL, in seconds, for the URL. For how long it should
     * be valid.
     */
    ttl: pulumi.Input<number>;
}
