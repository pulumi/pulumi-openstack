# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'GetShareNetworkResult',
    'AwaitableGetShareNetworkResult',
    'get_share_network',
]

@pulumi.output_type
class GetShareNetworkResult:
    """
    A collection of values returned by getShareNetwork.
    """
    def __init__(__self__, cidr=None, description=None, id=None, ip_version=None, name=None, network_type=None, neutron_net_id=None, neutron_subnet_id=None, project_id=None, region=None, security_service_id=None, security_service_ids=None, segmentation_id=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_version and not isinstance(ip_version, int):
            raise TypeError("Expected argument 'ip_version' to be a int")
        pulumi.set(__self__, "ip_version", ip_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if neutron_net_id and not isinstance(neutron_net_id, str):
            raise TypeError("Expected argument 'neutron_net_id' to be a str")
        pulumi.set(__self__, "neutron_net_id", neutron_net_id)
        if neutron_subnet_id and not isinstance(neutron_subnet_id, str):
            raise TypeError("Expected argument 'neutron_subnet_id' to be a str")
        pulumi.set(__self__, "neutron_subnet_id", neutron_subnet_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_service_id and not isinstance(security_service_id, str):
            raise TypeError("Expected argument 'security_service_id' to be a str")
        pulumi.set(__self__, "security_service_id", security_service_id)
        if security_service_ids and not isinstance(security_service_ids, list):
            raise TypeError("Expected argument 'security_service_ids' to be a list")
        pulumi.set(__self__, "security_service_ids", security_service_ids)
        if segmentation_id and not isinstance(segmentation_id, int):
            raise TypeError("Expected argument 'segmentation_id' to be a int")
        pulumi.set(__self__, "segmentation_id", segmentation_id)

    @property
    @pulumi.getter
    def cidr(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> int:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="neutronNetId")
    def neutron_net_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "neutron_net_id")

    @property
    @pulumi.getter(name="neutronSubnetId")
    def neutron_subnet_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "neutron_subnet_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The owner of the Share Network.
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
    @pulumi.getter(name="securityServiceId")
    def security_service_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "security_service_id")

    @property
    @pulumi.getter(name="securityServiceIds")
    def security_service_ids(self) -> Sequence[str]:
        """
        The list of security service IDs associated with
        the share network.
        """
        return pulumi.get(self, "security_service_ids")

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> int:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "segmentation_id")


class AwaitableGetShareNetworkResult(GetShareNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetShareNetworkResult(
            cidr=self.cidr,
            description=self.description,
            id=self.id,
            ip_version=self.ip_version,
            name=self.name,
            network_type=self.network_type,
            neutron_net_id=self.neutron_net_id,
            neutron_subnet_id=self.neutron_subnet_id,
            project_id=self.project_id,
            region=self.region,
            security_service_id=self.security_service_id,
            security_service_ids=self.security_service_ids,
            segmentation_id=self.segmentation_id)


def get_share_network(description: Optional[str] = None,
                      ip_version: Optional[int] = None,
                      name: Optional[str] = None,
                      network_type: Optional[str] = None,
                      neutron_net_id: Optional[str] = None,
                      neutron_subnet_id: Optional[str] = None,
                      region: Optional[str] = None,
                      security_service_id: Optional[str] = None,
                      segmentation_id: Optional[int] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetShareNetworkResult:
    """
    Use this data source to get the ID of an available Shared File System share network.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    sharenetwork1 = openstack.sharedfilesystem.get_share_network(name="sharenetwork_1")
    ```


    :param str description: The human-readable description of the share network.
    :param int ip_version: The IP version of the share network. Can either be 4 or 6.
    :param str name: The name of the share network.
    :param str network_type: The share network type. Can either be VLAN, VXLAN,
           GRE, or flat.
    :param str neutron_net_id: The neutron network UUID of the share network.
    :param str neutron_subnet_id: The neutron subnet UUID of the share network.
    :param str region: The region in which to obtain the V2 Shared File System client.
           A Shared File System client is needed to read a share network. If omitted, the
           `region` argument of the provider is used.
    :param str security_service_id: The security service IDs associated with
           the share network.
    :param int segmentation_id: The share network segmentation ID.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['ipVersion'] = ip_version
    __args__['name'] = name
    __args__['networkType'] = network_type
    __args__['neutronNetId'] = neutron_net_id
    __args__['neutronSubnetId'] = neutron_subnet_id
    __args__['region'] = region
    __args__['securityServiceId'] = security_service_id
    __args__['segmentationId'] = segmentation_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:sharedfilesystem/getShareNetwork:getShareNetwork', __args__, opts=opts, typ=GetShareNetworkResult).value

    return AwaitableGetShareNetworkResult(
        cidr=__ret__.cidr,
        description=__ret__.description,
        id=__ret__.id,
        ip_version=__ret__.ip_version,
        name=__ret__.name,
        network_type=__ret__.network_type,
        neutron_net_id=__ret__.neutron_net_id,
        neutron_subnet_id=__ret__.neutron_subnet_id,
        project_id=__ret__.project_id,
        region=__ret__.region,
        security_service_id=__ret__.security_service_id,
        security_service_ids=__ret__.security_service_ids,
        segmentation_id=__ret__.segmentation_id)
