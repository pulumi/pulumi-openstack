// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceNetwork {
    /**
     * @return Specifies if this network should be used for
     * provisioning access. Accepts true or false. Defaults to false.
     * 
     */
    private final @Nullable Boolean accessNetwork;
    /**
     * @return Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new server.
     * 
     */
    private final @Nullable String fixedIpV4;
    private final @Nullable String fixedIpV6;
    /**
     * @deprecated
     * Use the openstack_compute_floatingip_associate_v2 resource instead
     * 
     */
    @Deprecated /* Use the openstack_compute_floatingip_associate_v2 resource instead */
    private final @Nullable String floatingIp;
    private final @Nullable String mac;
    /**
     * @return The human-readable
     * name of the network. Changing this creates a new server.
     * 
     */
    private final @Nullable String name;
    /**
     * @return The port UUID of a
     * network to attach to the server. Changing this creates a new server.
     * 
     */
    private final @Nullable String port;
    /**
     * @return The UUID of
     * the image, volume, or snapshot. Changing this creates a new server.
     * 
     */
    private final @Nullable String uuid;

    @CustomType.Constructor
    private InstanceNetwork(
        @CustomType.Parameter("accessNetwork") @Nullable Boolean accessNetwork,
        @CustomType.Parameter("fixedIpV4") @Nullable String fixedIpV4,
        @CustomType.Parameter("fixedIpV6") @Nullable String fixedIpV6,
        @CustomType.Parameter("floatingIp") @Nullable String floatingIp,
        @CustomType.Parameter("mac") @Nullable String mac,
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("port") @Nullable String port,
        @CustomType.Parameter("uuid") @Nullable String uuid) {
        this.accessNetwork = accessNetwork;
        this.fixedIpV4 = fixedIpV4;
        this.fixedIpV6 = fixedIpV6;
        this.floatingIp = floatingIp;
        this.mac = mac;
        this.name = name;
        this.port = port;
        this.uuid = uuid;
    }

    /**
     * @return Specifies if this network should be used for
     * provisioning access. Accepts true or false. Defaults to false.
     * 
     */
    public Optional<Boolean> accessNetwork() {
        return Optional.ofNullable(this.accessNetwork);
    }
    /**
     * @return Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new server.
     * 
     */
    public Optional<String> fixedIpV4() {
        return Optional.ofNullable(this.fixedIpV4);
    }
    public Optional<String> fixedIpV6() {
        return Optional.ofNullable(this.fixedIpV6);
    }
    /**
     * @deprecated
     * Use the openstack_compute_floatingip_associate_v2 resource instead
     * 
     */
    @Deprecated /* Use the openstack_compute_floatingip_associate_v2 resource instead */
    public Optional<String> floatingIp() {
        return Optional.ofNullable(this.floatingIp);
    }
    public Optional<String> mac() {
        return Optional.ofNullable(this.mac);
    }
    /**
     * @return The human-readable
     * name of the network. Changing this creates a new server.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The port UUID of a
     * network to attach to the server. Changing this creates a new server.
     * 
     */
    public Optional<String> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return The UUID of
     * the image, volume, or snapshot. Changing this creates a new server.
     * 
     */
    public Optional<String> uuid() {
        return Optional.ofNullable(this.uuid);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceNetwork defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean accessNetwork;
        private @Nullable String fixedIpV4;
        private @Nullable String fixedIpV6;
        private @Nullable String floatingIp;
        private @Nullable String mac;
        private @Nullable String name;
        private @Nullable String port;
        private @Nullable String uuid;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessNetwork = defaults.accessNetwork;
    	      this.fixedIpV4 = defaults.fixedIpV4;
    	      this.fixedIpV6 = defaults.fixedIpV6;
    	      this.floatingIp = defaults.floatingIp;
    	      this.mac = defaults.mac;
    	      this.name = defaults.name;
    	      this.port = defaults.port;
    	      this.uuid = defaults.uuid;
        }

        public Builder accessNetwork(@Nullable Boolean accessNetwork) {
            this.accessNetwork = accessNetwork;
            return this;
        }
        public Builder fixedIpV4(@Nullable String fixedIpV4) {
            this.fixedIpV4 = fixedIpV4;
            return this;
        }
        public Builder fixedIpV6(@Nullable String fixedIpV6) {
            this.fixedIpV6 = fixedIpV6;
            return this;
        }
        public Builder floatingIp(@Nullable String floatingIp) {
            this.floatingIp = floatingIp;
            return this;
        }
        public Builder mac(@Nullable String mac) {
            this.mac = mac;
            return this;
        }
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder port(@Nullable String port) {
            this.port = port;
            return this;
        }
        public Builder uuid(@Nullable String uuid) {
            this.uuid = uuid;
            return this;
        }        public InstanceNetwork build() {
            return new InstanceNetwork(accessNetwork, fixedIpV4, fixedIpV6, floatingIp, mac, name, port, uuid);
        }
    }
}
