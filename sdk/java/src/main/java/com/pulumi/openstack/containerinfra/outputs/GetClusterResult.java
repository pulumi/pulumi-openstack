// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetClusterResult {
    /**
     * @return COE API address.
     * 
     */
    private String apiAddress;
    /**
     * @return The UUID of the V1 Container Infra cluster template.
     * 
     */
    private String clusterTemplateId;
    /**
     * @return COE software version.
     * 
     */
    private String coeVersion;
    private String containerVersion;
    /**
     * @return The timeout (in minutes) for creating the cluster.
     * 
     */
    private Integer createTimeout;
    /**
     * @return The time at which cluster was created.
     * 
     */
    private String createdAt;
    /**
     * @return The URL used for cluster node discovery.
     * 
     */
    private String discoveryUrl;
    /**
     * @return The size (in GB) of the Docker volume.
     * 
     */
    private Integer dockerVolumeSize;
    /**
     * @return The fixed network that is attached to the cluster.
     * 
     */
    private String fixedNetwork;
    /**
     * @return The fixed subnet that is attached to the cluster.
     * 
     */
    private String fixedSubnet;
    /**
     * @return The flavor for the nodes of the cluster.
     * 
     */
    private String flavor;
    private Boolean floatingIpEnabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the Compute service SSH keypair.
     * 
     */
    private String keypair;
    /**
     * @return The Kubernetes cluster&#39;s credentials
     * 
     */
    private Map<String,String> kubeconfig;
    /**
     * @return The list of key value pairs representing additional properties of
     * the cluster.
     * 
     */
    private Map<String,String> labels;
    /**
     * @return IP addresses of the master node of the cluster.
     * 
     */
    private List<String> masterAddresses;
    /**
     * @return The number of master nodes for the cluster.
     * 
     */
    private Integer masterCount;
    /**
     * @return The flavor for the master nodes.
     * 
     */
    private String masterFlavor;
    /**
     * @return Whether a load balancer is created for the master
     * cluster nodes.
     * 
     */
    private Boolean masterLbEnabled;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String name;
    /**
     * @return IP addresses of the node of the cluster.
     * 
     */
    private List<String> nodeAddresses;
    /**
     * @return The number of nodes for the cluster.
     * 
     */
    private Integer nodeCount;
    /**
     * @return The project of the cluster.
     * 
     */
    private String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String region;
    /**
     * @return UUID of the Orchestration service stack.
     * 
     */
    private String stackId;
    /**
     * @return The time at which cluster was updated.
     * 
     */
    private String updatedAt;
    /**
     * @return The user of the cluster.
     * 
     */
    private String userId;

    private GetClusterResult() {}
    /**
     * @return COE API address.
     * 
     */
    public String apiAddress() {
        return this.apiAddress;
    }
    /**
     * @return The UUID of the V1 Container Infra cluster template.
     * 
     */
    public String clusterTemplateId() {
        return this.clusterTemplateId;
    }
    /**
     * @return COE software version.
     * 
     */
    public String coeVersion() {
        return this.coeVersion;
    }
    public String containerVersion() {
        return this.containerVersion;
    }
    /**
     * @return The timeout (in minutes) for creating the cluster.
     * 
     */
    public Integer createTimeout() {
        return this.createTimeout;
    }
    /**
     * @return The time at which cluster was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The URL used for cluster node discovery.
     * 
     */
    public String discoveryUrl() {
        return this.discoveryUrl;
    }
    /**
     * @return The size (in GB) of the Docker volume.
     * 
     */
    public Integer dockerVolumeSize() {
        return this.dockerVolumeSize;
    }
    /**
     * @return The fixed network that is attached to the cluster.
     * 
     */
    public String fixedNetwork() {
        return this.fixedNetwork;
    }
    /**
     * @return The fixed subnet that is attached to the cluster.
     * 
     */
    public String fixedSubnet() {
        return this.fixedSubnet;
    }
    /**
     * @return The flavor for the nodes of the cluster.
     * 
     */
    public String flavor() {
        return this.flavor;
    }
    public Boolean floatingIpEnabled() {
        return this.floatingIpEnabled;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the Compute service SSH keypair.
     * 
     */
    public String keypair() {
        return this.keypair;
    }
    /**
     * @return The Kubernetes cluster&#39;s credentials
     * 
     */
    public Map<String,String> kubeconfig() {
        return this.kubeconfig;
    }
    /**
     * @return The list of key value pairs representing additional properties of
     * the cluster.
     * 
     */
    public Map<String,String> labels() {
        return this.labels;
    }
    /**
     * @return IP addresses of the master node of the cluster.
     * 
     */
    public List<String> masterAddresses() {
        return this.masterAddresses;
    }
    /**
     * @return The number of master nodes for the cluster.
     * 
     */
    public Integer masterCount() {
        return this.masterCount;
    }
    /**
     * @return The flavor for the master nodes.
     * 
     */
    public String masterFlavor() {
        return this.masterFlavor;
    }
    /**
     * @return Whether a load balancer is created for the master
     * cluster nodes.
     * 
     */
    public Boolean masterLbEnabled() {
        return this.masterLbEnabled;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return IP addresses of the node of the cluster.
     * 
     */
    public List<String> nodeAddresses() {
        return this.nodeAddresses;
    }
    /**
     * @return The number of nodes for the cluster.
     * 
     */
    public Integer nodeCount() {
        return this.nodeCount;
    }
    /**
     * @return The project of the cluster.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return UUID of the Orchestration service stack.
     * 
     */
    public String stackId() {
        return this.stackId;
    }
    /**
     * @return The time at which cluster was updated.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return The user of the cluster.
     * 
     */
    public String userId() {
        return this.userId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClusterResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String apiAddress;
        private String clusterTemplateId;
        private String coeVersion;
        private String containerVersion;
        private Integer createTimeout;
        private String createdAt;
        private String discoveryUrl;
        private Integer dockerVolumeSize;
        private String fixedNetwork;
        private String fixedSubnet;
        private String flavor;
        private Boolean floatingIpEnabled;
        private String id;
        private String keypair;
        private Map<String,String> kubeconfig;
        private Map<String,String> labels;
        private List<String> masterAddresses;
        private Integer masterCount;
        private String masterFlavor;
        private Boolean masterLbEnabled;
        private String name;
        private List<String> nodeAddresses;
        private Integer nodeCount;
        private String projectId;
        private String region;
        private String stackId;
        private String updatedAt;
        private String userId;
        public Builder() {}
        public Builder(GetClusterResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiAddress = defaults.apiAddress;
    	      this.clusterTemplateId = defaults.clusterTemplateId;
    	      this.coeVersion = defaults.coeVersion;
    	      this.containerVersion = defaults.containerVersion;
    	      this.createTimeout = defaults.createTimeout;
    	      this.createdAt = defaults.createdAt;
    	      this.discoveryUrl = defaults.discoveryUrl;
    	      this.dockerVolumeSize = defaults.dockerVolumeSize;
    	      this.fixedNetwork = defaults.fixedNetwork;
    	      this.fixedSubnet = defaults.fixedSubnet;
    	      this.flavor = defaults.flavor;
    	      this.floatingIpEnabled = defaults.floatingIpEnabled;
    	      this.id = defaults.id;
    	      this.keypair = defaults.keypair;
    	      this.kubeconfig = defaults.kubeconfig;
    	      this.labels = defaults.labels;
    	      this.masterAddresses = defaults.masterAddresses;
    	      this.masterCount = defaults.masterCount;
    	      this.masterFlavor = defaults.masterFlavor;
    	      this.masterLbEnabled = defaults.masterLbEnabled;
    	      this.name = defaults.name;
    	      this.nodeAddresses = defaults.nodeAddresses;
    	      this.nodeCount = defaults.nodeCount;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.stackId = defaults.stackId;
    	      this.updatedAt = defaults.updatedAt;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder apiAddress(String apiAddress) {
            if (apiAddress == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "apiAddress");
            }
            this.apiAddress = apiAddress;
            return this;
        }
        @CustomType.Setter
        public Builder clusterTemplateId(String clusterTemplateId) {
            if (clusterTemplateId == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "clusterTemplateId");
            }
            this.clusterTemplateId = clusterTemplateId;
            return this;
        }
        @CustomType.Setter
        public Builder coeVersion(String coeVersion) {
            if (coeVersion == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "coeVersion");
            }
            this.coeVersion = coeVersion;
            return this;
        }
        @CustomType.Setter
        public Builder containerVersion(String containerVersion) {
            if (containerVersion == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "containerVersion");
            }
            this.containerVersion = containerVersion;
            return this;
        }
        @CustomType.Setter
        public Builder createTimeout(Integer createTimeout) {
            if (createTimeout == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "createTimeout");
            }
            this.createTimeout = createTimeout;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder discoveryUrl(String discoveryUrl) {
            if (discoveryUrl == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "discoveryUrl");
            }
            this.discoveryUrl = discoveryUrl;
            return this;
        }
        @CustomType.Setter
        public Builder dockerVolumeSize(Integer dockerVolumeSize) {
            if (dockerVolumeSize == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "dockerVolumeSize");
            }
            this.dockerVolumeSize = dockerVolumeSize;
            return this;
        }
        @CustomType.Setter
        public Builder fixedNetwork(String fixedNetwork) {
            if (fixedNetwork == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "fixedNetwork");
            }
            this.fixedNetwork = fixedNetwork;
            return this;
        }
        @CustomType.Setter
        public Builder fixedSubnet(String fixedSubnet) {
            if (fixedSubnet == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "fixedSubnet");
            }
            this.fixedSubnet = fixedSubnet;
            return this;
        }
        @CustomType.Setter
        public Builder flavor(String flavor) {
            if (flavor == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "flavor");
            }
            this.flavor = flavor;
            return this;
        }
        @CustomType.Setter
        public Builder floatingIpEnabled(Boolean floatingIpEnabled) {
            if (floatingIpEnabled == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "floatingIpEnabled");
            }
            this.floatingIpEnabled = floatingIpEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder keypair(String keypair) {
            if (keypair == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "keypair");
            }
            this.keypair = keypair;
            return this;
        }
        @CustomType.Setter
        public Builder kubeconfig(Map<String,String> kubeconfig) {
            if (kubeconfig == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "kubeconfig");
            }
            this.kubeconfig = kubeconfig;
            return this;
        }
        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            if (labels == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "labels");
            }
            this.labels = labels;
            return this;
        }
        @CustomType.Setter
        public Builder masterAddresses(List<String> masterAddresses) {
            if (masterAddresses == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "masterAddresses");
            }
            this.masterAddresses = masterAddresses;
            return this;
        }
        public Builder masterAddresses(String... masterAddresses) {
            return masterAddresses(List.of(masterAddresses));
        }
        @CustomType.Setter
        public Builder masterCount(Integer masterCount) {
            if (masterCount == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "masterCount");
            }
            this.masterCount = masterCount;
            return this;
        }
        @CustomType.Setter
        public Builder masterFlavor(String masterFlavor) {
            if (masterFlavor == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "masterFlavor");
            }
            this.masterFlavor = masterFlavor;
            return this;
        }
        @CustomType.Setter
        public Builder masterLbEnabled(Boolean masterLbEnabled) {
            if (masterLbEnabled == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "masterLbEnabled");
            }
            this.masterLbEnabled = masterLbEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nodeAddresses(List<String> nodeAddresses) {
            if (nodeAddresses == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "nodeAddresses");
            }
            this.nodeAddresses = nodeAddresses;
            return this;
        }
        public Builder nodeAddresses(String... nodeAddresses) {
            return nodeAddresses(List.of(nodeAddresses));
        }
        @CustomType.Setter
        public Builder nodeCount(Integer nodeCount) {
            if (nodeCount == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "nodeCount");
            }
            this.nodeCount = nodeCount;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder stackId(String stackId) {
            if (stackId == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "stackId");
            }
            this.stackId = stackId;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            if (userId == null) {
              throw new MissingRequiredPropertyException("GetClusterResult", "userId");
            }
            this.userId = userId;
            return this;
        }
        public GetClusterResult build() {
            final var _resultValue = new GetClusterResult();
            _resultValue.apiAddress = apiAddress;
            _resultValue.clusterTemplateId = clusterTemplateId;
            _resultValue.coeVersion = coeVersion;
            _resultValue.containerVersion = containerVersion;
            _resultValue.createTimeout = createTimeout;
            _resultValue.createdAt = createdAt;
            _resultValue.discoveryUrl = discoveryUrl;
            _resultValue.dockerVolumeSize = dockerVolumeSize;
            _resultValue.fixedNetwork = fixedNetwork;
            _resultValue.fixedSubnet = fixedSubnet;
            _resultValue.flavor = flavor;
            _resultValue.floatingIpEnabled = floatingIpEnabled;
            _resultValue.id = id;
            _resultValue.keypair = keypair;
            _resultValue.kubeconfig = kubeconfig;
            _resultValue.labels = labels;
            _resultValue.masterAddresses = masterAddresses;
            _resultValue.masterCount = masterCount;
            _resultValue.masterFlavor = masterFlavor;
            _resultValue.masterLbEnabled = masterLbEnabled;
            _resultValue.name = name;
            _resultValue.nodeAddresses = nodeAddresses;
            _resultValue.nodeCount = nodeCount;
            _resultValue.projectId = projectId;
            _resultValue.region = region;
            _resultValue.stackId = stackId;
            _resultValue.updatedAt = updatedAt;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
