# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Zone(pulumi.CustomResource):
    attributes: pulumi.Output[dict]
    """
    Attributes for the DNS Service scheduler.
    Changing this creates a new zone.
    """
    description: pulumi.Output[str]
    """
    A description of the zone.
    """
    email: pulumi.Output[str]
    """
    The email contact for the zone record.
    """
    masters: pulumi.Output[list]
    """
    An array of master DNS servers. For when `type` is
    `SECONDARY`.
    """
    name: pulumi.Output[str]
    """
    The name of the zone. Note the `.` at the end of the name.
    Changing this creates a new DNS zone.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Compute client.
    Keypairs are associated with accounts, but a Compute client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new DNS zone.
    """
    ttl: pulumi.Output[float]
    """
    The time to live (TTL) of the zone.
    """
    type: pulumi.Output[str]
    """
    The type of zone. Can either be `PRIMARY` or `SECONDARY`.
    Changing this creates a new zone.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options. Changing this creates a
    new zone.
    """
    def __init__(__self__, resource_name, opts=None, attributes=None, description=None, email=None, masters=None, name=None, region=None, ttl=None, type=None, value_specs=None, __name__=None, __opts__=None):
        """
        Manages a DNS zone in the OpenStack DNS Service.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attributes: Attributes for the DNS Service scheduler.
               Changing this creates a new zone.
        :param pulumi.Input[str] description: A description of the zone.
        :param pulumi.Input[str] email: The email contact for the zone record.
        :param pulumi.Input[list] masters: An array of master DNS servers. For when `type` is
               `SECONDARY`.
        :param pulumi.Input[str] name: The name of the zone. Note the `.` at the end of the name.
               Changing this creates a new DNS zone.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[float] ttl: The time to live (TTL) of the zone.
        :param pulumi.Input[str] type: The type of zone. Can either be `PRIMARY` or `SECONDARY`.
               Changing this creates a new zone.
        :param pulumi.Input[dict] value_specs: Map of additional options. Changing this creates a
               new zone.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/dns_zone_v2.html.markdown.
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

        __props__['attributes'] = attributes

        __props__['description'] = description

        __props__['email'] = email

        __props__['masters'] = masters

        __props__['name'] = name

        __props__['region'] = region

        __props__['ttl'] = ttl

        __props__['type'] = type

        __props__['value_specs'] = value_specs

        super(Zone, __self__).__init__(
            'openstack:dns/zone:Zone',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

