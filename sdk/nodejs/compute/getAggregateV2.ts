// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about host aggregates
 * by name.
 */
export function getAggregateV2(args: GetAggregateV2Args, opts?: pulumi.InvokeOptions): Promise<GetAggregateV2Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("openstack:compute/getAggregateV2:getAggregateV2", {
        "hosts": args.hosts,
        "metadata": args.metadata,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAggregateV2.
 */
export interface GetAggregateV2Args {
    /**
     * List of Hypervisors contained in the Host Aggregate
     */
    hosts?: string[];
    /**
     * Metadata of the Host Aggregate
     */
    metadata?: {[key: string]: string};
    /**
     * The name of the host aggregate
     */
    name: string;
}

/**
 * A collection of values returned by getAggregateV2.
 */
export interface GetAggregateV2Result {
    /**
     * List of Hypervisors contained in the Host Aggregate
     */
    readonly hosts: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Metadata of the Host Aggregate
     */
    readonly metadata: {[key: string]: string};
    /**
     * See Argument Reference above.
     */
    readonly name: string;
    /**
     * Availability zone of the Host Aggregate
     */
    readonly zone: string;
}
/**
 * Use this data source to get information about host aggregates
 * by name.
 */
export function getAggregateV2Output(args: GetAggregateV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAggregateV2Result> {
    return pulumi.output(args).apply((a: any) => getAggregateV2(a, opts))
}

/**
 * A collection of arguments for invoking getAggregateV2.
 */
export interface GetAggregateV2OutputArgs {
    /**
     * List of Hypervisors contained in the Host Aggregate
     */
    hosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Metadata of the Host Aggregate
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the host aggregate
     */
    name: pulumi.Input<string>;
}
