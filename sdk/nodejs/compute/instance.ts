// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'openstack:compute/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The first detected Fixed IPv4 address.
     */
    public readonly accessIpV4!: pulumi.Output<string>;
    /**
     * The first detected Fixed IPv6 address.
     */
    public readonly accessIpV6!: pulumi.Output<string>;
    /**
     * The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     */
    public readonly adminPass!: pulumi.Output<string | undefined>;
    public /*out*/ readonly allMetadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * The collection of tags assigned on the instance, which have
     * been explicitly and implicitly added.
     */
    public /*out*/ readonly allTags!: pulumi.Output<string[]>;
    /**
     * The availability zone in which to create
     * the server. Conflicts with `availabilityZoneHints`. Changing this creates
     * a new server.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The availability zone in which to
     * create the server. This argument is preferred to `availabilityZone`, when
     * scheduling the server on a
     * [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
     * host or node. Conflicts with `availabilityZone`. Changing this creates a
     * new server.
     */
    public readonly availabilityZoneHints!: pulumi.Output<string | undefined>;
    /**
     * Configuration of block devices. The blockDevice
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
     * for more information.
     */
    public readonly blockDevices!: pulumi.Output<outputs.compute.InstanceBlockDevice[] | undefined>;
    /**
     * Whether to use the configDrive feature to
     * configure the instance. Changing this creates a new server.
     */
    public readonly configDrive!: pulumi.Output<boolean | undefined>;
    /**
     * The creation time of the instance.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     */
    public readonly flavorId!: pulumi.Output<string>;
    /**
     * The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     */
    public readonly flavorName!: pulumi.Output<string>;
    /**
     * @deprecated Use the openstack.compute.FloatingIpAssociate resource instead
     */
    public readonly floatingIp!: pulumi.Output<string | undefined>;
    /**
     * Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     */
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional; Required if `imageName` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this rebuilds the existing
     * server.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * (Optional; Required if `imageId` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this rebuilds the existing server.
     */
    public readonly imageName!: pulumi.Output<string>;
    /**
     * The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant's account.
     * Changing this creates a new server.
     */
    public readonly keyPair!: pulumi.Output<string | undefined>;
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A unique name for the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Special string for `network` option to create
     * the server. `networkMode` can be `"auto"` or `"none"`.
     * Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
     */
    public readonly networkMode!: pulumi.Output<string | undefined>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     */
    public readonly networks!: pulumi.Output<outputs.compute.InstanceNetwork[]>;
    /**
     * Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below. Changing this rebuilds the existing server.
     */
    public readonly personalities!: pulumi.Output<outputs.compute.InstancePersonality[] | undefined>;
    /**
     * Provide the VM state. Only 'active', 'shutoff'
     * and 'shelved_offloaded' are supported values.
     * *Note*: If the initial powerState is the shutoff
     * the VM will be stopped immediately after build and the provisioners like
     * remote-exec or files are not supported.
     */
    public readonly powerState!: pulumi.Output<string | undefined>;
    /**
     * The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     */
    public readonly schedulerHints!: pulumi.Output<outputs.compute.InstanceSchedulerHint[] | undefined>;
    /**
     * An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance. *Note*: Names should be used and not ids, as ids
     * trigger unnecessary updates.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn't stop within timeout, it will be destroyed anyway.
     */
    public readonly stopBeforeDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * A set of string tags for the instance. Changing this
     * updates the existing instance tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The time when the instance was last updated.
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;
    /**
     * The user data to provide when launching the instance.
     * Changing this creates a new server.
     */
    public readonly userData!: pulumi.Output<string | undefined>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    public readonly vendorOptions!: pulumi.Output<outputs.compute.InstanceVendorOptions | undefined>;
    /**
     * @deprecated Use blockDevice or openstack.compute.VolumeAttach instead
     */
    public readonly volumes!: pulumi.Output<outputs.compute.InstanceVolume[] | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["accessIpV4"] = state ? state.accessIpV4 : undefined;
            resourceInputs["accessIpV6"] = state ? state.accessIpV6 : undefined;
            resourceInputs["adminPass"] = state ? state.adminPass : undefined;
            resourceInputs["allMetadata"] = state ? state.allMetadata : undefined;
            resourceInputs["allTags"] = state ? state.allTags : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["availabilityZoneHints"] = state ? state.availabilityZoneHints : undefined;
            resourceInputs["blockDevices"] = state ? state.blockDevices : undefined;
            resourceInputs["configDrive"] = state ? state.configDrive : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["flavorId"] = state ? state.flavorId : undefined;
            resourceInputs["flavorName"] = state ? state.flavorName : undefined;
            resourceInputs["floatingIp"] = state ? state.floatingIp : undefined;
            resourceInputs["forceDelete"] = state ? state.forceDelete : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["imageName"] = state ? state.imageName : undefined;
            resourceInputs["keyPair"] = state ? state.keyPair : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkMode"] = state ? state.networkMode : undefined;
            resourceInputs["networks"] = state ? state.networks : undefined;
            resourceInputs["personalities"] = state ? state.personalities : undefined;
            resourceInputs["powerState"] = state ? state.powerState : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["schedulerHints"] = state ? state.schedulerHints : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["stopBeforeDestroy"] = state ? state.stopBeforeDestroy : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updated"] = state ? state.updated : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["vendorOptions"] = state ? state.vendorOptions : undefined;
            resourceInputs["volumes"] = state ? state.volumes : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            resourceInputs["accessIpV4"] = args ? args.accessIpV4 : undefined;
            resourceInputs["accessIpV6"] = args ? args.accessIpV6 : undefined;
            resourceInputs["adminPass"] = args?.adminPass ? pulumi.secret(args.adminPass) : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["availabilityZoneHints"] = args ? args.availabilityZoneHints : undefined;
            resourceInputs["blockDevices"] = args ? args.blockDevices : undefined;
            resourceInputs["configDrive"] = args ? args.configDrive : undefined;
            resourceInputs["flavorId"] = args ? args.flavorId : undefined;
            resourceInputs["flavorName"] = args ? args.flavorName : undefined;
            resourceInputs["floatingIp"] = args ? args.floatingIp : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["keyPair"] = args ? args.keyPair : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkMode"] = args ? args.networkMode : undefined;
            resourceInputs["networks"] = args ? args.networks : undefined;
            resourceInputs["personalities"] = args ? args.personalities : undefined;
            resourceInputs["powerState"] = args ? args.powerState : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["schedulerHints"] = args ? args.schedulerHints : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["stopBeforeDestroy"] = args ? args.stopBeforeDestroy : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["vendorOptions"] = args ? args.vendorOptions : undefined;
            resourceInputs["volumes"] = args ? args.volumes : undefined;
            resourceInputs["allMetadata"] = undefined /*out*/;
            resourceInputs["allTags"] = undefined /*out*/;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminPass"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The first detected Fixed IPv4 address.
     */
    accessIpV4?: pulumi.Input<string>;
    /**
     * The first detected Fixed IPv6 address.
     */
    accessIpV6?: pulumi.Input<string>;
    /**
     * The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     */
    adminPass?: pulumi.Input<string>;
    allMetadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The collection of tags assigned on the instance, which have
     * been explicitly and implicitly added.
     */
    allTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The availability zone in which to create
     * the server. Conflicts with `availabilityZoneHints`. Changing this creates
     * a new server.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The availability zone in which to
     * create the server. This argument is preferred to `availabilityZone`, when
     * scheduling the server on a
     * [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
     * host or node. Conflicts with `availabilityZone`. Changing this creates a
     * new server.
     */
    availabilityZoneHints?: pulumi.Input<string>;
    /**
     * Configuration of block devices. The blockDevice
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
     * for more information.
     */
    blockDevices?: pulumi.Input<pulumi.Input<inputs.compute.InstanceBlockDevice>[]>;
    /**
     * Whether to use the configDrive feature to
     * configure the instance. Changing this creates a new server.
     */
    configDrive?: pulumi.Input<boolean>;
    /**
     * The creation time of the instance.
     */
    created?: pulumi.Input<string>;
    /**
     * The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     */
    flavorName?: pulumi.Input<string>;
    /**
     * @deprecated Use the openstack.compute.FloatingIpAssociate resource instead
     */
    floatingIp?: pulumi.Input<string>;
    /**
     * Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * (Optional; Required if `imageName` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this rebuilds the existing
     * server.
     */
    imageId?: pulumi.Input<string>;
    /**
     * (Optional; Required if `imageId` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this rebuilds the existing server.
     */
    imageName?: pulumi.Input<string>;
    /**
     * The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant's account.
     * Changing this creates a new server.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Special string for `network` option to create
     * the server. `networkMode` can be `"auto"` or `"none"`.
     * Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
     */
    networkMode?: pulumi.Input<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceNetwork>[]>;
    /**
     * Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below. Changing this rebuilds the existing server.
     */
    personalities?: pulumi.Input<pulumi.Input<inputs.compute.InstancePersonality>[]>;
    /**
     * Provide the VM state. Only 'active', 'shutoff'
     * and 'shelved_offloaded' are supported values.
     * *Note*: If the initial powerState is the shutoff
     * the VM will be stopped immediately after build and the provisioners like
     * remote-exec or files are not supported.
     */
    powerState?: pulumi.Input<string>;
    /**
     * The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     */
    region?: pulumi.Input<string>;
    /**
     * Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     */
    schedulerHints?: pulumi.Input<pulumi.Input<inputs.compute.InstanceSchedulerHint>[]>;
    /**
     * An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance. *Note*: Names should be used and not ids, as ids
     * trigger unnecessary updates.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn't stop within timeout, it will be destroyed anyway.
     */
    stopBeforeDestroy?: pulumi.Input<boolean>;
    /**
     * A set of string tags for the instance. Changing this
     * updates the existing instance tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The time when the instance was last updated.
     */
    updated?: pulumi.Input<string>;
    /**
     * The user data to provide when launching the instance.
     * Changing this creates a new server.
     */
    userData?: pulumi.Input<string>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    vendorOptions?: pulumi.Input<inputs.compute.InstanceVendorOptions>;
    /**
     * @deprecated Use blockDevice or openstack.compute.VolumeAttach instead
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.compute.InstanceVolume>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The first detected Fixed IPv4 address.
     */
    accessIpV4?: pulumi.Input<string>;
    /**
     * The first detected Fixed IPv6 address.
     */
    accessIpV6?: pulumi.Input<string>;
    /**
     * The administrative password to assign to the server.
     * Changing this changes the root password on the existing server.
     */
    adminPass?: pulumi.Input<string>;
    /**
     * The availability zone in which to create
     * the server. Conflicts with `availabilityZoneHints`. Changing this creates
     * a new server.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The availability zone in which to
     * create the server. This argument is preferred to `availabilityZone`, when
     * scheduling the server on a
     * [particular](https://docs.openstack.org/nova/latest/admin/availability-zones.html)
     * host or node. Conflicts with `availabilityZone`. Changing this creates a
     * new server.
     */
    availabilityZoneHints?: pulumi.Input<string>;
    /**
     * Configuration of block devices. The blockDevice
     * structure is documented below. Changing this creates a new server.
     * You can specify multiple block devices which will create an instance with
     * multiple disks. This configuration is very flexible, so please see the
     * following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html)
     * for more information.
     */
    blockDevices?: pulumi.Input<pulumi.Input<inputs.compute.InstanceBlockDevice>[]>;
    /**
     * Whether to use the configDrive feature to
     * configure the instance. Changing this creates a new server.
     */
    configDrive?: pulumi.Input<boolean>;
    /**
     * The flavor ID of
     * the desired flavor for the server. Changing this resizes the existing server.
     */
    flavorId?: pulumi.Input<string>;
    /**
     * The name of the
     * desired flavor for the server. Changing this resizes the existing server.
     */
    flavorName?: pulumi.Input<string>;
    /**
     * @deprecated Use the openstack.compute.FloatingIpAssociate resource instead
     */
    floatingIp?: pulumi.Input<string>;
    /**
     * Whether to force the OpenStack instance to be
     * forcefully deleted. This is useful for environments that have reclaim / soft
     * deletion enabled.
     */
    forceDelete?: pulumi.Input<boolean>;
    /**
     * (Optional; Required if `imageName` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The image ID of
     * the desired image for the server. Changing this rebuilds the existing
     * server.
     */
    imageId?: pulumi.Input<string>;
    /**
     * (Optional; Required if `imageId` is empty and not booting
     * from a volume. Do not specify if booting from a volume.) The name of the
     * desired image for the server. Changing this rebuilds the existing server.
     */
    imageName?: pulumi.Input<string>;
    /**
     * The name of a key pair to put on the server. The key
     * pair must already be created and associated with the tenant's account.
     * Changing this creates a new server.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs to make available from
     * within the instance. Changing this updates the existing server metadata.
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Special string for `network` option to create
     * the server. `networkMode` can be `"auto"` or `"none"`.
     * Please see the following [reference](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail#id11) for more information. Conflicts with `network`.
     */
    networkMode?: pulumi.Input<string>;
    /**
     * An array of one or more networks to attach to the
     * instance. The network object structure is documented below. Changing this
     * creates a new server.
     */
    networks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceNetwork>[]>;
    /**
     * Customize the personality of an instance by
     * defining one or more files and their contents. The personality structure
     * is described below. Changing this rebuilds the existing server.
     */
    personalities?: pulumi.Input<pulumi.Input<inputs.compute.InstancePersonality>[]>;
    /**
     * Provide the VM state. Only 'active', 'shutoff'
     * and 'shelved_offloaded' are supported values.
     * *Note*: If the initial powerState is the shutoff
     * the VM will be stopped immediately after build and the provisioners like
     * remote-exec or files are not supported.
     */
    powerState?: pulumi.Input<string>;
    /**
     * The region in which to create the server instance. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new server.
     */
    region?: pulumi.Input<string>;
    /**
     * Provide the Nova scheduler with hints on how
     * the instance should be launched. The available hints are described below.
     */
    schedulerHints?: pulumi.Input<pulumi.Input<inputs.compute.InstanceSchedulerHint>[]>;
    /**
     * An array of one or more security group names
     * to associate with the server. Changing this results in adding/removing
     * security groups from the existing server. *Note*: When attaching the
     * instance to networks using Ports, place the security groups on the Port
     * and not the instance. *Note*: Names should be used and not ids, as ids
     * trigger unnecessary updates.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to try stop instance gracefully
     * before destroying it, thus giving chance for guest OS daemons to stop correctly.
     * If instance doesn't stop within timeout, it will be destroyed anyway.
     */
    stopBeforeDestroy?: pulumi.Input<boolean>;
    /**
     * A set of string tags for the instance. Changing this
     * updates the existing instance tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user data to provide when launching the instance.
     * Changing this creates a new server.
     */
    userData?: pulumi.Input<string>;
    /**
     * Map of additional vendor-specific options.
     * Supported options are described below.
     */
    vendorOptions?: pulumi.Input<inputs.compute.InstanceVendorOptions>;
    /**
     * @deprecated Use blockDevice or openstack.compute.VolumeAttach instead
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.compute.InstanceVolume>[]>;
}
