# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RbacPolicyV2Args', 'RbacPolicyV2']

@pulumi.input_type
class RbacPolicyV2Args:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 object_id: pulumi.Input[str],
                 object_type: pulumi.Input[str],
                 target_tenant: pulumi.Input[str],
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RbacPolicyV2 resource.
        :param pulumi.Input[str] action: Action for the RBAC policy. Can either be
               `access_as_external` or `access_as_shared`.
        :param pulumi.Input[str] object_id: The ID of the `object_type` resource. An
               `object_type` of `network` returns a network ID and an `object_type` of
               `qos_policy` returns a QoS ID.
        :param pulumi.Input[str] object_type: The type of the object that the RBAC policy
               affects. Can be one of the following: `address_scope`, `address_group`,
               `network`, `qos_policy`, `security_group` or `subnetpool`.
        :param pulumi.Input[str] target_tenant: The ID of the tenant to which the RBAC policy
               will be enforced.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to configure a routing entry on a subnet. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               routing entry.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "object_id", object_id)
        pulumi.set(__self__, "object_type", object_type)
        pulumi.set(__self__, "target_tenant", target_tenant)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Action for the RBAC policy. Can either be
        `access_as_external` or `access_as_shared`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Input[str]:
        """
        The ID of the `object_type` resource. An
        `object_type` of `network` returns a network ID and an `object_type` of
        `qos_policy` returns a QoS ID.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> pulumi.Input[str]:
        """
        The type of the object that the RBAC policy
        affects. Can be one of the following: `address_scope`, `address_group`,
        `network`, `qos_policy`, `security_group` or `subnetpool`.
        """
        return pulumi.get(self, "object_type")

    @object_type.setter
    def object_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_type", value)

    @property
    @pulumi.getter(name="targetTenant")
    def target_tenant(self) -> pulumi.Input[str]:
        """
        The ID of the tenant to which the RBAC policy
        will be enforced.
        """
        return pulumi.get(self, "target_tenant")

    @target_tenant.setter
    def target_tenant(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_tenant", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to configure a routing entry on a subnet. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        routing entry.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _RbacPolicyV2State:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 object_type: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target_tenant: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RbacPolicyV2 resources.
        :param pulumi.Input[str] action: Action for the RBAC policy. Can either be
               `access_as_external` or `access_as_shared`.
        :param pulumi.Input[str] object_id: The ID of the `object_type` resource. An
               `object_type` of `network` returns a network ID and an `object_type` of
               `qos_policy` returns a QoS ID.
        :param pulumi.Input[str] object_type: The type of the object that the RBAC policy
               affects. Can be one of the following: `address_scope`, `address_group`,
               `network`, `qos_policy`, `security_group` or `subnetpool`.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to configure a routing entry on a subnet. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               routing entry.
        :param pulumi.Input[str] target_tenant: The ID of the tenant to which the RBAC policy
               will be enforced.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if object_type is not None:
            pulumi.set(__self__, "object_type", object_type)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if target_tenant is not None:
            pulumi.set(__self__, "target_tenant", target_tenant)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Action for the RBAC policy. Can either be
        `access_as_external` or `access_as_shared`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the `object_type` resource. An
        `object_type` of `network` returns a network ID and an `object_type` of
        `qos_policy` returns a QoS ID.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the object that the RBAC policy
        affects. Can be one of the following: `address_scope`, `address_group`,
        `network`, `qos_policy`, `security_group` or `subnetpool`.
        """
        return pulumi.get(self, "object_type")

    @object_type.setter
    def object_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_type", value)

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
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to configure a routing entry on a subnet. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        routing entry.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="targetTenant")
    def target_tenant(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the tenant to which the RBAC policy
        will be enforced.
        """
        return pulumi.get(self, "target_tenant")

    @target_tenant.setter
    def target_tenant(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_tenant", value)


class RbacPolicyV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 object_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target_tenant: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The RBAC policy resource contains functionality for working with Neutron RBAC
        Policies. Role-Based Access Control (RBAC) policy framework enables both
        operators and users to grant access to resources for specific projects.

        Sharing an object with a specific project is accomplished by creating a
        policy entry that permits the target project the `access_as_shared` action
        on that object.

        To make a network available as an external network for specific projects
        rather than all projects, use the `access_as_external` action.
        If a network is marked as external during creation, it now implicitly creates
        a wildcard RBAC policy granting everyone access to preserve previous behavior
        before this feature was added.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network_1",
            name="network_1",
            admin_state_up=True)
        rbac_policy1 = openstack.networking.RbacPolicyV2("rbac_policy_1",
            action="access_as_shared",
            object_id=network1.id,
            object_type="network",
            target_tenant="20415a973c9e45d3917f078950644697")
        ```

        ## Import

        RBAC policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:networking/rbacPolicyV2:RbacPolicyV2 rbac_policy_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action for the RBAC policy. Can either be
               `access_as_external` or `access_as_shared`.
        :param pulumi.Input[str] object_id: The ID of the `object_type` resource. An
               `object_type` of `network` returns a network ID and an `object_type` of
               `qos_policy` returns a QoS ID.
        :param pulumi.Input[str] object_type: The type of the object that the RBAC policy
               affects. Can be one of the following: `address_scope`, `address_group`,
               `network`, `qos_policy`, `security_group` or `subnetpool`.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to configure a routing entry on a subnet. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               routing entry.
        :param pulumi.Input[str] target_tenant: The ID of the tenant to which the RBAC policy
               will be enforced.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RbacPolicyV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The RBAC policy resource contains functionality for working with Neutron RBAC
        Policies. Role-Based Access Control (RBAC) policy framework enables both
        operators and users to grant access to resources for specific projects.

        Sharing an object with a specific project is accomplished by creating a
        policy entry that permits the target project the `access_as_shared` action
        on that object.

        To make a network available as an external network for specific projects
        rather than all projects, use the `access_as_external` action.
        If a network is marked as external during creation, it now implicitly creates
        a wildcard RBAC policy granting everyone access to preserve previous behavior
        before this feature was added.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network_1",
            name="network_1",
            admin_state_up=True)
        rbac_policy1 = openstack.networking.RbacPolicyV2("rbac_policy_1",
            action="access_as_shared",
            object_id=network1.id,
            object_type="network",
            target_tenant="20415a973c9e45d3917f078950644697")
        ```

        ## Import

        RBAC policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:networking/rbacPolicyV2:RbacPolicyV2 rbac_policy_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
        ```

        :param str resource_name: The name of the resource.
        :param RbacPolicyV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RbacPolicyV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 object_type: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target_tenant: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RbacPolicyV2Args.__new__(RbacPolicyV2Args)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            if object_id is None and not opts.urn:
                raise TypeError("Missing required property 'object_id'")
            __props__.__dict__["object_id"] = object_id
            if object_type is None and not opts.urn:
                raise TypeError("Missing required property 'object_type'")
            __props__.__dict__["object_type"] = object_type
            __props__.__dict__["region"] = region
            if target_tenant is None and not opts.urn:
                raise TypeError("Missing required property 'target_tenant'")
            __props__.__dict__["target_tenant"] = target_tenant
            __props__.__dict__["project_id"] = None
        super(RbacPolicyV2, __self__).__init__(
            'openstack:networking/rbacPolicyV2:RbacPolicyV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            object_type: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            target_tenant: Optional[pulumi.Input[str]] = None) -> 'RbacPolicyV2':
        """
        Get an existing RbacPolicyV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Action for the RBAC policy. Can either be
               `access_as_external` or `access_as_shared`.
        :param pulumi.Input[str] object_id: The ID of the `object_type` resource. An
               `object_type` of `network` returns a network ID and an `object_type` of
               `qos_policy` returns a QoS ID.
        :param pulumi.Input[str] object_type: The type of the object that the RBAC policy
               affects. Can be one of the following: `address_scope`, `address_group`,
               `network`, `qos_policy`, `security_group` or `subnetpool`.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to configure a routing entry on a subnet. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               routing entry.
        :param pulumi.Input[str] target_tenant: The ID of the tenant to which the RBAC policy
               will be enforced.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RbacPolicyV2State.__new__(_RbacPolicyV2State)

        __props__.__dict__["action"] = action
        __props__.__dict__["object_id"] = object_id
        __props__.__dict__["object_type"] = object_type
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["target_tenant"] = target_tenant
        return RbacPolicyV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Action for the RBAC policy. Can either be
        `access_as_external` or `access_as_shared`.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[str]:
        """
        The ID of the `object_type` resource. An
        `object_type` of `network` returns a network ID and an `object_type` of
        `qos_policy` returns a QoS ID.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter(name="objectType")
    def object_type(self) -> pulumi.Output[str]:
        """
        The type of the object that the RBAC policy
        affects. Can be one of the following: `address_scope`, `address_group`,
        `network`, `qos_policy`, `security_group` or `subnetpool`.
        """
        return pulumi.get(self, "object_type")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to configure a routing entry on a subnet. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        routing entry.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="targetTenant")
    def target_tenant(self) -> pulumi.Output[str]:
        """
        The ID of the tenant to which the RBAC policy
        will be enforced.
        """
        return pulumi.get(self, "target_tenant")

