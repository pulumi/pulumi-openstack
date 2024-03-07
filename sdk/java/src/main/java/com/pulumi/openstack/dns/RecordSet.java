// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.dns;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.dns.RecordSetArgs;
import com.pulumi.openstack.dns.inputs.RecordSetState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a DNS record set in the OpenStack DNS Service.
 * 
 * ## Example Usage
 * 
 * ### Automatically detect the correct network
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.dns.Zone;
 * import com.pulumi.openstack.dns.ZoneArgs;
 * import com.pulumi.openstack.dns.RecordSet;
 * import com.pulumi.openstack.dns.RecordSetArgs;
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
 *         var exampleZone = new Zone(&#34;exampleZone&#34;, ZoneArgs.builder()        
 *             .email(&#34;email2@example.com&#34;)
 *             .description(&#34;a zone&#34;)
 *             .ttl(6000)
 *             .type(&#34;PRIMARY&#34;)
 *             .build());
 * 
 *         var rsExampleCom = new RecordSet(&#34;rsExampleCom&#34;, RecordSetArgs.builder()        
 *             .zoneId(exampleZone.id())
 *             .description(&#34;An example record set&#34;)
 *             .ttl(3000)
 *             .type(&#34;A&#34;)
 *             .records(&#34;10.0.0.1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource can be imported by specifying the zone ID and recordset ID,
 * separated by a forward slash.
 * 
 * ```sh
 * $ pulumi import openstack:dns/recordSet:RecordSet recordset_1 zone_id/recordset_id
 * ```
 * 
 */
@ResourceType(type="openstack:dns/recordSet:RecordSet")
public class RecordSet extends com.pulumi.resources.CustomResource {
    /**
     * A description of the  record set.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the  record set.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Disable wait for recordset to reach ACTIVE
     * status. This argumen is disabled by default. If it is set to true, the recordset
     * will be considered as created/updated/deleted if OpenStack request returned success.
     * 
     */
    @Export(name="disableStatusCheck", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableStatusCheck;

    /**
     * @return Disable wait for recordset to reach ACTIVE
     * status. This argumen is disabled by default. If it is set to true, the recordset
     * will be considered as created/updated/deleted if OpenStack request returned success.
     * 
     */
    public Output<Optional<Boolean>> disableStatusCheck() {
        return Codegen.optional(this.disableStatusCheck);
    }
    /**
     * The name of the record set. Note the `.` at the end of the name.
     * Changing this creates a new DNS  record set.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the record set. Note the `.` at the end of the name.
     * Changing this creates a new DNS  record set.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project DNS zone is created
     * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
     * user role in target project)
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return The ID of the project DNS zone is created
     * for, sets `X-Auth-Sudo-Tenant-ID` header (requires an assigned
     * user role in target project)
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * An array of DNS records.
     * 
     */
    @Export(name="records", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> records;

    /**
     * @return An array of DNS records.
     * 
     */
    public Output<List<String>> records() {
        return this.records;
    }
    /**
     * The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS  record set.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 DNS client.
     * If omitted, the `region` argument of the provider is used.
     * Changing this creates a new DNS  record set.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The time to live (TTL) of the record set.
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output<Integer> ttl;

    /**
     * @return The time to live (TTL) of the record set.
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }
    /**
     * The type of record set. Examples: &#34;A&#34;, &#34;MX&#34;.
     * Changing this creates a new DNS  record set.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of record set. Examples: &#34;A&#34;, &#34;MX&#34;.
     * Changing this creates a new DNS  record set.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Map of additional options. Changing this creates a
     * new record set.
     * 
     */
    @Export(name="valueSpecs", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> valueSpecs;

    /**
     * @return Map of additional options. Changing this creates a
     * new record set.
     * 
     */
    public Output<Optional<Map<String,Object>>> valueSpecs() {
        return Codegen.optional(this.valueSpecs);
    }
    /**
     * The ID of the zone in which to create the record set.
     * Changing this creates a new DNS  record set.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The ID of the zone in which to create the record set.
     * Changing this creates a new DNS  record set.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RecordSet(String name) {
        this(name, RecordSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RecordSet(String name, RecordSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RecordSet(String name, RecordSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:dns/recordSet:RecordSet", name, args == null ? RecordSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RecordSet(String name, Output<String> id, @Nullable RecordSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:dns/recordSet:RecordSet", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static RecordSet get(String name, Output<String> id, @Nullable RecordSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RecordSet(name, id, state, options);
    }
}
