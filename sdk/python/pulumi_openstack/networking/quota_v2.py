# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['QuotaV2Args', 'QuotaV2']

@pulumi.input_type
class QuotaV2Args:
    def __init__(__self__, *,
                 project_id: pulumi.Input[builtins.str],
                 floatingip: Optional[pulumi.Input[builtins.int]] = None,
                 network: Optional[pulumi.Input[builtins.int]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 rbac_policy: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router: Optional[pulumi.Input[builtins.int]] = None,
                 security_group: Optional[pulumi.Input[builtins.int]] = None,
                 security_group_rule: Optional[pulumi.Input[builtins.int]] = None,
                 subnet: Optional[pulumi.Input[builtins.int]] = None,
                 subnetpool: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a QuotaV2 resource.
        :param pulumi.Input[builtins.str] project_id: ID of the project to manage quota. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] floatingip: Quota value for floating IPs. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] network: Quota value for networks. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] port: Quota value for ports. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] rbac_policy: Quota value for RBAC policies.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.str] region: The region in which to create the quota. If
               omitted, the `region` argument of the provider is used. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] router: Quota value for routers. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] security_group: Quota value for security groups. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] security_group_rule: Quota value for security group rules.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.int] subnet: Quota value for subnets. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] subnetpool: Quota value for subnetpools.
               Changing this updates the existing quota.
        """
        pulumi.set(__self__, "project_id", project_id)
        if floatingip is not None:
            pulumi.set(__self__, "floatingip", floatingip)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if rbac_policy is not None:
            pulumi.set(__self__, "rbac_policy", rbac_policy)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if router is not None:
            pulumi.set(__self__, "router", router)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if security_group_rule is not None:
            pulumi.set(__self__, "security_group_rule", security_group_rule)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if subnetpool is not None:
            pulumi.set(__self__, "subnetpool", subnetpool)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        """
        ID of the project to manage quota. Changing this
        creates new quota.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def floatingip(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for floating IPs. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "floatingip")

    @floatingip.setter
    def floatingip(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "floatingip", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for networks. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for ports. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="rbacPolicy")
    def rbac_policy(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for RBAC policies.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "rbac_policy")

    @rbac_policy.setter
    def rbac_policy(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "rbac_policy", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to create the quota. If
        omitted, the `region` argument of the provider is used. Changing this
        creates new quota.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def router(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for routers. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "router")

    @router.setter
    def router(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "router", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for security groups. Changing
        this updates the existing quota.
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="securityGroupRule")
    def security_group_rule(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for security group rules.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "security_group_rule")

    @security_group_rule.setter
    def security_group_rule(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "security_group_rule", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for subnets. Changing
        this updates the existing quota.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def subnetpool(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for subnetpools.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "subnetpool")

    @subnetpool.setter
    def subnetpool(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "subnetpool", value)


@pulumi.input_type
class _QuotaV2State:
    def __init__(__self__, *,
                 floatingip: Optional[pulumi.Input[builtins.int]] = None,
                 network: Optional[pulumi.Input[builtins.int]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 rbac_policy: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router: Optional[pulumi.Input[builtins.int]] = None,
                 security_group: Optional[pulumi.Input[builtins.int]] = None,
                 security_group_rule: Optional[pulumi.Input[builtins.int]] = None,
                 subnet: Optional[pulumi.Input[builtins.int]] = None,
                 subnetpool: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering QuotaV2 resources.
        :param pulumi.Input[builtins.int] floatingip: Quota value for floating IPs. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] network: Quota value for networks. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] port: Quota value for ports. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.str] project_id: ID of the project to manage quota. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] rbac_policy: Quota value for RBAC policies.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.str] region: The region in which to create the quota. If
               omitted, the `region` argument of the provider is used. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] router: Quota value for routers. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] security_group: Quota value for security groups. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] security_group_rule: Quota value for security group rules.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.int] subnet: Quota value for subnets. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] subnetpool: Quota value for subnetpools.
               Changing this updates the existing quota.
        """
        if floatingip is not None:
            pulumi.set(__self__, "floatingip", floatingip)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if rbac_policy is not None:
            pulumi.set(__self__, "rbac_policy", rbac_policy)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if router is not None:
            pulumi.set(__self__, "router", router)
        if security_group is not None:
            pulumi.set(__self__, "security_group", security_group)
        if security_group_rule is not None:
            pulumi.set(__self__, "security_group_rule", security_group_rule)
        if subnet is not None:
            pulumi.set(__self__, "subnet", subnet)
        if subnetpool is not None:
            pulumi.set(__self__, "subnetpool", subnetpool)

    @property
    @pulumi.getter
    def floatingip(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for floating IPs. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "floatingip")

    @floatingip.setter
    def floatingip(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "floatingip", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for networks. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for ports. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ID of the project to manage quota. Changing this
        creates new quota.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="rbacPolicy")
    def rbac_policy(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for RBAC policies.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "rbac_policy")

    @rbac_policy.setter
    def rbac_policy(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "rbac_policy", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to create the quota. If
        omitted, the `region` argument of the provider is used. Changing this
        creates new quota.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def router(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for routers. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "router")

    @router.setter
    def router(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "router", value)

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for security groups. Changing
        this updates the existing quota.
        """
        return pulumi.get(self, "security_group")

    @security_group.setter
    def security_group(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "security_group", value)

    @property
    @pulumi.getter(name="securityGroupRule")
    def security_group_rule(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for security group rules.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "security_group_rule")

    @security_group_rule.setter
    def security_group_rule(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "security_group_rule", value)

    @property
    @pulumi.getter
    def subnet(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for subnets. Changing
        this updates the existing quota.
        """
        return pulumi.get(self, "subnet")

    @subnet.setter
    def subnet(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "subnet", value)

    @property
    @pulumi.getter
    def subnetpool(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Quota value for subnetpools.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "subnetpool")

    @subnetpool.setter
    def subnetpool(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "subnetpool", value)


class QuotaV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 floatingip: Optional[pulumi.Input[builtins.int]] = None,
                 network: Optional[pulumi.Input[builtins.int]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 rbac_policy: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router: Optional[pulumi.Input[builtins.int]] = None,
                 security_group: Optional[pulumi.Input[builtins.int]] = None,
                 security_group_rule: Optional[pulumi.Input[builtins.int]] = None,
                 subnet: Optional[pulumi.Input[builtins.int]] = None,
                 subnetpool: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Manages a V2 networking quota resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
            in case of delete call.

        > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
            created with zero value.

        ## Import

        Quotas can be imported using the `project_id/region_name`, e.g.

        ```sh
        $ pulumi import openstack:networking/quotaV2:QuotaV2 quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] floatingip: Quota value for floating IPs. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] network: Quota value for networks. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] port: Quota value for ports. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.str] project_id: ID of the project to manage quota. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] rbac_policy: Quota value for RBAC policies.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.str] region: The region in which to create the quota. If
               omitted, the `region` argument of the provider is used. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] router: Quota value for routers. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] security_group: Quota value for security groups. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] security_group_rule: Quota value for security group rules.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.int] subnet: Quota value for subnets. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] subnetpool: Quota value for subnetpools.
               Changing this updates the existing quota.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QuotaV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 networking quota resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API
            in case of delete call.

        > **Note:** This resource has all-in creation so all optional quota arguments that were not specified are
            created with zero value.

        ## Import

        Quotas can be imported using the `project_id/region_name`, e.g.

        ```sh
        $ pulumi import openstack:networking/quotaV2:QuotaV2 quota_1 2a0f2240-c5e6-41de-896d-e80d97428d6b/region_1
        ```

        :param str resource_name: The name of the resource.
        :param QuotaV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QuotaV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 floatingip: Optional[pulumi.Input[builtins.int]] = None,
                 network: Optional[pulumi.Input[builtins.int]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 rbac_policy: Optional[pulumi.Input[builtins.int]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 router: Optional[pulumi.Input[builtins.int]] = None,
                 security_group: Optional[pulumi.Input[builtins.int]] = None,
                 security_group_rule: Optional[pulumi.Input[builtins.int]] = None,
                 subnet: Optional[pulumi.Input[builtins.int]] = None,
                 subnetpool: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QuotaV2Args.__new__(QuotaV2Args)

            __props__.__dict__["floatingip"] = floatingip
            __props__.__dict__["network"] = network
            __props__.__dict__["port"] = port
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["rbac_policy"] = rbac_policy
            __props__.__dict__["region"] = region
            __props__.__dict__["router"] = router
            __props__.__dict__["security_group"] = security_group
            __props__.__dict__["security_group_rule"] = security_group_rule
            __props__.__dict__["subnet"] = subnet
            __props__.__dict__["subnetpool"] = subnetpool
        super(QuotaV2, __self__).__init__(
            'openstack:networking/quotaV2:QuotaV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            floatingip: Optional[pulumi.Input[builtins.int]] = None,
            network: Optional[pulumi.Input[builtins.int]] = None,
            port: Optional[pulumi.Input[builtins.int]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            rbac_policy: Optional[pulumi.Input[builtins.int]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            router: Optional[pulumi.Input[builtins.int]] = None,
            security_group: Optional[pulumi.Input[builtins.int]] = None,
            security_group_rule: Optional[pulumi.Input[builtins.int]] = None,
            subnet: Optional[pulumi.Input[builtins.int]] = None,
            subnetpool: Optional[pulumi.Input[builtins.int]] = None) -> 'QuotaV2':
        """
        Get an existing QuotaV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] floatingip: Quota value for floating IPs. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] network: Quota value for networks. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] port: Quota value for ports. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.str] project_id: ID of the project to manage quota. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] rbac_policy: Quota value for RBAC policies.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.str] region: The region in which to create the quota. If
               omitted, the `region` argument of the provider is used. Changing this
               creates new quota.
        :param pulumi.Input[builtins.int] router: Quota value for routers. Changing this updates the
               existing quota.
        :param pulumi.Input[builtins.int] security_group: Quota value for security groups. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] security_group_rule: Quota value for security group rules.
               Changing this updates the existing quota.
        :param pulumi.Input[builtins.int] subnet: Quota value for subnets. Changing
               this updates the existing quota.
        :param pulumi.Input[builtins.int] subnetpool: Quota value for subnetpools.
               Changing this updates the existing quota.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QuotaV2State.__new__(_QuotaV2State)

        __props__.__dict__["floatingip"] = floatingip
        __props__.__dict__["network"] = network
        __props__.__dict__["port"] = port
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["rbac_policy"] = rbac_policy
        __props__.__dict__["region"] = region
        __props__.__dict__["router"] = router
        __props__.__dict__["security_group"] = security_group
        __props__.__dict__["security_group_rule"] = security_group_rule
        __props__.__dict__["subnet"] = subnet
        __props__.__dict__["subnetpool"] = subnetpool
        return QuotaV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def floatingip(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for floating IPs. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "floatingip")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for networks. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for ports. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        ID of the project to manage quota. Changing this
        creates new quota.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="rbacPolicy")
    def rbac_policy(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for RBAC policies.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "rbac_policy")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to create the quota. If
        omitted, the `region` argument of the provider is used. Changing this
        creates new quota.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def router(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for routers. Changing this updates the
        existing quota.
        """
        return pulumi.get(self, "router")

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for security groups. Changing
        this updates the existing quota.
        """
        return pulumi.get(self, "security_group")

    @property
    @pulumi.getter(name="securityGroupRule")
    def security_group_rule(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for security group rules.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "security_group_rule")

    @property
    @pulumi.getter
    def subnet(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for subnets. Changing
        this updates the existing quota.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def subnetpool(self) -> pulumi.Output[builtins.int]:
        """
        Quota value for subnetpools.
        Changing this updates the existing quota.
        """
        return pulumi.get(self, "subnetpool")

