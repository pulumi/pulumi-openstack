# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class RecordSet(pulumi.CustomResource):
    """
    Manages a DNS record set in the OpenStack DNS Service.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, name=None, records=None, region=None, ttl=None, type=None, value_specs=None, zone_id=None):
        """Create a RecordSet resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['name'] = name

        __props__['records'] = records

        __props__['region'] = region

        __props__['ttl'] = ttl

        __props__['type'] = type

        __props__['valueSpecs'] = value_specs

        if not zone_id:
            raise TypeError('Missing required property zone_id')
        __props__['zoneId'] = zone_id

        super(RecordSet, __self__).__init__(
            'openstack:dns/recordSet:RecordSet',
            __name__,
            __props__,
            __opts__)

