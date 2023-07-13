# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SubnetPoolArgs', 'SubnetPool']

@pulumi.input_type
class SubnetPoolArgs:
    def __init__(__self__, *,
                 prefixes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 address_scope_id: Optional[pulumi.Input[str]] = None,
                 default_prefixlen: Optional[pulumi.Input[int]] = None,
                 default_quota: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[int]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 max_prefixlen: Optional[pulumi.Input[int]] = None,
                 min_prefixlen: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a SubnetPool resource.
        """
        pulumi.set(__self__, "prefixes", prefixes)
        if address_scope_id is not None:
            pulumi.set(__self__, "address_scope_id", address_scope_id)
        if default_prefixlen is not None:
            pulumi.set(__self__, "default_prefixlen", default_prefixlen)
        if default_quota is not None:
            pulumi.set(__self__, "default_quota", default_quota)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if max_prefixlen is not None:
            pulumi.set(__self__, "max_prefixlen", max_prefixlen)
        if min_prefixlen is not None:
            pulumi.set(__self__, "min_prefixlen", min_prefixlen)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def prefixes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "prefixes")

    @prefixes.setter
    def prefixes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "prefixes", value)

    @property
    @pulumi.getter(name="addressScopeId")
    def address_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "address_scope_id")

    @address_scope_id.setter
    def address_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_scope_id", value)

    @property
    @pulumi.getter(name="defaultPrefixlen")
    def default_prefixlen(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_prefixlen")

    @default_prefixlen.setter
    def default_prefixlen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_prefixlen", value)

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_quota")

    @default_quota.setter
    def default_quota(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_quota", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter(name="maxPrefixlen")
    def max_prefixlen(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_prefixlen")

    @max_prefixlen.setter
    def max_prefixlen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_prefixlen", value)

    @property
    @pulumi.getter(name="minPrefixlen")
    def min_prefixlen(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_prefixlen")

    @min_prefixlen.setter
    def min_prefixlen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_prefixlen", value)

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
    def shared(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _SubnetPoolState:
    def __init__(__self__, *,
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
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering SubnetPool resources.
        """
        if address_scope_id is not None:
            pulumi.set(__self__, "address_scope_id", address_scope_id)
        if all_tags is not None:
            pulumi.set(__self__, "all_tags", all_tags)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if default_prefixlen is not None:
            pulumi.set(__self__, "default_prefixlen", default_prefixlen)
        if default_quota is not None:
            pulumi.set(__self__, "default_quota", default_quota)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if max_prefixlen is not None:
            pulumi.set(__self__, "max_prefixlen", max_prefixlen)
        if min_prefixlen is not None:
            pulumi.set(__self__, "min_prefixlen", min_prefixlen)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if prefixes is not None:
            pulumi.set(__self__, "prefixes", prefixes)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if revision_number is not None:
            pulumi.set(__self__, "revision_number", revision_number)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter(name="addressScopeId")
    def address_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "address_scope_id")

    @address_scope_id.setter
    def address_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_scope_id", value)

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "all_tags")

    @all_tags.setter
    def all_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "all_tags", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="defaultPrefixlen")
    def default_prefixlen(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_prefixlen")

    @default_prefixlen.setter
    def default_prefixlen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_prefixlen", value)

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "default_quota")

    @default_quota.setter
    def default_quota(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_quota", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter(name="maxPrefixlen")
    def max_prefixlen(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_prefixlen")

    @max_prefixlen.setter
    def max_prefixlen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_prefixlen", value)

    @property
    @pulumi.getter(name="minPrefixlen")
    def min_prefixlen(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "min_prefixlen")

    @min_prefixlen.setter
    def min_prefixlen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_prefixlen", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "prefixes")

    @prefixes.setter
    def prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "prefixes", value)

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
    @pulumi.getter(name="revisionNumber")
    def revision_number(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "revision_number")

    @revision_number.setter
    def revision_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision_number", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


class SubnetPool(pulumi.CustomResource):
    @overload
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
                 __props__=None):
        """
        Create a SubnetPool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubnetPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SubnetPool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SubnetPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
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
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetPoolArgs.__new__(SubnetPoolArgs)

            __props__.__dict__["address_scope_id"] = address_scope_id
            __props__.__dict__["default_prefixlen"] = default_prefixlen
            __props__.__dict__["default_quota"] = default_quota
            __props__.__dict__["description"] = description
            __props__.__dict__["ip_version"] = ip_version
            __props__.__dict__["is_default"] = is_default
            __props__.__dict__["max_prefixlen"] = max_prefixlen
            __props__.__dict__["min_prefixlen"] = min_prefixlen
            __props__.__dict__["name"] = name
            if prefixes is None and not opts.urn:
                raise TypeError("Missing required property 'prefixes'")
            __props__.__dict__["prefixes"] = prefixes
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["shared"] = shared
            __props__.__dict__["tags"] = tags
            __props__.__dict__["value_specs"] = value_specs
            __props__.__dict__["all_tags"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["revision_number"] = None
            __props__.__dict__["updated_at"] = None
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
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubnetPoolState.__new__(_SubnetPoolState)

        __props__.__dict__["address_scope_id"] = address_scope_id
        __props__.__dict__["all_tags"] = all_tags
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["default_prefixlen"] = default_prefixlen
        __props__.__dict__["default_quota"] = default_quota
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_version"] = ip_version
        __props__.__dict__["is_default"] = is_default
        __props__.__dict__["max_prefixlen"] = max_prefixlen
        __props__.__dict__["min_prefixlen"] = min_prefixlen
        __props__.__dict__["name"] = name
        __props__.__dict__["prefixes"] = prefixes
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["revision_number"] = revision_number
        __props__.__dict__["shared"] = shared
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["value_specs"] = value_specs
        return SubnetPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressScopeId")
    def address_scope_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "address_scope_id")

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="defaultPrefixlen")
    def default_prefixlen(self) -> pulumi.Output[int]:
        return pulumi.get(self, "default_prefixlen")

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "default_quota")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[int]:
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="maxPrefixlen")
    def max_prefixlen(self) -> pulumi.Output[int]:
        return pulumi.get(self, "max_prefixlen")

    @property
    @pulumi.getter(name="minPrefixlen")
    def min_prefixlen(self) -> pulumi.Output[int]:
        return pulumi.get(self, "min_prefixlen")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def prefixes(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "prefixes")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="revisionNumber")
    def revision_number(self) -> pulumi.Output[int]:
        return pulumi.get(self, "revision_number")

    @property
    @pulumi.getter
    def shared(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "value_specs")

