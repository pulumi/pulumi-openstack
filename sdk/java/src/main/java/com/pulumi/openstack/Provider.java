// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.ProviderArgs;
import com.pulumi.openstack.Utilities;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the openstack package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:openstack")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * Application Credential ID to login with.
     * 
     */
    @Export(name="applicationCredentialId", type=String.class, parameters={})
    private Output</* @Nullable */ String> applicationCredentialId;

    /**
     * @return Application Credential ID to login with.
     * 
     */
    public Output<Optional<String>> applicationCredentialId() {
        return Codegen.optional(this.applicationCredentialId);
    }
    /**
     * Application Credential name to login with.
     * 
     */
    @Export(name="applicationCredentialName", type=String.class, parameters={})
    private Output</* @Nullable */ String> applicationCredentialName;

    /**
     * @return Application Credential name to login with.
     * 
     */
    public Output<Optional<String>> applicationCredentialName() {
        return Codegen.optional(this.applicationCredentialName);
    }
    /**
     * Application Credential secret to login with.
     * 
     */
    @Export(name="applicationCredentialSecret", type=String.class, parameters={})
    private Output</* @Nullable */ String> applicationCredentialSecret;

    /**
     * @return Application Credential secret to login with.
     * 
     */
    public Output<Optional<String>> applicationCredentialSecret() {
        return Codegen.optional(this.applicationCredentialSecret);
    }
    /**
     * The Identity authentication URL.
     * 
     */
    @Export(name="authUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> authUrl;

    /**
     * @return The Identity authentication URL.
     * 
     */
    public Output<Optional<String>> authUrl() {
        return Codegen.optional(this.authUrl);
    }
    /**
     * A Custom CA certificate.
     * 
     */
    @Export(name="cacertFile", type=String.class, parameters={})
    private Output</* @Nullable */ String> cacertFile;

    /**
     * @return A Custom CA certificate.
     * 
     */
    public Output<Optional<String>> cacertFile() {
        return Codegen.optional(this.cacertFile);
    }
    /**
     * A client certificate to authenticate with.
     * 
     */
    @Export(name="cert", type=String.class, parameters={})
    private Output</* @Nullable */ String> cert;

    /**
     * @return A client certificate to authenticate with.
     * 
     */
    public Output<Optional<String>> cert() {
        return Codegen.optional(this.cert);
    }
    /**
     * An entry in a `clouds.yaml` file to use.
     * 
     */
    @Export(name="cloud", type=String.class, parameters={})
    private Output</* @Nullable */ String> cloud;

    /**
     * @return An entry in a `clouds.yaml` file to use.
     * 
     */
    public Output<Optional<String>> cloud() {
        return Codegen.optional(this.cloud);
    }
    /**
     * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
     * 
     */
    @Export(name="defaultDomain", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultDomain;

    /**
     * @return The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
     * 
     */
    public Output<Optional<String>> defaultDomain() {
        return Codegen.optional(this.defaultDomain);
    }
    /**
     * The ID of the Domain to scope to (Identity v3).
     * 
     */
    @Export(name="domainId", type=String.class, parameters={})
    private Output</* @Nullable */ String> domainId;

    /**
     * @return The ID of the Domain to scope to (Identity v3).
     * 
     */
    public Output<Optional<String>> domainId() {
        return Codegen.optional(this.domainId);
    }
    /**
     * The name of the Domain to scope to (Identity v3).
     * 
     */
    @Export(name="domainName", type=String.class, parameters={})
    private Output</* @Nullable */ String> domainName;

    /**
     * @return The name of the Domain to scope to (Identity v3).
     * 
     */
    public Output<Optional<String>> domainName() {
        return Codegen.optional(this.domainName);
    }
    @Export(name="endpointType", type=String.class, parameters={})
    private Output</* @Nullable */ String> endpointType;

    public Output<Optional<String>> endpointType() {
        return Codegen.optional(this.endpointType);
    }
    /**
     * A client private key to authenticate with.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output</* @Nullable */ String> key;

    /**
     * @return A client private key to authenticate with.
     * 
     */
    public Output<Optional<String>> key() {
        return Codegen.optional(this.key);
    }
    /**
     * Password to login with.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return Password to login with.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * The ID of the domain where the proejct resides (Identity v3).
     * 
     */
    @Export(name="projectDomainId", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectDomainId;

    /**
     * @return The ID of the domain where the proejct resides (Identity v3).
     * 
     */
    public Output<Optional<String>> projectDomainId() {
        return Codegen.optional(this.projectDomainId);
    }
    /**
     * The name of the domain where the project resides (Identity v3).
     * 
     */
    @Export(name="projectDomainName", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectDomainName;

    /**
     * @return The name of the domain where the project resides (Identity v3).
     * 
     */
    public Output<Optional<String>> projectDomainName() {
        return Codegen.optional(this.projectDomainName);
    }
    /**
     * The OpenStack region to connect to.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output</* @Nullable */ String> region;

    /**
     * @return The OpenStack region to connect to.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    @Export(name="tenantId", type=String.class, parameters={})
    private Output</* @Nullable */ String> tenantId;

    /**
     * @return The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    public Output<Optional<String>> tenantId() {
        return Codegen.optional(this.tenantId);
    }
    /**
     * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    @Export(name="tenantName", type=String.class, parameters={})
    private Output</* @Nullable */ String> tenantName;

    /**
     * @return The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    public Output<Optional<String>> tenantName() {
        return Codegen.optional(this.tenantName);
    }
    /**
     * Authentication token to use as an alternative to username/password.
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output</* @Nullable */ String> token;

    /**
     * @return Authentication token to use as an alternative to username/password.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
    }
    /**
     * The ID of the domain where the user resides (Identity v3).
     * 
     */
    @Export(name="userDomainId", type=String.class, parameters={})
    private Output</* @Nullable */ String> userDomainId;

    /**
     * @return The ID of the domain where the user resides (Identity v3).
     * 
     */
    public Output<Optional<String>> userDomainId() {
        return Codegen.optional(this.userDomainId);
    }
    /**
     * The name of the domain where the user resides (Identity v3).
     * 
     */
    @Export(name="userDomainName", type=String.class, parameters={})
    private Output</* @Nullable */ String> userDomainName;

    /**
     * @return The name of the domain where the user resides (Identity v3).
     * 
     */
    public Output<Optional<String>> userDomainName() {
        return Codegen.optional(this.userDomainName);
    }
    /**
     * Username to login with.
     * 
     */
    @Export(name="userId", type=String.class, parameters={})
    private Output</* @Nullable */ String> userId;

    /**
     * @return Username to login with.
     * 
     */
    public Output<Optional<String>> userId() {
        return Codegen.optional(this.userId);
    }
    /**
     * Username to login with.
     * 
     */
    @Export(name="userName", type=String.class, parameters={})
    private Output</* @Nullable */ String> userName;

    /**
     * @return Username to login with.
     * 
     */
    public Output<Optional<String>> userName() {
        return Codegen.optional(this.userName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
