# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecretV1(pulumi.CustomResource):
    algorithm: pulumi.Output[str]
    """
    Metadata provided by a user or system for informational purposes.
    """
    all_metadata: pulumi.Output[dict]
    """
    The map of metadata, assigned on the secret, which has been
    explicitly and implicitly added.
    """
    bit_length: pulumi.Output[float]
    """
    Metadata provided by a user or system for informational purposes.
    """
    content_types: pulumi.Output[dict]
    """
    The map of the content types, assigned on the secret.
    """
    created_at: pulumi.Output[str]
    """
    The date the secret was created.
    """
    creator_id: pulumi.Output[str]
    """
    The creator of the secret.
    """
    expiration: pulumi.Output[str]
    """
    The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
    """
    metadata: pulumi.Output[dict]
    """
    Additional Metadata for the secret.
    """
    mode: pulumi.Output[str]
    """
    Metadata provided by a user or system for informational purposes.
    """
    name: pulumi.Output[str]
    """
    Human-readable name for the Secret. Does not have
    to be unique.
    """
    payload: pulumi.Output[str]
    """
    The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
    """
    payload_content_encoding: pulumi.Output[str]
    """
    (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
    """
    payload_content_type: pulumi.Output[str]
    """
    (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
    """
    region: pulumi.Output[str]
    """
    The region in which to obtain the V1 KeyManager client.
    A KeyManager client is needed to create a secret. If omitted, the
    `region` argument of the provider is used. Changing this creates a new
    V1 secret.
    """
    secret_ref: pulumi.Output[str]
    """
    The secret reference / where to find the secret.
    """
    secret_type: pulumi.Output[str]
    """
    Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
    """
    status: pulumi.Output[str]
    """
    The status of the secret.
    """
    updated_at: pulumi.Output[str]
    """
    The date the secret was last updated.
    """
    def __init__(__self__, resource_name, opts=None, algorithm=None, bit_length=None, expiration=None, metadata=None, mode=None, name=None, payload=None, payload_content_encoding=None, payload_content_type=None, region=None, secret_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a SecretV1 resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: Metadata provided by a user or system for informational purposes.
        :param pulumi.Input[float] bit_length: Metadata provided by a user or system for informational purposes.
        :param pulumi.Input[str] expiration: The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
        :param pulumi.Input[dict] metadata: Additional Metadata for the secret.
        :param pulumi.Input[str] mode: Metadata provided by a user or system for informational purposes.
        :param pulumi.Input[str] name: Human-readable name for the Secret. Does not have
               to be unique.
        :param pulumi.Input[str] payload: The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
        :param pulumi.Input[str] payload_content_encoding: (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
        :param pulumi.Input[str] payload_content_type: (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        :param pulumi.Input[str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a secret. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 secret.
        :param pulumi.Input[str] secret_type: Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/keymanager_secret_v1.html.markdown.
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

            __props__['algorithm'] = algorithm
            __props__['bit_length'] = bit_length
            __props__['expiration'] = expiration
            __props__['metadata'] = metadata
            __props__['mode'] = mode
            __props__['name'] = name
            __props__['payload'] = payload
            __props__['payload_content_encoding'] = payload_content_encoding
            __props__['payload_content_type'] = payload_content_type
            __props__['region'] = region
            __props__['secret_type'] = secret_type
            __props__['all_metadata'] = None
            __props__['content_types'] = None
            __props__['created_at'] = None
            __props__['creator_id'] = None
            __props__['secret_ref'] = None
            __props__['status'] = None
            __props__['updated_at'] = None
        super(SecretV1, __self__).__init__(
            'openstack:keymanager/secretV1:SecretV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, algorithm=None, all_metadata=None, bit_length=None, content_types=None, created_at=None, creator_id=None, expiration=None, metadata=None, mode=None, name=None, payload=None, payload_content_encoding=None, payload_content_type=None, region=None, secret_ref=None, secret_type=None, status=None, updated_at=None):
        """
        Get an existing SecretV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: Metadata provided by a user or system for informational purposes.
        :param pulumi.Input[dict] all_metadata: The map of metadata, assigned on the secret, which has been
               explicitly and implicitly added.
        :param pulumi.Input[float] bit_length: Metadata provided by a user or system for informational purposes.
        :param pulumi.Input[dict] content_types: The map of the content types, assigned on the secret.
        :param pulumi.Input[str] created_at: The date the secret was created.
        :param pulumi.Input[str] creator_id: The creator of the secret.
        :param pulumi.Input[str] expiration: The expiration time of the secret in the RFC3339 timestamp format (e.g. `2019-03-09T12:58:49Z`). If omitted, a secret will never expire. Changing this creates a new secret.
        :param pulumi.Input[dict] metadata: Additional Metadata for the secret.
        :param pulumi.Input[str] mode: Metadata provided by a user or system for informational purposes.
        :param pulumi.Input[str] name: Human-readable name for the Secret. Does not have
               to be unique.
        :param pulumi.Input[str] payload: The secret's data to be stored. **payload\_content\_type** must also be supplied if **payload** is included.
        :param pulumi.Input[str] payload_content_encoding: (required if **payload** is encoded) The encoding used for the payload to be able to include it in the JSON request. Must be either `base64` or `binary`.
        :param pulumi.Input[str] payload_content_type: (required if **payload** is included) The media type for the content of the payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
        :param pulumi.Input[str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a secret. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 secret.
        :param pulumi.Input[str] secret_ref: The secret reference / where to find the secret.
        :param pulumi.Input[str] secret_type: Used to indicate the type of secret being stored. For more information see [Secret types](https://docs.openstack.org/barbican/latest/api/reference/secret_types.html).
        :param pulumi.Input[str] status: The status of the secret.
        :param pulumi.Input[str] updated_at: The date the secret was last updated.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/keymanager_secret_v1.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["algorithm"] = algorithm
        __props__["all_metadata"] = all_metadata
        __props__["bit_length"] = bit_length
        __props__["content_types"] = content_types
        __props__["created_at"] = created_at
        __props__["creator_id"] = creator_id
        __props__["expiration"] = expiration
        __props__["metadata"] = metadata
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["payload"] = payload
        __props__["payload_content_encoding"] = payload_content_encoding
        __props__["payload_content_type"] = payload_content_type
        __props__["region"] = region
        __props__["secret_ref"] = secret_ref
        __props__["secret_type"] = secret_type
        __props__["status"] = status
        __props__["updated_at"] = updated_at
        return SecretV1(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
