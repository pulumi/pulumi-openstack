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
    'MembersMemberArgs',
    'MembersMemberArgsDict',
    'PoolPersistenceArgs',
    'PoolPersistenceArgsDict',
]

MYPY = False

if not MYPY:
    class MembersMemberArgsDict(TypedDict):
        address: pulumi.Input[_builtins.str]
        """
        The IP address of the members to receive traffic from
        the load balancer.
        """
        protocol_port: pulumi.Input[_builtins.int]
        """
        The port on which to listen for client traffic.
        """
        admin_state_up: NotRequired[pulumi.Input[_builtins.bool]]
        """
        The administrative state of the member.
        A valid value is true (UP) or false (DOWN). Defaults to true.
        """
        backup: NotRequired[pulumi.Input[_builtins.bool]]
        """
        A bool that indicates whether the member is
        backup. **Requires octavia minor version 2.1 or later**.
        """
        id: NotRequired[pulumi.Input[_builtins.str]]
        """
        The unique ID for the members.
        """
        monitor_address: NotRequired[pulumi.Input[_builtins.str]]
        """
        An alternate IP address used for health 
        monitoring a backend member.
        """
        monitor_port: NotRequired[pulumi.Input[_builtins.int]]
        """
        An alternate protocol port used for health 
        monitoring a backend member.
        """
        name: NotRequired[pulumi.Input[_builtins.str]]
        """
        Human-readable name for the member.
        """
        subnet_id: NotRequired[pulumi.Input[_builtins.str]]
        """
        The subnet in which to access the member.
        """
        weight: NotRequired[pulumi.Input[_builtins.int]]
        """
        A positive integer value that indicates the relative
        portion of traffic that this members should receive from the pool. For
        example, a member with a weight of 10 receives five times as much traffic
        as a member with a weight of 2. Defaults to 1.
        """
elif False:
    MembersMemberArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MembersMemberArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[_builtins.str],
                 protocol_port: pulumi.Input[_builtins.int],
                 admin_state_up: Optional[pulumi.Input[_builtins.bool]] = None,
                 backup: Optional[pulumi.Input[_builtins.bool]] = None,
                 id: Optional[pulumi.Input[_builtins.str]] = None,
                 monitor_address: Optional[pulumi.Input[_builtins.str]] = None,
                 monitor_port: Optional[pulumi.Input[_builtins.int]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 subnet_id: Optional[pulumi.Input[_builtins.str]] = None,
                 weight: Optional[pulumi.Input[_builtins.int]] = None):
        """
        :param pulumi.Input[_builtins.str] address: The IP address of the members to receive traffic from
               the load balancer.
        :param pulumi.Input[_builtins.int] protocol_port: The port on which to listen for client traffic.
        :param pulumi.Input[_builtins.bool] admin_state_up: The administrative state of the member.
               A valid value is true (UP) or false (DOWN). Defaults to true.
        :param pulumi.Input[_builtins.bool] backup: A bool that indicates whether the member is
               backup. **Requires octavia minor version 2.1 or later**.
        :param pulumi.Input[_builtins.str] id: The unique ID for the members.
        :param pulumi.Input[_builtins.str] monitor_address: An alternate IP address used for health 
               monitoring a backend member.
        :param pulumi.Input[_builtins.int] monitor_port: An alternate protocol port used for health 
               monitoring a backend member.
        :param pulumi.Input[_builtins.str] name: Human-readable name for the member.
        :param pulumi.Input[_builtins.str] subnet_id: The subnet in which to access the member.
        :param pulumi.Input[_builtins.int] weight: A positive integer value that indicates the relative
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
        if monitor_address is not None:
            pulumi.set(__self__, "monitor_address", monitor_address)
        if monitor_port is not None:
            pulumi.set(__self__, "monitor_port", monitor_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @_builtins.property
    @pulumi.getter
    def address(self) -> pulumi.Input[_builtins.str]:
        """
        The IP address of the members to receive traffic from
        the load balancer.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "address", value)

    @_builtins.property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Input[_builtins.int]:
        """
        The port on which to listen for client traffic.
        """
        return pulumi.get(self, "protocol_port")

    @protocol_port.setter
    def protocol_port(self, value: pulumi.Input[_builtins.int]):
        pulumi.set(self, "protocol_port", value)

    @_builtins.property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        The administrative state of the member.
        A valid value is true (UP) or false (DOWN). Defaults to true.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "admin_state_up", value)

    @_builtins.property
    @pulumi.getter
    def backup(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        A bool that indicates whether the member is
        backup. **Requires octavia minor version 2.1 or later**.
        """
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "backup", value)

    @_builtins.property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The unique ID for the members.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "id", value)

    @_builtins.property
    @pulumi.getter(name="monitorAddress")
    def monitor_address(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        An alternate IP address used for health 
        monitoring a backend member.
        """
        return pulumi.get(self, "monitor_address")

    @monitor_address.setter
    def monitor_address(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "monitor_address", value)

    @_builtins.property
    @pulumi.getter(name="monitorPort")
    def monitor_port(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        An alternate protocol port used for health 
        monitoring a backend member.
        """
        return pulumi.get(self, "monitor_port")

    @monitor_port.setter
    def monitor_port(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "monitor_port", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Human-readable name for the member.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The subnet in which to access the member.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "subnet_id", value)

    @_builtins.property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        A positive integer value that indicates the relative
        portion of traffic that this members should receive from the pool. For
        example, a member with a weight of 10 receives five times as much traffic
        as a member with a weight of 2. Defaults to 1.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "weight", value)


if not MYPY:
    class PoolPersistenceArgsDict(TypedDict):
        type: pulumi.Input[_builtins.str]
        """
        The type of persistence mode. The current specification
        supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        """
        cookie_name: NotRequired[pulumi.Input[_builtins.str]]
        """
        The name of the cookie if persistence mode is set
        appropriately. Required if `type = APP_COOKIE`.
        """
elif False:
    PoolPersistenceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class PoolPersistenceArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[_builtins.str],
                 cookie_name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        :param pulumi.Input[_builtins.str] type: The type of persistence mode. The current specification
               supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        :param pulumi.Input[_builtins.str] cookie_name: The name of the cookie if persistence mode is set
               appropriately. Required if `type = APP_COOKIE`.
        """
        pulumi.set(__self__, "type", type)
        if cookie_name is not None:
            pulumi.set(__self__, "cookie_name", cookie_name)

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Input[_builtins.str]:
        """
        The type of persistence mode. The current specification
        supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the cookie if persistence mode is set
        appropriately. Required if `type = APP_COOKIE`.
        """
        return pulumi.get(self, "cookie_name")

    @cookie_name.setter
    def cookie_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "cookie_name", value)


