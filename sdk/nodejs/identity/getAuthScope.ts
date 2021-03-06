// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get authentication information about the current
 * auth scope in use. This can be used as self-discovery or introspection of
 * the username or project name currently in use as well as the service catalog.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const scope = pulumi.output(openstack.identity.getAuthScope({
 *     name: "my_scope",
 * }, { async: true }));
 * ```
 *
 * To find the the public object storage endpoint for "region1" as listed in the
 * service catalog:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 *
 * const objectStoreService = .filter(entry => entry.type == "object-store").map(entry => entry)[0];
 * const objectStoreEndpoint = .filter(endpoint => endpoint["interface"] == "public" && endpoint.region == "region1").map(endpoint => endpoint)[0];
 * const objectStorePublicUrl = objectStoreEndpoint.url;
 * ```
 */
export function getAuthScope(args: GetAuthScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthScopeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:identity/getAuthScope:getAuthScope", {
        "name": args.name,
        "region": args.region,
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
    readonly name: string;
    /**
     * The region in which to obtain the V3 Identity client.
     * A Identity client is needed to retrieve tokens IDs. If omitted, the
     * `region` argument of the provider is used.
     */
    readonly region?: string;
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
