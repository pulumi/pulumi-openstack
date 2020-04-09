# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class QuoteSetV2(pulumi.CustomResource):
    backup_gigabytes: pulumi.Output[float]
    """
    Quota value for backup gigabytes. Changing
    this updates the existing quotaset.
    """
    backups: pulumi.Output[float]
    """
    Quota value for backups. Changing this updates the
    existing quotaset.
    """
    gigabytes: pulumi.Output[float]
    """
    Quota value for gigabytes. Changing this updates the
    existing quotaset.
    """
    groups: pulumi.Output[float]
    """
    Quota value for groups. Changing this updates the
    existing quotaset.
    """
    per_volume_gigabytes: pulumi.Output[float]
    """
    Quota value for gigabytes per volume .
    Changing this updates the existing quotaset.
    """
    project_id: pulumi.Output[str]
    """
    ID of the project to manage quotas. Changing this
    creates a new quotaset.
    """
    region: pulumi.Output[str]
    """
    The region in which to create the volume. If
    omitted, the `region` argument of the provider is used. Changing this
    creates a new quotaset.
    """
    snapshots: pulumi.Output[float]
    """
    Quota value for snapshots. Changing this updates the
    existing quotaset.
    """
    volumes: pulumi.Output[float]
    """
    Quota value for volumes. Changing this updates the
    existing quotaset.
    """
    def __init__(__self__, resource_name, opts=None, backup_gigabytes=None, backups=None, gigabytes=None, groups=None, per_volume_gigabytes=None, project_id=None, region=None, snapshots=None, volumes=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a V2 block storage quotaset resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        > **Note:** This resource has a no-op deletion so no actual actions will be done against the OpenStack API 
            in case of delete call.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/blockstorage_quotaset_v2.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] backup_gigabytes: Quota value for backup gigabytes. Changing
               this updates the existing quotaset.
        :param pulumi.Input[float] backups: Quota value for backups. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] gigabytes: Quota value for gigabytes. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] groups: Quota value for groups. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] per_volume_gigabytes: Quota value for gigabytes per volume .
               Changing this updates the existing quotaset.
        :param pulumi.Input[str] project_id: ID of the project to manage quotas. Changing this
               creates a new quotaset.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
        :param pulumi.Input[float] snapshots: Quota value for snapshots. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] volumes: Quota value for volumes. Changing this updates the
               existing quotaset.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['backup_gigabytes'] = backup_gigabytes
            __props__['backups'] = backups
            __props__['gigabytes'] = gigabytes
            __props__['groups'] = groups
            __props__['per_volume_gigabytes'] = per_volume_gigabytes
            if project_id is None:
                raise TypeError("Missing required property 'project_id'")
            __props__['project_id'] = project_id
            __props__['region'] = region
            __props__['snapshots'] = snapshots
            __props__['volumes'] = volumes
        super(QuoteSetV2, __self__).__init__(
            'openstack:blockstorage/quoteSetV2:QuoteSetV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backup_gigabytes=None, backups=None, gigabytes=None, groups=None, per_volume_gigabytes=None, project_id=None, region=None, snapshots=None, volumes=None):
        """
        Get an existing QuoteSetV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] backup_gigabytes: Quota value for backup gigabytes. Changing
               this updates the existing quotaset.
        :param pulumi.Input[float] backups: Quota value for backups. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] gigabytes: Quota value for gigabytes. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] groups: Quota value for groups. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] per_volume_gigabytes: Quota value for gigabytes per volume .
               Changing this updates the existing quotaset.
        :param pulumi.Input[str] project_id: ID of the project to manage quotas. Changing this
               creates a new quotaset.
        :param pulumi.Input[str] region: The region in which to create the volume. If
               omitted, the `region` argument of the provider is used. Changing this
               creates a new quotaset.
        :param pulumi.Input[float] snapshots: Quota value for snapshots. Changing this updates the
               existing quotaset.
        :param pulumi.Input[float] volumes: Quota value for volumes. Changing this updates the
               existing quotaset.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backup_gigabytes"] = backup_gigabytes
        __props__["backups"] = backups
        __props__["gigabytes"] = gigabytes
        __props__["groups"] = groups
        __props__["per_volume_gigabytes"] = per_volume_gigabytes
        __props__["project_id"] = project_id
        __props__["region"] = region
        __props__["snapshots"] = snapshots
        __props__["volumes"] = volumes
        return QuoteSetV2(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

