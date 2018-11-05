# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetEndpointResult(object):
    """
    A collection of values returned by getEndpoint.
    """
    def __init__(__self__, region=None, url=None, id=None):
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
        """
        The region the endpoint is located in.
        """
        if url and not isinstance(url, str):
            raise TypeError('Expected argument url to be a str')
        __self__.url = url
        """
        The endpoint URL
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_endpoint(interface=None, region=None, service_id=None, service_name=None):
    """
    Use this data source to get the ID of an OpenStack endpoint.
    
    Note: This usually requires admin privileges.
    """
    __args__ = dict()

    __args__['interface'] = interface
    __args__['region'] = region
    __args__['serviceId'] = service_id
    __args__['serviceName'] = service_name
    __ret__ = pulumi.runtime.invoke('openstack:identity/getEndpoint:getEndpoint', __args__)

    return GetEndpointResult(
        region=__ret__.get('region'),
        url=__ret__.get('url'),
        id=__ret__.get('id'))
