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

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 default_project_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 extra: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 ignore_change_password_upon_first_use: Optional[pulumi.Input[bool]] = None,
                 ignore_lockout_failure_attempts: Optional[pulumi.Input[bool]] = None,
                 ignore_password_expiry: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_rules: Optional[pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] default_project_id: The default project this user belongs to.
        :param pulumi.Input[str] description: A description of the user.
        :param pulumi.Input[str] domain_id: The domain this user belongs to.
        :param pulumi.Input[bool] enabled: Whether the user is enabled or disabled. Valid
               values are `true` and `false`.
        :param pulumi.Input[Mapping[str, Any]] extra: Free-form key/value pairs of extra information.
        :param pulumi.Input[bool] ignore_change_password_upon_first_use: User will not have to
               change their password upon first use. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_lockout_failure_attempts: User will not have a failure
               lockout placed on their account. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_password_expiry: User's password will not expire.
               Valid values are `true` and `false`.
        :param pulumi.Input[bool] multi_factor_auth_enabled: Whether to enable multi-factor
               authentication. Valid values are `true` and `false`.
        :param pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]] multi_factor_auth_rules: A multi-factor authentication rule.
               The structure is documented below. Please see the
               [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
               for more information on how to use mulit-factor rules.
        :param pulumi.Input[str] name: The name of the user.
        :param pulumi.Input[str] password: The password for the user.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new User.
        """
        if default_project_id is not None:
            pulumi.set(__self__, "default_project_id", default_project_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if extra is not None:
            pulumi.set(__self__, "extra", extra)
        if ignore_change_password_upon_first_use is not None:
            pulumi.set(__self__, "ignore_change_password_upon_first_use", ignore_change_password_upon_first_use)
        if ignore_lockout_failure_attempts is not None:
            pulumi.set(__self__, "ignore_lockout_failure_attempts", ignore_lockout_failure_attempts)
        if ignore_password_expiry is not None:
            pulumi.set(__self__, "ignore_password_expiry", ignore_password_expiry)
        if multi_factor_auth_enabled is not None:
            pulumi.set(__self__, "multi_factor_auth_enabled", multi_factor_auth_enabled)
        if multi_factor_auth_rules is not None:
            pulumi.set(__self__, "multi_factor_auth_rules", multi_factor_auth_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="defaultProjectId")
    def default_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The default project this user belongs to.
        """
        return pulumi.get(self, "default_project_id")

    @default_project_id.setter
    def default_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_project_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the user.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[str]]:
        """
        The domain this user belongs to.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the user is enabled or disabled. Valid
        values are `true` and `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def extra(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Free-form key/value pairs of extra information.
        """
        return pulumi.get(self, "extra")

    @extra.setter
    def extra(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra", value)

    @property
    @pulumi.getter(name="ignoreChangePasswordUponFirstUse")
    def ignore_change_password_upon_first_use(self) -> Optional[pulumi.Input[bool]]:
        """
        User will not have to
        change their password upon first use. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_change_password_upon_first_use")

    @ignore_change_password_upon_first_use.setter
    def ignore_change_password_upon_first_use(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_change_password_upon_first_use", value)

    @property
    @pulumi.getter(name="ignoreLockoutFailureAttempts")
    def ignore_lockout_failure_attempts(self) -> Optional[pulumi.Input[bool]]:
        """
        User will not have a failure
        lockout placed on their account. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_lockout_failure_attempts")

    @ignore_lockout_failure_attempts.setter
    def ignore_lockout_failure_attempts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_lockout_failure_attempts", value)

    @property
    @pulumi.getter(name="ignorePasswordExpiry")
    def ignore_password_expiry(self) -> Optional[pulumi.Input[bool]]:
        """
        User's password will not expire.
        Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_password_expiry")

    @ignore_password_expiry.setter
    def ignore_password_expiry(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_password_expiry", value)

    @property
    @pulumi.getter(name="multiFactorAuthEnabled")
    def multi_factor_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable multi-factor
        authentication. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "multi_factor_auth_enabled")

    @multi_factor_auth_enabled.setter
    def multi_factor_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_factor_auth_enabled", value)

    @property
    @pulumi.getter(name="multiFactorAuthRules")
    def multi_factor_auth_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]]]:
        """
        A multi-factor authentication rule.
        The structure is documented below. Please see the
        [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        for more information on how to use mulit-factor rules.
        """
        return pulumi.get(self, "multi_factor_auth_rules")

    @multi_factor_auth_rules.setter
    def multi_factor_auth_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]]]):
        pulumi.set(self, "multi_factor_auth_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new User.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 default_project_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 extra: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 ignore_change_password_upon_first_use: Optional[pulumi.Input[bool]] = None,
                 ignore_lockout_failure_attempts: Optional[pulumi.Input[bool]] = None,
                 ignore_password_expiry: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_rules: Optional[pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[str] default_project_id: The default project this user belongs to.
        :param pulumi.Input[str] description: A description of the user.
        :param pulumi.Input[str] domain_id: The domain this user belongs to.
        :param pulumi.Input[bool] enabled: Whether the user is enabled or disabled. Valid
               values are `true` and `false`.
        :param pulumi.Input[Mapping[str, Any]] extra: Free-form key/value pairs of extra information.
        :param pulumi.Input[bool] ignore_change_password_upon_first_use: User will not have to
               change their password upon first use. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_lockout_failure_attempts: User will not have a failure
               lockout placed on their account. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_password_expiry: User's password will not expire.
               Valid values are `true` and `false`.
        :param pulumi.Input[bool] multi_factor_auth_enabled: Whether to enable multi-factor
               authentication. Valid values are `true` and `false`.
        :param pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]] multi_factor_auth_rules: A multi-factor authentication rule.
               The structure is documented below. Please see the
               [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
               for more information on how to use mulit-factor rules.
        :param pulumi.Input[str] name: The name of the user.
        :param pulumi.Input[str] password: The password for the user.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new User.
        """
        if default_project_id is not None:
            pulumi.set(__self__, "default_project_id", default_project_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if extra is not None:
            pulumi.set(__self__, "extra", extra)
        if ignore_change_password_upon_first_use is not None:
            pulumi.set(__self__, "ignore_change_password_upon_first_use", ignore_change_password_upon_first_use)
        if ignore_lockout_failure_attempts is not None:
            pulumi.set(__self__, "ignore_lockout_failure_attempts", ignore_lockout_failure_attempts)
        if ignore_password_expiry is not None:
            pulumi.set(__self__, "ignore_password_expiry", ignore_password_expiry)
        if multi_factor_auth_enabled is not None:
            pulumi.set(__self__, "multi_factor_auth_enabled", multi_factor_auth_enabled)
        if multi_factor_auth_rules is not None:
            pulumi.set(__self__, "multi_factor_auth_rules", multi_factor_auth_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="defaultProjectId")
    def default_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The default project this user belongs to.
        """
        return pulumi.get(self, "default_project_id")

    @default_project_id.setter
    def default_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_project_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the user.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[str]]:
        """
        The domain this user belongs to.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the user is enabled or disabled. Valid
        values are `true` and `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def extra(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Free-form key/value pairs of extra information.
        """
        return pulumi.get(self, "extra")

    @extra.setter
    def extra(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra", value)

    @property
    @pulumi.getter(name="ignoreChangePasswordUponFirstUse")
    def ignore_change_password_upon_first_use(self) -> Optional[pulumi.Input[bool]]:
        """
        User will not have to
        change their password upon first use. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_change_password_upon_first_use")

    @ignore_change_password_upon_first_use.setter
    def ignore_change_password_upon_first_use(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_change_password_upon_first_use", value)

    @property
    @pulumi.getter(name="ignoreLockoutFailureAttempts")
    def ignore_lockout_failure_attempts(self) -> Optional[pulumi.Input[bool]]:
        """
        User will not have a failure
        lockout placed on their account. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_lockout_failure_attempts")

    @ignore_lockout_failure_attempts.setter
    def ignore_lockout_failure_attempts(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_lockout_failure_attempts", value)

    @property
    @pulumi.getter(name="ignorePasswordExpiry")
    def ignore_password_expiry(self) -> Optional[pulumi.Input[bool]]:
        """
        User's password will not expire.
        Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_password_expiry")

    @ignore_password_expiry.setter
    def ignore_password_expiry(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_password_expiry", value)

    @property
    @pulumi.getter(name="multiFactorAuthEnabled")
    def multi_factor_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable multi-factor
        authentication. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "multi_factor_auth_enabled")

    @multi_factor_auth_enabled.setter
    def multi_factor_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_factor_auth_enabled", value)

    @property
    @pulumi.getter(name="multiFactorAuthRules")
    def multi_factor_auth_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]]]:
        """
        A multi-factor authentication rule.
        The structure is documented below. Please see the
        [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        for more information on how to use mulit-factor rules.
        """
        return pulumi.get(self, "multi_factor_auth_rules")

    @multi_factor_auth_rules.setter
    def multi_factor_auth_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserMultiFactorAuthRuleArgs']]]]):
        pulumi.set(self, "multi_factor_auth_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the user.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new User.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_project_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 extra: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 ignore_change_password_upon_first_use: Optional[pulumi.Input[bool]] = None,
                 ignore_lockout_failure_attempts: Optional[pulumi.Input[bool]] = None,
                 ignore_password_expiry: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserMultiFactorAuthRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V3 User resource within OpenStack Keystone.

        > **Note:** All arguments including the user password will be stored in the
        raw state as plain-text. Read more about sensitive data in
        state.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        user1 = openstack.identity.User("user1",
            default_project_id=project1.id,
            description="A user",
            password="password123",
            ignore_change_password_upon_first_use=True,
            multi_factor_auth_enabled=True,
            multi_factor_auth_rules=[
                openstack.identity.UserMultiFactorAuthRuleArgs(
                    rules=[
                        "password",
                        "totp",
                    ],
                ),
                openstack.identity.UserMultiFactorAuthRuleArgs(
                    rules=["password"],
                ),
            ],
            extra={
                "email": "user_1@foobar.com",
            })
        ```

        ## Import

        Users can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:identity/user:User user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_project_id: The default project this user belongs to.
        :param pulumi.Input[str] description: A description of the user.
        :param pulumi.Input[str] domain_id: The domain this user belongs to.
        :param pulumi.Input[bool] enabled: Whether the user is enabled or disabled. Valid
               values are `true` and `false`.
        :param pulumi.Input[Mapping[str, Any]] extra: Free-form key/value pairs of extra information.
        :param pulumi.Input[bool] ignore_change_password_upon_first_use: User will not have to
               change their password upon first use. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_lockout_failure_attempts: User will not have a failure
               lockout placed on their account. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_password_expiry: User's password will not expire.
               Valid values are `true` and `false`.
        :param pulumi.Input[bool] multi_factor_auth_enabled: Whether to enable multi-factor
               authentication. Valid values are `true` and `false`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserMultiFactorAuthRuleArgs']]]] multi_factor_auth_rules: A multi-factor authentication rule.
               The structure is documented below. Please see the
               [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
               for more information on how to use mulit-factor rules.
        :param pulumi.Input[str] name: The name of the user.
        :param pulumi.Input[str] password: The password for the user.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new User.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V3 User resource within OpenStack Keystone.

        > **Note:** All arguments including the user password will be stored in the
        raw state as plain-text. Read more about sensitive data in
        state.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        user1 = openstack.identity.User("user1",
            default_project_id=project1.id,
            description="A user",
            password="password123",
            ignore_change_password_upon_first_use=True,
            multi_factor_auth_enabled=True,
            multi_factor_auth_rules=[
                openstack.identity.UserMultiFactorAuthRuleArgs(
                    rules=[
                        "password",
                        "totp",
                    ],
                ),
                openstack.identity.UserMultiFactorAuthRuleArgs(
                    rules=["password"],
                ),
            ],
            extra={
                "email": "user_1@foobar.com",
            })
        ```

        ## Import

        Users can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:identity/user:User user_1 89c60255-9bd6-460c-822a-e2b959ede9d2
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_project_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 extra: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 ignore_change_password_upon_first_use: Optional[pulumi.Input[bool]] = None,
                 ignore_lockout_failure_attempts: Optional[pulumi.Input[bool]] = None,
                 ignore_password_expiry: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 multi_factor_auth_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserMultiFactorAuthRuleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["default_project_id"] = default_project_id
            __props__.__dict__["description"] = description
            __props__.__dict__["domain_id"] = domain_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["extra"] = extra
            __props__.__dict__["ignore_change_password_upon_first_use"] = ignore_change_password_upon_first_use
            __props__.__dict__["ignore_lockout_failure_attempts"] = ignore_lockout_failure_attempts
            __props__.__dict__["ignore_password_expiry"] = ignore_password_expiry
            __props__.__dict__["multi_factor_auth_enabled"] = multi_factor_auth_enabled
            __props__.__dict__["multi_factor_auth_rules"] = multi_factor_auth_rules
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["region"] = region
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(User, __self__).__init__(
            'openstack:identity/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_project_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            domain_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            extra: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            ignore_change_password_upon_first_use: Optional[pulumi.Input[bool]] = None,
            ignore_lockout_failure_attempts: Optional[pulumi.Input[bool]] = None,
            ignore_password_expiry: Optional[pulumi.Input[bool]] = None,
            multi_factor_auth_enabled: Optional[pulumi.Input[bool]] = None,
            multi_factor_auth_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserMultiFactorAuthRuleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_project_id: The default project this user belongs to.
        :param pulumi.Input[str] description: A description of the user.
        :param pulumi.Input[str] domain_id: The domain this user belongs to.
        :param pulumi.Input[bool] enabled: Whether the user is enabled or disabled. Valid
               values are `true` and `false`.
        :param pulumi.Input[Mapping[str, Any]] extra: Free-form key/value pairs of extra information.
        :param pulumi.Input[bool] ignore_change_password_upon_first_use: User will not have to
               change their password upon first use. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_lockout_failure_attempts: User will not have a failure
               lockout placed on their account. Valid values are `true` and `false`.
        :param pulumi.Input[bool] ignore_password_expiry: User's password will not expire.
               Valid values are `true` and `false`.
        :param pulumi.Input[bool] multi_factor_auth_enabled: Whether to enable multi-factor
               authentication. Valid values are `true` and `false`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserMultiFactorAuthRuleArgs']]]] multi_factor_auth_rules: A multi-factor authentication rule.
               The structure is documented below. Please see the
               [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
               for more information on how to use mulit-factor rules.
        :param pulumi.Input[str] name: The name of the user.
        :param pulumi.Input[str] password: The password for the user.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new User.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["default_project_id"] = default_project_id
        __props__.__dict__["description"] = description
        __props__.__dict__["domain_id"] = domain_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["extra"] = extra
        __props__.__dict__["ignore_change_password_upon_first_use"] = ignore_change_password_upon_first_use
        __props__.__dict__["ignore_lockout_failure_attempts"] = ignore_lockout_failure_attempts
        __props__.__dict__["ignore_password_expiry"] = ignore_password_expiry
        __props__.__dict__["multi_factor_auth_enabled"] = multi_factor_auth_enabled
        __props__.__dict__["multi_factor_auth_rules"] = multi_factor_auth_rules
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["region"] = region
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultProjectId")
    def default_project_id(self) -> pulumi.Output[str]:
        """
        The default project this user belongs to.
        """
        return pulumi.get(self, "default_project_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the user.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[str]:
        """
        The domain this user belongs to.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user is enabled or disabled. Valid
        values are `true` and `false`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def extra(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Free-form key/value pairs of extra information.
        """
        return pulumi.get(self, "extra")

    @property
    @pulumi.getter(name="ignoreChangePasswordUponFirstUse")
    def ignore_change_password_upon_first_use(self) -> pulumi.Output[Optional[bool]]:
        """
        User will not have to
        change their password upon first use. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_change_password_upon_first_use")

    @property
    @pulumi.getter(name="ignoreLockoutFailureAttempts")
    def ignore_lockout_failure_attempts(self) -> pulumi.Output[Optional[bool]]:
        """
        User will not have a failure
        lockout placed on their account. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_lockout_failure_attempts")

    @property
    @pulumi.getter(name="ignorePasswordExpiry")
    def ignore_password_expiry(self) -> pulumi.Output[Optional[bool]]:
        """
        User's password will not expire.
        Valid values are `true` and `false`.
        """
        return pulumi.get(self, "ignore_password_expiry")

    @property
    @pulumi.getter(name="multiFactorAuthEnabled")
    def multi_factor_auth_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable multi-factor
        authentication. Valid values are `true` and `false`.
        """
        return pulumi.get(self, "multi_factor_auth_enabled")

    @property
    @pulumi.getter(name="multiFactorAuthRules")
    def multi_factor_auth_rules(self) -> pulumi.Output[Optional[Sequence['outputs.UserMultiFactorAuthRule']]]:
        """
        A multi-factor authentication rule.
        The structure is documented below. Please see the
        [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html)
        for more information on how to use mulit-factor rules.
        """
        return pulumi.get(self, "multi_factor_auth_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password for the user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new User.
        """
        return pulumi.get(self, "region")

