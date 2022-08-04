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
    'GetRoleResult',
    'AwaitableGetRoleResult',
    'get_role',
    'get_role_output',
]

@pulumi.output_type
class GetRoleResult:
    """
    A collection of values returned by getRole.
    """
    def __init__(__self__, domain_id=None, id=None, name=None, region=None):
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        pulumi.set(__self__, "domain_id", domain_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")


class AwaitableGetRoleResult(GetRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleResult(
            domain_id=self.domain_id,
            id=self.id,
            name=self.name,
            region=self.region)


def get_role(domain_id: Optional[str] = None,
             name: Optional[str] = None,
             region: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleResult:
    """
    Use this data source to get the ID of an OpenStack role.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    admin = openstack.identity.get_role(name="admin")
    ```


    :param str domain_id: The domain the role belongs to.
    :param str name: The name of the role.
    :param str region: The region in which to obtain the V3 Keystone client.
           If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:identity/getRole:getRole', __args__, opts=opts, typ=GetRoleResult).value

    return AwaitableGetRoleResult(
        domain_id=__ret__.domain_id,
        id=__ret__.id,
        name=__ret__.name,
        region=__ret__.region)


@_utilities.lift_output_func(get_role)
def get_role_output(domain_id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[str]] = None,
                    region: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoleResult]:
    """
    Use this data source to get the ID of an OpenStack role.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    admin = openstack.identity.get_role(name="admin")
    ```


    :param str domain_id: The domain the role belongs to.
    :param str name: The name of the role.
    :param str region: The region in which to obtain the V3 Keystone client.
           If omitted, the `region` argument of the provider is used.
    """
    ...
