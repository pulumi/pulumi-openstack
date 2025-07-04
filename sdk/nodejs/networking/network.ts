// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a V2 Neutron network resource within OpenStack.
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
 *     networkId: network1.id,
 *     cidr: "192.168.199.0/24",
 *     ipVersion: 4,
 * });
 * const secgroup1 = new openstack.networking.SecGroup("secgroup_1", {
 *     name: "secgroup_1",
 *     description: "a security group",
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
 * const port1 = new openstack.networking.Port("port_1", {
 *     name: "port_1",
 *     networkId: network1.id,
 *     adminStateUp: true,
 *     securityGroupIds: [secgroup1.id],
 *     fixedIps: [{
 *         subnetId: subnet1.id,
 *         ipAddress: "192.168.199.10",
 *     }],
 * });
 * const instance1 = new openstack.compute.Instance("instance_1", {
 *     name: "instance_1",
 *     securityGroups: [secgroup1.name],
 *     networks: [{
 *         port: port1.id,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Networks can be imported using the `id`, e.g.
 *
 * ```sh
 * $ pulumi import openstack:networking/network:Network network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9
 * ```
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/network:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing network.
     */
    public readonly adminStateUp!: pulumi.Output<boolean>;
    /**
     * The collection of tags assigned on the network, which have been
     * explicitly and implicitly added.
     */
    public /*out*/ readonly allTags!: pulumi.Output<string[]>;
    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     */
    public readonly availabilityZoneHints!: pulumi.Output<string[]>;
    /**
     * Human-readable description of the network. Changing this
     * updates the name of the existing network.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The network DNS domain. Available, when Neutron DNS
     * extension is enabled. The `dnsDomain` of a network in conjunction with the
     * `dnsName` attribute of its ports will be published in an external DNS
     * service when Neutron is configured to integrate with such a service.
     */
    public readonly dnsDomain!: pulumi.Output<string>;
    /**
     * Specifies whether the network resource has the
     * external routing facility. Valid values are true and false. Defaults to
     * false. Changing this updates the external attribute of the existing network.
     */
    public readonly external!: pulumi.Output<boolean>;
    /**
     * The network MTU. Available for read-only, when Neutron
     * `net-mtu` extension is enabled. Available for the modification, when
     * Neutron `net-mtu-writable` extension is enabled.
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * The name of the network. Changing this updates the name of
     * the existing network.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether to explicitly enable or disable
     * port security on the network. Port Security is usually enabled by default, so
     * omitting this argument will usually result in a value of "true". Setting this
     * explicitly to `false` will disable port security. Valid values are `true` and
     * `false`.
     */
    public readonly portSecurityEnabled!: pulumi.Output<boolean>;
    /**
     * Reference to the associated QoS policy.
     */
    public readonly qosPolicyId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * An array of one or more provider segment objects.
     * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
     * updating any provider related segments attributes. Check your plug-in whether
     * it supports updating.
     */
    public readonly segments!: pulumi.Output<outputs.networking.NetworkSegment[]>;
    /**
     * Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabilities of the
     * existing network.
     */
    public readonly shared!: pulumi.Output<boolean>;
    /**
     * A set of string tags for the network.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Specifies whether the network resource has the
     * VLAN transparent attribute set. Valid values are true and false. Defaults to
     * false. Changing this updates the `transparentVlan` attribute of the existing
     * network.
     */
    public readonly transparentVlan!: pulumi.Output<boolean>;
    /**
     * Map of additional options.
     */
    public readonly valueSpecs!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["allTags"] = state ? state.allTags : undefined;
            resourceInputs["availabilityZoneHints"] = state ? state.availabilityZoneHints : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dnsDomain"] = state ? state.dnsDomain : undefined;
            resourceInputs["external"] = state ? state.external : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["portSecurityEnabled"] = state ? state.portSecurityEnabled : undefined;
            resourceInputs["qosPolicyId"] = state ? state.qosPolicyId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["segments"] = state ? state.segments : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["transparentVlan"] = state ? state.transparentVlan : undefined;
            resourceInputs["valueSpecs"] = state ? state.valueSpecs : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["availabilityZoneHints"] = args ? args.availabilityZoneHints : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsDomain"] = args ? args.dnsDomain : undefined;
            resourceInputs["external"] = args ? args.external : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["portSecurityEnabled"] = args ? args.portSecurityEnabled : undefined;
            resourceInputs["qosPolicyId"] = args ? args.qosPolicyId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["segments"] = args ? args.segments : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["transparentVlan"] = args ? args.transparentVlan : undefined;
            resourceInputs["valueSpecs"] = args ? args.valueSpecs : undefined;
            resourceInputs["allTags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Network.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing network.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The collection of tags assigned on the network, which have been
     * explicitly and implicitly added.
     */
    allTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     */
    availabilityZoneHints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable description of the network. Changing this
     * updates the name of the existing network.
     */
    description?: pulumi.Input<string>;
    /**
     * The network DNS domain. Available, when Neutron DNS
     * extension is enabled. The `dnsDomain` of a network in conjunction with the
     * `dnsName` attribute of its ports will be published in an external DNS
     * service when Neutron is configured to integrate with such a service.
     */
    dnsDomain?: pulumi.Input<string>;
    /**
     * Specifies whether the network resource has the
     * external routing facility. Valid values are true and false. Defaults to
     * false. Changing this updates the external attribute of the existing network.
     */
    external?: pulumi.Input<boolean>;
    /**
     * The network MTU. Available for read-only, when Neutron
     * `net-mtu` extension is enabled. Available for the modification, when
     * Neutron `net-mtu-writable` extension is enabled.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name of the network. Changing this updates the name of
     * the existing network.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to explicitly enable or disable
     * port security on the network. Port Security is usually enabled by default, so
     * omitting this argument will usually result in a value of "true". Setting this
     * explicitly to `false` will disable port security. Valid values are `true` and
     * `false`.
     */
    portSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * Reference to the associated QoS policy.
     */
    qosPolicyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     */
    region?: pulumi.Input<string>;
    /**
     * An array of one or more provider segment objects.
     * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
     * updating any provider related segments attributes. Check your plug-in whether
     * it supports updating.
     */
    segments?: pulumi.Input<pulumi.Input<inputs.networking.NetworkSegment>[]>;
    /**
     * Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabilities of the
     * existing network.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * A set of string tags for the network.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Specifies whether the network resource has the
     * VLAN transparent attribute set. Valid values are true and false. Defaults to
     * false. Changing this updates the `transparentVlan` attribute of the existing
     * network.
     */
    transparentVlan?: pulumi.Input<boolean>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing network.
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     */
    availabilityZoneHints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Human-readable description of the network. Changing this
     * updates the name of the existing network.
     */
    description?: pulumi.Input<string>;
    /**
     * The network DNS domain. Available, when Neutron DNS
     * extension is enabled. The `dnsDomain` of a network in conjunction with the
     * `dnsName` attribute of its ports will be published in an external DNS
     * service when Neutron is configured to integrate with such a service.
     */
    dnsDomain?: pulumi.Input<string>;
    /**
     * Specifies whether the network resource has the
     * external routing facility. Valid values are true and false. Defaults to
     * false. Changing this updates the external attribute of the existing network.
     */
    external?: pulumi.Input<boolean>;
    /**
     * The network MTU. Available for read-only, when Neutron
     * `net-mtu` extension is enabled. Available for the modification, when
     * Neutron `net-mtu-writable` extension is enabled.
     */
    mtu?: pulumi.Input<number>;
    /**
     * The name of the network. Changing this updates the name of
     * the existing network.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether to explicitly enable or disable
     * port security on the network. Port Security is usually enabled by default, so
     * omitting this argument will usually result in a value of "true". Setting this
     * explicitly to `false` will disable port security. Valid values are `true` and
     * `false`.
     */
    portSecurityEnabled?: pulumi.Input<boolean>;
    /**
     * Reference to the associated QoS policy.
     */
    qosPolicyId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     */
    region?: pulumi.Input<string>;
    /**
     * An array of one or more provider segment objects.
     * Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
     * updating any provider related segments attributes. Check your plug-in whether
     * it supports updating.
     */
    segments?: pulumi.Input<pulumi.Input<inputs.networking.NetworkSegment>[]>;
    /**
     * Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabilities of the
     * existing network.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * A set of string tags for the network.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Specifies whether the network resource has the
     * VLAN transparent attribute set. Valid values are true and false. Defaults to
     * false. Changing this updates the `transparentVlan` attribute of the existing
     * network.
     */
    transparentVlan?: pulumi.Input<boolean>;
    /**
     * Map of additional options.
     */
    valueSpecs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
