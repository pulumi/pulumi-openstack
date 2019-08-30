# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EndpointV3(pulumi.CustomResource):
    endpoint_region: pulumi.Output[str]
    """
    The endpoint region. The `region` and
    `endpoint_region` can be different.
    """
    interface: pulumi.Output[str]
    """
    The endpoint interface. Valid values are `public`,
    `internal` and `admin`. Default value is `public`
    """
    name: pulumi.Output[str]
    """
    The endpoint name.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V3 Keystone client.
    If omitted, the `region` argument of the provider is used.
    """
    service_id: pulumi.Output[str]
    """
    The endpoint service ID.
    """
    service_name: pulumi.Output[str]
    """
    The service name of the endpoint.
    """
    service_type: pulumi.Output[str]
    """
    The service type of the endpoint.
    """
    url: pulumi.Output[str]
    """
    The endpoint url.
    """
    def __init__(__self__, resource_name, opts=None, endpoint_region=None, interface=None, name=None, region=None, service_id=None, url=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V3 Endpoint resource within OpenStack Keystone.
        
        > **Note:** This usually requires admin privileges.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_region: The endpoint region. The `region` and
               `endpoint_region` can be different.
        :param pulumi.Input[str] interface: The endpoint interface. Valid values are `public`,
               `internal` and `admin`. Default value is `public`
        :param pulumi.Input[str] name: The endpoint name.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used.
        :param pulumi.Input[str] service_id: The endpoint service ID.
        :param pulumi.Input[str] url: The endpoint url.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_endpoint_v3.html.markdown.
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

            if endpoint_region is None:
                raise TypeError("Missing required property 'endpoint_region'")
            __props__['endpoint_region'] = endpoint_region
            __props__['interface'] = interface
            __props__['name'] = name
            __props__['region'] = region
            if service_id is None:
                raise TypeError("Missing required property 'service_id'")
            __props__['service_id'] = service_id
            if url is None:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['service_name'] = None
            __props__['service_type'] = None
        super(EndpointV3, __self__).__init__(
            'openstack:identity/endpointV3:EndpointV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, endpoint_region=None, interface=None, name=None, region=None, service_id=None, service_name=None, service_type=None, url=None):
        """
        Get an existing EndpointV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_region: The endpoint region. The `region` and
               `endpoint_region` can be different.
        :param pulumi.Input[str] interface: The endpoint interface. Valid values are `public`,
               `internal` and `admin`. Default value is `public`
        :param pulumi.Input[str] name: The endpoint name.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used.
        :param pulumi.Input[str] service_id: The endpoint service ID.
        :param pulumi.Input[str] service_name: The service name of the endpoint.
        :param pulumi.Input[str] service_type: The service type of the endpoint.
        :param pulumi.Input[str] url: The endpoint url.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/identity_endpoint_v3.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["endpoint_region"] = endpoint_region
        __props__["interface"] = interface
        __props__["name"] = name
        __props__["region"] = region
        __props__["service_id"] = service_id
        __props__["service_name"] = service_name
        __props__["service_type"] = service_type
        __props__["url"] = url
        return EndpointV3(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
