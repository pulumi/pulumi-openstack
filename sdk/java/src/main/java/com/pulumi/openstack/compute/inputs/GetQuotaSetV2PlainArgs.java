// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetQuotaSetV2PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetQuotaSetV2PlainArgs Empty = new GetQuotaSetV2PlainArgs();

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
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetQuotaSetV2PlainArgs() {}

    private GetQuotaSetV2PlainArgs(GetQuotaSetV2PlainArgs $) {
        this.projectId = $.projectId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetQuotaSetV2PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetQuotaSetV2PlainArgs $;

        public Builder() {
            $ = new GetQuotaSetV2PlainArgs();
        }

        public Builder(GetQuotaSetV2PlainArgs defaults) {
            $ = new GetQuotaSetV2PlainArgs(Objects.requireNonNull(defaults));
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
         * @param region The region in which to obtain the V2 Compute client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetQuotaSetV2PlainArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetQuotaSetV2PlainArgs", "projectId");
            }
            return $;
        }
    }

}
