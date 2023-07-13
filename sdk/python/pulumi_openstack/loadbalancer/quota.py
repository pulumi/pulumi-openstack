# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['QuotaArgs', 'Quota']

@pulumi.input_type
class QuotaArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 health_monitor: Optional[pulumi.Input[int]] = None,
                 l7_policy: Optional[pulumi.Input[int]] = None,
                 l7_rule: Optional[pulumi.Input[int]] = None,
                 listener: Optional[pulumi.Input[int]] = None,
                 loadbalancer: Optional[pulumi.Input[int]] = None,
                 member: Optional[pulumi.Input[int]] = None,
                 pool: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Quota resource.
        """
        pulumi.set(__self__, "project_id", project_id)
        if health_monitor is not None:
            pulumi.set(__self__, "health_monitor", health_monitor)
        if l7_policy is not None:
            pulumi.set(__self__, "l7_policy", l7_policy)
        if l7_rule is not None:
            pulumi.set(__self__, "l7_rule", l7_rule)
        if listener is not None:
            pulumi.set(__self__, "listener", listener)
        if loadbalancer is not None:
            pulumi.set(__self__, "loadbalancer", loadbalancer)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if pool is not None:
            pulumi.set(__self__, "pool", pool)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="healthMonitor")
    def health_monitor(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_monitor")

    @health_monitor.setter
    def health_monitor(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_monitor", value)

    @property
    @pulumi.getter(name="l7Policy")
    def l7_policy(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "l7_policy")

    @l7_policy.setter
    def l7_policy(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_policy", value)

    @property
    @pulumi.getter(name="l7Rule")
    def l7_rule(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "l7_rule")

    @l7_rule.setter
    def l7_rule(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_rule", value)

    @property
    @pulumi.getter
    def listener(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "listener")

    @listener.setter
    def listener(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listener", value)

    @property
    @pulumi.getter
    def loadbalancer(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "loadbalancer")

    @loadbalancer.setter
    def loadbalancer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "loadbalancer", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def pool(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "pool")

    @pool.setter
    def pool(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pool", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _QuotaState:
    def __init__(__self__, *,
                 health_monitor: Optional[pulumi.Input[int]] = None,
                 l7_policy: Optional[pulumi.Input[int]] = None,
                 l7_rule: Optional[pulumi.Input[int]] = None,
                 listener: Optional[pulumi.Input[int]] = None,
                 loadbalancer: Optional[pulumi.Input[int]] = None,
                 member: Optional[pulumi.Input[int]] = None,
                 pool: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Quota resources.
        """
        if health_monitor is not None:
            pulumi.set(__self__, "health_monitor", health_monitor)
        if l7_policy is not None:
            pulumi.set(__self__, "l7_policy", l7_policy)
        if l7_rule is not None:
            pulumi.set(__self__, "l7_rule", l7_rule)
        if listener is not None:
            pulumi.set(__self__, "listener", listener)
        if loadbalancer is not None:
            pulumi.set(__self__, "loadbalancer", loadbalancer)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if pool is not None:
            pulumi.set(__self__, "pool", pool)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="healthMonitor")
    def health_monitor(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_monitor")

    @health_monitor.setter
    def health_monitor(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_monitor", value)

    @property
    @pulumi.getter(name="l7Policy")
    def l7_policy(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "l7_policy")

    @l7_policy.setter
    def l7_policy(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_policy", value)

    @property
    @pulumi.getter(name="l7Rule")
    def l7_rule(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "l7_rule")

    @l7_rule.setter
    def l7_rule(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_rule", value)

    @property
    @pulumi.getter
    def listener(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "listener")

    @listener.setter
    def listener(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listener", value)

    @property
    @pulumi.getter
    def loadbalancer(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "loadbalancer")

    @loadbalancer.setter
    def loadbalancer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "loadbalancer", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def pool(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "pool")

    @pool.setter
    def pool(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pool", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class Quota(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 health_monitor: Optional[pulumi.Input[int]] = None,
                 l7_policy: Optional[pulumi.Input[int]] = None,
                 l7_rule: Optional[pulumi.Input[int]] = None,
                 listener: Optional[pulumi.Input[int]] = None,
                 loadbalancer: Optional[pulumi.Input[int]] = None,
                 member: Optional[pulumi.Input[int]] = None,
                 pool: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Quota resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QuotaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Quota resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param QuotaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QuotaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 health_monitor: Optional[pulumi.Input[int]] = None,
                 l7_policy: Optional[pulumi.Input[int]] = None,
                 l7_rule: Optional[pulumi.Input[int]] = None,
                 listener: Optional[pulumi.Input[int]] = None,
                 loadbalancer: Optional[pulumi.Input[int]] = None,
                 member: Optional[pulumi.Input[int]] = None,
                 pool: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QuotaArgs.__new__(QuotaArgs)

            __props__.__dict__["health_monitor"] = health_monitor
            __props__.__dict__["l7_policy"] = l7_policy
            __props__.__dict__["l7_rule"] = l7_rule
            __props__.__dict__["listener"] = listener
            __props__.__dict__["loadbalancer"] = loadbalancer
            __props__.__dict__["member"] = member
            __props__.__dict__["pool"] = pool
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
        super(Quota, __self__).__init__(
            'openstack:loadbalancer/quota:Quota',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            health_monitor: Optional[pulumi.Input[int]] = None,
            l7_policy: Optional[pulumi.Input[int]] = None,
            l7_rule: Optional[pulumi.Input[int]] = None,
            listener: Optional[pulumi.Input[int]] = None,
            loadbalancer: Optional[pulumi.Input[int]] = None,
            member: Optional[pulumi.Input[int]] = None,
            pool: Optional[pulumi.Input[int]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'Quota':
        """
        Get an existing Quota resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QuotaState.__new__(_QuotaState)

        __props__.__dict__["health_monitor"] = health_monitor
        __props__.__dict__["l7_policy"] = l7_policy
        __props__.__dict__["l7_rule"] = l7_rule
        __props__.__dict__["listener"] = listener
        __props__.__dict__["loadbalancer"] = loadbalancer
        __props__.__dict__["member"] = member
        __props__.__dict__["pool"] = pool
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return Quota(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="healthMonitor")
    def health_monitor(self) -> pulumi.Output[int]:
        return pulumi.get(self, "health_monitor")

    @property
    @pulumi.getter(name="l7Policy")
    def l7_policy(self) -> pulumi.Output[int]:
        return pulumi.get(self, "l7_policy")

    @property
    @pulumi.getter(name="l7Rule")
    def l7_rule(self) -> pulumi.Output[int]:
        return pulumi.get(self, "l7_rule")

    @property
    @pulumi.getter
    def listener(self) -> pulumi.Output[int]:
        return pulumi.get(self, "listener")

    @property
    @pulumi.getter
    def loadbalancer(self) -> pulumi.Output[int]:
        return pulumi.get(self, "loadbalancer")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[int]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def pool(self) -> pulumi.Output[int]:
        return pulumi.get(self, "pool")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

