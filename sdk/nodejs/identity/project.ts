// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 Project resource within OpenStack Keystone.
 *
 * > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
 * this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const project1 = new openstack.identity.Project("project_1", {
 *     description: "A project",
 * });
 * ```
 *
 * ## Import
 *
 * Projects can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:identity/project:Project project_1 89c60255-9bd6-460c-822a-e2b959ede9d2
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:identity/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * A description of the project.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The domain this project belongs to.
     */
    public readonly domainId!: pulumi.Output<string>;
    /**
     * Whether the project is enabled or disabled. Valid
     * values are `true` and `false`. Default is `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Whether this project is a domain. Valid values
     * are `true` and `false`. Default is `false`. Changing this creates a new
     * project/domain.
     */
    public readonly isDomain!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parent of this project. Changing this creates
     * a new project.
     */
    public readonly parentId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new project.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Tags for the project. Changing this updates the existing
     * project.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["domainId"] = state ? state.domainId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["isDomain"] = state ? state.isDomain : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parentId"] = state ? state.parentId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["domainId"] = args ? args.domainId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["isDomain"] = args ? args.isDomain : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parentId"] = args ? args.parentId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Project.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * A description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain this project belongs to.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * Whether the project is enabled or disabled. Valid
     * values are `true` and `false`. Default is `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Whether this project is a domain. Valid values
     * are `true` and `false`. Default is `false`. Changing this creates a new
     * project/domain.
     */
    readonly isDomain?: pulumi.Input<boolean>;
    /**
     * The name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent of this project. Changing this creates
     * a new project.
     */
    readonly parentId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new project.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Tags for the project. Changing this updates the existing
     * project.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * A description of the project.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The domain this project belongs to.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * Whether the project is enabled or disabled. Valid
     * values are `true` and `false`. Default is `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Whether this project is a domain. Valid values
     * are `true` and `false`. Default is `false`. Changing this creates a new
     * project/domain.
     */
    readonly isDomain?: pulumi.Input<boolean>;
    /**
     * The name of the project.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The parent of this project. Changing this creates
     * a new project.
     */
    readonly parentId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new project.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Tags for the project. Changing this updates the existing
     * project.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
