# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetFloatingIpResult:
    """
    A collection of values returned by getFloatingIp.
    """
    def __init__(__self__, all_tags=None, id=None):
        if all_tags and not isinstance(all_tags, list):
            raise TypeError('Expected argument all_tags to be a list')
        __self__.all_tags = all_tags
        """
        A set of string tags applied on the floating IP.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_floating_ip(address=None,description=None,fixed_ip=None,pool=None,port_id=None,region=None,status=None,tags=None,tenant_id=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack floating IP.
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
    __ret__ = await pulumi.runtime.invoke('openstack:networking/getFloatingIp:getFloatingIp', __args__, opts=opts)

    return GetFloatingIpResult(
        all_tags=__ret__.get('allTags'),
        id=__ret__.get('id'))
