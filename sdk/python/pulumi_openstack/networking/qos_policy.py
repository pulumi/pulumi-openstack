# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['QosPolicyArgs', 'QosPolicy']

@pulumi.input_type
class QosPolicyArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a QosPolicy resource.
        :param pulumi.Input[str] description: The human-readable description for the QoS policy.
               Changing this updates the description of the existing QoS policy.
        :param pulumi.Input[bool] is_default: Indicates whether the QoS policy is default
               QoS policy or not. Changing this updates the default status of the existing
               QoS policy.
        :param pulumi.Input[str] name: The name of the QoS policy. Changing this updates the name of
               the existing QoS policy.
        :param pulumi.Input[str] project_id: The owner of the QoS policy. Required if admin wants to
               create a QoS policy for another project. Changing this creates a new QoS policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron Qos policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               QoS policy.
        :param pulumi.Input[bool] shared: Indicates whether this QoS policy is shared across
               all projects. Changing this updates the shared status of the existing
               QoS policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the QoS policy.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        QosPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            is_default=is_default,
            name=name,
            project_id=project_id,
            region=region,
            shared=shared,
            tags=tags,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             is_default: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             shared: Optional[pulumi.Input[bool]] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if is_default is None and 'isDefault' in kwargs:
            is_default = kwargs['isDefault']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']

        if description is not None:
            _setter("description", description)
        if is_default is not None:
            _setter("is_default", is_default)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)
        if region is not None:
            _setter("region", region)
        if shared is not None:
            _setter("shared", shared)
        if tags is not None:
            _setter("tags", tags)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable description for the QoS policy.
        Changing this updates the description of the existing QoS policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the QoS policy is default
        QoS policy or not. Changing this updates the default status of the existing
        QoS policy.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the QoS policy. Changing this updates the name of
        the existing QoS policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the QoS policy. Required if admin wants to
        create a QoS policy for another project. Changing this creates a new QoS policy.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron Qos policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        QoS policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this QoS policy is shared across
        all projects. Changing this updates the shared status of the existing
        QoS policy.
        """
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of string tags for the QoS policy.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _QosPolicyState:
    def __init__(__self__, *,
                 all_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 revision_number: Optional[pulumi.Input[int]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering QosPolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] all_tags: The collection of tags assigned on the QoS policy, which have been
               explicitly and implicitly added.
        :param pulumi.Input[str] created_at: The time at which QoS policy was created.
        :param pulumi.Input[str] description: The human-readable description for the QoS policy.
               Changing this updates the description of the existing QoS policy.
        :param pulumi.Input[bool] is_default: Indicates whether the QoS policy is default
               QoS policy or not. Changing this updates the default status of the existing
               QoS policy.
        :param pulumi.Input[str] name: The name of the QoS policy. Changing this updates the name of
               the existing QoS policy.
        :param pulumi.Input[str] project_id: The owner of the QoS policy. Required if admin wants to
               create a QoS policy for another project. Changing this creates a new QoS policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron Qos policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               QoS policy.
        :param pulumi.Input[int] revision_number: The revision number of the QoS policy.
        :param pulumi.Input[bool] shared: Indicates whether this QoS policy is shared across
               all projects. Changing this updates the shared status of the existing
               QoS policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the QoS policy.
        :param pulumi.Input[str] updated_at: The time at which QoS policy was created.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        _QosPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            all_tags=all_tags,
            created_at=created_at,
            description=description,
            is_default=is_default,
            name=name,
            project_id=project_id,
            region=region,
            revision_number=revision_number,
            shared=shared,
            tags=tags,
            updated_at=updated_at,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             all_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             created_at: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             is_default: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             revision_number: Optional[pulumi.Input[int]] = None,
             shared: Optional[pulumi.Input[bool]] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             updated_at: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if all_tags is None and 'allTags' in kwargs:
            all_tags = kwargs['allTags']
        if created_at is None and 'createdAt' in kwargs:
            created_at = kwargs['createdAt']
        if is_default is None and 'isDefault' in kwargs:
            is_default = kwargs['isDefault']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if revision_number is None and 'revisionNumber' in kwargs:
            revision_number = kwargs['revisionNumber']
        if updated_at is None and 'updatedAt' in kwargs:
            updated_at = kwargs['updatedAt']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']

        if all_tags is not None:
            _setter("all_tags", all_tags)
        if created_at is not None:
            _setter("created_at", created_at)
        if description is not None:
            _setter("description", description)
        if is_default is not None:
            _setter("is_default", is_default)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)
        if region is not None:
            _setter("region", region)
        if revision_number is not None:
            _setter("revision_number", revision_number)
        if shared is not None:
            _setter("shared", shared)
        if tags is not None:
            _setter("tags", tags)
        if updated_at is not None:
            _setter("updated_at", updated_at)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The collection of tags assigned on the QoS policy, which have been
        explicitly and implicitly added.
        """
        return pulumi.get(self, "all_tags")

    @all_tags.setter
    def all_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "all_tags", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which QoS policy was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable description for the QoS policy.
        Changing this updates the description of the existing QoS policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the QoS policy is default
        QoS policy or not. Changing this updates the default status of the existing
        QoS policy.
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the QoS policy. Changing this updates the name of
        the existing QoS policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the QoS policy. Required if admin wants to
        create a QoS policy for another project. Changing this creates a new QoS policy.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron Qos policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        QoS policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="revisionNumber")
    def revision_number(self) -> Optional[pulumi.Input[int]]:
        """
        The revision number of the QoS policy.
        """
        return pulumi.get(self, "revision_number")

    @revision_number.setter
    def revision_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision_number", value)

    @property
    @pulumi.getter
    def shared(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this QoS policy is shared across
        all projects. Changing this updates the shared status of the existing
        QoS policy.
        """
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of string tags for the QoS policy.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which QoS policy was created.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


class QosPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Manages a V2 Neutron QoS policy resource within OpenStack.

        ## Example Usage
        ### Create a QoS Policy

        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos_policy1 = openstack.networking.QosPolicy("qosPolicy1", description="bw_limit")
        ```

        ## Import

        QoS Policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:networking/qosPolicy:QosPolicy qos_policy_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The human-readable description for the QoS policy.
               Changing this updates the description of the existing QoS policy.
        :param pulumi.Input[bool] is_default: Indicates whether the QoS policy is default
               QoS policy or not. Changing this updates the default status of the existing
               QoS policy.
        :param pulumi.Input[str] name: The name of the QoS policy. Changing this updates the name of
               the existing QoS policy.
        :param pulumi.Input[str] project_id: The owner of the QoS policy. Required if admin wants to
               create a QoS policy for another project. Changing this creates a new QoS policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron Qos policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               QoS policy.
        :param pulumi.Input[bool] shared: Indicates whether this QoS policy is shared across
               all projects. Changing this updates the shared status of the existing
               QoS policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the QoS policy.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[QosPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 Neutron QoS policy resource within OpenStack.

        ## Example Usage
        ### Create a QoS Policy

        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos_policy1 = openstack.networking.QosPolicy("qosPolicy1", description="bw_limit")
        ```

        ## Import

        QoS Policies can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:networking/qosPolicy:QosPolicy qos_policy_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae
        ```

        :param str resource_name: The name of the resource.
        :param QosPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QosPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            QosPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = QosPolicyArgs.__new__(QosPolicyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["is_default"] = is_default
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["shared"] = shared
            __props__.__dict__["tags"] = tags
            __props__.__dict__["value_specs"] = value_specs
            __props__.__dict__["all_tags"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["revision_number"] = None
            __props__.__dict__["updated_at"] = None
        super(QosPolicy, __self__).__init__(
            'openstack:networking/qosPolicy:QosPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            revision_number: Optional[pulumi.Input[int]] = None,
            shared: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'QosPolicy':
        """
        Get an existing QosPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] all_tags: The collection of tags assigned on the QoS policy, which have been
               explicitly and implicitly added.
        :param pulumi.Input[str] created_at: The time at which QoS policy was created.
        :param pulumi.Input[str] description: The human-readable description for the QoS policy.
               Changing this updates the description of the existing QoS policy.
        :param pulumi.Input[bool] is_default: Indicates whether the QoS policy is default
               QoS policy or not. Changing this updates the default status of the existing
               QoS policy.
        :param pulumi.Input[str] name: The name of the QoS policy. Changing this updates the name of
               the existing QoS policy.
        :param pulumi.Input[str] project_id: The owner of the QoS policy. Required if admin wants to
               create a QoS policy for another project. Changing this creates a new QoS policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron Qos policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               QoS policy.
        :param pulumi.Input[int] revision_number: The revision number of the QoS policy.
        :param pulumi.Input[bool] shared: Indicates whether this QoS policy is shared across
               all projects. Changing this updates the shared status of the existing
               QoS policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A set of string tags for the QoS policy.
        :param pulumi.Input[str] updated_at: The time at which QoS policy was created.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QosPolicyState.__new__(_QosPolicyState)

        __props__.__dict__["all_tags"] = all_tags
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["is_default"] = is_default
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["revision_number"] = revision_number
        __props__.__dict__["shared"] = shared
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["value_specs"] = value_specs
        return QosPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allTags")
    def all_tags(self) -> pulumi.Output[Sequence[str]]:
        """
        The collection of tags assigned on the QoS policy, which have been
        explicitly and implicitly added.
        """
        return pulumi.get(self, "all_tags")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The time at which QoS policy was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The human-readable description for the QoS policy.
        Changing this updates the description of the existing QoS policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether the QoS policy is default
        QoS policy or not. Changing this updates the default status of the existing
        QoS policy.
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the QoS policy. Changing this updates the name of
        the existing QoS policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The owner of the QoS policy. Required if admin wants to
        create a QoS policy for another project. Changing this creates a new QoS policy.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron Qos policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        QoS policy.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="revisionNumber")
    def revision_number(self) -> pulumi.Output[int]:
        """
        The revision number of the QoS policy.
        """
        return pulumi.get(self, "revision_number")

    @property
    @pulumi.getter
    def shared(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether this QoS policy is shared across
        all projects. Changing this updates the shared status of the existing
        QoS policy.
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of string tags for the QoS policy.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The time at which QoS policy was created.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

