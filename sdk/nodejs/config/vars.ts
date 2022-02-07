// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("openstack");

/**
 * If set to `false`, OpenStack authorization won't be perfomed automatically, if the initial auth token get expired.
 * Defaults to `true`
 */
export declare const allowReauth: boolean | undefined;
Object.defineProperty(exports, "allowReauth", {
    get() {
        return __config.getObject<boolean>("allowReauth") ?? utilities.getEnvBoolean("OS_ALLOW_REAUTH");
    },
    enumerable: true,
});

/**
 * Application Credential ID to login with.
 */
export declare const applicationCredentialId: string | undefined;
Object.defineProperty(exports, "applicationCredentialId", {
    get() {
        return __config.get("applicationCredentialId");
    },
    enumerable: true,
});

/**
 * Application Credential name to login with.
 */
export declare const applicationCredentialName: string | undefined;
Object.defineProperty(exports, "applicationCredentialName", {
    get() {
        return __config.get("applicationCredentialName");
    },
    enumerable: true,
});

/**
 * Application Credential secret to login with.
 */
export declare const applicationCredentialSecret: string | undefined;
Object.defineProperty(exports, "applicationCredentialSecret", {
    get() {
        return __config.get("applicationCredentialSecret");
    },
    enumerable: true,
});

/**
 * The Identity authentication URL.
 */
export declare const authUrl: string | undefined;
Object.defineProperty(exports, "authUrl", {
    get() {
        return __config.get("authUrl");
    },
    enumerable: true,
});

/**
 * A Custom CA certificate.
 */
export declare const cacertFile: string | undefined;
Object.defineProperty(exports, "cacertFile", {
    get() {
        return __config.get("cacertFile");
    },
    enumerable: true,
});

/**
 * A client certificate to authenticate with.
 */
export declare const cert: string | undefined;
Object.defineProperty(exports, "cert", {
    get() {
        return __config.get("cert");
    },
    enumerable: true,
});

/**
 * An entry in a `clouds.yaml` file to use.
 */
export declare const cloud: string | undefined;
Object.defineProperty(exports, "cloud", {
    get() {
        return __config.get("cloud") ?? utilities.getEnv("OS_CLOUD");
    },
    enumerable: true,
});

/**
 * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
 */
export declare const defaultDomain: string | undefined;
Object.defineProperty(exports, "defaultDomain", {
    get() {
        return __config.get("defaultDomain");
    },
    enumerable: true,
});

/**
 * If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
 * to `true`.
 */
export declare const delayedAuth: boolean | undefined;
Object.defineProperty(exports, "delayedAuth", {
    get() {
        return __config.getObject<boolean>("delayedAuth") ?? utilities.getEnvBoolean("OS_DELAYED_AUTH");
    },
    enumerable: true,
});

/**
 * If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
 */
export declare const disableNoCacheHeader: boolean | undefined;
Object.defineProperty(exports, "disableNoCacheHeader", {
    get() {
        return __config.getObject<boolean>("disableNoCacheHeader");
    },
    enumerable: true,
});

/**
 * The ID of the Domain to scope to (Identity v3).
 */
export declare const domainId: string | undefined;
Object.defineProperty(exports, "domainId", {
    get() {
        return __config.get("domainId");
    },
    enumerable: true,
});

/**
 * The name of the Domain to scope to (Identity v3).
 */
export declare const domainName: string | undefined;
Object.defineProperty(exports, "domainName", {
    get() {
        return __config.get("domainName");
    },
    enumerable: true,
});

/**
 * Outputs very verbose logs with all calls made to and responses from OpenStack
 */
export declare const enableLogging: boolean | undefined;
Object.defineProperty(exports, "enableLogging", {
    get() {
        return __config.getObject<boolean>("enableLogging");
    },
    enumerable: true,
});

/**
 * A map of services with an endpoint to override what was from the Keystone catalog
 */
export declare const endpointOverrides: {[key: string]: any} | undefined;
Object.defineProperty(exports, "endpointOverrides", {
    get() {
        return __config.getObject<{[key: string]: any}>("endpointOverrides");
    },
    enumerable: true,
});

export declare const endpointType: string | undefined;
Object.defineProperty(exports, "endpointType", {
    get() {
        return __config.get("endpointType") ?? utilities.getEnv("OS_ENDPOINT_TYPE");
    },
    enumerable: true,
});

/**
 * Trust self-signed certificates.
 */
export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure") ?? utilities.getEnvBoolean("OS_INSECURE");
    },
    enumerable: true,
});

/**
 * A client private key to authenticate with.
 */
export declare const key: string | undefined;
Object.defineProperty(exports, "key", {
    get() {
        return __config.get("key");
    },
    enumerable: true,
});

/**
 * How many times HTTP connection should be retried until giving up.
 */
export declare const maxRetries: number | undefined;
Object.defineProperty(exports, "maxRetries", {
    get() {
        return __config.getObject<number>("maxRetries");
    },
    enumerable: true,
});

/**
 * Password to login with.
 */
export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password");
    },
    enumerable: true,
});

/**
 * The ID of the domain where the proejct resides (Identity v3).
 */
export declare const projectDomainId: string | undefined;
Object.defineProperty(exports, "projectDomainId", {
    get() {
        return __config.get("projectDomainId");
    },
    enumerable: true,
});

/**
 * The name of the domain where the project resides (Identity v3).
 */
export declare const projectDomainName: string | undefined;
Object.defineProperty(exports, "projectDomainName", {
    get() {
        return __config.get("projectDomainName");
    },
    enumerable: true,
});

/**
 * The OpenStack region to connect to.
 */
export declare const region: string | undefined;
Object.defineProperty(exports, "region", {
    get() {
        return __config.get("region") ?? utilities.getEnv("OS_REGION_NAME");
    },
    enumerable: true,
});

/**
 * Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
 */
export declare const swauth: boolean | undefined;
Object.defineProperty(exports, "swauth", {
    get() {
        return __config.getObject<boolean>("swauth") ?? utilities.getEnvBoolean("OS_SWAUTH");
    },
    enumerable: true,
});

/**
 * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
export declare const tenantId: string | undefined;
Object.defineProperty(exports, "tenantId", {
    get() {
        return __config.get("tenantId");
    },
    enumerable: true,
});

/**
 * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
export declare const tenantName: string | undefined;
Object.defineProperty(exports, "tenantName", {
    get() {
        return __config.get("tenantName");
    },
    enumerable: true,
});

/**
 * Authentication token to use as an alternative to username/password.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

/**
 * If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
 */
export declare const useOctavia: boolean | undefined;
Object.defineProperty(exports, "useOctavia", {
    get() {
        return __config.getObject<boolean>("useOctavia") ?? utilities.getEnvBoolean("OS_USE_OCTAVIA");
    },
    enumerable: true,
});

/**
 * The ID of the domain where the user resides (Identity v3).
 */
export declare const userDomainId: string | undefined;
Object.defineProperty(exports, "userDomainId", {
    get() {
        return __config.get("userDomainId");
    },
    enumerable: true,
});

/**
 * The name of the domain where the user resides (Identity v3).
 */
export declare const userDomainName: string | undefined;
Object.defineProperty(exports, "userDomainName", {
    get() {
        return __config.get("userDomainName");
    },
    enumerable: true,
});

/**
 * Username to login with.
 */
export declare const userId: string | undefined;
Object.defineProperty(exports, "userId", {
    get() {
        return __config.get("userId");
    },
    enumerable: true,
});

/**
 * Username to login with.
 */
export declare const userName: string | undefined;
Object.defineProperty(exports, "userName", {
    get() {
        return __config.get("userName");
    },
    enumerable: true,
});

