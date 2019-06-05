// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V1 container resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const container1 = new openstack.objectstorage.Container("container_1", {
 *     contentType: "application/json",
 *     metadata: {
 *         test: "true",
 *     },
 *     region: "RegionOne",
 *     versioning: {
 *         location: "tf-test-container-versions",
 *         type: "versions",
 *     },
 * });
 * ```
 */
export class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerState, opts?: pulumi.CustomResourceOptions): Container {
        return new Container(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:objectstorage/container:Container';

    /**
     * Returns true if the given object is an instance of Container.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Container {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Container.__pulumiType;
    }

    /**
     * Sets an access control list (ACL) that grants
     * read access. This header can contain a comma-delimited list of users that
     * can read the container (allows the GET method for all objects in the
     * container). Changing this updates the access control list read access.
     */
    public readonly containerRead!: pulumi.Output<string | undefined>;
    /**
     * The secret key for container synchronization.
     * Changing this updates container synchronization.
     */
    public readonly containerSyncKey!: pulumi.Output<string | undefined>;
    /**
     * The destination for container synchronization.
     * Changing this updates container synchronization.
     */
    public readonly containerSyncTo!: pulumi.Output<string | undefined>;
    /**
     * Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     */
    public readonly containerWrite!: pulumi.Output<string | undefined>;
    /**
     * The MIME type for the container. Changing this
     * updates the MIME type.
     */
    public readonly contentType!: pulumi.Output<string | undefined>;
    /**
     * A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A unique name for the container. Changing this creates a
     * new container.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Enable object versioning. The structure is described below.
     */
    public readonly versioning!: pulumi.Output<{ location: string, type: string } | undefined>;

    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerArgs | ContainerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ContainerState | undefined;
            inputs["containerRead"] = state ? state.containerRead : undefined;
            inputs["containerSyncKey"] = state ? state.containerSyncKey : undefined;
            inputs["containerSyncTo"] = state ? state.containerSyncTo : undefined;
            inputs["containerWrite"] = state ? state.containerWrite : undefined;
            inputs["contentType"] = state ? state.contentType : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["versioning"] = state ? state.versioning : undefined;
        } else {
            const args = argsOrState as ContainerArgs | undefined;
            inputs["containerRead"] = args ? args.containerRead : undefined;
            inputs["containerSyncKey"] = args ? args.containerSyncKey : undefined;
            inputs["containerSyncTo"] = args ? args.containerSyncTo : undefined;
            inputs["containerWrite"] = args ? args.containerWrite : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["versioning"] = args ? args.versioning : undefined;
        }
        super(Container.__pulumiType, name, inputs, opts);
    }
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
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
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
    /**
     * Enable object versioning. The structure is described below.
     */
    readonly versioning?: pulumi.Input<{ location: pulumi.Input<string>, type: pulumi.Input<string> }>;
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
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
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
    /**
     * Enable object versioning. The structure is described below.
     */
    readonly versioning?: pulumi.Input<{ location: pulumi.Input<string>, type: pulumi.Input<string> }>;
}
