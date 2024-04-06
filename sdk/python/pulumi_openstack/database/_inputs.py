# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ConfigurationConfigurationArgs',
    'ConfigurationDatastoreArgs',
    'InstanceDatabaseArgs',
    'InstanceDatastoreArgs',
    'InstanceNetworkArgs',
    'InstanceUserArgs',
]

@pulumi.input_type
class ConfigurationConfigurationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str],
                 string_type: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] name: Configuration parameter name. Changing this creates a new resource.
        :param pulumi.Input[str] value: Configuration parameter value. Changing this creates a new resource.
        :param pulumi.Input[bool] string_type: Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if string_type is not None:
            pulumi.set(__self__, "string_type", string_type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Configuration parameter name. Changing this creates a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Configuration parameter value. Changing this creates a new resource.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="stringType")
    def string_type(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        """
        return pulumi.get(self, "string_type")

    @string_type.setter
    def string_type(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "string_type", value)


@pulumi.input_type
class ConfigurationDatastoreArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 version: pulumi.Input[str]):
        """
        :param pulumi.Input[str] type: Database engine type to be used with this configuration. Changing this creates a new resource.
        :param pulumi.Input[str] version: Version of database engine type to be used with this configuration. Changing this creates a new resource.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Database engine type to be used with this configuration. Changing this creates a new resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Version of database engine type to be used with this configuration. Changing this creates a new resource.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class InstanceDatabaseArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 charset: Optional[pulumi.Input[str]] = None,
                 collate: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Database to be created on new instance. Changing this creates a
               new instance.
        :param pulumi.Input[str] charset: Database character set. Changing this creates a
               new instance.
        :param pulumi.Input[str] collate: Database collation. Changing this creates a new instance.
        """
        pulumi.set(__self__, "name", name)
        if charset is not None:
            pulumi.set(__self__, "charset", charset)
        if collate is not None:
            pulumi.set(__self__, "collate", collate)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Database to be created on new instance. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def charset(self) -> Optional[pulumi.Input[str]]:
        """
        Database character set. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "charset")

    @charset.setter
    def charset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charset", value)

    @property
    @pulumi.getter
    def collate(self) -> Optional[pulumi.Input[str]]:
        """
        Database collation. Changing this creates a new instance.
        """
        return pulumi.get(self, "collate")

    @collate.setter
    def collate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collate", value)


@pulumi.input_type
class InstanceDatastoreArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 version: pulumi.Input[str]):
        """
        :param pulumi.Input[str] type: Database engine type to be used in new instance. Changing this
               creates a new instance.
        :param pulumi.Input[str] version: Version of database engine type to be used in new instance.
               Changing this creates a new instance.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Database engine type to be used in new instance. Changing this
        creates a new instance.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Version of database engine type to be used in new instance.
        Changing this creates a new instance.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class InstanceNetworkArgs:
    def __init__(__self__, *,
                 fixed_ip_v4: Optional[pulumi.Input[str]] = None,
                 fixed_ip_v6: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] fixed_ip_v4: Specifies a fixed IPv4 address to be used on this
               network. Changing this creates a new instance.
        :param pulumi.Input[str] fixed_ip_v6: Specifies a fixed IPv6 address to be used on this
               network. Changing this creates a new instance.
        :param pulumi.Input[str] port: The port UUID of a
               network to attach to the instance. Changing this creates a new instance.
        :param pulumi.Input[str] uuid: The network UUID to
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
    def fixed_ip_v4(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a fixed IPv4 address to be used on this
        network. Changing this creates a new instance.
        """
        return pulumi.get(self, "fixed_ip_v4")

    @fixed_ip_v4.setter
    def fixed_ip_v4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip_v4", value)

    @property
    @pulumi.getter(name="fixedIpV6")
    def fixed_ip_v6(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a fixed IPv6 address to be used on this
        network. Changing this creates a new instance.
        """
        return pulumi.get(self, "fixed_ip_v6")

    @fixed_ip_v6.setter
    def fixed_ip_v6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fixed_ip_v6", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        The port UUID of a
        network to attach to the instance. Changing this creates a new instance.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        The network UUID to
        attach to the instance. Changing this creates a new instance.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


@pulumi.input_type
class InstanceUserArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 databases: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Username to be created on new instance. Changing this creates a
               new instance.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] databases: A list of databases that user will have access to. If not specified,
               user has access to all databases on th einstance. Changing this creates a new instance.
        :param pulumi.Input[str] host: An ip address or % sign indicating what ip addresses can connect with
               this user credentials. Changing this creates a new instance.
        :param pulumi.Input[str] password: User's password. Changing this creates a
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
    def name(self) -> pulumi.Input[str]:
        """
        Username to be created on new instance. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def databases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of databases that user will have access to. If not specified,
        user has access to all databases on th einstance. Changing this creates a new instance.
        """
        return pulumi.get(self, "databases")

    @databases.setter
    def databases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "databases", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        An ip address or % sign indicating what ip addresses can connect with
        this user credentials. Changing this creates a new instance.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        User's password. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)


