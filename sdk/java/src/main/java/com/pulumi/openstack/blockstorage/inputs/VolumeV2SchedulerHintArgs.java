// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeV2SchedulerHintArgs extends com.pulumi.resources.ResourceArgs {

    public static final VolumeV2SchedulerHintArgs Empty = new VolumeV2SchedulerHintArgs();

    /**
     * Arbitrary key/value pairs of additional
     * properties to pass to the scheduler.
     * 
     */
    @Import(name="additionalProperties")
    private @Nullable Output<Map<String,Object>> additionalProperties;

    /**
     * @return Arbitrary key/value pairs of additional
     * properties to pass to the scheduler.
     * 
     */
    public Optional<Output<Map<String,Object>>> additionalProperties() {
        return Optional.ofNullable(this.additionalProperties);
    }

    /**
     * The volume should be scheduled on a
     * different host from the set of volumes specified in the list provided.
     * 
     */
    @Import(name="differentHosts")
    private @Nullable Output<List<String>> differentHosts;

    /**
     * @return The volume should be scheduled on a
     * different host from the set of volumes specified in the list provided.
     * 
     */
    public Optional<Output<List<String>>> differentHosts() {
        return Optional.ofNullable(this.differentHosts);
    }

    /**
     * An instance UUID. The volume should be
     * scheduled on the same host as the instance.
     * 
     */
    @Import(name="localToInstance")
    private @Nullable Output<String> localToInstance;

    /**
     * @return An instance UUID. The volume should be
     * scheduled on the same host as the instance.
     * 
     */
    public Optional<Output<String>> localToInstance() {
        return Optional.ofNullable(this.localToInstance);
    }

    /**
     * A conditional query that a back-end must pass in
     * order to host a volume. The query must use the `JsonFilter` syntax
     * which is described
     * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
     * At this time, only simple queries are supported. Compound queries using
     * `and`, `or`, or `not` are not supported. An example of a simple query is:
     * 
     */
    @Import(name="query")
    private @Nullable Output<String> query;

    /**
     * @return A conditional query that a back-end must pass in
     * order to host a volume. The query must use the `JsonFilter` syntax
     * which is described
     * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
     * At this time, only simple queries are supported. Compound queries using
     * `and`, `or`, or `not` are not supported. An example of a simple query is:
     * 
     */
    public Optional<Output<String>> query() {
        return Optional.ofNullable(this.query);
    }

    /**
     * A list of volume UUIDs. The volume should be
     * scheduled on the same host as another volume specified in the list provided.
     * 
     */
    @Import(name="sameHosts")
    private @Nullable Output<List<String>> sameHosts;

    /**
     * @return A list of volume UUIDs. The volume should be
     * scheduled on the same host as another volume specified in the list provided.
     * 
     */
    public Optional<Output<List<String>>> sameHosts() {
        return Optional.ofNullable(this.sameHosts);
    }

    private VolumeV2SchedulerHintArgs() {}

    private VolumeV2SchedulerHintArgs(VolumeV2SchedulerHintArgs $) {
        this.additionalProperties = $.additionalProperties;
        this.differentHosts = $.differentHosts;
        this.localToInstance = $.localToInstance;
        this.query = $.query;
        this.sameHosts = $.sameHosts;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeV2SchedulerHintArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeV2SchedulerHintArgs $;

        public Builder() {
            $ = new VolumeV2SchedulerHintArgs();
        }

        public Builder(VolumeV2SchedulerHintArgs defaults) {
            $ = new VolumeV2SchedulerHintArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalProperties Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         * 
         * @return builder
         * 
         */
        public Builder additionalProperties(@Nullable Output<Map<String,Object>> additionalProperties) {
            $.additionalProperties = additionalProperties;
            return this;
        }

        /**
         * @param additionalProperties Arbitrary key/value pairs of additional
         * properties to pass to the scheduler.
         * 
         * @return builder
         * 
         */
        public Builder additionalProperties(Map<String,Object> additionalProperties) {
            return additionalProperties(Output.of(additionalProperties));
        }

        /**
         * @param differentHosts The volume should be scheduled on a
         * different host from the set of volumes specified in the list provided.
         * 
         * @return builder
         * 
         */
        public Builder differentHosts(@Nullable Output<List<String>> differentHosts) {
            $.differentHosts = differentHosts;
            return this;
        }

        /**
         * @param differentHosts The volume should be scheduled on a
         * different host from the set of volumes specified in the list provided.
         * 
         * @return builder
         * 
         */
        public Builder differentHosts(List<String> differentHosts) {
            return differentHosts(Output.of(differentHosts));
        }

        /**
         * @param differentHosts The volume should be scheduled on a
         * different host from the set of volumes specified in the list provided.
         * 
         * @return builder
         * 
         */
        public Builder differentHosts(String... differentHosts) {
            return differentHosts(List.of(differentHosts));
        }

        /**
         * @param localToInstance An instance UUID. The volume should be
         * scheduled on the same host as the instance.
         * 
         * @return builder
         * 
         */
        public Builder localToInstance(@Nullable Output<String> localToInstance) {
            $.localToInstance = localToInstance;
            return this;
        }

        /**
         * @param localToInstance An instance UUID. The volume should be
         * scheduled on the same host as the instance.
         * 
         * @return builder
         * 
         */
        public Builder localToInstance(String localToInstance) {
            return localToInstance(Output.of(localToInstance));
        }

        /**
         * @param query A conditional query that a back-end must pass in
         * order to host a volume. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         * 
         * @return builder
         * 
         */
        public Builder query(@Nullable Output<String> query) {
            $.query = query;
            return this;
        }

        /**
         * @param query A conditional query that a back-end must pass in
         * order to host a volume. The query must use the `JsonFilter` syntax
         * which is described
         * [here](https://docs.openstack.org/cinder/latest/configuration/block-storage/scheduler-filters.html#jsonfilter).
         * At this time, only simple queries are supported. Compound queries using
         * `and`, `or`, or `not` are not supported. An example of a simple query is:
         * 
         * @return builder
         * 
         */
        public Builder query(String query) {
            return query(Output.of(query));
        }

        /**
         * @param sameHosts A list of volume UUIDs. The volume should be
         * scheduled on the same host as another volume specified in the list provided.
         * 
         * @return builder
         * 
         */
        public Builder sameHosts(@Nullable Output<List<String>> sameHosts) {
            $.sameHosts = sameHosts;
            return this;
        }

        /**
         * @param sameHosts A list of volume UUIDs. The volume should be
         * scheduled on the same host as another volume specified in the list provided.
         * 
         * @return builder
         * 
         */
        public Builder sameHosts(List<String> sameHosts) {
            return sameHosts(Output.of(sameHosts));
        }

        /**
         * @param sameHosts A list of volume UUIDs. The volume should be
         * scheduled on the same host as another volume specified in the list provided.
         * 
         * @return builder
         * 
         */
        public Builder sameHosts(String... sameHosts) {
            return sameHosts(List.of(sameHosts));
        }

        public VolumeV2SchedulerHintArgs build() {
            return $;
        }
    }

}