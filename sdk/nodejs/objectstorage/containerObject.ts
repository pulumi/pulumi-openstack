// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 container object resource within OpenStack.
 *
 * ## Example Usage
 */
export class ContainerObject extends pulumi.CustomResource {
    /**
     * Get an existing ContainerObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerObjectState, opts?: pulumi.CustomResourceOptions): ContainerObject {
        return new ContainerObject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:objectstorage/containerObject:ContainerObject';

    /**
     * Returns true if the given object is an instance of ContainerObject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerObject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerObject.__pulumiType;
    }

    /**
     * A unique (within an account) name for the container. 
     * The container name must be from 1 to 256 characters long and can start
     * with any character and contain any pattern. Character set must be UTF-8.
     * The container name cannot contain a slash (/) character because this
     * character delimits the container and object name. For example, the path
     * /v1/account/www/pages specifies the www container, not the www/pages container.
     */
    public readonly containerName!: pulumi.Output<string>;
    /**
     * A string representing the content of the object. Conflicts with
     * `source` and `copyFrom`.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * A string which specifies the override behavior for 
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     */
    public readonly contentDisposition!: pulumi.Output<string>;
    /**
     * A string representing the value of the Content-Encoding
     * metadata.
     */
    public readonly contentEncoding!: pulumi.Output<string>;
    /**
     * If the operation succeeds, this value is zero (0) or the 
     * length of informational or error text in the response body.
     */
    public /*out*/ readonly contentLength!: pulumi.Output<number>;
    /**
     * A string which sets the MIME type for the object.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * A string representing the name of an object 
     * used to create the new object by copying the `copyFrom` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     */
    public readonly copyFrom!: pulumi.Output<string | undefined>;
    /**
     * The date and time the system responded to the request, using the preferred 
     * format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
     * time is always in UTC.
     */
    public /*out*/ readonly date!: pulumi.Output<string>;
    /**
     * An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     */
    public readonly deleteAfter!: pulumi.Output<number | undefined>;
    /**
     * An string representing the date when the system removes the object. 
     * For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     */
    public readonly deleteAt!: pulumi.Output<string>;
    /**
     * If set to true, Object Storage guesses the content 
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     */
    public readonly detectContentType!: pulumi.Output<boolean | undefined>;
    /**
     * Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The date and time when the object was last modified. The date and time 
     * stamp format is ISO 8601:
     * CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00.
     * The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
     * example, the offset value is -05:00.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique name for the object.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A string set to specify that this is a dynamic large 
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     */
    public readonly objectManifest!: pulumi.Output<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A string representing the local path of a file which will be used
     * as the object's content. Conflicts with `source` and `copyFrom`.
     */
    public readonly source!: pulumi.Output<string | undefined>;
    /**
     * A unique transaction ID for this request. Your service provider might 
     * need this value if you report a problem.
     */
    public /*out*/ readonly transId!: pulumi.Output<string>;

    /**
     * Create a ContainerObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerObjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerObjectArgs | ContainerObjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerObjectState | undefined;
            resourceInputs["containerName"] = state ? state.containerName : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["contentDisposition"] = state ? state.contentDisposition : undefined;
            resourceInputs["contentEncoding"] = state ? state.contentEncoding : undefined;
            resourceInputs["contentLength"] = state ? state.contentLength : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["copyFrom"] = state ? state.copyFrom : undefined;
            resourceInputs["date"] = state ? state.date : undefined;
            resourceInputs["deleteAfter"] = state ? state.deleteAfter : undefined;
            resourceInputs["deleteAt"] = state ? state.deleteAt : undefined;
            resourceInputs["detectContentType"] = state ? state.detectContentType : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["lastModified"] = state ? state.lastModified : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["objectManifest"] = state ? state.objectManifest : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
            resourceInputs["transId"] = state ? state.transId : undefined;
        } else {
            const args = argsOrState as ContainerObjectArgs | undefined;
            if ((!args || args.containerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerName'");
            }
            resourceInputs["containerName"] = args ? args.containerName : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["contentDisposition"] = args ? args.contentDisposition : undefined;
            resourceInputs["contentEncoding"] = args ? args.contentEncoding : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["copyFrom"] = args ? args.copyFrom : undefined;
            resourceInputs["deleteAfter"] = args ? args.deleteAfter : undefined;
            resourceInputs["deleteAt"] = args ? args.deleteAt : undefined;
            resourceInputs["detectContentType"] = args ? args.detectContentType : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["objectManifest"] = args ? args.objectManifest : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["contentLength"] = undefined /*out*/;
            resourceInputs["date"] = undefined /*out*/;
            resourceInputs["lastModified"] = undefined /*out*/;
            resourceInputs["transId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerObject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerObject resources.
 */
export interface ContainerObjectState {
    /**
     * A unique (within an account) name for the container. 
     * The container name must be from 1 to 256 characters long and can start
     * with any character and contain any pattern. Character set must be UTF-8.
     * The container name cannot contain a slash (/) character because this
     * character delimits the container and object name. For example, the path
     * /v1/account/www/pages specifies the www container, not the www/pages container.
     */
    containerName?: pulumi.Input<string>;
    /**
     * A string representing the content of the object. Conflicts with
     * `source` and `copyFrom`.
     */
    content?: pulumi.Input<string>;
    /**
     * A string which specifies the override behavior for 
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     */
    contentDisposition?: pulumi.Input<string>;
    /**
     * A string representing the value of the Content-Encoding
     * metadata.
     */
    contentEncoding?: pulumi.Input<string>;
    /**
     * If the operation succeeds, this value is zero (0) or the 
     * length of informational or error text in the response body.
     */
    contentLength?: pulumi.Input<number>;
    /**
     * A string which sets the MIME type for the object.
     */
    contentType?: pulumi.Input<string>;
    /**
     * A string representing the name of an object 
     * used to create the new object by copying the `copyFrom` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     */
    copyFrom?: pulumi.Input<string>;
    /**
     * The date and time the system responded to the request, using the preferred 
     * format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
     * time is always in UTC.
     */
    date?: pulumi.Input<string>;
    /**
     * An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     */
    deleteAfter?: pulumi.Input<number>;
    /**
     * An string representing the date when the system removes the object. 
     * For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     */
    deleteAt?: pulumi.Input<string>;
    /**
     * If set to true, Object Storage guesses the content 
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     */
    detectContentType?: pulumi.Input<boolean>;
    /**
     * Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
     */
    etag?: pulumi.Input<string>;
    /**
     * The date and time when the object was last modified. The date and time 
     * stamp format is ISO 8601:
     * CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00.
     * The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
     * example, the offset value is -05:00.
     */
    lastModified?: pulumi.Input<string>;
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the object.
     */
    name?: pulumi.Input<string>;
    /**
     * A string set to specify that this is a dynamic large 
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     */
    objectManifest?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    region?: pulumi.Input<string>;
    /**
     * A string representing the local path of a file which will be used
     * as the object's content. Conflicts with `source` and `copyFrom`.
     */
    source?: pulumi.Input<string>;
    /**
     * A unique transaction ID for this request. Your service provider might 
     * need this value if you report a problem.
     */
    transId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerObject resource.
 */
export interface ContainerObjectArgs {
    /**
     * A unique (within an account) name for the container. 
     * The container name must be from 1 to 256 characters long and can start
     * with any character and contain any pattern. Character set must be UTF-8.
     * The container name cannot contain a slash (/) character because this
     * character delimits the container and object name. For example, the path
     * /v1/account/www/pages specifies the www container, not the www/pages container.
     */
    containerName: pulumi.Input<string>;
    /**
     * A string representing the content of the object. Conflicts with
     * `source` and `copyFrom`.
     */
    content?: pulumi.Input<string>;
    /**
     * A string which specifies the override behavior for 
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     */
    contentDisposition?: pulumi.Input<string>;
    /**
     * A string representing the value of the Content-Encoding
     * metadata.
     */
    contentEncoding?: pulumi.Input<string>;
    /**
     * A string which sets the MIME type for the object.
     */
    contentType?: pulumi.Input<string>;
    /**
     * A string representing the name of an object 
     * used to create the new object by copying the `copyFrom` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     */
    copyFrom?: pulumi.Input<string>;
    /**
     * An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     */
    deleteAfter?: pulumi.Input<number>;
    /**
     * An string representing the date when the system removes the object. 
     * For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     */
    deleteAt?: pulumi.Input<string>;
    /**
     * If set to true, Object Storage guesses the content 
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     */
    detectContentType?: pulumi.Input<boolean>;
    /**
     * Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
     */
    etag?: pulumi.Input<string>;
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the object.
     */
    name?: pulumi.Input<string>;
    /**
     * A string set to specify that this is a dynamic large 
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     */
    objectManifest?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    region?: pulumi.Input<string>;
    /**
     * A string representing the local path of a file which will be used
     * as the object's content. Conflicts with `source` and `copyFrom`.
     */
    source?: pulumi.Input<string>;
}
