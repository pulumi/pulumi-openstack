# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
    'ShareExportLocation',
    'GetShareExportLocationResult',
]

@pulumi.output_type
class ShareExportLocation(dict):
    def __init__(__self__, *,
                 path: Optional[_builtins.str] = None,
                 preferred: Optional[_builtins.str] = None):
        if path is not None:
            pulumi.set(__self__, "path", path)
        if preferred is not None:
            pulumi.set(__self__, "preferred", preferred)

    @_builtins.property
    @pulumi.getter
    def path(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "path")

    @_builtins.property
    @pulumi.getter
    def preferred(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "preferred")


@pulumi.output_type
class GetShareExportLocationResult(dict):
    def __init__(__self__, *,
                 path: _builtins.str,
                 preferred: _builtins.str):
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "preferred", preferred)

    @_builtins.property
    @pulumi.getter
    def path(self) -> _builtins.str:
        return pulumi.get(self, "path")

    @_builtins.property
    @pulumi.getter
    def preferred(self) -> _builtins.str:
        return pulumi.get(self, "preferred")


