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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, default_project_id=None, description=None, domain_id=None, enabled=None, id=None, idp_id=None, name=None, password_expires_at=None, protocol_id=None, region=None, unique_id=None):
        if default_project_id and not isinstance(default_project_id, str):
            raise TypeError("Expected argument 'default_project_id' to be a str")
        pulumi.set(__self__, "default_project_id", default_project_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        pulumi.set(__self__, "domain_id", domain_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idp_id and not isinstance(idp_id, str):
            raise TypeError("Expected argument 'idp_id' to be a str")
        pulumi.set(__self__, "idp_id", idp_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if password_expires_at and not isinstance(password_expires_at, str):
            raise TypeError("Expected argument 'password_expires_at' to be a str")
        pulumi.set(__self__, "password_expires_at", password_expires_at)
        if protocol_id and not isinstance(protocol_id, str):
            raise TypeError("Expected argument 'protocol_id' to be a str")
        pulumi.set(__self__, "protocol_id", protocol_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if unique_id and not isinstance(unique_id, str):
            raise TypeError("Expected argument 'unique_id' to be a str")
        pulumi.set(__self__, "unique_id", unique_id)

    @property
    @pulumi.getter(name="defaultProjectId")
    def default_project_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "default_project_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the user.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "idp_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="passwordExpiresAt")
    def password_expires_at(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "password_expires_at")

    @property
    @pulumi.getter(name="protocolId")
    def protocol_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "protocol_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region the user is located in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "unique_id")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            default_project_id=self.default_project_id,
            description=self.description,
            domain_id=self.domain_id,
            enabled=self.enabled,
            id=self.id,
            idp_id=self.idp_id,
            name=self.name,
            password_expires_at=self.password_expires_at,
            protocol_id=self.protocol_id,
            region=self.region,
            unique_id=self.unique_id)


def get_user(domain_id: Optional[str] = None,
             enabled: Optional[bool] = None,
             idp_id: Optional[str] = None,
             name: Optional[str] = None,
             password_expires_at: Optional[str] = None,
             protocol_id: Optional[str] = None,
             region: Optional[str] = None,
             unique_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Use this data source to get the ID of an OpenStack user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    user1 = openstack.identity.get_user(name="user_1")
    ```


    :param str domain_id: The domain this user belongs to.
    :param bool enabled: Whether the user is enabled or disabled. Valid
           values are `true` and `false`.
    :param str idp_id: The identity provider ID of the user.
    :param str name: The name of the user.
    :param str password_expires_at: Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
    :param str protocol_id: The protocol ID of the user.
    :param str region: The region the user is located in.
    :param str unique_id: The unique ID of the user.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['enabled'] = enabled
    __args__['idpId'] = idp_id
    __args__['name'] = name
    __args__['passwordExpiresAt'] = password_expires_at
    __args__['protocolId'] = protocol_id
    __args__['region'] = region
    __args__['uniqueId'] = unique_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:identity/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        default_project_id=pulumi.get(__ret__, 'default_project_id'),
        description=pulumi.get(__ret__, 'description'),
        domain_id=pulumi.get(__ret__, 'domain_id'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        idp_id=pulumi.get(__ret__, 'idp_id'),
        name=pulumi.get(__ret__, 'name'),
        password_expires_at=pulumi.get(__ret__, 'password_expires_at'),
        protocol_id=pulumi.get(__ret__, 'protocol_id'),
        region=pulumi.get(__ret__, 'region'),
        unique_id=pulumi.get(__ret__, 'unique_id'))


@_utilities.lift_output_func(get_user)
def get_user_output(domain_id: Optional[pulumi.Input[Optional[str]]] = None,
                    enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                    idp_id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    password_expires_at: Optional[pulumi.Input[Optional[str]]] = None,
                    protocol_id: Optional[pulumi.Input[Optional[str]]] = None,
                    region: Optional[pulumi.Input[Optional[str]]] = None,
                    unique_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Use this data source to get the ID of an OpenStack user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    user1 = openstack.identity.get_user(name="user_1")
    ```


    :param str domain_id: The domain this user belongs to.
    :param bool enabled: Whether the user is enabled or disabled. Valid
           values are `true` and `false`.
    :param str idp_id: The identity provider ID of the user.
    :param str name: The name of the user.
    :param str password_expires_at: Query for expired passwords. See the [OpenStack API docs](https://developer.openstack.org/api-ref/identity/v3/#list-users) for more information on the query format.
    :param str protocol_id: The protocol ID of the user.
    :param str region: The region the user is located in.
    :param str unique_id: The unique ID of the user.
    """
    ...
