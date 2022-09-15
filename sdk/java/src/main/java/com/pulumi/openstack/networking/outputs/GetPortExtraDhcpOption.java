// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPortExtraDhcpOption {
    /**
     * @return IP protocol version
     * 
     */
    private final Integer ipVersion;
    /**
     * @return The name of the port.
     * 
     */
    private final String name;
    /**
     * @return Value of the DHCP option.
     * 
     */
    private final String value;

    @CustomType.Constructor
    private GetPortExtraDhcpOption(
        @CustomType.Parameter("ipVersion") Integer ipVersion,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("value") String value) {
        this.ipVersion = ipVersion;
        this.name = name;
        this.value = value;
    }

    /**
     * @return IP protocol version
     * 
     */
    public Integer ipVersion() {
        return this.ipVersion;
    }
    /**
     * @return The name of the port.
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

    public static Builder builder(GetPortExtraDhcpOption defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Integer ipVersion;
        private String name;
        private String value;

        public Builder() {
    	      // Empty
        }

        public Builder(GetPortExtraDhcpOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipVersion = defaults.ipVersion;
    	      this.name = defaults.name;
    	      this.value = defaults.value;
        }

        public Builder ipVersion(Integer ipVersion) {
            this.ipVersion = Objects.requireNonNull(ipVersion);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }        public GetPortExtraDhcpOption build() {
            return new GetPortExtraDhcpOption(ipVersion, name, value);
        }
    }
}