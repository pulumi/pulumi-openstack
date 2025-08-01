# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'GetFwRuleV2Result',
    'AwaitableGetFwRuleV2Result',
    'get_fw_rule_v2',
    'get_fw_rule_v2_output',
]

warnings.warn("""openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2""", DeprecationWarning)

@pulumi.output_type
class GetFwRuleV2Result:
    """
    A collection of values returned by getFwRuleV2.
    """
    def __init__(__self__, action=None, description=None, destination_ip_address=None, destination_port=None, enabled=None, firewall_policy_ids=None, id=None, ip_version=None, name=None, project_id=None, protocol=None, region=None, rule_id=None, shared=None, source_ip_address=None, source_port=None, tenant_id=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if destination_ip_address and not isinstance(destination_ip_address, str):
            raise TypeError("Expected argument 'destination_ip_address' to be a str")
        pulumi.set(__self__, "destination_ip_address", destination_ip_address)
        if destination_port and not isinstance(destination_port, str):
            raise TypeError("Expected argument 'destination_port' to be a str")
        pulumi.set(__self__, "destination_port", destination_port)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if firewall_policy_ids and not isinstance(firewall_policy_ids, list):
            raise TypeError("Expected argument 'firewall_policy_ids' to be a list")
        pulumi.set(__self__, "firewall_policy_ids", firewall_policy_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_version and not isinstance(ip_version, int):
            raise TypeError("Expected argument 'ip_version' to be a int")
        pulumi.set(__self__, "ip_version", ip_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if rule_id and not isinstance(rule_id, str):
            raise TypeError("Expected argument 'rule_id' to be a str")
        pulumi.set(__self__, "rule_id", rule_id)
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        pulumi.set(__self__, "shared", shared)
        if source_ip_address and not isinstance(source_ip_address, str):
            raise TypeError("Expected argument 'source_ip_address' to be a str")
        pulumi.set(__self__, "source_ip_address", source_ip_address)
        if source_port and not isinstance(source_port, str):
            raise TypeError("Expected argument 'source_port' to be a str")
        pulumi.set(__self__, "source_port", source_port)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)

    @_builtins.property
    @pulumi.getter
    def action(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "action")

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter(name="destinationIpAddress")
    def destination_ip_address(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "destination_ip_address")

    @_builtins.property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "destination_port")

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> _builtins.bool:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "enabled")

    @_builtins.property
    @pulumi.getter(name="firewallPolicyIds")
    def firewall_policy_ids(self) -> Sequence[_builtins.str]:
        """
        The ID of the firewall policy the rule belongs to.
        """
        return pulumi.get(self, "firewall_policy_ids")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[_builtins.int]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "ip_version")

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> _builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter
    def protocol(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "protocol")

    @_builtins.property
    @pulumi.getter
    def region(self) -> _builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "region")

    @_builtins.property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "rule_id")

    @_builtins.property
    @pulumi.getter
    def shared(self) -> _builtins.bool:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "shared")

    @_builtins.property
    @pulumi.getter(name="sourceIpAddress")
    def source_ip_address(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "source_ip_address")

    @_builtins.property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[_builtins.str]:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "source_port")

    @_builtins.property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> _builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "tenant_id")


class AwaitableGetFwRuleV2Result(GetFwRuleV2Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFwRuleV2Result(
            action=self.action,
            description=self.description,
            destination_ip_address=self.destination_ip_address,
            destination_port=self.destination_port,
            enabled=self.enabled,
            firewall_policy_ids=self.firewall_policy_ids,
            id=self.id,
            ip_version=self.ip_version,
            name=self.name,
            project_id=self.project_id,
            protocol=self.protocol,
            region=self.region,
            rule_id=self.rule_id,
            shared=self.shared,
            source_ip_address=self.source_ip_address,
            source_port=self.source_port,
            tenant_id=self.tenant_id)


def get_fw_rule_v2(action: Optional[_builtins.str] = None,
                   description: Optional[_builtins.str] = None,
                   destination_ip_address: Optional[_builtins.str] = None,
                   destination_port: Optional[_builtins.str] = None,
                   enabled: Optional[_builtins.bool] = None,
                   firewall_policy_ids: Optional[Sequence[_builtins.str]] = None,
                   ip_version: Optional[_builtins.int] = None,
                   name: Optional[_builtins.str] = None,
                   project_id: Optional[_builtins.str] = None,
                   protocol: Optional[_builtins.str] = None,
                   region: Optional[_builtins.str] = None,
                   rule_id: Optional[_builtins.str] = None,
                   shared: Optional[_builtins.bool] = None,
                   source_ip_address: Optional[_builtins.str] = None,
                   source_port: Optional[_builtins.str] = None,
                   tenant_id: Optional[_builtins.str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFwRuleV2Result:
    """
    Use this data source to get information of an available OpenStack firewall rule v2.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    rule = openstack.firewall.get_rule_v2(name="tf_test_rule")
    ```


    :param _builtins.str action: Action to be taken when the firewall rule matches.
    :param _builtins.str description: The description of the firewall rule.
    :param _builtins.str destination_ip_address: The destination IP address on which the
           firewall rule operates.
    :param _builtins.str destination_port: The destination port on which the firewall
           rule operates.
    :param _builtins.bool enabled: Enabled status for the firewall rule.
    :param Sequence[_builtins.str] firewall_policy_ids: The ID of the firewall policy the rule belongs to.
    :param _builtins.int ip_version: IP version, either 4 (default) or 6.
    :param _builtins.str name: The name of the firewall rule.
    :param _builtins.str project_id: This argument conflicts and is interchangeable
           with `tenant_id`. The owner of the firewall rule.
    :param _builtins.str protocol: The protocol type on which the firewall rule operates.
    :param _builtins.str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve firewall policy ids. If omitted, the
           `region` argument of the provider is used.
    :param _builtins.str rule_id: The ID of the firewall rule.
    :param _builtins.bool shared: The sharing status of the firewall policy.
    :param _builtins.str source_ip_address: The source IP address on which the firewall
           rule operates.
    :param _builtins.str source_port: The source port on which the firewall
           rule operates.
    :param _builtins.str tenant_id: This argument conflicts and is interchangeable
           with `project_id`. The owner of the firewall rule.
    """
    pulumi.log.warn("""get_fw_rule_v2 is deprecated: openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2""")
    __args__ = dict()
    __args__['action'] = action
    __args__['description'] = description
    __args__['destinationIpAddress'] = destination_ip_address
    __args__['destinationPort'] = destination_port
    __args__['enabled'] = enabled
    __args__['firewallPolicyIds'] = firewall_policy_ids
    __args__['ipVersion'] = ip_version
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['protocol'] = protocol
    __args__['region'] = region
    __args__['ruleId'] = rule_id
    __args__['shared'] = shared
    __args__['sourceIpAddress'] = source_ip_address
    __args__['sourcePort'] = source_port
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('openstack:index/getFwRuleV2:getFwRuleV2', __args__, opts=opts, typ=GetFwRuleV2Result).value

    return AwaitableGetFwRuleV2Result(
        action=pulumi.get(__ret__, 'action'),
        description=pulumi.get(__ret__, 'description'),
        destination_ip_address=pulumi.get(__ret__, 'destination_ip_address'),
        destination_port=pulumi.get(__ret__, 'destination_port'),
        enabled=pulumi.get(__ret__, 'enabled'),
        firewall_policy_ids=pulumi.get(__ret__, 'firewall_policy_ids'),
        id=pulumi.get(__ret__, 'id'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        protocol=pulumi.get(__ret__, 'protocol'),
        region=pulumi.get(__ret__, 'region'),
        rule_id=pulumi.get(__ret__, 'rule_id'),
        shared=pulumi.get(__ret__, 'shared'),
        source_ip_address=pulumi.get(__ret__, 'source_ip_address'),
        source_port=pulumi.get(__ret__, 'source_port'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'))
def get_fw_rule_v2_output(action: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          description: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          destination_ip_address: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          destination_port: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          enabled: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                          firewall_policy_ids: Optional[pulumi.Input[Optional[Sequence[_builtins.str]]]] = None,
                          ip_version: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                          name: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          project_id: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          protocol: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          region: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          rule_id: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          shared: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                          source_ip_address: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          source_port: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          tenant_id: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFwRuleV2Result]:
    """
    Use this data source to get information of an available OpenStack firewall rule v2.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_openstack as openstack

    rule = openstack.firewall.get_rule_v2(name="tf_test_rule")
    ```


    :param _builtins.str action: Action to be taken when the firewall rule matches.
    :param _builtins.str description: The description of the firewall rule.
    :param _builtins.str destination_ip_address: The destination IP address on which the
           firewall rule operates.
    :param _builtins.str destination_port: The destination port on which the firewall
           rule operates.
    :param _builtins.bool enabled: Enabled status for the firewall rule.
    :param Sequence[_builtins.str] firewall_policy_ids: The ID of the firewall policy the rule belongs to.
    :param _builtins.int ip_version: IP version, either 4 (default) or 6.
    :param _builtins.str name: The name of the firewall rule.
    :param _builtins.str project_id: This argument conflicts and is interchangeable
           with `tenant_id`. The owner of the firewall rule.
    :param _builtins.str protocol: The protocol type on which the firewall rule operates.
    :param _builtins.str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve firewall policy ids. If omitted, the
           `region` argument of the provider is used.
    :param _builtins.str rule_id: The ID of the firewall rule.
    :param _builtins.bool shared: The sharing status of the firewall policy.
    :param _builtins.str source_ip_address: The source IP address on which the firewall
           rule operates.
    :param _builtins.str source_port: The source port on which the firewall
           rule operates.
    :param _builtins.str tenant_id: This argument conflicts and is interchangeable
           with `project_id`. The owner of the firewall rule.
    """
    pulumi.log.warn("""get_fw_rule_v2 is deprecated: openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2""")
    __args__ = dict()
    __args__['action'] = action
    __args__['description'] = description
    __args__['destinationIpAddress'] = destination_ip_address
    __args__['destinationPort'] = destination_port
    __args__['enabled'] = enabled
    __args__['firewallPolicyIds'] = firewall_policy_ids
    __args__['ipVersion'] = ip_version
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['protocol'] = protocol
    __args__['region'] = region
    __args__['ruleId'] = rule_id
    __args__['shared'] = shared
    __args__['sourceIpAddress'] = source_ip_address
    __args__['sourcePort'] = source_port
    __args__['tenantId'] = tenant_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('openstack:index/getFwRuleV2:getFwRuleV2', __args__, opts=opts, typ=GetFwRuleV2Result)
    return __ret__.apply(lambda __response__: GetFwRuleV2Result(
        action=pulumi.get(__response__, 'action'),
        description=pulumi.get(__response__, 'description'),
        destination_ip_address=pulumi.get(__response__, 'destination_ip_address'),
        destination_port=pulumi.get(__response__, 'destination_port'),
        enabled=pulumi.get(__response__, 'enabled'),
        firewall_policy_ids=pulumi.get(__response__, 'firewall_policy_ids'),
        id=pulumi.get(__response__, 'id'),
        ip_version=pulumi.get(__response__, 'ip_version'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        protocol=pulumi.get(__response__, 'protocol'),
        region=pulumi.get(__response__, 'region'),
        rule_id=pulumi.get(__response__, 'rule_id'),
        shared=pulumi.get(__response__, 'shared'),
        source_ip_address=pulumi.get(__response__, 'source_ip_address'),
        source_port=pulumi.get(__response__, 'source_port'),
        tenant_id=pulumi.get(__response__, 'tenant_id')))
