# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Listener']


class Listener(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 allowed_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 default_pool_id: Optional[pulumi.Input[str]] = None,
                 default_tls_container_ref: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 insert_headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 loadbalancer_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_port: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sni_container_refs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 timeout_client_data: Optional[pulumi.Input[int]] = None,
                 timeout_member_connect: Optional[pulumi.Input[int]] = None,
                 timeout_member_data: Optional[pulumi.Input[int]] = None,
                 timeout_tcp_inspect: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a V2 listener resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        listener1 = openstack.loadbalancer.Listener("listener1",
            insert_headers={
                "X-Forwarded-For": "true",
            },
            loadbalancer_id="d9415786-5f1a-428b-b35f-2f1523e146d2",
            protocol="HTTP",
            protocol_port=8080)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the Listener.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_cidrs: A list of CIDR blocks that are permitted to connect to this listener, denying
               all other source addresses. If not present, defaults to allow all.
        :param pulumi.Input[int] connection_limit: The maximum number of connections allowed
               for the Listener.
        :param pulumi.Input[str] default_pool_id: The ID of the default pool with which the
               Listener is associated.
        :param pulumi.Input[str] default_tls_container_ref: A reference to a Barbican Secrets
               container which stores TLS information. This is required if the protocol
               is `TERMINATED_HTTPS`. See
               [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
               for more information.
        :param pulumi.Input[str] description: Human-readable description for the Listener.
        :param pulumi.Input[Mapping[str, Any]] insert_headers: The list of key value pairs representing headers to insert
               into the request before it is sent to the backend members. Changing this updates the headers of the
               existing listener.
        :param pulumi.Input[str] loadbalancer_id: The load balancer on which to provision this
               Listener. Changing this creates a new Listener.
        :param pulumi.Input[str] name: Human-readable name for the Listener. Does not have
               to be unique.
        :param pulumi.Input[str] protocol: The protocol - can either be TCP, HTTP, HTTPS,
               TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
               new Listener.
        :param pulumi.Input[int] protocol_port: The port on which to listen for client traffic.
               Changing this creates a new Listener.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               Listener.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sni_container_refs: A list of references to Barbican Secrets
               containers which store SNI information. See
               [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
               for more information.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the Listener.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new Listener.
        :param pulumi.Input[int] timeout_client_data: The client inactivity timeout in milliseconds.
        :param pulumi.Input[int] timeout_member_connect: The member connection timeout in milliseconds.
        :param pulumi.Input[int] timeout_member_data: The member inactivity timeout in milliseconds.
        :param pulumi.Input[int] timeout_tcp_inspect: The time in milliseconds, to wait for additional
               TCP packets for content inspection.
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

            __props__['admin_state_up'] = admin_state_up
            __props__['allowed_cidrs'] = allowed_cidrs
            __props__['connection_limit'] = connection_limit
            __props__['default_pool_id'] = default_pool_id
            __props__['default_tls_container_ref'] = default_tls_container_ref
            __props__['description'] = description
            __props__['insert_headers'] = insert_headers
            if loadbalancer_id is None:
                raise TypeError("Missing required property 'loadbalancer_id'")
            __props__['loadbalancer_id'] = loadbalancer_id
            __props__['name'] = name
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            if protocol_port is None:
                raise TypeError("Missing required property 'protocol_port'")
            __props__['protocol_port'] = protocol_port
            __props__['region'] = region
            __props__['sni_container_refs'] = sni_container_refs
            __props__['tenant_id'] = tenant_id
            __props__['timeout_client_data'] = timeout_client_data
            __props__['timeout_member_connect'] = timeout_member_connect
            __props__['timeout_member_data'] = timeout_member_data
            __props__['timeout_tcp_inspect'] = timeout_tcp_inspect
        super(Listener, __self__).__init__(
            'openstack:loadbalancer/listener:Listener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            allowed_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            connection_limit: Optional[pulumi.Input[int]] = None,
            default_pool_id: Optional[pulumi.Input[str]] = None,
            default_tls_container_ref: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            insert_headers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            loadbalancer_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            protocol_port: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            sni_container_refs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            timeout_client_data: Optional[pulumi.Input[int]] = None,
            timeout_member_connect: Optional[pulumi.Input[int]] = None,
            timeout_member_data: Optional[pulumi.Input[int]] = None,
            timeout_tcp_inspect: Optional[pulumi.Input[int]] = None) -> 'Listener':
        """
        Get an existing Listener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the Listener.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_cidrs: A list of CIDR blocks that are permitted to connect to this listener, denying
               all other source addresses. If not present, defaults to allow all.
        :param pulumi.Input[int] connection_limit: The maximum number of connections allowed
               for the Listener.
        :param pulumi.Input[str] default_pool_id: The ID of the default pool with which the
               Listener is associated.
        :param pulumi.Input[str] default_tls_container_ref: A reference to a Barbican Secrets
               container which stores TLS information. This is required if the protocol
               is `TERMINATED_HTTPS`. See
               [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
               for more information.
        :param pulumi.Input[str] description: Human-readable description for the Listener.
        :param pulumi.Input[Mapping[str, Any]] insert_headers: The list of key value pairs representing headers to insert
               into the request before it is sent to the backend members. Changing this updates the headers of the
               existing listener.
        :param pulumi.Input[str] loadbalancer_id: The load balancer on which to provision this
               Listener. Changing this creates a new Listener.
        :param pulumi.Input[str] name: Human-readable name for the Listener. Does not have
               to be unique.
        :param pulumi.Input[str] protocol: The protocol - can either be TCP, HTTP, HTTPS,
               TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
               new Listener.
        :param pulumi.Input[int] protocol_port: The port on which to listen for client traffic.
               Changing this creates a new Listener.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               Listener.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sni_container_refs: A list of references to Barbican Secrets
               containers which store SNI information. See
               [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
               for more information.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the Listener.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new Listener.
        :param pulumi.Input[int] timeout_client_data: The client inactivity timeout in milliseconds.
        :param pulumi.Input[int] timeout_member_connect: The member connection timeout in milliseconds.
        :param pulumi.Input[int] timeout_member_data: The member inactivity timeout in milliseconds.
        :param pulumi.Input[int] timeout_tcp_inspect: The time in milliseconds, to wait for additional
               TCP packets for content inspection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_state_up"] = admin_state_up
        __props__["allowed_cidrs"] = allowed_cidrs
        __props__["connection_limit"] = connection_limit
        __props__["default_pool_id"] = default_pool_id
        __props__["default_tls_container_ref"] = default_tls_container_ref
        __props__["description"] = description
        __props__["insert_headers"] = insert_headers
        __props__["loadbalancer_id"] = loadbalancer_id
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["protocol_port"] = protocol_port
        __props__["region"] = region
        __props__["sni_container_refs"] = sni_container_refs
        __props__["tenant_id"] = tenant_id
        __props__["timeout_client_data"] = timeout_client_data
        __props__["timeout_member_connect"] = timeout_member_connect
        __props__["timeout_member_data"] = timeout_member_data
        __props__["timeout_tcp_inspect"] = timeout_tcp_inspect
        return Listener(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        """
        The administrative state of the Listener.
        A valid value is true (UP) or false (DOWN).
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter(name="allowedCidrs")
    def allowed_cidrs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of CIDR blocks that are permitted to connect to this listener, denying
        all other source addresses. If not present, defaults to allow all.
        """
        return pulumi.get(self, "allowed_cidrs")

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> pulumi.Output[int]:
        """
        The maximum number of connections allowed
        for the Listener.
        """
        return pulumi.get(self, "connection_limit")

    @property
    @pulumi.getter(name="defaultPoolId")
    def default_pool_id(self) -> pulumi.Output[str]:
        """
        The ID of the default pool with which the
        Listener is associated.
        """
        return pulumi.get(self, "default_pool_id")

    @property
    @pulumi.getter(name="defaultTlsContainerRef")
    def default_tls_container_ref(self) -> pulumi.Output[Optional[str]]:
        """
        A reference to a Barbican Secrets
        container which stores TLS information. This is required if the protocol
        is `TERMINATED_HTTPS`. See
        [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        for more information.
        """
        return pulumi.get(self, "default_tls_container_ref")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human-readable description for the Listener.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="insertHeaders")
    def insert_headers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The list of key value pairs representing headers to insert
        into the request before it is sent to the backend members. Changing this updates the headers of the
        existing listener.
        """
        return pulumi.get(self, "insert_headers")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> pulumi.Output[str]:
        """
        The load balancer on which to provision this
        Listener. Changing this creates a new Listener.
        """
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Human-readable name for the Listener. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol - can either be TCP, HTTP, HTTPS,
        TERMINATED_HTTPS or UDP (supported only in Octavia). Changing this creates a
        new Listener.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolPort")
    def protocol_port(self) -> pulumi.Output[int]:
        """
        The port on which to listen for client traffic.
        Changing this creates a new Listener.
        """
        return pulumi.get(self, "protocol_port")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an . If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        Listener.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sniContainerRefs")
    def sni_container_refs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of references to Barbican Secrets
        containers which store SNI information. See
        [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
        for more information.
        """
        return pulumi.get(self, "sni_container_refs")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Required for admins. The UUID of the tenant who owns
        the Listener.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new Listener.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="timeoutClientData")
    def timeout_client_data(self) -> pulumi.Output[int]:
        """
        The client inactivity timeout in milliseconds.
        """
        return pulumi.get(self, "timeout_client_data")

    @property
    @pulumi.getter(name="timeoutMemberConnect")
    def timeout_member_connect(self) -> pulumi.Output[int]:
        """
        The member connection timeout in milliseconds.
        """
        return pulumi.get(self, "timeout_member_connect")

    @property
    @pulumi.getter(name="timeoutMemberData")
    def timeout_member_data(self) -> pulumi.Output[int]:
        """
        The member inactivity timeout in milliseconds.
        """
        return pulumi.get(self, "timeout_member_data")

    @property
    @pulumi.getter(name="timeoutTcpInspect")
    def timeout_tcp_inspect(self) -> pulumi.Output[int]:
        """
        The time in milliseconds, to wait for additional
        TCP packets for content inspection.
        """
        return pulumi.get(self, "timeout_tcp_inspect")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

