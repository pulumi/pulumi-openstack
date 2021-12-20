# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetQuotaV2Result',
    'AwaitableGetQuotaV2Result',
    'get_quota_v2',
    'get_quota_v2_output',
]

@pulumi.output_type
class GetQuotaV2Result:
    """
    A collection of values returned by getQuotaV2.
    """
    def __init__(__self__, floatingip=None, id=None, network=None, port=None, project_id=None, rbac_policy=None, region=None, router=None, security_group=None, security_group_rule=None, subnet=None, subnetpool=None):
        if floatingip and not isinstance(floatingip, int):
            raise TypeError("Expected argument 'floatingip' to be a int")
        pulumi.set(__self__, "floatingip", floatingip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if network and not isinstance(network, int):
            raise TypeError("Expected argument 'network' to be a int")
        pulumi.set(__self__, "network", network)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if rbac_policy and not isinstance(rbac_policy, int):
            raise TypeError("Expected argument 'rbac_policy' to be a int")
        pulumi.set(__self__, "rbac_policy", rbac_policy)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if router and not isinstance(router, int):
            raise TypeError("Expected argument 'router' to be a int")
        pulumi.set(__self__, "router", router)
        if security_group and not isinstance(security_group, int):
            raise TypeError("Expected argument 'security_group' to be a int")
        pulumi.set(__self__, "security_group", security_group)
        if security_group_rule and not isinstance(security_group_rule, int):
            raise TypeError("Expected argument 'security_group_rule' to be a int")
        pulumi.set(__self__, "security_group_rule", security_group_rule)
        if subnet and not isinstance(subnet, int):
            raise TypeError("Expected argument 'subnet' to be a int")
        pulumi.set(__self__, "subnet", subnet)
        if subnetpool and not isinstance(subnetpool, int):
            raise TypeError("Expected argument 'subnetpool' to be a int")
        pulumi.set(__self__, "subnetpool", subnetpool)

    @property
    @pulumi.getter
    def floatingip(self) -> int:
        """
        The number of allowed floating ips.
        """
        return pulumi.get(self, "floatingip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def network(self) -> int:
        """
        The number of allowed networks.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        The number of allowed ports.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="rbacPolicy")
    def rbac_policy(self) -> int:
        """
        The number of allowed rbac policies.
        """
        return pulumi.get(self, "rbac_policy")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def router(self) -> int:
        """
        The amount of allowed routers.
        """
        return pulumi.get(self, "router")

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> int:
        """
        The number of allowed security groups.
        """
        return pulumi.get(self, "security_group")

    @property
    @pulumi.getter(name="securityGroupRule")
    def security_group_rule(self) -> int:
        """
        The number of allowed security group rules.
        """
        return pulumi.get(self, "security_group_rule")

    @property
    @pulumi.getter
    def subnet(self) -> int:
        """
        The number of allowed subnets.
        * `subnetpool-` - The number of allowed subnet pools.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def subnetpool(self) -> int:
        return pulumi.get(self, "subnetpool")


class AwaitableGetQuotaV2Result(GetQuotaV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQuotaV2Result(
            floatingip=self.floatingip,
            id=self.id,
            network=self.network,
            port=self.port,
            project_id=self.project_id,
            rbac_policy=self.rbac_policy,
            region=self.region,
            router=self.router,
            security_group=self.security_group,
            security_group_rule=self.security_group_rule,
            subnet=self.subnet,
            subnetpool=self.subnetpool)


def get_quota_v2(project_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQuotaV2Result:
    """
    Use this data source to get the networking quota of an OpenStack project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    quota = openstack.networking.get_quota_v2(project_id="2e367a3d29f94fd988e6ec54e305ec9d")
    ```


    :param str project_id: The id of the project to retrieve the quota.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:networking/getQuotaV2:getQuotaV2', __args__, opts=opts, typ=GetQuotaV2Result).value

    return AwaitableGetQuotaV2Result(
        floatingip=__ret__.floatingip,
        id=__ret__.id,
        network=__ret__.network,
        port=__ret__.port,
        project_id=__ret__.project_id,
        rbac_policy=__ret__.rbac_policy,
        region=__ret__.region,
        router=__ret__.router,
        security_group=__ret__.security_group,
        security_group_rule=__ret__.security_group_rule,
        subnet=__ret__.subnet,
        subnetpool=__ret__.subnetpool)


@_utilities.lift_output_func(get_quota_v2)
def get_quota_v2_output(project_id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQuotaV2Result]:
    """
    Use this data source to get the networking quota of an OpenStack project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    quota = openstack.networking.get_quota_v2(project_id="2e367a3d29f94fd988e6ec54e305ec9d")
    ```


    :param str project_id: The id of the project to retrieve the quota.
    """
    ...
