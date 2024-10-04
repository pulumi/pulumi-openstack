# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['OrderV1Args', 'OrderV1']

@pulumi.input_type
class OrderV1Args:
    def __init__(__self__, *,
                 meta: pulumi.Input['OrderV1MetaArgs'],
                 type: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrderV1 resource.
        """
        pulumi.set(__self__, "meta", meta)
        pulumi.set(__self__, "type", type)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def meta(self) -> pulumi.Input['OrderV1MetaArgs']:
        return pulumi.get(self, "meta")

    @meta.setter
    def meta(self, value: pulumi.Input['OrderV1MetaArgs']):
        pulumi.set(self, "meta", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _OrderV1State:
    def __init__(__self__, *,
                 container_ref: Optional[pulumi.Input[str]] = None,
                 created: Optional[pulumi.Input[str]] = None,
                 creator_id: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input['OrderV1MetaArgs']] = None,
                 order_ref: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_ref: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 sub_status: Optional[pulumi.Input[str]] = None,
                 sub_status_message: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrderV1 resources.
        """
        if container_ref is not None:
            pulumi.set(__self__, "container_ref", container_ref)
        if created is not None:
            pulumi.set(__self__, "created", created)
        if creator_id is not None:
            pulumi.set(__self__, "creator_id", creator_id)
        if meta is not None:
            pulumi.set(__self__, "meta", meta)
        if order_ref is not None:
            pulumi.set(__self__, "order_ref", order_ref)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_ref is not None:
            pulumi.set(__self__, "secret_ref", secret_ref)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if sub_status is not None:
            pulumi.set(__self__, "sub_status", sub_status)
        if sub_status_message is not None:
            pulumi.set(__self__, "sub_status_message", sub_status_message)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated is not None:
            pulumi.set(__self__, "updated", updated)

    @property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container_ref")

    @container_ref.setter
    def container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_ref", value)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "creator_id")

    @creator_id.setter
    def creator_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creator_id", value)

    @property
    @pulumi.getter
    def meta(self) -> Optional[pulumi.Input['OrderV1MetaArgs']]:
        return pulumi.get(self, "meta")

    @meta.setter
    def meta(self, value: Optional[pulumi.Input['OrderV1MetaArgs']]):
        pulumi.set(self, "meta", value)

    @property
    @pulumi.getter(name="orderRef")
    def order_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "order_ref")

    @order_ref.setter
    def order_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "order_ref", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_ref")

    @secret_ref.setter
    def secret_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_ref", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subStatus")
    def sub_status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sub_status")

    @sub_status.setter
    def sub_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_status", value)

    @property
    @pulumi.getter(name="subStatusMessage")
    def sub_status_message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "sub_status_message")

    @sub_status_message.setter
    def sub_status_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_status_message", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def updated(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated")

    @updated.setter
    def updated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated", value)


class OrderV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 meta: Optional[pulumi.Input[Union['OrderV1MetaArgs', 'OrderV1MetaArgsDict']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a OrderV1 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrderV1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a OrderV1 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param OrderV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrderV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 meta: Optional[pulumi.Input[Union['OrderV1MetaArgs', 'OrderV1MetaArgsDict']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrderV1Args.__new__(OrderV1Args)

            if meta is None and not opts.urn:
                raise TypeError("Missing required property 'meta'")
            __props__.__dict__["meta"] = meta
            __props__.__dict__["region"] = region
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["container_ref"] = None
            __props__.__dict__["created"] = None
            __props__.__dict__["creator_id"] = None
            __props__.__dict__["order_ref"] = None
            __props__.__dict__["secret_ref"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["sub_status"] = None
            __props__.__dict__["sub_status_message"] = None
            __props__.__dict__["updated"] = None
        super(OrderV1, __self__).__init__(
            'openstack:keymanager/orderV1:OrderV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_ref: Optional[pulumi.Input[str]] = None,
            created: Optional[pulumi.Input[str]] = None,
            creator_id: Optional[pulumi.Input[str]] = None,
            meta: Optional[pulumi.Input[Union['OrderV1MetaArgs', 'OrderV1MetaArgsDict']]] = None,
            order_ref: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_ref: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            sub_status: Optional[pulumi.Input[str]] = None,
            sub_status_message: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated: Optional[pulumi.Input[str]] = None) -> 'OrderV1':
        """
        Get an existing OrderV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrderV1State.__new__(_OrderV1State)

        __props__.__dict__["container_ref"] = container_ref
        __props__.__dict__["created"] = created
        __props__.__dict__["creator_id"] = creator_id
        __props__.__dict__["meta"] = meta
        __props__.__dict__["order_ref"] = order_ref
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_ref"] = secret_ref
        __props__.__dict__["status"] = status
        __props__.__dict__["sub_status"] = sub_status
        __props__.__dict__["sub_status_message"] = sub_status_message
        __props__.__dict__["type"] = type
        __props__.__dict__["updated"] = updated
        return OrderV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "container_ref")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creator_id")

    @property
    @pulumi.getter
    def meta(self) -> pulumi.Output['outputs.OrderV1Meta']:
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter(name="orderRef")
    def order_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "order_ref")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret_ref")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subStatus")
    def sub_status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sub_status")

    @property
    @pulumi.getter(name="subStatusMessage")
    def sub_status_message(self) -> pulumi.Output[str]:
        return pulumi.get(self, "sub_status_message")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated")

