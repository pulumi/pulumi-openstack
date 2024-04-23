# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MemberV1Args', 'MemberV1']

@pulumi.input_type
class MemberV1Args:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 pool_id: pulumi.Input[str],
                 port: pulumi.Input[int],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a MemberV1 resource.
        :param pulumi.Input[str] address: The IP address of the member. Changing this creates a
               new member.
        :param pulumi.Input[str] pool_id: The ID of the LB pool. Changing this creates a new
               member.
        :param pulumi.Input[int] port: An integer representing the port on which the member is
               hosted. Changing this creates a new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               Acceptable values are 'true' and 'false'. Changing this value updates the
               state of the existing member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB member. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB member.
        :param pulumi.Input[str] tenant_id: The owner of the member. Required if admin wants to
               create a member for another tenant. Changing this creates a new member.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "pool_id", pool_id)
        pulumi.set(__self__, "port", port)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        The IP address of the member. Changing this creates a
        new member.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Input[str]:
        """
        The ID of the LB pool. Changing this creates a new
        member.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        An integer representing the port on which the member is
        hosted. Changing this creates a new member.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the member.
        Acceptable values are 'true' and 'false'. Changing this value updates the
        state of the existing member.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB member. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB member.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the member. Required if admin wants to
        create a member for another tenant. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _MemberV1State:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering MemberV1 resources.
        :param pulumi.Input[str] address: The IP address of the member. Changing this creates a
               new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               Acceptable values are 'true' and 'false'. Changing this value updates the
               state of the existing member.
        :param pulumi.Input[str] pool_id: The ID of the LB pool. Changing this creates a new
               member.
        :param pulumi.Input[int] port: An integer representing the port on which the member is
               hosted. Changing this creates a new member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB member. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB member.
        :param pulumi.Input[str] tenant_id: The owner of the member. Required if admin wants to
               create a member for another tenant. Changing this creates a new member.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if pool_id is not None:
            pulumi.set(__self__, "pool_id", pool_id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the member. Changing this creates a
        new member.
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
        Acceptable values are 'true' and 'false'. Changing this value updates the
        state of the existing member.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the LB pool. Changing this creates a new
        member.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        An integer representing the port on which the member is
        hosted. Changing this creates a new member.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB member. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB member.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the member. Required if admin wants to
        create a member for another tenant. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class MemberV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a V1 load balancer member resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        member1 = openstack.loadbalancer.MemberV1("member_1",
            pool_id="d9415786-5f1a-428b-b35f-2f1523e146d2",
            address="192.168.0.10",
            port=80)
        ```

        ## Import

        Load Balancer Members can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:loadbalancer/memberV1:MemberV1 member_1 a7498676-4fe4-4243-a864-2eaaf18c73df
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address of the member. Changing this creates a
               new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               Acceptable values are 'true' and 'false'. Changing this value updates the
               state of the existing member.
        :param pulumi.Input[str] pool_id: The ID of the LB pool. Changing this creates a new
               member.
        :param pulumi.Input[int] port: An integer representing the port on which the member is
               hosted. Changing this creates a new member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB member. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB member.
        :param pulumi.Input[str] tenant_id: The owner of the member. Required if admin wants to
               create a member for another tenant. Changing this creates a new member.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MemberV1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V1 load balancer member resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        member1 = openstack.loadbalancer.MemberV1("member_1",
            pool_id="d9415786-5f1a-428b-b35f-2f1523e146d2",
            address="192.168.0.10",
            port=80)
        ```

        ## Import

        Load Balancer Members can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:loadbalancer/memberV1:MemberV1 member_1 a7498676-4fe4-4243-a864-2eaaf18c73df
        ```

        :param str resource_name: The name of the resource.
        :param MemberV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MemberV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MemberV1Args.__new__(MemberV1Args)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["admin_state_up"] = admin_state_up
            if pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'pool_id'")
            __props__.__dict__["pool_id"] = pool_id
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["region"] = region
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["weight"] = weight
        super(MemberV1, __self__).__init__(
            'openstack:loadbalancer/memberV1:MemberV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            pool_id: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[int]] = None) -> 'MemberV1':
        """
        Get an existing MemberV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address of the member. Changing this creates a
               new member.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the member.
               Acceptable values are 'true' and 'false'. Changing this value updates the
               state of the existing member.
        :param pulumi.Input[str] pool_id: The ID of the LB pool. Changing this creates a new
               member.
        :param pulumi.Input[int] port: An integer representing the port on which the member is
               hosted. Changing this creates a new member.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB member. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB member.
        :param pulumi.Input[str] tenant_id: The owner of the member. Required if admin wants to
               create a member for another tenant. Changing this creates a new member.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MemberV1State.__new__(_MemberV1State)

        __props__.__dict__["address"] = address
        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["pool_id"] = pool_id
        __props__.__dict__["port"] = port
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["weight"] = weight
        return MemberV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The IP address of the member. Changing this creates a
        new member.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[bool]:
        """
        The administrative state of the member.
        Acceptable values are 'true' and 'false'. Changing this value updates the
        state of the existing member.
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Output[str]:
        """
        The ID of the LB pool. Changing this creates a new
        member.
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        An integer representing the port on which the member is
        hosted. Changing this creates a new member.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB member. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB member.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The owner of the member. Required if admin wants to
        create a member for another tenant. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[int]:
        return pulumi.get(self, "weight")

