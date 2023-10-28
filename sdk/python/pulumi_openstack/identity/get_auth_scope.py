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
    'GetAuthScopeResult',
    'AwaitableGetAuthScopeResult',
    'get_auth_scope',
    'get_auth_scope_output',
]

@pulumi.output_type
class GetAuthScopeResult:
    """
    A collection of values returned by getAuthScope.
    """
    def __init__(__self__, domain_id=None, domain_name=None, id=None, name=None, project_domain_id=None, project_domain_name=None, project_id=None, project_name=None, region=None, roles=None, service_catalogs=None, set_token_id=None, token_id=None, user_domain_id=None, user_domain_name=None, user_id=None, user_name=None):
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
        if service_catalogs and not isinstance(service_catalogs, list):
            raise TypeError("Expected argument 'service_catalogs' to be a list")
        pulumi.set(__self__, "service_catalogs", service_catalogs)
        if set_token_id and not isinstance(set_token_id, bool):
            raise TypeError("Expected argument 'set_token_id' to be a bool")
        pulumi.set(__self__, "set_token_id", set_token_id)
        if token_id and not isinstance(token_id, str):
            raise TypeError("Expected argument 'token_id' to be a str")
        pulumi.set(__self__, "token_id", token_id)
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
        """
        The name of the service.
        """
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
        """
        The region of the endpoint.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetAuthScopeRoleResult']:
        """
        A list of roles in the current scope. See reference below.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="serviceCatalogs")
    def service_catalogs(self) -> Sequence['outputs.GetAuthScopeServiceCatalogResult']:
        """
        A list of service catalog entries returned with the token.
        """
        return pulumi.get(self, "service_catalogs")

    @property
    @pulumi.getter(name="setTokenId")
    def set_token_id(self) -> Optional[bool]:
        return pulumi.get(self, "set_token_id")

    @property
    @pulumi.getter(name="tokenId")
    def token_id(self) -> str:
        """
        The token ID of the scope.
        """
        return pulumi.get(self, "token_id")

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
            service_catalogs=self.service_catalogs,
            set_token_id=self.set_token_id,
            token_id=self.token_id,
            user_domain_id=self.user_domain_id,
            user_domain_name=self.user_domain_name,
            user_id=self.user_id,
            user_name=self.user_name)


def get_auth_scope(name: Optional[str] = None,
                   region: Optional[str] = None,
                   set_token_id: Optional[bool] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthScopeResult:
    """
    ## Example Usage
    ### Simple

    ```python
    import pulumi
    import pulumi_openstack as openstack

    scope = openstack.identity.get_auth_scope(name="my_scope")
    ```

    To find the the public object storage endpoint for "region1" as listed in the
    service catalog:

    ```python
    import pulumi

    object_store_service = [entry for entry in data["openstack_identity_auth_scope_v3"]["scope"]["service_catalog"] if entry["type"] == "object-store"][0]
    object_store_endpoint = [endpoint for endpoint in object_store_service["endpoints"] if endpoint["interface"] == "public" and endpoint["region"] == "region1"][0]
    object_store_public_url = object_store_endpoint["url"]
    ```
    ### In a combination with an http data source provider

    See [http](https://www.terraform.io/providers/hashicorp/http/latest/docs/data-sources/http) provider for reference.

    ```python
    import pulumi
    import pulumi_openstack as openstack

    scope = openstack.identity.get_auth_scope(name="my_scope")
    ```

    ```python
    import pulumi
    import pulumi_http as http

    object_store_service = [entry for entry in data["openstack_identity_auth_scope_v3"]["scope"]["service_catalog"] if entry["type"] == "object-store"][0]
    object_store_endpoint = [endpoint for endpoint in object_store_service["endpoints"] if endpoint["interface"] == "public" and endpoint["region"] == "region1"][0]
    object_store_public_url = object_store_endpoint["url"]
    example = http.get_http(url=object_store_public_url,
        request_headers={
            "Accept": "application/json",
            "X-Auth-Token": data["openstack_identity_auth_scope_v3"]["scope"]["token_id"],
        })
    pulumi.export("containers", example.response_body)
    ```


    :param str name: The name of the scope. This is an arbitrary name which is
           only used as a unique identifier so an actual token isn't used as the ID.
    :param str region: The region in which to obtain the V3 Identity client.
           A Identity client is needed to retrieve tokens IDs. If omitted, the
           `region` argument of the provider is used.
    :param bool set_token_id: A boolean argument that determines whether to
           export the current auth scope token ID. When set to `true`, the `token_id`
           attribute will contain an unencrypted token that can be used for further API
           calls. **Warning**: please note that the leaked token may allow unauthorized
           access to other OpenStack services within the current auth scope, so use this
           option with caution.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['region'] = region
    __args__['setTokenId'] = set_token_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:identity/getAuthScope:getAuthScope', __args__, opts=opts, typ=GetAuthScopeResult).value

    return AwaitableGetAuthScopeResult(
        domain_id=pulumi.get(__ret__, 'domain_id'),
        domain_name=pulumi.get(__ret__, 'domain_name'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_domain_id=pulumi.get(__ret__, 'project_domain_id'),
        project_domain_name=pulumi.get(__ret__, 'project_domain_name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        project_name=pulumi.get(__ret__, 'project_name'),
        region=pulumi.get(__ret__, 'region'),
        roles=pulumi.get(__ret__, 'roles'),
        service_catalogs=pulumi.get(__ret__, 'service_catalogs'),
        set_token_id=pulumi.get(__ret__, 'set_token_id'),
        token_id=pulumi.get(__ret__, 'token_id'),
        user_domain_id=pulumi.get(__ret__, 'user_domain_id'),
        user_domain_name=pulumi.get(__ret__, 'user_domain_name'),
        user_id=pulumi.get(__ret__, 'user_id'),
        user_name=pulumi.get(__ret__, 'user_name'))


@_utilities.lift_output_func(get_auth_scope)
def get_auth_scope_output(name: Optional[pulumi.Input[str]] = None,
                          region: Optional[pulumi.Input[Optional[str]]] = None,
                          set_token_id: Optional[pulumi.Input[Optional[bool]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAuthScopeResult]:
    """
    ## Example Usage
    ### Simple

    ```python
    import pulumi
    import pulumi_openstack as openstack

    scope = openstack.identity.get_auth_scope(name="my_scope")
    ```

    To find the the public object storage endpoint for "region1" as listed in the
    service catalog:

    ```python
    import pulumi

    object_store_service = [entry for entry in data["openstack_identity_auth_scope_v3"]["scope"]["service_catalog"] if entry["type"] == "object-store"][0]
    object_store_endpoint = [endpoint for endpoint in object_store_service["endpoints"] if endpoint["interface"] == "public" and endpoint["region"] == "region1"][0]
    object_store_public_url = object_store_endpoint["url"]
    ```
    ### In a combination with an http data source provider

    See [http](https://www.terraform.io/providers/hashicorp/http/latest/docs/data-sources/http) provider for reference.

    ```python
    import pulumi
    import pulumi_openstack as openstack

    scope = openstack.identity.get_auth_scope(name="my_scope")
    ```

    ```python
    import pulumi
    import pulumi_http as http

    object_store_service = [entry for entry in data["openstack_identity_auth_scope_v3"]["scope"]["service_catalog"] if entry["type"] == "object-store"][0]
    object_store_endpoint = [endpoint for endpoint in object_store_service["endpoints"] if endpoint["interface"] == "public" and endpoint["region"] == "region1"][0]
    object_store_public_url = object_store_endpoint["url"]
    example = http.get_http(url=object_store_public_url,
        request_headers={
            "Accept": "application/json",
            "X-Auth-Token": data["openstack_identity_auth_scope_v3"]["scope"]["token_id"],
        })
    pulumi.export("containers", example.response_body)
    ```


    :param str name: The name of the scope. This is an arbitrary name which is
           only used as a unique identifier so an actual token isn't used as the ID.
    :param str region: The region in which to obtain the V3 Identity client.
           A Identity client is needed to retrieve tokens IDs. If omitted, the
           `region` argument of the provider is used.
    :param bool set_token_id: A boolean argument that determines whether to
           export the current auth scope token ID. When set to `true`, the `token_id`
           attribute will contain an unencrypted token that can be used for further API
           calls. **Warning**: please note that the leaked token may allow unauthorized
           access to other OpenStack services within the current auth scope, so use this
           option with caution.
    """
    ...
