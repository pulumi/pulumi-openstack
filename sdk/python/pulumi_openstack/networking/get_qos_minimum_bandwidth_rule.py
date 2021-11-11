# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetQosMinimumBandwidthRuleResult',
    'AwaitableGetQosMinimumBandwidthRuleResult',
    'get_qos_minimum_bandwidth_rule',
    'get_qos_minimum_bandwidth_rule_output',
]

@pulumi.output_type
class GetQosMinimumBandwidthRuleResult:
    """
    A collection of values returned by getQosMinimumBandwidthRule.
    """
    def __init__(__self__, direction=None, id=None, min_kbps=None, qos_policy_id=None, region=None):
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if min_kbps and not isinstance(min_kbps, int):
            raise TypeError("Expected argument 'min_kbps' to be a int")
        pulumi.set(__self__, "min_kbps", min_kbps)
        if qos_policy_id and not isinstance(qos_policy_id, str):
            raise TypeError("Expected argument 'qos_policy_id' to be a str")
        pulumi.set(__self__, "qos_policy_id", qos_policy_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def direction(self) -> str:
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="minKbps")
    def min_kbps(self) -> int:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "min_kbps")

    @property
    @pulumi.getter(name="qosPolicyId")
    def qos_policy_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "qos_policy_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")


class AwaitableGetQosMinimumBandwidthRuleResult(GetQosMinimumBandwidthRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQosMinimumBandwidthRuleResult(
            direction=self.direction,
            id=self.id,
            min_kbps=self.min_kbps,
            qos_policy_id=self.qos_policy_id,
            region=self.region)


def get_qos_minimum_bandwidth_rule(direction: Optional[str] = None,
                                   min_kbps: Optional[int] = None,
                                   qos_policy_id: Optional[str] = None,
                                   region: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQosMinimumBandwidthRuleResult:
    """
    Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    qos_min_bw_rule1 = openstack.networking.get_qos_minimum_bandwidth_rule(min_kbps=2000)
    ```


    :param int min_kbps: The value of a minimum kbps bandwidth.
    :param str qos_policy_id: The QoS policy reference.
    :param str region: The region in which to obtain the V2 Networking client.
           A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
           `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['direction'] = direction
    __args__['minKbps'] = min_kbps
    __args__['qosPolicyId'] = qos_policy_id
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule', __args__, opts=opts, typ=GetQosMinimumBandwidthRuleResult).value

    return AwaitableGetQosMinimumBandwidthRuleResult(
        direction=__ret__.direction,
        id=__ret__.id,
        min_kbps=__ret__.min_kbps,
        qos_policy_id=__ret__.qos_policy_id,
        region=__ret__.region)


@_utilities.lift_output_func(get_qos_minimum_bandwidth_rule)
def get_qos_minimum_bandwidth_rule_output(direction: Optional[pulumi.Input[Optional[str]]] = None,
                                          min_kbps: Optional[pulumi.Input[Optional[int]]] = None,
                                          qos_policy_id: Optional[pulumi.Input[str]] = None,
                                          region: Optional[pulumi.Input[Optional[str]]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQosMinimumBandwidthRuleResult]:
    """
    Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    qos_min_bw_rule1 = openstack.networking.get_qos_minimum_bandwidth_rule(min_kbps=2000)
    ```


    :param int min_kbps: The value of a minimum kbps bandwidth.
    :param str qos_policy_id: The QoS policy reference.
    :param str region: The region in which to obtain the V2 Networking client.
           A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the
           `region` argument of the provider is used.
    """
    ...
