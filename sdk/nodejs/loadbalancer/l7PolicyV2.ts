// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Load Balancer L7 Policy resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const network1 = new openstack.networking.Network("network_1", {
 *     adminStateUp: true,
 * });
 * const subnet1 = new openstack.networking.Subnet("subnet_1", {
 *     cidr: "192.168.199.0/24",
 *     ipVersion: 4,
 *     networkId: network1.id,
 * });
 * const loadbalancer1 = new openstack.loadbalancer.LoadBalancer("loadbalancer_1", {
 *     vipSubnetId: subnet1.id,
 * });
 * const listener1 = new openstack.loadbalancer.Listener("listener_1", {
 *     loadbalancerId: loadbalancer1.id,
 *     protocol: "HTTP",
 *     protocolPort: 8080,
 * });
 * const pool1 = new openstack.loadbalancer.Pool("pool_1", {
 *     lbMethod: "ROUND_ROBIN",
 *     loadbalancerId: loadbalancer1.id,
 *     protocol: "HTTP",
 * });
 * const l7policy1 = new openstack.loadbalancer.L7PolicyV2("l7policy_1", {
 *     action: "REDIRECT_TO_POOL",
 *     description: "test l7 policy",
 *     listenerId: listener1.id,
 *     position: 1,
 *     redirectPoolId: pool1.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_l7policy_v2.html.markdown.
 */
export class L7PolicyV2 extends pulumi.CustomResource {
    /**
     * Get an existing L7PolicyV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: L7PolicyV2State, opts?: pulumi.CustomResourceOptions): L7PolicyV2 {
        return new L7PolicyV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/l7PolicyV2:L7PolicyV2';

    /**
     * Returns true if the given object is an instance of L7PolicyV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L7PolicyV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L7PolicyV2.__pulumiType;
    }

    /**
     * The L7 Policy action - can either be REDIRECT\_TO\_POOL,
     * REDIRECT\_TO\_URL or REJECT.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The administrative state of the L7 Policy.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Human-readable description for the L7 Policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Listener on which the L7 Policy will be associated with.
     * Changing this creates a new L7 Policy.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * Human-readable name for the L7 Policy. Does not have
     * to be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The position of this policy on the listener. Positions start at 1.
     */
    public readonly position!: pulumi.Output<number>;
    /**
     * Requests matching this policy will be redirected to the
     * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
     */
    public readonly redirectPoolId!: pulumi.Output<string | undefined>;
    /**
     * Requests matching this policy will be redirected to this URL.
     * Only valid if action is REDIRECT\_TO\_URL.
     */
    public readonly redirectUrl!: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Policy.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Policy.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Policy.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a L7PolicyV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L7PolicyV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: L7PolicyV2Args | L7PolicyV2State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as L7PolicyV2State | undefined;
            inputs["action"] = state ? state.action : undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["listenerId"] = state ? state.listenerId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["position"] = state ? state.position : undefined;
            inputs["redirectPoolId"] = state ? state.redirectPoolId : undefined;
            inputs["redirectUrl"] = state ? state.redirectUrl : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as L7PolicyV2Args | undefined;
            if (!args || args.action === undefined) {
                throw new Error("Missing required property 'action'");
            }
            if (!args || args.listenerId === undefined) {
                throw new Error("Missing required property 'listenerId'");
            }
            inputs["action"] = args ? args.action : undefined;
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["listenerId"] = args ? args.listenerId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["position"] = args ? args.position : undefined;
            inputs["redirectPoolId"] = args ? args.redirectPoolId : undefined;
            inputs["redirectUrl"] = args ? args.redirectUrl : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        super(L7PolicyV2.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering L7PolicyV2 resources.
 */
export interface L7PolicyV2State {
    /**
     * The L7 Policy action - can either be REDIRECT\_TO\_POOL,
     * REDIRECT\_TO\_URL or REJECT.
     */
    readonly action?: pulumi.Input<string>;
    /**
     * The administrative state of the L7 Policy.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the L7 Policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Listener on which the L7 Policy will be associated with.
     * Changing this creates a new L7 Policy.
     */
    readonly listenerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the L7 Policy. Does not have
     * to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The position of this policy on the listener. Positions start at 1.
     */
    readonly position?: pulumi.Input<number>;
    /**
     * Requests matching this policy will be redirected to the
     * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
     */
    readonly redirectPoolId?: pulumi.Input<string>;
    /**
     * Requests matching this policy will be redirected to this URL.
     * Only valid if action is REDIRECT\_TO\_URL.
     */
    readonly redirectUrl?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Policy.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Policy.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Policy.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a L7PolicyV2 resource.
 */
export interface L7PolicyV2Args {
    /**
     * The L7 Policy action - can either be REDIRECT\_TO\_POOL,
     * REDIRECT\_TO\_URL or REJECT.
     */
    readonly action: pulumi.Input<string>;
    /**
     * The administrative state of the L7 Policy.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the L7 Policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The Listener on which the L7 Policy will be associated with.
     * Changing this creates a new L7 Policy.
     */
    readonly listenerId: pulumi.Input<string>;
    /**
     * Human-readable name for the L7 Policy. Does not have
     * to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The position of this policy on the listener. Positions start at 1.
     */
    readonly position?: pulumi.Input<number>;
    /**
     * Requests matching this policy will be redirected to the
     * pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.
     */
    readonly redirectPoolId?: pulumi.Input<string>;
    /**
     * Requests matching this policy will be redirected to this URL.
     * Only valid if action is REDIRECT\_TO\_URL.
     */
    readonly redirectUrl?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Policy.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Policy.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Policy.
     */
    readonly tenantId?: pulumi.Input<string>;
}
