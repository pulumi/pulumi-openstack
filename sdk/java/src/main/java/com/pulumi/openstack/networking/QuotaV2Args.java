// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaV2Args extends com.pulumi.resources.ResourceArgs {

    public static final QuotaV2Args Empty = new QuotaV2Args();

    /**
     * Quota value for floating IPs. Changing this updates the
     * existing quota.
     * 
     */
    @Import(name="floatingip")
    private @Nullable Output<Integer> floatingip;

    /**
     * @return Quota value for floating IPs. Changing this updates the
     * existing quota.
     * 
     */
    public Optional<Output<Integer>> floatingip() {
        return Optional.ofNullable(this.floatingip);
    }

    /**
     * Quota value for networks. Changing this updates the
     * existing quota.
     * 
     */
    @Import(name="network")
    private @Nullable Output<Integer> network;

    /**
     * @return Quota value for networks. Changing this updates the
     * existing quota.
     * 
     */
    public Optional<Output<Integer>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * Quota value for ports. Changing this updates the
     * existing quota.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Quota value for ports. Changing this updates the
     * existing quota.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * ID of the project to manage quota. Changing this
     * creates new quota.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return ID of the project to manage quota. Changing this
     * creates new quota.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * Quota value for RBAC policies.
     * Changing this updates the existing quota.
     * 
     */
    @Import(name="rbacPolicy")
    private @Nullable Output<Integer> rbacPolicy;

    /**
     * @return Quota value for RBAC policies.
     * Changing this updates the existing quota.
     * 
     */
    public Optional<Output<Integer>> rbacPolicy() {
        return Optional.ofNullable(this.rbacPolicy);
    }

    /**
     * The region in which to create the quota. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates new quota.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to create the quota. If
     * omitted, the `region` argument of the provider is used. Changing this
     * creates new quota.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Quota value for routers. Changing this updates the
     * existing quota.
     * 
     */
    @Import(name="router")
    private @Nullable Output<Integer> router;

    /**
     * @return Quota value for routers. Changing this updates the
     * existing quota.
     * 
     */
    public Optional<Output<Integer>> router() {
        return Optional.ofNullable(this.router);
    }

    /**
     * Quota value for security groups. Changing
     * this updates the existing quota.
     * 
     */
    @Import(name="securityGroup")
    private @Nullable Output<Integer> securityGroup;

    /**
     * @return Quota value for security groups. Changing
     * this updates the existing quota.
     * 
     */
    public Optional<Output<Integer>> securityGroup() {
        return Optional.ofNullable(this.securityGroup);
    }

    /**
     * Quota value for security group rules.
     * Changing this updates the existing quota.
     * 
     */
    @Import(name="securityGroupRule")
    private @Nullable Output<Integer> securityGroupRule;

    /**
     * @return Quota value for security group rules.
     * Changing this updates the existing quota.
     * 
     */
    public Optional<Output<Integer>> securityGroupRule() {
        return Optional.ofNullable(this.securityGroupRule);
    }

    /**
     * Quota value for subnets. Changing
     * this updates the existing quota.
     * 
     */
    @Import(name="subnet")
    private @Nullable Output<Integer> subnet;

    /**
     * @return Quota value for subnets. Changing
     * this updates the existing quota.
     * 
     */
    public Optional<Output<Integer>> subnet() {
        return Optional.ofNullable(this.subnet);
    }

    /**
     * Quota value for subnetpools.
     * Changing this updates the existing quota.
     * 
     */
    @Import(name="subnetpool")
    private @Nullable Output<Integer> subnetpool;

    /**
     * @return Quota value for subnetpools.
     * Changing this updates the existing quota.
     * 
     */
    public Optional<Output<Integer>> subnetpool() {
        return Optional.ofNullable(this.subnetpool);
    }

    private QuotaV2Args() {}

    private QuotaV2Args(QuotaV2Args $) {
        this.floatingip = $.floatingip;
        this.network = $.network;
        this.port = $.port;
        this.projectId = $.projectId;
        this.rbacPolicy = $.rbacPolicy;
        this.region = $.region;
        this.router = $.router;
        this.securityGroup = $.securityGroup;
        this.securityGroupRule = $.securityGroupRule;
        this.subnet = $.subnet;
        this.subnetpool = $.subnetpool;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaV2Args $;

        public Builder() {
            $ = new QuotaV2Args();
        }

        public Builder(QuotaV2Args defaults) {
            $ = new QuotaV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param floatingip Quota value for floating IPs. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder floatingip(@Nullable Output<Integer> floatingip) {
            $.floatingip = floatingip;
            return this;
        }

        /**
         * @param floatingip Quota value for floating IPs. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder floatingip(Integer floatingip) {
            return floatingip(Output.of(floatingip));
        }

        /**
         * @param network Quota value for networks. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<Integer> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network Quota value for networks. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder network(Integer network) {
            return network(Output.of(network));
        }

        /**
         * @param port Quota value for ports. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Quota value for ports. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param projectId ID of the project to manage quota. Changing this
         * creates new quota.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId ID of the project to manage quota. Changing this
         * creates new quota.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param rbacPolicy Quota value for RBAC policies.
         * Changing this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder rbacPolicy(@Nullable Output<Integer> rbacPolicy) {
            $.rbacPolicy = rbacPolicy;
            return this;
        }

        /**
         * @param rbacPolicy Quota value for RBAC policies.
         * Changing this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder rbacPolicy(Integer rbacPolicy) {
            return rbacPolicy(Output.of(rbacPolicy));
        }

        /**
         * @param region The region in which to create the quota. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates new quota.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to create the quota. If
         * omitted, the `region` argument of the provider is used. Changing this
         * creates new quota.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param router Quota value for routers. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder router(@Nullable Output<Integer> router) {
            $.router = router;
            return this;
        }

        /**
         * @param router Quota value for routers. Changing this updates the
         * existing quota.
         * 
         * @return builder
         * 
         */
        public Builder router(Integer router) {
            return router(Output.of(router));
        }

        /**
         * @param securityGroup Quota value for security groups. Changing
         * this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder securityGroup(@Nullable Output<Integer> securityGroup) {
            $.securityGroup = securityGroup;
            return this;
        }

        /**
         * @param securityGroup Quota value for security groups. Changing
         * this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder securityGroup(Integer securityGroup) {
            return securityGroup(Output.of(securityGroup));
        }

        /**
         * @param securityGroupRule Quota value for security group rules.
         * Changing this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupRule(@Nullable Output<Integer> securityGroupRule) {
            $.securityGroupRule = securityGroupRule;
            return this;
        }

        /**
         * @param securityGroupRule Quota value for security group rules.
         * Changing this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupRule(Integer securityGroupRule) {
            return securityGroupRule(Output.of(securityGroupRule));
        }

        /**
         * @param subnet Quota value for subnets. Changing
         * this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder subnet(@Nullable Output<Integer> subnet) {
            $.subnet = subnet;
            return this;
        }

        /**
         * @param subnet Quota value for subnets. Changing
         * this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder subnet(Integer subnet) {
            return subnet(Output.of(subnet));
        }

        /**
         * @param subnetpool Quota value for subnetpools.
         * Changing this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder subnetpool(@Nullable Output<Integer> subnetpool) {
            $.subnetpool = subnetpool;
            return this;
        }

        /**
         * @param subnetpool Quota value for subnetpools.
         * Changing this updates the existing quota.
         * 
         * @return builder
         * 
         */
        public Builder subnetpool(Integer subnetpool) {
            return subnetpool(Output.of(subnetpool));
        }

        public QuotaV2Args build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("QuotaV2Args", "projectId");
            }
            return $;
        }
    }

}
