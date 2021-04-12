# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['RuleArgs', 'Rule']

@pulumi.input_type
class RuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 destination_ip_address: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_version: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source_ip_address: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Rule resource.
        :param pulumi.Input[str] action: Action to be taken ( must be "allow" or "deny") when the
               firewall rule matches. Changing this updates the `action` of an existing
               firewall rule.
        :param pulumi.Input[str] protocol: The protocol type on which the firewall rule operates.
               Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
               `protocol` of an existing firewall rule.
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
        :param pulumi.Input[int] ip_version: IP version, either 4 (default) or 6. Changing this
               updates the `ip_version` of an existing firewall rule.
        :param pulumi.Input[str] name: A unique name for the firewall rule. Changing this
               updates the `name` of an existing firewall rule.
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
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "protocol", protocol)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_ip_address is not None:
            pulumi.set(__self__, "destination_ip_address", destination_ip_address)
        if destination_port is not None:
            pulumi.set(__self__, "destination_port", destination_port)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if source_ip_address is not None:
            pulumi.set(__self__, "source_ip_address", source_ip_address)
        if source_port is not None:
            pulumi.set(__self__, "source_port", source_port)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Action to be taken ( must be "allow" or "deny") when the
        firewall rule matches. Changing this updates the `action` of an existing
        firewall rule.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The protocol type on which the firewall rule operates.
        Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        `protocol` of an existing firewall rule.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the firewall rule. Changing this
        updates the `description` of an existing firewall rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationIpAddress")
    def destination_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The destination IP address on which the
        firewall rule operates. Changing this updates the `destination_ip_address`
        of an existing firewall rule.
        """
        return pulumi.get(self, "destination_ip_address")

    @destination_ip_address.setter
    def destination_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_ip_address", value)

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[pulumi.Input[str]]:
        """
        The destination port on which the firewall
        rule operates. Changing this updates the `destination_port` of an existing
        firewall rule.
        """
        return pulumi.get(self, "destination_port")

    @destination_port.setter
    def destination_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_port", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enabled status for the firewall rule (must be "true"
        or "false" if provided - defaults to "true"). Changing this updates the
        `enabled` status of an existing firewall rule.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[int]]:
        """
        IP version, either 4 (default) or 6. Changing this
        updates the `ip_version` of an existing firewall rule.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the firewall rule. Changing this
        updates the `name` of an existing firewall rule.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the v1 Compute client.
        A Compute client is needed to create a firewall rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sourceIpAddress")
    def source_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The source IP address on which the firewall
        rule operates. Changing this updates the `source_ip_address` of an existing
        firewall rule.
        """
        return pulumi.get(self, "source_ip_address")

    @source_ip_address.setter
    def source_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip_address", value)

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[pulumi.Input[str]]:
        """
        The source port on which the firewall
        rule operates. Changing this updates the `source_port` of an existing
        firewall rule.
        """
        return pulumi.get(self, "source_port")

    @source_port.setter
    def source_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_port", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the firewall rule. Required if admin
        wants to create a firewall rule for another tenant. Changing this creates a
        new firewall rule.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


class Rule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_ip_address: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_version: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source_ip_address: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a v1 firewall rule resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rule1 = openstack.firewall.Rule("rule1",
            action="deny",
            description="drop TELNET traffic",
            destination_port="23",
            enabled=True,
            protocol="tcp")
        ```

        ## Import

        Firewall Rules can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:firewall/rule:Rule rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
        ```

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
        :param pulumi.Input[int] ip_version: IP version, either 4 (default) or 6. Changing this
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
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a v1 firewall rule resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        rule1 = openstack.firewall.Rule("rule1",
            action="deny",
            description="drop TELNET traffic",
            destination_port="23",
            enabled=True,
            protocol="tcp")
        ```

        ## Import

        Firewall Rules can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:firewall/rule:Rule rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327
        ```

        :param str resource_name: The name of the resource.
        :param RuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_ip_address: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 ip_version: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 source_ip_address: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            __props__['description'] = description
            __props__['destination_ip_address'] = destination_ip_address
            __props__['destination_port'] = destination_port
            __props__['enabled'] = enabled
            __props__['ip_version'] = ip_version
            __props__['name'] = name
            if protocol is None and not opts.urn:
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

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination_ip_address: Optional[pulumi.Input[str]] = None,
            destination_port: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            ip_version: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            source_ip_address: Optional[pulumi.Input[str]] = None,
            source_port: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Rule':
        """
        Get an existing Rule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[int] ip_version: IP version, either 4 (default) or 6. Changing this
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
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["action"] = action
        __props__["description"] = description
        __props__["destination_ip_address"] = destination_ip_address
        __props__["destination_port"] = destination_port
        __props__["enabled"] = enabled
        __props__["ip_version"] = ip_version
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["region"] = region
        __props__["source_ip_address"] = source_ip_address
        __props__["source_port"] = source_port
        __props__["tenant_id"] = tenant_id
        __props__["value_specs"] = value_specs
        return Rule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Action to be taken ( must be "allow" or "deny") when the
        firewall rule matches. Changing this updates the `action` of an existing
        firewall rule.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the firewall rule. Changing this
        updates the `description` of an existing firewall rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationIpAddress")
    def destination_ip_address(self) -> pulumi.Output[Optional[str]]:
        """
        The destination IP address on which the
        firewall rule operates. Changing this updates the `destination_ip_address`
        of an existing firewall rule.
        """
        return pulumi.get(self, "destination_ip_address")

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> pulumi.Output[Optional[str]]:
        """
        The destination port on which the firewall
        rule operates. Changing this updates the `destination_port` of an existing
        firewall rule.
        """
        return pulumi.get(self, "destination_port")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enabled status for the firewall rule (must be "true"
        or "false" if provided - defaults to "true"). Changing this updates the
        `enabled` status of an existing firewall rule.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[Optional[int]]:
        """
        IP version, either 4 (default) or 6. Changing this
        updates the `ip_version` of an existing firewall rule.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the firewall rule. Changing this
        updates the `name` of an existing firewall rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol type on which the firewall rule operates.
        Valid values are: `tcp`, `udp`, `icmp`, and `any`. Changing this updates the
        `protocol` of an existing firewall rule.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the v1 Compute client.
        A Compute client is needed to create a firewall rule. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        firewall rule.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sourceIpAddress")
    def source_ip_address(self) -> pulumi.Output[Optional[str]]:
        """
        The source IP address on which the firewall
        rule operates. Changing this updates the `source_ip_address` of an existing
        firewall rule.
        """
        return pulumi.get(self, "source_ip_address")

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> pulumi.Output[Optional[str]]:
        """
        The source port on which the firewall
        rule operates. Changing this updates the `source_port` of an existing
        firewall rule.
        """
        return pulumi.get(self, "source_port")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The owner of the firewall rule. Required if admin
        wants to create a firewall rule for another tenant. Changing this creates a
        new firewall rule.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

