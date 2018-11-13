# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetFloatingIpResult(object):
    """
    A collection of values returned by getFloatingIp.
    """
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_floating_ip(address=None, fixed_ip=None, pool=None, port_id=None, region=None, status=None, tenant_id=None):
    """
    Use this data source to get the ID of an available OpenStack floating IP.
    """
    __args__ = dict()

    __args__['address'] = address
    __args__['fixedIp'] = fixed_ip
    __args__['pool'] = pool
    __args__['portId'] = port_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['tenantId'] = tenant_id
    __ret__ = await pulumi.runtime.invoke('openstack:networking/getFloatingIp:getFloatingIp', __args__)

    return GetFloatingIpResult(
        id=__ret__.get('id'))
