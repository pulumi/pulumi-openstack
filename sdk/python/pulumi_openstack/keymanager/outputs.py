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

__all__ = [
    'ContainerV1Acl',
    'ContainerV1AclRead',
    'ContainerV1Consumer',
    'ContainerV1SecretRef',
    'OrderV1Meta',
    'SecretV1Acl',
    'SecretV1AclRead',
    'GetContainerAclResult',
    'GetContainerAclReadResult',
    'GetContainerConsumerResult',
    'GetContainerSecretRefResult',
    'GetSecretAclResult',
    'GetSecretAclReadResult',
]

@pulumi.output_type
class ContainerV1Acl(dict):
    def __init__(__self__, *,
                 read: Optional['outputs.ContainerV1AclRead'] = None):
        if read is not None:
            pulumi.set(__self__, "read", read)

    @property
    @pulumi.getter
    def read(self) -> Optional['outputs.ContainerV1AclRead']:
        return pulumi.get(self, "read")


@pulumi.output_type
class ContainerV1AclRead(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "projectAccess":
            suggest = "project_access"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerV1AclRead. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerV1AclRead.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerV1AclRead.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 project_access: Optional[bool] = None,
                 updated_at: Optional[str] = None,
                 users: Optional[Sequence[str]] = None):
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if project_access is not None:
            pulumi.set(__self__, "project_access", project_access)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="projectAccess")
    def project_access(self) -> Optional[bool]:
        return pulumi.get(self, "project_access")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[str]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def users(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "users")


@pulumi.output_type
class ContainerV1Consumer(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 url: Optional[str] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        return pulumi.get(self, "url")


@pulumi.output_type
class ContainerV1SecretRef(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "secretRef":
            suggest = "secret_ref"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerV1SecretRef. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerV1SecretRef.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerV1SecretRef.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 secret_ref: str,
                 name: Optional[str] = None):
        pulumi.set(__self__, "secret_ref", secret_ref)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> str:
        return pulumi.get(self, "secret_ref")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


@pulumi.output_type
class OrderV1Meta(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "bitLength":
            suggest = "bit_length"
        elif key == "payloadContentType":
            suggest = "payload_content_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OrderV1Meta. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OrderV1Meta.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OrderV1Meta.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 algorithm: str,
                 bit_length: int,
                 expiration: Optional[str] = None,
                 mode: Optional[str] = None,
                 name: Optional[str] = None,
                 payload_content_type: Optional[str] = None):
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "bit_length", bit_length)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if payload_content_type is not None:
            pulumi.set(__self__, "payload_content_type", payload_content_type)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="bitLength")
    def bit_length(self) -> int:
        return pulumi.get(self, "bit_length")

    @property
    @pulumi.getter
    def expiration(self) -> Optional[str]:
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter
    def mode(self) -> Optional[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="payloadContentType")
    def payload_content_type(self) -> Optional[str]:
        return pulumi.get(self, "payload_content_type")


@pulumi.output_type
class SecretV1Acl(dict):
    def __init__(__self__, *,
                 read: Optional['outputs.SecretV1AclRead'] = None):
        if read is not None:
            pulumi.set(__self__, "read", read)

    @property
    @pulumi.getter
    def read(self) -> Optional['outputs.SecretV1AclRead']:
        return pulumi.get(self, "read")


@pulumi.output_type
class SecretV1AclRead(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "projectAccess":
            suggest = "project_access"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecretV1AclRead. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecretV1AclRead.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecretV1AclRead.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 project_access: Optional[bool] = None,
                 updated_at: Optional[str] = None,
                 users: Optional[Sequence[str]] = None):
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if project_access is not None:
            pulumi.set(__self__, "project_access", project_access)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="projectAccess")
    def project_access(self) -> Optional[bool]:
        return pulumi.get(self, "project_access")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[str]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def users(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "users")


@pulumi.output_type
class GetContainerAclResult(dict):
    def __init__(__self__, *,
                 read: 'outputs.GetContainerAclReadResult'):
        pulumi.set(__self__, "read", read)

    @property
    @pulumi.getter
    def read(self) -> 'outputs.GetContainerAclReadResult':
        return pulumi.get(self, "read")


@pulumi.output_type
class GetContainerAclReadResult(dict):
    def __init__(__self__, *,
                 created_at: str,
                 updated_at: str,
                 project_access: Optional[bool] = None,
                 users: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "updated_at", updated_at)
        if project_access is not None:
            pulumi.set(__self__, "project_access", project_access)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="projectAccess")
    def project_access(self) -> Optional[bool]:
        return pulumi.get(self, "project_access")

    @property
    @pulumi.getter
    def users(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "users")


@pulumi.output_type
class GetContainerConsumerResult(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 url: Optional[str] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        return pulumi.get(self, "url")


@pulumi.output_type
class GetContainerSecretRefResult(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 secret_ref: Optional[str] = None):
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret_ref is not None:
            pulumi.set(__self__, "secret_ref", secret_ref)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="secretRef")
    def secret_ref(self) -> Optional[str]:
        return pulumi.get(self, "secret_ref")


@pulumi.output_type
class GetSecretAclResult(dict):
    def __init__(__self__, *,
                 read: 'outputs.GetSecretAclReadResult'):
        pulumi.set(__self__, "read", read)

    @property
    @pulumi.getter
    def read(self) -> 'outputs.GetSecretAclReadResult':
        return pulumi.get(self, "read")


@pulumi.output_type
class GetSecretAclReadResult(dict):
    def __init__(__self__, *,
                 created_at: str,
                 updated_at: str,
                 project_access: Optional[bool] = None,
                 users: Optional[Sequence[str]] = None):
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "updated_at", updated_at)
        if project_access is not None:
            pulumi.set(__self__, "project_access", project_access)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="projectAccess")
    def project_access(self) -> Optional[bool]:
        return pulumi.get(self, "project_access")

    @property
    @pulumi.getter
    def users(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "users")


