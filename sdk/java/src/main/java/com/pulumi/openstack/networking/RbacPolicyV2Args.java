// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RbacPolicyV2Args extends com.pulumi.resources.ResourceArgs {

    public static final RbacPolicyV2Args Empty = new RbacPolicyV2Args();

    /**
     * Action for the RBAC policy. Can either be
     * `access_as_external` or `access_as_shared`.
     * 
     */
    @Import(name="action", required=true)
    private Output<String> action;

    /**
     * @return Action for the RBAC policy. Can either be
     * `access_as_external` or `access_as_shared`.
     * 
     */
    public Output<String> action() {
        return this.action;
    }

    /**
     * The ID of the `object_type` resource. An
     * `object_type` of `network` returns a network ID and an `object_type` of
     * `qos_policy` returns a QoS ID.
     * 
     */
    @Import(name="objectId", required=true)
    private Output<String> objectId;

    /**
     * @return The ID of the `object_type` resource. An
     * `object_type` of `network` returns a network ID and an `object_type` of
     * `qos_policy` returns a QoS ID.
     * 
     */
    public Output<String> objectId() {
        return this.objectId;
    }

    /**
     * The type of the object that the RBAC policy
     * affects. Can be one of the following: `address_scope`, `address_group`,
     * `network`, `qos_policy`, `security_group`, `subnetpool` or `bgpvpn`.
     * 
     */
    @Import(name="objectType", required=true)
    private Output<String> objectType;

    /**
     * @return The type of the object that the RBAC policy
     * affects. Can be one of the following: `address_scope`, `address_group`,
     * `network`, `qos_policy`, `security_group`, `subnetpool` or `bgpvpn`.
     * 
     */
    public Output<String> objectType() {
        return this.objectType;
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
    @Import(name="targetTenant", required=true)
    private Output<String> targetTenant;

    /**
     * @return The ID of the tenant to which the RBAC policy
     * will be enforced.
     * 
     */
    public Output<String> targetTenant() {
        return this.targetTenant;
    }

    private RbacPolicyV2Args() {}

    private RbacPolicyV2Args(RbacPolicyV2Args $) {
        this.action = $.action;
        this.objectId = $.objectId;
        this.objectType = $.objectType;
        this.region = $.region;
        this.targetTenant = $.targetTenant;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RbacPolicyV2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RbacPolicyV2Args $;

        public Builder() {
            $ = new RbacPolicyV2Args();
        }

        public Builder(RbacPolicyV2Args defaults) {
            $ = new RbacPolicyV2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param action Action for the RBAC policy. Can either be
         * `access_as_external` or `access_as_shared`.
         * 
         * @return builder
         * 
         */
        public Builder action(Output<String> action) {
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
        public Builder objectId(Output<String> objectId) {
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
         * `network`, `qos_policy`, `security_group`, `subnetpool` or `bgpvpn`.
         * 
         * @return builder
         * 
         */
        public Builder objectType(Output<String> objectType) {
            $.objectType = objectType;
            return this;
        }

        /**
         * @param objectType The type of the object that the RBAC policy
         * affects. Can be one of the following: `address_scope`, `address_group`,
         * `network`, `qos_policy`, `security_group`, `subnetpool` or `bgpvpn`.
         * 
         * @return builder
         * 
         */
        public Builder objectType(String objectType) {
            return objectType(Output.of(objectType));
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
        public Builder targetTenant(Output<String> targetTenant) {
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

        public RbacPolicyV2Args build() {
            if ($.action == null) {
                throw new MissingRequiredPropertyException("RbacPolicyV2Args", "action");
            }
            if ($.objectId == null) {
                throw new MissingRequiredPropertyException("RbacPolicyV2Args", "objectId");
            }
            if ($.objectType == null) {
                throw new MissingRequiredPropertyException("RbacPolicyV2Args", "objectType");
            }
            if ($.targetTenant == null) {
                throw new MissingRequiredPropertyException("RbacPolicyV2Args", "targetTenant");
            }
            return $;
        }
    }

}
