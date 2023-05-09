// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.objectstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ContainerVersioningLegacyArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerVersioningLegacyArgs Empty = new ContainerVersioningLegacyArgs();

    /**
     * Container in which versions will be stored.
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return Container in which versions will be stored.
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private ContainerVersioningLegacyArgs() {}

    private ContainerVersioningLegacyArgs(ContainerVersioningLegacyArgs $) {
        this.location = $.location;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerVersioningLegacyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerVersioningLegacyArgs $;

        public Builder() {
            $ = new ContainerVersioningLegacyArgs();
        }

        public Builder(ContainerVersioningLegacyArgs defaults) {
            $ = new ContainerVersioningLegacyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param location Container in which versions will be stored.
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Container in which versions will be stored.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param type Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Versioning type which can be `versions` or `history` according to [Openstack documentation](https://docs.openstack.org/swift/latest/api/object_versioning.html).
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ContainerVersioningLegacyArgs build() {
            $.location = Objects.requireNonNull($.location, "expected parameter 'location' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}