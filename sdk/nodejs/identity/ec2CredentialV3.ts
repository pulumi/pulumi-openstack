// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * EC2 Credentials can be imported using the `access`, e.g.
 *
 * ```sh
 *  $ pulumi import openstack:identity/ec2CredentialV3:Ec2CredentialV3 ec2_cred_1 2d0ac4a2f81b4b0f9513ee49e780647d
 * ```
 */
export class Ec2CredentialV3 extends pulumi.CustomResource {
    /**
     * Get an existing Ec2CredentialV3 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ec2CredentialV3State, opts?: pulumi.CustomResourceOptions): Ec2CredentialV3 {
        return new Ec2CredentialV3(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:identity/ec2CredentialV3:Ec2CredentialV3';

    /**
     * Returns true if the given object is an instance of Ec2CredentialV3.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ec2CredentialV3 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ec2CredentialV3.__pulumiType;
    }

    /**
     * contains an EC2 credential access UUID
     */
    public /*out*/ readonly access!: pulumi.Output<string>;
    /**
     * The ID of the project the EC2 credential is created
     * for and that authentication requests using this EC2 credential will
     * be scoped to.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new EC2 credential.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * contains an EC2 credential secret UUID
     */
    public /*out*/ readonly secret!: pulumi.Output<string>;
    /**
     * contains an EC2 credential trust ID scope
     */
    public /*out*/ readonly trustId!: pulumi.Output<string>;
    /**
     * The ID of the user the EC2 credential is created for.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a Ec2CredentialV3 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Ec2CredentialV3Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ec2CredentialV3Args | Ec2CredentialV3State, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as Ec2CredentialV3State | undefined;
            inputs["access"] = state ? state.access : undefined;
            inputs["projectId"] = state ? state.projectId : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["secret"] = state ? state.secret : undefined;
            inputs["trustId"] = state ? state.trustId : undefined;
            inputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as Ec2CredentialV3Args | undefined;
            inputs["projectId"] = args ? args.projectId : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["userId"] = args ? args.userId : undefined;
            inputs["access"] = undefined /*out*/;
            inputs["secret"] = undefined /*out*/;
            inputs["trustId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Ec2CredentialV3.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ec2CredentialV3 resources.
 */
export interface Ec2CredentialV3State {
    /**
     * contains an EC2 credential access UUID
     */
    readonly access?: pulumi.Input<string>;
    /**
     * The ID of the project the EC2 credential is created
     * for and that authentication requests using this EC2 credential will
     * be scoped to.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new EC2 credential.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * contains an EC2 credential secret UUID
     */
    readonly secret?: pulumi.Input<string>;
    /**
     * contains an EC2 credential trust ID scope
     */
    readonly trustId?: pulumi.Input<string>;
    /**
     * The ID of the user the EC2 credential is created for.
     */
    readonly userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ec2CredentialV3 resource.
 */
export interface Ec2CredentialV3Args {
    /**
     * The ID of the project the EC2 credential is created
     * for and that authentication requests using this EC2 credential will
     * be scoped to.
     */
    readonly projectId?: pulumi.Input<string>;
    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new EC2 credential.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The ID of the user the EC2 credential is created for.
     */
    readonly userId?: pulumi.Input<string>;
}
