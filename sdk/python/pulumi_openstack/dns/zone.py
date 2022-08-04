# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ZoneArgs', 'Zone']

@pulumi.input_type
class ZoneArgs:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Zone resource.
        :param pulumi.Input[Mapping[str, Any]] attributes: Attributes for the DNS Service scheduler.
               Changing this creates a new zone.
        :param pulumi.Input[str] description: A description of the zone.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
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
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_status_check is not None:
            pulumi.set(__self__, "disable_status_check", disable_status_check)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if masters is not None:
            pulumi.set(__self__, "masters", masters)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Attributes for the DNS Service scheduler.
        Changing this creates a new zone.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the zone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable wait for zone to reach ACTIVE
        status. The check is enabled by default. If this argument is true, zone
        will be considered as created/updated if OpenStack request returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @disable_status_check.setter
    def disable_status_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_status_check", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email contact for the zone record.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def masters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of master DNS servers. For when `type` is
        `SECONDARY`.
        """
        return pulumi.get(self, "masters")

    @masters.setter
    def masters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "masters", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the zone. Note the `.` at the end of the name.
        Changing this creates a new DNS zone.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project DNS zone is created
        for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
        user role in target project)
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new DNS zone.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The time to live (TTL) of the zone.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of zone. Can either be `PRIMARY` or `SECONDARY`.
        Changing this creates a new zone.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new zone.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _ZoneState:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Zone resources.
        :param pulumi.Input[Mapping[str, Any]] attributes: Attributes for the DNS Service scheduler.
               Changing this creates a new zone.
        :param pulumi.Input[str] description: A description of the zone.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
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
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_status_check is not None:
            pulumi.set(__self__, "disable_status_check", disable_status_check)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if masters is not None:
            pulumi.set(__self__, "masters", masters)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Attributes for the DNS Service scheduler.
        Changing this creates a new zone.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the zone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable wait for zone to reach ACTIVE
        status. The check is enabled by default. If this argument is true, zone
        will be considered as created/updated if OpenStack request returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @disable_status_check.setter
    def disable_status_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_status_check", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The email contact for the zone record.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def masters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of master DNS servers. For when `type` is
        `SECONDARY`.
        """
        return pulumi.get(self, "masters")

    @masters.setter
    def masters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "masters", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the zone. Note the `.` at the end of the name.
        Changing this creates a new DNS zone.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project DNS zone is created
        for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
        user role in target project)
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new DNS zone.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The time to live (TTL) of the zone.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of zone. Can either be `PRIMARY` or `SECONDARY`.
        Changing this creates a new zone.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new zone.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


class Zone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
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

        ## Import

        This resource can be imported by specifying the zone ID with optional project ID

        ```sh
         $ pulumi import openstack:dns/zone:Zone zone_1 <zone_id>
        ```

        ```sh
         $ pulumi import openstack:dns/zone:Zone zone_1 <zone_id>:<project_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] attributes: Attributes for the DNS Service scheduler.
               Changing this creates a new zone.
        :param pulumi.Input[str] description: A description of the zone.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ZoneArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        ## Import

        This resource can be imported by specifying the zone ID with optional project ID

        ```sh
         $ pulumi import openstack:dns/zone:Zone zone_1 <zone_id>
        ```

        ```sh
         $ pulumi import openstack:dns/zone:Zone zone_1 <zone_id>:<project_id>
        ```

        :param str resource_name: The name of the resource.
        :param ZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 masters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneArgs.__new__(ZoneArgs)

            __props__.__dict__["attributes"] = attributes
            __props__.__dict__["description"] = description
            __props__.__dict__["disable_status_check"] = disable_status_check
            __props__.__dict__["email"] = email
            __props__.__dict__["masters"] = masters
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["type"] = type
            __props__.__dict__["value_specs"] = value_specs
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
            disable_status_check: Optional[pulumi.Input[bool]] = None,
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
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
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

        __props__ = _ZoneState.__new__(_ZoneState)

        __props__.__dict__["attributes"] = attributes
        __props__.__dict__["description"] = description
        __props__.__dict__["disable_status_check"] = disable_status_check
        __props__.__dict__["email"] = email
        __props__.__dict__["masters"] = masters
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["value_specs"] = value_specs
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
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable wait for zone to reach ACTIVE
        status. The check is enabled by default. If this argument is true, zone
        will be considered as created/updated if OpenStack request returned success.
        """
        return pulumi.get(self, "disable_status_check")

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

