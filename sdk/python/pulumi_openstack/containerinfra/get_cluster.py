# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, api_address=None, cluster_template_id=None, coe_version=None, container_version=None, create_timeout=None, created_at=None, discovery_url=None, docker_volume_size=None, flavor=None, keypair=None, labels=None, master_addresses=None, master_count=None, master_flavor=None, name=None, node_addresses=None, node_count=None, project_id=None, region=None, stack_id=None, updated_at=None, user_id=None, id=None):
        if api_address and not isinstance(api_address, str):
            raise TypeError("Expected argument 'api_address' to be a str")
        __self__.api_address = api_address
        """
        COE API address.
        """
        if cluster_template_id and not isinstance(cluster_template_id, str):
            raise TypeError("Expected argument 'cluster_template_id' to be a str")
        __self__.cluster_template_id = cluster_template_id
        """
        The UUID of the V1 Container Infra cluster template.
        """
        if coe_version and not isinstance(coe_version, str):
            raise TypeError("Expected argument 'coe_version' to be a str")
        __self__.coe_version = coe_version
        """
        COE software version.
        """
        if container_version and not isinstance(container_version, str):
            raise TypeError("Expected argument 'container_version' to be a str")
        __self__.container_version = container_version
        if create_timeout and not isinstance(create_timeout, float):
            raise TypeError("Expected argument 'create_timeout' to be a float")
        __self__.create_timeout = create_timeout
        """
        The timeout (in minutes) for creating the cluster.
        """
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        The time at which cluster was created.
        """
        if discovery_url and not isinstance(discovery_url, str):
            raise TypeError("Expected argument 'discovery_url' to be a str")
        __self__.discovery_url = discovery_url
        """
        The URL used for cluster node discovery.
        """
        if docker_volume_size and not isinstance(docker_volume_size, float):
            raise TypeError("Expected argument 'docker_volume_size' to be a float")
        __self__.docker_volume_size = docker_volume_size
        """
        The size (in GB) of the Docker volume.
        """
        if flavor and not isinstance(flavor, str):
            raise TypeError("Expected argument 'flavor' to be a str")
        __self__.flavor = flavor
        """
        The flavor for the nodes of the cluster.
        """
        if keypair and not isinstance(keypair, str):
            raise TypeError("Expected argument 'keypair' to be a str")
        __self__.keypair = keypair
        """
        The name of the Compute service SSH keypair.
        """
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        """
        The list of key value pairs representing additional properties of
        the cluster.
        """
        if master_addresses and not isinstance(master_addresses, str):
            raise TypeError("Expected argument 'master_addresses' to be a str")
        __self__.master_addresses = master_addresses
        """
        IP addresses of the master node of the cluster.
        """
        if master_count and not isinstance(master_count, float):
            raise TypeError("Expected argument 'master_count' to be a float")
        __self__.master_count = master_count
        """
        The number of master nodes for the cluster.
        """
        if master_flavor and not isinstance(master_flavor, str):
            raise TypeError("Expected argument 'master_flavor' to be a str")
        __self__.master_flavor = master_flavor
        """
        The flavor for the master nodes.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if node_addresses and not isinstance(node_addresses, str):
            raise TypeError("Expected argument 'node_addresses' to be a str")
        __self__.node_addresses = node_addresses
        """
        IP addresses of the node of the cluster.
        """
        if node_count and not isinstance(node_count, float):
            raise TypeError("Expected argument 'node_count' to be a float")
        __self__.node_count = node_count
        """
        The number of nodes for the cluster.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        """
        The project of the cluster.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        __self__.stack_id = stack_id
        """
        UUID of the Orchestration service stack.
        """
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        __self__.updated_at = updated_at
        """
        The time at which cluster was updated.
        """
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        __self__.user_id = user_id
        """
        The user of the cluster.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_cluster(name=None,region=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack Magnum cluster.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/containerinfra_cluster_v1.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['region'] = region
    __ret__ = await pulumi.runtime.invoke('openstack:containerinfra/getCluster:getCluster', __args__, opts=opts)

    return GetClusterResult(
        api_address=__ret__.get('apiAddress'),
        cluster_template_id=__ret__.get('clusterTemplateId'),
        coe_version=__ret__.get('coeVersion'),
        container_version=__ret__.get('containerVersion'),
        create_timeout=__ret__.get('createTimeout'),
        created_at=__ret__.get('createdAt'),
        discovery_url=__ret__.get('discoveryUrl'),
        docker_volume_size=__ret__.get('dockerVolumeSize'),
        flavor=__ret__.get('flavor'),
        keypair=__ret__.get('keypair'),
        labels=__ret__.get('labels'),
        master_addresses=__ret__.get('masterAddresses'),
        master_count=__ret__.get('masterCount'),
        master_flavor=__ret__.get('masterFlavor'),
        name=__ret__.get('name'),
        node_addresses=__ret__.get('nodeAddresses'),
        node_count=__ret__.get('nodeCount'),
        project_id=__ret__.get('projectId'),
        region=__ret__.get('region'),
        stack_id=__ret__.get('stackId'),
        updated_at=__ret__.get('updatedAt'),
        user_id=__ret__.get('userId'),
        id=__ret__.get('id'))
