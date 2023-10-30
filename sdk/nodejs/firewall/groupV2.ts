// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a v2 firewall group resource within OpenStack.
 *
 * > **Note:** Firewall v2 has no support for OVN currently.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const rule1 = new openstack.firewall.RuleV2("rule1", {
 *     description: "drop TELNET traffic",
 *     action: "deny",
 *     protocol: "tcp",
 *     destinationPort: "23",
 *     enabled: true,
 * });
 * const rule2 = new openstack.firewall.RuleV2("rule2", {
 *     description: "drop NTP traffic",
 *     action: "deny",
 *     protocol: "udp",
 *     destinationPort: "123",
 *     enabled: false,
 * });
 * const policy1 = new openstack.firewall.PolicyV2("policy1", {rules: [rule1.id]});
 * const policy2 = new openstack.firewall.PolicyV2("policy2", {rules: [rule2.id]});
 * const group1 = new openstack.firewall.GroupV2("group1", {
 *     ingressFirewallPolicyId: policy1.id,
 *     egressFirewallPolicyId: policy2.id,
 * });
 * ```
 *
 * ## Import
 *
 * Firewall groups can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:firewall/groupV2:GroupV2 group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97
 * ```
 */
export class GroupV2 extends pulumi.CustomResource {
    /**
     * Get an existing GroupV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupV2State, opts?: pulumi.CustomResourceOptions): GroupV2 {
        return new GroupV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:firewall/groupV2:GroupV2';

    /**
     * Returns true if the given object is an instance of GroupV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupV2.__pulumiType;
    }

    /**
     * Administrative up/down status for the firewall
     * group (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `adminStateUp` of an existing firewall group.
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * A description for the firewall group. Changing this
     * updates the `description` of an existing firewall group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The egress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `egressFirewallPolicyId` of an existing firewall group.
     */
    public readonly egressFirewallPolicyId!: pulumi.Output<string | undefined>;
    /**
     * The ingress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `ingressFirewallPolicyId` of an existing firewall group.
     */
    public readonly ingressFirewallPolicyId!: pulumi.Output<string | undefined>;
    /**
     * A name for the firewall group. Changing this
     * updates the `name` of an existing firewall.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Port(s) to associate this firewall group
     * with. Must be a list of strings. Changing this updates the associated ports
     * of an existing firewall group.
     */
    public readonly ports!: pulumi.Output<string[] | undefined>;
    /**
     * This argument conflicts and  is interchangeable
     * with `tenantId`. The owner of the firewall group. Required if admin wants
     * to create a firewall group for another project. Changing this creates a new
     * firewall group.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall group.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Sharing status of the firewall group (must be "true"
     * or "false" if provided). If this is "true" the firewall group is visible to,
     * and can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall group. Only administrative users
     * can specify if the firewall group should be shared.
     */
    public readonly shared!: pulumi.Output<boolean | undefined>;
    /**
     * The status of the firewall group.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * This argument conflicts and is interchangeable with
     * `projectId`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another tenant. Changing this creates a new
     * firewall group.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a GroupV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupV2Args | GroupV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupV2State | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["egressFirewallPolicyId"] = state ? state.egressFirewallPolicyId : undefined;
            resourceInputs["ingressFirewallPolicyId"] = state ? state.ingressFirewallPolicyId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as GroupV2Args | undefined;
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["egressFirewallPolicyId"] = args ? args.egressFirewallPolicyId : undefined;
            resourceInputs["ingressFirewallPolicyId"] = args ? args.ingressFirewallPolicyId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupV2 resources.
 */
export interface GroupV2State {
    /**
     * Administrative up/down status for the firewall
     * group (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `adminStateUp` of an existing firewall group.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * A description for the firewall group. Changing this
     * updates the `description` of an existing firewall group.
     */
    description?: pulumi.Input<string>;
    /**
     * The egress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `egressFirewallPolicyId` of an existing firewall group.
     */
    egressFirewallPolicyId?: pulumi.Input<string>;
    /**
     * The ingress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `ingressFirewallPolicyId` of an existing firewall group.
     */
    ingressFirewallPolicyId?: pulumi.Input<string>;
    /**
     * A name for the firewall group. Changing this
     * updates the `name` of an existing firewall.
     */
    name?: pulumi.Input<string>;
    /**
     * Port(s) to associate this firewall group
     * with. Must be a list of strings. Changing this updates the associated ports
     * of an existing firewall group.
     */
    ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This argument conflicts and  is interchangeable
     * with `tenantId`. The owner of the firewall group. Required if admin wants
     * to create a firewall group for another project. Changing this creates a new
     * firewall group.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall group.
     */
    region?: pulumi.Input<string>;
    /**
     * Sharing status of the firewall group (must be "true"
     * or "false" if provided). If this is "true" the firewall group is visible to,
     * and can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall group. Only administrative users
     * can specify if the firewall group should be shared.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * The status of the firewall group.
     */
    status?: pulumi.Input<string>;
    /**
     * This argument conflicts and is interchangeable with
     * `projectId`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another tenant. Changing this creates a new
     * firewall group.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupV2 resource.
 */
export interface GroupV2Args {
    /**
     * Administrative up/down status for the firewall
     * group (must be "true" or "false" if provided - defaults to "true").
     * Changing this updates the `adminStateUp` of an existing firewall group.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * A description for the firewall group. Changing this
     * updates the `description` of an existing firewall group.
     */
    description?: pulumi.Input<string>;
    /**
     * The egress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `egressFirewallPolicyId` of an existing firewall group.
     */
    egressFirewallPolicyId?: pulumi.Input<string>;
    /**
     * The ingress firewall policy resource
     * id for the firewall group. Changing this updates the
     * `ingressFirewallPolicyId` of an existing firewall group.
     */
    ingressFirewallPolicyId?: pulumi.Input<string>;
    /**
     * A name for the firewall group. Changing this
     * updates the `name` of an existing firewall.
     */
    name?: pulumi.Input<string>;
    /**
     * Port(s) to associate this firewall group
     * with. Must be a list of strings. Changing this updates the associated ports
     * of an existing firewall group.
     */
    ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This argument conflicts and  is interchangeable
     * with `tenantId`. The owner of the firewall group. Required if admin wants
     * to create a firewall group for another project. Changing this creates a new
     * firewall group.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the v2 networking client.
     * A networking client is needed to create a firewall group. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * firewall group.
     */
    region?: pulumi.Input<string>;
    /**
     * Sharing status of the firewall group (must be "true"
     * or "false" if provided). If this is "true" the firewall group is visible to,
     * and can be used in, firewalls in other tenants. Changing this updates the
     * `shared` status of an existing firewall group. Only administrative users
     * can specify if the firewall group should be shared.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * This argument conflicts and is interchangeable with
     * `projectId`. The owner of the firewall group. Required if admin wants to
     * create a firewall group for another tenant. Changing this creates a new
     * firewall group.
     */
    tenantId?: pulumi.Input<string>;
}
