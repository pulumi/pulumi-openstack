// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetQuotasetV3PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetQuotasetV3PlainArgs Empty = new GetQuotasetV3PlainArgs();

    /**
     * The id of the project to retrieve the quotaset.
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return The id of the project to retrieve the quotaset.
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    /**
     * The region in which to obtain the V3 Blockstorage client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V3 Blockstorage client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetQuotasetV3PlainArgs() {}

    private GetQuotasetV3PlainArgs(GetQuotasetV3PlainArgs $) {
        this.projectId = $.projectId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetQuotasetV3PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetQuotasetV3PlainArgs $;

        public Builder() {
            $ = new GetQuotasetV3PlainArgs();
        }

        public Builder(GetQuotasetV3PlainArgs defaults) {
            $ = new GetQuotasetV3PlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The id of the project to retrieve the quotaset.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Blockstorage client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetQuotasetV3PlainArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetQuotasetV3PlainArgs", "projectId");
            }
            return $;
        }
    }

}
