// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getAuthScope(args: GetAuthScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthScopeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:identity/getAuthScope:getAuthScope", {
        "name": args.name,
        "region": args.region,
        "setTokenId": args.setTokenId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthScope.
 */
export interface GetAuthScopeArgs {
    /**
     * The name of the scope. This is an arbitrary name which is
     * only used as a unique identifier so an actual token isn't used as the ID.
     */
    name: string;
    /**
     * The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: string;
    /**
     * A boolean argument that determines whether to
     * export the current auth scope token ID. When set to `true`, the `tokenId`
     * attribute will contain an unencrypted token that can be used for further API
     * calls. **Warning**: please note that the leaked token may allow unauthorized
     * access to other OpenStack services within the current auth scope, so use this
     * option with caution.
     */
    setTokenId?: boolean;
}

/**
 * A collection of values returned by getAuthScope.
 */
export interface GetAuthScopeResult {
    /**
     * The domain ID of the scope.
     */
    readonly domainId: string;
    /**
     * The domain name of the scope.
     */
    readonly domainName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the service.
     */
    readonly name: string;
    /**
     * The domain ID of the project.
     */
    readonly projectDomainId: string;
    /**
     * The domain name of the project.
     */
    readonly projectDomainName: string;
    /**
     * The project ID of the scope.
     */
    readonly projectId: string;
    /**
     * The project name of the scope.
     */
    readonly projectName: string;
    /**
     * The region of the endpoint.
     */
    readonly region: string;
    /**
     * A list of roles in the current scope. See reference below.
     */
    readonly roles: outputs.identity.GetAuthScopeRole[];
    /**
     * A list of service catalog entries returned with the token.
     */
    readonly serviceCatalogs: outputs.identity.GetAuthScopeServiceCatalog[];
    readonly setTokenId?: boolean;
    /**
     * The token ID of the scope.
     */
    readonly tokenId: string;
    /**
     * The domain ID of the user.
     */
    readonly userDomainId: string;
    /**
     * The domain name of the user.
     */
    readonly userDomainName: string;
    /**
     * The user ID the of the scope.
     */
    readonly userId: string;
    /**
     * The username of the scope.
     */
    readonly userName: string;
}
export function getAuthScopeOutput(args: GetAuthScopeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAuthScopeResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("openstack:identity/getAuthScope:getAuthScope", {
        "name": args.name,
        "region": args.region,
        "setTokenId": args.setTokenId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthScope.
 */
export interface GetAuthScopeOutputArgs {
    /**
     * The name of the scope. This is an arbitrary name which is
     * only used as a unique identifier so an actual token isn't used as the ID.
     */
    name: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     */
    region?: pulumi.Input<string>;
    /**
     * A boolean argument that determines whether to
     * export the current auth scope token ID. When set to `true`, the `tokenId`
     * attribute will contain an unencrypted token that can be used for further API
     * calls. **Warning**: please note that the leaked token may allow unauthorized
     * access to other OpenStack services within the current auth scope, so use this
     * option with caution.
     */
    setTokenId?: pulumi.Input<boolean>;
}
