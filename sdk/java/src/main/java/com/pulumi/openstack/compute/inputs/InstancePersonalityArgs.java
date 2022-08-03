// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class InstancePersonalityArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstancePersonalityArgs Empty = new InstancePersonalityArgs();

    /**
     * The contents of the file. Limited to 255 bytes.
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return The contents of the file. Limited to 255 bytes.
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * The absolute path of the destination file.
     * 
     */
    @Import(name="file", required=true)
    private Output<String> file;

    /**
     * @return The absolute path of the destination file.
     * 
     */
    public Output<String> file() {
        return this.file;
    }

    private InstancePersonalityArgs() {}

    private InstancePersonalityArgs(InstancePersonalityArgs $) {
        this.content = $.content;
        this.file = $.file;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstancePersonalityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstancePersonalityArgs $;

        public Builder() {
            $ = new InstancePersonalityArgs();
        }

        public Builder(InstancePersonalityArgs defaults) {
            $ = new InstancePersonalityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param content The contents of the file. Limited to 255 bytes.
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The contents of the file. Limited to 255 bytes.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param file The absolute path of the destination file.
         * 
         * @return builder
         * 
         */
        public Builder file(Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file The absolute path of the destination file.
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        public InstancePersonalityArgs build() {
            $.content = Objects.requireNonNull($.content, "expected parameter 'content' to be non-null");
            $.file = Objects.requireNonNull($.file, "expected parameter 'file' to be non-null");
            return $;
        }
    }

}
