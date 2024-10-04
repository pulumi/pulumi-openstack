// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 L7 Rule resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const network1 = new openstack.networking.Network("network_1", {
 *     name: "network_1",
 *     adminStateUp: true,
 * });
 * const subnet1 = new openstack.networking.Subnet("subnet_1", {
 *     name: "subnet_1",
 *     cidr: "192.168.199.0/24",
 *     ipVersion: 4,
 *     networkId: network1.id,
 * });
 * const loadbalancer1 = new openstack.loadbalancer.LoadBalancer("loadbalancer_1", {
 *     name: "loadbalancer_1",
 *     vipSubnetId: subnet1.id,
 * });
 * const listener1 = new openstack.loadbalancer.Listener("listener_1", {
 *     name: "listener_1",
 *     protocol: "HTTP",
 *     protocolPort: 8080,
 *     loadbalancerId: loadbalancer1.id,
 * });
 * const pool1 = new openstack.loadbalancer.Pool("pool_1", {
 *     name: "pool_1",
 *     protocol: "HTTP",
 *     lbMethod: "ROUND_ROBIN",
 *     loadbalancerId: loadbalancer1.id,
 * });
 * const l7policy1 = new openstack.loadbalancer.L7PolicyV2("l7policy_1", {
 *     name: "test",
 *     action: "REDIRECT_TO_URL",
 *     description: "test description",
 *     position: 1,
 *     listenerId: listener1.id,
 *     redirectUrl: "http://www.example.com",
 * });
 * const l7rule1 = new openstack.loadbalancer.L7RuleV2("l7rule_1", {
 *     l7policyId: l7policy1.id,
 *     type: "PATH",
 *     compareType: "EQUAL_TO",
 *     value: "/api",
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID
 * separated by a slash, e.g.:
 *
 * ```sh
 * $ pulumi import openstack:loadbalancer/l7RuleV2:L7RuleV2 l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
 * ```
 */
export class L7RuleV2 extends pulumi.CustomResource {
    /**
     * Get an existing L7RuleV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: L7RuleV2State, opts?: pulumi.CustomResourceOptions): L7RuleV2 {
        return new L7RuleV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/l7RuleV2:L7RuleV2';

    /**
     * Returns true if the given object is an instance of L7RuleV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L7RuleV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L7RuleV2.__pulumiType;
    }

    /**
     * The administrative state of the L7 Rule.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * The comparison type for the L7 rule - can either be
     * CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
     */
    public readonly compareType!: pulumi.Output<string>;
    /**
     * When true the logic of the rule is inverted. For example, with invert
     * true, equal to would become not equal to. Default is false.
     */
    public readonly invert!: pulumi.Output<boolean | undefined>;
    /**
     * The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * The ID of the L7 Policy to query. Changing this creates a new
     * L7 Rule.
     */
    public readonly l7policyId!: pulumi.Output<string>;
    /**
     * The ID of the Listener owning this resource.
     */
    public /*out*/ readonly listenerId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an L7 rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Rule.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Rule.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Rule.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
     * HOST\_NAME, PATH, SSL\_CONN\_HAS\_CERT, SSL\_VERIFY\_RESULT or SSL\_DN\_FIELD.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The value to use for the comparison. For example, the file type to
     * compare.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a L7RuleV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L7RuleV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: L7RuleV2Args | L7RuleV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as L7RuleV2State | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["compareType"] = state ? state.compareType : undefined;
            resourceInputs["invert"] = state ? state.invert : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["l7policyId"] = state ? state.l7policyId : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as L7RuleV2Args | undefined;
            if ((!args || args.compareType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compareType'");
            }
            if ((!args || args.l7policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'l7policyId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["compareType"] = args ? args.compareType : undefined;
            resourceInputs["invert"] = args ? args.invert : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["l7policyId"] = args ? args.l7policyId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["listenerId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(L7RuleV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering L7RuleV2 resources.
 */
export interface L7RuleV2State {
    /**
     * The administrative state of the L7 Rule.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The comparison type for the L7 rule - can either be
     * CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
     */
    compareType?: pulumi.Input<string>;
    /**
     * When true the logic of the rule is inverted. For example, with invert
     * true, equal to would become not equal to. Default is false.
     */
    invert?: pulumi.Input<boolean>;
    /**
     * The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the L7 Policy to query. Changing this creates a new
     * L7 Rule.
     */
    l7policyId?: pulumi.Input<string>;
    /**
     * The ID of the Listener owning this resource.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an L7 rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Rule.
     */
    region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Rule.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Rule.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
     * HOST\_NAME, PATH, SSL\_CONN\_HAS\_CERT, SSL\_VERIFY\_RESULT or SSL\_DN\_FIELD.
     */
    type?: pulumi.Input<string>;
    /**
     * The value to use for the comparison. For example, the file type to
     * compare.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a L7RuleV2 resource.
 */
export interface L7RuleV2Args {
    /**
     * The administrative state of the L7 Rule.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The comparison type for the L7 rule - can either be
     * CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX
     */
    compareType: pulumi.Input<string>;
    /**
     * When true the logic of the rule is inverted. For example, with invert
     * true, equal to would become not equal to. Default is false.
     */
    invert?: pulumi.Input<boolean>;
    /**
     * The key to use for the comparison. For example, the name of the cookie to
     * evaluate. Valid when `type` is set to COOKIE or HEADER.
     */
    key?: pulumi.Input<string>;
    /**
     * The ID of the L7 Policy to query. Changing this creates a new
     * L7 Rule.
     */
    l7policyId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an L7 rule. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * L7 Rule.
     */
    region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the L7 Rule.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new L7 Rule.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER,
     * HOST\_NAME, PATH, SSL\_CONN\_HAS\_CERT, SSL\_VERIFY\_RESULT or SSL\_DN\_FIELD.
     */
    type: pulumi.Input<string>;
    /**
     * The value to use for the comparison. For example, the file type to
     * compare.
     */
    value: pulumi.Input<string>;
}
