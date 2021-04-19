# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FlavorArgs', 'Flavor']

@pulumi.input_type
class FlavorArgs:
    def __init__(__self__, *,
                 disk: pulumi.Input[int],
                 ram: pulumi.Input[int],
                 vcpus: pulumi.Input[int],
                 ephemeral: Optional[pulumi.Input[int]] = None,
                 extra_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rx_tx_factor: Optional[pulumi.Input[float]] = None,
                 swap: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Flavor resource.
        :param pulumi.Input[int] disk: The amount of disk space in GiB to use for the root
               (/) partition. Changing this creates a new flavor.
        :param pulumi.Input[int] ram: The amount of RAM to use, in megabytes. Changing this
               creates a new flavor.
        :param pulumi.Input[int] vcpus: The number of virtual CPUs to use. Changing this creates
               a new flavor.
        :param pulumi.Input[int] ephemeral: The amount of ephemeral in GiB. If unspecified,
               the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[Mapping[str, Any]] extra_specs: Key/Value pairs of metadata for the flavor.
        :param pulumi.Input[str] flavor_id: Unique ID (integer or UUID) of flavor to create. Changing
               this creates a new flavor.
        :param pulumi.Input[bool] is_public: Whether the flavor is public. Changing this creates
               a new flavor.
        :param pulumi.Input[str] name: A unique name for the flavor. Changing this creates a new
               flavor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Flavors are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor.
        :param pulumi.Input[float] rx_tx_factor: RX/TX bandwith factor. The default is 1. Changing
               this creates a new flavor.
        :param pulumi.Input[int] swap: The amount of disk space in megabytes to use. If
               unspecified, the default is 0. Changing this creates a new flavor.
        """
        pulumi.set(__self__, "disk", disk)
        pulumi.set(__self__, "ram", ram)
        pulumi.set(__self__, "vcpus", vcpus)
        if ephemeral is not None:
            pulumi.set(__self__, "ephemeral", ephemeral)
        if extra_specs is not None:
            pulumi.set(__self__, "extra_specs", extra_specs)
        if flavor_id is not None:
            pulumi.set(__self__, "flavor_id", flavor_id)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rx_tx_factor is not None:
            pulumi.set(__self__, "rx_tx_factor", rx_tx_factor)
        if swap is not None:
            pulumi.set(__self__, "swap", swap)

    @property
    @pulumi.getter
    def disk(self) -> pulumi.Input[int]:
        """
        The amount of disk space in GiB to use for the root
        (/) partition. Changing this creates a new flavor.
        """
        return pulumi.get(self, "disk")

    @disk.setter
    def disk(self, value: pulumi.Input[int]):
        pulumi.set(self, "disk", value)

    @property
    @pulumi.getter
    def ram(self) -> pulumi.Input[int]:
        """
        The amount of RAM to use, in megabytes. Changing this
        creates a new flavor.
        """
        return pulumi.get(self, "ram")

    @ram.setter
    def ram(self, value: pulumi.Input[int]):
        pulumi.set(self, "ram", value)

    @property
    @pulumi.getter
    def vcpus(self) -> pulumi.Input[int]:
        """
        The number of virtual CPUs to use. Changing this creates
        a new flavor.
        """
        return pulumi.get(self, "vcpus")

    @vcpus.setter
    def vcpus(self, value: pulumi.Input[int]):
        pulumi.set(self, "vcpus", value)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of ephemeral in GiB. If unspecified,
        the default is 0. Changing this creates a new flavor.
        """
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/Value pairs of metadata for the flavor.
        """
        return pulumi.get(self, "extra_specs")

    @extra_specs.setter
    def extra_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_specs", value)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique ID (integer or UUID) of flavor to create. Changing
        this creates a new flavor.
        """
        return pulumi.get(self, "flavor_id")

    @flavor_id.setter
    def flavor_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor_id", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the flavor is public. Changing this creates
        a new flavor.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the flavor. Changing this creates a new
        flavor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Flavors are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new flavor.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="rxTxFactor")
    def rx_tx_factor(self) -> Optional[pulumi.Input[float]]:
        """
        RX/TX bandwith factor. The default is 1. Changing
        this creates a new flavor.
        """
        return pulumi.get(self, "rx_tx_factor")

    @rx_tx_factor.setter
    def rx_tx_factor(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "rx_tx_factor", value)

    @property
    @pulumi.getter
    def swap(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of disk space in megabytes to use. If
        unspecified, the default is 0. Changing this creates a new flavor.
        """
        return pulumi.get(self, "swap")

    @swap.setter
    def swap(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "swap", value)


@pulumi.input_type
class _FlavorState:
    def __init__(__self__, *,
                 disk: Optional[pulumi.Input[int]] = None,
                 ephemeral: Optional[pulumi.Input[int]] = None,
                 extra_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ram: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rx_tx_factor: Optional[pulumi.Input[float]] = None,
                 swap: Optional[pulumi.Input[int]] = None,
                 vcpus: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Flavor resources.
        :param pulumi.Input[int] disk: The amount of disk space in GiB to use for the root
               (/) partition. Changing this creates a new flavor.
        :param pulumi.Input[int] ephemeral: The amount of ephemeral in GiB. If unspecified,
               the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[Mapping[str, Any]] extra_specs: Key/Value pairs of metadata for the flavor.
        :param pulumi.Input[str] flavor_id: Unique ID (integer or UUID) of flavor to create. Changing
               this creates a new flavor.
        :param pulumi.Input[bool] is_public: Whether the flavor is public. Changing this creates
               a new flavor.
        :param pulumi.Input[str] name: A unique name for the flavor. Changing this creates a new
               flavor.
        :param pulumi.Input[int] ram: The amount of RAM to use, in megabytes. Changing this
               creates a new flavor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Flavors are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor.
        :param pulumi.Input[float] rx_tx_factor: RX/TX bandwith factor. The default is 1. Changing
               this creates a new flavor.
        :param pulumi.Input[int] swap: The amount of disk space in megabytes to use. If
               unspecified, the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[int] vcpus: The number of virtual CPUs to use. Changing this creates
               a new flavor.
        """
        if disk is not None:
            pulumi.set(__self__, "disk", disk)
        if ephemeral is not None:
            pulumi.set(__self__, "ephemeral", ephemeral)
        if extra_specs is not None:
            pulumi.set(__self__, "extra_specs", extra_specs)
        if flavor_id is not None:
            pulumi.set(__self__, "flavor_id", flavor_id)
        if is_public is not None:
            pulumi.set(__self__, "is_public", is_public)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ram is not None:
            pulumi.set(__self__, "ram", ram)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rx_tx_factor is not None:
            pulumi.set(__self__, "rx_tx_factor", rx_tx_factor)
        if swap is not None:
            pulumi.set(__self__, "swap", swap)
        if vcpus is not None:
            pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter
    def disk(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of disk space in GiB to use for the root
        (/) partition. Changing this creates a new flavor.
        """
        return pulumi.get(self, "disk")

    @disk.setter
    def disk(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk", value)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of ephemeral in GiB. If unspecified,
        the default is 0. Changing this creates a new flavor.
        """
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/Value pairs of metadata for the flavor.
        """
        return pulumi.get(self, "extra_specs")

    @extra_specs.setter
    def extra_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_specs", value)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique ID (integer or UUID) of flavor to create. Changing
        this creates a new flavor.
        """
        return pulumi.get(self, "flavor_id")

    @flavor_id.setter
    def flavor_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor_id", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the flavor is public. Changing this creates
        a new flavor.
        """
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the flavor. Changing this creates a new
        flavor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def ram(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of RAM to use, in megabytes. Changing this
        creates a new flavor.
        """
        return pulumi.get(self, "ram")

    @ram.setter
    def ram(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ram", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Flavors are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new flavor.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="rxTxFactor")
    def rx_tx_factor(self) -> Optional[pulumi.Input[float]]:
        """
        RX/TX bandwith factor. The default is 1. Changing
        this creates a new flavor.
        """
        return pulumi.get(self, "rx_tx_factor")

    @rx_tx_factor.setter
    def rx_tx_factor(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "rx_tx_factor", value)

    @property
    @pulumi.getter
    def swap(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of disk space in megabytes to use. If
        unspecified, the default is 0. Changing this creates a new flavor.
        """
        return pulumi.get(self, "swap")

    @swap.setter
    def swap(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "swap", value)

    @property
    @pulumi.getter
    def vcpus(self) -> Optional[pulumi.Input[int]]:
        """
        The number of virtual CPUs to use. Changing this creates
        a new flavor.
        """
        return pulumi.get(self, "vcpus")

    @vcpus.setter
    def vcpus(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vcpus", value)


class Flavor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk: Optional[pulumi.Input[int]] = None,
                 ephemeral: Optional[pulumi.Input[int]] = None,
                 extra_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ram: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rx_tx_factor: Optional[pulumi.Input[float]] = None,
                 swap: Optional[pulumi.Input[int]] = None,
                 vcpus: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a V2 flavor resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_flavor = openstack.compute.Flavor("test-flavor",
            disk=20,
            extra_specs={
                "hw:cpu_policy": "CPU-POLICY",
                "hw:cpu_thread_policy": "CPU-THREAD-POLICY",
            },
            ram=8096,
            vcpus=2)
        ```

        ## Import

        Flavors can be imported using the `ID`, e.g.

        ```sh
         $ pulumi import openstack:compute/flavor:Flavor my-flavor 4142e64b-1b35-44a0-9b1e-5affc7af1106
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] disk: The amount of disk space in GiB to use for the root
               (/) partition. Changing this creates a new flavor.
        :param pulumi.Input[int] ephemeral: The amount of ephemeral in GiB. If unspecified,
               the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[Mapping[str, Any]] extra_specs: Key/Value pairs of metadata for the flavor.
        :param pulumi.Input[str] flavor_id: Unique ID (integer or UUID) of flavor to create. Changing
               this creates a new flavor.
        :param pulumi.Input[bool] is_public: Whether the flavor is public. Changing this creates
               a new flavor.
        :param pulumi.Input[str] name: A unique name for the flavor. Changing this creates a new
               flavor.
        :param pulumi.Input[int] ram: The amount of RAM to use, in megabytes. Changing this
               creates a new flavor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Flavors are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor.
        :param pulumi.Input[float] rx_tx_factor: RX/TX bandwith factor. The default is 1. Changing
               this creates a new flavor.
        :param pulumi.Input[int] swap: The amount of disk space in megabytes to use. If
               unspecified, the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[int] vcpus: The number of virtual CPUs to use. Changing this creates
               a new flavor.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlavorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 flavor resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_flavor = openstack.compute.Flavor("test-flavor",
            disk=20,
            extra_specs={
                "hw:cpu_policy": "CPU-POLICY",
                "hw:cpu_thread_policy": "CPU-THREAD-POLICY",
            },
            ram=8096,
            vcpus=2)
        ```

        ## Import

        Flavors can be imported using the `ID`, e.g.

        ```sh
         $ pulumi import openstack:compute/flavor:Flavor my-flavor 4142e64b-1b35-44a0-9b1e-5affc7af1106
        ```

        :param str resource_name: The name of the resource.
        :param FlavorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlavorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk: Optional[pulumi.Input[int]] = None,
                 ephemeral: Optional[pulumi.Input[int]] = None,
                 extra_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ram: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rx_tx_factor: Optional[pulumi.Input[float]] = None,
                 swap: Optional[pulumi.Input[int]] = None,
                 vcpus: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlavorArgs.__new__(FlavorArgs)

            if disk is None and not opts.urn:
                raise TypeError("Missing required property 'disk'")
            __props__.__dict__["disk"] = disk
            __props__.__dict__["ephemeral"] = ephemeral
            __props__.__dict__["extra_specs"] = extra_specs
            __props__.__dict__["flavor_id"] = flavor_id
            __props__.__dict__["is_public"] = is_public
            __props__.__dict__["name"] = name
            if ram is None and not opts.urn:
                raise TypeError("Missing required property 'ram'")
            __props__.__dict__["ram"] = ram
            __props__.__dict__["region"] = region
            __props__.__dict__["rx_tx_factor"] = rx_tx_factor
            __props__.__dict__["swap"] = swap
            if vcpus is None and not opts.urn:
                raise TypeError("Missing required property 'vcpus'")
            __props__.__dict__["vcpus"] = vcpus
        super(Flavor, __self__).__init__(
            'openstack:compute/flavor:Flavor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disk: Optional[pulumi.Input[int]] = None,
            ephemeral: Optional[pulumi.Input[int]] = None,
            extra_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            flavor_id: Optional[pulumi.Input[str]] = None,
            is_public: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ram: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rx_tx_factor: Optional[pulumi.Input[float]] = None,
            swap: Optional[pulumi.Input[int]] = None,
            vcpus: Optional[pulumi.Input[int]] = None) -> 'Flavor':
        """
        Get an existing Flavor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] disk: The amount of disk space in GiB to use for the root
               (/) partition. Changing this creates a new flavor.
        :param pulumi.Input[int] ephemeral: The amount of ephemeral in GiB. If unspecified,
               the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[Mapping[str, Any]] extra_specs: Key/Value pairs of metadata for the flavor.
        :param pulumi.Input[str] flavor_id: Unique ID (integer or UUID) of flavor to create. Changing
               this creates a new flavor.
        :param pulumi.Input[bool] is_public: Whether the flavor is public. Changing this creates
               a new flavor.
        :param pulumi.Input[str] name: A unique name for the flavor. Changing this creates a new
               flavor.
        :param pulumi.Input[int] ram: The amount of RAM to use, in megabytes. Changing this
               creates a new flavor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Flavors are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor.
        :param pulumi.Input[float] rx_tx_factor: RX/TX bandwith factor. The default is 1. Changing
               this creates a new flavor.
        :param pulumi.Input[int] swap: The amount of disk space in megabytes to use. If
               unspecified, the default is 0. Changing this creates a new flavor.
        :param pulumi.Input[int] vcpus: The number of virtual CPUs to use. Changing this creates
               a new flavor.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlavorState.__new__(_FlavorState)

        __props__.__dict__["disk"] = disk
        __props__.__dict__["ephemeral"] = ephemeral
        __props__.__dict__["extra_specs"] = extra_specs
        __props__.__dict__["flavor_id"] = flavor_id
        __props__.__dict__["is_public"] = is_public
        __props__.__dict__["name"] = name
        __props__.__dict__["ram"] = ram
        __props__.__dict__["region"] = region
        __props__.__dict__["rx_tx_factor"] = rx_tx_factor
        __props__.__dict__["swap"] = swap
        __props__.__dict__["vcpus"] = vcpus
        return Flavor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def disk(self) -> pulumi.Output[int]:
        """
        The amount of disk space in GiB to use for the root
        (/) partition. Changing this creates a new flavor.
        """
        return pulumi.get(self, "disk")

    @property
    @pulumi.getter
    def ephemeral(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of ephemeral in GiB. If unspecified,
        the default is 0. Changing this creates a new flavor.
        """
        return pulumi.get(self, "ephemeral")

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Key/Value pairs of metadata for the flavor.
        """
        return pulumi.get(self, "extra_specs")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> pulumi.Output[str]:
        """
        Unique ID (integer or UUID) of flavor to create. Changing
        this creates a new flavor.
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the flavor is public. Changing this creates
        a new flavor.
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the flavor. Changing this creates a new
        flavor.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def ram(self) -> pulumi.Output[int]:
        """
        The amount of RAM to use, in megabytes. Changing this
        creates a new flavor.
        """
        return pulumi.get(self, "ram")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        Flavors are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new flavor.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="rxTxFactor")
    def rx_tx_factor(self) -> pulumi.Output[Optional[float]]:
        """
        RX/TX bandwith factor. The default is 1. Changing
        this creates a new flavor.
        """
        return pulumi.get(self, "rx_tx_factor")

    @property
    @pulumi.getter
    def swap(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of disk space in megabytes to use. If
        unspecified, the default is 0. Changing this creates a new flavor.
        """
        return pulumi.get(self, "swap")

    @property
    @pulumi.getter
    def vcpus(self) -> pulumi.Output[int]:
        """
        The number of virtual CPUs to use. Changing this creates
        a new flavor.
        """
        return pulumi.get(self, "vcpus")

