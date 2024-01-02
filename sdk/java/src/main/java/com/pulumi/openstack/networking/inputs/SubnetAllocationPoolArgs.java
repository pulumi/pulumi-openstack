// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class SubnetAllocationPoolArgs extends com.pulumi.resources.ResourceArgs {

    public static final SubnetAllocationPoolArgs Empty = new SubnetAllocationPoolArgs();

    /**
     * The ending address.
     * 
     */
    @Import(name="end", required=true)
    private Output<String> end;

    /**
     * @return The ending address.
     * 
     */
    public Output<String> end() {
        return this.end;
    }

    /**
     * The starting address.
     * 
     */
    @Import(name="start", required=true)
    private Output<String> start;

    /**
     * @return The starting address.
     * 
     */
    public Output<String> start() {
        return this.start;
    }

    private SubnetAllocationPoolArgs() {}

    private SubnetAllocationPoolArgs(SubnetAllocationPoolArgs $) {
        this.end = $.end;
        this.start = $.start;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubnetAllocationPoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubnetAllocationPoolArgs $;

        public Builder() {
            $ = new SubnetAllocationPoolArgs();
        }

        public Builder(SubnetAllocationPoolArgs defaults) {
            $ = new SubnetAllocationPoolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param end The ending address.
         * 
         * @return builder
         * 
         */
        public Builder end(Output<String> end) {
            $.end = end;
            return this;
        }

        /**
         * @param end The ending address.
         * 
         * @return builder
         * 
         */
        public Builder end(String end) {
            return end(Output.of(end));
        }

        /**
         * @param start The starting address.
         * 
         * @return builder
         * 
         */
        public Builder start(Output<String> start) {
            $.start = start;
            return this;
        }

        /**
         * @param start The starting address.
         * 
         * @return builder
         * 
         */
        public Builder start(String start) {
            return start(Output.of(start));
        }

        public SubnetAllocationPoolArgs build() {
            if ($.end == null) {
                throw new MissingRequiredPropertyException("SubnetAllocationPoolArgs", "end");
            }
            if ($.start == null) {
                throw new MissingRequiredPropertyException("SubnetAllocationPoolArgs", "start");
            }
            return $;
        }
    }

}
