# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetShareNetworkResult:
    """
    A collection of values returned by getShareNetwork.
    """
    def __init__(__self__, cidr=None, description=None, id=None, ip_version=None, name=None, network_type=None, neutron_net_id=None, neutron_subnet_id=None, project_id=None, region=None, security_service_id=None, security_service_ids=None, segmentation_id=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        __self__.cidr = cidr
        """
        See Argument Reference above.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ip_version and not isinstance(ip_version, float):
            raise TypeError("Expected argument 'ip_version' to be a float")
        __self__.ip_version = ip_version
        """
        See Argument Reference above.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        __self__.network_type = network_type
        """
        See Argument Reference above.
        """
        if neutron_net_id and not isinstance(neutron_net_id, str):
            raise TypeError("Expected argument 'neutron_net_id' to be a str")
        __self__.neutron_net_id = neutron_net_id
        """
        See Argument Reference above.
        """
        if neutron_subnet_id and not isinstance(neutron_subnet_id, str):
            raise TypeError("Expected argument 'neutron_subnet_id' to be a str")
        __self__.neutron_subnet_id = neutron_subnet_id
        """
        See Argument Reference above.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        """
        The owner of the Share Network.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if security_service_id and not isinstance(security_service_id, str):
            raise TypeError("Expected argument 'security_service_id' to be a str")
        __self__.security_service_id = security_service_id
        """
        See Argument Reference above.
        """
        if security_service_ids and not isinstance(security_service_ids, list):
            raise TypeError("Expected argument 'security_service_ids' to be a list")
        __self__.security_service_ids = security_service_ids
        """
        The list of security service IDs associated with
        the share network.
        """
        if segmentation_id and not isinstance(segmentation_id, float):
            raise TypeError("Expected argument 'segmentation_id' to be a float")
        __self__.segmentation_id = segmentation_id
        """
        See Argument Reference above.
        """
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

def get_share_network(description=None,ip_version=None,name=None,network_type=None,neutron_net_id=None,neutron_subnet_id=None,region=None,security_service_id=None,segmentation_id=None,opts=None):
    """
    Use this data source to get the ID of an available Shared File System share network.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/sharedfilesystem_sharenetwork_v2.html.markdown.


    :param str description: The human-readable description of the share network.
    :param float ip_version: The IP version of the share network. Can either be 4 or 6.
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
    :param float segmentation_id: The share network segmentation ID.
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
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:sharedfilesystem/getShareNetwork:getShareNetwork', __args__, opts=opts).value

    return AwaitableGetShareNetworkResult(
        cidr=__ret__.get('cidr'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        ip_version=__ret__.get('ipVersion'),
        name=__ret__.get('name'),
        network_type=__ret__.get('networkType'),
        neutron_net_id=__ret__.get('neutronNetId'),
        neutron_subnet_id=__ret__.get('neutronSubnetId'),
        project_id=__ret__.get('projectId'),
        region=__ret__.get('region'),
        security_service_id=__ret__.get('securityServiceId'),
        security_service_ids=__ret__.get('securityServiceIds'),
        segmentation_id=__ret__.get('segmentationId'))
