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

__all__ = ['SecretV1Args', 'SecretV1']

@pulumi.input_type
class SecretV1Args:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['SecretV1AclArgs']] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 bit_length: Optional[pulumi.Input[int]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload: Optional[pulumi.Input[str]] = None,
                 payload_content_encoding: Optional[pulumi.Input[str]] = None,
                 payload_content_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretV1 resource.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if bit_length is not None:
            pulumi.set(__self__, "bit_length", bit_length)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if payload is not None:
            pulumi.set(__self__, "payload", payload)
        if payload_content_encoding is not None:
            pulumi.set(__self__, "payload_content_encoding", payload_content_encoding)
        if payload_content_type is not None:
            pulumi.set(__self__, "payload_content_type", payload_content_type)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['SecretV1AclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['SecretV1AclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="bitLength")
    def bit_length(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "bit_length")

    @bit_length.setter
    def bit_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bit_length", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def payload(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload")

    @payload.setter
    def payload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload", value)

    @property
    @pulumi.getter(name="payloadContentEncoding")
    def payload_content_encoding(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload_content_encoding")

    @payload_content_encoding.setter
    def payload_content_encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_content_encoding", value)

    @property
    @pulumi.getter(name="payloadContentType")
    def payload_content_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload_content_type")

    @payload_content_type.setter
    def payload_content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_content_type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_type", value)


@pulumi.input_type
class _SecretV1State:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['SecretV1AclArgs']] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 all_metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 bit_length: Optional[pulumi.Input[int]] = None,
                 content_types: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 creator_id: Optional[pulumi.Input[str]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload: Optional[pulumi.Input[str]] = None,
                 payload_content_encoding: Optional[pulumi.Input[str]] = None,
                 payload_content_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_ref: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretV1 resources.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if all_metadata is not None:
            pulumi.set(__self__, "all_metadata", all_metadata)
        if bit_length is not None:
            pulumi.set(__self__, "bit_length", bit_length)
        if content_types is not None:
            pulumi.set(__self__, "content_types", content_types)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if creator_id is not None:
            pulumi.set(__self__, "creator_id", creator_id)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if payload is not None:
            pulumi.set(__self__, "payload", payload)
        if payload_content_encoding is not None:
            pulumi.set(__self__, "payload_content_encoding", payload_content_encoding)
        if payload_content_type is not None:
            pulumi.set(__self__, "payload_content_type", payload_content_type)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_ref is not None:
            pulumi.set(__self__, "secret_ref", secret_ref)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['SecretV1AclArgs']]:
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['SecretV1AclArgs']]):
        pulumi.set(self, "acl", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="allMetadata")
    def all_metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "all_metadata")

    @all_metadata.setter
    def all_metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "all_metadata", value)

    @property
    @pulumi.getter(name="bitLength")
    def bit_length(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "bit_length")

    @bit_length.setter
    def bit_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bit_length", value)

    @property
    @pulumi.getter(name="contentTypes")
    def content_types(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "content_types")

    @content_types.setter
    def content_types(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "content_types", value)

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
    def expiration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def payload(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload")

    @payload.setter
    def payload(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload", value)

    @property
    @pulumi.getter(name="payloadContentEncoding")
    def payload_content_encoding(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload_content_encoding")

    @payload_content_encoding.setter
    def payload_content_encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_content_encoding", value)

    @property
    @pulumi.getter(name="payloadContentType")
    def payload_content_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "payload_content_type")

    @payload_content_type.setter
    def payload_content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_content_type", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_ref")

    @secret_ref.setter
    def secret_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_ref", value)

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_type", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class SecretV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['SecretV1AclArgs', 'SecretV1AclArgsDict']]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 bit_length: Optional[pulumi.Input[int]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload: Optional[pulumi.Input[str]] = None,
                 payload_content_encoding: Optional[pulumi.Input[str]] = None,
                 payload_content_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SecretV1 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretV1Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecretV1 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SecretV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['SecretV1AclArgs', 'SecretV1AclArgsDict']]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 bit_length: Optional[pulumi.Input[int]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload: Optional[pulumi.Input[str]] = None,
                 payload_content_encoding: Optional[pulumi.Input[str]] = None,
                 payload_content_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretV1Args.__new__(SecretV1Args)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["bit_length"] = bit_length
            __props__.__dict__["expiration"] = expiration
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["mode"] = mode
            __props__.__dict__["name"] = name
            __props__.__dict__["payload"] = None if payload is None else pulumi.Output.secret(payload)
            __props__.__dict__["payload_content_encoding"] = payload_content_encoding
            __props__.__dict__["payload_content_type"] = payload_content_type
            __props__.__dict__["region"] = region
            __props__.__dict__["secret_type"] = secret_type
            __props__.__dict__["all_metadata"] = None
            __props__.__dict__["content_types"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["creator_id"] = None
            __props__.__dict__["secret_ref"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["payload"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SecretV1, __self__).__init__(
            'openstack:keymanager/secretV1:SecretV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[Union['SecretV1AclArgs', 'SecretV1AclArgsDict']]] = None,
            algorithm: Optional[pulumi.Input[str]] = None,
            all_metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            bit_length: Optional[pulumi.Input[int]] = None,
            content_types: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            creator_id: Optional[pulumi.Input[str]] = None,
            expiration: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            payload: Optional[pulumi.Input[str]] = None,
            payload_content_encoding: Optional[pulumi.Input[str]] = None,
            payload_content_type: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_ref: Optional[pulumi.Input[str]] = None,
            secret_type: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'SecretV1':
        """
        Get an existing SecretV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretV1State.__new__(_SecretV1State)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["all_metadata"] = all_metadata
        __props__.__dict__["bit_length"] = bit_length
        __props__.__dict__["content_types"] = content_types
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["creator_id"] = creator_id
        __props__.__dict__["expiration"] = expiration
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["mode"] = mode
        __props__.__dict__["name"] = name
        __props__.__dict__["payload"] = payload
        __props__.__dict__["payload_content_encoding"] = payload_content_encoding
        __props__.__dict__["payload_content_type"] = payload_content_type
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_ref"] = secret_ref
        __props__.__dict__["secret_type"] = secret_type
        __props__.__dict__["status"] = status
        __props__.__dict__["updated_at"] = updated_at
        return SecretV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.SecretV1Acl']:
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[str]:
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="allMetadata")
    def all_metadata(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "all_metadata")

    @property
    @pulumi.getter(name="bitLength")
    def bit_length(self) -> pulumi.Output[int]:
        return pulumi.get(self, "bit_length")

    @property
    @pulumi.getter(name="contentTypes")
    def content_types(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "content_types")

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
    def expiration(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def payload(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "payload")

    @property
    @pulumi.getter(name="payloadContentEncoding")
    def payload_content_encoding(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "payload_content_encoding")

    @property
    @pulumi.getter(name="payloadContentType")
    def payload_content_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "payload_content_type")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret_ref")

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret_type")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

