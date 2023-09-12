// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
 * const certificate1 = new openstack.keymanager.SecretV1("certificate1", {
 *     payload: fs.readFileSync("cert.pem"),
 *     secretType: "certificate",
 *     payloadContentType: "text/plain",
 * });
 * const privateKey1 = new openstack.keymanager.SecretV1("privateKey1", {
 *     payload: fs.readFileSync("cert-key.pem"),
 *     secretType: "private",
 *     payloadContentType: "text/plain",
 * });
 * const intermediate1 = new openstack.keymanager.SecretV1("intermediate1", {
 *     payload: fs.readFileSync("intermediate-ca.pem"),
 *     secretType: "certificate",
 *     payloadContentType: "text/plain",
 * });
 * const tls1 = new openstack.keymanager.ContainerV1("tls1", {
 *     type: "certificate",
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
 * });
 * const subnet1 = openstack.networking.getSubnet({
 *     name: "my-subnet",
 * });
 * const lb1 = new openstack.loadbalancer.LoadBalancer("lb1", {vipSubnetId: subnet1.then(subnet1 => subnet1.id)});
 * const listener1 = new openstack.loadbalancer.Listener("listener1", {
 *     protocol: "TERMINATED_HTTPS",
 *     protocolPort: 443,
 *     loadbalancerId: lb1.id,
 *     defaultTlsContainerRef: tls1.containerRef,
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
 * const tls1 = new openstack.keymanager.ContainerV1("tls1", {
 *     type: "certificate",
 *     secretRefs: [
 *         {
 *             name: "certificate",
 *             secretRef: openstack_keymanager_secret_v1.certificate_1.secret_ref,
 *         },
 *         {
 *             name: "private_key",
 *             secretRef: openstack_keymanager_secret_v1.private_key_1.secret_ref,
 *         },
 *         {
 *             name: "intermediates",
 *             secretRef: openstack_keymanager_secret_v1.intermediate_1.secret_ref,
 *         },
 *     ],
 *     acl: {
 *         read: {
 *             projectAccess: false,
 *             users: [
 *                 "userid1",
 *                 "userid2",
 *             ],
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Containers can be imported using the container id (the last part of the container reference), e.g.:
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
     * Human-readable name for the Container. Does not have
     * to be unique.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerV1State | undefined;
            resourceInputs["acl"] = state ? state.acl : undefined;
            resourceInputs["consumers"] = state ? state.consumers : undefined;
            resourceInputs["containerRef"] = state ? state.containerRef : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["creatorId"] = state ? state.creatorId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretRefs"] = state ? state.secretRefs : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ContainerV1Args | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretRefs"] = args ? args.secretRefs : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["consumers"] = undefined /*out*/;
            resourceInputs["containerRef"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["creatorId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerV1.__pulumiType, name, resourceInputs, opts);
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
    acl?: pulumi.Input<inputs.keymanager.ContainerV1Acl>;
    /**
     * The list of the container consumers. The structure is described below.
     */
    consumers?: pulumi.Input<pulumi.Input<inputs.keymanager.ContainerV1Consumer>[]>;
    /**
     * The container reference / where to find the container.
     */
    containerRef?: pulumi.Input<string>;
    /**
     * The date the container ACL was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The creator of the container.
     */
    creatorId?: pulumi.Input<string>;
    /**
     * Human-readable name for the Container. Does not have
     * to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a container. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 container.
     */
    region?: pulumi.Input<string>;
    /**
     * A set of dictionaries containing references to secrets. The structure is described
     * below.
     */
    secretRefs?: pulumi.Input<pulumi.Input<inputs.keymanager.ContainerV1SecretRef>[]>;
    /**
     * The status of the container.
     */
    status?: pulumi.Input<string>;
    /**
     * Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
     */
    type?: pulumi.Input<string>;
    /**
     * The date the container ACL was last updated.
     */
    updatedAt?: pulumi.Input<string>;
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
    acl?: pulumi.Input<inputs.keymanager.ContainerV1Acl>;
    /**
     * Human-readable name for the Container. Does not have
     * to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a container. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 container.
     */
    region?: pulumi.Input<string>;
    /**
     * A set of dictionaries containing references to secrets. The structure is described
     * below.
     */
    secretRefs?: pulumi.Input<pulumi.Input<inputs.keymanager.ContainerV1SecretRef>[]>;
    /**
     * Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
     */
    type: pulumi.Input<string>;
}
