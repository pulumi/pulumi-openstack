# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SubnetPool']


class SubnetPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_scope_id: Optional[pulumi.Input[str]] = None,
                 default_prefixlen: Optional[pulumi.Input[int]] = None,
                 default_quota: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[int]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 max_prefixlen: Optional[pulumi.Input[int]] = None,
                 min_prefixlen: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V2 Neutron subnetpool resource within OpenStack.

        ## Example Usage
        ### Create a Subnet Pool

        ```python
        import pulumi
        import pulumi_openstack as openstack

        subnetpool1 = openstack.networking.SubnetPool("subnetpool1",
            ip_version=6,
            prefixes=[
                "fdf7:b13d:dead:beef::/64",
                "fd65:86cc:a334:39b7::/64",
            ])
        ```
        ### Create a Subnet from a Subnet Pool

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        subnetpool1 = openstack.networking.SubnetPool("subnetpool1", prefixes=["10.11.12.0/24"])
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="10.11.12.0/25",
            network_id=network1.id,
            subnetpool_id=subnetpool1.id)
        ```

        ## Import

        Subnetpools can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:networking/subnetPool:SubnetPool subnetpool_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_scope_id: The Neutron address scope to assign to the
               subnetpool. Changing this updates the address scope id of the existing
               subnetpool.
        :param pulumi.Input[int] default_prefixlen: The size of the prefix to allocate when the cidr
               or prefixlen attributes are omitted when you create the subnet. Defaults to the
               MinPrefixLen. Changing this updates the default prefixlen of the existing
               subnetpool.
        :param pulumi.Input[int] default_quota: The per-project quota on the prefix space that can be
               allocated from the subnetpool for project subnets. Changing this updates the
               default quota of the existing subnetpool.
        :param pulumi.Input[str] description: The human-readable description for the subnetpool.
               Changing this updates the description of the existing subnetpool.
        :param pulumi.Input[int] ip_version: The IP protocol version.
        :param pulumi.Input[bool] is_default: Indicates whether the subnetpool is default
               subnetpool or not. Changing this updates the default status of the existing
               subnetpool.
        :param pulumi.Input[int] max_prefixlen: The maximum prefix size that can be allocated from
               the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
               default is 128. Changing this updates the max prefixlen of the existing
               subnetpool.
        :param pulumi.Input[int] min_prefixlen: The smallest prefix that can be allocated from a
               subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
               is 64. Changing this updates the min prefixlen of the existing subnetpool.
        :param pulumi.Input[str] name: The name of the subnetpool. Changing this updates the name of
               the existing subnetpool.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] prefixes: A list of subnet prefixes to assign to the subnetpool.
               Neutron API merges adjacent prefixes and treats them as a single prefix. Each
               subnet prefix must be unique among all subnet prefixes in all subnetpools that
               are associated with the address scope. Changing this updates the prefixes list
               of the existing subnetpool.
        :param pulumi.Input[str] project_id: The owner of the subnetpool. Required if admin wants to
               create a subnetpool for another project. Changing this creates a new subnetpool.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron subnetpool. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               subnetpool.
        :param pulumi.Input[bool] shared: Indicates whether this subnetpool is shared across
               all projects. Changing this updates the shared status of the existing
               subnetpool.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the subnetpool.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
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

            __props__['address_scope_id'] = address_scope_id
            __props__['default_prefixlen'] = default_prefixlen
            __props__['default_quota'] = default_quota
            __props__['description'] = description
            __props__['ip_version'] = ip_version
            __props__['is_default'] = is_default
            __props__['max_prefixlen'] = max_prefixlen
            __props__['min_prefixlen'] = min_prefixlen
            __props__['name'] = name
            if prefixes is None and not opts.urn:
                raise TypeError("Missing required property 'prefixes'")
            __props__['prefixes'] = prefixes
            __props__['project_id'] = project_id
            __props__['region'] = region
            __props__['shared'] = shared
            __props__['tags'] = tags
            __props__['value_specs'] = value_specs
            __props__['all_tags'] = None
            __props__['created_at'] = None
            __props__['revision_number'] = None
            __props__['updated_at'] = None
        super(SubnetPool, __self__).__init__(
            'openstack:networking/subnetPool:SubnetPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_scope_id: Optional[pulumi.Input[str]] = None,
            all_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            default_prefixlen: Optional[pulumi.Input[int]] = None,
            default_quota: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_version: Optional[pulumi.Input[int]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            max_prefixlen: Optional[pulumi.Input[int]] = None,
            min_prefixlen: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            revision_number: Optional[pulumi.Input[int]] = None,
            shared: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'SubnetPool':
        """
        Get an existing SubnetPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_scope_id: The Neutron address scope to assign to the
               subnetpool. Changing this updates the address scope id of the existing
               subnetpool.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] all_tags: The collection of tags assigned on the subnetpool, which have been
               explicitly and implicitly added.
        :param pulumi.Input[str] created_at: The time at which subnetpool was created.
        :param pulumi.Input[int] default_prefixlen: The size of the prefix to allocate when the cidr
               or prefixlen attributes are omitted when you create the subnet. Defaults to the
               MinPrefixLen. Changing this updates the default prefixlen of the existing
               subnetpool.
        :param pulumi.Input[int] default_quota: The per-project quota on the prefix space that can be
               allocated from the subnetpool for project subnets. Changing this updates the
               default quota of the existing subnetpool.
        :param pulumi.Input[str] description: The human-readable description for the subnetpool.
               Changing this updates the description of the existing subnetpool.
        :param pulumi.Input[int] ip_version: The IP protocol version.
        :param pulumi.Input[bool] is_default: Indicates whether the subnetpool is default
               subnetpool or not. Changing this updates the default status of the existing
               subnetpool.
        :param pulumi.Input[int] max_prefixlen: The maximum prefix size that can be allocated from
               the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
               default is 128. Changing this updates the max prefixlen of the existing
               subnetpool.
        :param pulumi.Input[int] min_prefixlen: The smallest prefix that can be allocated from a
               subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
               is 64. Changing this updates the min prefixlen of the existing subnetpool.
        :param pulumi.Input[str] name: The name of the subnetpool. Changing this updates the name of
               the existing subnetpool.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] prefixes: A list of subnet prefixes to assign to the subnetpool.
               Neutron API merges adjacent prefixes and treats them as a single prefix. Each
               subnet prefix must be unique among all subnet prefixes in all subnetpools that
               are associated with the address scope. Changing this updates the prefixes list
               of the existing subnetpool.
        :param pulumi.Input[str] project_id: The owner of the subnetpool. Required if admin wants to
               create a subnetpool for another project. Changing this creates a new subnetpool.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron subnetpool. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               subnetpool.
        :param pulumi.Input[int] revision_number: The revision number of the subnetpool.
        :param pulumi.Input[bool] shared: Indicates whether this subnetpool is shared across
               all projects. Changing this updates the shared status of the existing
               subnetpool.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the subnetpool.
        :param pulumi.Input[str] updated_at: The time at which subnetpool was created.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_scope_id"] = address_scope_id
        __props__["all_tags"] = all_tags
        __props__["created_at"] = created_at
        __props__["default_prefixlen"] = default_prefixlen
        __props__["default_quota"] = default_quota
        __props__["description"] = description
        __props__["ip_version"] = ip_version
        __props__["is_default"] = is_default
        __props__["max_prefixlen"] = max_prefixlen
        __props__["min_prefixlen"] = min_prefixlen
        __props__["name"] = name
        __props__["prefixes"] = prefixes
        __props__["project_id"] = project_id
        __props__["region"] = region
        __props__["revision_number"] = revision_number
        __props__["shared"] = shared
        __props__["tags"] = tags
        __props__["updated_at"] = updated_at
        __props__["value_specs"] = value_specs
        return SubnetPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressScopeId")
    def address_scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Neutron address scope to assign to the
        subnetpool. Changing this updates the address scope id of the existing
        subnetpool.
        """
        return pulumi.get(self, "address_scope_id")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> pulumi.Output[Sequence[str]]:
        """
        The collection of tags assigned on the subnetpool, which have been
        explicitly and implicitly added.
        """
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The time at which subnetpool was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultPrefixlen")
    def default_prefixlen(self) -> pulumi.Output[int]:
        """
        The size of the prefix to allocate when the cidr
        or prefixlen attributes are omitted when you create the subnet. Defaults to the
        MinPrefixLen. Changing this updates the default prefixlen of the existing
        subnetpool.
        """
        return pulumi.get(self, "default_prefixlen")

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> pulumi.Output[Optional[int]]:
        """
        The per-project quota on the prefix space that can be
        allocated from the subnetpool for project subnets. Changing this updates the
        default quota of the existing subnetpool.
        """
        return pulumi.get(self, "default_quota")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The human-readable description for the subnetpool.
        Changing this updates the description of the existing subnetpool.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[int]:
        """
        The IP protocol version.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the subnetpool is default
        subnetpool or not. Changing this updates the default status of the existing
        subnetpool.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="maxPrefixlen")
    def max_prefixlen(self) -> pulumi.Output[int]:
        """
        The maximum prefix size that can be allocated from
        the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
        default is 128. Changing this updates the max prefixlen of the existing
        subnetpool.
        """
        return pulumi.get(self, "max_prefixlen")

    @property
    @pulumi.getter(name="minPrefixlen")
    def min_prefixlen(self) -> pulumi.Output[int]:
        """
        The smallest prefix that can be allocated from a
        subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
        is 64. Changing this updates the min prefixlen of the existing subnetpool.
        """
        return pulumi.get(self, "min_prefixlen")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the subnetpool. Changing this updates the name of
        the existing subnetpool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def prefixes(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of subnet prefixes to assign to the subnetpool.
        Neutron API merges adjacent prefixes and treats them as a single prefix. Each
        subnet prefix must be unique among all subnet prefixes in all subnetpools that
        are associated with the address scope. Changing this updates the prefixes list
        of the existing subnetpool.
        """
        return pulumi.get(self, "prefixes")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The owner of the subnetpool. Required if admin wants to
        create a subnetpool for another project. Changing this creates a new subnetpool.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron subnetpool. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        subnetpool.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="revisionNumber")
    def revision_number(self) -> pulumi.Output[int]:
        """
        The revision number of the subnetpool.
        """
        return pulumi.get(self, "revision_number")

    @property
    @pulumi.getter
    def shared(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this subnetpool is shared across
        all projects. Changing this updates the shared status of the existing
        subnetpool.
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of string tags for the subnetpool.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The time at which subnetpool was created.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

