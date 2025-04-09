# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'GetAvailbilityZonesResult',
    'AwaitableGetAvailbilityZonesResult',
    'get_availbility_zones',
    'get_availbility_zones_output',
]

@pulumi.output_type
class GetAvailbilityZonesResult:
    """
    A collection of values returned by getAvailbilityZones.
    """
    def __init__(__self__, id=None, names=None, region=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[builtins.str]:
        """
        The names of the availability zones, ordered alphanumerically.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")


class AwaitableGetAvailbilityZonesResult(GetAvailbilityZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAvailbilityZonesResult(
            id=self.id,
            names=self.names,
            region=self.region)


def get_availbility_zones(region: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAvailbilityZonesResult:
    """
    Use this data source to get a list of Shared File System availability zones
    from OpenStack

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    zones = openstack.sharedfilesystem.get_availbility_zones()
    ```


    :param builtins.str region: The region in which to obtain the V2 Shared File System
           client. If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones', __args__, opts=opts, typ=GetAvailbilityZonesResult).value

    return AwaitableGetAvailbilityZonesResult(
        id=pulumi.get(__ret__, 'id'),
        names=pulumi.get(__ret__, 'names'),
        region=pulumi.get(__ret__, 'region'))
def get_availbility_zones_output(region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAvailbilityZonesResult]:
    """
    Use this data source to get a list of Shared File System availability zones
    from OpenStack

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    zones = openstack.sharedfilesystem.get_availbility_zones()
    ```


    :param builtins.str region: The region in which to obtain the V2 Shared File System
           client. If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones', __args__, opts=opts, typ=GetAvailbilityZonesResult)
    return __ret__.apply(lambda __response__: GetAvailbilityZonesResult(
        id=pulumi.get(__response__, 'id'),
        names=pulumi.get(__response__, 'names'),
        region=pulumi.get(__response__, 'region')))
