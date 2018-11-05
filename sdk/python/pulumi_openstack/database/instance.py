# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Instance(pulumi.CustomResource):
    """
    Manages a V1 DB instance resource within OpenStack.
    """
    def __init__(__self__, __name__, __opts__=None, configuration_id=None, databases=None, datastore=None, flavor_id=None, name=None, networks=None, region=None, size=None, users=None):
        """Create a Instance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['configurationId'] = configuration_id

        __props__['databases'] = databases

        if not datastore:
            raise TypeError('Missing required property datastore')
        __props__['datastore'] = datastore

        __props__['flavorId'] = flavor_id

        __props__['name'] = name

        __props__['networks'] = networks

        if not region:
            raise TypeError('Missing required property region')
        __props__['region'] = region

        if not size:
            raise TypeError('Missing required property size')
        __props__['size'] = size

        __props__['users'] = users

        super(Instance, __self__).__init__(
            'openstack:database/instance:Instance',
            __name__,
            __props__,
            __opts__)

