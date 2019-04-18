# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SecGroupRule(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description of the rule. Changing this creates a new security group rule.
    """
    direction: pulumi.Output[str]
    """
    The direction of the rule, valid values are __ingress__
    or __egress__. Changing this creates a new security group rule.
    """
    ethertype: pulumi.Output[str]
    """
    The layer 3 protocol type, valid values are __IPv4__
    or __IPv6__. Changing this creates a new security group rule.
    """
    port_range_max: pulumi.Output[float]
    """
    The higher part of the allowed port range, valid
    integer value needs to be between 1 and 65535. Changing this creates a new
    security group rule.
    """
    port_range_min: pulumi.Output[float]
    """
    The lower part of the allowed port range, valid
    integer value needs to be between 1 and 65535. Changing this creates a new
    security group rule.
    """
    protocol: pulumi.Output[str]
    """
    The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.
    * __tcp__
    * __udp__
    * __icmp__
    * __ah__
    * __dccp__
    * __egp__
    * __esp__
    * __gre__
    * __igmp__
    * __ipv6-encap__
    * __ipv6-frag__
    * __ipv6-icmp__
    * __ipv6-nonxt__
    * __ipv6-opts__
    * __ipv6-route__
    * __ospf__
    * __pgm__
    * __rsvp__
    * __sctp__
    * __udplite__
    * __vrrp__
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 networking client.
    A networking client is needed to create a port. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    security group rule.
    """
    remote_group_id: pulumi.Output[str]
    """
    The remote group id, the value needs to be an
    Openstack ID of a security group in the same tenant. Changing this creates
    a new security group rule.
    """
    remote_ip_prefix: pulumi.Output[str]
    """
    The remote CIDR, the value needs to be a valid
    CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
    """
    security_group_id: pulumi.Output[str]
    """
    The security group id the rule should belong
    to, the value needs to be an Openstack ID of a security group in the same
    tenant. Changing this creates a new security group rule.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the security group. Required if admin
    wants to create a port for another tenant. Changing this creates a new
    security group rule.
    """
    def __init__(__self__, resource_name, opts=None, description=None, direction=None, ethertype=None, port_range_max=None, port_range_min=None, protocol=None, region=None, remote_group_id=None, remote_ip_prefix=None, security_group_id=None, tenant_id=None, __name__=None, __opts__=None):
        """
        Manages a V2 neutron security group rule resource within OpenStack.
        Unlike Nova security groups, neutron separates the group from the rules
        and also allows an admin to target a specific tenant_id.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the rule. Changing this creates a new security group rule.
        :param pulumi.Input[str] direction: The direction of the rule, valid values are __ingress__
               or __egress__. Changing this creates a new security group rule.
        :param pulumi.Input[str] ethertype: The layer 3 protocol type, valid values are __IPv4__
               or __IPv6__. Changing this creates a new security group rule.
        :param pulumi.Input[float] port_range_max: The higher part of the allowed port range, valid
               integer value needs to be between 1 and 65535. Changing this creates a new
               security group rule.
        :param pulumi.Input[float] port_range_min: The lower part of the allowed port range, valid
               integer value needs to be between 1 and 65535. Changing this creates a new
               security group rule.
        :param pulumi.Input[str] protocol: The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.
               * __tcp__
               * __udp__
               * __icmp__
               * __ah__
               * __dccp__
               * __egp__
               * __esp__
               * __gre__
               * __igmp__
               * __ipv6-encap__
               * __ipv6-frag__
               * __ipv6-icmp__
               * __ipv6-nonxt__
               * __ipv6-opts__
               * __ipv6-route__
               * __ospf__
               * __pgm__
               * __rsvp__
               * __sctp__
               * __udplite__
               * __vrrp__
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group rule.
        :param pulumi.Input[str] remote_group_id: The remote group id, the value needs to be an
               Openstack ID of a security group in the same tenant. Changing this creates
               a new security group rule.
        :param pulumi.Input[str] remote_ip_prefix: The remote CIDR, the value needs to be a valid
               CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
        :param pulumi.Input[str] security_group_id: The security group id the rule should belong
               to, the value needs to be an Openstack ID of a security group in the same
               tenant. Changing this creates a new security group rule.
        :param pulumi.Input[str] tenant_id: The owner of the security group. Required if admin
               wants to create a port for another tenant. Changing this creates a new
               security group rule.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        if direction is None:
            raise TypeError("Missing required property 'direction'")
        __props__['direction'] = direction

        if ethertype is None:
            raise TypeError("Missing required property 'ethertype'")
        __props__['ethertype'] = ethertype

        __props__['port_range_max'] = port_range_max

        __props__['port_range_min'] = port_range_min

        __props__['protocol'] = protocol

        __props__['region'] = region

        __props__['remote_group_id'] = remote_group_id

        __props__['remote_ip_prefix'] = remote_ip_prefix

        if security_group_id is None:
            raise TypeError("Missing required property 'security_group_id'")
        __props__['security_group_id'] = security_group_id

        __props__['tenant_id'] = tenant_id

        super(SecGroupRule, __self__).__init__(
            'openstack:networking/secGroupRule:SecGroupRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

