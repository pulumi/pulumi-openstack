# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetTrunkResult',
    'AwaitableGetTrunkResult',
    'get_trunk',
    'get_trunk_output',
]

@pulumi.output_type
class GetTrunkResult:
    """
    A collection of values returned by getTrunk.
    """
    def __init__(__self__, admin_state_up=None, all_tags=None, description=None, id=None, name=None, port_id=None, project_id=None, region=None, status=None, sub_ports=None, tags=None, trunk_id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        pulumi.set(__self__, "admin_state_up", admin_state_up)
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        pulumi.set(__self__, "all_tags", all_tags)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        pulumi.set(__self__, "port_id", port_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if sub_ports and not isinstance(sub_ports, list):
            raise TypeError("Expected argument 'sub_ports' to be a list")
        pulumi.set(__self__, "sub_ports", sub_ports)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if trunk_id and not isinstance(trunk_id, str):
            raise TypeError("Expected argument 'trunk_id' to be a str")
        pulumi.set(__self__, "trunk_id", trunk_id)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[bool]:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Sequence[str]:
        """
        The set of string tags applied on the trunk.
        """
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[str]:
        """
        The ID of the trunk subport.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subPorts")
    def sub_ports(self) -> Sequence['outputs.GetTrunkSubPortResult']:
        """
        The set of the trunk subports. The structure of each subport is
        described below.
        """
        return pulumi.get(self, "sub_ports")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trunkId")
    def trunk_id(self) -> Optional[str]:
        return pulumi.get(self, "trunk_id")


class AwaitableGetTrunkResult(GetTrunkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrunkResult(
            admin_state_up=self.admin_state_up,
            all_tags=self.all_tags,
            description=self.description,
            id=self.id,
            name=self.name,
            port_id=self.port_id,
            project_id=self.project_id,
            region=self.region,
            status=self.status,
            sub_ports=self.sub_ports,
            tags=self.tags,
            trunk_id=self.trunk_id)


def get_trunk(admin_state_up: Optional[bool] = None,
              description: Optional[str] = None,
              name: Optional[str] = None,
              port_id: Optional[str] = None,
              project_id: Optional[str] = None,
              region: Optional[str] = None,
              status: Optional[str] = None,
              tags: Optional[Sequence[str]] = None,
              trunk_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrunkResult:
    """
    Use this data source to get the ID of an available OpenStack trunk.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    trunk1 = openstack.networking.get_trunk(name="trunk_1")
    ```


    :param bool admin_state_up: The administrative state of the trunk.
    :param str description: Human-readable description of the trunk.
    :param str name: The name of the trunk.
    :param str port_id: The ID of the trunk parent port.
    :param str project_id: The owner of the trunk.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve trunk ids. If omitted, the
           `region` argument of the provider is used.
    :param str status: The status of the trunk.
    :param Sequence[str] tags: The list of trunk tags to filter.
    :param str trunk_id: The ID of the trunk.
    """
    __args__ = dict()
    __args__['adminStateUp'] = admin_state_up
    __args__['description'] = description
    __args__['name'] = name
    __args__['portId'] = port_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['trunkId'] = trunk_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getTrunk:getTrunk', __args__, opts=opts, typ=GetTrunkResult).value

    return AwaitableGetTrunkResult(
        admin_state_up=pulumi.get(__ret__, 'admin_state_up'),
        all_tags=pulumi.get(__ret__, 'all_tags'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        port_id=pulumi.get(__ret__, 'port_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        sub_ports=pulumi.get(__ret__, 'sub_ports'),
        tags=pulumi.get(__ret__, 'tags'),
        trunk_id=pulumi.get(__ret__, 'trunk_id'))


@_utilities.lift_output_func(get_trunk)
def get_trunk_output(admin_state_up: Optional[pulumi.Input[Optional[bool]]] = None,
                     description: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     port_id: Optional[pulumi.Input[Optional[str]]] = None,
                     project_id: Optional[pulumi.Input[Optional[str]]] = None,
                     region: Optional[pulumi.Input[Optional[str]]] = None,
                     status: Optional[pulumi.Input[Optional[str]]] = None,
                     tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     trunk_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrunkResult]:
    """
    Use this data source to get the ID of an available OpenStack trunk.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    trunk1 = openstack.networking.get_trunk(name="trunk_1")
    ```


    :param bool admin_state_up: The administrative state of the trunk.
    :param str description: Human-readable description of the trunk.
    :param str name: The name of the trunk.
    :param str port_id: The ID of the trunk parent port.
    :param str project_id: The owner of the trunk.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve trunk ids. If omitted, the
           `region` argument of the provider is used.
    :param str status: The status of the trunk.
    :param Sequence[str] tags: The list of trunk tags to filter.
    :param str trunk_id: The ID of the trunk.
    """
    ...
