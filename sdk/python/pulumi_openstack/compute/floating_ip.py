# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['FloatingIp']


class FloatingIp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pool: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V2 floating IP resource within OpenStack Nova (compute)
        that can be used for compute instances.

        Please note that managing floating IPs through the OpenStack Compute API has
        been deprecated. Unless you are using an older OpenStack environment, it is
        recommended to use the `networking.FloatingIp`
        resource instead, which uses the OpenStack Networking API.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        floatip1 = openstack.compute.FloatingIp("floatip1", pool="public")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] pool: The name of the pool from which to obtain the floating
               IP. Changing this creates a new floating IP.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a floating IP that can be used with
               a compute instance. If omitted, the `region` argument of the provider
               is used. Changing this creates a new floating IP (which may or may not
               have a different address).
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if pool is None:
                raise TypeError("Missing required property 'pool'")
            __props__['pool'] = pool
            __props__['region'] = region
            __props__['address'] = None
            __props__['fixed_ip'] = None
            __props__['instance_id'] = None
        super(FloatingIp, __self__).__init__(
            'openstack:compute/floatingIp:FloatingIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            fixed_ip: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            pool: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'FloatingIp':
        """
        Get an existing FloatingIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The actual floating IP address itself.
        :param pulumi.Input[str] fixed_ip: The fixed IP address corresponding to the floating IP.
        :param pulumi.Input[str] instance_id: UUID of the compute instance associated with the floating IP.
        :param pulumi.Input[str] pool: The name of the pool from which to obtain the floating
               IP. Changing this creates a new floating IP.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a floating IP that can be used with
               a compute instance. If omitted, the `region` argument of the provider
               is used. Changing this creates a new floating IP (which may or may not
               have a different address).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address"] = address
        __props__["fixed_ip"] = fixed_ip
        __props__["instance_id"] = instance_id
        __props__["pool"] = pool
        __props__["region"] = region
        return FloatingIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The actual floating IP address itself.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> pulumi.Output[str]:
        """
        The fixed IP address corresponding to the floating IP.
        """
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the compute instance associated with the floating IP.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def pool(self) -> pulumi.Output[str]:
        """
        The name of the pool from which to obtain the floating
        IP. Changing this creates a new floating IP.
        """
        return pulumi.get(self, "pool")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        A Compute client is needed to create a floating IP that can be used with
        a compute instance. If omitted, the `region` argument of the provider
        is used. Changing this creates a new floating IP (which may or may not
        have a different address).
        """
        return pulumi.get(self, "region")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

