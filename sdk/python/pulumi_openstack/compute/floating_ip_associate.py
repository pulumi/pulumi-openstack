# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class FloatingIpAssociate(pulumi.CustomResource):
    fixed_ip: pulumi.Output[str]
    """
    The specific IP address to direct traffic to.
    """
    floating_ip: pulumi.Output[str]
    """
    The floating IP to associate.
    """
    instance_id: pulumi.Output[str]
    """
    The instance to associte the floating IP with.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Compute client.
    Keypairs are associated with accounts, but a Compute client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new floatingip_associate.
    """
    wait_until_associated: pulumi.Output[bool]
    def __init__(__self__, resource_name, opts=None, fixed_ip=None, floating_ip=None, instance_id=None, region=None, wait_until_associated=None, __props__=None, __name__=None, __opts__=None):
        """
        Associate a floating IP to an instance. This can be used instead of the
        `floating_ip` options in `compute.Instance`.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_floatingip_associate_v2.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fixed_ip: The specific IP address to direct traffic to.
        :param pulumi.Input[str] floating_ip: The floating IP to associate.
        :param pulumi.Input[str] instance_id: The instance to associte the floating IP with.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new floatingip_associate.
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

            __props__['fixed_ip'] = fixed_ip
            if floating_ip is None:
                raise TypeError("Missing required property 'floating_ip'")
            __props__['floating_ip'] = floating_ip
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['region'] = region
            __props__['wait_until_associated'] = wait_until_associated
        super(FloatingIpAssociate, __self__).__init__(
            'openstack:compute/floatingIpAssociate:FloatingIpAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, fixed_ip=None, floating_ip=None, instance_id=None, region=None, wait_until_associated=None):
        """
        Get an existing FloatingIpAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fixed_ip: The specific IP address to direct traffic to.
        :param pulumi.Input[str] floating_ip: The floating IP to associate.
        :param pulumi.Input[str] instance_id: The instance to associte the floating IP with.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new floatingip_associate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["fixed_ip"] = fixed_ip
        __props__["floating_ip"] = floating_ip
        __props__["instance_id"] = instance_id
        __props__["region"] = region
        __props__["wait_until_associated"] = wait_until_associated
        return FloatingIpAssociate(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

