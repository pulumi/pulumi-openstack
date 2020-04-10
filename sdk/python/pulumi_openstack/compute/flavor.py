# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Flavor(pulumi.CustomResource):
    disk: pulumi.Output[float]
    """
    The amount of disk space in gigabytes to use for the root
    (/) partition. Changing this creates a new flavor.
    """
    ephemeral: pulumi.Output[float]
    extra_specs: pulumi.Output[dict]
    """
    Key/Value pairs of metadata for the flavor.
    """
    is_public: pulumi.Output[bool]
    """
    Whether the flavor is public. Changing this creates
    a new flavor.
    """
    name: pulumi.Output[str]
    """
    A unique name for the flavor. Changing this creates a new
    flavor.
    """
    ram: pulumi.Output[float]
    """
    The amount of RAM to use, in megabytes. Changing this
    creates a new flavor.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Compute client.
    Flavors are associated with accounts, but a Compute client is needed to
    create one. If omitted, the `region` argument of the provider is used.
    Changing this creates a new flavor.
    """
    rx_tx_factor: pulumi.Output[float]
    """
    RX/TX bandwith factor. The default is 1. Changing
    this creates a new flavor.
    """
    swap: pulumi.Output[float]
    """
    The amount of disk space in megabytes to use. If
    unspecified, the default is 0. Changing this creates a new flavor.
    """
    vcpus: pulumi.Output[float]
    """
    The number of virtual CPUs to use. Changing this creates
    a new flavor.
    """
    def __init__(__self__, resource_name, opts=None, disk=None, ephemeral=None, extra_specs=None, is_public=None, name=None, ram=None, region=None, rx_tx_factor=None, swap=None, vcpus=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V2 flavor resource within OpenStack.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] disk: The amount of disk space in gigabytes to use for the root
               (/) partition. Changing this creates a new flavor.
        :param pulumi.Input[dict] extra_specs: Key/Value pairs of metadata for the flavor.
        :param pulumi.Input[bool] is_public: Whether the flavor is public. Changing this creates
               a new flavor.
        :param pulumi.Input[str] name: A unique name for the flavor. Changing this creates a new
               flavor.
        :param pulumi.Input[float] ram: The amount of RAM to use, in megabytes. Changing this
               creates a new flavor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Flavors are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor.
        :param pulumi.Input[float] rx_tx_factor: RX/TX bandwith factor. The default is 1. Changing
               this creates a new flavor.
        :param pulumi.Input[float] swap: The amount of disk space in megabytes to use. If
               unspecified, the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[float] vcpus: The number of virtual CPUs to use. Changing this creates
               a new flavor.
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

            if disk is None:
                raise TypeError("Missing required property 'disk'")
            __props__['disk'] = disk
            __props__['ephemeral'] = ephemeral
            __props__['extra_specs'] = extra_specs
            __props__['is_public'] = is_public
            __props__['name'] = name
            if ram is None:
                raise TypeError("Missing required property 'ram'")
            __props__['ram'] = ram
            __props__['region'] = region
            __props__['rx_tx_factor'] = rx_tx_factor
            __props__['swap'] = swap
            if vcpus is None:
                raise TypeError("Missing required property 'vcpus'")
            __props__['vcpus'] = vcpus
        super(Flavor, __self__).__init__(
            'openstack:compute/flavor:Flavor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, disk=None, ephemeral=None, extra_specs=None, is_public=None, name=None, ram=None, region=None, rx_tx_factor=None, swap=None, vcpus=None):
        """
        Get an existing Flavor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] disk: The amount of disk space in gigabytes to use for the root
               (/) partition. Changing this creates a new flavor.
        :param pulumi.Input[dict] extra_specs: Key/Value pairs of metadata for the flavor.
        :param pulumi.Input[bool] is_public: Whether the flavor is public. Changing this creates
               a new flavor.
        :param pulumi.Input[str] name: A unique name for the flavor. Changing this creates a new
               flavor.
        :param pulumi.Input[float] ram: The amount of RAM to use, in megabytes. Changing this
               creates a new flavor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Flavors are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor.
        :param pulumi.Input[float] rx_tx_factor: RX/TX bandwith factor. The default is 1. Changing
               this creates a new flavor.
        :param pulumi.Input[float] swap: The amount of disk space in megabytes to use. If
               unspecified, the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[float] vcpus: The number of virtual CPUs to use. Changing this creates
               a new flavor.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["disk"] = disk
        __props__["ephemeral"] = ephemeral
        __props__["extra_specs"] = extra_specs
        __props__["is_public"] = is_public
        __props__["name"] = name
        __props__["ram"] = ram
        __props__["region"] = region
        __props__["rx_tx_factor"] = rx_tx_factor
        __props__["swap"] = swap
        __props__["vcpus"] = vcpus
        return Flavor(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

