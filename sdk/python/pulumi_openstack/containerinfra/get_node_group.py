# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetNodeGroupResult',
    'AwaitableGetNodeGroupResult',
    'get_node_group',
    'get_node_group_output',
]

@pulumi.output_type
class GetNodeGroupResult:
    """
    A collection of values returned by getNodeGroup.
    """
    def __init__(__self__, cluster_id=None, created_at=None, docker_volume_size=None, flavor=None, id=None, image=None, labels=None, max_node_count=None, min_node_count=None, name=None, node_count=None, project_id=None, region=None, role=None, updated_at=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if docker_volume_size and not isinstance(docker_volume_size, int):
            raise TypeError("Expected argument 'docker_volume_size' to be a int")
        pulumi.set(__self__, "docker_volume_size", docker_volume_size)
        if flavor and not isinstance(flavor, str):
            raise TypeError("Expected argument 'flavor' to be a str")
        pulumi.set(__self__, "flavor", flavor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image and not isinstance(image, str):
            raise TypeError("Expected argument 'image' to be a str")
        pulumi.set(__self__, "image", image)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if max_node_count and not isinstance(max_node_count, int):
            raise TypeError("Expected argument 'max_node_count' to be a int")
        pulumi.set(__self__, "max_node_count", max_node_count)
        if min_node_count and not isinstance(min_node_count, int):
            raise TypeError("Expected argument 'min_node_count' to be a int")
        pulumi.set(__self__, "min_node_count", min_node_count)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_count and not isinstance(node_count, int):
            raise TypeError("Expected argument 'node_count' to be a int")
        pulumi.set(__self__, "node_count", node_count)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The time at which the node group was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dockerVolumeSize")
    def docker_volume_size(self) -> int:
        """
        The size (in GB) of the Docker volume.
        """
        return pulumi.get(self, "docker_volume_size")

    @property
    @pulumi.getter
    def flavor(self) -> str:
        """
        The flavor for the nodes of the node group.
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def image(self) -> str:
        """
        The reference to an image that is used for nodes of the node group.
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, Any]:
        """
        The list of key value pairs representing additional properties of
        the node group.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maxNodeCount")
    def max_node_count(self) -> int:
        """
        The maximum number of nodes for the node group.
        """
        return pulumi.get(self, "max_node_count")

    @property
    @pulumi.getter(name="minNodeCount")
    def min_node_count(self) -> int:
        """
        The minimum number of nodes for the node group.
        """
        return pulumi.get(self, "min_node_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> int:
        """
        The number of nodes for the node group.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The project of the node group.
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
    def role(self) -> str:
        """
        The role of the node group.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The time at which the node group was updated.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetNodeGroupResult(GetNodeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodeGroupResult(
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            docker_volume_size=self.docker_volume_size,
            flavor=self.flavor,
            id=self.id,
            image=self.image,
            labels=self.labels,
            max_node_count=self.max_node_count,
            min_node_count=self.min_node_count,
            name=self.name,
            node_count=self.node_count,
            project_id=self.project_id,
            region=self.region,
            role=self.role,
            updated_at=self.updated_at)


def get_node_group(cluster_id: Optional[str] = None,
                   name: Optional[str] = None,
                   region: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNodeGroupResult:
    """
    Use this data source to get information of an available OpenStack Magnum node group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    nodegroup1 = openstack.containerinfra.get_node_group(cluster_id="cluster_1",
        name="nodegroup_1")
    ```


    :param str cluster_id: The name of the OpenStack Magnum cluster.
    :param str name: The name of the node group.
    :param str region: The region in which to obtain the V1 Container Infra
           client.
           If omitted, the `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:containerinfra/getNodeGroup:getNodeGroup', __args__, opts=opts, typ=GetNodeGroupResult).value

    return AwaitableGetNodeGroupResult(
        cluster_id=__ret__.cluster_id,
        created_at=__ret__.created_at,
        docker_volume_size=__ret__.docker_volume_size,
        flavor=__ret__.flavor,
        id=__ret__.id,
        image=__ret__.image,
        labels=__ret__.labels,
        max_node_count=__ret__.max_node_count,
        min_node_count=__ret__.min_node_count,
        name=__ret__.name,
        node_count=__ret__.node_count,
        project_id=__ret__.project_id,
        region=__ret__.region,
        role=__ret__.role,
        updated_at=__ret__.updated_at)


@_utilities.lift_output_func(get_node_group)
def get_node_group_output(cluster_id: Optional[pulumi.Input[str]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          region: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNodeGroupResult]:
    """
    Use this data source to get information of an available OpenStack Magnum node group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    nodegroup1 = openstack.containerinfra.get_node_group(cluster_id="cluster_1",
        name="nodegroup_1")
    ```


    :param str cluster_id: The name of the OpenStack Magnum cluster.
    :param str name: The name of the node group.
    :param str region: The region in which to obtain the V1 Container Infra
           client.
           If omitted, the `region` argument of the provider is used.
    """
    ...
