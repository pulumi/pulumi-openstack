// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.loadbalancer;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.openstack.loadbalancer.inputs.MembersMemberArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MembersArgs extends com.pulumi.resources.ResourceArgs {

    public static final MembersArgs Empty = new MembersArgs();

    /**
     * A set of dictionaries containing member parameters. The
     * structure is described below.
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<MembersMemberArgs>> members;

    /**
     * @return A set of dictionaries containing member parameters. The
     * structure is described below.
     * 
     */
    public Optional<Output<List<MembersMemberArgs>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The id of the pool that members will be assigned to.
     * Changing this creates a new members resource.
     * 
     */
    @Import(name="poolId", required=true)
    private Output<String> poolId;

    /**
     * @return The id of the pool that members will be assigned to.
     * Changing this creates a new members resource.
     * 
     */
    public Output<String> poolId() {
        return this.poolId;
    }

    /**
     * The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create pool members. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * members resource.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Networking client.
     * A Networking client is needed to create pool members. If omitted, the
     * `region` argument of the provider is used. Changing this creates a new
     * members resource.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private MembersArgs() {}

    private MembersArgs(MembersArgs $) {
        this.members = $.members;
        this.poolId = $.poolId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MembersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MembersArgs $;

        public Builder() {
            $ = new MembersArgs();
        }

        public Builder(MembersArgs defaults) {
            $ = new MembersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param members A set of dictionaries containing member parameters. The
         * structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<MembersMemberArgs>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members A set of dictionaries containing member parameters. The
         * structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder members(List<MembersMemberArgs> members) {
            return members(Output.of(members));
        }

        /**
         * @param members A set of dictionaries containing member parameters. The
         * structure is described below.
         * 
         * @return builder
         * 
         */
        public Builder members(MembersMemberArgs... members) {
            return members(List.of(members));
        }

        /**
         * @param poolId The id of the pool that members will be assigned to.
         * Changing this creates a new members resource.
         * 
         * @return builder
         * 
         */
        public Builder poolId(Output<String> poolId) {
            $.poolId = poolId;
            return this;
        }

        /**
         * @param poolId The id of the pool that members will be assigned to.
         * Changing this creates a new members resource.
         * 
         * @return builder
         * 
         */
        public Builder poolId(String poolId) {
            return poolId(Output.of(poolId));
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create pool members. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * members resource.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Networking client.
         * A Networking client is needed to create pool members. If omitted, the
         * `region` argument of the provider is used. Changing this creates a new
         * members resource.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public MembersArgs build() {
            $.poolId = Objects.requireNonNull($.poolId, "expected parameter 'poolId' to be non-null");
            return $;
        }
    }

}
