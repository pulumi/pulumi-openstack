// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointV3Args extends com.pulumi.resources.ResourceArgs {

    public static final EndpointV3Args Empty = new EndpointV3Args();

    /**
     * The endpoint region. The `region` and
     * `endpoint_region` can be different.
     * 
     */
    @Import(name="endpointRegion", required=true)
    private Output<String> endpointRegion;

    /**
     * @return The endpoint region. The `region` and
     * `endpoint_region` can be different.
     * 
     */
    public Output<String> endpointRegion() {
        return this.endpointRegion;
    }

    /**
     * The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     * 
     */
    @Import(name="interface")
    private @Nullable Output<String> interface_;

    /**
     * @return The endpoint interface. Valid values are `public`,
     * `internal` and `admin`. Default value is `public`
     * 
     */
    public Optional<Output<String>> interface_() {
        return Optional.ofNullable(this.interface_);
    }

    /**
     * The endpoint name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The endpoint name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The endpoint service ID.
     * 
     */
    @Import(name="serviceId", required=true)
    private Output<String> serviceId;

    /**
     * @return The endpoint service ID.
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }

    /**
     * The endpoint url.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The endpoint url.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private EndpointV3Args() {}

    private EndpointV3Args(EndpointV3Args $) {
        this.endpointRegion = $.endpointRegion;
        this.interface_ = $.interface_;
        this.name = $.name;
        this.region = $.region;
        this.serviceId = $.serviceId;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointV3Args $;

        public Builder() {
            $ = new EndpointV3Args();
        }

        public Builder(EndpointV3Args defaults) {
            $ = new EndpointV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param endpointRegion The endpoint region. The `region` and
         * `endpoint_region` can be different.
         * 
         * @return builder
         * 
         */
        public Builder endpointRegion(Output<String> endpointRegion) {
            $.endpointRegion = endpointRegion;
            return this;
        }

        /**
         * @param endpointRegion The endpoint region. The `region` and
         * `endpoint_region` can be different.
         * 
         * @return builder
         * 
         */
        public Builder endpointRegion(String endpointRegion) {
            return endpointRegion(Output.of(endpointRegion));
        }

        /**
         * @param interface_ The endpoint interface. Valid values are `public`,
         * `internal` and `admin`. Default value is `public`
         * 
         * @return builder
         * 
         */
        public Builder interface_(@Nullable Output<String> interface_) {
            $.interface_ = interface_;
            return this;
        }

        /**
         * @param interface_ The endpoint interface. Valid values are `public`,
         * `internal` and `admin`. Default value is `public`
         * 
         * @return builder
         * 
         */
        public Builder interface_(String interface_) {
            return interface_(Output.of(interface_));
        }

        /**
         * @param name The endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The endpoint name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V3 Keystone client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Keystone client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param serviceId The endpoint service ID.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(Output<String> serviceId) {
            $.serviceId = serviceId;
            return this;
        }

        /**
         * @param serviceId The endpoint service ID.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(String serviceId) {
            return serviceId(Output.of(serviceId));
        }

        /**
         * @param url The endpoint url.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The endpoint url.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public EndpointV3Args build() {
            if ($.endpointRegion == null) {
                throw new MissingRequiredPropertyException("EndpointV3Args", "endpointRegion");
            }
            if ($.serviceId == null) {
                throw new MissingRequiredPropertyException("EndpointV3Args", "serviceId");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("EndpointV3Args", "url");
            }
            return $;
        }
    }

}
