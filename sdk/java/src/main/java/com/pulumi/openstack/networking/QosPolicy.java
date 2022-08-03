// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.networking.QosPolicyArgs;
import com.pulumi.openstack.networking.inputs.QosPolicyState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron QoS policy resource within OpenStack.
 * 
 * ## Example Usage
 * ### Create a QoS Policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.QosPolicy;
 * import com.pulumi.openstack.networking.QosPolicyArgs;
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
 *         var qosPolicy1 = new QosPolicy(&#34;qosPolicy1&#34;, QosPolicyArgs.builder()        
 *             .description(&#34;bw_limit&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * QoS Policies can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:networking/qosPolicy:QosPolicy qos_policy_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae
 * ```
 * 
 */
@ResourceType(type="openstack:networking/qosPolicy:QosPolicy")
public class QosPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The collection of tags assigned on the QoS policy, which have been
     * explicitly and implicitly added.
     * 
     */
    @Export(name="allTags", type=List.class, parameters={String.class})
    private Output<List<String>> allTags;

    /**
     * @return The collection of tags assigned on the QoS policy, which have been
     * explicitly and implicitly added.
     * 
     */
    public Output<List<String>> allTags() {
        return this.allTags;
    }
    /**
     * The time at which QoS policy was created.
     * 
     */
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    /**
     * @return The time at which QoS policy was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The human-readable description for the QoS policy.
     * Changing this updates the description of the existing QoS policy.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The human-readable description for the QoS policy.
     * Changing this updates the description of the existing QoS policy.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Indicates whether the QoS policy is default
     * QoS policy or not. Changing this updates the default status of the existing
     * QoS policy.
     * 
     */
    @Export(name="isDefault", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isDefault;

    /**
     * @return Indicates whether the QoS policy is default
     * QoS policy or not. Changing this updates the default status of the existing
     * QoS policy.
     * 
     */
    public Output<Optional<Boolean>> isDefault() {
        return Codegen.optional(this.isDefault);
    }
    /**
     * The name of the QoS policy. Changing this updates the name of
     * the existing QoS policy.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the QoS policy. Changing this updates the name of
     * the existing QoS policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The owner of the QoS policy. Required if admin wants to
     * create a QoS policy for another project. Changing this creates a new QoS policy.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The owner of the QoS policy. Required if admin wants to
     * create a QoS policy for another project. Changing this creates a new QoS policy.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron Qos policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * QoS policy.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron Qos policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * QoS policy.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The revision number of the QoS policy.
     * 
     */
    @Export(name="revisionNumber", type=Integer.class, parameters={})
    private Output<Integer> revisionNumber;

    /**
     * @return The revision number of the QoS policy.
     * 
     */
    public Output<Integer> revisionNumber() {
        return this.revisionNumber;
    }
    /**
     * Indicates whether this QoS policy is shared across
     * all projects. Changing this updates the shared status of the existing
     * QoS policy.
     * 
     */
    @Export(name="shared", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> shared;

    /**
     * @return Indicates whether this QoS policy is shared across
     * all projects. Changing this updates the shared status of the existing
     * QoS policy.
     * 
     */
    public Output<Optional<Boolean>> shared() {
        return Codegen.optional(this.shared);
    }
    /**
     * A set of string tags for the QoS policy.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A set of string tags for the QoS policy.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The time at which QoS policy was created.
     * 
     */
    @Export(name="updatedAt", type=String.class, parameters={})
    private Output<String> updatedAt;

    /**
     * @return The time at which QoS policy was created.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * Map of additional options.
     * 
     */
    @Export(name="valueSpecs", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options.
     * 
     */
    public Output<Optional<Map<String,Object>>> valueSpecs() {
        return Codegen.optional(this.valueSpecs);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QosPolicy(String name) {
        this(name, QosPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QosPolicy(String name, @Nullable QosPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QosPolicy(String name, @Nullable QosPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosPolicy:QosPolicy", name, args == null ? QosPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QosPolicy(String name, Output<String> id, @Nullable QosPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:networking/qosPolicy:QosPolicy", name, state, makeResourceOptions(options, id));
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
    public static QosPolicy get(String name, Output<String> id, @Nullable QosPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QosPolicy(name, id, state, options);
    }
}
