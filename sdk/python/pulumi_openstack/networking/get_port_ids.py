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
    'GetPortIdsResult',
    'AwaitableGetPortIdsResult',
    'get_port_ids',
    'get_port_ids_output',
]

@pulumi.output_type
class GetPortIdsResult:
    """
    A collection of values returned by getPortIds.
    """
    def __init__(__self__, admin_state_up=None, description=None, device_id=None, device_owner=None, dns_name=None, fixed_ip=None, id=None, ids=None, mac_address=None, name=None, network_id=None, project_id=None, region=None, security_group_ids=None, sort_direction=None, sort_key=None, status=None, tags=None, tenant_id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        pulumi.set(__self__, "admin_state_up", admin_state_up)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device_id and not isinstance(device_id, str):
            raise TypeError("Expected argument 'device_id' to be a str")
        pulumi.set(__self__, "device_id", device_id)
        if device_owner and not isinstance(device_owner, str):
            raise TypeError("Expected argument 'device_owner' to be a str")
        pulumi.set(__self__, "device_owner", device_owner)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if fixed_ip and not isinstance(fixed_ip, str):
            raise TypeError("Expected argument 'fixed_ip' to be a str")
        pulumi.set(__self__, "fixed_ip", fixed_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        pulumi.set(__self__, "sort_direction", sort_direction)
        if sort_key and not isinstance(sort_key, str):
            raise TypeError("Expected argument 'sort_key' to be a str")
        pulumi.set(__self__, "sort_key", sort_key)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[bool]:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[str]:
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="deviceOwner")
    def device_owner(self) -> Optional[str]:
        return pulumi.get(self, "device_owner")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> Optional[str]:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="fixedIp")
    def fixed_ip(self) -> Optional[str]:
        return pulumi.get(self, "fixed_ip")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> Optional[str]:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[str]:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="sortDirection")
    def sort_direction(self) -> Optional[str]:
        return pulumi.get(self, "sort_direction")

    @property
    @pulumi.getter(name="sortKey")
    def sort_key(self) -> Optional[str]:
        return pulumi.get(self, "sort_key")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")


class AwaitableGetPortIdsResult(GetPortIdsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortIdsResult(
            admin_state_up=self.admin_state_up,
            description=self.description,
            device_id=self.device_id,
            device_owner=self.device_owner,
            dns_name=self.dns_name,
            fixed_ip=self.fixed_ip,
            id=self.id,
            ids=self.ids,
            mac_address=self.mac_address,
            name=self.name,
            network_id=self.network_id,
            project_id=self.project_id,
            region=self.region,
            security_group_ids=self.security_group_ids,
            sort_direction=self.sort_direction,
            sort_key=self.sort_key,
            status=self.status,
            tags=self.tags,
            tenant_id=self.tenant_id)


def get_port_ids(admin_state_up: Optional[bool] = None,
                 description: Optional[str] = None,
                 device_id: Optional[str] = None,
                 device_owner: Optional[str] = None,
                 dns_name: Optional[str] = None,
                 fixed_ip: Optional[str] = None,
                 mac_address: Optional[str] = None,
                 name: Optional[str] = None,
                 network_id: Optional[str] = None,
                 project_id: Optional[str] = None,
                 region: Optional[str] = None,
                 security_group_ids: Optional[Sequence[str]] = None,
                 sort_direction: Optional[str] = None,
                 sort_key: Optional[str] = None,
                 status: Optional[str] = None,
                 tags: Optional[Sequence[str]] = None,
                 tenant_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortIdsResult:
    """
    Use this data source to get a list of Openstack Port IDs matching the
    specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    ports = openstack.networking.get_port_ids(name="port")
    ```


    :param bool admin_state_up: The administrative state of the port.
    :param str description: Human-readable description of the port.
    :param str device_id: The ID of the device the port belongs to.
    :param str device_owner: The device owner of the port.
    :param str fixed_ip: The port IP address filter.
    :param str mac_address: The MAC address of the port.
    :param str name: The name of the port.
    :param str network_id: The ID of the network the port belongs to.
    :param str project_id: The owner of the port.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve port ids. If omitted, the
           `region` argument of the provider is used.
    :param Sequence[str] security_group_ids: The list of port security group IDs to filter.
    :param str sort_direction: Order the results in either `asc` or `desc`.
           Defaults to none.
    :param str sort_key: Sort ports based on a certain key. Defaults to none.
    :param str status: The status of the port.
    :param Sequence[str] tags: The list of port tags to filter.
    """
    __args__ = dict()
    __args__['adminStateUp'] = admin_state_up
    __args__['description'] = description
    __args__['deviceId'] = device_id
    __args__['deviceOwner'] = device_owner
    __args__['dnsName'] = dns_name
    __args__['fixedIp'] = fixed_ip
    __args__['macAddress'] = mac_address
    __args__['name'] = name
    __args__['networkId'] = network_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['securityGroupIds'] = security_group_ids
    __args__['sortDirection'] = sort_direction
    __args__['sortKey'] = sort_key
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getPortIds:getPortIds', __args__, opts=opts, typ=GetPortIdsResult).value

    return AwaitableGetPortIdsResult(
        admin_state_up=pulumi.get(__ret__, 'admin_state_up'),
        description=pulumi.get(__ret__, 'description'),
        device_id=pulumi.get(__ret__, 'device_id'),
        device_owner=pulumi.get(__ret__, 'device_owner'),
        dns_name=pulumi.get(__ret__, 'dns_name'),
        fixed_ip=pulumi.get(__ret__, 'fixed_ip'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        mac_address=pulumi.get(__ret__, 'mac_address'),
        name=pulumi.get(__ret__, 'name'),
        network_id=pulumi.get(__ret__, 'network_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        sort_direction=pulumi.get(__ret__, 'sort_direction'),
        sort_key=pulumi.get(__ret__, 'sort_key'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_port_ids)
def get_port_ids_output(admin_state_up: Optional[pulumi.Input[Optional[bool]]] = None,
                        description: Optional[pulumi.Input[Optional[str]]] = None,
                        device_id: Optional[pulumi.Input[Optional[str]]] = None,
                        device_owner: Optional[pulumi.Input[Optional[str]]] = None,
                        dns_name: Optional[pulumi.Input[Optional[str]]] = None,
                        fixed_ip: Optional[pulumi.Input[Optional[str]]] = None,
                        mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        network_id: Optional[pulumi.Input[Optional[str]]] = None,
                        project_id: Optional[pulumi.Input[Optional[str]]] = None,
                        region: Optional[pulumi.Input[Optional[str]]] = None,
                        security_group_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        sort_direction: Optional[pulumi.Input[Optional[str]]] = None,
                        sort_key: Optional[pulumi.Input[Optional[str]]] = None,
                        status: Optional[pulumi.Input[Optional[str]]] = None,
                        tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortIdsResult]:
    """
    Use this data source to get a list of Openstack Port IDs matching the
    specified criteria.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    ports = openstack.networking.get_port_ids(name="port")
    ```


    :param bool admin_state_up: The administrative state of the port.
    :param str description: Human-readable description of the port.
    :param str device_id: The ID of the device the port belongs to.
    :param str device_owner: The device owner of the port.
    :param str fixed_ip: The port IP address filter.
    :param str mac_address: The MAC address of the port.
    :param str name: The name of the port.
    :param str network_id: The ID of the network the port belongs to.
    :param str project_id: The owner of the port.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve port ids. If omitted, the
           `region` argument of the provider is used.
    :param Sequence[str] security_group_ids: The list of port security group IDs to filter.
    :param str sort_direction: Order the results in either `asc` or `desc`.
           Defaults to none.
    :param str sort_key: Sort ports based on a certain key. Defaults to none.
    :param str status: The status of the port.
    :param Sequence[str] tags: The list of port tags to filter.
    """
    ...
