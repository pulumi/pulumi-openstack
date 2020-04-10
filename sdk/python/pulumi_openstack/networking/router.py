# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Router(pulumi.CustomResource):
    admin_state_up: pulumi.Output[bool]
    """
    Administrative up/down status for the router
    (must be "true" or "false" if provided). Changing this updates the
    `admin_state_up` of an existing router.
    """
    all_tags: pulumi.Output[list]
    """
    The collection of tags assigned on the router, which have been
    explicitly and implicitly added.
    """
    availability_zone_hints: pulumi.Output[list]
    """
    An availability zone is used to make 
    network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
    this creates a new router.
    """
    description: pulumi.Output[str]
    """
    Human-readable description for the router.
    """
    distributed: pulumi.Output[bool]
    """
    Indicates whether or not to create a
    distributed router. The default policy setting in Neutron restricts
    usage of this property to administrative users only.
    """
    enable_snat: pulumi.Output[bool]
    """
    Enable Source NAT for the router. Valid values are
    "true" or "false". An `external_network_id` has to be set in order to
    set this property. Changing this updates the `enable_snat` of the router.
    Setting this value **requires** an **ext-gw-mode** extension to be enabled
    in OpenStack Neutron.
    """
    external_fixed_ips: pulumi.Output[list]
    """
    An external fixed IP for the router. This
    can be repeated. The structure is described below. An `external_network_id`
    has to be set in order to set this property. Changing this updates the
    external fixed IPs of the router.

      * `ip_address` (`str`) - The IP address to set on the router.
      * `subnet_id` (`str`) - Subnet in which the fixed IP belongs to.
    """
    external_gateway: pulumi.Output[str]
    """
    The
    network UUID of an external gateway for the router. A router with an
    external gateway is required if any compute instances or load balancers
    will be using floating IPs. Changing this updates the external gateway
    of an existing router.
    """
    external_network_id: pulumi.Output[str]
    """
    The network UUID of an external gateway
    for the router. A router with an external gateway is required if any
    compute instances or load balancers will be using floating IPs. Changing
    this updates the external gateway of the router.
    """
    name: pulumi.Output[str]
    """
    A unique name for the router. Changing this
    updates the `name` of an existing router.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 networking client.
    A networking client is needed to create a router. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    router.
    """
    tags: pulumi.Output[list]
    """
    A set of string tags for the router.
    """
    tenant_id: pulumi.Output[str]
    """
    The owner of the floating IP. Required if admin wants
    to create a router for another tenant. Changing this creates a new router.
    """
    value_specs: pulumi.Output[dict]
    """
    Map of additional driver-specific options.
    """
    vendor_options: pulumi.Output[dict]
    """
    Map of additional vendor-specific options.
    Supported options are described below.

      * `setRouterGatewayAfterCreate` (`bool`) - Boolean to control whether
        the Router gateway is assigned during creation or updated after creation.
    """
    def __init__(__self__, resource_name, opts=None, admin_state_up=None, availability_zone_hints=None, description=None, distributed=None, enable_snat=None, external_fixed_ips=None, external_gateway=None, external_network_id=None, name=None, region=None, tags=None, tenant_id=None, value_specs=None, vendor_options=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V2 router resource within OpenStack.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the router
               (must be "true" or "false" if provided). Changing this updates the
               `admin_state_up` of an existing router.
        :param pulumi.Input[list] availability_zone_hints: An availability zone is used to make 
               network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
               this creates a new router.
        :param pulumi.Input[str] description: Human-readable description for the router.
        :param pulumi.Input[bool] distributed: Indicates whether or not to create a
               distributed router. The default policy setting in Neutron restricts
               usage of this property to administrative users only.
        :param pulumi.Input[bool] enable_snat: Enable Source NAT for the router. Valid values are
               "true" or "false". An `external_network_id` has to be set in order to
               set this property. Changing this updates the `enable_snat` of the router.
               Setting this value **requires** an **ext-gw-mode** extension to be enabled
               in OpenStack Neutron.
        :param pulumi.Input[list] external_fixed_ips: An external fixed IP for the router. This
               can be repeated. The structure is described below. An `external_network_id`
               has to be set in order to set this property. Changing this updates the
               external fixed IPs of the router.
        :param pulumi.Input[str] external_gateway: The
               network UUID of an external gateway for the router. A router with an
               external gateway is required if any compute instances or load balancers
               will be using floating IPs. Changing this updates the external gateway
               of an existing router.
        :param pulumi.Input[str] external_network_id: The network UUID of an external gateway
               for the router. A router with an external gateway is required if any
               compute instances or load balancers will be using floating IPs. Changing
               this updates the external gateway of the router.
        :param pulumi.Input[str] name: A unique name for the router. Changing this
               updates the `name` of an existing router.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               router.
        :param pulumi.Input[list] tags: A set of string tags for the router.
        :param pulumi.Input[str] tenant_id: The owner of the floating IP. Required if admin wants
               to create a router for another tenant. Changing this creates a new router.
        :param pulumi.Input[dict] value_specs: Map of additional driver-specific options.
        :param pulumi.Input[dict] vendor_options: Map of additional vendor-specific options.
               Supported options are described below.

        The **external_fixed_ips** object supports the following:

          * `ip_address` (`pulumi.Input[str]`) - The IP address to set on the router.
          * `subnet_id` (`pulumi.Input[str]`) - Subnet in which the fixed IP belongs to.

        The **vendor_options** object supports the following:

          * `setRouterGatewayAfterCreate` (`pulumi.Input[bool]`) - Boolean to control whether
            the Router gateway is assigned during creation or updated after creation.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin_state_up'] = admin_state_up
            __props__['availability_zone_hints'] = availability_zone_hints
            __props__['description'] = description
            __props__['distributed'] = distributed
            __props__['enable_snat'] = enable_snat
            __props__['external_fixed_ips'] = external_fixed_ips
            __props__['external_gateway'] = external_gateway
            __props__['external_network_id'] = external_network_id
            __props__['name'] = name
            __props__['region'] = region
            __props__['tags'] = tags
            __props__['tenant_id'] = tenant_id
            __props__['value_specs'] = value_specs
            __props__['vendor_options'] = vendor_options
            __props__['all_tags'] = None
        super(Router, __self__).__init__(
            'openstack:networking/router:Router',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admin_state_up=None, all_tags=None, availability_zone_hints=None, description=None, distributed=None, enable_snat=None, external_fixed_ips=None, external_gateway=None, external_network_id=None, name=None, region=None, tags=None, tenant_id=None, value_specs=None, vendor_options=None):
        """
        Get an existing Router resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: Administrative up/down status for the router
               (must be "true" or "false" if provided). Changing this updates the
               `admin_state_up` of an existing router.
        :param pulumi.Input[list] all_tags: The collection of tags assigned on the router, which have been
               explicitly and implicitly added.
        :param pulumi.Input[list] availability_zone_hints: An availability zone is used to make 
               network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing
               this creates a new router.
        :param pulumi.Input[str] description: Human-readable description for the router.
        :param pulumi.Input[bool] distributed: Indicates whether or not to create a
               distributed router. The default policy setting in Neutron restricts
               usage of this property to administrative users only.
        :param pulumi.Input[bool] enable_snat: Enable Source NAT for the router. Valid values are
               "true" or "false". An `external_network_id` has to be set in order to
               set this property. Changing this updates the `enable_snat` of the router.
               Setting this value **requires** an **ext-gw-mode** extension to be enabled
               in OpenStack Neutron.
        :param pulumi.Input[list] external_fixed_ips: An external fixed IP for the router. This
               can be repeated. The structure is described below. An `external_network_id`
               has to be set in order to set this property. Changing this updates the
               external fixed IPs of the router.
        :param pulumi.Input[str] external_gateway: The
               network UUID of an external gateway for the router. A router with an
               external gateway is required if any compute instances or load balancers
               will be using floating IPs. Changing this updates the external gateway
               of an existing router.
        :param pulumi.Input[str] external_network_id: The network UUID of an external gateway
               for the router. A router with an external gateway is required if any
               compute instances or load balancers will be using floating IPs. Changing
               this updates the external gateway of the router.
        :param pulumi.Input[str] name: A unique name for the router. Changing this
               updates the `name` of an existing router.
        :param pulumi.Input[str] region: The region in which to obtain the V2 networking client.
               A networking client is needed to create a router. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               router.
        :param pulumi.Input[list] tags: A set of string tags for the router.
        :param pulumi.Input[str] tenant_id: The owner of the floating IP. Required if admin wants
               to create a router for another tenant. Changing this creates a new router.
        :param pulumi.Input[dict] value_specs: Map of additional driver-specific options.
        :param pulumi.Input[dict] vendor_options: Map of additional vendor-specific options.
               Supported options are described below.

        The **external_fixed_ips** object supports the following:

          * `ip_address` (`pulumi.Input[str]`) - The IP address to set on the router.
          * `subnet_id` (`pulumi.Input[str]`) - Subnet in which the fixed IP belongs to.

        The **vendor_options** object supports the following:

          * `setRouterGatewayAfterCreate` (`pulumi.Input[bool]`) - Boolean to control whether
            the Router gateway is assigned during creation or updated after creation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_state_up"] = admin_state_up
        __props__["all_tags"] = all_tags
        __props__["availability_zone_hints"] = availability_zone_hints
        __props__["description"] = description
        __props__["distributed"] = distributed
        __props__["enable_snat"] = enable_snat
        __props__["external_fixed_ips"] = external_fixed_ips
        __props__["external_gateway"] = external_gateway
        __props__["external_network_id"] = external_network_id
        __props__["name"] = name
        __props__["region"] = region
        __props__["tags"] = tags
        __props__["tenant_id"] = tenant_id
        __props__["value_specs"] = value_specs
        __props__["vendor_options"] = vendor_options
        return Router(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

