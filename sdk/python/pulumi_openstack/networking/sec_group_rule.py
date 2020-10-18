# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SecGroupRule']


class SecGroupRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 ethertype: Optional[pulumi.Input[str]] = None,
                 port_range_max: Optional[pulumi.Input[int]] = None,
                 port_range_min: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 remote_group_id: Optional[pulumi.Input[str]] = None,
                 remote_ip_prefix: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V2 neutron security group rule resource within OpenStack.
        Unlike Nova security groups, neutron separates the group from the rules
        and also allows an admin to target a specific tenant_id.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        secgroup1 = openstack.networking.SecGroup("secgroup1", description="My neutron security group")
        secgroup_rule1 = openstack.networking.SecGroupRule("secgroupRule1",
            direction="ingress",
            ethertype="IPv4",
            port_range_max=22,
            port_range_min=22,
            protocol="tcp",
            remote_ip_prefix="0.0.0.0/0",
            security_group_id=secgroup1.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the rule. Changing this creates a new security group rule.
        :param pulumi.Input[str] direction: The direction of the rule, valid values are __ingress__
               or __egress__. Changing this creates a new security group rule.
        :param pulumi.Input[str] ethertype: The layer 3 protocol type, valid values are __IPv4__
               or __IPv6__. Changing this creates a new security group rule.
        :param pulumi.Input[int] port_range_max: The higher part of the allowed port range, valid
               integer value needs to be between 1 and 65535. Changing this creates a new
               security group rule.
        :param pulumi.Input[int] port_range_min: The lower part of the allowed port range, valid
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
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            direction: Optional[pulumi.Input[str]] = None,
            ethertype: Optional[pulumi.Input[str]] = None,
            port_range_max: Optional[pulumi.Input[int]] = None,
            port_range_min: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            remote_group_id: Optional[pulumi.Input[str]] = None,
            remote_ip_prefix: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'SecGroupRule':
        """
        Get an existing SecGroupRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the rule. Changing this creates a new security group rule.
        :param pulumi.Input[str] direction: The direction of the rule, valid values are __ingress__
               or __egress__. Changing this creates a new security group rule.
        :param pulumi.Input[str] ethertype: The layer 3 protocol type, valid values are __IPv4__
               or __IPv6__. Changing this creates a new security group rule.
        :param pulumi.Input[int] port_range_max: The higher part of the allowed port range, valid
               integer value needs to be between 1 and 65535. Changing this creates a new
               security group rule.
        :param pulumi.Input[int] port_range_min: The lower part of the allowed port range, valid
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
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["direction"] = direction
        __props__["ethertype"] = ethertype
        __props__["port_range_max"] = port_range_max
        __props__["port_range_min"] = port_range_min
        __props__["protocol"] = protocol
        __props__["region"] = region
        __props__["remote_group_id"] = remote_group_id
        __props__["remote_ip_prefix"] = remote_ip_prefix
        __props__["security_group_id"] = security_group_id
        __props__["tenant_id"] = tenant_id
        return SecGroupRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the rule. Changing this creates a new security group rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[str]:
        """
        The direction of the rule, valid values are __ingress__
        or __egress__. Changing this creates a new security group rule.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def ethertype(self) -> pulumi.Output[str]:
        """
        The layer 3 protocol type, valid values are __IPv4__
        or __IPv6__. Changing this creates a new security group rule.
        """
        return pulumi.get(self, "ethertype")

    @property
    @pulumi.getter(name="portRangeMax")
    def port_range_max(self) -> pulumi.Output[int]:
        """
        The higher part of the allowed port range, valid
        integer value needs to be between 1 and 65535. Changing this creates a new
        security group rule.
        """
        return pulumi.get(self, "port_range_max")

    @property
    @pulumi.getter(name="portRangeMin")
    def port_range_min(self) -> pulumi.Output[int]:
        """
        The lower part of the allowed port range, valid
        integer value needs to be between 1 and 65535. Changing this creates a new
        security group rule.
        """
        return pulumi.get(self, "port_range_min")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
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
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to create a port. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        security group rule.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="remoteGroupId")
    def remote_group_id(self) -> pulumi.Output[str]:
        """
        The remote group id, the value needs to be an
        Openstack ID of a security group in the same tenant. Changing this creates
        a new security group rule.
        """
        return pulumi.get(self, "remote_group_id")

    @property
    @pulumi.getter(name="remoteIpPrefix")
    def remote_ip_prefix(self) -> pulumi.Output[str]:
        """
        The remote CIDR, the value needs to be a valid
        CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.
        """
        return pulumi.get(self, "remote_ip_prefix")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The security group id the rule should belong
        to, the value needs to be an Openstack ID of a security group in the same
        tenant. Changing this creates a new security group rule.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The owner of the security group. Required if admin
        wants to create a port for another tenant. Changing this creates a new
        security group rule.
        """
        return pulumi.get(self, "tenant_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

