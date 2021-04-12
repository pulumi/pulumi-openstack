# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['FlavorAccessArgs', 'FlavorAccess']

@pulumi.input_type
class FlavorAccessArgs:
    def __init__(__self__, *,
                 flavor_id: pulumi.Input[str],
                 tenant_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FlavorAccess resource.
        :param pulumi.Input[str] flavor_id: The UUID of flavor to use. Changing this creates a new flavor access.
        :param pulumi.Input[str] tenant_id: The UUID of tenant which is allowed to use the flavor.
               Changing this creates a new flavor access.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor access.
        """
        pulumi.set(__self__, "flavor_id", flavor_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> pulumi.Input[str]:
        """
        The UUID of flavor to use. Changing this creates a new flavor access.
        """
        return pulumi.get(self, "flavor_id")

    @flavor_id.setter
    def flavor_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "flavor_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Input[str]:
        """
        The UUID of tenant which is allowed to use the flavor.
        Changing this creates a new flavor access.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new flavor access.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class FlavorAccess(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a project access for flavor V2 resource within OpenStack.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ***

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        flavor1 = openstack.compute.Flavor("flavor1",
            disk=20,
            is_public=False,
            ram=8096,
            vcpus=2)
        access1 = openstack.compute.FlavorAccess("access1",
            flavor_id=flavor1.id,
            tenant_id=project1.id)
        ```

        ## Import

        This resource can be imported by specifying all two arguments, separated by a forward slash

        ```sh
         $ pulumi import openstack:compute/flavorAccess:FlavorAccess access_1 <flavor_id>/<tenant_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] flavor_id: The UUID of flavor to use. Changing this creates a new flavor access.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor access.
        :param pulumi.Input[str] tenant_id: The UUID of tenant which is allowed to use the flavor.
               Changing this creates a new flavor access.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlavorAccessArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a project access for flavor V2 resource within OpenStack.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ***

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        flavor1 = openstack.compute.Flavor("flavor1",
            disk=20,
            is_public=False,
            ram=8096,
            vcpus=2)
        access1 = openstack.compute.FlavorAccess("access1",
            flavor_id=flavor1.id,
            tenant_id=project1.id)
        ```

        ## Import

        This resource can be imported by specifying all two arguments, separated by a forward slash

        ```sh
         $ pulumi import openstack:compute/flavorAccess:FlavorAccess access_1 <flavor_id>/<tenant_id>
        ```

        :param str resource_name: The name of the resource.
        :param FlavorAccessArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlavorAccessArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flavor_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = dict()

            if flavor_id is None and not opts.urn:
                raise TypeError("Missing required property 'flavor_id'")
            __props__['flavor_id'] = flavor_id
            __props__['region'] = region
            if tenant_id is None and not opts.urn:
                raise TypeError("Missing required property 'tenant_id'")
            __props__['tenant_id'] = tenant_id
        super(FlavorAccess, __self__).__init__(
            'openstack:compute/flavorAccess:FlavorAccess',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            flavor_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'FlavorAccess':
        """
        Get an existing FlavorAccess resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] flavor_id: The UUID of flavor to use. Changing this creates a new flavor access.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new flavor access.
        :param pulumi.Input[str] tenant_id: The UUID of tenant which is allowed to use the flavor.
               Changing this creates a new flavor access.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["flavor_id"] = flavor_id
        __props__["region"] = region
        __props__["tenant_id"] = tenant_id
        return FlavorAccess(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> pulumi.Output[str]:
        """
        The UUID of flavor to use. Changing this creates a new flavor access.
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new flavor access.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The UUID of tenant which is allowed to use the flavor.
        Changing this creates a new flavor access.
        """
        return pulumi.get(self, "tenant_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

