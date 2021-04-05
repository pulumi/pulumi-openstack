# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['VolumeTypeV3']


class VolumeTypeV3(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 extra_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 is_public: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V3 block storage volume type resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        volume_type1 = openstack.blockstorage.VolumeTypeV3("volumeType1",
            description="Volume type 1",
            extra_specs={
                "capabilities": "gpu",
                "volume_backend_name": "ssd",
            })
        ```

        ## Import

        Volume types can be imported using the `volume_type_id`, e.g.

        ```sh
         $ pulumi import openstack:blockstorage/volumeTypeV3:VolumeTypeV3 volume_type_1 941793f0-0a34-4bc4-b72e-a6326ae58283
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Human-readable description of the port. Changing
               this updates the `description` of an existing volume type.
        :param pulumi.Input[Mapping[str, Any]] extra_specs: Key/Value pairs of metadata for the volume type.
        :param pulumi.Input[bool] is_public: Whether the volume type is public. Changing
               this updates the `is_public` of an existing volume type.
        :param pulumi.Input[str] name: Name of the volume type.  Changing this
               updates the `name` of an existing volume type.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
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

            __props__['description'] = description
            __props__['extra_specs'] = extra_specs
            __props__['is_public'] = is_public
            __props__['name'] = name
            __props__['region'] = region
        super(VolumeTypeV3, __self__).__init__(
            'openstack:blockstorage/volumeTypeV3:VolumeTypeV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            extra_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            is_public: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'VolumeTypeV3':
        """
        Get an existing VolumeTypeV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Human-readable description of the port. Changing
               this updates the `description` of an existing volume type.
        :param pulumi.Input[Mapping[str, Any]] extra_specs: Key/Value pairs of metadata for the volume type.
        :param pulumi.Input[bool] is_public: Whether the volume type is public. Changing
               this updates the `is_public` of an existing volume type.
        :param pulumi.Input[str] name: Name of the volume type.  Changing this
               updates the `name` of an existing volume type.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["extra_specs"] = extra_specs
        __props__["is_public"] = is_public
        __props__["name"] = name
        __props__["region"] = region
        return VolumeTypeV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Human-readable description of the port. Changing
        this updates the `description` of an existing volume type.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="extraSpecs")
    def extra_specs(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Key/Value pairs of metadata for the volume type.
        """
        return pulumi.get(self, "extra_specs")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> pulumi.Output[bool]:
        """
        Whether the volume type is public. Changing
        this updates the `is_public` of an existing volume type.
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the volume type.  Changing this
        updates the `name` of an existing volume type.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the volume. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new quotaset.
        """
        return pulumi.get(self, "region")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

