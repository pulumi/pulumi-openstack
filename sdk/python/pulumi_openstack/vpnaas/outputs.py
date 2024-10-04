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
    'IkePolicyLifetime',
    'IpSecPolicyLifetime',
    'SiteConnectionDpd',
]

@pulumi.output_type
class IkePolicyLifetime(dict):
    def __init__(__self__, *,
                 units: Optional[str] = None,
                 value: Optional[int] = None):
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[str]:
        return pulumi.get(self, "units")

    @property
    @pulumi.getter
    def value(self) -> Optional[int]:
        return pulumi.get(self, "value")


@pulumi.output_type
class IpSecPolicyLifetime(dict):
    def __init__(__self__, *,
                 units: Optional[str] = None,
                 value: Optional[int] = None):
        if units is not None:
            pulumi.set(__self__, "units", units)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def units(self) -> Optional[str]:
        return pulumi.get(self, "units")

    @property
    @pulumi.getter
    def value(self) -> Optional[int]:
        return pulumi.get(self, "value")


@pulumi.output_type
class SiteConnectionDpd(dict):
    def __init__(__self__, *,
                 action: Optional[str] = None,
                 interval: Optional[int] = None,
                 timeout: Optional[int] = None):
        if action is not None:
            pulumi.set(__self__, "action", action)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def action(self) -> Optional[str]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        return pulumi.get(self, "timeout")


