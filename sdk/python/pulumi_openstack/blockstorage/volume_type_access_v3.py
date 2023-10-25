# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VolumeTypeAccessV3Args', 'VolumeTypeAccessV3']

@pulumi.input_type
class VolumeTypeAccessV3Args:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 volume_type_id: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VolumeTypeAccessV3 resource.
        :param pulumi.Input[str] project_id: ID of the project to give access to. Changing this
               creates a new resource.
        :param pulumi.Input[str] volume_type_id: ID of the volume type to give access to. Changing
               this creates a new resource.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
        """
        VolumeTypeAccessV3Args._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project_id=project_id,
            volume_type_id=volume_type_id,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project_id: Optional[pulumi.Input[str]] = None,
             volume_type_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if project_id is None:
            raise TypeError("Missing 'project_id' argument")
        if volume_type_id is None and 'volumeTypeId' in kwargs:
            volume_type_id = kwargs['volumeTypeId']
        if volume_type_id is None:
            raise TypeError("Missing 'volume_type_id' argument")

        _setter("project_id", project_id)
        _setter("volume_type_id", volume_type_id)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        ID of the project to give access to. Changing this
        creates a new resource.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="volumeTypeId")
    def volume_type_id(self) -> pulumi.Input[str]:
        """
        ID of the volume type to give access to. Changing
        this creates a new resource.
        """
        return pulumi.get(self, "volume_type_id")

    @volume_type_id.setter
    def volume_type_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "volume_type_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the volume. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new quotaset.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _VolumeTypeAccessV3State:
    def __init__(__self__, *,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 volume_type_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VolumeTypeAccessV3 resources.
        :param pulumi.Input[str] project_id: ID of the project to give access to. Changing this
               creates a new resource.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
        :param pulumi.Input[str] volume_type_id: ID of the volume type to give access to. Changing
               this creates a new resource.
        """
        _VolumeTypeAccessV3State._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            project_id=project_id,
            region=region,
            volume_type_id=volume_type_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             project_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             volume_type_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if volume_type_id is None and 'volumeTypeId' in kwargs:
            volume_type_id = kwargs['volumeTypeId']

        if project_id is not None:
            _setter("project_id", project_id)
        if region is not None:
            _setter("region", region)
        if volume_type_id is not None:
            _setter("volume_type_id", volume_type_id)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the project to give access to. Changing this
        creates a new resource.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the volume. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new quotaset.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="volumeTypeId")
    def volume_type_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the volume type to give access to. Changing
        this creates a new resource.
        """
        return pulumi.get(self, "volume_type_id")

    @volume_type_id.setter
    def volume_type_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type_id", value)


class VolumeTypeAccessV3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 volume_type_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V3 block storage volume type access resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        ## Import

        Volume types access can be imported using the `volume_type_id/project_id`, e.g.

        ```sh
         $ pulumi import openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3 volume_type_access 941793f0-0a34-4bc4-b72e-a6326ae58283/ed498e81f0cc448bae0ad4f8f21bf67f
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: ID of the project to give access to. Changing this
               creates a new resource.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
        :param pulumi.Input[str] volume_type_id: ID of the volume type to give access to. Changing
               this creates a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VolumeTypeAccessV3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V3 block storage volume type access resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        ## Import

        Volume types access can be imported using the `volume_type_id/project_id`, e.g.

        ```sh
         $ pulumi import openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3 volume_type_access 941793f0-0a34-4bc4-b72e-a6326ae58283/ed498e81f0cc448bae0ad4f8f21bf67f
        ```

        :param str resource_name: The name of the resource.
        :param VolumeTypeAccessV3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VolumeTypeAccessV3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            VolumeTypeAccessV3Args._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 volume_type_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VolumeTypeAccessV3Args.__new__(VolumeTypeAccessV3Args)

            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            if volume_type_id is None and not opts.urn:
                raise TypeError("Missing required property 'volume_type_id'")
            __props__.__dict__["volume_type_id"] = volume_type_id
        super(VolumeTypeAccessV3, __self__).__init__(
            'openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            volume_type_id: Optional[pulumi.Input[str]] = None) -> 'VolumeTypeAccessV3':
        """
        Get an existing VolumeTypeAccessV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: ID of the project to give access to. Changing this
               creates a new resource.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
        :param pulumi.Input[str] volume_type_id: ID of the volume type to give access to. Changing
               this creates a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VolumeTypeAccessV3State.__new__(_VolumeTypeAccessV3State)

        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["volume_type_id"] = volume_type_id
        return VolumeTypeAccessV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        ID of the project to give access to. Changing this
        creates a new resource.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the volume. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new quotaset.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="volumeTypeId")
    def volume_type_id(self) -> pulumi.Output[str]:
        """
        ID of the volume type to give access to. Changing
        this creates a new resource.
        """
        return pulumi.get(self, "volume_type_id")

