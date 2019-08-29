# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Vip(pulumi.CustomResource):
    address: pulumi.Output[str]
    """
    The IP address of the vip. Changing this creates a new
    vip.
    """
    admin_state_up: pulumi.Output[bool]
    """
    The administrative state of the vip.
    Acceptable values are "true" and "false". Changing this value updates the
    state of the existing vip.
    """
    conn_limit: pulumi.Output[float]
    """
    The maximum number of connections allowed for the
    vip. Default is -1, meaning no limit. Changing this updates the conn_limit
    of the existing vip.
    """
    description: pulumi.Output[str]
    """
    Human-readable description for the vip. Changing
    this updates the description of the existing vip.
    """
    floating_ip: pulumi.Output[str]
    """
    A *Networking* Floating IP that will be associated
    with the vip. The Floating IP must be provisioned already.
    """
    name: pulumi.Output[str]
    """
    The name of the vip. Changing this updates the name of
    the existing vip.
    """
    persistence: pulumi.Output[dict]
    """
    Omit this field to prevent session persistence.
    The persistence object structure is documented below. Changing this updates
    the persistence of the existing vip.
    """
    pool_id: pulumi.Output[str]
    """
    The ID of the pool with which the vip is associated.
    Changing this updates the pool_id of the existing vip.
    """
    port: pulumi.Output[float]
    """
    The port on which to listen for client traffic. Changing
    this creates a new vip.
    """
    port_id: pulumi.Output[str]
    """
    Port UUID for this VIP at associated floating IP (if any).
    """
    protocol: pulumi.Output[str]
    """
    The protocol - can be either 'TCP, 'HTTP', or
    HTTPS'. Changing this creates a new vip.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a VIP. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    VIP.
    """
    subnet_id: pulumi.Output[str]
    """
    The network on which to allocate the vip's address. A
    tenant can only create vips on networks authorized by policy (e.g. networks
    that belong to them or networks that are shared). Changing this creates a
    new vip.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the vip. Required if admin wants to
    create a vip member for another tenant. Changing this creates a new vip.
    """
    def __init__(__self__, resource_name, opts=None, address=None, admin_state_up=None, conn_limit=None, description=None, floating_ip=None, name=None, persistence=None, pool_id=None, port=None, protocol=None, region=None, subnet_id=None, tenant_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V1 load balancer vip resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address of the vip. Changing this creates a new
               vip.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the vip.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing vip.
        :param pulumi.Input[float] conn_limit: The maximum number of connections allowed for the
               vip. Default is -1, meaning no limit. Changing this updates the conn_limit
               of the existing vip.
        :param pulumi.Input[str] description: Human-readable description for the vip. Changing
               this updates the description of the existing vip.
        :param pulumi.Input[str] floating_ip: A *Networking* Floating IP that will be associated
               with the vip. The Floating IP must be provisioned already.
        :param pulumi.Input[str] name: The name of the vip. Changing this updates the name of
               the existing vip.
        :param pulumi.Input[dict] persistence: Omit this field to prevent session persistence.
               The persistence object structure is documented below. Changing this updates
               the persistence of the existing vip.
        :param pulumi.Input[str] pool_id: The ID of the pool with which the vip is associated.
               Changing this updates the pool_id of the existing vip.
        :param pulumi.Input[float] port: The port on which to listen for client traffic. Changing
               this creates a new vip.
        :param pulumi.Input[str] protocol: The protocol - can be either 'TCP, 'HTTP', or
               HTTPS'. Changing this creates a new vip.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a VIP. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               VIP.
        :param pulumi.Input[str] subnet_id: The network on which to allocate the vip's address. A
               tenant can only create vips on networks authorized by policy (e.g. networks
               that belong to them or networks that are shared). Changing this creates a
               new vip.
        :param pulumi.Input[str] tenant_id: The owner of the vip. Required if admin wants to
               create a vip member for another tenant. Changing this creates a new vip.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_vip_v1.html.markdown.
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

            __props__['address'] = address
            __props__['admin_state_up'] = admin_state_up
            __props__['conn_limit'] = conn_limit
            __props__['description'] = description
            __props__['floating_ip'] = floating_ip
            __props__['name'] = name
            __props__['persistence'] = persistence
            if pool_id is None:
                raise TypeError("Missing required property 'pool_id'")
            __props__['pool_id'] = pool_id
            if port is None:
                raise TypeError("Missing required property 'port'")
            __props__['port'] = port
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            __props__['region'] = region
            if subnet_id is None:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            __props__['tenant_id'] = tenant_id
            __props__['port_id'] = None
        super(Vip, __self__).__init__(
            'openstack:loadbalancer/vip:Vip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address=None, admin_state_up=None, conn_limit=None, description=None, floating_ip=None, name=None, persistence=None, pool_id=None, port=None, port_id=None, protocol=None, region=None, subnet_id=None, tenant_id=None):
        """
        Get an existing Vip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address of the vip. Changing this creates a new
               vip.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the vip.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing vip.
        :param pulumi.Input[float] conn_limit: The maximum number of connections allowed for the
               vip. Default is -1, meaning no limit. Changing this updates the conn_limit
               of the existing vip.
        :param pulumi.Input[str] description: Human-readable description for the vip. Changing
               this updates the description of the existing vip.
        :param pulumi.Input[str] floating_ip: A *Networking* Floating IP that will be associated
               with the vip. The Floating IP must be provisioned already.
        :param pulumi.Input[str] name: The name of the vip. Changing this updates the name of
               the existing vip.
        :param pulumi.Input[dict] persistence: Omit this field to prevent session persistence.
               The persistence object structure is documented below. Changing this updates
               the persistence of the existing vip.
        :param pulumi.Input[str] pool_id: The ID of the pool with which the vip is associated.
               Changing this updates the pool_id of the existing vip.
        :param pulumi.Input[float] port: The port on which to listen for client traffic. Changing
               this creates a new vip.
        :param pulumi.Input[str] port_id: Port UUID for this VIP at associated floating IP (if any).
        :param pulumi.Input[str] protocol: The protocol - can be either 'TCP, 'HTTP', or
               HTTPS'. Changing this creates a new vip.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a VIP. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               VIP.
        :param pulumi.Input[str] subnet_id: The network on which to allocate the vip's address. A
               tenant can only create vips on networks authorized by policy (e.g. networks
               that belong to them or networks that are shared). Changing this creates a
               new vip.
        :param pulumi.Input[str] tenant_id: The owner of the vip. Required if admin wants to
               create a vip member for another tenant. Changing this creates a new vip.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/lb_vip_v1.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["address"] = address
        __props__["admin_state_up"] = admin_state_up
        __props__["conn_limit"] = conn_limit
        __props__["description"] = description
        __props__["floating_ip"] = floating_ip
        __props__["name"] = name
        __props__["persistence"] = persistence
        __props__["pool_id"] = pool_id
        __props__["port"] = port
        __props__["port_id"] = port_id
        __props__["protocol"] = protocol
        __props__["region"] = region
        __props__["subnet_id"] = subnet_id
        __props__["tenant_id"] = tenant_id
        return Vip(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

