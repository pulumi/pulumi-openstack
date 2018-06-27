import * as pulumi from "@pulumi/pulumi";
/**
 * Use this data source to get the ID of an OpenStack project.
 */
export declare function getProject(args?: GetProjectArgs): Promise<GetProjectResult>;
/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * The domain this project belongs to.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * Whether the project is enabled or disabled. Valid
     * values are `true` and `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Whether this project is a domain. Valid values
     * are `true` and `false`.
     */
    readonly isDomain?: pulumi.Input<boolean>;
    /**
     * The name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent of this project.
     */
    readonly parentId?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
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
     * The region the project is located in.
     */
    readonly region: string;
}
