# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['StackV1']


class StackV1(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capabilities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_rollback: Optional[pulumi.Input[bool]] = None,
                 environment_opts: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackV1OutputArgs']]]]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_reason: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 template_description: Optional[pulumi.Input[str]] = None,
                 template_opts: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 updated_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V1 stack resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        stack1 = openstack.orchestration.StackV1("stack1",
            disable_rollback=True,
            environment_opts={
                "Bin": "\n\n",
            },
            parameters={
                "length": 4,
            },
            template_opts={
                "Bin": \"\"\"heat_template_version: 2013-05-23
        parameters:
          length:
            type: number
        resources:
          test_res:
            type: OS::Heat::TestResource
          random:
            type: OS::Heat::RandomString
            properties:
              length: {get_param: length}

        \"\"\",
            },
            timeout=30)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] capabilities: List of stack capabilities for stack.
        :param pulumi.Input[str] creation_time: The date and time when the resource was created. The date
               and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
               For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
               is the time zone as an offset from UTC.
        :param pulumi.Input[str] description: The description of the stack resource.
        :param pulumi.Input[bool] disable_rollback: Enables or disables deletion of all stack
               resources when a stack creation fails. Default is true, meaning all
               resources are not deleted when stack creation fails.
        :param pulumi.Input[Mapping[str, Any]] environment_opts: Environment key/value pairs to associate with
               the stack which contains details for the environment of the stack.
               Allowed keys: Bin, URL, Files. Changing this updates the existing stack
               Environment Opts.
        :param pulumi.Input[str] name: A unique name for the stack. It must start with an
               alphabetic character. Changing this updates the stack's name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_topics: List of notification topics for stack.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackV1OutputArgs']]]] outputs: A list of stack outputs.
        :param pulumi.Input[Mapping[str, Any]] parameters: User-defined key/value pairs as parameters to pass
               to the template. Changing this updates the existing stack parameters.
        :param pulumi.Input[str] region: The region in which to create the stack. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new stack.
        :param pulumi.Input[str] status: The status of the stack.
        :param pulumi.Input[str] status_reason: The reason for the current status of the stack.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to assosciate with the Stack
        :param pulumi.Input[str] template_description: The description of the stack template.
        :param pulumi.Input[Mapping[str, Any]] template_opts: Template key/value pairs to associate with the
               stack which contains either the template file or url.
               Allowed keys: Bin, URL, Files. Changing this updates the existing stack
               Template Opts.
        :param pulumi.Input[int] timeout: The timeout for stack action in minutes.
        :param pulumi.Input[str] updated_time: The date and time when the resource was updated. The date
               and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
               For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
               is the time zone as an offset from UTC.
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

            __props__['capabilities'] = capabilities
            __props__['creation_time'] = creation_time
            __props__['description'] = description
            __props__['disable_rollback'] = disable_rollback
            __props__['environment_opts'] = environment_opts
            __props__['name'] = name
            __props__['notification_topics'] = notification_topics
            __props__['outputs'] = outputs
            __props__['parameters'] = parameters
            __props__['region'] = region
            __props__['status'] = status
            __props__['status_reason'] = status_reason
            __props__['tags'] = tags
            __props__['template_description'] = template_description
            if template_opts is None:
                raise TypeError("Missing required property 'template_opts'")
            __props__['template_opts'] = template_opts
            __props__['timeout'] = timeout
            __props__['updated_time'] = updated_time
        super(StackV1, __self__).__init__(
            'openstack:orchestration/stackV1:StackV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            capabilities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_rollback: Optional[pulumi.Input[bool]] = None,
            environment_opts: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackV1OutputArgs']]]]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            status_reason: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            template_description: Optional[pulumi.Input[str]] = None,
            template_opts: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            updated_time: Optional[pulumi.Input[str]] = None) -> 'StackV1':
        """
        Get an existing StackV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] capabilities: List of stack capabilities for stack.
        :param pulumi.Input[str] creation_time: The date and time when the resource was created. The date
               and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
               For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
               is the time zone as an offset from UTC.
        :param pulumi.Input[str] description: The description of the stack resource.
        :param pulumi.Input[bool] disable_rollback: Enables or disables deletion of all stack
               resources when a stack creation fails. Default is true, meaning all
               resources are not deleted when stack creation fails.
        :param pulumi.Input[Mapping[str, Any]] environment_opts: Environment key/value pairs to associate with
               the stack which contains details for the environment of the stack.
               Allowed keys: Bin, URL, Files. Changing this updates the existing stack
               Environment Opts.
        :param pulumi.Input[str] name: A unique name for the stack. It must start with an
               alphabetic character. Changing this updates the stack's name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_topics: List of notification topics for stack.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StackV1OutputArgs']]]] outputs: A list of stack outputs.
        :param pulumi.Input[Mapping[str, Any]] parameters: User-defined key/value pairs as parameters to pass
               to the template. Changing this updates the existing stack parameters.
        :param pulumi.Input[str] region: The region in which to create the stack. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new stack.
        :param pulumi.Input[str] status: The status of the stack.
        :param pulumi.Input[str] status_reason: The reason for the current status of the stack.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to assosciate with the Stack
        :param pulumi.Input[str] template_description: The description of the stack template.
        :param pulumi.Input[Mapping[str, Any]] template_opts: Template key/value pairs to associate with the
               stack which contains either the template file or url.
               Allowed keys: Bin, URL, Files. Changing this updates the existing stack
               Template Opts.
        :param pulumi.Input[int] timeout: The timeout for stack action in minutes.
        :param pulumi.Input[str] updated_time: The date and time when the resource was updated. The date
               and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
               For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
               is the time zone as an offset from UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["capabilities"] = capabilities
        __props__["creation_time"] = creation_time
        __props__["description"] = description
        __props__["disable_rollback"] = disable_rollback
        __props__["environment_opts"] = environment_opts
        __props__["name"] = name
        __props__["notification_topics"] = notification_topics
        __props__["outputs"] = outputs
        __props__["parameters"] = parameters
        __props__["region"] = region
        __props__["status"] = status
        __props__["status_reason"] = status_reason
        __props__["tags"] = tags
        __props__["template_description"] = template_description
        __props__["template_opts"] = template_opts
        __props__["timeout"] = timeout
        __props__["updated_time"] = updated_time
        return StackV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capabilities(self) -> pulumi.Output[Sequence[str]]:
        """
        List of stack capabilities for stack.
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The date and time when the resource was created. The date
        and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        is the time zone as an offset from UTC.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the stack resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableRollback")
    def disable_rollback(self) -> pulumi.Output[bool]:
        """
        Enables or disables deletion of all stack
        resources when a stack creation fails. Default is true, meaning all
        resources are not deleted when stack creation fails.
        """
        return pulumi.get(self, "disable_rollback")

    @property
    @pulumi.getter(name="environmentOpts")
    def environment_opts(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Environment key/value pairs to associate with
        the stack which contains details for the environment of the stack.
        Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        Environment Opts.
        """
        return pulumi.get(self, "environment_opts")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the stack. It must start with an
        alphabetic character. Changing this updates the stack's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationTopics")
    def notification_topics(self) -> pulumi.Output[Sequence[str]]:
        """
        List of notification topics for stack.
        """
        return pulumi.get(self, "notification_topics")

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Sequence['outputs.StackV1Output']]:
        """
        A list of stack outputs.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        User-defined key/value pairs as parameters to pass
        to the template. Changing this updates the existing stack parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the stack. If
        omitted, the `region` argument of the provider is used. Changing this
        creates a new stack.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the stack.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> pulumi.Output[str]:
        """
        The reason for the current status of the stack.
        """
        return pulumi.get(self, "status_reason")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of tags to assosciate with the Stack
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateDescription")
    def template_description(self) -> pulumi.Output[str]:
        """
        The description of the stack template.
        """
        return pulumi.get(self, "template_description")

    @property
    @pulumi.getter(name="templateOpts")
    def template_opts(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Template key/value pairs to associate with the
        stack which contains either the template file or url.
        Allowed keys: Bin, URL, Files. Changing this updates the existing stack
        Template Opts.
        """
        return pulumi.get(self, "template_opts")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[int]:
        """
        The timeout for stack action in minutes.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> pulumi.Output[str]:
        """
        The date and time when the resource was updated. The date
        and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
        For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
        is the time zone as an offset from UTC.
        """
        return pulumi.get(self, "updated_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

