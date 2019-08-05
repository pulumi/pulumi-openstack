// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 pool resource within OpenStack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 * 
 * const pool1 = new openstack.loadbalancer.Pool("pool_1", {
 *     lbMethod: "ROUND_ROBIN",
 *     listenerId: "d9415786-5f1a-428b-b35f-2f1523e146d2",
 *     persistences: [{
 *         cookieName: "testCookie",
 *         type: "HTTP_COOKIE",
 *     }],
 *     protocol: "HTTP",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_pool_v2.html.markdown.
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
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
    public readonly persistences!: pulumi.Output<{ cookieName?: string, type: string }[] | undefined>;
    /**
     * See Argument Reference above.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PoolState | undefined;
            inputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["lbMethod"] = state ? state.lbMethod : undefined;
            inputs["listenerId"] = state ? state.listenerId : undefined;
            inputs["loadbalancerId"] = state ? state.loadbalancerId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["persistences"] = state ? state.persistences : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            if (!args || args.lbMethod === undefined) {
                throw new Error("Missing required property 'lbMethod'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            inputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["lbMethod"] = args ? args.lbMethod : undefined;
            inputs["listenerId"] = args ? args.listenerId : undefined;
            inputs["loadbalancerId"] = args ? args.loadbalancerId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["persistences"] = args ? args.persistences : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Pool.__pulumiType, name, inputs, opts);
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
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the pool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The load balancing algorithm to
     * distribute traffic to the pool's members. Must be one of
     * ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
     */
    readonly lbMethod?: pulumi.Input<string>;
    /**
     * The Listener on which the members of the pool
     * will be associated with. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    readonly listenerId?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * pool. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    readonly loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence.  Indicates
     * whether connections in the same session will be processed by the same Pool
     * member or not. Changing this creates a new pool.
     */
    readonly persistences?: pulumi.Input<pulumi.Input<{ cookieName?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * See Argument Reference above.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * pool.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the pool.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new pool.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * The administrative state of the pool.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * Human-readable description for the pool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The load balancing algorithm to
     * distribute traffic to the pool's members. Must be one of
     * ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.
     */
    readonly lbMethod: pulumi.Input<string>;
    /**
     * The Listener on which the members of the pool
     * will be associated with. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    readonly listenerId?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * pool. Changing this creates a new pool.
     * Note:  One of LoadbalancerID or ListenerID must be provided.
     */
    readonly loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Omit this field to prevent session persistence.  Indicates
     * whether connections in the same session will be processed by the same Pool
     * member or not. Changing this creates a new pool.
     */
    readonly persistences?: pulumi.Input<pulumi.Input<{ cookieName?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * See Argument Reference above.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * pool.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the pool.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new pool.
     */
    readonly tenantId?: pulumi.Input<string>;
}
