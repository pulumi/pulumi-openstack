# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSecGroupResult:
    """
    A collection of values returned by getSecGroup.
    """
    def __init__(__self__, all_tags=None, region=None, tenant_id=None, id=None):
        if all_tags and not isinstance(all_tags, list):
            raise TypeError('Expected argument all_tags to be a list')
        __self__.all_tags = all_tags
        """
        The set of string tags applied on the security group.
        """
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
        """
        See Argument Reference above.
        """
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError('Expected argument tenant_id to be a str')
        __self__.tenant_id = tenant_id
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_sec_group(description=None,name=None,region=None,secgroup_id=None,tags=None,tenant_id=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack security group.
    """
    __args__ = dict()

    __args__['description'] = description
    __args__['name'] = name
    __args__['region'] = region
    __args__['secgroupId'] = secgroup_id
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    __ret__ = await pulumi.runtime.invoke('openstack:networking/getSecGroup:getSecGroup', __args__, opts=opts)

    return GetSecGroupResult(
        all_tags=__ret__.get('allTags'),
        region=__ret__.get('region'),
        tenant_id=__ret__.get('tenantId'),
        id=__ret__.get('id'))
