# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Port(pulumi.CustomResource):
    admin_state_up: pulumi.Output[bool]
    """
    Administrative up/down status for the port
    (must be "true" or "false" if provided). Changing this updates the
    `admin_state_up` of an existing port.
    """
    all_fixed_ips: pulumi.Output[list]
    """
    The collection of Fixed IP addresses on the port in the
    order returned by the Network v2 API.
    """
    all_security_group_ids: pulumi.Output[list]
    """
    The collection of Security Group IDs on the port
    which have been explicitly and implicitly added.
    """
    all_tags: pulumi.Output[list]
    """
    The collection of tags assigned on the port, which have been
    explicitly and implicitly added.
    """
    allowed_address_pairs: pulumi.Output[list]
    """
    An IP/MAC Address pair of additional IP
    addresses that can be active on this port. The structure is described
    below.
    """
    binding: pulumi.Output[dict]
    """
    The port binding allows to specify binding information
    for the port. The structure is described below.
    """
    description: pulumi.Output[str]
    """
    Human-readable description of the floating IP. Changing
    this updates the `description` of an existing port.
    """
    device_id: pulumi.Output[str]
    """
    The ID of the device attached to the port. Changing this
    creates a new port.
    """
    device_owner: pulumi.Output[str]
    """
    The device owner of the Port. Changing this creates
    a new port.
    """
    dns_assignments: pulumi.Output[list]
    """
    The list of maps representing port DNS assignments.
    """
    dns_name: pulumi.Output[str]
    """
    The port DNS name. Available, when Neutron DNS extension
    is enabled.
    """
    extra_dhcp_options: pulumi.Output[list]
    """
    An extra DHCP option that needs to be configured
    on the port. The structure is described below. Can be specified multiple
    times.
    """
    fixed_ips: pulumi.Output[list]
    """
    An array of desired IPs for
    this port. The structure is described below.
    """
    mac_address: pulumi.Output[str]
    """
    Specify a specific MAC address for the port. Changing
    this creates a new port.
    """
    name: pulumi.Output[str]
    """
    A unique name for the port. Changing this
    updates the `name` of an existing port.
    """
    network_id: pulumi.Output[str]
    """
    The ID of the network to attach the port to. Changing
    this creates a new port.
    """
    no_fixed_ip: pulumi.Output[bool]
    """
    Create a port with no fixed
    IP address. This will also remove any fixed IPs previously set on a port. `true`
    is the only valid value for this argument.
    """
    no_security_groups: pulumi.Output[bool]
    """
    If set to
    `true`, then no security groups are applied to the port. If set to `false` and
    no `security_group_ids` are specified, then the Port will yield to the default
    behavior of the Networking service, which is to usually apply the "default"
    security group.
    """
    port_security_enabled: pulumi.Output[bool]
    """
    Whether to explicitly enable or disable
    port security on the port. Port Security is usually enabled by default, so
    omitting argument will usually result in a value of "true". Setting this
    explicitly to `false` will disable port security. In order to disable port
    security, the port must not have any security groups. Valid values are `true`
    and `false`.
    """
    qos_policy_id: pulumi.Output[str]
    """
    Reference to the associated QoS policy.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 networking client.
    A networking client is needed to create a port. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    port.
    """
    security_group_ids: pulumi.Output[list]
    """
    A list
    of security group IDs to apply to the port. The security groups must be
    specified by ID and not name (as opposed to how they are configured with
    the Compute Instance).
    """
    tags: pulumi.Output[list]
    """
    A set of string tags for the port.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the Port. Required if admin wants
    to create a port for another tenant. Changing this creates a new port.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional options.
    """
    def __init__(__self__, resource_name, opts=None, admin_state_up=None, allowed_address_pairs=None, binding=None, description=None, device_id=None, device_owner=None, dns_name=None, extra_dhcp_options=None, fixed_ips=None, mac_address=None, name=None, network_id=None, no_fixed_ip=None, no_security_groups=None, port_security_enabled=None, qos_policy_id=None, region=None, security_group_ids=None, tags=None, tenant_id=None, value_specs=None, __name__=None, __opts__=None):
        """
        Manages a V2 port resource within OpenStack.
        
        ## Notes
        
        ### Ports and Instances
        
        There are some notes to consider when connecting Instances to networks using
        Ports. Please see the `openstack_compute_instance_v2` documentation for further
        documentation.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the port
               (must be "true" or "false" if provided). Changing this updates the
               `admin_state_up` of an existing port.
        :param pulumi.Input[list] allowed_address_pairs: An IP/MAC Address pair of additional IP
               addresses that can be active on this port. The structure is described
               below.
        :param pulumi.Input[dict] binding: The port binding allows to specify binding information
               for the port. The structure is described below.
        :param pulumi.Input[str] description: Human-readable description of the floating IP. Changing
               this updates the `description` of an existing port.
        :param pulumi.Input[str] device_id: The ID of the device attached to the port. Changing this
               creates a new port.
        :param pulumi.Input[str] device_owner: The device owner of the Port. Changing this creates
               a new port.
        :param pulumi.Input[str] dns_name: The port DNS name. Available, when Neutron DNS extension
               is enabled.
        :param pulumi.Input[list] extra_dhcp_options: An extra DHCP option that needs to be configured
               on the port. The structure is described below. Can be specified multiple
               times.
        :param pulumi.Input[list] fixed_ips: An array of desired IPs for
               this port. The structure is described below.
        :param pulumi.Input[str] mac_address: Specify a specific MAC address for the port. Changing
               this creates a new port.
        :param pulumi.Input[str] name: A unique name for the port. Changing this
               updates the `name` of an existing port.
        :param pulumi.Input[str] network_id: The ID of the network to attach the port to. Changing
               this creates a new port.
        :param pulumi.Input[bool] no_fixed_ip: Create a port with no fixed
               IP address. This will also remove any fixed IPs previously set on a port. `true`
               is the only valid value for this argument.
        :param pulumi.Input[bool] no_security_groups: If set to
               `true`, then no security groups are applied to the port. If set to `false` and
               no `security_group_ids` are specified, then the Port will yield to the default
               behavior of the Networking service, which is to usually apply the "default"
               security group.
        :param pulumi.Input[bool] port_security_enabled: Whether to explicitly enable or disable
               port security on the port. Port Security is usually enabled by default, so
               omitting argument will usually result in a value of "true". Setting this
               explicitly to `false` will disable port security. In order to disable port
               security, the port must not have any security groups. Valid values are `true`
               and `false`.
        :param pulumi.Input[str] qos_policy_id: Reference to the associated QoS policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a port. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               port.
        :param pulumi.Input[list] security_group_ids: A list
               of security group IDs to apply to the port. The security groups must be
               specified by ID and not name (as opposed to how they are configured with
               the Compute Instance).
        :param pulumi.Input[list] tags: A set of string tags for the port.
        :param pulumi.Input[str] tenant_id: The owner of the Port. Required if admin wants
               to create a port for another tenant. Changing this creates a new port.
        :param pulumi.Input[dict] value_specs: Map of additional options.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/networking_port_v2.html.markdown.
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

        __props__['allowed_address_pairs'] = allowed_address_pairs

        __props__['binding'] = binding

        __props__['description'] = description

        __props__['device_id'] = device_id

        __props__['device_owner'] = device_owner

        __props__['dns_name'] = dns_name

        __props__['extra_dhcp_options'] = extra_dhcp_options

        __props__['fixed_ips'] = fixed_ips

        __props__['mac_address'] = mac_address

        __props__['name'] = name

        if network_id is None:
            raise TypeError("Missing required property 'network_id'")
        __props__['network_id'] = network_id

        __props__['no_fixed_ip'] = no_fixed_ip

        __props__['no_security_groups'] = no_security_groups

        __props__['port_security_enabled'] = port_security_enabled

        __props__['qos_policy_id'] = qos_policy_id

        __props__['region'] = region

        __props__['security_group_ids'] = security_group_ids

        __props__['tags'] = tags

        __props__['tenant_id'] = tenant_id

        __props__['value_specs'] = value_specs

        __props__['all_fixed_ips'] = None
        __props__['all_security_group_ids'] = None
        __props__['all_tags'] = None
        __props__['dns_assignments'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Port, __self__).__init__(
            'openstack:networking/port:Port',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

