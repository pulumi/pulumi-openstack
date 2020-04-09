# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Network(pulumi.CustomResource):
    admin_state_up: pulumi.Output[bool]
    """
    The administrative state of the network.
    Acceptable values are "true" and "false". Changing this value updates the
    state of the existing network.
    """
    all_tags: pulumi.Output[list]
    """
    The collection of tags assigned on the network, which have been
    explicitly and implicitly added.
    """
    availability_zone_hints: pulumi.Output[list]
    """
    An availability zone is used to make
    network resources highly available. Used for resources with high availability
    so that they are scheduled on different availability zones. Changing this
    creates a new network.
    """
    description: pulumi.Output[str]
    """
    Human-readable description of the network. Changing this
    updates the name of the existing network.
    """
    dns_domain: pulumi.Output[str]
    """
    The network DNS domain. Available, when Neutron DNS
    extension is enabled. The `dns_domain` of a network in conjunction with the
    `dns_name` attribute of its ports will be published in an external DNS
    service when Neutron is configured to integrate with such a service.
    """
    external: pulumi.Output[bool]
    """
    Specifies whether the network resource has the
    external routing facility. Valid values are true and false. Defaults to
    false. Changing this updates the external attribute of the existing network.
    """
    mtu: pulumi.Output[float]
    """
    The network MTU. Available for read-only, when Neutron
    `net-mtu` extension is enabled. Available for the modification, when
    Neutron `net-mtu-writable` extension is enabled.
    """
    name: pulumi.Output[str]
    """
    The name of the network. Changing this updates the name of
    the existing network.
    """
    port_security_enabled: pulumi.Output[bool]
    """
    Whether to explicitly enable or disable
    port security on the network. Port Security is usually enabled by default, so
    omitting this argument will usually result in a value of "true". Setting this
    explicitly to `false` will disable port security. Valid values are `true` and
    `false`.
    """
    qos_policy_id: pulumi.Output[str]
    """
    Reference to the associated QoS policy.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron network. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    network.
    """
    segments: pulumi.Output[list]
    """
    An array of one or more provider segment objects.

      * `network_type` (`str`) - The type of physical network.
      * `physicalNetwork` (`str`) - The physical network where this network is implemented.
      * `segmentation_id` (`float`) - An isolated segment on the physical network.
    """
    shared: pulumi.Output[bool]
    """
    Specifies whether the network resource can be accessed
    by any tenant or not. Changing this updates the sharing capabilities of the
    existing network.
    """
    tags: pulumi.Output[list]
    """
    A set of string tags for the network.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the network. Required if admin wants to
    create a network for another tenant. Changing this creates a new network.
    """
    transparent_vlan: pulumi.Output[bool]
    """
    Specifies whether the network resource has the
    VLAN transparent attribute set. Valid values are true and false. Defaults to
    false. Changing this updates the `transparent_vlan` attribute of the existing
    network.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, admin_state_up=None, availability_zone_hints=None, description=None, dns_domain=None, external=None, mtu=None, name=None, port_security_enabled=None, qos_policy_id=None, region=None, segments=None, shared=None, tags=None, tenant_id=None, transparent_vlan=None, value_specs=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V2 Neutron network resource within OpenStack.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_network_v2.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the network.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing network.
        :param pulumi.Input[list] availability_zone_hints: An availability zone is used to make
               network resources highly available. Used for resources with high availability
               so that they are scheduled on different availability zones. Changing this
               creates a new network.
        :param pulumi.Input[str] description: Human-readable description of the network. Changing this
               updates the name of the existing network.
        :param pulumi.Input[str] dns_domain: The network DNS domain. Available, when Neutron DNS
               extension is enabled. The `dns_domain` of a network in conjunction with the
               `dns_name` attribute of its ports will be published in an external DNS
               service when Neutron is configured to integrate with such a service.
        :param pulumi.Input[bool] external: Specifies whether the network resource has the
               external routing facility. Valid values are true and false. Defaults to
               false. Changing this updates the external attribute of the existing network.
        :param pulumi.Input[float] mtu: The network MTU. Available for read-only, when Neutron
               `net-mtu` extension is enabled. Available for the modification, when
               Neutron `net-mtu-writable` extension is enabled.
        :param pulumi.Input[str] name: The name of the network. Changing this updates the name of
               the existing network.
        :param pulumi.Input[bool] port_security_enabled: Whether to explicitly enable or disable
               port security on the network. Port Security is usually enabled by default, so
               omitting this argument will usually result in a value of "true". Setting this
               explicitly to `false` will disable port security. Valid values are `true` and
               `false`.
        :param pulumi.Input[str] qos_policy_id: Reference to the associated QoS policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               network.
        :param pulumi.Input[list] segments: An array of one or more provider segment objects.
        :param pulumi.Input[bool] shared: Specifies whether the network resource can be accessed
               by any tenant or not. Changing this updates the sharing capabilities of the
               existing network.
        :param pulumi.Input[list] tags: A set of string tags for the network.
        :param pulumi.Input[str] tenant_id: The owner of the network. Required if admin wants to
               create a network for another tenant. Changing this creates a new network.
        :param pulumi.Input[bool] transparent_vlan: Specifies whether the network resource has the
               VLAN transparent attribute set. Valid values are true and false. Defaults to
               false. Changing this updates the `transparent_vlan` attribute of the existing
               network.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        The **segments** object supports the following:

          * `network_type` (`pulumi.Input[str]`) - The type of physical network.
          * `physicalNetwork` (`pulumi.Input[str]`) - The physical network where this network is implemented.
          * `segmentation_id` (`pulumi.Input[float]`) - An isolated segment on the physical network.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin_state_up'] = admin_state_up
            __props__['availability_zone_hints'] = availability_zone_hints
            __props__['description'] = description
            __props__['dns_domain'] = dns_domain
            __props__['external'] = external
            __props__['mtu'] = mtu
            __props__['name'] = name
            __props__['port_security_enabled'] = port_security_enabled
            __props__['qos_policy_id'] = qos_policy_id
            __props__['region'] = region
            __props__['segments'] = segments
            __props__['shared'] = shared
            __props__['tags'] = tags
            __props__['tenant_id'] = tenant_id
            __props__['transparent_vlan'] = transparent_vlan
            __props__['value_specs'] = value_specs
            __props__['all_tags'] = None
        super(Network, __self__).__init__(
            'openstack:networking/network:Network',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admin_state_up=None, all_tags=None, availability_zone_hints=None, description=None, dns_domain=None, external=None, mtu=None, name=None, port_security_enabled=None, qos_policy_id=None, region=None, segments=None, shared=None, tags=None, tenant_id=None, transparent_vlan=None, value_specs=None):
        """
        Get an existing Network resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the network.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing network.
        :param pulumi.Input[list] all_tags: The collection of tags assigned on the network, which have been
               explicitly and implicitly added.
        :param pulumi.Input[list] availability_zone_hints: An availability zone is used to make
               network resources highly available. Used for resources with high availability
               so that they are scheduled on different availability zones. Changing this
               creates a new network.
        :param pulumi.Input[str] description: Human-readable description of the network. Changing this
               updates the name of the existing network.
        :param pulumi.Input[str] dns_domain: The network DNS domain. Available, when Neutron DNS
               extension is enabled. The `dns_domain` of a network in conjunction with the
               `dns_name` attribute of its ports will be published in an external DNS
               service when Neutron is configured to integrate with such a service.
        :param pulumi.Input[bool] external: Specifies whether the network resource has the
               external routing facility. Valid values are true and false. Defaults to
               false. Changing this updates the external attribute of the existing network.
        :param pulumi.Input[float] mtu: The network MTU. Available for read-only, when Neutron
               `net-mtu` extension is enabled. Available for the modification, when
               Neutron `net-mtu-writable` extension is enabled.
        :param pulumi.Input[str] name: The name of the network. Changing this updates the name of
               the existing network.
        :param pulumi.Input[bool] port_security_enabled: Whether to explicitly enable or disable
               port security on the network. Port Security is usually enabled by default, so
               omitting this argument will usually result in a value of "true". Setting this
               explicitly to `false` will disable port security. Valid values are `true` and
               `false`.
        :param pulumi.Input[str] qos_policy_id: Reference to the associated QoS policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               network.
        :param pulumi.Input[list] segments: An array of one or more provider segment objects.
        :param pulumi.Input[bool] shared: Specifies whether the network resource can be accessed
               by any tenant or not. Changing this updates the sharing capabilities of the
               existing network.
        :param pulumi.Input[list] tags: A set of string tags for the network.
        :param pulumi.Input[str] tenant_id: The owner of the network. Required if admin wants to
               create a network for another tenant. Changing this creates a new network.
        :param pulumi.Input[bool] transparent_vlan: Specifies whether the network resource has the
               VLAN transparent attribute set. Valid values are true and false. Defaults to
               false. Changing this updates the `transparent_vlan` attribute of the existing
               network.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        The **segments** object supports the following:

          * `network_type` (`pulumi.Input[str]`) - The type of physical network.
          * `physicalNetwork` (`pulumi.Input[str]`) - The physical network where this network is implemented.
          * `segmentation_id` (`pulumi.Input[float]`) - An isolated segment on the physical network.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_state_up"] = admin_state_up
        __props__["all_tags"] = all_tags
        __props__["availability_zone_hints"] = availability_zone_hints
        __props__["description"] = description
        __props__["dns_domain"] = dns_domain
        __props__["external"] = external
        __props__["mtu"] = mtu
        __props__["name"] = name
        __props__["port_security_enabled"] = port_security_enabled
        __props__["qos_policy_id"] = qos_policy_id
        __props__["region"] = region
        __props__["segments"] = segments
        __props__["shared"] = shared
        __props__["tags"] = tags
        __props__["tenant_id"] = tenant_id
        __props__["transparent_vlan"] = transparent_vlan
        __props__["value_specs"] = value_specs
        return Network(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

