# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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
                 description: Optional[pulumi.Input[str]] = None,
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
        """
        pulumi.set(__self__, "disk", disk)
        pulumi.set(__self__, "ram", ram)
        pulumi.set(__self__, "vcpus", vcpus)
        if description is not None:
            pulumi.set(__self__, "description", description)
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
        return pulumi.get(self, "disk")

    @disk.setter
    def disk(self, value: pulumi.Input[int]):
        pulumi.set(self, "disk", value)

    @property
    @pulumi.getter
    def ram(self) -> pulumi.Input[int]:
        return pulumi.get(self, "ram")

    @ram.setter
    def ram(self, value: pulumi.Input[int]):
        pulumi.set(self, "ram", value)

    @property
    @pulumi.getter
    def vcpus(self) -> pulumi.Input[int]:
        return pulumi.get(self, "vcpus")

    @vcpus.setter
    def vcpus(self, value: pulumi.Input[int]):
        pulumi.set(self, "vcpus", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "extra_specs")

    @extra_specs.setter
    def extra_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_specs", value)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "flavor_id")

    @flavor_id.setter
    def flavor_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor_id", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="rxTxFactor")
    def rx_tx_factor(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "rx_tx_factor")

    @rx_tx_factor.setter
    def rx_tx_factor(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "rx_tx_factor", value)

    @property
    @pulumi.getter
    def swap(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "swap")

    @swap.setter
    def swap(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "swap", value)


@pulumi.input_type
class _FlavorState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
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
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
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
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disk(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "disk")

    @disk.setter
    def disk(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk", value)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "extra_specs")

    @extra_specs.setter
    def extra_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_specs", value)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "flavor_id")

    @flavor_id.setter
    def flavor_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor_id", value)

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_public")

    @is_public.setter
    def is_public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_public", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def ram(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ram")

    @ram.setter
    def ram(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ram", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="rxTxFactor")
    def rx_tx_factor(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "rx_tx_factor")

    @rx_tx_factor.setter
    def rx_tx_factor(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "rx_tx_factor", value)

    @property
    @pulumi.getter
    def swap(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "swap")

    @swap.setter
    def swap(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "swap", value)

    @property
    @pulumi.getter
    def vcpus(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "vcpus")

    @vcpus.setter
    def vcpus(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vcpus", value)


class Flavor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
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
        Create a Flavor resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlavorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Flavor resource with the given unique name, props, and options.
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
                 description: Optional[pulumi.Input[str]] = None,
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
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlavorArgs.__new__(FlavorArgs)

            __props__.__dict__["description"] = description
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
            description: Optional[pulumi.Input[str]] = None,
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
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlavorState.__new__(_FlavorState)

        __props__.__dict__["description"] = description
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
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disk(self) -> pulumi.Output[int]:
        return pulumi.get(self, "disk")

    @property
    @pulumi.getter
    def ephemeral(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "ephemeral")

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> pulumi.Output[Mapping[str, Any]]:
        return pulumi.get(self, "extra_specs")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def ram(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ram")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="rxTxFactor")
    def rx_tx_factor(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "rx_tx_factor")

    @property
    @pulumi.getter
    def swap(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "swap")

    @property
    @pulumi.getter
    def vcpus(self) -> pulumi.Output[int]:
        return pulumi.get(self, "vcpus")

