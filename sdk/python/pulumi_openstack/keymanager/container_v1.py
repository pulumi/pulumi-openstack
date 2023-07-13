# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ContainerV1Args', 'ContainerV1']

@pulumi.input_type
class ContainerV1Args:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 acl: Optional[pulumi.Input['ContainerV1AclArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]] = None):
        """
        The set of arguments for constructing a ContainerV1 resource.
        """
        pulumi.set(__self__, "type", type)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_refs is not None:
            pulumi.set(__self__, "secret_refs", secret_refs)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ContainerV1AclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ContainerV1AclArgs']]):
        pulumi.set(self, "acl", value)

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
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]:
        return pulumi.get(self, "secret_refs")

    @secret_refs.setter
    def secret_refs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]):
        pulumi.set(self, "secret_refs", value)


@pulumi.input_type
class _ContainerV1State:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['ContainerV1AclArgs']] = None,
                 consumers: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]] = None,
                 container_ref: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 creator_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ContainerV1 resources.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if consumers is not None:
            pulumi.set(__self__, "consumers", consumers)
        if container_ref is not None:
            pulumi.set(__self__, "container_ref", container_ref)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if creator_id is not None:
            pulumi.set(__self__, "creator_id", creator_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_refs is not None:
            pulumi.set(__self__, "secret_refs", secret_refs)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ContainerV1AclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ContainerV1AclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def consumers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]]:
        return pulumi.get(self, "consumers")

    @consumers.setter
    def consumers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]]):
        pulumi.set(self, "consumers", value)

    @property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container_ref")

    @container_ref.setter
    def container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_ref", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "creator_id")

    @creator_id.setter
    def creator_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creator_id", value)

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
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]:
        return pulumi.get(self, "secret_refs")

    @secret_refs.setter
    def secret_refs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]):
        pulumi.set(self, "secret_refs", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class ContainerV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['ContainerV1AclArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ContainerV1 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerV1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ContainerV1 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ContainerV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['ContainerV1AclArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerV1Args.__new__(ContainerV1Args)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["secret_refs"] = secret_refs
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["consumers"] = None
            __props__.__dict__["container_ref"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["creator_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(ContainerV1, __self__).__init__(
            'openstack:keymanager/containerV1:ContainerV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[pulumi.InputType['ContainerV1AclArgs']]] = None,
            consumers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1ConsumerArgs']]]]] = None,
            container_ref: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            creator_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'ContainerV1':
        """
        Get an existing ContainerV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerV1State.__new__(_ContainerV1State)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["consumers"] = consumers
        __props__.__dict__["container_ref"] = container_ref
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["creator_id"] = creator_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_refs"] = secret_refs
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        return ContainerV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.ContainerV1Acl']:
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def consumers(self) -> pulumi.Output[Sequence['outputs.ContainerV1Consumer']]:
        return pulumi.get(self, "consumers")

    @property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "container_ref")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creator_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> pulumi.Output[Optional[Sequence['outputs.ContainerV1SecretRef']]]:
        return pulumi.get(self, "secret_refs")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

