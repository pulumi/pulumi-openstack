# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_id: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 is_domain: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parent_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[builtins.str] description: A description of the project.
        :param pulumi.Input[builtins.str] domain_id: The domain this project belongs to.
        :param pulumi.Input[builtins.bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`. Default is `true`.
        :param pulumi.Input[builtins.bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`. Default is `false`. Changing this creates a new
               project/domain.
        :param pulumi.Input[builtins.str] name: The name of the project.
        :param pulumi.Input[builtins.str] parent_id: The parent of this project. Changing this creates
               a new project.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new project.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags for the project. Changing this updates the existing
               project.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if is_domain is not None:
            pulumi.set(__self__, "is_domain", is_domain)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description of the project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The domain this project belongs to.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the project is enabled or disabled. Valid
        values are `true` and `false`. Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="isDomain")
    def is_domain(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether this project is a domain. Valid values
        are `true` and `false`. Default is `false`. Changing this creates a new
        project/domain.
        """
        return pulumi.get(self, "is_domain")

    @is_domain.setter
    def is_domain(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_domain", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The parent of this project. Changing this creates
        a new project.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new project.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Tags for the project. Changing this updates the existing
        project.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _ProjectState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_id: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 is_domain: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parent_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering Project resources.
        :param pulumi.Input[builtins.str] description: A description of the project.
        :param pulumi.Input[builtins.str] domain_id: The domain this project belongs to.
        :param pulumi.Input[builtins.bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`. Default is `true`.
        :param pulumi.Input[builtins.bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`. Default is `false`. Changing this creates a new
               project/domain.
        :param pulumi.Input[builtins.str] name: The name of the project.
        :param pulumi.Input[builtins.str] parent_id: The parent of this project. Changing this creates
               a new project.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new project.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags for the project. Changing this updates the existing
               project.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if is_domain is not None:
            pulumi.set(__self__, "is_domain", is_domain)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A description of the project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The domain this project belongs to.
        """
        return pulumi.get(self, "domain_id")

    @domain_id.setter
    def domain_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the project is enabled or disabled. Valid
        values are `true` and `false`. Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="isDomain")
    def is_domain(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether this project is a domain. Valid values
        are `true` and `false`. Default is `false`. Changing this creates a new
        project/domain.
        """
        return pulumi.get(self, "is_domain")

    @is_domain.setter
    def is_domain(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_domain", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The parent of this project. Changing this creates
        a new project.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new project.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Tags for the project. Changing this updates the existing
        project.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


class Project(pulumi.CustomResource):

    pulumi_type = "openstack:identity/project:Project"

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_id: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 is_domain: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parent_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Manages a V3 Project resource within OpenStack Keystone.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project_1",
            name="project_1",
            description="A project")
        ```

        ## Import

        Projects can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:identity/project:Project project_1 89c60255-9bd6-460c-822a-e2b959ede9d2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: A description of the project.
        :param pulumi.Input[builtins.str] domain_id: The domain this project belongs to.
        :param pulumi.Input[builtins.bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`. Default is `true`.
        :param pulumi.Input[builtins.bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`. Default is `false`. Changing this creates a new
               project/domain.
        :param pulumi.Input[builtins.str] name: The name of the project.
        :param pulumi.Input[builtins.str] parent_id: The parent of this project. Changing this creates
               a new project.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new project.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags for the project. Changing this updates the existing
               project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProjectArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V3 Project resource within OpenStack Keystone.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project_1",
            name="project_1",
            description="A project")
        ```

        ## Import

        Projects can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:identity/project:Project project_1 89c60255-9bd6-460c-822a-e2b959ede9d2
        ```

        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain_id: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 is_domain: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 parent_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["domain_id"] = domain_id
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["is_domain"] = is_domain
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_id"] = parent_id
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
        super(Project, __self__).__init__(
            'openstack:identity/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            domain_id: Optional[pulumi.Input[builtins.str]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            is_domain: Optional[pulumi.Input[builtins.bool]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            parent_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: A description of the project.
        :param pulumi.Input[builtins.str] domain_id: The domain this project belongs to.
        :param pulumi.Input[builtins.bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`. Default is `true`.
        :param pulumi.Input[builtins.bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`. Default is `false`. Changing this creates a new
               project/domain.
        :param pulumi.Input[builtins.str] name: The name of the project.
        :param pulumi.Input[builtins.str] parent_id: The parent of this project. Changing this creates
               a new project.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new project.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags for the project. Changing this updates the existing
               project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectState.__new__(_ProjectState)

        __props__.__dict__["description"] = description
        __props__.__dict__["domain_id"] = domain_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["is_domain"] = is_domain
        __props__.__dict__["name"] = name
        __props__.__dict__["parent_id"] = parent_id
        __props__.__dict__["region"] = region
        __props__.__dict__["tags"] = tags
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A description of the project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[builtins.str]:
        """
        The domain this project belongs to.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Whether the project is enabled or disabled. Valid
        values are `true` and `false`. Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="isDomain")
    def is_domain(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Whether this project is a domain. Valid values
        are `true` and `false`. Default is `false`. Changing this creates a new
        project/domain.
        """
        return pulumi.get(self, "is_domain")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[builtins.str]:
        """
        The parent of this project. Changing this creates
        a new project.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new project.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Tags for the project. Changing this updates the existing
        project.
        """
        return pulumi.get(self, "tags")

