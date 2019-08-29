# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RecordSet(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description of the  record set.
    """
    name: pulumi.Output[str]
    """
    The name of the record set. Note the `.` at the end of the name.
    Changing this creates a new DNS  record set.
    """
    records: pulumi.Output[list]
    """
    An array of DNS records. _Note:_ if an IPv6 address
    contains brackets (`[ ]`), the brackets will be stripped and the modified
    address will be recorded in the state.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 DNS client.
    If omitted, the `region` argument of the provider is used.
    Changing this creates a new DNS  record set.
    """
    ttl: pulumi.Output[float]
    """
    The time to live (TTL) of the record set.
    """
    type: pulumi.Output[str]
    """
    The type of record set. Examples: "A", "MX".
    Changing this creates a new DNS  record set.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options. Changing this creates a
    new record set.
    """
    zone_id: pulumi.Output[str]
    """
    The ID of the zone in which to create the record set.
    Changing this creates a new DNS  record set.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, records=None, region=None, ttl=None, type=None, value_specs=None, zone_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a DNS record set in the OpenStack DNS Service.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the  record set.
        :param pulumi.Input[str] name: The name of the record set. Note the `.` at the end of the name.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[list] records: An array of DNS records. _Note:_ if an IPv6 address
               contains brackets (`[ ]`), the brackets will be stripped and the modified
               address will be recorded in the state.
        :param pulumi.Input[str] region: The region in which to obtain the V2 DNS client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[float] ttl: The time to live (TTL) of the record set.
        :param pulumi.Input[str] type: The type of record set. Examples: "A", "MX".
               Changing this creates a new DNS  record set.
        :param pulumi.Input[dict] value_specs: Map of additional options. Changing this creates a
               new record set.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the record set.
               Changing this creates a new DNS  record set.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/dns_recordset_v2.html.markdown.
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

            __props__['description'] = description
            __props__['name'] = name
            __props__['records'] = records
            __props__['region'] = region
            __props__['ttl'] = ttl
            __props__['type'] = type
            __props__['value_specs'] = value_specs
            if zone_id is None:
                raise TypeError("Missing required property 'zone_id'")
            __props__['zone_id'] = zone_id
        super(RecordSet, __self__).__init__(
            'openstack:dns/recordSet:RecordSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, name=None, records=None, region=None, ttl=None, type=None, value_specs=None, zone_id=None):
        """
        Get an existing RecordSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the  record set.
        :param pulumi.Input[str] name: The name of the record set. Note the `.` at the end of the name.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[list] records: An array of DNS records. _Note:_ if an IPv6 address
               contains brackets (`[ ]`), the brackets will be stripped and the modified
               address will be recorded in the state.
        :param pulumi.Input[str] region: The region in which to obtain the V2 DNS client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[float] ttl: The time to live (TTL) of the record set.
        :param pulumi.Input[str] type: The type of record set. Examples: "A", "MX".
               Changing this creates a new DNS  record set.
        :param pulumi.Input[dict] value_specs: Map of additional options. Changing this creates a
               new record set.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the record set.
               Changing this creates a new DNS  record set.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/dns_recordset_v2.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["description"] = description
        __props__["name"] = name
        __props__["records"] = records
        __props__["region"] = region
        __props__["ttl"] = ttl
        __props__["type"] = type
        __props__["value_specs"] = value_specs
        __props__["zone_id"] = zone_id
        return RecordSet(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

