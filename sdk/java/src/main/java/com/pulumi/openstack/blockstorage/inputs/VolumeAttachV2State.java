// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeAttachV2State extends com.pulumi.resources.ResourceArgs {

    public static final VolumeAttachV2State Empty = new VolumeAttachV2State();

    /**
     * Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     * 
     */
    @Import(name="attachMode")
    private @Nullable Output<String> attachMode;

    /**
     * @return Specify whether to attach the volume as Read-Only
     * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
     * If left unspecified, the Block Storage API will apply a default of `rw`.
     * 
     */
    public Optional<Output<String>> attachMode() {
        return Optional.ofNullable(this.attachMode);
    }

    /**
     * This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     * 
     */
    @Import(name="data")
    private @Nullable Output<Map<String,Object>> data;

    /**
     * @return This is a map of key/value pairs that contain the connection
     * information. You will want to pass this information to a provisioner
     * script to finalize the connection. See below for more information.
     * 
     */
    public Optional<Output<Map<String,Object>>> data() {
        return Optional.ofNullable(this.data);
    }

    /**
     * The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     * 
     */
    @Import(name="device")
    private @Nullable Output<String> device;

    /**
     * @return The device to tell the Block Storage service this
     * volume will be attached as. This is purely for informational purposes.
     * You can specify `auto` or a device such as `/dev/vdc`.
     * 
     */
    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * The storage driver that the volume is based on.
     * 
     */
    @Import(name="driverVolumeType")
    private @Nullable Output<String> driverVolumeType;

    /**
     * @return The storage driver that the volume is based on.
     * 
     */
    public Optional<Output<String>> driverVolumeType() {
        return Optional.ofNullable(this.driverVolumeType);
    }

    /**
     * The host to attach the volume to.
     * 
     */
    @Import(name="hostName")
    private @Nullable Output<String> hostName;

    /**
     * @return The host to attach the volume to.
     * 
     */
    public Optional<Output<String>> hostName() {
        return Optional.ofNullable(this.hostName);
    }

    /**
     * The iSCSI initiator string to make the connection.
     * 
     */
    @Import(name="initiator")
    private @Nullable Output<String> initiator;

    /**
     * @return The iSCSI initiator string to make the connection.
     * 
     */
    public Optional<Output<String>> initiator() {
        return Optional.ofNullable(this.initiator);
    }

    /**
     * @deprecated
     * instance_id is no longer used in this resource
     * 
     */
    @Deprecated /* instance_id is no longer used in this resource */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @deprecated
     * instance_id is no longer used in this resource
     * 
     */
    @Deprecated /* instance_id is no longer used in this resource */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The IP address of the `host_name` above.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The IP address of the `host_name` above.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * A mount point base name for shared storage.
     * 
     */
    @Import(name="mountPointBase")
    private @Nullable Output<String> mountPointBase;

    /**
     * @return A mount point base name for shared storage.
     * 
     */
    public Optional<Output<String>> mountPointBase() {
        return Optional.ofNullable(this.mountPointBase);
    }

    /**
     * Whether to connect to this volume via multipath.
     * 
     */
    @Import(name="multipath")
    private @Nullable Output<Boolean> multipath;

    /**
     * @return Whether to connect to this volume via multipath.
     * 
     */
    public Optional<Output<Boolean>> multipath() {
        return Optional.ofNullable(this.multipath);
    }

    /**
     * The iSCSI initiator OS type.
     * 
     */
    @Import(name="osType")
    private @Nullable Output<String> osType;

    /**
     * @return The iSCSI initiator OS type.
     * 
     */
    public Optional<Output<String>> osType() {
        return Optional.ofNullable(this.osType);
    }

    /**
     * The iSCSI initiator platform.
     * 
     */
    @Import(name="platform")
    private @Nullable Output<String> platform;

    /**
     * @return The iSCSI initiator platform.
     * 
     */
    public Optional<Output<String>> platform() {
        return Optional.ofNullable(this.platform);
    }

    /**
     * The region in which to obtain the V2 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Block Storage
     * client. A Block Storage client is needed to create a volume attachment.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new volume attachment.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The ID of the Volume to attach to an Instance.
     * 
     */
    @Import(name="volumeId")
    private @Nullable Output<String> volumeId;

    /**
     * @return The ID of the Volume to attach to an Instance.
     * 
     */
    public Optional<Output<String>> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    /**
     * A wwnn name. Used for Fibre Channel connections.
     * 
     */
    @Import(name="wwnn")
    private @Nullable Output<String> wwnn;

    /**
     * @return A wwnn name. Used for Fibre Channel connections.
     * 
     */
    public Optional<Output<String>> wwnn() {
        return Optional.ofNullable(this.wwnn);
    }

    /**
     * An array of wwpn strings. Used for Fibre Channel
     * connections.
     * 
     */
    @Import(name="wwpns")
    private @Nullable Output<List<String>> wwpns;

    /**
     * @return An array of wwpn strings. Used for Fibre Channel
     * connections.
     * 
     */
    public Optional<Output<List<String>>> wwpns() {
        return Optional.ofNullable(this.wwpns);
    }

    private VolumeAttachV2State() {}

    private VolumeAttachV2State(VolumeAttachV2State $) {
        this.attachMode = $.attachMode;
        this.data = $.data;
        this.device = $.device;
        this.driverVolumeType = $.driverVolumeType;
        this.hostName = $.hostName;
        this.initiator = $.initiator;
        this.instanceId = $.instanceId;
        this.ipAddress = $.ipAddress;
        this.mountPointBase = $.mountPointBase;
        this.multipath = $.multipath;
        this.osType = $.osType;
        this.platform = $.platform;
        this.region = $.region;
        this.volumeId = $.volumeId;
        this.wwnn = $.wwnn;
        this.wwpns = $.wwpns;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeAttachV2State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeAttachV2State $;

        public Builder() {
            $ = new VolumeAttachV2State();
        }

        public Builder(VolumeAttachV2State defaults) {
            $ = new VolumeAttachV2State(Objects.requireNonNull(defaults));
        }

        /**
         * @param attachMode Specify whether to attach the volume as Read-Only
         * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
         * If left unspecified, the Block Storage API will apply a default of `rw`.
         * 
         * @return builder
         * 
         */
        public Builder attachMode(@Nullable Output<String> attachMode) {
            $.attachMode = attachMode;
            return this;
        }

        /**
         * @param attachMode Specify whether to attach the volume as Read-Only
         * (`ro`) or Read-Write (`rw`). Only values of `ro` and `rw` are accepted.
         * If left unspecified, the Block Storage API will apply a default of `rw`.
         * 
         * @return builder
         * 
         */
        public Builder attachMode(String attachMode) {
            return attachMode(Output.of(attachMode));
        }

        /**
         * @param data This is a map of key/value pairs that contain the connection
         * information. You will want to pass this information to a provisioner
         * script to finalize the connection. See below for more information.
         * 
         * @return builder
         * 
         */
        public Builder data(@Nullable Output<Map<String,Object>> data) {
            $.data = data;
            return this;
        }

        /**
         * @param data This is a map of key/value pairs that contain the connection
         * information. You will want to pass this information to a provisioner
         * script to finalize the connection. See below for more information.
         * 
         * @return builder
         * 
         */
        public Builder data(Map<String,Object> data) {
            return data(Output.of(data));
        }

        /**
         * @param device The device to tell the Block Storage service this
         * volume will be attached as. This is purely for informational purposes.
         * You can specify `auto` or a device such as `/dev/vdc`.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        /**
         * @param device The device to tell the Block Storage service this
         * volume will be attached as. This is purely for informational purposes.
         * You can specify `auto` or a device such as `/dev/vdc`.
         * 
         * @return builder
         * 
         */
        public Builder device(String device) {
            return device(Output.of(device));
        }

        /**
         * @param driverVolumeType The storage driver that the volume is based on.
         * 
         * @return builder
         * 
         */
        public Builder driverVolumeType(@Nullable Output<String> driverVolumeType) {
            $.driverVolumeType = driverVolumeType;
            return this;
        }

        /**
         * @param driverVolumeType The storage driver that the volume is based on.
         * 
         * @return builder
         * 
         */
        public Builder driverVolumeType(String driverVolumeType) {
            return driverVolumeType(Output.of(driverVolumeType));
        }

        /**
         * @param hostName The host to attach the volume to.
         * 
         * @return builder
         * 
         */
        public Builder hostName(@Nullable Output<String> hostName) {
            $.hostName = hostName;
            return this;
        }

        /**
         * @param hostName The host to attach the volume to.
         * 
         * @return builder
         * 
         */
        public Builder hostName(String hostName) {
            return hostName(Output.of(hostName));
        }

        /**
         * @param initiator The iSCSI initiator string to make the connection.
         * 
         * @return builder
         * 
         */
        public Builder initiator(@Nullable Output<String> initiator) {
            $.initiator = initiator;
            return this;
        }

        /**
         * @param initiator The iSCSI initiator string to make the connection.
         * 
         * @return builder
         * 
         */
        public Builder initiator(String initiator) {
            return initiator(Output.of(initiator));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * instance_id is no longer used in this resource
         * 
         */
        @Deprecated /* instance_id is no longer used in this resource */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * instance_id is no longer used in this resource
         * 
         */
        @Deprecated /* instance_id is no longer used in this resource */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param ipAddress The IP address of the `host_name` above.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The IP address of the `host_name` above.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param mountPointBase A mount point base name for shared storage.
         * 
         * @return builder
         * 
         */
        public Builder mountPointBase(@Nullable Output<String> mountPointBase) {
            $.mountPointBase = mountPointBase;
            return this;
        }

        /**
         * @param mountPointBase A mount point base name for shared storage.
         * 
         * @return builder
         * 
         */
        public Builder mountPointBase(String mountPointBase) {
            return mountPointBase(Output.of(mountPointBase));
        }

        /**
         * @param multipath Whether to connect to this volume via multipath.
         * 
         * @return builder
         * 
         */
        public Builder multipath(@Nullable Output<Boolean> multipath) {
            $.multipath = multipath;
            return this;
        }

        /**
         * @param multipath Whether to connect to this volume via multipath.
         * 
         * @return builder
         * 
         */
        public Builder multipath(Boolean multipath) {
            return multipath(Output.of(multipath));
        }

        /**
         * @param osType The iSCSI initiator OS type.
         * 
         * @return builder
         * 
         */
        public Builder osType(@Nullable Output<String> osType) {
            $.osType = osType;
            return this;
        }

        /**
         * @param osType The iSCSI initiator OS type.
         * 
         * @return builder
         * 
         */
        public Builder osType(String osType) {
            return osType(Output.of(osType));
        }

        /**
         * @param platform The iSCSI initiator platform.
         * 
         * @return builder
         * 
         */
        public Builder platform(@Nullable Output<String> platform) {
            $.platform = platform;
            return this;
        }

        /**
         * @param platform The iSCSI initiator platform.
         * 
         * @return builder
         * 
         */
        public Builder platform(String platform) {
            return platform(Output.of(platform));
        }

        /**
         * @param region The region in which to obtain the V2 Block Storage
         * client. A Block Storage client is needed to create a volume attachment.
         * If omitted, the `region` argument of the provider is used. Changing this
         * creates a new volume attachment.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Block Storage
         * client. A Block Storage client is needed to create a volume attachment.
         * If omitted, the `region` argument of the provider is used. Changing this
         * creates a new volume attachment.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param volumeId The ID of the Volume to attach to an Instance.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(@Nullable Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        /**
         * @param volumeId The ID of the Volume to attach to an Instance.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        /**
         * @param wwnn A wwnn name. Used for Fibre Channel connections.
         * 
         * @return builder
         * 
         */
        public Builder wwnn(@Nullable Output<String> wwnn) {
            $.wwnn = wwnn;
            return this;
        }

        /**
         * @param wwnn A wwnn name. Used for Fibre Channel connections.
         * 
         * @return builder
         * 
         */
        public Builder wwnn(String wwnn) {
            return wwnn(Output.of(wwnn));
        }

        /**
         * @param wwpns An array of wwpn strings. Used for Fibre Channel
         * connections.
         * 
         * @return builder
         * 
         */
        public Builder wwpns(@Nullable Output<List<String>> wwpns) {
            $.wwpns = wwpns;
            return this;
        }

        /**
         * @param wwpns An array of wwpn strings. Used for Fibre Channel
         * connections.
         * 
         * @return builder
         * 
         */
        public Builder wwpns(List<String> wwpns) {
            return wwpns(Output.of(wwpns));
        }

        /**
         * @param wwpns An array of wwpn strings. Used for Fibre Channel
         * connections.
         * 
         * @return builder
         * 
         */
        public Builder wwpns(String... wwpns) {
            return wwpns(List.of(wwpns));
        }

        public VolumeAttachV2State build() {
            return $;
        }
    }

}
