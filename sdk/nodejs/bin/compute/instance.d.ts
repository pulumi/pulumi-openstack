import * as pulumi from "@pulumi/pulumi";
/**
 * Manages a V2 VM instance resource within OpenStack.
 */
export declare class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState): Instance;
    /**
     * The first detected Fixed IPv4 address _or_ the
     * Floating IP.
     */
    readonly accessIpV4: pulumi.Output<string>;
    /**
     * The first detected Fixed IPv6 address.
     */
    readonly accessIpV6: pulumi.Output<string>;
    /**
     * The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     */
    readonly adminPass: pulumi.Output<string | undefined>;
    /**
     * Contains all instance metadata, even metadata not set
     * by Terraform.
     */
    readonly allMetadata: pulumi.Output<{
        [key: string]: any;
    }>;
    /**
     * The availability zone in which to create
     * the server. Changing this creates a new server.
     */
    readonly availabilityZone: pulumi.Output<string>;
    /**
     * Configuration of block devices. The block_device
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](http://docs.openstack.org/developer/nova/block_device_mapping.html)
     * for more information.
     */
    readonly blockDevices: pulumi.Output<{
        bootIndex?: number;
        deleteOnTermination?: boolean;
        destinationType?: string;
        guestFormat?: string;
        sourceType: string;
        uuid?: string;
        volumeSize?: number;
    }[] | undefined>;
    /**
     * Whether to use the config_drive feature to
     * configure the instance. Changing this creates a new server.
     */
    readonly configDrive: pulumi.Output<boolean | undefined>;
    /**
     * The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     */
    readonly flavorId: pulumi.Output<string>;
    /**
     * The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     */
    readonly flavorName: pulumi.Output<string>;
    /**
     * Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     */
    readonly forceDelete: pulumi.Output<boolean | undefined>;
    /**
     * (Optional; Required if `image_name` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this creates a new server.
     */
    readonly imageId: pulumi.Output<string>;
    /**
     * (Optional; Required if `image_id` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this creates a new server.
     */
    readonly imageName: pulumi.Output<string>;
    /**
     * The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant's account.
     * Changing this creates a new server.
     */
    readonly keyPair: pulumi.Output<string | undefined>;
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     */
    readonly metadata: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The human-readable
     * name of the network. Changing this creates a new server.
     */
    readonly name: pulumi.Output<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     */
    readonly networks: pulumi.Output<{
        accessNetwork?: boolean;
        fixedIpV4: string;
        fixedIpV6: string;
        floatingIp: string;
        mac: string;
        name: string;
        port: string;
        uuid: string;
    }[]>;
    /**
     * Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below.
     */
    readonly personalities: pulumi.Output<{
        content: string;
        file: string;
    }[] | undefined>;
    /**
     * The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     */
    readonly region: pulumi.Output<string>;
    /**
     * Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     */
    readonly schedulerHints: pulumi.Output<{
        additionalProperties?: {
            [key: string]: any;
        };
        buildNearHostIp?: string;
        differentHosts?: string[];
        group?: string;
        queries?: string[];
        sameHosts?: string[];
        targetCell?: string;
    }[] | undefined>;
    /**
     * An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance.
     */
    readonly securityGroups: pulumi.Output<string[]>;
    /**
     * Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn't stop within timeout, it will be destroyed anyway.
     */
    readonly stopBeforeDestroy: pulumi.Output<boolean | undefined>;
    /**
     * The user data to provide when launching the instance.
     * Changing this creates a new server.
     */
    readonly userData: pulumi.Output<string | undefined>;
    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.ResourceOptions);
}
/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The first detected Fixed IPv4 address _or_ the
     * Floating IP.
     */
    readonly accessIpV4?: pulumi.Input<string>;
    /**
     * The first detected Fixed IPv6 address.
     */
    readonly accessIpV6?: pulumi.Input<string>;
    /**
     * The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     */
    readonly adminPass?: pulumi.Input<string>;
    /**
     * Contains all instance metadata, even metadata not set
     * by Terraform.
     */
    readonly allMetadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The availability zone in which to create
     * the server. Changing this creates a new server.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Configuration of block devices. The block_device
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](http://docs.openstack.org/developer/nova/block_device_mapping.html)
     * for more information.
     */
    readonly blockDevices?: pulumi.Input<{
        bootIndex?: pulumi.Input<number>;
        deleteOnTermination?: pulumi.Input<boolean>;
        destinationType?: pulumi.Input<string>;
        guestFormat?: pulumi.Input<string>;
        sourceType: pulumi.Input<string>;
        uuid?: pulumi.Input<string>;
        volumeSize?: pulumi.Input<number>;
    }[]>;
    /**
     * Whether to use the config_drive feature to
     * configure the instance. Changing this creates a new server.
     */
    readonly configDrive?: pulumi.Input<boolean>;
    /**
     * The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     */
    readonly flavorId?: pulumi.Input<string>;
    /**
     * The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     */
    readonly flavorName?: pulumi.Input<string>;
    /**
     * Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     */
    readonly forceDelete?: pulumi.Input<boolean>;
    /**
     * (Optional; Required if `image_name` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this creates a new server.
     */
    readonly imageId?: pulumi.Input<string>;
    /**
     * (Optional; Required if `image_id` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this creates a new server.
     */
    readonly imageName?: pulumi.Input<string>;
    /**
     * The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant's account.
     * Changing this creates a new server.
     */
    readonly keyPair?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     */
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The human-readable
     * name of the network. Changing this creates a new server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     */
    readonly networks?: pulumi.Input<{
        accessNetwork?: pulumi.Input<boolean>;
        fixedIpV4?: pulumi.Input<string>;
        fixedIpV6?: pulumi.Input<string>;
        floatingIp?: pulumi.Input<string>;
        mac?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
        port?: pulumi.Input<string>;
        uuid?: pulumi.Input<string>;
    }[]>;
    /**
     * Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below.
     */
    readonly personalities?: pulumi.Input<{
        content: pulumi.Input<string>;
        file: pulumi.Input<string>;
    }[]>;
    /**
     * The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     */
    readonly schedulerHints?: pulumi.Input<{
        additionalProperties?: pulumi.Input<{
            [key: string]: any;
        }>;
        buildNearHostIp?: pulumi.Input<string>;
        differentHosts?: pulumi.Input<pulumi.Input<string>[]>;
        group?: pulumi.Input<string>;
        queries?: pulumi.Input<pulumi.Input<string>[]>;
        sameHosts?: pulumi.Input<pulumi.Input<string>[]>;
        targetCell?: pulumi.Input<string>;
    }[]>;
    /**
     * An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn't stop within timeout, it will be destroyed anyway.
     */
    readonly stopBeforeDestroy?: pulumi.Input<boolean>;
    /**
     * The user data to provide when launching the instance.
     * Changing this creates a new server.
     */
    readonly userData?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The first detected Fixed IPv4 address _or_ the
     * Floating IP.
     */
    readonly accessIpV4?: pulumi.Input<string>;
    /**
     * The first detected Fixed IPv6 address.
     */
    readonly accessIpV6?: pulumi.Input<string>;
    /**
     * The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     */
    readonly adminPass?: pulumi.Input<string>;
    /**
     * The availability zone in which to create
     * the server. Changing this creates a new server.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Configuration of block devices. The block_device
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](http://docs.openstack.org/developer/nova/block_device_mapping.html)
     * for more information.
     */
    readonly blockDevices?: pulumi.Input<{
        bootIndex?: pulumi.Input<number>;
        deleteOnTermination?: pulumi.Input<boolean>;
        destinationType?: pulumi.Input<string>;
        guestFormat?: pulumi.Input<string>;
        sourceType: pulumi.Input<string>;
        uuid?: pulumi.Input<string>;
        volumeSize?: pulumi.Input<number>;
    }[]>;
    /**
     * Whether to use the config_drive feature to
     * configure the instance. Changing this creates a new server.
     */
    readonly configDrive?: pulumi.Input<boolean>;
    /**
     * The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     */
    readonly flavorId?: pulumi.Input<string>;
    /**
     * The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     */
    readonly flavorName?: pulumi.Input<string>;
    /**
     * Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     */
    readonly forceDelete?: pulumi.Input<boolean>;
    /**
     * (Optional; Required if `image_name` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this creates a new server.
     */
    readonly imageId?: pulumi.Input<string>;
    /**
     * (Optional; Required if `image_id` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this creates a new server.
     */
    readonly imageName?: pulumi.Input<string>;
    /**
     * The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant's account.
     * Changing this creates a new server.
     */
    readonly keyPair?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     */
    readonly metadata?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The human-readable
     * name of the network. Changing this creates a new server.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     */
    readonly networks?: pulumi.Input<{
        accessNetwork?: pulumi.Input<boolean>;
        fixedIpV4?: pulumi.Input<string>;
        fixedIpV6?: pulumi.Input<string>;
        floatingIp?: pulumi.Input<string>;
        mac?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
        port?: pulumi.Input<string>;
        uuid?: pulumi.Input<string>;
    }[]>;
    /**
     * Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below.
     */
    readonly personalities?: pulumi.Input<{
        content: pulumi.Input<string>;
        file: pulumi.Input<string>;
    }[]>;
    /**
     * The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     */
    readonly schedulerHints?: pulumi.Input<{
        additionalProperties?: pulumi.Input<{
            [key: string]: any;
        }>;
        buildNearHostIp?: pulumi.Input<string>;
        differentHosts?: pulumi.Input<pulumi.Input<string>[]>;
        group?: pulumi.Input<string>;
        queries?: pulumi.Input<pulumi.Input<string>[]>;
        sameHosts?: pulumi.Input<pulumi.Input<string>[]>;
        targetCell?: pulumi.Input<string>;
    }[]>;
    /**
     * An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn't stop within timeout, it will be destroyed anyway.
     */
    readonly stopBeforeDestroy?: pulumi.Input<boolean>;
    /**
     * The user data to provide when launching the instance.
     * Changing this creates a new server.
     */
    readonly userData?: pulumi.Input<string>;
}
