// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSnapshotResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String name;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String projectId;
    private String region;
    /**
     * @return The UUID of the source share that was used to create the snapshot.
     * 
     */
    private String shareId;
    /**
     * @return The file system protocol of a share snapshot.
     * 
     */
    private String shareProto;
    /**
     * @return The share snapshot size, in GBs.
     * 
     */
    private Integer shareSize;
    /**
     * @return The snapshot size, in GBs.
     * 
     */
    private Integer size;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String status;

    private GetSnapshotResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    public String region() {
        return this.region;
    }
    /**
     * @return The UUID of the source share that was used to create the snapshot.
     * 
     */
    public String shareId() {
        return this.shareId;
    }
    /**
     * @return The file system protocol of a share snapshot.
     * 
     */
    public String shareProto() {
        return this.shareProto;
    }
    /**
     * @return The share snapshot size, in GBs.
     * 
     */
    public Integer shareSize() {
        return this.shareSize;
    }
    /**
     * @return The snapshot size, in GBs.
     * 
     */
    public Integer size() {
        return this.size;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSnapshotResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String name;
        private String projectId;
        private String region;
        private String shareId;
        private String shareProto;
        private Integer shareSize;
        private Integer size;
        private String status;
        public Builder() {}
        public Builder(GetSnapshotResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.shareId = defaults.shareId;
    	      this.shareProto = defaults.shareProto;
    	      this.shareSize = defaults.shareSize;
    	      this.size = defaults.size;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder shareId(String shareId) {
            this.shareId = Objects.requireNonNull(shareId);
            return this;
        }
        @CustomType.Setter
        public Builder shareProto(String shareProto) {
            this.shareProto = Objects.requireNonNull(shareProto);
            return this;
        }
        @CustomType.Setter
        public Builder shareSize(Integer shareSize) {
            this.shareSize = Objects.requireNonNull(shareSize);
            return this;
        }
        @CustomType.Setter
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        public GetSnapshotResult build() {
            final var _resultValue = new GetSnapshotResult();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.projectId = projectId;
            _resultValue.region = region;
            _resultValue.shareId = shareId;
            _resultValue.shareProto = shareProto;
            _resultValue.shareSize = shareSize;
            _resultValue.size = size;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
