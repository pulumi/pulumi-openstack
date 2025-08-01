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
from . import outputs

__all__ = [
    'GetServergroupV2Result',
    'AwaitableGetServergroupV2Result',
    'get_servergroup_v2',
    'get_servergroup_v2_output',
]

@pulumi.output_type
class GetServergroupV2Result:
    """
    A collection of values returned by getServergroupV2.
    """
    def __init__(__self__, id=None, members=None, metadata=None, name=None, policies=None, project_id=None, region=None, rules=None, user_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def members(self) -> Sequence[_builtins.str]:
        """
        The instances that are part of this server group.
        """
        return pulumi.get(self, "members")

    @_builtins.property
    @pulumi.getter
    def metadata(self) -> Mapping[str, _builtins.str]:
        """
        Metadata of the server group.
        """
        return pulumi.get(self, "metadata")

    @_builtins.property
    @pulumi.getter
    def name(self) -> _builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def policies(self) -> Sequence[_builtins.str]:
        """
        Policy name associated with the server group.
        """
        return pulumi.get(self, "policies")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> _builtins.str:
        """
        Project ID of the server group.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter
    def region(self) -> _builtins.str:
        return pulumi.get(self, "region")

    @_builtins.property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetServergroupV2RuleResult']:
        """
        Rules which are applied to specified policy.
        """
        return pulumi.get(self, "rules")

    @_builtins.property
    @pulumi.getter(name="userId")
    def user_id(self) -> _builtins.str:
        """
        User ID of the server group.
        """
        return pulumi.get(self, "user_id")


class AwaitableGetServergroupV2Result(GetServergroupV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServergroupV2Result(
            id=self.id,
            members=self.members,
            metadata=self.metadata,
            name=self.name,
            policies=self.policies,
            project_id=self.project_id,
            region=self.region,
            rules=self.rules,
            user_id=self.user_id)


def get_servergroup_v2(name: Optional[_builtins.str] = None,
                       region: Optional[_builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServergroupV2Result:
    """
    Use this data source to get information about server groups
    by name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    test = openstack.compute.get_servergroup_v2(name="test")
    ```


    :param _builtins.str name: The name of the server group.
    :param _builtins.str region: The region in which to obtain the V2 Compute client.
           If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:compute/getServergroupV2:getServergroupV2', __args__, opts=opts, typ=GetServergroupV2Result).value

    return AwaitableGetServergroupV2Result(
        id=pulumi.get(__ret__, 'id'),
        members=pulumi.get(__ret__, 'members'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        policies=pulumi.get(__ret__, 'policies'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        rules=pulumi.get(__ret__, 'rules'),
        user_id=pulumi.get(__ret__, 'user_id'))
def get_servergroup_v2_output(name: Optional[pulumi.Input[_builtins.str]] = None,
                              region: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServergroupV2Result]:
    """
    Use this data source to get information about server groups
    by name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    test = openstack.compute.get_servergroup_v2(name="test")
    ```


    :param _builtins.str name: The name of the server group.
    :param _builtins.str region: The region in which to obtain the V2 Compute client.
           If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:compute/getServergroupV2:getServergroupV2', __args__, opts=opts, typ=GetServergroupV2Result)
    return __ret__.apply(lambda __response__: GetServergroupV2Result(
        id=pulumi.get(__response__, 'id'),
        members=pulumi.get(__response__, 'members'),
        metadata=pulumi.get(__response__, 'metadata'),
        name=pulumi.get(__response__, 'name'),
        policies=pulumi.get(__response__, 'policies'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        rules=pulumi.get(__response__, 'rules'),
        user_id=pulumi.get(__response__, 'user_id')))
