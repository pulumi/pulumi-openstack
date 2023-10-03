# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MemberArgs', 'Member']

@pulumi.input_type
class MemberArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 pool_id: pulumi.Input[str],
                 protocol_port: pulumi.Input[int],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 monitor_address: Optional[pulumi.Input[str]] = None,
                 monitor_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Member resource.
        :param pulumi.Input[str] address: The IP address of the member to receive traffic from
               the load balancer. Changing this creates a new member.
        :param pulumi.Input[str] pool_id: The id of the pool that this member will be assigned
               to. Changing this creates a new member.
        :param pulumi.Input[int] protocol_port: The port on which to listen for client traffic.
               Changing this creates a new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               A valid value is true (UP) or false (DOWN). Defaults to true.
        :param pulumi.Input[bool] backup: Boolean that indicates whether that member works as a backup or not. Available 
               only for Octavia >= 2.1.
        :param pulumi.Input[str] monitor_address: An alternate IP address used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[int] monitor_port: An alternate protocol port used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[str] name: Human-readable name for the member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a member. If omitted, the `region`
               argument of the provider is used. Changing this creates a new member.
        :param pulumi.Input[str] subnet_id: The subnet in which to access the member. Changing
               this creates a new member.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the member.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new member.
        :param pulumi.Input[int] weight: A positive integer value that indicates the relative
               portion of traffic that this member should receive from the pool. For
               example, a member with a weight of 10 receives five times as much traffic
               as a member with a weight of 2. Defaults to 1.
        """
        MemberArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address=address,
            pool_id=pool_id,
            protocol_port=protocol_port,
            admin_state_up=admin_state_up,
            backup=backup,
            monitor_address=monitor_address,
            monitor_port=monitor_port,
            name=name,
            region=region,
            subnet_id=subnet_id,
            tenant_id=tenant_id,
            weight=weight,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address: pulumi.Input[str],
             pool_id: pulumi.Input[str],
             protocol_port: pulumi.Input[int],
             admin_state_up: Optional[pulumi.Input[bool]] = None,
             backup: Optional[pulumi.Input[bool]] = None,
             monitor_address: Optional[pulumi.Input[str]] = None,
             monitor_port: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             subnet_id: Optional[pulumi.Input[str]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             weight: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("address", address)
        _setter("pool_id", pool_id)
        _setter("protocol_port", protocol_port)
        if admin_state_up is not None:
            _setter("admin_state_up", admin_state_up)
        if backup is not None:
            _setter("backup", backup)
        if monitor_address is not None:
            _setter("monitor_address", monitor_address)
        if monitor_port is not None:
            _setter("monitor_port", monitor_port)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)
        if subnet_id is not None:
            _setter("subnet_id", subnet_id)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if weight is not None:
            _setter("weight", weight)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        The IP address of the member to receive traffic from
        the load balancer. Changing this creates a new member.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Input[str]:
        """
        The id of the pool that this member will be assigned
        to. Changing this creates a new member.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Input[int]:
        """
        The port on which to listen for client traffic.
        Changing this creates a new member.
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
        Boolean that indicates whether that member works as a backup or not. Available 
        only for Octavia >= 2.1.
        """
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter(name="monitorAddress")
    def monitor_address(self) -> Optional[pulumi.Input[str]]:
        """
        An alternate IP address used for health monitoring a backend member.
        Available only for Octavia
        """
        return pulumi.get(self, "monitor_address")

    @monitor_address.setter
    def monitor_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor_address", value)

    @property
    @pulumi.getter(name="monitorPort")
    def monitor_port(self) -> Optional[pulumi.Input[int]]:
        """
        An alternate protocol port used for health monitoring a backend member.
        Available only for Octavia
        """
        return pulumi.get(self, "monitor_port")

    @monitor_port.setter
    def monitor_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "monitor_port", value)

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
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a member. If omitted, the `region`
        argument of the provider is used. Changing this creates a new member.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet in which to access the member. Changing
        this creates a new member.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required for admins. The UUID of the tenant who owns
        the member.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        A positive integer value that indicates the relative
        portion of traffic that this member should receive from the pool. For
        example, a member with a weight of 10 receives five times as much traffic
        as a member with a weight of 2. Defaults to 1.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _MemberState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 monitor_address: Optional[pulumi.Input[str]] = None,
                 monitor_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 protocol_port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Member resources.
        :param pulumi.Input[str] address: The IP address of the member to receive traffic from
               the load balancer. Changing this creates a new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               A valid value is true (UP) or false (DOWN). Defaults to true.
        :param pulumi.Input[bool] backup: Boolean that indicates whether that member works as a backup or not. Available 
               only for Octavia >= 2.1.
        :param pulumi.Input[str] monitor_address: An alternate IP address used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[int] monitor_port: An alternate protocol port used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[str] name: Human-readable name for the member.
        :param pulumi.Input[str] pool_id: The id of the pool that this member will be assigned
               to. Changing this creates a new member.
        :param pulumi.Input[int] protocol_port: The port on which to listen for client traffic.
               Changing this creates a new member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a member. If omitted, the `region`
               argument of the provider is used. Changing this creates a new member.
        :param pulumi.Input[str] subnet_id: The subnet in which to access the member. Changing
               this creates a new member.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the member.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new member.
        :param pulumi.Input[int] weight: A positive integer value that indicates the relative
               portion of traffic that this member should receive from the pool. For
               example, a member with a weight of 10 receives five times as much traffic
               as a member with a weight of 2. Defaults to 1.
        """
        _MemberState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address=address,
            admin_state_up=admin_state_up,
            backup=backup,
            monitor_address=monitor_address,
            monitor_port=monitor_port,
            name=name,
            pool_id=pool_id,
            protocol_port=protocol_port,
            region=region,
            subnet_id=subnet_id,
            tenant_id=tenant_id,
            weight=weight,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address: Optional[pulumi.Input[str]] = None,
             admin_state_up: Optional[pulumi.Input[bool]] = None,
             backup: Optional[pulumi.Input[bool]] = None,
             monitor_address: Optional[pulumi.Input[str]] = None,
             monitor_port: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             pool_id: Optional[pulumi.Input[str]] = None,
             protocol_port: Optional[pulumi.Input[int]] = None,
             region: Optional[pulumi.Input[str]] = None,
             subnet_id: Optional[pulumi.Input[str]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             weight: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if address is not None:
            _setter("address", address)
        if admin_state_up is not None:
            _setter("admin_state_up", admin_state_up)
        if backup is not None:
            _setter("backup", backup)
        if monitor_address is not None:
            _setter("monitor_address", monitor_address)
        if monitor_port is not None:
            _setter("monitor_port", monitor_port)
        if name is not None:
            _setter("name", name)
        if pool_id is not None:
            _setter("pool_id", pool_id)
        if protocol_port is not None:
            _setter("protocol_port", protocol_port)
        if region is not None:
            _setter("region", region)
        if subnet_id is not None:
            _setter("subnet_id", subnet_id)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if weight is not None:
            _setter("weight", weight)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the member to receive traffic from
        the load balancer. Changing this creates a new member.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

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
        Boolean that indicates whether that member works as a backup or not. Available 
        only for Octavia >= 2.1.
        """
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter(name="monitorAddress")
    def monitor_address(self) -> Optional[pulumi.Input[str]]:
        """
        An alternate IP address used for health monitoring a backend member.
        Available only for Octavia
        """
        return pulumi.get(self, "monitor_address")

    @monitor_address.setter
    def monitor_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor_address", value)

    @property
    @pulumi.getter(name="monitorPort")
    def monitor_port(self) -> Optional[pulumi.Input[int]]:
        """
        An alternate protocol port used for health monitoring a backend member.
        Available only for Octavia
        """
        return pulumi.get(self, "monitor_port")

    @monitor_port.setter
    def monitor_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "monitor_port", value)

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
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the pool that this member will be assigned
        to. Changing this creates a new member.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> Optional[pulumi.Input[int]]:
        """
        The port on which to listen for client traffic.
        Changing this creates a new member.
        """
        return pulumi.get(self, "protocol_port")

    @protocol_port.setter
    def protocol_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "protocol_port", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a member. If omitted, the `region`
        argument of the provider is used. Changing this creates a new member.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet in which to access the member. Changing
        this creates a new member.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required for admins. The UUID of the tenant who owns
        the member.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        A positive integer value that indicates the relative
        portion of traffic that this member should receive from the pool. For
        example, a member with a weight of 10 receives five times as much traffic
        as a member with a weight of 2. Defaults to 1.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class Member(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 monitor_address: Optional[pulumi.Input[str]] = None,
                 monitor_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 protocol_port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a V2 member resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        member1 = openstack.loadbalancer.Member("member1",
            address="192.168.199.23",
            pool_id="935685fb-a896-40f9-9ff4-ae531a3a00fe",
            protocol_port=8080)
        ```

        ## Import

        Load Balancer Pool Member can be imported using the Pool ID and Member ID separated by a slash, e.g.:

        ```sh
         $ pulumi import openstack:loadbalancer/member:Member member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address of the member to receive traffic from
               the load balancer. Changing this creates a new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               A valid value is true (UP) or false (DOWN). Defaults to true.
        :param pulumi.Input[bool] backup: Boolean that indicates whether that member works as a backup or not. Available 
               only for Octavia >= 2.1.
        :param pulumi.Input[str] monitor_address: An alternate IP address used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[int] monitor_port: An alternate protocol port used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[str] name: Human-readable name for the member.
        :param pulumi.Input[str] pool_id: The id of the pool that this member will be assigned
               to. Changing this creates a new member.
        :param pulumi.Input[int] protocol_port: The port on which to listen for client traffic.
               Changing this creates a new member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a member. If omitted, the `region`
               argument of the provider is used. Changing this creates a new member.
        :param pulumi.Input[str] subnet_id: The subnet in which to access the member. Changing
               this creates a new member.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the member.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new member.
        :param pulumi.Input[int] weight: A positive integer value that indicates the relative
               portion of traffic that this member should receive from the pool. For
               example, a member with a weight of 10 receives five times as much traffic
               as a member with a weight of 2. Defaults to 1.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 member resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        member1 = openstack.loadbalancer.Member("member1",
            address="192.168.199.23",
            pool_id="935685fb-a896-40f9-9ff4-ae531a3a00fe",
            protocol_port=8080)
        ```

        ## Import

        Load Balancer Pool Member can be imported using the Pool ID and Member ID separated by a slash, e.g.:

        ```sh
         $ pulumi import openstack:loadbalancer/member:Member member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f
        ```

        :param str resource_name: The name of the resource.
        :param MemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            MemberArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 monitor_address: Optional[pulumi.Input[str]] = None,
                 monitor_port: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 protocol_port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MemberArgs.__new__(MemberArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["backup"] = backup
            __props__.__dict__["monitor_address"] = monitor_address
            __props__.__dict__["monitor_port"] = monitor_port
            __props__.__dict__["name"] = name
            if pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'pool_id'")
            __props__.__dict__["pool_id"] = pool_id
            if protocol_port is None and not opts.urn:
                raise TypeError("Missing required property 'protocol_port'")
            __props__.__dict__["protocol_port"] = protocol_port
            __props__.__dict__["region"] = region
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["weight"] = weight
        super(Member, __self__).__init__(
            'openstack:loadbalancer/member:Member',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            backup: Optional[pulumi.Input[bool]] = None,
            monitor_address: Optional[pulumi.Input[str]] = None,
            monitor_port: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pool_id: Optional[pulumi.Input[str]] = None,
            protocol_port: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[int]] = None) -> 'Member':
        """
        Get an existing Member resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address of the member to receive traffic from
               the load balancer. Changing this creates a new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               A valid value is true (UP) or false (DOWN). Defaults to true.
        :param pulumi.Input[bool] backup: Boolean that indicates whether that member works as a backup or not. Available 
               only for Octavia >= 2.1.
        :param pulumi.Input[str] monitor_address: An alternate IP address used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[int] monitor_port: An alternate protocol port used for health monitoring a backend member.
               Available only for Octavia
        :param pulumi.Input[str] name: Human-readable name for the member.
        :param pulumi.Input[str] pool_id: The id of the pool that this member will be assigned
               to. Changing this creates a new member.
        :param pulumi.Input[int] protocol_port: The port on which to listen for client traffic.
               Changing this creates a new member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a member. If omitted, the `region`
               argument of the provider is used. Changing this creates a new member.
        :param pulumi.Input[str] subnet_id: The subnet in which to access the member. Changing
               this creates a new member.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the member.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new member.
        :param pulumi.Input[int] weight: A positive integer value that indicates the relative
               portion of traffic that this member should receive from the pool. For
               example, a member with a weight of 10 receives five times as much traffic
               as a member with a weight of 2. Defaults to 1.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MemberState.__new__(_MemberState)

        __props__.__dict__["address"] = address
        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["backup"] = backup
        __props__.__dict__["monitor_address"] = monitor_address
        __props__.__dict__["monitor_port"] = monitor_port
        __props__.__dict__["name"] = name
        __props__.__dict__["pool_id"] = pool_id
        __props__.__dict__["protocol_port"] = protocol_port
        __props__.__dict__["region"] = region
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["weight"] = weight
        return Member(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The IP address of the member to receive traffic from
        the load balancer. Changing this creates a new member.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        """
        The administrative state of the member.
        A valid value is true (UP) or false (DOWN). Defaults to true.
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def backup(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean that indicates whether that member works as a backup or not. Available 
        only for Octavia >= 2.1.
        """
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter(name="monitorAddress")
    def monitor_address(self) -> pulumi.Output[Optional[str]]:
        """
        An alternate IP address used for health monitoring a backend member.
        Available only for Octavia
        """
        return pulumi.get(self, "monitor_address")

    @property
    @pulumi.getter(name="monitorPort")
    def monitor_port(self) -> pulumi.Output[Optional[int]]:
        """
        An alternate protocol port used for health monitoring a backend member.
        Available only for Octavia
        """
        return pulumi.get(self, "monitor_port")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Human-readable name for the member.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Output[str]:
        """
        The id of the pool that this member will be assigned
        to. Changing this creates a new member.
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Output[int]:
        """
        The port on which to listen for client traffic.
        Changing this creates a new member.
        """
        return pulumi.get(self, "protocol_port")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a member. If omitted, the `region`
        argument of the provider is used. Changing this creates a new member.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        The subnet in which to access the member. Changing
        this creates a new member.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Required for admins. The UUID of the tenant who owns
        the member.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[int]:
        """
        A positive integer value that indicates the relative
        portion of traffic that this member should receive from the pool. For
        example, a member with a weight of 10 receives five times as much traffic
        as a member with a weight of 2. Defaults to 1.
        """
        return pulumi.get(self, "weight")

