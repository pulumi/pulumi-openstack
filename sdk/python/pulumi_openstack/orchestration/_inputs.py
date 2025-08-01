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
    'StackV1StackOutputArgs',
    'StackV1StackOutputArgsDict',
]

MYPY = False

if not MYPY:
    class StackV1StackOutputArgsDict(TypedDict):
        output_key: pulumi.Input[_builtins.str]
        output_value: pulumi.Input[_builtins.str]
        description: NotRequired[pulumi.Input[_builtins.str]]
        """
        The description of the stack resource.
        """
elif False:
    StackV1StackOutputArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class StackV1StackOutputArgs:
    def __init__(__self__, *,
                 output_key: pulumi.Input[_builtins.str],
                 output_value: pulumi.Input[_builtins.str],
                 description: Optional[pulumi.Input[_builtins.str]] = None):
        """
        :param pulumi.Input[_builtins.str] description: The description of the stack resource.
        """
        pulumi.set(__self__, "output_key", output_key)
        pulumi.set(__self__, "output_value", output_value)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @_builtins.property
    @pulumi.getter(name="outputKey")
    def output_key(self) -> pulumi.Input[_builtins.str]:
        return pulumi.get(self, "output_key")

    @output_key.setter
    def output_key(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "output_key", value)

    @_builtins.property
    @pulumi.getter(name="outputValue")
    def output_value(self) -> pulumi.Input[_builtins.str]:
        return pulumi.get(self, "output_value")

    @output_value.setter
    def output_value(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "output_value", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The description of the stack resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)


