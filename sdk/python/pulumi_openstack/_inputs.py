# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'BgpvpnPortAssociateV2RouteArgs',
]

@pulumi.input_type
class BgpvpnPortAssociateV2RouteArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 bgpvpn_id: Optional[pulumi.Input[str]] = None,
                 local_pref: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: Can be `prefix` or `bgpvpn`. For the `prefix` type, the
               CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
               `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
        :param pulumi.Input[str] bgpvpn_id: The ID of the BGP VPN to be advertised. Required
               if `type` is `bgpvpn`. Conflicts with `prefix`.
        :param pulumi.Input[int] local_pref: The BGP LOCAL\\_PREF value of the routes that will
               be advertised.
        :param pulumi.Input[str] prefix: The CIDR prefix (v4 or v6) to be advertised. Required
               if `type` is `prefix`. Conflicts with `bgpvpn_id`.
        """
        pulumi.set(__self__, "type", type)
        if bgpvpn_id is not None:
            pulumi.set(__self__, "bgpvpn_id", bgpvpn_id)
        if local_pref is not None:
            pulumi.set(__self__, "local_pref", local_pref)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Can be `prefix` or `bgpvpn`. For the `prefix` type, the
        CIDR prefix (v4 or v6) must be specified in the `prefix` key. For the
        `bgpvpn` type, the BGP VPN ID must be specified in the `bgpvpn_id` key.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="bgpvpnId")
    def bgpvpn_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the BGP VPN to be advertised. Required
        if `type` is `bgpvpn`. Conflicts with `prefix`.
        """
        return pulumi.get(self, "bgpvpn_id")

    @bgpvpn_id.setter
    def bgpvpn_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgpvpn_id", value)

    @property
    @pulumi.getter(name="localPref")
    def local_pref(self) -> Optional[pulumi.Input[int]]:
        """
        The BGP LOCAL\\_PREF value of the routes that will
        be advertised.
        """
        return pulumi.get(self, "local_pref")

    @local_pref.setter
    def local_pref(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "local_pref", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR prefix (v4 or v6) to be advertised. Required
        if `type` is `prefix`. Conflicts with `bgpvpn_id`.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)


