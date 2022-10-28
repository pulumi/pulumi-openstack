// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PortExtraDhcpOption {
    /**
     * @return IP protocol version. Defaults to 4.
     * 
     */
    private @Nullable Integer ipVersion;
    /**
     * @return Name of the DHCP option.
     * 
     */
    private String name;
    /**
     * @return Value of the DHCP option.
     * 
     */
    private String value;

    private PortExtraDhcpOption() {}
    /**
     * @return IP protocol version. Defaults to 4.
     * 
     */
    public Optional<Integer> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }
    /**
     * @return Name of the DHCP option.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Value of the DHCP option.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PortExtraDhcpOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer ipVersion;
        private String name;
        private String value;
        public Builder() {}
        public Builder(PortExtraDhcpOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipVersion = defaults.ipVersion;
    	      this.name = defaults.name;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder ipVersion(@Nullable Integer ipVersion) {
            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public PortExtraDhcpOption build() {
            final var o = new PortExtraDhcpOption();
            o.ipVersion = ipVersion;
            o.name = name;
            o.value = value;
            return o;
        }
    }
}
