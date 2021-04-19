# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ContainerArgs', 'Container']

@pulumi.input_type
class ContainerArgs:
    def __init__(__self__, *,
                 container_read: Optional[pulumi.Input[str]] = None,
                 container_sync_key: Optional[pulumi.Input[str]] = None,
                 container_sync_to: Optional[pulumi.Input[str]] = None,
                 container_write: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 versioning: Optional[pulumi.Input['ContainerVersioningArgs']] = None):
        """
        The set of arguments for constructing a Container resource.
        :param pulumi.Input[str] container_read: Sets an access control list (ACL) that grants
               read access. This header can contain a comma-delimited list of users that
               can read the container (allows the GET method for all objects in the
               container). Changing this updates the access control list read access.
        :param pulumi.Input[str] container_sync_key: The secret key for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_sync_to: The destination for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_write: Sets an ACL that grants write access.
               Changing this updates the access control list write access.
        :param pulumi.Input[str] content_type: The MIME type for the container. Changing this
               updates the MIME type.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        :param pulumi.Input[Mapping[str, Any]] metadata: Custom key/value pairs to associate with the container.
               Changing this updates the existing container metadata.
        :param pulumi.Input[str] name: A unique name for the container. Changing this creates a
               new container.
        :param pulumi.Input[str] region: The region in which to create the container. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new container.
        :param pulumi.Input['ContainerVersioningArgs'] versioning: Enable object versioning. The structure is described below.
        """
        if container_read is not None:
            pulumi.set(__self__, "container_read", container_read)
        if container_sync_key is not None:
            pulumi.set(__self__, "container_sync_key", container_sync_key)
        if container_sync_to is not None:
            pulumi.set(__self__, "container_sync_to", container_sync_to)
        if container_write is not None:
            pulumi.set(__self__, "container_write", container_write)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if versioning is not None:
            pulumi.set(__self__, "versioning", versioning)

    @property
    @pulumi.getter(name="containerRead")
    def container_read(self) -> Optional[pulumi.Input[str]]:
        """
        Sets an access control list (ACL) that grants
        read access. This header can contain a comma-delimited list of users that
        can read the container (allows the GET method for all objects in the
        container). Changing this updates the access control list read access.
        """
        return pulumi.get(self, "container_read")

    @container_read.setter
    def container_read(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_read", value)

    @property
    @pulumi.getter(name="containerSyncKey")
    def container_sync_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key for container synchronization.
        Changing this updates container synchronization.
        """
        return pulumi.get(self, "container_sync_key")

    @container_sync_key.setter
    def container_sync_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_sync_key", value)

    @property
    @pulumi.getter(name="containerSyncTo")
    def container_sync_to(self) -> Optional[pulumi.Input[str]]:
        """
        The destination for container synchronization.
        Changing this updates container synchronization.
        """
        return pulumi.get(self, "container_sync_to")

    @container_sync_to.setter
    def container_sync_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_sync_to", value)

    @property
    @pulumi.getter(name="containerWrite")
    def container_write(self) -> Optional[pulumi.Input[str]]:
        """
        Sets an ACL that grants write access.
        Changing this updates the access control list write access.
        """
        return pulumi.get(self, "container_write")

    @container_write.setter
    def container_write(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_write", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The MIME type for the container. Changing this
        updates the MIME type.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Custom key/value pairs to associate with the container.
        Changing this updates the existing container metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the container. Changing this creates a
        new container.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the container. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new container.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def versioning(self) -> Optional[pulumi.Input['ContainerVersioningArgs']]:
        """
        Enable object versioning. The structure is described below.
        """
        return pulumi.get(self, "versioning")

    @versioning.setter
    def versioning(self, value: Optional[pulumi.Input['ContainerVersioningArgs']]):
        pulumi.set(self, "versioning", value)


@pulumi.input_type
class _ContainerState:
    def __init__(__self__, *,
                 container_read: Optional[pulumi.Input[str]] = None,
                 container_sync_key: Optional[pulumi.Input[str]] = None,
                 container_sync_to: Optional[pulumi.Input[str]] = None,
                 container_write: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 versioning: Optional[pulumi.Input['ContainerVersioningArgs']] = None):
        """
        Input properties used for looking up and filtering Container resources.
        :param pulumi.Input[str] container_read: Sets an access control list (ACL) that grants
               read access. This header can contain a comma-delimited list of users that
               can read the container (allows the GET method for all objects in the
               container). Changing this updates the access control list read access.
        :param pulumi.Input[str] container_sync_key: The secret key for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_sync_to: The destination for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_write: Sets an ACL that grants write access.
               Changing this updates the access control list write access.
        :param pulumi.Input[str] content_type: The MIME type for the container. Changing this
               updates the MIME type.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        :param pulumi.Input[Mapping[str, Any]] metadata: Custom key/value pairs to associate with the container.
               Changing this updates the existing container metadata.
        :param pulumi.Input[str] name: A unique name for the container. Changing this creates a
               new container.
        :param pulumi.Input[str] region: The region in which to create the container. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new container.
        :param pulumi.Input['ContainerVersioningArgs'] versioning: Enable object versioning. The structure is described below.
        """
        if container_read is not None:
            pulumi.set(__self__, "container_read", container_read)
        if container_sync_key is not None:
            pulumi.set(__self__, "container_sync_key", container_sync_key)
        if container_sync_to is not None:
            pulumi.set(__self__, "container_sync_to", container_sync_to)
        if container_write is not None:
            pulumi.set(__self__, "container_write", container_write)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if versioning is not None:
            pulumi.set(__self__, "versioning", versioning)

    @property
    @pulumi.getter(name="containerRead")
    def container_read(self) -> Optional[pulumi.Input[str]]:
        """
        Sets an access control list (ACL) that grants
        read access. This header can contain a comma-delimited list of users that
        can read the container (allows the GET method for all objects in the
        container). Changing this updates the access control list read access.
        """
        return pulumi.get(self, "container_read")

    @container_read.setter
    def container_read(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_read", value)

    @property
    @pulumi.getter(name="containerSyncKey")
    def container_sync_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key for container synchronization.
        Changing this updates container synchronization.
        """
        return pulumi.get(self, "container_sync_key")

    @container_sync_key.setter
    def container_sync_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_sync_key", value)

    @property
    @pulumi.getter(name="containerSyncTo")
    def container_sync_to(self) -> Optional[pulumi.Input[str]]:
        """
        The destination for container synchronization.
        Changing this updates container synchronization.
        """
        return pulumi.get(self, "container_sync_to")

    @container_sync_to.setter
    def container_sync_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_sync_to", value)

    @property
    @pulumi.getter(name="containerWrite")
    def container_write(self) -> Optional[pulumi.Input[str]]:
        """
        Sets an ACL that grants write access.
        Changing this updates the access control list write access.
        """
        return pulumi.get(self, "container_write")

    @container_write.setter
    def container_write(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_write", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The MIME type for the container. Changing this
        updates the MIME type.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Custom key/value pairs to associate with the container.
        Changing this updates the existing container metadata.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the container. Changing this creates a
        new container.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the container. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new container.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def versioning(self) -> Optional[pulumi.Input['ContainerVersioningArgs']]:
        """
        Enable object versioning. The structure is described below.
        """
        return pulumi.get(self, "versioning")

    @versioning.setter
    def versioning(self, value: Optional[pulumi.Input['ContainerVersioningArgs']]):
        pulumi.set(self, "versioning", value)


class Container(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_read: Optional[pulumi.Input[str]] = None,
                 container_sync_key: Optional[pulumi.Input[str]] = None,
                 container_sync_to: Optional[pulumi.Input[str]] = None,
                 container_write: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 versioning: Optional[pulumi.Input[pulumi.InputType['ContainerVersioningArgs']]] = None,
                 __props__=None):
        """
        Manages a V1 container resource within OpenStack.

        ## Example Usage
        ### Basic Container

        ```python
        import pulumi
        import pulumi_openstack as openstack

        container1 = openstack.objectstorage.Container("container1",
            content_type="application/json",
            metadata={
                "test": "true",
            },
            region="RegionOne",
            versioning=openstack.objectstorage.ContainerVersioningArgs(
                location="tf-test-container-versions",
                type="versions",
            ))
        ```
        ### Global Read Access

        ```python
        import pulumi
        import pulumi_openstack as openstack

        container1 = openstack.objectstorage.Container("container1",
            container_read=".r:*",
            region="RegionOne")
        ```
        ### Global Read and List Access

        ```python
        import pulumi
        import pulumi_openstack as openstack

        container1 = openstack.objectstorage.Container("container1",
            container_read=".r:*,.rlistings",
            region="RegionOne")
        ```
        ### Write-Only Access for a User

        ```python
        import pulumi
        import pulumi_openstack as openstack

        current = openstack.identity.get_auth_scope(name="current")
        container1 = openstack.objectstorage.Container("container1",
            container_read=f".r:-{var['username']}",
            container_write=f"{current.project_id}:{var['username']}",
            region="RegionOne")
        ```

        ## Import

        This resource can be imported by specifying the name of the containerSome attributes can't be imported * `force_destroy` * `content_type` * `metadata` * `container_sync_to` * `container_sync_key` So you'll have to `terraform plan` and `terraform apply` after the import to fix those missing attributes.

        ```sh
         $ pulumi import openstack:objectstorage/container:Container container_1 <name>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_read: Sets an access control list (ACL) that grants
               read access. This header can contain a comma-delimited list of users that
               can read the container (allows the GET method for all objects in the
               container). Changing this updates the access control list read access.
        :param pulumi.Input[str] container_sync_key: The secret key for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_sync_to: The destination for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_write: Sets an ACL that grants write access.
               Changing this updates the access control list write access.
        :param pulumi.Input[str] content_type: The MIME type for the container. Changing this
               updates the MIME type.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        :param pulumi.Input[Mapping[str, Any]] metadata: Custom key/value pairs to associate with the container.
               Changing this updates the existing container metadata.
        :param pulumi.Input[str] name: A unique name for the container. Changing this creates a
               new container.
        :param pulumi.Input[str] region: The region in which to create the container. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new container.
        :param pulumi.Input[pulumi.InputType['ContainerVersioningArgs']] versioning: Enable object versioning. The structure is described below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContainerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V1 container resource within OpenStack.

        ## Example Usage
        ### Basic Container

        ```python
        import pulumi
        import pulumi_openstack as openstack

        container1 = openstack.objectstorage.Container("container1",
            content_type="application/json",
            metadata={
                "test": "true",
            },
            region="RegionOne",
            versioning=openstack.objectstorage.ContainerVersioningArgs(
                location="tf-test-container-versions",
                type="versions",
            ))
        ```
        ### Global Read Access

        ```python
        import pulumi
        import pulumi_openstack as openstack

        container1 = openstack.objectstorage.Container("container1",
            container_read=".r:*",
            region="RegionOne")
        ```
        ### Global Read and List Access

        ```python
        import pulumi
        import pulumi_openstack as openstack

        container1 = openstack.objectstorage.Container("container1",
            container_read=".r:*,.rlistings",
            region="RegionOne")
        ```
        ### Write-Only Access for a User

        ```python
        import pulumi
        import pulumi_openstack as openstack

        current = openstack.identity.get_auth_scope(name="current")
        container1 = openstack.objectstorage.Container("container1",
            container_read=f".r:-{var['username']}",
            container_write=f"{current.project_id}:{var['username']}",
            region="RegionOne")
        ```

        ## Import

        This resource can be imported by specifying the name of the containerSome attributes can't be imported * `force_destroy` * `content_type` * `metadata` * `container_sync_to` * `container_sync_key` So you'll have to `terraform plan` and `terraform apply` after the import to fix those missing attributes.

        ```sh
         $ pulumi import openstack:objectstorage/container:Container container_1 <name>
        ```

        :param str resource_name: The name of the resource.
        :param ContainerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_read: Optional[pulumi.Input[str]] = None,
                 container_sync_key: Optional[pulumi.Input[str]] = None,
                 container_sync_to: Optional[pulumi.Input[str]] = None,
                 container_write: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 versioning: Optional[pulumi.Input[pulumi.InputType['ContainerVersioningArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerArgs.__new__(ContainerArgs)

            __props__.__dict__["container_read"] = container_read
            __props__.__dict__["container_sync_key"] = container_sync_key
            __props__.__dict__["container_sync_to"] = container_sync_to
            __props__.__dict__["container_write"] = container_write
            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["force_destroy"] = force_destroy
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["versioning"] = versioning
        super(Container, __self__).__init__(
            'openstack:objectstorage/container:Container',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            container_read: Optional[pulumi.Input[str]] = None,
            container_sync_key: Optional[pulumi.Input[str]] = None,
            container_sync_to: Optional[pulumi.Input[str]] = None,
            container_write: Optional[pulumi.Input[str]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            versioning: Optional[pulumi.Input[pulumi.InputType['ContainerVersioningArgs']]] = None) -> 'Container':
        """
        Get an existing Container resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_read: Sets an access control list (ACL) that grants
               read access. This header can contain a comma-delimited list of users that
               can read the container (allows the GET method for all objects in the
               container). Changing this updates the access control list read access.
        :param pulumi.Input[str] container_sync_key: The secret key for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_sync_to: The destination for container synchronization.
               Changing this updates container synchronization.
        :param pulumi.Input[str] container_write: Sets an ACL that grants write access.
               Changing this updates the access control list write access.
        :param pulumi.Input[str] content_type: The MIME type for the container. Changing this
               updates the MIME type.
        :param pulumi.Input[bool] force_destroy: A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        :param pulumi.Input[Mapping[str, Any]] metadata: Custom key/value pairs to associate with the container.
               Changing this updates the existing container metadata.
        :param pulumi.Input[str] name: A unique name for the container. Changing this creates a
               new container.
        :param pulumi.Input[str] region: The region in which to create the container. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new container.
        :param pulumi.Input[pulumi.InputType['ContainerVersioningArgs']] versioning: Enable object versioning. The structure is described below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerState.__new__(_ContainerState)

        __props__.__dict__["container_read"] = container_read
        __props__.__dict__["container_sync_key"] = container_sync_key
        __props__.__dict__["container_sync_to"] = container_sync_to
        __props__.__dict__["container_write"] = container_write
        __props__.__dict__["content_type"] = content_type
        __props__.__dict__["force_destroy"] = force_destroy
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["versioning"] = versioning
        return Container(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="containerRead")
    def container_read(self) -> pulumi.Output[Optional[str]]:
        """
        Sets an access control list (ACL) that grants
        read access. This header can contain a comma-delimited list of users that
        can read the container (allows the GET method for all objects in the
        container). Changing this updates the access control list read access.
        """
        return pulumi.get(self, "container_read")

    @property
    @pulumi.getter(name="containerSyncKey")
    def container_sync_key(self) -> pulumi.Output[Optional[str]]:
        """
        The secret key for container synchronization.
        Changing this updates container synchronization.
        """
        return pulumi.get(self, "container_sync_key")

    @property
    @pulumi.getter(name="containerSyncTo")
    def container_sync_to(self) -> pulumi.Output[Optional[str]]:
        """
        The destination for container synchronization.
        Changing this updates container synchronization.
        """
        return pulumi.get(self, "container_sync_to")

    @property
    @pulumi.getter(name="containerWrite")
    def container_write(self) -> pulumi.Output[Optional[str]]:
        """
        Sets an ACL that grants write access.
        Changing this updates the access control list write access.
        """
        return pulumi.get(self, "container_write")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        """
        The MIME type for the container. Changing this
        updates the MIME type.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Custom key/value pairs to associate with the container.
        Changing this updates the existing container metadata.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the container. Changing this creates a
        new container.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the container. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new container.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def versioning(self) -> pulumi.Output[Optional['outputs.ContainerVersioning']]:
        """
        Enable object versioning. The structure is described below.
        """
        return pulumi.get(self, "versioning")

