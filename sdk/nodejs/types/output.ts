// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace blockstorage {
    export interface GetVolumeV3Attachment {
        device: string;
        id: string;
        instanceId: string;
    }

    export interface VolumeAttachment {
        device: string;
        id: string;
        instanceId: string;
    }

    export interface VolumeSchedulerHint {
        /**
         * Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         */
        additionalProperties?: {[key: string]: any};
        /**
         * The volume should be scheduled on a 
         * different host from the set of volumes specified in the list provided.
         */
        differentHosts?: string[];
        /**
         * An instance UUID. The volume should be 
         * scheduled on the same host as the instance.
         */
        localToInstance?: string;
        /**
         * A conditional query that a back-end must pass in
         * order to host a volume. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         *
         * ```
         * [“=”, “$backend_id”, “rbd:vol@ceph#cloud”]
         * ```
         */
        query?: string;
        /**
         * A list of volume UUIDs. The volume should be
         * scheduled on the same host as another volume specified in the list provided.
         */
        sameHosts?: string[];
    }

    export interface VolumeV1Attachment {
        device: string;
        id: string;
        instanceId: string;
    }

    export interface VolumeV2Attachment {
        device: string;
        id: string;
        instanceId: string;
    }

    export interface VolumeV2SchedulerHint {
        /**
         * Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         */
        additionalProperties?: {[key: string]: any};
        /**
         * The volume should be scheduled on a 
         * different host from the set of volumes specified in the list provided.
         */
        differentHosts?: string[];
        /**
         * An instance UUID. The volume should be 
         * scheduled on the same host as the instance.
         */
        localToInstance?: string;
        /**
         * A conditional query that a back-end must pass in
         * order to host a volume. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         *
         * ```
         * [“=”, “$backend_id”, “rbd:vol@ceph#cloud”]
         * ```
         */
        query?: string;
        /**
         * A list of volume UUIDs. The volume should be
         * scheduled on the same host as another volume specified in the list provided.
         */
        sameHosts?: string[];
    }

}

export namespace compute {
    export interface GetInstanceV2Network {
        /**
         * The IPv4 address assigned to this network port.
         */
        fixedIpV4: string;
        /**
         * The IPv6 address assigned to this network port.
         */
        fixedIpV6: string;
        /**
         * The MAC address assigned to this network interface.
         */
        mac: string;
        /**
         * The name of the network
         */
        name: string;
        /**
         * The port UUID for this network
         */
        port: string;
        /**
         * The UUID of the network
         */
        uuid: string;
    }

    export interface InstanceBlockDevice {
        /**
         * The boot index of the volume. It defaults to 0.
         * Changing this creates a new server.
         */
        bootIndex?: number;
        /**
         * Delete the volume / block device upon
         * termination of the instance. Defaults to false. Changing this creates a
         * new server.
         */
        deleteOnTermination?: boolean;
        /**
         * The type that gets created. Possible values
         * are "volume" and "local". Changing this creates a new server.
         */
        destinationType?: string;
        /**
         * The low-level device type that will be used. Most
         * common thing is to leave this empty. Changing this creates a new server.
         */
        deviceType?: string;
        /**
         * The low-level disk bus that will be used. Most common
         * thing is to leave this empty. Changing this creates a new server.
         */
        diskBus?: string;
        /**
         * Specifies the guest server disk file system format,
         * such as `ext2`, `ext3`, `ext4`, `xfs` or `swap`. Swap block device mappings
         * have the following restrictions: sourceType must be blank and destinationType
         * must be local and only one swap disk per server and the size of the swap disk
         * must be less than or equal to the swap size of the flavor. Changing this
         * creates a new server.
         */
        guestFormat?: string;
        /**
         * Enable the attachment of multiattach-capable
         * volumes.
         */
        multiattach?: boolean;
        /**
         * The source type of the device. Must be one of
         * "blank", "image", "volume", or "snapshot". Changing this creates a new
         * server.
         */
        sourceType: string;
        /**
         * The UUID of
         * the image, volume, or snapshot. Changing this creates a new server.
         */
        uuid?: string;
        /**
         * The size of the volume to create (in gigabytes). Required
         * in the following combinations: source=image and destination=volume,
         * source=blank and destination=local, and source=blank and destination=volume.
         * Changing this creates a new server.
         */
        volumeSize?: number;
        /**
         * The volume type that will be used, for example SSD
         * or HDD storage. The available options depend on how your specific OpenStack
         * cloud is configured and what classes of storage are provided. Changing this
         * creates a new server.
         */
        volumeType?: string;
    }

    export interface InstanceNetwork {
        /**
         * Specifies if this network should be used for
         * provisioning access. Accepts true or false. Defaults to false.
         */
        accessNetwork?: boolean;
        /**
         * Specifies a fixed IPv4 address to be used on this
         * network. Changing this creates a new server.
         */
        fixedIpV4: string;
        fixedIpV6: string;
        /**
         * @deprecated Use the openstack.compute.FloatingIpAssociate resource instead
         */
        floatingIp: string;
        mac: string;
        /**
         * The human-readable
         * name of the network. Changing this creates a new server.
         */
        name: string;
        /**
         * The port UUID of a
         * network to attach to the server. Changing this creates a new server.
         */
        port: string;
        /**
         * The network UUID to
         * attach to the server. Changing this creates a new server.
         */
        uuid: string;
    }

    export interface InstancePersonality {
        /**
         * The contents of the file. Limited to 255 bytes.
         */
        content: string;
        /**
         * The absolute path of the destination file.
         */
        file: string;
    }

    export interface InstanceSchedulerHint {
        /**
         * Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         */
        additionalProperties?: {[key: string]: any};
        /**
         * An IP Address in CIDR form. The instance
         * will be placed on a compute node that is in the same subnet.
         */
        buildNearHostIp?: string;
        /**
         * The names of cells where not to build the instance.
         */
        differentCells?: string[];
        /**
         * A list of instance UUIDs. The instance will
         * be scheduled on a different host than all other instances.
         */
        differentHosts?: string[];
        /**
         * A UUID of a Server Group. The instance will be placed
         * into that group.
         */
        group?: string;
        /**
         * A conditional query that a compute node must pass in
         * order to host an instance. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         *
         * ```
         * [">=", "$free_ram_mb", "1024"]
         * ```
         */
        queries?: string[];
        /**
         * A list of instance UUIDs. The instance will be
         * scheduled on the same host of those specified.
         */
        sameHosts?: string[];
        /**
         * The name of a cell to host the instance.
         */
        targetCell?: string;
    }

    export interface InstanceVendorOptions {
        /**
         * Whether to try to detach all attached
         * ports to the vm before destroying it to make sure the port state is correct
         * after the vm destruction. This is helpful when the port is not deleted.
         */
        detachPortsBeforeDestroy?: boolean;
        /**
         * Boolean to control whether
         * to ignore manual confirmation of the instance resizing. This can be helpful
         * to work with some OpenStack clouds which automatically confirm resizing of
         * instances after some timeout.
         */
        ignoreResizeConfirmation?: boolean;
    }

    export interface InstanceVolume {
        device: string;
        id: string;
        volumeId: string;
    }

    export interface SecGroupRule {
        /**
         * Required if `fromGroupId` or `self` is empty. The IP range
         * that will be the source of network traffic to the security group. Use 0.0.0.0/0
         * to allow all IP addresses. Changing this creates a new security group rule. Cannot
         * be combined with `fromGroupId` or `self`.
         */
        cidr?: string;
        /**
         * Required if `cidr` or `self` is empty. The ID of a
         * group from which to forward traffic to the parent group. Changing this creates a
         * new security group rule. Cannot be combined with `cidr` or `self`.
         */
        fromGroupId?: string;
        /**
         * An integer representing the lower bound of the port
         * range to open. Changing this creates a new security group rule.
         */
        fromPort: number;
        id: string;
        /**
         * The protocol type that will be allowed. Changing
         * this creates a new security group rule.
         */
        ipProtocol: string;
        /**
         * Required if `cidr` and `fromGroupId` is empty. If true,
         * the security group itself will be added as a source to this ingress rule. Cannot
         * be combined with `cidr` or `fromGroupId`.
         */
        self?: boolean;
        /**
         * An integer representing the upper bound of the port
         * range to open. Changing this creates a new security group rule.
         */
        toPort: number;
    }

    export interface ServerGroupRules {
        maxServerPerHost?: number;
    }

    export interface VolumeAttachVendorOptions {
        /**
         * Boolean to control whether
         * to ignore volume status confirmation of the attached volume. This can be helpful
         * to work with some OpenStack clouds which don't have the Block Storage V3 API available.
         */
        ignoreVolumeConfirmation?: boolean;
    }

}

export namespace database {
    export interface ConfigurationConfiguration {
        /**
         * Configuration parameter name. Changing this creates a new resource.
         */
        name: string;
        /**
         * Whether or not to store configuration parameter value as string. Changing this creates a new resource. See the below note for more information.
         */
        stringType?: boolean;
        /**
         * Configuration parameter value. Changing this creates a new resource.
         */
        value: string;
    }

    export interface ConfigurationDatastore {
        /**
         * Database engine type to be used with this configuration. Changing this creates a new resource.
         */
        type: string;
        /**
         * Version of database engine type to be used with this configuration. Changing this creates a new resource.
         */
        version: string;
    }

    export interface InstanceDatabase {
        /**
         * Database character set. Changing this creates a
         * new instance.
         */
        charset?: string;
        /**
         * Database collation. Changing this creates a new instance.
         */
        collate?: string;
        /**
         * Database to be created on new instance. Changing this creates a
         * new instance.
         */
        name: string;
    }

    export interface InstanceDatastore {
        /**
         * Database engine type to be used in new instance. Changing this
         * creates a new instance.
         */
        type: string;
        /**
         * Version of database engine type to be used in new instance.
         * Changing this creates a new instance.
         */
        version: string;
    }

    export interface InstanceNetwork {
        /**
         * Specifies a fixed IPv4 address to be used on this
         * network. Changing this creates a new instance.
         */
        fixedIpV4?: string;
        /**
         * Specifies a fixed IPv6 address to be used on this
         * network. Changing this creates a new instance.
         */
        fixedIpV6?: string;
        /**
         * The port UUID of a
         * network to attach to the instance. Changing this creates a new instance.
         */
        port?: string;
        /**
         * The network UUID to
         * attach to the instance. Changing this creates a new instance.
         */
        uuid?: string;
    }

    export interface InstanceUser {
        /**
         * A list of databases that user will have access to. If not specified,
         * user has access to all databases on th einstance. Changing this creates a new instance.
         */
        databases?: string[];
        /**
         * An ip address or % sign indicating what ip addresses can connect with
         * this user credentials. Changing this creates a new instance.
         */
        host?: string;
        /**
         * Username to be created on new instance. Changing this creates a
         * new instance.
         */
        name: string;
        /**
         * User's password. Changing this creates a
         * new instance.
         */
        password?: string;
    }

}

export namespace identity {
    export interface ApplicationCredentialAccessRule {
        /**
         * The ID of the existing access rule. The access rule ID of
         * another application credential can be provided.
         */
        id: string;
        /**
         * The request method that the application credential is
         * permitted to use for a given API endpoint. Allowed values: `POST`, `GET`,
         * `HEAD`, `PATCH`, `PUT` and `DELETE`.
         */
        method: string;
        /**
         * The API path that the application credential is permitted
         * to access. May use named wildcards such as **{tag}** or the unnamed wildcard
         * **\*** to match against any string in the path up to a **&#47;**, or the recursive
         * wildcard **\*\*** to include **&#47;** in the matched path.
         */
        path: string;
        /**
         * The service type identifier for the service that the
         * application credential is granted to access. Must be a service type that is
         * listed in the service catalog and not a code name for a service. E.g.
         * **identity**, **compute**, **volumev3**, **image**, **network**,
         * **object-store**, **sharev2**, **dns**, **key-manager**, **monitoring**, etc.
         */
        service: string;
    }

    export interface GetAuthScopeRole {
        /**
         * The ID of the role.
         */
        roleId: string;
        /**
         * The name of the role.
         */
        roleName: string;
    }

    export interface GetAuthScopeServiceCatalog {
        /**
         * A list of endpoints for the service.
         */
        endpoints: outputs.identity.GetAuthScopeServiceCatalogEndpoint[];
        /**
         * The ID of the endpoint.
         */
        id: string;
        /**
         * The name of the scope. This is an arbitrary name which is
         * only used as a unique identifier so an actual token isn't used as the ID.
         */
        name: string;
        /**
         * The type of the service.
         */
        type: string;
    }

    export interface GetAuthScopeServiceCatalogEndpoint {
        /**
         * The ID of the endpoint.
         */
        id: string;
        /**
         * The interface of the endpoint.
         */
        interface: string;
        /**
         * The region in which to obtain the V3 Identity client.
         * A Identity client is needed to retrieve tokens IDs. If omitted, the
         * `region` argument of the provider is used.
         */
        region: string;
        /**
         * The region ID of the endpoint.
         */
        regionId: string;
        /**
         * The URL of the endpoint.
         */
        url: string;
    }

    export interface UserMultiFactorAuthRule {
        /**
         * A list of authentication plugins that the user must
         * authenticate with.
         */
        rules: string[];
    }

}

export namespace keymanager {
    export interface ContainerV1Acl {
        read: outputs.keymanager.ContainerV1AclRead;
    }

    export interface ContainerV1AclRead {
        /**
         * The date the container was created.
         */
        createdAt: string;
        /**
         * Whether the container is accessible project wide.
         * Defaults to `true`.
         */
        projectAccess?: boolean;
        /**
         * The date the container was last updated.
         */
        updatedAt: string;
        /**
         * The list of user IDs, which are allowed to access the
         * container, when `projectAccess` is set to `false`.
         */
        users?: string[];
    }

    export interface ContainerV1Consumer {
        /**
         * Human-readable name for the Container. Does not have
         * to be unique.
         */
        name?: string;
        /**
         * The consumer URL.
         */
        url?: string;
    }

    export interface ContainerV1SecretRef {
        /**
         * The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
         */
        name?: string;
        /**
         * The secret reference / where to find the secret, URL.
         */
        secretRef: string;
    }

    export interface GetContainerAcl {
        read: outputs.keymanager.GetContainerAclRead;
    }

    export interface GetContainerAclRead {
        /**
         * The date the container ACL was created.
         */
        createdAt: string;
        /**
         * Whether the container is accessible project wide.
         */
        projectAccess?: boolean;
        /**
         * The date the container ACL was last updated.
         */
        updatedAt: string;
        /**
         * The list of user IDs, which are allowed to access the container,
         * when `projectAccess` is set to `false`.
         */
        users?: string[];
    }

    export interface GetContainerConsumer {
        /**
         * The Container name.
         */
        name?: string;
        /**
         * The consumer URL.
         */
        url?: string;
    }

    export interface GetContainerSecretRef {
        /**
         * The Container name.
         */
        name?: string;
        /**
         * The secret reference / where to find the secret, URL.
         */
        secretRef?: string;
    }

    export interface GetSecretAcl {
        read: outputs.keymanager.GetSecretAclRead;
    }

    export interface GetSecretAclRead {
        /**
         * The date the secret ACL was created.
         */
        createdAt: string;
        /**
         * Whether the secret is accessible project wide.
         */
        projectAccess?: boolean;
        /**
         * The date the secret ACL was last updated.
         */
        updatedAt: string;
        /**
         * The list of user IDs, which are allowed to access the secret, when
         * `projectAccess` is set to `false`.
         */
        users?: string[];
    }

    export interface OrderV1Meta {
        /**
         * Algorithm to use for key generation.
         */
        algorithm: string;
        /**
         * Bit lenght of key to be generated.
         */
        bitLength: number;
        /**
         * This is a UTC timestamp in ISO 8601 format YYYY-MM-DDTHH:MM:SSZ. If set, the secret will not be available after this time.
         */
        expiration?: string;
        /**
         * The mode to use for key generation.
         */
        mode?: string;
        /**
         * The name of the secret set by the user.
         */
        name?: string;
        /**
         * The media type for the content of the secrets payload. Must be one of `text/plain`, `text/plain;charset=utf-8`, `text/plain; charset=utf-8`, `application/octet-stream`, `application/pkcs8`.
         */
        payloadContentType?: string;
    }

    export interface SecretV1Acl {
        read: outputs.keymanager.SecretV1AclRead;
    }

    export interface SecretV1AclRead {
        /**
         * The date the secret was created.
         */
        createdAt: string;
        /**
         * Whether the secret is accessible project wide.
         * Defaults to `true`.
         */
        projectAccess?: boolean;
        /**
         * The date the secret was last updated.
         */
        updatedAt: string;
        /**
         * The list of user IDs, which are allowed to access the
         * secret, when `projectAccess` is set to `false`.
         */
        users?: string[];
    }

}

export namespace loadbalancer {
    export interface MembersMember {
        /**
         * The IP address of the members to receive traffic from
         * the load balancer.
         */
        address: string;
        /**
         * The administrative state of the member.
         * A valid value is true (UP) or false (DOWN). Defaults to true.
         */
        adminStateUp?: boolean;
        /**
         * A bool that indicates whether the member is
         * backup. **Requires octavia minor version 2.1 or later**.
         */
        backup?: boolean;
        /**
         * The unique ID for the members.
         */
        id: string;
        /**
         * An alternate IP address used for health 
         * monitoring a backend member.
         */
        monitorAddress?: string;
        /**
         * An alternate protocol port used for health 
         * monitoring a backend member.
         */
        monitorPort?: number;
        /**
         * Human-readable name for the member.
         */
        name?: string;
        /**
         * The port on which to listen for client traffic.
         */
        protocolPort: number;
        /**
         * The subnet in which to access the member.
         */
        subnetId?: string;
        /**
         * A positive integer value that indicates the relative
         * portion of traffic that this members should receive from the pool. For
         * example, a member with a weight of 10 receives five times as much traffic
         * as a member with a weight of 2. Defaults to 1.
         */
        weight?: number;
    }

    export interface PoolPersistence {
        /**
         * The name of the cookie if persistence mode is set
         * appropriately. Required if `type = APP_COOKIE`.
         */
        cookieName?: string;
        /**
         * The type of persistence mode. The current specification
         * supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.
         */
        type: string;
    }

}

export namespace networking {
    export interface GetNetworkSegment {
        networkType: string;
        physicalNetwork: string;
        segmentationId: number;
    }

    export interface GetPortAllowedAddressPair {
        /**
         * The additional IP address.
         */
        ipAddress: string;
        /**
         * The MAC address of the port.
         */
        macAddress: string;
    }

    export interface GetPortBinding {
        /**
         * The ID of the host, which has the allocatee port.
         */
        hostId: string;
        /**
         * A JSON string containing the binding profile information.
         */
        profile: string;
        /**
         * A map of JSON strings containing additional details for this
         * specific binding.
         */
        vifDetails: {[key: string]: any};
        /**
         * The VNIC type of the port binding.
         */
        vifType: string;
        /**
         * VNIC type for the port.
         */
        vnicType: string;
    }

    export interface GetPortExtraDhcpOption {
        /**
         * IP protocol version
         */
        ipVersion: number;
        /**
         * The name of the port.
         */
        name: string;
        /**
         * Value of the DHCP option.
         */
        value: string;
    }

    export interface GetRouterExternalFixedIp {
        /**
         * The IP address to set on the router.
         */
        ipAddress?: string;
        /**
         * Subnet in which the fixed IP belongs to.
         */
        subnetId?: string;
    }

    export interface GetSubnetAllocationPool {
        end: string;
        start: string;
    }

    export interface GetSubnetHostRoute {
        destinationCidr: string;
        nextHop: string;
    }

    export interface GetTrunkSubPort {
        /**
         * The ID of the trunk parent port.
         */
        portId: string;
        /**
         * The numeric id of the subport segment.
         */
        segmentationId: number;
        /**
         * The segmenation tecnology used, e.g., "vlan".
         */
        segmentationType: string;
    }

    export interface NetworkSegment {
        /**
         * The type of physical network.
         */
        networkType?: string;
        /**
         * The physical network where this network is implemented.
         */
        physicalNetwork?: string;
        /**
         * An isolated segment on the physical network.
         */
        segmentationId: number;
    }

    export interface PortAllowedAddressPair {
        /**
         * The additional IP address.
         */
        ipAddress: string;
        /**
         * The additional MAC address.
         */
        macAddress?: string;
    }

    export interface PortBinding {
        /**
         * The ID of the host to allocate port on.
         */
        hostId: string;
        /**
         * Custom data to be passed as `binding:profile`. Data
         * must be passed as JSON.
         */
        profile?: string;
        /**
         * A map of JSON strings containing additional
         * details for this specific binding.
         */
        vifDetails: {[key: string]: any};
        /**
         * The VNIC type of the port binding.
         */
        vifType: string;
        /**
         * VNIC type for the port. Can either be `direct`,
         * `direct-physical`, `macvtap`, `normal`, `baremetal` or `virtio-forwarder`.
         * Default value is `normal`. It can be updated on unbound ports only.
         */
        vnicType?: string;
    }

    export interface PortExtraDhcpOption {
        /**
         * IP protocol version. Defaults to 4.
         */
        ipVersion?: number;
        /**
         * Name of the DHCP option.
         */
        name: string;
        /**
         * Value of the DHCP option.
         */
        value: string;
    }

    export interface PortFixedIp {
        /**
         * IP address desired in the subnet for this port. If
         * you don't specify `ipAddress`, an available IP address from the specified
         * subnet will be allocated to this port. This field will not be populated if it
         * is left blank or omitted. To retrieve the assigned IP address, use the
         * `allFixedIps` attribute.
         */
        ipAddress?: string;
        /**
         * Subnet in which to allocate IP address for
         * this port.
         */
        subnetId: string;
    }

    export interface RouterExternalFixedIp {
        /**
         * The IP address to set on the router.
         */
        ipAddress: string;
        /**
         * Subnet in which the fixed IP belongs to.
         */
        subnetId?: string;
    }

    export interface RouterVendorOptions {
        /**
         * Boolean to control whether
         * the Router gateway is assigned during creation or updated after creation.
         */
        setRouterGatewayAfterCreate?: boolean;
    }

    export interface SubnetAllocationPool {
        /**
         * The ending address.
         */
        end: string;
        /**
         * The starting address.
         */
        start: string;
    }

    export interface SubnetHostRoute {
        /**
         * The destination CIDR.
         */
        destinationCidr: string;
        /**
         * The next hop in the route.
         */
        nextHop: string;
    }

    export interface TrunkSubPort {
        /**
         * The ID of the port to be made a subport of the trunk.
         */
        portId: string;
        /**
         * The numeric id of the subport segment.
         */
        segmentationId: number;
        /**
         * The segmentation technology to use, e.g., "vlan".
         */
        segmentationType: string;
    }

}

export namespace objectstorage {
    export interface ContainerVersioningLegacy {
        /**
         * Container in which versions will be stored.
         */
        location: string;
        /**
         * Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
         */
        type: string;
    }

}

export namespace orchestration {
    export interface StackV1StackOutput {
        /**
         * The description of the stack resource.
         */
        description: string;
        outputKey: string;
        outputValue: string;
    }

}

export namespace sharedfilesystem {
    export interface GetShareExportLocation {
        path: string;
        preferred: string;
    }

    export interface ShareExportLocation {
        path: string;
        preferred: string;
    }

}

export namespace vpnaas {
    export interface IkePolicyLifetime {
        units: string;
        /**
         * The value for the lifetime of the security association. Must be a positive integer.
         * Default is 3600.
         */
        value: number;
    }

    export interface IpSecPolicyLifetime {
        units: string;
        /**
         * The value for the lifetime of the security association. Must be a positive integer.
         * Default is 3600.
         */
        value: number;
    }

    export interface SiteConnectionDpd {
        /**
         * The dead peer detection (DPD) action.
         * A valid value is clear, hold, restart, disabled, or restart-by-peer.
         * Default value is hold.
         */
        action: string;
        /**
         * The dead peer detection (DPD) interval, in seconds.
         * A valid value is a positive integer.
         * Default is 30.
         */
        interval: number;
        /**
         * The dead peer detection (DPD) timeout in seconds.
         * A valid value is a positive integer that is greater than the DPD interval value.
         * Default is 120.
         */
        timeout: number;
    }

}
