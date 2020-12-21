// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export namespace blockstorage {
    export interface VolumeAttachment {
        device?: pulumi.Input<string>;
        id?: pulumi.Input<string>;
        instanceId?: pulumi.Input<string>;
    }

    export interface VolumeSchedulerHint {
        /**
         * Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         */
        additionalProperties?: pulumi.Input<{[key: string]: any}>;
        /**
         * The volume should be scheduled on a 
         * different host from the set of volumes specified in the list provided.
         */
        differentHosts?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * An instance UUID. The volume should be 
         * scheduled on the same host as the instance.
         */
        localToInstance?: pulumi.Input<string>;
        /**
         * A conditional query that a back-end must pass in
         * order to host a volume. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         */
        query?: pulumi.Input<string>;
        /**
         * A list of volume UUIDs. The volume should be
         * scheduled on the same host as another volume specified in the list provided.
         */
        sameHosts?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface VolumeV1Attachment {
        device?: pulumi.Input<string>;
        id?: pulumi.Input<string>;
        instanceId?: pulumi.Input<string>;
    }

    export interface VolumeV2Attachment {
        device?: pulumi.Input<string>;
        id?: pulumi.Input<string>;
        instanceId?: pulumi.Input<string>;
    }

    export interface VolumeV2SchedulerHint {
        /**
         * Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         */
        additionalProperties?: pulumi.Input<{[key: string]: any}>;
        /**
         * The volume should be scheduled on a 
         * different host from the set of volumes specified in the list provided.
         */
        differentHosts?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * An instance UUID. The volume should be 
         * scheduled on the same host as the instance.
         */
        localToInstance?: pulumi.Input<string>;
        /**
         * A conditional query that a back-end must pass in
         * order to host a volume. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         */
        query?: pulumi.Input<string>;
        /**
         * A list of volume UUIDs. The volume should be
         * scheduled on the same host as another volume specified in the list provided.
         */
        sameHosts?: pulumi.Input<pulumi.Input<string>[]>;
    }
}

export namespace compute {
    export interface GetInstanceV2Network {
        /**
         * The IPv4 address assigned to this network port.
         */
        fixedIpV4?: string;
        /**
         * The IPv6 address assigned to this network port.
         */
        fixedIpV6?: string;
        /**
         * The MAC address assigned to this network interface.
         */
        mac?: string;
        /**
         * The name of the network
         */
        name?: string;
        /**
         * The port UUID for this network
         */
        port?: string;
        /**
         * The UUID of the network
         */
        uuid?: string;
    }

    export interface InstanceBlockDevice {
        /**
         * The boot index of the volume. It defaults to 0.
         * Changing this creates a new server.
         */
        bootIndex?: pulumi.Input<number>;
        /**
         * Delete the volume / block device upon
         * termination of the instance. Defaults to false. Changing this creates a
         * new server.
         */
        deleteOnTermination?: pulumi.Input<boolean>;
        /**
         * The type that gets created. Possible values
         * are "volume" and "local". Changing this creates a new server.
         */
        destinationType?: pulumi.Input<string>;
        /**
         * The low-level device type that will be used. Most
         * common thing is to leave this empty. Changing this creates a new server.
         */
        deviceType?: pulumi.Input<string>;
        /**
         * The low-level disk bus that will be used. Most common
         * thing is to leave this empty. Changing this creates a new server.
         */
        diskBus?: pulumi.Input<string>;
        /**
         * Specifies the guest server disk file system format,
         * such as `ext2`, `ext3`, `ext4`, `xfs` or `swap`. Swap block device mappings
         * have the following restrictions: sourceType must be blank and destinationType
         * must be local and only one swap disk per server and the size of the swap disk
         * must be less than or equal to the swap size of the flavor. Changing this
         * creates a new server.
         */
        guestFormat?: pulumi.Input<string>;
        /**
         * The source type of the device. Must be one of
         * "blank", "image", "volume", or "snapshot". Changing this creates a new
         * server.
         */
        sourceType: pulumi.Input<string>;
        /**
         * The UUID of
         * the image, volume, or snapshot. Changing this creates a new server.
         */
        uuid?: pulumi.Input<string>;
        /**
         * The size of the volume to create (in gigabytes). Required
         * in the following combinations: source=image and destination=volume,
         * source=blank and destination=local, and source=blank and destination=volume.
         * Changing this creates a new server.
         */
        volumeSize?: pulumi.Input<number>;
        /**
         * The volume type that will be used, for example SSD
         * or HDD storage. The available options depend on how your specific OpenStack
         * cloud is configured and what classes of storage are provided. Changing this
         * creates a new server.
         */
        volumeType?: pulumi.Input<string>;
    }

    export interface InstanceNetwork {
        /**
         * Specifies if this network should be used for
         * provisioning access. Accepts true or false. Defaults to false.
         */
        accessNetwork?: pulumi.Input<boolean>;
        /**
         * Specifies a fixed IPv4 address to be used on this
         * network. Changing this creates a new server.
         */
        fixedIpV4?: pulumi.Input<string>;
        fixedIpV6?: pulumi.Input<string>;
        floatingIp?: pulumi.Input<string>;
        mac?: pulumi.Input<string>;
        /**
         * The human-readable
         * name of the network. Changing this creates a new server.
         */
        name?: pulumi.Input<string>;
        /**
         * The port UUID of a
         * network to attach to the server. Changing this creates a new server.
         */
        port?: pulumi.Input<string>;
        /**
         * The UUID of
         * the image, volume, or snapshot. Changing this creates a new server.
         */
        uuid?: pulumi.Input<string>;
    }

    export interface InstancePersonality {
        /**
         * The contents of the file. Limited to 255 bytes.
         */
        content: pulumi.Input<string>;
        /**
         * The absolute path of the destination file.
         */
        file: pulumi.Input<string>;
    }

    export interface InstanceSchedulerHint {
        /**
         * Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         */
        additionalProperties?: pulumi.Input<{[key: string]: any}>;
        /**
         * An IP Address in CIDR form. The instance
         * will be placed on a compute node that is in the same subnet.
         */
        buildNearHostIp?: pulumi.Input<string>;
        /**
         * The names of cells where not to build the instance.
         */
        differentCells?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * A list of instance UUIDs. The instance will
         * be scheduled on a different host than all other instances.
         */
        differentHosts?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * A UUID of a Server Group. The instance will be placed
         * into that group.
         */
        group?: pulumi.Input<string>;
        /**
         * A conditional query that a compute node must pass in
         * order to host an instance. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         */
        queries?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * A list of instance UUIDs. The instance will be
         * scheduled on the same host of those specified.
         */
        sameHosts?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * The name of a cell to host the instance.
         */
        targetCell?: pulumi.Input<string>;
    }

    export interface InstanceVendorOptions {
        /**
         * Whether to try to detach all attached
         * ports to the vm before destroying it to make sure the port state is correct
         * after the vm destruction. This is helpful when the port is not deleted.
         */
        detachPortsBeforeDestroy?: pulumi.Input<boolean>;
        /**
         * Boolean to control whether
         * to ignore manual confirmation of the instance resizing. This can be helpful
         * to work with some OpenStack clouds which automatically confirm resizing of
         * instances after some timeout.
         */
        ignoreResizeConfirmation?: pulumi.Input<boolean>;
    }

    export interface SecGroupRule {
        /**
         * Required if `fromGroupId` or `self` is empty. The IP range
         * that will be the source of network traffic to the security group. Use 0.0.0.0/0
         * to allow all IP addresses. Changing this creates a new security group rule. Cannot
         * be combined with `fromGroupId` or `self`.
         */
        cidr?: pulumi.Input<string>;
        /**
         * Required if `cidr` or `self` is empty. The ID of a
         * group from which to forward traffic to the parent group. Changing this creates a
         * new security group rule. Cannot be combined with `cidr` or `self`.
         */
        fromGroupId?: pulumi.Input<string>;
        /**
         * An integer representing the lower bound of the port
         * range to open. Changing this creates a new security group rule.
         */
        fromPort: pulumi.Input<number>;
        id?: pulumi.Input<string>;
        /**
         * The protocol type that will be allowed. Changing
         * this creates a new security group rule.
         */
        ipProtocol: pulumi.Input<string>;
        /**
         * Required if `cidr` and `fromGroupId` is empty. If true,
         * the security group itself will be added as a source to this ingress rule. Cannot
         * be combined with `cidr` or `fromGroupId`.
         */
        self?: pulumi.Input<boolean>;
        /**
         * An integer representing the upper bound of the port
         * range to open. Changing this creates a new security group rule.
         */
        toPort: pulumi.Input<number>;
    }
}

export namespace containerinfra {
    export interface ClusterKubeconfig {
        clientCertificate?: pulumi.Input<string>;
        clientKey?: pulumi.Input<string>;
        clusterCaCertificate?: pulumi.Input<string>;
        host?: pulumi.Input<string>;
        rawConfig?: pulumi.Input<string>;
    }
}

export namespace database {
    export interface ConfigurationConfiguration {
        /**
         * Configuration parameter name. Changing this creates a new resource.
         */
        name: pulumi.Input<string>;
        /**
         * Configuration parameter value. Changing this creates a new resource.
         */
        value: pulumi.Input<string>;
    }

    export interface ConfigurationDatastore {
        /**
         * Database engine type to be used with this configuration. Changing this creates a new resource.
         */
        type: pulumi.Input<string>;
        /**
         * Version of database engine type to be used with this configuration. Changing this creates a new resource.
         */
        version: pulumi.Input<string>;
    }

    export interface InstanceDatabase {
        /**
         * Database character set. Changing this creates a
         * new instance.
         */
        charset?: pulumi.Input<string>;
        /**
         * Database collation. Changing this creates a new instance.
         */
        collate?: pulumi.Input<string>;
        /**
         * Database to be created on new instance. Changing this creates a
         * new instance.
         */
        name: pulumi.Input<string>;
    }

    export interface InstanceDatastore {
        /**
         * Database engine type to be used in new instance. Changing this
         * creates a new instance.
         */
        type: pulumi.Input<string>;
        /**
         * Version of database engine type to be used in new instance.
         * Changing this creates a new instance.
         */
        version: pulumi.Input<string>;
    }

    export interface InstanceNetwork {
        /**
         * Specifies a fixed IPv4 address to be used on this
         * network. Changing this creates a new instance.
         */
        fixedIpV4?: pulumi.Input<string>;
        /**
         * Specifies a fixed IPv6 address to be used on this
         * network. Changing this creates a new instance.
         */
        fixedIpV6?: pulumi.Input<string>;
        /**
         * The port UUID of a
         * network to attach to the instance. Changing this creates a new instance.
         */
        port?: pulumi.Input<string>;
        /**
         * The network UUID to
         * attach to the instance. Changing this creates a new instance.
         */
        uuid?: pulumi.Input<string>;
    }

    export interface InstanceUser {
        /**
         * A list of databases that user will have access to. If not specified, 
         * user has access to all databases on th einstance. Changing this creates a new instance.
         */
        databases?: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * An ip address or % sign indicating what ip addresses can connect with
         * this user credentials. Changing this creates a new instance.
         */
        host?: pulumi.Input<string>;
        /**
         * Database to be created on new instance. Changing this creates a
         * new instance.
         */
        name: pulumi.Input<string>;
        /**
         * User's password. Changing this creates a
         * new instance.
         */
        password?: pulumi.Input<string>;
    }
}

export namespace identity {
    export interface ApplicationCredentialAccessRule {
        /**
         * The ID of the existing access rule. The access rule ID of
         * another application credential can be provided.
         */
        id?: pulumi.Input<string>;
        /**
         * The request method that the application credential is
         * permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
         * `HEAD`, `PATCH`, `PUT` and `DELETE`.
         */
        method: pulumi.Input<string>;
        /**
         * The API path that the application credential is permitted
         * to access. May use named wildcards such as **{tag}** or the unnamed wildcard
         * **\*** to match against any string in the path up to a **&#47;**, or the recursive
         * wildcard **\*\*** to include **&#47;** in the matched path.
         */
        path: pulumi.Input<string>;
        /**
         * The service type identifier for the service that the
         * application credential is granted to access. Must be a service type that is
         * listed in the service catalog and not a code name for a service. E.g.
         * **identity**, **compute**, **volumev3**, **image**, **network**,
         * **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
         */
        service: pulumi.Input<string>;
    }

    export interface UserMultiFactorAuthRule {
        /**
         * A list of authentication plugins that the user must
         * authenticate with.
         */
        rules: pulumi.Input<pulumi.Input<string>[]>;
    }
}

export namespace keymanager {
    export interface ContainerV1Acl {
        read?: pulumi.Input<inputs.keymanager.ContainerV1AclRead>;
    }

    export interface ContainerV1AclRead {
        /**
         * The date the container ACL was created.
         */
        createdAt?: pulumi.Input<string>;
        /**
         * Whether the container is accessible project wide.
         * Defaults to `true`.
         */
        projectAccess?: pulumi.Input<boolean>;
        /**
         * The date the container ACL was last updated.
         */
        updatedAt?: pulumi.Input<string>;
        /**
         * The list of user IDs, which are allowed to access the
         * container, when `projectAccess` is set to `false`.
         */
        users?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ContainerV1Consumer {
        /**
         * The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
         */
        name?: pulumi.Input<string>;
        /**
         * The consumer URL.
         */
        url?: pulumi.Input<string>;
    }

    export interface ContainerV1SecretRef {
        /**
         * The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
         */
        name?: pulumi.Input<string>;
        /**
         * The secret reference / where to find the secret, URL.
         */
        secretRef: pulumi.Input<string>;
    }

    export interface OrderV1Meta {
        /**
         * Algorithm to use for key generation.
         */
        algorithm: pulumi.Input<string>;
        /**
         * - Bit lenght of key to be generated.
         */
        bitLength: pulumi.Input<number>;
        /**
         * This is a UTC timestamp in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ. If set, the secret will not be available after this time.
         */
        expiration?: pulumi.Input<string>;
        /**
         * The mode to use for key generation.
         */
        mode?: pulumi.Input<string>;
        /**
         * The name of the secret set by the user.
         */
        name?: pulumi.Input<string>;
        /**
         * The media type for the content of the secrets payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
         */
        payloadContentType?: pulumi.Input<string>;
    }

    export interface SecretV1Acl {
        read?: pulumi.Input<inputs.keymanager.SecretV1AclRead>;
    }

    export interface SecretV1AclRead {
        /**
         * The date the secret ACL was created.
         */
        createdAt?: pulumi.Input<string>;
        /**
         * Whether the secret is accessible project wide.
         * Defaults to `true`.
         */
        projectAccess?: pulumi.Input<boolean>;
        /**
         * The date the secret ACL was last updated.
         */
        updatedAt?: pulumi.Input<string>;
        /**
         * The list of user IDs, which are allowed to access the
         * secret, when `projectAccess` is set to `false`.
         */
        users?: pulumi.Input<pulumi.Input<string>[]>;
    }
}

export namespace loadbalancer {
    export interface MembersMember {
        /**
         * The IP address of the members to receive traffic from
         * the load balancer.
         */
        address: pulumi.Input<string>;
        /**
         * The administrative state of the member.
         * A valid value is true (UP) or false (DOWN). Defaults to true.
         */
        adminStateUp?: pulumi.Input<boolean>;
        /**
         * The unique ID for the members.
         */
        id?: pulumi.Input<string>;
        /**
         * Human-readable name for the member.
         */
        name?: pulumi.Input<string>;
        /**
         * The port on which to listen for client traffic.
         */
        protocolPort: pulumi.Input<number>;
        /**
         * The subnet in which to access the member.
         */
        subnetId?: pulumi.Input<string>;
        /**
         * A positive integer value that indicates the relative
         * portion of traffic that this members should receive from the pool. For
         * example, a member with a weight of 10 receives five times as much traffic
         * as a member with a weight of 2. Defaults to 1.
         */
        weight?: pulumi.Input<number>;
    }

    export interface PoolPersistence {
        /**
         * The name of the cookie if persistence mode is set
         * appropriately. Required if `type = APP_COOKIE`.
         */
        cookieName?: pulumi.Input<string>;
        /**
         * The type of persistence mode. The current specification
         * supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
         */
        type: pulumi.Input<string>;
    }
}

export namespace networking {
    export interface NetworkSegment {
        /**
         * The type of physical network.
         */
        networkType?: pulumi.Input<string>;
        /**
         * The physical network where this network is implemented.
         */
        physicalNetwork?: pulumi.Input<string>;
        /**
         * An isolated segment on the physical network.
         */
        segmentationId?: pulumi.Input<number>;
    }

    export interface PortAllowedAddressPair {
        /**
         * The additional IP address.
         */
        ipAddress: pulumi.Input<string>;
        /**
         * The additional MAC address.
         */
        macAddress?: pulumi.Input<string>;
    }

    export interface PortBinding {
        /**
         * The ID of the host to allocate port on.
         */
        hostId?: pulumi.Input<string>;
        /**
         * Custom data to be passed as `binding:profile`. Data
         * must be passed as JSON.
         */
        profile?: pulumi.Input<string>;
        /**
         * A map of JSON strings containing additional
         * details for this specific binding.
         */
        vifDetails?: pulumi.Input<{[key: string]: any}>;
        /**
         * The VNIC type of the port binding.
         */
        vifType?: pulumi.Input<string>;
        /**
         * VNIC type for the port. Can either be `direct`,
         * `direct-physical`, `macvtap`, `normal`, `baremetal` or `virtio-forwarder`.
         * Default value is `normal`.
         */
        vnicType?: pulumi.Input<string>;
    }

    export interface PortExtraDhcpOption {
        /**
         * IP protocol version. Defaults to 4.
         */
        ipVersion?: pulumi.Input<number>;
        /**
         * Name of the DHCP option.
         */
        name: pulumi.Input<string>;
        /**
         * Value of the DHCP option.
         */
        value: pulumi.Input<string>;
    }

    export interface PortFixedIp {
        /**
         * The additional IP address.
         */
        ipAddress?: pulumi.Input<string>;
        /**
         * Subnet in which to allocate IP address for
         * this port.
         */
        subnetId: pulumi.Input<string>;
    }

    export interface RouterExternalFixedIp {
        /**
         * The IP address to set on the router.
         */
        ipAddress?: pulumi.Input<string>;
        /**
         * Subnet in which the fixed IP belongs to.
         */
        subnetId?: pulumi.Input<string>;
    }

    export interface RouterVendorOptions {
        /**
         * Boolean to control whether
         * the Router gateway is assigned during creation or updated after creation.
         */
        setRouterGatewayAfterCreate?: pulumi.Input<boolean>;
    }

    export interface SubnetAllocationPool {
        /**
         * The ending address.
         */
        end: pulumi.Input<string>;
        /**
         * The starting address.
         */
        start: pulumi.Input<string>;
    }

    export interface SubnetAllocationPoolsCollection {
        /**
         * The ending address.
         */
        end: pulumi.Input<string>;
        /**
         * The starting address.
         */
        start: pulumi.Input<string>;
    }

    export interface SubnetHostRoute {
        /**
         * The destination CIDR.
         */
        destinationCidr: pulumi.Input<string>;
        /**
         * The next hop in the route.
         */
        nextHop: pulumi.Input<string>;
    }

    export interface TrunkSubPort {
        /**
         * The ID of the port to be made a subport of the trunk.
         */
        portId: pulumi.Input<string>;
        /**
         * The numeric id of the subport segment.
         */
        segmentationId: pulumi.Input<number>;
        /**
         * The segmentation technology to use, e.g., "vlan".
         */
        segmentationType: pulumi.Input<string>;
    }
}

export namespace objectstorage {
    export interface ContainerVersioning {
        /**
         * Container in which versions will be stored.
         */
        location: pulumi.Input<string>;
        /**
         * Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
         */
        type: pulumi.Input<string>;
    }
}

export namespace orchestration {
    export interface StackV1StackOutput {
        /**
         * The description of the stack resource.
         */
        description?: pulumi.Input<string>;
        outputKey: pulumi.Input<string>;
        outputValue: pulumi.Input<string>;
    }
}

export namespace sharedfilesystem {
    export interface ShareExportLocation {
        path?: pulumi.Input<string>;
        preferred?: pulumi.Input<string>;
    }
}

export namespace vpnaas {
    export interface IkePolicyLifetime {
        units?: pulumi.Input<string>;
        /**
         * The value for the lifetime of the security association. Must be a positive integer.
         * Default is 3600.
         */
        value?: pulumi.Input<number>;
    }

    export interface IpSecPolicyLifetime {
        units?: pulumi.Input<string>;
        /**
         * The value for the lifetime of the security association. Must be a positive integer.
         * Default is 3600.
         */
        value?: pulumi.Input<number>;
    }

    export interface SiteConnectionDpd {
        /**
         * The dead peer detection (DPD) action.
         * A valid value is clear, hold, restart, disabled, or restart-by-peer.
         * Default value is hold.
         */
        action?: pulumi.Input<string>;
        /**
         * The dead peer detection (DPD) interval, in seconds.
         * A valid value is a positive integer.
         * Default is 30.
         */
        interval?: pulumi.Input<number>;
        /**
         * The dead peer detection (DPD) timeout in seconds.
         * A valid value is a positive integer that is greater than the DPD interval value.
         * Default is 120.
         */
        timeout?: pulumi.Input<number>;
    }
}
