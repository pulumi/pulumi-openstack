# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'IkePolicyLifetimeArgs',
    'IpSecPolicyLifetimeArgs',
    'SiteConnectionDpdArgs',
]

@pulumi.input_type
class IkePolicyLifetimeArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[int]] = None):
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "units")

    @units.setter
    def units(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "units", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class IpSecPolicyLifetimeArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[int]] = None):
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "units")

    @units.setter
    def units(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "units", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class SiteConnectionDpdArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        if action is not None:
            pulumi.set(__self__, "action", action)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


