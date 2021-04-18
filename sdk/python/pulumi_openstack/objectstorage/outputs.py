# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ContainerVersioning',
]

@pulumi.output_type
class ContainerVersioning(dict):
    def __init__(__self__, *,
                 location: str,
                 type: str):
        """
        :param str location: Container in which versions will be stored.
        :param str type: Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Container in which versions will be stored.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        """
        return pulumi.get(self, "type")


