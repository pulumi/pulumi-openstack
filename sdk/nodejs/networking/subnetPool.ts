// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron subnetpool resource within OpenStack.
 *
 * ## Example Usage
 *
 * ### Create a Subnet Pool
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const subnetpool1 = new openstack.networking.SubnetPool("subnetpool_1", {
 *     name: "subnetpool_1",
 *     ipVersion: 6,
 *     prefixes: [
 *         "fdf7:b13d:dead:beef::/64",
 *         "fd65:86cc:a334:39b7::/64",
 *     ],
 * });
 * ```
 *
 * ### Create a Subnet from a Subnet Pool
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     name: "network_1",
 *     adminStateUp: true,
 * });
 * const subnetpool1 = new openstack.networking.SubnetPool("subnetpool_1", {
 *     name: "subnetpool_1",
 *     prefixes: ["10.11.12.0/24"],
 * });
 * const subnet1 = new openstack.networking.Subnet("subnet_1", {
 *     name: "subnet_1",
 *     cidr: "10.11.12.0/25",
 *     networkId: network1.id,
 *     subnetpoolId: subnetpool1.id,
 * });
 * ```
 *
 * ## Import
 *
 * Subnetpools can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:networking/subnetPool:SubnetPool subnetpool_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
 * ```
 */
export class SubnetPool extends pulumi.CustomResource {
    /**
     * Get an existing SubnetPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetPoolState, opts?: pulumi.CustomResourceOptions): SubnetPool {
        return new SubnetPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/subnetPool:SubnetPool';

    /**
     * Returns true if the given object is an instance of SubnetPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubnetPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubnetPool.__pulumiType;
    }

    /**
     * The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     */
    public readonly addressScopeId!: pulumi.Output<string | undefined>;
    /**
     * The collection of tags assigned on the subnetpool, which have been
     * explicitly and implicitly added.
     */
    public /*out*/ readonly allTags!: pulumi.Output<string[]>;
    /**
     * The time at which subnetpool was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     */
    public readonly defaultPrefixlen!: pulumi.Output<number>;
    /**
     * The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     */
    public readonly defaultQuota!: pulumi.Output<number | undefined>;
    /**
     * The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IP protocol version.
     */
    public readonly ipVersion!: pulumi.Output<number>;
    /**
     * Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     */
    public readonly isDefault!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     */
    public readonly maxPrefixlen!: pulumi.Output<number>;
    /**
     * The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     */
    public readonly minPrefixlen!: pulumi.Output<number>;
    /**
     * The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     */
    public readonly prefixes!: pulumi.Output<string[]>;
    /**
     * The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The revision number of the subnetpool.
     */
    public /*out*/ readonly revisionNumber!: pulumi.Output<number>;
    /**
     * Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     */
    public readonly shared!: pulumi.Output<boolean | undefined>;
    /**
     * A set of string tags for the subnetpool.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The time at which subnetpool was created.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a SubnetPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetPoolArgs | SubnetPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubnetPoolState | undefined;
            resourceInputs["addressScopeId"] = state ? state.addressScopeId : undefined;
            resourceInputs["allTags"] = state ? state.allTags : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["defaultPrefixlen"] = state ? state.defaultPrefixlen : undefined;
            resourceInputs["defaultQuota"] = state ? state.defaultQuota : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipVersion"] = state ? state.ipVersion : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["maxPrefixlen"] = state ? state.maxPrefixlen : undefined;
            resourceInputs["minPrefixlen"] = state ? state.minPrefixlen : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["prefixes"] = state ? state.prefixes : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["revisionNumber"] = state ? state.revisionNumber : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as SubnetPoolArgs | undefined;
            if ((!args || args.prefixes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefixes'");
            }
            resourceInputs["addressScopeId"] = args ? args.addressScopeId : undefined;
            resourceInputs["defaultPrefixlen"] = args ? args.defaultPrefixlen : undefined;
            resourceInputs["defaultQuota"] = args ? args.defaultQuota : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipVersion"] = args ? args.ipVersion : undefined;
            resourceInputs["isDefault"] = args ? args.isDefault : undefined;
            resourceInputs["maxPrefixlen"] = args ? args.maxPrefixlen : undefined;
            resourceInputs["minPrefixlen"] = args ? args.minPrefixlen : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["prefixes"] = args ? args.prefixes : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["allTags"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["revisionNumber"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SubnetPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubnetPool resources.
 */
export interface SubnetPoolState {
    /**
     * The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     */
    addressScopeId?: pulumi.Input<string>;
    /**
     * The collection of tags assigned on the subnetpool, which have been
     * explicitly and implicitly added.
     */
    allTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time at which subnetpool was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     */
    defaultPrefixlen?: pulumi.Input<number>;
    /**
     * The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     */
    defaultQuota?: pulumi.Input<number>;
    /**
     * The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP protocol version.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     */
    maxPrefixlen?: pulumi.Input<number>;
    /**
     * The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     */
    minPrefixlen?: pulumi.Input<number>;
    /**
     * The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     */
    prefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     */
    region?: pulumi.Input<string>;
    /**
     * The revision number of the subnetpool.
     */
    revisionNumber?: pulumi.Input<number>;
    /**
     * Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * A set of string tags for the subnetpool.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time at which subnetpool was created.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a SubnetPool resource.
 */
export interface SubnetPoolArgs {
    /**
     * The Neutron address scope to assign to the
     * subnetpool. Changing this updates the address scope id of the existing
     * subnetpool.
     */
    addressScopeId?: pulumi.Input<string>;
    /**
     * The size of the prefix to allocate when the cidr
     * or prefixlen attributes are omitted when you create the subnet. Defaults to the
     * MinPrefixLen. Changing this updates the default prefixlen of the existing
     * subnetpool.
     */
    defaultPrefixlen?: pulumi.Input<number>;
    /**
     * The per-project quota on the prefix space that can be
     * allocated from the subnetpool for project subnets. Changing this updates the
     * default quota of the existing subnetpool.
     */
    defaultQuota?: pulumi.Input<number>;
    /**
     * The human-readable description for the subnetpool.
     * Changing this updates the description of the existing subnetpool.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP protocol version.
     */
    ipVersion?: pulumi.Input<number>;
    /**
     * Indicates whether the subnetpool is default
     * subnetpool or not. Changing this updates the default status of the existing
     * subnetpool.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The maximum prefix size that can be allocated from
     * the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
     * default is 128. Changing this updates the max prefixlen of the existing
     * subnetpool.
     */
    maxPrefixlen?: pulumi.Input<number>;
    /**
     * The smallest prefix that can be allocated from a
     * subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
     * is 64. Changing this updates the min prefixlen of the existing subnetpool.
     */
    minPrefixlen?: pulumi.Input<number>;
    /**
     * The name of the subnetpool. Changing this updates the name of
     * the existing subnetpool.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of subnet prefixes to assign to the subnetpool.
     * Neutron API merges adjacent prefixes and treats them as a single prefix. Each
     * subnet prefix must be unique among all subnet prefixes in all subnetpools that
     * are associated with the address scope. Changing this updates the prefixes list
     * of the existing subnetpool.
     */
    prefixes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the subnetpool. Required if admin wants to
     * create a subnetpool for another project. Changing this creates a new subnetpool.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron subnetpool. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * subnetpool.
     */
    region?: pulumi.Input<string>;
    /**
     * Indicates whether this subnetpool is shared across
     * all projects. Changing this updates the shared status of the existing
     * subnetpool.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * A set of string tags for the subnetpool.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
