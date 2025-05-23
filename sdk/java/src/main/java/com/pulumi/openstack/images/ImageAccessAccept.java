// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.images.ImageAccessAcceptArgs;
import com.pulumi.openstack.images.inputs.ImageAccessAcceptState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages memberships status for the shared OpenStack Glance V2 Image within the
 * destination project, which has a member proposal.
 * 
 * ## Example Usage
 * 
 * Accept a shared image membershipship proposal within the current project.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.images.ImagesFunctions;
 * import com.pulumi.openstack.images.inputs.GetImageArgs;
 * import com.pulumi.openstack.images.ImageAccessAccept;
 * import com.pulumi.openstack.images.ImageAccessAcceptArgs;
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
 *         final var rancheros = ImagesFunctions.getImage(GetImageArgs.builder()
 *             .name("RancherOS")
 *             .visibility("shared")
 *             .memberStatus("all")
 *             .build());
 * 
 *         var rancherosMember = new ImageAccessAccept("rancherosMember", ImageAccessAcceptArgs.builder()
 *             .imageId(rancheros.id())
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
 * Image access acceptance status can be imported using the `image_id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:images/imageAccessAccept:ImageAccessAccept openstack_images_image_access_accept_v2 89c60255-9bd6-460c-822a-e2b959ede9d2
 * ```
 * 
 */
@ResourceType(type="openstack:images/imageAccessAccept:ImageAccessAccept")
public class ImageAccessAccept extends com.pulumi.resources.CustomResource {
    /**
     * The date the image membership was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date the image membership was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The proposed image ID.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output<String> imageId;

    /**
     * @return The proposed image ID.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The member ID, e.g. the target project ID. Optional
     * for admin accounts. Defaults to the current scope project ID.
     * 
     */
    @Export(name="memberId", refs={String.class}, tree="[0]")
    private Output<String> memberId;

    /**
     * @return The member ID, e.g. the target project ID. Optional
     * for admin accounts. Defaults to the current scope project ID.
     * 
     */
    public Output<String> memberId() {
        return this.memberId;
    }
    /**
     * The region in which to obtain the V2 Glance client.
     * A Glance client is needed to manage Image memberships. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * membership.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Glance client.
     * A Glance client is needed to manage Image memberships. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * membership.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The membership schema.
     * 
     */
    @Export(name="schema", refs={String.class}, tree="[0]")
    private Output<String> schema;

    /**
     * @return The membership schema.
     * 
     */
    public Output<String> schema() {
        return this.schema;
    }
    /**
     * The membership proposal status. Can either be
     * `accepted`, `rejected` or `pending`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The membership proposal status. Can either be
     * `accepted`, `rejected` or `pending`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The date the image membership was last updated.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The date the image membership was last updated.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageAccessAccept(java.lang.String name) {
        this(name, ImageAccessAcceptArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageAccessAccept(java.lang.String name, ImageAccessAcceptArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageAccessAccept(java.lang.String name, ImageAccessAcceptArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:images/imageAccessAccept:ImageAccessAccept", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ImageAccessAccept(java.lang.String name, Output<java.lang.String> id, @Nullable ImageAccessAcceptState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:images/imageAccessAccept:ImageAccessAccept", name, state, makeResourceOptions(options, id), false);
    }

    private static ImageAccessAcceptArgs makeArgs(ImageAccessAcceptArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ImageAccessAcceptArgs.Empty : args;
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
    public static ImageAccessAccept get(java.lang.String name, Output<java.lang.String> id, @Nullable ImageAccessAcceptState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageAccessAccept(name, id, state, options);
    }
}
