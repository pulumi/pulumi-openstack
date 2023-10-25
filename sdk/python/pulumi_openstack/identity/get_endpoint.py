# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetEndpointResult',
    'AwaitableGetEndpointResult',
    'get_endpoint',
    'get_endpoint_output',
]

@pulumi.output_type
class GetEndpointResult:
    """
    A collection of values returned by getEndpoint.
    """
    def __init__(__self__, endpoint_region=None, id=None, interface=None, name=None, region=None, service_id=None, service_name=None, service_type=None, url=None):
        if endpoint_region and not isinstance(endpoint_region, str):
            raise TypeError("Expected argument 'endpoint_region' to be a str")
        pulumi.set(__self__, "endpoint_region", endpoint_region)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if service_type and not isinstance(service_type, str):
            raise TypeError("Expected argument 'service_type' to be a str")
        pulumi.set(__self__, "service_type", service_type)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="endpointRegion")
    def endpoint_region(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "endpoint_region")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The endpoint URL.
        """
        return pulumi.get(self, "url")


class AwaitableGetEndpointResult(GetEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointResult(
            endpoint_region=self.endpoint_region,
            id=self.id,
            interface=self.interface,
            name=self.name,
            region=self.region,
            service_id=self.service_id,
            service_name=self.service_name,
            service_type=self.service_type,
            url=self.url)


def get_endpoint(endpoint_region: Optional[str] = None,
                 interface: Optional[str] = None,
                 name: Optional[str] = None,
                 region: Optional[str] = None,
                 service_id: Optional[str] = None,
                 service_name: Optional[str] = None,
                 service_type: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointResult:
    """
    Use this data source to get the ID of an OpenStack endpoint.

    > **Note:** This usually requires admin privileges.


    :param str endpoint_region: The region the endpoint is assigned to. The
           `region` and `endpoint_region` can be different.
    :param str interface: The endpoint interface. Valid values are `public`,
           `internal`, and `admin`. Default value is `public`
    :param str name: The name of the endpoint.
    :param str region: The region in which to obtain the V3 Keystone client.
           If omitted, the `region` argument of the provider is used.
    :param str service_id: The service id this endpoint belongs to.
    :param str service_name: The service name of the endpoint.
    :param str service_type: The service type of the endpoint.
    """
    __args__ = dict()
    __args__['endpointRegion'] = endpoint_region
    __args__['interface'] = interface
    __args__['name'] = name
    __args__['region'] = region
    __args__['serviceId'] = service_id
    __args__['serviceName'] = service_name
    __args__['serviceType'] = service_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:identity/getEndpoint:getEndpoint', __args__, opts=opts, typ=GetEndpointResult).value

    return AwaitableGetEndpointResult(
        endpoint_region=pulumi.get(__ret__, 'endpoint_region'),
        id=pulumi.get(__ret__, 'id'),
        interface=pulumi.get(__ret__, 'interface'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        service_id=pulumi.get(__ret__, 'service_id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        service_type=pulumi.get(__ret__, 'service_type'),
        url=pulumi.get(__ret__, 'url'))


@_utilities.lift_output_func(get_endpoint)
def get_endpoint_output(endpoint_region: Optional[pulumi.Input[Optional[str]]] = None,
                        interface: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        service_id: Optional[pulumi.Input[Optional[str]]] = None,
                        service_name: Optional[pulumi.Input[Optional[str]]] = None,
                        service_type: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEndpointResult]:
    """
    Use this data source to get the ID of an OpenStack endpoint.

    > **Note:** This usually requires admin privileges.


    :param str endpoint_region: The region the endpoint is assigned to. The
           `region` and `endpoint_region` can be different.
    :param str interface: The endpoint interface. Valid values are `public`,
           `internal`, and `admin`. Default value is `public`
    :param str name: The name of the endpoint.
    :param str region: The region in which to obtain the V3 Keystone client.
           If omitted, the `region` argument of the provider is used.
    :param str service_id: The service id this endpoint belongs to.
    :param str service_name: The service name of the endpoint.
    :param str service_type: The service type of the endpoint.
    """
    ...
