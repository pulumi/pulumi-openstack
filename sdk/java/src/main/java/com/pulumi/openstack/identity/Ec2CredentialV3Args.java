// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Ec2CredentialV3Args extends com.pulumi.resources.ResourceArgs {

    public static final Ec2CredentialV3Args Empty = new Ec2CredentialV3Args();

    /**
     * The ID of the project the EC2 credential is created
     * for and that authentication requests using this EC2 credential will
     * be scoped to.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The ID of the project the EC2 credential is created
     * for and that authentication requests using this EC2 credential will
     * be scoped to.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new EC2 credential.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V3 Keystone client.
     * If omitted, the `region` argument of the provider is used. Changing this
     * creates a new EC2 credential.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The ID of the user the EC2 credential is created for.
     * 
     */
    @Import(name="userId")
    private @Nullable Output<String> userId;

    /**
     * @return The ID of the user the EC2 credential is created for.
     * 
     */
    public Optional<Output<String>> userId() {
        return Optional.ofNullable(this.userId);
    }

    private Ec2CredentialV3Args() {}

    private Ec2CredentialV3Args(Ec2CredentialV3Args $) {
        this.projectId = $.projectId;
        this.region = $.region;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Ec2CredentialV3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Ec2CredentialV3Args $;

        public Builder() {
            $ = new Ec2CredentialV3Args();
        }

        public Builder(Ec2CredentialV3Args defaults) {
            $ = new Ec2CredentialV3Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId The ID of the project the EC2 credential is created
         * for and that authentication requests using this EC2 credential will
         * be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The ID of the project the EC2 credential is created
         * for and that authentication requests using this EC2 credential will
         * be scoped to.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param region The region in which to obtain the V3 Keystone client.
         * If omitted, the `region` argument of the provider is used. Changing this
         * creates a new EC2 credential.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V3 Keystone client.
         * If omitted, the `region` argument of the provider is used. Changing this
         * creates a new EC2 credential.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param userId The ID of the user the EC2 credential is created for.
         * 
         * @return builder
         * 
         */
        public Builder userId(@Nullable Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of the user the EC2 credential is created for.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public Ec2CredentialV3Args build() {
            return $;
        }
    }

}