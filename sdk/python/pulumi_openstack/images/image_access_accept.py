# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ImageAccessAccept(pulumi.CustomResource):
    created_at: pulumi.Output[str]
    """
    The date the image membership was created.
    """
    image_id: pulumi.Output[str]
    """
    The proposed image ID.
    """
    member_id: pulumi.Output[str]
    """
    The member ID, e.g. the target project ID. Optional
    for admin accounts. Defaults to the current scope project ID.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Glance client.
    A Glance client is needed to manage Image memberships. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    membership.
    """
    schema: pulumi.Output[str]
    """
    The membership schema.
    """
    status: pulumi.Output[str]
    """
    The membership proposal status. Can either be
    `accepted`, `rejected` or `pending`.
    """
    updated_at: pulumi.Output[str]
    """
    The date the image membership was last updated.
    """
    def __init__(__self__, resource_name, opts=None, image_id=None, member_id=None, region=None, status=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages memberships status for the shared OpenStack Glance V2 Image within the
        destination project, which has a member proposal.

        ## Example Usage

        Accept a shared image membershipship proposal within the current project.

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rancheros = openstack.images.get_image(member_status="all",
            name="RancherOS",
            visibility="shared")
        rancheros_member = openstack.images.ImageAccessAccept("rancherosMember",
            image_id=rancheros.id,
            status="accepted")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_id: The proposed image ID.
        :param pulumi.Input[str] member_id: The member ID, e.g. the target project ID. Optional
               for admin accounts. Defaults to the current scope project ID.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image memberships. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               membership.
        :param pulumi.Input[str] status: The membership proposal status. Can either be
               `accepted`, `rejected` or `pending`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if image_id is None:
                raise TypeError("Missing required property 'image_id'")
            __props__['image_id'] = image_id
            __props__['member_id'] = member_id
            __props__['region'] = region
            if status is None:
                raise TypeError("Missing required property 'status'")
            __props__['status'] = status
            __props__['created_at'] = None
            __props__['schema'] = None
            __props__['updated_at'] = None
        super(ImageAccessAccept, __self__).__init__(
            'openstack:images/imageAccessAccept:ImageAccessAccept',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, created_at=None, image_id=None, member_id=None, region=None, schema=None, status=None, updated_at=None):
        """
        Get an existing ImageAccessAccept resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date the image membership was created.
        :param pulumi.Input[str] image_id: The proposed image ID.
        :param pulumi.Input[str] member_id: The member ID, e.g. the target project ID. Optional
               for admin accounts. Defaults to the current scope project ID.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image memberships. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               membership.
        :param pulumi.Input[str] schema: The membership schema.
        :param pulumi.Input[str] status: The membership proposal status. Can either be
               `accepted`, `rejected` or `pending`.
        :param pulumi.Input[str] updated_at: The date the image membership was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["created_at"] = created_at
        __props__["image_id"] = image_id
        __props__["member_id"] = member_id
        __props__["region"] = region
        __props__["schema"] = schema
        __props__["status"] = status
        __props__["updated_at"] = updated_at
        return ImageAccessAccept(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
