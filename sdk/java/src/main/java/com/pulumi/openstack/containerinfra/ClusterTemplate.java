// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.containerinfra.ClusterTemplateArgs;
import com.pulumi.openstack.containerinfra.inputs.ClusterTemplateState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a V1 Magnum cluster template resource within OpenStack.
 * 
 * ## Example Usage
 * ### Create a Cluster template
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.containerinfra.ClusterTemplate;
 * import com.pulumi.openstack.containerinfra.ClusterTemplateArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var clustertemplate1 = new ClusterTemplate(&#34;clustertemplate1&#34;, ClusterTemplateArgs.builder()        
 *             .coe(&#34;kubernetes&#34;)
 *             .dnsNameserver(&#34;1.1.1.1&#34;)
 *             .dockerStorageDriver(&#34;devicemapper&#34;)
 *             .dockerVolumeSize(10)
 *             .flavor(&#34;m1.small&#34;)
 *             .floatingIpEnabled(false)
 *             .image(&#34;Fedora-Atomic-27&#34;)
 *             .labels(Map.ofEntries(
 *                 Map.entry(&#34;influx_grafana_dashboard_enabled&#34;, &#34;true&#34;),
 *                 Map.entry(&#34;kube_dashboard_enabled&#34;, &#34;true&#34;),
 *                 Map.entry(&#34;kube_tag&#34;, &#34;1.11.1&#34;),
 *                 Map.entry(&#34;prometheus_monitoring&#34;, &#34;true&#34;)
 *             ))
 *             .masterFlavor(&#34;m1.medium&#34;)
 *             .masterLbEnabled(true)
 *             .networkDriver(&#34;flannel&#34;)
 *             .serverType(&#34;vm&#34;)
 *             .volumeDriver(&#34;cinder&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Attributes reference
 * 
 * The following attributes are exported:
 * 
 * * `region` - See Argument Reference above.
 * * `name` - See Argument Reference above.
 * * `project_id` - See Argument Reference above.
 * * `created_at` - The time at which cluster template was created.
 * * `updated_at` - The time at which cluster template was created.
 * * `apiserver_port` - See Argument Reference above.
 * * `coe` - See Argument Reference above.
 * * `cluster_distro` - See Argument Reference above.
 * * `dns_nameserver` - See Argument Reference above.
 * * `docker_storage_driver` - See Argument Reference above.
 * * `docker_volume_size` - See Argument Reference above.
 * * `external_network_id` - See Argument Reference above.
 * * `fixed_network` - See Argument Reference above.
 * * `fixed_subnet` - See Argument Reference above.
 * * `flavor` - See Argument Reference above.
 * * `master_flavor` - See Argument Reference above.
 * * `floating_ip_enabled` - See Argument Reference above.
 * * `http_proxy` - See Argument Reference above.
 * * `https_proxy` - See Argument Reference above.
 * * `image` - See Argument Reference above.
 * * `insecure_registry` - See Argument Reference above.
 * * `keypair_id` - See Argument Reference above.
 * * `labels` - See Argument Reference above.
 * * `links` - A list containing associated cluster template links.
 * * `master_lb_enabled` - See Argument Reference above.
 * * `network_driver` - See Argument Reference above.
 * * `no_proxy` - See Argument Reference above.
 * * `public` - See Argument Reference above.
 * * `registry_enabled` - See Argument Reference above.
 * * `server_type` - See Argument Reference above.
 * * `tls_disabled` - See Argument Reference above.
 * * `volume_driver` - See Argument Reference above.
 * * `hidden` - See Argument Reference above.
 * 
 * ## Import
 * 
 * Cluster templates can be imported using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:containerinfra/clusterTemplate:ClusterTemplate clustertemplate_1 b9a45c5c-cd03-4958-82aa-b80bf93cb922
 * ```
 * 
 */
@ResourceType(type="openstack:containerinfra/clusterTemplate:ClusterTemplate")
public class ClusterTemplate extends com.pulumi.resources.CustomResource {
    /**
     * The API server port for the Container
     * Orchestration Engine for this cluster template. Changing this updates the
     * API server port of the existing cluster template.
     * 
     */
    @Export(name="apiserverPort", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> apiserverPort;

    /**
     * @return The API server port for the Container
     * Orchestration Engine for this cluster template. Changing this updates the
     * API server port of the existing cluster template.
     * 
     */
    public Output<Optional<Integer>> apiserverPort() {
        return Codegen.optional(this.apiserverPort);
    }
    /**
     * The distro for the cluster (fedora-atomic,
     * coreos, etc.). Changing this updates the cluster distro of the existing
     * cluster template.
     * 
     */
    @Export(name="clusterDistro", type=String.class, parameters={})
    private Output<String> clusterDistro;

    /**
     * @return The distro for the cluster (fedora-atomic,
     * coreos, etc.). Changing this updates the cluster distro of the existing
     * cluster template.
     * 
     */
    public Output<String> clusterDistro() {
        return this.clusterDistro;
    }
    /**
     * The Container Orchestration Engine for this cluster
     * template. Changing this updates the engine of the existing cluster
     * template.
     * 
     */
    @Export(name="coe", type=String.class, parameters={})
    private Output<String> coe;

    /**
     * @return The Container Orchestration Engine for this cluster
     * template. Changing this updates the engine of the existing cluster
     * template.
     * 
     */
    public Output<String> coe() {
        return this.coe;
    }
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Address of the DNS nameserver that is used in
     * nodes of the cluster. Changing this updates the DNS nameserver of the
     * existing cluster template.
     * 
     */
    @Export(name="dnsNameserver", type=String.class, parameters={})
    private Output</* @Nullable */ String> dnsNameserver;

    /**
     * @return Address of the DNS nameserver that is used in
     * nodes of the cluster. Changing this updates the DNS nameserver of the
     * existing cluster template.
     * 
     */
    public Output<Optional<String>> dnsNameserver() {
        return Codegen.optional(this.dnsNameserver);
    }
    /**
     * Docker storage driver. Changing this
     * updates the Docker storage driver of the existing cluster template.
     * 
     */
    @Export(name="dockerStorageDriver", type=String.class, parameters={})
    private Output</* @Nullable */ String> dockerStorageDriver;

    /**
     * @return Docker storage driver. Changing this
     * updates the Docker storage driver of the existing cluster template.
     * 
     */
    public Output<Optional<String>> dockerStorageDriver() {
        return Codegen.optional(this.dockerStorageDriver);
    }
    /**
     * The size (in GB) of the Docker volume.
     * Changing this updates the Docker volume size of the existing cluster
     * template.
     * 
     */
    @Export(name="dockerVolumeSize", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> dockerVolumeSize;

    /**
     * @return The size (in GB) of the Docker volume.
     * Changing this updates the Docker volume size of the existing cluster
     * template.
     * 
     */
    public Output<Optional<Integer>> dockerVolumeSize() {
        return Codegen.optional(this.dockerVolumeSize);
    }
    /**
     * The ID of the external network that will
     * be used for the cluster. Changing this updates the external network ID of
     * the existing cluster template.
     * 
     */
    @Export(name="externalNetworkId", type=String.class, parameters={})
    private Output</* @Nullable */ String> externalNetworkId;

    /**
     * @return The ID of the external network that will
     * be used for the cluster. Changing this updates the external network ID of
     * the existing cluster template.
     * 
     */
    public Output<Optional<String>> externalNetworkId() {
        return Codegen.optional(this.externalNetworkId);
    }
    /**
     * The fixed network that will be attached to the
     * cluster. Changing this updates the fixed network of the existing cluster
     * template.
     * 
     */
    @Export(name="fixedNetwork", type=String.class, parameters={})
    private Output</* @Nullable */ String> fixedNetwork;

    /**
     * @return The fixed network that will be attached to the
     * cluster. Changing this updates the fixed network of the existing cluster
     * template.
     * 
     */
    public Output<Optional<String>> fixedNetwork() {
        return Codegen.optional(this.fixedNetwork);
    }
    /**
     * The fixed subnet that will be attached to the
     * cluster. Changing this updates the fixed subnet of the existing cluster
     * template.
     * 
     */
    @Export(name="fixedSubnet", type=String.class, parameters={})
    private Output</* @Nullable */ String> fixedSubnet;

    /**
     * @return The fixed subnet that will be attached to the
     * cluster. Changing this updates the fixed subnet of the existing cluster
     * template.
     * 
     */
    public Output<Optional<String>> fixedSubnet() {
        return Codegen.optional(this.fixedSubnet);
    }
    /**
     * The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this updates the
     * flavor of the existing cluster template.
     * 
     */
    @Export(name="flavor", type=String.class, parameters={})
    private Output</* @Nullable */ String> flavor;

    /**
     * @return The flavor for the nodes of the cluster. Can be set via
     * the `OS_MAGNUM_FLAVOR` environment variable. Changing this updates the
     * flavor of the existing cluster template.
     * 
     */
    public Output<Optional<String>> flavor() {
        return Codegen.optional(this.flavor);
    }
    /**
     * Indicates whether created cluster should
     * create floating IP for every node or not. Changing this updates the
     * floating IP enabled attribute of the existing cluster template.
     * 
     */
    @Export(name="floatingIpEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> floatingIpEnabled;

    /**
     * @return Indicates whether created cluster should
     * create floating IP for every node or not. Changing this updates the
     * floating IP enabled attribute of the existing cluster template.
     * 
     */
    public Output<Optional<Boolean>> floatingIpEnabled() {
        return Codegen.optional(this.floatingIpEnabled);
    }
    /**
     * Indicates whether the ClusterTemplate is hidden or not.
     * Changing this updates the hidden attribute of the existing cluster
     * template.
     * 
     */
    @Export(name="hidden", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> hidden;

    /**
     * @return Indicates whether the ClusterTemplate is hidden or not.
     * Changing this updates the hidden attribute of the existing cluster
     * template.
     * 
     */
    public Output<Optional<Boolean>> hidden() {
        return Codegen.optional(this.hidden);
    }
    /**
     * The address of a proxy for receiving all HTTP
     * requests and relay them. Changing this updates the HTTP proxy address of
     * the existing cluster template.
     * 
     */
    @Export(name="httpProxy", type=String.class, parameters={})
    private Output</* @Nullable */ String> httpProxy;

    /**
     * @return The address of a proxy for receiving all HTTP
     * requests and relay them. Changing this updates the HTTP proxy address of
     * the existing cluster template.
     * 
     */
    public Output<Optional<String>> httpProxy() {
        return Codegen.optional(this.httpProxy);
    }
    /**
     * The address of a proxy for receiving all HTTPS
     * requests and relay them. Changing this updates the HTTPS proxy address of
     * the existing cluster template.
     * 
     */
    @Export(name="httpsProxy", type=String.class, parameters={})
    private Output</* @Nullable */ String> httpsProxy;

    /**
     * @return The address of a proxy for receiving all HTTPS
     * requests and relay them. Changing this updates the HTTPS proxy address of
     * the existing cluster template.
     * 
     */
    public Output<Optional<String>> httpsProxy() {
        return Codegen.optional(this.httpsProxy);
    }
    /**
     * The reference to an image that is used for nodes of the
     * cluster. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
     * Changing this updates the image attribute of the existing cluster template.
     * 
     */
    @Export(name="image", type=String.class, parameters={})
    private Output<String> image;

    /**
     * @return The reference to an image that is used for nodes of the
     * cluster. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
     * Changing this updates the image attribute of the existing cluster template.
     * 
     */
    public Output<String> image() {
        return this.image;
    }
    /**
     * The insecure registry URL for the cluster
     * template. Changing this updates the insecure registry attribute of the
     * existing cluster template.
     * 
     */
    @Export(name="insecureRegistry", type=String.class, parameters={})
    private Output</* @Nullable */ String> insecureRegistry;

    /**
     * @return The insecure registry URL for the cluster
     * template. Changing this updates the insecure registry attribute of the
     * existing cluster template.
     * 
     */
    public Output<Optional<String>> insecureRegistry() {
        return Codegen.optional(this.insecureRegistry);
    }
    /**
     * The name of the Compute service SSH keypair.
     * Changing this updates the keypair of the existing cluster template.
     * 
     */
    @Export(name="keypairId", type=String.class, parameters={})
    private Output</* @Nullable */ String> keypairId;

    /**
     * @return The name of the Compute service SSH keypair.
     * Changing this updates the keypair of the existing cluster template.
     * 
     */
    public Output<Optional<String>> keypairId() {
        return Codegen.optional(this.keypairId);
    }
    /**
     * The list of key value pairs representing additional
     * properties of the cluster template. Changing this updates the labels of the
     * existing cluster template.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> labels;

    /**
     * @return The list of key value pairs representing additional
     * properties of the cluster template. Changing this updates the labels of the
     * existing cluster template.
     * 
     */
    public Output<Optional<Map<String,Object>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this updates
     * the master flavor of the existing cluster template.
     * 
     */
    @Export(name="masterFlavor", type=String.class, parameters={})
    private Output</* @Nullable */ String> masterFlavor;

    /**
     * @return The flavor for the master nodes. Can be set via
     * the `OS_MAGNUM_MASTER_FLAVOR` environment variable. Changing this updates
     * the master flavor of the existing cluster template.
     * 
     */
    public Output<Optional<String>> masterFlavor() {
        return Codegen.optional(this.masterFlavor);
    }
    /**
     * Indicates whether created cluster should
     * has a loadbalancer for master nodes or not. Changing this updates the
     * attribute of the existing cluster template.
     * 
     */
    @Export(name="masterLbEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> masterLbEnabled;

    /**
     * @return Indicates whether created cluster should
     * has a loadbalancer for master nodes or not. Changing this updates the
     * attribute of the existing cluster template.
     * 
     */
    public Output<Optional<Boolean>> masterLbEnabled() {
        return Codegen.optional(this.masterLbEnabled);
    }
    /**
     * The name of the cluster template. Changing this updates
     * the name of the existing cluster template.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the cluster template. Changing this updates
     * the name of the existing cluster template.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the driver for the container
     * network. Changing this updates the network driver of the existing cluster
     * template.
     * 
     */
    @Export(name="networkDriver", type=String.class, parameters={})
    private Output<String> networkDriver;

    /**
     * @return The name of the driver for the container
     * network. Changing this updates the network driver of the existing cluster
     * template.
     * 
     */
    public Output<String> networkDriver() {
        return this.networkDriver;
    }
    /**
     * A comma-separated list of IP addresses that shouldn&#39;t
     * be used in the cluster. Changing this updates the no proxy list of the
     * existing cluster template.
     * 
     */
    @Export(name="noProxy", type=String.class, parameters={})
    private Output</* @Nullable */ String> noProxy;

    /**
     * @return A comma-separated list of IP addresses that shouldn&#39;t
     * be used in the cluster. Changing this updates the no proxy list of the
     * existing cluster template.
     * 
     */
    public Output<Optional<String>> noProxy() {
        return Codegen.optional(this.noProxy);
    }
    /**
     * The project of the cluster template. Required if
     * admin wants to create a cluster template in another project. Changing this
     * creates a new cluster template.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output<String> projectId;

    /**
     * @return The project of the cluster template. Required if
     * admin wants to create a cluster template in another project. Changing this
     * creates a new cluster template.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Indicates whether cluster template should be public.
     * Changing this updates the public attribute of the existing cluster
     * template.
     * 
     */
    @Export(name="public", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> public_;

    /**
     * @return Indicates whether cluster template should be public.
     * Changing this updates the public attribute of the existing cluster
     * template.
     * 
     */
    public Output<Optional<Boolean>> public_() {
        return Codegen.optional(this.public_);
    }
    /**
     * The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster template. If
     * omitted,the `region` argument of the provider is used. Changing this
     * creates a new cluster template.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V1 Container Infra
     * client. A Container Infra client is needed to create a cluster template. If
     * omitted,the `region` argument of the provider is used. Changing this
     * creates a new cluster template.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Indicates whether Docker registry is enabled
     * in the cluster. Changing this updates the registry enabled attribute of the
     * existing cluster template.
     * 
     */
    @Export(name="registryEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> registryEnabled;

    /**
     * @return Indicates whether Docker registry is enabled
     * in the cluster. Changing this updates the registry enabled attribute of the
     * existing cluster template.
     * 
     */
    public Output<Optional<Boolean>> registryEnabled() {
        return Codegen.optional(this.registryEnabled);
    }
    /**
     * The server type for the cluster template. Changing
     * this updates the server type of the existing cluster template.
     * 
     */
    @Export(name="serverType", type=String.class, parameters={})
    private Output<String> serverType;

    /**
     * @return The server type for the cluster template. Changing
     * this updates the server type of the existing cluster template.
     * 
     */
    public Output<String> serverType() {
        return this.serverType;
    }
    /**
     * Indicates whether the TLS should be disabled in
     * the cluster. Changing this updates the attribute of the existing cluster.
     * 
     */
    @Export(name="tlsDisabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> tlsDisabled;

    /**
     * @return Indicates whether the TLS should be disabled in
     * the cluster. Changing this updates the attribute of the existing cluster.
     * 
     */
    public Output<Optional<Boolean>> tlsDisabled() {
        return Codegen.optional(this.tlsDisabled);
    }
    @Export(name="updatedAt", type=String.class, parameters={})
    private Output<String> updatedAt;

    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * The user of the cluster template. Required if admin
     * wants to create a cluster template for another user. Changing this creates
     * a new cluster template.
     * 
     */
    @Export(name="userId", type=String.class, parameters={})
    private Output<String> userId;

    /**
     * @return The user of the cluster template. Required if admin
     * wants to create a cluster template for another user. Changing this creates
     * a new cluster template.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }
    /**
     * The name of the driver that is used for the
     * volumes of the cluster nodes. Changing this updates the volume driver of
     * the existing cluster template.
     * 
     */
    @Export(name="volumeDriver", type=String.class, parameters={})
    private Output</* @Nullable */ String> volumeDriver;

    /**
     * @return The name of the driver that is used for the
     * volumes of the cluster nodes. Changing this updates the volume driver of
     * the existing cluster template.
     * 
     */
    public Output<Optional<String>> volumeDriver() {
        return Codegen.optional(this.volumeDriver);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClusterTemplate(String name) {
        this(name, ClusterTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClusterTemplate(String name, ClusterTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClusterTemplate(String name, ClusterTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:containerinfra/clusterTemplate:ClusterTemplate", name, args == null ? ClusterTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClusterTemplate(String name, Output<String> id, @Nullable ClusterTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:containerinfra/clusterTemplate:ClusterTemplate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ClusterTemplate get(String name, Output<String> id, @Nullable ClusterTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClusterTemplate(name, id, state, options);
    }
}
