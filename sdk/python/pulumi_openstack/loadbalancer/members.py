# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
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
from . import outputs
from ._inputs import *

__all__ = ['MembersArgs', 'Members']

@pulumi.input_type
class MembersArgs:
    def __init__(__self__, *,
                 pool_id: pulumi.Input[builtins.str],
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Members resource.
        :param pulumi.Input[builtins.str] pool_id: The id of the pool that members will be assigned to.
               Changing this creates a new members resource.
        :param pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]] members: A set of dictionaries containing member parameters. The
               structure is described below.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create pool members. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               members resource.
        """
        pulumi.set(__self__, "pool_id", pool_id)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Input[builtins.str]:
        """
        The id of the pool that members will be assigned to.
        Changing this creates a new members resource.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]]]:
        """
        A set of dictionaries containing member parameters. The
        structure is described below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create pool members. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        members resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _MembersState:
    def __init__(__self__, *,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Members resources.
        :param pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]] members: A set of dictionaries containing member parameters. The
               structure is described below.
        :param pulumi.Input[builtins.str] pool_id: The id of the pool that members will be assigned to.
               Changing this creates a new members resource.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create pool members. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               members resource.
        """
        if members is not None:
            pulumi.set(__self__, "members", members)
        if pool_id is not None:
            pulumi.set(__self__, "pool_id", pool_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]]]:
        """
        A set of dictionaries containing member parameters. The
        structure is described below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MembersMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the pool that members will be assigned to.
        Changing this creates a new members resource.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create pool members. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        members resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


class Members(pulumi.CustomResource):

    pulumi_type = "openstack:loadbalancer/members:Members"

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MembersMemberArgs', 'MembersMemberArgsDict']]]]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages a V2 members resource within OpenStack (batch members update).

        > **Note:** This resource has attributes that depend on octavia minor versions.
        Please ensure your Openstack cloud supports the required minor version.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        members1 = openstack.loadbalancer.Members("members_1",
            pool_id="935685fb-a896-40f9-9ff4-ae531a3a00fe",
            members=[
                {
                    "address": "192.168.199.23",
                    "protocol_port": 8080,
                },
                {
                    "address": "192.168.199.24",
                    "protocol_port": 8080,
                },
            ])
        ```

        ## Import

        Load Balancer Pool Members can be imported using the Pool ID, e.g.:

        ```sh
        $ pulumi import openstack:loadbalancer/members:Members members_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MembersMemberArgs', 'MembersMemberArgsDict']]]] members: A set of dictionaries containing member parameters. The
               structure is described below.
        :param pulumi.Input[builtins.str] pool_id: The id of the pool that members will be assigned to.
               Changing this creates a new members resource.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create pool members. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               members resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MembersArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 members resource within OpenStack (batch members update).

        > **Note:** This resource has attributes that depend on octavia minor versions.
        Please ensure your Openstack cloud supports the required minor version.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        members1 = openstack.loadbalancer.Members("members_1",
            pool_id="935685fb-a896-40f9-9ff4-ae531a3a00fe",
            members=[
                {
                    "address": "192.168.199.23",
                    "protocol_port": 8080,
                },
                {
                    "address": "192.168.199.24",
                    "protocol_port": 8080,
                },
            ])
        ```

        ## Import

        Load Balancer Pool Members can be imported using the Pool ID, e.g.:

        ```sh
        $ pulumi import openstack:loadbalancer/members:Members members_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5
        ```

        :param str resource_name: The name of the resource.
        :param MembersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MembersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MembersMemberArgs', 'MembersMemberArgsDict']]]]] = None,
                 pool_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MembersArgs.__new__(MembersArgs)

            __props__.__dict__["members"] = members
            if pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'pool_id'")
            __props__.__dict__["pool_id"] = pool_id
            __props__.__dict__["region"] = region
        super(Members, __self__).__init__(
            'openstack:loadbalancer/members:Members',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[Union['MembersMemberArgs', 'MembersMemberArgsDict']]]]] = None,
            pool_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None) -> 'Members':
        """
        Get an existing Members resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['MembersMemberArgs', 'MembersMemberArgsDict']]]] members: A set of dictionaries containing member parameters. The
               structure is described below.
        :param pulumi.Input[builtins.str] pool_id: The id of the pool that members will be assigned to.
               Changing this creates a new members resource.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create pool members. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               members resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MembersState.__new__(_MembersState)

        __props__.__dict__["members"] = members
        __props__.__dict__["pool_id"] = pool_id
        __props__.__dict__["region"] = region
        return Members(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Optional[Sequence['outputs.MembersMember']]]:
        """
        A set of dictionaries containing member parameters. The
        structure is described below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Output[builtins.str]:
        """
        The id of the pool that members will be assigned to.
        Changing this creates a new members resource.
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create pool members. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        members resource.
        """
        return pulumi.get(self, "region")

