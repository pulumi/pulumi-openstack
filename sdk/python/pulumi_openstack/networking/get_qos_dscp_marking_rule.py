# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetQosDscpMarkingRuleResult:
    """
    A collection of values returned by getQosDscpMarkingRule.
    """
    def __init__(__self__, dscp_mark=None, qos_policy_id=None, region=None, id=None):
        if dscp_mark and not isinstance(dscp_mark, float):
            raise TypeError("Expected argument 'dscp_mark' to be a float")
        __self__.dscp_mark = dscp_mark
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
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_qos_dscp_marking_rule(dscp_mark=None,qos_policy_id=None,region=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack QoS DSCP marking rule.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_dscp_marking_rule_v2.html.markdown.
    """
    __args__ = dict()

    __args__['dscpMark'] = dscp_mark
    __args__['qosPolicyId'] = qos_policy_id
    __args__['region'] = region
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:networking/getQosDscpMarkingRule:getQosDscpMarkingRule', __args__, opts=opts).value

    return GetQosDscpMarkingRuleResult(
        dscp_mark=__ret__.get('dscpMark'),
        qos_policy_id=__ret__.get('qosPolicyId'),
        region=__ret__.get('region'),
        id=__ret__.get('id'))
