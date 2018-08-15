# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class FloatingIp(pulumi.CustomResource):
    """
    Manages a V2 floating IP resource within OpenStack Nova (compute)
    that can be used for compute instances.
    
    Please note that managing floating IPs through the OpenStack Compute API has
    been deprecated. Unless you are using an older OpenStack environment, it is
    recommended to use the [`openstack_networking_floatingip_v2`](networking_floatingip_v2.html)
    resource instead, which uses the OpenStack Networking API.
    """
    def __init__(__self__, __name__, __opts__=None, pool=None, region=None):
        """Create a FloatingIp resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not pool:
            raise TypeError('Missing required property pool')
        elif not isinstance(pool, basestring):
            raise TypeError('Expected property pool to be a basestring')
        __self__.pool = pool
        """
        The name of the pool from which to obtain the floating
        IP. Changing this creates a new floating IP.
        """
        __props__['pool'] = pool

        if region and not isinstance(region, basestring):
            raise TypeError('Expected property region to be a basestring')
        __self__.region = region
        """
        The region in which to obtain the V2 Compute client.
        A Compute client is needed to create a floating IP that can be used with
        a compute instance. If omitted, the `region` argument of the provider
        is used. Changing this creates a new floating IP (which may or may not
        have a different address).
        """
        __props__['region'] = region

        __self__.address = pulumi.runtime.UNKNOWN
        """
        The actual floating IP address itself.
        """
        __self__.fixed_ip = pulumi.runtime.UNKNOWN
        """
        The fixed IP address corresponding to the floating IP.
        """
        __self__.instance_id = pulumi.runtime.UNKNOWN
        """
        UUID of the compute instance associated with the floating IP.
        """

        super(FloatingIp, __self__).__init__(
            'openstack:compute/floatingIp:FloatingIp',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'address' in outs:
            self.address = outs['address']
        if 'fixedIp' in outs:
            self.fixed_ip = outs['fixedIp']
        if 'instanceId' in outs:
            self.instance_id = outs['instanceId']
        if 'pool' in outs:
            self.pool = outs['pool']
        if 'region' in outs:
            self.region = outs['region']
