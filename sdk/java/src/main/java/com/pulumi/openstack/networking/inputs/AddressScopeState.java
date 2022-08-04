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


public final class AddressScopeState extends com.pulumi.resources.ResourceArgs {

    public static final AddressScopeState Empty = new AddressScopeState();

    /**
     * IP version, either 4 (default) or 6. Changing this
     * creates a new address-scope.
     * 
     */
    @Import(name="ipVersion")
    private @Nullable Output<Integer> ipVersion;

    /**
     * @return IP version, either 4 (default) or 6. Changing this
     * creates a new address-scope.
     * 
     */
    public Optional<Output<Integer>> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }

    /**
     * The name of the address-scope. Changing this updates the
     * name of the existing address-scope.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the address-scope. Changing this updates the
     * name of the existing address-scope.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The owner of the address-scope. Required if admin
     * wants to create a address-scope for another project. Changing this creates a
     * new address-scope.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The owner of the address-scope. Required if admin
     * wants to create a address-scope for another project. Changing this creates a
     * new address-scope.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron address-scope. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * address-scope.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create a Neutron address-scope. If omitted,
     * the `region` argument of the provider is used. Changing this creates a new
     * address-scope.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Indicates whether this address-scope is shared across
     * all projects. Changing this updates the shared status of the existing
     * address-scope.
     * 
     */
    @Import(name="shared")
    private @Nullable Output<Boolean> shared;

    /**
     * @return Indicates whether this address-scope is shared across
     * all projects. Changing this updates the shared status of the existing
     * address-scope.
     * 
     */
    public Optional<Output<Boolean>> shared() {
        return Optional.ofNullable(this.shared);
    }

    private AddressScopeState() {}

    private AddressScopeState(AddressScopeState $) {
        this.ipVersion = $.ipVersion;
        this.name = $.name;
        this.projectId = $.projectId;
        this.region = $.region;
        this.shared = $.shared;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AddressScopeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AddressScopeState $;

        public Builder() {
            $ = new AddressScopeState();
        }

        public Builder(AddressScopeState defaults) {
            $ = new AddressScopeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipVersion IP version, either 4 (default) or 6. Changing this
         * creates a new address-scope.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(@Nullable Output<Integer> ipVersion) {
            $.ipVersion = ipVersion;
            return this;
        }

        /**
         * @param ipVersion IP version, either 4 (default) or 6. Changing this
         * creates a new address-scope.
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(Integer ipVersion) {
            return ipVersion(Output.of(ipVersion));
        }

        /**
         * @param name The name of the address-scope. Changing this updates the
         * name of the existing address-scope.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the address-scope. Changing this updates the
         * name of the existing address-scope.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param projectId The owner of the address-scope. Required if admin
         * wants to create a address-scope for another project. Changing this creates a
         * new address-scope.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The owner of the address-scope. Required if admin
         * wants to create a address-scope for another project. Changing this creates a
         * new address-scope.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron address-scope. If omitted,
         * the `region` argument of the provider is used. Changing this creates a new
         * address-scope.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create a Neutron address-scope. If omitted,
         * the `region` argument of the provider is used. Changing this creates a new
         * address-scope.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param shared Indicates whether this address-scope is shared across
         * all projects. Changing this updates the shared status of the existing
         * address-scope.
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
         * all projects. Changing this updates the shared status of the existing
         * address-scope.
         * 
         * @return builder
         * 
         */
        public Builder shared(Boolean shared) {
            return shared(Output.of(shared));
        }

        public AddressScopeState build() {
            return $;
        }
    }

}
