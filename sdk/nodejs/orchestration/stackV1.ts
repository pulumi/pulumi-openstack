// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V1 stack resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const stack1 = new openstack.orchestration.StackV1("stack_1", {
 *     name: "stack_1",
 *     parameters: {
 *         length: "4",
 *     },
 *     templateOpts: {
 *         Bin: `heat_template_version: 2013-05-23
 * parameters:
 *   length:
 *     type: number
 * resources:
 *   test_res:
 *     type: OS::Heat::TestResource
 *   random:
 *     type: OS::Heat::RandomString
 *     properties:
 *       length: {get_param: length}
 * `,
 *     },
 *     environmentOpts: {
 *         Bin: "\n",
 *     },
 *     disableRollback: true,
 *     timeout: 30,
 * });
 * ```
 *
 * ## Import
 *
 * stacks can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:orchestration/stackV1:StackV1 stack_1 ea257959-eeb1-4c10-8d33-26f0409a755d
 * ```
 */
export class StackV1 extends pulumi.CustomResource {
    /**
     * Get an existing StackV1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackV1State, opts?: pulumi.CustomResourceOptions): StackV1 {
        return new StackV1(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:orchestration/stackV1:StackV1';

    /**
     * Returns true if the given object is an instance of StackV1.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StackV1 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StackV1.__pulumiType;
    }

    /**
     * A list of stack outputs.
     */
    public readonly StackOutputs!: pulumi.Output<outputs.orchestration.StackV1StackOutput[]>;
    /**
     * List of stack capabilities for stack.
     */
    public readonly capabilities!: pulumi.Output<string[]>;
    /**
     * The date and time when the resource was created. The date
     * and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
     * is the time zone as an offset from UTC.
     */
    public readonly creationTime!: pulumi.Output<string>;
    /**
     * The description of the stack resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Enables or disables deletion of all stack
     * resources when a stack creation fails. Default is true, meaning all
     * resources are not deleted when stack creation fails.
     */
    public readonly disableRollback!: pulumi.Output<boolean>;
    /**
     * Environment key/value pairs to associate with
     * the stack which contains details for the environment of the stack.
     * Allowed keys: Bin, URL, Files. Changing this updates the existing stack
     * Environment Opts.
     */
    public readonly environmentOpts!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique name for the stack. It must start with an
     * alphabetic character. Changing this updates the stack's name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of notification topics for stack.
     */
    public readonly notificationTopics!: pulumi.Output<string[]>;
    /**
     * User-defined key/value pairs as parameters to pass
     * to the template. Changing this updates the existing stack parameters.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The region in which to create the stack. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new stack.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The status of the stack.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * The reason for the current status of the stack.
     */
    public readonly statusReason!: pulumi.Output<string>;
    /**
     * A list of tags to assosciate with the Stack
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The description of the stack template.
     */
    public readonly templateDescription!: pulumi.Output<string>;
    /**
     * Template key/value pairs to associate with the
     * stack which contains either the template file or url.
     * Allowed keys: Bin, URL, Files. Changing this updates the existing stack
     * Template Opts.
     */
    public readonly templateOpts!: pulumi.Output<{[key: string]: string}>;
    /**
     * The timeout for stack action in minutes.
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * The date and time when the resource was updated. The date
     * and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
     * is the time zone as an offset from UTC.
     */
    public readonly updatedTime!: pulumi.Output<string>;

    /**
     * Create a StackV1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StackV1Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackV1Args | StackV1State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackV1State | undefined;
            resourceInputs["StackOutputs"] = state ? state.StackOutputs : undefined;
            resourceInputs["capabilities"] = state ? state.capabilities : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableRollback"] = state ? state.disableRollback : undefined;
            resourceInputs["environmentOpts"] = state ? state.environmentOpts : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationTopics"] = state ? state.notificationTopics : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusReason"] = state ? state.statusReason : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["templateDescription"] = state ? state.templateDescription : undefined;
            resourceInputs["templateOpts"] = state ? state.templateOpts : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["updatedTime"] = state ? state.updatedTime : undefined;
        } else {
            const args = argsOrState as StackV1Args | undefined;
            if ((!args || args.templateOpts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateOpts'");
            }
            resourceInputs["StackOutputs"] = args ? args.StackOutputs : undefined;
            resourceInputs["capabilities"] = args ? args.capabilities : undefined;
            resourceInputs["creationTime"] = args ? args.creationTime : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableRollback"] = args ? args.disableRollback : undefined;
            resourceInputs["environmentOpts"] = args ? args.environmentOpts : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationTopics"] = args ? args.notificationTopics : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["statusReason"] = args ? args.statusReason : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["templateDescription"] = args ? args.templateDescription : undefined;
            resourceInputs["templateOpts"] = args ? args.templateOpts : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["updatedTime"] = args ? args.updatedTime : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StackV1.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StackV1 resources.
 */
export interface StackV1State {
    /**
     * A list of stack outputs.
     */
    StackOutputs?: pulumi.Input<pulumi.Input<inputs.orchestration.StackV1StackOutput>[]>;
    /**
     * List of stack capabilities for stack.
     */
    capabilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time when the resource was created. The date
     * and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
     * is the time zone as an offset from UTC.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The description of the stack resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Enables or disables deletion of all stack
     * resources when a stack creation fails. Default is true, meaning all
     * resources are not deleted when stack creation fails.
     */
    disableRollback?: pulumi.Input<boolean>;
    /**
     * Environment key/value pairs to associate with
     * the stack which contains details for the environment of the stack.
     * Allowed keys: Bin, URL, Files. Changing this updates the existing stack
     * Environment Opts.
     */
    environmentOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the stack. It must start with an
     * alphabetic character. Changing this updates the stack's name.
     */
    name?: pulumi.Input<string>;
    /**
     * List of notification topics for stack.
     */
    notificationTopics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User-defined key/value pairs as parameters to pass
     * to the template. Changing this updates the existing stack parameters.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The region in which to create the stack. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new stack.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the stack.
     */
    status?: pulumi.Input<string>;
    /**
     * The reason for the current status of the stack.
     */
    statusReason?: pulumi.Input<string>;
    /**
     * A list of tags to assosciate with the Stack
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the stack template.
     */
    templateDescription?: pulumi.Input<string>;
    /**
     * Template key/value pairs to associate with the
     * stack which contains either the template file or url.
     * Allowed keys: Bin, URL, Files. Changing this updates the existing stack
     * Template Opts.
     */
    templateOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The timeout for stack action in minutes.
     */
    timeout?: pulumi.Input<number>;
    /**
     * The date and time when the resource was updated. The date
     * and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
     * is the time zone as an offset from UTC.
     */
    updatedTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StackV1 resource.
 */
export interface StackV1Args {
    /**
     * A list of stack outputs.
     */
    StackOutputs?: pulumi.Input<pulumi.Input<inputs.orchestration.StackV1StackOutput>[]>;
    /**
     * List of stack capabilities for stack.
     */
    capabilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time when the resource was created. The date
     * and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
     * is the time zone as an offset from UTC.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The description of the stack resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Enables or disables deletion of all stack
     * resources when a stack creation fails. Default is true, meaning all
     * resources are not deleted when stack creation fails.
     */
    disableRollback?: pulumi.Input<boolean>;
    /**
     * Environment key/value pairs to associate with
     * the stack which contains details for the environment of the stack.
     * Allowed keys: Bin, URL, Files. Changing this updates the existing stack
     * Environment Opts.
     */
    environmentOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the stack. It must start with an
     * alphabetic character. Changing this updates the stack's name.
     */
    name?: pulumi.Input<string>;
    /**
     * List of notification topics for stack.
     */
    notificationTopics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User-defined key/value pairs as parameters to pass
     * to the template. Changing this updates the existing stack parameters.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The region in which to create the stack. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new stack.
     */
    region?: pulumi.Input<string>;
    /**
     * The status of the stack.
     */
    status?: pulumi.Input<string>;
    /**
     * The reason for the current status of the stack.
     */
    statusReason?: pulumi.Input<string>;
    /**
     * A list of tags to assosciate with the Stack
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the stack template.
     */
    templateDescription?: pulumi.Input<string>;
    /**
     * Template key/value pairs to associate with the
     * stack which contains either the template file or url.
     * Allowed keys: Bin, URL, Files. Changing this updates the existing stack
     * Template Opts.
     */
    templateOpts: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The timeout for stack action in minutes.
     */
    timeout?: pulumi.Input<number>;
    /**
     * The date and time when the resource was updated. The date
     * and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
     * For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
     * is the time zone as an offset from UTC.
     */
    updatedTime?: pulumi.Input<string>;
}
