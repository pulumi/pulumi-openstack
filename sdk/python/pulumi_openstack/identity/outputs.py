# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'ApplicationCredentialAccessRule',
    'UserMultiFactorAuthRule',
    'GetAuthScopeRoleResult',
    'GetAuthScopeServiceCatalogResult',
    'GetAuthScopeServiceCatalogEndpointResult',
]

@pulumi.output_type
class ApplicationCredentialAccessRule(dict):
    def __init__(__self__, *,
                 method: str,
                 path: str,
                 service: str,
                 id: Optional[str] = None):
        """
        :param str method: The request method that the application credential is
               permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
               `HEAD`, `PATCH`, `PUT` and `DELETE`.
        :param str path: The API path that the application credential is permitted
               to access. May use named wildcards such as **{tag}** or the unnamed wildcard
               **\\*** to match against any string in the path up to a **/**, or the recursive
               wildcard **\\*\\*** to include **/** in the matched path.
        :param str service: The service type identifier for the service that the
               application credential is granted to access. Must be a service type that is
               listed in the service catalog and not a code name for a service. E.g.
               **identity**, **compute**, **volumev3**, **image**, **network**,
               **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        :param str id: The ID of the existing access rule. The access rule ID of
               another application credential can be provided.
        """
        ApplicationCredentialAccessRule._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            method=method,
            path=path,
            service=service,
            id=id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             method: str,
             path: str,
             service: str,
             id: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("method", method)
        _setter("path", path)
        _setter("service", service)
        if id is not None:
            _setter("id", id)

    @property
    @pulumi.getter
    def method(self) -> str:
        """
        The request method that the application credential is
        permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
        `HEAD`, `PATCH`, `PUT` and `DELETE`.
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        The API path that the application credential is permitted
        to access. May use named wildcards such as **{tag}** or the unnamed wildcard
        **\\*** to match against any string in the path up to a **/**, or the recursive
        wildcard **\\*\\*** to include **/** in the matched path.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        The service type identifier for the service that the
        application credential is granted to access. Must be a service type that is
        listed in the service catalog and not a code name for a service. E.g.
        **identity**, **compute**, **volumev3**, **image**, **network**,
        **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the existing access rule. The access rule ID of
        another application credential can be provided.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class UserMultiFactorAuthRule(dict):
    def __init__(__self__, *,
                 rules: Sequence[str]):
        """
        :param Sequence[str] rules: A list of authentication plugins that the user must
               authenticate with.
        """
        UserMultiFactorAuthRule._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            rules=rules,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             rules: Sequence[str],
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Sequence[str]:
        """
        A list of authentication plugins that the user must
        authenticate with.
        """
        return pulumi.get(self, "rules")


@pulumi.output_type
class GetAuthScopeRoleResult(dict):
    def __init__(__self__, *,
                 role_id: str,
                 role_name: str):
        """
        :param str role_id: The ID of the role.
        :param str role_name: The name of the role.
        """
        GetAuthScopeRoleResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            role_id=role_id,
            role_name=role_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             role_id: str,
             role_name: str,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("role_id", role_id)
        _setter("role_name", role_name)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> str:
        """
        The ID of the role.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        """
        The name of the role.
        """
        return pulumi.get(self, "role_name")


@pulumi.output_type
class GetAuthScopeServiceCatalogResult(dict):
    def __init__(__self__, *,
                 endpoints: Sequence['outputs.GetAuthScopeServiceCatalogEndpointResult'],
                 id: str,
                 name: str,
                 type: str):
        """
        :param Sequence['GetAuthScopeServiceCatalogEndpointArgs'] endpoints: A list of endpoints for the service.
        :param str id: The ID of the endpoint.
        :param str name: The name of the scope. This is an arbitrary name which is
               only used as a unique identifier so an actual token isn't used as the ID.
        :param str type: The type of the service.
        """
        GetAuthScopeServiceCatalogResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            endpoints=endpoints,
            id=id,
            name=name,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             endpoints: Sequence['outputs.GetAuthScopeServiceCatalogEndpointResult'],
             id: str,
             name: str,
             type: str,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("endpoints", endpoints)
        _setter("id", id)
        _setter("name", name)
        _setter("type", type)

    @property
    @pulumi.getter
    def endpoints(self) -> Sequence['outputs.GetAuthScopeServiceCatalogEndpointResult']:
        """
        A list of endpoints for the service.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the endpoint.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the scope. This is an arbitrary name which is
        only used as a unique identifier so an actual token isn't used as the ID.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the service.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetAuthScopeServiceCatalogEndpointResult(dict):
    def __init__(__self__, *,
                 id: str,
                 interface: str,
                 region: str,
                 region_id: str,
                 url: str):
        """
        :param str id: The ID of the endpoint.
        :param str interface: The interface of the endpoint.
        :param str region: The region in which to obtain the V3 Identity client.
               A Identity client is needed to retrieve tokens IDs. If omitted, the
               `region` argument of the provider is used.
        :param str region_id: The region ID of the endpoint.
        :param str url: The URL of the endpoint.
        """
        GetAuthScopeServiceCatalogEndpointResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            id=id,
            interface=interface,
            region=region,
            region_id=region_id,
            url=url,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             id: str,
             interface: str,
             region: str,
             region_id: str,
             url: str,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("id", id)
        _setter("interface", interface)
        _setter("region", region)
        _setter("region_id", region_id)
        _setter("url", url)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the endpoint.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        The interface of the endpoint.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region in which to obtain the V3 Identity client.
        A Identity client is needed to retrieve tokens IDs. If omitted, the
        `region` argument of the provider is used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        """
        The region ID of the endpoint.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The URL of the endpoint.
        """
        return pulumi.get(self, "url")


