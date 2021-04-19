# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserMembershipV3Args', 'UserMembershipV3']

@pulumi.input_type
class UserMembershipV3Args:
    def __init__(__self__, *,
                 group_id: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserMembershipV3 resource.
        :param pulumi.Input[str] group_id: The UUID of group to which the user will be added.
               Changing this creates a new user membership.
        :param pulumi.Input[str] user_id: The UUID of user to use. Changing this creates a new user membership.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Identity client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new user membership.
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "user_id", user_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[str]:
        """
        The UUID of group to which the user will be added.
        Changing this creates a new user membership.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        The UUID of user to use. Changing this creates a new user membership.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V3 Identity client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new user membership.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _UserMembershipV3State:
    def __init__(__self__, *,
                 group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserMembershipV3 resources.
        :param pulumi.Input[str] group_id: The UUID of group to which the user will be added.
               Changing this creates a new user membership.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Identity client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new user membership.
        :param pulumi.Input[str] user_id: The UUID of user to use. Changing this creates a new user membership.
        """
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of group to which the user will be added.
        Changing this creates a new user membership.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V3 Identity client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new user membership.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of user to use. Changing this creates a new user membership.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class UserMembershipV3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a user membership to group V3 resource within OpenStack.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ***

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        user1 = openstack.identity.User("user1", default_project_id=project1.id)
        group1 = openstack.identity.GroupV3("group1", description="group 1")
        role1 = openstack.identity.Role("role1")
        user_membership1 = openstack.identity.UserMembershipV3("userMembership1",
            group_id=group1.id,
            user_id=user1.id)
        role_assignment1 = openstack.identity.RoleAssignment("roleAssignment1",
            group_id=group1.id,
            project_id=project1.id,
            role_id=role1.id)
        ```

        ## Import

        This resource can be imported by specifying all two arguments, separated by a forward slash

        ```sh
         $ pulumi import openstack:identity/userMembershipV3:UserMembershipV3 user_membership_1 <user_id>/<group_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: The UUID of group to which the user will be added.
               Changing this creates a new user membership.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Identity client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new user membership.
        :param pulumi.Input[str] user_id: The UUID of user to use. Changing this creates a new user membership.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserMembershipV3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a user membership to group V3 resource within OpenStack.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ***

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        user1 = openstack.identity.User("user1", default_project_id=project1.id)
        group1 = openstack.identity.GroupV3("group1", description="group 1")
        role1 = openstack.identity.Role("role1")
        user_membership1 = openstack.identity.UserMembershipV3("userMembership1",
            group_id=group1.id,
            user_id=user1.id)
        role_assignment1 = openstack.identity.RoleAssignment("roleAssignment1",
            group_id=group1.id,
            project_id=project1.id,
            role_id=role1.id)
        ```

        ## Import

        This resource can be imported by specifying all two arguments, separated by a forward slash

        ```sh
         $ pulumi import openstack:identity/userMembershipV3:UserMembershipV3 user_membership_1 <user_id>/<group_id>
        ```

        :param str resource_name: The name of the resource.
        :param UserMembershipV3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserMembershipV3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserMembershipV3Args.__new__(UserMembershipV3Args)

            if group_id is None and not opts.urn:
                raise TypeError("Missing required property 'group_id'")
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["region"] = region
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(UserMembershipV3, __self__).__init__(
            'openstack:identity/userMembershipV3:UserMembershipV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'UserMembershipV3':
        """
        Get an existing UserMembershipV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: The UUID of group to which the user will be added.
               Changing this creates a new user membership.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Identity client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new user membership.
        :param pulumi.Input[str] user_id: The UUID of user to use. Changing this creates a new user membership.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserMembershipV3State.__new__(_UserMembershipV3State)

        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["region"] = region
        __props__.__dict__["user_id"] = user_id
        return UserMembershipV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[str]:
        """
        The UUID of group to which the user will be added.
        Changing this creates a new user membership.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V3 Identity client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new user membership.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        The UUID of user to use. Changing this creates a new user membership.
        """
        return pulumi.get(self, "user_id")

