# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MonitorV1Args', 'MonitorV1']

@pulumi.input_type
class MonitorV1Args:
    def __init__(__self__, *,
                 delay: pulumi.Input[int],
                 max_retries: pulumi.Input[int],
                 timeout: pulumi.Input[int],
                 type: pulumi.Input[str],
                 admin_state_up: Optional[pulumi.Input[str]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MonitorV1 resource.
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
               Changing this creates a new monitor.
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before changing
               the member's status to INACTIVE. Must be a number between 1 and 10. Changing
               this updates the max_retries of the existing monitor.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay value.
               Changing this updates the timeout of the existing monitor.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, or HTTPS,
               that is sent by the monitor to verify the member state. Changing this
               creates a new monitor.
        :param pulumi.Input[str] admin_state_up: The administrative state of the monitor.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing monitor.
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", or a range like "200-202". Changing this updates the expected_codes
               of the existing monitor.
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method used
               for requests by the monitor. If this attribute is not specified, it defaults
               to "GET". Changing this updates the http_method of the existing monitor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB monitor. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB monitor.
        :param pulumi.Input[str] tenant_id: The owner of the monitor. Required if admin wants to
               create a monitor for another tenant. Changing this creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Changing this updates the
               url_path of the existing monitor.
        """
        MonitorV1Args._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            delay=delay,
            max_retries=max_retries,
            timeout=timeout,
            type=type,
            admin_state_up=admin_state_up,
            expected_codes=expected_codes,
            http_method=http_method,
            region=region,
            tenant_id=tenant_id,
            url_path=url_path,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             delay: pulumi.Input[int],
             max_retries: pulumi.Input[int],
             timeout: pulumi.Input[int],
             type: pulumi.Input[str],
             admin_state_up: Optional[pulumi.Input[str]] = None,
             expected_codes: Optional[pulumi.Input[str]] = None,
             http_method: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             url_path: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("delay", delay)
        _setter("max_retries", max_retries)
        _setter("timeout", timeout)
        _setter("type", type)
        if admin_state_up is not None:
            _setter("admin_state_up", admin_state_up)
        if expected_codes is not None:
            _setter("expected_codes", expected_codes)
        if http_method is not None:
            _setter("http_method", http_method)
        if region is not None:
            _setter("region", region)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if url_path is not None:
            _setter("url_path", url_path)

    @property
    @pulumi.getter
    def delay(self) -> pulumi.Input[int]:
        """
        The time, in seconds, between sending probes to members.
        Changing this creates a new monitor.
        """
        return pulumi.get(self, "delay")

    @delay.setter
    def delay(self, value: pulumi.Input[int]):
        pulumi.set(self, "delay", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Input[int]:
        """
        Number of permissible ping failures before changing
        the member's status to INACTIVE. Must be a number between 1 and 10. Changing
        this updates the max_retries of the existing monitor.
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Input[int]:
        """
        Maximum number of seconds for a monitor to wait for a
        ping reply before it times out. The value must be less than the delay value.
        Changing this updates the timeout of the existing monitor.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: pulumi.Input[int]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of probe, which is PING, TCP, HTTP, or HTTPS,
        that is sent by the monitor to verify the member state. Changing this
        creates a new monitor.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[str]]:
        """
        The administrative state of the monitor.
        Acceptable values are "true" and "false". Changing this value updates the
        state of the existing monitor.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="expectedCodes")
    def expected_codes(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. Expected HTTP codes
        for a passing HTTP(S) monitor. You can either specify a single status like
        "200", or a range like "200-202". Changing this updates the expected_codes
        of the existing monitor.
        """
        return pulumi.get(self, "expected_codes")

    @expected_codes.setter
    def expected_codes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expected_codes", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. The HTTP method used
        for requests by the monitor. If this attribute is not specified, it defaults
        to "GET". Changing this updates the http_method of the existing monitor.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB monitor. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB monitor.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the monitor. Required if admin wants to
        create a monitor for another tenant. Changing this creates a new monitor.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. URI path that will be
        accessed if monitor type is HTTP or HTTPS. Changing this updates the
        url_path of the existing monitor.
        """
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_path", value)


@pulumi.input_type
class _MonitorV1State:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[str]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MonitorV1 resources.
        :param pulumi.Input[str] admin_state_up: The administrative state of the monitor.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing monitor.
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
               Changing this creates a new monitor.
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", or a range like "200-202". Changing this updates the expected_codes
               of the existing monitor.
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method used
               for requests by the monitor. If this attribute is not specified, it defaults
               to "GET". Changing this updates the http_method of the existing monitor.
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before changing
               the member's status to INACTIVE. Must be a number between 1 and 10. Changing
               this updates the max_retries of the existing monitor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB monitor. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB monitor.
        :param pulumi.Input[str] tenant_id: The owner of the monitor. Required if admin wants to
               create a monitor for another tenant. Changing this creates a new monitor.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay value.
               Changing this updates the timeout of the existing monitor.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, or HTTPS,
               that is sent by the monitor to verify the member state. Changing this
               creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Changing this updates the
               url_path of the existing monitor.
        """
        _MonitorV1State._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            admin_state_up=admin_state_up,
            delay=delay,
            expected_codes=expected_codes,
            http_method=http_method,
            max_retries=max_retries,
            region=region,
            tenant_id=tenant_id,
            timeout=timeout,
            type=type,
            url_path=url_path,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             admin_state_up: Optional[pulumi.Input[str]] = None,
             delay: Optional[pulumi.Input[int]] = None,
             expected_codes: Optional[pulumi.Input[str]] = None,
             http_method: Optional[pulumi.Input[str]] = None,
             max_retries: Optional[pulumi.Input[int]] = None,
             region: Optional[pulumi.Input[str]] = None,
             tenant_id: Optional[pulumi.Input[str]] = None,
             timeout: Optional[pulumi.Input[int]] = None,
             type: Optional[pulumi.Input[str]] = None,
             url_path: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if admin_state_up is not None:
            _setter("admin_state_up", admin_state_up)
        if delay is not None:
            _setter("delay", delay)
        if expected_codes is not None:
            _setter("expected_codes", expected_codes)
        if http_method is not None:
            _setter("http_method", http_method)
        if max_retries is not None:
            _setter("max_retries", max_retries)
        if region is not None:
            _setter("region", region)
        if tenant_id is not None:
            _setter("tenant_id", tenant_id)
        if timeout is not None:
            _setter("timeout", timeout)
        if type is not None:
            _setter("type", type)
        if url_path is not None:
            _setter("url_path", url_path)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[str]]:
        """
        The administrative state of the monitor.
        Acceptable values are "true" and "false". Changing this value updates the
        state of the existing monitor.
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def delay(self) -> Optional[pulumi.Input[int]]:
        """
        The time, in seconds, between sending probes to members.
        Changing this creates a new monitor.
        """
        return pulumi.get(self, "delay")

    @delay.setter
    def delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delay", value)

    @property
    @pulumi.getter(name="expectedCodes")
    def expected_codes(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. Expected HTTP codes
        for a passing HTTP(S) monitor. You can either specify a single status like
        "200", or a range like "200-202". Changing this updates the expected_codes
        of the existing monitor.
        """
        return pulumi.get(self, "expected_codes")

    @expected_codes.setter
    def expected_codes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expected_codes", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. The HTTP method used
        for requests by the monitor. If this attribute is not specified, it defaults
        to "GET". Changing this updates the http_method of the existing monitor.
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[int]]:
        """
        Number of permissible ping failures before changing
        the member's status to INACTIVE. Must be a number between 1 and 10. Changing
        this updates the max_retries of the existing monitor.
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB monitor. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB monitor.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the monitor. Required if admin wants to
        create a monitor for another tenant. Changing this creates a new monitor.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of seconds for a monitor to wait for a
        ping reply before it times out. The value must be less than the delay value.
        Changing this updates the timeout of the existing monitor.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of probe, which is PING, TCP, HTTP, or HTTPS,
        that is sent by the monitor to verify the member state. Changing this
        creates a new monitor.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. URI path that will be
        accessed if monitor type is HTTP or HTTPS. Changing this updates the
        url_path of the existing monitor.
        """
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_path", value)


class MonitorV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[str]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V1 load balancer monitor resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        monitor1 = openstack.loadbalancer.MonitorV1("monitor1",
            admin_state_up="true",
            delay=30,
            max_retries=3,
            timeout=5,
            type="PING")
        ```

        ## Import

        Load Balancer Members can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:loadbalancer/monitorV1:MonitorV1 monitor_1 119d7530-72e9-449a-aa97-124a5ef1992c
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state_up: The administrative state of the monitor.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing monitor.
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
               Changing this creates a new monitor.
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", or a range like "200-202". Changing this updates the expected_codes
               of the existing monitor.
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method used
               for requests by the monitor. If this attribute is not specified, it defaults
               to "GET". Changing this updates the http_method of the existing monitor.
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before changing
               the member's status to INACTIVE. Must be a number between 1 and 10. Changing
               this updates the max_retries of the existing monitor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB monitor. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB monitor.
        :param pulumi.Input[str] tenant_id: The owner of the monitor. Required if admin wants to
               create a monitor for another tenant. Changing this creates a new monitor.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay value.
               Changing this updates the timeout of the existing monitor.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, or HTTPS,
               that is sent by the monitor to verify the member state. Changing this
               creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Changing this updates the
               url_path of the existing monitor.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MonitorV1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V1 load balancer monitor resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        monitor1 = openstack.loadbalancer.MonitorV1("monitor1",
            admin_state_up="true",
            delay=30,
            max_retries=3,
            timeout=5,
            type="PING")
        ```

        ## Import

        Load Balancer Members can be imported using the `id`, e.g.

        ```sh
         $ pulumi import openstack:loadbalancer/monitorV1:MonitorV1 monitor_1 119d7530-72e9-449a-aa97-124a5ef1992c
        ```

        :param str resource_name: The name of the resource.
        :param MonitorV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MonitorV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            MonitorV1Args._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[str]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MonitorV1Args.__new__(MonitorV1Args)

            __props__.__dict__["admin_state_up"] = admin_state_up
            if delay is None and not opts.urn:
                raise TypeError("Missing required property 'delay'")
            __props__.__dict__["delay"] = delay
            __props__.__dict__["expected_codes"] = expected_codes
            __props__.__dict__["http_method"] = http_method
            if max_retries is None and not opts.urn:
                raise TypeError("Missing required property 'max_retries'")
            __props__.__dict__["max_retries"] = max_retries
            __props__.__dict__["region"] = region
            __props__.__dict__["tenant_id"] = tenant_id
            if timeout is None and not opts.urn:
                raise TypeError("Missing required property 'timeout'")
            __props__.__dict__["timeout"] = timeout
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["url_path"] = url_path
        super(MonitorV1, __self__).__init__(
            'openstack:loadbalancer/monitorV1:MonitorV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[str]] = None,
            delay: Optional[pulumi.Input[int]] = None,
            expected_codes: Optional[pulumi.Input[str]] = None,
            http_method: Optional[pulumi.Input[str]] = None,
            max_retries: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url_path: Optional[pulumi.Input[str]] = None) -> 'MonitorV1':
        """
        Get an existing MonitorV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state_up: The administrative state of the monitor.
               Acceptable values are "true" and "false". Changing this value updates the
               state of the existing monitor.
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
               Changing this creates a new monitor.
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", or a range like "200-202". Changing this updates the expected_codes
               of the existing monitor.
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method used
               for requests by the monitor. If this attribute is not specified, it defaults
               to "GET". Changing this updates the http_method of the existing monitor.
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before changing
               the member's status to INACTIVE. Must be a number between 1 and 10. Changing
               this updates the max_retries of the existing monitor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an LB monitor. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               LB monitor.
        :param pulumi.Input[str] tenant_id: The owner of the monitor. Required if admin wants to
               create a monitor for another tenant. Changing this creates a new monitor.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay value.
               Changing this updates the timeout of the existing monitor.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, or HTTPS,
               that is sent by the monitor to verify the member state. Changing this
               creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Changing this updates the
               url_path of the existing monitor.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MonitorV1State.__new__(_MonitorV1State)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["delay"] = delay
        __props__.__dict__["expected_codes"] = expected_codes
        __props__.__dict__["http_method"] = http_method
        __props__.__dict__["max_retries"] = max_retries
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["type"] = type
        __props__.__dict__["url_path"] = url_path
        return MonitorV1(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[str]:
        """
        The administrative state of the monitor.
        Acceptable values are "true" and "false". Changing this value updates the
        state of the existing monitor.
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def delay(self) -> pulumi.Output[int]:
        """
        The time, in seconds, between sending probes to members.
        Changing this creates a new monitor.
        """
        return pulumi.get(self, "delay")

    @property
    @pulumi.getter(name="expectedCodes")
    def expected_codes(self) -> pulumi.Output[Optional[str]]:
        """
        Required for HTTP(S) types. Expected HTTP codes
        for a passing HTTP(S) monitor. You can either specify a single status like
        "200", or a range like "200-202". Changing this updates the expected_codes
        of the existing monitor.
        """
        return pulumi.get(self, "expected_codes")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Output[Optional[str]]:
        """
        Required for HTTP(S) types. The HTTP method used
        for requests by the monitor. If this attribute is not specified, it defaults
        to "GET". Changing this updates the http_method of the existing monitor.
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Output[int]:
        """
        Number of permissible ping failures before changing
        the member's status to INACTIVE. Must be a number between 1 and 10. Changing
        this updates the max_retries of the existing monitor.
        """
        return pulumi.get(self, "max_retries")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an LB monitor. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        LB monitor.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The owner of the monitor. Required if admin wants to
        create a monitor for another tenant. Changing this creates a new monitor.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[int]:
        """
        Maximum number of seconds for a monitor to wait for a
        ping reply before it times out. The value must be less than the delay value.
        Changing this updates the timeout of the existing monitor.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of probe, which is PING, TCP, HTTP, or HTTPS,
        that is sent by the monitor to verify the member state. Changing this
        creates a new monitor.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> pulumi.Output[Optional[str]]:
        """
        Required for HTTP(S) types. URI path that will be
        accessed if monitor type is HTTP or HTTPS. Changing this updates the
        url_path of the existing monitor.
        """
        return pulumi.get(self, "url_path")

