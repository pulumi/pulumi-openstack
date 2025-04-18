// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { BgpvpnNetworkAssociateV2Args, BgpvpnNetworkAssociateV2State } from "./bgpvpnNetworkAssociateV2";
export type BgpvpnNetworkAssociateV2 = import("./bgpvpnNetworkAssociateV2").BgpvpnNetworkAssociateV2;
export const BgpvpnNetworkAssociateV2: typeof import("./bgpvpnNetworkAssociateV2").BgpvpnNetworkAssociateV2 = null as any;
utilities.lazyLoad(exports, ["BgpvpnNetworkAssociateV2"], () => require("./bgpvpnNetworkAssociateV2"));

export { BgpvpnPortAssociateV2Args, BgpvpnPortAssociateV2State } from "./bgpvpnPortAssociateV2";
export type BgpvpnPortAssociateV2 = import("./bgpvpnPortAssociateV2").BgpvpnPortAssociateV2;
export const BgpvpnPortAssociateV2: typeof import("./bgpvpnPortAssociateV2").BgpvpnPortAssociateV2 = null as any;
utilities.lazyLoad(exports, ["BgpvpnPortAssociateV2"], () => require("./bgpvpnPortAssociateV2"));

export { BgpvpnRouterAssociateV2Args, BgpvpnRouterAssociateV2State } from "./bgpvpnRouterAssociateV2";
export type BgpvpnRouterAssociateV2 = import("./bgpvpnRouterAssociateV2").BgpvpnRouterAssociateV2;
export const BgpvpnRouterAssociateV2: typeof import("./bgpvpnRouterAssociateV2").BgpvpnRouterAssociateV2 = null as any;
utilities.lazyLoad(exports, ["BgpvpnRouterAssociateV2"], () => require("./bgpvpnRouterAssociateV2"));

export { BgpvpnV2Args, BgpvpnV2State } from "./bgpvpnV2";
export type BgpvpnV2 = import("./bgpvpnV2").BgpvpnV2;
export const BgpvpnV2: typeof import("./bgpvpnV2").BgpvpnV2 = null as any;
utilities.lazyLoad(exports, ["BgpvpnV2"], () => require("./bgpvpnV2"));

export { GetFwGroupV2Args, GetFwGroupV2Result, GetFwGroupV2OutputArgs } from "./getFwGroupV2";
export const getFwGroupV2: typeof import("./getFwGroupV2").getFwGroupV2 = null as any;
export const getFwGroupV2Output: typeof import("./getFwGroupV2").getFwGroupV2Output = null as any;
utilities.lazyLoad(exports, ["getFwGroupV2","getFwGroupV2Output"], () => require("./getFwGroupV2"));

export { GetFwPolicyV2Args, GetFwPolicyV2Result, GetFwPolicyV2OutputArgs } from "./getFwPolicyV2";
export const getFwPolicyV2: typeof import("./getFwPolicyV2").getFwPolicyV2 = null as any;
export const getFwPolicyV2Output: typeof import("./getFwPolicyV2").getFwPolicyV2Output = null as any;
utilities.lazyLoad(exports, ["getFwPolicyV2","getFwPolicyV2Output"], () => require("./getFwPolicyV2"));

export { GetFwRuleV2Args, GetFwRuleV2Result, GetFwRuleV2OutputArgs } from "./getFwRuleV2";
export const getFwRuleV2: typeof import("./getFwRuleV2").getFwRuleV2 = null as any;
export const getFwRuleV2Output: typeof import("./getFwRuleV2").getFwRuleV2Output = null as any;
utilities.lazyLoad(exports, ["getFwRuleV2","getFwRuleV2Output"], () => require("./getFwRuleV2"));

export { LbFlavorprofileV2Args, LbFlavorprofileV2State } from "./lbFlavorprofileV2";
export type LbFlavorprofileV2 = import("./lbFlavorprofileV2").LbFlavorprofileV2;
export const LbFlavorprofileV2: typeof import("./lbFlavorprofileV2").LbFlavorprofileV2 = null as any;
utilities.lazyLoad(exports, ["LbFlavorprofileV2"], () => require("./lbFlavorprofileV2"));

export { LbLoadbalancerV2Args, LbLoadbalancerV2State } from "./lbLoadbalancerV2";
export type LbLoadbalancerV2 = import("./lbLoadbalancerV2").LbLoadbalancerV2;
export const LbLoadbalancerV2: typeof import("./lbLoadbalancerV2").LbLoadbalancerV2 = null as any;
utilities.lazyLoad(exports, ["LbLoadbalancerV2"], () => require("./lbLoadbalancerV2"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as bgpvpn from "./bgpvpn";
import * as blockstorage from "./blockstorage";
import * as compute from "./compute";
import * as config from "./config";
import * as containerinfra from "./containerinfra";
import * as database from "./database";
import * as dns from "./dns";
import * as firewall from "./firewall";
import * as identity from "./identity";
import * as images from "./images";
import * as keymanager from "./keymanager";
import * as loadbalancer from "./loadbalancer";
import * as networking from "./networking";
import * as objectstorage from "./objectstorage";
import * as orchestration from "./orchestration";
import * as sharedfilesystem from "./sharedfilesystem";
import * as types from "./types";
import * as vpnaas from "./vpnaas";

export {
    bgpvpn,
    blockstorage,
    compute,
    config,
    containerinfra,
    database,
    dns,
    firewall,
    identity,
    images,
    keymanager,
    loadbalancer,
    networking,
    objectstorage,
    orchestration,
    sharedfilesystem,
    types,
    vpnaas,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "openstack:index/bgpvpnNetworkAssociateV2:BgpvpnNetworkAssociateV2":
                return new BgpvpnNetworkAssociateV2(name, <any>undefined, { urn })
            case "openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2":
                return new BgpvpnPortAssociateV2(name, <any>undefined, { urn })
            case "openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2":
                return new BgpvpnRouterAssociateV2(name, <any>undefined, { urn })
            case "openstack:index/bgpvpnV2:BgpvpnV2":
                return new BgpvpnV2(name, <any>undefined, { urn })
            case "openstack:index/lbFlavorprofileV2:LbFlavorprofileV2":
                return new LbFlavorprofileV2(name, <any>undefined, { urn })
            case "openstack:index/lbLoadbalancerV2:LbLoadbalancerV2":
                return new LbLoadbalancerV2(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("openstack", "index/bgpvpnNetworkAssociateV2", _module)
pulumi.runtime.registerResourceModule("openstack", "index/bgpvpnPortAssociateV2", _module)
pulumi.runtime.registerResourceModule("openstack", "index/bgpvpnRouterAssociateV2", _module)
pulumi.runtime.registerResourceModule("openstack", "index/bgpvpnV2", _module)
pulumi.runtime.registerResourceModule("openstack", "index/lbFlavorprofileV2", _module)
pulumi.runtime.registerResourceModule("openstack", "index/lbLoadbalancerV2", _module)
pulumi.runtime.registerResourcePackage("openstack", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:openstack") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
