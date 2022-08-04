// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetQuotasetV3Args extends com.pulumi.resources.InvokeArgs {

    public static final GetQuotasetV3Args Empty = new GetQuotasetV3Args();

    /**
     * The id of the project to retrieve the quotaset.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The id of the project to retrieve the quotaset.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * The region in which to obtain the V3 Blockstorage client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V3 Blockstorage client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetQuotasetV3Args() {}

    private GetQuotasetV3Args(GetQuotasetV3Args $) {
        this.projectId = $.projectId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetQuotasetV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetQuotasetV3Args $;

        public Builder() {
            $ = new GetQuotasetV3Args();
        }

        public Builder(GetQuotasetV3Args defaults) {
            $ = new GetQuotasetV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The id of the project to retrieve the quotaset.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The id of the project to retrieve the quotaset.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V3 Blockstorage client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Blockstorage client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetQuotasetV3Args build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
