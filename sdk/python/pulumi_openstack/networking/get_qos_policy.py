# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetQosPolicyResult:
    """
    A collection of values returned by getQosPolicy.
    """
    def __init__(__self__, all_tags=None, created_at=None, description=None, is_default=None, name=None, project_id=None, region=None, revision_number=None, shared=None, tags=None, updated_at=None, id=None):
        if all_tags and not isinstance(all_tags, list):
            raise TypeError("Expected argument 'all_tags' to be a list")
        __self__.all_tags = all_tags
        """
        The set of string tags applied on the QoS policy.
        """
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        The time at which QoS policy was created.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        See Argument Reference above.
        """
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        __self__.is_default = is_default
        """
        See Argument Reference above.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if revision_number and not isinstance(revision_number, float):
            raise TypeError("Expected argument 'revision_number' to be a float")
        __self__.revision_number = revision_number
        """
        The revision number of the QoS policy.
        """
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        __self__.shared = shared
        """
        See Argument Reference above.
        """
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        __self__.tags = tags
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        __self__.updated_at = updated_at
        """
        The time at which QoS policy was created.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_qos_policy(description=None,is_default=None,name=None,project_id=None,region=None,shared=None,tags=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack QoS policy.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/d/networking_qos_policy_v2.html.markdown.
    """
    __args__ = dict()

    __args__['description'] = description
    __args__['isDefault'] = is_default
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['shared'] = shared
    __args__['tags'] = tags
    __ret__ = await pulumi.runtime.invoke('openstack:networking/getQosPolicy:getQosPolicy', __args__, opts=opts)

    return GetQosPolicyResult(
        all_tags=__ret__.get('allTags'),
        created_at=__ret__.get('createdAt'),
        description=__ret__.get('description'),
        is_default=__ret__.get('isDefault'),
        name=__ret__.get('name'),
        project_id=__ret__.get('projectId'),
        region=__ret__.get('region'),
        revision_number=__ret__.get('revisionNumber'),
        shared=__ret__.get('shared'),
        tags=__ret__.get('tags'),
        updated_at=__ret__.get('updatedAt'),
        id=__ret__.get('id'))