// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.identity.Ec2CredentialV3Args;
import com.pulumi.openstack.identity.inputs.Ec2CredentialV3State;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### EC2 credential in current project scope
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.Ec2CredentialV3;
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
 *         var ec2Key1 = new Ec2CredentialV3(&#34;ec2Key1&#34;);
 * 
 *     }
 * }
 * ```
 * ### EC2 credential in pre-defined project scope
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.identity.Ec2CredentialV3;
 * import com.pulumi.openstack.identity.Ec2CredentialV3Args;
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
 *         var ec2Key1 = new Ec2CredentialV3(&#34;ec2Key1&#34;, Ec2CredentialV3Args.builder()        
 *             .projectId(&#34;f7ac731cc11f40efbc03a9f9e1d1d21f&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * EC2 Credentials can be imported using the `access`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:identity/ec2CredentialV3:Ec2CredentialV3 ec2_cred_1 2d0ac4a2f81b4b0f9513ee49e780647d
 * ```
 * 
 */
@ResourceType(type="openstack:identity/ec2CredentialV3:Ec2CredentialV3")
public class Ec2CredentialV3 extends com.pulumi.resources.CustomResource {
    /**
     * contains an EC2 credential access UUID
     * 
     */
    @Export(name="access", type=String.class, parameters={})
    private Output<String> access;

    /**
     * @return contains an EC2 credential access UUID
     * 
     */
    public Output<String> access() {
        return this.access;
    }
    /**
     * The ID of the project the EC2 credential is created
     * for and that authentication requests using this EC2 credential will
     * be scoped to.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The ID of the project the EC2 credential is created
     * for and that authentication requests using this EC2 credential will
     * be scoped to.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new EC2 credential.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new EC2 credential.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * contains an EC2 credential secret UUID
     * 
     */
    @Export(name="secret", type=String.class, parameters={})
    private Output<String> secret;

    /**
     * @return contains an EC2 credential secret UUID
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }
    /**
     * contains an EC2 credential trust ID scope
     * 
     */
    @Export(name="trustId", type=String.class, parameters={})
    private Output<String> trustId;

    /**
     * @return contains an EC2 credential trust ID scope
     * 
     */
    public Output<String> trustId() {
        return this.trustId;
    }
    /**
     * The ID of the user the EC2 credential is created for.
     * 
     */
    @Export(name="userId", type=String.class, parameters={})
    private Output<String> userId;

    /**
     * @return The ID of the user the EC2 credential is created for.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ec2CredentialV3(String name) {
        this(name, Ec2CredentialV3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ec2CredentialV3(String name, @Nullable Ec2CredentialV3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ec2CredentialV3(String name, @Nullable Ec2CredentialV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/ec2CredentialV3:Ec2CredentialV3", name, args == null ? Ec2CredentialV3Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ec2CredentialV3(String name, Output<String> id, @Nullable Ec2CredentialV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:identity/ec2CredentialV3:Ec2CredentialV3", name, state, makeResourceOptions(options, id));
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
    public static Ec2CredentialV3 get(String name, Output<String> id, @Nullable Ec2CredentialV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ec2CredentialV3(name, id, state, options);
    }
}
