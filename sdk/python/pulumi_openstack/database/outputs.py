# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'ConfigurationConfiguration',
    'ConfigurationDatastore',
    'InstanceDatabase',
    'InstanceDatastore',
    'InstanceNetwork',
    'InstanceUser',
]

@pulumi.output_type
class ConfigurationConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "stringType":
            suggest = "string_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ConfigurationConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ConfigurationConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ConfigurationConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: str,
                 value: str,
                 string_type: Optional[bool] = None):
        """
        :param str name: Configuration parameter name. Changing this creates a new resource.
        :param str value: Configuration parameter value. Changing this creates a new resource.
        :param bool string_type: Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if string_type is not None:
            pulumi.set(__self__, "string_type", string_type)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Configuration parameter name. Changing this creates a new resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Configuration parameter value. Changing this creates a new resource.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="stringType")
    def string_type(self) -> Optional[bool]:
        """
        Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        """
        return pulumi.get(self, "string_type")


@pulumi.output_type
class ConfigurationDatastore(dict):
    def __init__(__self__, *,
                 type: str,
                 version: str):
        """
        :param str type: Database engine type to be used with this configuration. Changing this creates a new resource.
        :param str version: Version of database engine type to be used with this configuration. Changing this creates a new resource.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Database engine type to be used with this configuration. Changing this creates a new resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version of database engine type to be used with this configuration. Changing this creates a new resource.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class InstanceDatabase(dict):
    def __init__(__self__, *,
                 name: str,
                 charset: Optional[str] = None,
                 collate: Optional[str] = None):
        """
        :param str name: Database to be created on new instance. Changing this creates a
               new instance.
        :param str charset: Database character set. Changing this creates a
               new instance.
        :param str collate: Database collation. Changing this creates a new instance.
        """
        pulumi.set(__self__, "name", name)
        if charset is not None:
            pulumi.set(__self__, "charset", charset)
        if collate is not None:
            pulumi.set(__self__, "collate", collate)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Database to be created on new instance. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def charset(self) -> Optional[str]:
        """
        Database character set. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "charset")

    @property
    @pulumi.getter
    def collate(self) -> Optional[str]:
        """
        Database collation. Changing this creates a new instance.
        """
        return pulumi.get(self, "collate")


@pulumi.output_type
class InstanceDatastore(dict):
    def __init__(__self__, *,
                 type: str,
                 version: str):
        """
        :param str type: Database engine type to be used in new instance. Changing this
               creates a new instance.
        :param str version: Version of database engine type to be used in new instance.
               Changing this creates a new instance.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Database engine type to be used in new instance. Changing this
        creates a new instance.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version of database engine type to be used in new instance.
        Changing this creates a new instance.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class InstanceNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fixedIpV4":
            suggest = "fixed_ip_v4"
        elif key == "fixedIpV6":
            suggest = "fixed_ip_v6"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 fixed_ip_v4: Optional[str] = None,
                 fixed_ip_v6: Optional[str] = None,
                 port: Optional[str] = None,
                 uuid: Optional[str] = None):
        """
        :param str fixed_ip_v4: Specifies a fixed IPv4 address to be used on this
               network. Changing this creates a new instance.
        :param str fixed_ip_v6: Specifies a fixed IPv6 address to be used on this
               network. Changing this creates a new instance.
        :param str port: The port UUID of a
               network to attach to the instance. Changing this creates a new instance.
        :param str uuid: The network UUID to
               attach to the instance. Changing this creates a new instance.
        """
        if fixed_ip_v4 is not None:
            pulumi.set(__self__, "fixed_ip_v4", fixed_ip_v4)
        if fixed_ip_v6 is not None:
            pulumi.set(__self__, "fixed_ip_v6", fixed_ip_v6)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="fixedIpV4")
    def fixed_ip_v4(self) -> Optional[str]:
        """
        Specifies a fixed IPv4 address to be used on this
        network. Changing this creates a new instance.
        """
        return pulumi.get(self, "fixed_ip_v4")

    @property
    @pulumi.getter(name="fixedIpV6")
    def fixed_ip_v6(self) -> Optional[str]:
        """
        Specifies a fixed IPv6 address to be used on this
        network. Changing this creates a new instance.
        """
        return pulumi.get(self, "fixed_ip_v6")

    @property
    @pulumi.getter
    def port(self) -> Optional[str]:
        """
        The port UUID of a
        network to attach to the instance. Changing this creates a new instance.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[str]:
        """
        The network UUID to
        attach to the instance. Changing this creates a new instance.
        """
        return pulumi.get(self, "uuid")


@pulumi.output_type
class InstanceUser(dict):
    def __init__(__self__, *,
                 name: str,
                 databases: Optional[Sequence[str]] = None,
                 host: Optional[str] = None,
                 password: Optional[str] = None):
        """
        :param str name: Username to be created on new instance. Changing this creates a
               new instance.
        :param Sequence[str] databases: A list of databases that user will have access to. If not specified,
               user has access to all databases on th einstance. Changing this creates a new instance.
        :param str host: An ip address or % sign indicating what ip addresses can connect with
               this user credentials. Changing this creates a new instance.
        :param str password: User's password. Changing this creates a
               new instance.
        """
        pulumi.set(__self__, "name", name)
        if databases is not None:
            pulumi.set(__self__, "databases", databases)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if password is not None:
            pulumi.set(__self__, "password", password)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Username to be created on new instance. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def databases(self) -> Optional[Sequence[str]]:
        """
        A list of databases that user will have access to. If not specified,
        user has access to all databases on th einstance. Changing this creates a new instance.
        """
        return pulumi.get(self, "databases")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        """
        An ip address or % sign indicating what ip addresses can connect with
        this user credentials. Changing this creates a new instance.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        """
        User's password. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "password")


