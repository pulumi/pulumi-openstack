// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the details of a running server
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const instance = pulumi.output(openstack.compute.getInstanceV2({
 *     // Randomly generated UUID, for demonstration purposes
 *     id: "2ba26dc6-a12d-4889-8f25-794ea5bf4453",
 * }, { async: true }));
 * ```
 */
export function getInstanceV2(args: GetInstanceV2Args, opts?: pulumi.InvokeOptions): Promise<GetInstanceV2Result> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("openstack:compute/getInstanceV2:getInstanceV2", {
        "id": args.id,
        "networks": args.networks,
        "region": args.region,
        "userData": args.userData,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceV2.
 */
export interface GetInstanceV2Args {
    /**
     * The UUID of the instance
     */
    readonly id: string;
    /**
     * An array of maps, detailed below.
     */
    readonly networks?: inputs.compute.GetInstanceV2Network[];
    readonly region?: string;
    /**
     * The user data added when the server was created.
     */
    readonly userData?: string;
}

/**
 * A collection of values returned by getInstanceV2.
 */
export interface GetInstanceV2Result {
    /**
     * The first IPv4 address assigned to this server.
     */
    readonly accessIpV4: string;
    /**
     * The first IPv6 address assigned to this server.
     */
    readonly accessIpV6: string;
    /**
     * The availability zone of this server.
     */
    readonly availabilityZone: string;
    /**
     * The flavor ID used to create the server.
     */
    readonly flavorId: string;
    /**
     * The flavor name used to create the server.
     */
    readonly flavorName: string;
    readonly id: string;
    /**
     * The image ID used to create the server.
     */
    readonly imageId: string;
    /**
     * The image name used to create the server.
     */
    readonly imageName: string;
    /**
     * The name of the key pair assigned to this server.
     */
    readonly keyPair: string;
    /**
     * A set of key/value pairs made available to the server.
     */
    readonly metadata: {[key: string]: any};
    /**
     * The name of the network
     */
    readonly name: string;
    /**
     * An array of maps, detailed below.
     */
    readonly networks: outputs.compute.GetInstanceV2Network[];
    readonly powerState: string;
    readonly region: string;
    /**
     * An array of security group names associated with this server.
     */
    readonly securityGroups: string[];
    /**
     * A set of string tags assigned to this server.
     */
    readonly tags: string[];
    /**
     * The user data added when the server was created.
     */
    readonly userData: string;
}
