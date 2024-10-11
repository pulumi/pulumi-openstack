# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetInstanceV2Result',
    'AwaitableGetInstanceV2Result',
    'get_instance_v2',
    'get_instance_v2_output',
]

@pulumi.output_type
class GetInstanceV2Result:
    """
    A collection of values returned by getInstanceV2.
    """
    def __init__(__self__, access_ip_v4=None, access_ip_v6=None, availability_zone=None, created=None, flavor_id=None, flavor_name=None, id=None, image_id=None, image_name=None, key_pair=None, metadata=None, name=None, networks=None, power_state=None, region=None, security_groups=None, tags=None, updated=None, user_data=None):
        if access_ip_v4 and not isinstance(access_ip_v4, str):
            raise TypeError("Expected argument 'access_ip_v4' to be a str")
        pulumi.set(__self__, "access_ip_v4", access_ip_v4)
        if access_ip_v6 and not isinstance(access_ip_v6, str):
            raise TypeError("Expected argument 'access_ip_v6' to be a str")
        pulumi.set(__self__, "access_ip_v6", access_ip_v6)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if flavor_name and not isinstance(flavor_name, str):
            raise TypeError("Expected argument 'flavor_name' to be a str")
        pulumi.set(__self__, "flavor_name", flavor_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if image_name and not isinstance(image_name, str):
            raise TypeError("Expected argument 'image_name' to be a str")
        pulumi.set(__self__, "image_name", image_name)
        if key_pair and not isinstance(key_pair, str):
            raise TypeError("Expected argument 'key_pair' to be a str")
        pulumi.set(__self__, "key_pair", key_pair)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if power_state and not isinstance(power_state, str):
            raise TypeError("Expected argument 'power_state' to be a str")
        pulumi.set(__self__, "power_state", power_state)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated and not isinstance(updated, str):
            raise TypeError("Expected argument 'updated' to be a str")
        pulumi.set(__self__, "updated", updated)
        if user_data and not isinstance(user_data, str):
            raise TypeError("Expected argument 'user_data' to be a str")
        pulumi.set(__self__, "user_data", user_data)

    @property
    @pulumi.getter(name="accessIpV4")
    def access_ip_v4(self) -> str:
        """
        The first IPv4 address assigned to this server.
        """
        return pulumi.get(self, "access_ip_v4")

    @property
    @pulumi.getter(name="accessIpV6")
    def access_ip_v6(self) -> str:
        """
        The first IPv6 address assigned to this server.
        """
        return pulumi.get(self, "access_ip_v6")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The availability zone of this server.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def created(self) -> str:
        """
        The creation time of the instance.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> str:
        """
        The flavor ID used to create the server.
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="flavorName")
    def flavor_name(self) -> str:
        """
        The flavor name used to create the server.
        """
        return pulumi.get(self, "flavor_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        The image ID used to create the server.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> str:
        """
        The image name used to create the server.
        """
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="keyPair")
    def key_pair(self) -> str:
        """
        The name of the key pair assigned to this server.
        """
        return pulumi.get(self, "key_pair")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, str]:
        """
        A set of key/value pairs made available to the server.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the network
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetInstanceV2NetworkResult']:
        """
        An array of maps, detailed below.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="powerState")
    def power_state(self) -> str:
        return pulumi.get(self, "power_state")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[str]:
        """
        An array of security group names associated with this server.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        A set of string tags assigned to this server.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def updated(self) -> str:
        """
        The time when the instance was last updated.
        """
        return pulumi.get(self, "updated")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> str:
        """
        The user data added when the server was created.
        """
        return pulumi.get(self, "user_data")


class AwaitableGetInstanceV2Result(GetInstanceV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceV2Result(
            access_ip_v4=self.access_ip_v4,
            access_ip_v6=self.access_ip_v6,
            availability_zone=self.availability_zone,
            created=self.created,
            flavor_id=self.flavor_id,
            flavor_name=self.flavor_name,
            id=self.id,
            image_id=self.image_id,
            image_name=self.image_name,
            key_pair=self.key_pair,
            metadata=self.metadata,
            name=self.name,
            networks=self.networks,
            power_state=self.power_state,
            region=self.region,
            security_groups=self.security_groups,
            tags=self.tags,
            updated=self.updated,
            user_data=self.user_data)


def get_instance_v2(id: Optional[str] = None,
                    networks: Optional[Sequence[Union['GetInstanceV2NetworkArgs', 'GetInstanceV2NetworkArgsDict']]] = None,
                    region: Optional[str] = None,
                    user_data: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceV2Result:
    """
    Use this data source to get the details of a running server

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    instance = openstack.compute.get_instance_v2(id="2ba26dc6-a12d-4889-8f25-794ea5bf4453")
    ```


    :param str id: The UUID of the instance
    :param Sequence[Union['GetInstanceV2NetworkArgs', 'GetInstanceV2NetworkArgsDict']] networks: An array of maps, detailed below.
    :param str user_data: The user data added when the server was created.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['networks'] = networks
    __args__['region'] = region
    __args__['userData'] = user_data
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:compute/getInstanceV2:getInstanceV2', __args__, opts=opts, typ=GetInstanceV2Result).value

    return AwaitableGetInstanceV2Result(
        access_ip_v4=pulumi.get(__ret__, 'access_ip_v4'),
        access_ip_v6=pulumi.get(__ret__, 'access_ip_v6'),
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        created=pulumi.get(__ret__, 'created'),
        flavor_id=pulumi.get(__ret__, 'flavor_id'),
        flavor_name=pulumi.get(__ret__, 'flavor_name'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        image_name=pulumi.get(__ret__, 'image_name'),
        key_pair=pulumi.get(__ret__, 'key_pair'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        networks=pulumi.get(__ret__, 'networks'),
        power_state=pulumi.get(__ret__, 'power_state'),
        region=pulumi.get(__ret__, 'region'),
        security_groups=pulumi.get(__ret__, 'security_groups'),
        tags=pulumi.get(__ret__, 'tags'),
        updated=pulumi.get(__ret__, 'updated'),
        user_data=pulumi.get(__ret__, 'user_data'))
def get_instance_v2_output(id: Optional[pulumi.Input[str]] = None,
                           networks: Optional[pulumi.Input[Optional[Sequence[Union['GetInstanceV2NetworkArgs', 'GetInstanceV2NetworkArgsDict']]]]] = None,
                           region: Optional[pulumi.Input[Optional[str]]] = None,
                           user_data: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceV2Result]:
    """
    Use this data source to get the details of a running server

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    instance = openstack.compute.get_instance_v2(id="2ba26dc6-a12d-4889-8f25-794ea5bf4453")
    ```


    :param str id: The UUID of the instance
    :param Sequence[Union['GetInstanceV2NetworkArgs', 'GetInstanceV2NetworkArgsDict']] networks: An array of maps, detailed below.
    :param str user_data: The user data added when the server was created.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['networks'] = networks
    __args__['region'] = region
    __args__['userData'] = user_data
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:compute/getInstanceV2:getInstanceV2', __args__, opts=opts, typ=GetInstanceV2Result)
    return __ret__.apply(lambda __response__: GetInstanceV2Result(
        access_ip_v4=pulumi.get(__response__, 'access_ip_v4'),
        access_ip_v6=pulumi.get(__response__, 'access_ip_v6'),
        availability_zone=pulumi.get(__response__, 'availability_zone'),
        created=pulumi.get(__response__, 'created'),
        flavor_id=pulumi.get(__response__, 'flavor_id'),
        flavor_name=pulumi.get(__response__, 'flavor_name'),
        id=pulumi.get(__response__, 'id'),
        image_id=pulumi.get(__response__, 'image_id'),
        image_name=pulumi.get(__response__, 'image_name'),
        key_pair=pulumi.get(__response__, 'key_pair'),
        metadata=pulumi.get(__response__, 'metadata'),
        name=pulumi.get(__response__, 'name'),
        networks=pulumi.get(__response__, 'networks'),
        power_state=pulumi.get(__response__, 'power_state'),
        region=pulumi.get(__response__, 'region'),
        security_groups=pulumi.get(__response__, 'security_groups'),
        tags=pulumi.get(__response__, 'tags'),
        updated=pulumi.get(__response__, 'updated'),
        user_data=pulumi.get(__response__, 'user_data')))
