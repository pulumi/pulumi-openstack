# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class TempUrl(pulumi.CustomResource):
    container: pulumi.Output[str]
    """
    The container name the object belongs to.
    """
    method: pulumi.Output[str]
    """
    The method allowed when accessing this URL.
    Valid values are `GET`, and `POST`. Default is `GET`.
    """
    object: pulumi.Output[str]
    """
    The object name the tempurl is for.
    """
    regenerate: pulumi.Output[bool]
    """
    Whether to automatically regenerate the URL when
    it has expired. If set to true, this will create a new resource with a new
    ID and new URL. Defaults to false.
    """
    region: pulumi.Output[str]
    """
    The region the tempurl is located in.
    """
    split: pulumi.Output[str]
    ttl: pulumi.Output[float]
    """
    The TTL, in seconds, for the URL. For how long it should
    be valid.
    """
    url: pulumi.Output[str]
    """
    The URL
    """
    def __init__(__self__, resource_name, opts=None, container=None, method=None, object=None, regenerate=None, region=None, split=None, ttl=None, __name__=None, __opts__=None):
        """
        Use this resource to generate an OpenStack Object Storage temporary URL.
        
        The temporary URL will be valid for as long as TTL is set to (in seconds).
        Once the URL has expired, it will no longer be valid, but the resource
        will remain in place. If you wish to automatically regenerate a URL, set
        the `regenerate` argument to `true`. This will create a new resource with
        a new ID and URL.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container: The container name the object belongs to.
        :param pulumi.Input[str] method: The method allowed when accessing this URL.
               Valid values are `GET`, and `POST`. Default is `GET`.
        :param pulumi.Input[str] object: The object name the tempurl is for.
        :param pulumi.Input[bool] regenerate: Whether to automatically regenerate the URL when
               it has expired. If set to true, this will create a new resource with a new
               ID and new URL. Defaults to false.
        :param pulumi.Input[str] region: The region the tempurl is located in.
        :param pulumi.Input[float] ttl: The TTL, in seconds, for the URL. For how long it should
               be valid.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/objectstorage_tempurl_v1.html.markdown.
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

        if container is None:
            raise TypeError("Missing required property 'container'")
        __props__['container'] = container

        __props__['method'] = method

        if object is None:
            raise TypeError("Missing required property 'object'")
        __props__['object'] = object

        __props__['regenerate'] = regenerate

        __props__['region'] = region

        __props__['split'] = split

        if ttl is None:
            raise TypeError("Missing required property 'ttl'")
        __props__['ttl'] = ttl

        __props__['url'] = None

        super(TempUrl, __self__).__init__(
            'openstack:objectstorage/tempUrl:TempUrl',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

