# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .endpoint_group import *
from .ike_policy import *
from .ip_sec_policy import *
from .service import *
from .site_connection import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "openstack:vpnaas/endpointGroup:EndpointGroup":
                return EndpointGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "openstack:vpnaas/ikePolicy:IkePolicy":
                return IkePolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "openstack:vpnaas/ipSecPolicy:IpSecPolicy":
                return IpSecPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "openstack:vpnaas/service:Service":
                return Service(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "openstack:vpnaas/siteConnection:SiteConnection":
                return SiteConnection(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("openstack", "vpnaas/endpointGroup", _module_instance)
    pulumi.runtime.register_resource_module("openstack", "vpnaas/ikePolicy", _module_instance)
    pulumi.runtime.register_resource_module("openstack", "vpnaas/ipSecPolicy", _module_instance)
    pulumi.runtime.register_resource_module("openstack", "vpnaas/service", _module_instance)
    pulumi.runtime.register_resource_module("openstack", "vpnaas/siteConnection", _module_instance)

_register_module()
