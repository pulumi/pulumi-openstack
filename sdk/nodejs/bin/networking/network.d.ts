import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 Neutron network resource within OpenStack.
 */
export declare class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState): Network;
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing network.
     */
    readonly adminStateUp: pulumi.Output<string>;
    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     */
    readonly availabilityZoneHints: pulumi.Output<string[]>;
    /**
     * The name of the network. Changing this updates the name of
     * the existing network.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     */
    readonly region: pulumi.Output<string>;
    /**
     * An array of one or more provider segment objects.
     */
    readonly segments: pulumi.Output<{
        networkType?: string;
        physicalNetwork?: string;
        segmentationId?: number;
    }[] | undefined>;
    /**
     * Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabalities of the
     * existing network.
     */
    readonly shared: pulumi.Output<string>;
    /**
     * The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing network.
     */
    readonly adminStateUp?: pulumi.Input<string>;
    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     */
    readonly availabilityZoneHints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the network. Changing this updates the name of
     * the existing network.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * An array of one or more provider segment objects.
     */
    readonly segments?: pulumi.Input<{
        networkType?: pulumi.Input<string>;
        physicalNetwork?: pulumi.Input<string>;
        segmentationId?: pulumi.Input<number>;
    }[]>;
    /**
     * Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabalities of the
     * existing network.
     */
    readonly shared?: pulumi.Input<string>;
    /**
     * The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * The administrative state of the network.
     * Acceptable values are "true" and "false". Changing this value updates the
     * state of the existing network.
     */
    readonly adminStateUp?: pulumi.Input<string>;
    /**
     * An availability zone is used to make
     * network resources highly available. Used for resources with high availability
     * so that they are scheduled on different availability zones. Changing this
     * creates a new network.
     */
    readonly availabilityZoneHints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the network. Changing this updates the name of
     * the existing network.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron network. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * network.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * An array of one or more provider segment objects.
     */
    readonly segments?: pulumi.Input<{
        networkType?: pulumi.Input<string>;
        physicalNetwork?: pulumi.Input<string>;
        segmentationId?: pulumi.Input<number>;
    }[]>;
    /**
     * Specifies whether the network resource can be accessed
     * by any tenant or not. Changing this updates the sharing capabalities of the
     * existing network.
     */
    readonly shared?: pulumi.Input<string>;
    /**
     * The owner of the network. Required if admin wants to
     * create a network for another tenant. Changing this creates a new network.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * Map of additional options.
     */
    readonly valueSpecs?: pulumi.Input<{
        [key: string]: any;
    }>;
}
