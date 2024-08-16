// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuoteSetV2State extends com.pulumi.resources.ResourceArgs {

    public static final QuoteSetV2State Empty = new QuoteSetV2State();

    /**
     * Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     * 
     */
    @Import(name="backupGigabytes")
    private @Nullable Output<Integer> backupGigabytes;

    /**
     * @return Quota value for backup gigabytes. Changing
     * this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> backupGigabytes() {
        return Optional.ofNullable(this.backupGigabytes);
    }

    /**
     * Quota value for backups. Changing this updates the
     * existing quotaset.
     * 
     */
    @Import(name="backups")
    private @Nullable Output<Integer> backups;

    /**
     * @return Quota value for backups. Changing this updates the
     * existing quotaset.
     * 
     */
    public Optional<Output<Integer>> backups() {
        return Optional.ofNullable(this.backups);
    }

    /**
     * Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     * 
     */
    @Import(name="gigabytes")
    private @Nullable Output<Integer> gigabytes;

    /**
     * @return Quota value for gigabytes. Changing this updates the
     * existing quotaset.
     * 
     */
    public Optional<Output<Integer>> gigabytes() {
        return Optional.ofNullable(this.gigabytes);
    }

    /**
     * Quota value for groups. Changing this updates the
     * existing quotaset.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<Integer> groups;

    /**
     * @return Quota value for groups. Changing this updates the
     * existing quotaset.
     * 
     */
    public Optional<Output<Integer>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="perVolumeGigabytes")
    private @Nullable Output<Integer> perVolumeGigabytes;

    /**
     * @return Quota value for gigabytes per volume .
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> perVolumeGigabytes() {
        return Optional.ofNullable(this.perVolumeGigabytes);
    }

    /**
     * ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return ID of the project to manage quotas. Changing this
     * creates a new quotaset.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Quota value for snapshots. Changing this updates the
     * existing quotaset.
     * 
     */
    @Import(name="snapshots")
    private @Nullable Output<Integer> snapshots;

    /**
     * @return Quota value for snapshots. Changing this updates the
     * existing quotaset.
     * 
     */
    public Optional<Output<Integer>> snapshots() {
        return Optional.ofNullable(this.snapshots);
    }

    /**
     * Key/Value pairs for setting quota for
     * volumes types. Possible keys are `snapshots_&lt;volume_type_name&gt;`,
     * `volumes_&lt;volume_type_name&gt;` and `gigabytes_&lt;volume_type_name&gt;`.
     * 
     */
    @Import(name="volumeTypeQuota")
    private @Nullable Output<Map<String,String>> volumeTypeQuota;

    /**
     * @return Key/Value pairs for setting quota for
     * volumes types. Possible keys are `snapshots_&lt;volume_type_name&gt;`,
     * `volumes_&lt;volume_type_name&gt;` and `gigabytes_&lt;volume_type_name&gt;`.
     * 
     */
    public Optional<Output<Map<String,String>>> volumeTypeQuota() {
        return Optional.ofNullable(this.volumeTypeQuota);
    }

    /**
     * Quota value for volumes. Changing this updates the
     * existing quotaset.
     * 
     */
    @Import(name="volumes")
    private @Nullable Output<Integer> volumes;

    /**
     * @return Quota value for volumes. Changing this updates the
     * existing quotaset.
     * 
     */
    public Optional<Output<Integer>> volumes() {
        return Optional.ofNullable(this.volumes);
    }

    private QuoteSetV2State() {}

    private QuoteSetV2State(QuoteSetV2State $) {
        this.backupGigabytes = $.backupGigabytes;
        this.backups = $.backups;
        this.gigabytes = $.gigabytes;
        this.groups = $.groups;
        this.perVolumeGigabytes = $.perVolumeGigabytes;
        this.projectId = $.projectId;
        this.region = $.region;
        this.snapshots = $.snapshots;
        this.volumeTypeQuota = $.volumeTypeQuota;
        this.volumes = $.volumes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuoteSetV2State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuoteSetV2State $;

        public Builder() {
            $ = new QuoteSetV2State();
        }

        public Builder(QuoteSetV2State defaults) {
            $ = new QuoteSetV2State(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupGigabytes Quota value for backup gigabytes. Changing
         * this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder backupGigabytes(@Nullable Output<Integer> backupGigabytes) {
            $.backupGigabytes = backupGigabytes;
            return this;
        }

        /**
         * @param backupGigabytes Quota value for backup gigabytes. Changing
         * this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder backupGigabytes(Integer backupGigabytes) {
            return backupGigabytes(Output.of(backupGigabytes));
        }

        /**
         * @param backups Quota value for backups. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder backups(@Nullable Output<Integer> backups) {
            $.backups = backups;
            return this;
        }

        /**
         * @param backups Quota value for backups. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder backups(Integer backups) {
            return backups(Output.of(backups));
        }

        /**
         * @param gigabytes Quota value for gigabytes. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder gigabytes(@Nullable Output<Integer> gigabytes) {
            $.gigabytes = gigabytes;
            return this;
        }

        /**
         * @param gigabytes Quota value for gigabytes. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder gigabytes(Integer gigabytes) {
            return gigabytes(Output.of(gigabytes));
        }

        /**
         * @param groups Quota value for groups. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<Integer> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups Quota value for groups. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder groups(Integer groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param perVolumeGigabytes Quota value for gigabytes per volume .
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder perVolumeGigabytes(@Nullable Output<Integer> perVolumeGigabytes) {
            $.perVolumeGigabytes = perVolumeGigabytes;
            return this;
        }

        /**
         * @param perVolumeGigabytes Quota value for gigabytes per volume .
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder perVolumeGigabytes(Integer perVolumeGigabytes) {
            return perVolumeGigabytes(Output.of(perVolumeGigabytes));
        }

        /**
         * @param projectId ID of the project to manage quotas. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId ID of the project to manage quotas. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to create the volume. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to create the volume. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param snapshots Quota value for snapshots. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder snapshots(@Nullable Output<Integer> snapshots) {
            $.snapshots = snapshots;
            return this;
        }

        /**
         * @param snapshots Quota value for snapshots. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder snapshots(Integer snapshots) {
            return snapshots(Output.of(snapshots));
        }

        /**
         * @param volumeTypeQuota Key/Value pairs for setting quota for
         * volumes types. Possible keys are `snapshots_&lt;volume_type_name&gt;`,
         * `volumes_&lt;volume_type_name&gt;` and `gigabytes_&lt;volume_type_name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder volumeTypeQuota(@Nullable Output<Map<String,String>> volumeTypeQuota) {
            $.volumeTypeQuota = volumeTypeQuota;
            return this;
        }

        /**
         * @param volumeTypeQuota Key/Value pairs for setting quota for
         * volumes types. Possible keys are `snapshots_&lt;volume_type_name&gt;`,
         * `volumes_&lt;volume_type_name&gt;` and `gigabytes_&lt;volume_type_name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder volumeTypeQuota(Map<String,String> volumeTypeQuota) {
            return volumeTypeQuota(Output.of(volumeTypeQuota));
        }

        /**
         * @param volumes Quota value for volumes. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder volumes(@Nullable Output<Integer> volumes) {
            $.volumes = volumes;
            return this;
        }

        /**
         * @param volumes Quota value for volumes. Changing this updates the
         * existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder volumes(Integer volumes) {
            return volumes(Output.of(volumes));
        }

        public QuoteSetV2State build() {
            return $;
        }
    }

}
