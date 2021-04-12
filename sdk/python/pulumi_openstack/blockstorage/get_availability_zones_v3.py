# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetAvailabilityZonesV3Result',
    'AwaitableGetAvailabilityZonesV3Result',
    'get_availability_zones_v3',
]

@pulumi.output_type
class GetAvailabilityZonesV3Result:
    """
    A collection of values returned by getAvailabilityZonesV3.
    """
    def __init__(__self__, id=None, names=None, region=None, state=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        The names of the availability zones, ordered alphanumerically, that
        match the queried `state`.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "state")


class AwaitableGetAvailabilityZonesV3Result(GetAvailabilityZonesV3Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAvailabilityZonesV3Result(
            id=self.id,
            names=self.names,
            region=self.region,
            state=self.state)


def get_availability_zones_v3(region: Optional[str] = None,
                              state: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAvailabilityZonesV3Result:
    """
    Use this data source to get a list of Block Storage availability zones from OpenStack

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    zones = openstack.blockstorage.get_availability_zones_v3()
    ```


    :param str region: The region in which to obtain the Block Storage client.
           If omitted, the `region` argument of the provider is used.
    :param str state: The `state` of the availability zones to match. Can
           either be `available` or `unavailable`. Default is `available`.
    """
    __args__ = dict()
    __args__['region'] = region
    __args__['state'] = state
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:blockstorage/getAvailabilityZonesV3:getAvailabilityZonesV3', __args__, opts=opts, typ=GetAvailabilityZonesV3Result).value

    return AwaitableGetAvailabilityZonesV3Result(
        id=__ret__.id,
        names=__ret__.names,
        region=__ret__.region,
        state=__ret__.state)
