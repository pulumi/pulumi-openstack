# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
    'GetAddressScopeResult',
    'AwaitableGetAddressScopeResult',
    'get_address_scope',
    'get_address_scope_output',
]

@pulumi.output_type
class GetAddressScopeResult:
    """
    A collection of values returned by getAddressScope.
    """
    def __init__(__self__, id=None, ip_version=None, name=None, project_id=None, region=None, shared=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_version and not isinstance(ip_version, int):
            raise TypeError("Expected argument 'ip_version' to be a int")
        pulumi.set(__self__, "ip_version", ip_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        pulumi.set(__self__, "shared", shared)

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[_builtins.int]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "ip_version")

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter
    def region(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "region")

    @_builtins.property
    @pulumi.getter
    def shared(self) -> Optional[_builtins.bool]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "shared")


class AwaitableGetAddressScopeResult(GetAddressScopeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAddressScopeResult(
            id=self.id,
            ip_version=self.ip_version,
            name=self.name,
            project_id=self.project_id,
            region=self.region,
            shared=self.shared)


def get_address_scope(ip_version: Optional[_builtins.int] = None,
                      name: Optional[_builtins.str] = None,
                      project_id: Optional[_builtins.str] = None,
                      region: Optional[_builtins.str] = None,
                      shared: Optional[_builtins.bool] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAddressScopeResult:
    """
    Use this data source to get the ID of an available OpenStack address-scope.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    public_addressscope = openstack.networking.get_address_scope(name="public_addressscope",
        shared=True,
        ip_version=4)
    ```


    :param _builtins.int ip_version: IP version.
    :param _builtins.str name: Name of the address-scope.
    :param _builtins.str project_id: The owner of the address-scope.
    :param _builtins.str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve address-scopes. If omitted, the
           `region` argument of the provider is used.
    :param _builtins.bool shared: Indicates whether this address-scope is shared across
           all projects.
    """
    __args__ = dict()
    __args__['ipVersion'] = ip_version
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['shared'] = shared
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getAddressScope:getAddressScope', __args__, opts=opts, typ=GetAddressScopeResult).value

    return AwaitableGetAddressScopeResult(
        id=pulumi.get(__ret__, 'id'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        shared=pulumi.get(__ret__, 'shared'))
def get_address_scope_output(ip_version: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                             name: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                             project_id: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                             region: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                             shared: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAddressScopeResult]:
    """
    Use this data source to get the ID of an available OpenStack address-scope.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    public_addressscope = openstack.networking.get_address_scope(name="public_addressscope",
        shared=True,
        ip_version=4)
    ```


    :param _builtins.int ip_version: IP version.
    :param _builtins.str name: Name of the address-scope.
    :param _builtins.str project_id: The owner of the address-scope.
    :param _builtins.str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve address-scopes. If omitted, the
           `region` argument of the provider is used.
    :param _builtins.bool shared: Indicates whether this address-scope is shared across
           all projects.
    """
    __args__ = dict()
    __args__['ipVersion'] = ip_version
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['shared'] = shared
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:networking/getAddressScope:getAddressScope', __args__, opts=opts, typ=GetAddressScopeResult)
    return __ret__.apply(lambda __response__: GetAddressScopeResult(
        id=pulumi.get(__response__, 'id'),
        ip_version=pulumi.get(__response__, 'ip_version'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        shared=pulumi.get(__response__, 'shared')))
