# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransferAcceptArgs', 'TransferAccept']

@pulumi.input_type
class TransferAcceptArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 zone_transfer_request_id: pulumi.Input[str],
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a TransferAccept resource.
        :param pulumi.Input[str] key: The transfer key.
        :param pulumi.Input[str] zone_transfer_request_id: The ID of the zone transfer request.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack accept returned success.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer accept.
        """
        TransferAcceptArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            key=key,
            zone_transfer_request_id=zone_transfer_request_id,
            disable_status_check=disable_status_check,
            region=region,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             key: Optional[pulumi.Input[str]] = None,
             zone_transfer_request_id: Optional[pulumi.Input[str]] = None,
             disable_status_check: Optional[pulumi.Input[bool]] = None,
             region: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if key is None:
            raise TypeError("Missing 'key' argument")
        if zone_transfer_request_id is None and 'zoneTransferRequestId' in kwargs:
            zone_transfer_request_id = kwargs['zoneTransferRequestId']
        if zone_transfer_request_id is None:
            raise TypeError("Missing 'zone_transfer_request_id' argument")
        if disable_status_check is None and 'disableStatusCheck' in kwargs:
            disable_status_check = kwargs['disableStatusCheck']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']

        _setter("key", key)
        _setter("zone_transfer_request_id", zone_transfer_request_id)
        if disable_status_check is not None:
            _setter("disable_status_check", disable_status_check)
        if region is not None:
            _setter("region", region)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The transfer key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="zoneTransferRequestId")
    def zone_transfer_request_id(self) -> pulumi.Input[str]:
        """
        The ID of the zone transfer request.
        """
        return pulumi.get(self, "zone_transfer_request_id")

    @zone_transfer_request_id.setter
    def zone_transfer_request_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_transfer_request_id", value)

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable wait for zone to reach ACTIVE
        status. The check is enabled by default. If this argument is true, zone
        will be considered as created/updated if OpenStack accept returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @disable_status_check.setter
    def disable_status_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_status_check", value)

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
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new transfer accept.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _TransferAcceptState:
    def __init__(__self__, *,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_transfer_request_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransferAccept resources.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack accept returned success.
        :param pulumi.Input[str] key: The transfer key.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer accept.
        :param pulumi.Input[str] zone_transfer_request_id: The ID of the zone transfer request.
        """
        _TransferAcceptState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            disable_status_check=disable_status_check,
            key=key,
            region=region,
            value_specs=value_specs,
            zone_transfer_request_id=zone_transfer_request_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             disable_status_check: Optional[pulumi.Input[bool]] = None,
             key: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             zone_transfer_request_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if disable_status_check is None and 'disableStatusCheck' in kwargs:
            disable_status_check = kwargs['disableStatusCheck']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']
        if zone_transfer_request_id is None and 'zoneTransferRequestId' in kwargs:
            zone_transfer_request_id = kwargs['zoneTransferRequestId']

        if disable_status_check is not None:
            _setter("disable_status_check", disable_status_check)
        if key is not None:
            _setter("key", key)
        if region is not None:
            _setter("region", region)
        if value_specs is not None:
            _setter("value_specs", value_specs)
        if zone_transfer_request_id is not None:
            _setter("zone_transfer_request_id", zone_transfer_request_id)

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable wait for zone to reach ACTIVE
        status. The check is enabled by default. If this argument is true, zone
        will be considered as created/updated if OpenStack accept returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @disable_status_check.setter
    def disable_status_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_status_check", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The transfer key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

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
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new transfer accept.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)

    @property
    @pulumi.getter(name="zoneTransferRequestId")
    def zone_transfer_request_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the zone transfer request.
        """
        return pulumi.get(self, "zone_transfer_request_id")

    @zone_transfer_request_id.setter
    def zone_transfer_request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_transfer_request_id", value)


class TransferAccept(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_transfer_request_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a DNS zone transfer accept in the OpenStack DNS Service.

        ## Example Usage
        ### Automatically detect the correct network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        example_zone = openstack.dns.Zone("exampleZone",
            email="jdoe@example.com",
            description="An example zone",
            ttl=3000,
            type="PRIMARY")
        request1 = openstack.dns.TransferRequest("request1",
            zone_id=example_zone.id,
            description="a transfer accept")
        accept1 = openstack.dns.TransferAccept("accept1",
            zone_transfer_request_id=request1.id,
            key=request1.key)
        ```

        ## Import

        This resource can be imported by specifying the transferAccept ID:

        ```sh
         $ pulumi import openstack:dns/transferAccept:TransferAccept accept_1 accept_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack accept returned success.
        :param pulumi.Input[str] key: The transfer key.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer accept.
        :param pulumi.Input[str] zone_transfer_request_id: The ID of the zone transfer request.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransferAcceptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a DNS zone transfer accept in the OpenStack DNS Service.

        ## Example Usage
        ### Automatically detect the correct network

        ```python
        import pulumi
        import pulumi_openstack as openstack

        example_zone = openstack.dns.Zone("exampleZone",
            email="jdoe@example.com",
            description="An example zone",
            ttl=3000,
            type="PRIMARY")
        request1 = openstack.dns.TransferRequest("request1",
            zone_id=example_zone.id,
            description="a transfer accept")
        accept1 = openstack.dns.TransferAccept("accept1",
            zone_transfer_request_id=request1.id,
            key=request1.key)
        ```

        ## Import

        This resource can be imported by specifying the transferAccept ID:

        ```sh
         $ pulumi import openstack:dns/transferAccept:TransferAccept accept_1 accept_id
        ```

        :param str resource_name: The name of the resource.
        :param TransferAcceptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransferAcceptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TransferAcceptArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_transfer_request_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransferAcceptArgs.__new__(TransferAcceptArgs)

            __props__.__dict__["disable_status_check"] = disable_status_check
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["region"] = region
            __props__.__dict__["value_specs"] = value_specs
            if zone_transfer_request_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_transfer_request_id'")
            __props__.__dict__["zone_transfer_request_id"] = zone_transfer_request_id
        super(TransferAccept, __self__).__init__(
            'openstack:dns/transferAccept:TransferAccept',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_status_check: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            zone_transfer_request_id: Optional[pulumi.Input[str]] = None) -> 'TransferAccept':
        """
        Get an existing TransferAccept resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack accept returned success.
        :param pulumi.Input[str] key: The transfer key.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer accept.
        :param pulumi.Input[str] zone_transfer_request_id: The ID of the zone transfer request.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransferAcceptState.__new__(_TransferAcceptState)

        __props__.__dict__["disable_status_check"] = disable_status_check
        __props__.__dict__["key"] = key
        __props__.__dict__["region"] = region
        __props__.__dict__["value_specs"] = value_specs
        __props__.__dict__["zone_transfer_request_id"] = zone_transfer_request_id
        return TransferAccept(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableStatusCheck")
    def disable_status_check(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable wait for zone to reach ACTIVE
        status. The check is enabled by default. If this argument is true, zone
        will be considered as created/updated if OpenStack accept returned success.
        """
        return pulumi.get(self, "disable_status_check")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The transfer key.
        """
        return pulumi.get(self, "key")

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
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new transfer accept.
        """
        return pulumi.get(self, "value_specs")

    @property
    @pulumi.getter(name="zoneTransferRequestId")
    def zone_transfer_request_id(self) -> pulumi.Output[str]:
        """
        The ID of the zone transfer request.
        """
        return pulumi.get(self, "zone_transfer_request_id")

