// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ShareArgs extends com.pulumi.resources.ResourceArgs {

    public static final ShareArgs Empty = new ShareArgs();

    /**
     * The share availability zone. Changing this creates a
     * new share.
     * 
     */
    @Import(name="availabilityZone")
    private @Nullable Output<String> availabilityZone;

    /**
     * @return The share availability zone. Changing this creates a
     * new share.
     * 
     */
    public Optional<Output<String>> availabilityZone() {
        return Optional.ofNullable(this.availabilityZone);
    }

    /**
     * The human-readable description for the share.
     * Changing this updates the description of the existing share.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The human-readable description for the share.
     * Changing this updates the description of the existing share.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The level of visibility for the share. Set to true to make
     * share public. Set to false to make it private. Default value is false. Changing this
     * updates the existing share.
     * 
     */
    @Import(name="isPublic")
    private @Nullable Output<Boolean> isPublic;

    /**
     * @return The level of visibility for the share. Set to true to make
     * share public. Set to false to make it private. Default value is false. Changing this
     * updates the existing share.
     * 
     */
    public Optional<Output<Boolean>> isPublic() {
        return Optional.ofNullable(this.isPublic);
    }

    /**
     * One or more metadata key and value pairs as a dictionary of
     * strings.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,Object>> metadata;

    /**
     * @return One or more metadata key and value pairs as a dictionary of
     * strings.
     * 
     */
    public Optional<Output<Map<String,Object>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * The name of the share. Changing this updates the name
     * of the existing share.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the share. Changing this updates the name
     * of the existing share.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share. Changing this
     * creates a new share.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share. Changing this
     * creates a new share.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The UUID of a share network where the share server exists
     * or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
     * the share_network_id value from the snapshot is used. Changing this creates a new share.
     * 
     */
    @Import(name="shareNetworkId")
    private @Nullable Output<String> shareNetworkId;

    /**
     * @return The UUID of a share network where the share server exists
     * or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
     * the share_network_id value from the snapshot is used. Changing this creates a new share.
     * 
     */
    public Optional<Output<String>> shareNetworkId() {
        return Optional.ofNullable(this.shareNetworkId);
    }

    /**
     * The share protocol - can either be NFS, CIFS,
     * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
     * 
     */
    @Import(name="shareProto", required=true)
    private Output<String> shareProto;

    /**
     * @return The share protocol - can either be NFS, CIFS,
     * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
     * 
     */
    public Output<String> shareProto() {
        return this.shareProto;
    }

    /**
     * The share type name. If you omit this parameter, the default
     * share type is used.
     * 
     */
    @Import(name="shareType")
    private @Nullable Output<String> shareType;

    /**
     * @return The share type name. If you omit this parameter, the default
     * share type is used.
     * 
     */
    public Optional<Output<String>> shareType() {
        return Optional.ofNullable(this.shareType);
    }

    /**
     * The share size, in GBs. The requested share size cannot be greater
     * than the allowed GB quota. Changing this resizes the existing share.
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return The share size, in GBs. The requested share size cannot be greater
     * than the allowed GB quota. Changing this resizes the existing share.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * The UUID of the share&#39;s base snapshot. Changing this creates
     * a new share.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return The UUID of the share&#39;s base snapshot. Changing this creates
     * a new share.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
    }

    private ShareArgs() {}

    private ShareArgs(ShareArgs $) {
        this.availabilityZone = $.availabilityZone;
        this.description = $.description;
        this.isPublic = $.isPublic;
        this.metadata = $.metadata;
        this.name = $.name;
        this.region = $.region;
        this.shareNetworkId = $.shareNetworkId;
        this.shareProto = $.shareProto;
        this.shareType = $.shareType;
        this.size = $.size;
        this.snapshotId = $.snapshotId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ShareArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ShareArgs $;

        public Builder() {
            $ = new ShareArgs();
        }

        public Builder(ShareArgs defaults) {
            $ = new ShareArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param availabilityZone The share availability zone. Changing this creates a
         * new share.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(@Nullable Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The share availability zone. Changing this creates a
         * new share.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param description The human-readable description for the share.
         * Changing this updates the description of the existing share.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The human-readable description for the share.
         * Changing this updates the description of the existing share.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param isPublic The level of visibility for the share. Set to true to make
         * share public. Set to false to make it private. Default value is false. Changing this
         * updates the existing share.
         * 
         * @return builder
         * 
         */
        public Builder isPublic(@Nullable Output<Boolean> isPublic) {
            $.isPublic = isPublic;
            return this;
        }

        /**
         * @param isPublic The level of visibility for the share. Set to true to make
         * share public. Set to false to make it private. Default value is false. Changing this
         * updates the existing share.
         * 
         * @return builder
         * 
         */
        public Builder isPublic(Boolean isPublic) {
            return isPublic(Output.of(isPublic));
        }

        /**
         * @param metadata One or more metadata key and value pairs as a dictionary of
         * strings.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,Object>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata One or more metadata key and value pairs as a dictionary of
         * strings.
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,Object> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param name The name of the share. Changing this updates the name
         * of the existing share.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the share. Changing this updates the name
         * of the existing share.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V2 Shared File System client.
         * A Shared File System client is needed to create a share. Changing this
         * creates a new share.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Shared File System client.
         * A Shared File System client is needed to create a share. Changing this
         * creates a new share.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param shareNetworkId The UUID of a share network where the share server exists
         * or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
         * the share_network_id value from the snapshot is used. Changing this creates a new share.
         * 
         * @return builder
         * 
         */
        public Builder shareNetworkId(@Nullable Output<String> shareNetworkId) {
            $.shareNetworkId = shareNetworkId;
            return this;
        }

        /**
         * @param shareNetworkId The UUID of a share network where the share server exists
         * or will be created. If `share_network_id` is not set and you provide a `snapshot_id`,
         * the share_network_id value from the snapshot is used. Changing this creates a new share.
         * 
         * @return builder
         * 
         */
        public Builder shareNetworkId(String shareNetworkId) {
            return shareNetworkId(Output.of(shareNetworkId));
        }

        /**
         * @param shareProto The share protocol - can either be NFS, CIFS,
         * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
         * 
         * @return builder
         * 
         */
        public Builder shareProto(Output<String> shareProto) {
            $.shareProto = shareProto;
            return this;
        }

        /**
         * @param shareProto The share protocol - can either be NFS, CIFS,
         * CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.
         * 
         * @return builder
         * 
         */
        public Builder shareProto(String shareProto) {
            return shareProto(Output.of(shareProto));
        }

        /**
         * @param shareType The share type name. If you omit this parameter, the default
         * share type is used.
         * 
         * @return builder
         * 
         */
        public Builder shareType(@Nullable Output<String> shareType) {
            $.shareType = shareType;
            return this;
        }

        /**
         * @param shareType The share type name. If you omit this parameter, the default
         * share type is used.
         * 
         * @return builder
         * 
         */
        public Builder shareType(String shareType) {
            return shareType(Output.of(shareType));
        }

        /**
         * @param size The share size, in GBs. The requested share size cannot be greater
         * than the allowed GB quota. Changing this resizes the existing share.
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The share size, in GBs. The requested share size cannot be greater
         * than the allowed GB quota. Changing this resizes the existing share.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param snapshotId The UUID of the share&#39;s base snapshot. Changing this creates
         * a new share.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId The UUID of the share&#39;s base snapshot. Changing this creates
         * a new share.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
        }

        public ShareArgs build() {
            if ($.shareProto == null) {
                throw new MissingRequiredPropertyException("ShareArgs", "shareProto");
            }
            if ($.size == null) {
                throw new MissingRequiredPropertyException("ShareArgs", "size");
            }
            return $;
        }
    }

}
