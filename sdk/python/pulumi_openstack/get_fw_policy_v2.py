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
from . import _utilities

__all__ = [
    'GetFwPolicyV2Result',
    'AwaitableGetFwPolicyV2Result',
    'get_fw_policy_v2',
    'get_fw_policy_v2_output',
]

@pulumi.output_type
class GetFwPolicyV2Result:
    """
    A collection of values returned by getFwPolicyV2.
    """
    def __init__(__self__, audited=None, description=None, id=None, name=None, policy_id=None, project_id=None, region=None, rules=None, shared=None, tenant_id=None):
        if audited and not isinstance(audited, bool):
            raise TypeError("Expected argument 'audited' to be a bool")
        pulumi.set(__self__, "audited", audited)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        pulumi.set(__self__, "policy_id", policy_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        pulumi.set(__self__, "shared", shared)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def audited(self) -> builtins.bool:
        """
        The audit status of the firewall policy.
        """
        return pulumi.get(self, "audited")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> Sequence[builtins.str]:
        """
        The array of one or more firewall rules that comprise the policy.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def shared(self) -> builtins.bool:
        """
        The sharing status of the firewall policy.
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetFwPolicyV2Result(GetFwPolicyV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFwPolicyV2Result(
            audited=self.audited,
            description=self.description,
            id=self.id,
            name=self.name,
            policy_id=self.policy_id,
            project_id=self.project_id,
            region=self.region,
            rules=self.rules,
            shared=self.shared,
            tenant_id=self.tenant_id)


def get_fw_policy_v2(audited: Optional[builtins.bool] = None,
                     description: Optional[builtins.str] = None,
                     name: Optional[builtins.str] = None,
                     policy_id: Optional[builtins.str] = None,
                     project_id: Optional[builtins.str] = None,
                     region: Optional[builtins.str] = None,
                     shared: Optional[builtins.bool] = None,
                     tenant_id: Optional[builtins.str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFwPolicyV2Result:
    """
    Use this data source to get information of an available OpenStack firewall policy v2.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    policy = openstack.get_fw_policy_v2(name="tf_test_policy")
    ```


    :param builtins.bool audited: Whether this policy has been audited.
    :param builtins.str description: Human-readable description of the policy.
    :param builtins.str name: The name of the firewall policy.
    :param builtins.str policy_id: The ID of the firewall policy.
    :param builtins.str project_id: This argument conflicts and is interchangeable
           with `tenant_id`. The owner of the firewall policy.
    :param builtins.str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve firewall policy ids. If omitted, the
           `region` argument of the provider is used.
    :param builtins.bool shared: Whether this policy is shared across all projects.
    :param builtins.str tenant_id: This argument conflicts and is interchangeable
           with `project_id`. The owner of the firewall policy.
    """
    __args__ = dict()
    __args__['audited'] = audited
    __args__['description'] = description
    __args__['name'] = name
    __args__['policyId'] = policy_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['shared'] = shared
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:index/getFwPolicyV2:getFwPolicyV2', __args__, opts=opts, typ=GetFwPolicyV2Result).value

    return AwaitableGetFwPolicyV2Result(
        audited=pulumi.get(__ret__, 'audited'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        policy_id=pulumi.get(__ret__, 'policy_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        rules=pulumi.get(__ret__, 'rules'),
        shared=pulumi.get(__ret__, 'shared'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_fw_policy_v2_output(audited: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                            description: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            policy_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            project_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            shared: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                            tenant_id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFwPolicyV2Result]:
    """
    Use this data source to get information of an available OpenStack firewall policy v2.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    policy = openstack.get_fw_policy_v2(name="tf_test_policy")
    ```


    :param builtins.bool audited: Whether this policy has been audited.
    :param builtins.str description: Human-readable description of the policy.
    :param builtins.str name: The name of the firewall policy.
    :param builtins.str policy_id: The ID of the firewall policy.
    :param builtins.str project_id: This argument conflicts and is interchangeable
           with `tenant_id`. The owner of the firewall policy.
    :param builtins.str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve firewall policy ids. If omitted, the
           `region` argument of the provider is used.
    :param builtins.bool shared: Whether this policy is shared across all projects.
    :param builtins.str tenant_id: This argument conflicts and is interchangeable
           with `project_id`. The owner of the firewall policy.
    """
    __args__ = dict()
    __args__['audited'] = audited
    __args__['description'] = description
    __args__['name'] = name
    __args__['policyId'] = policy_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['shared'] = shared
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:index/getFwPolicyV2:getFwPolicyV2', __args__, opts=opts, typ=GetFwPolicyV2Result)
    return __ret__.apply(lambda __response__: GetFwPolicyV2Result(
        audited=pulumi.get(__response__, 'audited'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        policy_id=pulumi.get(__response__, 'policy_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        rules=pulumi.get(__response__, 'rules'),
        shared=pulumi.get(__response__, 'shared'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
