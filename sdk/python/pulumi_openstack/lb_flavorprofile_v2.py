# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['LbFlavorprofileV2Args', 'LbFlavorprofileV2']

@pulumi.input_type
class LbFlavorprofileV2Args:
    def __init__(__self__, *,
                 flavor_data: pulumi.Input[builtins.str],
                 provider_name: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LbFlavorprofileV2 resource.
        :param pulumi.Input[builtins.str] flavor_data: String that passes the flavor_data for the flavorprofile.
               The data that are allowed depend on the `provider_name` that is passed. jsonencode
               can be used for readability as shown in the example above.
               Changing this updates the existing flavorprofile.
        :param pulumi.Input[builtins.str] provider_name: The provider_name that the flavor_profile will use.
               Changing this updates the existing flavorprofile.
        :param pulumi.Input[builtins.str] name: Name of the flavorprofile. Changing this updates the existing
               flavorprofile.
        """
        pulumi.set(__self__, "flavor_data", flavor_data)
        pulumi.set(__self__, "provider_name", provider_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="flavorData")
    def flavor_data(self) -> pulumi.Input[builtins.str]:
        """
        String that passes the flavor_data for the flavorprofile.
        The data that are allowed depend on the `provider_name` that is passed. jsonencode
        can be used for readability as shown in the example above.
        Changing this updates the existing flavorprofile.
        """
        return pulumi.get(self, "flavor_data")

    @flavor_data.setter
    def flavor_data(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "flavor_data", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Input[builtins.str]:
        """
        The provider_name that the flavor_profile will use.
        Changing this updates the existing flavorprofile.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the flavorprofile. Changing this updates the existing
        flavorprofile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _LbFlavorprofileV2State:
    def __init__(__self__, *,
                 flavor_data: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 provider_name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LbFlavorprofileV2 resources.
        :param pulumi.Input[builtins.str] flavor_data: String that passes the flavor_data for the flavorprofile.
               The data that are allowed depend on the `provider_name` that is passed. jsonencode
               can be used for readability as shown in the example above.
               Changing this updates the existing flavorprofile.
        :param pulumi.Input[builtins.str] name: Name of the flavorprofile. Changing this updates the existing
               flavorprofile.
        :param pulumi.Input[builtins.str] provider_name: The provider_name that the flavor_profile will use.
               Changing this updates the existing flavorprofile.
        """
        if flavor_data is not None:
            pulumi.set(__self__, "flavor_data", flavor_data)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if provider_name is not None:
            pulumi.set(__self__, "provider_name", provider_name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="flavorData")
    def flavor_data(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        String that passes the flavor_data for the flavorprofile.
        The data that are allowed depend on the `provider_name` that is passed. jsonencode
        can be used for readability as shown in the example above.
        Changing this updates the existing flavorprofile.
        """
        return pulumi.get(self, "flavor_data")

    @flavor_data.setter
    def flavor_data(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "flavor_data", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the flavorprofile. Changing this updates the existing
        flavorprofile.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The provider_name that the flavor_profile will use.
        Changing this updates the existing flavorprofile.
        """
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


class LbFlavorprofileV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flavor_data: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 provider_name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages a V2 load balancer flavorprofile resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        ## Example Usage

        ### Using jsonencode

        ```python
        import pulumi
        import json
        import pulumi_openstack as openstack

        flavorprofile1 = openstack.LbFlavorprofileV2("flavorprofile_1",
            name="amphora-single-profile",
            provider_name="amphora",
            flavor_data=json.dumps({
                "loadbalancer_topology": "SINGLE",
            }))
        ```

        ### Using plain string

        ```python
        import pulumi
        import pulumi_openstack as openstack

        flavorprofile1 = openstack.LbFlavorprofileV2("flavorprofile_1",
            name="amphora-single-profile",
            provider_name="amphora",
            flavor_data="{\\"loadbalancer_topology\\": \\"SINGLE\\"}")
        ```

        ## Import

        flavorprofiles can be imported using their `id`. Example:

        ```sh
        $ pulumi import openstack:index/lbFlavorprofileV2:LbFlavorprofileV2 flavorprofile_1 2a0f2240-c5e6-41de-896d-e80d97428d6b
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] flavor_data: String that passes the flavor_data for the flavorprofile.
               The data that are allowed depend on the `provider_name` that is passed. jsonencode
               can be used for readability as shown in the example above.
               Changing this updates the existing flavorprofile.
        :param pulumi.Input[builtins.str] name: Name of the flavorprofile. Changing this updates the existing
               flavorprofile.
        :param pulumi.Input[builtins.str] provider_name: The provider_name that the flavor_profile will use.
               Changing this updates the existing flavorprofile.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LbFlavorprofileV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 load balancer flavorprofile resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        ## Example Usage

        ### Using jsonencode

        ```python
        import pulumi
        import json
        import pulumi_openstack as openstack

        flavorprofile1 = openstack.LbFlavorprofileV2("flavorprofile_1",
            name="amphora-single-profile",
            provider_name="amphora",
            flavor_data=json.dumps({
                "loadbalancer_topology": "SINGLE",
            }))
        ```

        ### Using plain string

        ```python
        import pulumi
        import pulumi_openstack as openstack

        flavorprofile1 = openstack.LbFlavorprofileV2("flavorprofile_1",
            name="amphora-single-profile",
            provider_name="amphora",
            flavor_data="{\\"loadbalancer_topology\\": \\"SINGLE\\"}")
        ```

        ## Import

        flavorprofiles can be imported using their `id`. Example:

        ```sh
        $ pulumi import openstack:index/lbFlavorprofileV2:LbFlavorprofileV2 flavorprofile_1 2a0f2240-c5e6-41de-896d-e80d97428d6b
        ```

        :param str resource_name: The name of the resource.
        :param LbFlavorprofileV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbFlavorprofileV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flavor_data: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 provider_name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbFlavorprofileV2Args.__new__(LbFlavorprofileV2Args)

            if flavor_data is None and not opts.urn:
                raise TypeError("Missing required property 'flavor_data'")
            __props__.__dict__["flavor_data"] = flavor_data
            __props__.__dict__["name"] = name
            if provider_name is None and not opts.urn:
                raise TypeError("Missing required property 'provider_name'")
            __props__.__dict__["provider_name"] = provider_name
            __props__.__dict__["region"] = region
        super(LbFlavorprofileV2, __self__).__init__(
            'openstack:index/lbFlavorprofileV2:LbFlavorprofileV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            flavor_data: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            provider_name: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None) -> 'LbFlavorprofileV2':
        """
        Get an existing LbFlavorprofileV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] flavor_data: String that passes the flavor_data for the flavorprofile.
               The data that are allowed depend on the `provider_name` that is passed. jsonencode
               can be used for readability as shown in the example above.
               Changing this updates the existing flavorprofile.
        :param pulumi.Input[builtins.str] name: Name of the flavorprofile. Changing this updates the existing
               flavorprofile.
        :param pulumi.Input[builtins.str] provider_name: The provider_name that the flavor_profile will use.
               Changing this updates the existing flavorprofile.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbFlavorprofileV2State.__new__(_LbFlavorprofileV2State)

        __props__.__dict__["flavor_data"] = flavor_data
        __props__.__dict__["name"] = name
        __props__.__dict__["provider_name"] = provider_name
        __props__.__dict__["region"] = region
        return LbFlavorprofileV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="flavorData")
    def flavor_data(self) -> pulumi.Output[builtins.str]:
        """
        String that passes the flavor_data for the flavorprofile.
        The data that are allowed depend on the `provider_name` that is passed. jsonencode
        can be used for readability as shown in the example above.
        Changing this updates the existing flavorprofile.
        """
        return pulumi.get(self, "flavor_data")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the flavorprofile. Changing this updates the existing
        flavorprofile.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[builtins.str]:
        """
        The provider_name that the flavor_profile will use.
        Changing this updates the existing flavorprofile.
        """
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "region")

