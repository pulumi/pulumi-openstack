# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDnsZoneResult:
    """
    A collection of values returned by getDnsZone.
    """
    def __init__(__self__, attributes=None, created_at=None, description=None, email=None, id=None, masters=None, name=None, pool_id=None, project_id=None, region=None, serial=None, status=None, transferred_at=None, ttl=None, type=None, updated_at=None, version=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        __self__.attributes = attributes
        """
        Attributes of the DNS Service scheduler.
        """
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        The time the zone was created.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        See Argument Reference above.
        """
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        __self__.email = email
        """
        See Argument Reference above.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if masters and not isinstance(masters, list):
            raise TypeError("Expected argument 'masters' to be a list")
        __self__.masters = masters
        """
        An array of master DNS servers. When `type` is  `SECONDARY`.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        See Argument Reference above.
        """
        if pool_id and not isinstance(pool_id, str):
            raise TypeError("Expected argument 'pool_id' to be a str")
        __self__.pool_id = pool_id
        """
        The ID of the pool hosting the zone.
        """
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        __self__.project_id = project_id
        """
        The project ID that owns the zone.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        See Argument Reference above.
        """
        if serial and not isinstance(serial, float):
            raise TypeError("Expected argument 'serial' to be a float")
        __self__.serial = serial
        """
        The serial number of the zone.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        See Argument Reference above.
        """
        if transferred_at and not isinstance(transferred_at, str):
            raise TypeError("Expected argument 'transferred_at' to be a str")
        __self__.transferred_at = transferred_at
        """
        The time the zone was transferred.
        """
        if ttl and not isinstance(ttl, float):
            raise TypeError("Expected argument 'ttl' to be a float")
        __self__.ttl = ttl
        """
        See Argument Reference above.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        See Argument Reference above.
        """
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        __self__.updated_at = updated_at
        """
        The time the zone was last updated.
        """
        if version and not isinstance(version, float):
            raise TypeError("Expected argument 'version' to be a float")
        __self__.version = version
        """
        The version of the zone.
        """
class AwaitableGetDnsZoneResult(GetDnsZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDnsZoneResult(
            attributes=self.attributes,
            created_at=self.created_at,
            description=self.description,
            email=self.email,
            id=self.id,
            masters=self.masters,
            name=self.name,
            pool_id=self.pool_id,
            project_id=self.project_id,
            region=self.region,
            serial=self.serial,
            status=self.status,
            transferred_at=self.transferred_at,
            ttl=self.ttl,
            type=self.type,
            updated_at=self.updated_at,
            version=self.version)

def get_dns_zone(attributes=None,created_at=None,description=None,email=None,masters=None,name=None,pool_id=None,project_id=None,region=None,serial=None,status=None,transferred_at=None,ttl=None,type=None,updated_at=None,version=None,opts=None):
    """
    Use this data source to get the ID of an available OpenStack DNS zone.




    :param dict attributes: Attributes of the DNS Service scheduler.
    :param str created_at: The time the zone was created.
    :param str description: A description of the zone.
    :param str email: The email contact for the zone record.
    :param list masters: An array of master DNS servers. When `type` is  `SECONDARY`.
    :param str name: The name of the zone.
    :param str pool_id: The ID of the pool hosting the zone.
    :param str project_id: The project ID that owns the zone.
    :param str region: The region in which to obtain the V2 DNS client.
           A DNS client is needed to retrieve zone ids. If omitted, the
           `region` argument of the provider is used.
    :param float serial: The serial number of the zone.
    :param str status: The zone's status.
    :param str transferred_at: The time the zone was transferred.
    :param float ttl: The time to live (TTL) of the zone.
    :param str type: The type of the zone. Can either be `PRIMARY` or `SECONDARY`.
    :param str updated_at: The time the zone was last updated.
    :param float version: The version of the zone.
    """
    __args__ = dict()


    __args__['attributes'] = attributes
    __args__['createdAt'] = created_at
    __args__['description'] = description
    __args__['email'] = email
    __args__['masters'] = masters
    __args__['name'] = name
    __args__['poolId'] = pool_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['serial'] = serial
    __args__['status'] = status
    __args__['transferredAt'] = transferred_at
    __args__['ttl'] = ttl
    __args__['type'] = type
    __args__['updatedAt'] = updated_at
    __args__['version'] = version
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('openstack:dns/getDnsZone:getDnsZone', __args__, opts=opts).value

    return AwaitableGetDnsZoneResult(
        attributes=__ret__.get('attributes'),
        created_at=__ret__.get('createdAt'),
        description=__ret__.get('description'),
        email=__ret__.get('email'),
        id=__ret__.get('id'),
        masters=__ret__.get('masters'),
        name=__ret__.get('name'),
        pool_id=__ret__.get('poolId'),
        project_id=__ret__.get('projectId'),
        region=__ret__.get('region'),
        serial=__ret__.get('serial'),
        status=__ret__.get('status'),
        transferred_at=__ret__.get('transferredAt'),
        ttl=__ret__.get('ttl'),
        type=__ret__.get('type'),
        updated_at=__ret__.get('updatedAt'),
        version=__ret__.get('version'))
