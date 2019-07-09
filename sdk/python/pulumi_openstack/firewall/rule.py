# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Rule(pulumi.CustomResource):
    action: pulumi.Output[str]
    """
    Action to be taken ( must be "allow" or "deny") when the
    firewall rule matches. Changing this updates the `action` of an existing
    firewall rule.
    """
    description: pulumi.Output[str]
    """
    A description for the firewall rule. Changing this
    updates the `description` of an existing firewall rule.
    """
    destination_ip_address: pulumi.Output[str]
    """
    The destination IP address on which the
    firewall rule operates. Changing this updates the `destination_ip_address`
    of an existing firewall rule.
    """
    destination_port: pulumi.Output[str]
    """
    The destination port on which the firewall
    rule operates. Changing this updates the `destination_port` of an existing
    firewall rule.
    """
    enabled: pulumi.Output[bool]
    """
    Enabled status for the firewall rule (must be "true"
    or "false" if provided - defaults to "true"). Changing this updates the
    `enabled` status of an existing firewall rule.
    """
    ip_version: pulumi.Output[float]
    """
    IP version, either 4 (default) or 6. Changing this
    updates the `ip_version` of an existing firewall rule.
    """
    name: pulumi.Output[str]
    """
    A unique name for the firewall rule. Changing this
    updates the `name` of an existing firewall rule.
    """
    protocol: pulumi.Output[str]
    """
    The protocol type on which the firewall rule operates.
    Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
    `protocol` of an existing firewall rule.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the v1 Compute client.
    A Compute client is needed to create a firewall rule. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    firewall rule.
    """
    source_ip_address: pulumi.Output[str]
    """
    The source IP address on which the firewall
    rule operates. Changing this updates the `source_ip_address` of an existing
    firewall rule.
    """
    source_port: pulumi.Output[str]
    """
    The source port on which the firewall
    rule operates. Changing this updates the `source_port` of an existing
    firewall rule.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the firewall rule. Required if admin
    wants to create a firewall rule for another tenant. Changing this creates a
    new firewall rule.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, action=None, description=None, destination_ip_address=None, destination_port=None, enabled=None, ip_version=None, name=None, protocol=None, region=None, source_ip_address=None, source_port=None, tenant_id=None, value_specs=None, __name__=None, __opts__=None):
        """
        Manages a v1 firewall rule resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action to be taken ( must be "allow" or "deny") when the
               firewall rule matches. Changing this updates the `action` of an existing
               firewall rule.
        :param pulumi.Input[str] description: A description for the firewall rule. Changing this
               updates the `description` of an existing firewall rule.
        :param pulumi.Input[str] destination_ip_address: The destination IP address on which the
               firewall rule operates. Changing this updates the `destination_ip_address`
               of an existing firewall rule.
        :param pulumi.Input[str] destination_port: The destination port on which the firewall
               rule operates. Changing this updates the `destination_port` of an existing
               firewall rule.
        :param pulumi.Input[bool] enabled: Enabled status for the firewall rule (must be "true"
               or "false" if provided - defaults to "true"). Changing this updates the
               `enabled` status of an existing firewall rule.
        :param pulumi.Input[float] ip_version: IP version, either 4 (default) or 6. Changing this
               updates the `ip_version` of an existing firewall rule.
        :param pulumi.Input[str] name: A unique name for the firewall rule. Changing this
               updates the `name` of an existing firewall rule.
        :param pulumi.Input[str] protocol: The protocol type on which the firewall rule operates.
               Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
               `protocol` of an existing firewall rule.
        :param pulumi.Input[str] region: The region in which to obtain the v1 Compute client.
               A Compute client is needed to create a firewall rule. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               firewall rule.
        :param pulumi.Input[str] source_ip_address: The source IP address on which the firewall
               rule operates. Changing this updates the `source_ip_address` of an existing
               firewall rule.
        :param pulumi.Input[str] source_port: The source port on which the firewall
               rule operates. Changing this updates the `source_port` of an existing
               firewall rule.
        :param pulumi.Input[str] tenant_id: The owner of the firewall rule. Required if admin
               wants to create a firewall rule for another tenant. Changing this creates a
               new firewall rule.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/fw_rule_v1.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if action is None:
            raise TypeError("Missing required property 'action'")
        __props__['action'] = action

        __props__['description'] = description

        __props__['destination_ip_address'] = destination_ip_address

        __props__['destination_port'] = destination_port

        __props__['enabled'] = enabled

        __props__['ip_version'] = ip_version

        __props__['name'] = name

        if protocol is None:
            raise TypeError("Missing required property 'protocol'")
        __props__['protocol'] = protocol

        __props__['region'] = region

        __props__['source_ip_address'] = source_ip_address

        __props__['source_port'] = source_port

        __props__['tenant_id'] = tenant_id

        __props__['value_specs'] = value_specs

        super(Rule, __self__).__init__(
            'openstack:firewall/rule:Rule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

