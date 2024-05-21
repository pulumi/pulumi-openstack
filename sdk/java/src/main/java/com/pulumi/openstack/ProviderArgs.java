// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * If set to `false`, OpenStack authorization won&#39;t be perfomed automatically, if the initial auth token get expired.
     * Defaults to `true`
     * 
     */
    @Import(name="allowReauth", json=true)
    private @Nullable Output<Boolean> allowReauth;

    /**
     * @return If set to `false`, OpenStack authorization won&#39;t be perfomed automatically, if the initial auth token get expired.
     * Defaults to `true`
     * 
     */
    public Optional<Output<Boolean>> allowReauth() {
        return Optional.ofNullable(this.allowReauth);
    }

    /**
     * Application Credential ID to login with.
     * 
     */
    @Import(name="applicationCredentialId")
    private @Nullable Output<String> applicationCredentialId;

    /**
     * @return Application Credential ID to login with.
     * 
     */
    public Optional<Output<String>> applicationCredentialId() {
        return Optional.ofNullable(this.applicationCredentialId);
    }

    /**
     * Application Credential name to login with.
     * 
     */
    @Import(name="applicationCredentialName")
    private @Nullable Output<String> applicationCredentialName;

    /**
     * @return Application Credential name to login with.
     * 
     */
    public Optional<Output<String>> applicationCredentialName() {
        return Optional.ofNullable(this.applicationCredentialName);
    }

    /**
     * Application Credential secret to login with.
     * 
     */
    @Import(name="applicationCredentialSecret")
    private @Nullable Output<String> applicationCredentialSecret;

    /**
     * @return Application Credential secret to login with.
     * 
     */
    public Optional<Output<String>> applicationCredentialSecret() {
        return Optional.ofNullable(this.applicationCredentialSecret);
    }

    /**
     * The Identity authentication URL.
     * 
     */
    @Import(name="authUrl")
    private @Nullable Output<String> authUrl;

    /**
     * @return The Identity authentication URL.
     * 
     */
    public Optional<Output<String>> authUrl() {
        return Optional.ofNullable(this.authUrl);
    }

    /**
     * A Custom CA certificate.
     * 
     */
    @Import(name="cacertFile")
    private @Nullable Output<String> cacertFile;

    /**
     * @return A Custom CA certificate.
     * 
     */
    public Optional<Output<String>> cacertFile() {
        return Optional.ofNullable(this.cacertFile);
    }

    /**
     * A client certificate to authenticate with.
     * 
     */
    @Import(name="cert")
    private @Nullable Output<String> cert;

    /**
     * @return A client certificate to authenticate with.
     * 
     */
    public Optional<Output<String>> cert() {
        return Optional.ofNullable(this.cert);
    }

    /**
     * An entry in a `clouds.yaml` file to use.
     * 
     */
    @Import(name="cloud")
    private @Nullable Output<String> cloud;

    /**
     * @return An entry in a `clouds.yaml` file to use.
     * 
     */
    public Optional<Output<String>> cloud() {
        return Optional.ofNullable(this.cloud);
    }

    /**
     * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
     * 
     */
    @Import(name="defaultDomain")
    private @Nullable Output<String> defaultDomain;

    /**
     * @return The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
     * 
     */
    public Optional<Output<String>> defaultDomain() {
        return Optional.ofNullable(this.defaultDomain);
    }

    /**
     * If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
     * to `true`.
     * 
     */
    @Import(name="delayedAuth", json=true)
    private @Nullable Output<Boolean> delayedAuth;

    /**
     * @return If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
     * to `true`.
     * 
     */
    public Optional<Output<Boolean>> delayedAuth() {
        return Optional.ofNullable(this.delayedAuth);
    }

    /**
     * If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
     * 
     */
    @Import(name="disableNoCacheHeader", json=true)
    private @Nullable Output<Boolean> disableNoCacheHeader;

    /**
     * @return If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
     * 
     */
    public Optional<Output<Boolean>> disableNoCacheHeader() {
        return Optional.ofNullable(this.disableNoCacheHeader);
    }

    /**
     * The ID of the Domain to scope to (Identity v3).
     * 
     */
    @Import(name="domainId")
    private @Nullable Output<String> domainId;

    /**
     * @return The ID of the Domain to scope to (Identity v3).
     * 
     */
    public Optional<Output<String>> domainId() {
        return Optional.ofNullable(this.domainId);
    }

    /**
     * The name of the Domain to scope to (Identity v3).
     * 
     */
    @Import(name="domainName")
    private @Nullable Output<String> domainName;

    /**
     * @return The name of the Domain to scope to (Identity v3).
     * 
     */
    public Optional<Output<String>> domainName() {
        return Optional.ofNullable(this.domainName);
    }

    /**
     * Outputs very verbose logs with all calls made to and responses from OpenStack
     * 
     */
    @Import(name="enableLogging", json=true)
    private @Nullable Output<Boolean> enableLogging;

    /**
     * @return Outputs very verbose logs with all calls made to and responses from OpenStack
     * 
     */
    public Optional<Output<Boolean>> enableLogging() {
        return Optional.ofNullable(this.enableLogging);
    }

    /**
     * A map of services with an endpoint to override what was from the Keystone catalog
     * 
     */
    @Import(name="endpointOverrides", json=true)
    private @Nullable Output<Map<String,Object>> endpointOverrides;

    /**
     * @return A map of services with an endpoint to override what was from the Keystone catalog
     * 
     */
    public Optional<Output<Map<String,Object>>> endpointOverrides() {
        return Optional.ofNullable(this.endpointOverrides);
    }

    @Import(name="endpointType")
    private @Nullable Output<String> endpointType;

    public Optional<Output<String>> endpointType() {
        return Optional.ofNullable(this.endpointType);
    }

    /**
     * Trust self-signed certificates.
     * 
     */
    @Import(name="insecure", json=true)
    private @Nullable Output<Boolean> insecure;

    /**
     * @return Trust self-signed certificates.
     * 
     */
    public Optional<Output<Boolean>> insecure() {
        return Optional.ofNullable(this.insecure);
    }

    /**
     * A client private key to authenticate with.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return A client private key to authenticate with.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * How many times HTTP connection should be retried until giving up.
     * 
     */
    @Import(name="maxRetries", json=true)
    private @Nullable Output<Integer> maxRetries;

    /**
     * @return How many times HTTP connection should be retried until giving up.
     * 
     */
    public Optional<Output<Integer>> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }

    /**
     * Password to login with.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password to login with.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The ID of the domain where the proejct resides (Identity v3).
     * 
     */
    @Import(name="projectDomainId")
    private @Nullable Output<String> projectDomainId;

    /**
     * @return The ID of the domain where the proejct resides (Identity v3).
     * 
     */
    public Optional<Output<String>> projectDomainId() {
        return Optional.ofNullable(this.projectDomainId);
    }

    /**
     * The name of the domain where the project resides (Identity v3).
     * 
     */
    @Import(name="projectDomainName")
    private @Nullable Output<String> projectDomainName;

    /**
     * @return The name of the domain where the project resides (Identity v3).
     * 
     */
    public Optional<Output<String>> projectDomainName() {
        return Optional.ofNullable(this.projectDomainName);
    }

    /**
     * The OpenStack region to connect to.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The OpenStack region to connect to.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Use Swift&#39;s authentication system instead of Keystone. Only used for interaction with Swift.
     * 
     */
    @Import(name="swauth", json=true)
    private @Nullable Output<Boolean> swauth;

    /**
     * @return Use Swift&#39;s authentication system instead of Keystone. Only used for interaction with Swift.
     * 
     */
    public Optional<Output<Boolean>> swauth() {
        return Optional.ofNullable(this.swauth);
    }

    /**
     * If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).
     * 
     */
    @Import(name="systemScope", json=true)
    private @Nullable Output<Boolean> systemScope;

    /**
     * @return If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).
     * 
     */
    public Optional<Output<Boolean>> systemScope() {
        return Optional.ofNullable(this.systemScope);
    }

    /**
     * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    /**
     * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    @Import(name="tenantName")
    private @Nullable Output<String> tenantName;

    /**
     * @return The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
     * 
     */
    public Optional<Output<String>> tenantName() {
        return Optional.ofNullable(this.tenantName);
    }

    /**
     * Authentication token to use as an alternative to username/password.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Authentication token to use as an alternative to username/password.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * The ID of the domain where the user resides (Identity v3).
     * 
     */
    @Import(name="userDomainId")
    private @Nullable Output<String> userDomainId;

    /**
     * @return The ID of the domain where the user resides (Identity v3).
     * 
     */
    public Optional<Output<String>> userDomainId() {
        return Optional.ofNullable(this.userDomainId);
    }

    /**
     * The name of the domain where the user resides (Identity v3).
     * 
     */
    @Import(name="userDomainName")
    private @Nullable Output<String> userDomainName;

    /**
     * @return The name of the domain where the user resides (Identity v3).
     * 
     */
    public Optional<Output<String>> userDomainName() {
        return Optional.ofNullable(this.userDomainName);
    }

    /**
     * User ID to login with.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<String> userId;

    /**
     * @return User ID to login with.
     * 
     */
    public Optional<Output<String>> userId() {
        return Optional.ofNullable(this.userId);
    }

    /**
     * Username to login with.
     * 
     */
    @Import(name="userName")
    private @Nullable Output<String> userName;

    /**
     * @return Username to login with.
     * 
     */
    public Optional<Output<String>> userName() {
        return Optional.ofNullable(this.userName);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.allowReauth = $.allowReauth;
        this.applicationCredentialId = $.applicationCredentialId;
        this.applicationCredentialName = $.applicationCredentialName;
        this.applicationCredentialSecret = $.applicationCredentialSecret;
        this.authUrl = $.authUrl;
        this.cacertFile = $.cacertFile;
        this.cert = $.cert;
        this.cloud = $.cloud;
        this.defaultDomain = $.defaultDomain;
        this.delayedAuth = $.delayedAuth;
        this.disableNoCacheHeader = $.disableNoCacheHeader;
        this.domainId = $.domainId;
        this.domainName = $.domainName;
        this.enableLogging = $.enableLogging;
        this.endpointOverrides = $.endpointOverrides;
        this.endpointType = $.endpointType;
        this.insecure = $.insecure;
        this.key = $.key;
        this.maxRetries = $.maxRetries;
        this.password = $.password;
        this.projectDomainId = $.projectDomainId;
        this.projectDomainName = $.projectDomainName;
        this.region = $.region;
        this.swauth = $.swauth;
        this.systemScope = $.systemScope;
        this.tenantId = $.tenantId;
        this.tenantName = $.tenantName;
        this.token = $.token;
        this.userDomainId = $.userDomainId;
        this.userDomainName = $.userDomainName;
        this.userId = $.userId;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowReauth If set to `false`, OpenStack authorization won&#39;t be perfomed automatically, if the initial auth token get expired.
         * Defaults to `true`
         * 
         * @return builder
         * 
         */
        public Builder allowReauth(@Nullable Output<Boolean> allowReauth) {
            $.allowReauth = allowReauth;
            return this;
        }

        /**
         * @param allowReauth If set to `false`, OpenStack authorization won&#39;t be perfomed automatically, if the initial auth token get expired.
         * Defaults to `true`
         * 
         * @return builder
         * 
         */
        public Builder allowReauth(Boolean allowReauth) {
            return allowReauth(Output.of(allowReauth));
        }

        /**
         * @param applicationCredentialId Application Credential ID to login with.
         * 
         * @return builder
         * 
         */
        public Builder applicationCredentialId(@Nullable Output<String> applicationCredentialId) {
            $.applicationCredentialId = applicationCredentialId;
            return this;
        }

        /**
         * @param applicationCredentialId Application Credential ID to login with.
         * 
         * @return builder
         * 
         */
        public Builder applicationCredentialId(String applicationCredentialId) {
            return applicationCredentialId(Output.of(applicationCredentialId));
        }

        /**
         * @param applicationCredentialName Application Credential name to login with.
         * 
         * @return builder
         * 
         */
        public Builder applicationCredentialName(@Nullable Output<String> applicationCredentialName) {
            $.applicationCredentialName = applicationCredentialName;
            return this;
        }

        /**
         * @param applicationCredentialName Application Credential name to login with.
         * 
         * @return builder
         * 
         */
        public Builder applicationCredentialName(String applicationCredentialName) {
            return applicationCredentialName(Output.of(applicationCredentialName));
        }

        /**
         * @param applicationCredentialSecret Application Credential secret to login with.
         * 
         * @return builder
         * 
         */
        public Builder applicationCredentialSecret(@Nullable Output<String> applicationCredentialSecret) {
            $.applicationCredentialSecret = applicationCredentialSecret;
            return this;
        }

        /**
         * @param applicationCredentialSecret Application Credential secret to login with.
         * 
         * @return builder
         * 
         */
        public Builder applicationCredentialSecret(String applicationCredentialSecret) {
            return applicationCredentialSecret(Output.of(applicationCredentialSecret));
        }

        /**
         * @param authUrl The Identity authentication URL.
         * 
         * @return builder
         * 
         */
        public Builder authUrl(@Nullable Output<String> authUrl) {
            $.authUrl = authUrl;
            return this;
        }

        /**
         * @param authUrl The Identity authentication URL.
         * 
         * @return builder
         * 
         */
        public Builder authUrl(String authUrl) {
            return authUrl(Output.of(authUrl));
        }

        /**
         * @param cacertFile A Custom CA certificate.
         * 
         * @return builder
         * 
         */
        public Builder cacertFile(@Nullable Output<String> cacertFile) {
            $.cacertFile = cacertFile;
            return this;
        }

        /**
         * @param cacertFile A Custom CA certificate.
         * 
         * @return builder
         * 
         */
        public Builder cacertFile(String cacertFile) {
            return cacertFile(Output.of(cacertFile));
        }

        /**
         * @param cert A client certificate to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder cert(@Nullable Output<String> cert) {
            $.cert = cert;
            return this;
        }

        /**
         * @param cert A client certificate to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder cert(String cert) {
            return cert(Output.of(cert));
        }

        /**
         * @param cloud An entry in a `clouds.yaml` file to use.
         * 
         * @return builder
         * 
         */
        public Builder cloud(@Nullable Output<String> cloud) {
            $.cloud = cloud;
            return this;
        }

        /**
         * @param cloud An entry in a `clouds.yaml` file to use.
         * 
         * @return builder
         * 
         */
        public Builder cloud(String cloud) {
            return cloud(Output.of(cloud));
        }

        /**
         * @param defaultDomain The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder defaultDomain(@Nullable Output<String> defaultDomain) {
            $.defaultDomain = defaultDomain;
            return this;
        }

        /**
         * @param defaultDomain The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder defaultDomain(String defaultDomain) {
            return defaultDomain(Output.of(defaultDomain));
        }

        /**
         * @param delayedAuth If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
         * to `true`.
         * 
         * @return builder
         * 
         */
        public Builder delayedAuth(@Nullable Output<Boolean> delayedAuth) {
            $.delayedAuth = delayedAuth;
            return this;
        }

        /**
         * @param delayedAuth If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
         * to `true`.
         * 
         * @return builder
         * 
         */
        public Builder delayedAuth(Boolean delayedAuth) {
            return delayedAuth(Output.of(delayedAuth));
        }

        /**
         * @param disableNoCacheHeader If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
         * 
         * @return builder
         * 
         */
        public Builder disableNoCacheHeader(@Nullable Output<Boolean> disableNoCacheHeader) {
            $.disableNoCacheHeader = disableNoCacheHeader;
            return this;
        }

        /**
         * @param disableNoCacheHeader If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
         * 
         * @return builder
         * 
         */
        public Builder disableNoCacheHeader(Boolean disableNoCacheHeader) {
            return disableNoCacheHeader(Output.of(disableNoCacheHeader));
        }

        /**
         * @param domainId The ID of the Domain to scope to (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder domainId(@Nullable Output<String> domainId) {
            $.domainId = domainId;
            return this;
        }

        /**
         * @param domainId The ID of the Domain to scope to (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder domainId(String domainId) {
            return domainId(Output.of(domainId));
        }

        /**
         * @param domainName The name of the Domain to scope to (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder domainName(@Nullable Output<String> domainName) {
            $.domainName = domainName;
            return this;
        }

        /**
         * @param domainName The name of the Domain to scope to (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder domainName(String domainName) {
            return domainName(Output.of(domainName));
        }

        /**
         * @param enableLogging Outputs very verbose logs with all calls made to and responses from OpenStack
         * 
         * @return builder
         * 
         */
        public Builder enableLogging(@Nullable Output<Boolean> enableLogging) {
            $.enableLogging = enableLogging;
            return this;
        }

        /**
         * @param enableLogging Outputs very verbose logs with all calls made to and responses from OpenStack
         * 
         * @return builder
         * 
         */
        public Builder enableLogging(Boolean enableLogging) {
            return enableLogging(Output.of(enableLogging));
        }

        /**
         * @param endpointOverrides A map of services with an endpoint to override what was from the Keystone catalog
         * 
         * @return builder
         * 
         */
        public Builder endpointOverrides(@Nullable Output<Map<String,Object>> endpointOverrides) {
            $.endpointOverrides = endpointOverrides;
            return this;
        }

        /**
         * @param endpointOverrides A map of services with an endpoint to override what was from the Keystone catalog
         * 
         * @return builder
         * 
         */
        public Builder endpointOverrides(Map<String,Object> endpointOverrides) {
            return endpointOverrides(Output.of(endpointOverrides));
        }

        public Builder endpointType(@Nullable Output<String> endpointType) {
            $.endpointType = endpointType;
            return this;
        }

        public Builder endpointType(String endpointType) {
            return endpointType(Output.of(endpointType));
        }

        /**
         * @param insecure Trust self-signed certificates.
         * 
         * @return builder
         * 
         */
        public Builder insecure(@Nullable Output<Boolean> insecure) {
            $.insecure = insecure;
            return this;
        }

        /**
         * @param insecure Trust self-signed certificates.
         * 
         * @return builder
         * 
         */
        public Builder insecure(Boolean insecure) {
            return insecure(Output.of(insecure));
        }

        /**
         * @param key A client private key to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key A client private key to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param maxRetries How many times HTTP connection should be retried until giving up.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(@Nullable Output<Integer> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        /**
         * @param maxRetries How many times HTTP connection should be retried until giving up.
         * 
         * @return builder
         * 
         */
        public Builder maxRetries(Integer maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        /**
         * @param password Password to login with.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password to login with.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param projectDomainId The ID of the domain where the proejct resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder projectDomainId(@Nullable Output<String> projectDomainId) {
            $.projectDomainId = projectDomainId;
            return this;
        }

        /**
         * @param projectDomainId The ID of the domain where the proejct resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder projectDomainId(String projectDomainId) {
            return projectDomainId(Output.of(projectDomainId));
        }

        /**
         * @param projectDomainName The name of the domain where the project resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder projectDomainName(@Nullable Output<String> projectDomainName) {
            $.projectDomainName = projectDomainName;
            return this;
        }

        /**
         * @param projectDomainName The name of the domain where the project resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder projectDomainName(String projectDomainName) {
            return projectDomainName(Output.of(projectDomainName));
        }

        /**
         * @param region The OpenStack region to connect to.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The OpenStack region to connect to.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param swauth Use Swift&#39;s authentication system instead of Keystone. Only used for interaction with Swift.
         * 
         * @return builder
         * 
         */
        public Builder swauth(@Nullable Output<Boolean> swauth) {
            $.swauth = swauth;
            return this;
        }

        /**
         * @param swauth Use Swift&#39;s authentication system instead of Keystone. Only used for interaction with Swift.
         * 
         * @return builder
         * 
         */
        public Builder swauth(Boolean swauth) {
            return swauth(Output.of(swauth));
        }

        /**
         * @param systemScope If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder systemScope(@Nullable Output<Boolean> systemScope) {
            $.systemScope = systemScope;
            return this;
        }

        /**
         * @param systemScope If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder systemScope(Boolean systemScope) {
            return systemScope(Output.of(systemScope));
        }

        /**
         * @param tenantId The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        /**
         * @param tenantName The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
         * 
         * @return builder
         * 
         */
        public Builder tenantName(@Nullable Output<String> tenantName) {
            $.tenantName = tenantName;
            return this;
        }

        /**
         * @param tenantName The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
         * 
         * @return builder
         * 
         */
        public Builder tenantName(String tenantName) {
            return tenantName(Output.of(tenantName));
        }

        /**
         * @param token Authentication token to use as an alternative to username/password.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Authentication token to use as an alternative to username/password.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param userDomainId The ID of the domain where the user resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder userDomainId(@Nullable Output<String> userDomainId) {
            $.userDomainId = userDomainId;
            return this;
        }

        /**
         * @param userDomainId The ID of the domain where the user resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder userDomainId(String userDomainId) {
            return userDomainId(Output.of(userDomainId));
        }

        /**
         * @param userDomainName The name of the domain where the user resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder userDomainName(@Nullable Output<String> userDomainName) {
            $.userDomainName = userDomainName;
            return this;
        }

        /**
         * @param userDomainName The name of the domain where the user resides (Identity v3).
         * 
         * @return builder
         * 
         */
        public Builder userDomainName(String userDomainName) {
            return userDomainName(Output.of(userDomainName));
        }

        /**
         * @param userId User ID to login with.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId User ID to login with.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        /**
         * @param userName Username to login with.
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName Username to login with.
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public ProviderArgs build() {
            $.allowReauth = Codegen.booleanProp("allowReauth").output().arg($.allowReauth).env("OS_ALLOW_REAUTH").getNullable();
            $.cloud = Codegen.stringProp("cloud").output().arg($.cloud).env("OS_CLOUD").getNullable();
            $.delayedAuth = Codegen.booleanProp("delayedAuth").output().arg($.delayedAuth).env("OS_DELAYED_AUTH").getNullable();
            $.endpointType = Codegen.stringProp("endpointType").output().arg($.endpointType).env("OS_ENDPOINT_TYPE").getNullable();
            $.insecure = Codegen.booleanProp("insecure").output().arg($.insecure).env("OS_INSECURE").getNullable();
            $.region = Codegen.stringProp("region").output().arg($.region).env("OS_REGION_NAME").getNullable();
            $.swauth = Codegen.booleanProp("swauth").output().arg($.swauth).env("OS_SWAUTH").getNullable();
            return $;
        }
    }

}
