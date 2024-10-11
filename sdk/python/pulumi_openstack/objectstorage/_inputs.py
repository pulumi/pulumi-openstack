# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'ContainerVersioningLegacyArgs',
    'ContainerVersioningLegacyArgsDict',
]

MYPY = False

if not MYPY:
    class ContainerVersioningLegacyArgsDict(TypedDict):
        location: pulumi.Input[str]
        """
        Container in which versions will be stored.
        """
        type: pulumi.Input[str]
        """
        Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        """
elif False:
    ContainerVersioningLegacyArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ContainerVersioningLegacyArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] location: Container in which versions will be stored.
        :param pulumi.Input[str] type: Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        Container in which versions will be stored.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


