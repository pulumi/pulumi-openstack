// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the openstack package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'openstack';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * Application Credential ID to login with.
     */
    public readonly applicationCredentialId!: pulumi.Output<string | undefined>;
    /**
     * Application Credential name to login with.
     */
    public readonly applicationCredentialName!: pulumi.Output<string | undefined>;
    /**
     * Application Credential secret to login with.
     */
    public readonly applicationCredentialSecret!: pulumi.Output<string | undefined>;
    /**
     * The Identity authentication URL.
     */
    public readonly authUrl!: pulumi.Output<string | undefined>;
    /**
     * A Custom CA certificate.
     */
    public readonly cacertFile!: pulumi.Output<string | undefined>;
    /**
     * A client certificate to authenticate with.
     */
    public readonly cert!: pulumi.Output<string | undefined>;
    /**
     * An entry in a `clouds.yaml` file to use.
     */
    public readonly cloud!: pulumi.Output<string | undefined>;
    /**
     * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
     */
    public readonly defaultDomain!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Domain to scope to (Identity v3).
     */
    public readonly domainId!: pulumi.Output<string | undefined>;
    /**
     * The name of the Domain to scope to (Identity v3).
     */
    public readonly domainName!: pulumi.Output<string | undefined>;
    public readonly endpointType!: pulumi.Output<string | undefined>;
    /**
     * A client private key to authenticate with.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Password to login with.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The ID of the domain where the proejct resides (Identity v3).
     */
    public readonly projectDomainId!: pulumi.Output<string | undefined>;
    /**
     * The name of the domain where the project resides (Identity v3).
     */
    public readonly projectDomainName!: pulumi.Output<string | undefined>;
    /**
     * The OpenStack region to connect to.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
     */
    public readonly tenantName!: pulumi.Output<string | undefined>;
    /**
     * Authentication token to use as an alternative to username/password.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * The ID of the domain where the user resides (Identity v3).
     */
    public readonly userDomainId!: pulumi.Output<string | undefined>;
    /**
     * The name of the domain where the user resides (Identity v3).
     */
    public readonly userDomainName!: pulumi.Output<string | undefined>;
    /**
     * Username to login with.
     */
    public readonly userId!: pulumi.Output<string | undefined>;
    /**
     * Username to login with.
     */
    public readonly userName!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["allowReauth"] = pulumi.output((args ? args.allowReauth : undefined) ?? utilities.getEnvBoolean("OS_ALLOW_REAUTH")).apply(JSON.stringify);
            resourceInputs["applicationCredentialId"] = args ? args.applicationCredentialId : undefined;
            resourceInputs["applicationCredentialName"] = args ? args.applicationCredentialName : undefined;
            resourceInputs["applicationCredentialSecret"] = args ? args.applicationCredentialSecret : undefined;
            resourceInputs["authUrl"] = args ? args.authUrl : undefined;
            resourceInputs["cacertFile"] = args ? args.cacertFile : undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["cloud"] = (args ? args.cloud : undefined) ?? utilities.getEnv("OS_CLOUD");
            resourceInputs["defaultDomain"] = args ? args.defaultDomain : undefined;
            resourceInputs["delayedAuth"] = pulumi.output((args ? args.delayedAuth : undefined) ?? utilities.getEnvBoolean("OS_DELAYED_AUTH")).apply(JSON.stringify);
            resourceInputs["disableNoCacheHeader"] = pulumi.output(args ? args.disableNoCacheHeader : undefined).apply(JSON.stringify);
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["endpointOverrides"] = pulumi.output(args ? args.endpointOverrides : undefined).apply(JSON.stringify);
            resourceInputs["endpointType"] = (args ? args.endpointType : undefined) ?? utilities.getEnv("OS_ENDPOINT_TYPE");
            resourceInputs["insecure"] = pulumi.output((args ? args.insecure : undefined) ?? utilities.getEnvBoolean("OS_INSECURE")).apply(JSON.stringify);
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["maxRetries"] = pulumi.output(args ? args.maxRetries : undefined).apply(JSON.stringify);
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["projectDomainId"] = args ? args.projectDomainId : undefined;
            resourceInputs["projectDomainName"] = args ? args.projectDomainName : undefined;
            resourceInputs["region"] = (args ? args.region : undefined) ?? utilities.getEnv("OS_REGION_NAME");
            resourceInputs["swauth"] = pulumi.output((args ? args.swauth : undefined) ?? utilities.getEnvBoolean("OS_SWAUTH")).apply(JSON.stringify);
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["tenantName"] = args ? args.tenantName : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["useOctavia"] = pulumi.output((args ? args.useOctavia : undefined) ?? utilities.getEnvBoolean("OS_USE_OCTAVIA")).apply(JSON.stringify);
            resourceInputs["userDomainId"] = args ? args.userDomainId : undefined;
            resourceInputs["userDomainName"] = args ? args.userDomainName : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * If set to `false`, OpenStack authorization won't be perfomed automatically, if the initial auth token get expired.
     * Defaults to `true`
     */
    allowReauth?: pulumi.Input<boolean>;
    /**
     * Application Credential ID to login with.
     */
    applicationCredentialId?: pulumi.Input<string>;
    /**
     * Application Credential name to login with.
     */
    applicationCredentialName?: pulumi.Input<string>;
    /**
     * Application Credential secret to login with.
     */
    applicationCredentialSecret?: pulumi.Input<string>;
    /**
     * The Identity authentication URL.
     */
    authUrl?: pulumi.Input<string>;
    /**
     * A Custom CA certificate.
     */
    cacertFile?: pulumi.Input<string>;
    /**
     * A client certificate to authenticate with.
     */
    cert?: pulumi.Input<string>;
    /**
     * An entry in a `clouds.yaml` file to use.
     */
    cloud?: pulumi.Input<string>;
    /**
     * The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
     */
    defaultDomain?: pulumi.Input<string>;
    /**
     * If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
     * to `true`.
     */
    delayedAuth?: pulumi.Input<boolean>;
    /**
     * If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
     */
    disableNoCacheHeader?: pulumi.Input<boolean>;
    /**
     * The ID of the Domain to scope to (Identity v3).
     */
    domainId?: pulumi.Input<string>;
    /**
     * The name of the Domain to scope to (Identity v3).
     */
    domainName?: pulumi.Input<string>;
    /**
     * A map of services with an endpoint to override what was from the Keystone catalog
     */
    endpointOverrides?: pulumi.Input<{[key: string]: any}>;
    endpointType?: pulumi.Input<string>;
    /**
     * Trust self-signed certificates.
     */
    insecure?: pulumi.Input<boolean>;
    /**
     * A client private key to authenticate with.
     */
    key?: pulumi.Input<string>;
    /**
     * How many times HTTP connection should be retried until giving up.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Password to login with.
     */
    password?: pulumi.Input<string>;
    /**
     * The ID of the domain where the proejct resides (Identity v3).
     */
    projectDomainId?: pulumi.Input<string>;
    /**
     * The name of the domain where the project resides (Identity v3).
     */
    projectDomainName?: pulumi.Input<string>;
    /**
     * The OpenStack region to connect to.
     */
    region?: pulumi.Input<string>;
    /**
     * Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
     */
    swauth?: pulumi.Input<boolean>;
    /**
     * The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
     */
    tenantName?: pulumi.Input<string>;
    /**
     * Authentication token to use as an alternative to username/password.
     */
    token?: pulumi.Input<string>;
    /**
     * If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
     */
    useOctavia?: pulumi.Input<boolean>;
    /**
     * The ID of the domain where the user resides (Identity v3).
     */
    userDomainId?: pulumi.Input<string>;
    /**
     * The name of the domain where the user resides (Identity v3).
     */
    userDomainName?: pulumi.Input<string>;
    /**
     * Username to login with.
     */
    userId?: pulumi.Input<string>;
    /**
     * Username to login with.
     */
    userName?: pulumi.Input<string>;
}
