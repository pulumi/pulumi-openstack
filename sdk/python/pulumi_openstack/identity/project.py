# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Project']


class Project(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 is_domain: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V3 Project resource within OpenStack Keystone.

        > **Note:** You _must_ have admin privileges in your OpenStack cloud to use
        this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1", description="A project")
        ```

        ## Import

        Projects can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:identity/project:Project project_1 89c60255-9bd6-460c-822a-e2b959ede9d2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the project.
        :param pulumi.Input[str] domain_id: The domain this project belongs to.
        :param pulumi.Input[bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`. Default is `true`.
        :param pulumi.Input[bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`. Default is `false`. Changing this creates a new
               project/domain.
        :param pulumi.Input[str] name: The name of the project.
        :param pulumi.Input[str] parent_id: The parent of this project. Changing this creates
               a new project.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new project.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for the project. Changing this updates the existing
               project.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['domain_id'] = domain_id
            __props__['enabled'] = enabled
            __props__['is_domain'] = is_domain
            __props__['name'] = name
            __props__['parent_id'] = parent_id
            __props__['region'] = region
            __props__['tags'] = tags
        super(Project, __self__).__init__(
            'openstack:identity/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            domain_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            is_domain: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the project.
        :param pulumi.Input[str] domain_id: The domain this project belongs to.
        :param pulumi.Input[bool] enabled: Whether the project is enabled or disabled. Valid
               values are `true` and `false`. Default is `true`.
        :param pulumi.Input[bool] is_domain: Whether this project is a domain. Valid values
               are `true` and `false`. Default is `false`. Changing this creates a new
               project/domain.
        :param pulumi.Input[str] name: The name of the project.
        :param pulumi.Input[str] parent_id: The parent of this project. Changing this creates
               a new project.
        :param pulumi.Input[str] region: The region in which to obtain the V3 Keystone client.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new project.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags for the project. Changing this updates the existing
               project.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["domain_id"] = domain_id
        __props__["enabled"] = enabled
        __props__["is_domain"] = is_domain
        __props__["name"] = name
        __props__["parent_id"] = parent_id
        __props__["region"] = region
        __props__["tags"] = tags
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> pulumi.Output[str]:
        """
        The domain this project belongs to.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the project is enabled or disabled. Valid
        values are `true` and `false`. Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="isDomain")
    def is_domain(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether this project is a domain. Valid values
        are `true` and `false`. Default is `false`. Changing this creates a new
        project/domain.
        """
        return pulumi.get(self, "is_domain")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[str]:
        """
        The parent of this project. Changing this creates
        a new project.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V3 Keystone client.
        If omitted, the `region` argument of the provider is used. Changing this
        creates a new project.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Tags for the project. Changing this updates the existing
        project.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

