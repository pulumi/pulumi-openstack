// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.objectstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.objectstorage.ContainerObjectArgs;
import com.pulumi.openstack.objectstorage.inputs.ContainerObjectState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 container object resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### Example with simple content
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.objectstorage.Container;
 * import com.pulumi.openstack.objectstorage.ContainerArgs;
 * import com.pulumi.openstack.objectstorage.ContainerObject;
 * import com.pulumi.openstack.objectstorage.ContainerObjectArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var container1 = new Container("container1", ContainerArgs.builder()
 *             .region("RegionOne")
 *             .name("tf-test-container-1")
 *             .metadata(Map.of("test", "true"))
 *             .contentType("application/json")
 *             .build());
 * 
 *         var doc1 = new ContainerObject("doc1", ContainerObjectArgs.builder()
 *             .region("RegionOne")
 *             .containerName(container1.name())
 *             .name("test/default.json")
 *             .metadata(Map.of("test", "true"))
 *             .contentType("application/json")
 *             .content("""
 *                {
 *                  "foo" : "bar"
 *                }
 *             """)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Example with content from file
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.objectstorage.Container;
 * import com.pulumi.openstack.objectstorage.ContainerArgs;
 * import com.pulumi.openstack.objectstorage.ContainerObject;
 * import com.pulumi.openstack.objectstorage.ContainerObjectArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var container1 = new Container("container1", ContainerArgs.builder()
 *             .region("RegionOne")
 *             .name("tf-test-container-1")
 *             .metadata(Map.of("test", "true"))
 *             .contentType("application/json")
 *             .build());
 * 
 *         var doc1 = new ContainerObject("doc1", ContainerObjectArgs.builder()
 *             .region("RegionOne")
 *             .containerName(container1.name())
 *             .name("test/default.json")
 *             .metadata(Map.of("test", "true"))
 *             .contentType("application/json")
 *             .source("./default.json")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="openstack:objectstorage/containerObject:ContainerObject")
public class ContainerObject extends com.pulumi.resources.CustomResource {
    /**
     * A unique (within an account) name for the container.
     * The container name must be from 1 to 256 characters long and can start
     * with any character and contain any pattern. Character set must be UTF-8.
     * The container name cannot contain a slash (/) character because this
     * character delimits the container and object name. For example, the path
     * /v1/account/www/pages specifies the www container, not the www/pages container.
     * 
     */
    @Export(name="containerName", refs={String.class}, tree="[0]")
    private Output<String> containerName;

    /**
     * @return A unique (within an account) name for the container.
     * The container name must be from 1 to 256 characters long and can start
     * with any character and contain any pattern. Character set must be UTF-8.
     * The container name cannot contain a slash (/) character because this
     * character delimits the container and object name. For example, the path
     * /v1/account/www/pages specifies the www container, not the www/pages container.
     * 
     */
    public Output<String> containerName() {
        return this.containerName;
    }
    /**
     * A string representing the content of the object. Conflicts with
     * `source` and `copy_from`.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> content;

    /**
     * @return A string representing the content of the object. Conflicts with
     * `source` and `copy_from`.
     * 
     */
    public Output<Optional<String>> content() {
        return Codegen.optional(this.content);
    }
    /**
     * A string which specifies the override behavior for
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     * 
     */
    @Export(name="contentDisposition", refs={String.class}, tree="[0]")
    private Output<String> contentDisposition;

    /**
     * @return A string which specifies the override behavior for
     * the browser. For example, this header might specify that the browser use a download
     * program to save this file rather than show the file, which is the default.
     * 
     */
    public Output<String> contentDisposition() {
        return this.contentDisposition;
    }
    /**
     * A string representing the value of the Content-Encoding
     * metadata.
     * 
     */
    @Export(name="contentEncoding", refs={String.class}, tree="[0]")
    private Output<String> contentEncoding;

    /**
     * @return A string representing the value of the Content-Encoding
     * metadata.
     * 
     */
    public Output<String> contentEncoding() {
        return this.contentEncoding;
    }
    /**
     * If the operation succeeds, this value is zero (0) or the
     * length of informational or error text in the response body.
     * 
     */
    @Export(name="contentLength", refs={Integer.class}, tree="[0]")
    private Output<Integer> contentLength;

    /**
     * @return If the operation succeeds, this value is zero (0) or the
     * length of informational or error text in the response body.
     * 
     */
    public Output<Integer> contentLength() {
        return this.contentLength;
    }
    /**
     * A string which sets the MIME type for the object.
     * 
     */
    @Export(name="contentType", refs={String.class}, tree="[0]")
    private Output<String> contentType;

    /**
     * @return A string which sets the MIME type for the object.
     * 
     */
    public Output<String> contentType() {
        return this.contentType;
    }
    /**
     * A string representing the name of an object
     * used to create the new object by copying the `copy_from` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     * 
     */
    @Export(name="copyFrom", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> copyFrom;

    /**
     * @return A string representing the name of an object
     * used to create the new object by copying the `copy_from` object. The value is in form
     * {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
     * container and object before you include them in the header. Conflicts with `source` and
     * `content`.
     * 
     */
    public Output<Optional<String>> copyFrom() {
        return Codegen.optional(this.copyFrom);
    }
    /**
     * The date and time the system responded to the request, using the preferred
     * format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
     * time is always in UTC.
     * 
     */
    @Export(name="date", refs={String.class}, tree="[0]")
    private Output<String> date;

    /**
     * @return The date and time the system responded to the request, using the preferred
     * format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
     * time is always in UTC.
     * 
     */
    public Output<String> date() {
        return this.date;
    }
    /**
     * An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     * 
     */
    @Export(name="deleteAfter", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> deleteAfter;

    /**
     * @return An integer representing the number of seconds after which the
     * system removes the object. Internally, the Object Storage system stores this value in
     * the X-Delete-At metadata item.
     * 
     */
    public Output<Optional<Integer>> deleteAfter() {
        return Codegen.optional(this.deleteAfter);
    }
    /**
     * An string representing the date when the system removes the object.
     * For example, &#34;2015-08-26&#34; is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     * 
     */
    @Export(name="deleteAt", refs={String.class}, tree="[0]")
    private Output<String> deleteAt;

    /**
     * @return An string representing the date when the system removes the object.
     * For example, &#34;2015-08-26&#34; is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
     * 
     */
    public Output<String> deleteAt() {
        return this.deleteAt;
    }
    /**
     * If set to true, Object Storage guesses the content
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     * 
     */
    @Export(name="detectContentType", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> detectContentType;

    /**
     * @return If set to true, Object Storage guesses the content
     * type based on the file extension and ignores the value sent in the Content-Type
     * header, if present.
     * 
     */
    public Output<Optional<Boolean>> detectContentType() {
        return Codegen.optional(this.detectContentType);
    }
    /**
     * Used to trigger updates. The only meaningful value is ${md5(file(&#34;path/to/file&#34;))}.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return Used to trigger updates. The only meaningful value is ${md5(file(&#34;path/to/file&#34;))}.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The date and time when the object was last modified. The date and time
     * stamp format is ISO 8601:
     * CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00.
     * The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
     * example, the offset value is -05:00.
     * 
     */
    @Export(name="lastModified", refs={String.class}, tree="[0]")
    private Output<String> lastModified;

    /**
     * @return The date and time when the object was last modified. The date and time
     * stamp format is ISO 8601:
     * CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00.
     * The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
     * example, the offset value is -05:00.
     * 
     */
    public Output<String> lastModified() {
        return this.lastModified;
    }
    @Export(name="metadata", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> metadata;

    public Output<Optional<Map<String,Object>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * A unique name for the object.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the object.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A string set to specify that this is a dynamic large
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     * 
     */
    @Export(name="objectManifest", refs={String.class}, tree="[0]")
    private Output<String> objectManifest;

    /**
     * @return A string set to specify that this is a dynamic large
     * object manifest object. The value is the container and object name prefix of the
     * segment objects in the form container/prefix. You must UTF-8-encode and then
     * URL-encode the names of the container and prefix before you include them in this
     * header.
     * 
     */
    public Output<String> objectManifest() {
        return this.objectManifest;
    }
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A string representing the local path of a file which will be used
     * as the object&#39;s content. Conflicts with `source` and `copy_from`.
     * 
     */
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> source;

    /**
     * @return A string representing the local path of a file which will be used
     * as the object&#39;s content. Conflicts with `source` and `copy_from`.
     * 
     */
    public Output<Optional<String>> source() {
        return Codegen.optional(this.source);
    }
    /**
     * A unique transaction ID for this request. Your service provider might
     * need this value if you report a problem.
     * 
     */
    @Export(name="transId", refs={String.class}, tree="[0]")
    private Output<String> transId;

    /**
     * @return A unique transaction ID for this request. Your service provider might
     * need this value if you report a problem.
     * 
     */
    public Output<String> transId() {
        return this.transId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerObject(String name) {
        this(name, ContainerObjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerObject(String name, ContainerObjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerObject(String name, ContainerObjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:objectstorage/containerObject:ContainerObject", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerObject(String name, Output<String> id, @Nullable ContainerObjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:objectstorage/containerObject:ContainerObject", name, state, makeResourceOptions(options, id));
    }

    private static ContainerObjectArgs makeArgs(ContainerObjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ContainerObjectArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ContainerObject get(String name, Output<String> id, @Nullable ContainerObjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerObject(name, id, state, options);
    }
}
