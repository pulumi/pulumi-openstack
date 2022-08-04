// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceNetwork {
    /**
     * @return Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    private final @Nullable String fixedIpV4;
    /**
     * @return Specifies a fixed IPv6 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    private final @Nullable String fixedIpV6;
    /**
     * @return The port UUID of a
     * network to attach to the instance. Changing this creates a new instance.
     * 
     */
    private final @Nullable String port;
    /**
     * @return The network UUID to
     * attach to the instance. Changing this creates a new instance.
     * 
     */
    private final @Nullable String uuid;

    @CustomType.Constructor
    private InstanceNetwork(
        @CustomType.Parameter("fixedIpV4") @Nullable String fixedIpV4,
        @CustomType.Parameter("fixedIpV6") @Nullable String fixedIpV6,
        @CustomType.Parameter("port") @Nullable String port,
        @CustomType.Parameter("uuid") @Nullable String uuid) {
        this.fixedIpV4 = fixedIpV4;
        this.fixedIpV6 = fixedIpV6;
        this.port = port;
        this.uuid = uuid;
    }

    /**
     * @return Specifies a fixed IPv4 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    public Optional<String> fixedIpV4() {
        return Optional.ofNullable(this.fixedIpV4);
    }
    /**
     * @return Specifies a fixed IPv6 address to be used on this
     * network. Changing this creates a new instance.
     * 
     */
    public Optional<String> fixedIpV6() {
        return Optional.ofNullable(this.fixedIpV6);
    }
    /**
     * @return The port UUID of a
     * network to attach to the instance. Changing this creates a new instance.
     * 
     */
    public Optional<String> port() {
        return Optional.ofNullable(this.port);
    }
    /**
     * @return The network UUID to
     * attach to the instance. Changing this creates a new instance.
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
        private @Nullable String fixedIpV4;
        private @Nullable String fixedIpV6;
        private @Nullable String port;
        private @Nullable String uuid;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fixedIpV4 = defaults.fixedIpV4;
    	      this.fixedIpV6 = defaults.fixedIpV6;
    	      this.port = defaults.port;
    	      this.uuid = defaults.uuid;
        }

        public Builder fixedIpV4(@Nullable String fixedIpV4) {
            this.fixedIpV4 = fixedIpV4;
            return this;
        }
        public Builder fixedIpV6(@Nullable String fixedIpV6) {
            this.fixedIpV6 = fixedIpV6;
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
            return new InstanceNetwork(fixedIpV4, fixedIpV6, port, uuid);
        }
    }
}
