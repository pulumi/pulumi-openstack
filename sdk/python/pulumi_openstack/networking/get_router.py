# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRouterResult:
    """
    A collection of values returned by getRouter.
    """
    def __init__(__self__, admin_state_up=None, all_tags=None, availability_zone_hints=None, description=None, distributed=None, enable_snat=None, external_fixed_ips=None, external_network_id=None, name=None, region=None, router_id=None, status=None, tags=None, tenant_id=None, id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        __self__.admin_state_up = admin_state_up
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        __self__.all_tags = all_tags
        """
        The set of string tags applied on the router.
        """
        if availability_zone_hints and not isinstance(availability_zone_hints, list):
            raise TypeError("Expected argument 'availability_zone_hints' to be a list")
        __self__.availability_zone_hints = availability_zone_hints
        """
        The availability zone that is used to make router resources highly available.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if distributed and not isinstance(distributed, bool):
            raise TypeError("Expected argument 'distributed' to be a bool")
        __self__.distributed = distributed
        if enable_snat and not isinstance(enable_snat, bool):
            raise TypeError("Expected argument 'enable_snat' to be a bool")
        __self__.enable_snat = enable_snat
        """
        The value that points out if the Source NAT is enabled on the router.
        """
        if external_fixed_ips and not isinstance(external_fixed_ips, list):
            raise TypeError("Expected argument 'external_fixed_ips' to be a list")
        __self__.external_fixed_ips = external_fixed_ips
        """
        The external fixed IPs of the router.
        """
        if external_network_id and not isinstance(external_network_id, str):
            raise TypeError("Expected argument 'external_network_id' to be a str")
        __self__.external_network_id = external_network_id
        """
        The network UUID of an external gateway for the router.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if router_id and not isinstance(router_id, str):
            raise TypeError("Expected argument 'router_id' to be a str")
        __self__.router_id = router_id
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        __self__.tenant_id = tenant_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetRouterResult(GetRouterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterResult(
            admin_state_up=self.admin_state_up,
            all_tags=self.all_tags,
            availability_zone_hints=self.availability_zone_hints,
            description=self.description,
            distributed=self.distributed,
            enable_snat=self.enable_snat,
            external_fixed_ips=self.external_fixed_ips,
            external_network_id=self.external_network_id,
            name=self.name,
            region=self.region,
            router_id=self.router_id,
            status=self.status,
            tags=self.tags,
            tenant_id=self.tenant_id,
            id=self.id)

def get_router(admin_state_up=None,description=None,distributed=None,enable_snat=None,name=None,region=None,router_id=None,status=None,tags=None,tenant_id=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack router.
    
    :param bool admin_state_up: Administrative up/down status for the router (must be "true" or "false" if provided).
    :param str description: Human-readable description of the router.
    :param bool distributed: Indicates whether or not to get a distributed router.
    :param str name: The name of the router.
    :param str region: The region in which to obtain the V2 Neutron client.
           A Neutron client is needed to retrieve router ids. If omitted, the
           `region` argument of the provider is used.
    :param str router_id: The UUID of the router resource.
    :param str status: The status of the router (ACTIVE/DOWN).
    :param list tags: The list of router tags to filter.
    :param str tenant_id: The owner of the router.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_router_v2.html.markdown.
    """
    __args__ = dict()

    __args__['adminStateUp'] = admin_state_up
    __args__['description'] = description
    __args__['distributed'] = distributed
    __args__['enableSnat'] = enable_snat
    __args__['name'] = name
    __args__['region'] = region
    __args__['routerId'] = router_id
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['tenantId'] = tenant_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:networking/getRouter:getRouter', __args__, opts=opts).value

    return AwaitableGetRouterResult(
        admin_state_up=__ret__.get('adminStateUp'),
        all_tags=__ret__.get('allTags'),
        availability_zone_hints=__ret__.get('availabilityZoneHints'),
        description=__ret__.get('description'),
        distributed=__ret__.get('distributed'),
        enable_snat=__ret__.get('enableSnat'),
        external_fixed_ips=__ret__.get('externalFixedIps'),
        external_network_id=__ret__.get('externalNetworkId'),
        name=__ret__.get('name'),
        region=__ret__.get('region'),
        router_id=__ret__.get('routerId'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'),
        tenant_id=__ret__.get('tenantId'),
        id=__ret__.get('id'))
