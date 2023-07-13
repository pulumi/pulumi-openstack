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
from ._inputs import *

__all__ = ['ApplicationCredentialArgs', 'ApplicationCredential']

@pulumi.input_type
class ApplicationCredentialArgs:
    def __init__(__self__, *,
                 access_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCredentialAccessRuleArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 unrestricted: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ApplicationCredential resource.
        """
        if access_rules is not None:
            pulumi.set(__self__, "access_rules", access_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if unrestricted is not None:
            pulumi.set(__self__, "unrestricted", unrestricted)

    @property
    @pulumi.getter(name="accessRules")
    def access_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCredentialAccessRuleArgs']]]]:
        return pulumi.get(self, "access_rules")

    @access_rules.setter
    def access_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCredentialAccessRuleArgs']]]]):
        pulumi.set(self, "access_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def unrestricted(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "unrestricted")

    @unrestricted.setter
    def unrestricted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unrestricted", value)


@pulumi.input_type
class _ApplicationCredentialState:
    def __init__(__self__, *,
                 access_rules: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCredentialAccessRuleArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 unrestricted: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ApplicationCredential resources.
        """
        if access_rules is not None:
            pulumi.set(__self__, "access_rules", access_rules)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if unrestricted is not None:
            pulumi.set(__self__, "unrestricted", unrestricted)

    @property
    @pulumi.getter(name="accessRules")
    def access_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCredentialAccessRuleArgs']]]]:
        return pulumi.get(self, "access_rules")

    @access_rules.setter
    def access_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApplicationCredentialAccessRuleArgs']]]]):
        pulumi.set(self, "access_rules", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def unrestricted(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "unrestricted")

    @unrestricted.setter
    def unrestricted(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unrestricted", value)


class ApplicationCredential(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationCredentialAccessRuleArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 unrestricted: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a ApplicationCredential resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ApplicationCredentialArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApplicationCredential resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApplicationCredentialArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationCredentialArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationCredentialAccessRuleArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 unrestricted: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationCredentialArgs.__new__(ApplicationCredentialArgs)

            __props__.__dict__["access_rules"] = access_rules
            __props__.__dict__["description"] = description
            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["roles"] = roles
            __props__.__dict__["secret"] = None if secret is None else pulumi.Output.secret(secret)
            __props__.__dict__["unrestricted"] = unrestricted
            __props__.__dict__["project_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ApplicationCredential, __self__).__init__(
            'openstack:identity/applicationCredential:ApplicationCredential',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationCredentialAccessRuleArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            unrestricted: Optional[pulumi.Input[bool]] = None) -> 'ApplicationCredential':
        """
        Get an existing ApplicationCredential resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApplicationCredentialState.__new__(_ApplicationCredentialState)

        __props__.__dict__["access_rules"] = access_rules
        __props__.__dict__["description"] = description
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["roles"] = roles
        __props__.__dict__["secret"] = secret
        __props__.__dict__["unrestricted"] = unrestricted
        return ApplicationCredential(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessRules")
    def access_rules(self) -> pulumi.Output[Optional[Sequence['outputs.ApplicationCredentialAccessRule']]]:
        return pulumi.get(self, "access_rules")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def unrestricted(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "unrestricted")

