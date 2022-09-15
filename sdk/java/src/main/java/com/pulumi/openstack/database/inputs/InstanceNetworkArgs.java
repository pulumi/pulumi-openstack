// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceNetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceNetworkArgs Empty = new InstanceNetworkArgs();

    /**
     * Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    @Import(name="fixedIpV4")
    private @Nullable Output<String> fixedIpV4;

    /**
     * @return Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    public Optional<Output<String>> fixedIpV4() {
        return Optional.ofNullable(this.fixedIpV4);
    }

    /**
     * Specifies a fixed IPv6 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    @Import(name="fixedIpV6")
    private @Nullable Output<String> fixedIpV6;

    /**
     * @return Specifies a fixed IPv6 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    public Optional<Output<String>> fixedIpV6() {
        return Optional.ofNullable(this.fixedIpV6);
    }

    /**
     * The port UUID of a
     * network to attach to the instance. Changing this creates a new instance.
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return The port UUID of a
     * network to attach to the instance. Changing this creates a new instance.
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The network UUID to
     * attach to the instance. Changing this creates a new instance.
     * 
     */
    @Import(name="uuid")
    private @Nullable Output<String> uuid;

    /**
     * @return The network UUID to
     * attach to the instance. Changing this creates a new instance.
     * 
     */
    public Optional<Output<String>> uuid() {
        return Optional.ofNullable(this.uuid);
    }

    private InstanceNetworkArgs() {}

    private InstanceNetworkArgs(InstanceNetworkArgs $) {
        this.fixedIpV4 = $.fixedIpV4;
        this.fixedIpV6 = $.fixedIpV6;
        this.port = $.port;
        this.uuid = $.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceNetworkArgs $;

        public Builder() {
            $ = new InstanceNetworkArgs();
        }

        public Builder(InstanceNetworkArgs defaults) {
            $ = new InstanceNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param fixedIpV4 Specifies a fixed IPv4 address to be used on this
         * network. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder fixedIpV4(@Nullable Output<String> fixedIpV4) {
            $.fixedIpV4 = fixedIpV4;
            return this;
        }

        /**
         * @param fixedIpV4 Specifies a fixed IPv4 address to be used on this
         * network. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder fixedIpV4(String fixedIpV4) {
            return fixedIpV4(Output.of(fixedIpV4));
        }

        /**
         * @param fixedIpV6 Specifies a fixed IPv6 address to be used on this
         * network. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder fixedIpV6(@Nullable Output<String> fixedIpV6) {
            $.fixedIpV6 = fixedIpV6;
            return this;
        }

        /**
         * @param fixedIpV6 Specifies a fixed IPv6 address to be used on this
         * network. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder fixedIpV6(String fixedIpV6) {
            return fixedIpV6(Output.of(fixedIpV6));
        }

        /**
         * @param port The port UUID of a
         * network to attach to the instance. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port UUID of a
         * network to attach to the instance. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param uuid The network UUID to
         * attach to the instance. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder uuid(@Nullable Output<String> uuid) {
            $.uuid = uuid;
            return this;
        }

        /**
         * @param uuid The network UUID to
         * attach to the instance. Changing this creates a new instance.
         * 
         * @return builder
         * 
         */
        public Builder uuid(String uuid) {
            return uuid(Output.of(uuid));
        }

        public InstanceNetworkArgs build() {
            return $;
        }
    }

}