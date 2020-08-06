# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ShareNetwork(pulumi.CustomResource):
    cidr: pulumi.Output[str]
    """
    The share network CIDR.
    """
    description: pulumi.Output[str]
    """
    The human-readable description for the share network.
    Changing this updates the description of the existing share network.
    """
    ip_version: pulumi.Output[float]
    """
    The IP version of the share network. Can either be 4 or 6.
    """
    name: pulumi.Output[str]
    """
    The name for the share network. Changing this updates the name
    of the existing share network.
    """
    network_type: pulumi.Output[str]
    """
    The share network type. Can either be VLAN, VXLAN, GRE, or flat.
    """
    neutron_net_id: pulumi.Output[str]
    """
    The UUID of a neutron network when setting up or updating
    a share network. Changing this updates the existing share network if it's not used by
    shares.
    """
    neutron_subnet_id: pulumi.Output[str]
    """
    The UUID of the neutron subnet when setting up or
    updating a share network. Changing this updates the existing share network if it's
    not used by shares.
    """
    project_id: pulumi.Output[str]
    """
    The owner of the Share Network.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Shared File System client.
    A Shared File System client is needed to create a share network. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    share network.
    """
    security_service_ids: pulumi.Output[list]
    """
    The list of security service IDs to associate with
    the share network. The security service must be specified by ID and not name.
    """
    segmentation_id: pulumi.Output[float]
    """
    The share network segmentation ID.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, neutron_net_id=None, neutron_subnet_id=None, region=None, security_service_ids=None, __props__=None, __name__=None, __opts__=None):
        """
        Use this resource to configure a share network.

        A share network stores network information that share servers can use when
        shares are created.

        ## Example Usage
        ### Basic share network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up="true")
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        sharenetwork1 = openstack.sharedfilesystem.ShareNetwork("sharenetwork1",
            description="test share network",
            neutron_net_id=network1.id,
            neutron_subnet_id=subnet1.id)
        ```
        ### Share network with associated security services

        ```python
        import pulumi
        import pulumi_openstack as openstack

        network1 = openstack.networking.Network("network1", admin_state_up="true")
        subnet1 = openstack.networking.Subnet("subnet1",
            cidr="192.168.199.0/24",
            ip_version=4,
            network_id=network1.id)
        securityservice1 = openstack.sharedfilesystem.SecurityService("securityservice1",
            description="created by terraform",
            dns_ip="192.168.199.10",
            domain="example.com",
            ou="CN=Computers,DC=example,DC=com",
            password="s8cret",
            server="192.168.199.10",
            type="active_directory",
            user="joinDomainUser")
        sharenetwork1 = openstack.sharedfilesystem.ShareNetwork("sharenetwork1",
            description="test share network with security services",
            neutron_net_id=network1.id,
            neutron_subnet_id=subnet1.id,
            security_service_ids=[securityservice1.id])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The human-readable description for the share network.
               Changing this updates the description of the existing share network.
        :param pulumi.Input[str] name: The name for the share network. Changing this updates the name
               of the existing share network.
        :param pulumi.Input[str] neutron_net_id: The UUID of a neutron network when setting up or updating
               a share network. Changing this updates the existing share network if it's not used by
               shares.
        :param pulumi.Input[str] neutron_subnet_id: The UUID of the neutron subnet when setting up or
               updating a share network. Changing this updates the existing share network if it's
               not used by shares.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Shared File System client.
               A Shared File System client is needed to create a share network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               share network.
        :param pulumi.Input[list] security_service_ids: The list of security service IDs to associate with
               the share network. The security service must be specified by ID and not name.
        """
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

            __props__['description'] = description
            __props__['name'] = name
            if neutron_net_id is None:
                raise TypeError("Missing required property 'neutron_net_id'")
            __props__['neutron_net_id'] = neutron_net_id
            if neutron_subnet_id is None:
                raise TypeError("Missing required property 'neutron_subnet_id'")
            __props__['neutron_subnet_id'] = neutron_subnet_id
            __props__['region'] = region
            __props__['security_service_ids'] = security_service_ids
            __props__['cidr'] = None
            __props__['ip_version'] = None
            __props__['network_type'] = None
            __props__['project_id'] = None
            __props__['segmentation_id'] = None
        super(ShareNetwork, __self__).__init__(
            'openstack:sharedfilesystem/shareNetwork:ShareNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cidr=None, description=None, ip_version=None, name=None, network_type=None, neutron_net_id=None, neutron_subnet_id=None, project_id=None, region=None, security_service_ids=None, segmentation_id=None):
        """
        Get an existing ShareNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The share network CIDR.
        :param pulumi.Input[str] description: The human-readable description for the share network.
               Changing this updates the description of the existing share network.
        :param pulumi.Input[float] ip_version: The IP version of the share network. Can either be 4 or 6.
        :param pulumi.Input[str] name: The name for the share network. Changing this updates the name
               of the existing share network.
        :param pulumi.Input[str] network_type: The share network type. Can either be VLAN, VXLAN, GRE, or flat.
        :param pulumi.Input[str] neutron_net_id: The UUID of a neutron network when setting up or updating
               a share network. Changing this updates the existing share network if it's not used by
               shares.
        :param pulumi.Input[str] neutron_subnet_id: The UUID of the neutron subnet when setting up or
               updating a share network. Changing this updates the existing share network if it's
               not used by shares.
        :param pulumi.Input[str] project_id: The owner of the Share Network.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Shared File System client.
               A Shared File System client is needed to create a share network. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               share network.
        :param pulumi.Input[list] security_service_ids: The list of security service IDs to associate with
               the share network. The security service must be specified by ID and not name.
        :param pulumi.Input[float] segmentation_id: The share network segmentation ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cidr"] = cidr
        __props__["description"] = description
        __props__["ip_version"] = ip_version
        __props__["name"] = name
        __props__["network_type"] = network_type
        __props__["neutron_net_id"] = neutron_net_id
        __props__["neutron_subnet_id"] = neutron_subnet_id
        __props__["project_id"] = project_id
        __props__["region"] = region
        __props__["security_service_ids"] = security_service_ids
        __props__["segmentation_id"] = segmentation_id
        return ShareNetwork(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
