# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Flavor(pulumi.CustomResource):
    """
    Manages a V2 flavor resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, disk=None, ephemeral=None, extra_specs=None, is_public=None, name=None, ram=None, region=None, rx_tx_factor=None, swap=None, vcpus=None):
        """Create a Flavor resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not disk:
            raise TypeError('Missing required property disk')
        __props__['disk'] = disk

        __props__['ephemeral'] = ephemeral

        __props__['extraSpecs'] = extra_specs

        __props__['isPublic'] = is_public

        __props__['name'] = name

        if not ram:
            raise TypeError('Missing required property ram')
        __props__['ram'] = ram

        __props__['region'] = region

        __props__['rxTxFactor'] = rx_tx_factor

        __props__['swap'] = swap

        if not vcpus:
            raise TypeError('Missing required property vcpus')
        __props__['vcpus'] = vcpus

        super(Flavor, __self__).__init__(
            'openstack:compute/flavor:Flavor',
            __name__,
            __props__,
            __opts__)

