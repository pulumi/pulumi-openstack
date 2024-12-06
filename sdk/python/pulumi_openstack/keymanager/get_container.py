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

__all__ = [
    'GetContainerResult',
    'AwaitableGetContainerResult',
    'get_container',
    'get_container_output',
]

@pulumi.output_type
class GetContainerResult:
    """
    A collection of values returned by getContainer.
    """
    def __init__(__self__, acls=None, consumers=None, container_ref=None, created_at=None, creator_id=None, id=None, name=None, region=None, secret_refs=None, status=None, type=None, updated_at=None):
        if acls and not isinstance(acls, list):
            raise TypeError("Expected argument 'acls' to be a list")
        pulumi.set(__self__, "acls", acls)
        if consumers and not isinstance(consumers, list):
            raise TypeError("Expected argument 'consumers' to be a list")
        pulumi.set(__self__, "consumers", consumers)
        if container_ref and not isinstance(container_ref, str):
            raise TypeError("Expected argument 'container_ref' to be a str")
        pulumi.set(__self__, "container_ref", container_ref)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if creator_id and not isinstance(creator_id, str):
            raise TypeError("Expected argument 'creator_id' to be a str")
        pulumi.set(__self__, "creator_id", creator_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if secret_refs and not isinstance(secret_refs, list):
            raise TypeError("Expected argument 'secret_refs' to be a list")
        pulumi.set(__self__, "secret_refs", secret_refs)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def acls(self) -> Sequence['outputs.GetContainerAclResult']:
        """
        The list of ACLs assigned to a container. The `read` structure is
        described below.
        """
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter
    def consumers(self) -> Sequence['outputs.GetContainerConsumerResult']:
        """
        The list of the container consumers. The structure is described
        below.
        """
        return pulumi.get(self, "consumers")

    @property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> str:
        """
        The container reference / where to find the container.
        """
        return pulumi.get(self, "container_ref")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date the container ACL was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> str:
        """
        The creator of the container.
        """
        return pulumi.get(self, "creator_id")

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
        """
        The name of the consumer.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> Sequence['outputs.GetContainerSecretRefResult']:
        """
        A set of dictionaries containing references to secrets. The
        structure is described below.
        """
        return pulumi.get(self, "secret_refs")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the container.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The container type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date the container ACL was last updated.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetContainerResult(GetContainerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerResult(
            acls=self.acls,
            consumers=self.consumers,
            container_ref=self.container_ref,
            created_at=self.created_at,
            creator_id=self.creator_id,
            id=self.id,
            name=self.name,
            region=self.region,
            secret_refs=self.secret_refs,
            status=self.status,
            type=self.type,
            updated_at=self.updated_at)


def get_container(name: Optional[str] = None,
                  region: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerResult:
    """
    Use this data source to get the ID of an available Barbican container.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    example = openstack.keymanager.get_container(name="my_container")
    ```


    :param str name: The Container name.
    :param str region: The region in which to obtain the V1 KeyManager client.
           A KeyManager client is needed to fetch a container. If omitted, the `region`
           argument of the provider is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:keymanager/getContainer:getContainer', __args__, opts=opts, typ=GetContainerResult).value

    return AwaitableGetContainerResult(
        acls=pulumi.get(__ret__, 'acls'),
        consumers=pulumi.get(__ret__, 'consumers'),
        container_ref=pulumi.get(__ret__, 'container_ref'),
        created_at=pulumi.get(__ret__, 'created_at'),
        creator_id=pulumi.get(__ret__, 'creator_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        secret_refs=pulumi.get(__ret__, 'secret_refs'),
        status=pulumi.get(__ret__, 'status'),
        type=pulumi.get(__ret__, 'type'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_container_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetContainerResult]:
    """
    Use this data source to get the ID of an available Barbican container.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    example = openstack.keymanager.get_container(name="my_container")
    ```


    :param str name: The Container name.
    :param str region: The region in which to obtain the V1 KeyManager client.
           A KeyManager client is needed to fetch a container. If omitted, the `region`
           argument of the provider is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:keymanager/getContainer:getContainer', __args__, opts=opts, typ=GetContainerResult)
    return __ret__.apply(lambda __response__: GetContainerResult(
        acls=pulumi.get(__response__, 'acls'),
        consumers=pulumi.get(__response__, 'consumers'),
        container_ref=pulumi.get(__response__, 'container_ref'),
        created_at=pulumi.get(__response__, 'created_at'),
        creator_id=pulumi.get(__response__, 'creator_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        secret_refs=pulumi.get(__response__, 'secret_refs'),
        status=pulumi.get(__response__, 'status'),
        type=pulumi.get(__response__, 'type'),
        updated_at=pulumi.get(__response__, 'updated_at')))
