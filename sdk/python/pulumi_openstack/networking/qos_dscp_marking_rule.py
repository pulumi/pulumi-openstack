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

__all__ = ['QosDscpMarkingRuleArgs', 'QosDscpMarkingRule']

@pulumi.input_type
class QosDscpMarkingRuleArgs:
    def __init__(__self__, *,
                 dscp_mark: pulumi.Input[builtins.int],
                 qos_policy_id: pulumi.Input[builtins.str],
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a QosDscpMarkingRule resource.
        :param pulumi.Input[builtins.int] dscp_mark: The value of DSCP mark. Changing this updates the DSCP mark value existing
               QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
        """
        pulumi.set(__self__, "dscp_mark", dscp_mark)
        pulumi.set(__self__, "qos_policy_id", qos_policy_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="dscpMark")
    def dscp_mark(self) -> pulumi.Input[builtins.int]:
        """
        The value of DSCP mark. Changing this updates the DSCP mark value existing
        QoS DSCP marking rule.
        """
        return pulumi.get(self, "dscp_mark")

    @dscp_mark.setter
    def dscp_mark(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "dscp_mark", value)

    @property
    @pulumi.getter(name="qosPolicyId")
    def qos_policy_id(self) -> pulumi.Input[builtins.str]:
        """
        The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
        """
        return pulumi.get(self, "qos_policy_id")

    @qos_policy_id.setter
    def qos_policy_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "qos_policy_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _QosDscpMarkingRuleState:
    def __init__(__self__, *,
                 dscp_mark: Optional[pulumi.Input[builtins.int]] = None,
                 qos_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering QosDscpMarkingRule resources.
        :param pulumi.Input[builtins.int] dscp_mark: The value of DSCP mark. Changing this updates the DSCP mark value existing
               QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
        """
        if dscp_mark is not None:
            pulumi.set(__self__, "dscp_mark", dscp_mark)
        if qos_policy_id is not None:
            pulumi.set(__self__, "qos_policy_id", qos_policy_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="dscpMark")
    def dscp_mark(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The value of DSCP mark. Changing this updates the DSCP mark value existing
        QoS DSCP marking rule.
        """
        return pulumi.get(self, "dscp_mark")

    @dscp_mark.setter
    def dscp_mark(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "dscp_mark", value)

    @property
    @pulumi.getter(name="qosPolicyId")
    def qos_policy_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
        """
        return pulumi.get(self, "qos_policy_id")

    @qos_policy_id.setter
    def qos_policy_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "qos_policy_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


class QosDscpMarkingRule(pulumi.CustomResource):

    pulumi_type = "openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule"

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dscp_mark: Optional[pulumi.Input[builtins.int]] = None,
                 qos_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages a V2 Neutron QoS DSCP marking rule resource within OpenStack.

        ## Example Usage

        ### Create a QoS Policy with some DSCP marking rule

        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos_policy1 = openstack.networking.QosPolicy("qos_policy_1",
            name="qos_policy_1",
            description="dscp_mark")
        dscp_marking_rule1 = openstack.networking.QosDscpMarkingRule("dscp_marking_rule_1",
            qos_policy_id=qos_policy1.id,
            dscp_mark=26)
        ```

        ## Import

        QoS DSCP marking rules can be imported using the `qos_policy_id/dscp_marking_rule_id` format, e.g.

        ```sh
        $ pulumi import openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule dscp_marking_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] dscp_mark: The value of DSCP mark. Changing this updates the DSCP mark value existing
               QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QosDscpMarkingRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 Neutron QoS DSCP marking rule resource within OpenStack.

        ## Example Usage

        ### Create a QoS Policy with some DSCP marking rule

        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos_policy1 = openstack.networking.QosPolicy("qos_policy_1",
            name="qos_policy_1",
            description="dscp_mark")
        dscp_marking_rule1 = openstack.networking.QosDscpMarkingRule("dscp_marking_rule_1",
            qos_policy_id=qos_policy1.id,
            dscp_mark=26)
        ```

        ## Import

        QoS DSCP marking rules can be imported using the `qos_policy_id/dscp_marking_rule_id` format, e.g.

        ```sh
        $ pulumi import openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule dscp_marking_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
        ```

        :param str resource_name: The name of the resource.
        :param QosDscpMarkingRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QosDscpMarkingRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dscp_mark: Optional[pulumi.Input[builtins.int]] = None,
                 qos_policy_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QosDscpMarkingRuleArgs.__new__(QosDscpMarkingRuleArgs)

            if dscp_mark is None and not opts.urn:
                raise TypeError("Missing required property 'dscp_mark'")
            __props__.__dict__["dscp_mark"] = dscp_mark
            if qos_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'qos_policy_id'")
            __props__.__dict__["qos_policy_id"] = qos_policy_id
            __props__.__dict__["region"] = region
        super(QosDscpMarkingRule, __self__).__init__(
            'openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dscp_mark: Optional[pulumi.Input[builtins.int]] = None,
            qos_policy_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None) -> 'QosDscpMarkingRule':
        """
        Get an existing QosDscpMarkingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] dscp_mark: The value of DSCP mark. Changing this updates the DSCP mark value existing
               QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QosDscpMarkingRuleState.__new__(_QosDscpMarkingRuleState)

        __props__.__dict__["dscp_mark"] = dscp_mark
        __props__.__dict__["qos_policy_id"] = qos_policy_id
        __props__.__dict__["region"] = region
        return QosDscpMarkingRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dscpMark")
    def dscp_mark(self) -> pulumi.Output[builtins.int]:
        """
        The value of DSCP mark. Changing this updates the DSCP mark value existing
        QoS DSCP marking rule.
        """
        return pulumi.get(self, "dscp_mark")

    @property
    @pulumi.getter(name="qosPolicyId")
    def qos_policy_id(self) -> pulumi.Output[builtins.str]:
        """
        The QoS policy reference. Changing this creates a new QoS DSCP marking rule.
        """
        return pulumi.get(self, "qos_policy_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.
        """
        return pulumi.get(self, "region")

