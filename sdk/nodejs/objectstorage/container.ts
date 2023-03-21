// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V1 container resource within OpenStack.
 *
 * ## Example Usage
 * ### Basic Container
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const container1 = new openstack.objectstorage.Container("container1", {
 *     contentType: "application/json",
 *     metadata: {
 *         test: "true",
 *     },
 *     region: "RegionOne",
 *     versioning: true,
 * });
 * ```
 * ### Basic Container with legacy versioning
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const container1 = new openstack.objectstorage.Container("container1", {
 *     contentType: "application/json",
 *     metadata: {
 *         test: "true",
 *     },
 *     region: "RegionOne",
 *     versioningLegacy: {
 *         location: "tf-test-container-versions",
 *         type: "versions",
 *     },
 * });
 * ```
 * ### Global Read Access
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const container1 = new openstack.objectstorage.Container("container1", {
 *     containerRead: ".r:*",
 *     region: "RegionOne",
 * });
 * ```
 * ### Global Read and List Access
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const container1 = new openstack.objectstorage.Container("container1", {
 *     containerRead: ".r:*,.rlistings",
 *     region: "RegionOne",
 * });
 * ```
 * ### Write-Only Access for a User
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const current = openstack.identity.getAuthScope({
 *     name: "current",
 * });
 * const container1 = new openstack.objectstorage.Container("container1", {
 *     containerRead: `.r:-${_var.username}`,
 *     containerWrite: current.then(current => `${current.projectId}:${_var.username}`),
 *     region: "RegionOne",
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported by specifying the name of the containerSome attributes can't be imported * `force_destroy` * `content_type` * `metadata` * `container_sync_to` * `container_sync_key` So you'll have to `terraform plan` and `terraform apply` after the import to fix those missing attributes.
 *
 * ```sh
 *  $ pulumi import openstack:objectstorage/container:Container container_1 <name>
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
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * The storage policy to be used for the container. 
     * Changing this creates a new container.
     */
    public readonly storagePolicy!: pulumi.Output<string>;
    /**
     * A boolean that can enable or disable object
     * versioning. The default value is `false`. To use this feature, your Swift
     * version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
     * and a cloud administrator must have set the `allowObjectVersioning = true`
     * configuration option in Swift. If you cannot set this versioning type, you may
     * want to consider using `versioningLegacy` instead.
     */
    public readonly versioning!: pulumi.Output<boolean | undefined>;
    /**
     * Enable legacy object versioning. The structure is described below.
     *
     * @deprecated Use newer "versioning" implementation
     */
    public readonly versioningLegacy!: pulumi.Output<outputs.objectstorage.ContainerVersioningLegacy | undefined>;

    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerArgs | ContainerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerState | undefined;
            resourceInputs["containerRead"] = state ? state.containerRead : undefined;
            resourceInputs["containerSyncKey"] = state ? state.containerSyncKey : undefined;
            resourceInputs["containerSyncTo"] = state ? state.containerSyncTo : undefined;
            resourceInputs["containerWrite"] = state ? state.containerWrite : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["storagePolicy"] = state ? state.storagePolicy : undefined;
            resourceInputs["versioning"] = state ? state.versioning : undefined;
            resourceInputs["versioningLegacy"] = state ? state.versioningLegacy : undefined;
        } else {
            const args = argsOrState as ContainerArgs | undefined;
            resourceInputs["containerRead"] = args ? args.containerRead : undefined;
            resourceInputs["containerSyncKey"] = args ? args.containerSyncKey : undefined;
            resourceInputs["containerSyncTo"] = args ? args.containerSyncTo : undefined;
            resourceInputs["containerWrite"] = args ? args.containerWrite : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["storagePolicy"] = args ? args.storagePolicy : undefined;
            resourceInputs["versioning"] = args ? args.versioning : undefined;
            resourceInputs["versioningLegacy"] = args ? args.versioningLegacy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Container.__pulumiType, name, resourceInputs, opts);
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
    containerRead?: pulumi.Input<string>;
    /**
     * The secret key for container synchronization.
     * Changing this updates container synchronization.
     */
    containerSyncKey?: pulumi.Input<string>;
    /**
     * The destination for container synchronization.
     * Changing this updates container synchronization.
     */
    containerSyncTo?: pulumi.Input<string>;
    /**
     * Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     */
    containerWrite?: pulumi.Input<string>;
    /**
     * The MIME type for the container. Changing this
     * updates the MIME type.
     */
    contentType?: pulumi.Input<string>;
    /**
     * A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * A unique name for the container. Changing this creates a
     * new container.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    region?: pulumi.Input<string>;
    /**
     * The storage policy to be used for the container. 
     * Changing this creates a new container.
     */
    storagePolicy?: pulumi.Input<string>;
    /**
     * A boolean that can enable or disable object
     * versioning. The default value is `false`. To use this feature, your Swift
     * version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
     * and a cloud administrator must have set the `allowObjectVersioning = true`
     * configuration option in Swift. If you cannot set this versioning type, you may
     * want to consider using `versioningLegacy` instead.
     */
    versioning?: pulumi.Input<boolean>;
    /**
     * Enable legacy object versioning. The structure is described below.
     *
     * @deprecated Use newer "versioning" implementation
     */
    versioningLegacy?: pulumi.Input<inputs.objectstorage.ContainerVersioningLegacy>;
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
    containerRead?: pulumi.Input<string>;
    /**
     * The secret key for container synchronization.
     * Changing this updates container synchronization.
     */
    containerSyncKey?: pulumi.Input<string>;
    /**
     * The destination for container synchronization.
     * Changing this updates container synchronization.
     */
    containerSyncTo?: pulumi.Input<string>;
    /**
     * Sets an ACL that grants write access.
     * Changing this updates the access control list write access.
     */
    containerWrite?: pulumi.Input<string>;
    /**
     * The MIME type for the container. Changing this
     * updates the MIME type.
     */
    contentType?: pulumi.Input<string>;
    /**
     * A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * Custom key/value pairs to associate with the container.
     * Changing this updates the existing container metadata.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * A unique name for the container. Changing this creates a
     * new container.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to create the container. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new container.
     */
    region?: pulumi.Input<string>;
    /**
     * The storage policy to be used for the container. 
     * Changing this creates a new container.
     */
    storagePolicy?: pulumi.Input<string>;
    /**
     * A boolean that can enable or disable object
     * versioning. The default value is `false`. To use this feature, your Swift
     * version must be 2.24 or higher (as described in the [OpenStack Swift Ussuri release notes](https://docs.openstack.org/releasenotes/swift/ussuri.html#relnotes-2-24-0-stable-ussuri)),
     * and a cloud administrator must have set the `allowObjectVersioning = true`
     * configuration option in Swift. If you cannot set this versioning type, you may
     * want to consider using `versioningLegacy` instead.
     */
    versioning?: pulumi.Input<boolean>;
    /**
     * Enable legacy object versioning. The structure is described below.
     *
     * @deprecated Use newer "versioning" implementation
     */
    versioningLegacy?: pulumi.Input<inputs.objectstorage.ContainerVersioningLegacy>;
}
