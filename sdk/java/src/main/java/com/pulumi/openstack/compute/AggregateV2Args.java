// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AggregateV2Args extends com.pulumi.resources.ResourceArgs {

    public static final AggregateV2Args Empty = new AggregateV2Args();

    /**
     * The list of hosts contained in the Host Aggregate. The hosts must be added
     * to Openstack and visible in the web interface, or the provider will fail to add them to the host
     * aggregate.
     * 
     */
    @Import(name="hosts")
    private @Nullable Output<List<String>> hosts;

    /**
     * @return The list of hosts contained in the Host Aggregate. The hosts must be added
     * to Openstack and visible in the web interface, or the provider will fail to add them to the host
     * aggregate.
     * 
     */
    public Optional<Output<List<String>>> hosts() {
        return Optional.ofNullable(this.hosts);
    }

    /**
     * The metadata of the Host Aggregate. Can be useful to indicate scheduler hints.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,String>> metadata;

    /**
     * @return The metadata of the Host Aggregate. Can be useful to indicate scheduler hints.
     * 
     */
    public Optional<Output<Map<String,String>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * The name of the Host Aggregate
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Host Aggregate
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to create the Host Aggregate. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new Host Aggregate.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the Host Aggregate. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new Host Aggregate.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The name of the Availability Zone to use. If ommited, it will take the default
     * availability zone.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return The name of the Availability Zone to use. If ommited, it will take the default
     * availability zone.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private AggregateV2Args() {}

    private AggregateV2Args(AggregateV2Args $) {
        this.hosts = $.hosts;
        this.metadata = $.metadata;
        this.name = $.name;
        this.region = $.region;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AggregateV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AggregateV2Args $;

        public Builder() {
            $ = new AggregateV2Args();
        }

        public Builder(AggregateV2Args defaults) {
            $ = new AggregateV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param hosts The list of hosts contained in the Host Aggregate. The hosts must be added
         * to Openstack and visible in the web interface, or the provider will fail to add them to the host
         * aggregate.
         * 
         * @return builder
         * 
         */
        public Builder hosts(@Nullable Output<List<String>> hosts) {
            $.hosts = hosts;
            return this;
        }

        /**
         * @param hosts The list of hosts contained in the Host Aggregate. The hosts must be added
         * to Openstack and visible in the web interface, or the provider will fail to add them to the host
         * aggregate.
         * 
         * @return builder
         * 
         */
        public Builder hosts(List<String> hosts) {
            return hosts(Output.of(hosts));
        }

        /**
         * @param hosts The list of hosts contained in the Host Aggregate. The hosts must be added
         * to Openstack and visible in the web interface, or the provider will fail to add them to the host
         * aggregate.
         * 
         * @return builder
         * 
         */
        public Builder hosts(String... hosts) {
            return hosts(List.of(hosts));
        }

        /**
         * @param metadata The metadata of the Host Aggregate. Can be useful to indicate scheduler hints.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,String>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata The metadata of the Host Aggregate. Can be useful to indicate scheduler hints.
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,String> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param name The name of the Host Aggregate
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Host Aggregate
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to create the Host Aggregate. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new Host Aggregate.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to create the Host Aggregate. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new Host Aggregate.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param zone The name of the Availability Zone to use. If ommited, it will take the default
         * availability zone.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone The name of the Availability Zone to use. If ommited, it will take the default
         * availability zone.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public AggregateV2Args build() {
            return $;
        }
    }

}
