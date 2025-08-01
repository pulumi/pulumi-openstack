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
from .. import _utilities

__all__ = ['PortSecGroupAssociateArgs', 'PortSecGroupAssociate']

@pulumi.input_type
class PortSecGroupAssociateArgs:
    def __init__(__self__, *,
                 port_id: pulumi.Input[_builtins.str],
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]],
                 enforce: Optional[pulumi.Input[_builtins.bool]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a PortSecGroupAssociate resource.
        :param pulumi.Input[_builtins.str] port_id: An UUID of the port to apply security groups to.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] security_group_ids: A list of security group IDs to apply to
               the port. The security groups must be specified by ID and not name (as
               opposed to how they are configured with the Compute Instance).
        :param pulumi.Input[_builtins.bool] enforce: Whether to replace or append the list of security
               groups, specified in the `security_group_ids`. Defaults to `false`.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to manage a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               resource.
        """
        pulumi.set(__self__, "port_id", port_id)
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        if enforce is not None:
            pulumi.set(__self__, "enforce", enforce)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @_builtins.property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Input[_builtins.str]:
        """
        An UUID of the port to apply security groups to.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "port_id", value)

    @_builtins.property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]:
        """
        A list of security group IDs to apply to
        the port. The security groups must be specified by ID and not name (as
        opposed to how they are configured with the Compute Instance).
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]):
        pulumi.set(self, "security_group_ids", value)

    @_builtins.property
    @pulumi.getter
    def enforce(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Whether to replace or append the list of security
        groups, specified in the `security_group_ids`. Defaults to `false`.
        """
        return pulumi.get(self, "enforce")

    @enforce.setter
    def enforce(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enforce", value)

    @_builtins.property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to manage a port. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _PortSecGroupAssociateState:
    def __init__(__self__, *,
                 all_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 enforce: Optional[pulumi.Input[_builtins.bool]] = None,
                 port_id: Optional[pulumi.Input[_builtins.str]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering PortSecGroupAssociate resources.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] all_security_group_ids: The collection of Security Group IDs on the port
               which have been explicitly and implicitly added.
        :param pulumi.Input[_builtins.bool] enforce: Whether to replace or append the list of security
               groups, specified in the `security_group_ids`. Defaults to `false`.
        :param pulumi.Input[_builtins.str] port_id: An UUID of the port to apply security groups to.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to manage a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] security_group_ids: A list of security group IDs to apply to
               the port. The security groups must be specified by ID and not name (as
               opposed to how they are configured with the Compute Instance).
        """
        if all_security_group_ids is not None:
            pulumi.set(__self__, "all_security_group_ids", all_security_group_ids)
        if enforce is not None:
            pulumi.set(__self__, "enforce", enforce)
        if port_id is not None:
            pulumi.set(__self__, "port_id", port_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)

    @_builtins.property
    @pulumi.getter(name="allSecurityGroupIds")
    def all_security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        The collection of Security Group IDs on the port
        which have been explicitly and implicitly added.
        """
        return pulumi.get(self, "all_security_group_ids")

    @all_security_group_ids.setter
    def all_security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "all_security_group_ids", value)

    @_builtins.property
    @pulumi.getter
    def enforce(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Whether to replace or append the list of security
        groups, specified in the `security_group_ids`. Defaults to `false`.
        """
        return pulumi.get(self, "enforce")

    @enforce.setter
    def enforce(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "enforce", value)

    @_builtins.property
    @pulumi.getter(name="portId")
    def port_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        An UUID of the port to apply security groups to.
        """
        return pulumi.get(self, "port_id")

    @port_id.setter
    def port_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "port_id", value)

    @_builtins.property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to manage a port. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "region", value)

    @_builtins.property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A list of security group IDs to apply to
        the port. The security groups must be specified by ID and not name (as
        opposed to how they are configured with the Compute Instance).
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "security_group_ids", value)


@pulumi.type_token("openstack:networking/portSecGroupAssociate:PortSecGroupAssociate")
class PortSecGroupAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enforce: Optional[pulumi.Input[_builtins.bool]] = None,
                 port_id: Optional[pulumi.Input[_builtins.str]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Append a security group to an existing port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        system_port = openstack.networking.get_port(fixed_ip="10.0.0.10")
        secgroup = openstack.networking.get_sec_group(name="secgroup")
        port1 = openstack.networking.PortSecGroupAssociate("port_1",
            port_id=system_port.id,
            security_group_ids=[secgroup.id])
        ```

        ### Enforce a security group to an existing port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        system_port = openstack.networking.get_port(fixed_ip="10.0.0.10")
        secgroup = openstack.networking.get_sec_group(name="secgroup")
        port1 = openstack.networking.PortSecGroupAssociate("port_1",
            port_id=system_port.id,
            enforce=True,
            security_group_ids=[secgroup.id])
        ```

        ### Remove all security groups from an existing port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        system_port = openstack.networking.get_port(fixed_ip="10.0.0.10")
        port1 = openstack.networking.PortSecGroupAssociate("port_1",
            port_id=system_port.id,
            enforce=True,
            security_group_ids=[])
        ```

        ## Import

        Port security group association can be imported using the `id` of the port, e.g.

        ```sh
        $ pulumi import openstack:networking/portSecGroupAssociate:PortSecGroupAssociate port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] enforce: Whether to replace or append the list of security
               groups, specified in the `security_group_ids`. Defaults to `false`.
        :param pulumi.Input[_builtins.str] port_id: An UUID of the port to apply security groups to.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to manage a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] security_group_ids: A list of security group IDs to apply to
               the port. The security groups must be specified by ID and not name (as
               opposed to how they are configured with the Compute Instance).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PortSecGroupAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Append a security group to an existing port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        system_port = openstack.networking.get_port(fixed_ip="10.0.0.10")
        secgroup = openstack.networking.get_sec_group(name="secgroup")
        port1 = openstack.networking.PortSecGroupAssociate("port_1",
            port_id=system_port.id,
            security_group_ids=[secgroup.id])
        ```

        ### Enforce a security group to an existing port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        system_port = openstack.networking.get_port(fixed_ip="10.0.0.10")
        secgroup = openstack.networking.get_sec_group(name="secgroup")
        port1 = openstack.networking.PortSecGroupAssociate("port_1",
            port_id=system_port.id,
            enforce=True,
            security_group_ids=[secgroup.id])
        ```

        ### Remove all security groups from an existing port

        ```python
        import pulumi
        import pulumi_openstack as openstack

        system_port = openstack.networking.get_port(fixed_ip="10.0.0.10")
        port1 = openstack.networking.PortSecGroupAssociate("port_1",
            port_id=system_port.id,
            enforce=True,
            security_group_ids=[])
        ```

        ## Import

        Port security group association can be imported using the `id` of the port, e.g.

        ```sh
        $ pulumi import openstack:networking/portSecGroupAssociate:PortSecGroupAssociate port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1
        ```

        :param str resource_name: The name of the resource.
        :param PortSecGroupAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PortSecGroupAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enforce: Optional[pulumi.Input[_builtins.bool]] = None,
                 port_id: Optional[pulumi.Input[_builtins.str]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PortSecGroupAssociateArgs.__new__(PortSecGroupAssociateArgs)

            __props__.__dict__["enforce"] = enforce
            if port_id is None and not opts.urn:
                raise TypeError("Missing required property 'port_id'")
            __props__.__dict__["port_id"] = port_id
            __props__.__dict__["region"] = region
            if security_group_ids is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_ids'")
            __props__.__dict__["security_group_ids"] = security_group_ids
            __props__.__dict__["all_security_group_ids"] = None
        super(PortSecGroupAssociate, __self__).__init__(
            'openstack:networking/portSecGroupAssociate:PortSecGroupAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            enforce: Optional[pulumi.Input[_builtins.bool]] = None,
            port_id: Optional[pulumi.Input[_builtins.str]] = None,
            region: Optional[pulumi.Input[_builtins.str]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None) -> 'PortSecGroupAssociate':
        """
        Get an existing PortSecGroupAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] all_security_group_ids: The collection of Security Group IDs on the port
               which have been explicitly and implicitly added.
        :param pulumi.Input[_builtins.bool] enforce: Whether to replace or append the list of security
               groups, specified in the `security_group_ids`. Defaults to `false`.
        :param pulumi.Input[_builtins.str] port_id: An UUID of the port to apply security groups to.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to manage a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] security_group_ids: A list of security group IDs to apply to
               the port. The security groups must be specified by ID and not name (as
               opposed to how they are configured with the Compute Instance).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PortSecGroupAssociateState.__new__(_PortSecGroupAssociateState)

        __props__.__dict__["all_security_group_ids"] = all_security_group_ids
        __props__.__dict__["enforce"] = enforce
        __props__.__dict__["port_id"] = port_id
        __props__.__dict__["region"] = region
        __props__.__dict__["security_group_ids"] = security_group_ids
        return PortSecGroupAssociate(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="allSecurityGroupIds")
    def all_security_group_ids(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        The collection of Security Group IDs on the port
        which have been explicitly and implicitly added.
        """
        return pulumi.get(self, "all_security_group_ids")

    @_builtins.property
    @pulumi.getter
    def enforce(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Whether to replace or append the list of security
        groups, specified in the `security_group_ids`. Defaults to `false`.
        """
        return pulumi.get(self, "enforce")

    @_builtins.property
    @pulumi.getter(name="portId")
    def port_id(self) -> pulumi.Output[_builtins.str]:
        """
        An UUID of the port to apply security groups to.
        """
        return pulumi.get(self, "port_id")

    @_builtins.property
    @pulumi.getter
    def region(self) -> pulumi.Output[_builtins.str]:
        """
        The region in which to obtain the V2 networking client.
        A networking client is needed to manage a port. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        resource.
        """
        return pulumi.get(self, "region")

    @_builtins.property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        A list of security group IDs to apply to
        the port. The security groups must be specified by ID and not name (as
        opposed to how they are configured with the Compute Instance).
        """
        return pulumi.get(self, "security_group_ids")

