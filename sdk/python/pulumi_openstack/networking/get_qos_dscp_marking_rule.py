# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetQosDscpMarkingRuleResult',
    'AwaitableGetQosDscpMarkingRuleResult',
    'get_qos_dscp_marking_rule',
    'get_qos_dscp_marking_rule_output',
]

@pulumi.output_type
class GetQosDscpMarkingRuleResult:
    """
    A collection of values returned by getQosDscpMarkingRule.
    """
    def __init__(__self__, dscp_mark=None, id=None, qos_policy_id=None, region=None):
        if dscp_mark and not isinstance(dscp_mark, int):
            raise TypeError("Expected argument 'dscp_mark' to be a int")
        pulumi.set(__self__, "dscp_mark", dscp_mark)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if qos_policy_id and not isinstance(qos_policy_id, str):
            raise TypeError("Expected argument 'qos_policy_id' to be a str")
        pulumi.set(__self__, "qos_policy_id", qos_policy_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="dscpMark")
    def dscp_mark(self) -> int:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "dscp_mark")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

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


class AwaitableGetQosDscpMarkingRuleResult(GetQosDscpMarkingRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQosDscpMarkingRuleResult(
            dscp_mark=self.dscp_mark,
            id=self.id,
            qos_policy_id=self.qos_policy_id,
            region=self.region)


def get_qos_dscp_marking_rule(dscp_mark: Optional[int] = None,
                              qos_policy_id: Optional[str] = None,
                              region: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQosDscpMarkingRuleResult:
    """
    Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    qos_dscp_marking_rule1 = openstack.networking.get_qos_dscp_marking_rule(dscp_mark=26)
    ```


    :param int dscp_mark: The value of a DSCP mark.
    :param str qos_policy_id: The QoS policy reference.
    :param str region: The region in which to obtain the V2 Networking client.
           A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
           `region` argument of the provider is used.
    """
    __args__ = dict()
    __args__['dscpMark'] = dscp_mark
    __args__['qosPolicyId'] = qos_policy_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule', __args__, opts=opts, typ=GetQosDscpMarkingRuleResult).value

    return AwaitableGetQosDscpMarkingRuleResult(
        dscp_mark=pulumi.get(__ret__, 'dscp_mark'),
        id=pulumi.get(__ret__, 'id'),
        qos_policy_id=pulumi.get(__ret__, 'qos_policy_id'),
        region=pulumi.get(__ret__, 'region'))


@_utilities.lift_output_func(get_qos_dscp_marking_rule)
def get_qos_dscp_marking_rule_output(dscp_mark: Optional[pulumi.Input[Optional[int]]] = None,
                                     qos_policy_id: Optional[pulumi.Input[str]] = None,
                                     region: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQosDscpMarkingRuleResult]:
    """
    Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    qos_dscp_marking_rule1 = openstack.networking.get_qos_dscp_marking_rule(dscp_mark=26)
    ```


    :param int dscp_mark: The value of a DSCP mark.
    :param str qos_policy_id: The QoS policy reference.
    :param str region: The region in which to obtain the V2 Networking client.
           A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the
           `region` argument of the provider is used.
    """
    ...
