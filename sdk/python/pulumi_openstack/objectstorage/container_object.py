# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ContainerObjectArgs', 'ContainerObject']

@pulumi.input_type
class ContainerObjectArgs:
    def __init__(__self__, *,
                 container_name: pulumi.Input[str],
                 content: Optional[pulumi.Input[str]] = None,
                 content_disposition: Optional[pulumi.Input[str]] = None,
                 content_encoding: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 copy_from: Optional[pulumi.Input[str]] = None,
                 delete_after: Optional[pulumi.Input[int]] = None,
                 delete_at: Optional[pulumi.Input[str]] = None,
                 detect_content_type: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_manifest: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ContainerObject resource.
        """
        pulumi.set(__self__, "container_name", container_name)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_disposition is not None:
            pulumi.set(__self__, "content_disposition", content_disposition)
        if content_encoding is not None:
            pulumi.set(__self__, "content_encoding", content_encoding)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if copy_from is not None:
            pulumi.set(__self__, "copy_from", copy_from)
        if delete_after is not None:
            pulumi.set(__self__, "delete_after", delete_after)
        if delete_at is not None:
            pulumi.set(__self__, "delete_at", delete_at)
        if detect_content_type is not None:
            pulumi.set(__self__, "detect_content_type", detect_content_type)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object_manifest is not None:
            pulumi.set(__self__, "object_manifest", object_manifest)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "container_name")

    @container_name.setter
    def container_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_name", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentDisposition")
    def content_disposition(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_disposition")

    @content_disposition.setter
    def content_disposition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_disposition", value)

    @property
    @pulumi.getter(name="contentEncoding")
    def content_encoding(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_encoding")

    @content_encoding.setter
    def content_encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_encoding", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="copyFrom")
    def copy_from(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "copy_from")

    @copy_from.setter
    def copy_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "copy_from", value)

    @property
    @pulumi.getter(name="deleteAfter")
    def delete_after(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "delete_after")

    @delete_after.setter
    def delete_after(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delete_after", value)

    @property
    @pulumi.getter(name="deleteAt")
    def delete_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "delete_at")

    @delete_at.setter
    def delete_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_at", value)

    @property
    @pulumi.getter(name="detectContentType")
    def detect_content_type(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "detect_content_type")

    @detect_content_type.setter
    def detect_content_type(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "detect_content_type", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectManifest")
    def object_manifest(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "object_manifest")

    @object_manifest.setter
    def object_manifest(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_manifest", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)


@pulumi.input_type
class _ContainerObjectState:
    def __init__(__self__, *,
                 container_name: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_disposition: Optional[pulumi.Input[str]] = None,
                 content_encoding: Optional[pulumi.Input[str]] = None,
                 content_length: Optional[pulumi.Input[int]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 copy_from: Optional[pulumi.Input[str]] = None,
                 date: Optional[pulumi.Input[str]] = None,
                 delete_after: Optional[pulumi.Input[int]] = None,
                 delete_at: Optional[pulumi.Input[str]] = None,
                 detect_content_type: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 last_modified: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_manifest: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 trans_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ContainerObject resources.
        """
        if container_name is not None:
            pulumi.set(__self__, "container_name", container_name)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_disposition is not None:
            pulumi.set(__self__, "content_disposition", content_disposition)
        if content_encoding is not None:
            pulumi.set(__self__, "content_encoding", content_encoding)
        if content_length is not None:
            pulumi.set(__self__, "content_length", content_length)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if copy_from is not None:
            pulumi.set(__self__, "copy_from", copy_from)
        if date is not None:
            pulumi.set(__self__, "date", date)
        if delete_after is not None:
            pulumi.set(__self__, "delete_after", delete_after)
        if delete_at is not None:
            pulumi.set(__self__, "delete_at", delete_at)
        if detect_content_type is not None:
            pulumi.set(__self__, "detect_content_type", detect_content_type)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if last_modified is not None:
            pulumi.set(__self__, "last_modified", last_modified)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object_manifest is not None:
            pulumi.set(__self__, "object_manifest", object_manifest)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if trans_id is not None:
            pulumi.set(__self__, "trans_id", trans_id)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container_name")

    @container_name.setter
    def container_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_name", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentDisposition")
    def content_disposition(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_disposition")

    @content_disposition.setter
    def content_disposition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_disposition", value)

    @property
    @pulumi.getter(name="contentEncoding")
    def content_encoding(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_encoding")

    @content_encoding.setter
    def content_encoding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_encoding", value)

    @property
    @pulumi.getter(name="contentLength")
    def content_length(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "content_length")

    @content_length.setter
    def content_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "content_length", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="copyFrom")
    def copy_from(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "copy_from")

    @copy_from.setter
    def copy_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "copy_from", value)

    @property
    @pulumi.getter
    def date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "date")

    @date.setter
    def date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "date", value)

    @property
    @pulumi.getter(name="deleteAfter")
    def delete_after(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "delete_after")

    @delete_after.setter
    def delete_after(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delete_after", value)

    @property
    @pulumi.getter(name="deleteAt")
    def delete_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "delete_at")

    @delete_at.setter
    def delete_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_at", value)

    @property
    @pulumi.getter(name="detectContentType")
    def detect_content_type(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "detect_content_type")

    @detect_content_type.setter
    def detect_content_type(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "detect_content_type", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "last_modified")

    @last_modified.setter
    def last_modified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modified", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectManifest")
    def object_manifest(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "object_manifest")

    @object_manifest.setter
    def object_manifest(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_manifest", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter(name="transId")
    def trans_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trans_id")

    @trans_id.setter
    def trans_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trans_id", value)


class ContainerObject(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_name: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_disposition: Optional[pulumi.Input[str]] = None,
                 content_encoding: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 copy_from: Optional[pulumi.Input[str]] = None,
                 delete_after: Optional[pulumi.Input[int]] = None,
                 delete_at: Optional[pulumi.Input[str]] = None,
                 detect_content_type: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_manifest: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ContainerObject resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerObjectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ContainerObject resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ContainerObjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerObjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_name: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_disposition: Optional[pulumi.Input[str]] = None,
                 content_encoding: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 copy_from: Optional[pulumi.Input[str]] = None,
                 delete_after: Optional[pulumi.Input[int]] = None,
                 delete_at: Optional[pulumi.Input[str]] = None,
                 detect_content_type: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_manifest: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerObjectArgs.__new__(ContainerObjectArgs)

            if container_name is None and not opts.urn:
                raise TypeError("Missing required property 'container_name'")
            __props__.__dict__["container_name"] = container_name
            __props__.__dict__["content"] = content
            __props__.__dict__["content_disposition"] = content_disposition
            __props__.__dict__["content_encoding"] = content_encoding
            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["copy_from"] = copy_from
            __props__.__dict__["delete_after"] = delete_after
            __props__.__dict__["delete_at"] = delete_at
            __props__.__dict__["detect_content_type"] = detect_content_type
            __props__.__dict__["etag"] = etag
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["object_manifest"] = object_manifest
            __props__.__dict__["region"] = region
            __props__.__dict__["source"] = source
            __props__.__dict__["content_length"] = None
            __props__.__dict__["date"] = None
            __props__.__dict__["last_modified"] = None
            __props__.__dict__["trans_id"] = None
        super(ContainerObject, __self__).__init__(
            'openstack:objectstorage/containerObject:ContainerObject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_name: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            content_disposition: Optional[pulumi.Input[str]] = None,
            content_encoding: Optional[pulumi.Input[str]] = None,
            content_length: Optional[pulumi.Input[int]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            copy_from: Optional[pulumi.Input[str]] = None,
            date: Optional[pulumi.Input[str]] = None,
            delete_after: Optional[pulumi.Input[int]] = None,
            delete_at: Optional[pulumi.Input[str]] = None,
            detect_content_type: Optional[pulumi.Input[bool]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            last_modified: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            object_manifest: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None,
            trans_id: Optional[pulumi.Input[str]] = None) -> 'ContainerObject':
        """
        Get an existing ContainerObject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerObjectState.__new__(_ContainerObjectState)

        __props__.__dict__["container_name"] = container_name
        __props__.__dict__["content"] = content
        __props__.__dict__["content_disposition"] = content_disposition
        __props__.__dict__["content_encoding"] = content_encoding
        __props__.__dict__["content_length"] = content_length
        __props__.__dict__["content_type"] = content_type
        __props__.__dict__["copy_from"] = copy_from
        __props__.__dict__["date"] = date
        __props__.__dict__["delete_after"] = delete_after
        __props__.__dict__["delete_at"] = delete_at
        __props__.__dict__["detect_content_type"] = detect_content_type
        __props__.__dict__["etag"] = etag
        __props__.__dict__["last_modified"] = last_modified
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["object_manifest"] = object_manifest
        __props__.__dict__["region"] = region
        __props__.__dict__["source"] = source
        __props__.__dict__["trans_id"] = trans_id
        return ContainerObject(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "container_name")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentDisposition")
    def content_disposition(self) -> pulumi.Output[str]:
        return pulumi.get(self, "content_disposition")

    @property
    @pulumi.getter(name="contentEncoding")
    def content_encoding(self) -> pulumi.Output[str]:
        return pulumi.get(self, "content_encoding")

    @property
    @pulumi.getter(name="contentLength")
    def content_length(self) -> pulumi.Output[int]:
        return pulumi.get(self, "content_length")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="copyFrom")
    def copy_from(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "copy_from")

    @property
    @pulumi.getter
    def date(self) -> pulumi.Output[str]:
        return pulumi.get(self, "date")

    @property
    @pulumi.getter(name="deleteAfter")
    def delete_after(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "delete_after")

    @property
    @pulumi.getter(name="deleteAt")
    def delete_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "delete_at")

    @property
    @pulumi.getter(name="detectContentType")
    def detect_content_type(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "detect_content_type")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectManifest")
    def object_manifest(self) -> pulumi.Output[str]:
        return pulumi.get(self, "object_manifest")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="transId")
    def trans_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trans_id")

