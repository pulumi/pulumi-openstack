# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Zone']


class Zone(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a DNS zone in the OpenStack DNS Service.

        ## Example Usage
        ### Automatically detect the correct network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        example_com = openstack.dns.Zone("example.com",
            description="An example zone",
            email="jdoe@example.com",
            ttl=3000,
            type="PRIMARY")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] attributes: Attributes for the DNS Service scheduler.
               Changing this creates a new zone.
        :param pulumi.Input[str] description: A description of the zone.
        :param pulumi.Input[str] email: The email contact for the zone record.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] masters: An array of master DNS servers. For when `type` is
               `SECONDARY`.
        :param pulumi.Input[str] name: The name of the zone. Note the `.` at the end of the name.
               Changing this creates a new DNS zone.
        :param pulumi.Input[str] project_id: The ID of the project DNS zone is created
               for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
               user role in target project)
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[int] ttl: The time to live (TTL) of the zone.
        :param pulumi.Input[str] type: The type of zone. Can either be `PRIMARY` or `SECONDARY`.
               Changing this creates a new zone.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new zone.
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

            __props__['attributes'] = attributes
            __props__['description'] = description
            __props__['email'] = email
            __props__['masters'] = masters
            __props__['name'] = name
            __props__['project_id'] = project_id
            __props__['region'] = region
            __props__['ttl'] = ttl
            __props__['type'] = type
            __props__['value_specs'] = value_specs
        super(Zone, __self__).__init__(
            'openstack:dns/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Zone':
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] attributes: Attributes for the DNS Service scheduler.
               Changing this creates a new zone.
        :param pulumi.Input[str] description: A description of the zone.
        :param pulumi.Input[str] email: The email contact for the zone record.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] masters: An array of master DNS servers. For when `type` is
               `SECONDARY`.
        :param pulumi.Input[str] name: The name of the zone. Note the `.` at the end of the name.
               Changing this creates a new DNS zone.
        :param pulumi.Input[str] project_id: The ID of the project DNS zone is created
               for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
               user role in target project)
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[int] ttl: The time to live (TTL) of the zone.
        :param pulumi.Input[str] type: The type of zone. Can either be `PRIMARY` or `SECONDARY`.
               Changing this creates a new zone.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["description"] = description
        __props__["email"] = email
        __props__["masters"] = masters
        __props__["name"] = name
        __props__["project_id"] = project_id
        __props__["region"] = region
        __props__["ttl"] = ttl
        __props__["type"] = type
        __props__["value_specs"] = value_specs
        return Zone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Attributes for the DNS Service scheduler.
        Changing this creates a new zone.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the zone.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[str]]:
        """
        The email contact for the zone record.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def masters(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of master DNS servers. For when `type` is
        `SECONDARY`.
        """
        return pulumi.get(self, "masters")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the zone. Note the `.` at the end of the name.
        Changing this creates a new DNS zone.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project DNS zone is created
        for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
        user role in target project)
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new DNS zone.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[int]:
        """
        The time to live (TTL) of the zone.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of zone. Can either be `PRIMARY` or `SECONDARY`.
        Changing this creates a new zone.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new zone.
        """
        return pulumi.get(self, "value_specs")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

