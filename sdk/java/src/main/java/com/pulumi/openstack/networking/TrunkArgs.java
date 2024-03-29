// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.networking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.openstack.networking.inputs.TrunkSubPortArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TrunkArgs extends com.pulumi.resources.ResourceArgs {

    public static final TrunkArgs Empty = new TrunkArgs();

    /**
     * Administrative up/down status for the trunk
     * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
     * `admin_state_up` of an existing trunk.
     * 
     */
    @Import(name="adminStateUp")
    private @Nullable Output<Boolean> adminStateUp;

    /**
     * @return Administrative up/down status for the trunk
     * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
     * `admin_state_up` of an existing trunk.
     * 
     */
    public Optional<Output<Boolean>> adminStateUp() {
        return Optional.ofNullable(this.adminStateUp);
    }

    /**
     * Human-readable description of the trunk. Changing this
     * updates the name of the existing trunk.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description of the trunk. Changing this
     * updates the name of the existing trunk.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A unique name for the trunk. Changing this
     * updates the `name` of an existing trunk.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A unique name for the trunk. Changing this
     * updates the `name` of an existing trunk.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the port to be used as the parent port of the
     * trunk. This is the port that should be used as the compute instance network
     * port. Changing this creates a new trunk.
     * 
     */
    @Import(name="portId", required=true)
    private Output<String> portId;

    /**
     * @return The ID of the port to be used as the parent port of the
     * trunk. This is the port that should be used as the compute instance network
     * port. Changing this creates a new trunk.
     * 
     */
    public Output<String> portId() {
        return this.portId;
    }

    /**
     * The region in which to obtain the V2 networking client.
     * A networking client is needed to create a trunk. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * trunk.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 networking client.
     * A networking client is needed to create a trunk. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * trunk.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The set of ports that will be made subports of the trunk.
     * The structure of each subport is described below.
     * 
     */
    @Import(name="subPorts")
    private @Nullable Output<List<TrunkSubPortArgs>> subPorts;

    /**
     * @return The set of ports that will be made subports of the trunk.
     * The structure of each subport is described below.
     * 
     */
    public Optional<Output<List<TrunkSubPortArgs>>> subPorts() {
        return Optional.ofNullable(this.subPorts);
    }

    /**
     * A set of string tags for the port.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of string tags for the port.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The owner of the Trunk. Required if admin wants
     * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return The owner of the Trunk. Required if admin wants
     * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private TrunkArgs() {}

    private TrunkArgs(TrunkArgs $) {
        this.adminStateUp = $.adminStateUp;
        this.description = $.description;
        this.name = $.name;
        this.portId = $.portId;
        this.region = $.region;
        this.subPorts = $.subPorts;
        this.tags = $.tags;
        this.tenantId = $.tenantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrunkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrunkArgs $;

        public Builder() {
            $ = new TrunkArgs();
        }

        public Builder(TrunkArgs defaults) {
            $ = new TrunkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminStateUp Administrative up/down status for the trunk
         * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
         * `admin_state_up` of an existing trunk.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(@Nullable Output<Boolean> adminStateUp) {
            $.adminStateUp = adminStateUp;
            return this;
        }

        /**
         * @param adminStateUp Administrative up/down status for the trunk
         * (must be &#34;true&#34; or &#34;false&#34; if provided). Changing this updates the
         * `admin_state_up` of an existing trunk.
         * 
         * @return builder
         * 
         */
        public Builder adminStateUp(Boolean adminStateUp) {
            return adminStateUp(Output.of(adminStateUp));
        }

        /**
         * @param description Human-readable description of the trunk. Changing this
         * updates the name of the existing trunk.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description of the trunk. Changing this
         * updates the name of the existing trunk.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name A unique name for the trunk. Changing this
         * updates the `name` of an existing trunk.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A unique name for the trunk. Changing this
         * updates the `name` of an existing trunk.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param portId The ID of the port to be used as the parent port of the
         * trunk. This is the port that should be used as the compute instance network
         * port. Changing this creates a new trunk.
         * 
         * @return builder
         * 
         */
        public Builder portId(Output<String> portId) {
            $.portId = portId;
            return this;
        }

        /**
         * @param portId The ID of the port to be used as the parent port of the
         * trunk. This is the port that should be used as the compute instance network
         * port. Changing this creates a new trunk.
         * 
         * @return builder
         * 
         */
        public Builder portId(String portId) {
            return portId(Output.of(portId));
        }

        /**
         * @param region The region in which to obtain the V2 networking client.
         * A networking client is needed to create a trunk. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * trunk.
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
         * A networking client is needed to create a trunk. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * trunk.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param subPorts The set of ports that will be made subports of the trunk.
         * The structure of each subport is described below.
         * 
         * @return builder
         * 
         */
        public Builder subPorts(@Nullable Output<List<TrunkSubPortArgs>> subPorts) {
            $.subPorts = subPorts;
            return this;
        }

        /**
         * @param subPorts The set of ports that will be made subports of the trunk.
         * The structure of each subport is described below.
         * 
         * @return builder
         * 
         */
        public Builder subPorts(List<TrunkSubPortArgs> subPorts) {
            return subPorts(Output.of(subPorts));
        }

        /**
         * @param subPorts The set of ports that will be made subports of the trunk.
         * The structure of each subport is described below.
         * 
         * @return builder
         * 
         */
        public Builder subPorts(TrunkSubPortArgs... subPorts) {
            return subPorts(List.of(subPorts));
        }

        /**
         * @param tags A set of string tags for the port.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of string tags for the port.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of string tags for the port.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param tenantId The owner of the Trunk. Required if admin wants
         * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId The owner of the Trunk. Required if admin wants
         * to create a trunk on behalf of another tenant. Changing this creates a new trunk.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public TrunkArgs build() {
            if ($.portId == null) {
                throw new MissingRequiredPropertyException("TrunkArgs", "portId");
            }
            return $;
        }
    }

}
