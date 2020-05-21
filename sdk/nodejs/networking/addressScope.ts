// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron addressscope resource within OpenStack.
 *
 * ## Example Usage
 *
 * ### Create an Address-scope
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const addressscope1 = new openstack.networking.AddressScope("addressscope1", {
 *     ipVersion: 6,
 * });
 * ```
 *
 * ### Create a Subnet Pool from an Address-scope
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const addressscope1 = new openstack.networking.AddressScope("addressscope1", {
 *     ipVersion: 6,
 * });
 * const subnetpool1 = new openstack.networking.SubnetPool("subnetpool1", {
 *     addressScopeId: addressscope1.id,
 *     prefixes: [
 *         "fdf7:b13d:dead:beef::/64",
 *         "fd65:86cc:a334:39b7::/64",
 *     ],
 * });
 * ```
 */
export class AddressScope extends pulumi.CustomResource {
    /**
     * Get an existing AddressScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddressScopeState, opts?: pulumi.CustomResourceOptions): AddressScope {
        return new AddressScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/addressScope:AddressScope';

    /**
     * Returns true if the given object is an instance of AddressScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AddressScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AddressScope.__pulumiType;
    }

    /**
     * IP version, either 4 (default) or 6. Changing this
     * creates a new address-scope.
     */
    public readonly ipVersion!: pulumi.Output<number | undefined>;
    /**
     * The name of the address-scope. Changing this updates the
     * name of the existing address-scope.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The owner of the address-scope. Required if admin
     * wants to create a address-scope for another project. Changing this creates a
     * new address-scope.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron address-scope. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * address-scope.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Indicates whether this address-scope is shared across
     * all projects. Changing this updates the shared status of the existing
     * address-scope.
     */
    public readonly shared!: pulumi.Output<boolean>;

    /**
     * Create a AddressScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AddressScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddressScopeArgs | AddressScopeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AddressScopeState | undefined;
            inputs["ipVersion"] = state ? state.ipVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["shared"] = state ? state.shared : undefined;
        } else {
            const args = argsOrState as AddressScopeArgs | undefined;
            inputs["ipVersion"] = args ? args.ipVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["shared"] = args ? args.shared : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AddressScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AddressScope resources.
 */
export interface AddressScopeState {
    /**
     * IP version, either 4 (default) or 6. Changing this
     * creates a new address-scope.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * The name of the address-scope. Changing this updates the
     * name of the existing address-scope.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The owner of the address-scope. Required if admin
     * wants to create a address-scope for another project. Changing this creates a
     * new address-scope.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron address-scope. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * address-scope.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Indicates whether this address-scope is shared across
     * all projects. Changing this updates the shared status of the existing
     * address-scope.
     */
    readonly shared?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a AddressScope resource.
 */
export interface AddressScopeArgs {
    /**
     * IP version, either 4 (default) or 6. Changing this
     * creates a new address-scope.
     */
    readonly ipVersion?: pulumi.Input<number>;
    /**
     * The name of the address-scope. Changing this updates the
     * name of the existing address-scope.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The owner of the address-scope. Required if admin
     * wants to create a address-scope for another project. Changing this creates a
     * new address-scope.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron address-scope. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * address-scope.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Indicates whether this address-scope is shared across
     * all projects. Changing this updates the shared status of the existing
     * address-scope.
     */
    readonly shared?: pulumi.Input<boolean>;
}
