# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class QosPolicy(pulumi.CustomResource):
    all_tags: pulumi.Output[list]
    """
    The collection of tags assigned on the QoS policy, which have been
    explicitly and implicitly added.
    """
    created_at: pulumi.Output[str]
    """
    The time at which QoS policy was created.
    """
    description: pulumi.Output[str]
    """
    The human-readable description for the QoS policy.
    Changing this updates the description of the existing QoS policy.
    """
    is_default: pulumi.Output[bool]
    """
    Indicates whether the QoS policy is default
    QoS policy or not. Changing this updates the default status of the existing
    QoS policy.
    """
    name: pulumi.Output[str]
    """
    The name of the QoS policy. Changing this updates the name of
    the existing QoS policy.
    """
    project_id: pulumi.Output[str]
    """
    The owner of the QoS policy. Required if admin wants to
    create a QoS policy for another project. Changing this creates a new QoS policy.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create a Neutron Qos policy. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    QoS policy.
    """
    revision_number: pulumi.Output[float]
    """
    The revision number of the QoS policy.
    """
    shared: pulumi.Output[bool]
    """
    Indicates whether this QoS policy is shared across
    all projects. Changing this updates the shared status of the existing
    QoS policy.
    """
    tags: pulumi.Output[list]
    """
    A set of string tags for the QoS policy.
    """
    updated_at: pulumi.Output[str]
    """
    The time at which QoS policy was created.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, description=None, is_default=None, name=None, project_id=None, region=None, shared=None, tags=None, value_specs=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V2 Neutron QoS policy resource within OpenStack.
        
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
        :param pulumi.Input[list] tags: A set of string tags for the QoS policy.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_qos_policy_v2.html.markdown.
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

            __props__['description'] = description
            __props__['is_default'] = is_default
            __props__['name'] = name
            __props__['project_id'] = project_id
            __props__['region'] = region
            __props__['shared'] = shared
            __props__['tags'] = tags
            __props__['value_specs'] = value_specs
            __props__['all_tags'] = None
            __props__['created_at'] = None
            __props__['revision_number'] = None
            __props__['updated_at'] = None
        super(QosPolicy, __self__).__init__(
            'openstack:networking/qosPolicy:QosPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, all_tags=None, created_at=None, description=None, is_default=None, name=None, project_id=None, region=None, revision_number=None, shared=None, tags=None, updated_at=None, value_specs=None):
        """
        Get an existing QosPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] all_tags: The collection of tags assigned on the QoS policy, which have been
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
        :param pulumi.Input[float] revision_number: The revision number of the QoS policy.
        :param pulumi.Input[bool] shared: Indicates whether this QoS policy is shared across
               all projects. Changing this updates the shared status of the existing
               QoS policy.
        :param pulumi.Input[list] tags: A set of string tags for the QoS policy.
        :param pulumi.Input[str] updated_at: The time at which QoS policy was created.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_qos_policy_v2.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["all_tags"] = all_tags
        __props__["created_at"] = created_at
        __props__["description"] = description
        __props__["is_default"] = is_default
        __props__["name"] = name
        __props__["project_id"] = project_id
        __props__["region"] = region
        __props__["revision_number"] = revision_number
        __props__["shared"] = shared
        __props__["tags"] = tags
        __props__["updated_at"] = updated_at
        __props__["value_specs"] = value_specs
        return QosPolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

