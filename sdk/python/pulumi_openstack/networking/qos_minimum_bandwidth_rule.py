# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['QosMinimumBandwidthRuleArgs', 'QosMinimumBandwidthRule']

@pulumi.input_type
class QosMinimumBandwidthRuleArgs:
    def __init__(__self__, *,
                 min_kbps: pulumi.Input[int],
                 qos_policy_id: pulumi.Input[str],
                 direction: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a QosMinimumBandwidthRule resource.
        :param pulumi.Input[int] min_kbps: The minimum kilobits per second. Changing this updates the min kbps value of the existing
               QoS minimum bandwidth rule.
        :param pulumi.Input[str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        :param pulumi.Input[str] direction: The direction of traffic. Defaults to "egress". Changing this updates the direction of the
               existing QoS minimum bandwidth rule.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        """
        pulumi.set(__self__, "min_kbps", min_kbps)
        pulumi.set(__self__, "qos_policy_id", qos_policy_id)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="minKbps")
    def min_kbps(self) -> pulumi.Input[int]:
        """
        The minimum kilobits per second. Changing this updates the min kbps value of the existing
        QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "min_kbps")

    @min_kbps.setter
    def min_kbps(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_kbps", value)

    @property
    @pulumi.getter(name="qosPolicyId")
    def qos_policy_id(self) -> pulumi.Input[str]:
        """
        The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "qos_policy_id")

    @qos_policy_id.setter
    def qos_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "qos_policy_id", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        existing QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _QosMinimumBandwidthRuleState:
    def __init__(__self__, *,
                 direction: Optional[pulumi.Input[str]] = None,
                 min_kbps: Optional[pulumi.Input[int]] = None,
                 qos_policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering QosMinimumBandwidthRule resources.
        :param pulumi.Input[str] direction: The direction of traffic. Defaults to "egress". Changing this updates the direction of the
               existing QoS minimum bandwidth rule.
        :param pulumi.Input[int] min_kbps: The minimum kilobits per second. Changing this updates the min kbps value of the existing
               QoS minimum bandwidth rule.
        :param pulumi.Input[str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        """
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if min_kbps is not None:
            pulumi.set(__self__, "min_kbps", min_kbps)
        if qos_policy_id is not None:
            pulumi.set(__self__, "qos_policy_id", qos_policy_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        existing QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="minKbps")
    def min_kbps(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum kilobits per second. Changing this updates the min kbps value of the existing
        QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "min_kbps")

    @min_kbps.setter
    def min_kbps(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_kbps", value)

    @property
    @pulumi.getter(name="qosPolicyId")
    def qos_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "qos_policy_id")

    @qos_policy_id.setter
    def qos_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qos_policy_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class QosMinimumBandwidthRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 min_kbps: Optional[pulumi.Input[int]] = None,
                 qos_policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.

        ## Example Usage

        ### Create a QoS Policy with some minimum bandwidth rule

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos_policy1 = openstack.networking.QosPolicy("qos_policy_1",
            name="qos_policy_1",
            description="min_kbps")
        minimum_bandwidth_rule1 = openstack.networking.QosMinimumBandwidthRule("minimum_bandwidth_rule_1",
            qos_policy_id=qos_policy1.id,
            min_kbps=200)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        QoS minimum bandwidth rules can be imported using the `qos_policy_id/minimum_bandwidth_rule_id` format, e.g.

        ```sh
        $ pulumi import openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule minimum_bandwidth_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] direction: The direction of traffic. Defaults to "egress". Changing this updates the direction of the
               existing QoS minimum bandwidth rule.
        :param pulumi.Input[int] min_kbps: The minimum kilobits per second. Changing this updates the min kbps value of the existing
               QoS minimum bandwidth rule.
        :param pulumi.Input[str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QosMinimumBandwidthRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.

        ## Example Usage

        ### Create a QoS Policy with some minimum bandwidth rule

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos_policy1 = openstack.networking.QosPolicy("qos_policy_1",
            name="qos_policy_1",
            description="min_kbps")
        minimum_bandwidth_rule1 = openstack.networking.QosMinimumBandwidthRule("minimum_bandwidth_rule_1",
            qos_policy_id=qos_policy1.id,
            min_kbps=200)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        QoS minimum bandwidth rules can be imported using the `qos_policy_id/minimum_bandwidth_rule_id` format, e.g.

        ```sh
        $ pulumi import openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule minimum_bandwidth_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94
        ```

        :param str resource_name: The name of the resource.
        :param QosMinimumBandwidthRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QosMinimumBandwidthRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 min_kbps: Optional[pulumi.Input[int]] = None,
                 qos_policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QosMinimumBandwidthRuleArgs.__new__(QosMinimumBandwidthRuleArgs)

            __props__.__dict__["direction"] = direction
            if min_kbps is None and not opts.urn:
                raise TypeError("Missing required property 'min_kbps'")
            __props__.__dict__["min_kbps"] = min_kbps
            if qos_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'qos_policy_id'")
            __props__.__dict__["qos_policy_id"] = qos_policy_id
            __props__.__dict__["region"] = region
        super(QosMinimumBandwidthRule, __self__).__init__(
            'openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            direction: Optional[pulumi.Input[str]] = None,
            min_kbps: Optional[pulumi.Input[int]] = None,
            qos_policy_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'QosMinimumBandwidthRule':
        """
        Get an existing QosMinimumBandwidthRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] direction: The direction of traffic. Defaults to "egress". Changing this updates the direction of the
               existing QoS minimum bandwidth rule.
        :param pulumi.Input[int] min_kbps: The minimum kilobits per second. Changing this updates the min kbps value of the existing
               QoS minimum bandwidth rule.
        :param pulumi.Input[str] qos_policy_id: The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QosMinimumBandwidthRuleState.__new__(_QosMinimumBandwidthRuleState)

        __props__.__dict__["direction"] = direction
        __props__.__dict__["min_kbps"] = min_kbps
        __props__.__dict__["qos_policy_id"] = qos_policy_id
        __props__.__dict__["region"] = region
        return QosMinimumBandwidthRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[Optional[str]]:
        """
        The direction of traffic. Defaults to "egress". Changing this updates the direction of the
        existing QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter(name="minKbps")
    def min_kbps(self) -> pulumi.Output[int]:
        """
        The minimum kilobits per second. Changing this updates the min kbps value of the existing
        QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "min_kbps")

    @property
    @pulumi.getter(name="qosPolicyId")
    def qos_policy_id(self) -> pulumi.Output[str]:
        """
        The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "qos_policy_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.
        """
        return pulumi.get(self, "region")

