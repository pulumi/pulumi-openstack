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
    'StackV1StackOutput',
]

@pulumi.output_type
class StackV1StackOutput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "outputKey":
            suggest = "output_key"
        elif key == "outputValue":
            suggest = "output_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StackV1StackOutput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StackV1StackOutput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StackV1StackOutput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 output_key: _builtins.str,
                 output_value: _builtins.str,
                 description: Optional[_builtins.str] = None):
        """
        :param _builtins.str description: The description of the stack resource.
        """
        pulumi.set(__self__, "output_key", output_key)
        pulumi.set(__self__, "output_value", output_value)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @_builtins.property
    @pulumi.getter(name="outputKey")
    def output_key(self) -> _builtins.str:
        return pulumi.get(self, "output_key")

    @_builtins.property
    @pulumi.getter(name="outputValue")
    def output_value(self) -> _builtins.str:
        return pulumi.get(self, "output_value")

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[_builtins.str]:
        """
        The description of the stack resource.
        """
        return pulumi.get(self, "description")


