// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ShareExportLocation {
    private final @Nullable String path;
    private final @Nullable String preferred;

    @CustomType.Constructor
    private ShareExportLocation(
        @CustomType.Parameter("path") @Nullable String path,
        @CustomType.Parameter("preferred") @Nullable String preferred) {
        this.path = path;
        this.preferred = preferred;
    }

    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }
    public Optional<String> preferred() {
        return Optional.ofNullable(this.preferred);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ShareExportLocation defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String path;
        private @Nullable String preferred;

        public Builder() {
    	      // Empty
        }

        public Builder(ShareExportLocation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.path = defaults.path;
    	      this.preferred = defaults.preferred;
        }

        public Builder path(@Nullable String path) {
            this.path = path;
            return this;
        }
        public Builder preferred(@Nullable String preferred) {
            this.preferred = preferred;
            return this;
        }        public ShareExportLocation build() {
            return new ShareExportLocation(path, preferred);
        }
    }
}