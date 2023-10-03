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
    'GetFlavorResult',
    'AwaitableGetFlavorResult',
    'get_flavor',
    'get_flavor_output',
]

@pulumi.output_type
class GetFlavorResult:
    """
    A collection of values returned by getFlavor.
    """
    def __init__(__self__, description=None, disk=None, extra_specs=None, flavor_id=None, id=None, is_public=None, min_disk=None, min_ram=None, name=None, ram=None, region=None, rx_tx_factor=None, swap=None, vcpus=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk and not isinstance(disk, int):
            raise TypeError("Expected argument 'disk' to be a int")
        pulumi.set(__self__, "disk", disk)
        if extra_specs and not isinstance(extra_specs, dict):
            raise TypeError("Expected argument 'extra_specs' to be a dict")
        pulumi.set(__self__, "extra_specs", extra_specs)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_public and not isinstance(is_public, bool):
            raise TypeError("Expected argument 'is_public' to be a bool")
        pulumi.set(__self__, "is_public", is_public)
        if min_disk and not isinstance(min_disk, int):
            raise TypeError("Expected argument 'min_disk' to be a int")
        pulumi.set(__self__, "min_disk", min_disk)
        if min_ram and not isinstance(min_ram, int):
            raise TypeError("Expected argument 'min_ram' to be a int")
        pulumi.set(__self__, "min_ram", min_ram)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ram and not isinstance(ram, int):
            raise TypeError("Expected argument 'ram' to be a int")
        pulumi.set(__self__, "ram", ram)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if rx_tx_factor and not isinstance(rx_tx_factor, float):
            raise TypeError("Expected argument 'rx_tx_factor' to be a float")
        pulumi.set(__self__, "rx_tx_factor", rx_tx_factor)
        if swap and not isinstance(swap, int):
            raise TypeError("Expected argument 'swap' to be a int")
        pulumi.set(__self__, "swap", swap)
        if vcpus and not isinstance(vcpus, int):
            raise TypeError("Expected argument 'vcpus' to be a int")
        pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disk(self) -> Optional[int]:
        return pulumi.get(self, "disk")

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> Mapping[str, Any]:
        """
        Key/Value pairs of metadata for the flavor.
        """
        return pulumi.get(self, "extra_specs")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[str]:
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[bool]:
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter(name="minDisk")
    def min_disk(self) -> Optional[int]:
        return pulumi.get(self, "min_disk")

    @property
    @pulumi.getter(name="minRam")
    def min_ram(self) -> Optional[int]:
        return pulumi.get(self, "min_ram")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def ram(self) -> Optional[int]:
        return pulumi.get(self, "ram")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="rxTxFactor")
    def rx_tx_factor(self) -> Optional[float]:
        return pulumi.get(self, "rx_tx_factor")

    @property
    @pulumi.getter
    def swap(self) -> Optional[int]:
        return pulumi.get(self, "swap")

    @property
    @pulumi.getter
    def vcpus(self) -> Optional[int]:
        return pulumi.get(self, "vcpus")


class AwaitableGetFlavorResult(GetFlavorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlavorResult(
            description=self.description,
            disk=self.disk,
            extra_specs=self.extra_specs,
            flavor_id=self.flavor_id,
            id=self.id,
            is_public=self.is_public,
            min_disk=self.min_disk,
            min_ram=self.min_ram,
            name=self.name,
            ram=self.ram,
            region=self.region,
            rx_tx_factor=self.rx_tx_factor,
            swap=self.swap,
            vcpus=self.vcpus)


def get_flavor(description: Optional[str] = None,
               disk: Optional[int] = None,
               flavor_id: Optional[str] = None,
               is_public: Optional[bool] = None,
               min_disk: Optional[int] = None,
               min_ram: Optional[int] = None,
               name: Optional[str] = None,
               ram: Optional[int] = None,
               region: Optional[str] = None,
               rx_tx_factor: Optional[float] = None,
               swap: Optional[int] = None,
               vcpus: Optional[int] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlavorResult:
    """
    Use this data source to get the ID of an available OpenStack flavor.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    small = openstack.compute.get_flavor(ram=512,
        vcpus=1)
    ```


    :param str description: The description of the flavor.
    :param int disk: The exact amount of disk (in gigabytes).
    :param str flavor_id: The ID of the flavor. Conflicts with the `name`,
           `min_ram` and `min_disk`
    :param bool is_public: The flavor visibility.
    :param int min_disk: The minimum amount of disk (in gigabytes). Conflicts
           with the `flavor_id`.
    :param int min_ram: The minimum amount of RAM (in megabytes). Conflicts
           with the `flavor_id`.
    :param str name: The name of the flavor. Conflicts with the `flavor_id`.
    :param int ram: The exact amount of RAM (in megabytes).
    :param str region: The region in which to obtain the V2 Compute client.
           If omitted, the `region` argument of the provider is used.
    :param float rx_tx_factor: The `rx_tx_factor` of the flavor.
    :param int swap: The amount of swap (in gigabytes).
    :param int vcpus: The amount of VCPUs.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['disk'] = disk
    __args__['flavorId'] = flavor_id
    __args__['isPublic'] = is_public
    __args__['minDisk'] = min_disk
    __args__['minRam'] = min_ram
    __args__['name'] = name
    __args__['ram'] = ram
    __args__['region'] = region
    __args__['rxTxFactor'] = rx_tx_factor
    __args__['swap'] = swap
    __args__['vcpus'] = vcpus
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:compute/getFlavor:getFlavor', __args__, opts=opts, typ=GetFlavorResult).value

    return AwaitableGetFlavorResult(
        description=pulumi.get(__ret__, 'description'),
        disk=pulumi.get(__ret__, 'disk'),
        extra_specs=pulumi.get(__ret__, 'extra_specs'),
        flavor_id=pulumi.get(__ret__, 'flavor_id'),
        id=pulumi.get(__ret__, 'id'),
        is_public=pulumi.get(__ret__, 'is_public'),
        min_disk=pulumi.get(__ret__, 'min_disk'),
        min_ram=pulumi.get(__ret__, 'min_ram'),
        name=pulumi.get(__ret__, 'name'),
        ram=pulumi.get(__ret__, 'ram'),
        region=pulumi.get(__ret__, 'region'),
        rx_tx_factor=pulumi.get(__ret__, 'rx_tx_factor'),
        swap=pulumi.get(__ret__, 'swap'),
        vcpus=pulumi.get(__ret__, 'vcpus'))


@_utilities.lift_output_func(get_flavor)
def get_flavor_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                      disk: Optional[pulumi.Input[Optional[int]]] = None,
                      flavor_id: Optional[pulumi.Input[Optional[str]]] = None,
                      is_public: Optional[pulumi.Input[Optional[bool]]] = None,
                      min_disk: Optional[pulumi.Input[Optional[int]]] = None,
                      min_ram: Optional[pulumi.Input[Optional[int]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      ram: Optional[pulumi.Input[Optional[int]]] = None,
                      region: Optional[pulumi.Input[Optional[str]]] = None,
                      rx_tx_factor: Optional[pulumi.Input[Optional[float]]] = None,
                      swap: Optional[pulumi.Input[Optional[int]]] = None,
                      vcpus: Optional[pulumi.Input[Optional[int]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlavorResult]:
    """
    Use this data source to get the ID of an available OpenStack flavor.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    small = openstack.compute.get_flavor(ram=512,
        vcpus=1)
    ```


    :param str description: The description of the flavor.
    :param int disk: The exact amount of disk (in gigabytes).
    :param str flavor_id: The ID of the flavor. Conflicts with the `name`,
           `min_ram` and `min_disk`
    :param bool is_public: The flavor visibility.
    :param int min_disk: The minimum amount of disk (in gigabytes). Conflicts
           with the `flavor_id`.
    :param int min_ram: The minimum amount of RAM (in megabytes). Conflicts
           with the `flavor_id`.
    :param str name: The name of the flavor. Conflicts with the `flavor_id`.
    :param int ram: The exact amount of RAM (in megabytes).
    :param str region: The region in which to obtain the V2 Compute client.
           If omitted, the `region` argument of the provider is used.
    :param float rx_tx_factor: The `rx_tx_factor` of the flavor.
    :param int swap: The amount of swap (in gigabytes).
    :param int vcpus: The amount of VCPUs.
    """
    ...
