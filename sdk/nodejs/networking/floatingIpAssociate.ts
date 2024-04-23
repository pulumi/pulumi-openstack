// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates a floating IP to a port. This is useful for situations
 * where you have a pre-allocated floating IP or are unable to use the
 * `openstack.networking.FloatingIp` resource to create a floating IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const port1 = new openstack.networking.Port("port_1", {networkId: "a5bbd213-e1d3-49b6-aed1-9df60ea94b9a"});
 * const fip1 = new openstack.networking.FloatingIpAssociate("fip_1", {
 *     floatingIp: "1.2.3.4",
 *     portId: port1.id,
 * });
 * ```
 *
 * ## Import
 *
 * Floating IP associations can be imported using the `id` of the floating IP, e.g.
 *
 * ```sh
 * $ pulumi import openstack:networking/floatingIpAssociate:FloatingIpAssociate fip 2c7f39f3-702b-48d1-940c-b50384177ee1
 * ```
 */
export class FloatingIpAssociate extends pulumi.CustomResource {
    /**
     * Get an existing FloatingIpAssociate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FloatingIpAssociateState, opts?: pulumi.CustomResourceOptions): FloatingIpAssociate {
        return new FloatingIpAssociate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:networking/floatingIpAssociate:FloatingIpAssociate';

    /**
     * Returns true if the given object is an instance of FloatingIpAssociate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FloatingIpAssociate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FloatingIpAssociate.__pulumiType;
    }

    public readonly fixedIp!: pulumi.Output<string>;
    /**
     * IP Address of an existing floating IP.
     */
    public readonly floatingIp!: pulumi.Output<string>;
    /**
     * ID of an existing port with at least one IP address to
     * associate with this floating IP.
     */
    public readonly portId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a floating IP that can be used with
     * another networking resource, such as a load balancer. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * floating IP (which may or may not have a different address).
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a FloatingIpAssociate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FloatingIpAssociateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FloatingIpAssociateArgs | FloatingIpAssociateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FloatingIpAssociateState | undefined;
            resourceInputs["fixedIp"] = state ? state.fixedIp : undefined;
            resourceInputs["floatingIp"] = state ? state.floatingIp : undefined;
            resourceInputs["portId"] = state ? state.portId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as FloatingIpAssociateArgs | undefined;
            if ((!args || args.floatingIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'floatingIp'");
            }
            if ((!args || args.portId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'portId'");
            }
            resourceInputs["fixedIp"] = args ? args.fixedIp : undefined;
            resourceInputs["floatingIp"] = args ? args.floatingIp : undefined;
            resourceInputs["portId"] = args ? args.portId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FloatingIpAssociate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FloatingIpAssociate resources.
 */
export interface FloatingIpAssociateState {
    fixedIp?: pulumi.Input<string>;
    /**
     * IP Address of an existing floating IP.
     */
    floatingIp?: pulumi.Input<string>;
    /**
     * ID of an existing port with at least one IP address to
     * associate with this floating IP.
     */
    portId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a floating IP that can be used with
     * another networking resource, such as a load balancer. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * floating IP (which may or may not have a different address).
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FloatingIpAssociate resource.
 */
export interface FloatingIpAssociateArgs {
    fixedIp?: pulumi.Input<string>;
    /**
     * IP Address of an existing floating IP.
     */
    floatingIp: pulumi.Input<string>;
    /**
     * ID of an existing port with at least one IP address to
     * associate with this floating IP.
     */
    portId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a floating IP that can be used with
     * another networking resource, such as a load balancer. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * floating IP (which may or may not have a different address).
     */
    region?: pulumi.Input<string>;
}
