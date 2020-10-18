# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'MembersMember',
    'PoolPersistence',
]

@pulumi.output_type
class MembersMember(dict):
    def __init__(__self__, *,
                 address: str,
                 protocol_port: int,
                 admin_state_up: Optional[bool] = None,
                 id: Optional[str] = None,
                 name: Optional[str] = None,
                 subnet_id: Optional[str] = None,
                 weight: Optional[int] = None):
        """
        :param str address: The IP address of the members to receive traffic from
               the load balancer.
        :param int protocol_port: The port on which to listen for client traffic.
        :param bool admin_state_up: The administrative state of the member.
               A valid value is true (UP) or false (DOWN). Defaults to true.
        :param str id: The unique ID for the members.
        :param str name: Human-readable name for the member.
        :param str subnet_id: The subnet in which to access the member.
        :param int weight: A positive integer value that indicates the relative
               portion of traffic that this members should receive from the pool. For
               example, a member with a weight of 10 receives five times as much traffic
               as a member with a weight of 2. Defaults to 1.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "protocol_port", protocol_port)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
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
    def address(self) -> str:
        """
        The IP address of the members to receive traffic from
        the load balancer.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> int:
        """
        The port on which to listen for client traffic.
        """
        return pulumi.get(self, "protocol_port")

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[bool]:
        """
        The administrative state of the member.
        A valid value is true (UP) or false (DOWN). Defaults to true.
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The unique ID for the members.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Human-readable name for the member.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The subnet in which to access the member.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def weight(self) -> Optional[int]:
        """
        A positive integer value that indicates the relative
        portion of traffic that this members should receive from the pool. For
        example, a member with a weight of 10 receives five times as much traffic
        as a member with a weight of 2. Defaults to 1.
        """
        return pulumi.get(self, "weight")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PoolPersistence(dict):
    def __init__(__self__, *,
                 type: str,
                 cookie_name: Optional[str] = None):
        """
        :param str type: The type of persistence mode. The current specification
               supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        :param str cookie_name: The name of the cookie if persistence mode is set
               appropriately. Required if `type = APP_COOKIE`.
        """
        pulumi.set(__self__, "type", type)
        if cookie_name is not None:
            pulumi.set(__self__, "cookie_name", cookie_name)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of persistence mode. The current specification
        supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> Optional[str]:
        """
        The name of the cookie if persistence mode is set
        appropriately. Required if `type = APP_COOKIE`.
        """
        return pulumi.get(self, "cookie_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


