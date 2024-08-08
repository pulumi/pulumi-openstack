# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecGroupArgs', 'SecGroup']

@pulumi.input_type
class SecGroupArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]]] = None):
        """
        The set of arguments for constructing a SecGroup resource.
        :param pulumi.Input[str] description: A description for the security group. Changing this
               updates the `description` of an existing security group.
        :param pulumi.Input[str] name: A unique name for the security group. Changing this
               updates the `name` of an existing security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a security group. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.
        """
        pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        A description for the security group. Changing this
        updates the `description` of an existing security group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the security group. Changing this
        updates the `name` of an existing security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        A Compute client is needed to create a security group. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        security group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]]]:
        """
        A rule describing how the security group operates. The
        rule object structure is documented below. Changing this updates the
        security group rules. As shown in the example above, multiple rule blocks
        may be used.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class _SecGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering SecGroup resources.
        :param pulumi.Input[str] description: A description for the security group. Changing this
               updates the `description` of an existing security group.
        :param pulumi.Input[str] name: A unique name for the security group. Changing this
               updates the `name` of an existing security group.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               A Compute client is needed to create a security group. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               security group.
        :param pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the security group. Changing this
        updates the `description` of an existing security group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the security group. Changing this
        updates the `name` of an existing security group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        A Compute client is needed to create a security group. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        security group.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]]]:
        """
        A rule describing how the security group operates. The
        rule object structure is documented below. Changing this updates the
        security group rules. As shown in the example above, multiple rule blocks
        may be used.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecGroupRuleArgs']]]]):
        pulumi.set(self, "rules", value)


class SecGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecGroupRuleArgs', 'SecGroupRuleArgsDict']]]]] = None,
                 __props__=None):
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

        secgroup1 = openstack.compute.SecGroup("secgroup_1",
            name="my_secgroup",
            description="my security group",
            rules=[
                {
                    "from_port": 22,
                    "to_port": 22,
                    "ip_protocol": "tcp",
                    "cidr": "0.0.0.0/0",
                },
                {
                    "from_port": 80,
                    "to_port": 80,
                    "ip_protocol": "tcp",
                    "cidr": "0.0.0.0/0",
                },
            ])
        ```

        ## Notes

        ### Referencing Security Groups

        When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this:

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_server = openstack.compute.Instance("test-server",
            name="tf-test",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="3",
            key_pair="my_key_pair_name",
            security_groups=[secgroup1["name"]])
        ```

        ## Import

        Security Groups can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:compute/secGroup:SecGroup my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
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
        :param pulumi.Input[Sequence[pulumi.Input[Union['SecGroupRuleArgs', 'SecGroupRuleArgsDict']]]] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        secgroup1 = openstack.compute.SecGroup("secgroup_1",
            name="my_secgroup",
            description="my security group",
            rules=[
                {
                    "from_port": 22,
                    "to_port": 22,
                    "ip_protocol": "tcp",
                    "cidr": "0.0.0.0/0",
                },
                {
                    "from_port": 80,
                    "to_port": 80,
                    "ip_protocol": "tcp",
                    "cidr": "0.0.0.0/0",
                },
            ])
        ```

        ## Notes

        ### Referencing Security Groups

        When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this:

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_server = openstack.compute.Instance("test-server",
            name="tf-test",
            image_id="ad091b52-742f-469e-8f3c-fd81cadf0743",
            flavor_id="3",
            key_pair="my_key_pair_name",
            security_groups=[secgroup1["name"]])
        ```

        ## Import

        Security Groups can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:compute/secGroup:SecGroup my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
        ```

        :param str resource_name: The name of the resource.
        :param SecGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecGroupRuleArgs', 'SecGroupRuleArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecGroupArgs.__new__(SecGroupArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["rules"] = rules
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
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecGroupRuleArgs', 'SecGroupRuleArgsDict']]]]] = None) -> 'SecGroup':
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
        :param pulumi.Input[Sequence[pulumi.Input[Union['SecGroupRuleArgs', 'SecGroupRuleArgsDict']]]] rules: A rule describing how the security group operates. The
               rule object structure is documented below. Changing this updates the
               security group rules. As shown in the example above, multiple rule blocks
               may be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecGroupState.__new__(_SecGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["rules"] = rules
        return SecGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        A description for the security group. Changing this
        updates the `description` of an existing security group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the security group. Changing this
        updates the `name` of an existing security group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        A Compute client is needed to create a security group. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        security group.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.SecGroupRule']]:
        """
        A rule describing how the security group operates. The
        rule object structure is documented below. Changing this updates the
        security group rules. As shown in the example above, multiple rule blocks
        may be used.
        """
        return pulumi.get(self, "rules")

