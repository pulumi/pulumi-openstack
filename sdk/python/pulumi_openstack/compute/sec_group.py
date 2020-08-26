# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SecGroup']


class SecGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SecGroupRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V2 security group resource within OpenStack.

        Please note that managing security groups through the OpenStack Compute API
        has been deprecated. Unless you are using an older OpenStack environment, it is
        recommended to use the `networking.SecGroup`
        and `networking.SecGroupRule`
        resources instead, which uses the OpenStack Networking API.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        secgroup1 = openstack.compute.SecGroup("secgroup1",
            description="my security group",
            rules=[
                openstack.compute.SecGroupRuleArgs(
                    cidr="0.0.0.0/0",
                    from_port=22,
                    ip_protocol="tcp",
                    to_port=22,
                ),
                openstack.compute.SecGroupRuleArgs(
                    cidr="0.0.0.0/0",
                    from_port=80,
                    ip_protocol="tcp",
                    to_port=80,
                ),
            ])
        ```
        ## Notes

        ### ICMP Rules

        When using ICMP as the `ip_protocol`, the `from_port` sets the ICMP _type_ and the `to_port` sets the ICMP _code_. To allow all ICMP types, set each value to `-1`, like so:

        ```python
        import pulumi
        ```

        A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages).

        ### Referencing Security Groups

        When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this:

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_server = openstack.compute.Instance("test-server",
            flavor_id="3",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            key_pair="my_key_pair_name",
            security_groups=[openstack_compute_secgroup_v2["secgroup_1"]["name"]])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the security group. Changing this
               updates the `description` of an existing security group.
        :param pulumi.Input[str] name: A unique name for the security group. Changing this
               updates the `name` of an existing security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a security group. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SecGroupRuleArgs']]]] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.
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

            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            __props__['name'] = name
            __props__['region'] = region
            __props__['rules'] = rules
        super(SecGroup, __self__).__init__(
            'openstack:compute/secGroup:SecGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SecGroupRuleArgs']]]]] = None) -> 'SecGroup':
        """
        Get an existing SecGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the security group. Changing this
               updates the `description` of an existing security group.
        :param pulumi.Input[str] name: A unique name for the security group. Changing this
               updates the `name` of an existing security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a security group. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SecGroupRuleArgs']]]] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["region"] = region
        __props__["rules"] = rules
        return SecGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description for the security group. Changing this
        updates the `description` of an existing security group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A unique name for the security group. Changing this
        updates the `name` of an existing security group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region in which to obtain the V2 Compute client.
        A Compute client is needed to create a security group. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        security group.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> List['outputs.SecGroupRule']:
        """
        A rule describing how the security group operates. The
        rule object structure is documented below. Changing this updates the
        security group rules. As shown in the example above, multiple rule blocks
        may be used.
        """
        return pulumi.get(self, "rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

