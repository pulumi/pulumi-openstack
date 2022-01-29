# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['QosV3Args', 'QosV3']

@pulumi.input_type
class QosV3Args:
    def __init__(__self__, *,
                 consumer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a QosV3 resource.
        :param pulumi.Input[str] consumer: The consumer of qos. Can be one of `front-end`,
               `back-end` or `both`. Changing this updates the `consumer` of an
               existing qos.
        :param pulumi.Input[str] name: Name of the qos.  Changing this creates a new qos.
        :param pulumi.Input[str] region: The region in which to create the qos. If omitted,
               the `region` argument of the provider is used. Changing this creates
               a new qos.
        :param pulumi.Input[Mapping[str, Any]] specs: Key/Value pairs of specs for the qos.
        """
        if consumer is not None:
            pulumi.set(__self__, "consumer", consumer)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if specs is not None:
            pulumi.set(__self__, "specs", specs)

    @property
    @pulumi.getter
    def consumer(self) -> Optional[pulumi.Input[str]]:
        """
        The consumer of qos. Can be one of `front-end`,
        `back-end` or `both`. Changing this updates the `consumer` of an
        existing qos.
        """
        return pulumi.get(self, "consumer")

    @consumer.setter
    def consumer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the qos.  Changing this creates a new qos.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the qos. If omitted,
        the `region` argument of the provider is used. Changing this creates
        a new qos.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/Value pairs of specs for the qos.
        """
        return pulumi.get(self, "specs")

    @specs.setter
    def specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "specs", value)


@pulumi.input_type
class _QosV3State:
    def __init__(__self__, *,
                 consumer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 specs: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering QosV3 resources.
        :param pulumi.Input[str] consumer: The consumer of qos. Can be one of `front-end`,
               `back-end` or `both`. Changing this updates the `consumer` of an
               existing qos.
        :param pulumi.Input[str] name: Name of the qos.  Changing this creates a new qos.
        :param pulumi.Input[str] region: The region in which to create the qos. If omitted,
               the `region` argument of the provider is used. Changing this creates
               a new qos.
        :param pulumi.Input[Mapping[str, Any]] specs: Key/Value pairs of specs for the qos.
        """
        if consumer is not None:
            pulumi.set(__self__, "consumer", consumer)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if specs is not None:
            pulumi.set(__self__, "specs", specs)

    @property
    @pulumi.getter
    def consumer(self) -> Optional[pulumi.Input[str]]:
        """
        The consumer of qos. Can be one of `front-end`,
        `back-end` or `both`. Changing this updates the `consumer` of an
        existing qos.
        """
        return pulumi.get(self, "consumer")

    @consumer.setter
    def consumer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the qos.  Changing this creates a new qos.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the qos. If omitted,
        the `region` argument of the provider is used. Changing this creates
        a new qos.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def specs(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/Value pairs of specs for the qos.
        """
        return pulumi.get(self, "specs")

    @specs.setter
    def specs(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "specs", value)


class QosV3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        Manages a V3 block storage Quality-Of-Servirce (qos) resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos = openstack.blockstorage.QosV3("qos",
            consumer="back-end",
            specs={
                "read_iops_sec": "40000",
                "write_iops_sec": "40000",
            })
        ```

        ## Import

        Qos can be imported using the `qos_id`, e.g.

        ```sh
         $ pulumi import openstack:blockstorage/qosV3:QosV3 qos 941793f0-0a34-4bc4-b72e-a6326ae58283
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer: The consumer of qos. Can be one of `front-end`,
               `back-end` or `both`. Changing this updates the `consumer` of an
               existing qos.
        :param pulumi.Input[str] name: Name of the qos.  Changing this creates a new qos.
        :param pulumi.Input[str] region: The region in which to create the qos. If omitted,
               the `region` argument of the provider is used. Changing this creates
               a new qos.
        :param pulumi.Input[Mapping[str, Any]] specs: Key/Value pairs of specs for the qos.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[QosV3Args] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a V3 block storage Quality-Of-Servirce (qos) resource within OpenStack.

        > **Note:** This usually requires admin privileges.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_openstack as openstack

        qos = openstack.blockstorage.QosV3("qos",
            consumer="back-end",
            specs={
                "read_iops_sec": "40000",
                "write_iops_sec": "40000",
            })
        ```

        ## Import

        Qos can be imported using the `qos_id`, e.g.

        ```sh
         $ pulumi import openstack:blockstorage/qosV3:QosV3 qos 941793f0-0a34-4bc4-b72e-a6326ae58283
        ```

        :param str resource_name: The name of the resource.
        :param QosV3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QosV3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 specs: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QosV3Args.__new__(QosV3Args)

            __props__.__dict__["consumer"] = consumer
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["specs"] = specs
        super(QosV3, __self__).__init__(
            'openstack:blockstorage/qosV3:QosV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            consumer: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            specs: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'QosV3':
        """
        Get an existing QosV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer: The consumer of qos. Can be one of `front-end`,
               `back-end` or `both`. Changing this updates the `consumer` of an
               existing qos.
        :param pulumi.Input[str] name: Name of the qos.  Changing this creates a new qos.
        :param pulumi.Input[str] region: The region in which to create the qos. If omitted,
               the `region` argument of the provider is used. Changing this creates
               a new qos.
        :param pulumi.Input[Mapping[str, Any]] specs: Key/Value pairs of specs for the qos.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QosV3State.__new__(_QosV3State)

        __props__.__dict__["consumer"] = consumer
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["specs"] = specs
        return QosV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def consumer(self) -> pulumi.Output[Optional[str]]:
        """
        The consumer of qos. Can be one of `front-end`,
        `back-end` or `both`. Changing this updates the `consumer` of an
        existing qos.
        """
        return pulumi.get(self, "consumer")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the qos.  Changing this creates a new qos.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the qos. If omitted,
        the `region` argument of the provider is used. Changing this creates
        a new qos.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def specs(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Key/Value pairs of specs for the qos.
        """
        return pulumi.get(self, "specs")
