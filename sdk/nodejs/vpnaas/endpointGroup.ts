// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron Endpoint Group resource within OpenStack.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const group1 = new openstack.vpnaas.EndpointGroup("group_1", {
 *     name: "Group 1",
 *     type: "cidr",
 *     endpoints: [
 *         "10.2.0.0/24",
 *         "10.3.0.0/24",
 *     ],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Groups can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:vpnaas/endpointGroup:EndpointGroup group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 */
export class EndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing EndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointGroupState, opts?: pulumi.CustomResourceOptions): EndpointGroup {
        return new EndpointGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:vpnaas/endpointGroup:EndpointGroup';

    /**
     * Returns true if the given object is an instance of EndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointGroup.__pulumiType;
    }

    /**
     * The human-readable description for the group.
     * Changing this updates the description of the existing group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     */
    public readonly endpoints!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the group. Changing this updates the name of
     * the existing group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a EndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointGroupArgs | EndpointGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as EndpointGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpoints"] = args ? args.endpoints : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndpointGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointGroup resources.
 */
export interface EndpointGroupState {
    /**
     * The human-readable description for the group.
     * Changing this updates the description of the existing group.
     */
    description?: pulumi.Input<string>;
    /**
     * List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     */
    endpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the group. Changing this updates the name of
     * the existing group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     */
    region?: pulumi.Input<string>;
    /**
     * The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     */
    type?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EndpointGroup resource.
 */
export interface EndpointGroupArgs {
    /**
     * The human-readable description for the group.
     * Changing this updates the description of the existing group.
     */
    description?: pulumi.Input<string>;
    /**
     * List of endpoints of the same type, for the endpoint group. The values will depend on the type.
     * Changing this creates a new group.
     */
    endpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the group. Changing this updates the name of
     * the existing group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an endpoint group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * group.
     */
    region?: pulumi.Input<string>;
    /**
     * The owner of the group. Required if admin wants to
     * create an endpoint group for another project. Changing this creates a new group.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan.
     * Changing this creates a new group.
     */
    type?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: any}>;
}
