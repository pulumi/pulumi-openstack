# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetPolicyResult:
    """
    A collection of values returned by getPolicy.
    """
    def __init__(__self__, audited=None, description=None, name=None, policy_id=None, region=None, rules=None, shared=None, tenant_id=None, id=None):
        if audited and not isinstance(audited, bool):
            raise TypeError("Expected argument 'audited' to be a bool")
        __self__.audited = audited
        """
        The audit status of the firewall policy.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        The description of the firewall policy.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        __self__.policy_id = policy_id
        """
        See Argument Reference above.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        __self__.rules = rules
        """
        The array of one or more firewall rules that comprise the policy.
        """
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        __self__.shared = shared
        """
        The sharing status of the firewall policy.
        """
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        __self__.tenant_id = tenant_id
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_policy(name=None,policy_id=None,region=None,tenant_id=None,opts=None):
    """
    Use this data source to get firewall policy information of an available OpenStack firewall policy.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/fw_policy_v1.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['policyId'] = policy_id
    __args__['region'] = region
    __args__['tenantId'] = tenant_id
    __ret__ = await pulumi.runtime.invoke('openstack:firewall/getPolicy:getPolicy', __args__, opts=opts)

    return GetPolicyResult(
        audited=__ret__.get('audited'),
        description=__ret__.get('description'),
        name=__ret__.get('name'),
        policy_id=__ret__.get('policyId'),
        region=__ret__.get('region'),
        rules=__ret__.get('rules'),
        shared=__ret__.get('shared'),
        tenant_id=__ret__.get('tenantId'),
        id=__ret__.get('id'))
