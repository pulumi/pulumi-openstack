# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetTrunkResult:
    """
    A collection of values returned by getTrunk.
    """
    def __init__(__self__, admin_state_up=None, all_tags=None, description=None, name=None, port_id=None, project_id=None, region=None, status=None, sub_ports=None, tags=None, trunk_id=None, id=None):
        if admin_state_up and not isinstance(admin_state_up, bool):
            raise TypeError("Expected argument 'admin_state_up' to be a bool")
        __self__.admin_state_up = admin_state_up
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        __self__.all_tags = all_tags
        """
        The set of string tags applied on the trunk.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        __self__.port_id = port_id
        """
        The ID of the trunk subport.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        if sub_ports and not isinstance(sub_ports, list):
            raise TypeError("Expected argument 'sub_ports' to be a list")
        __self__.sub_ports = sub_ports
        """
        The set of the trunk subports. The structure of each subport is
        described below.
        """
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        if trunk_id and not isinstance(trunk_id, str):
            raise TypeError("Expected argument 'trunk_id' to be a str")
        __self__.trunk_id = trunk_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_trunk(admin_state_up=None,description=None,name=None,port_id=None,project_id=None,region=None,status=None,tags=None,trunk_id=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack trunk.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_trunk_v2.html.markdown.
    """
    __args__ = dict()

    __args__['adminStateUp'] = admin_state_up
    __args__['description'] = description
    __args__['name'] = name
    __args__['portId'] = port_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['status'] = status
    __args__['tags'] = tags
    __args__['trunkId'] = trunk_id
    __ret__ = await pulumi.runtime.invoke('openstack:networking/getTrunk:getTrunk', __args__, opts=opts)

    return GetTrunkResult(
        admin_state_up=__ret__.get('adminStateUp'),
        all_tags=__ret__.get('allTags'),
        description=__ret__.get('description'),
        name=__ret__.get('name'),
        port_id=__ret__.get('portId'),
        project_id=__ret__.get('projectId'),
        region=__ret__.get('region'),
        status=__ret__.get('status'),
        sub_ports=__ret__.get('subPorts'),
        tags=__ret__.get('tags'),
        trunk_id=__ret__.get('trunkId'),
        id=__ret__.get('id'))
