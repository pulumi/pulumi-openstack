# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'ShareExportLocation',
    'GetShareExportLocationResult',
]

@pulumi.output_type
class ShareExportLocation(dict):
    def __init__(__self__, *,
                 path: Optional[str] = None,
                 preferred: Optional[str] = None):
        if path is not None:
            pulumi.set(__self__, "path", path)
        if preferred is not None:
            pulumi.set(__self__, "preferred", preferred)

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def preferred(self) -> Optional[str]:
        return pulumi.get(self, "preferred")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetShareExportLocationResult(dict):
    def __init__(__self__, *,
                 path: str,
                 preferred: str):
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "preferred", preferred)

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def preferred(self) -> str:
        return pulumi.get(self, "preferred")


