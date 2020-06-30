# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetFlavorResult:
    """
    A collection of values returned by getFlavor.
    """
    def __init__(__self__, disk=None, extra_specs=None, flavor_id=None, id=None, is_public=None, min_disk=None, min_ram=None, name=None, ram=None, region=None, rx_tx_factor=None, swap=None, vcpus=None):
        if disk and not isinstance(disk, float):
            raise TypeError("Expected argument 'disk' to be a float")
        __self__.disk = disk
        if extra_specs and not isinstance(extra_specs, dict):
            raise TypeError("Expected argument 'extra_specs' to be a dict")
        __self__.extra_specs = extra_specs
        """
        Key/Value pairs of metadata for the flavor.
        """
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        __self__.flavor_id = flavor_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if is_public and not isinstance(is_public, bool):
            raise TypeError("Expected argument 'is_public' to be a bool")
        __self__.is_public = is_public
        if min_disk and not isinstance(min_disk, float):
            raise TypeError("Expected argument 'min_disk' to be a float")
        __self__.min_disk = min_disk
        if min_ram and not isinstance(min_ram, float):
            raise TypeError("Expected argument 'min_ram' to be a float")
        __self__.min_ram = min_ram
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if ram and not isinstance(ram, float):
            raise TypeError("Expected argument 'ram' to be a float")
        __self__.ram = ram
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if rx_tx_factor and not isinstance(rx_tx_factor, float):
            raise TypeError("Expected argument 'rx_tx_factor' to be a float")
        __self__.rx_tx_factor = rx_tx_factor
        if swap and not isinstance(swap, float):
            raise TypeError("Expected argument 'swap' to be a float")
        __self__.swap = swap
        if vcpus and not isinstance(vcpus, float):
            raise TypeError("Expected argument 'vcpus' to be a float")
        __self__.vcpus = vcpus
class AwaitableGetFlavorResult(GetFlavorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlavorResult(
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

def get_flavor(disk=None,flavor_id=None,is_public=None,min_disk=None,min_ram=None,name=None,ram=None,region=None,rx_tx_factor=None,swap=None,vcpus=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack flavor.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_openstack as openstack

    small = openstack.compute.get_flavor(ram=512,
        vcpus=1)
    ```



    :param float disk: The exact amount of disk (in gigabytes).
    :param str flavor_id: The ID of the flavor. Conflicts with the `name`,
           `min_ram` and `min_disk`
    :param bool is_public: The flavor visibility.
    :param float min_disk: The minimum amount of disk (in gigabytes). Conflicts
           with the `flavor_id`.
    :param float min_ram: The minimum amount of RAM (in megabytes). Conflicts
           with the `flavor_id`.
    :param str name: The name of the flavor. Conflicts with the `flavor_id`.
    :param float ram: The exact amount of RAM (in megabytes).
    :param str region: The region in which to obtain the V2 Compute client.
           If omitted, the `region` argument of the provider is used.
    :param float rx_tx_factor: The `rx_tx_factor` of the flavor.
    :param float swap: The amount of swap (in gigabytes).
    :param float vcpus: The amount of VCPUs.
    """
    __args__ = dict()


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
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:compute/getFlavor:getFlavor', __args__, opts=opts).value

    return AwaitableGetFlavorResult(
        disk=__ret__.get('disk'),
        extra_specs=__ret__.get('extraSpecs'),
        flavor_id=__ret__.get('flavorId'),
        id=__ret__.get('id'),
        is_public=__ret__.get('isPublic'),
        min_disk=__ret__.get('minDisk'),
        min_ram=__ret__.get('minRam'),
        name=__ret__.get('name'),
        ram=__ret__.get('ram'),
        region=__ret__.get('region'),
        rx_tx_factor=__ret__.get('rxTxFactor'),
        swap=__ret__.get('swap'),
        vcpus=__ret__.get('vcpus'))
