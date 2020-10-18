# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'StackV1Output',
]

@pulumi.output_type
class StackV1Output(dict):
    def __init__(__self__, *,
                 output_key: str,
                 output_value: str,
                 description: Optional[str] = None):
        """
        :param str description: The description of the stack resource.
        """
        pulumi.set(__self__, "output_key", output_key)
        pulumi.set(__self__, "output_value", output_value)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="outputKey")
    def output_key(self) -> str:
        return pulumi.get(self, "output_key")

    @property
    @pulumi.getter(name="outputValue")
    def output_value(self) -> str:
        return pulumi.get(self, "output_value")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the stack resource.
        """
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


