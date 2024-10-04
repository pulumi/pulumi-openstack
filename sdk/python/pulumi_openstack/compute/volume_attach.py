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

__all__ = ['VolumeAttachArgs', 'VolumeAttach']

@pulumi.input_type
class VolumeAttachArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 volume_id: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None,
                 multiattach: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vendor_options: Optional[pulumi.Input['VolumeAttachVendorOptionsArgs']] = None):
        """
        The set of arguments for constructing a VolumeAttach resource.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "volume_id", volume_id)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if multiattach is not None:
            pulumi.set(__self__, "multiattach", multiattach)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if vendor_options is not None:
            pulumi.set(__self__, "vendor_options", vendor_options)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "volume_id", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter
    def multiattach(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multiattach")

    @multiattach.setter
    def multiattach(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multiattach", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter(name="vendorOptions")
    def vendor_options(self) -> Optional[pulumi.Input['VolumeAttachVendorOptionsArgs']]:
        return pulumi.get(self, "vendor_options")

    @vendor_options.setter
    def vendor_options(self, value: Optional[pulumi.Input['VolumeAttachVendorOptionsArgs']]):
        pulumi.set(self, "vendor_options", value)


@pulumi.input_type
class _VolumeAttachState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 multiattach: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vendor_options: Optional[pulumi.Input['VolumeAttachVendorOptionsArgs']] = None,
                 volume_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VolumeAttach resources.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if multiattach is not None:
            pulumi.set(__self__, "multiattach", multiattach)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)
        if vendor_options is not None:
            pulumi.set(__self__, "vendor_options", vendor_options)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def multiattach(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multiattach")

    @multiattach.setter
    def multiattach(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multiattach", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)

    @property
    @pulumi.getter(name="vendorOptions")
    def vendor_options(self) -> Optional[pulumi.Input['VolumeAttachVendorOptionsArgs']]:
        return pulumi.get(self, "vendor_options")

    @vendor_options.setter
    def vendor_options(self, value: Optional[pulumi.Input['VolumeAttachVendorOptionsArgs']]):
        pulumi.set(self, "vendor_options", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_id", value)


class VolumeAttach(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 multiattach: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vendor_options: Optional[pulumi.Input[Union['VolumeAttachVendorOptionsArgs', 'VolumeAttachVendorOptionsArgsDict']]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VolumeAttach resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeAttachArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VolumeAttach resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VolumeAttachArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeAttachArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 multiattach: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 vendor_options: Optional[pulumi.Input[Union['VolumeAttachVendorOptionsArgs', 'VolumeAttachVendorOptionsArgsDict']]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeAttachArgs.__new__(VolumeAttachArgs)

            __props__.__dict__["device"] = device
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["multiattach"] = multiattach
            __props__.__dict__["region"] = region
            __props__.__dict__["tag"] = tag
            __props__.__dict__["vendor_options"] = vendor_options
            if volume_id is None and not opts.urn:
                raise TypeError("Missing required property 'volume_id'")
            __props__.__dict__["volume_id"] = volume_id
        super(VolumeAttach, __self__).__init__(
            'openstack:compute/volumeAttach:VolumeAttach',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            multiattach: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tag: Optional[pulumi.Input[str]] = None,
            vendor_options: Optional[pulumi.Input[Union['VolumeAttachVendorOptionsArgs', 'VolumeAttachVendorOptionsArgsDict']]] = None,
            volume_id: Optional[pulumi.Input[str]] = None) -> 'VolumeAttach':
        """
        Get an existing VolumeAttach resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VolumeAttachState.__new__(_VolumeAttachState)

        __props__.__dict__["device"] = device
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["multiattach"] = multiattach
        __props__.__dict__["region"] = region
        __props__.__dict__["tag"] = tag
        __props__.__dict__["vendor_options"] = vendor_options
        __props__.__dict__["volume_id"] = volume_id
        return VolumeAttach(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def multiattach(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "multiattach")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tag")

    @property
    @pulumi.getter(name="vendorOptions")
    def vendor_options(self) -> pulumi.Output[Optional['outputs.VolumeAttachVendorOptions']]:
        return pulumi.get(self, "vendor_options")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "volume_id")

