// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.sharedfilesystem.SecurityServiceArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.SecurityServiceState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Use this resource to configure a security service.
 * 
 * &gt; **Note:** All arguments including the security service password will be
 * stored in the raw state as plain-text. [Read more about sensitive data in
 * state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * A security service stores configuration information for clients for
 * authentication and authorization (AuthN/AuthZ). For example, a share server
 * will be the client for an existing service such as LDAP, Kerberos, or
 * Microsoft Active Directory.
 * 
 * Minimum supported Manila microversion is 2.7.
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
 * import com.pulumi.openstack.sharedfilesystem.SecurityService;
 * import com.pulumi.openstack.sharedfilesystem.SecurityServiceArgs;
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
 *         var securityservice1 = new SecurityService("securityservice1", SecurityServiceArgs.builder()
 *             .name("security")
 *             .description("created by terraform")
 *             .type("active_directory")
 *             .server("192.168.199.10")
 *             .dnsIp("192.168.199.10")
 *             .domain("example.com")
 *             .ou("CN=Computers,DC=example,DC=com")
 *             .user("joinDomainUser")
 *             .password("s8cret")
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
 * This resource can be imported by specifying the ID of the security service:
 * 
 * ```sh
 * $ pulumi import openstack:sharedfilesystem/securityService:SecurityService securityservice_1 id
 * ```
 * 
 */
@ResourceType(type="openstack:sharedfilesystem/securityService:SecurityService")
public class SecurityService extends com.pulumi.resources.CustomResource {
    /**
     * The human-readable description for the security service.
     * Changing this updates the description of the existing security service.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The human-readable description for the security service.
     * Changing this updates the description of the existing security service.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The security service DNS IP address that is used inside the
     * tenant network.
     * 
     */
    @Export(name="dnsIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dnsIp;

    /**
     * @return The security service DNS IP address that is used inside the
     * tenant network.
     * 
     */
    public Output<Optional<String>> dnsIp() {
        return Codegen.optional(this.dnsIp);
    }
    /**
     * The security service domain.
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> domain;

    /**
     * @return The security service domain.
     * 
     */
    public Output<Optional<String>> domain() {
        return Codegen.optional(this.domain);
    }
    /**
     * The name of the security service. Changing this updates the name
     * of the existing security service.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the security service. Changing this updates the name
     * of the existing security service.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The security service ou. An organizational unit can be added to
     * specify where the share ends up. New in Manila microversion 2.44.
     * 
     */
    @Export(name="ou", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ou;

    /**
     * @return The security service ou. An organizational unit can be added to
     * specify where the share ends up. New in Manila microversion 2.44.
     * 
     */
    public Output<Optional<String>> ou() {
        return Codegen.optional(this.ou);
    }
    /**
     * The user password, if you specify a user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The user password, if you specify a user.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The owner of the Security Service.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The owner of the Security Service.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a security service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security service.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a security service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security service.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The security service host name or IP address.
     * 
     */
    @Export(name="server", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> server;

    /**
     * @return The security service host name or IP address.
     * 
     */
    public Output<Optional<String>> server() {
        return Codegen.optional(this.server);
    }
    /**
     * The security service type - can either be active\_directory,
     * kerberos or ldap.  Changing this updates the existing security service.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The security service type - can either be active\_directory,
     * kerberos or ldap.  Changing this updates the existing security service.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The security service user or group name that is used by the
     * tenant.
     * 
     */
    @Export(name="user", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> user;

    /**
     * @return The security service user or group name that is used by the
     * tenant.
     * 
     */
    public Output<Optional<String>> user() {
        return Codegen.optional(this.user);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityService(java.lang.String name) {
        this(name, SecurityServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityService(java.lang.String name, SecurityServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityService(java.lang.String name, SecurityServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:sharedfilesystem/securityService:SecurityService", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecurityService(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:sharedfilesystem/securityService:SecurityService", name, state, makeResourceOptions(options, id), false);
    }

    private static SecurityServiceArgs makeArgs(SecurityServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecurityServiceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static SecurityService get(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityService(name, id, state, options);
    }
}
