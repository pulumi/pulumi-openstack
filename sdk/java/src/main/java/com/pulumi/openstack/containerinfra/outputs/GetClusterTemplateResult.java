// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetClusterTemplateResult {
    /**
     * @return The API server port for the Container Orchestration
     * Engine for this cluster template.
     * 
     */
    private Integer apiserverPort;
    /**
     * @return The distro for the cluster (fedora-atomic, coreos, etc.).
     * 
     */
    private String clusterDistro;
    /**
     * @return The Container Orchestration Engine for this cluster template.
     * 
     */
    private String coe;
    /**
     * @return The time at which cluster template was created.
     * 
     */
    private String createdAt;
    /**
     * @return Address of the DNS nameserver that is used in nodes of the
     * cluster.
     * 
     */
    private String dnsNameserver;
    /**
     * @return Docker storage driver. Changing this updates the
     * Docker storage driver of the existing cluster template.
     * 
     */
    private String dockerStorageDriver;
    /**
     * @return The size (in GB) of the Docker volume.
     * 
     */
    private Integer dockerVolumeSize;
    /**
     * @return The ID of the external network that will be used for
     * the cluster.
     * 
     */
    private String externalNetworkId;
    /**
     * @return The fixed network that will be attached to the cluster.
     * 
     */
    private String fixedNetwork;
    /**
     * @return =The fixed subnet that will be attached to the cluster.
     * 
     */
    private String fixedSubnet;
    /**
     * @return The flavor for the nodes of the cluster.
     * 
     */
    private String flavor;
    /**
     * @return Indicates whether created cluster should create IP
     * floating IP for every node or not.
     * 
     */
    private Boolean floatingIpEnabled;
    /**
     * @return Indicates whether the ClusterTemplate is hidden or not.
     * 
     */
    private Boolean hidden;
    /**
     * @return The address of a proxy for receiving all HTTP requests and
     * relay them.
     * 
     */
    private String httpProxy;
    /**
     * @return The address of a proxy for receiving all HTTPS requests and
     * relay them.
     * 
     */
    private String httpsProxy;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The reference to an image that is used for nodes of the cluster.
     * 
     */
    private String image;
    /**
     * @return The insecure registry URL for the cluster template.
     * 
     */
    private String insecureRegistry;
    /**
     * @return The name of the Compute service SSH keypair.
     * 
     */
    private String keypairId;
    /**
     * @return The list of key value pairs representing additional properties
     * of the cluster template.
     * 
     */
    private Map<String,Object> labels;
    /**
     * @return The flavor for the master nodes.
     * 
     */
    private String masterFlavor;
    /**
     * @return Indicates whether created cluster should has a
     * loadbalancer for master nodes or not.
     * 
     */
    private Boolean masterLbEnabled;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String name;
    /**
     * @return The name of the driver for the container network.
     * 
     */
    private String networkDriver;
    /**
     * @return A comma-separated list of IP addresses that shouldn&#39;t be used in
     * the cluster.
     * 
     */
    private String noProxy;
    /**
     * @return The project of the cluster template.
     * 
     */
    private String projectId;
    /**
     * @return Indicates whether cluster template should be public.
     * 
     */
    private Boolean public_;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String region;
    /**
     * @return Indicates whether Docker registry is enabled in the
     * cluster.
     * 
     */
    private Boolean registryEnabled;
    /**
     * @return The server type for the cluster template.
     * 
     */
    private String serverType;
    /**
     * @return Indicates whether the TLS should be disabled in the cluster.
     * 
     */
    private Boolean tlsDisabled;
    /**
     * @return The time at which cluster template was updated.
     * 
     */
    private String updatedAt;
    /**
     * @return The user of the cluster template.
     * 
     */
    private String userId;
    /**
     * @return The name of the driver that is used for the volumes of the
     * cluster nodes.
     * 
     */
    private String volumeDriver;

    private GetClusterTemplateResult() {}
    /**
     * @return The API server port for the Container Orchestration
     * Engine for this cluster template.
     * 
     */
    public Integer apiserverPort() {
        return this.apiserverPort;
    }
    /**
     * @return The distro for the cluster (fedora-atomic, coreos, etc.).
     * 
     */
    public String clusterDistro() {
        return this.clusterDistro;
    }
    /**
     * @return The Container Orchestration Engine for this cluster template.
     * 
     */
    public String coe() {
        return this.coe;
    }
    /**
     * @return The time at which cluster template was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Address of the DNS nameserver that is used in nodes of the
     * cluster.
     * 
     */
    public String dnsNameserver() {
        return this.dnsNameserver;
    }
    /**
     * @return Docker storage driver. Changing this updates the
     * Docker storage driver of the existing cluster template.
     * 
     */
    public String dockerStorageDriver() {
        return this.dockerStorageDriver;
    }
    /**
     * @return The size (in GB) of the Docker volume.
     * 
     */
    public Integer dockerVolumeSize() {
        return this.dockerVolumeSize;
    }
    /**
     * @return The ID of the external network that will be used for
     * the cluster.
     * 
     */
    public String externalNetworkId() {
        return this.externalNetworkId;
    }
    /**
     * @return The fixed network that will be attached to the cluster.
     * 
     */
    public String fixedNetwork() {
        return this.fixedNetwork;
    }
    /**
     * @return =The fixed subnet that will be attached to the cluster.
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
    /**
     * @return Indicates whether created cluster should create IP
     * floating IP for every node or not.
     * 
     */
    public Boolean floatingIpEnabled() {
        return this.floatingIpEnabled;
    }
    /**
     * @return Indicates whether the ClusterTemplate is hidden or not.
     * 
     */
    public Boolean hidden() {
        return this.hidden;
    }
    /**
     * @return The address of a proxy for receiving all HTTP requests and
     * relay them.
     * 
     */
    public String httpProxy() {
        return this.httpProxy;
    }
    /**
     * @return The address of a proxy for receiving all HTTPS requests and
     * relay them.
     * 
     */
    public String httpsProxy() {
        return this.httpsProxy;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The reference to an image that is used for nodes of the cluster.
     * 
     */
    public String image() {
        return this.image;
    }
    /**
     * @return The insecure registry URL for the cluster template.
     * 
     */
    public String insecureRegistry() {
        return this.insecureRegistry;
    }
    /**
     * @return The name of the Compute service SSH keypair.
     * 
     */
    public String keypairId() {
        return this.keypairId;
    }
    /**
     * @return The list of key value pairs representing additional properties
     * of the cluster template.
     * 
     */
    public Map<String,Object> labels() {
        return this.labels;
    }
    /**
     * @return The flavor for the master nodes.
     * 
     */
    public String masterFlavor() {
        return this.masterFlavor;
    }
    /**
     * @return Indicates whether created cluster should has a
     * loadbalancer for master nodes or not.
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
     * @return The name of the driver for the container network.
     * 
     */
    public String networkDriver() {
        return this.networkDriver;
    }
    /**
     * @return A comma-separated list of IP addresses that shouldn&#39;t be used in
     * the cluster.
     * 
     */
    public String noProxy() {
        return this.noProxy;
    }
    /**
     * @return The project of the cluster template.
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Indicates whether cluster template should be public.
     * 
     */
    public Boolean public_() {
        return this.public_;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return Indicates whether Docker registry is enabled in the
     * cluster.
     * 
     */
    public Boolean registryEnabled() {
        return this.registryEnabled;
    }
    /**
     * @return The server type for the cluster template.
     * 
     */
    public String serverType() {
        return this.serverType;
    }
    /**
     * @return Indicates whether the TLS should be disabled in the cluster.
     * 
     */
    public Boolean tlsDisabled() {
        return this.tlsDisabled;
    }
    /**
     * @return The time at which cluster template was updated.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return The user of the cluster template.
     * 
     */
    public String userId() {
        return this.userId;
    }
    /**
     * @return The name of the driver that is used for the volumes of the
     * cluster nodes.
     * 
     */
    public String volumeDriver() {
        return this.volumeDriver;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClusterTemplateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer apiserverPort;
        private String clusterDistro;
        private String coe;
        private String createdAt;
        private String dnsNameserver;
        private String dockerStorageDriver;
        private Integer dockerVolumeSize;
        private String externalNetworkId;
        private String fixedNetwork;
        private String fixedSubnet;
        private String flavor;
        private Boolean floatingIpEnabled;
        private Boolean hidden;
        private String httpProxy;
        private String httpsProxy;
        private String id;
        private String image;
        private String insecureRegistry;
        private String keypairId;
        private Map<String,Object> labels;
        private String masterFlavor;
        private Boolean masterLbEnabled;
        private String name;
        private String networkDriver;
        private String noProxy;
        private String projectId;
        private Boolean public_;
        private String region;
        private Boolean registryEnabled;
        private String serverType;
        private Boolean tlsDisabled;
        private String updatedAt;
        private String userId;
        private String volumeDriver;
        public Builder() {}
        public Builder(GetClusterTemplateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.apiserverPort = defaults.apiserverPort;
    	      this.clusterDistro = defaults.clusterDistro;
    	      this.coe = defaults.coe;
    	      this.createdAt = defaults.createdAt;
    	      this.dnsNameserver = defaults.dnsNameserver;
    	      this.dockerStorageDriver = defaults.dockerStorageDriver;
    	      this.dockerVolumeSize = defaults.dockerVolumeSize;
    	      this.externalNetworkId = defaults.externalNetworkId;
    	      this.fixedNetwork = defaults.fixedNetwork;
    	      this.fixedSubnet = defaults.fixedSubnet;
    	      this.flavor = defaults.flavor;
    	      this.floatingIpEnabled = defaults.floatingIpEnabled;
    	      this.hidden = defaults.hidden;
    	      this.httpProxy = defaults.httpProxy;
    	      this.httpsProxy = defaults.httpsProxy;
    	      this.id = defaults.id;
    	      this.image = defaults.image;
    	      this.insecureRegistry = defaults.insecureRegistry;
    	      this.keypairId = defaults.keypairId;
    	      this.labels = defaults.labels;
    	      this.masterFlavor = defaults.masterFlavor;
    	      this.masterLbEnabled = defaults.masterLbEnabled;
    	      this.name = defaults.name;
    	      this.networkDriver = defaults.networkDriver;
    	      this.noProxy = defaults.noProxy;
    	      this.projectId = defaults.projectId;
    	      this.public_ = defaults.public_;
    	      this.region = defaults.region;
    	      this.registryEnabled = defaults.registryEnabled;
    	      this.serverType = defaults.serverType;
    	      this.tlsDisabled = defaults.tlsDisabled;
    	      this.updatedAt = defaults.updatedAt;
    	      this.userId = defaults.userId;
    	      this.volumeDriver = defaults.volumeDriver;
        }

        @CustomType.Setter
        public Builder apiserverPort(Integer apiserverPort) {
            this.apiserverPort = Objects.requireNonNull(apiserverPort);
            return this;
        }
        @CustomType.Setter
        public Builder clusterDistro(String clusterDistro) {
            this.clusterDistro = Objects.requireNonNull(clusterDistro);
            return this;
        }
        @CustomType.Setter
        public Builder coe(String coe) {
            this.coe = Objects.requireNonNull(coe);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder dnsNameserver(String dnsNameserver) {
            this.dnsNameserver = Objects.requireNonNull(dnsNameserver);
            return this;
        }
        @CustomType.Setter
        public Builder dockerStorageDriver(String dockerStorageDriver) {
            this.dockerStorageDriver = Objects.requireNonNull(dockerStorageDriver);
            return this;
        }
        @CustomType.Setter
        public Builder dockerVolumeSize(Integer dockerVolumeSize) {
            this.dockerVolumeSize = Objects.requireNonNull(dockerVolumeSize);
            return this;
        }
        @CustomType.Setter
        public Builder externalNetworkId(String externalNetworkId) {
            this.externalNetworkId = Objects.requireNonNull(externalNetworkId);
            return this;
        }
        @CustomType.Setter
        public Builder fixedNetwork(String fixedNetwork) {
            this.fixedNetwork = Objects.requireNonNull(fixedNetwork);
            return this;
        }
        @CustomType.Setter
        public Builder fixedSubnet(String fixedSubnet) {
            this.fixedSubnet = Objects.requireNonNull(fixedSubnet);
            return this;
        }
        @CustomType.Setter
        public Builder flavor(String flavor) {
            this.flavor = Objects.requireNonNull(flavor);
            return this;
        }
        @CustomType.Setter
        public Builder floatingIpEnabled(Boolean floatingIpEnabled) {
            this.floatingIpEnabled = Objects.requireNonNull(floatingIpEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder hidden(Boolean hidden) {
            this.hidden = Objects.requireNonNull(hidden);
            return this;
        }
        @CustomType.Setter
        public Builder httpProxy(String httpProxy) {
            this.httpProxy = Objects.requireNonNull(httpProxy);
            return this;
        }
        @CustomType.Setter
        public Builder httpsProxy(String httpsProxy) {
            this.httpsProxy = Objects.requireNonNull(httpsProxy);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder image(String image) {
            this.image = Objects.requireNonNull(image);
            return this;
        }
        @CustomType.Setter
        public Builder insecureRegistry(String insecureRegistry) {
            this.insecureRegistry = Objects.requireNonNull(insecureRegistry);
            return this;
        }
        @CustomType.Setter
        public Builder keypairId(String keypairId) {
            this.keypairId = Objects.requireNonNull(keypairId);
            return this;
        }
        @CustomType.Setter
        public Builder labels(Map<String,Object> labels) {
            this.labels = Objects.requireNonNull(labels);
            return this;
        }
        @CustomType.Setter
        public Builder masterFlavor(String masterFlavor) {
            this.masterFlavor = Objects.requireNonNull(masterFlavor);
            return this;
        }
        @CustomType.Setter
        public Builder masterLbEnabled(Boolean masterLbEnabled) {
            this.masterLbEnabled = Objects.requireNonNull(masterLbEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder networkDriver(String networkDriver) {
            this.networkDriver = Objects.requireNonNull(networkDriver);
            return this;
        }
        @CustomType.Setter
        public Builder noProxy(String noProxy) {
            this.noProxy = Objects.requireNonNull(noProxy);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter("public")
        public Builder public_(Boolean public_) {
            this.public_ = Objects.requireNonNull(public_);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder registryEnabled(Boolean registryEnabled) {
            this.registryEnabled = Objects.requireNonNull(registryEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder serverType(String serverType) {
            this.serverType = Objects.requireNonNull(serverType);
            return this;
        }
        @CustomType.Setter
        public Builder tlsDisabled(Boolean tlsDisabled) {
            this.tlsDisabled = Objects.requireNonNull(tlsDisabled);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            this.userId = Objects.requireNonNull(userId);
            return this;
        }
        @CustomType.Setter
        public Builder volumeDriver(String volumeDriver) {
            this.volumeDriver = Objects.requireNonNull(volumeDriver);
            return this;
        }
        public GetClusterTemplateResult build() {
            final var o = new GetClusterTemplateResult();
            o.apiserverPort = apiserverPort;
            o.clusterDistro = clusterDistro;
            o.coe = coe;
            o.createdAt = createdAt;
            o.dnsNameserver = dnsNameserver;
            o.dockerStorageDriver = dockerStorageDriver;
            o.dockerVolumeSize = dockerVolumeSize;
            o.externalNetworkId = externalNetworkId;
            o.fixedNetwork = fixedNetwork;
            o.fixedSubnet = fixedSubnet;
            o.flavor = flavor;
            o.floatingIpEnabled = floatingIpEnabled;
            o.hidden = hidden;
            o.httpProxy = httpProxy;
            o.httpsProxy = httpsProxy;
            o.id = id;
            o.image = image;
            o.insecureRegistry = insecureRegistry;
            o.keypairId = keypairId;
            o.labels = labels;
            o.masterFlavor = masterFlavor;
            o.masterLbEnabled = masterLbEnabled;
            o.name = name;
            o.networkDriver = networkDriver;
            o.noProxy = noProxy;
            o.projectId = projectId;
            o.public_ = public_;
            o.region = region;
            o.registryEnabled = registryEnabled;
            o.serverType = serverType;
            o.tlsDisabled = tlsDisabled;
            o.updatedAt = updatedAt;
            o.userId = userId;
            o.volumeDriver = volumeDriver;
            return o;
        }
    }
}
