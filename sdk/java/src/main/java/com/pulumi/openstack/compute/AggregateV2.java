// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.AggregateV2Args;
import com.pulumi.openstack.compute.inputs.AggregateV2State;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a Host Aggregate within Openstack Nova.
 * 
 * ## Example Usage
 * 
 * ### Full example
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.AggregateV2;
 * import com.pulumi.openstack.compute.AggregateV2Args;
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
 *         var dellServers = new AggregateV2(&#34;dellServers&#34;, AggregateV2Args.builder()        
 *             .hosts(            
 *                 &#34;myhost01.example.com&#34;,
 *                 &#34;myhost02.example.com&#34;)
 *             .metadata(Map.of(&#34;cpus&#34;, &#34;56&#34;))
 *             .region(&#34;RegionOne&#34;)
 *             .zone(&#34;nova&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Minimum required example
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.AggregateV2;
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
 *         var test = new AggregateV2(&#34;test&#34;);
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * You can import an existing Host Aggregate by their ID.
 * 
 * ```sh
 * $ pulumi import openstack:compute/aggregateV2:AggregateV2 myaggregate 24
 * ```
 * 
 * The ID can be obtained with an openstack command:
 * 
 * $ openstack aggregate list
 * 
 * +----+------+-------------------+
 * 
 * | ID | Name | Availability Zone |
 * 
 * +----+------+-------------------+
 * 
 * | 59 | test | None              |
 * 
 * +----+------+-------------------+
 * 
 */
@ResourceType(type="openstack:compute/aggregateV2:AggregateV2")
public class AggregateV2 extends com.pulumi.resources.CustomResource {
    /**
     * The list of hosts contained in the Host Aggregate. The hosts must be added
     * to Openstack and visible in the web interface, or the provider will fail to add them to the host
     * aggregate.
     * 
     */
    @Export(name="hosts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> hosts;

    /**
     * @return The list of hosts contained in the Host Aggregate. The hosts must be added
     * to Openstack and visible in the web interface, or the provider will fail to add them to the host
     * aggregate.
     * 
     */
    public Output<Optional<List<String>>> hosts() {
        return Codegen.optional(this.hosts);
    }
    /**
     * The metadata of the Host Aggregate. Can be useful to indicate scheduler hints.
     * 
     */
    @Export(name="metadata", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> metadata;

    /**
     * @return The metadata of the Host Aggregate. Can be useful to indicate scheduler hints.
     * 
     */
    public Output<Optional<Map<String,String>>> metadata() {
        return Codegen.optional(this.metadata);
    }
    /**
     * The name of the Host Aggregate
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Host Aggregate
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The region in which to create the Host Aggregate. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new Host Aggregate.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the Host Aggregate. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new Host Aggregate.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The name of the Availability Zone to use. If ommited, it will take the default
     * availability zone.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> zone;

    /**
     * @return The name of the Availability Zone to use. If ommited, it will take the default
     * availability zone.
     * 
     */
    public Output<Optional<String>> zone() {
        return Codegen.optional(this.zone);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AggregateV2(String name) {
        this(name, AggregateV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AggregateV2(String name, @Nullable AggregateV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AggregateV2(String name, @Nullable AggregateV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/aggregateV2:AggregateV2", name, args == null ? AggregateV2Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AggregateV2(String name, Output<String> id, @Nullable AggregateV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/aggregateV2:AggregateV2", name, state, makeResourceOptions(options, id));
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
    public static AggregateV2 get(String name, Output<String> id, @Nullable AggregateV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AggregateV2(name, id, state, options);
    }
}
