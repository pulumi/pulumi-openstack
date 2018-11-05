# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class SecGroup(pulumi.CustomResource):
    """
    Manages a V2 neutron security group resource within OpenStack.
    Unlike Nova security groups, neutron separates the group from the rules
    and also allows an admin to target a specific tenant_id.
    """
    def __init__(__self__, __name__, __opts__=None, delete_default_rules=None, description=None, name=None, region=None, tenant_id=None):
        """Create a SecGroup resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['deleteDefaultRules'] = delete_default_rules

        __props__['description'] = description

        __props__['name'] = name

        __props__['region'] = region

        __props__['tenantId'] = tenant_id

        super(SecGroup, __self__).__init__(
            'openstack:networking/secGroup:SecGroup',
            __name__,
            __props__,
            __opts__)

