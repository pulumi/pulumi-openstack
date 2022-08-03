// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QosV3State extends com.pulumi.resources.ResourceArgs {

    public static final QosV3State Empty = new QosV3State();

    /**
     * The consumer of qos. Can be one of `front-end`,
     * `back-end` or `both`. Changing this updates the `consumer` of an
     * existing qos.
     * 
     */
    @Import(name="consumer")
    private @Nullable Output<String> consumer;

    /**
     * @return The consumer of qos. Can be one of `front-end`,
     * `back-end` or `both`. Changing this updates the `consumer` of an
     * existing qos.
     * 
     */
    public Optional<Output<String>> consumer() {
        return Optional.ofNullable(this.consumer);
    }

    /**
     * Name of the qos.  Changing this creates a new qos.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the qos.  Changing this creates a new qos.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to create the qos. If omitted,
     * the `region` argument of the provider is used. Changing this creates
     * a new qos.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the qos. If omitted,
     * the `region` argument of the provider is used. Changing this creates
     * a new qos.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Key/Value pairs of specs for the qos.
     * 
     */
    @Import(name="specs")
    private @Nullable Output<Map<String,Object>> specs;

    /**
     * @return Key/Value pairs of specs for the qos.
     * 
     */
    public Optional<Output<Map<String,Object>>> specs() {
        return Optional.ofNullable(this.specs);
    }

    private QosV3State() {}

    private QosV3State(QosV3State $) {
        this.consumer = $.consumer;
        this.name = $.name;
        this.region = $.region;
        this.specs = $.specs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QosV3State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QosV3State $;

        public Builder() {
            $ = new QosV3State();
        }

        public Builder(QosV3State defaults) {
            $ = new QosV3State(Objects.requireNonNull(defaults));
        }

        /**
         * @param consumer The consumer of qos. Can be one of `front-end`,
         * `back-end` or `both`. Changing this updates the `consumer` of an
         * existing qos.
         * 
         * @return builder
         * 
         */
        public Builder consumer(@Nullable Output<String> consumer) {
            $.consumer = consumer;
            return this;
        }

        /**
         * @param consumer The consumer of qos. Can be one of `front-end`,
         * `back-end` or `both`. Changing this updates the `consumer` of an
         * existing qos.
         * 
         * @return builder
         * 
         */
        public Builder consumer(String consumer) {
            return consumer(Output.of(consumer));
        }

        /**
         * @param name Name of the qos.  Changing this creates a new qos.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the qos.  Changing this creates a new qos.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to create the qos. If omitted,
         * the `region` argument of the provider is used. Changing this creates
         * a new qos.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to create the qos. If omitted,
         * the `region` argument of the provider is used. Changing this creates
         * a new qos.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param specs Key/Value pairs of specs for the qos.
         * 
         * @return builder
         * 
         */
        public Builder specs(@Nullable Output<Map<String,Object>> specs) {
            $.specs = specs;
            return this;
        }

        /**
         * @param specs Key/Value pairs of specs for the qos.
         * 
         * @return builder
         * 
         */
        public Builder specs(Map<String,Object> specs) {
            return specs(Output.of(specs));
        }

        public QosV3State build() {
            return $;
        }
    }

}
