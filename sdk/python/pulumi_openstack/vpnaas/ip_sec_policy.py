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

__all__ = ['IpSecPolicyArgs', 'IpSecPolicy']

@pulumi.input_type
class IpSecPolicyArgs:
    def __init__(__self__, *,
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encapsulation_mode: Optional[pulumi.Input[str]] = None,
                 encryption_algorithm: Optional[pulumi.Input[str]] = None,
                 lifetimes: Optional[pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pfs: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 transform_protocol: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a IpSecPolicy resource.
        :param pulumi.Input[str] auth_algorithm: The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
               Default is sha1. Changing this updates the algorithm of the existing policy.
        :param pulumi.Input[str] description: The human-readable description for the policy.
               Changing this updates the description of the existing policy.
        :param pulumi.Input[str] encapsulation_mode: The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
               Changing this updates the existing policy.
        :param pulumi.Input[str] encryption_algorithm: The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
               The default value is aes-128. Changing this updates the existing policy.
        :param pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]] lifetimes: The lifetime of the security association. Consists of Unit and Value.
        :param pulumi.Input[str] name: The name of the policy. Changing this updates the name of
               the existing policy.
        :param pulumi.Input[str] pfs: The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
               is group5. Changing this updates the existing policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an IPSec policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               policy.
        :param pulumi.Input[str] tenant_id: The owner of the policy. Required if admin wants to
               create a policy for another project. Changing this creates a new policy.
        :param pulumi.Input[str] transform_protocol: The transform protocol. Valid values are esp, ah and ah-esp.
               Changing this updates the existing policy. Default is ESP.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        if auth_algorithm is not None:
            pulumi.set(__self__, "auth_algorithm", auth_algorithm)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encapsulation_mode is not None:
            pulumi.set(__self__, "encapsulation_mode", encapsulation_mode)
        if encryption_algorithm is not None:
            pulumi.set(__self__, "encryption_algorithm", encryption_algorithm)
        if lifetimes is not None:
            pulumi.set(__self__, "lifetimes", lifetimes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pfs is not None:
            pulumi.set(__self__, "pfs", pfs)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if transform_protocol is not None:
            pulumi.set(__self__, "transform_protocol", transform_protocol)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter(name="authAlgorithm")
    def auth_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
        Default is sha1. Changing this updates the algorithm of the existing policy.
        """
        return pulumi.get(self, "auth_algorithm")

    @auth_algorithm.setter
    def auth_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_algorithm", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable description for the policy.
        Changing this updates the description of the existing policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encapsulationMode")
    def encapsulation_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
        Changing this updates the existing policy.
        """
        return pulumi.get(self, "encapsulation_mode")

    @encapsulation_mode.setter
    def encapsulation_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encapsulation_mode", value)

    @property
    @pulumi.getter(name="encryptionAlgorithm")
    def encryption_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        The default value is aes-128. Changing this updates the existing policy.
        """
        return pulumi.get(self, "encryption_algorithm")

    @encryption_algorithm.setter
    def encryption_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_algorithm", value)

    @property
    @pulumi.getter
    def lifetimes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]]]:
        """
        The lifetime of the security association. Consists of Unit and Value.
        """
        return pulumi.get(self, "lifetimes")

    @lifetimes.setter
    def lifetimes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]]]):
        pulumi.set(self, "lifetimes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the policy. Changing this updates the name of
        the existing policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def pfs(self) -> Optional[pulumi.Input[str]]:
        """
        The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
        is group5. Changing this updates the existing policy.
        """
        return pulumi.get(self, "pfs")

    @pfs.setter
    def pfs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pfs", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an IPSec policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the policy. Required if admin wants to
        create a policy for another project. Changing this creates a new policy.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="transformProtocol")
    def transform_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The transform protocol. Valid values are esp, ah and ah-esp.
        Changing this updates the existing policy. Default is ESP.
        """
        return pulumi.get(self, "transform_protocol")

    @transform_protocol.setter
    def transform_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transform_protocol", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


@pulumi.input_type
class _IpSecPolicyState:
    def __init__(__self__, *,
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encapsulation_mode: Optional[pulumi.Input[str]] = None,
                 encryption_algorithm: Optional[pulumi.Input[str]] = None,
                 lifetimes: Optional[pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pfs: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 transform_protocol: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering IpSecPolicy resources.
        :param pulumi.Input[str] auth_algorithm: The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
               Default is sha1. Changing this updates the algorithm of the existing policy.
        :param pulumi.Input[str] description: The human-readable description for the policy.
               Changing this updates the description of the existing policy.
        :param pulumi.Input[str] encapsulation_mode: The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
               Changing this updates the existing policy.
        :param pulumi.Input[str] encryption_algorithm: The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
               The default value is aes-128. Changing this updates the existing policy.
        :param pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]] lifetimes: The lifetime of the security association. Consists of Unit and Value.
        :param pulumi.Input[str] name: The name of the policy. Changing this updates the name of
               the existing policy.
        :param pulumi.Input[str] pfs: The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
               is group5. Changing this updates the existing policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an IPSec policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               policy.
        :param pulumi.Input[str] tenant_id: The owner of the policy. Required if admin wants to
               create a policy for another project. Changing this creates a new policy.
        :param pulumi.Input[str] transform_protocol: The transform protocol. Valid values are esp, ah and ah-esp.
               Changing this updates the existing policy. Default is ESP.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        if auth_algorithm is not None:
            pulumi.set(__self__, "auth_algorithm", auth_algorithm)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encapsulation_mode is not None:
            pulumi.set(__self__, "encapsulation_mode", encapsulation_mode)
        if encryption_algorithm is not None:
            pulumi.set(__self__, "encryption_algorithm", encryption_algorithm)
        if lifetimes is not None:
            pulumi.set(__self__, "lifetimes", lifetimes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pfs is not None:
            pulumi.set(__self__, "pfs", pfs)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if transform_protocol is not None:
            pulumi.set(__self__, "transform_protocol", transform_protocol)
        if value_specs is not None:
            pulumi.set(__self__, "value_specs", value_specs)

    @property
    @pulumi.getter(name="authAlgorithm")
    def auth_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
        Default is sha1. Changing this updates the algorithm of the existing policy.
        """
        return pulumi.get(self, "auth_algorithm")

    @auth_algorithm.setter
    def auth_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_algorithm", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The human-readable description for the policy.
        Changing this updates the description of the existing policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encapsulationMode")
    def encapsulation_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
        Changing this updates the existing policy.
        """
        return pulumi.get(self, "encapsulation_mode")

    @encapsulation_mode.setter
    def encapsulation_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encapsulation_mode", value)

    @property
    @pulumi.getter(name="encryptionAlgorithm")
    def encryption_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        The default value is aes-128. Changing this updates the existing policy.
        """
        return pulumi.get(self, "encryption_algorithm")

    @encryption_algorithm.setter
    def encryption_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_algorithm", value)

    @property
    @pulumi.getter
    def lifetimes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]]]:
        """
        The lifetime of the security association. Consists of Unit and Value.
        """
        return pulumi.get(self, "lifetimes")

    @lifetimes.setter
    def lifetimes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpSecPolicyLifetimeArgs']]]]):
        pulumi.set(self, "lifetimes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the policy. Changing this updates the name of
        the existing policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def pfs(self) -> Optional[pulumi.Input[str]]:
        """
        The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
        is group5. Changing this updates the existing policy.
        """
        return pulumi.get(self, "pfs")

    @pfs.setter
    def pfs(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pfs", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an IPSec policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the policy. Required if admin wants to
        create a policy for another project. Changing this creates a new policy.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="transformProtocol")
    def transform_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The transform protocol. Valid values are esp, ah and ah-esp.
        Changing this updates the existing policy. Default is ESP.
        """
        return pulumi.get(self, "transform_protocol")

    @transform_protocol.setter
    def transform_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transform_protocol", value)

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

    @value_specs.setter
    def value_specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "value_specs", value)


class IpSecPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encapsulation_mode: Optional[pulumi.Input[str]] = None,
                 encryption_algorithm: Optional[pulumi.Input[str]] = None,
                 lifetimes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpSecPolicyLifetimeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pfs: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 transform_protocol: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Manages a V2 Neutron IPSec policy resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        policy1 = openstack.vpnaas.IpSecPolicy("policy_1", name="my_policy")
        ```

        ## Import

        Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:vpnaas/ipSecPolicy:IpSecPolicy policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_algorithm: The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
               Default is sha1. Changing this updates the algorithm of the existing policy.
        :param pulumi.Input[str] description: The human-readable description for the policy.
               Changing this updates the description of the existing policy.
        :param pulumi.Input[str] encapsulation_mode: The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
               Changing this updates the existing policy.
        :param pulumi.Input[str] encryption_algorithm: The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
               The default value is aes-128. Changing this updates the existing policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpSecPolicyLifetimeArgs']]]] lifetimes: The lifetime of the security association. Consists of Unit and Value.
        :param pulumi.Input[str] name: The name of the policy. Changing this updates the name of
               the existing policy.
        :param pulumi.Input[str] pfs: The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
               is group5. Changing this updates the existing policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an IPSec policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               policy.
        :param pulumi.Input[str] tenant_id: The owner of the policy. Required if admin wants to
               create a policy for another project. Changing this creates a new policy.
        :param pulumi.Input[str] transform_protocol: The transform protocol. Valid values are esp, ah and ah-esp.
               Changing this updates the existing policy. Default is ESP.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IpSecPolicyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V2 Neutron IPSec policy resource within OpenStack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        policy1 = openstack.vpnaas.IpSecPolicy("policy_1", name="my_policy")
        ```

        ## Import

        Policies can be imported using the `id`, e.g.

        ```sh
        $ pulumi import openstack:vpnaas/ipSecPolicy:IpSecPolicy policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272
        ```

        :param str resource_name: The name of the resource.
        :param IpSecPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpSecPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encapsulation_mode: Optional[pulumi.Input[str]] = None,
                 encryption_algorithm: Optional[pulumi.Input[str]] = None,
                 lifetimes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpSecPolicyLifetimeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pfs: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 transform_protocol: Optional[pulumi.Input[str]] = None,
                 value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpSecPolicyArgs.__new__(IpSecPolicyArgs)

            __props__.__dict__["auth_algorithm"] = auth_algorithm
            __props__.__dict__["description"] = description
            __props__.__dict__["encapsulation_mode"] = encapsulation_mode
            __props__.__dict__["encryption_algorithm"] = encryption_algorithm
            __props__.__dict__["lifetimes"] = lifetimes
            __props__.__dict__["name"] = name
            __props__.__dict__["pfs"] = pfs
            __props__.__dict__["region"] = region
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["transform_protocol"] = transform_protocol
            __props__.__dict__["value_specs"] = value_specs
        super(IpSecPolicy, __self__).__init__(
            'openstack:vpnaas/ipSecPolicy:IpSecPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_algorithm: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encapsulation_mode: Optional[pulumi.Input[str]] = None,
            encryption_algorithm: Optional[pulumi.Input[str]] = None,
            lifetimes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpSecPolicyLifetimeArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pfs: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            transform_protocol: Optional[pulumi.Input[str]] = None,
            value_specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'IpSecPolicy':
        """
        Get an existing IpSecPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_algorithm: The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
               Default is sha1. Changing this updates the algorithm of the existing policy.
        :param pulumi.Input[str] description: The human-readable description for the policy.
               Changing this updates the description of the existing policy.
        :param pulumi.Input[str] encapsulation_mode: The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
               Changing this updates the existing policy.
        :param pulumi.Input[str] encryption_algorithm: The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
               The default value is aes-128. Changing this updates the existing policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpSecPolicyLifetimeArgs']]]] lifetimes: The lifetime of the security association. Consists of Unit and Value.
        :param pulumi.Input[str] name: The name of the policy. Changing this updates the name of
               the existing policy.
        :param pulumi.Input[str] pfs: The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
               is group5. Changing this updates the existing policy.
        :param pulumi.Input[str] region: The region in which to obtain the V2 Networking client.
               A Networking client is needed to create an IPSec policy. If omitted, the
               `region` argument of the provider is used. Changing this creates a new
               policy.
        :param pulumi.Input[str] tenant_id: The owner of the policy. Required if admin wants to
               create a policy for another project. Changing this creates a new policy.
        :param pulumi.Input[str] transform_protocol: The transform protocol. Valid values are esp, ah and ah-esp.
               Changing this updates the existing policy. Default is ESP.
        :param pulumi.Input[Mapping[str, Any]] value_specs: Map of additional options.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpSecPolicyState.__new__(_IpSecPolicyState)

        __props__.__dict__["auth_algorithm"] = auth_algorithm
        __props__.__dict__["description"] = description
        __props__.__dict__["encapsulation_mode"] = encapsulation_mode
        __props__.__dict__["encryption_algorithm"] = encryption_algorithm
        __props__.__dict__["lifetimes"] = lifetimes
        __props__.__dict__["name"] = name
        __props__.__dict__["pfs"] = pfs
        __props__.__dict__["region"] = region
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["transform_protocol"] = transform_protocol
        __props__.__dict__["value_specs"] = value_specs
        return IpSecPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authAlgorithm")
    def auth_algorithm(self) -> pulumi.Output[str]:
        """
        The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512.
        Default is sha1. Changing this updates the algorithm of the existing policy.
        """
        return pulumi.get(self, "auth_algorithm")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The human-readable description for the policy.
        Changing this updates the description of the existing policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encapsulationMode")
    def encapsulation_mode(self) -> pulumi.Output[str]:
        """
        The encapsulation mode. Valid values are tunnel and transport. Default is tunnel.
        Changing this updates the existing policy.
        """
        return pulumi.get(self, "encapsulation_mode")

    @property
    @pulumi.getter(name="encryptionAlgorithm")
    def encryption_algorithm(self) -> pulumi.Output[str]:
        """
        The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on.
        The default value is aes-128. Changing this updates the existing policy.
        """
        return pulumi.get(self, "encryption_algorithm")

    @property
    @pulumi.getter
    def lifetimes(self) -> pulumi.Output[Sequence['outputs.IpSecPolicyLifetime']]:
        """
        The lifetime of the security association. Consists of Unit and Value.
        """
        return pulumi.get(self, "lifetimes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy. Changing this updates the name of
        the existing policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pfs(self) -> pulumi.Output[str]:
        """
        The perfect forward secrecy mode. Valid values are group2, group5 and group14. Default 
        is group5. Changing this updates the existing policy.
        """
        return pulumi.get(self, "pfs")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to obtain the V2 Networking client.
        A Networking client is needed to create an IPSec policy. If omitted, the
        `region` argument of the provider is used. Changing this creates a new
        policy.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The owner of the policy. Required if admin wants to
        create a policy for another project. Changing this creates a new policy.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="transformProtocol")
    def transform_protocol(self) -> pulumi.Output[str]:
        """
        The transform protocol. Valid values are esp, ah and ah-esp.
        Changing this updates the existing policy. Default is ESP.
        """
        return pulumi.get(self, "transform_protocol")

    @property
    @pulumi.getter(name="valueSpecs")
    def value_specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Map of additional options.
        """
        return pulumi.get(self, "value_specs")

