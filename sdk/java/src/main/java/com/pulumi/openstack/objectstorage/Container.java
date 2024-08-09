// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.objectstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.objectstorage.ContainerArgs;
import com.pulumi.openstack.objectstorage.inputs.ContainerState;
import com.pulumi.openstack.objectstorage.outputs.ContainerVersioningLegacy;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 container resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ### Basic Container
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
 *             .versioning(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Basic Container with legacy versioning
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
 * import com.pulumi.openstack.objectstorage.inputs.ContainerVersioningLegacyArgs;
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
 *             .versioningLegacy(ContainerVersioningLegacyArgs.builder()
 *                 .type("versions")
 *                 .location("tf-test-container-versions")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Global Read Access
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
 *         // Requires that a user know the object name they are attempting to download
 *         var container1 = new Container("container1", ContainerArgs.builder()
 *             .region("RegionOne")
 *             .name("tf-test-container-1")
 *             .containerRead(".r:*")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Global Read and List Access
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
 *         // Any user can read any object, and list all objects in the container
 *         var container1 = new Container("container1", ContainerArgs.builder()
 *             .region("RegionOne")
 *             .name("tf-test-container-1")
 *             .containerRead(".r:*,.rlistings")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Write-Only Access for a User
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.IdentityFunctions;
 * import com.pulumi.openstack.identity.inputs.GetAuthScopeArgs;
 * import com.pulumi.openstack.objectstorage.Container;
 * import com.pulumi.openstack.objectstorage.ContainerArgs;
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
 *         final var current = IdentityFunctions.getAuthScope(GetAuthScopeArgs.builder()
 *             .name("current")
 *             .build());
 * 
 *         // The named user can only upload objects, not read objects or list the container
 *         var container1 = new Container("container1", ContainerArgs.builder()
 *             .region("RegionOne")
 *             .name("tf-test-container-1")
 *             .containerRead(String.format(".r:-%s", username))
 *             .containerWrite(String.format("%s:%s", current.applyValue(getAuthScopeResult -> getAuthScopeResult.projectId()),username))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource can be imported by specifying the name of the container:
 * 
 * Some attributes can&#39;t be imported :
 * * `force_destroy`
 * * `content_type`
 * * `metadata`
 * * `container_sync_to`
 * * `container_sync_key`
 * 
 * So you&#39;ll have to `pulumi preview` and `pulumi up` after the import to fix those missing attributes.
 * 
 * ```sh
 * $ pulumi import openstack:objectstorage/container:Container container_1 container_name
 * ```
 * 
 */
@ResourceType(type="openstack:objectstorage/container:Container")
public class Container extends com.pulumi.resources.CustomResource {
    /**
     * Sets an access control list (ACL) that grants
     * read access. This header can contain a comma-delimited list of users that
     * can read the container (allows the GET method for all objects in the
     * container). Changing this updates the access control list read access.
     * 
     */
    @Export(name="containerRead", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> containerRead;

    /**
     * @return Sets an access control list (ACL) that grants
     * read access. This header can contain a comma-delimited list of users that
     * can read the container (allows the GET method for all objects in the
     * container). Changing this updates the access control list read access.
     * 
     */
    public Output<Optional<String>> containerRead() {
        return Codegen.optional(this.containerRead);
    }
    /**
     * The secret key for container synchronization.
     * Changing this updates container synchronization.
     * 
     */
    @Export(name="containerSyncKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> containerSyncKey;

    /**
     * @return The secret key for container synchronization.
     * Changing this updates container synchronization.
     * 
     */
    public Output<Optional<String>> containerSyncKey() {
        return Codegen.optional(this.containerSyncKey);
    }
    /**
     * The destination for container synchronization.
     * Changing this updates container synchronization.
     * 
     */
    @Export(name="containerSyncTo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> containerSyncTo;

    /**
     * @return The destination for container synchronization.
     * Changing this updates container synchronization.
     * 
     */
    public Output<Optional<String>> containerSyncTo() {
        return Codegen.optional(this.containerSyncTo);
    }
    /**
     * Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     * 
     */
    @Export(name="containerWrite", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> containerWrite;

    /**
     * @return Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     * 
     */
    public Output<Optional<String>> containerWrite() {
        return Codegen.optional(this.containerWrite);
    }
    /**
     * The MIME type for the container. Changing this
     * updates the MIME type.
     * 
     */
    @Export(name="contentType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentType;

    /**
     * @return The MIME type for the container. Changing this
     * updates the MIME type.
     * 
     */
    public Output<Optional<String>> contentType() {
        return Codegen.optional(this.contentType);
    }
    /**
     * A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> metadata;

    /**
     * @return Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     * 
     */
    public Output<Optional<Map<String,Object>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * A unique name for the container. Changing this creates a
     * new container.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the container. Changing this creates a
     * new container.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * The storage policy to be used for the container.
     * Changing this creates a new container.
     * 
     */
    @Export(name="storagePolicy", refs={String.class}, tree="[0]")
    private Output<String> storagePolicy;

    /**
     * @return The storage policy to be used for the container.
     * Changing this creates a new container.
     * 
     */
    public Output<String> storagePolicy() {
        return this.storagePolicy;
    }
    /**
     * A boolean that can enable or disable object
     * versioning. The default value is `false`. To use this feature, your Swift
     * version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
     * and a cloud administrator must have set the `allow_object_versioning = true`
     * configuration option in Swift. If you cannot set this versioning type, you may
     * want to consider using `versioning_legacy` instead.
     * 
     */
    @Export(name="versioning", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> versioning;

    /**
     * @return A boolean that can enable or disable object
     * versioning. The default value is `false`. To use this feature, your Swift
     * version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
     * and a cloud administrator must have set the `allow_object_versioning = true`
     * configuration option in Swift. If you cannot set this versioning type, you may
     * want to consider using `versioning_legacy` instead.
     * 
     */
    public Output<Optional<Boolean>> versioning() {
        return Codegen.optional(this.versioning);
    }
    /**
     * Enable legacy object versioning. The structure is described below.
     * 
     * @deprecated
     * Use newer &#34;versioning&#34; implementation
     * 
     */
    @Deprecated /* Use newer ""versioning"" implementation */
    @Export(name="versioningLegacy", refs={ContainerVersioningLegacy.class}, tree="[0]")
    private Output</* @Nullable */ ContainerVersioningLegacy> versioningLegacy;

    /**
     * @return Enable legacy object versioning. The structure is described below.
     * 
     */
    public Output<Optional<ContainerVersioningLegacy>> versioningLegacy() {
        return Codegen.optional(this.versioningLegacy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Container(java.lang.String name) {
        this(name, ContainerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Container(java.lang.String name, @Nullable ContainerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Container(java.lang.String name, @Nullable ContainerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:objectstorage/container:Container", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Container(java.lang.String name, Output<java.lang.String> id, @Nullable ContainerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:objectstorage/container:Container", name, state, makeResourceOptions(options, id), false);
    }

    private static ContainerArgs makeArgs(@Nullable ContainerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ContainerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static Container get(java.lang.String name, Output<java.lang.String> id, @Nullable ContainerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Container(name, id, state, options);
    }
}
