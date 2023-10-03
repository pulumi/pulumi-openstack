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
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, description=None, domain_id=None, enabled=None, id=None, is_domain=None, name=None, parent_id=None, project_id=None, region=None, tags=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        pulumi.set(__self__, "domain_id", domain_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_domain and not isinstance(is_domain, bool):
            raise TypeError("Expected argument 'is_domain' to be a bool")
        pulumi.set(__self__, "is_domain", is_domain)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDomain")
    def is_domain(self) -> Optional[bool]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "is_domain")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region the project is located in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "tags")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            description=self.description,
            domain_id=self.domain_id,
            enabled=self.enabled,
            id=self.id,
            is_domain=self.is_domain,
            name=self.name,
            parent_id=self.parent_id,
            project_id=self.project_id,
            region=self.region,
            tags=self.tags)


def get_project(domain_id: Optional[str] = None,
                enabled: Optional[bool] = None,
                is_domain: Optional[bool] = None,
                name: Optional[str] = None,
                parent_id: Optional[str] = None,
                project_id: Optional[str] = None,
                region: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Use this data source to get the ID of an OpenStack project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    project1 = openstack.identity.get_project(name="demo")
    ```


    :param str domain_id: The domain this project belongs to.
    :param bool enabled: Whether the project is enabled or disabled. Valid
           values are `true` and `false`.
    :param bool is_domain: Whether this project is a domain. Valid values
           are `true` and `false`.
    :param str name: The name of the project.
    :param str parent_id: The parent of this project.
    :param str project_id: The id of the project. Conflicts with any of the
           above arguments.
    :param str region: The region the project is located in.
    """
    __args__ = dict()
    __args__['domainId'] = domain_id
    __args__['enabled'] = enabled
    __args__['isDomain'] = is_domain
    __args__['name'] = name
    __args__['parentId'] = parent_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:identity/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        description=pulumi.get(__ret__, 'description'),
        domain_id=pulumi.get(__ret__, 'domain_id'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        is_domain=pulumi.get(__ret__, 'is_domain'),
        name=pulumi.get(__ret__, 'name'),
        parent_id=pulumi.get(__ret__, 'parent_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_project)
def get_project_output(domain_id: Optional[pulumi.Input[Optional[str]]] = None,
                       enabled: Optional[pulumi.Input[Optional[bool]]] = None,
                       is_domain: Optional[pulumi.Input[Optional[bool]]] = None,
                       name: Optional[pulumi.Input[Optional[str]]] = None,
                       parent_id: Optional[pulumi.Input[Optional[str]]] = None,
                       project_id: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectResult]:
    """
    Use this data source to get the ID of an OpenStack project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    project1 = openstack.identity.get_project(name="demo")
    ```


    :param str domain_id: The domain this project belongs to.
    :param bool enabled: Whether the project is enabled or disabled. Valid
           values are `true` and `false`.
    :param bool is_domain: Whether this project is a domain. Valid values
           are `true` and `false`.
    :param str name: The name of the project.
    :param str parent_id: The parent of this project.
    :param str project_id: The id of the project. Conflicts with any of the
           above arguments.
    :param str region: The region the project is located in.
    """
    ...
