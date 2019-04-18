# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SiteConnection(pulumi.CustomResource):
    admin_state_up: pulumi.Output[bool]
    """
    The administrative state of the resource. Can either be up(true) or down(false).
    Changing this updates the administrative state of the existing connection.
    """
    description: pulumi.Output[str]
    """
    The human-readable description for the connection.
    Changing this updates the description of the existing connection.
    """
    dpds: pulumi.Output[list]
    """
    A dictionary with dead peer detection (DPD) protocol controls.
    - `action` - (Optional) The dead peer detection (DPD) action.
    A valid value is clear, hold, restart, disabled, or restart-by-peer.
    Default value is hold.
    """
    ikepolicy_id: pulumi.Output[str]
    """
    The ID of the IKE policy. Changing this creates a new connection.
    """
    initiator: pulumi.Output[str]
    """
    A valid value is response-only or bi-directional. Default is bi-directional.
    """
    ipsecpolicy_id: pulumi.Output[str]
    """
    The ID of the IPsec policy. Changing this creates a new connection.
    """
    local_ep_group_id: pulumi.Output[str]
    """
    The ID for the endpoint group that contains private subnets for the local side of the connection.
    You must specify this parameter with the peer_ep_group_id parameter unless
    in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
    Changing this updates the existing connection.
    """
    local_id: pulumi.Output[str]
    """
    An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
    Most often, local ID would be domain name, email address, etc.
    If this is not configured then the external IP address will be used as the ID.
    """
    mtu: pulumi.Output[float]
    """
    The maximum transmission unit (MTU) value to address fragmentation.
    Minimum value is 68 for IPv4, and 1280 for IPv6.
    """
    name: pulumi.Output[str]
    """
    The name of the connection. Changing this updates the name of
    the existing connection.
    """
    peer_address: pulumi.Output[str]
    """
    The peer gateway public IPv4 or IPv6 address or FQDN.
    """
    peer_cidrs: pulumi.Output[list]
    """
    Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
    """
    peer_ep_group_id: pulumi.Output[str]
    """
    The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
    You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
    where peer_cidrs is provided with a subnet_id for the VPN service.
    """
    peer_id: pulumi.Output[str]
    """
    The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
    Typically, this value matches the peer_address value.
    Changing this updates the existing policy.
    """
    psk: pulumi.Output[str]
    """
    The pre-shared key. A valid value is any string.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Networking client.
    A Networking client is needed to create an IPSec site connection. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    site connection.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the connection. Required if admin wants to
    create a connection for another project. Changing this creates a new connection.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    vpnservice_id: pulumi.Output[str]
    """
    The ID of the VPN service. Changing this creates a new connection.
    """
    def __init__(__self__, resource_name, opts=None, admin_state_up=None, description=None, dpds=None, ikepolicy_id=None, initiator=None, ipsecpolicy_id=None, local_ep_group_id=None, local_id=None, mtu=None, name=None, peer_address=None, peer_cidrs=None, peer_ep_group_id=None, peer_id=None, psk=None, region=None, tenant_id=None, value_specs=None, vpnservice_id=None, __name__=None, __opts__=None):
        """
        Manages a V2 Neutron IPSec site connection resource within OpenStack.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the resource. Can either be up(true) or down(false).
               Changing this updates the administrative state of the existing connection.
        :param pulumi.Input[str] description: The human-readable description for the connection.
               Changing this updates the description of the existing connection.
        :param pulumi.Input[list] dpds: A dictionary with dead peer detection (DPD) protocol controls.
               - `action` - (Optional) The dead peer detection (DPD) action.
               A valid value is clear, hold, restart, disabled, or restart-by-peer.
               Default value is hold.
        :param pulumi.Input[str] ikepolicy_id: The ID of the IKE policy. Changing this creates a new connection.
        :param pulumi.Input[str] initiator: A valid value is response-only or bi-directional. Default is bi-directional.
        :param pulumi.Input[str] ipsecpolicy_id: The ID of the IPsec policy. Changing this creates a new connection.
        :param pulumi.Input[str] local_ep_group_id: The ID for the endpoint group that contains private subnets for the local side of the connection.
               You must specify this parameter with the peer_ep_group_id parameter unless
               in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.
               Changing this updates the existing connection.
        :param pulumi.Input[str] local_id: An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic.
               Most often, local ID would be domain name, email address, etc.
               If this is not configured then the external IP address will be used as the ID.
        :param pulumi.Input[float] mtu: The maximum transmission unit (MTU) value to address fragmentation.
               Minimum value is 68 for IPv4, and 1280 for IPv6.
        :param pulumi.Input[str] name: The name of the connection. Changing this updates the name of
               the existing connection.
        :param pulumi.Input[str] peer_address: The peer gateway public IPv4 or IPv6 address or FQDN.
        :param pulumi.Input[list] peer_cidrs: Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .
        :param pulumi.Input[str] peer_ep_group_id: The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection.
               You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode
               where peer_cidrs is provided with a subnet_id for the VPN service.
        :param pulumi.Input[str] peer_id: The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN.
               Typically, this value matches the peer_address value.
               Changing this updates the existing policy.
        :param pulumi.Input[str] psk: The pre-shared key. A valid value is any string.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an IPSec site connection. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               site connection.
        :param pulumi.Input[str] tenant_id: The owner of the connection. Required if admin wants to
               create a connection for another project. Changing this creates a new connection.
        :param pulumi.Input[dict] value_specs: Map of additional options.
        :param pulumi.Input[str] vpnservice_id: The ID of the VPN service. Changing this creates a new connection.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['admin_state_up'] = admin_state_up

        __props__['description'] = description

        __props__['dpds'] = dpds

        if ikepolicy_id is None:
            raise TypeError("Missing required property 'ikepolicy_id'")
        __props__['ikepolicy_id'] = ikepolicy_id

        __props__['initiator'] = initiator

        if ipsecpolicy_id is None:
            raise TypeError("Missing required property 'ipsecpolicy_id'")
        __props__['ipsecpolicy_id'] = ipsecpolicy_id

        __props__['local_ep_group_id'] = local_ep_group_id

        __props__['local_id'] = local_id

        __props__['mtu'] = mtu

        __props__['name'] = name

        if peer_address is None:
            raise TypeError("Missing required property 'peer_address'")
        __props__['peer_address'] = peer_address

        __props__['peer_cidrs'] = peer_cidrs

        __props__['peer_ep_group_id'] = peer_ep_group_id

        if peer_id is None:
            raise TypeError("Missing required property 'peer_id'")
        __props__['peer_id'] = peer_id

        if psk is None:
            raise TypeError("Missing required property 'psk'")
        __props__['psk'] = psk

        __props__['region'] = region

        __props__['tenant_id'] = tenant_id

        __props__['value_specs'] = value_specs

        if vpnservice_id is None:
            raise TypeError("Missing required property 'vpnservice_id'")
        __props__['vpnservice_id'] = vpnservice_id

        super(SiteConnection, __self__).__init__(
            'openstack:vpnaas/siteConnection:SiteConnection',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

