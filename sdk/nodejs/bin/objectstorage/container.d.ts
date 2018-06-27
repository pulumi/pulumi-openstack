import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V1 container resource within OpenStack.
 */
export declare class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerState): Container;
    /**
     * Sets an access control list (ACL) that grants
     * read access. This header can contain a comma-delimited list of users that
     * can read the container (allows the GET method for all objects in the
     * container). Changing this updates the access control list read access.
     */
    readonly containerRead: pulumi.Output<string | undefined>;
    /**
     * The secret key for container synchronization.
     * Changing this updates container synchronization.
     */
    readonly containerSyncKey: pulumi.Output<string | undefined>;
    /**
     * The destination for container synchronization.
     * Changing this updates container synchronization.
     */
    readonly containerSyncTo: pulumi.Output<string | undefined>;
    /**
     * Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     */
    readonly containerWrite: pulumi.Output<string | undefined>;
    /**
     * The MIME type for the container. Changing this
     * updates the MIME type.
     */
    readonly contentType: pulumi.Output<string | undefined>;
    /**
     * A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     */
    readonly forceDestroy: pulumi.Output<boolean | undefined>;
    /**
     * Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     */
    readonly metadata: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * A unique name for the container. Changing this creates a
     * new container.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    readonly region: pulumi.Output<string>;
    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Container resources.
 */
export interface ContainerState {
    /**
     * Sets an access control list (ACL) that grants
     * read access. This header can contain a comma-delimited list of users that
     * can read the container (allows the GET method for all objects in the
     * container). Changing this updates the access control list read access.
     */
    readonly containerRead?: pulumi.Input<string>;
    /**
     * The secret key for container synchronization.
     * Changing this updates container synchronization.
     */
    readonly containerSyncKey?: pulumi.Input<string>;
    /**
     * The destination for container synchronization.
     * Changing this updates container synchronization.
     */
    readonly containerSyncTo?: pulumi.Input<string>;
    /**
     * Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     */
    readonly containerWrite?: pulumi.Input<string>;
    /**
     * The MIME type for the container. Changing this
     * updates the MIME type.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     */
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A unique name for the container. Changing this creates a
     * new container.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    readonly region?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Container resource.
 */
export interface ContainerArgs {
    /**
     * Sets an access control list (ACL) that grants
     * read access. This header can contain a comma-delimited list of users that
     * can read the container (allows the GET method for all objects in the
     * container). Changing this updates the access control list read access.
     */
    readonly containerRead?: pulumi.Input<string>;
    /**
     * The secret key for container synchronization.
     * Changing this updates container synchronization.
     */
    readonly containerSyncKey?: pulumi.Input<string>;
    /**
     * The destination for container synchronization.
     * Changing this updates container synchronization.
     */
    readonly containerSyncTo?: pulumi.Input<string>;
    /**
     * Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     */
    readonly containerWrite?: pulumi.Input<string>;
    /**
     * The MIME type for the container. Changing this
     * updates the MIME type.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     */
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A unique name for the container. Changing this creates a
     * new container.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    readonly region?: pulumi.Input<string>;
}
