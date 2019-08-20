// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an OpenStack project.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const project1 = openstack.identity.getProject({
 *     name: "demo",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_project_v3.html.markdown.
 */
export function getProject(args?: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> & GetProjectResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetProjectResult> = pulumi.runtime.invoke("openstack:identity/getProject:getProject", {
        "domainId": args.domainId,
        "enabled": args.enabled,
        "isDomain": args.isDomain,
        "name": args.name,
        "parentId": args.parentId,
        "region": args.region,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * The domain this project belongs to.
     */
    readonly domainId?: string;
    /**
     * Whether the project is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    readonly enabled?: boolean;
    /**
     * Whether this project is a domain. Valid values
     * are `true` and `false`.
     */
    readonly isDomain?: boolean;
    /**
     * The name of the project.
     */
    readonly name?: string;
    /**
     * The parent of this project.
     */
    readonly parentId?: string;
    readonly region?: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * The description of the project.
     */
    readonly description: string;
    /**
     * See Argument Reference above.
     */
    readonly domainId: string;
    /**
     * See Argument Reference above.
     */
    readonly enabled?: boolean;
    /**
     * See Argument Reference above.
     */
    readonly isDomain?: boolean;
    /**
     * See Argument Reference above.
     */
    readonly name?: string;
    /**
     * See Argument Reference above.
     */
    readonly parentId?: string;
    /**
     * The region the project is located in.
     */
    readonly region: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
