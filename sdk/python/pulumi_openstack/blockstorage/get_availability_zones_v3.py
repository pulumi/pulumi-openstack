# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetAvailabilityZonesV3Result:
    """
    A collection of values returned by getAvailabilityZonesV3.
    """
    def __init__(__self__, names=None, region=None, state=None, id=None):
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        __self__.names = names
        """
        The names of the availability zones, ordered alphanumerically, that
        match the queried `state`.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_availability_zones_v3(region=None,state=None,opts=None):
    """
    Use this data source to get a list of Block Storage availability zones from OpenStack
    """
    __args__ = dict()

    __args__['region'] = region
    __args__['state'] = state
    __ret__ = await pulumi.runtime.invoke('openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3', __args__, opts=opts)

    return GetAvailabilityZonesV3Result(
        names=__ret__.get('names'),
        region=__ret__.get('region'),
        state=__ret__.get('state'),
        id=__ret__.get('id'))
