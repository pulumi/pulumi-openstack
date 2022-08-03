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
        :param pulumi.Input[str] project_id: ID of the project to manage quotas. Changing this
               creates a new quota.
        :param pulumi.Input[int] health_monitor: Quota value for health_monitors. Changing
               this updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] l7_policy: Quota value for l7_policies. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] l7_rule: Quota value for l7_rules. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] listener: Quota value for listeners. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] loadbalancer: Quota value for loadbalancers. Changing this
               updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] member: Quota value for members. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] pool: Quota value for pools. Changing this updates the
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[str] region: Region in which to manage quotas. Changing this
               creates a new quota. If ommited, the region of the credentials is used.
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
        """
        ID of the project to manage quotas. Changing this
        creates a new quota.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="healthMonitor")
    def health_monitor(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for health_monitors. Changing
        this updates the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "health_monitor")

    @health_monitor.setter
    def health_monitor(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_monitor", value)

    @property
    @pulumi.getter(name="l7Policy")
    def l7_policy(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for l7_policies. Changing this
        updates the existing quota. Omitting it sets it to 0. Available in
        **Octavia minor version 2.19**.
        """
        return pulumi.get(self, "l7_policy")

    @l7_policy.setter
    def l7_policy(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_policy", value)

    @property
    @pulumi.getter(name="l7Rule")
    def l7_rule(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for l7_rules. Changing this
        updates the existing quota. Omitting it sets it to 0. Available in
        **Octavia minor version 2.19**.
        """
        return pulumi.get(self, "l7_rule")

    @l7_rule.setter
    def l7_rule(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_rule", value)

    @property
    @pulumi.getter
    def listener(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for listeners. Changing this updates
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "listener")

    @listener.setter
    def listener(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listener", value)

    @property
    @pulumi.getter
    def loadbalancer(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for loadbalancers. Changing this
        updates the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "loadbalancer")

    @loadbalancer.setter
    def loadbalancer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "loadbalancer", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for members. Changing this updates
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def pool(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for pools. Changing this updates the
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "pool")

    @pool.setter
    def pool(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pool", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region in which to manage quotas. Changing this
        creates a new quota. If ommited, the region of the credentials is used.
        """
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
        :param pulumi.Input[int] health_monitor: Quota value for health_monitors. Changing
               this updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] l7_policy: Quota value for l7_policies. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] l7_rule: Quota value for l7_rules. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] listener: Quota value for listeners. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] loadbalancer: Quota value for loadbalancers. Changing this
               updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] member: Quota value for members. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] pool: Quota value for pools. Changing this updates the
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[str] project_id: ID of the project to manage quotas. Changing this
               creates a new quota.
        :param pulumi.Input[str] region: Region in which to manage quotas. Changing this
               creates a new quota. If ommited, the region of the credentials is used.
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
        """
        Quota value for health_monitors. Changing
        this updates the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "health_monitor")

    @health_monitor.setter
    def health_monitor(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_monitor", value)

    @property
    @pulumi.getter(name="l7Policy")
    def l7_policy(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for l7_policies. Changing this
        updates the existing quota. Omitting it sets it to 0. Available in
        **Octavia minor version 2.19**.
        """
        return pulumi.get(self, "l7_policy")

    @l7_policy.setter
    def l7_policy(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_policy", value)

    @property
    @pulumi.getter(name="l7Rule")
    def l7_rule(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for l7_rules. Changing this
        updates the existing quota. Omitting it sets it to 0. Available in
        **Octavia minor version 2.19**.
        """
        return pulumi.get(self, "l7_rule")

    @l7_rule.setter
    def l7_rule(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "l7_rule", value)

    @property
    @pulumi.getter
    def listener(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for listeners. Changing this updates
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "listener")

    @listener.setter
    def listener(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listener", value)

    @property
    @pulumi.getter
    def loadbalancer(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for loadbalancers. Changing this
        updates the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "loadbalancer")

    @loadbalancer.setter
    def loadbalancer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "loadbalancer", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for members. Changing this updates
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def pool(self) -> Optional[pulumi.Input[int]]:
        """
        Quota value for pools. Changing this updates the
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "pool")

    @pool.setter
    def pool(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pool", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the project to manage quotas. Changing this
        creates a new quota.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region in which to manage quotas. Changing this
        creates a new quota. If ommited, the region of the credentials is used.
        """
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
        Manages a V2 load balancer quota resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        > **Note:** This resource is only available for Octavia.

        > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack
           API in case of delete call.

        > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
           created with zero value.

        > **Note:** This resource has attributes that depend on octavia minor versions.
        Please ensure your Openstack cloud supports the required minor version.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        quota1 = openstack.loadbalancer.Quota("quota1",
            health_monitor=10,
            l7_policy=11,
            l7_rule=12,
            listener=7,
            loadbalancer=6,
            member=8,
            pool=9,
            project_id=project1.id)
        ```

        ## Import

        Quotas can be imported using the `project_id/region_name`, where region_name is the one defined is the Openstack credentials that are in use. E.g.

        ```sh
         $ pulumi import openstack:loadbalancer/quota:Quota quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] health_monitor: Quota value for health_monitors. Changing
               this updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] l7_policy: Quota value for l7_policies. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] l7_rule: Quota value for l7_rules. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] listener: Quota value for listeners. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] loadbalancer: Quota value for loadbalancers. Changing this
               updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] member: Quota value for members. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] pool: Quota value for pools. Changing this updates the
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[str] project_id: ID of the project to manage quotas. Changing this
               creates a new quota.
        :param pulumi.Input[str] region: Region in which to manage quotas. Changing this
               creates a new quota. If ommited, the region of the credentials is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QuotaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 load balancer quota resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        > **Note:** This resource is only available for Octavia.

        > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack
           API in case of delete call.

        > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
           created with zero value.

        > **Note:** This resource has attributes that depend on octavia minor versions.
        Please ensure your Openstack cloud supports the required minor version.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        project1 = openstack.identity.Project("project1")
        quota1 = openstack.loadbalancer.Quota("quota1",
            health_monitor=10,
            l7_policy=11,
            l7_rule=12,
            listener=7,
            loadbalancer=6,
            member=8,
            pool=9,
            project_id=project1.id)
        ```

        ## Import

        Quotas can be imported using the `project_id/region_name`, where region_name is the one defined is the Openstack credentials that are in use. E.g.

        ```sh
         $ pulumi import openstack:loadbalancer/quota:Quota quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
        ```

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
        :param pulumi.Input[int] health_monitor: Quota value for health_monitors. Changing
               this updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] l7_policy: Quota value for l7_policies. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] l7_rule: Quota value for l7_rules. Changing this
               updates the existing quota. Omitting it sets it to 0. Available in
               **Octavia minor version 2.19**.
        :param pulumi.Input[int] listener: Quota value for listeners. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] loadbalancer: Quota value for loadbalancers. Changing this
               updates the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] member: Quota value for members. Changing this updates
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[int] pool: Quota value for pools. Changing this updates the
               the existing quota. Omitting it sets it to 0.
        :param pulumi.Input[str] project_id: ID of the project to manage quotas. Changing this
               creates a new quota.
        :param pulumi.Input[str] region: Region in which to manage quotas. Changing this
               creates a new quota. If ommited, the region of the credentials is used.
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
        """
        Quota value for health_monitors. Changing
        this updates the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "health_monitor")

    @property
    @pulumi.getter(name="l7Policy")
    def l7_policy(self) -> pulumi.Output[int]:
        """
        Quota value for l7_policies. Changing this
        updates the existing quota. Omitting it sets it to 0. Available in
        **Octavia minor version 2.19**.
        """
        return pulumi.get(self, "l7_policy")

    @property
    @pulumi.getter(name="l7Rule")
    def l7_rule(self) -> pulumi.Output[int]:
        """
        Quota value for l7_rules. Changing this
        updates the existing quota. Omitting it sets it to 0. Available in
        **Octavia minor version 2.19**.
        """
        return pulumi.get(self, "l7_rule")

    @property
    @pulumi.getter
    def listener(self) -> pulumi.Output[int]:
        """
        Quota value for listeners. Changing this updates
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "listener")

    @property
    @pulumi.getter
    def loadbalancer(self) -> pulumi.Output[int]:
        """
        Quota value for loadbalancers. Changing this
        updates the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "loadbalancer")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[int]:
        """
        Quota value for members. Changing this updates
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def pool(self) -> pulumi.Output[int]:
        """
        Quota value for pools. Changing this updates the
        the existing quota. Omitting it sets it to 0.
        """
        return pulumi.get(self, "pool")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        ID of the project to manage quotas. Changing this
        creates a new quota.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region in which to manage quotas. Changing this
        creates a new quota. If ommited, the region of the credentials is used.
        """
        return pulumi.get(self, "region")

