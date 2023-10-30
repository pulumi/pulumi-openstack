// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.SecGroupArgs;
import com.pulumi.openstack.networking.inputs.SecGroupState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Security Groups can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:networking/secGroup:SecGroup secgroup_1 38809219-5e8a-4852-9139-6f461c90e8bc
 * ```
 * 
 */
@ResourceType(type="openstack:networking/secGroup:SecGroup")
public class SecGroup extends com.pulumi.resources.CustomResource {
    /**
     * The collection of tags assigned on the security group, which have
     * been explicitly and implicitly added.
     * 
     */
    @Export(name="allTags", type=List.class, parameters={String.class})
    private Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the security group, which have
     * been explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allTags() {
        return this.allTags;
    }
    /**
     * Whether or not to delete the default
     * egress security rules. This is `false` by default. See the below note
     * for more information.
     * 
     */
    @Export(name="deleteDefaultRules", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> deleteDefaultRules;

    /**
     * @return Whether or not to delete the default
     * egress security rules. This is `false` by default. See the below note
     * for more information.
     * 
     */
    public Output<Optional<Boolean>> deleteDefaultRules() {
        return Codegen.optional(this.deleteDefaultRules);
    }
    /**
     * A unique name for the security group.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output<String> description;

    /**
     * @return A unique name for the security group.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * A unique name for the security group.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return A unique name for the security group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * A set of string tags for the security group.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A set of string tags for the security group.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group.
     * 
     */
    @Export(name="tenantId", type=String.class, parameters={})
    private Output<String> tenantId;

    /**
     * @return The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecGroup(String name) {
        this(name, SecGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecGroup(String name, @Nullable SecGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecGroup(String name, @Nullable SecGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/secGroup:SecGroup", name, args == null ? SecGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecGroup(String name, Output<String> id, @Nullable SecGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/secGroup:SecGroup", name, state, makeResourceOptions(options, id));
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
    public static SecGroup get(String name, Output<String> id, @Nullable SecGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecGroup(name, id, state, options);
    }
}
