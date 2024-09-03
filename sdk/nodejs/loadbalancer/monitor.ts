// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a V2 monitor resource within OpenStack.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const monitor1 = new openstack.loadbalancer.Monitor("monitor_1", {
 *     poolId: pool1.id,
 *     type: "PING",
 *     delay: 20,
 *     timeout: 10,
 *     maxRetries: 5,
 * });
 * ```
 *
 * ## Import
 *
 * Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.:
 *
 * ```sh
 * $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2
 * ```
 * In case of using OpenContrail, the import may not work properly. If you face an issue, try to import the monitor providing its parent pool ID:
 *
 * ```sh
 * $ pulumi import openstack:loadbalancer/monitor:Monitor monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2/708bc224-0f8c-4981-ac82-97095fe051b6
 * ```
 */
export class Monitor extends pulumi.CustomResource {
    /**
     * Get an existing Monitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitorState, opts?: pulumi.CustomResourceOptions): Monitor {
        return new Monitor(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:loadbalancer/monitor:Monitor';

    /**
     * Returns true if the given object is an instance of Monitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Monitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Monitor.__pulumiType;
    }

    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     */
    public readonly adminStateUp!: pulumi.Output<boolean | undefined>;
    /**
     * The time, in seconds, between sending probes to members.
     */
    public readonly delay!: pulumi.Output<number>;
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * "200", a list like "200, 202" or a range like "200-202". Default is "200".
     */
    public readonly expectedCodes!: pulumi.Output<string>;
    /**
     * Required for HTTP(S) types. The HTTP method that
     * the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
     * OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
     */
    public readonly httpMethod!: pulumi.Output<string>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to INACTIVE. Must be a number between 1
     * and 10.
     */
    public readonly maxRetries!: pulumi.Output<number>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to ERROR. Must be a number between 1 and 10.
     * The default is 3. Changing this updates the maxRetriesDown of the
     * existing monitor.
     */
    public readonly maxRetriesDown!: pulumi.Output<number>;
    /**
     * The Name of the Monitor.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the pool that this monitor will be assigned to.
     */
    public readonly poolId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     */
    public readonly timeout!: pulumi.Output<number>;
    /**
     * The type of probe, which is PING, TCP, HTTP, HTTPS,
     * TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
     * verify the member state. Changing this creates a new monitor.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Default is `/`.
     */
    public readonly urlPath!: pulumi.Output<string>;

    /**
     * Create a Monitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitorArgs | MonitorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitorState | undefined;
            resourceInputs["adminStateUp"] = state ? state.adminStateUp : undefined;
            resourceInputs["delay"] = state ? state.delay : undefined;
            resourceInputs["expectedCodes"] = state ? state.expectedCodes : undefined;
            resourceInputs["httpMethod"] = state ? state.httpMethod : undefined;
            resourceInputs["maxRetries"] = state ? state.maxRetries : undefined;
            resourceInputs["maxRetriesDown"] = state ? state.maxRetriesDown : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["poolId"] = state ? state.poolId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["urlPath"] = state ? state.urlPath : undefined;
        } else {
            const args = argsOrState as MonitorArgs | undefined;
            if ((!args || args.delay === undefined) && !opts.urn) {
                throw new Error("Missing required property 'delay'");
            }
            if ((!args || args.maxRetries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxRetries'");
            }
            if ((!args || args.poolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'poolId'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["adminStateUp"] = args ? args.adminStateUp : undefined;
            resourceInputs["delay"] = args ? args.delay : undefined;
            resourceInputs["expectedCodes"] = args ? args.expectedCodes : undefined;
            resourceInputs["httpMethod"] = args ? args.httpMethod : undefined;
            resourceInputs["maxRetries"] = args ? args.maxRetries : undefined;
            resourceInputs["maxRetriesDown"] = args ? args.maxRetriesDown : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["poolId"] = args ? args.poolId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["urlPath"] = args ? args.urlPath : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Monitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Monitor resources.
 */
export interface MonitorState {
    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, between sending probes to members.
     */
    delay?: pulumi.Input<number>;
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * "200", a list like "200, 202" or a range like "200-202". Default is "200".
     */
    expectedCodes?: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. The HTTP method that
     * the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
     * OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
     */
    httpMethod?: pulumi.Input<string>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to INACTIVE. Must be a number between 1
     * and 10.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to ERROR. Must be a number between 1 and 10.
     * The default is 3. Changing this updates the maxRetriesDown of the
     * existing monitor.
     */
    maxRetriesDown?: pulumi.Input<number>;
    /**
     * The Name of the Monitor.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the pool that this monitor will be assigned to.
     */
    poolId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     */
    region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     */
    timeout?: pulumi.Input<number>;
    /**
     * The type of probe, which is PING, TCP, HTTP, HTTPS,
     * TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
     * verify the member state. Changing this creates a new monitor.
     */
    type?: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Default is `/`.
     */
    urlPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Monitor resource.
 */
export interface MonitorArgs {
    /**
     * The administrative state of the monitor.
     * A valid value is true (UP) or false (DOWN).
     */
    adminStateUp?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, between sending probes to members.
     */
    delay: pulumi.Input<number>;
    /**
     * Required for HTTP(S) types. Expected HTTP codes
     * for a passing HTTP(S) monitor. You can either specify a single status like
     * "200", a list like "200, 202" or a range like "200-202". Default is "200".
     */
    expectedCodes?: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. The HTTP method that
     * the health monitor uses for requests. One of CONNECT, DELETE, GET, HEAD,
     * OPTIONS, PATCH, POST, PUT, or TRACE. The default is GET
     */
    httpMethod?: pulumi.Input<string>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to INACTIVE. Must be a number between 1
     * and 10.
     */
    maxRetries: pulumi.Input<number>;
    /**
     * Number of permissible ping failures before
     * changing the member's status to ERROR. Must be a number between 1 and 10.
     * The default is 3. Changing this updates the maxRetriesDown of the
     * existing monitor.
     */
    maxRetriesDown?: pulumi.Input<number>;
    /**
     * The Name of the Monitor.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the pool that this monitor will be assigned to.
     */
    poolId: pulumi.Input<string>;
    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create an . If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * monitor.
     */
    region?: pulumi.Input<string>;
    /**
     * Required for admins. The UUID of the tenant who owns
     * the monitor.  Only administrative users can specify a tenant UUID
     * other than their own. Changing this creates a new monitor.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * Maximum number of seconds for a monitor to wait for a
     * ping reply before it times out. The value must be less than the delay
     * value.
     */
    timeout: pulumi.Input<number>;
    /**
     * The type of probe, which is PING, TCP, HTTP, HTTPS,
     * TLS-HELLO, SCTP or UDP-CONNECT, that is sent by the loadbalancer to
     * verify the member state. Changing this creates a new monitor.
     */
    type: pulumi.Input<string>;
    /**
     * Required for HTTP(S) types. URI path that will be
     * accessed if monitor type is HTTP or HTTPS. Default is `/`.
     */
    urlPath?: pulumi.Input<string>;
}
