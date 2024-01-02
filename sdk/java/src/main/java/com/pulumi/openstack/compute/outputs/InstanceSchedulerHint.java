// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceSchedulerHint {
    /**
     * @return Arbitrary key/value pairs of additional
     * properties to pass to the scheduler.
     * 
     */
    private @Nullable Map<String,Object> additionalProperties;
    /**
     * @return An IP Address in CIDR form. The instance
     * will be placed on a compute node that is in the same subnet.
     * 
     */
    private @Nullable String buildNearHostIp;
    /**
     * @return The names of cells where not to build the instance.
     * 
     */
    private @Nullable List<String> differentCells;
    /**
     * @return A list of instance UUIDs. The instance will
     * be scheduled on a different host than all other instances.
     * 
     */
    private @Nullable List<String> differentHosts;
    /**
     * @return A UUID of a Server Group. The instance will be placed
     * into that group.
     * 
     */
    private @Nullable String group;
    /**
     * @return A conditional query that a compute node must pass in
     * order to host an instance. The query must use the `JsonFilter` syntax
     * which is described
     * [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter).
     * At this time, only simple queries are supported. Compound queries using
     * `and`, `or`, or `not` are not supported. An example of a simple query is:
     * 
     */
    private @Nullable List<String> queries;
    /**
     * @return A list of instance UUIDs. The instance will be
     * scheduled on the same host of those specified.
     * 
     */
    private @Nullable List<String> sameHosts;
    /**
     * @return The name of a cell to host the instance.
     * 
     */
    private @Nullable String targetCell;

    private InstanceSchedulerHint() {}
    /**
     * @return Arbitrary key/value pairs of additional
     * properties to pass to the scheduler.
     * 
     */
    public Map<String,Object> additionalProperties() {
        return this.additionalProperties == null ? Map.of() : this.additionalProperties;
    }
    /**
     * @return An IP Address in CIDR form. The instance
     * will be placed on a compute node that is in the same subnet.
     * 
     */
    public Optional<String> buildNearHostIp() {
        return Optional.ofNullable(this.buildNearHostIp);
    }
    /**
     * @return The names of cells where not to build the instance.
     * 
     */
    public List<String> differentCells() {
        return this.differentCells == null ? List.of() : this.differentCells;
    }
    /**
     * @return A list of instance UUIDs. The instance will
     * be scheduled on a different host than all other instances.
     * 
     */
    public List<String> differentHosts() {
        return this.differentHosts == null ? List.of() : this.differentHosts;
    }
    /**
     * @return A UUID of a Server Group. The instance will be placed
     * into that group.
     * 
     */
    public Optional<String> group() {
        return Optional.ofNullable(this.group);
    }
    /**
     * @return A conditional query that a compute node must pass in
     * order to host an instance. The query must use the `JsonFilter` syntax
     * which is described
     * [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter).
     * At this time, only simple queries are supported. Compound queries using
     * `and`, `or`, or `not` are not supported. An example of a simple query is:
     * 
     */
    public List<String> queries() {
        return this.queries == null ? List.of() : this.queries;
    }
    /**
     * @return A list of instance UUIDs. The instance will be
     * scheduled on the same host of those specified.
     * 
     */
    public List<String> sameHosts() {
        return this.sameHosts == null ? List.of() : this.sameHosts;
    }
    /**
     * @return The name of a cell to host the instance.
     * 
     */
    public Optional<String> targetCell() {
        return Optional.ofNullable(this.targetCell);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceSchedulerHint defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,Object> additionalProperties;
        private @Nullable String buildNearHostIp;
        private @Nullable List<String> differentCells;
        private @Nullable List<String> differentHosts;
        private @Nullable String group;
        private @Nullable List<String> queries;
        private @Nullable List<String> sameHosts;
        private @Nullable String targetCell;
        public Builder() {}
        public Builder(InstanceSchedulerHint defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.additionalProperties = defaults.additionalProperties;
    	      this.buildNearHostIp = defaults.buildNearHostIp;
    	      this.differentCells = defaults.differentCells;
    	      this.differentHosts = defaults.differentHosts;
    	      this.group = defaults.group;
    	      this.queries = defaults.queries;
    	      this.sameHosts = defaults.sameHosts;
    	      this.targetCell = defaults.targetCell;
        }

        @CustomType.Setter
        public Builder additionalProperties(@Nullable Map<String,Object> additionalProperties) {

            this.additionalProperties = additionalProperties;
            return this;
        }
        @CustomType.Setter
        public Builder buildNearHostIp(@Nullable String buildNearHostIp) {

            this.buildNearHostIp = buildNearHostIp;
            return this;
        }
        @CustomType.Setter
        public Builder differentCells(@Nullable List<String> differentCells) {

            this.differentCells = differentCells;
            return this;
        }
        public Builder differentCells(String... differentCells) {
            return differentCells(List.of(differentCells));
        }
        @CustomType.Setter
        public Builder differentHosts(@Nullable List<String> differentHosts) {

            this.differentHosts = differentHosts;
            return this;
        }
        public Builder differentHosts(String... differentHosts) {
            return differentHosts(List.of(differentHosts));
        }
        @CustomType.Setter
        public Builder group(@Nullable String group) {

            this.group = group;
            return this;
        }
        @CustomType.Setter
        public Builder queries(@Nullable List<String> queries) {

            this.queries = queries;
            return this;
        }
        public Builder queries(String... queries) {
            return queries(List.of(queries));
        }
        @CustomType.Setter
        public Builder sameHosts(@Nullable List<String> sameHosts) {

            this.sameHosts = sameHosts;
            return this;
        }
        public Builder sameHosts(String... sameHosts) {
            return sameHosts(List.of(sameHosts));
        }
        @CustomType.Setter
        public Builder targetCell(@Nullable String targetCell) {

            this.targetCell = targetCell;
            return this;
        }
        public InstanceSchedulerHint build() {
            final var _resultValue = new InstanceSchedulerHint();
            _resultValue.additionalProperties = additionalProperties;
            _resultValue.buildNearHostIp = buildNearHostIp;
            _resultValue.differentCells = differentCells;
            _resultValue.differentHosts = differentHosts;
            _resultValue.group = group;
            _resultValue.queries = queries;
            _resultValue.sameHosts = sameHosts;
            _resultValue.targetCell = targetCell;
            return _resultValue;
        }
    }
}
