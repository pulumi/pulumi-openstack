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
    'ShareExportLocationArgs',
    'ShareExportLocationArgsDict',
]

MYPY = False

if not MYPY:
    class ShareExportLocationArgsDict(TypedDict):
        path: NotRequired[pulumi.Input[str]]
        preferred: NotRequired[pulumi.Input[str]]
elif False:
    ShareExportLocationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ShareExportLocationArgs:
    def __init__(__self__, *,
                 path: Optional[pulumi.Input[str]] = None,
                 preferred: Optional[pulumi.Input[str]] = None):
        if path is not None:
            pulumi.set(__self__, "path", path)
        if preferred is not None:
            pulumi.set(__self__, "preferred", preferred)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def preferred(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "preferred")

    @preferred.setter
    def preferred(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred", value)


