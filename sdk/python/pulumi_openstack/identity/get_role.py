# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetRoleResult(object):
    """
    A collection of values returned by getRole.
    """
    def __init__(__self__, domain_id=None, region=None, id=None):
        if domain_id and not isinstance(domain_id, str):
            raise TypeError('Expected argument domain_id to be a str')
        __self__.domain_id = domain_id
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_role(domain_id=None, name=None, region=None):
    """
    Use this data source to get the ID of an OpenStack role.
    """
    __args__ = dict()

    __args__['domainId'] = domain_id
    __args__['name'] = name
    __args__['region'] = region
    __ret__ = pulumi.runtime.invoke('openstack:identity/getRole:getRole', __args__)

    return GetRoleResult(
        domain_id=__ret__.get('domainId'),
        region=__ret__.get('region'),
        id=__ret__.get('id'))
