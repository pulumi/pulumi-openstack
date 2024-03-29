// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPortExtraDhcpOption {
    /**
     * @return IP protocol version
     * 
     */
    private Integer ipVersion;
    /**
     * @return The name of the port.
     * 
     */
    private String name;
    /**
     * @return Value of the DHCP option.
     * 
     */
    private String value;

    private GetPortExtraDhcpOption() {}
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
    @CustomType.Builder
    public static final class Builder {
        private Integer ipVersion;
        private String name;
        private String value;
        public Builder() {}
        public Builder(GetPortExtraDhcpOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipVersion = defaults.ipVersion;
    	      this.name = defaults.name;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder ipVersion(Integer ipVersion) {
            if (ipVersion == null) {
              throw new MissingRequiredPropertyException("GetPortExtraDhcpOption", "ipVersion");
            }
            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetPortExtraDhcpOption", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("GetPortExtraDhcpOption", "value");
            }
            this.value = value;
            return this;
        }
        public GetPortExtraDhcpOption build() {
            final var _resultValue = new GetPortExtraDhcpOption();
            _resultValue.ipVersion = ipVersion;
            _resultValue.name = name;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
