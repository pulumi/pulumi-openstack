// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.keymanager.inputs.ContainerV1AclArgs;
import com.pulumi.openstack.keymanager.inputs.ContainerV1SecretRefArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerV1Args extends com.pulumi.resources.ResourceArgs {

    public static final ContainerV1Args Empty = new ContainerV1Args();

    /**
     * Allows to control an access to a container. Currently only
     * the `read` operation is supported. If not specified, the container is
     * accessible project wide. The `read` structure is described below.
     * 
     */
    @Import(name="acl")
    private @Nullable Output<ContainerV1AclArgs> acl;

    /**
     * @return Allows to control an access to a container. Currently only
     * the `read` operation is supported. If not specified, the container is
     * accessible project wide. The `read` structure is described below.
     * 
     */
    public Optional<Output<ContainerV1AclArgs>> acl() {
        return Optional.ofNullable(this.acl);
    }

    /**
     * The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a container. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 container.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V1 KeyManager client.
     * A KeyManager client is needed to create a container. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * V1 container.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * A set of dictionaries containing references to secrets. The structure is described
     * below.
     * 
     */
    @Import(name="secretRefs")
    private @Nullable Output<List<ContainerV1SecretRefArgs>> secretRefs;

    /**
     * @return A set of dictionaries containing references to secrets. The structure is described
     * below.
     * 
     */
    public Optional<Output<List<ContainerV1SecretRefArgs>>> secretRefs() {
        return Optional.ofNullable(this.secretRefs);
    }

    /**
     * Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private ContainerV1Args() {}

    private ContainerV1Args(ContainerV1Args $) {
        this.acl = $.acl;
        this.name = $.name;
        this.region = $.region;
        this.secretRefs = $.secretRefs;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerV1Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerV1Args $;

        public Builder() {
            $ = new ContainerV1Args();
        }

        public Builder(ContainerV1Args defaults) {
            $ = new ContainerV1Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl Allows to control an access to a container. Currently only
         * the `read` operation is supported. If not specified, the container is
         * accessible project wide. The `read` structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<ContainerV1AclArgs> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl Allows to control an access to a container. Currently only
         * the `read` operation is supported. If not specified, the container is
         * accessible project wide. The `read` structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder acl(ContainerV1AclArgs acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param name The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the secret reference. The reference names must correspond the container type, more details are available [here](https://docs.openstack.org/barbican/stein/api/reference/containers.html).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region The region in which to obtain the V1 KeyManager client.
         * A KeyManager client is needed to create a container. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * V1 container.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V1 KeyManager client.
         * A KeyManager client is needed to create a container. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * V1 container.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param secretRefs A set of dictionaries containing references to secrets. The structure is described
         * below.
         * 
         * @return builder
         * 
         */
        public Builder secretRefs(@Nullable Output<List<ContainerV1SecretRefArgs>> secretRefs) {
            $.secretRefs = secretRefs;
            return this;
        }

        /**
         * @param secretRefs A set of dictionaries containing references to secrets. The structure is described
         * below.
         * 
         * @return builder
         * 
         */
        public Builder secretRefs(List<ContainerV1SecretRefArgs> secretRefs) {
            return secretRefs(Output.of(secretRefs));
        }

        /**
         * @param secretRefs A set of dictionaries containing references to secrets. The structure is described
         * below.
         * 
         * @return builder
         * 
         */
        public Builder secretRefs(ContainerV1SecretRefArgs... secretRefs) {
            return secretRefs(List.of(secretRefs));
        }

        /**
         * @param type Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Used to indicate the type of container. Must be one of `generic`, `rsa` or `certificate`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ContainerV1Args build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}