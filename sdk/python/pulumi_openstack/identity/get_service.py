# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetServiceResult:
    """
    A collection of values returned by getService.
    """
    def __init__(__self__, description=None, enabled=None, id=None, name=None, region=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The service description.
        """
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
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
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        See Argument Reference above.
        """
class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            description=self.description,
            enabled=self.enabled,
            id=self.id,
            name=self.name,
            region=self.region,
            type=self.type)

def get_service(enabled=None,name=None,region=None,type=None,opts=None):
    """
    Use this data source to get the ID of an OpenStack service.

    > **Note:** This usually requires admin privileges.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/identity_service_v3.html.markdown.


    :param bool enabled: The service status.
    :param str name: The service name.
    :param str region: The region in which to obtain the V3 Keystone client.
           If omitted, the `region` argument of the provider is used.
    :param str type: The service type.
    """
    __args__ = dict()


    __args__['enabled'] = enabled
    __args__['name'] = name
    __args__['region'] = region
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:identity/getService:getService', __args__, opts=opts).value

    return AwaitableGetServiceResult(
        description=__ret__.get('description'),
        enabled=__ret__.get('enabled'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        region=__ret__.get('region'),
        type=__ret__.get('type'))
