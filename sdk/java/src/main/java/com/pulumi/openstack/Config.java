// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack;

import com.pulumi.core.TypeShape;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("openstack");
/**
 * If set to `false`, OpenStack authorization won&#39;t be perfomed automatically, if the initial auth token get expired.
 * Defaults to `true`
 * 
 */
    public Optional<Boolean> allowReauth() {
        return Codegen.booleanProp("allowReauth").config(config).env("OS_ALLOW_REAUTH").get();
    }
/**
 * Application Credential ID to login with.
 * 
 */
    public Optional<String> applicationCredentialId() {
        return Codegen.stringProp("applicationCredentialId").config(config).get();
    }
/**
 * Application Credential name to login with.
 * 
 */
    public Optional<String> applicationCredentialName() {
        return Codegen.stringProp("applicationCredentialName").config(config).get();
    }
/**
 * Application Credential secret to login with.
 * 
 */
    public Optional<String> applicationCredentialSecret() {
        return Codegen.stringProp("applicationCredentialSecret").config(config).get();
    }
/**
 * The Identity authentication URL.
 * 
 */
    public Optional<String> authUrl() {
        return Codegen.stringProp("authUrl").config(config).get();
    }
/**
 * A Custom CA certificate.
 * 
 */
    public Optional<String> cacertFile() {
        return Codegen.stringProp("cacertFile").config(config).get();
    }
/**
 * A client certificate to authenticate with.
 * 
 */
    public Optional<String> cert() {
        return Codegen.stringProp("cert").config(config).get();
    }
/**
 * An entry in a `clouds.yaml` file to use.
 * 
 */
    public Optional<String> cloud() {
        return Codegen.stringProp("cloud").config(config).env("OS_CLOUD").get();
    }
/**
 * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
 * 
 */
    public Optional<String> defaultDomain() {
        return Codegen.stringProp("defaultDomain").config(config).get();
    }
/**
 * If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
 * to `true`.
 * 
 */
    public Optional<Boolean> delayedAuth() {
        return Codegen.booleanProp("delayedAuth").config(config).env("OS_DELAYED_AUTH").get();
    }
/**
 * If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
 * 
 */
    public Optional<Boolean> disableNoCacheHeader() {
        return Codegen.booleanProp("disableNoCacheHeader").config(config).get();
    }
/**
 * The ID of the Domain to scope to (Identity v3).
 * 
 */
    public Optional<String> domainId() {
        return Codegen.stringProp("domainId").config(config).get();
    }
/**
 * The name of the Domain to scope to (Identity v3).
 * 
 */
    public Optional<String> domainName() {
        return Codegen.stringProp("domainName").config(config).get();
    }
/**
 * Outputs very verbose logs with all calls made to and responses from OpenStack
 * 
 */
    public Optional<Boolean> enableLogging() {
        return Codegen.booleanProp("enableLogging").config(config).get();
    }
/**
 * A map of services with an endpoint to override what was from the Keystone catalog
 * 
 */
    public Optional<Map<String,Object>> endpointOverrides() {
        return Codegen.objectProp("endpointOverrides", TypeShape.<Map<String,Object>>builder(Map.class).addParameter(String.class).addParameter(Object.class).build()).config(config).get();
    }
    public Optional<String> endpointType() {
        return Codegen.stringProp("endpointType").config(config).env("OS_ENDPOINT_TYPE").get();
    }
/**
 * Trust self-signed certificates.
 * 
 */
    public Optional<Boolean> insecure() {
        return Codegen.booleanProp("insecure").config(config).env("OS_INSECURE").get();
    }
/**
 * A client private key to authenticate with.
 * 
 */
    public Optional<String> key() {
        return Codegen.stringProp("key").config(config).get();
    }
/**
 * How many times HTTP connection should be retried until giving up.
 * 
 */
    public Optional<Integer> maxRetries() {
        return Codegen.integerProp("maxRetries").config(config).get();
    }
/**
 * Password to login with.
 * 
 */
    public Optional<String> password() {
        return Codegen.stringProp("password").config(config).get();
    }
/**
 * The ID of the domain where the proejct resides (Identity v3).
 * 
 */
    public Optional<String> projectDomainId() {
        return Codegen.stringProp("projectDomainId").config(config).get();
    }
/**
 * The name of the domain where the project resides (Identity v3).
 * 
 */
    public Optional<String> projectDomainName() {
        return Codegen.stringProp("projectDomainName").config(config).get();
    }
/**
 * The OpenStack region to connect to.
 * 
 */
    public Optional<String> region() {
        return Codegen.stringProp("region").config(config).env("OS_REGION_NAME").get();
    }
/**
 * Use Swift&#39;s authentication system instead of Keystone. Only used for interaction with Swift.
 * 
 */
    public Optional<Boolean> swauth() {
        return Codegen.booleanProp("swauth").config(config).env("OS_SWAUTH").get();
    }
/**
 * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
 * 
 */
    public Optional<String> tenantId() {
        return Codegen.stringProp("tenantId").config(config).get();
    }
/**
 * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
 * 
 */
    public Optional<String> tenantName() {
        return Codegen.stringProp("tenantName").config(config).get();
    }
/**
 * Authentication token to use as an alternative to username/password.
 * 
 */
    public Optional<String> token() {
        return Codegen.stringProp("token").config(config).get();
    }
/**
 * If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
 * 
 */
    public Optional<Boolean> useOctavia() {
        return Codegen.booleanProp("useOctavia").config(config).env("OS_USE_OCTAVIA").get();
    }
/**
 * The ID of the domain where the user resides (Identity v3).
 * 
 */
    public Optional<String> userDomainId() {
        return Codegen.stringProp("userDomainId").config(config).get();
    }
/**
 * The name of the domain where the user resides (Identity v3).
 * 
 */
    public Optional<String> userDomainName() {
        return Codegen.stringProp("userDomainName").config(config).get();
    }
/**
 * Username to login with.
 * 
 */
    public Optional<String> userId() {
        return Codegen.stringProp("userId").config(config).get();
    }
/**
 * Username to login with.
 * 
 */
    public Optional<String> userName() {
        return Codegen.stringProp("userName").config(config).get();
    }
}
