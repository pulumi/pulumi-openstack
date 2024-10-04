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

__all__ = ['PoolArgs', 'Pool']

@pulumi.input_type
class PoolArgs:
    def __init__(__self__, *,
                 lb_method: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 alpn_protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_tls_container_ref: Optional[pulumi.Input[str]] = None,
                 crl_container_ref: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistence: Optional[pulumi.Input['PoolPersistenceArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tls_ciphers: Optional[pulumi.Input[str]] = None,
                 tls_container_ref: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Pool resource.
        """
        pulumi.set(__self__, "lb_method", lb_method)
        pulumi.set(__self__, "protocol", protocol)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if alpn_protocols is not None:
            pulumi.set(__self__, "alpn_protocols", alpn_protocols)
        if ca_tls_container_ref is not None:
            pulumi.set(__self__, "ca_tls_container_ref", ca_tls_container_ref)
        if crl_container_ref is not None:
            pulumi.set(__self__, "crl_container_ref", crl_container_ref)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if loadbalancer_id is not None:
            pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if persistence is not None:
            pulumi.set(__self__, "persistence", persistence)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if tls_ciphers is not None:
            pulumi.set(__self__, "tls_ciphers", tls_ciphers)
        if tls_container_ref is not None:
            pulumi.set(__self__, "tls_container_ref", tls_container_ref)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)
        if tls_versions is not None:
            pulumi.set(__self__, "tls_versions", tls_versions)

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> pulumi.Input[str]:
        return pulumi.get(self, "lb_method")

    @lb_method.setter
    def lb_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "lb_method", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="alpnProtocols")
    def alpn_protocols(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "alpn_protocols")

    @alpn_protocols.setter
    def alpn_protocols(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alpn_protocols", value)

    @property
    @pulumi.getter(name="caTlsContainerRef")
    def ca_tls_container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ca_tls_container_ref")

    @ca_tls_container_ref.setter
    def ca_tls_container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_tls_container_ref", value)

    @property
    @pulumi.getter(name="crlContainerRef")
    def crl_container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "crl_container_ref")

    @crl_container_ref.setter
    def crl_container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crl_container_ref", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def persistence(self) -> Optional[pulumi.Input['PoolPersistenceArgs']]:
        return pulumi.get(self, "persistence")

    @persistence.setter
    def persistence(self, value: Optional[pulumi.Input['PoolPersistenceArgs']]):
        pulumi.set(self, "persistence", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="tlsCiphers")
    def tls_ciphers(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_ciphers")

    @tls_ciphers.setter
    def tls_ciphers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_ciphers", value)

    @property
    @pulumi.getter(name="tlsContainerRef")
    def tls_container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_container_ref")

    @tls_container_ref.setter
    def tls_container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_container_ref", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)

    @property
    @pulumi.getter(name="tlsVersions")
    def tls_versions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tls_versions")

    @tls_versions.setter
    def tls_versions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tls_versions", value)


@pulumi.input_type
class _PoolState:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 alpn_protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_tls_container_ref: Optional[pulumi.Input[str]] = None,
                 crl_container_ref: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lb_method: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistence: Optional[pulumi.Input['PoolPersistenceArgs']] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tls_ciphers: Optional[pulumi.Input[str]] = None,
                 tls_container_ref: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Pool resources.
        """
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if alpn_protocols is not None:
            pulumi.set(__self__, "alpn_protocols", alpn_protocols)
        if ca_tls_container_ref is not None:
            pulumi.set(__self__, "ca_tls_container_ref", ca_tls_container_ref)
        if crl_container_ref is not None:
            pulumi.set(__self__, "crl_container_ref", crl_container_ref)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if lb_method is not None:
            pulumi.set(__self__, "lb_method", lb_method)
        if listener_id is not None:
            pulumi.set(__self__, "listener_id", listener_id)
        if loadbalancer_id is not None:
            pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if persistence is not None:
            pulumi.set(__self__, "persistence", persistence)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if tls_ciphers is not None:
            pulumi.set(__self__, "tls_ciphers", tls_ciphers)
        if tls_container_ref is not None:
            pulumi.set(__self__, "tls_container_ref", tls_container_ref)
        if tls_enabled is not None:
            pulumi.set(__self__, "tls_enabled", tls_enabled)
        if tls_versions is not None:
            pulumi.set(__self__, "tls_versions", tls_versions)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="alpnProtocols")
    def alpn_protocols(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "alpn_protocols")

    @alpn_protocols.setter
    def alpn_protocols(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alpn_protocols", value)

    @property
    @pulumi.getter(name="caTlsContainerRef")
    def ca_tls_container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ca_tls_container_ref")

    @ca_tls_container_ref.setter
    def ca_tls_container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_tls_container_ref", value)

    @property
    @pulumi.getter(name="crlContainerRef")
    def crl_container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "crl_container_ref")

    @crl_container_ref.setter
    def crl_container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crl_container_ref", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "lb_method")

    @lb_method.setter
    def lb_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lb_method", value)

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "listener_id")

    @listener_id.setter
    def listener_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listener_id", value)

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "loadbalancer_id")

    @loadbalancer_id.setter
    def loadbalancer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "loadbalancer_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def persistence(self) -> Optional[pulumi.Input['PoolPersistenceArgs']]:
        return pulumi.get(self, "persistence")

    @persistence.setter
    def persistence(self, value: Optional[pulumi.Input['PoolPersistenceArgs']]):
        pulumi.set(self, "persistence", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="tlsCiphers")
    def tls_ciphers(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_ciphers")

    @tls_ciphers.setter
    def tls_ciphers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_ciphers", value)

    @property
    @pulumi.getter(name="tlsContainerRef")
    def tls_container_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_container_ref")

    @tls_container_ref.setter
    def tls_container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_container_ref", value)

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "tls_enabled")

    @tls_enabled.setter
    def tls_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_enabled", value)

    @property
    @pulumi.getter(name="tlsVersions")
    def tls_versions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "tls_versions")

    @tls_versions.setter
    def tls_versions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tls_versions", value)


class Pool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 alpn_protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_tls_container_ref: Optional[pulumi.Input[str]] = None,
                 crl_container_ref: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lb_method: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistence: Optional[pulumi.Input[Union['PoolPersistenceArgs', 'PoolPersistenceArgsDict']]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tls_ciphers: Optional[pulumi.Input[str]] = None,
                 tls_container_ref: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a Pool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Pool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 alpn_protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ca_tls_container_ref: Optional[pulumi.Input[str]] = None,
                 crl_container_ref: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 lb_method: Optional[pulumi.Input[str]] = None,
                 listener_id: Optional[pulumi.Input[str]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 persistence: Optional[pulumi.Input[Union['PoolPersistenceArgs', 'PoolPersistenceArgsDict']]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 tls_ciphers: Optional[pulumi.Input[str]] = None,
                 tls_container_ref: Optional[pulumi.Input[str]] = None,
                 tls_enabled: Optional[pulumi.Input[bool]] = None,
                 tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PoolArgs.__new__(PoolArgs)

            __props__.__dict__["admin_state_up"] = admin_state_up
            __props__.__dict__["alpn_protocols"] = alpn_protocols
            __props__.__dict__["ca_tls_container_ref"] = ca_tls_container_ref
            __props__.__dict__["crl_container_ref"] = crl_container_ref
            __props__.__dict__["description"] = description
            if lb_method is None and not opts.urn:
                raise TypeError("Missing required property 'lb_method'")
            __props__.__dict__["lb_method"] = lb_method
            __props__.__dict__["listener_id"] = listener_id
            __props__.__dict__["loadbalancer_id"] = loadbalancer_id
            __props__.__dict__["name"] = name
            __props__.__dict__["persistence"] = persistence
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["region"] = region
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["tls_ciphers"] = tls_ciphers
            __props__.__dict__["tls_container_ref"] = tls_container_ref
            __props__.__dict__["tls_enabled"] = tls_enabled
            __props__.__dict__["tls_versions"] = tls_versions
        super(Pool, __self__).__init__(
            'openstack:loadbalancer/pool:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            alpn_protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ca_tls_container_ref: Optional[pulumi.Input[str]] = None,
            crl_container_ref: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lb_method: Optional[pulumi.Input[str]] = None,
            listener_id: Optional[pulumi.Input[str]] = None,
            loadbalancer_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            persistence: Optional[pulumi.Input[Union['PoolPersistenceArgs', 'PoolPersistenceArgsDict']]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            tls_ciphers: Optional[pulumi.Input[str]] = None,
            tls_container_ref: Optional[pulumi.Input[str]] = None,
            tls_enabled: Optional[pulumi.Input[bool]] = None,
            tls_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Pool':
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PoolState.__new__(_PoolState)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["alpn_protocols"] = alpn_protocols
        __props__.__dict__["ca_tls_container_ref"] = ca_tls_container_ref
        __props__.__dict__["crl_container_ref"] = crl_container_ref
        __props__.__dict__["description"] = description
        __props__.__dict__["lb_method"] = lb_method
        __props__.__dict__["listener_id"] = listener_id
        __props__.__dict__["loadbalancer_id"] = loadbalancer_id
        __props__.__dict__["name"] = name
        __props__.__dict__["persistence"] = persistence
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["region"] = region
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["tls_ciphers"] = tls_ciphers
        __props__.__dict__["tls_container_ref"] = tls_container_ref
        __props__.__dict__["tls_enabled"] = tls_enabled
        __props__.__dict__["tls_versions"] = tls_versions
        return Pool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="alpnProtocols")
    def alpn_protocols(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "alpn_protocols")

    @property
    @pulumi.getter(name="caTlsContainerRef")
    def ca_tls_container_ref(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ca_tls_container_ref")

    @property
    @pulumi.getter(name="crlContainerRef")
    def crl_container_ref(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "crl_container_ref")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lbMethod")
    def lb_method(self) -> pulumi.Output[str]:
        return pulumi.get(self, "lb_method")

    @property
    @pulumi.getter(name="listenerId")
    def listener_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "listener_id")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def persistence(self) -> pulumi.Output[Optional['outputs.PoolPersistence']]:
        return pulumi.get(self, "persistence")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="tlsCiphers")
    def tls_ciphers(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tls_ciphers")

    @property
    @pulumi.getter(name="tlsContainerRef")
    def tls_container_ref(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "tls_container_ref")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "tls_enabled")

    @property
    @pulumi.getter(name="tlsVersions")
    def tls_versions(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "tls_versions")

