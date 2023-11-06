# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RecordSetArgs', 'RecordSet']

@pulumi.input_type
class RecordSetArgs:
    def __init__(__self__, *,
                 records: pulumi.Input[Sequence[pulumi.Input[str]]],
                 zone_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a RecordSet resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: An array of DNS records.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the record set.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[str] description: A description of the  record set.
        :param pulumi.Input[bool] disable_status_check: Disable wait for recordset to reach ACTIVE
               status. This argumen is disabled by default. If it is set to true, the recordset
               will be considered as created/updated/deleted if OpenStack request returned success.
        :param pulumi.Input[str] name: The name of the record set. Note the `.` at the end of the name.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[str] project_id: The ID of the project DNS zone is created
               for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
               user role in target project)
        :param pulumi.Input[str] region: The region in which to obtain the V2 DNS client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[int] ttl: The time to live (TTL) of the record set.
        :param pulumi.Input[str] type: The type of record set. Examples: "A", "MX".
               Changing this creates a new DNS  record set.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new record set.
        """
        RecordSetArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            records=records,
            zone_id=zone_id,
            description=description,
            disable_status_check=disable_status_check,
            name=name,
            project_id=project_id,
            region=region,
            ttl=ttl,
            type=type,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             zone_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             disable_status_check: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             ttl: Optional[pulumi.Input[int]] = None,
             type: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if records is None:
            raise TypeError("Missing 'records' argument")
        if zone_id is None and 'zoneId' in kwargs:
            zone_id = kwargs['zoneId']
        if zone_id is None:
            raise TypeError("Missing 'zone_id' argument")
        if disable_status_check is None and 'disableStatusCheck' in kwargs:
            disable_status_check = kwargs['disableStatusCheck']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']

        _setter("records", records)
        _setter("zone_id", zone_id)
        if description is not None:
            _setter("description", description)
        if disable_status_check is not None:
            _setter("disable_status_check", disable_status_check)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)
        if region is not None:
            _setter("region", region)
        if ttl is not None:
            _setter("ttl", ttl)
        if type is not None:
            _setter("type", type)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter
    def records(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        An array of DNS records.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "records", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The ID of the zone in which to create the record set.
        Changing this creates a new DNS  record set.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the  record set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable wait for recordset to reach ACTIVE
        status. This argumen is disabled by default. If it is set to true, the recordset
        will be considered as created/updated/deleted if OpenStack request returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @disable_status_check.setter
    def disable_status_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_status_check", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the record set. Note the `.` at the end of the name.
        Changing this creates a new DNS  record set.
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
        The region in which to obtain the V2 DNS client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new DNS  record set.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The time to live (TTL) of the record set.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of record set. Examples: "A", "MX".
        Changing this creates a new DNS  record set.
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
        new record set.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _RecordSetState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RecordSet resources.
        :param pulumi.Input[str] description: A description of the  record set.
        :param pulumi.Input[bool] disable_status_check: Disable wait for recordset to reach ACTIVE
               status. This argumen is disabled by default. If it is set to true, the recordset
               will be considered as created/updated/deleted if OpenStack request returned success.
        :param pulumi.Input[str] name: The name of the record set. Note the `.` at the end of the name.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[str] project_id: The ID of the project DNS zone is created
               for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
               user role in target project)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: An array of DNS records.
        :param pulumi.Input[str] region: The region in which to obtain the V2 DNS client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[int] ttl: The time to live (TTL) of the record set.
        :param pulumi.Input[str] type: The type of record set. Examples: "A", "MX".
               Changing this creates a new DNS  record set.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new record set.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the record set.
               Changing this creates a new DNS  record set.
        """
        _RecordSetState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            disable_status_check=disable_status_check,
            name=name,
            project_id=project_id,
            records=records,
            region=region,
            ttl=ttl,
            type=type,
            value_specs=value_specs,
            zone_id=zone_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             disable_status_check: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             region: Optional[pulumi.Input[str]] = None,
             ttl: Optional[pulumi.Input[int]] = None,
             type: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             zone_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if disable_status_check is None and 'disableStatusCheck' in kwargs:
            disable_status_check = kwargs['disableStatusCheck']
        if project_id is None and 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']
        if zone_id is None and 'zoneId' in kwargs:
            zone_id = kwargs['zoneId']

        if description is not None:
            _setter("description", description)
        if disable_status_check is not None:
            _setter("disable_status_check", disable_status_check)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)
        if records is not None:
            _setter("records", records)
        if region is not None:
            _setter("region", region)
        if ttl is not None:
            _setter("ttl", ttl)
        if type is not None:
            _setter("type", type)
        if value_specs is not None:
            _setter("value_specs", value_specs)
        if zone_id is not None:
            _setter("zone_id", zone_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the  record set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable wait for recordset to reach ACTIVE
        status. This argumen is disabled by default. If it is set to true, the recordset
        will be considered as created/updated/deleted if OpenStack request returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @disable_status_check.setter
    def disable_status_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_status_check", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the record set. Note the `.` at the end of the name.
        Changing this creates a new DNS  record set.
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
    def records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of DNS records.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "records", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 DNS client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new DNS  record set.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The time to live (TTL) of the record set.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of record set. Examples: "A", "MX".
        Changing this creates a new DNS  record set.
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
        new record set.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the zone in which to create the record set.
        Changing this creates a new DNS  record set.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class RecordSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a DNS record set in the OpenStack DNS Service.

        ## Example Usage
        ### Automatically detect the correct network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        example_zone = openstack.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a zone",
            ttl=6000,
            type="PRIMARY")
        rs_example_com = openstack.dns.RecordSet("rsExampleCom",
            zone_id=example_zone.id,
            description="An example record set",
            ttl=3000,
            type="A",
            records=["10.0.0.1"])
        ```

        ## Import

        This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash.

        ```sh
         $ pulumi import openstack:dns/recordSet:RecordSet recordset_1 zone_id/recordset_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the  record set.
        :param pulumi.Input[bool] disable_status_check: Disable wait for recordset to reach ACTIVE
               status. This argumen is disabled by default. If it is set to true, the recordset
               will be considered as created/updated/deleted if OpenStack request returned success.
        :param pulumi.Input[str] name: The name of the record set. Note the `.` at the end of the name.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[str] project_id: The ID of the project DNS zone is created
               for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
               user role in target project)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: An array of DNS records.
        :param pulumi.Input[str] region: The region in which to obtain the V2 DNS client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[int] ttl: The time to live (TTL) of the record set.
        :param pulumi.Input[str] type: The type of record set. Examples: "A", "MX".
               Changing this creates a new DNS  record set.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new record set.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the record set.
               Changing this creates a new DNS  record set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RecordSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a DNS record set in the OpenStack DNS Service.

        ## Example Usage
        ### Automatically detect the correct network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        example_zone = openstack.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a zone",
            ttl=6000,
            type="PRIMARY")
        rs_example_com = openstack.dns.RecordSet("rsExampleCom",
            zone_id=example_zone.id,
            description="An example record set",
            ttl=3000,
            type="A",
            records=["10.0.0.1"])
        ```

        ## Import

        This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash.

        ```sh
         $ pulumi import openstack:dns/recordSet:RecordSet recordset_1 zone_id/recordset_id
        ```

        :param str resource_name: The name of the resource.
        :param RecordSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RecordSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RecordSetArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RecordSetArgs.__new__(RecordSetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disable_status_check"] = disable_status_check
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            if records is None and not opts.urn:
                raise TypeError("Missing required property 'records'")
            __props__.__dict__["records"] = records
            __props__.__dict__["region"] = region
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["type"] = type
            __props__.__dict__["value_specs"] = value_specs
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(RecordSet, __self__).__init__(
            'openstack:dns/recordSet:RecordSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_status_check: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'RecordSet':
        """
        Get an existing RecordSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the  record set.
        :param pulumi.Input[bool] disable_status_check: Disable wait for recordset to reach ACTIVE
               status. This argumen is disabled by default. If it is set to true, the recordset
               will be considered as created/updated/deleted if OpenStack request returned success.
        :param pulumi.Input[str] name: The name of the record set. Note the `.` at the end of the name.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[str] project_id: The ID of the project DNS zone is created
               for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
               user role in target project)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: An array of DNS records.
        :param pulumi.Input[str] region: The region in which to obtain the V2 DNS client.
               If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS  record set.
        :param pulumi.Input[int] ttl: The time to live (TTL) of the record set.
        :param pulumi.Input[str] type: The type of record set. Examples: "A", "MX".
               Changing this creates a new DNS  record set.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new record set.
        :param pulumi.Input[str] zone_id: The ID of the zone in which to create the record set.
               Changing this creates a new DNS  record set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RecordSetState.__new__(_RecordSetState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disable_status_check"] = disable_status_check
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["records"] = records
        __props__.__dict__["region"] = region
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["value_specs"] = value_specs
        __props__.__dict__["zone_id"] = zone_id
        return RecordSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the  record set.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable wait for recordset to reach ACTIVE
        status. This argumen is disabled by default. If it is set to true, the recordset
        will be considered as created/updated/deleted if OpenStack request returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the record set. Note the `.` at the end of the name.
        Changing this creates a new DNS  record set.
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
    def records(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of DNS records.
        """
        return pulumi.get(self, "records")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 DNS client.
        If omitted, the `region` argument of the provider is used.
        Changing this creates a new DNS  record set.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[int]:
        """
        The time to live (TTL) of the record set.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of record set. Examples: "A", "MX".
        Changing this creates a new DNS  record set.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new record set.
        """
        return pulumi.get(self, "value_specs")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The ID of the zone in which to create the record set.
        Changing this creates a new DNS  record set.
        """
        return pulumi.get(self, "zone_id")

