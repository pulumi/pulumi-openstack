/**
 * The Identity authentication URL.
 */
export declare let authUrl: string | undefined;
/**
 * A Custom CA certificate.
 */
export declare let cacertFile: string | undefined;
/**
 * A client certificate to authenticate with.
 */
export declare let cert: string | undefined;
/**
 * An entry in a `clouds.yaml` file to use.
 */
export declare let cloud: string | undefined;
/**
 * The ID of the Domain to scope to (Identity v3).
 */
export declare let domainId: string | undefined;
/**
 * The name of the Domain to scope to (Identity v3).
 */
export declare let domainName: string | undefined;
export declare let endpointType: string | undefined;
/**
 * Trust self-signed certificates.
 */
export declare let insecure: boolean | undefined;
/**
 * A client private key to authenticate with.
 */
export declare let key: string | undefined;
/**
 * Password to login with.
 */
export declare let password: string | undefined;
/**
 * The ID of the domain where the proejct resides (Identity v3).
 */
export declare let projectDomainId: string | undefined;
/**
 * The name of the domain where the project resides (Identity v3).
 */
export declare let projectDomainName: string | undefined;
/**
 * The OpenStack region to connect to.
 */
export declare let region: string | undefined;
/**
 * Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
 */
export declare let swauth: boolean | undefined;
/**
 * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
export declare let tenantId: string | undefined;
/**
 * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
 */
export declare let tenantName: string | undefined;
/**
 * Authentication token to use as an alternative to username/password.
 */
export declare let token: string | undefined;
/**
 * If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
 */
export declare let useOctavia: boolean | undefined;
/**
 * The ID of the domain where the user resides (Identity v3).
 */
export declare let userDomainId: string | undefined;
/**
 * The name of the domain where the user resides (Identity v3).
 */
export declare let userDomainName: string | undefined;
/**
 * Username to login with.
 */
export declare let userId: string | undefined;
/**
 * Username to login with.
 */
export declare let userName: string | undefined;
