# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFwGroupV2Result',
    'AwaitableGetFwGroupV2Result',
    'get_fw_group_v2',
    'get_fw_group_v2_output',
]

@pulumi.output_type
class GetFwGroupV2Result:
    """
    A collection of values returned by getFwGroupV2.
    """
    def __init__(__self__, admin_state_up=None, description=None, egress_firewall_policy_id=None, group_id=None, id=None, ingress_firewall_policy_id=None, name=None, ports=None, project_id=None, region=None, shared=None, status=None, tenant_id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if egress_firewall_policy_id and not isinstance(egress_firewall_policy_id, str):
            raise TypeError("Expected argument 'egress_firewall_policy_id' to be a str")
        pulumi.set(__self__, "egress_firewall_policy_id", egress_firewall_policy_id)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ingress_firewall_policy_id and not isinstance(ingress_firewall_policy_id, str):
            raise TypeError("Expected argument 'ingress_firewall_policy_id' to be a str")
        pulumi.set(__self__, "ingress_firewall_policy_id", ingress_firewall_policy_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ports and not isinstance(ports, list):
            raise TypeError("Expected argument 'ports' to be a list")
        pulumi.set(__self__, "ports", ports)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        pulumi.set(__self__, "shared", shared)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> bool:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="egressFirewallPolicyId")
    def egress_firewall_policy_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "egress_firewall_policy_id")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ingressFirewallPolicyId")
    def ingress_firewall_policy_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "ingress_firewall_policy_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def ports(self) -> Sequence[str]:
        """
        Ports associated with the firewall group.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def shared(self) -> bool:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetFwGroupV2Result(GetFwGroupV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFwGroupV2Result(
            admin_state_up=self.admin_state_up,
            description=self.description,
            egress_firewall_policy_id=self.egress_firewall_policy_id,
            group_id=self.group_id,
            id=self.id,
            ingress_firewall_policy_id=self.ingress_firewall_policy_id,
            name=self.name,
            ports=self.ports,
            project_id=self.project_id,
            region=self.region,
            shared=self.shared,
            status=self.status,
            tenant_id=self.tenant_id)


def get_fw_group_v2(admin_state_up: Optional[bool] = None,
                    description: Optional[str] = None,
                    egress_firewall_policy_id: Optional[str] = None,
                    group_id: Optional[str] = None,
                    ingress_firewall_policy_id: Optional[str] = None,
                    name: Optional[str] = None,
                    project_id: Optional[str] = None,
                    region: Optional[str] = None,
                    shared: Optional[bool] = None,
                    status: Optional[str] = None,
                    tenant_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFwGroupV2Result:
    """
    Use this data source to get information of an available OpenStack firewall group v2.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    group = openstack.get_fw_group_v2(name="tf_test_group")
    ```


    :param bool admin_state_up: Administrative up/down status for the firewall group.
    :param str description: Human-readable description of the firewall group.
    :param str egress_firewall_policy_id: The egress policy ID of the firewall group.
    :param str group_id: The ID of the firewall group.
    :param str ingress_firewall_policy_id: The ingress policy ID of the firewall group.
    :param str name: The name of the firewall group.
    :param str project_id: This argument conflict and interchangeable with
           `tenant_id`. The owner of the firewall group.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve firewall group ids. If omitted, the
           `region` argument of the provider is used.
    :param bool shared: The sharing status of the firewall group.
    :param str status: Enabled status for the firewall group.
    :param str tenant_id: This argument conflict and interchangeable with
           `project_id`. The owner of the firewall group.
    """
    __args__ = dict()
    __args__['adminStateUp'] = admin_state_up
    __args__['description'] = description
    __args__['egressFirewallPolicyId'] = egress_firewall_policy_id
    __args__['groupId'] = group_id
    __args__['ingressFirewallPolicyId'] = ingress_firewall_policy_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['shared'] = shared
    __args__['status'] = status
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:index/getFwGroupV2:getFwGroupV2', __args__, opts=opts, typ=GetFwGroupV2Result).value

    return AwaitableGetFwGroupV2Result(
        admin_state_up=pulumi.get(__ret__, 'admin_state_up'),
        description=pulumi.get(__ret__, 'description'),
        egress_firewall_policy_id=pulumi.get(__ret__, 'egress_firewall_policy_id'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        ingress_firewall_policy_id=pulumi.get(__ret__, 'ingress_firewall_policy_id'),
        name=pulumi.get(__ret__, 'name'),
        ports=pulumi.get(__ret__, 'ports'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        shared=pulumi.get(__ret__, 'shared'),
        status=pulumi.get(__ret__, 'status'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_fw_group_v2)
def get_fw_group_v2_output(admin_state_up: Optional[pulumi.Input[Optional[bool]]] = None,
                           description: Optional[pulumi.Input[Optional[str]]] = None,
                           egress_firewall_policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                           group_id: Optional[pulumi.Input[Optional[str]]] = None,
                           ingress_firewall_policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                           name: Optional[pulumi.Input[Optional[str]]] = None,
                           project_id: Optional[pulumi.Input[Optional[str]]] = None,
                           region: Optional[pulumi.Input[Optional[str]]] = None,
                           shared: Optional[pulumi.Input[Optional[bool]]] = None,
                           status: Optional[pulumi.Input[Optional[str]]] = None,
                           tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFwGroupV2Result]:
    """
    Use this data source to get information of an available OpenStack firewall group v2.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    group = openstack.get_fw_group_v2(name="tf_test_group")
    ```


    :param bool admin_state_up: Administrative up/down status for the firewall group.
    :param str description: Human-readable description of the firewall group.
    :param str egress_firewall_policy_id: The egress policy ID of the firewall group.
    :param str group_id: The ID of the firewall group.
    :param str ingress_firewall_policy_id: The ingress policy ID of the firewall group.
    :param str name: The name of the firewall group.
    :param str project_id: This argument conflict and interchangeable with
           `tenant_id`. The owner of the firewall group.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve firewall group ids. If omitted, the
           `region` argument of the provider is used.
    :param bool shared: The sharing status of the firewall group.
    :param str status: Enabled status for the firewall group.
    :param str tenant_id: This argument conflict and interchangeable with
           `project_id`. The owner of the firewall group.
    """
    ...