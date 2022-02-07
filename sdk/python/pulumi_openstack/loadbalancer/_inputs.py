# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'MembersMemberArgs',
    'PoolPersistenceArgs',
]

@pulumi.input_type
class MembersMemberArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 protocol_port: pulumi.Input[int],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] address: The IP address of the members to receive traffic from
               the load balancer.
        :param pulumi.Input[int] protocol_port: The port on which to listen for client traffic.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               A valid value is true (UP) or false (DOWN). Defaults to true.
        :param pulumi.Input[bool] backup: A bool that indicates whether the member is
               backup. **Requires octavia minor version 2.1 or later**.
        :param pulumi.Input[str] id: The unique ID for the members.
        :param pulumi.Input[str] name: Human-readable name for the member.
        :param pulumi.Input[str] subnet_id: The subnet in which to access the member.
        :param pulumi.Input[int] weight: A positive integer value that indicates the relative
               portion of traffic that this members should receive from the pool. For
               example, a member with a weight of 10 receives five times as much traffic
               as a member with a weight of 2. Defaults to 1.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "protocol_port", protocol_port)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if backup is not None:
            pulumi.set(__self__, "backup", backup)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        The IP address of the members to receive traffic from
        the load balancer.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Input[int]:
        """
        The port on which to listen for client traffic.
        """
        return pulumi.get(self, "protocol_port")

    @protocol_port.setter
    def protocol_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "protocol_port", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the member.
        A valid value is true (UP) or false (DOWN). Defaults to true.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def backup(self) -> Optional[pulumi.Input[bool]]:
        """
        A bool that indicates whether the member is
        backup. **Requires octavia minor version 2.1 or later**.
        """
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID for the members.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable name for the member.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet in which to access the member.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        A positive integer value that indicates the relative
        portion of traffic that this members should receive from the pool. For
        example, a member with a weight of 10 receives five times as much traffic
        as a member with a weight of 2. Defaults to 1.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class PoolPersistenceArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 cookie_name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The type of persistence mode. The current specification
               supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        :param pulumi.Input[str] cookie_name: The name of the cookie if persistence mode is set
               appropriately. Required if `type = APP_COOKIE`.
        """
        pulumi.set(__self__, "type", type)
        if cookie_name is not None:
            pulumi.set(__self__, "cookie_name", cookie_name)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of persistence mode. The current specification
        supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cookie if persistence mode is set
        appropriately. Required if `type = APP_COOKIE`.
        """
        return pulumi.get(self, "cookie_name")

    @cookie_name.setter
    def cookie_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cookie_name", value)


