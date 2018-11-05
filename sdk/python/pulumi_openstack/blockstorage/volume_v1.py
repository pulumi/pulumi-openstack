# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class VolumeV1(pulumi.CustomResource):
    """
    Manages a V1 volume resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, availability_zone=None, description=None, image_id=None, metadata=None, name=None, region=None, size=None, snapshot_id=None, source_vol_id=None, volume_type=None):
        """Create a VolumeV1 resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['availabilityZone'] = availability_zone

        __props__['description'] = description

        __props__['imageId'] = image_id

        __props__['metadata'] = metadata

        __props__['name'] = name

        __props__['region'] = region

        if not size:
            raise TypeError('Missing required property size')
        __props__['size'] = size

        __props__['snapshotId'] = snapshot_id

        __props__['sourceVolId'] = source_vol_id

        __props__['volumeType'] = volume_type

        __props__['attachments'] = None

        super(VolumeV1, __self__).__init__(
            'openstack:blockstorage/volumeV1:VolumeV1',
            __name__,
            __props__,
            __opts__)

