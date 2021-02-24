# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Instance']


class Instance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_ip_v4: Optional[pulumi.Input[str]] = None,
                 access_ip_v6: Optional[pulumi.Input[str]] = None,
                 admin_pass: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 availability_zone_hints: Optional[pulumi.Input[str]] = None,
                 block_devices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBlockDeviceArgs']]]]] = None,
                 config_drive: Optional[pulumi.Input[bool]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 flavor_name: Optional[pulumi.Input[str]] = None,
                 force_delete: Optional[pulumi.Input[bool]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 key_pair: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_mode: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceNetworkArgs']]]]] = None,
                 personalities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstancePersonalityArgs']]]]] = None,
                 power_state: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 scheduler_hints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSchedulerHintArgs']]]]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 stop_before_destroy: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_data: Optional[pulumi.Input[str]] = None,
                 vendor_options: Optional[pulumi.Input[pulumi.InputType['InstanceVendorOptionsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Instance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_ip_v4: The first detected Fixed IPv4 address.
        :param pulumi.Input[str] access_ip_v6: The first detected Fixed IPv6 address.
        :param pulumi.Input[str] admin_pass: The administrative password to assign to the server.
               Changing this changes the root password on the existing server.
        :param pulumi.Input[str] availability_zone: The availability zone in which to create
               the server. Conflicts with `availability_zone_hints`. Changing this creates
               a new server.
        :param pulumi.Input[str] availability_zone_hints: The availability zone in which to
               create the server. This argument is preferred to `availability_zone`, when
               scheduling the server on a
               [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
               host or node. Conflicts with `availability_zone`. Changing this creates a
               new server.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBlockDeviceArgs']]]] block_devices: Configuration of block devices. The block_device
               structure is documented below. Changing this creates a new server.
               You can specify multiple block devices which will create an instance with
               multiple disks. This configuration is very flexible, so please see the
               following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
               for more information.
        :param pulumi.Input[bool] config_drive: Whether to use the config_drive feature to
               configure the instance. Changing this creates a new server.
        :param pulumi.Input[str] flavor_id: The flavor ID of
               the desired flavor for the server. Changing this resizes the existing server.
        :param pulumi.Input[str] flavor_name: The name of the
               desired flavor for the server. Changing this resizes the existing server.
        :param pulumi.Input[bool] force_delete: Whether to force the OpenStack instance to be
               forcefully deleted. This is useful for environments that have reclaim / soft
               deletion enabled.
        :param pulumi.Input[str] image_id: (Optional; Required if `image_name` is empty and not booting
               from a volume. Do not specify if booting from a volume.) The image ID of
               the desired image for the server. Changing this creates a new server.
        :param pulumi.Input[str] image_name: (Optional; Required if `image_id` is empty and not booting
               from a volume. Do not specify if booting from a volume.) The name of the
               desired image for the server. Changing this creates a new server.
        :param pulumi.Input[str] key_pair: The name of a key pair to put on the server. The key
               pair must already be created and associated with the tenant's account.
               Changing this creates a new server.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to make available from
               within the instance. Changing this updates the existing server metadata.
        :param pulumi.Input[str] name: The human-readable
               name of the network. Changing this creates a new server.
        :param pulumi.Input[str] network_mode: Special string for `network` option to create
               the server. `network_mode` can be `"auto"` or `"none"`.
               Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceNetworkArgs']]]] networks: An array of one or more networks to attach to the
               instance. The network object structure is documented below. Changing this
               creates a new server.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstancePersonalityArgs']]]] personalities: Customize the personality of an instance by
               defining one or more files and their contents. The personality structure
               is described below.
        :param pulumi.Input[str] power_state: Provide the VM state. Only 'active' and 'shutoff'
               are supported values. *Note*: If the initial power_state is the shutoff
               the VM will be stopped immediately after build and the provisioners like
               remote-exec or files are not supported.
        :param pulumi.Input[str] region: The region in which to create the server instance. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new server.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSchedulerHintArgs']]]] scheduler_hints: Provide the Nova scheduler with hints on how
               the instance should be launched. The available hints are described below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: An array of one or more security group names
               to associate with the server. Changing this results in adding/removing
               security groups from the existing server. *Note*: When attaching the
               instance to networks using Ports, place the security groups on the Port
               and not the instance. *Note*: Names should be used and not ids, as ids
               trigger unnecessary updates.
        :param pulumi.Input[bool] stop_before_destroy: Whether to try stop instance gracefully
               before destroying it, thus giving chance for guest OS daemons to stop correctly.
               If instance doesn't stop within timeout, it will be destroyed anyway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the instance. Changing this
               updates the existing instance tags.
        :param pulumi.Input[str] user_data: The user data to provide when launching the instance.
               Changing this creates a new server.
        :param pulumi.Input[pulumi.InputType['InstanceVendorOptionsArgs']] vendor_options: Map of additional vendor-specific options.
               Supported options are described below.
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

            __props__['access_ip_v4'] = access_ip_v4
            __props__['access_ip_v6'] = access_ip_v6
            __props__['admin_pass'] = admin_pass
            __props__['availability_zone'] = availability_zone
            __props__['availability_zone_hints'] = availability_zone_hints
            __props__['block_devices'] = block_devices
            __props__['config_drive'] = config_drive
            __props__['flavor_id'] = flavor_id
            __props__['flavor_name'] = flavor_name
            __props__['force_delete'] = force_delete
            __props__['image_id'] = image_id
            __props__['image_name'] = image_name
            __props__['key_pair'] = key_pair
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['network_mode'] = network_mode
            __props__['networks'] = networks
            __props__['personalities'] = personalities
            __props__['power_state'] = power_state
            __props__['region'] = region
            __props__['scheduler_hints'] = scheduler_hints
            __props__['security_groups'] = security_groups
            __props__['stop_before_destroy'] = stop_before_destroy
            __props__['tags'] = tags
            __props__['user_data'] = user_data
            __props__['vendor_options'] = vendor_options
            __props__['all_metadata'] = None
            __props__['all_tags'] = None
        super(Instance, __self__).__init__(
            'openstack:compute/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_ip_v4: Optional[pulumi.Input[str]] = None,
            access_ip_v6: Optional[pulumi.Input[str]] = None,
            admin_pass: Optional[pulumi.Input[str]] = None,
            all_metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            all_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            availability_zone_hints: Optional[pulumi.Input[str]] = None,
            block_devices: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBlockDeviceArgs']]]]] = None,
            config_drive: Optional[pulumi.Input[bool]] = None,
            flavor_id: Optional[pulumi.Input[str]] = None,
            flavor_name: Optional[pulumi.Input[str]] = None,
            force_delete: Optional[pulumi.Input[bool]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            image_name: Optional[pulumi.Input[str]] = None,
            key_pair: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_mode: Optional[pulumi.Input[str]] = None,
            networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceNetworkArgs']]]]] = None,
            personalities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstancePersonalityArgs']]]]] = None,
            power_state: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            scheduler_hints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSchedulerHintArgs']]]]] = None,
            security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            stop_before_destroy: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user_data: Optional[pulumi.Input[str]] = None,
            vendor_options: Optional[pulumi.Input[pulumi.InputType['InstanceVendorOptionsArgs']]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_ip_v4: The first detected Fixed IPv4 address.
        :param pulumi.Input[str] access_ip_v6: The first detected Fixed IPv6 address.
        :param pulumi.Input[str] admin_pass: The administrative password to assign to the server.
               Changing this changes the root password on the existing server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] all_tags: The collection of tags assigned on the instance, which have
               been explicitly and implicitly added.
        :param pulumi.Input[str] availability_zone: The availability zone in which to create
               the server. Conflicts with `availability_zone_hints`. Changing this creates
               a new server.
        :param pulumi.Input[str] availability_zone_hints: The availability zone in which to
               create the server. This argument is preferred to `availability_zone`, when
               scheduling the server on a
               [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
               host or node. Conflicts with `availability_zone`. Changing this creates a
               new server.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceBlockDeviceArgs']]]] block_devices: Configuration of block devices. The block_device
               structure is documented below. Changing this creates a new server.
               You can specify multiple block devices which will create an instance with
               multiple disks. This configuration is very flexible, so please see the
               following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
               for more information.
        :param pulumi.Input[bool] config_drive: Whether to use the config_drive feature to
               configure the instance. Changing this creates a new server.
        :param pulumi.Input[str] flavor_id: The flavor ID of
               the desired flavor for the server. Changing this resizes the existing server.
        :param pulumi.Input[str] flavor_name: The name of the
               desired flavor for the server. Changing this resizes the existing server.
        :param pulumi.Input[bool] force_delete: Whether to force the OpenStack instance to be
               forcefully deleted. This is useful for environments that have reclaim / soft
               deletion enabled.
        :param pulumi.Input[str] image_id: (Optional; Required if `image_name` is empty and not booting
               from a volume. Do not specify if booting from a volume.) The image ID of
               the desired image for the server. Changing this creates a new server.
        :param pulumi.Input[str] image_name: (Optional; Required if `image_id` is empty and not booting
               from a volume. Do not specify if booting from a volume.) The name of the
               desired image for the server. Changing this creates a new server.
        :param pulumi.Input[str] key_pair: The name of a key pair to put on the server. The key
               pair must already be created and associated with the tenant's account.
               Changing this creates a new server.
        :param pulumi.Input[Mapping[str, Any]] metadata: Metadata key/value pairs to make available from
               within the instance. Changing this updates the existing server metadata.
        :param pulumi.Input[str] name: The human-readable
               name of the network. Changing this creates a new server.
        :param pulumi.Input[str] network_mode: Special string for `network` option to create
               the server. `network_mode` can be `"auto"` or `"none"`.
               Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceNetworkArgs']]]] networks: An array of one or more networks to attach to the
               instance. The network object structure is documented below. Changing this
               creates a new server.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstancePersonalityArgs']]]] personalities: Customize the personality of an instance by
               defining one or more files and their contents. The personality structure
               is described below.
        :param pulumi.Input[str] power_state: Provide the VM state. Only 'active' and 'shutoff'
               are supported values. *Note*: If the initial power_state is the shutoff
               the VM will be stopped immediately after build and the provisioners like
               remote-exec or files are not supported.
        :param pulumi.Input[str] region: The region in which to create the server instance. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new server.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSchedulerHintArgs']]]] scheduler_hints: Provide the Nova scheduler with hints on how
               the instance should be launched. The available hints are described below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: An array of one or more security group names
               to associate with the server. Changing this results in adding/removing
               security groups from the existing server. *Note*: When attaching the
               instance to networks using Ports, place the security groups on the Port
               and not the instance. *Note*: Names should be used and not ids, as ids
               trigger unnecessary updates.
        :param pulumi.Input[bool] stop_before_destroy: Whether to try stop instance gracefully
               before destroying it, thus giving chance for guest OS daemons to stop correctly.
               If instance doesn't stop within timeout, it will be destroyed anyway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the instance. Changing this
               updates the existing instance tags.
        :param pulumi.Input[str] user_data: The user data to provide when launching the instance.
               Changing this creates a new server.
        :param pulumi.Input[pulumi.InputType['InstanceVendorOptionsArgs']] vendor_options: Map of additional vendor-specific options.
               Supported options are described below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_ip_v4"] = access_ip_v4
        __props__["access_ip_v6"] = access_ip_v6
        __props__["admin_pass"] = admin_pass
        __props__["all_metadata"] = all_metadata
        __props__["all_tags"] = all_tags
        __props__["availability_zone"] = availability_zone
        __props__["availability_zone_hints"] = availability_zone_hints
        __props__["block_devices"] = block_devices
        __props__["config_drive"] = config_drive
        __props__["flavor_id"] = flavor_id
        __props__["flavor_name"] = flavor_name
        __props__["force_delete"] = force_delete
        __props__["image_id"] = image_id
        __props__["image_name"] = image_name
        __props__["key_pair"] = key_pair
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["network_mode"] = network_mode
        __props__["networks"] = networks
        __props__["personalities"] = personalities
        __props__["power_state"] = power_state
        __props__["region"] = region
        __props__["scheduler_hints"] = scheduler_hints
        __props__["security_groups"] = security_groups
        __props__["stop_before_destroy"] = stop_before_destroy
        __props__["tags"] = tags
        __props__["user_data"] = user_data
        __props__["vendor_options"] = vendor_options
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessIpV4")
    def access_ip_v4(self) -> pulumi.Output[str]:
        """
        The first detected Fixed IPv4 address.
        """
        return pulumi.get(self, "access_ip_v4")

    @property
    @pulumi.getter(name="accessIpV6")
    def access_ip_v6(self) -> pulumi.Output[str]:
        """
        The first detected Fixed IPv6 address.
        """
        return pulumi.get(self, "access_ip_v6")

    @property
    @pulumi.getter(name="adminPass")
    def admin_pass(self) -> pulumi.Output[Optional[str]]:
        """
        The administrative password to assign to the server.
        Changing this changes the root password on the existing server.
        """
        return pulumi.get(self, "admin_pass")

    @property
    @pulumi.getter(name="allMetadata")
    def all_metadata(self) -> pulumi.Output[Mapping[str, Any]]:
        return pulumi.get(self, "all_metadata")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> pulumi.Output[Sequence[str]]:
        """
        The collection of tags assigned on the instance, which have
        been explicitly and implicitly added.
        """
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The availability zone in which to create
        the server. Conflicts with `availability_zone_hints`. Changing this creates
        a new server.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="availabilityZoneHints")
    def availability_zone_hints(self) -> pulumi.Output[Optional[str]]:
        """
        The availability zone in which to
        create the server. This argument is preferred to `availability_zone`, when
        scheduling the server on a
        [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
        host or node. Conflicts with `availability_zone`. Changing this creates a
        new server.
        """
        return pulumi.get(self, "availability_zone_hints")

    @property
    @pulumi.getter(name="blockDevices")
    def block_devices(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceBlockDevice']]]:
        """
        Configuration of block devices. The block_device
        structure is documented below. Changing this creates a new server.
        You can specify multiple block devices which will create an instance with
        multiple disks. This configuration is very flexible, so please see the
        following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
        for more information.
        """
        return pulumi.get(self, "block_devices")

    @property
    @pulumi.getter(name="configDrive")
    def config_drive(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to use the config_drive feature to
        configure the instance. Changing this creates a new server.
        """
        return pulumi.get(self, "config_drive")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> pulumi.Output[str]:
        """
        The flavor ID of
        the desired flavor for the server. Changing this resizes the existing server.
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="flavorName")
    def flavor_name(self) -> pulumi.Output[str]:
        """
        The name of the
        desired flavor for the server. Changing this resizes the existing server.
        """
        return pulumi.get(self, "flavor_name")

    @property
    @pulumi.getter(name="forceDelete")
    def force_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to force the OpenStack instance to be
        forcefully deleted. This is useful for environments that have reclaim / soft
        deletion enabled.
        """
        return pulumi.get(self, "force_delete")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        (Optional; Required if `image_name` is empty and not booting
        from a volume. Do not specify if booting from a volume.) The image ID of
        the desired image for the server. Changing this creates a new server.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Output[str]:
        """
        (Optional; Required if `image_id` is empty and not booting
        from a volume. Do not specify if booting from a volume.) The name of the
        desired image for the server. Changing this creates a new server.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> pulumi.Output[Optional[str]]:
        """
        The name of a key pair to put on the server. The key
        pair must already be created and associated with the tenant's account.
        Changing this creates a new server.
        """
        return pulumi.get(self, "key_pair")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Metadata key/value pairs to make available from
        within the instance. Changing this updates the existing server metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The human-readable
        name of the network. Changing this creates a new server.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkMode")
    def network_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Special string for `network` option to create
        the server. `network_mode` can be `"auto"` or `"none"`.
        Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
        """
        return pulumi.get(self, "network_mode")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Sequence['outputs.InstanceNetwork']]:
        """
        An array of one or more networks to attach to the
        instance. The network object structure is documented below. Changing this
        creates a new server.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter
    def personalities(self) -> pulumi.Output[Optional[Sequence['outputs.InstancePersonality']]]:
        """
        Customize the personality of an instance by
        defining one or more files and their contents. The personality structure
        is described below.
        """
        return pulumi.get(self, "personalities")

    @property
    @pulumi.getter(name="powerState")
    def power_state(self) -> pulumi.Output[Optional[str]]:
        """
        Provide the VM state. Only 'active' and 'shutoff'
        are supported values. *Note*: If the initial power_state is the shutoff
        the VM will be stopped immediately after build and the provisioners like
        remote-exec or files are not supported.
        """
        return pulumi.get(self, "power_state")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the server instance. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new server.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="schedulerHints")
    def scheduler_hints(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceSchedulerHint']]]:
        """
        Provide the Nova scheduler with hints on how
        the instance should be launched. The available hints are described below.
        """
        return pulumi.get(self, "scheduler_hints")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of one or more security group names
        to associate with the server. Changing this results in adding/removing
        security groups from the existing server. *Note*: When attaching the
        instance to networks using Ports, place the security groups on the Port
        and not the instance. *Note*: Names should be used and not ids, as ids
        trigger unnecessary updates.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="stopBeforeDestroy")
    def stop_before_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to try stop instance gracefully
        before destroying it, thus giving chance for guest OS daemons to stop correctly.
        If instance doesn't stop within timeout, it will be destroyed anyway.
        """
        return pulumi.get(self, "stop_before_destroy")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of string tags for the instance. Changing this
        updates the existing instance tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> pulumi.Output[Optional[str]]:
        """
        The user data to provide when launching the instance.
        Changing this creates a new server.
        """
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter(name="vendorOptions")
    def vendor_options(self) -> pulumi.Output[Optional['outputs.InstanceVendorOptions']]:
        """
        Map of additional vendor-specific options.
        Supported options are described below.
        """
        return pulumi.get(self, "vendor_options")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

