# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VolumeAttach(pulumi.CustomResource):
    device: pulumi.Output[str]
    instance_id: pulumi.Output[str]
    """
    The ID of the Instance to attach the Volume to.
    """
    multiattach: pulumi.Output[bool]
    """
    Enable attachment of multiattach-capable volumes.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Compute client.
    A Compute client is needed to create a volume attachment. If omitted, the
    `region` argument of the provider is used. Changing this creates a
    new volume attachment.
    """
    volume_id: pulumi.Output[str]
    """
    The ID of the Volume to attach to an Instance.
    """
    def __init__(__self__, resource_name, opts=None, device=None, instance_id=None, multiattach=None, region=None, volume_id=None, __name__=None, __opts__=None):
        """
        Attaches a Block Storage Volume to an Instance using the OpenStack
        Compute (Nova) v2 API.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Volume to.
        :param pulumi.Input[bool] multiattach: Enable attachment of multiattach-capable volumes.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a volume attachment. If omitted, the
               `region` argument of the provider is used. Changing this creates a
               new volume attachment.
        :param pulumi.Input[str] volume_id: The ID of the Volume to attach to an Instance.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_volume_attach_v2.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['device'] = device

        if instance_id is None:
            raise TypeError("Missing required property 'instance_id'")
        __props__['instance_id'] = instance_id

        __props__['multiattach'] = multiattach

        __props__['region'] = region

        if volume_id is None:
            raise TypeError("Missing required property 'volume_id'")
        __props__['volume_id'] = volume_id

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(VolumeAttach, __self__).__init__(
            'openstack:compute/volumeAttach:VolumeAttach',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

