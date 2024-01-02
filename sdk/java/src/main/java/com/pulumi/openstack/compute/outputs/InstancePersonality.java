// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class InstancePersonality {
    /**
     * @return The contents of the file. Limited to 255 bytes.
     * 
     */
    private String content;
    /**
     * @return The absolute path of the destination file.
     * 
     */
    private String file;

    private InstancePersonality() {}
    /**
     * @return The contents of the file. Limited to 255 bytes.
     * 
     */
    public String content() {
        return this.content;
    }
    /**
     * @return The absolute path of the destination file.
     * 
     */
    public String file() {
        return this.file;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstancePersonality defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String content;
        private String file;
        public Builder() {}
        public Builder(InstancePersonality defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.content = defaults.content;
    	      this.file = defaults.file;
        }

        @CustomType.Setter
        public Builder content(String content) {
            if (content == null) {
              throw new MissingRequiredPropertyException("InstancePersonality", "content");
            }
            this.content = content;
            return this;
        }
        @CustomType.Setter
        public Builder file(String file) {
            if (file == null) {
              throw new MissingRequiredPropertyException("InstancePersonality", "file");
            }
            this.file = file;
            return this;
        }
        public InstancePersonality build() {
            final var _resultValue = new InstancePersonality();
            _resultValue.content = content;
            _resultValue.file = file;
            return _resultValue;
        }
    }
}
