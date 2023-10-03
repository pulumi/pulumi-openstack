# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FloatingIpAssociateArgs', 'FloatingIpAssociate']

@pulumi.input_type
class FloatingIpAssociateArgs:
    def __init__(__self__, *,
                 floating_ip: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 wait_until_associated: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a FloatingIpAssociate resource.
        :param pulumi.Input[str] floating_ip: The floating IP to associate.
        :param pulumi.Input[str] instance_id: The instance to associte the floating IP with.
        :param pulumi.Input[str] fixed_ip: The specific IP address to direct traffic to.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new floatingip_associate.
        """
        FloatingIpAssociateArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            floating_ip=floating_ip,
            instance_id=instance_id,
            fixed_ip=fixed_ip,
            region=region,
            wait_until_associated=wait_until_associated,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             floating_ip: pulumi.Input[str],
             instance_id: pulumi.Input[str],
             fixed_ip: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             wait_until_associated: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("floating_ip", floating_ip)
        _setter("instance_id", instance_id)
        if fixed_ip is not None:
            _setter("fixed_ip", fixed_ip)
        if region is not None:
            _setter("region", region)
        if wait_until_associated is not None:
            _setter("wait_until_associated", wait_until_associated)

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> pulumi.Input[str]:
        """
        The floating IP to associate.
        """
        return pulumi.get(self, "floating_ip")

    @floating_ip.setter
    def floating_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "floating_ip", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The instance to associte the floating IP with.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The specific IP address to direct traffic to.
        """
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new floatingip_associate.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="waitUntilAssociated")
    def wait_until_associated(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "wait_until_associated")

    @wait_until_associated.setter
    def wait_until_associated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_associated", value)


@pulumi.input_type
class _FloatingIpAssociateState:
    def __init__(__self__, *,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 floating_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 wait_until_associated: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering FloatingIpAssociate resources.
        :param pulumi.Input[str] fixed_ip: The specific IP address to direct traffic to.
        :param pulumi.Input[str] floating_ip: The floating IP to associate.
        :param pulumi.Input[str] instance_id: The instance to associte the floating IP with.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new floatingip_associate.
        """
        _FloatingIpAssociateState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            fixed_ip=fixed_ip,
            floating_ip=floating_ip,
            instance_id=instance_id,
            region=region,
            wait_until_associated=wait_until_associated,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             fixed_ip: Optional[pulumi.Input[str]] = None,
             floating_ip: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             wait_until_associated: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if fixed_ip is not None:
            _setter("fixed_ip", fixed_ip)
        if floating_ip is not None:
            _setter("floating_ip", floating_ip)
        if instance_id is not None:
            _setter("instance_id", instance_id)
        if region is not None:
            _setter("region", region)
        if wait_until_associated is not None:
            _setter("wait_until_associated", wait_until_associated)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The specific IP address to direct traffic to.
        """
        return pulumi.get(self, "fixed_ip")

    @fixed_ip.setter
    def fixed_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip", value)

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The floating IP to associate.
        """
        return pulumi.get(self, "floating_ip")

    @floating_ip.setter
    def floating_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "floating_ip", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance to associte the floating IP with.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new floatingip_associate.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="waitUntilAssociated")
    def wait_until_associated(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "wait_until_associated")

    @wait_until_associated.setter
    def wait_until_associated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "wait_until_associated", value)


class FloatingIpAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 floating_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 wait_until_associated: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Associate a floating IP to an instance.

        ## Example Usage
        ### Automatically detect the correct network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        instance1 = openstack.compute.Instance("instance1",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="3",
            key_pair="my_key_pair_name",
            security_groups=["default"])
        fip1_floating_ip = openstack.networking.FloatingIp("fip1FloatingIp", pool="my_pool")
        fip1_floating_ip_associate = openstack.compute.FloatingIpAssociate("fip1FloatingIpAssociate",
            floating_ip=fip1_floating_ip.address,
            instance_id=instance1.id)
        ```
        ### Explicitly set the network to attach to

        ```python
        import pulumi
        import pulumi_openstack as openstack

        instance1 = openstack.compute.Instance("instance1",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="3",
            key_pair="my_key_pair_name",
            security_groups=["default"],
            networks=[
                openstack.compute.InstanceNetworkArgs(
                    name="my_network",
                ),
                openstack.compute.InstanceNetworkArgs(
                    name="default",
                ),
            ])
        fip1_floating_ip = openstack.networking.FloatingIp("fip1FloatingIp", pool="my_pool")
        fip1_floating_ip_associate = openstack.compute.FloatingIpAssociate("fip1FloatingIpAssociate",
            floating_ip=fip1_floating_ip.address,
            instance_id=instance1.id,
            fixed_ip=instance1.networks[1].fixed_ip_v4)
        ```

        ## Import

        This resource can be imported by specifying all three arguments, separated by a forward slash:

        ```sh
         $ pulumi import openstack:compute/floatingIpAssociate:FloatingIpAssociate fip_1 floating_ip/instance_id/fixed_ip
        ```

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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FloatingIpAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Associate a floating IP to an instance.

        ## Example Usage
        ### Automatically detect the correct network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        instance1 = openstack.compute.Instance("instance1",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="3",
            key_pair="my_key_pair_name",
            security_groups=["default"])
        fip1_floating_ip = openstack.networking.FloatingIp("fip1FloatingIp", pool="my_pool")
        fip1_floating_ip_associate = openstack.compute.FloatingIpAssociate("fip1FloatingIpAssociate",
            floating_ip=fip1_floating_ip.address,
            instance_id=instance1.id)
        ```
        ### Explicitly set the network to attach to

        ```python
        import pulumi
        import pulumi_openstack as openstack

        instance1 = openstack.compute.Instance("instance1",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="3",
            key_pair="my_key_pair_name",
            security_groups=["default"],
            networks=[
                openstack.compute.InstanceNetworkArgs(
                    name="my_network",
                ),
                openstack.compute.InstanceNetworkArgs(
                    name="default",
                ),
            ])
        fip1_floating_ip = openstack.networking.FloatingIp("fip1FloatingIp", pool="my_pool")
        fip1_floating_ip_associate = openstack.compute.FloatingIpAssociate("fip1FloatingIpAssociate",
            floating_ip=fip1_floating_ip.address,
            instance_id=instance1.id,
            fixed_ip=instance1.networks[1].fixed_ip_v4)
        ```

        ## Import

        This resource can be imported by specifying all three arguments, separated by a forward slash:

        ```sh
         $ pulumi import openstack:compute/floatingIpAssociate:FloatingIpAssociate fip_1 floating_ip/instance_id/fixed_ip
        ```

        :param str resource_name: The name of the resource.
        :param FloatingIpAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FloatingIpAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            FloatingIpAssociateArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fixed_ip: Optional[pulumi.Input[str]] = None,
                 floating_ip: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 wait_until_associated: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FloatingIpAssociateArgs.__new__(FloatingIpAssociateArgs)

            __props__.__dict__["fixed_ip"] = fixed_ip
            if floating_ip is None and not opts.urn:
                raise TypeError("Missing required property 'floating_ip'")
            __props__.__dict__["floating_ip"] = floating_ip
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region"] = region
            __props__.__dict__["wait_until_associated"] = wait_until_associated
        super(FloatingIpAssociate, __self__).__init__(
            'openstack:compute/floatingIpAssociate:FloatingIpAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fixed_ip: Optional[pulumi.Input[str]] = None,
            floating_ip: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            wait_until_associated: Optional[pulumi.Input[bool]] = None) -> 'FloatingIpAssociate':
        """
        Get an existing FloatingIpAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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

        __props__ = _FloatingIpAssociateState.__new__(_FloatingIpAssociateState)

        __props__.__dict__["fixed_ip"] = fixed_ip
        __props__.__dict__["floating_ip"] = floating_ip
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["region"] = region
        __props__.__dict__["wait_until_associated"] = wait_until_associated
        return FloatingIpAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The specific IP address to direct traffic to.
        """
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter(name="floatingIp")
    def floating_ip(self) -> pulumi.Output[str]:
        """
        The floating IP to associate.
        """
        return pulumi.get(self, "floating_ip")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance to associte the floating IP with.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new floatingip_associate.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="waitUntilAssociated")
    def wait_until_associated(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "wait_until_associated")

