// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAddressScopeArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAddressScopeArgs Empty = new GetAddressScopeArgs();

    /**
     * IP version.
     * 
     */
    @Import(name="ipVersion")
    private @Nullable Output<Integer> ipVersion;

    /**
     * @return IP version.
     * 
     */
    public Optional<Output<Integer>> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }

    /**
     * Name of the address-scope.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the address-scope.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The owner of the address-scope.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The owner of the address-scope.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve address-scopes. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Neutron client.
     * A Neutron client is needed to retrieve address-scopes. If omitted, the
     * `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Indicates whether this address-scope is shared across
     * all projects.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return Indicates whether this address-scope is shared across
     * all projects.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    private GetAddressScopeArgs() {}

    private GetAddressScopeArgs(GetAddressScopeArgs $) {
        this.ipVersion = $.ipVersion;
        this.name = $.name;
        this.projectId = $.projectId;
        this.region = $.region;
        this.shared = $.shared;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAddressScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAddressScopeArgs $;

        public Builder() {
            $ = new GetAddressScopeArgs();
        }

        public Builder(GetAddressScopeArgs defaults) {
            $ = new GetAddressScopeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipVersion IP version.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(@Nullable Output<Integer> ipVersion) {
            $.ipVersion = ipVersion;
            return this;
        }

        /**
         * @param ipVersion IP version.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(Integer ipVersion) {
            return ipVersion(Output.of(ipVersion));
        }

        /**
         * @param name Name of the address-scope.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the address-scope.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The owner of the address-scope.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The owner of the address-scope.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V2 Neutron client.
         * A Neutron client is needed to retrieve address-scopes. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Neutron client.
         * A Neutron client is needed to retrieve address-scopes. If omitted, the
         * `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param shared Indicates whether this address-scope is shared across
         * all projects.
         * 
         * @return builder
         * 
         */
        public Builder shared(@Nullable Output<Boolean> shared) {
            $.shared = shared;
            return this;
        }

        /**
         * @param shared Indicates whether this address-scope is shared across
         * all projects.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        public GetAddressScopeArgs build() {
            return $;
        }
    }

}
