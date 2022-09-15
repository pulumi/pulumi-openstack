// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RbacPolicyV2State extends com.pulumi.resources.ResourceArgs {

    public static final RbacPolicyV2State Empty = new RbacPolicyV2State();

    /**
     * Action for the RBAC policy. Can either be
     * `access_as_external` or `access_as_shared`.
     * 
     */
    @Import(name="action")
    private @Nullable Output<String> action;

    /**
     * @return Action for the RBAC policy. Can either be
     * `access_as_external` or `access_as_shared`.
     * 
     */
    public Optional<Output<String>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * The ID of the `object_type` resource. An
     * `object_type` of `network` returns a network ID and an `object_type` of
     * `qos_policy` returns a QoS ID.
     * 
     */
    @Import(name="objectId")
    private @Nullable Output<String> objectId;

    /**
     * @return The ID of the `object_type` resource. An
     * `object_type` of `network` returns a network ID and an `object_type` of
     * `qos_policy` returns a QoS ID.
     * 
     */
    public Optional<Output<String>> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    /**
     * The type of the object that the RBAC policy
     * affects. Can be one of the following: `address_scope`, `address_group`,
     * `network`, `qos_policy`, `security_group` or `subnetpool`.
     * 
     */
    @Import(name="objectType")
    private @Nullable Output<String> objectType;

    /**
     * @return The type of the object that the RBAC policy
     * affects. Can be one of the following: `address_scope`, `address_group`,
     * `network`, `qos_policy`, `security_group` or `subnetpool`.
     * 
     */
    public Optional<Output<String>> objectType() {
        return Optional.ofNullable(this.objectType);
    }

    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to configure a routing entry on a subnet. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * routing entry.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The ID of the tenant to which the RBAC policy
     * will be enforced.
     * 
     */
    @Import(name="targetTenant")
    private @Nullable Output<String> targetTenant;

    /**
     * @return The ID of the tenant to which the RBAC policy
     * will be enforced.
     * 
     */
    public Optional<Output<String>> targetTenant() {
        return Optional.ofNullable(this.targetTenant);
    }

    private RbacPolicyV2State() {}

    private RbacPolicyV2State(RbacPolicyV2State $) {
        this.action = $.action;
        this.objectId = $.objectId;
        this.objectType = $.objectType;
        this.projectId = $.projectId;
        this.region = $.region;
        this.targetTenant = $.targetTenant;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RbacPolicyV2State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RbacPolicyV2State $;

        public Builder() {
            $ = new RbacPolicyV2State();
        }

        public Builder(RbacPolicyV2State defaults) {
            $ = new RbacPolicyV2State(Objects.requireNonNull(defaults));
        }

        /**
         * @param action Action for the RBAC policy. Can either be
         * `access_as_external` or `access_as_shared`.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action Action for the RBAC policy. Can either be
         * `access_as_external` or `access_as_shared`.
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param objectId The ID of the `object_type` resource. An
         * `object_type` of `network` returns a network ID and an `object_type` of
         * `qos_policy` returns a QoS ID.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable Output<String> objectId) {
            $.objectId = objectId;
            return this;
        }

        /**
         * @param objectId The ID of the `object_type` resource. An
         * `object_type` of `network` returns a network ID and an `object_type` of
         * `qos_policy` returns a QoS ID.
         * 
         * @return builder
         * 
         */
        public Builder objectId(String objectId) {
            return objectId(Output.of(objectId));
        }

        /**
         * @param objectType The type of the object that the RBAC policy
         * affects. Can be one of the following: `address_scope`, `address_group`,
         * `network`, `qos_policy`, `security_group` or `subnetpool`.
         * 
         * @return builder
         * 
         */
        public Builder objectType(@Nullable Output<String> objectType) {
            $.objectType = objectType;
            return this;
        }

        /**
         * @param objectType The type of the object that the RBAC policy
         * affects. Can be one of the following: `address_scope`, `address_group`,
         * `network`, `qos_policy`, `security_group` or `subnetpool`.
         * 
         * @return builder
         * 
         */
        public Builder objectType(String objectType) {
            return objectType(Output.of(objectType));
        }

        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V2 networking client.
         * A networking client is needed to configure a routing entry on a subnet. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * routing entry.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 networking client.
         * A networking client is needed to configure a routing entry on a subnet. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * routing entry.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param targetTenant The ID of the tenant to which the RBAC policy
         * will be enforced.
         * 
         * @return builder
         * 
         */
        public Builder targetTenant(@Nullable Output<String> targetTenant) {
            $.targetTenant = targetTenant;
            return this;
        }

        /**
         * @param targetTenant The ID of the tenant to which the RBAC policy
         * will be enforced.
         * 
         * @return builder
         * 
         */
        public Builder targetTenant(String targetTenant) {
            return targetTenant(Output.of(targetTenant));
        }

        public RbacPolicyV2State build() {
            return $;
        }
    }

}