// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFlavorArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFlavorArgs Empty = new GetFlavorArgs();

    /**
     * The exact amount of disk (in gigabytes).
     * 
     */
    @Import(name="disk")
    private @Nullable Output<Integer> disk;

    /**
     * @return The exact amount of disk (in gigabytes).
     * 
     */
    public Optional<Output<Integer>> disk() {
        return Optional.ofNullable(this.disk);
    }

    /**
     * The ID of the flavor. Conflicts with the `name`,
     * `min_ram` and `min_disk`
     * 
     */
    @Import(name="flavorId")
    private @Nullable Output<String> flavorId;

    /**
     * @return The ID of the flavor. Conflicts with the `name`,
     * `min_ram` and `min_disk`
     * 
     */
    public Optional<Output<String>> flavorId() {
        return Optional.ofNullable(this.flavorId);
    }

    /**
     * The flavor visibility.
     * 
     */
    @Import(name="isPublic")
    private @Nullable Output<Boolean> isPublic;

    /**
     * @return The flavor visibility.
     * 
     */
    public Optional<Output<Boolean>> isPublic() {
        return Optional.ofNullable(this.isPublic);
    }

    /**
     * The minimum amount of disk (in gigabytes). Conflicts
     * with the `flavor_id`.
     * 
     */
    @Import(name="minDisk")
    private @Nullable Output<Integer> minDisk;

    /**
     * @return The minimum amount of disk (in gigabytes). Conflicts
     * with the `flavor_id`.
     * 
     */
    public Optional<Output<Integer>> minDisk() {
        return Optional.ofNullable(this.minDisk);
    }

    /**
     * The minimum amount of RAM (in megabytes). Conflicts
     * with the `flavor_id`.
     * 
     */
    @Import(name="minRam")
    private @Nullable Output<Integer> minRam;

    /**
     * @return The minimum amount of RAM (in megabytes). Conflicts
     * with the `flavor_id`.
     * 
     */
    public Optional<Output<Integer>> minRam() {
        return Optional.ofNullable(this.minRam);
    }

    /**
     * The name of the flavor. Conflicts with the `flavor_id`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the flavor. Conflicts with the `flavor_id`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The exact amount of RAM (in megabytes).
     * 
     */
    @Import(name="ram")
    private @Nullable Output<Integer> ram;

    /**
     * @return The exact amount of RAM (in megabytes).
     * 
     */
    public Optional<Output<Integer>> ram() {
        return Optional.ofNullable(this.ram);
    }

    /**
     * The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region in which to obtain the V2 Compute client.
     * If omitted, the `region` argument of the provider is used.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The `rx_tx_factor` of the flavor.
     * 
     */
    @Import(name="rxTxFactor")
    private @Nullable Output<Double> rxTxFactor;

    /**
     * @return The `rx_tx_factor` of the flavor.
     * 
     */
    public Optional<Output<Double>> rxTxFactor() {
        return Optional.ofNullable(this.rxTxFactor);
    }

    /**
     * The amount of swap (in gigabytes).
     * 
     */
    @Import(name="swap")
    private @Nullable Output<Integer> swap;

    /**
     * @return The amount of swap (in gigabytes).
     * 
     */
    public Optional<Output<Integer>> swap() {
        return Optional.ofNullable(this.swap);
    }

    /**
     * The amount of VCPUs.
     * 
     */
    @Import(name="vcpus")
    private @Nullable Output<Integer> vcpus;

    /**
     * @return The amount of VCPUs.
     * 
     */
    public Optional<Output<Integer>> vcpus() {
        return Optional.ofNullable(this.vcpus);
    }

    private GetFlavorArgs() {}

    private GetFlavorArgs(GetFlavorArgs $) {
        this.disk = $.disk;
        this.flavorId = $.flavorId;
        this.isPublic = $.isPublic;
        this.minDisk = $.minDisk;
        this.minRam = $.minRam;
        this.name = $.name;
        this.ram = $.ram;
        this.region = $.region;
        this.rxTxFactor = $.rxTxFactor;
        this.swap = $.swap;
        this.vcpus = $.vcpus;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFlavorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFlavorArgs $;

        public Builder() {
            $ = new GetFlavorArgs();
        }

        public Builder(GetFlavorArgs defaults) {
            $ = new GetFlavorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disk The exact amount of disk (in gigabytes).
         * 
         * @return builder
         * 
         */
        public Builder disk(@Nullable Output<Integer> disk) {
            $.disk = disk;
            return this;
        }

        /**
         * @param disk The exact amount of disk (in gigabytes).
         * 
         * @return builder
         * 
         */
        public Builder disk(Integer disk) {
            return disk(Output.of(disk));
        }

        /**
         * @param flavorId The ID of the flavor. Conflicts with the `name`,
         * `min_ram` and `min_disk`
         * 
         * @return builder
         * 
         */
        public Builder flavorId(@Nullable Output<String> flavorId) {
            $.flavorId = flavorId;
            return this;
        }

        /**
         * @param flavorId The ID of the flavor. Conflicts with the `name`,
         * `min_ram` and `min_disk`
         * 
         * @return builder
         * 
         */
        public Builder flavorId(String flavorId) {
            return flavorId(Output.of(flavorId));
        }

        /**
         * @param isPublic The flavor visibility.
         * 
         * @return builder
         * 
         */
        public Builder isPublic(@Nullable Output<Boolean> isPublic) {
            $.isPublic = isPublic;
            return this;
        }

        /**
         * @param isPublic The flavor visibility.
         * 
         * @return builder
         * 
         */
        public Builder isPublic(Boolean isPublic) {
            return isPublic(Output.of(isPublic));
        }

        /**
         * @param minDisk The minimum amount of disk (in gigabytes). Conflicts
         * with the `flavor_id`.
         * 
         * @return builder
         * 
         */
        public Builder minDisk(@Nullable Output<Integer> minDisk) {
            $.minDisk = minDisk;
            return this;
        }

        /**
         * @param minDisk The minimum amount of disk (in gigabytes). Conflicts
         * with the `flavor_id`.
         * 
         * @return builder
         * 
         */
        public Builder minDisk(Integer minDisk) {
            return minDisk(Output.of(minDisk));
        }

        /**
         * @param minRam The minimum amount of RAM (in megabytes). Conflicts
         * with the `flavor_id`.
         * 
         * @return builder
         * 
         */
        public Builder minRam(@Nullable Output<Integer> minRam) {
            $.minRam = minRam;
            return this;
        }

        /**
         * @param minRam The minimum amount of RAM (in megabytes). Conflicts
         * with the `flavor_id`.
         * 
         * @return builder
         * 
         */
        public Builder minRam(Integer minRam) {
            return minRam(Output.of(minRam));
        }

        /**
         * @param name The name of the flavor. Conflicts with the `flavor_id`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the flavor. Conflicts with the `flavor_id`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param ram The exact amount of RAM (in megabytes).
         * 
         * @return builder
         * 
         */
        public Builder ram(@Nullable Output<Integer> ram) {
            $.ram = ram;
            return this;
        }

        /**
         * @param ram The exact amount of RAM (in megabytes).
         * 
         * @return builder
         * 
         */
        public Builder ram(Integer ram) {
            return ram(Output.of(ram));
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region in which to obtain the V2 Compute client.
         * If omitted, the `region` argument of the provider is used.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param rxTxFactor The `rx_tx_factor` of the flavor.
         * 
         * @return builder
         * 
         */
        public Builder rxTxFactor(@Nullable Output<Double> rxTxFactor) {
            $.rxTxFactor = rxTxFactor;
            return this;
        }

        /**
         * @param rxTxFactor The `rx_tx_factor` of the flavor.
         * 
         * @return builder
         * 
         */
        public Builder rxTxFactor(Double rxTxFactor) {
            return rxTxFactor(Output.of(rxTxFactor));
        }

        /**
         * @param swap The amount of swap (in gigabytes).
         * 
         * @return builder
         * 
         */
        public Builder swap(@Nullable Output<Integer> swap) {
            $.swap = swap;
            return this;
        }

        /**
         * @param swap The amount of swap (in gigabytes).
         * 
         * @return builder
         * 
         */
        public Builder swap(Integer swap) {
            return swap(Output.of(swap));
        }

        /**
         * @param vcpus The amount of VCPUs.
         * 
         * @return builder
         * 
         */
        public Builder vcpus(@Nullable Output<Integer> vcpus) {
            $.vcpus = vcpus;
            return this;
        }

        /**
         * @param vcpus The amount of VCPUs.
         * 
         * @return builder
         * 
         */
        public Builder vcpus(Integer vcpus) {
            return vcpus(Output.of(vcpus));
        }

        public GetFlavorArgs build() {
            return $;
        }
    }

}
