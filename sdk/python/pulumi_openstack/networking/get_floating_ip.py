# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetFloatingIpResult:
    """
    A collection of values returned by getFloatingIp.
    """
    def __init__(__self__, address=None, all_tags=None, description=None, dns_domain=None, dns_name=None, fixed_ip=None, id=None, pool=None, port_id=None, region=None, status=None, tags=None, tenant_id=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        __self__.address = address
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        __self__.all_tags = all_tags
        """
        A set of string tags applied on the floating IP.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if dns_domain and not isinstance(dns_domain, str):
            raise TypeError("Expected argument 'dns_domain' to be a str")
        __self__.dns_domain = dns_domain
        """
        The floating IP DNS domain. Available, when Neutron DNS
        extension is enabled.
        """
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        __self__.dns_name = dns_name
        """
        The floating IP DNS name. Available, when Neutron DNS extension
        is enabled.
        """
        if fixed_ip and not isinstance(fixed_ip, str):
            raise TypeError("Expected argument 'fixed_ip' to be a str")
        __self__.fixed_ip = fixed_ip
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if pool and not isinstance(pool, str):
            raise TypeError("Expected argument 'pool' to be a str")
        __self__.pool = pool
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        __self__.port_id = port_id
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        __self__.tenant_id = tenant_id
class AwaitableGetFloatingIpResult(GetFloatingIpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFloatingIpResult(
            address=self.address,
            all_tags=self.all_tags,
            description=self.description,
            dns_domain=self.dns_domain,
            dns_name=self.dns_name,
            fixed_ip=self.fixed_ip,
            id=self.id,
            pool=self.pool,
            port_id=self.port_id,
            region=self.region,
            status=self.status,
            tags=self.tags,
            tenant_id=self.tenant_id)

def get_floating_ip(address=None,description=None,fixed_ip=None,pool=None,port_id=None,region=None,status=None,tags=None,tenant_id=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack floating IP.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_floatingip_v2.html.markdown.


    :param str address: The IP address of the floating IP.
    :param str description: Human-readable description of the floating IP.
    :param str fixed_ip: The specific IP address of the internal port which should be associated with the floating IP.
    :param str pool: The name of the pool from which the floating IP belongs to.
    :param str port_id: The ID of the port the floating IP is attached.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve floating IP ids. If omitted, the
           `region` argument of the provider is used.
    :param str status: status of the floating IP (ACTIVE/DOWN).
    :param list tags: The list of floating IP tags to filter.
    :param str tenant_id: The owner of the floating IP.
    """
    __args__ = dict()


    __args__['address'] = address
    __args__['description'] = description
    __args__['fixedIp'] = fixed_ip
    __args__['pool'] = pool
    __args__['portId'] = port_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:networking/getFloatingIp:getFloatingIp', __args__, opts=opts).value

    return AwaitableGetFloatingIpResult(
        address=__ret__.get('address'),
        all_tags=__ret__.get('allTags'),
        description=__ret__.get('description'),
        dns_domain=__ret__.get('dnsDomain'),
        dns_name=__ret__.get('dnsName'),
        fixed_ip=__ret__.get('fixedIp'),
        id=__ret__.get('id'),
        pool=__ret__.get('pool'),
        port_id=__ret__.get('portId'),
        region=__ret__.get('region'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        tenant_id=__ret__.get('tenantId'))
