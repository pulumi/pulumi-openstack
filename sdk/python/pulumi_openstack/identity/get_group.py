# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, domain_id=None, name=None, region=None, id=None):
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        __self__.domain_id = domain_id
        """
        See Argument Reference above.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_group(domain_id=None,name=None,region=None,opts=None):
    """
    Use this data source to get the ID of an OpenStack group.
    
    Note: This usually requires admin privileges.
    """
    __args__ = dict()

    __args__['domainId'] = domain_id
    __args__['name'] = name
    __args__['region'] = region
    __ret__ = await pulumi.runtime.invoke('openstack:identity/getGroup:getGroup', __args__, opts=opts)

    return GetGroupResult(
        domain_id=__ret__.get('domainId'),
        name=__ret__.get('name'),
        region=__ret__.get('region'),
        id=__ret__.get('id'))
