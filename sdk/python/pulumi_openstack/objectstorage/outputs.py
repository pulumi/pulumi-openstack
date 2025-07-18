# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
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
    'ContainerVersioningLegacy',
]

@pulumi.output_type
class ContainerVersioningLegacy(dict):
    def __init__(__self__, *,
                 location: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str location: Container in which versions will be stored.
        :param builtins.str type: Versioning type which can be `versions` or `history`
               according to [OpenStack
               documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> builtins.str:
        """
        Container in which versions will be stored.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Versioning type which can be `versions` or `history`
        according to [OpenStack
        documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
        """
        return pulumi.get(self, "type")


