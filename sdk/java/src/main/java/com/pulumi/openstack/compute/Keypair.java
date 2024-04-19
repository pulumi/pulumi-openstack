// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.compute.KeypairArgs;
import com.pulumi.openstack.compute.inputs.KeypairState;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * ### Import an Existing Public Key
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.Keypair;
 * import com.pulumi.openstack.compute.KeypairArgs;
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
 *         var test_keypair = new Keypair(&#34;test-keypair&#34;, KeypairArgs.builder()        
 *             .name(&#34;my-keypair&#34;)
 *             .publicKey(&#34;ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLotBCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAnOfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZqd9LvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TaIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIF61p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Generate a Public/Private Key Pair
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.compute.Keypair;
 * import com.pulumi.openstack.compute.KeypairArgs;
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
 *         var test_keypair = new Keypair(&#34;test-keypair&#34;, KeypairArgs.builder()        
 *             .name(&#34;my-keypair&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Keypairs can be imported using the `name`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:compute/keypair:Keypair my-keypair test-keypair
 * ```
 * 
 */
@ResourceType(type="openstack:compute/keypair:Keypair")
public class Keypair extends com.pulumi.resources.CustomResource {
    /**
     * The fingerprint of the public key.
     * 
     */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return The fingerprint of the public key.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * A unique name for the keypair. Changing this creates a new
     * keypair.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return A unique name for the keypair. Changing this creates a new
     * keypair.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The generated private key when no public key is specified.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return The generated private key when no public key is specified.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * A pregenerated OpenSSH-formatted public key.
     * Changing this creates a new keypair. If a public key is not specified, then
     * a public/private key pair will be automatically generated. If a pair is
     * created, then destroying this resource means you will lose access to that
     * keypair forever.
     * 
     */
    @Export(name="publicKey", refs={String.class}, tree="[0]")
    private Output<String> publicKey;

    /**
     * @return A pregenerated OpenSSH-formatted public key.
     * Changing this creates a new keypair. If a public key is not specified, then
     * a public/private key pair will be automatically generated. If a pair is
     * created, then destroying this resource means you will lose access to that
     * keypair forever.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }
    /**
     * The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new keypair.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * Keypairs are associated with accounts, but a Compute client is needed to
     * create one. If omitted, the `region` argument of the provider is used.
     * Changing this creates a new keypair.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * This allows administrative users to operate key-pairs
     * of specified user ID. For this feature your need to have openstack microversion
     * 2.10 (Liberty) or later.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return This allows administrative users to operate key-pairs
     * of specified user ID. For this feature your need to have openstack microversion
     * 2.10 (Liberty) or later.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }
    /**
     * Map of additional options.
     * 
     */
    @Export(name="valueSpecs", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
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
    public Keypair(String name) {
        this(name, KeypairArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Keypair(String name, @Nullable KeypairArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Keypair(String name, @Nullable KeypairArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/keypair:Keypair", name, args == null ? KeypairArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Keypair(String name, Output<String> id, @Nullable KeypairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:compute/keypair:Keypair", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "privateKey"
            ))
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
    public static Keypair get(String name, Output<String> id, @Nullable KeypairState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Keypair(name, id, state, options);
    }
}
