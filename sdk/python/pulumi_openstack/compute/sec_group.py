# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class SecGroup(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description for the security group. Changing this
    updates the `description` of an existing security group.
    """
    name: pulumi.Output[str]
    """
    A unique name for the security group. Changing this
    updates the `name` of an existing security group.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V2 Compute client.
    A Compute client is needed to create a security group. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    security group.
    """
    rules: pulumi.Output[list]
    """
    A rule describing how the security group operates. The
    rule object structure is documented below. Changing this updates the
    security group rules. As shown in the example above, multiple rule blocks
    may be used.

      * `cidr` (`str`) - Required if `from_group_id` or `self` is empty. The IP range
        that will be the source of network traffic to the security group. Use 0.0.0.0/0
        to allow all IP addresses. Changing this creates a new security group rule. Cannot
        be combined with `from_group_id` or `self`.
      * `fromGroupId` (`str`) - Required if `cidr` or `self` is empty. The ID of a
        group from which to forward traffic to the parent group. Changing this creates a
        new security group rule. Cannot be combined with `cidr` or `self`.
      * `fromPort` (`float`) - An integer representing the lower bound of the port
        range to open. Changing this creates a new security group rule.
      * `id` (`str`)
      * `ipProtocol` (`str`) - The protocol type that will be allowed. Changing
        this creates a new security group rule.
      * `self` (`bool`) - Required if `cidr` and `from_group_id` is empty. If true,
        the security group itself will be added as a source to this ingress rule. Cannot
        be combined with `cidr` or `from_group_id`.
      * `toPort` (`float`) - An integer representing the upper bound of the port
        range to open. Changing this creates a new security group rule.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, region=None, rules=None, __props__=None, __name__=None, __opts__=None):
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
                {
                    "cidr": "0.0.0.0/0",
                    "fromPort": 22,
                    "ipProtocol": "tcp",
                    "toPort": 22,
                },
                {
                    "cidr": "0.0.0.0/0",
                    "fromPort": 80,
                    "ipProtocol": "tcp",
                    "toPort": 80,
                },
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
        :param pulumi.Input[list] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.

        The **rules** object supports the following:

          * `cidr` (`pulumi.Input[str]`) - Required if `from_group_id` or `self` is empty. The IP range
            that will be the source of network traffic to the security group. Use 0.0.0.0/0
            to allow all IP addresses. Changing this creates a new security group rule. Cannot
            be combined with `from_group_id` or `self`.
          * `fromGroupId` (`pulumi.Input[str]`) - Required if `cidr` or `self` is empty. The ID of a
            group from which to forward traffic to the parent group. Changing this creates a
            new security group rule. Cannot be combined with `cidr` or `self`.
          * `fromPort` (`pulumi.Input[float]`) - An integer representing the lower bound of the port
            range to open. Changing this creates a new security group rule.
          * `id` (`pulumi.Input[str]`)
          * `ipProtocol` (`pulumi.Input[str]`) - The protocol type that will be allowed. Changing
            this creates a new security group rule.
          * `self` (`pulumi.Input[bool]`) - Required if `cidr` and `from_group_id` is empty. If true,
            the security group itself will be added as a source to this ingress rule. Cannot
            be combined with `cidr` or `from_group_id`.
          * `toPort` (`pulumi.Input[float]`) - An integer representing the upper bound of the port
            range to open. Changing this creates a new security group rule.
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
    def get(resource_name, id, opts=None, description=None, name=None, region=None, rules=None):
        """
        Get an existing SecGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the security group. Changing this
               updates the `description` of an existing security group.
        :param pulumi.Input[str] name: A unique name for the security group. Changing this
               updates the `name` of an existing security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a security group. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[list] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.

        The **rules** object supports the following:

          * `cidr` (`pulumi.Input[str]`) - Required if `from_group_id` or `self` is empty. The IP range
            that will be the source of network traffic to the security group. Use 0.0.0.0/0
            to allow all IP addresses. Changing this creates a new security group rule. Cannot
            be combined with `from_group_id` or `self`.
          * `fromGroupId` (`pulumi.Input[str]`) - Required if `cidr` or `self` is empty. The ID of a
            group from which to forward traffic to the parent group. Changing this creates a
            new security group rule. Cannot be combined with `cidr` or `self`.
          * `fromPort` (`pulumi.Input[float]`) - An integer representing the lower bound of the port
            range to open. Changing this creates a new security group rule.
          * `id` (`pulumi.Input[str]`)
          * `ipProtocol` (`pulumi.Input[str]`) - The protocol type that will be allowed. Changing
            this creates a new security group rule.
          * `self` (`pulumi.Input[bool]`) - Required if `cidr` and `from_group_id` is empty. If true,
            the security group itself will be added as a source to this ingress rule. Cannot
            be combined with `cidr` or `from_group_id`.
          * `toPort` (`pulumi.Input[float]`) - An integer representing the upper bound of the port
            range to open. Changing this creates a new security group rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["region"] = region
        __props__["rules"] = rules
        return SecGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
