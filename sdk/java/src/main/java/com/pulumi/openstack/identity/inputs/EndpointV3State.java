// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointV3State extends com.pulumi.resources.ResourceArgs {

    public static final EndpointV3State Empty = new EndpointV3State();

    /**
     * The endpoint region. The `region` and
     * `endpoint_region` can be different.
     * 
     */
    @Import(name="endpointRegion")
    private @Nullable Output<String> endpointRegion;

    /**
     * @return The endpoint region. The `region` and
     * `endpoint_region` can be different.
     * 
     */
    public Optional<Output<String>> endpointRegion() {
        return Optional.ofNullable(this.endpointRegion);
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
    @Import(name="serviceId")
    private @Nullable Output<String> serviceId;

    /**
     * @return The endpoint service ID.
     * 
     */
    public Optional<Output<String>> serviceId() {
        return Optional.ofNullable(this.serviceId);
    }

    /**
     * The service name of the endpoint.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The service name of the endpoint.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * The service type of the endpoint.
     * 
     */
    @Import(name="serviceType")
    private @Nullable Output<String> serviceType;

    /**
     * @return The service type of the endpoint.
     * 
     */
    public Optional<Output<String>> serviceType() {
        return Optional.ofNullable(this.serviceType);
    }

    /**
     * The endpoint url.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The endpoint url.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private EndpointV3State() {}

    private EndpointV3State(EndpointV3State $) {
        this.endpointRegion = $.endpointRegion;
        this.interface_ = $.interface_;
        this.name = $.name;
        this.region = $.region;
        this.serviceId = $.serviceId;
        this.serviceName = $.serviceName;
        this.serviceType = $.serviceType;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointV3State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointV3State $;

        public Builder() {
            $ = new EndpointV3State();
        }

        public Builder(EndpointV3State defaults) {
            $ = new EndpointV3State(Objects.requireNonNull(defaults));
        }

        /**
         * @param endpointRegion The endpoint region. The `region` and
         * `endpoint_region` can be different.
         * 
         * @return builder
         * 
         */
        public Builder endpointRegion(@Nullable Output<String> endpointRegion) {
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
        public Builder serviceId(@Nullable Output<String> serviceId) {
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
         * @param serviceName The service name of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param serviceType The service type of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceType(@Nullable Output<String> serviceType) {
            $.serviceType = serviceType;
            return this;
        }

        /**
         * @param serviceType The service type of the endpoint.
         * 
         * @return builder
         * 
         */
        public Builder serviceType(String serviceType) {
            return serviceType(Output.of(serviceType));
        }

        /**
         * @param url The endpoint url.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
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

        public EndpointV3State build() {
            return $;
        }
    }

}
