# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
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
        """
        :param pulumi.Input[int] value: The value for the lifetime of the security association. Must be a positive integer.
               Default is 3600.
        """
        IkePolicyLifetimeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            units=units,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             units: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if units is not None:
            _setter("units", units)
        if value is not None:
            _setter("value", value)

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
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class IpSecPolicyLifetimeArgs:
    def __init__(__self__, *,
                 units: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] value: The value for the lifetime of the security association. Must be a positive integer.
               Default is 3600.
        """
        IpSecPolicyLifetimeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            units=units,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             units: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if units is not None:
            _setter("units", units)
        if value is not None:
            _setter("value", value)

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
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
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
        """
        :param pulumi.Input[str] action: The dead peer detection (DPD) action.
               A valid value is clear, hold, restart, disabled, or restart-by-peer.
               Default value is hold.
        :param pulumi.Input[int] interval: The dead peer detection (DPD) interval, in seconds.
               A valid value is a positive integer.
               Default is 30.
        :param pulumi.Input[int] timeout: The dead peer detection (DPD) timeout in seconds.
               A valid value is a positive integer that is greater than the DPD interval value.
               Default is 120.
        """
        SiteConnectionDpdArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            interval=interval,
            timeout=timeout,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional[pulumi.Input[str]] = None,
             interval: Optional[pulumi.Input[int]] = None,
             timeout: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if action is not None:
            _setter("action", action)
        if interval is not None:
            _setter("interval", interval)
        if timeout is not None:
            _setter("timeout", timeout)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The dead peer detection (DPD) action.
        A valid value is clear, hold, restart, disabled, or restart-by-peer.
        Default value is hold.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The dead peer detection (DPD) interval, in seconds.
        A valid value is a positive integer.
        Default is 30.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The dead peer detection (DPD) timeout in seconds.
        A valid value is a positive integer that is greater than the DPD interval value.
        Default is 120.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


