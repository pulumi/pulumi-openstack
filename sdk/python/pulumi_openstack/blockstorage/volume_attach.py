# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VolumeAttach(pulumi.CustomResource):
    attach_mode: pulumi.Output[str]
    """
    Specify whether to attach the volume as Read-Only
    (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
    If left unspecified, the Block Storage API will apply a default of `rw`.
    """
    data: pulumi.Output[dict]
    """
    This is a map of key/value pairs that contain the connection
    information. You will want to pass this information to a provisioner
    script to finalize the connection. See below for more information.
    """
    device: pulumi.Output[str]
    """
    The device to tell the Block Storage service this
    volume will be attached as. This is purely for informational purposes.
    You can specify `auto` or a device such as `/dev/vdc`.
    """
    driver_volume_type: pulumi.Output[str]
    """
    The storage driver that the volume is based on.
    """
    host_name: pulumi.Output[str]
    """
    The host to attach the volume to.
    """
    initiator: pulumi.Output[str]
    """
    The iSCSI initiator string to make the connection.
    """
    ip_address: pulumi.Output[str]
    """
    The IP address of the `host_name` above.
    """
    mount_point_base: pulumi.Output[str]
    """
    A mount point base name for shared storage.
    """
    multipath: pulumi.Output[bool]
    """
    Whether to connect to this volume via multipath.
    """
    os_type: pulumi.Output[str]
    """
    The iSCSI initiator OS type.
    """
    platform: pulumi.Output[str]
    """
    The iSCSI initiator platform.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V3 Block Storage
    client. A Block Storage client is needed to create a volume attachment.
    If omitted, the `region` argument of the provider is used. Changing this
    creates a new volume attachment.
    """
    volume_id: pulumi.Output[str]
    """
    The ID of the Volume to attach to an Instance.
    """
    wwnn: pulumi.Output[str]
    """
    A wwnn name. Used for Fibre Channel connections.
    """
    wwpns: pulumi.Output[list]
    """
    An array of wwpn strings. Used for Fibre Channel
    connections.
    """
    def __init__(__self__, resource_name, opts=None, attach_mode=None, device=None, host_name=None, initiator=None, ip_address=None, multipath=None, os_type=None, platform=None, region=None, volume_id=None, wwnn=None, wwpns=None, __name__=None, __opts__=None):
        """
        This resource is experimental and may be removed in the future! Feedback
        is requested if you find this resource useful or if you find any problems
        with it.
        
        Creates a general purpose attachment connection to a Block
        Storage volume using the OpenStack Block Storage (Cinder) v3 API.
        Depending on your Block Storage service configuration, this
        resource can assist in attaching a volume to a non-OpenStack resource
        such as a bare-metal server or a remote virtual machine in a
        different cloud provider.
        
        This does not actually attach a volume to an instance. Please use
        the `openstack_compute_volume_attach_v3` resource for that.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attach_mode: Specify whether to attach the volume as Read-Only
               (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
               If left unspecified, the Block Storage API will apply a default of `rw`.
        :param pulumi.Input[str] device: The device to tell the Block Storage service this
               volume will be attached as. This is purely for informational purposes.
               You can specify `auto` or a device such as `/dev/vdc`.
        :param pulumi.Input[str] host_name: The host to attach the volume to.
        :param pulumi.Input[str] initiator: The iSCSI initiator string to make the connection.
        :param pulumi.Input[str] ip_address: The IP address of the `host_name` above.
        :param pulumi.Input[bool] multipath: Whether to connect to this volume via multipath.
        :param pulumi.Input[str] os_type: The iSCSI initiator OS type.
        :param pulumi.Input[str] platform: The iSCSI initiator platform.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Block Storage
               client. A Block Storage client is needed to create a volume attachment.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new volume attachment.
        :param pulumi.Input[str] volume_id: The ID of the Volume to attach to an Instance.
        :param pulumi.Input[str] wwnn: A wwnn name. Used for Fibre Channel connections.
        :param pulumi.Input[list] wwpns: An array of wwpn strings. Used for Fibre Channel
               connections.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/blockstorage_volume_attach_v3.html.markdown.
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

        __props__['attach_mode'] = attach_mode

        __props__['device'] = device

        if host_name is None:
            raise TypeError("Missing required property 'host_name'")
        __props__['host_name'] = host_name

        __props__['initiator'] = initiator

        __props__['ip_address'] = ip_address

        __props__['multipath'] = multipath

        __props__['os_type'] = os_type

        __props__['platform'] = platform

        __props__['region'] = region

        if volume_id is None:
            raise TypeError("Missing required property 'volume_id'")
        __props__['volume_id'] = volume_id

        __props__['wwnn'] = wwnn

        __props__['wwpns'] = wwpns

        __props__['data'] = None
        __props__['driver_volume_type'] = None
        __props__['mount_point_base'] = None

        super(VolumeAttach, __self__).__init__(
            'openstack:blockstorage/volumeAttach:VolumeAttach',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

