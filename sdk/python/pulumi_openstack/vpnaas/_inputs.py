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
    'IkePolicyLifetimeArgs',
    'IkePolicyLifetimeArgsDict',
    'IpSecPolicyLifetimeArgs',
    'IpSecPolicyLifetimeArgsDict',
    'SiteConnectionDpdArgs',
    'SiteConnectionDpdArgsDict',
]

MYPY = False

if not MYPY:
    class IkePolicyLifetimeArgsDict(TypedDict):
        units: NotRequired[pulumi.Input[_builtins.str]]
        value: NotRequired[pulumi.Input[_builtins.int]]
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
elif False:
    IkePolicyLifetimeArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IkePolicyLifetimeArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input[_builtins.str]] = None,
                 value: Optional[pulumi.Input[_builtins.int]] = None):
        """
        :param pulumi.Input[_builtins.int] value: The value for the lifetime of the security association. Must be a positive integer.
               Default is 3600.
        """
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @_builtins.property
    @pulumi.getter
    def units(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "units")

    @units.setter
    def units(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "units", value)

    @_builtins.property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "value", value)


if not MYPY:
    class IpSecPolicyLifetimeArgsDict(TypedDict):
        units: NotRequired[pulumi.Input[_builtins.str]]
        value: NotRequired[pulumi.Input[_builtins.int]]
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
elif False:
    IpSecPolicyLifetimeArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpSecPolicyLifetimeArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input[_builtins.str]] = None,
                 value: Optional[pulumi.Input[_builtins.int]] = None):
        """
        :param pulumi.Input[_builtins.int] value: The value for the lifetime of the security association. Must be a positive integer.
               Default is 3600.
        """
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @_builtins.property
    @pulumi.getter
    def units(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "units")

    @units.setter
    def units(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "units", value)

    @_builtins.property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "value", value)


if not MYPY:
    class SiteConnectionDpdArgsDict(TypedDict):
        action: NotRequired[pulumi.Input[_builtins.str]]
        """
        The dead peer detection (DPD) action.
        A valid value is clear, hold, restart, disabled, or restart-by-peer.
        Default value is hold.
        """
        interval: NotRequired[pulumi.Input[_builtins.int]]
        """
        The dead peer detection (DPD) interval, in seconds.
        A valid value is a positive integer.
        Default is 30.
        """
        timeout: NotRequired[pulumi.Input[_builtins.int]]
        """
        The dead peer detection (DPD) timeout in seconds.
        A valid value is a positive integer that is greater than the DPD interval value.
        Default is 120.
        """
elif False:
    SiteConnectionDpdArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SiteConnectionDpdArgs:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[_builtins.str]] = None,
                 interval: Optional[pulumi.Input[_builtins.int]] = None,
                 timeout: Optional[pulumi.Input[_builtins.int]] = None):
        """
        :param pulumi.Input[_builtins.str] action: The dead peer detection (DPD) action.
               A valid value is clear, hold, restart, disabled, or restart-by-peer.
               Default value is hold.
        :param pulumi.Input[_builtins.int] interval: The dead peer detection (DPD) interval, in seconds.
               A valid value is a positive integer.
               Default is 30.
        :param pulumi.Input[_builtins.int] timeout: The dead peer detection (DPD) timeout in seconds.
               A valid value is a positive integer that is greater than the DPD interval value.
               Default is 120.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @_builtins.property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The dead peer detection (DPD) action.
        A valid value is clear, hold, restart, disabled, or restart-by-peer.
        Default value is hold.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "action", value)

    @_builtins.property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The dead peer detection (DPD) interval, in seconds.
        A valid value is a positive integer.
        Default is 30.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "interval", value)

    @_builtins.property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The dead peer detection (DPD) timeout in seconds.
        A valid value is a positive integer that is greater than the DPD interval value.
        Default is 120.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "timeout", value)


