// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.images.ImageAccessArgs;
import com.pulumi.openstack.images.inputs.ImageAccessState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages members for the shared OpenStack Glance V2 Image within the source
 * project, which owns the Image.
 * 
 * ## Example Usage
 * 
 * ### Unprivileged user
 * 
 * Create a shared image and propose a membership to the
 * `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.images.Image;
 * import com.pulumi.openstack.images.ImageArgs;
 * import com.pulumi.openstack.images.ImageAccess;
 * import com.pulumi.openstack.images.ImageAccessArgs;
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
 *         var rancheros = new Image("rancheros", ImageArgs.builder()
 *             .name("RancherOS")
 *             .imageSourceUrl("https://releases.rancher.com/os/latest/rancheros-openstack.img")
 *             .containerFormat("bare")
 *             .diskFormat("qcow2")
 *             .visibility("shared")
 *             .properties(Map.of("key", "value"))
 *             .build());
 * 
 *         var rancherosMember = new ImageAccess("rancherosMember", ImageAccessArgs.builder()
 *             .imageId(rancheros.id())
 *             .memberId("bed6b6cbb86a4e2d8dc2735c2f1000e4")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Privileged user
 * 
 * Create a shared image and set a membership to the
 * `bed6b6cbb86a4e2d8dc2735c2f1000e4` project ID.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.images.Image;
 * import com.pulumi.openstack.images.ImageArgs;
 * import com.pulumi.openstack.images.ImageAccess;
 * import com.pulumi.openstack.images.ImageAccessArgs;
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
 *         var rancheros = new Image("rancheros", ImageArgs.builder()
 *             .name("RancherOS")
 *             .imageSourceUrl("https://releases.rancher.com/os/latest/rancheros-openstack.img")
 *             .containerFormat("bare")
 *             .diskFormat("qcow2")
 *             .visibility("shared")
 *             .properties(Map.of("key", "value"))
 *             .build());
 * 
 *         var rancherosMember = new ImageAccess("rancherosMember", ImageAccessArgs.builder()
 *             .imageId(rancheros.id())
 *             .memberId("bed6b6cbb86a4e2d8dc2735c2f1000e4")
 *             .status("accepted")
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
 * Image access can be imported using the `image_id` and the `member_id`,
 * 
 * separated by a slash, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:images/imageAccess:ImageAccess openstack_images_image_access_v2 89c60255-9bd6-460c-822a-e2b959ede9d2/bed6b6cbb86a4e2d8dc2735c2f1000e4
 * ```
 * 
 */
@ResourceType(type="openstack:images/imageAccess:ImageAccess")
public class ImageAccess extends com.pulumi.resources.CustomResource {
    /**
     * The date the image access was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date the image access was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The image ID.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return The image ID.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The member ID, e.g. the target project ID.
     * 
     */
    @Export(name="memberId", refs={String.class}, tree="[0]")
    private Output<String> memberId;

    /**
     * @return The member ID, e.g. the target project ID.
     * 
     */
    public Output<String> memberId() {
        return this.memberId;
    }
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to manage Image members. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new resource.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Glance client.
     * A Glance client is needed to manage Image members. If omitted, the `region`
     * argument of the provider is used. Changing this creates a new resource.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The member schema.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output<String> schema;

    /**
     * @return The member schema.
     * 
     */
    public Output<String> schema() {
        return this.schema;
    }
    /**
     * The member proposal status. Optional if admin wants to
     * force the member proposal acceptance. Can either be `accepted`, `rejected` or
     * `pending`. Defaults to `pending`. Foridden for non-admin users.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The member proposal status. Optional if admin wants to
     * force the member proposal acceptance. Can either be `accepted`, `rejected` or
     * `pending`. Defaults to `pending`. Foridden for non-admin users.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The date the image access was last updated.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The date the image access was last updated.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageAccess(String name) {
        this(name, ImageAccessArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageAccess(String name, ImageAccessArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageAccess(String name, ImageAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:images/imageAccess:ImageAccess", name, args == null ? ImageAccessArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ImageAccess(String name, Output<String> id, @Nullable ImageAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:images/imageAccess:ImageAccess", name, state, makeResourceOptions(options, id));
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
    public static ImageAccess get(String name, Output<String> id, @Nullable ImageAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageAccess(name, id, state, options);
    }
}
