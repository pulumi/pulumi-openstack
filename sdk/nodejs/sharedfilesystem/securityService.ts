// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to configure a security service.
 *
 * > **Note:** All arguments including the security service password will be
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
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const securityservice1 = new openstack.sharedfilesystem.SecurityService("securityservice_1", {
 *     name: "security",
 *     description: "created by terraform",
 *     type: "active_directory",
 *     server: "192.168.199.10",
 *     dnsIp: "192.168.199.10",
 *     domain: "example.com",
 *     ou: "CN=Computers,DC=example,DC=com",
 *     user: "joinDomainUser",
 *     password: "s8cret",
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported by specifying the ID of the security service:
 *
 * ```sh
 * $ pulumi import openstack:sharedfilesystem/securityService:SecurityService securityservice_1 id
 * ```
 */
export class SecurityService extends pulumi.CustomResource {
    /**
     * Get an existing SecurityService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityServiceState, opts?: pulumi.CustomResourceOptions): SecurityService {
        return new SecurityService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:sharedfilesystem/securityService:SecurityService';

    /**
     * Returns true if the given object is an instance of SecurityService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityService.__pulumiType;
    }

    /**
     * The human-readable description for the security service.
     * Changing this updates the description of the existing security service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The security service DNS IP address that is used inside the
     * tenant network.
     */
    public readonly dnsIp!: pulumi.Output<string | undefined>;
    /**
     * The security service domain.
     */
    public readonly domain!: pulumi.Output<string | undefined>;
    /**
     * The name of the security service. Changing this updates the name
     * of the existing security service.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The security service ou. An organizational unit can be added to
     * specify where the share ends up. New in Manila microversion 2.44.
     */
    public readonly ou!: pulumi.Output<string | undefined>;
    /**
     * The user password, if you specify a user.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The owner of the Security Service.
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a security service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security service.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The security service host name or IP address.
     */
    public readonly server!: pulumi.Output<string | undefined>;
    /**
     * The security service type - can either be active\_directory,
     * kerberos or ldap.  Changing this updates the existing security service.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The security service user or group name that is used by the
     * tenant.
     */
    public readonly user!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityServiceArgs | SecurityServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityServiceState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dnsIp"] = state ? state.dnsIp : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ou"] = state ? state.ou : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as SecurityServiceArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsIp"] = args ? args.dnsIp : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ou"] = args ? args.ou : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["projectId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SecurityService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityService resources.
 */
export interface SecurityServiceState {
    /**
     * The human-readable description for the security service.
     * Changing this updates the description of the existing security service.
     */
    description?: pulumi.Input<string>;
    /**
     * The security service DNS IP address that is used inside the
     * tenant network.
     */
    dnsIp?: pulumi.Input<string>;
    /**
     * The security service domain.
     */
    domain?: pulumi.Input<string>;
    /**
     * The name of the security service. Changing this updates the name
     * of the existing security service.
     */
    name?: pulumi.Input<string>;
    /**
     * The security service ou. An organizational unit can be added to
     * specify where the share ends up. New in Manila microversion 2.44.
     */
    ou?: pulumi.Input<string>;
    /**
     * The user password, if you specify a user.
     */
    password?: pulumi.Input<string>;
    /**
     * The owner of the Security Service.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a security service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security service.
     */
    region?: pulumi.Input<string>;
    /**
     * The security service host name or IP address.
     */
    server?: pulumi.Input<string>;
    /**
     * The security service type - can either be active\_directory,
     * kerberos or ldap.  Changing this updates the existing security service.
     */
    type?: pulumi.Input<string>;
    /**
     * The security service user or group name that is used by the
     * tenant.
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityService resource.
 */
export interface SecurityServiceArgs {
    /**
     * The human-readable description for the security service.
     * Changing this updates the description of the existing security service.
     */
    description?: pulumi.Input<string>;
    /**
     * The security service DNS IP address that is used inside the
     * tenant network.
     */
    dnsIp?: pulumi.Input<string>;
    /**
     * The security service domain.
     */
    domain?: pulumi.Input<string>;
    /**
     * The name of the security service. Changing this updates the name
     * of the existing security service.
     */
    name?: pulumi.Input<string>;
    /**
     * The security service ou. An organizational unit can be added to
     * specify where the share ends up. New in Manila microversion 2.44.
     */
    ou?: pulumi.Input<string>;
    /**
     * The user password, if you specify a user.
     */
    password?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a security service. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security service.
     */
    region?: pulumi.Input<string>;
    /**
     * The security service host name or IP address.
     */
    server?: pulumi.Input<string>;
    /**
     * The security service type - can either be active\_directory,
     * kerberos or ldap.  Changing this updates the existing security service.
     */
    type: pulumi.Input<string>;
    /**
     * The security service user or group name that is used by the
     * tenant.
     */
    user?: pulumi.Input<string>;
}
