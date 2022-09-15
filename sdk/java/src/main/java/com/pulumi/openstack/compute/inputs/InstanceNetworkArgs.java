// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceNetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceNetworkArgs Empty = new InstanceNetworkArgs();

    /**
     * Specifies if this network should be used for
     * provisioning access. Accepts true or false. Defaults to false.
     * 
     */
    @Import(name="accessNetwork")
    private @Nullable Output<Boolean> accessNetwork;

    /**
     * @return Specifies if this network should be used for
     * provisioning access. Accepts true or false. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> accessNetwork() {
        return Optional.ofNullable(this.accessNetwork);
    }

    /**
     * Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new server.
     * 
     */
    @Import(name="fixedIpV4")
    private @Nullable Output<String> fixedIpV4;

    /**
     * @return Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new server.
     * 
     */
    public Optional<Output<String>> fixedIpV4() {
        return Optional.ofNullable(this.fixedIpV4);
    }

    @Import(name="fixedIpV6")
    private @Nullable Output<String> fixedIpV6;

    public Optional<Output<String>> fixedIpV6() {
        return Optional.ofNullable(this.fixedIpV6);
    }

    /**
     * @deprecated
     * Use the openstack_compute_floatingip_associate_v2 resource instead
     * 
     */
    @Deprecated /* Use the openstack_compute_floatingip_associate_v2 resource instead */
    @Import(name="floatingIp")
    private @Nullable Output<String> floatingIp;

    /**
     * @deprecated
     * Use the openstack_compute_floatingip_associate_v2 resource instead
     * 
     */
    @Deprecated /* Use the openstack_compute_floatingip_associate_v2 resource instead */
    public Optional<Output<String>> floatingIp() {
        return Optional.ofNullable(this.floatingIp);
    }

    @Import(name="mac")
    private @Nullable Output<String> mac;

    public Optional<Output<String>> mac() {
        return Optional.ofNullable(this.mac);
    }

    /**
     * The human-readable
     * name of the network. Changing this creates a new server.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The human-readable
     * name of the network. Changing this creates a new server.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The port UUID of a
     * network to attach to the server. Changing this creates a new server.
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return The port UUID of a
     * network to attach to the server. Changing this creates a new server.
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The UUID of
     * the image, volume, or snapshot. Changing this creates a new server.
     * 
     */
    @Import(name="uuid")
    private @Nullable Output<String> uuid;

    /**
     * @return The UUID of
     * the image, volume, or snapshot. Changing this creates a new server.
     * 
     */
    public Optional<Output<String>> uuid() {
        return Optional.ofNullable(this.uuid);
    }

    private InstanceNetworkArgs() {}

    private InstanceNetworkArgs(InstanceNetworkArgs $) {
        this.accessNetwork = $.accessNetwork;
        this.fixedIpV4 = $.fixedIpV4;
        this.fixedIpV6 = $.fixedIpV6;
        this.floatingIp = $.floatingIp;
        this.mac = $.mac;
        this.name = $.name;
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
         * @param accessNetwork Specifies if this network should be used for
         * provisioning access. Accepts true or false. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder accessNetwork(@Nullable Output<Boolean> accessNetwork) {
            $.accessNetwork = accessNetwork;
            return this;
        }

        /**
         * @param accessNetwork Specifies if this network should be used for
         * provisioning access. Accepts true or false. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder accessNetwork(Boolean accessNetwork) {
            return accessNetwork(Output.of(accessNetwork));
        }

        /**
         * @param fixedIpV4 Specifies a fixed IPv4 address to be used on this
         * network. Changing this creates a new server.
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
         * network. Changing this creates a new server.
         * 
         * @return builder
         * 
         */
        public Builder fixedIpV4(String fixedIpV4) {
            return fixedIpV4(Output.of(fixedIpV4));
        }

        public Builder fixedIpV6(@Nullable Output<String> fixedIpV6) {
            $.fixedIpV6 = fixedIpV6;
            return this;
        }

        public Builder fixedIpV6(String fixedIpV6) {
            return fixedIpV6(Output.of(fixedIpV6));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Use the openstack_compute_floatingip_associate_v2 resource instead
         * 
         */
        @Deprecated /* Use the openstack_compute_floatingip_associate_v2 resource instead */
        public Builder floatingIp(@Nullable Output<String> floatingIp) {
            $.floatingIp = floatingIp;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Use the openstack_compute_floatingip_associate_v2 resource instead
         * 
         */
        @Deprecated /* Use the openstack_compute_floatingip_associate_v2 resource instead */
        public Builder floatingIp(String floatingIp) {
            return floatingIp(Output.of(floatingIp));
        }

        public Builder mac(@Nullable Output<String> mac) {
            $.mac = mac;
            return this;
        }

        public Builder mac(String mac) {
            return mac(Output.of(mac));
        }

        /**
         * @param name The human-readable
         * name of the network. Changing this creates a new server.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The human-readable
         * name of the network. Changing this creates a new server.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param port The port UUID of a
         * network to attach to the server. Changing this creates a new server.
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
         * network to attach to the server. Changing this creates a new server.
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param uuid The UUID of
         * the image, volume, or snapshot. Changing this creates a new server.
         * 
         * @return builder
         * 
         */
        public Builder uuid(@Nullable Output<String> uuid) {
            $.uuid = uuid;
            return this;
        }

        /**
         * @param uuid The UUID of
         * the image, volume, or snapshot. Changing this creates a new server.
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