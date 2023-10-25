// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 loadbalancer resource within OpenStack.
 *
 * > **Note:** This resource has attributes that depend on octavia minor versions.
 * Please ensure your Openstack cloud supports the required minor version.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const lb1 = new openstack.loadbalancer.LoadBalancer("lb1", {vipSubnetId: "d9415786-5f1a-428b-b35f-2f1523e146d2"});
 * ```
 *
 * ## Import
 *
 * Load Balancer can be imported using the Load Balancer ID, e.g.:
 *
 * ```sh
 *  $ pulumi import openstack:loadbalancer/loadBalancer:LoadBalancer loadbalancer_1 19bcfdc7-c521-4a7e-9459-6750bd16df76
 * ```
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * The administrative state of the Loadbalancer.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * The availability zone of the Loadbalancer.
     * Changing this creates a new loadbalancer. Available only for Octavia
     * **minor version 2.14 or later**.
     */
    public readonly availabilityZone!: pulumi.Output<string | undefined>;
    /**
     * Human-readable description for the Loadbalancer.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The UUID of a flavor. Changing this creates a new
     * loadbalancer.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * The name of the provider. Changing this
     * creates a new loadbalancer.
     */
    public readonly loadbalancerProvider!: pulumi.Output<string>;
    /**
     * Human-readable name for the Loadbalancer. Does not have
     * to be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A list of security group IDs to apply to the
     * loadbalancer. The security groups must be specified by ID and not name (as
     * opposed to how they are configured with the Compute Instance).
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * A list of simple strings assigned to the loadbalancer.
     * Available only for Octavia **minor version 2.5 or later**.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Loadbalancer.  Only administrative users can specify a tenant UUID
     * other than their own.  Changing this creates a new loadbalancer.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The ip address of the load balancer.
     * Changing this creates a new loadbalancer.
     */
    public readonly vipAddress!: pulumi.Output<string>;
    /**
     * The network on which to allocate the
     * Loadbalancer's address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is available only for Octavia.
     */
    public readonly vipNetworkId!: pulumi.Output<string>;
    /**
     * The port UUID that the loadbalancer will use.
     * Changing this creates a new loadbalancer. It is available only for Octavia.
     */
    public readonly vipPortId!: pulumi.Output<string>;
    /**
     * The subnet on which to allocate the
     * Loadbalancer's address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is required to Neutron LBaaS but optional for Octavia.
     */
    public readonly vipSubnetId!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["loadbalancerProvider"] = state ? state.loadbalancerProvider : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["vipAddress"] = state ? state.vipAddress : undefined;
            resourceInputs["vipNetworkId"] = state ? state.vipNetworkId : undefined;
            resourceInputs["vipPortId"] = state ? state.vipPortId : undefined;
            resourceInputs["vipSubnetId"] = state ? state.vipSubnetId : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["loadbalancerProvider"] = args ? args.loadbalancerProvider : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["vipAddress"] = args ? args.vipAddress : undefined;
            resourceInputs["vipNetworkId"] = args ? args.vipNetworkId : undefined;
            resourceInputs["vipPortId"] = args ? args.vipPortId : undefined;
            resourceInputs["vipSubnetId"] = args ? args.vipSubnetId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoadBalancer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * The administrative state of the Loadbalancer.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The availability zone of the Loadbalancer.
     * Changing this creates a new loadbalancer. Available only for Octavia
     * **minor version 2.14 or later**.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Human-readable description for the Loadbalancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The UUID of a flavor. Changing this creates a new
     * loadbalancer.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The name of the provider. Changing this
     * creates a new loadbalancer.
     */
    loadbalancerProvider?: pulumi.Input<string>;
    /**
     * Human-readable name for the Loadbalancer. Does not have
     * to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of security group IDs to apply to the
     * loadbalancer. The security groups must be specified by ID and not name (as
     * opposed to how they are configured with the Compute Instance).
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of simple strings assigned to the loadbalancer.
     * Available only for Octavia **minor version 2.5 or later**.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Loadbalancer.  Only administrative users can specify a tenant UUID
     * other than their own.  Changing this creates a new loadbalancer.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The ip address of the load balancer.
     * Changing this creates a new loadbalancer.
     */
    vipAddress?: pulumi.Input<string>;
    /**
     * The network on which to allocate the
     * Loadbalancer's address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is available only for Octavia.
     */
    vipNetworkId?: pulumi.Input<string>;
    /**
     * The port UUID that the loadbalancer will use.
     * Changing this creates a new loadbalancer. It is available only for Octavia.
     */
    vipPortId?: pulumi.Input<string>;
    /**
     * The subnet on which to allocate the
     * Loadbalancer's address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is required to Neutron LBaaS but optional for Octavia.
     */
    vipSubnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * The administrative state of the Loadbalancer.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The availability zone of the Loadbalancer.
     * Changing this creates a new loadbalancer. Available only for Octavia
     * **minor version 2.14 or later**.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Human-readable description for the Loadbalancer.
     */
    description?: pulumi.Input<string>;
    /**
     * The UUID of a flavor. Changing this creates a new
     * loadbalancer.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The name of the provider. Changing this
     * creates a new loadbalancer.
     */
    loadbalancerProvider?: pulumi.Input<string>;
    /**
     * Human-readable name for the Loadbalancer. Does not have
     * to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an LB member. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * LB member.
     */
    region?: pulumi.Input<string>;
    /**
     * A list of security group IDs to apply to the
     * loadbalancer. The security groups must be specified by ID and not name (as
     * opposed to how they are configured with the Compute Instance).
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of simple strings assigned to the loadbalancer.
     * Available only for Octavia **minor version 2.5 or later**.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Loadbalancer.  Only administrative users can specify a tenant UUID
     * other than their own.  Changing this creates a new loadbalancer.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The ip address of the load balancer.
     * Changing this creates a new loadbalancer.
     */
    vipAddress?: pulumi.Input<string>;
    /**
     * The network on which to allocate the
     * Loadbalancer's address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is available only for Octavia.
     */
    vipNetworkId?: pulumi.Input<string>;
    /**
     * The port UUID that the loadbalancer will use.
     * Changing this creates a new loadbalancer. It is available only for Octavia.
     */
    vipPortId?: pulumi.Input<string>;
    /**
     * The subnet on which to allocate the
     * Loadbalancer's address. A tenant can only create Loadbalancers on networks
     * authorized by policy (e.g. networks that belong to them or networks that
     * are shared).  Changing this creates a new loadbalancer.
     * It is required to Neutron LBaaS but optional for Octavia.
     */
    vipSubnetId?: pulumi.Input<string>;
}
