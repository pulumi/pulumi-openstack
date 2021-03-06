// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a V1 Barbican container resource within OpenStack.
 *
 * ## Example Usage
 * ### Simple secret
 *
 * The container with the TLS certificates, which can be used by the loadbalancer HTTPS listener.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as openstack from "@pulumi/openstack";
 *
 * const certificate1 = new openstack.keymanager.SecretV1("certificate_1", {
 *     payload: fs.readFileSync("cert.pem", "utf-8"),
 *     payloadContentType: "text/plain",
 *     secretType: "certificate",
 * });
 * const privateKey1 = new openstack.keymanager.SecretV1("private_key_1", {
 *     payload: fs.readFileSync("cert-key.pem", "utf-8"),
 *     payloadContentType: "text/plain",
 *     secretType: "private",
 * });
 * const intermediate1 = new openstack.keymanager.SecretV1("intermediate_1", {
 *     payload: fs.readFileSync("intermediate-ca.pem", "utf-8"),
 *     payloadContentType: "text/plain",
 *     secretType: "certificate",
 * });
 * const tls1 = new openstack.keymanager.ContainerV1("tls_1", {
 *     secretRefs: [
 *         {
 *             name: "certificate",
 *             secretRef: certificate1.secretRef,
 *         },
 *         {
 *             name: "private_key",
 *             secretRef: privateKey1.secretRef,
 *         },
 *         {
 *             name: "intermediates",
 *             secretRef: intermediate1.secretRef,
 *         },
 *     ],
 *     type: "certificate",
 * });
 * const subnet1 = pulumi.output(openstack.networking.getSubnet({
 *     name: "my-subnet",
 * }, { async: true }));
 * const lb1 = new openstack.loadbalancer.LoadBalancer("lb_1", {
 *     vipSubnetId: subnet1.id,
 * });
 * const listener1 = new openstack.loadbalancer.Listener("listener_1", {
 *     defaultTlsContainerRef: tls1.containerRef,
 *     loadbalancerId: lb1.id,
 *     protocol: "TERMINATED_HTTPS",
 *     protocolPort: 443,
 * });
 * ```
 * ### Container with the ACL
 *
 * > **Note** Only read ACLs are supported
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as openstack from "@pulumi/openstack";
 *
 * const tls1 = new openstack.keymanager.ContainerV1("tls_1", {
 *     acl: {
 *         read: {
 *             projectAccess: false,
 *             users: [
 *                 "userid1",
 *                 "userid2",
 *             ],
 *         },
 *     },
 *     secretRefs: [
 *         {
 *             name: "certificate",
 *             secretRef: openstack_keymanager_secret_v1_certificate_1.secretRef,
 *         },
 *         {
 *             name: "private_key",
 *             secretRef: openstack_keymanager_secret_v1_private_key_1.secretRef,
 *         },
 *         {
 *             name: "intermediates",
 *             secretRef: openstack_keymanager_secret_v1_intermediate_1.secretRef,
 *         },
 *     ],
 *     type: "certificate",
 * });
 * ```
 *
 * ## Import
 *
 * Containers can be imported using the container id (the last part of the container reference), e.g.
 *
 * ```sh
 *  $ pulumi import openstack:keymanager/containerV1:ContainerV1 container_1 0c6cd26a-c012-4d7b-8034-057c0f1c2953
 * ```
 */
export class ContainerV1 extends pulumi.CustomResource {
    /**
     * Get an existing ContainerV1 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerV1State, opts?: pulumi.CustomResourceOptions): ContainerV1 {
        return new ContainerV1(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:keymanager/containerV1:ContainerV1';

    /**
     * Returns true if the given object is an instance of ContainerV1.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerV1 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerV1.__pulumiType;
    }

    /**
     * Allows to control an access to a container. Currently only
     * the `read` operation is supported. If not specified, the container is
     * accessible project wide. The `read` structure is described below.
     */
    public readonly acl!: pulumi.Output<outputs.keymanager.ContainerV1Acl>;
    /**
     * The list of the container consumers. The structure is described below.
     */
    public /*out*/ readonly consumers!: pulumi.Output<outputs.keymanager.ContainerV1Consumer[]>;
    /**
     * The container reference / where to find the container.
     */
    public /*out*/ readonly containerRef!: pulumi.Output<string>;
    /**
     * The date the container ACL was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The creator of the container.
     */
    public /*out*/ readonly creatorId!: pulumi.Output<string>;
    /**
     * The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a container. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 container.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A set of dictionaries containing references to secrets. The structure is described
     * below.
     */
    public readonly secretRefs!: pulumi.Output<outputs.keymanager.ContainerV1SecretRef[] | undefined>;
    /**
     * The status of the container.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The date the container ACL was last updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ContainerV1 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerV1Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerV1Args | ContainerV1State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerV1State | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["consumers"] = state ? state.consumers : undefined;
            inputs["containerRef"] = state ? state.containerRef : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["creatorId"] = state ? state.creatorId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["secretRefs"] = state ? state.secretRefs : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ContainerV1Args | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["acl"] = args ? args.acl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["secretRefs"] = args ? args.secretRefs : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["consumers"] = undefined /*out*/;
            inputs["containerRef"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["creatorId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ContainerV1.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerV1 resources.
 */
export interface ContainerV1State {
    /**
     * Allows to control an access to a container. Currently only
     * the `read` operation is supported. If not specified, the container is
     * accessible project wide. The `read` structure is described below.
     */
    readonly acl?: pulumi.Input<inputs.keymanager.ContainerV1Acl>;
    /**
     * The list of the container consumers. The structure is described below.
     */
    readonly consumers?: pulumi.Input<pulumi.Input<inputs.keymanager.ContainerV1Consumer>[]>;
    /**
     * The container reference / where to find the container.
     */
    readonly containerRef?: pulumi.Input<string>;
    /**
     * The date the container ACL was created.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * The creator of the container.
     */
    readonly creatorId?: pulumi.Input<string>;
    /**
     * The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a container. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 container.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A set of dictionaries containing references to secrets. The structure is described
     * below.
     */
    readonly secretRefs?: pulumi.Input<pulumi.Input<inputs.keymanager.ContainerV1SecretRef>[]>;
    /**
     * The status of the container.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * The date the container ACL was last updated.
     */
    readonly updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerV1 resource.
 */
export interface ContainerV1Args {
    /**
     * Allows to control an access to a container. Currently only
     * the `read` operation is supported. If not specified, the container is
     * accessible project wide. The `read` structure is described below.
     */
    readonly acl?: pulumi.Input<inputs.keymanager.ContainerV1Acl>;
    /**
     * The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a container. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 container.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A set of dictionaries containing references to secrets. The structure is described
     * below.
     */
    readonly secretRefs?: pulumi.Input<pulumi.Input<inputs.keymanager.ContainerV1SecretRef>[]>;
    /**
     * Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
     */
    readonly type: pulumi.Input<string>;
}
