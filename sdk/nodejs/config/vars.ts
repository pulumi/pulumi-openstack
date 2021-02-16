// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("openstack");

/**
 * If set to `false`, OpenStack authorization won't be perfomed automatically, if the initial auth token get expired.
 * Defaults to `true`
 */
export let allowReauth: boolean | undefined = __config.getObject<boolean>("allowReauth") || <any>utilities.getEnvBoolean("OS_ALLOW_REAUTH");
/**
 * Application Credential ID to login with.
 */
export let applicationCredentialId: string | undefined = __config.get("applicationCredentialId");
/**
 * Application Credential name to login with.
 */
export let applicationCredentialName: string | undefined = __config.get("applicationCredentialName");
/**
 * Application Credential secret to login with.
 */
export let applicationCredentialSecret: string | undefined = __config.get("applicationCredentialSecret");
/**
 * The Identity authentication URL.
 */
export let authUrl: string | undefined = __config.get("authUrl");
/**
 * A Custom CA certificate.
 */
export let cacertFile: string | undefined = __config.get("cacertFile");
/**
 * A client certificate to authenticate with.
 */
export let cert: string | undefined = __config.get("cert");
/**
 * An entry in a `clouds.yaml` file to use.
 */
export let cloud: string | undefined = __config.get("cloud") || utilities.getEnv("OS_CLOUD");
/**
 * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
 */
export let defaultDomain: string | undefined = __config.get("defaultDomain");
/**
 * If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
 * to `true`.
 */
export let delayedAuth: boolean | undefined = __config.getObject<boolean>("delayedAuth") || <any>utilities.getEnvBoolean("OS_DELAYED_AUTH");
/**
 * If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
 */
export let disableNoCacheHeader: boolean | undefined = __config.getObject<boolean>("disableNoCacheHeader");
/**
 * The ID of the Domain to scope to (Identity v3).
 */
export let domainId: string | undefined = __config.get("domainId");
/**
 * The name of the Domain to scope to (Identity v3).
 */
export let domainName: string | undefined = __config.get("domainName");
/**
 * A map of services with an endpoint to override what was from the Keystone catalog
 */
export let endpointOverrides: {[key: string]: any} | undefined = __config.getObject<{[key: string]: any}>("endpointOverrides");
export let endpointType: string | undefined = __config.get("endpointType") || utilities.getEnv("OS_ENDPOINT_TYPE");
/**
 * Trust self-signed certificates.
 */
export let insecure: boolean | undefined = __config.getObject<boolean>("insecure") || <any>utilities.getEnvBoolean("OS_INSECURE");
/**
 * A client private key to authenticate with.
 */
export let key: string | undefined = __config.get("key");
/**
 * How many times HTTP connection should be retried until giving up.
 */
export let maxRetries: number | undefined = __config.getObject<number>("maxRetries");
/**
 * Password to login with.
 */
export let password: string | undefined = __config.get("password");
/**
 * The ID of the domain where the proejct resides (Identity v3).
 */
export let projectDomainId: string | undefined = __config.get("projectDomainId");
/**
 * The name of the domain where the project resides (Identity v3).
 */
export let projectDomainName: string | undefined = __config.get("projectDomainName");
/**
 * The OpenStack region to connect to.
 */
export let region: string | undefined = __config.get("region") || utilities.getEnv("OS_REGION_NAME");
/**
 * Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
 */
export let swauth: boolean | undefined = __config.getObject<boolean>("swauth") || <any>utilities.getEnvBoolean("OS_SWAUTH");
/**
 * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
export let tenantId: string | undefined = __config.get("tenantId");
/**
 * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
export let tenantName: string | undefined = __config.get("tenantName");
/**
 * Authentication token to use as an alternative to username/password.
 */
export let token: string | undefined = __config.get("token");
/**
 * If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
 */
export let useOctavia: boolean | undefined = __config.getObject<boolean>("useOctavia") || <any>utilities.getEnvBoolean("OS_USE_OCTAVIA");
/**
 * The ID of the domain where the user resides (Identity v3).
 */
export let userDomainId: string | undefined = __config.get("userDomainId");
/**
 * The name of the domain where the user resides (Identity v3).
 */
export let userDomainName: string | undefined = __config.get("userDomainName");
/**
 * Username to login with.
 */
export let userId: string | undefined = __config.get("userId");
/**
 * Username to login with.
 */
export let userName: string | undefined = __config.get("userName");
