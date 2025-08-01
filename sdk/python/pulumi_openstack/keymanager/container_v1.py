# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ContainerV1Args', 'ContainerV1']

@pulumi.input_type
class ContainerV1Args:
    def __init__(__self__, *,
                 type: pulumi.Input[_builtins.str],
                 acl: Optional[pulumi.Input['ContainerV1AclArgs']] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]] = None):
        """
        The set of arguments for constructing a ContainerV1 resource.
        :param pulumi.Input[_builtins.str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        :param pulumi.Input['ContainerV1AclArgs'] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[_builtins.str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a container. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 container.
        :param pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]] secret_refs: A set of dictionaries containing references to secrets. The structure is described
               below.
        """
        pulumi.set(__self__, "type", type)
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_refs is not None:
            pulumi.set(__self__, "secret_refs", secret_refs)

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Input[_builtins.str]:
        """
        Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ContainerV1AclArgs']]:
        """
        Allows to control an access to a container. Currently only
        the `read` operation is supported. If not specified, the container is
        accessible project wide. The `read` structure is described below.
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ContainerV1AclArgs']]):
        pulumi.set(self, "acl", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Human-readable name for the Container. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The region in which to obtain the V1 KeyManager client.
        A KeyManager client is needed to create a container. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        V1 container.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "region", value)

    @_builtins.property
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]:
        """
        A set of dictionaries containing references to secrets. The structure is described
        below.
        """
        return pulumi.get(self, "secret_refs")

    @secret_refs.setter
    def secret_refs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]):
        pulumi.set(self, "secret_refs", value)


@pulumi.input_type
class _ContainerV1State:
    def __init__(__self__, *,
                 acl: Optional[pulumi.Input['ContainerV1AclArgs']] = None,
                 consumers: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]] = None,
                 container_ref: Optional[pulumi.Input[_builtins.str]] = None,
                 created_at: Optional[pulumi.Input[_builtins.str]] = None,
                 creator_id: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]] = None,
                 status: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering ContainerV1 resources.
        :param pulumi.Input['ContainerV1AclArgs'] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]] consumers: The list of the container consumers. The structure is described below.
        :param pulumi.Input[_builtins.str] container_ref: The container reference / where to find the container.
        :param pulumi.Input[_builtins.str] created_at: The date the container was created.
        :param pulumi.Input[_builtins.str] creator_id: The creator of the container.
        :param pulumi.Input[_builtins.str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a container. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 container.
        :param pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]] secret_refs: A set of dictionaries containing references to secrets. The structure is described
               below.
        :param pulumi.Input[_builtins.str] status: The status of the container.
        :param pulumi.Input[_builtins.str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        :param pulumi.Input[_builtins.str] updated_at: The date the container was last updated.
        """
        if acl is not None:
            pulumi.set(__self__, "acl", acl)
        if consumers is not None:
            pulumi.set(__self__, "consumers", consumers)
        if container_ref is not None:
            pulumi.set(__self__, "container_ref", container_ref)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if creator_id is not None:
            pulumi.set(__self__, "creator_id", creator_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_refs is not None:
            pulumi.set(__self__, "secret_refs", secret_refs)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> Optional[pulumi.Input['ContainerV1AclArgs']]:
        """
        Allows to control an access to a container. Currently only
        the `read` operation is supported. If not specified, the container is
        accessible project wide. The `read` structure is described below.
        """
        return pulumi.get(self, "acl")

    @acl.setter
    def acl(self, value: Optional[pulumi.Input['ContainerV1AclArgs']]):
        pulumi.set(self, "acl", value)

    @_builtins.property
    @pulumi.getter
    def consumers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]]:
        """
        The list of the container consumers. The structure is described below.
        """
        return pulumi.get(self, "consumers")

    @consumers.setter
    def consumers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]]):
        pulumi.set(self, "consumers", value)

    @_builtins.property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The container reference / where to find the container.
        """
        return pulumi.get(self, "container_ref")

    @container_ref.setter
    def container_ref(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "container_ref", value)

    @_builtins.property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The date the container was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "created_at", value)

    @_builtins.property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The creator of the container.
        """
        return pulumi.get(self, "creator_id")

    @creator_id.setter
    def creator_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "creator_id", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Human-readable name for the Container. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The region in which to obtain the V1 KeyManager client.
        A KeyManager client is needed to create a container. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        V1 container.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "region", value)

    @_builtins.property
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]:
        """
        A set of dictionaries containing references to secrets. The structure is described
        below.
        """
        return pulumi.get(self, "secret_refs")

    @secret_refs.setter
    def secret_refs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]]):
        pulumi.set(self, "secret_refs", value)

    @_builtins.property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The status of the container.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "status", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The date the container was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "updated_at", value)


@pulumi.type_token("openstack:keymanager/containerV1:ContainerV1")
class ContainerV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['ContainerV1AclArgs', 'ContainerV1AclArgsDict']]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ContainerV1SecretRefArgs', 'ContainerV1SecretRefArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Manages a V1 Barbican container resource within OpenStack.

        ## Example Usage

        ### Simple container

        A container with the TLS certificates.

        ```python
        import pulumi
        import pulumi_openstack as openstack
        import pulumi_std as std

        certificate1 = openstack.keymanager.SecretV1("certificate_1",
            name="certificate",
            payload=std.file(input="cert.pem").result,
            secret_type="certificate",
            payload_content_type="text/plain")
        private_key1 = openstack.keymanager.SecretV1("private_key_1",
            name="private_key",
            payload=std.file(input="cert-key.pem").result,
            secret_type="private",
            payload_content_type="text/plain")
        intermediate1 = openstack.keymanager.SecretV1("intermediate_1",
            name="intermediate",
            payload=std.file(input="intermediate-ca.pem").result,
            secret_type="certificate",
            payload_content_type="text/plain")
        tls1 = openstack.keymanager.ContainerV1("tls_1",
            name="tls",
            type="certificate",
            secret_refs=[
                {
                    "name": "certificate",
                    "secret_ref": certificate1.secret_ref,
                },
                {
                    "name": "private_key",
                    "secret_ref": private_key1.secret_ref,
                },
                {
                    "name": "intermediates",
                    "secret_ref": intermediate1.secret_ref,
                },
            ])
        ```

        ### Container with the ACL

        > **Note** Only read ACLs are supported

        ```python
        import pulumi
        import pulumi_openstack as openstack

        tls1 = openstack.keymanager.ContainerV1("tls_1",
            name="tls",
            type="certificate",
            secret_refs=[
                {
                    "name": "certificate",
                    "secret_ref": certificate1["secretRef"],
                },
                {
                    "name": "private_key",
                    "secret_ref": private_key1["secretRef"],
                },
                {
                    "name": "intermediates",
                    "secret_ref": intermediate1["secretRef"],
                },
            ],
            acl={
                "read": {
                    "project_access": False,
                    "users": [
                        "userid1",
                        "userid2",
                    ],
                },
            })
        ```

        ## Import

        Containers can be imported using the container id (the last part of the container reference), e.g.:

        ```sh
        $ pulumi import openstack:keymanager/containerV1:ContainerV1 container_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ContainerV1AclArgs', 'ContainerV1AclArgsDict']] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[_builtins.str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a container. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 container.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ContainerV1SecretRefArgs', 'ContainerV1SecretRefArgsDict']]]] secret_refs: A set of dictionaries containing references to secrets. The structure is described
               below.
        :param pulumi.Input[_builtins.str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ContainerV1Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V1 Barbican container resource within OpenStack.

        ## Example Usage

        ### Simple container

        A container with the TLS certificates.

        ```python
        import pulumi
        import pulumi_openstack as openstack
        import pulumi_std as std

        certificate1 = openstack.keymanager.SecretV1("certificate_1",
            name="certificate",
            payload=std.file(input="cert.pem").result,
            secret_type="certificate",
            payload_content_type="text/plain")
        private_key1 = openstack.keymanager.SecretV1("private_key_1",
            name="private_key",
            payload=std.file(input="cert-key.pem").result,
            secret_type="private",
            payload_content_type="text/plain")
        intermediate1 = openstack.keymanager.SecretV1("intermediate_1",
            name="intermediate",
            payload=std.file(input="intermediate-ca.pem").result,
            secret_type="certificate",
            payload_content_type="text/plain")
        tls1 = openstack.keymanager.ContainerV1("tls_1",
            name="tls",
            type="certificate",
            secret_refs=[
                {
                    "name": "certificate",
                    "secret_ref": certificate1.secret_ref,
                },
                {
                    "name": "private_key",
                    "secret_ref": private_key1.secret_ref,
                },
                {
                    "name": "intermediates",
                    "secret_ref": intermediate1.secret_ref,
                },
            ])
        ```

        ### Container with the ACL

        > **Note** Only read ACLs are supported

        ```python
        import pulumi
        import pulumi_openstack as openstack

        tls1 = openstack.keymanager.ContainerV1("tls_1",
            name="tls",
            type="certificate",
            secret_refs=[
                {
                    "name": "certificate",
                    "secret_ref": certificate1["secretRef"],
                },
                {
                    "name": "private_key",
                    "secret_ref": private_key1["secretRef"],
                },
                {
                    "name": "intermediates",
                    "secret_ref": intermediate1["secretRef"],
                },
            ],
            acl={
                "read": {
                    "project_access": False,
                    "users": [
                        "userid1",
                        "userid2",
                    ],
                },
            })
        ```

        ## Import

        Containers can be imported using the container id (the last part of the container reference), e.g.:

        ```sh
        $ pulumi import openstack:keymanager/containerV1:ContainerV1 container_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
        ```

        :param str resource_name: The name of the resource.
        :param ContainerV1Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContainerV1Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[Union['ContainerV1AclArgs', 'ContainerV1AclArgsDict']]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 region: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ContainerV1SecretRefArgs', 'ContainerV1SecretRefArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ContainerV1Args.__new__(ContainerV1Args)

            __props__.__dict__["acl"] = acl
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["secret_refs"] = secret_refs
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["consumers"] = None
            __props__.__dict__["container_ref"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["creator_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(ContainerV1, __self__).__init__(
            'openstack:keymanager/containerV1:ContainerV1',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl: Optional[pulumi.Input[Union['ContainerV1AclArgs', 'ContainerV1AclArgsDict']]] = None,
            consumers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ContainerV1ConsumerArgs', 'ContainerV1ConsumerArgsDict']]]]] = None,
            container_ref: Optional[pulumi.Input[_builtins.str]] = None,
            created_at: Optional[pulumi.Input[_builtins.str]] = None,
            creator_id: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            region: Optional[pulumi.Input[_builtins.str]] = None,
            secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ContainerV1SecretRefArgs', 'ContainerV1SecretRefArgsDict']]]]] = None,
            status: Optional[pulumi.Input[_builtins.str]] = None,
            type: Optional[pulumi.Input[_builtins.str]] = None,
            updated_at: Optional[pulumi.Input[_builtins.str]] = None) -> 'ContainerV1':
        """
        Get an existing ContainerV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['ContainerV1AclArgs', 'ContainerV1AclArgsDict']] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ContainerV1ConsumerArgs', 'ContainerV1ConsumerArgsDict']]]] consumers: The list of the container consumers. The structure is described below.
        :param pulumi.Input[_builtins.str] container_ref: The container reference / where to find the container.
        :param pulumi.Input[_builtins.str] created_at: The date the container was created.
        :param pulumi.Input[_builtins.str] creator_id: The creator of the container.
        :param pulumi.Input[_builtins.str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[_builtins.str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a container. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 container.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ContainerV1SecretRefArgs', 'ContainerV1SecretRefArgsDict']]]] secret_refs: A set of dictionaries containing references to secrets. The structure is described
               below.
        :param pulumi.Input[_builtins.str] status: The status of the container.
        :param pulumi.Input[_builtins.str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        :param pulumi.Input[_builtins.str] updated_at: The date the container was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContainerV1State.__new__(_ContainerV1State)

        __props__.__dict__["acl"] = acl
        __props__.__dict__["consumers"] = consumers
        __props__.__dict__["container_ref"] = container_ref
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["creator_id"] = creator_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_refs"] = secret_refs
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        return ContainerV1(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.ContainerV1Acl']:
        """
        Allows to control an access to a container. Currently only
        the `read` operation is supported. If not specified, the container is
        accessible project wide. The `read` structure is described below.
        """
        return pulumi.get(self, "acl")

    @_builtins.property
    @pulumi.getter
    def consumers(self) -> pulumi.Output[Sequence['outputs.ContainerV1Consumer']]:
        """
        The list of the container consumers. The structure is described below.
        """
        return pulumi.get(self, "consumers")

    @_builtins.property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> pulumi.Output[_builtins.str]:
        """
        The container reference / where to find the container.
        """
        return pulumi.get(self, "container_ref")

    @_builtins.property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[_builtins.str]:
        """
        The date the container was created.
        """
        return pulumi.get(self, "created_at")

    @_builtins.property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> pulumi.Output[_builtins.str]:
        """
        The creator of the container.
        """
        return pulumi.get(self, "creator_id")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Human-readable name for the Container. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def region(self) -> pulumi.Output[_builtins.str]:
        """
        The region in which to obtain the V1 KeyManager client.
        A KeyManager client is needed to create a container. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        V1 container.
        """
        return pulumi.get(self, "region")

    @_builtins.property
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> pulumi.Output[Optional[Sequence['outputs.ContainerV1SecretRef']]]:
        """
        A set of dictionaries containing references to secrets. The structure is described
        below.
        """
        return pulumi.get(self, "secret_refs")

    @_builtins.property
    @pulumi.getter
    def status(self) -> pulumi.Output[_builtins.str]:
        """
        The status of the container.
        """
        return pulumi.get(self, "status")

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Output[_builtins.str]:
        """
        Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        """
        return pulumi.get(self, "type")

    @_builtins.property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[_builtins.str]:
        """
        The date the container was last updated.
        """
        return pulumi.get(self, "updated_at")

