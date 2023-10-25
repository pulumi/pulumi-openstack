# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransferRequestArgs', 'TransferRequest']

@pulumi.input_type
class TransferRequestArgs:
    def __init__(__self__, *,
                 zone_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target_project_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a TransferRequest resource.
        :param pulumi.Input[str] zone_id: The ID of the zone for which to create the transfer
               request.
        :param pulumi.Input[str] description: A description of the zone tranfer request.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[str] target_project_id: The target Project ID to transfer to.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer request.
        """
        TransferRequestArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            zone_id=zone_id,
            description=description,
            disable_status_check=disable_status_check,
            key=key,
            region=region,
            target_project_id=target_project_id,
            value_specs=value_specs,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             zone_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             disable_status_check: Optional[pulumi.Input[bool]] = None,
             key: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             target_project_id: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if zone_id is None and 'zoneId' in kwargs:
            zone_id = kwargs['zoneId']
        if zone_id is None:
            raise TypeError("Missing 'zone_id' argument")
        if disable_status_check is None and 'disableStatusCheck' in kwargs:
            disable_status_check = kwargs['disableStatusCheck']
        if target_project_id is None and 'targetProjectId' in kwargs:
            target_project_id = kwargs['targetProjectId']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']

        _setter("zone_id", zone_id)
        if description is not None:
            _setter("description", description)
        if disable_status_check is not None:
            _setter("disable_status_check", disable_status_check)
        if key is not None:
            _setter("key", key)
        if region is not None:
            _setter("region", region)
        if target_project_id is not None:
            _setter("target_project_id", target_project_id)
        if value_specs is not None:
            _setter("value_specs", value_specs)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The ID of the zone for which to create the transfer
        request.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the zone tranfer request.
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
    def key(self) -> Optional[pulumi.Input[str]]:
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
    @pulumi.getter(name="targetProjectId")
    def target_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The target Project ID to transfer to.
        """
        return pulumi.get(self, "target_project_id")

    @target_project_id.setter
    def target_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_project_id", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new transfer request.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _TransferRequestState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target_project_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransferRequest resources.
        :param pulumi.Input[str] description: A description of the zone tranfer request.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[str] target_project_id: The target Project ID to transfer to.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer request.
        :param pulumi.Input[str] zone_id: The ID of the zone for which to create the transfer
               request.
        """
        _TransferRequestState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            disable_status_check=disable_status_check,
            key=key,
            region=region,
            target_project_id=target_project_id,
            value_specs=value_specs,
            zone_id=zone_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             disable_status_check: Optional[pulumi.Input[bool]] = None,
             key: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             target_project_id: Optional[pulumi.Input[str]] = None,
             value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             zone_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if disable_status_check is None and 'disableStatusCheck' in kwargs:
            disable_status_check = kwargs['disableStatusCheck']
        if target_project_id is None and 'targetProjectId' in kwargs:
            target_project_id = kwargs['targetProjectId']
        if value_specs is None and 'valueSpecs' in kwargs:
            value_specs = kwargs['valueSpecs']
        if zone_id is None and 'zoneId' in kwargs:
            zone_id = kwargs['zoneId']

        if description is not None:
            _setter("description", description)
        if disable_status_check is not None:
            _setter("disable_status_check", disable_status_check)
        if key is not None:
            _setter("key", key)
        if region is not None:
            _setter("region", region)
        if target_project_id is not None:
            _setter("target_project_id", target_project_id)
        if value_specs is not None:
            _setter("value_specs", value_specs)
        if zone_id is not None:
            _setter("zone_id", zone_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the zone tranfer request.
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
    def key(self) -> Optional[pulumi.Input[str]]:
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
    @pulumi.getter(name="targetProjectId")
    def target_project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The target Project ID to transfer to.
        """
        return pulumi.get(self, "target_project_id")

    @target_project_id.setter
    def target_project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_project_id", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new transfer request.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the zone for which to create the transfer
        request.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class TransferRequest(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target_project_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a DNS zone transfer request in the OpenStack DNS Service.

        ## Example Usage

        ## Import

        This resource can be imported by specifying the transferRequest ID:

        ```sh
         $ pulumi import openstack:dns/transferRequest:TransferRequest request_1 request_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the zone tranfer request.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[str] target_project_id: The target Project ID to transfer to.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer request.
        :param pulumi.Input[str] zone_id: The ID of the zone for which to create the transfer
               request.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransferRequestArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a DNS zone transfer request in the OpenStack DNS Service.

        ## Example Usage

        ## Import

        This resource can be imported by specifying the transferRequest ID:

        ```sh
         $ pulumi import openstack:dns/transferRequest:TransferRequest request_1 request_id
        ```

        :param str resource_name: The name of the resource.
        :param TransferRequestArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransferRequestArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TransferRequestArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_status_check: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 target_project_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransferRequestArgs.__new__(TransferRequestArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disable_status_check"] = disable_status_check
            __props__.__dict__["key"] = key
            __props__.__dict__["region"] = region
            __props__.__dict__["target_project_id"] = target_project_id
            __props__.__dict__["value_specs"] = value_specs
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
        super(TransferRequest, __self__).__init__(
            'openstack:dns/transferRequest:TransferRequest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_status_check: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            target_project_id: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'TransferRequest':
        """
        Get an existing TransferRequest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the zone tranfer request.
        :param pulumi.Input[bool] disable_status_check: Disable wait for zone to reach ACTIVE
               status. The check is enabled by default. If this argument is true, zone
               will be considered as created/updated if OpenStack request returned success.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new DNS zone.
        :param pulumi.Input[str] target_project_id: The target Project ID to transfer to.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options. Changing this creates a
               new transfer request.
        :param pulumi.Input[str] zone_id: The ID of the zone for which to create the transfer
               request.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransferRequestState.__new__(_TransferRequestState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disable_status_check"] = disable_status_check
        __props__.__dict__["key"] = key
        __props__.__dict__["region"] = region
        __props__.__dict__["target_project_id"] = target_project_id
        __props__.__dict__["value_specs"] = value_specs
        __props__.__dict__["zone_id"] = zone_id
        return TransferRequest(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the zone tranfer request.
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
    def key(self) -> pulumi.Output[str]:
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
    @pulumi.getter(name="targetProjectId")
    def target_project_id(self) -> pulumi.Output[str]:
        """
        The target Project ID to transfer to.
        """
        return pulumi.get(self, "target_project_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options. Changing this creates a
        new transfer request.
        """
        return pulumi.get(self, "value_specs")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The ID of the zone for which to create the transfer
        request.
        """
        return pulumi.get(self, "zone_id")

