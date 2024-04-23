# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['L7RuleV2Args', 'L7RuleV2']

@pulumi.input_type
class L7RuleV2Args:
    def __init__(__self__, *,
                 compare_type: pulumi.Input[str],
                 l7policy_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 value: pulumi.Input[str],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 invert: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a L7RuleV2 resource.
        :param pulumi.Input[str] compare_type: The comparison type for the L7 rule - can either be
               CONTAINS, STARTS\\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        :param pulumi.Input[str] l7policy_id: The ID of the L7 Policy to query. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] type: The L7 Rule type - can either be COOKIE, FILE\\_TYPE, HEADER,
               HOST\\_NAME or PATH.
        :param pulumi.Input[str] value: The value to use for the comparison. For example, the file type to
               compare.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the L7 Rule.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[bool] invert: When true the logic of the rule is inverted. For example, with invert
               true, equal to would become not equal to. Default is false.
        :param pulumi.Input[str] key: The key to use for the comparison. For example, the name of the cookie to
               evaluate. Valid when `type` is set to COOKIE or HEADER.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the L7 Rule.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new L7 Rule.
        """
        pulumi.set(__self__, "compare_type", compare_type)
        pulumi.set(__self__, "l7policy_id", l7policy_id)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if invert is not None:
            pulumi.set(__self__, "invert", invert)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="compareType")
    def compare_type(self) -> pulumi.Input[str]:
        """
        The comparison type for the L7 rule - can either be
        CONTAINS, STARTS\\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        """
        return pulumi.get(self, "compare_type")

    @compare_type.setter
    def compare_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "compare_type", value)

    @property
    @pulumi.getter(name="l7policyId")
    def l7policy_id(self) -> pulumi.Input[str]:
        """
        The ID of the L7 Policy to query. Changing this creates a new
        L7 Rule.
        """
        return pulumi.get(self, "l7policy_id")

    @l7policy_id.setter
    def l7policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "l7policy_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The L7 Rule type - can either be COOKIE, FILE\\_TYPE, HEADER,
        HOST\\_NAME or PATH.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value to use for the comparison. For example, the file type to
        compare.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the L7 Rule.
        A valid value is true (UP) or false (DOWN).
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def invert(self) -> Optional[pulumi.Input[bool]]:
        """
        When true the logic of the rule is inverted. For example, with invert
        true, equal to would become not equal to. Default is false.
        """
        return pulumi.get(self, "invert")

    @invert.setter
    def invert(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "invert", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The key to use for the comparison. For example, the name of the cookie to
        evaluate. Valid when `type` is set to COOKIE or HEADER.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an . If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        L7 Rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required for admins. The UUID of the tenant who owns
        the L7 Rule.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new L7 Rule.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _L7RuleV2State:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 compare_type: Optional[pulumi.Input[str]] = None,
                 invert: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 l7policy_id: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering L7RuleV2 resources.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the L7 Rule.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[str] compare_type: The comparison type for the L7 rule - can either be
               CONTAINS, STARTS\\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        :param pulumi.Input[bool] invert: When true the logic of the rule is inverted. For example, with invert
               true, equal to would become not equal to. Default is false.
        :param pulumi.Input[str] key: The key to use for the comparison. For example, the name of the cookie to
               evaluate. Valid when `type` is set to COOKIE or HEADER.
        :param pulumi.Input[str] l7policy_id: The ID of the L7 Policy to query. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] listener_id: The ID of the Listener owning this resource.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the L7 Rule.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new L7 Rule.
        :param pulumi.Input[str] type: The L7 Rule type - can either be COOKIE, FILE\\_TYPE, HEADER,
               HOST\\_NAME or PATH.
        :param pulumi.Input[str] value: The value to use for the comparison. For example, the file type to
               compare.
        """
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if compare_type is not None:
            pulumi.set(__self__, "compare_type", compare_type)
        if invert is not None:
            pulumi.set(__self__, "invert", invert)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if l7policy_id is not None:
            pulumi.set(__self__, "l7policy_id", l7policy_id)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the L7 Rule.
        A valid value is true (UP) or false (DOWN).
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="compareType")
    def compare_type(self) -> Optional[pulumi.Input[str]]:
        """
        The comparison type for the L7 rule - can either be
        CONTAINS, STARTS\\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        """
        return pulumi.get(self, "compare_type")

    @compare_type.setter
    def compare_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compare_type", value)

    @property
    @pulumi.getter
    def invert(self) -> Optional[pulumi.Input[bool]]:
        """
        When true the logic of the rule is inverted. For example, with invert
        true, equal to would become not equal to. Default is false.
        """
        return pulumi.get(self, "invert")

    @invert.setter
    def invert(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "invert", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The key to use for the comparison. For example, the name of the cookie to
        evaluate. Valid when `type` is set to COOKIE or HEADER.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="l7policyId")
    def l7policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the L7 Policy to query. Changing this creates a new
        L7 Rule.
        """
        return pulumi.get(self, "l7policy_id")

    @l7policy_id.setter
    def l7policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "l7policy_id", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Listener owning this resource.
        """
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an . If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        L7 Rule.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required for admins. The UUID of the tenant who owns
        the L7 Rule.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new L7 Rule.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The L7 Rule type - can either be COOKIE, FILE\\_TYPE, HEADER,
        HOST\\_NAME or PATH.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The value to use for the comparison. For example, the file type to
        compare.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class L7RuleV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 compare_type: Optional[pulumi.Input[str]] = None,
                 invert: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 l7policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V2 L7 Rule resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network_1",
            name="network_1",
            admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet_1",
            name="subnet_1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        loadbalancer1 = openstack.loadbalancer.LoadBalancer("loadbalancer_1",
            name="loadbalancer_1",
            vip_subnet_id=subnet1.id)
        listener1 = openstack.loadbalancer.Listener("listener_1",
            name="listener_1",
            protocol="HTTP",
            protocol_port=8080,
            loadbalancer_id=loadbalancer1.id)
        pool1 = openstack.loadbalancer.Pool("pool_1",
            name="pool_1",
            protocol="HTTP",
            lb_method="ROUND_ROBIN",
            loadbalancer_id=loadbalancer1.id)
        l7policy1 = openstack.loadbalancer.L7PolicyV2("l7policy_1",
            name="test",
            action="REDIRECT_TO_URL",
            description="test description",
            position=1,
            listener_id=listener1.id,
            redirect_url="http://www.example.com")
        l7rule1 = openstack.loadbalancer.L7RuleV2("l7rule_1",
            l7policy_id=l7policy1.id,
            type="PATH",
            compare_type="EQUAL_TO",
            value="/api")
        ```

        ## Import

        Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID
        separated by a slash, e.g.:

        ```sh
        $ pulumi import openstack:loadbalancer/l7RuleV2:L7RuleV2 l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the L7 Rule.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[str] compare_type: The comparison type for the L7 rule - can either be
               CONTAINS, STARTS\\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        :param pulumi.Input[bool] invert: When true the logic of the rule is inverted. For example, with invert
               true, equal to would become not equal to. Default is false.
        :param pulumi.Input[str] key: The key to use for the comparison. For example, the name of the cookie to
               evaluate. Valid when `type` is set to COOKIE or HEADER.
        :param pulumi.Input[str] l7policy_id: The ID of the L7 Policy to query. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the L7 Rule.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new L7 Rule.
        :param pulumi.Input[str] type: The L7 Rule type - can either be COOKIE, FILE\\_TYPE, HEADER,
               HOST\\_NAME or PATH.
        :param pulumi.Input[str] value: The value to use for the comparison. For example, the file type to
               compare.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: L7RuleV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 L7 Rule resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network_1",
            name="network_1",
            admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet_1",
            name="subnet_1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        loadbalancer1 = openstack.loadbalancer.LoadBalancer("loadbalancer_1",
            name="loadbalancer_1",
            vip_subnet_id=subnet1.id)
        listener1 = openstack.loadbalancer.Listener("listener_1",
            name="listener_1",
            protocol="HTTP",
            protocol_port=8080,
            loadbalancer_id=loadbalancer1.id)
        pool1 = openstack.loadbalancer.Pool("pool_1",
            name="pool_1",
            protocol="HTTP",
            lb_method="ROUND_ROBIN",
            loadbalancer_id=loadbalancer1.id)
        l7policy1 = openstack.loadbalancer.L7PolicyV2("l7policy_1",
            name="test",
            action="REDIRECT_TO_URL",
            description="test description",
            position=1,
            listener_id=listener1.id,
            redirect_url="http://www.example.com")
        l7rule1 = openstack.loadbalancer.L7RuleV2("l7rule_1",
            l7policy_id=l7policy1.id,
            type="PATH",
            compare_type="EQUAL_TO",
            value="/api")
        ```

        ## Import

        Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID
        separated by a slash, e.g.:

        ```sh
        $ pulumi import openstack:loadbalancer/l7RuleV2:L7RuleV2 l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e
        ```

        :param str resource_name: The name of the resource.
        :param L7RuleV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(L7RuleV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 compare_type: Optional[pulumi.Input[str]] = None,
                 invert: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 l7policy_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = L7RuleV2Args.__new__(L7RuleV2Args)

            __props__.__dict__["admin_state_up"] = admin_state_up
            if compare_type is None and not opts.urn:
                raise TypeError("Missing required property 'compare_type'")
            __props__.__dict__["compare_type"] = compare_type
            __props__.__dict__["invert"] = invert
            __props__.__dict__["key"] = key
            if l7policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'l7policy_id'")
            __props__.__dict__["l7policy_id"] = l7policy_id
            __props__.__dict__["region"] = region
            __props__.__dict__["tenant_id"] = tenant_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["listener_id"] = None
        super(L7RuleV2, __self__).__init__(
            'openstack:loadbalancer/l7RuleV2:L7RuleV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            compare_type: Optional[pulumi.Input[str]] = None,
            invert: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            l7policy_id: Optional[pulumi.Input[str]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'L7RuleV2':
        """
        Get an existing L7RuleV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the L7 Rule.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[str] compare_type: The comparison type for the L7 rule - can either be
               CONTAINS, STARTS\\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        :param pulumi.Input[bool] invert: When true the logic of the rule is inverted. For example, with invert
               true, equal to would become not equal to. Default is false.
        :param pulumi.Input[str] key: The key to use for the comparison. For example, the name of the cookie to
               evaluate. Valid when `type` is set to COOKIE or HEADER.
        :param pulumi.Input[str] l7policy_id: The ID of the L7 Policy to query. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] listener_id: The ID of the Listener owning this resource.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               L7 Rule.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the L7 Rule.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new L7 Rule.
        :param pulumi.Input[str] type: The L7 Rule type - can either be COOKIE, FILE\\_TYPE, HEADER,
               HOST\\_NAME or PATH.
        :param pulumi.Input[str] value: The value to use for the comparison. For example, the file type to
               compare.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _L7RuleV2State.__new__(_L7RuleV2State)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["compare_type"] = compare_type
        __props__.__dict__["invert"] = invert
        __props__.__dict__["key"] = key
        __props__.__dict__["l7policy_id"] = l7policy_id
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["type"] = type
        __props__.__dict__["value"] = value
        return L7RuleV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        """
        The administrative state of the L7 Rule.
        A valid value is true (UP) or false (DOWN).
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="compareType")
    def compare_type(self) -> pulumi.Output[str]:
        """
        The comparison type for the L7 rule - can either be
        CONTAINS, STARTS\\_WITH, ENDS_WITH, EQUAL_TO or REGEX
        """
        return pulumi.get(self, "compare_type")

    @property
    @pulumi.getter
    def invert(self) -> pulumi.Output[Optional[bool]]:
        """
        When true the logic of the rule is inverted. For example, with invert
        true, equal to would become not equal to. Default is false.
        """
        return pulumi.get(self, "invert")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[str]]:
        """
        The key to use for the comparison. For example, the name of the cookie to
        evaluate. Valid when `type` is set to COOKIE or HEADER.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="l7policyId")
    def l7policy_id(self) -> pulumi.Output[str]:
        """
        The ID of the L7 Policy to query. Changing this creates a new
        L7 Rule.
        """
        return pulumi.get(self, "l7policy_id")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[str]:
        """
        The ID of the Listener owning this resource.
        """
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an . If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        L7 Rule.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Required for admins. The UUID of the tenant who owns
        the L7 Rule.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new L7 Rule.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The L7 Rule type - can either be COOKIE, FILE\\_TYPE, HEADER,
        HOST\\_NAME or PATH.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        The value to use for the comparison. For example, the file type to
        compare.
        """
        return pulumi.get(self, "value")

