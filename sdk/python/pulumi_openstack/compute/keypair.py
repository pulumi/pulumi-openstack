# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeypairArgs', 'Keypair']

@pulumi.input_type
class KeypairArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Keypair resource.
        :param pulumi.Input[str] name: A unique name for the keypair. Changing this creates a new
               keypair.
        :param pulumi.Input[str] public_key: A pregenerated OpenSSH-formatted public key.
               Changing this creates a new keypair. If a public key is not specified, then
               a public/private key pair will be automatically generated. If a pair is
               created, then destroying this resource means you will lose access to that
               keypair forever.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new keypair.
        :param pulumi.Input[str] user_id: This allows administrative users to operate key-pairs
               of specified user ID. For this feature your need to have openstack microversion
               2.10 (Liberty) or later.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] value_specs: Map of additional options.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the keypair. Changing this creates a new
        keypair.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        A pregenerated OpenSSH-formatted public key.
        Changing this creates a new keypair. If a public key is not specified, then
        a public/private key pair will be automatically generated. If a pair is
        created, then destroying this resource means you will lose access to that
        keypair forever.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new keypair.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        This allows administrative users to operate key-pairs
        of specified user ID. For this feature your need to have openstack microversion
        2.10 (Liberty) or later.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _KeypairState:
    def __init__(__self__, *,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Keypair resources.
        :param pulumi.Input[str] fingerprint: The fingerprint of the public key.
        :param pulumi.Input[str] name: A unique name for the keypair. Changing this creates a new
               keypair.
        :param pulumi.Input[str] private_key: The generated private key when no public key is specified.
        :param pulumi.Input[str] public_key: A pregenerated OpenSSH-formatted public key.
               Changing this creates a new keypair. If a public key is not specified, then
               a public/private key pair will be automatically generated. If a pair is
               created, then destroying this resource means you will lose access to that
               keypair forever.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new keypair.
        :param pulumi.Input[str] user_id: This allows administrative users to operate key-pairs
               of specified user ID. For this feature your need to have openstack microversion
               2.10 (Liberty) or later.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] value_specs: Map of additional options.
        """
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        The fingerprint of the public key.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the keypair. Changing this creates a new
        keypair.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The generated private key when no public key is specified.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        A pregenerated OpenSSH-formatted public key.
        Changing this creates a new keypair. If a public key is not specified, then
        a public/private key pair will be automatically generated. If a pair is
        created, then destroying this resource means you will lose access to that
        keypair forever.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new keypair.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        This allows administrative users to operate key-pairs
        of specified user ID. For this feature your need to have openstack microversion
        2.10 (Liberty) or later.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "value_specs", value)


class Keypair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Import an Existing Public Key

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_keypair = openstack.compute.Keypair("test-keypair",
            name="my-keypair",
            public_key="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLotBCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAnOfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZqd9LvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TaIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIF61p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB")
        ```

        ### Generate a Public/Private Key Pair

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_keypair = openstack.compute.Keypair("test-keypair", name="my-keypair")
        ```

        ## Import

        Keypairs can be imported using the `name`, e.g.

        ```sh
        $ pulumi import openstack:compute/keypair:Keypair my-keypair test-keypair
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A unique name for the keypair. Changing this creates a new
               keypair.
        :param pulumi.Input[str] public_key: A pregenerated OpenSSH-formatted public key.
               Changing this creates a new keypair. If a public key is not specified, then
               a public/private key pair will be automatically generated. If a pair is
               created, then destroying this resource means you will lose access to that
               keypair forever.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new keypair.
        :param pulumi.Input[str] user_id: This allows administrative users to operate key-pairs
               of specified user ID. For this feature your need to have openstack microversion
               2.10 (Liberty) or later.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[KeypairArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Import an Existing Public Key

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_keypair = openstack.compute.Keypair("test-keypair",
            name="my-keypair",
            public_key="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLotBCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAnOfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZqd9LvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TaIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIF61p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB")
        ```

        ### Generate a Public/Private Key Pair

        ```python
        import pulumi
        import pulumi_openstack as openstack

        test_keypair = openstack.compute.Keypair("test-keypair", name="my-keypair")
        ```

        ## Import

        Keypairs can be imported using the `name`, e.g.

        ```sh
        $ pulumi import openstack:compute/keypair:Keypair my-keypair test-keypair
        ```

        :param str resource_name: The name of the resource.
        :param KeypairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeypairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeypairArgs.__new__(KeypairArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["region"] = region
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["value_specs"] = value_specs
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["private_key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Keypair, __self__).__init__(
            'openstack:compute/keypair:Keypair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Keypair':
        """
        Get an existing Keypair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fingerprint: The fingerprint of the public key.
        :param pulumi.Input[str] name: A unique name for the keypair. Changing this creates a new
               keypair.
        :param pulumi.Input[str] private_key: The generated private key when no public key is specified.
        :param pulumi.Input[str] public_key: A pregenerated OpenSSH-formatted public key.
               Changing this creates a new keypair. If a public key is not specified, then
               a public/private key pair will be automatically generated. If a pair is
               created, then destroying this resource means you will lose access to that
               keypair forever.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Compute client.
               Keypairs are associated with accounts, but a Compute client is needed to
               create one. If omitted, the `region` argument of the provider is used.
               Changing this creates a new keypair.
        :param pulumi.Input[str] user_id: This allows administrative users to operate key-pairs
               of specified user ID. For this feature your need to have openstack microversion
               2.10 (Liberty) or later.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeypairState.__new__(_KeypairState)

        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["name"] = name
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["region"] = region
        __props__.__dict__["user_id"] = user_id
        __props__.__dict__["value_specs"] = value_specs
        return Keypair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint of the public key.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the keypair. Changing this creates a new
        keypair.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        The generated private key when no public key is specified.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        A pregenerated OpenSSH-formatted public key.
        Changing this creates a new keypair. If a public key is not specified, then
        a public/private key pair will be automatically generated. If a pair is
        created, then destroying this resource means you will lose access to that
        keypair forever.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Compute client.
        Keypairs are associated with accounts, but a Compute client is needed to
        create one. If omitted, the `region` argument of the provider is used.
        Changing this creates a new keypair.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        This allows administrative users to operate key-pairs
        of specified user ID. For this feature your need to have openstack microversion
        2.10 (Liberty) or later.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

