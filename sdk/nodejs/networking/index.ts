// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./addressScope";
export * from "./floatingIp";
export * from "./floatingIpAssociate";
export * from "./getAddressScope";
export * from "./getFloatingIp";
export * from "./getNetwork";
export * from "./getPort";
export * from "./getPortIds";
export * from "./getQosBandwidthLimitRule";
export * from "./getQosDscpMarkingRule";
export * from "./getQosMinimumBandwidthRule";
export * from "./getQosPolicy";
export * from "./getRouter";
export * from "./getSecGroup";
export * from "./getSubnet";
export * from "./getSubnetPool";
export * from "./getTrunk";
export * from "./network";
export * from "./port";
export * from "./portSecGroupAssociate";
export * from "./qosBandwidthLimitRule";
export * from "./qosDscpMarkingRule";
export * from "./qosMinimumBandwidthRule";
export * from "./qosPolicy";
export * from "./quotaV2";
export * from "./rbacPolicyV2";
export * from "./router";
export * from "./routerInterface";
export * from "./routerRoute";
export * from "./secGroup";
export * from "./secGroupRule";
export * from "./subnet";
export * from "./subnetPool";
export * from "./subnetRoute";
export * from "./trunk";

// Import resources to register:
import { AddressScope } from "./addressScope";
import { FloatingIp } from "./floatingIp";
import { FloatingIpAssociate } from "./floatingIpAssociate";
import { Network } from "./network";
import { Port } from "./port";
import { PortSecGroupAssociate } from "./portSecGroupAssociate";
import { QosBandwidthLimitRule } from "./qosBandwidthLimitRule";
import { QosDscpMarkingRule } from "./qosDscpMarkingRule";
import { QosMinimumBandwidthRule } from "./qosMinimumBandwidthRule";
import { QosPolicy } from "./qosPolicy";
import { QuotaV2 } from "./quotaV2";
import { RbacPolicyV2 } from "./rbacPolicyV2";
import { Router } from "./router";
import { RouterInterface } from "./routerInterface";
import { RouterRoute } from "./routerRoute";
import { SecGroup } from "./secGroup";
import { SecGroupRule } from "./secGroupRule";
import { Subnet } from "./subnet";
import { SubnetPool } from "./subnetPool";
import { SubnetRoute } from "./subnetRoute";
import { Trunk } from "./trunk";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "openstack:networking/addressScope:AddressScope":
                return new AddressScope(name, <any>undefined, { urn })
            case "openstack:networking/floatingIp:FloatingIp":
                return new FloatingIp(name, <any>undefined, { urn })
            case "openstack:networking/floatingIpAssociate:FloatingIpAssociate":
                return new FloatingIpAssociate(name, <any>undefined, { urn })
            case "openstack:networking/network:Network":
                return new Network(name, <any>undefined, { urn })
            case "openstack:networking/port:Port":
                return new Port(name, <any>undefined, { urn })
            case "openstack:networking/portSecGroupAssociate:PortSecGroupAssociate":
                return new PortSecGroupAssociate(name, <any>undefined, { urn })
            case "openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule":
                return new QosBandwidthLimitRule(name, <any>undefined, { urn })
            case "openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule":
                return new QosDscpMarkingRule(name, <any>undefined, { urn })
            case "openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule":
                return new QosMinimumBandwidthRule(name, <any>undefined, { urn })
            case "openstack:networking/qosPolicy:QosPolicy":
                return new QosPolicy(name, <any>undefined, { urn })
            case "openstack:networking/quotaV2:QuotaV2":
                return new QuotaV2(name, <any>undefined, { urn })
            case "openstack:networking/rbacPolicyV2:RbacPolicyV2":
                return new RbacPolicyV2(name, <any>undefined, { urn })
            case "openstack:networking/router:Router":
                return new Router(name, <any>undefined, { urn })
            case "openstack:networking/routerInterface:RouterInterface":
                return new RouterInterface(name, <any>undefined, { urn })
            case "openstack:networking/routerRoute:RouterRoute":
                return new RouterRoute(name, <any>undefined, { urn })
            case "openstack:networking/secGroup:SecGroup":
                return new SecGroup(name, <any>undefined, { urn })
            case "openstack:networking/secGroupRule:SecGroupRule":
                return new SecGroupRule(name, <any>undefined, { urn })
            case "openstack:networking/subnet:Subnet":
                return new Subnet(name, <any>undefined, { urn })
            case "openstack:networking/subnetPool:SubnetPool":
                return new SubnetPool(name, <any>undefined, { urn })
            case "openstack:networking/subnetRoute:SubnetRoute":
                return new SubnetRoute(name, <any>undefined, { urn })
            case "openstack:networking/trunk:Trunk":
                return new Trunk(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("openstack", "networking/addressScope", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/floatingIp", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/floatingIpAssociate", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/network", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/port", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/portSecGroupAssociate", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/qosBandwidthLimitRule", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/qosDscpMarkingRule", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/qosMinimumBandwidthRule", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/qosPolicy", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/quotaV2", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/rbacPolicyV2", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/router", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/routerInterface", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/routerRoute", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/secGroup", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/secGroupRule", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/subnet", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/subnetPool", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/subnetRoute", _module)
pulumi.runtime.registerResourceModule("openstack", "networking/trunk", _module)
