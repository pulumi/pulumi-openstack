// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.blockstorage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.blockstorage.QosAssociationV3Args;
import com.pulumi.openstack.blockstorage.inputs.QosAssociationV3State;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages a V3 block storage Qos Association resource within OpenStack.
 * 
 * &gt; **Note:** This usually requires admin privileges.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.blockstorage.QosV3;
 * import com.pulumi.openstack.blockstorage.QosV3Args;
 * import com.pulumi.openstack.blockstorage.VolumeTypeV3;
 * import com.pulumi.openstack.blockstorage.VolumeTypeV3Args;
 * import com.pulumi.openstack.blockstorage.QosAssociationV3;
 * import com.pulumi.openstack.blockstorage.QosAssociationV3Args;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var qos = new QosV3("qos", QosV3Args.builder()
 *             .name("%s")
 *             .consumer("front-end")
 *             .specs(Map.of("read_iops_sec", "20000"))
 *             .build());
 * 
 *         var volumeType = new VolumeTypeV3("volumeType", VolumeTypeV3Args.builder()
 *             .name("%s")
 *             .build());
 * 
 *         var qosAssociation = new QosAssociationV3("qosAssociation", QosAssociationV3Args.builder()
 *             .qosId(qos.id())
 *             .volumeTypeId(volumeType.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Qos association can be imported using the `qos_id/volume_type_id`, e.g.
 * 
 * ```sh
 * $ pulumi import openstack:blockstorage/qosAssociationV3:QosAssociationV3 qos_association 941793f0-0a34-4bc4-b72e-a6326ae58283/ea257959-eeb1-4c10-8d33-26f0409a755d
 * ```
 * 
 */
@ResourceType(type="openstack:blockstorage/qosAssociationV3:QosAssociationV3")
public class QosAssociationV3 extends com.pulumi.resources.CustomResource {
    /**
     * ID of the qos to associate. Changing this creates
     * a new qos association.
     * 
     */
    @Export(name="qosId", refs={String.class}, tree="[0]")
    private Output<String> qosId;

    /**
     * @return ID of the qos to associate. Changing this creates
     * a new qos association.
     * 
     */
    public Output<String> qosId() {
        return this.qosId;
    }
    /**
     * The region in which to create the qos association.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new qos association.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to create the qos association.
     * If omitted, the `region` argument of the provider is used. Changing
     * this creates a new qos association.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * ID of the volume_type to associate.
     * Changing this creates a new qos association.
     * 
     */
    @Export(name="volumeTypeId", refs={String.class}, tree="[0]")
    private Output<String> volumeTypeId;

    /**
     * @return ID of the volume_type to associate.
     * Changing this creates a new qos association.
     * 
     */
    public Output<String> volumeTypeId() {
        return this.volumeTypeId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QosAssociationV3(java.lang.String name) {
        this(name, QosAssociationV3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QosAssociationV3(java.lang.String name, QosAssociationV3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QosAssociationV3(java.lang.String name, QosAssociationV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/qosAssociationV3:QosAssociationV3", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private QosAssociationV3(java.lang.String name, Output<java.lang.String> id, @Nullable QosAssociationV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:blockstorage/qosAssociationV3:QosAssociationV3", name, state, makeResourceOptions(options, id), false);
    }

    private static QosAssociationV3Args makeArgs(QosAssociationV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? QosAssociationV3Args.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static QosAssociationV3 get(java.lang.String name, Output<java.lang.String> id, @Nullable QosAssociationV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QosAssociationV3(name, id, state, options);
    }
}
