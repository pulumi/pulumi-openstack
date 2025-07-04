# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'ConfigurationConfigurationArgs',
    'ConfigurationConfigurationArgsDict',
    'ConfigurationDatastoreArgs',
    'ConfigurationDatastoreArgsDict',
    'InstanceDatabaseArgs',
    'InstanceDatabaseArgsDict',
    'InstanceDatastoreArgs',
    'InstanceDatastoreArgsDict',
    'InstanceNetworkArgs',
    'InstanceNetworkArgsDict',
    'InstanceUserArgs',
    'InstanceUserArgsDict',
]

MYPY = False

if not MYPY:
    class ConfigurationConfigurationArgsDict(TypedDict):
        name: pulumi.Input[builtins.str]
        """
        Configuration parameter name. Changing this creates a new resource.
        """
        value: pulumi.Input[builtins.str]
        """
        Configuration parameter value. Changing this creates a new resource.
        """
        string_type: NotRequired[pulumi.Input[builtins.bool]]
        """
        Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        """
elif False:
    ConfigurationConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ConfigurationConfigurationArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 value: pulumi.Input[builtins.str],
                 string_type: Optional[pulumi.Input[builtins.bool]] = None):
        """
        :param pulumi.Input[builtins.str] name: Configuration parameter name. Changing this creates a new resource.
        :param pulumi.Input[builtins.str] value: Configuration parameter value. Changing this creates a new resource.
        :param pulumi.Input[builtins.bool] string_type: Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)
        if string_type is not None:
            pulumi.set(__self__, "string_type", string_type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Configuration parameter name. Changing this creates a new resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[builtins.str]:
        """
        Configuration parameter value. Changing this creates a new resource.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="stringType")
    def string_type(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
        """
        return pulumi.get(self, "string_type")

    @string_type.setter
    def string_type(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "string_type", value)


if not MYPY:
    class ConfigurationDatastoreArgsDict(TypedDict):
        type: pulumi.Input[builtins.str]
        """
        Database engine type to be used with this configuration. Changing this creates a new resource.
        """
        version: pulumi.Input[builtins.str]
        """
        Version of database engine type to be used with this configuration. Changing this creates a new resource.
        """
elif False:
    ConfigurationDatastoreArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ConfigurationDatastoreArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[builtins.str],
                 version: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] type: Database engine type to be used with this configuration. Changing this creates a new resource.
        :param pulumi.Input[builtins.str] version: Version of database engine type to be used with this configuration. Changing this creates a new resource.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        """
        Database engine type to be used with this configuration. Changing this creates a new resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[builtins.str]:
        """
        Version of database engine type to be used with this configuration. Changing this creates a new resource.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "version", value)


if not MYPY:
    class InstanceDatabaseArgsDict(TypedDict):
        name: pulumi.Input[builtins.str]
        """
        Database to be created on new instance. Changing this creates a
        new instance.
        """
        charset: NotRequired[pulumi.Input[builtins.str]]
        """
        Database character set. Changing this creates a
        new instance.
        """
        collate: NotRequired[pulumi.Input[builtins.str]]
        """
        Database collation. Changing this creates a new instance.
        """
elif False:
    InstanceDatabaseArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InstanceDatabaseArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 charset: Optional[pulumi.Input[builtins.str]] = None,
                 collate: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] name: Database to be created on new instance. Changing this creates a
               new instance.
        :param pulumi.Input[builtins.str] charset: Database character set. Changing this creates a
               new instance.
        :param pulumi.Input[builtins.str] collate: Database collation. Changing this creates a new instance.
        """
        pulumi.set(__self__, "name", name)
        if charset is not None:
            pulumi.set(__self__, "charset", charset)
        if collate is not None:
            pulumi.set(__self__, "collate", collate)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Database to be created on new instance. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def charset(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Database character set. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "charset")

    @charset.setter
    def charset(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "charset", value)

    @property
    @pulumi.getter
    def collate(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Database collation. Changing this creates a new instance.
        """
        return pulumi.get(self, "collate")

    @collate.setter
    def collate(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "collate", value)


if not MYPY:
    class InstanceDatastoreArgsDict(TypedDict):
        type: pulumi.Input[builtins.str]
        """
        Database engine type to be used in new instance. Changing this
        creates a new instance.
        """
        version: pulumi.Input[builtins.str]
        """
        Version of database engine type to be used in new instance.
        Changing this creates a new instance.
        """
elif False:
    InstanceDatastoreArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InstanceDatastoreArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[builtins.str],
                 version: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] type: Database engine type to be used in new instance. Changing this
               creates a new instance.
        :param pulumi.Input[builtins.str] version: Version of database engine type to be used in new instance.
               Changing this creates a new instance.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        """
        Database engine type to be used in new instance. Changing this
        creates a new instance.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[builtins.str]:
        """
        Version of database engine type to be used in new instance.
        Changing this creates a new instance.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "version", value)


if not MYPY:
    class InstanceNetworkArgsDict(TypedDict):
        fixed_ip_v4: NotRequired[pulumi.Input[builtins.str]]
        """
        Specifies a fixed IPv4 address to be used on this
        network. Changing this creates a new instance.
        """
        fixed_ip_v6: NotRequired[pulumi.Input[builtins.str]]
        """
        Specifies a fixed IPv6 address to be used on this
        network. Changing this creates a new instance.
        """
        port: NotRequired[pulumi.Input[builtins.str]]
        """
        The port UUID of a
        network to attach to the instance. Changing this creates a new instance.
        """
        uuid: NotRequired[pulumi.Input[builtins.str]]
        """
        The network UUID to
        attach to the instance. Changing this creates a new instance.
        """
elif False:
    InstanceNetworkArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InstanceNetworkArgs:
    def __init__(__self__, *,
                 fixed_ip_v4: Optional[pulumi.Input[builtins.str]] = None,
                 fixed_ip_v6: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.str]] = None,
                 uuid: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] fixed_ip_v4: Specifies a fixed IPv4 address to be used on this
               network. Changing this creates a new instance.
        :param pulumi.Input[builtins.str] fixed_ip_v6: Specifies a fixed IPv6 address to be used on this
               network. Changing this creates a new instance.
        :param pulumi.Input[builtins.str] port: The port UUID of a
               network to attach to the instance. Changing this creates a new instance.
        :param pulumi.Input[builtins.str] uuid: The network UUID to
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
    def fixed_ip_v4(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies a fixed IPv4 address to be used on this
        network. Changing this creates a new instance.
        """
        return pulumi.get(self, "fixed_ip_v4")

    @fixed_ip_v4.setter
    def fixed_ip_v4(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "fixed_ip_v4", value)

    @property
    @pulumi.getter(name="fixedIpV6")
    def fixed_ip_v6(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies a fixed IPv6 address to be used on this
        network. Changing this creates a new instance.
        """
        return pulumi.get(self, "fixed_ip_v6")

    @fixed_ip_v6.setter
    def fixed_ip_v6(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "fixed_ip_v6", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The port UUID of a
        network to attach to the instance. Changing this creates a new instance.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The network UUID to
        attach to the instance. Changing this creates a new instance.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "uuid", value)


if not MYPY:
    class InstanceUserArgsDict(TypedDict):
        name: pulumi.Input[builtins.str]
        """
        Username to be created on new instance. Changing this creates a
        new instance.
        """
        databases: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        A list of databases that user will have access to. If not specified,
        user has access to all databases on th einstance. Changing this creates a new instance.
        """
        host: NotRequired[pulumi.Input[builtins.str]]
        """
        An ip address or % sign indicating what ip addresses can connect with
        this user credentials. Changing this creates a new instance.
        """
        password: NotRequired[pulumi.Input[builtins.str]]
        """
        User's password. Changing this creates a
        new instance.
        """
elif False:
    InstanceUserArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class InstanceUserArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 databases: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 host: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] name: Username to be created on new instance. Changing this creates a
               new instance.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] databases: A list of databases that user will have access to. If not specified,
               user has access to all databases on th einstance. Changing this creates a new instance.
        :param pulumi.Input[builtins.str] host: An ip address or % sign indicating what ip addresses can connect with
               this user credentials. Changing this creates a new instance.
        :param pulumi.Input[builtins.str] password: User's password. Changing this creates a
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
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Username to be created on new instance. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def databases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of databases that user will have access to. If not specified,
        user has access to all databases on th einstance. Changing this creates a new instance.
        """
        return pulumi.get(self, "databases")

    @databases.setter
    def databases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "databases", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An ip address or % sign indicating what ip addresses can connect with
        this user credentials. Changing this creates a new instance.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User's password. Changing this creates a
        new instance.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password", value)


