// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 neutron security group rule resource within OpenStack.
 * Unlike Nova security groups, neutron separates the group from the rules
 * and also allows an admin to target a specific tenant_id.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const secgroup1 = new openstack.networking.SecGroup("secgroup_1", {
 *     name: "secgroup_1",
 *     description: "My neutron security group",
 * });
 * const secgroupRule1 = new openstack.networking.SecGroupRule("secgroup_rule_1", {
 *     direction: "ingress",
 *     ethertype: "IPv4",
 *     protocol: "tcp",
 *     portRangeMin: 22,
 *     portRangeMax: 22,
 *     remoteIpPrefix: "0.0.0.0/0",
 *     securityGroupId: secgroup1.id,
 * });
 * ```
 *
 * > **Note:** To expose the full port-range 1:65535, use `0` for `portRangeMin`
 * and `portRangeMax`.
 *
 * ## Import
 *
 * Security Group Rules can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:networking/secGroupRule:SecGroupRule secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745
 * ```
 */
export class SecGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing SecGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecGroupRuleState, opts?: pulumi.CustomResourceOptions): SecGroupRule {
        return new SecGroupRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/secGroupRule:SecGroupRule';

    /**
     * Returns true if the given object is an instance of SecGroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecGroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecGroupRule.__pulumiType;
    }

    /**
     * A description of the rule. Changing this creates a new security group rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The direction of the rule, valid values are __ingress__
     * or __egress__. Changing this creates a new security group rule.
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * The layer 3 protocol type, valid values are __IPv4__
     * or __IPv6__. Changing this creates a new security group rule.
     */
    public readonly ethertype!: pulumi.Output<string>;
    /**
     * The higher part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     */
    public readonly portRangeMax!: pulumi.Output<number | undefined>;
    /**
     * The lower part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     */
    public readonly portRangeMin!: pulumi.Output<number | undefined>;
    /**
     * The layer 4 protocol type, valid values are
     * following. Changing this creates a new security group rule. This is required
     * if you want to specify a port range.
     * * empty string or omitted (any protocol)
     * * integer value between 0 and 255 (valid IP protocol number)
     * * __tcp__
     * * __udp__
     * * __icmp__
     * * __ah__
     * * __dccp__
     * * __egp__
     * * __esp__
     * * __gre__
     * * __igmp__
     * * __ipv6-encap__
     * * __ipv6-frag__
     * * __ipv6-icmp__
     * * __ipv6-nonxt__
     * * __ipv6-opts__
     * * __ipv6-route__
     * * __ospf__
     * * __pgm__
     * * __rsvp__
     * * __sctp__
     * * __udplite__
     * * __vrrp__
     * * __ipip__
     */
    public readonly protocol!: pulumi.Output<string | undefined>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group rule.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The remote address group id, the value
     * needs to be an OpenStack ID of an address group in the same tenant. Changing
     * this creates a new security group rule. This argument is mutually exclusive
     * with `remoteIpPrefix` and `remoteGroupId`.
     */
    public readonly remoteAddressGroupId!: pulumi.Output<string>;
    /**
     * The remote group id, the value needs to be an
     * Openstack ID of a security group in the same tenant. Changing this creates
     * a new security group rule.
     */
    public readonly remoteGroupId!: pulumi.Output<string>;
    /**
     * The remote CIDR, the value needs to be a valid
     * CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
     */
    public readonly remoteIpPrefix!: pulumi.Output<string>;
    /**
     * The security group id the rule should belong
     * to, the value needs to be an Openstack ID of a security group in the same
     * tenant. Changing this creates a new security group rule.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group rule.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a SecGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecGroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecGroupRuleArgs | SecGroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecGroupRuleState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["ethertype"] = state ? state.ethertype : undefined;
            resourceInputs["portRangeMax"] = state ? state.portRangeMax : undefined;
            resourceInputs["portRangeMin"] = state ? state.portRangeMin : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["remoteAddressGroupId"] = state ? state.remoteAddressGroupId : undefined;
            resourceInputs["remoteGroupId"] = state ? state.remoteGroupId : undefined;
            resourceInputs["remoteIpPrefix"] = state ? state.remoteIpPrefix : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as SecGroupRuleArgs | undefined;
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.ethertype === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ethertype'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["ethertype"] = args ? args.ethertype : undefined;
            resourceInputs["portRangeMax"] = args ? args.portRangeMax : undefined;
            resourceInputs["portRangeMin"] = args ? args.portRangeMin : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["remoteAddressGroupId"] = args ? args.remoteAddressGroupId : undefined;
            resourceInputs["remoteGroupId"] = args ? args.remoteGroupId : undefined;
            resourceInputs["remoteIpPrefix"] = args ? args.remoteIpPrefix : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecGroupRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecGroupRule resources.
 */
export interface SecGroupRuleState {
    /**
     * A description of the rule. Changing this creates a new security group rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The direction of the rule, valid values are __ingress__
     * or __egress__. Changing this creates a new security group rule.
     */
    direction?: pulumi.Input<string>;
    /**
     * The layer 3 protocol type, valid values are __IPv4__
     * or __IPv6__. Changing this creates a new security group rule.
     */
    ethertype?: pulumi.Input<string>;
    /**
     * The higher part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     */
    portRangeMax?: pulumi.Input<number>;
    /**
     * The lower part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     */
    portRangeMin?: pulumi.Input<number>;
    /**
     * The layer 4 protocol type, valid values are
     * following. Changing this creates a new security group rule. This is required
     * if you want to specify a port range.
     * * empty string or omitted (any protocol)
     * * integer value between 0 and 255 (valid IP protocol number)
     * * __tcp__
     * * __udp__
     * * __icmp__
     * * __ah__
     * * __dccp__
     * * __egp__
     * * __esp__
     * * __gre__
     * * __igmp__
     * * __ipv6-encap__
     * * __ipv6-frag__
     * * __ipv6-icmp__
     * * __ipv6-nonxt__
     * * __ipv6-opts__
     * * __ipv6-route__
     * * __ospf__
     * * __pgm__
     * * __rsvp__
     * * __sctp__
     * * __udplite__
     * * __vrrp__
     * * __ipip__
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group rule.
     */
    region?: pulumi.Input<string>;
    /**
     * The remote address group id, the value
     * needs to be an OpenStack ID of an address group in the same tenant. Changing
     * this creates a new security group rule. This argument is mutually exclusive
     * with `remoteIpPrefix` and `remoteGroupId`.
     */
    remoteAddressGroupId?: pulumi.Input<string>;
    /**
     * The remote group id, the value needs to be an
     * Openstack ID of a security group in the same tenant. Changing this creates
     * a new security group rule.
     */
    remoteGroupId?: pulumi.Input<string>;
    /**
     * The remote CIDR, the value needs to be a valid
     * CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
     */
    remoteIpPrefix?: pulumi.Input<string>;
    /**
     * The security group id the rule should belong
     * to, the value needs to be an Openstack ID of a security group in the same
     * tenant. Changing this creates a new security group rule.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group rule.
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecGroupRule resource.
 */
export interface SecGroupRuleArgs {
    /**
     * A description of the rule. Changing this creates a new security group rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The direction of the rule, valid values are __ingress__
     * or __egress__. Changing this creates a new security group rule.
     */
    direction: pulumi.Input<string>;
    /**
     * The layer 3 protocol type, valid values are __IPv4__
     * or __IPv6__. Changing this creates a new security group rule.
     */
    ethertype: pulumi.Input<string>;
    /**
     * The higher part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     */
    portRangeMax?: pulumi.Input<number>;
    /**
     * The lower part of the allowed port range, valid
     * integer value needs to be between 1 and 65535. Changing this creates a new
     * security group rule.
     */
    portRangeMin?: pulumi.Input<number>;
    /**
     * The layer 4 protocol type, valid values are
     * following. Changing this creates a new security group rule. This is required
     * if you want to specify a port range.
     * * empty string or omitted (any protocol)
     * * integer value between 0 and 255 (valid IP protocol number)
     * * __tcp__
     * * __udp__
     * * __icmp__
     * * __ah__
     * * __dccp__
     * * __egp__
     * * __esp__
     * * __gre__
     * * __igmp__
     * * __ipv6-encap__
     * * __ipv6-frag__
     * * __ipv6-icmp__
     * * __ipv6-nonxt__
     * * __ipv6-opts__
     * * __ipv6-route__
     * * __ospf__
     * * __pgm__
     * * __rsvp__
     * * __sctp__
     * * __udplite__
     * * __vrrp__
     * * __ipip__
     */
    protocol?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a port. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * security group rule.
     */
    region?: pulumi.Input<string>;
    /**
     * The remote address group id, the value
     * needs to be an OpenStack ID of an address group in the same tenant. Changing
     * this creates a new security group rule. This argument is mutually exclusive
     * with `remoteIpPrefix` and `remoteGroupId`.
     */
    remoteAddressGroupId?: pulumi.Input<string>;
    /**
     * The remote group id, the value needs to be an
     * Openstack ID of a security group in the same tenant. Changing this creates
     * a new security group rule.
     */
    remoteGroupId?: pulumi.Input<string>;
    /**
     * The remote CIDR, the value needs to be a valid
     * CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
     */
    remoteIpPrefix?: pulumi.Input<string>;
    /**
     * The security group id the rule should belong
     * to, the value needs to be an Openstack ID of a security group in the same
     * tenant. Changing this creates a new security group rule.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * The owner of the security group. Required if admin
     * wants to create a port for another tenant. Changing this creates a new
     * security group rule.
     */
    tenantId?: pulumi.Input<string>;
}
