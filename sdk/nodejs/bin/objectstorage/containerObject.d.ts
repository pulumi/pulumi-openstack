import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 container object resource within OpenStack.
 */
export declare class ContainerObject extends pulumi.CustomResource {
    /**
     * Get an existing ContainerObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerObjectState): ContainerObject;
    /**
     * A unique (within an account) name for the container.
     * The container name must be from 1 to 256 characters long and can start
     * with any character and contain any pattern. Character set must be UTF-8.
     * The container name cannot contain a slash (/) character because this
     * character delimits the container and object name. For example, the path
     * /v1/account/www/pages specifies the www container, not the www/pages container.
     */
    readonly containerName: pulumi.Output<string>;
    /**
     * A string representing the content of the object. Conflicts with
     * `source` and `copy_from`.
     */
    readonly content: pulumi.Output<string | undefined>;
    /**
     * A string which specifies the override behavior for
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     */
    readonly contentDisposition: pulumi.Output<string>;
    /**
     * A string representing the value of the Content-Encoding
     * metadata.
     */
    readonly contentEncoding: pulumi.Output<string>;
    /**
     * If the operation succeeds, this value is zero (0) or the
     * length of informational or error text in the response body.
     */
    readonly contentLength: pulumi.Output<number>;
    /**
     * A string which sets the MIME type for the object.
     */
    readonly contentType: pulumi.Output<string>;
    /**
     * A string representing the name of an object
     * used to create the new object by copying the `copy_from` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     */
    readonly copyFrom: pulumi.Output<string | undefined>;
    /**
     * The date and time the system responded to the request, using the preferred
     * format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
     * time is always in UTC.
     */
    readonly date: pulumi.Output<string>;
    /**
     * An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     */
    readonly deleteAfter: pulumi.Output<number | undefined>;
    /**
     * An string representing the date when the system removes the object.
     * For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     */
    readonly deleteAt: pulumi.Output<string>;
    /**
     * If set to true, Object Storage guesses the content
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     */
    readonly detectContentType: pulumi.Output<boolean | undefined>;
    /**
     * Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
     */
    readonly etag: pulumi.Output<string>;
    /**
     * The date and time when the object was last modified. The date and time
     * stamp format is ISO 8601:
     * CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00.
     * The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
     * example, the offset value is -05:00.
     */
    readonly lastModified: pulumi.Output<string>;
    readonly metadata: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * A unique name for the object.
     */
    readonly name: pulumi.Output<string>;
    /**
     * A string set to specify that this is a dynamic large
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     */
    readonly objectManifest: pulumi.Output<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    readonly region: pulumi.Output<string>;
    /**
     * A string representing the local path of a file which will be used
     * as the object's content. Conflicts with `source` and `copy_from`.
     */
    readonly source: pulumi.Output<string | undefined>;
    /**
     * A unique transaction ID for this request. Your service provider might
     * need this value if you report a problem.
     */
    readonly transId: pulumi.Output<string>;
    /**
     * Create a ContainerObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerObjectArgs, opts?: pulumi.ResourceOptions);
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
    readonly containerName?: pulumi.Input<string>;
    /**
     * A string representing the content of the object. Conflicts with
     * `source` and `copy_from`.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * A string which specifies the override behavior for
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * A string representing the value of the Content-Encoding
     * metadata.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * If the operation succeeds, this value is zero (0) or the
     * length of informational or error text in the response body.
     */
    readonly contentLength?: pulumi.Input<number>;
    /**
     * A string which sets the MIME type for the object.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * A string representing the name of an object
     * used to create the new object by copying the `copy_from` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     */
    readonly copyFrom?: pulumi.Input<string>;
    /**
     * The date and time the system responded to the request, using the preferred
     * format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
     * time is always in UTC.
     */
    readonly date?: pulumi.Input<string>;
    /**
     * An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     */
    readonly deleteAfter?: pulumi.Input<number>;
    /**
     * An string representing the date when the system removes the object.
     * For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     */
    readonly deleteAt?: pulumi.Input<string>;
    /**
     * If set to true, Object Storage guesses the content
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     */
    readonly detectContentType?: pulumi.Input<boolean>;
    /**
     * Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The date and time when the object was last modified. The date and time
     * stamp format is ISO 8601:
     * CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00.
     * The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
     * example, the offset value is -05:00.
     */
    readonly lastModified?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A unique name for the object.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A string set to specify that this is a dynamic large
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     */
    readonly objectManifest?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A string representing the local path of a file which will be used
     * as the object's content. Conflicts with `source` and `copy_from`.
     */
    readonly source?: pulumi.Input<string>;
    /**
     * A unique transaction ID for this request. Your service provider might
     * need this value if you report a problem.
     */
    readonly transId?: pulumi.Input<string>;
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
    readonly containerName: pulumi.Input<string>;
    /**
     * A string representing the content of the object. Conflicts with
     * `source` and `copy_from`.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * A string which specifies the override behavior for
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * A string representing the value of the Content-Encoding
     * metadata.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * A string which sets the MIME type for the object.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * A string representing the name of an object
     * used to create the new object by copying the `copy_from` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     */
    readonly copyFrom?: pulumi.Input<string>;
    /**
     * An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     */
    readonly deleteAfter?: pulumi.Input<number>;
    /**
     * An string representing the date when the system removes the object.
     * For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     */
    readonly deleteAt?: pulumi.Input<string>;
    /**
     * If set to true, Object Storage guesses the content
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     */
    readonly detectContentType?: pulumi.Input<boolean>;
    /**
     * Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
     */
    readonly etag?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A unique name for the object.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A string set to specify that this is a dynamic large
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     */
    readonly objectManifest?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A string representing the local path of a file which will be used
     * as the object's content. Conflicts with `source` and `copy_from`.
     */
    readonly source?: pulumi.Input<string>;
}
