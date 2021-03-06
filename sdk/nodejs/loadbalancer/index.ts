// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./l7policyV2";
export * from "./l7ruleV2";
export * from "./listener";
export * from "./loadBalancer";
export * from "./member";
export * from "./memberV1";
export * from "./members";
export * from "./monitor";
export * from "./monitorV1";
export * from "./pool";
export * from "./poolV1";
export * from "./quota";
export * from "./vip";

// Import resources to register:
import { L7PolicyV2 } from "./l7policyV2";
import { L7RuleV2 } from "./l7ruleV2";
import { Listener } from "./listener";
import { LoadBalancer } from "./loadBalancer";
import { Member } from "./member";
import { MemberV1 } from "./memberV1";
import { Members } from "./members";
import { Monitor } from "./monitor";
import { MonitorV1 } from "./monitorV1";
import { Pool } from "./pool";
import { PoolV1 } from "./poolV1";
import { Quota } from "./quota";
import { Vip } from "./vip";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "openstack:loadbalancer/l7PolicyV2:L7PolicyV2":
                return new L7PolicyV2(name, <any>undefined, { urn })
            case "openstack:loadbalancer/l7RuleV2:L7RuleV2":
                return new L7RuleV2(name, <any>undefined, { urn })
            case "openstack:loadbalancer/listener:Listener":
                return new Listener(name, <any>undefined, { urn })
            case "openstack:loadbalancer/loadBalancer:LoadBalancer":
                return new LoadBalancer(name, <any>undefined, { urn })
            case "openstack:loadbalancer/member:Member":
                return new Member(name, <any>undefined, { urn })
            case "openstack:loadbalancer/memberV1:MemberV1":
                return new MemberV1(name, <any>undefined, { urn })
            case "openstack:loadbalancer/members:Members":
                return new Members(name, <any>undefined, { urn })
            case "openstack:loadbalancer/monitor:Monitor":
                return new Monitor(name, <any>undefined, { urn })
            case "openstack:loadbalancer/monitorV1:MonitorV1":
                return new MonitorV1(name, <any>undefined, { urn })
            case "openstack:loadbalancer/pool:Pool":
                return new Pool(name, <any>undefined, { urn })
            case "openstack:loadbalancer/poolV1:PoolV1":
                return new PoolV1(name, <any>undefined, { urn })
            case "openstack:loadbalancer/quota:Quota":
                return new Quota(name, <any>undefined, { urn })
            case "openstack:loadbalancer/vip:Vip":
                return new Vip(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/l7PolicyV2", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/l7RuleV2", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/listener", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/loadBalancer", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/member", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/memberV1", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/members", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/monitor", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/monitorV1", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/pool", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/poolV1", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/quota", _module)
pulumi.runtime.registerResourceModule("openstack", "loadbalancer/vip", _module)
