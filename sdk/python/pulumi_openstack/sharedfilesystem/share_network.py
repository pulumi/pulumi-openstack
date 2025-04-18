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

__all__ = ['ShareNetworkArgs', 'ShareNetwork']

@pulumi.input_type
class ShareNetworkArgs:
    def __init__(__self__, *,
                 neutron_net_id: pulumi.Input[builtins.str],
                 neutron_subnet_id: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a ShareNetwork resource.
        :param pulumi.Input[builtins.str] neutron_net_id: The UUID of a neutron network when setting up or updating
               a share network. Changing this updates the existing share network if it's not used by
               shares.
        :param pulumi.Input[builtins.str] neutron_subnet_id: The UUID of the neutron subnet when setting up or
               updating a share network. Changing this updates the existing share network if it's
               not used by shares.
        :param pulumi.Input[builtins.str] description: The human-readable description for the share network.
               Changing this updates the description of the existing share network.
        :param pulumi.Input[builtins.str] name: The name for the share network. Changing this updates the name
               of the existing share network.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Shared File System client.
               A Shared File System client is needed to create a share network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               share network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] security_service_ids: The list of security service IDs to associate with
               the share network. The security service must be specified by ID and not name.
        """
        pulumi.set(__self__, "neutron_net_id", neutron_net_id)
        pulumi.set(__self__, "neutron_subnet_id", neutron_subnet_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_service_ids is not None:
            pulumi.set(__self__, "security_service_ids", security_service_ids)

    @property
    @pulumi.getter(name="neutronNetId")
    def neutron_net_id(self) -> pulumi.Input[builtins.str]:
        """
        The UUID of a neutron network when setting up or updating
        a share network. Changing this updates the existing share network if it's not used by
        shares.
        """
        return pulumi.get(self, "neutron_net_id")

    @neutron_net_id.setter
    def neutron_net_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "neutron_net_id", value)

    @property
    @pulumi.getter(name="neutronSubnetId")
    def neutron_subnet_id(self) -> pulumi.Input[builtins.str]:
        """
        The UUID of the neutron subnet when setting up or
        updating a share network. Changing this updates the existing share network if it's
        not used by shares.
        """
        return pulumi.get(self, "neutron_subnet_id")

    @neutron_subnet_id.setter
    def neutron_subnet_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "neutron_subnet_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The human-readable description for the share network.
        Changing this updates the description of the existing share network.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name for the share network. Changing this updates the name
        of the existing share network.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Shared File System client.
        A Shared File System client is needed to create a share network. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        share network.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityServiceIds")
    def security_service_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The list of security service IDs to associate with
        the share network. The security service must be specified by ID and not name.
        """
        return pulumi.get(self, "security_service_ids")

    @security_service_ids.setter
    def security_service_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "security_service_ids", value)


@pulumi.input_type
class _ShareNetworkState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 ip_version: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 network_type: Optional[pulumi.Input[builtins.str]] = None,
                 neutron_net_id: Optional[pulumi.Input[builtins.str]] = None,
                 neutron_subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 segmentation_id: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering ShareNetwork resources.
        :param pulumi.Input[builtins.str] cidr: The share network CIDR.
        :param pulumi.Input[builtins.str] description: The human-readable description for the share network.
               Changing this updates the description of the existing share network.
        :param pulumi.Input[builtins.int] ip_version: The IP version of the share network. Can either be 4 or 6.
        :param pulumi.Input[builtins.str] name: The name for the share network. Changing this updates the name
               of the existing share network.
        :param pulumi.Input[builtins.str] network_type: The share network type. Can either be VLAN, VXLAN, GRE, or flat.
        :param pulumi.Input[builtins.str] neutron_net_id: The UUID of a neutron network when setting up or updating
               a share network. Changing this updates the existing share network if it's not used by
               shares.
        :param pulumi.Input[builtins.str] neutron_subnet_id: The UUID of the neutron subnet when setting up or
               updating a share network. Changing this updates the existing share network if it's
               not used by shares.
        :param pulumi.Input[builtins.str] project_id: The owner of the Share Network.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Shared File System client.
               A Shared File System client is needed to create a share network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               share network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] security_service_ids: The list of security service IDs to associate with
               the share network. The security service must be specified by ID and not name.
        :param pulumi.Input[builtins.int] segmentation_id: The share network segmentation ID.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if neutron_net_id is not None:
            pulumi.set(__self__, "neutron_net_id", neutron_net_id)
        if neutron_subnet_id is not None:
            pulumi.set(__self__, "neutron_subnet_id", neutron_subnet_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_service_ids is not None:
            pulumi.set(__self__, "security_service_ids", security_service_ids)
        if segmentation_id is not None:
            pulumi.set(__self__, "segmentation_id", segmentation_id)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The share network CIDR.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The human-readable description for the share network.
        Changing this updates the description of the existing share network.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The IP version of the share network. Can either be 4 or 6.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name for the share network. Changing this updates the name
        of the existing share network.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The share network type. Can either be VLAN, VXLAN, GRE, or flat.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter(name="neutronNetId")
    def neutron_net_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The UUID of a neutron network when setting up or updating
        a share network. Changing this updates the existing share network if it's not used by
        shares.
        """
        return pulumi.get(self, "neutron_net_id")

    @neutron_net_id.setter
    def neutron_net_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "neutron_net_id", value)

    @property
    @pulumi.getter(name="neutronSubnetId")
    def neutron_subnet_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The UUID of the neutron subnet when setting up or
        updating a share network. Changing this updates the existing share network if it's
        not used by shares.
        """
        return pulumi.get(self, "neutron_subnet_id")

    @neutron_subnet_id.setter
    def neutron_subnet_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "neutron_subnet_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The owner of the Share Network.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The region in which to obtain the V2 Shared File System client.
        A Shared File System client is needed to create a share network. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        share network.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityServiceIds")
    def security_service_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The list of security service IDs to associate with
        the share network. The security service must be specified by ID and not name.
        """
        return pulumi.get(self, "security_service_ids")

    @security_service_ids.setter
    def security_service_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "security_service_ids", value)

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The share network segmentation ID.
        """
        return pulumi.get(self, "segmentation_id")

    @segmentation_id.setter
    def segmentation_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "segmentation_id", value)


class ShareNetwork(pulumi.CustomResource):

    pulumi_type = "openstack:sharedfilesystem/shareNetwork:ShareNetwork"

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 neutron_net_id: Optional[pulumi.Input[builtins.str]] = None,
                 neutron_subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Use this resource to configure a share network.

        A share network stores network information that share servers can use when
        shares are created.

        ## Example Usage

        ### Basic share network

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
        sharenetwork1 = openstack.sharedfilesystem.ShareNetwork("sharenetwork_1",
            name="test_sharenetwork",
            description="test share network",
            neutron_net_id=network1.id,
            neutron_subnet_id=subnet1.id)
        ```

        ### Share network with associated security services

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
        securityservice1 = openstack.sharedfilesystem.SecurityService("securityservice_1",
            name="security",
            description="created by terraform",
            type="active_directory",
            server="192.168.199.10",
            dns_ip="192.168.199.10",
            domain="example.com",
            ou="CN=Computers,DC=example,DC=com",
            user="joinDomainUser",
            password="s8cret")
        sharenetwork1 = openstack.sharedfilesystem.ShareNetwork("sharenetwork_1",
            name="test_sharenetwork",
            description="test share network with security services",
            neutron_net_id=network1.id,
            neutron_subnet_id=subnet1.id,
            security_service_ids=[securityservice1.id])
        ```

        ## Import

        This resource can be imported by specifying the ID of the share network:

        ```sh
        $ pulumi import openstack:sharedfilesystem/shareNetwork:ShareNetwork sharenetwork_1 id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The human-readable description for the share network.
               Changing this updates the description of the existing share network.
        :param pulumi.Input[builtins.str] name: The name for the share network. Changing this updates the name
               of the existing share network.
        :param pulumi.Input[builtins.str] neutron_net_id: The UUID of a neutron network when setting up or updating
               a share network. Changing this updates the existing share network if it's not used by
               shares.
        :param pulumi.Input[builtins.str] neutron_subnet_id: The UUID of the neutron subnet when setting up or
               updating a share network. Changing this updates the existing share network if it's
               not used by shares.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Shared File System client.
               A Shared File System client is needed to create a share network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               share network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] security_service_ids: The list of security service IDs to associate with
               the share network. The security service must be specified by ID and not name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ShareNetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to configure a share network.

        A share network stores network information that share servers can use when
        shares are created.

        ## Example Usage

        ### Basic share network

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
        sharenetwork1 = openstack.sharedfilesystem.ShareNetwork("sharenetwork_1",
            name="test_sharenetwork",
            description="test share network",
            neutron_net_id=network1.id,
            neutron_subnet_id=subnet1.id)
        ```

        ### Share network with associated security services

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
        securityservice1 = openstack.sharedfilesystem.SecurityService("securityservice_1",
            name="security",
            description="created by terraform",
            type="active_directory",
            server="192.168.199.10",
            dns_ip="192.168.199.10",
            domain="example.com",
            ou="CN=Computers,DC=example,DC=com",
            user="joinDomainUser",
            password="s8cret")
        sharenetwork1 = openstack.sharedfilesystem.ShareNetwork("sharenetwork_1",
            name="test_sharenetwork",
            description="test share network with security services",
            neutron_net_id=network1.id,
            neutron_subnet_id=subnet1.id,
            security_service_ids=[securityservice1.id])
        ```

        ## Import

        This resource can be imported by specifying the ID of the share network:

        ```sh
        $ pulumi import openstack:sharedfilesystem/shareNetwork:ShareNetwork sharenetwork_1 id
        ```

        :param str resource_name: The name of the resource.
        :param ShareNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShareNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 neutron_net_id: Optional[pulumi.Input[builtins.str]] = None,
                 neutron_subnet_id: Optional[pulumi.Input[builtins.str]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 security_service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ShareNetworkArgs.__new__(ShareNetworkArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if neutron_net_id is None and not opts.urn:
                raise TypeError("Missing required property 'neutron_net_id'")
            __props__.__dict__["neutron_net_id"] = neutron_net_id
            if neutron_subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'neutron_subnet_id'")
            __props__.__dict__["neutron_subnet_id"] = neutron_subnet_id
            __props__.__dict__["region"] = region
            __props__.__dict__["security_service_ids"] = security_service_ids
            __props__.__dict__["cidr"] = None
            __props__.__dict__["ip_version"] = None
            __props__.__dict__["network_type"] = None
            __props__.__dict__["project_id"] = None
            __props__.__dict__["segmentation_id"] = None
        super(ShareNetwork, __self__).__init__(
            'openstack:sharedfilesystem/shareNetwork:ShareNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            ip_version: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            network_type: Optional[pulumi.Input[builtins.str]] = None,
            neutron_net_id: Optional[pulumi.Input[builtins.str]] = None,
            neutron_subnet_id: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None,
            security_service_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            segmentation_id: Optional[pulumi.Input[builtins.int]] = None) -> 'ShareNetwork':
        """
        Get an existing ShareNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cidr: The share network CIDR.
        :param pulumi.Input[builtins.str] description: The human-readable description for the share network.
               Changing this updates the description of the existing share network.
        :param pulumi.Input[builtins.int] ip_version: The IP version of the share network. Can either be 4 or 6.
        :param pulumi.Input[builtins.str] name: The name for the share network. Changing this updates the name
               of the existing share network.
        :param pulumi.Input[builtins.str] network_type: The share network type. Can either be VLAN, VXLAN, GRE, or flat.
        :param pulumi.Input[builtins.str] neutron_net_id: The UUID of a neutron network when setting up or updating
               a share network. Changing this updates the existing share network if it's not used by
               shares.
        :param pulumi.Input[builtins.str] neutron_subnet_id: The UUID of the neutron subnet when setting up or
               updating a share network. Changing this updates the existing share network if it's
               not used by shares.
        :param pulumi.Input[builtins.str] project_id: The owner of the Share Network.
        :param pulumi.Input[builtins.str] region: The region in which to obtain the V2 Shared File System client.
               A Shared File System client is needed to create a share network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               share network.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] security_service_ids: The list of security service IDs to associate with
               the share network. The security service must be specified by ID and not name.
        :param pulumi.Input[builtins.int] segmentation_id: The share network segmentation ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ShareNetworkState.__new__(_ShareNetworkState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["description"] = description
        __props__.__dict__["ip_version"] = ip_version
        __props__.__dict__["name"] = name
        __props__.__dict__["network_type"] = network_type
        __props__.__dict__["neutron_net_id"] = neutron_net_id
        __props__.__dict__["neutron_subnet_id"] = neutron_subnet_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["security_service_ids"] = security_service_ids
        __props__.__dict__["segmentation_id"] = segmentation_id
        return ShareNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[builtins.str]:
        """
        The share network CIDR.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The human-readable description for the share network.
        Changing this updates the description of the existing share network.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[builtins.int]:
        """
        The IP version of the share network. Can either be 4 or 6.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name for the share network. Changing this updates the name
        of the existing share network.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> pulumi.Output[builtins.str]:
        """
        The share network type. Can either be VLAN, VXLAN, GRE, or flat.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter(name="neutronNetId")
    def neutron_net_id(self) -> pulumi.Output[builtins.str]:
        """
        The UUID of a neutron network when setting up or updating
        a share network. Changing this updates the existing share network if it's not used by
        shares.
        """
        return pulumi.get(self, "neutron_net_id")

    @property
    @pulumi.getter(name="neutronSubnetId")
    def neutron_subnet_id(self) -> pulumi.Output[builtins.str]:
        """
        The UUID of the neutron subnet when setting up or
        updating a share network. Changing this updates the existing share network if it's
        not used by shares.
        """
        return pulumi.get(self, "neutron_subnet_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The owner of the Share Network.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        The region in which to obtain the V2 Shared File System client.
        A Shared File System client is needed to create a share network. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        share network.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityServiceIds")
    def security_service_ids(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The list of security service IDs to associate with
        the share network. The security service must be specified by ID and not name.
        """
        return pulumi.get(self, "security_service_ids")

    @property
    @pulumi.getter(name="segmentationId")
    def segmentation_id(self) -> pulumi.Output[builtins.int]:
        """
        The share network segmentation ID.
        """
        return pulumi.get(self, "segmentation_id")

