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
    'GetPortResult',
    'AwaitableGetPortResult',
    'get_port',
    'get_port_output',
]

@pulumi.output_type
class GetPortResult:
    """
    A collection of values returned by getPort.
    """
    def __init__(__self__, admin_state_up=None, all_fixed_ips=None, all_security_group_ids=None, all_tags=None, allowed_address_pairs=None, bindings=None, description=None, device_id=None, device_owner=None, dns_assignments=None, dns_name=None, extra_dhcp_options=None, fixed_ip=None, id=None, mac_address=None, name=None, network_id=None, port_id=None, project_id=None, region=None, security_group_ids=None, status=None, tags=None, tenant_id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        pulumi.set(__self__, "admin_state_up", admin_state_up)
        if all_fixed_ips and not isinstance(all_fixed_ips, list):
            raise TypeError("Expected argument 'all_fixed_ips' to be a list")
        pulumi.set(__self__, "all_fixed_ips", all_fixed_ips)
        if all_security_group_ids and not isinstance(all_security_group_ids, list):
            raise TypeError("Expected argument 'all_security_group_ids' to be a list")
        pulumi.set(__self__, "all_security_group_ids", all_security_group_ids)
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        pulumi.set(__self__, "all_tags", all_tags)
        if allowed_address_pairs and not isinstance(allowed_address_pairs, list):
            raise TypeError("Expected argument 'allowed_address_pairs' to be a list")
        pulumi.set(__self__, "allowed_address_pairs", allowed_address_pairs)
        if bindings and not isinstance(bindings, list):
            raise TypeError("Expected argument 'bindings' to be a list")
        pulumi.set(__self__, "bindings", bindings)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device_id and not isinstance(device_id, str):
            raise TypeError("Expected argument 'device_id' to be a str")
        pulumi.set(__self__, "device_id", device_id)
        if device_owner and not isinstance(device_owner, str):
            raise TypeError("Expected argument 'device_owner' to be a str")
        pulumi.set(__self__, "device_owner", device_owner)
        if dns_assignments and not isinstance(dns_assignments, list):
            raise TypeError("Expected argument 'dns_assignments' to be a list")
        pulumi.set(__self__, "dns_assignments", dns_assignments)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if extra_dhcp_options and not isinstance(extra_dhcp_options, list):
            raise TypeError("Expected argument 'extra_dhcp_options' to be a list")
        pulumi.set(__self__, "extra_dhcp_options", extra_dhcp_options)
        if fixed_ip and not isinstance(fixed_ip, str):
            raise TypeError("Expected argument 'fixed_ip' to be a str")
        pulumi.set(__self__, "fixed_ip", fixed_ip)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        pulumi.set(__self__, "port_id", port_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        pulumi.set(__self__, "security_group_ids", security_group_ids)
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
    @pulumi.getter(name="allFixedIps")
    def all_fixed_ips(self) -> Sequence[str]:
        return pulumi.get(self, "all_fixed_ips")

    @property
    @pulumi.getter(name="allSecurityGroupIds")
    def all_security_group_ids(self) -> Sequence[str]:
        return pulumi.get(self, "all_security_group_ids")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Sequence[str]:
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="allowedAddressPairs")
    def allowed_address_pairs(self) -> Sequence['outputs.GetPortAllowedAddressPairResult']:
        return pulumi.get(self, "allowed_address_pairs")

    @property
    @pulumi.getter
    def bindings(self) -> Sequence['outputs.GetPortBindingResult']:
        return pulumi.get(self, "bindings")

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
    @pulumi.getter(name="dnsAssignments")
    def dns_assignments(self) -> Sequence[Mapping[str, Any]]:
        return pulumi.get(self, "dns_assignments")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> Optional[str]:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="extraDhcpOptions")
    def extra_dhcp_options(self) -> Sequence['outputs.GetPortExtraDhcpOptionResult']:
        return pulumi.get(self, "extra_dhcp_options")

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
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[str]:
        return pulumi.get(self, "port_id")

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


class AwaitableGetPortResult(GetPortResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortResult(
            admin_state_up=self.admin_state_up,
            all_fixed_ips=self.all_fixed_ips,
            all_security_group_ids=self.all_security_group_ids,
            all_tags=self.all_tags,
            allowed_address_pairs=self.allowed_address_pairs,
            bindings=self.bindings,
            description=self.description,
            device_id=self.device_id,
            device_owner=self.device_owner,
            dns_assignments=self.dns_assignments,
            dns_name=self.dns_name,
            extra_dhcp_options=self.extra_dhcp_options,
            fixed_ip=self.fixed_ip,
            id=self.id,
            mac_address=self.mac_address,
            name=self.name,
            network_id=self.network_id,
            port_id=self.port_id,
            project_id=self.project_id,
            region=self.region,
            security_group_ids=self.security_group_ids,
            status=self.status,
            tags=self.tags,
            tenant_id=self.tenant_id)


def get_port(admin_state_up: Optional[bool] = None,
             description: Optional[str] = None,
             device_id: Optional[str] = None,
             device_owner: Optional[str] = None,
             dns_name: Optional[str] = None,
             fixed_ip: Optional[str] = None,
             mac_address: Optional[str] = None,
             name: Optional[str] = None,
             network_id: Optional[str] = None,
             port_id: Optional[str] = None,
             project_id: Optional[str] = None,
             region: Optional[str] = None,
             security_group_ids: Optional[Sequence[str]] = None,
             status: Optional[str] = None,
             tags: Optional[Sequence[str]] = None,
             tenant_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortResult:
    """
    Use this data source to access information about an existing resource.
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
    __args__['portId'] = port_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['securityGroupIds'] = security_group_ids
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getPort:getPort', __args__, opts=opts, typ=GetPortResult).value

    return AwaitableGetPortResult(
        admin_state_up=pulumi.get(__ret__, 'admin_state_up'),
        all_fixed_ips=pulumi.get(__ret__, 'all_fixed_ips'),
        all_security_group_ids=pulumi.get(__ret__, 'all_security_group_ids'),
        all_tags=pulumi.get(__ret__, 'all_tags'),
        allowed_address_pairs=pulumi.get(__ret__, 'allowed_address_pairs'),
        bindings=pulumi.get(__ret__, 'bindings'),
        description=pulumi.get(__ret__, 'description'),
        device_id=pulumi.get(__ret__, 'device_id'),
        device_owner=pulumi.get(__ret__, 'device_owner'),
        dns_assignments=pulumi.get(__ret__, 'dns_assignments'),
        dns_name=pulumi.get(__ret__, 'dns_name'),
        extra_dhcp_options=pulumi.get(__ret__, 'extra_dhcp_options'),
        fixed_ip=pulumi.get(__ret__, 'fixed_ip'),
        id=pulumi.get(__ret__, 'id'),
        mac_address=pulumi.get(__ret__, 'mac_address'),
        name=pulumi.get(__ret__, 'name'),
        network_id=pulumi.get(__ret__, 'network_id'),
        port_id=pulumi.get(__ret__, 'port_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        security_group_ids=pulumi.get(__ret__, 'security_group_ids'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))


@_utilities.lift_output_func(get_port)
def get_port_output(admin_state_up: Optional[pulumi.Input[Optional[bool]]] = None,
                    description: Optional[pulumi.Input[Optional[str]]] = None,
                    device_id: Optional[pulumi.Input[Optional[str]]] = None,
                    device_owner: Optional[pulumi.Input[Optional[str]]] = None,
                    dns_name: Optional[pulumi.Input[Optional[str]]] = None,
                    fixed_ip: Optional[pulumi.Input[Optional[str]]] = None,
                    mac_address: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    network_id: Optional[pulumi.Input[Optional[str]]] = None,
                    port_id: Optional[pulumi.Input[Optional[str]]] = None,
                    project_id: Optional[pulumi.Input[Optional[str]]] = None,
                    region: Optional[pulumi.Input[Optional[str]]] = None,
                    security_group_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    status: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
