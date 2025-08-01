// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 neutron address group resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const group1 = new openstack.networking.AddressGroupV2("group_1", {
 *     name: "group_1",
 *     description: "My neutron address group",
 *     addresses: [
 *         "192.168.0.1/32",
 *         "2001:db8::1/128",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Address Groups can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:networking/addressGroupV2:AddressGroupV2 group_1 782fef9c-d03c-400a-9735-2f9af5681cb3
 * ```
 */
export class AddressGroupV2 extends pulumi.CustomResource {
    /**
     * Get an existing AddressGroupV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddressGroupV2State, opts?: pulumi.CustomResourceOptions): AddressGroupV2 {
        return new AddressGroupV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/addressGroupV2:AddressGroupV2';

    /**
     * Returns true if the given object is an instance of AddressGroupV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AddressGroupV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AddressGroupV2.__pulumiType;
    }

    /**
     * A list of CIDR blocks that define the addresses in
     * the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
     */
    public readonly addresses!: pulumi.Output<string[]>;
    /**
     * A description of the address group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A name of the address group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The owner of the address group. Required if admin
     * wants to create a group for a specific project. Changing this creates a new
     * address group.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new address group.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a AddressGroupV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AddressGroupV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddressGroupV2Args | AddressGroupV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AddressGroupV2State | undefined;
            resourceInputs["addresses"] = state ? state.addresses : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as AddressGroupV2Args | undefined;
            if ((!args || args.addresses === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addresses'");
            }
            resourceInputs["addresses"] = args ? args.addresses : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AddressGroupV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AddressGroupV2 resources.
 */
export interface AddressGroupV2State {
    /**
     * A list of CIDR blocks that define the addresses in
     * the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description of the address group.
     */
    description?: pulumi.Input<string>;
    /**
     * A name of the address group.
     */
    name?: pulumi.Input<string>;
    /**
     * The owner of the address group. Required if admin
     * wants to create a group for a specific project. Changing this creates a new
     * address group.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new address group.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AddressGroupV2 resource.
 */
export interface AddressGroupV2Args {
    /**
     * A list of CIDR blocks that define the addresses in
     * the address group. Each address must be a valid IPv4 or IPv6 CIDR block.
     */
    addresses: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description of the address group.
     */
    description?: pulumi.Input<string>;
    /**
     * A name of the address group.
     */
    name?: pulumi.Input<string>;
    /**
     * The owner of the address group. Required if admin
     * wants to create a group for a specific project. Changing this creates a new
     * address group.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new address group.
     */
    region?: pulumi.Input<string>;
}
