# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
]

@pulumi.output_type
class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, api_address=None, cluster_template_id=None, coe_version=None, container_version=None, create_timeout=None, created_at=None, discovery_url=None, docker_volume_size=None, fixed_network=None, fixed_subnet=None, flavor=None, floating_ip_enabled=None, id=None, keypair=None, labels=None, master_addresses=None, master_count=None, master_flavor=None, name=None, node_addresses=None, node_count=None, project_id=None, region=None, stack_id=None, updated_at=None, user_id=None):
        if api_address and not isinstance(api_address, str):
            raise TypeError("Expected argument 'api_address' to be a str")
        pulumi.set(__self__, "api_address", api_address)
        if cluster_template_id and not isinstance(cluster_template_id, str):
            raise TypeError("Expected argument 'cluster_template_id' to be a str")
        pulumi.set(__self__, "cluster_template_id", cluster_template_id)
        if coe_version and not isinstance(coe_version, str):
            raise TypeError("Expected argument 'coe_version' to be a str")
        pulumi.set(__self__, "coe_version", coe_version)
        if container_version and not isinstance(container_version, str):
            raise TypeError("Expected argument 'container_version' to be a str")
        pulumi.set(__self__, "container_version", container_version)
        if create_timeout and not isinstance(create_timeout, int):
            raise TypeError("Expected argument 'create_timeout' to be a int")
        pulumi.set(__self__, "create_timeout", create_timeout)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if discovery_url and not isinstance(discovery_url, str):
            raise TypeError("Expected argument 'discovery_url' to be a str")
        pulumi.set(__self__, "discovery_url", discovery_url)
        if docker_volume_size and not isinstance(docker_volume_size, int):
            raise TypeError("Expected argument 'docker_volume_size' to be a int")
        pulumi.set(__self__, "docker_volume_size", docker_volume_size)
        if fixed_network and not isinstance(fixed_network, str):
            raise TypeError("Expected argument 'fixed_network' to be a str")
        pulumi.set(__self__, "fixed_network", fixed_network)
        if fixed_subnet and not isinstance(fixed_subnet, str):
            raise TypeError("Expected argument 'fixed_subnet' to be a str")
        pulumi.set(__self__, "fixed_subnet", fixed_subnet)
        if flavor and not isinstance(flavor, str):
            raise TypeError("Expected argument 'flavor' to be a str")
        pulumi.set(__self__, "flavor", flavor)
        if floating_ip_enabled and not isinstance(floating_ip_enabled, bool):
            raise TypeError("Expected argument 'floating_ip_enabled' to be a bool")
        pulumi.set(__self__, "floating_ip_enabled", floating_ip_enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keypair and not isinstance(keypair, str):
            raise TypeError("Expected argument 'keypair' to be a str")
        pulumi.set(__self__, "keypair", keypair)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if master_addresses and not isinstance(master_addresses, list):
            raise TypeError("Expected argument 'master_addresses' to be a list")
        pulumi.set(__self__, "master_addresses", master_addresses)
        if master_count and not isinstance(master_count, int):
            raise TypeError("Expected argument 'master_count' to be a int")
        pulumi.set(__self__, "master_count", master_count)
        if master_flavor and not isinstance(master_flavor, str):
            raise TypeError("Expected argument 'master_flavor' to be a str")
        pulumi.set(__self__, "master_flavor", master_flavor)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_addresses and not isinstance(node_addresses, list):
            raise TypeError("Expected argument 'node_addresses' to be a list")
        pulumi.set(__self__, "node_addresses", node_addresses)
        if node_count and not isinstance(node_count, int):
            raise TypeError("Expected argument 'node_count' to be a int")
        pulumi.set(__self__, "node_count", node_count)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if stack_id and not isinstance(stack_id, str):
            raise TypeError("Expected argument 'stack_id' to be a str")
        pulumi.set(__self__, "stack_id", stack_id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="apiAddress")
    def api_address(self) -> str:
        """
        COE API address.
        """
        return pulumi.get(self, "api_address")

    @property
    @pulumi.getter(name="clusterTemplateId")
    def cluster_template_id(self) -> str:
        """
        The UUID of the V1 Container Infra cluster template.
        """
        return pulumi.get(self, "cluster_template_id")

    @property
    @pulumi.getter(name="coeVersion")
    def coe_version(self) -> str:
        """
        COE software version.
        """
        return pulumi.get(self, "coe_version")

    @property
    @pulumi.getter(name="containerVersion")
    def container_version(self) -> str:
        return pulumi.get(self, "container_version")

    @property
    @pulumi.getter(name="createTimeout")
    def create_timeout(self) -> int:
        """
        The timeout (in minutes) for creating the cluster.
        """
        return pulumi.get(self, "create_timeout")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The time at which cluster was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="discoveryUrl")
    def discovery_url(self) -> str:
        """
        The URL used for cluster node discovery.
        """
        return pulumi.get(self, "discovery_url")

    @property
    @pulumi.getter(name="dockerVolumeSize")
    def docker_volume_size(self) -> int:
        """
        The size (in GB) of the Docker volume.
        """
        return pulumi.get(self, "docker_volume_size")

    @property
    @pulumi.getter(name="fixedNetwork")
    def fixed_network(self) -> str:
        """
        The fixed network that is attached to the cluster.
        """
        return pulumi.get(self, "fixed_network")

    @property
    @pulumi.getter(name="fixedSubnet")
    def fixed_subnet(self) -> str:
        """
        The fixed subnet that is attached to the cluster.
        """
        return pulumi.get(self, "fixed_subnet")

    @property
    @pulumi.getter
    def flavor(self) -> str:
        """
        The flavor for the nodes of the cluster.
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter(name="floatingIpEnabled")
    def floating_ip_enabled(self) -> bool:
        return pulumi.get(self, "floating_ip_enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keypair(self) -> str:
        """
        The name of the Compute service SSH keypair.
        """
        return pulumi.get(self, "keypair")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, Any]:
        """
        The list of key value pairs representing additional properties of
        the cluster.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="masterAddresses")
    def master_addresses(self) -> Sequence[str]:
        """
        IP addresses of the master node of the cluster.
        """
        return pulumi.get(self, "master_addresses")

    @property
    @pulumi.getter(name="masterCount")
    def master_count(self) -> int:
        """
        The number of master nodes for the cluster.
        """
        return pulumi.get(self, "master_count")

    @property
    @pulumi.getter(name="masterFlavor")
    def master_flavor(self) -> str:
        """
        The flavor for the master nodes.
        """
        return pulumi.get(self, "master_flavor")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeAddresses")
    def node_addresses(self) -> Sequence[str]:
        """
        IP addresses of the node of the cluster.
        """
        return pulumi.get(self, "node_addresses")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> int:
        """
        The number of nodes for the cluster.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The project of the cluster.
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
    @pulumi.getter(name="stackId")
    def stack_id(self) -> str:
        """
        UUID of the Orchestration service stack.
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The time at which cluster was updated.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The user of the cluster.
        """
        return pulumi.get(self, "user_id")


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            api_address=self.api_address,
            cluster_template_id=self.cluster_template_id,
            coe_version=self.coe_version,
            container_version=self.container_version,
            create_timeout=self.create_timeout,
            created_at=self.created_at,
            discovery_url=self.discovery_url,
            docker_volume_size=self.docker_volume_size,
            fixed_network=self.fixed_network,
            fixed_subnet=self.fixed_subnet,
            flavor=self.flavor,
            floating_ip_enabled=self.floating_ip_enabled,
            id=self.id,
            keypair=self.keypair,
            labels=self.labels,
            master_addresses=self.master_addresses,
            master_count=self.master_count,
            master_flavor=self.master_flavor,
            name=self.name,
            node_addresses=self.node_addresses,
            node_count=self.node_count,
            project_id=self.project_id,
            region=self.region,
            stack_id=self.stack_id,
            updated_at=self.updated_at,
            user_id=self.user_id)


def get_cluster(name: Optional[str] = None,
                region: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Use this data source to get the ID of an available OpenStack Magnum cluster.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    cluster1 = openstack.containerinfra.get_cluster(name="cluster_1")
    ```


    :param str name: The name of the cluster.
    :param str region: The region in which to obtain the V1 Container Infra
           client.
           If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:containerinfra/getCluster:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        api_address=__ret__.api_address,
        cluster_template_id=__ret__.cluster_template_id,
        coe_version=__ret__.coe_version,
        container_version=__ret__.container_version,
        create_timeout=__ret__.create_timeout,
        created_at=__ret__.created_at,
        discovery_url=__ret__.discovery_url,
        docker_volume_size=__ret__.docker_volume_size,
        fixed_network=__ret__.fixed_network,
        fixed_subnet=__ret__.fixed_subnet,
        flavor=__ret__.flavor,
        floating_ip_enabled=__ret__.floating_ip_enabled,
        id=__ret__.id,
        keypair=__ret__.keypair,
        labels=__ret__.labels,
        master_addresses=__ret__.master_addresses,
        master_count=__ret__.master_count,
        master_flavor=__ret__.master_flavor,
        name=__ret__.name,
        node_addresses=__ret__.node_addresses,
        node_count=__ret__.node_count,
        project_id=__ret__.project_id,
        region=__ret__.region,
        stack_id=__ret__.stack_id,
        updated_at=__ret__.updated_at,
        user_id=__ret__.user_id)
