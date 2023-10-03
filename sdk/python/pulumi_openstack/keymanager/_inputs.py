# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ContainerV1AclArgs',
    'ContainerV1AclReadArgs',
    'ContainerV1ConsumerArgs',
    'ContainerV1SecretRefArgs',
    'OrderV1MetaArgs',
    'SecretV1AclArgs',
    'SecretV1AclReadArgs',
]

@pulumi.input_type
class ContainerV1AclArgs:
    def __init__(__self__, *,
                 read: Optional[pulumi.Input['ContainerV1AclReadArgs']] = None):
        ContainerV1AclArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            read=read,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             read: Optional[pulumi.Input['ContainerV1AclReadArgs']] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if read is not None:
            _setter("read", read)

    @property
    @pulumi.getter
    def read(self) -> Optional[pulumi.Input['ContainerV1AclReadArgs']]:
        return pulumi.get(self, "read")

    @read.setter
    def read(self, value: Optional[pulumi.Input['ContainerV1AclReadArgs']]):
        pulumi.set(self, "read", value)


@pulumi.input_type
class ContainerV1AclReadArgs:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 project_access: Optional[pulumi.Input[bool]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] created_at: The date the container ACL was created.
        :param pulumi.Input[bool] project_access: Whether the container is accessible project wide.
               Defaults to `true`.
        :param pulumi.Input[str] updated_at: The date the container ACL was last updated.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: The list of user IDs, which are allowed to access the
               container, when `project_access` is set to `false`.
        """
        ContainerV1AclReadArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            created_at=created_at,
            project_access=project_access,
            updated_at=updated_at,
            users=users,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             created_at: Optional[pulumi.Input[str]] = None,
             project_access: Optional[pulumi.Input[bool]] = None,
             updated_at: Optional[pulumi.Input[str]] = None,
             users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if created_at is not None:
            _setter("created_at", created_at)
        if project_access is not None:
            _setter("project_access", project_access)
        if updated_at is not None:
            _setter("updated_at", updated_at)
        if users is not None:
            _setter("users", users)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date the container ACL was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="projectAccess")
    def project_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the container is accessible project wide.
        Defaults to `true`.
        """
        return pulumi.get(self, "project_access")

    @project_access.setter
    def project_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "project_access", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date the container ACL was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of user IDs, which are allowed to access the
        container, when `project_access` is set to `false`.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users", value)


@pulumi.input_type
class ContainerV1ConsumerArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[str] url: The consumer URL.
        """
        ContainerV1ConsumerArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            url=url,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[pulumi.Input[str]] = None,
             url: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if name is not None:
            _setter("name", name)
        if url is not None:
            _setter("url", url)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable name for the Container. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The consumer URL.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class ContainerV1SecretRefArgs:
    def __init__(__self__, *,
                 secret_ref: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] secret_ref: The secret reference / where to find the secret, URL.
        :param pulumi.Input[str] name: The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
        """
        ContainerV1SecretRefArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            secret_ref=secret_ref,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             secret_ref: pulumi.Input[str],
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("secret_ref", secret_ref)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> pulumi.Input[str]:
        """
        The secret reference / where to find the secret, URL.
        """
        return pulumi.get(self, "secret_ref")

    @secret_ref.setter
    def secret_ref(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_ref", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class OrderV1MetaArgs:
    def __init__(__self__, *,
                 algorithm: pulumi.Input[str],
                 bit_length: pulumi.Input[int],
                 expiration: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_content_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] algorithm: Algorithm to use for key generation.
        :param pulumi.Input[int] bit_length: Bit lenght of key to be generated.
        :param pulumi.Input[str] expiration: This is a UTC timestamp in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ. If set, the secret will not be available after this time.
        :param pulumi.Input[str] mode: The mode to use for key generation.
        :param pulumi.Input[str] name: The name of the secret set by the user.
        :param pulumi.Input[str] payload_content_type: The media type for the content of the secrets payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        """
        OrderV1MetaArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            algorithm=algorithm,
            bit_length=bit_length,
            expiration=expiration,
            mode=mode,
            name=name,
            payload_content_type=payload_content_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             algorithm: pulumi.Input[str],
             bit_length: pulumi.Input[int],
             expiration: Optional[pulumi.Input[str]] = None,
             mode: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             payload_content_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("algorithm", algorithm)
        _setter("bit_length", bit_length)
        if expiration is not None:
            _setter("expiration", expiration)
        if mode is not None:
            _setter("mode", mode)
        if name is not None:
            _setter("name", name)
        if payload_content_type is not None:
            _setter("payload_content_type", payload_content_type)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Input[str]:
        """
        Algorithm to use for key generation.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="bitLength")
    def bit_length(self) -> pulumi.Input[int]:
        """
        Bit lenght of key to be generated.
        """
        return pulumi.get(self, "bit_length")

    @bit_length.setter
    def bit_length(self, value: pulumi.Input[int]):
        pulumi.set(self, "bit_length", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        """
        This is a UTC timestamp in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ. If set, the secret will not be available after this time.
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode to use for key generation.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the secret set by the user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="payloadContentType")
    def payload_content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The media type for the content of the secrets payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        """
        return pulumi.get(self, "payload_content_type")

    @payload_content_type.setter
    def payload_content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_content_type", value)


@pulumi.input_type
class SecretV1AclArgs:
    def __init__(__self__, *,
                 read: Optional[pulumi.Input['SecretV1AclReadArgs']] = None):
        SecretV1AclArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            read=read,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             read: Optional[pulumi.Input['SecretV1AclReadArgs']] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if read is not None:
            _setter("read", read)

    @property
    @pulumi.getter
    def read(self) -> Optional[pulumi.Input['SecretV1AclReadArgs']]:
        return pulumi.get(self, "read")

    @read.setter
    def read(self, value: Optional[pulumi.Input['SecretV1AclReadArgs']]):
        pulumi.set(self, "read", value)


@pulumi.input_type
class SecretV1AclReadArgs:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 project_access: Optional[pulumi.Input[bool]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] created_at: The date the secret ACL was created.
        :param pulumi.Input[bool] project_access: Whether the secret is accessible project wide.
               Defaults to `true`.
        :param pulumi.Input[str] updated_at: The date the secret ACL was last updated.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: The list of user IDs, which are allowed to access the
               secret, when `project_access` is set to `false`.
        """
        SecretV1AclReadArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            created_at=created_at,
            project_access=project_access,
            updated_at=updated_at,
            users=users,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             created_at: Optional[pulumi.Input[str]] = None,
             project_access: Optional[pulumi.Input[bool]] = None,
             updated_at: Optional[pulumi.Input[str]] = None,
             users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if created_at is not None:
            _setter("created_at", created_at)
        if project_access is not None:
            _setter("project_access", project_access)
        if updated_at is not None:
            _setter("updated_at", updated_at)
        if users is not None:
            _setter("users", users)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date the secret ACL was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="projectAccess")
    def project_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the secret is accessible project wide.
        Defaults to `true`.
        """
        return pulumi.get(self, "project_access")

    @project_access.setter
    def project_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "project_access", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date the secret ACL was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of user IDs, which are allowed to access the
        secret, when `project_access` is set to `false`.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users", value)


