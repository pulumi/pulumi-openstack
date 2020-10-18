# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetAuthScopeResult',
    'AwaitableGetAuthScopeResult',
    'get_auth_scope',
]

@pulumi.output_type
class GetAuthScopeResult:
    """
    A collection of values returned by getAuthScope.
    """
    def __init__(__self__, domain_id=None, domain_name=None, id=None, name=None, project_domain_id=None, project_domain_name=None, project_id=None, project_name=None, region=None, roles=None, user_domain_id=None, user_domain_name=None, user_id=None, user_name=None):
        if domain_id and not isinstance(domain_id, str):
            raise TypeError("Expected argument 'domain_id' to be a str")
        pulumi.set(__self__, "domain_id", domain_id)
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_domain_id and not isinstance(project_domain_id, str):
            raise TypeError("Expected argument 'project_domain_id' to be a str")
        pulumi.set(__self__, "project_domain_id", project_domain_id)
        if project_domain_name and not isinstance(project_domain_name, str):
            raise TypeError("Expected argument 'project_domain_name' to be a str")
        pulumi.set(__self__, "project_domain_name", project_domain_name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if user_domain_id and not isinstance(user_domain_id, str):
            raise TypeError("Expected argument 'user_domain_id' to be a str")
        pulumi.set(__self__, "user_domain_id", user_domain_id)
        if user_domain_name and not isinstance(user_domain_name, str):
            raise TypeError("Expected argument 'user_domain_name' to be a str")
        pulumi.set(__self__, "user_domain_name", user_domain_name)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> str:
        """
        The domain ID of the scope.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> str:
        """
        The domain name of the scope.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectDomainId")
    def project_domain_id(self) -> str:
        """
        The domain ID of the project.
        """
        return pulumi.get(self, "project_domain_id")

    @property
    @pulumi.getter(name="projectDomainName")
    def project_domain_name(self) -> str:
        """
        The domain name of the project.
        """
        return pulumi.get(self, "project_domain_name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The project ID of the scope.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> str:
        """
        The project name of the scope.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetAuthScopeRoleResult']:
        """
        A list of roles in the current scope. See reference below.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="userDomainId")
    def user_domain_id(self) -> str:
        """
        The domain ID of the user.
        """
        return pulumi.get(self, "user_domain_id")

    @property
    @pulumi.getter(name="userDomainName")
    def user_domain_name(self) -> str:
        """
        The domain name of the user.
        """
        return pulumi.get(self, "user_domain_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The user ID the of the scope.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The username of the scope.
        """
        return pulumi.get(self, "user_name")


class AwaitableGetAuthScopeResult(GetAuthScopeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthScopeResult(
            domain_id=self.domain_id,
            domain_name=self.domain_name,
            id=self.id,
            name=self.name,
            project_domain_id=self.project_domain_id,
            project_domain_name=self.project_domain_name,
            project_id=self.project_id,
            project_name=self.project_name,
            region=self.region,
            roles=self.roles,
            user_domain_id=self.user_domain_id,
            user_domain_name=self.user_domain_name,
            user_id=self.user_id,
            user_name=self.user_name)


def get_auth_scope(name: Optional[str] = None,
                   region: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthScopeResult:
    """
    Use this data source to get authentication information about the current
    auth scope in use. This can be used as self-discovery or introspection of
    the username or project name currently in use.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    scope = openstack.identity.get_auth_scope(name="my_scope")
    ```


    :param str name: The name of the scope. This is an arbitrary name which is
           only used as a unique identifier so an actual token isn't used as the ID.
    :param str region: The region in which to obtain the V3 Identity client.
           A Identity client is needed to retrieve tokens IDs. If omitted, the
           `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:identity/getAuthScope:getAuthScope', __args__, opts=opts, typ=GetAuthScopeResult).value

    return AwaitableGetAuthScopeResult(
        domain_id=__ret__.domain_id,
        domain_name=__ret__.domain_name,
        id=__ret__.id,
        name=__ret__.name,
        project_domain_id=__ret__.project_domain_id,
        project_domain_name=__ret__.project_domain_name,
        project_id=__ret__.project_id,
        project_name=__ret__.project_name,
        region=__ret__.region,
        roles=__ret__.roles,
        user_domain_id=__ret__.user_domain_id,
        user_domain_name=__ret__.user_domain_name,
        user_id=__ret__.user_id,
        user_name=__ret__.user_name)
