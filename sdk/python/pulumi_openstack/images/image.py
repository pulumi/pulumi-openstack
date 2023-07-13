# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 container_format: pulumi.Input[str],
                 disk_format: pulumi.Input[str],
                 decompress: Optional[pulumi.Input[bool]] = None,
                 hidden: Optional[pulumi.Input[bool]] = None,
                 image_cache_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_source_password: Optional[pulumi.Input[str]] = None,
                 image_source_url: Optional[pulumi.Input[str]] = None,
                 image_source_username: Optional[pulumi.Input[str]] = None,
                 local_file_path: Optional[pulumi.Input[str]] = None,
                 min_disk_gb: Optional[pulumi.Input[int]] = None,
                 min_ram_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 verify_checksum: Optional[pulumi.Input[bool]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 web_download: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Image resource.
        """
        pulumi.set(__self__, "container_format", container_format)
        pulumi.set(__self__, "disk_format", disk_format)
        if decompress is not None:
            pulumi.set(__self__, "decompress", decompress)
        if hidden is not None:
            pulumi.set(__self__, "hidden", hidden)
        if image_cache_path is not None:
            pulumi.set(__self__, "image_cache_path", image_cache_path)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if image_source_password is not None:
            pulumi.set(__self__, "image_source_password", image_source_password)
        if image_source_url is not None:
            pulumi.set(__self__, "image_source_url", image_source_url)
        if image_source_username is not None:
            pulumi.set(__self__, "image_source_username", image_source_username)
        if local_file_path is not None:
            pulumi.set(__self__, "local_file_path", local_file_path)
        if min_disk_gb is not None:
            pulumi.set(__self__, "min_disk_gb", min_disk_gb)
        if min_ram_mb is not None:
            pulumi.set(__self__, "min_ram_mb", min_ram_mb)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if protected is not None:
            pulumi.set(__self__, "protected", protected)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if verify_checksum is not None:
            pulumi.set(__self__, "verify_checksum", verify_checksum)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)
        if web_download is not None:
            pulumi.set(__self__, "web_download", web_download)

    @property
    @pulumi.getter(name="containerFormat")
    def container_format(self) -> pulumi.Input[str]:
        return pulumi.get(self, "container_format")

    @container_format.setter
    def container_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_format", value)

    @property
    @pulumi.getter(name="diskFormat")
    def disk_format(self) -> pulumi.Input[str]:
        return pulumi.get(self, "disk_format")

    @disk_format.setter
    def disk_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk_format", value)

    @property
    @pulumi.getter
    def decompress(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "decompress")

    @decompress.setter
    def decompress(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "decompress", value)

    @property
    @pulumi.getter
    def hidden(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "hidden")

    @hidden.setter
    def hidden(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hidden", value)

    @property
    @pulumi.getter(name="imageCachePath")
    def image_cache_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_cache_path")

    @image_cache_path.setter
    def image_cache_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_cache_path", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="imageSourcePassword")
    def image_source_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_source_password")

    @image_source_password.setter
    def image_source_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_source_password", value)

    @property
    @pulumi.getter(name="imageSourceUrl")
    def image_source_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_source_url")

    @image_source_url.setter
    def image_source_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_source_url", value)

    @property
    @pulumi.getter(name="imageSourceUsername")
    def image_source_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_source_username")

    @image_source_username.setter
    def image_source_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_source_username", value)

    @property
    @pulumi.getter(name="localFilePath")
    def local_file_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_file_path")

    @local_file_path.setter
    def local_file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_file_path", value)

    @property
    @pulumi.getter(name="minDiskGb")
    def min_disk_gb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_disk_gb")

    @min_disk_gb.setter
    def min_disk_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_disk_gb", value)

    @property
    @pulumi.getter(name="minRamMb")
    def min_ram_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_ram_mb")

    @min_ram_mb.setter
    def min_ram_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ram_mb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def protected(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "protected")

    @protected.setter
    def protected(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protected", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="verifyChecksum")
    def verify_checksum(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "verify_checksum")

    @verify_checksum.setter
    def verify_checksum(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_checksum", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)

    @property
    @pulumi.getter(name="webDownload")
    def web_download(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "web_download")

    @web_download.setter
    def web_download(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "web_download", value)


@pulumi.input_type
class _ImageState:
    def __init__(__self__, *,
                 checksum: Optional[pulumi.Input[str]] = None,
                 container_format: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 decompress: Optional[pulumi.Input[bool]] = None,
                 disk_format: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 hidden: Optional[pulumi.Input[bool]] = None,
                 image_cache_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_source_password: Optional[pulumi.Input[str]] = None,
                 image_source_url: Optional[pulumi.Input[str]] = None,
                 image_source_username: Optional[pulumi.Input[str]] = None,
                 local_file_path: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 min_disk_gb: Optional[pulumi.Input[int]] = None,
                 min_ram_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 size_bytes: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 update_at: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 verify_checksum: Optional[pulumi.Input[bool]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 web_download: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Image resources.
        """
        if checksum is not None:
            pulumi.set(__self__, "checksum", checksum)
        if container_format is not None:
            pulumi.set(__self__, "container_format", container_format)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if decompress is not None:
            pulumi.set(__self__, "decompress", decompress)
        if disk_format is not None:
            pulumi.set(__self__, "disk_format", disk_format)
        if file is not None:
            pulumi.set(__self__, "file", file)
        if hidden is not None:
            pulumi.set(__self__, "hidden", hidden)
        if image_cache_path is not None:
            pulumi.set(__self__, "image_cache_path", image_cache_path)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if image_source_password is not None:
            pulumi.set(__self__, "image_source_password", image_source_password)
        if image_source_url is not None:
            pulumi.set(__self__, "image_source_url", image_source_url)
        if image_source_username is not None:
            pulumi.set(__self__, "image_source_username", image_source_username)
        if local_file_path is not None:
            pulumi.set(__self__, "local_file_path", local_file_path)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if min_disk_gb is not None:
            pulumi.set(__self__, "min_disk_gb", min_disk_gb)
        if min_ram_mb is not None:
            pulumi.set(__self__, "min_ram_mb", min_ram_mb)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if protected is not None:
            pulumi.set(__self__, "protected", protected)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if schema is not None:
            pulumi.set(__self__, "schema", schema)
        if size_bytes is not None:
            pulumi.set(__self__, "size_bytes", size_bytes)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if update_at is not None:
            warnings.warn("""Use updated_at instead""", DeprecationWarning)
            pulumi.log.warn("""update_at is deprecated: Use updated_at instead""")
        if update_at is not None:
            pulumi.set(__self__, "update_at", update_at)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if verify_checksum is not None:
            pulumi.set(__self__, "verify_checksum", verify_checksum)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)
        if web_download is not None:
            pulumi.set(__self__, "web_download", web_download)

    @property
    @pulumi.getter
    def checksum(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "checksum")

    @checksum.setter
    def checksum(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "checksum", value)

    @property
    @pulumi.getter(name="containerFormat")
    def container_format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container_format")

    @container_format.setter
    def container_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_format", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def decompress(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "decompress")

    @decompress.setter
    def decompress(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "decompress", value)

    @property
    @pulumi.getter(name="diskFormat")
    def disk_format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "disk_format")

    @disk_format.setter
    def disk_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_format", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def hidden(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "hidden")

    @hidden.setter
    def hidden(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hidden", value)

    @property
    @pulumi.getter(name="imageCachePath")
    def image_cache_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_cache_path")

    @image_cache_path.setter
    def image_cache_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_cache_path", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="imageSourcePassword")
    def image_source_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_source_password")

    @image_source_password.setter
    def image_source_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_source_password", value)

    @property
    @pulumi.getter(name="imageSourceUrl")
    def image_source_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_source_url")

    @image_source_url.setter
    def image_source_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_source_url", value)

    @property
    @pulumi.getter(name="imageSourceUsername")
    def image_source_username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_source_username")

    @image_source_username.setter
    def image_source_username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_source_username", value)

    @property
    @pulumi.getter(name="localFilePath")
    def local_file_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_file_path")

    @local_file_path.setter
    def local_file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_file_path", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter(name="minDiskGb")
    def min_disk_gb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_disk_gb")

    @min_disk_gb.setter
    def min_disk_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_disk_gb", value)

    @property
    @pulumi.getter(name="minRamMb")
    def min_ram_mb(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_ram_mb")

    @min_ram_mb.setter
    def min_ram_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_ram_mb", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def protected(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "protected")

    @protected.setter
    def protected(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protected", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def schema(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter(name="sizeBytes")
    def size_bytes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "size_bytes")

    @size_bytes.setter
    def size_bytes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_bytes", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updateAt")
    def update_at(self) -> Optional[pulumi.Input[str]]:
        warnings.warn("""Use updated_at instead""", DeprecationWarning)
        pulumi.log.warn("""update_at is deprecated: Use updated_at instead""")

        return pulumi.get(self, "update_at")

    @update_at.setter
    def update_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_at", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="verifyChecksum")
    def verify_checksum(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "verify_checksum")

    @verify_checksum.setter
    def verify_checksum(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_checksum", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)

    @property
    @pulumi.getter(name="webDownload")
    def web_download(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "web_download")

    @web_download.setter
    def web_download(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "web_download", value)


class Image(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_format: Optional[pulumi.Input[str]] = None,
                 decompress: Optional[pulumi.Input[bool]] = None,
                 disk_format: Optional[pulumi.Input[str]] = None,
                 hidden: Optional[pulumi.Input[bool]] = None,
                 image_cache_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_source_password: Optional[pulumi.Input[str]] = None,
                 image_source_url: Optional[pulumi.Input[str]] = None,
                 image_source_username: Optional[pulumi.Input[str]] = None,
                 local_file_path: Optional[pulumi.Input[str]] = None,
                 min_disk_gb: Optional[pulumi.Input[int]] = None,
                 min_ram_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 verify_checksum: Optional[pulumi.Input[bool]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 web_download: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a Image resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Image resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_format: Optional[pulumi.Input[str]] = None,
                 decompress: Optional[pulumi.Input[bool]] = None,
                 disk_format: Optional[pulumi.Input[str]] = None,
                 hidden: Optional[pulumi.Input[bool]] = None,
                 image_cache_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 image_source_password: Optional[pulumi.Input[str]] = None,
                 image_source_url: Optional[pulumi.Input[str]] = None,
                 image_source_username: Optional[pulumi.Input[str]] = None,
                 local_file_path: Optional[pulumi.Input[str]] = None,
                 min_disk_gb: Optional[pulumi.Input[int]] = None,
                 min_ram_mb: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 protected: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 verify_checksum: Optional[pulumi.Input[bool]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 web_download: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageArgs.__new__(ImageArgs)

            if container_format is None and not opts.urn:
                raise TypeError("Missing required property 'container_format'")
            __props__.__dict__["container_format"] = container_format
            __props__.__dict__["decompress"] = decompress
            if disk_format is None and not opts.urn:
                raise TypeError("Missing required property 'disk_format'")
            __props__.__dict__["disk_format"] = disk_format
            __props__.__dict__["hidden"] = hidden
            __props__.__dict__["image_cache_path"] = image_cache_path
            __props__.__dict__["image_id"] = image_id
            __props__.__dict__["image_source_password"] = None if image_source_password is None else pulumi.Output.secret(image_source_password)
            __props__.__dict__["image_source_url"] = image_source_url
            __props__.__dict__["image_source_username"] = image_source_username
            __props__.__dict__["local_file_path"] = local_file_path
            __props__.__dict__["min_disk_gb"] = min_disk_gb
            __props__.__dict__["min_ram_mb"] = min_ram_mb
            __props__.__dict__["name"] = name
            __props__.__dict__["properties"] = properties
            __props__.__dict__["protected"] = protected
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
            __props__.__dict__["verify_checksum"] = verify_checksum
            __props__.__dict__["visibility"] = visibility
            __props__.__dict__["web_download"] = web_download
            __props__.__dict__["checksum"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["file"] = None
            __props__.__dict__["metadata"] = None
            __props__.__dict__["owner"] = None
            __props__.__dict__["schema"] = None
            __props__.__dict__["size_bytes"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_at"] = None
            __props__.__dict__["updated_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["imageSourcePassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Image, __self__).__init__(
            'openstack:images/image:Image',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            checksum: Optional[pulumi.Input[str]] = None,
            container_format: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            decompress: Optional[pulumi.Input[bool]] = None,
            disk_format: Optional[pulumi.Input[str]] = None,
            file: Optional[pulumi.Input[str]] = None,
            hidden: Optional[pulumi.Input[bool]] = None,
            image_cache_path: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            image_source_password: Optional[pulumi.Input[str]] = None,
            image_source_url: Optional[pulumi.Input[str]] = None,
            image_source_username: Optional[pulumi.Input[str]] = None,
            local_file_path: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            min_disk_gb: Optional[pulumi.Input[int]] = None,
            min_ram_mb: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            properties: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            protected: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            schema: Optional[pulumi.Input[str]] = None,
            size_bytes: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            update_at: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            verify_checksum: Optional[pulumi.Input[bool]] = None,
            visibility: Optional[pulumi.Input[str]] = None,
            web_download: Optional[pulumi.Input[bool]] = None) -> 'Image':
        """
        Get an existing Image resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImageState.__new__(_ImageState)

        __props__.__dict__["checksum"] = checksum
        __props__.__dict__["container_format"] = container_format
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["decompress"] = decompress
        __props__.__dict__["disk_format"] = disk_format
        __props__.__dict__["file"] = file
        __props__.__dict__["hidden"] = hidden
        __props__.__dict__["image_cache_path"] = image_cache_path
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["image_source_password"] = image_source_password
        __props__.__dict__["image_source_url"] = image_source_url
        __props__.__dict__["image_source_username"] = image_source_username
        __props__.__dict__["local_file_path"] = local_file_path
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["min_disk_gb"] = min_disk_gb
        __props__.__dict__["min_ram_mb"] = min_ram_mb
        __props__.__dict__["name"] = name
        __props__.__dict__["owner"] = owner
        __props__.__dict__["properties"] = properties
        __props__.__dict__["protected"] = protected
        __props__.__dict__["region"] = region
        __props__.__dict__["schema"] = schema
        __props__.__dict__["size_bytes"] = size_bytes
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["update_at"] = update_at
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["verify_checksum"] = verify_checksum
        __props__.__dict__["visibility"] = visibility
        __props__.__dict__["web_download"] = web_download
        return Image(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def checksum(self) -> pulumi.Output[str]:
        return pulumi.get(self, "checksum")

    @property
    @pulumi.getter(name="containerFormat")
    def container_format(self) -> pulumi.Output[str]:
        return pulumi.get(self, "container_format")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def decompress(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "decompress")

    @property
    @pulumi.getter(name="diskFormat")
    def disk_format(self) -> pulumi.Output[str]:
        return pulumi.get(self, "disk_format")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output[str]:
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def hidden(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "hidden")

    @property
    @pulumi.getter(name="imageCachePath")
    def image_cache_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_cache_path")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageSourcePassword")
    def image_source_password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_source_password")

    @property
    @pulumi.getter(name="imageSourceUrl")
    def image_source_url(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_source_url")

    @property
    @pulumi.getter(name="imageSourceUsername")
    def image_source_username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_source_username")

    @property
    @pulumi.getter(name="localFilePath")
    def local_file_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "local_file_path")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, Any]]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="minDiskGb")
    def min_disk_gb(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "min_disk_gb")

    @property
    @pulumi.getter(name="minRamMb")
    def min_ram_mb(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "min_ram_mb")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output[Mapping[str, Any]]:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def protected(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "protected")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[str]:
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter(name="sizeBytes")
    def size_bytes(self) -> pulumi.Output[int]:
        return pulumi.get(self, "size_bytes")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateAt")
    def update_at(self) -> pulumi.Output[str]:
        warnings.warn("""Use updated_at instead""", DeprecationWarning)
        pulumi.log.warn("""update_at is deprecated: Use updated_at instead""")

        return pulumi.get(self, "update_at")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="verifyChecksum")
    def verify_checksum(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "verify_checksum")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter(name="webDownload")
    def web_download(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "web_download")

