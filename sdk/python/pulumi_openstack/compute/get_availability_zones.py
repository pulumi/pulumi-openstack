# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetAvailabilityZonesResult:
    """
    A collection of values returned by getAvailabilityZones.
    """
    def __init__(__self__, names=None, region=None, id=None):
        if names and not isinstance(names, list):
            raise TypeError('Expected argument names to be a list')
        __self__.names = names
        """
        The names of the availability zones, ordered alphanumerically, that match the queried `state`
        """
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_availability_zones(region=None,state=None,opts=None):
    """
    Use this data source to get a list of availability zones from OpenStack
    """
    __args__ = dict()

    __args__['region'] = region
    __args__['state'] = state
    __ret__ = await pulumi.runtime.invoke('openstack:compute/getAvailabilityZones:getAvailabilityZones', __args__, opts=opts)

    return GetAvailabilityZonesResult(
        names=__ret__.get('names'),
        region=__ret__.get('region'),
        id=__ret__.get('id'))
