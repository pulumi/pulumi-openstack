// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./endpointGroup";
export * from "./ikePolicy";
export * from "./ipSecPolicy";
export * from "./service";
export * from "./siteConnection";

// Import resources to register:
import { EndpointGroup } from "./endpointGroup";
import { IkePolicy } from "./ikePolicy";
import { IpSecPolicy } from "./ipSecPolicy";
import { Service } from "./service";
import { SiteConnection } from "./siteConnection";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "openstack:vpnaas/endpointGroup:EndpointGroup":
                return new EndpointGroup(name, <any>undefined, { urn })
            case "openstack:vpnaas/ikePolicy:IkePolicy":
                return new IkePolicy(name, <any>undefined, { urn })
            case "openstack:vpnaas/ipSecPolicy:IpSecPolicy":
                return new IpSecPolicy(name, <any>undefined, { urn })
            case "openstack:vpnaas/service:Service":
                return new Service(name, <any>undefined, { urn })
            case "openstack:vpnaas/siteConnection:SiteConnection":
                return new SiteConnection(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("openstack", "vpnaas/endpointGroup", _module)
pulumi.runtime.registerResourceModule("openstack", "vpnaas/ikePolicy", _module)
pulumi.runtime.registerResourceModule("openstack", "vpnaas/ipSecPolicy", _module)
pulumi.runtime.registerResourceModule("openstack", "vpnaas/service", _module)
pulumi.runtime.registerResourceModule("openstack", "vpnaas/siteConnection", _module)
