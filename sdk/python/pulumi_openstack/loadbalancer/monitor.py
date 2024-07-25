# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MonitorArgs', 'Monitor']

@pulumi.input_type
class MonitorArgs:
    def __init__(__self__, *,
                 delay: pulumi.Input[int],
                 max_retries: pulumi.Input[int],
                 pool_id: pulumi.Input[str],
                 timeout: pulumi.Input[int],
                 type: pulumi.Input[str],
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 max_retries_down: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Monitor resource.
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before
               changing the member's status to INACTIVE. Must be a number between 1
               and 10.
        :param pulumi.Input[str] pool_id: The id of the pool that this monitor will be assigned to.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay
               value.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, HTTPS,
               TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
               verify the member state. Changing this creates a new monitor.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the monitor.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", a list like "200, 202" or a range like "200-202". Default is "200".
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method that 
               the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
               OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
        :param pulumi.Input[int] max_retries_down: Number of permissible ping failures before 
               changing the member's status to ERROR. Must be a number between 1 and 10.
               The default is 3. Changing this updates the max_retries_down of the
               existing monitor.
        :param pulumi.Input[str] name: The Name of the Monitor.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               monitor.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the monitor.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Default is `/`.
        """
        pulumi.set(__self__, "delay", delay)
        pulumi.set(__self__, "max_retries", max_retries)
        pulumi.set(__self__, "pool_id", pool_id)
        pulumi.set(__self__, "timeout", timeout)
        pulumi.set(__self__, "type", type)
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if expected_codes is not None:
            pulumi.set(__self__, "expected_codes", expected_codes)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if max_retries_down is not None:
            pulumi.set(__self__, "max_retries_down", max_retries_down)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)

    @property
    @pulumi.getter
    def delay(self) -> pulumi.Input[int]:
        """
        The time, in seconds, between sending probes to members.
        """
        return pulumi.get(self, "delay")

    @delay.setter
    def delay(self, value: pulumi.Input[int]):
        pulumi.set(self, "delay", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Input[int]:
        """
        Number of permissible ping failures before
        changing the member's status to INACTIVE. Must be a number between 1
        and 10.
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Input[str]:
        """
        The id of the pool that this monitor will be assigned to.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Input[int]:
        """
        Maximum number of seconds for a monitor to wait for a
        ping reply before it times out. The value must be less than the delay
        value.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: pulumi.Input[int]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of probe, which is PING, TCP, HTTP, HTTPS,
        TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
        verify the member state. Changing this creates a new monitor.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the monitor.
        A valid value is true (UP) or false (DOWN).
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter(name="expectedCodes")
    def expected_codes(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. Expected HTTP codes
        for a passing HTTP(S) monitor. You can either specify a single status like
        "200", a list like "200, 202" or a range like "200-202". Default is "200".
        """
        return pulumi.get(self, "expected_codes")

    @expected_codes.setter
    def expected_codes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expected_codes", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. The HTTP method that 
        the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
        OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="maxRetriesDown")
    def max_retries_down(self) -> Optional[pulumi.Input[int]]:
        """
        Number of permissible ping failures before 
        changing the member's status to ERROR. Must be a number between 1 and 10.
        The default is 3. Changing this updates the max_retries_down of the
        existing monitor.
        """
        return pulumi.get(self, "max_retries_down")

    @max_retries_down.setter
    def max_retries_down(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries_down", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the Monitor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an . If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        monitor.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required for admins. The UUID of the tenant who owns
        the monitor.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new monitor.
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
        accessed if monitor type is HTTP or HTTPS. Default is `/`.
        """
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_path", value)


@pulumi.input_type
class _MonitorState:
    def __init__(__self__, *,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 max_retries_down: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Monitor resources.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the monitor.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", a list like "200, 202" or a range like "200-202". Default is "200".
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method that 
               the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
               OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before
               changing the member's status to INACTIVE. Must be a number between 1
               and 10.
        :param pulumi.Input[int] max_retries_down: Number of permissible ping failures before 
               changing the member's status to ERROR. Must be a number between 1 and 10.
               The default is 3. Changing this updates the max_retries_down of the
               existing monitor.
        :param pulumi.Input[str] name: The Name of the Monitor.
        :param pulumi.Input[str] pool_id: The id of the pool that this monitor will be assigned to.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               monitor.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the monitor.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new monitor.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay
               value.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, HTTPS,
               TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
               verify the member state. Changing this creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Default is `/`.
        """
        if admin_state_up is not None:
            pulumi.set(__self__, "admin_state_up", admin_state_up)
        if delay is not None:
            pulumi.set(__self__, "delay", delay)
        if expected_codes is not None:
            pulumi.set(__self__, "expected_codes", expected_codes)
        if http_method is not None:
            pulumi.set(__self__, "http_method", http_method)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if max_retries_down is not None:
            pulumi.set(__self__, "max_retries_down", max_retries_down)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pool_id is not None:
            pulumi.set(__self__, "pool_id", pool_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if url_path is not None:
            pulumi.set(__self__, "url_path", url_path)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> Optional[pulumi.Input[bool]]:
        """
        The administrative state of the monitor.
        A valid value is true (UP) or false (DOWN).
        """
        return pulumi.get(self, "admin_state_up")

    @admin_state_up.setter
    def admin_state_up(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_state_up", value)

    @property
    @pulumi.getter
    def delay(self) -> Optional[pulumi.Input[int]]:
        """
        The time, in seconds, between sending probes to members.
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
        "200", a list like "200, 202" or a range like "200-202". Default is "200".
        """
        return pulumi.get(self, "expected_codes")

    @expected_codes.setter
    def expected_codes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expected_codes", value)

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> Optional[pulumi.Input[str]]:
        """
        Required for HTTP(S) types. The HTTP method that 
        the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
        OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
        """
        return pulumi.get(self, "http_method")

    @http_method.setter
    def http_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_method", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[int]]:
        """
        Number of permissible ping failures before
        changing the member's status to INACTIVE. Must be a number between 1
        and 10.
        """
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter(name="maxRetriesDown")
    def max_retries_down(self) -> Optional[pulumi.Input[int]]:
        """
        Number of permissible ping failures before 
        changing the member's status to ERROR. Must be a number between 1 and 10.
        The default is 3. Changing this updates the max_retries_down of the
        existing monitor.
        """
        return pulumi.get(self, "max_retries_down")

    @max_retries_down.setter
    def max_retries_down(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries_down", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The Name of the Monitor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the pool that this monitor will be assigned to.
        """
        return pulumi.get(self, "pool_id")

    @pool_id.setter
    def pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an . If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        monitor.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required for admins. The UUID of the tenant who owns
        the monitor.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new monitor.
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
        ping reply before it times out. The value must be less than the delay
        value.
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of probe, which is PING, TCP, HTTP, HTTPS,
        TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
        verify the member state. Changing this creates a new monitor.
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
        accessed if monitor type is HTTP or HTTPS. Default is `/`.
        """
        return pulumi.get(self, "url_path")

    @url_path.setter
    def url_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_path", value)


class Monitor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 max_retries_down: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url_path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V2 monitor resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        monitor1 = openstack.loadbalancer.Monitor("monitor_1",
            pool_id=pool1["id"],
            type="PING",
            delay=20,
            timeout=10,
            max_retries=5)
        ```

        ## Import

        Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.:

        ```sh
        $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2
        ```
        In case of using OpenContrail, the import may not work properly. If you face an issue, try to import the monitor providing its parent pool ID:

        ```sh
        $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2/708bc224-0f8c-4981-ac82-97095fe051b6
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the monitor.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", a list like "200, 202" or a range like "200-202". Default is "200".
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method that 
               the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
               OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before
               changing the member's status to INACTIVE. Must be a number between 1
               and 10.
        :param pulumi.Input[int] max_retries_down: Number of permissible ping failures before 
               changing the member's status to ERROR. Must be a number between 1 and 10.
               The default is 3. Changing this updates the max_retries_down of the
               existing monitor.
        :param pulumi.Input[str] name: The Name of the Monitor.
        :param pulumi.Input[str] pool_id: The id of the pool that this monitor will be assigned to.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               monitor.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the monitor.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new monitor.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay
               value.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, HTTPS,
               TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
               verify the member state. Changing this creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Default is `/`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MonitorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 monitor resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        monitor1 = openstack.loadbalancer.Monitor("monitor_1",
            pool_id=pool1["id"],
            type="PING",
            delay=20,
            timeout=10,
            max_retries=5)
        ```

        ## Import

        Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.:

        ```sh
        $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2
        ```
        In case of using OpenContrail, the import may not work properly. If you face an issue, try to import the monitor providing its parent pool ID:

        ```sh
        $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2/708bc224-0f8c-4981-ac82-97095fe051b6
        ```

        :param str resource_name: The name of the resource.
        :param MonitorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MonitorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state_up: Optional[pulumi.Input[bool]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 expected_codes: Optional[pulumi.Input[str]] = None,
                 http_method: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 max_retries_down: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = MonitorArgs.__new__(MonitorArgs)

            __props__.__dict__["admin_state_up"] = admin_state_up
            if delay is None and not opts.urn:
                raise TypeError("Missing required property 'delay'")
            __props__.__dict__["delay"] = delay
            __props__.__dict__["expected_codes"] = expected_codes
            __props__.__dict__["http_method"] = http_method
            if max_retries is None and not opts.urn:
                raise TypeError("Missing required property 'max_retries'")
            __props__.__dict__["max_retries"] = max_retries
            __props__.__dict__["max_retries_down"] = max_retries_down
            __props__.__dict__["name"] = name
            if pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'pool_id'")
            __props__.__dict__["pool_id"] = pool_id
            __props__.__dict__["region"] = region
            __props__.__dict__["tenant_id"] = tenant_id
            if timeout is None and not opts.urn:
                raise TypeError("Missing required property 'timeout'")
            __props__.__dict__["timeout"] = timeout
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["url_path"] = url_path
        super(Monitor, __self__).__init__(
            'openstack:loadbalancer/monitor:Monitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state_up: Optional[pulumi.Input[bool]] = None,
            delay: Optional[pulumi.Input[int]] = None,
            expected_codes: Optional[pulumi.Input[str]] = None,
            http_method: Optional[pulumi.Input[str]] = None,
            max_retries: Optional[pulumi.Input[int]] = None,
            max_retries_down: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pool_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url_path: Optional[pulumi.Input[str]] = None) -> 'Monitor':
        """
        Get an existing Monitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_state_up: The administrative state of the monitor.
               A valid value is true (UP) or false (DOWN).
        :param pulumi.Input[int] delay: The time, in seconds, between sending probes to members.
        :param pulumi.Input[str] expected_codes: Required for HTTP(S) types. Expected HTTP codes
               for a passing HTTP(S) monitor. You can either specify a single status like
               "200", a list like "200, 202" or a range like "200-202". Default is "200".
        :param pulumi.Input[str] http_method: Required for HTTP(S) types. The HTTP method that 
               the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
               OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
        :param pulumi.Input[int] max_retries: Number of permissible ping failures before
               changing the member's status to INACTIVE. Must be a number between 1
               and 10.
        :param pulumi.Input[int] max_retries_down: Number of permissible ping failures before 
               changing the member's status to ERROR. Must be a number between 1 and 10.
               The default is 3. Changing this updates the max_retries_down of the
               existing monitor.
        :param pulumi.Input[str] name: The Name of the Monitor.
        :param pulumi.Input[str] pool_id: The id of the pool that this monitor will be assigned to.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an . If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               monitor.
        :param pulumi.Input[str] tenant_id: Required for admins. The UUID of the tenant who owns
               the monitor.  Only administrative users can specify a tenant UUID
               other than their own. Changing this creates a new monitor.
        :param pulumi.Input[int] timeout: Maximum number of seconds for a monitor to wait for a
               ping reply before it times out. The value must be less than the delay
               value.
        :param pulumi.Input[str] type: The type of probe, which is PING, TCP, HTTP, HTTPS,
               TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
               verify the member state. Changing this creates a new monitor.
        :param pulumi.Input[str] url_path: Required for HTTP(S) types. URI path that will be
               accessed if monitor type is HTTP or HTTPS. Default is `/`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MonitorState.__new__(_MonitorState)

        __props__.__dict__["admin_state_up"] = admin_state_up
        __props__.__dict__["delay"] = delay
        __props__.__dict__["expected_codes"] = expected_codes
        __props__.__dict__["http_method"] = http_method
        __props__.__dict__["max_retries"] = max_retries
        __props__.__dict__["max_retries_down"] = max_retries_down
        __props__.__dict__["name"] = name
        __props__.__dict__["pool_id"] = pool_id
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["timeout"] = timeout
        __props__.__dict__["type"] = type
        __props__.__dict__["url_path"] = url_path
        return Monitor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminStateUp")
    def admin_state_up(self) -> pulumi.Output[Optional[bool]]:
        """
        The administrative state of the monitor.
        A valid value is true (UP) or false (DOWN).
        """
        return pulumi.get(self, "admin_state_up")

    @property
    @pulumi.getter
    def delay(self) -> pulumi.Output[int]:
        """
        The time, in seconds, between sending probes to members.
        """
        return pulumi.get(self, "delay")

    @property
    @pulumi.getter(name="expectedCodes")
    def expected_codes(self) -> pulumi.Output[str]:
        """
        Required for HTTP(S) types. Expected HTTP codes
        for a passing HTTP(S) monitor. You can either specify a single status like
        "200", a list like "200, 202" or a range like "200-202". Default is "200".
        """
        return pulumi.get(self, "expected_codes")

    @property
    @pulumi.getter(name="httpMethod")
    def http_method(self) -> pulumi.Output[str]:
        """
        Required for HTTP(S) types. The HTTP method that 
        the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
        OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
        """
        return pulumi.get(self, "http_method")

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> pulumi.Output[int]:
        """
        Number of permissible ping failures before
        changing the member's status to INACTIVE. Must be a number between 1
        and 10.
        """
        return pulumi.get(self, "max_retries")

    @property
    @pulumi.getter(name="maxRetriesDown")
    def max_retries_down(self) -> pulumi.Output[int]:
        """
        Number of permissible ping failures before 
        changing the member's status to ERROR. Must be a number between 1 and 10.
        The default is 3. Changing this updates the max_retries_down of the
        existing monitor.
        """
        return pulumi.get(self, "max_retries_down")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The Name of the Monitor.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolId")
    def pool_id(self) -> pulumi.Output[str]:
        """
        The id of the pool that this monitor will be assigned to.
        """
        return pulumi.get(self, "pool_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an . If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        monitor.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Required for admins. The UUID of the tenant who owns
        the monitor.  Only administrative users can specify a tenant UUID
        other than their own. Changing this creates a new monitor.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[int]:
        """
        Maximum number of seconds for a monitor to wait for a
        ping reply before it times out. The value must be less than the delay
        value.
        """
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of probe, which is PING, TCP, HTTP, HTTPS,
        TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
        verify the member state. Changing this creates a new monitor.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="urlPath")
    def url_path(self) -> pulumi.Output[str]:
        """
        Required for HTTP(S) types. URI path that will be
        accessed if monitor type is HTTP or HTTPS. Default is `/`.
        """
        return pulumi.get(self, "url_path")

