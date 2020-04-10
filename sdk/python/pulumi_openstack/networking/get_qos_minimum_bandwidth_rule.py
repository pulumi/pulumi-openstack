# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetQosMinimumBandwidthRuleResult:
    """
    A collection of values returned by getQosMinimumBandwidthRule.
    """
    def __init__(__self__, direction=None, id=None, min_kbps=None, qos_policy_id=None, region=None):
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        __self__.direction = direction
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if min_kbps and not isinstance(min_kbps, float):
            raise TypeError("Expected argument 'min_kbps' to be a float")
        __self__.min_kbps = min_kbps
        """
        See Argument Reference above.
        """
        if qos_policy_id and not isinstance(qos_policy_id, str):
            raise TypeError("Expected argument 'qos_policy_id' to be a str")
        __self__.qos_policy_id = qos_policy_id
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
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

def get_qos_minimum_bandwidth_rule(direction=None,min_kbps=None,qos_policy_id=None,region=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack QoS minimum bandwidth rule.




    :param float min_kbps: The value of a minimum kbps bandwidth.
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
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:networking/getQosMinimumBandwidthRule:getQosMinimumBandwidthRule', __args__, opts=opts).value

    return AwaitableGetQosMinimumBandwidthRuleResult(
        direction=__ret__.get('direction'),
        id=__ret__.get('id'),
        min_kbps=__ret__.get('minKbps'),
        qos_policy_id=__ret__.get('qosPolicyId'),
        region=__ret__.get('region'))
