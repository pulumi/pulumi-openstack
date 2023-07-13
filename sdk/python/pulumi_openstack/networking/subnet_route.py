# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SubnetRouteArgs', 'SubnetRoute']

@pulumi.input_type
class SubnetRouteArgs:
    def __init__(__self__, *,
                 destination_cidr: pulumi.Input[str],
                 next_hop: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SubnetRoute resource.
        """
        pulumi.set(__self__, "destination_cidr", destination_cidr)
        pulumi.set(__self__, "next_hop", next_hop)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> pulumi.Input[str]:
        return pulumi.get(self, "destination_cidr")

    @destination_cidr.setter
    def destination_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr", value)

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> pulumi.Input[str]:
        return pulumi.get(self, "next_hop")

    @next_hop.setter
    def next_hop(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_hop", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _SubnetRouteState:
    def __init__(__self__, *,
                 destination_cidr: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SubnetRoute resources.
        """
        if destination_cidr is not None:
            pulumi.set(__self__, "destination_cidr", destination_cidr)
        if next_hop is not None:
            pulumi.set(__self__, "next_hop", next_hop)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination_cidr")

    @destination_cidr.setter
    def destination_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr", value)

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "next_hop")

    @next_hop.setter
    def next_hop(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class SubnetRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SubnetRoute resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubnetRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SubnetRoute resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SubnetRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr: Optional[pulumi.Input[str]] = None,
                 next_hop: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetRouteArgs.__new__(SubnetRouteArgs)

            if destination_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr'")
            __props__.__dict__["destination_cidr"] = destination_cidr
            if next_hop is None and not opts.urn:
                raise TypeError("Missing required property 'next_hop'")
            __props__.__dict__["next_hop"] = next_hop
            __props__.__dict__["region"] = region
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
        super(SubnetRoute, __self__).__init__(
            'openstack:networking/subnetRoute:SubnetRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_cidr: Optional[pulumi.Input[str]] = None,
            next_hop: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'SubnetRoute':
        """
        Get an existing SubnetRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubnetRouteState.__new__(_SubnetRouteState)

        __props__.__dict__["destination_cidr"] = destination_cidr
        __props__.__dict__["next_hop"] = next_hop
        __props__.__dict__["region"] = region
        __props__.__dict__["subnet_id"] = subnet_id
        return SubnetRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "destination_cidr")

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> pulumi.Output[str]:
        return pulumi.get(self, "next_hop")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subnet_id")

