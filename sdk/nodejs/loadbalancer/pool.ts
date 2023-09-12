// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 pool resource within OpenStack.
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
 * const pool1 = new openstack.loadbalancer.Pool("pool1", {
 *     lbMethod: "ROUND_ROBIN",
 *     listenerId: "d9415786-5f1a-428b-b35f-2f1523e146d2",
 *     persistence: {
 *         cookieName: "testCookie",
 *         type: "APP_COOKIE",
 *     },
 *     protocol: "HTTP",
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancer Pool can be imported using the Pool ID, e.g.:
 *
 * ```sh
 *  $ pulumi import openstack:loadbalancer/pool:Pool pool_1 60ad9ee4-249a-4d60-a45b-aa60e046c513
 * ```
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/pool:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    /**
     * The administrative state of the pool.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * Human-readable description for the pool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The load balancing algorithm to
     * distribute traffic to the pool's members. Must be one of
     * ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
     * in Octavia).
     */
    public readonly lbMethod!: pulumi.Output<string>;
    /**
     * The Listener on which the members of the pool
     * will be associated with. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    public readonly listenerId!: pulumi.Output<string | undefined>;
    /**
     * The load balancer on which to provision this
     * pool. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    public readonly loadbalancerId!: pulumi.Output<string | undefined>;
    /**
     * Human-readable name for the pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Omit this field to prevent session persistence.  Indicates
     * whether connections in the same session will be processed by the same Pool
     * member or not. Changing this creates a new pool.
     */
    public readonly persistence!: pulumi.Output<outputs.loadbalancer.PoolPersistence>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS, PROXY,
     * UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
     * or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * pool.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the pool.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new pool.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PoolArgs | PoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PoolState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lbMethod"] = state ? state.lbMethod : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["loadbalancerId"] = state ? state.loadbalancerId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["persistence"] = state ? state.persistence : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            if ((!args || args.lbMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lbMethod'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["lbMethod"] = args ? args.lbMethod : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["loadbalancerId"] = args ? args.loadbalancerId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["persistence"] = args ? args.persistence : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    /**
     * The administrative state of the pool.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the pool.
     */
    description?: pulumi.Input<string>;
    /**
     * The load balancing algorithm to
     * distribute traffic to the pool's members. Must be one of
     * ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
     * in Octavia).
     */
    lbMethod?: pulumi.Input<string>;
    /**
     * The Listener on which the members of the pool
     * will be associated with. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * pool. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the pool.
     */
    name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence.  Indicates
     * whether connections in the same session will be processed by the same Pool
     * member or not. Changing this creates a new pool.
     */
    persistence?: pulumi.Input<inputs.loadbalancer.PoolPersistence>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS, PROXY,
     * UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
     * or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * pool.
     */
    region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the pool.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new pool.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * The administrative state of the pool.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the pool.
     */
    description?: pulumi.Input<string>;
    /**
     * The load balancing algorithm to
     * distribute traffic to the pool's members. Must be one of
     * ROUND_ROBIN, LEAST_CONNECTIONS, SOURCE_IP, or SOURCE_IP_PORT (supported only
     * in Octavia).
     */
    lbMethod: pulumi.Input<string>;
    /**
     * The Listener on which the members of the pool
     * will be associated with. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * pool. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the pool.
     */
    name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence.  Indicates
     * whether connections in the same session will be processed by the same Pool
     * member or not. Changing this creates a new pool.
     */
    persistence?: pulumi.Input<inputs.loadbalancer.PoolPersistence>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS, PROXY,
     * UDP (supported only in Octavia), PROXYV2 (**Octavia minor version >= 2.22**)
     * or SCTP (**Octavia minor version >= 2.23**). Changing this creates a new pool.
     */
    protocol: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * pool.
     */
    region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the pool.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new pool.
     */
    tenantId?: pulumi.Input<string>;
}
