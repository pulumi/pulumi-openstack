# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Volume(pulumi.CustomResource):
    """
    Manages a V3 volume resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, availability_zone=None, consistency_group_id=None, description=None, enable_online_resize=None, image_id=None, metadata=None, multiattach=None, name=None, region=None, size=None, snapshot_id=None, source_replica=None, source_vol_id=None, volume_type=None):
        """Create a Volume resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['availability_zone'] = availability_zone

        __props__['consistency_group_id'] = consistency_group_id

        __props__['description'] = description

        __props__['enable_online_resize'] = enable_online_resize

        __props__['image_id'] = image_id

        __props__['metadata'] = metadata

        __props__['multiattach'] = multiattach

        __props__['name'] = name

        __props__['region'] = region

        if not size:
            raise TypeError('Missing required property size')
        __props__['size'] = size

        __props__['snapshot_id'] = snapshot_id

        __props__['source_replica'] = source_replica

        __props__['source_vol_id'] = source_vol_id

        __props__['volume_type'] = volume_type

        __props__['attachments'] = None

        super(Volume, __self__).__init__(
            'openstack:blockstorage/volume:Volume',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

