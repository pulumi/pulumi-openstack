// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.vpnaas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.vpnaas.IpSecPolicyArgs;
import com.pulumi.openstack.vpnaas.inputs.IpSecPolicyState;
import com.pulumi.openstack.vpnaas.outputs.IpSecPolicyLifetime;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V2 Neutron IPSec policy resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.vpnaas.IpSecPolicy;
 * import com.pulumi.openstack.vpnaas.IpSecPolicyArgs;
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
 *         var policy1 = new IpSecPolicy("policy1", IpSecPolicyArgs.builder()
 *             .name("my_policy")
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
 * Policies can be imported using the `id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:vpnaas/ipSecPolicy:IpSecPolicy policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 * 
 */
@ResourceType(type="openstack:vpnaas/ipSecPolicy:IpSecPolicy")
public class IpSecPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     * 
     */
    @Export(name="authAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> authAlgorithm;

    /**
     * @return The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
     * Default is sha1. Changing this updates the algorithm of the existing policy.
     * 
     */
    public Output<String> authAlgorithm() {
        return this.authAlgorithm;
    }
    /**
     * The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The human-readable description for the policy.
     * Changing this updates the description of the existing policy.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     * 
     */
    @Export(name="encapsulationMode", refs={String.class}, tree="[0]")
    private Output<String> encapsulationMode;

    /**
     * @return The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
     * Changing this updates the existing policy.
     * 
     */
    public Output<String> encapsulationMode() {
        return this.encapsulationMode;
    }
    /**
     * The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     * 
     */
    @Export(name="encryptionAlgorithm", refs={String.class}, tree="[0]")
    private Output<String> encryptionAlgorithm;

    /**
     * @return The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
     * The default value is aes-128. Changing this updates the existing policy.
     * 
     */
    public Output<String> encryptionAlgorithm() {
        return this.encryptionAlgorithm;
    }
    /**
     * The lifetime of the security association. Consists of Unit and Value.
     * 
     */
    @Export(name="lifetimes", refs={List.class,IpSecPolicyLifetime.class}, tree="[0,1]")
    private Output<List<IpSecPolicyLifetime>> lifetimes;

    /**
     * @return The lifetime of the security association. Consists of Unit and Value.
     * 
     */
    public Output<List<IpSecPolicyLifetime>> lifetimes() {
        return this.lifetimes;
    }
    /**
     * The name of the policy. Changing this updates the name of
     * the existing policy.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the policy. Changing this updates the name of
     * the existing policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
     * is group5. Changing this updates the existing policy.
     * 
     */
    @Export(name="pfs", refs={String.class}, tree="[0]")
    private Output<String> pfs;

    /**
     * @return The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default
     * is group5. Changing this updates the existing policy.
     * 
     */
    public Output<String> pfs() {
        return this.pfs;
    }
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an IPSec policy. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * policy.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     * 
     */
    @Export(name="tenantId", refs={String.class}, tree="[0]")
    private Output<String> tenantId;

    /**
     * @return The owner of the policy. Required if admin wants to
     * create a policy for another project. Changing this creates a new policy.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }
    /**
     * The transform protocol. Valid values are esp, ah and ah-esp.
     * Changing this updates the existing policy. Default is ESP.
     * 
     */
    @Export(name="transformProtocol", refs={String.class}, tree="[0]")
    private Output<String> transformProtocol;

    /**
     * @return The transform protocol. Valid values are esp, ah and ah-esp.
     * Changing this updates the existing policy. Default is ESP.
     * 
     */
    public Output<String> transformProtocol() {
        return this.transformProtocol;
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
    public IpSecPolicy(String name) {
        this(name, IpSecPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpSecPolicy(String name, @Nullable IpSecPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpSecPolicy(String name, @Nullable IpSecPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/ipSecPolicy:IpSecPolicy", name, args == null ? IpSecPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpSecPolicy(String name, Output<String> id, @Nullable IpSecPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:vpnaas/ipSecPolicy:IpSecPolicy", name, state, makeResourceOptions(options, id));
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
    public static IpSecPolicy get(String name, Output<String> id, @Nullable IpSecPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpSecPolicy(name, id, state, options);
    }
}
