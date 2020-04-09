# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ImageAccess(pulumi.CustomResource):
    created_at: pulumi.Output[str]
    """
    The date the image access was created.
    """
    image_id: pulumi.Output[str]
    """
    The image ID.
    """
    member_id: pulumi.Output[str]
    """
    The member ID, e.g. the target project ID.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Glance client.
    A Glance client is needed to manage Image members. If omitted, the `region`
    argument of the provider is used. Changing this creates a new resource.
    """
    schema: pulumi.Output[str]
    """
    The member schema.
    """
    status: pulumi.Output[str]
    """
    The member proposal status. Optional if admin wants to
    force the member proposal acceptance. Can either be `accepted`, `rejected` or
    `pending`. Defaults to `pending`. Foridden for non-admin users.
    """
    updated_at: pulumi.Output[str]
    """
    The date the image access was last updated.
    """
    def __init__(__self__, resource_name, opts=None, image_id=None, member_id=None, region=None, status=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages members for the shared OpenStack Glance V2 Image within the source
        project, which owns the Image.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/images_image_access_v2.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] image_id: The image ID.
        :param pulumi.Input[str] member_id: The member ID, e.g. the target project ID.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image members. If omitted, the `region`
               argument of the provider is used. Changing this creates a new resource.
        :param pulumi.Input[str] status: The member proposal status. Optional if admin wants to
               force the member proposal acceptance. Can either be `accepted`, `rejected` or
               `pending`. Defaults to `pending`. Foridden for non-admin users.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if image_id is None:
                raise TypeError("Missing required property 'image_id'")
            __props__['image_id'] = image_id
            if member_id is None:
                raise TypeError("Missing required property 'member_id'")
            __props__['member_id'] = member_id
            __props__['region'] = region
            __props__['status'] = status
            __props__['created_at'] = None
            __props__['schema'] = None
            __props__['updated_at'] = None
        super(ImageAccess, __self__).__init__(
            'openstack:images/imageAccess:ImageAccess',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, created_at=None, image_id=None, member_id=None, region=None, schema=None, status=None, updated_at=None):
        """
        Get an existing ImageAccess resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date the image access was created.
        :param pulumi.Input[str] image_id: The image ID.
        :param pulumi.Input[str] member_id: The member ID, e.g. the target project ID.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Glance client.
               A Glance client is needed to manage Image members. If omitted, the `region`
               argument of the provider is used. Changing this creates a new resource.
        :param pulumi.Input[str] schema: The member schema.
        :param pulumi.Input[str] status: The member proposal status. Optional if admin wants to
               force the member proposal acceptance. Can either be `accepted`, `rejected` or
               `pending`. Defaults to `pending`. Foridden for non-admin users.
        :param pulumi.Input[str] updated_at: The date the image access was last updated.
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
        return ImageAccess(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

