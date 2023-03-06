# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ContainerV1Args', 'ContainerV1']

@pulumi.input_type
class ContainerV1Args:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 acl: Optional[pulumi.Input['ContainerV1AclArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]] = None):
        """
        The set of arguments for constructing a ContainerV1 resource.
        :param pulumi.Input[str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        :param pulumi.Input['ContainerV1AclArgs'] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[str] region: The region in which to obtain the V1 KeyManager client.
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

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
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

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable name for the Container. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V1 KeyManager client.
        A KeyManager client is needed to create a container. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        V1 container.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
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
                 container_ref: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 creator_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ContainerV1 resources.
        :param pulumi.Input['ContainerV1AclArgs'] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]] consumers: The list of the container consumers. The structure is described below.
        :param pulumi.Input[str] container_ref: The container reference / where to find the container.
        :param pulumi.Input[str] created_at: The date the container ACL was created.
        :param pulumi.Input[str] creator_id: The creator of the container.
        :param pulumi.Input[str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a container. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 container.
        :param pulumi.Input[Sequence[pulumi.Input['ContainerV1SecretRefArgs']]] secret_refs: A set of dictionaries containing references to secrets. The structure is described
               below.
        :param pulumi.Input[str] status: The status of the container.
        :param pulumi.Input[str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        :param pulumi.Input[str] updated_at: The date the container ACL was last updated.
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

    @property
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

    @property
    @pulumi.getter
    def consumers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]]:
        """
        The list of the container consumers. The structure is described below.
        """
        return pulumi.get(self, "consumers")

    @consumers.setter
    def consumers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ContainerV1ConsumerArgs']]]]):
        pulumi.set(self, "consumers", value)

    @property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> Optional[pulumi.Input[str]]:
        """
        The container reference / where to find the container.
        """
        return pulumi.get(self, "container_ref")

    @container_ref.setter
    def container_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_ref", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date the container ACL was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> Optional[pulumi.Input[str]]:
        """
        The creator of the container.
        """
        return pulumi.get(self, "creator_id")

    @creator_id.setter
    def creator_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creator_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable name for the Container. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V1 KeyManager client.
        A KeyManager client is needed to create a container. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        V1 container.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
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

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the container.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date the container ACL was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class ContainerV1(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl: Optional[pulumi.Input[pulumi.InputType['ContainerV1AclArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a V1 Barbican container resource within OpenStack.

        ## Example Usage
        ### Simple secret

        The container with the TLS certificates, which can be used by the loadbalancer HTTPS listener.

        ```python
        import pulumi
        import pulumi_openstack as openstack

        certificate1 = openstack.keymanager.SecretV1("certificate1",
            payload=(lambda path: open(path).read())("cert.pem"),
            payload_content_type="text/plain",
            secret_type="certificate")
        private_key1 = openstack.keymanager.SecretV1("privateKey1",
            payload=(lambda path: open(path).read())("cert-key.pem"),
            payload_content_type="text/plain",
            secret_type="private")
        intermediate1 = openstack.keymanager.SecretV1("intermediate1",
            payload=(lambda path: open(path).read())("intermediate-ca.pem"),
            payload_content_type="text/plain",
            secret_type="certificate")
        tls1 = openstack.keymanager.ContainerV1("tls1",
            secret_refs=[
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="certificate",
                    secret_ref=certificate1.secret_ref,
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="private_key",
                    secret_ref=private_key1.secret_ref,
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="intermediates",
                    secret_ref=intermediate1.secret_ref,
                ),
            ],
            type="certificate")
        subnet1 = openstack.networking.get_subnet(name="my-subnet")
        lb1 = openstack.loadbalancer.LoadBalancer("lb1", vip_subnet_id=subnet1.id)
        listener1 = openstack.loadbalancer.Listener("listener1",
            default_tls_container_ref=tls1.container_ref,
            loadbalancer_id=lb1.id,
            protocol="TERMINATED_HTTPS",
            protocol_port=443)
        ```
        ### Container with the ACL

        > **Note** Only read ACLs are supported

        ```python
        import pulumi
        import pulumi_openstack as openstack

        tls1 = openstack.keymanager.ContainerV1("tls1",
            acl=openstack.keymanager.ContainerV1AclArgs(
                read=openstack.keymanager.ContainerV1AclReadArgs(
                    project_access=False,
                    users=[
                        "userid1",
                        "userid2",
                    ],
                ),
            ),
            secret_refs=[
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="certificate",
                    secret_ref=openstack_keymanager_secret_v1["certificate_1"]["secret_ref"],
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="private_key",
                    secret_ref=openstack_keymanager_secret_v1["private_key_1"]["secret_ref"],
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="intermediates",
                    secret_ref=openstack_keymanager_secret_v1["intermediate_1"]["secret_ref"],
                ),
            ],
            type="certificate")
        ```

        ## Import

        Containers can be imported using the container id (the last part of the container reference), e.g.

        ```sh
         $ pulumi import openstack:keymanager/containerV1:ContainerV1 container_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ContainerV1AclArgs']] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a container. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 container.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]] secret_refs: A set of dictionaries containing references to secrets. The structure is described
               below.
        :param pulumi.Input[str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
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
        ### Simple secret

        The container with the TLS certificates, which can be used by the loadbalancer HTTPS listener.

        ```python
        import pulumi
        import pulumi_openstack as openstack

        certificate1 = openstack.keymanager.SecretV1("certificate1",
            payload=(lambda path: open(path).read())("cert.pem"),
            payload_content_type="text/plain",
            secret_type="certificate")
        private_key1 = openstack.keymanager.SecretV1("privateKey1",
            payload=(lambda path: open(path).read())("cert-key.pem"),
            payload_content_type="text/plain",
            secret_type="private")
        intermediate1 = openstack.keymanager.SecretV1("intermediate1",
            payload=(lambda path: open(path).read())("intermediate-ca.pem"),
            payload_content_type="text/plain",
            secret_type="certificate")
        tls1 = openstack.keymanager.ContainerV1("tls1",
            secret_refs=[
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="certificate",
                    secret_ref=certificate1.secret_ref,
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="private_key",
                    secret_ref=private_key1.secret_ref,
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="intermediates",
                    secret_ref=intermediate1.secret_ref,
                ),
            ],
            type="certificate")
        subnet1 = openstack.networking.get_subnet(name="my-subnet")
        lb1 = openstack.loadbalancer.LoadBalancer("lb1", vip_subnet_id=subnet1.id)
        listener1 = openstack.loadbalancer.Listener("listener1",
            default_tls_container_ref=tls1.container_ref,
            loadbalancer_id=lb1.id,
            protocol="TERMINATED_HTTPS",
            protocol_port=443)
        ```
        ### Container with the ACL

        > **Note** Only read ACLs are supported

        ```python
        import pulumi
        import pulumi_openstack as openstack

        tls1 = openstack.keymanager.ContainerV1("tls1",
            acl=openstack.keymanager.ContainerV1AclArgs(
                read=openstack.keymanager.ContainerV1AclReadArgs(
                    project_access=False,
                    users=[
                        "userid1",
                        "userid2",
                    ],
                ),
            ),
            secret_refs=[
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="certificate",
                    secret_ref=openstack_keymanager_secret_v1["certificate_1"]["secret_ref"],
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="private_key",
                    secret_ref=openstack_keymanager_secret_v1["private_key_1"]["secret_ref"],
                ),
                openstack.keymanager.ContainerV1SecretRefArgs(
                    name="intermediates",
                    secret_ref=openstack_keymanager_secret_v1["intermediate_1"]["secret_ref"],
                ),
            ],
            type="certificate")
        ```

        ## Import

        Containers can be imported using the container id (the last part of the container reference), e.g.

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
                 acl: Optional[pulumi.Input[pulumi.InputType['ContainerV1AclArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            acl: Optional[pulumi.Input[pulumi.InputType['ContainerV1AclArgs']]] = None,
            consumers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1ConsumerArgs']]]]] = None,
            container_ref: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            creator_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_refs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'ContainerV1':
        """
        Get an existing ContainerV1 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ContainerV1AclArgs']] acl: Allows to control an access to a container. Currently only
               the `read` operation is supported. If not specified, the container is
               accessible project wide. The `read` structure is described below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1ConsumerArgs']]]] consumers: The list of the container consumers. The structure is described below.
        :param pulumi.Input[str] container_ref: The container reference / where to find the container.
        :param pulumi.Input[str] created_at: The date the container ACL was created.
        :param pulumi.Input[str] creator_id: The creator of the container.
        :param pulumi.Input[str] name: Human-readable name for the Container. Does not have
               to be unique.
        :param pulumi.Input[str] region: The region in which to obtain the V1 KeyManager client.
               A KeyManager client is needed to create a container. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               V1 container.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ContainerV1SecretRefArgs']]]] secret_refs: A set of dictionaries containing references to secrets. The structure is described
               below.
        :param pulumi.Input[str] status: The status of the container.
        :param pulumi.Input[str] type: Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        :param pulumi.Input[str] updated_at: The date the container ACL was last updated.
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

    @property
    @pulumi.getter
    def acl(self) -> pulumi.Output['outputs.ContainerV1Acl']:
        """
        Allows to control an access to a container. Currently only
        the `read` operation is supported. If not specified, the container is
        accessible project wide. The `read` structure is described below.
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def consumers(self) -> pulumi.Output[Sequence['outputs.ContainerV1Consumer']]:
        """
        The list of the container consumers. The structure is described below.
        """
        return pulumi.get(self, "consumers")

    @property
    @pulumi.getter(name="containerRef")
    def container_ref(self) -> pulumi.Output[str]:
        """
        The container reference / where to find the container.
        """
        return pulumi.get(self, "container_ref")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date the container ACL was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="creatorId")
    def creator_id(self) -> pulumi.Output[str]:
        """
        The creator of the container.
        """
        return pulumi.get(self, "creator_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Human-readable name for the Container. Does not have
        to be unique.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V1 KeyManager client.
        A KeyManager client is needed to create a container. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        V1 container.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretRefs")
    def secret_refs(self) -> pulumi.Output[Optional[Sequence['outputs.ContainerV1SecretRef']]]:
        """
        A set of dictionaries containing references to secrets. The structure is described
        below.
        """
        return pulumi.get(self, "secret_refs")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the container.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date the container ACL was last updated.
        """
        return pulumi.get(self, "updated_at")

