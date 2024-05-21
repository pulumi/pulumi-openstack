# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetFlavorV2Result',
    'AwaitableGetFlavorV2Result',
    'get_flavor_v2',
    'get_flavor_v2_output',
]

@pulumi.output_type
class GetFlavorV2Result:
    """
    A collection of values returned by getFlavorV2.
    """
    def __init__(__self__, description=None, enabled=None, flavor_id=None, flavor_profile_id=None, id=None, name=None, region=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if flavor_profile_id and not isinstance(flavor_profile_id, str):
            raise TypeError("Expected argument 'flavor_profile_id' to be a str")
        pulumi.set(__self__, "flavor_profile_id", flavor_profile_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> str:
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="flavorProfileId")
    def flavor_profile_id(self) -> str:
        return pulumi.get(self, "flavor_profile_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")


class AwaitableGetFlavorV2Result(GetFlavorV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlavorV2Result(
            description=self.description,
            enabled=self.enabled,
            flavor_id=self.flavor_id,
            flavor_profile_id=self.flavor_profile_id,
            id=self.id,
            name=self.name,
            region=self.region)


def get_flavor_v2(flavor_id: Optional[str] = None,
                  name: Optional[str] = None,
                  region: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlavorV2Result:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['flavorId'] = flavor_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:loadbalancer/getFlavorV2:getFlavorV2', __args__, opts=opts, typ=GetFlavorV2Result).value

    return AwaitableGetFlavorV2Result(
        description=pulumi.get(__ret__, 'description'),
        enabled=pulumi.get(__ret__, 'enabled'),
        flavor_id=pulumi.get(__ret__, 'flavor_id'),
        flavor_profile_id=pulumi.get(__ret__, 'flavor_profile_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'))


@_utilities.lift_output_func(get_flavor_v2)
def get_flavor_v2_output(flavor_id: Optional[pulumi.Input[Optional[str]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlavorV2Result]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
