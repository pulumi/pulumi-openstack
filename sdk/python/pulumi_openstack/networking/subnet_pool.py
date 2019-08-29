# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SubnetPool(pulumi.CustomResource):
    address_scope_id: pulumi.Output[str]
    """
    The Neutron address scope to assign to the
    subnetpool. Changing this updates the address scope id of the existing
    subnetpool.
    """
    all_tags: pulumi.Output[list]
    """
    The collection of tags assigned on the subnetpool, which have been
    explicitly and implicitly added.
    """
    created_at: pulumi.Output[str]
    """
    The time at which subnetpool was created.
    """
    default_prefixlen: pulumi.Output[float]
    """
    The size of the prefix to allocate when the cidr
    or prefixlen attributes are omitted when you create the subnet. Defaults to the
    MinPrefixLen. Changing this updates the default prefixlen of the existing
    subnetpool.
    """
    default_quota: pulumi.Output[float]
    """
    The per-project quota on the prefix space that can be
    allocated from the subnetpool for project subnets. Changing this updates the
    default quota of the existing subnetpool.
    """
    description: pulumi.Output[str]
    """
    The human-readable description for the subnetpool.
    Changing this updates the description of the existing subnetpool.
    """
    ip_version: pulumi.Output[float]
    """
    The IP protocol version.
    """
    is_default: pulumi.Output[bool]
    """
    Indicates whether the subnetpool is default
    subnetpool or not. Changing this updates the default status of the existing
    subnetpool.
    """
    max_prefixlen: pulumi.Output[float]
    """
    The maximum prefix size that can be allocated from
    the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
    default is 128. Changing this updates the max prefixlen of the existing
    subnetpool.
    """
    min_prefixlen: pulumi.Output[float]
    """
    The smallest prefix that can be allocated from a
    subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
    is 64. Changing this updates the min prefixlen of the existing subnetpool.
    """
    name: pulumi.Output[str]
    """
    The name of the subnetpool. Changing this updates the name of
    the existing subnetpool.
    """
    prefixes: pulumi.Output[list]
    """
    A list of subnet prefixes to assign to the subnetpool.
    Neutron API merges adjacent prefixes and treats them as a single prefix. Each
    subnet prefix must be unique among all subnet prefixes in all subnetpools that
    are associated with the address scope. Changing this updates the prefixes list
    of the existing subnetpool.
    """
    project_id: pulumi.Output[str]
    """
    The owner of the subnetpool. Required if admin wants to
    create a subnetpool for another project. Changing this creates a new subnetpool.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron subnetpool. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    subnetpool.
    """
    revision_number: pulumi.Output[float]
    """
    The revision number of the subnetpool.
    """
    shared: pulumi.Output[bool]
    """
    Indicates whether this subnetpool is shared across
    all projects. Changing this updates the shared status of the existing
    subnetpool.
    """
    tags: pulumi.Output[list]
    """
    A set of string tags for the subnetpool.
    """
    updated_at: pulumi.Output[str]
    """
    The time at which subnetpool was created.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, address_scope_id=None, default_prefixlen=None, default_quota=None, description=None, ip_version=None, is_default=None, max_prefixlen=None, min_prefixlen=None, name=None, prefixes=None, project_id=None, region=None, shared=None, tags=None, value_specs=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V2 Neutron subnetpool resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_scope_id: The Neutron address scope to assign to the
               subnetpool. Changing this updates the address scope id of the existing
               subnetpool.
        :param pulumi.Input[float] default_prefixlen: The size of the prefix to allocate when the cidr
               or prefixlen attributes are omitted when you create the subnet. Defaults to the
               MinPrefixLen. Changing this updates the default prefixlen of the existing
               subnetpool.
        :param pulumi.Input[float] default_quota: The per-project quota on the prefix space that can be
               allocated from the subnetpool for project subnets. Changing this updates the
               default quota of the existing subnetpool.
        :param pulumi.Input[str] description: The human-readable description for the subnetpool.
               Changing this updates the description of the existing subnetpool.
        :param pulumi.Input[float] ip_version: The IP protocol version.
        :param pulumi.Input[bool] is_default: Indicates whether the subnetpool is default
               subnetpool or not. Changing this updates the default status of the existing
               subnetpool.
        :param pulumi.Input[float] max_prefixlen: The maximum prefix size that can be allocated from
               the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
               default is 128. Changing this updates the max prefixlen of the existing
               subnetpool.
        :param pulumi.Input[float] min_prefixlen: The smallest prefix that can be allocated from a
               subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
               is 64. Changing this updates the min prefixlen of the existing subnetpool.
        :param pulumi.Input[str] name: The name of the subnetpool. Changing this updates the name of
               the existing subnetpool.
        :param pulumi.Input[list] prefixes: A list of subnet prefixes to assign to the subnetpool.
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
        :param pulumi.Input[list] tags: A set of string tags for the subnetpool.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_subnetpool_v2.html.markdown.
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
            opts.version = utilities.get_version()
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
            if prefixes is None:
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
    def get(resource_name, id, opts=None, address_scope_id=None, all_tags=None, created_at=None, default_prefixlen=None, default_quota=None, description=None, ip_version=None, is_default=None, max_prefixlen=None, min_prefixlen=None, name=None, prefixes=None, project_id=None, region=None, revision_number=None, shared=None, tags=None, updated_at=None, value_specs=None):
        """
        Get an existing SubnetPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_scope_id: The Neutron address scope to assign to the
               subnetpool. Changing this updates the address scope id of the existing
               subnetpool.
        :param pulumi.Input[list] all_tags: The collection of tags assigned on the subnetpool, which have been
               explicitly and implicitly added.
        :param pulumi.Input[str] created_at: The time at which subnetpool was created.
        :param pulumi.Input[float] default_prefixlen: The size of the prefix to allocate when the cidr
               or prefixlen attributes are omitted when you create the subnet. Defaults to the
               MinPrefixLen. Changing this updates the default prefixlen of the existing
               subnetpool.
        :param pulumi.Input[float] default_quota: The per-project quota on the prefix space that can be
               allocated from the subnetpool for project subnets. Changing this updates the
               default quota of the existing subnetpool.
        :param pulumi.Input[str] description: The human-readable description for the subnetpool.
               Changing this updates the description of the existing subnetpool.
        :param pulumi.Input[float] ip_version: The IP protocol version.
        :param pulumi.Input[bool] is_default: Indicates whether the subnetpool is default
               subnetpool or not. Changing this updates the default status of the existing
               subnetpool.
        :param pulumi.Input[float] max_prefixlen: The maximum prefix size that can be allocated from
               the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools,
               default is 128. Changing this updates the max prefixlen of the existing
               subnetpool.
        :param pulumi.Input[float] min_prefixlen: The smallest prefix that can be allocated from a
               subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default
               is 64. Changing this updates the min prefixlen of the existing subnetpool.
        :param pulumi.Input[str] name: The name of the subnetpool. Changing this updates the name of
               the existing subnetpool.
        :param pulumi.Input[list] prefixes: A list of subnet prefixes to assign to the subnetpool.
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
        :param pulumi.Input[float] revision_number: The revision number of the subnetpool.
        :param pulumi.Input[bool] shared: Indicates whether this subnetpool is shared across
               all projects. Changing this updates the shared status of the existing
               subnetpool.
        :param pulumi.Input[list] tags: A set of string tags for the subnetpool.
        :param pulumi.Input[str] updated_at: The time at which subnetpool was created.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_subnetpool_v2.html.markdown.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

