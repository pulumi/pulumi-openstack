// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaSetV2Args extends com.pulumi.resources.ResourceArgs {

    public static final QuotaSetV2Args Empty = new QuotaSetV2Args();

    /**
     * Quota value for cores.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="cores")
    private @Nullable Output<Integer> cores;

    /**
     * @return Quota value for cores.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> cores() {
        return Optional.ofNullable(this.cores);
    }

    /**
     * Quota value for fixed IPs.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="fixedIps")
    private @Nullable Output<Integer> fixedIps;

    /**
     * @return Quota value for fixed IPs.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> fixedIps() {
        return Optional.ofNullable(this.fixedIps);
    }

    /**
     * Quota value for floating IPs.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="floatingIps")
    private @Nullable Output<Integer> floatingIps;

    /**
     * @return Quota value for floating IPs.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> floatingIps() {
        return Optional.ofNullable(this.floatingIps);
    }

    /**
     * Quota value for content bytes
     * of injected files. Changing this updates the existing quotaset.
     * 
     */
    @Import(name="injectedFileContentBytes")
    private @Nullable Output<Integer> injectedFileContentBytes;

    /**
     * @return Quota value for content bytes
     * of injected files. Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> injectedFileContentBytes() {
        return Optional.ofNullable(this.injectedFileContentBytes);
    }

    /**
     * Quota value for path bytes of
     * injected files. Changing this updates the existing quotaset.
     * 
     */
    @Import(name="injectedFilePathBytes")
    private @Nullable Output<Integer> injectedFilePathBytes;

    /**
     * @return Quota value for path bytes of
     * injected files. Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> injectedFilePathBytes() {
        return Optional.ofNullable(this.injectedFilePathBytes);
    }

    /**
     * Quota value for injected files.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="injectedFiles")
    private @Nullable Output<Integer> injectedFiles;

    /**
     * @return Quota value for injected files.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> injectedFiles() {
        return Optional.ofNullable(this.injectedFiles);
    }

    /**
     * Quota value for instances.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="instances")
    private @Nullable Output<Integer> instances;

    /**
     * @return Quota value for instances.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> instances() {
        return Optional.ofNullable(this.instances);
    }

    /**
     * Quota value for key pairs.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="keyPairs")
    private @Nullable Output<Integer> keyPairs;

    /**
     * @return Quota value for key pairs.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> keyPairs() {
        return Optional.ofNullable(this.keyPairs);
    }

    /**
     * Quota value for metadata items.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="metadataItems")
    private @Nullable Output<Integer> metadataItems;

    /**
     * @return Quota value for metadata items.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> metadataItems() {
        return Optional.ofNullable(this.metadataItems);
    }

    /**
     * ID of the project to manage quotas.
     * Changing this creates a new quotaset.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return ID of the project to manage quotas.
     * Changing this creates a new quotaset.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * Quota value for RAM.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="ram")
    private @Nullable Output<Integer> ram;

    /**
     * @return Quota value for RAM.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> ram() {
        return Optional.ofNullable(this.ram);
    }

    /**
     * The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the volume. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates a new quotaset.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Quota value for security group rules.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="securityGroupRules")
    private @Nullable Output<Integer> securityGroupRules;

    /**
     * @return Quota value for security group rules.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> securityGroupRules() {
        return Optional.ofNullable(this.securityGroupRules);
    }

    /**
     * Quota value for security groups.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="securityGroups")
    private @Nullable Output<Integer> securityGroups;

    /**
     * @return Quota value for security groups.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> securityGroups() {
        return Optional.ofNullable(this.securityGroups);
    }

    /**
     * Quota value for server groups members.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="serverGroupMembers")
    private @Nullable Output<Integer> serverGroupMembers;

    /**
     * @return Quota value for server groups members.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> serverGroupMembers() {
        return Optional.ofNullable(this.serverGroupMembers);
    }

    /**
     * Quota value for server groups.
     * Changing this updates the existing quotaset.
     * 
     */
    @Import(name="serverGroups")
    private @Nullable Output<Integer> serverGroups;

    /**
     * @return Quota value for server groups.
     * Changing this updates the existing quotaset.
     * 
     */
    public Optional<Output<Integer>> serverGroups() {
        return Optional.ofNullable(this.serverGroups);
    }

    private QuotaSetV2Args() {}

    private QuotaSetV2Args(QuotaSetV2Args $) {
        this.cores = $.cores;
        this.fixedIps = $.fixedIps;
        this.floatingIps = $.floatingIps;
        this.injectedFileContentBytes = $.injectedFileContentBytes;
        this.injectedFilePathBytes = $.injectedFilePathBytes;
        this.injectedFiles = $.injectedFiles;
        this.instances = $.instances;
        this.keyPairs = $.keyPairs;
        this.metadataItems = $.metadataItems;
        this.projectId = $.projectId;
        this.ram = $.ram;
        this.region = $.region;
        this.securityGroupRules = $.securityGroupRules;
        this.securityGroups = $.securityGroups;
        this.serverGroupMembers = $.serverGroupMembers;
        this.serverGroups = $.serverGroups;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaSetV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaSetV2Args $;

        public Builder() {
            $ = new QuotaSetV2Args();
        }

        public Builder(QuotaSetV2Args defaults) {
            $ = new QuotaSetV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param cores Quota value for cores.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder cores(@Nullable Output<Integer> cores) {
            $.cores = cores;
            return this;
        }

        /**
         * @param cores Quota value for cores.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder cores(Integer cores) {
            return cores(Output.of(cores));
        }

        /**
         * @param fixedIps Quota value for fixed IPs.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder fixedIps(@Nullable Output<Integer> fixedIps) {
            $.fixedIps = fixedIps;
            return this;
        }

        /**
         * @param fixedIps Quota value for fixed IPs.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder fixedIps(Integer fixedIps) {
            return fixedIps(Output.of(fixedIps));
        }

        /**
         * @param floatingIps Quota value for floating IPs.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder floatingIps(@Nullable Output<Integer> floatingIps) {
            $.floatingIps = floatingIps;
            return this;
        }

        /**
         * @param floatingIps Quota value for floating IPs.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder floatingIps(Integer floatingIps) {
            return floatingIps(Output.of(floatingIps));
        }

        /**
         * @param injectedFileContentBytes Quota value for content bytes
         * of injected files. Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder injectedFileContentBytes(@Nullable Output<Integer> injectedFileContentBytes) {
            $.injectedFileContentBytes = injectedFileContentBytes;
            return this;
        }

        /**
         * @param injectedFileContentBytes Quota value for content bytes
         * of injected files. Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder injectedFileContentBytes(Integer injectedFileContentBytes) {
            return injectedFileContentBytes(Output.of(injectedFileContentBytes));
        }

        /**
         * @param injectedFilePathBytes Quota value for path bytes of
         * injected files. Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder injectedFilePathBytes(@Nullable Output<Integer> injectedFilePathBytes) {
            $.injectedFilePathBytes = injectedFilePathBytes;
            return this;
        }

        /**
         * @param injectedFilePathBytes Quota value for path bytes of
         * injected files. Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder injectedFilePathBytes(Integer injectedFilePathBytes) {
            return injectedFilePathBytes(Output.of(injectedFilePathBytes));
        }

        /**
         * @param injectedFiles Quota value for injected files.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder injectedFiles(@Nullable Output<Integer> injectedFiles) {
            $.injectedFiles = injectedFiles;
            return this;
        }

        /**
         * @param injectedFiles Quota value for injected files.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder injectedFiles(Integer injectedFiles) {
            return injectedFiles(Output.of(injectedFiles));
        }

        /**
         * @param instances Quota value for instances.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder instances(@Nullable Output<Integer> instances) {
            $.instances = instances;
            return this;
        }

        /**
         * @param instances Quota value for instances.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder instances(Integer instances) {
            return instances(Output.of(instances));
        }

        /**
         * @param keyPairs Quota value for key pairs.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder keyPairs(@Nullable Output<Integer> keyPairs) {
            $.keyPairs = keyPairs;
            return this;
        }

        /**
         * @param keyPairs Quota value for key pairs.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder keyPairs(Integer keyPairs) {
            return keyPairs(Output.of(keyPairs));
        }

        /**
         * @param metadataItems Quota value for metadata items.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder metadataItems(@Nullable Output<Integer> metadataItems) {
            $.metadataItems = metadataItems;
            return this;
        }

        /**
         * @param metadataItems Quota value for metadata items.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder metadataItems(Integer metadataItems) {
            return metadataItems(Output.of(metadataItems));
        }

        /**
         * @param projectId ID of the project to manage quotas.
         * Changing this creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId ID of the project to manage quotas.
         * Changing this creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param ram Quota value for RAM.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder ram(@Nullable Output<Integer> ram) {
            $.ram = ram;
            return this;
        }

        /**
         * @param ram Quota value for RAM.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder ram(Integer ram) {
            return ram(Output.of(ram));
        }

        /**
         * @param region The region in which to create the volume. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to create the volume. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates a new quotaset.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param securityGroupRules Quota value for security group rules.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupRules(@Nullable Output<Integer> securityGroupRules) {
            $.securityGroupRules = securityGroupRules;
            return this;
        }

        /**
         * @param securityGroupRules Quota value for security group rules.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupRules(Integer securityGroupRules) {
            return securityGroupRules(Output.of(securityGroupRules));
        }

        /**
         * @param securityGroups Quota value for security groups.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(@Nullable Output<Integer> securityGroups) {
            $.securityGroups = securityGroups;
            return this;
        }

        /**
         * @param securityGroups Quota value for security groups.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(Integer securityGroups) {
            return securityGroups(Output.of(securityGroups));
        }

        /**
         * @param serverGroupMembers Quota value for server groups members.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupMembers(@Nullable Output<Integer> serverGroupMembers) {
            $.serverGroupMembers = serverGroupMembers;
            return this;
        }

        /**
         * @param serverGroupMembers Quota value for server groups members.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder serverGroupMembers(Integer serverGroupMembers) {
            return serverGroupMembers(Output.of(serverGroupMembers));
        }

        /**
         * @param serverGroups Quota value for server groups.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder serverGroups(@Nullable Output<Integer> serverGroups) {
            $.serverGroups = serverGroups;
            return this;
        }

        /**
         * @param serverGroups Quota value for server groups.
         * Changing this updates the existing quotaset.
         * 
         * @return builder
         * 
         */
        public Builder serverGroups(Integer serverGroups) {
            return serverGroups(Output.of(serverGroups));
        }

        public QuotaSetV2Args build() {
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}