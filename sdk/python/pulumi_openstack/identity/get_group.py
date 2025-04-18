# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
    'get_group_output',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, description=None, domain_id=None, id=None, name=None, region=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
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
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        A description of the group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            description=self.description,
            domain_id=self.domain_id,
            id=self.id,
            name=self.name,
            region=self.region)


def get_group(domain_id: Optional[builtins.str] = None,
              name: Optional[builtins.str] = None,
              region: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    Use this data source to get the ID of an OpenStack group.

    > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
    this resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    admins = openstack.identity.get_group(name="admins")
    ```


    :param builtins.str domain_id: The domain the group belongs to.
    :param builtins.str name: The name of the group.
    :param builtins.str region: The region in which to obtain the V3 Keystone client.
           If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:identity/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        description=pulumi.get(__ret__, 'description'),
        domain_id=pulumi.get(__ret__, 'domain_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'))
def get_group_output(domain_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     name: Optional[pulumi.Input[builtins.str]] = None,
                     region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGroupResult]:
    """
    Use this data source to get the ID of an OpenStack group.

    > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
    this resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    admins = openstack.identity.get_group(name="admins")
    ```


    :param builtins.str domain_id: The domain the group belongs to.
    :param builtins.str name: The name of the group.
    :param builtins.str region: The region in which to obtain the V3 Keystone client.
           If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:identity/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult)
    return __ret__.apply(lambda __response__: GetGroupResult(
        description=pulumi.get(__response__, 'description'),
        domain_id=pulumi.get(__response__, 'domain_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region')))
