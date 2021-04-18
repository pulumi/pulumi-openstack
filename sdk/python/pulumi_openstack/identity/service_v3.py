# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceV3Args', 'ServiceV3']

@pulumi.input_type
class ServiceV3Args:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceV3 resource.
        :param pulumi.Input[str] type: The service type.
        :param pulumi.Input[str] description: The service description.
        :param pulumi.Input[bool] enabled: The service status. Defaults to `true`.
        :param pulumi.Input[str] name: The service name.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used.
        """
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The service type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The service description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        The service status. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ServiceV3State:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceV3 resources.
        :param pulumi.Input[str] description: The service description.
        :param pulumi.Input[bool] enabled: The service status. Defaults to `true`.
        :param pulumi.Input[str] name: The service name.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used.
        :param pulumi.Input[str] type: The service type.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The service description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        The service status. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The service name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The service type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class ServiceV3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V3 Service resource within OpenStack Keystone.

        > **Note:** This usually requires admin privileges.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        service1 = openstack.identity.ServiceV3("service1", type="custom")
        ```

        ## Import

        Services can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:identity/serviceV3:ServiceV3 service_1 6688e967-158a-496f-a224-cae3414e6b61
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The service description.
        :param pulumi.Input[bool] enabled: The service status. Defaults to `true`.
        :param pulumi.Input[str] name: The service name.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used.
        :param pulumi.Input[str] type: The service type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceV3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V3 Service resource within OpenStack Keystone.

        > **Note:** This usually requires admin privileges.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        service1 = openstack.identity.ServiceV3("service1", type="custom")
        ```

        ## Import

        Services can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:identity/serviceV3:ServiceV3 service_1 6688e967-158a-496f-a224-cae3414e6b61
        ```

        :param str resource_name: The name of the resource.
        :param ServiceV3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceV3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = ServiceV3Args.__new__(ServiceV3Args)

            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(ServiceV3, __self__).__init__(
            'openstack:identity/serviceV3:ServiceV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ServiceV3':
        """
        Get an existing ServiceV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The service description.
        :param pulumi.Input[bool] enabled: The service status. Defaults to `true`.
        :param pulumi.Input[str] name: The service name.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used.
        :param pulumi.Input[str] type: The service type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceV3State.__new__(_ServiceV3State)

        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["type"] = type
        return ServiceV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The service description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        The service status. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The service name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The service type.
        """
        return pulumi.get(self, "type")

