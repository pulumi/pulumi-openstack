# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class InterfaceAttach(pulumi.CustomResource):
    fixed_ip: pulumi.Output[str]
    """
    An IP address to assosciate with the port.
    _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
    """
    instance_id: pulumi.Output[str]
    """
    The ID of the Instance to attach the Port or Network to.
    """
    network_id: pulumi.Output[str]
    """
    The ID of the Network to attach to an Instance. A port will be created automatically.
    _NOTE_: This option and `port_id` are mutually exclusive.
    """
    port_id: pulumi.Output[str]
    """
    The ID of the Port to attach to an Instance.
    _NOTE_: This option and `network_id` are mutually exclusive.
    """
    region: pulumi.Output[str]
    """
    The region in which to create the interface attachment.
    If omitted, the `region` argument of the provider is used. Changing this
    creates a new attachment.
    """
    def __init__(__self__, resource_name, opts=None, fixed_ip=None, instance_id=None, network_id=None, port_id=None, region=None, __name__=None, __opts__=None):
        """
        Attaches a Network Interface (a Port) to an Instance using the OpenStack
        Compute (Nova) v2 API.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fixed_ip: An IP address to assosciate with the port.
               _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network.
        :param pulumi.Input[str] instance_id: The ID of the Instance to attach the Port or Network to.
        :param pulumi.Input[str] network_id: The ID of the Network to attach to an Instance. A port will be created automatically.
               _NOTE_: This option and `port_id` are mutually exclusive.
        :param pulumi.Input[str] port_id: The ID of the Port to attach to an Instance.
               _NOTE_: This option and `network_id` are mutually exclusive.
        :param pulumi.Input[str] region: The region in which to create the interface attachment.
               If omitted, the `region` argument of the provider is used. Changing this
               creates a new attachment.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/compute_interface_attach_v2.html.markdown.
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

        __props__['fixed_ip'] = fixed_ip

        if instance_id is None:
            raise TypeError("Missing required property 'instance_id'")
        __props__['instance_id'] = instance_id

        __props__['network_id'] = network_id

        __props__['port_id'] = port_id

        __props__['region'] = region

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(InterfaceAttach, __self__).__init__(
            'openstack:compute/interfaceAttach:InterfaceAttach',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

