// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V3 Service resource within OpenStack Keystone.
 * 
 * > **Note:** This usually requires admin privileges.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const service1 = new openstack.identity.ServiceV3("service1", {
 *     type: "custom",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_service_v3.html.markdown.
 */
export class ServiceV3 extends pulumi.CustomResource {
    /**
     * Get an existing ServiceV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceV3State, opts?: pulumi.CustomResourceOptions): ServiceV3 {
        return new ServiceV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:identity/serviceV3:ServiceV3';

    /**
     * Returns true if the given object is an instance of ServiceV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceV3.__pulumiType;
    }

    /**
     * The service description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The service status. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The service name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The service type.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ServiceV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceV3Args | ServiceV3State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceV3State | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ServiceV3Args | undefined;
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceV3.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceV3 resources.
 */
export interface ServiceV3State {
    /**
     * The service description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The service status. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The service name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The service type.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceV3 resource.
 */
export interface ServiceV3Args {
    /**
     * The service description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The service status. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The service name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The service type.
     */
    readonly type: pulumi.Input<string>;
}
