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
    'IkePolicyLifetime',
    'IpSecPolicyLifetime',
    'SiteConnectionDpd',
]

@pulumi.output_type
class IkePolicyLifetime(dict):
    def __init__(__self__, *,
                 units: Optional[str] = None,
                 value: Optional[int] = None):
        """
        :param int value: The value for the lifetime of the security association. Must be a positive integer.
               Default is 3600.
        """
        IkePolicyLifetime._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            units=units,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             units: Optional[str] = None,
             value: Optional[int] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if units is not None:
            _setter("units", units)
        if value is not None:
            _setter("value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[str]:
        return pulumi.get(self, "units")

    @property
    @pulumi.getter
    def value(self) -> Optional[int]:
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class IpSecPolicyLifetime(dict):
    def __init__(__self__, *,
                 units: Optional[str] = None,
                 value: Optional[int] = None):
        """
        :param int value: The value for the lifetime of the security association. Must be a positive integer.
               Default is 3600.
        """
        IpSecPolicyLifetime._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            units=units,
            value=value,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             units: Optional[str] = None,
             value: Optional[int] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if units is not None:
            _setter("units", units)
        if value is not None:
            _setter("value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[str]:
        return pulumi.get(self, "units")

    @property
    @pulumi.getter
    def value(self) -> Optional[int]:
        """
        The value for the lifetime of the security association. Must be a positive integer.
        Default is 3600.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class SiteConnectionDpd(dict):
    def __init__(__self__, *,
                 action: Optional[str] = None,
                 interval: Optional[int] = None,
                 timeout: Optional[int] = None):
        """
        :param str action: The dead peer detection (DPD) action.
               A valid value is clear, hold, restart, disabled, or restart-by-peer.
               Default value is hold.
        :param int interval: The dead peer detection (DPD) interval, in seconds.
               A valid value is a positive integer.
               Default is 30.
        :param int timeout: The dead peer detection (DPD) timeout in seconds.
               A valid value is a positive integer that is greater than the DPD interval value.
               Default is 120.
        """
        SiteConnectionDpd._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            interval=interval,
            timeout=timeout,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional[str] = None,
             interval: Optional[int] = None,
             timeout: Optional[int] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if action is not None:
            _setter("action", action)
        if interval is not None:
            _setter("interval", interval)
        if timeout is not None:
            _setter("timeout", timeout)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        """
        The dead peer detection (DPD) action.
        A valid value is clear, hold, restart, disabled, or restart-by-peer.
        Default value is hold.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        """
        The dead peer detection (DPD) interval, in seconds.
        A valid value is a positive integer.
        Default is 30.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        """
        The dead peer detection (DPD) timeout in seconds.
        A valid value is a positive integer that is greater than the DPD interval value.
        Default is 120.
        """
        return pulumi.get(self, "timeout")


