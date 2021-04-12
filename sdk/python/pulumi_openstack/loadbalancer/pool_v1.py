# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['PoolV1Args', 'PoolV1']

@pulumi.input_type
class PoolV1Args:
    def __init__(__self__, *,
                 lb_method: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 lb_provider: Optional[pulumi.Input[str]] = None,
                 monitor_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PoolV1 resource.
        :param pulumi.Input[str] lb_method: The algorithm used to distribute load between the
               members of the pool. The current specification supports 'ROUND_ROBIN' and
               'LEAST_CONNECTIONS' as valid values for this attribute.
        :param pulumi.Input[str] protocol: The protocol used by the pool members, you can use
               either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        :param pulumi.Input[str] subnet_id: The network on which the members of the pool will be
               located. Only members that are on this network can be added to the pool.
               Changing this creates a new pool.
        :param pulumi.Input[str] lb_provider: The backend load balancing provider. For example:
               `haproxy`, `F5`, etc.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitor_ids: A list of IDs of monitors to associate with the
               pool.
        :param pulumi.Input[str] name: The name of the pool. Changing this updates the name of
               the existing pool.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB pool. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB pool.
        :param pulumi.Input[str] tenant_id: The owner of the member. Required if admin wants to
               create a pool member for another tenant. Changing this creates a new member.
        """
        pulumi.set(__self__, "lb_method", lb_method)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if lb_provider is not None:
            pulumi.set(__self__, "lb_provider", lb_provider)
        if monitor_ids is not None:
            pulumi.set(__self__, "monitor_ids", monitor_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> pulumi.Input[str]:
        """
        The algorithm used to distribute load between the
        members of the pool. The current specification supports 'ROUND_ROBIN' and
        'LEAST_CONNECTIONS' as valid values for this attribute.
        """
        return pulumi.get(self, "lb_method")

    @lb_method.setter
    def lb_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "lb_method", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        The protocol used by the pool members, you can use
        either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The network on which the members of the pool will be
        located. Only members that are on this network can be added to the pool.
        Changing this creates a new pool.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="lbProvider")
    def lb_provider(self) -> Optional[pulumi.Input[str]]:
        """
        The backend load balancing provider. For example:
        `haproxy`, `F5`, etc.
        """
        return pulumi.get(self, "lb_provider")

    @lb_provider.setter
    def lb_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lb_provider", value)

    @property
    @pulumi.getter(name="monitorIds")
    def monitor_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of IDs of monitors to associate with the
        pool.
        """
        return pulumi.get(self, "monitor_ids")

    @monitor_ids.setter
    def monitor_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "monitor_ids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the pool. Changing this updates the name of
        the existing pool.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB pool. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB pool.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the member. Required if admin wants to
        create a pool member for another tenant. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class PoolV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lb_method: Optional[pulumi.Input[str]] = None,
                 lb_provider: Optional[pulumi.Input[str]] = None,
                 monitor_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V1 load balancer pool resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        pool1 = openstack.loadbalancer.PoolV1("pool1",
            lb_method="ROUND_ROBIN",
            lb_provider="haproxy",
            monitor_ids=["67890"],
            protocol="HTTP",
            subnet_id="12345")
        ```
        ## Complete Load Balancing Stack Example

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        secgroup1 = openstack.compute.SecGroup("secgroup1",
            description="Rules for secgroup_1",
            rules=[
                openstack.compute.SecGroupRuleArgs(
                    cidr="0.0.0.0/0",
                    from_port=-1,
                    ip_protocol="icmp",
                    to_port=-1,
                ),
                openstack.compute.SecGroupRuleArgs(
                    cidr="0.0.0.0/0",
                    from_port=80,
                    ip_protocol="tcp",
                    to_port=80,
                ),
            ])
        instance1 = openstack.compute.Instance("instance1",
            networks=[openstack.compute.InstanceNetworkArgs(
                uuid=network1.id,
            )],
            security_groups=[
                "default",
                secgroup1.name,
            ])
        instance2 = openstack.compute.Instance("instance2",
            networks=[openstack.compute.InstanceNetworkArgs(
                uuid=network1.id,
            )],
            security_groups=[
                "default",
                secgroup1.name,
            ])
        monitor1 = openstack.loadbalancer.MonitorV1("monitor1",
            admin_state_up="true",
            delay=30,
            max_retries=3,
            timeout=5,
            type="TCP")
        pool1 = openstack.loadbalancer.PoolV1("pool1",
            lb_method="ROUND_ROBIN",
            monitor_ids=[monitor1.id],
            protocol="TCP",
            subnet_id=subnet1.id)
        member1 = openstack.loadbalancer.MemberV1("member1",
            address=instance1.access_ip_v4,
            pool_id=pool1.id,
            port=80)
        member2 = openstack.loadbalancer.MemberV1("member2",
            address=instance2.access_ip_v4,
            pool_id=pool1.id,
            port=80)
        vip1 = openstack.loadbalancer.Vip("vip1",
            pool_id=pool1.id,
            port=80,
            protocol="TCP",
            subnet_id=subnet1.id)
        ```

        ## Notes

        The `member` block is deprecated in favor of the `loadbalancer.MemberV1` resource.

        ## Import

        Load Balancer Pools can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:loadbalancer/poolV1:PoolV1 pool_1 b255e6ba-02ad-43e6-8951-3428ca26b713
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lb_method: The algorithm used to distribute load between the
               members of the pool. The current specification supports 'ROUND_ROBIN' and
               'LEAST_CONNECTIONS' as valid values for this attribute.
        :param pulumi.Input[str] lb_provider: The backend load balancing provider. For example:
               `haproxy`, `F5`, etc.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitor_ids: A list of IDs of monitors to associate with the
               pool.
        :param pulumi.Input[str] name: The name of the pool. Changing this updates the name of
               the existing pool.
        :param pulumi.Input[str] protocol: The protocol used by the pool members, you can use
               either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB pool. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB pool.
        :param pulumi.Input[str] subnet_id: The network on which the members of the pool will be
               located. Only members that are on this network can be added to the pool.
               Changing this creates a new pool.
        :param pulumi.Input[str] tenant_id: The owner of the member. Required if admin wants to
               create a pool member for another tenant. Changing this creates a new member.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PoolV1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V1 load balancer pool resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        pool1 = openstack.loadbalancer.PoolV1("pool1",
            lb_method="ROUND_ROBIN",
            lb_provider="haproxy",
            monitor_ids=["67890"],
            protocol="HTTP",
            subnet_id="12345")
        ```
        ## Complete Load Balancing Stack Example

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up=True)
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        secgroup1 = openstack.compute.SecGroup("secgroup1",
            description="Rules for secgroup_1",
            rules=[
                openstack.compute.SecGroupRuleArgs(
                    cidr="0.0.0.0/0",
                    from_port=-1,
                    ip_protocol="icmp",
                    to_port=-1,
                ),
                openstack.compute.SecGroupRuleArgs(
                    cidr="0.0.0.0/0",
                    from_port=80,
                    ip_protocol="tcp",
                    to_port=80,
                ),
            ])
        instance1 = openstack.compute.Instance("instance1",
            networks=[openstack.compute.InstanceNetworkArgs(
                uuid=network1.id,
            )],
            security_groups=[
                "default",
                secgroup1.name,
            ])
        instance2 = openstack.compute.Instance("instance2",
            networks=[openstack.compute.InstanceNetworkArgs(
                uuid=network1.id,
            )],
            security_groups=[
                "default",
                secgroup1.name,
            ])
        monitor1 = openstack.loadbalancer.MonitorV1("monitor1",
            admin_state_up="true",
            delay=30,
            max_retries=3,
            timeout=5,
            type="TCP")
        pool1 = openstack.loadbalancer.PoolV1("pool1",
            lb_method="ROUND_ROBIN",
            monitor_ids=[monitor1.id],
            protocol="TCP",
            subnet_id=subnet1.id)
        member1 = openstack.loadbalancer.MemberV1("member1",
            address=instance1.access_ip_v4,
            pool_id=pool1.id,
            port=80)
        member2 = openstack.loadbalancer.MemberV1("member2",
            address=instance2.access_ip_v4,
            pool_id=pool1.id,
            port=80)
        vip1 = openstack.loadbalancer.Vip("vip1",
            pool_id=pool1.id,
            port=80,
            protocol="TCP",
            subnet_id=subnet1.id)
        ```

        ## Notes

        The `member` block is deprecated in favor of the `loadbalancer.MemberV1` resource.

        ## Import

        Load Balancer Pools can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:loadbalancer/poolV1:PoolV1 pool_1 b255e6ba-02ad-43e6-8951-3428ca26b713
        ```

        :param str resource_name: The name of the resource.
        :param PoolV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PoolV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 lb_method: Optional[pulumi.Input[str]] = None,
                 lb_provider: Optional[pulumi.Input[str]] = None,
                 monitor_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
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

            if lb_method is None and not opts.urn:
                raise TypeError("Missing required property 'lb_method'")
            __props__['lb_method'] = lb_method
            __props__['lb_provider'] = lb_provider
            __props__['monitor_ids'] = monitor_ids
            __props__['name'] = name
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            __props__['region'] = region
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            __props__['tenant_id'] = tenant_id
        super(PoolV1, __self__).__init__(
            'openstack:loadbalancer/poolV1:PoolV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            lb_method: Optional[pulumi.Input[str]] = None,
            lb_provider: Optional[pulumi.Input[str]] = None,
            monitor_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'PoolV1':
        """
        Get an existing PoolV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lb_method: The algorithm used to distribute load between the
               members of the pool. The current specification supports 'ROUND_ROBIN' and
               'LEAST_CONNECTIONS' as valid values for this attribute.
        :param pulumi.Input[str] lb_provider: The backend load balancing provider. For example:
               `haproxy`, `F5`, etc.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitor_ids: A list of IDs of monitors to associate with the
               pool.
        :param pulumi.Input[str] name: The name of the pool. Changing this updates the name of
               the existing pool.
        :param pulumi.Input[str] protocol: The protocol used by the pool members, you can use
               either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB pool. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB pool.
        :param pulumi.Input[str] subnet_id: The network on which the members of the pool will be
               located. Only members that are on this network can be added to the pool.
               Changing this creates a new pool.
        :param pulumi.Input[str] tenant_id: The owner of the member. Required if admin wants to
               create a pool member for another tenant. Changing this creates a new member.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["lb_method"] = lb_method
        __props__["lb_provider"] = lb_provider
        __props__["monitor_ids"] = monitor_ids
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["region"] = region
        __props__["subnet_id"] = subnet_id
        __props__["tenant_id"] = tenant_id
        return PoolV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> pulumi.Output[str]:
        """
        The algorithm used to distribute load between the
        members of the pool. The current specification supports 'ROUND_ROBIN' and
        'LEAST_CONNECTIONS' as valid values for this attribute.
        """
        return pulumi.get(self, "lb_method")

    @property
    @pulumi.getter(name="lbProvider")
    def lb_provider(self) -> pulumi.Output[str]:
        """
        The backend load balancing provider. For example:
        `haproxy`, `F5`, etc.
        """
        return pulumi.get(self, "lb_provider")

    @property
    @pulumi.getter(name="monitorIds")
    def monitor_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of IDs of monitors to associate with the
        pool.
        """
        return pulumi.get(self, "monitor_ids")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the pool. Changing this updates the name of
        the existing pool.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol used by the pool members, you can use
        either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB pool. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB pool.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The network on which the members of the pool will be
        located. Only members that are on this network can be added to the pool.
        Changing this creates a new pool.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The owner of the member. Required if admin wants to
        create a pool member for another tenant. Changing this creates a new member.
        """
        return pulumi.get(self, "tenant_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

