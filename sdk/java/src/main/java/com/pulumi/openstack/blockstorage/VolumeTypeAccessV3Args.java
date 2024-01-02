// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeTypeAccessV3Args extends com.pulumi.resources.ResourceArgs {

    public static final VolumeTypeAccessV3Args Empty = new VolumeTypeAccessV3Args();

    /**
     * ID of the project to give access to. Changing this
     * creates a new resource.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return ID of the project to give access to. Changing this
     * creates a new resource.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
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
     * ID of the volume type to give access to. Changing
     * this creates a new resource.
     * 
     */
    @Import(name="volumeTypeId", required=true)
    private Output<String> volumeTypeId;

    /**
     * @return ID of the volume type to give access to. Changing
     * this creates a new resource.
     * 
     */
    public Output<String> volumeTypeId() {
        return this.volumeTypeId;
    }

    private VolumeTypeAccessV3Args() {}

    private VolumeTypeAccessV3Args(VolumeTypeAccessV3Args $) {
        this.projectId = $.projectId;
        this.region = $.region;
        this.volumeTypeId = $.volumeTypeId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeTypeAccessV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeTypeAccessV3Args $;

        public Builder() {
            $ = new VolumeTypeAccessV3Args();
        }

        public Builder(VolumeTypeAccessV3Args defaults) {
            $ = new VolumeTypeAccessV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId ID of the project to give access to. Changing this
         * creates a new resource.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId ID of the project to give access to. Changing this
         * creates a new resource.
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
         * @param volumeTypeId ID of the volume type to give access to. Changing
         * this creates a new resource.
         * 
         * @return builder
         * 
         */
        public Builder volumeTypeId(Output<String> volumeTypeId) {
            $.volumeTypeId = volumeTypeId;
            return this;
        }

        /**
         * @param volumeTypeId ID of the volume type to give access to. Changing
         * this creates a new resource.
         * 
         * @return builder
         * 
         */
        public Builder volumeTypeId(String volumeTypeId) {
            return volumeTypeId(Output.of(volumeTypeId));
        }

        public VolumeTypeAccessV3Args build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("VolumeTypeAccessV3Args", "projectId");
            }
            if ($.volumeTypeId == null) {
                throw new MissingRequiredPropertyException("VolumeTypeAccessV3Args", "volumeTypeId");
            }
            return $;
        }
    }

}
