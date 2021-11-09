// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
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
 *     disableRollback: true,
 *     environmentOpts: {
 *         Bin: "\n",
 *     },
 *     parameters: {
 *         length: 4,
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
 *     timeout: 30,
 * });
 * ```
 *
 * ## Import
 *
 * stacks can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:orchestration/stackV1:StackV1 stack_1 ea257959-eeb1-4c10-8d33-26f0409a755d
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
    public readonly environmentOpts!: pulumi.Output<{[key: string]: any} | undefined>;
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
    public readonly parameters!: pulumi.Output<{[key: string]: any} | undefined>;
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
    public readonly templateOpts!: pulumi.Output<{[key: string]: any}>;
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackV1State | undefined;
            inputs["StackOutputs"] = state ? state.StackOutputs : undefined;
            inputs["capabilities"] = state ? state.capabilities : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disableRollback"] = state ? state.disableRollback : undefined;
            inputs["environmentOpts"] = state ? state.environmentOpts : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationTopics"] = state ? state.notificationTopics : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["statusReason"] = state ? state.statusReason : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["templateDescription"] = state ? state.templateDescription : undefined;
            inputs["templateOpts"] = state ? state.templateOpts : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["updatedTime"] = state ? state.updatedTime : undefined;
        } else {
            const args = argsOrState as StackV1Args | undefined;
            if ((!args || args.templateOpts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateOpts'");
            }
            inputs["StackOutputs"] = args ? args.StackOutputs : undefined;
            inputs["capabilities"] = args ? args.capabilities : undefined;
            inputs["creationTime"] = args ? args.creationTime : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableRollback"] = args ? args.disableRollback : undefined;
            inputs["environmentOpts"] = args ? args.environmentOpts : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationTopics"] = args ? args.notificationTopics : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["statusReason"] = args ? args.statusReason : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateDescription"] = args ? args.templateDescription : undefined;
            inputs["templateOpts"] = args ? args.templateOpts : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["updatedTime"] = args ? args.updatedTime : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StackV1.__pulumiType, name, inputs, opts);
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
    environmentOpts?: pulumi.Input<{[key: string]: any}>;
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
    parameters?: pulumi.Input<{[key: string]: any}>;
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
    templateOpts?: pulumi.Input<{[key: string]: any}>;
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
    environmentOpts?: pulumi.Input<{[key: string]: any}>;
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
    parameters?: pulumi.Input<{[key: string]: any}>;
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
    templateOpts: pulumi.Input<{[key: string]: any}>;
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
