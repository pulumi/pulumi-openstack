# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterKubeconfigArgs',
]

@pulumi.input_type
class ClusterKubeconfigArgs:
    def __init__(__self__, *,
                 client_certificate: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 cluster_ca_certificate: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 raw_config: Optional[pulumi.Input[str]] = None):
        if client_certificate is not None:
            pulumi.set(__self__, "client_certificate", client_certificate)
        if client_key is not None:
            pulumi.set(__self__, "client_key", client_key)
        if cluster_ca_certificate is not None:
            pulumi.set(__self__, "cluster_ca_certificate", cluster_ca_certificate)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if raw_config is not None:
            pulumi.set(__self__, "raw_config", raw_config)

    @property
    @pulumi.getter(name="clientCertificate")
    def client_certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_certificate")

    @client_certificate.setter
    def client_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_certificate", value)

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_key")

    @client_key.setter
    def client_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key", value)

    @property
    @pulumi.getter(name="clusterCaCertificate")
    def cluster_ca_certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cluster_ca_certificate")

    @cluster_ca_certificate.setter
    def cluster_ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_ca_certificate", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="rawConfig")
    def raw_config(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "raw_config")

    @raw_config.setter
    def raw_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "raw_config", value)


