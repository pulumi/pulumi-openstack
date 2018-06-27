import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 listener resource within OpenStack.
 */
export declare class Listener extends pulumi.CustomResource {
    /**
     * Get an existing Listener resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerState): Listener;
    /**
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp: pulumi.Output<boolean | undefined>;
    /**
     * The maximum number of connections allowed
     * for the Listener.
     */
    readonly connectionLimit: pulumi.Output<number>;
    /**
     * The ID of the default pool with which the
     * Listener is associated. Changing this creates a new Listener.
     */
    readonly defaultPoolId: pulumi.Output<string>;
    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly defaultTlsContainerRef: pulumi.Output<string | undefined>;
    /**
     * Human-readable description for the Listener.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     */
    readonly loadbalancerId: pulumi.Output<string>;
    /**
     * Human-readable name for the Listener. Does not have
     * to be unique.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
     * Changing this creates a new Listener.
     */
    readonly protocol: pulumi.Output<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     */
    readonly protocolPort: pulumi.Output<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     */
    readonly region: pulumi.Output<string>;
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly sniContainerRefs: pulumi.Output<string[] | undefined>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     */
    readonly tenantId: pulumi.Output<string>;
    /**
     * Create a Listener resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Listener resources.
 */
export interface ListenerState {
    /**
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The maximum number of connections allowed
     * for the Listener.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * The ID of the default pool with which the
     * Listener is associated. Changing this creates a new Listener.
     */
    readonly defaultPoolId?: pulumi.Input<string>;
    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly defaultTlsContainerRef?: pulumi.Input<string>;
    /**
     * Human-readable description for the Listener.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     */
    readonly loadbalancerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the Listener. Does not have
     * to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
     * Changing this creates a new Listener.
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     */
    readonly protocolPort?: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly sniContainerRefs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     */
    readonly tenantId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Listener resource.
 */
export interface ListenerArgs {
    /**
     * The administrative state of the Listener.
     * A valid value is true (UP) or false (DOWN).
     */
    readonly adminStateUp?: pulumi.Input<boolean>;
    /**
     * The maximum number of connections allowed
     * for the Listener.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * The ID of the default pool with which the
     * Listener is associated. Changing this creates a new Listener.
     */
    readonly defaultPoolId?: pulumi.Input<string>;
    /**
     * A reference to a Barbican Secrets
     * container which stores TLS information. This is required if the protocol
     * is `TERMINATED_HTTPS`. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly defaultTlsContainerRef?: pulumi.Input<string>;
    /**
     * Human-readable description for the Listener.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The load balancer on which to provision this
     * Listener. Changing this creates a new Listener.
     */
    readonly loadbalancerId: pulumi.Input<string>;
    /**
     * Human-readable name for the Listener. Does not have
     * to be unique.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS.
     * Changing this creates a new Listener.
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The port on which to listen for client traffic.
     * Changing this creates a new Listener.
     */
    readonly protocolPort: pulumi.Input<number>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * Listener.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A list of references to Barbican Secrets
     * containers which store SNI information. See
     * [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer)
     * for more information.
     */
    readonly sniContainerRefs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the Listener.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new Listener.
     */
    readonly tenantId?: pulumi.Input<string>;
}
