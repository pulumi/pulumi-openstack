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
    'ApplicationCredentialAccessRuleArgs',
    'UserMultiFactorAuthRuleArgs',
]

@pulumi.input_type
class ApplicationCredentialAccessRuleArgs:
    def __init__(__self__, *,
                 method: pulumi.Input[str],
                 path: pulumi.Input[str],
                 service: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] method: The request method that the application credential is
               permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
               `HEAD`, `PATCH`, `PUT` and `DELETE`.
        :param pulumi.Input[str] path: The API path that the application credential is permitted
               to access. May use named wildcards such as **{tag}** or the unnamed wildcard
               **\\*** to match against any string in the path up to a **/**, or the recursive
               wildcard **\\*\\*** to include **/** in the matched path.
        :param pulumi.Input[str] service: The service type identifier for the service that the
               application credential is granted to access. Must be a service type that is
               listed in the service catalog and not a code name for a service. E.g.
               **identity**, **compute**, **volumev3**, **image**, **network**,
               **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        :param pulumi.Input[str] id: The ID of the existing access rule. The access rule ID of
               another application credential can be provided.
        """
        ApplicationCredentialAccessRuleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            method=method,
            path=path,
            service=service,
            id=id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             method: pulumi.Input[str],
             path: pulumi.Input[str],
             service: pulumi.Input[str],
             id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("method", method)
        _setter("path", path)
        _setter("service", service)
        if id is not None:
            _setter("id", id)

    @property
    @pulumi.getter
    def method(self) -> pulumi.Input[str]:
        """
        The request method that the application credential is
        permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
        `HEAD`, `PATCH`, `PUT` and `DELETE`.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: pulumi.Input[str]):
        pulumi.set(self, "method", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The API path that the application credential is permitted
        to access. May use named wildcards such as **{tag}** or the unnamed wildcard
        **\\*** to match against any string in the path up to a **/**, or the recursive
        wildcard **\\*\\*** to include **/** in the matched path.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def service(self) -> pulumi.Input[str]:
        """
        The service type identifier for the service that the
        application credential is granted to access. Must be a service type that is
        listed in the service catalog and not a code name for a service. E.g.
        **identity**, **compute**, **volumev3**, **image**, **network**,
        **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: pulumi.Input[str]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the existing access rule. The access rule ID of
        another application credential can be provided.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class UserMultiFactorAuthRuleArgs:
    def __init__(__self__, *,
                 rules: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rules: A list of authentication plugins that the user must
               authenticate with.
        """
        UserMultiFactorAuthRuleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            rules=rules,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             rules: pulumi.Input[Sequence[pulumi.Input[str]]],
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of authentication plugins that the user must
        authenticate with.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "rules", value)


