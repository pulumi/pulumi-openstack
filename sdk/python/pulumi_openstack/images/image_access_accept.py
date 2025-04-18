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

__all__ = ['ImageAccessAcceptArgs', 'ImageAccessAccept']

@pulumi.input_type
class ImageAccessAcceptArgs:
    def __init__(__self__, *,
                 image_id: pulumi.Input[builtins.str],
                 status: pulumi.Input[builtins.str],
                 member_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ImageAccessAccept resource.
        :param pulumi.Input[builtins.str] image_id: The proposed image ID.
        :param pulumi.Input[builtins.str] status: The membership proposal status. Can either be
               `accepted`, `rejected` or `pending`.
        :param pulumi.Input[builtins.str] member_id: The member ID, e.g. the target project ID. Optional
               for admin accounts. Defaults to the current scope project ID.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image memberships. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               membership.
        """
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "status", status)
        if member_id is not None:
            pulumi.set(__self__, "member_id", member_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Input[builtins.str]:
        """
        The proposed image ID.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[builtins.str]:
        """
        The membership proposal status. Can either be
        `accepted`, `rejected` or `pending`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="memberId")
    def member_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The member ID, e.g. the target project ID. Optional
        for admin accounts. Defaults to the current scope project ID.
        """
        return pulumi.get(self, "member_id")

    @member_id.setter
    def member_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "member_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Glance client.
        A Glance client is needed to manage Image memberships. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        membership.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ImageAccessAcceptState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 image_id: Optional[pulumi.Input[builtins.str]] = None,
                 member_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 schema: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ImageAccessAccept resources.
        :param pulumi.Input[builtins.str] created_at: The date the image membership was created.
        :param pulumi.Input[builtins.str] image_id: The proposed image ID.
        :param pulumi.Input[builtins.str] member_id: The member ID, e.g. the target project ID. Optional
               for admin accounts. Defaults to the current scope project ID.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image memberships. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               membership.
        :param pulumi.Input[builtins.str] schema: The membership schema.
        :param pulumi.Input[builtins.str] status: The membership proposal status. Can either be
               `accepted`, `rejected` or `pending`.
        :param pulumi.Input[builtins.str] updated_at: The date the image membership was last updated.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if member_id is not None:
            pulumi.set(__self__, "member_id", member_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date the image membership was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The proposed image ID.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="memberId")
    def member_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The member ID, e.g. the target project ID. Optional
        for admin accounts. Defaults to the current scope project ID.
        """
        return pulumi.get(self, "member_id")

    @member_id.setter
    def member_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "member_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Glance client.
        A Glance client is needed to manage Image memberships. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        membership.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The membership schema.
        """
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The membership proposal status. Can either be
        `accepted`, `rejected` or `pending`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date the image membership was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


class ImageAccessAccept(pulumi.CustomResource):

    pulumi_type = "openstack:images/imageAccessAccept:ImageAccessAccept"

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[builtins.str]] = None,
                 member_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages memberships status for the shared OpenStack Glance V2 Image within the
        destination project, which has a member proposal.

        ## Example Usage

        Accept a shared image membershipship proposal within the current project.

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rancheros = openstack.images.get_image(name="RancherOS",
            visibility="shared",
            member_status="all")
        rancheros_member = openstack.images.ImageAccessAccept("rancheros_member",
            image_id=rancheros.id,
            status="accepted")
        ```

        ## Import

        Image access acceptance status can be imported using the `image_id`, e.g.

        ```sh
        $ pulumi import openstack:images/imageAccessAccept:ImageAccessAccept openstack_images_image_access_accept_v2 89c60255-9bd6-460c-822a-e2b959ede9d2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] image_id: The proposed image ID.
        :param pulumi.Input[builtins.str] member_id: The member ID, e.g. the target project ID. Optional
               for admin accounts. Defaults to the current scope project ID.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image memberships. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               membership.
        :param pulumi.Input[builtins.str] status: The membership proposal status. Can either be
               `accepted`, `rejected` or `pending`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageAccessAcceptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages memberships status for the shared OpenStack Glance V2 Image within the
        destination project, which has a member proposal.

        ## Example Usage

        Accept a shared image membershipship proposal within the current project.

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rancheros = openstack.images.get_image(name="RancherOS",
            visibility="shared",
            member_status="all")
        rancheros_member = openstack.images.ImageAccessAccept("rancheros_member",
            image_id=rancheros.id,
            status="accepted")
        ```

        ## Import

        Image access acceptance status can be imported using the `image_id`, e.g.

        ```sh
        $ pulumi import openstack:images/imageAccessAccept:ImageAccessAccept openstack_images_image_access_accept_v2 89c60255-9bd6-460c-822a-e2b959ede9d2
        ```

        :param str resource_name: The name of the resource.
        :param ImageAccessAcceptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageAccessAcceptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 image_id: Optional[pulumi.Input[builtins.str]] = None,
                 member_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageAccessAcceptArgs.__new__(ImageAccessAcceptArgs)

            if image_id is None and not opts.urn:
                raise TypeError("Missing required property 'image_id'")
            __props__.__dict__["image_id"] = image_id
            __props__.__dict__["member_id"] = member_id
            __props__.__dict__["region"] = region
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["created_at"] = None
            __props__.__dict__["schema"] = None
            __props__.__dict__["updated_at"] = None
        super(ImageAccessAccept, __self__).__init__(
            'openstack:images/imageAccessAccept:ImageAccessAccept',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            image_id: Optional[pulumi.Input[builtins.str]] = None,
            member_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            schema: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None) -> 'ImageAccessAccept':
        """
        Get an existing ImageAccessAccept resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: The date the image membership was created.
        :param pulumi.Input[builtins.str] image_id: The proposed image ID.
        :param pulumi.Input[builtins.str] member_id: The member ID, e.g. the target project ID. Optional
               for admin accounts. Defaults to the current scope project ID.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image memberships. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               membership.
        :param pulumi.Input[builtins.str] schema: The membership schema.
        :param pulumi.Input[builtins.str] status: The membership proposal status. Can either be
               `accepted`, `rejected` or `pending`.
        :param pulumi.Input[builtins.str] updated_at: The date the image membership was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImageAccessAcceptState.__new__(_ImageAccessAcceptState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["member_id"] = member_id
        __props__.__dict__["region"] = region
        __props__.__dict__["schema"] = schema
        __props__.__dict__["status"] = status
        __props__.__dict__["updated_at"] = updated_at
        return ImageAccessAccept(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The date the image membership was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[builtins.str]:
        """
        The proposed image ID.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="memberId")
    def member_id(self) -> pulumi.Output[builtins.str]:
        """
        The member ID, e.g. the target project ID. Optional
        for admin accounts. Defaults to the current scope project ID.
        """
        return pulumi.get(self, "member_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to obtain the V2 Glance client.
        A Glance client is needed to manage Image memberships. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        membership.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[builtins.str]:
        """
        The membership schema.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        The membership proposal status. Can either be
        `accepted`, `rejected` or `pending`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        The date the image membership was last updated.
        """
        return pulumi.get(self, "updated_at")

