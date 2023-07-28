// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.sharedfilesystem.ShareAccessArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.ShareAccessState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### NFS
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Subnet;
 * import com.pulumi.openstack.networking.SubnetArgs;
 * import com.pulumi.openstack.sharedfilesystem.ShareNetwork;
 * import com.pulumi.openstack.sharedfilesystem.ShareNetworkArgs;
 * import com.pulumi.openstack.sharedfilesystem.Share;
 * import com.pulumi.openstack.sharedfilesystem.ShareArgs;
 * import com.pulumi.openstack.sharedfilesystem.ShareAccess;
 * import com.pulumi.openstack.sharedfilesystem.ShareAccessArgs;
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
 *         var network1 = new Network(&#34;network1&#34;, NetworkArgs.builder()        
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
 *         var subnet1 = new Subnet(&#34;subnet1&#34;, SubnetArgs.builder()        
 *             .cidr(&#34;192.168.199.0/24&#34;)
 *             .ipVersion(4)
 *             .networkId(network1.id())
 *             .build());
 * 
 *         var sharenetwork1 = new ShareNetwork(&#34;sharenetwork1&#34;, ShareNetworkArgs.builder()        
 *             .description(&#34;test share network with security services&#34;)
 *             .neutronNetId(network1.id())
 *             .neutronSubnetId(subnet1.id())
 *             .build());
 * 
 *         var share1 = new Share(&#34;share1&#34;, ShareArgs.builder()        
 *             .description(&#34;test share description&#34;)
 *             .shareProto(&#34;NFS&#34;)
 *             .size(1)
 *             .shareNetworkId(sharenetwork1.id())
 *             .build());
 * 
 *         var shareAccess1 = new ShareAccess(&#34;shareAccess1&#34;, ShareAccessArgs.builder()        
 *             .shareId(share1.id())
 *             .accessType(&#34;ip&#34;)
 *             .accessTo(&#34;192.168.199.10&#34;)
 *             .accessLevel(&#34;rw&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### CIFS
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.openstack.networking.Network;
 * import com.pulumi.openstack.networking.NetworkArgs;
 * import com.pulumi.openstack.networking.Subnet;
 * import com.pulumi.openstack.networking.SubnetArgs;
 * import com.pulumi.openstack.sharedfilesystem.SecurityService;
 * import com.pulumi.openstack.sharedfilesystem.SecurityServiceArgs;
 * import com.pulumi.openstack.sharedfilesystem.ShareNetwork;
 * import com.pulumi.openstack.sharedfilesystem.ShareNetworkArgs;
 * import com.pulumi.openstack.sharedfilesystem.Share;
 * import com.pulumi.openstack.sharedfilesystem.ShareArgs;
 * import com.pulumi.openstack.sharedfilesystem.ShareAccess;
 * import com.pulumi.openstack.sharedfilesystem.ShareAccessArgs;
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
 *         var network1 = new Network(&#34;network1&#34;, NetworkArgs.builder()        
 *             .adminStateUp(&#34;true&#34;)
 *             .build());
 * 
 *         var subnet1 = new Subnet(&#34;subnet1&#34;, SubnetArgs.builder()        
 *             .cidr(&#34;192.168.199.0/24&#34;)
 *             .ipVersion(4)
 *             .networkId(network1.id())
 *             .build());
 * 
 *         var securityservice1 = new SecurityService(&#34;securityservice1&#34;, SecurityServiceArgs.builder()        
 *             .description(&#34;created by terraform&#34;)
 *             .type(&#34;active_directory&#34;)
 *             .server(&#34;192.168.199.10&#34;)
 *             .dnsIp(&#34;192.168.199.10&#34;)
 *             .domain(&#34;example.com&#34;)
 *             .ou(&#34;CN=Computers,DC=example,DC=com&#34;)
 *             .user(&#34;joinDomainUser&#34;)
 *             .password(&#34;s8cret&#34;)
 *             .build());
 * 
 *         var sharenetwork1 = new ShareNetwork(&#34;sharenetwork1&#34;, ShareNetworkArgs.builder()        
 *             .description(&#34;share the secure love&#34;)
 *             .neutronNetId(network1.id())
 *             .neutronSubnetId(subnet1.id())
 *             .securityServiceIds(securityservice1.id())
 *             .build());
 * 
 *         var share1 = new Share(&#34;share1&#34;, ShareArgs.builder()        
 *             .shareProto(&#34;CIFS&#34;)
 *             .size(1)
 *             .shareNetworkId(sharenetwork1.id())
 *             .build());
 * 
 *         var shareAccess1 = new ShareAccess(&#34;shareAccess1&#34;, ShareAccessArgs.builder()        
 *             .shareId(share1.id())
 *             .accessType(&#34;user&#34;)
 *             .accessTo(&#34;windows&#34;)
 *             .accessLevel(&#34;ro&#34;)
 *             .build());
 * 
 *         var shareAccess2 = new ShareAccess(&#34;shareAccess2&#34;, ShareAccessArgs.builder()        
 *             .shareId(share1.id())
 *             .accessType(&#34;user&#34;)
 *             .accessTo(&#34;linux&#34;)
 *             .accessLevel(&#34;rw&#34;)
 *             .build());
 * 
 *         ctx.export(&#34;exportLocations&#34;, share1.exportLocations());
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported by specifying the ID of the share and the ID of the share access, separated by a slash, e.g.
 * 
 * ```sh
 *  $ pulumi import openstack:sharedfilesystem/shareAccess:ShareAccess share_access_1 share_id/share_access_id
 * ```
 * 
 */
@ResourceType(type="openstack:sharedfilesystem/shareAccess:ShareAccess")
public class ShareAccess extends com.pulumi.resources.CustomResource {
    /**
     * The access credential of the entity granted access.
     * 
     */
    @Export(name="accessKey", type=String.class, parameters={})
    private Output<String> accessKey;

    /**
     * @return The access credential of the entity granted access.
     * 
     */
    public Output<String> accessKey() {
        return this.accessKey;
    }
    /**
     * The access level to the share. Can either be `rw` or `ro`.
     * 
     */
    @Export(name="accessLevel", type=String.class, parameters={})
    private Output<String> accessLevel;

    /**
     * @return The access level to the share. Can either be `rw` or `ro`.
     * 
     */
    public Output<String> accessLevel() {
        return this.accessLevel;
    }
    /**
     * The value that defines the access. Can either be an IP
     * address or a username verified by configured Security Service of the Share Network.
     * 
     */
    @Export(name="accessTo", type=String.class, parameters={})
    private Output<String> accessTo;

    /**
     * @return The value that defines the access. Can either be an IP
     * address or a username verified by configured Security Service of the Share Network.
     * 
     */
    public Output<String> accessTo() {
        return this.accessTo;
    }
    /**
     * The access rule type. Can either be an ip, user,
     * cert, or cephx. cephx support requires an OpenStack environment that supports
     * Shared Filesystem microversion 2.13 (Mitaka) or later.
     * 
     */
    @Export(name="accessType", type=String.class, parameters={})
    private Output<String> accessType;

    /**
     * @return The access rule type. Can either be an ip, user,
     * cert, or cephx. cephx support requires an OpenStack environment that supports
     * Shared Filesystem microversion 2.13 (Mitaka) or later.
     * 
     */
    public Output<String> accessType() {
        return this.accessType;
    }
    /**
     * The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share access. Changing this
     * creates a new share access.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The region in which to obtain the V2 Shared File System client.
     * A Shared File System client is needed to create a share access. Changing this
     * creates a new share access.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The UUID of the share to which you are granted access.
     * 
     */
    @Export(name="shareId", type=String.class, parameters={})
    private Output<String> shareId;

    /**
     * @return The UUID of the share to which you are granted access.
     * 
     */
    public Output<String> shareId() {
        return this.shareId;
    }
    /**
     * The share access state.
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return The share access state.
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ShareAccess(String name) {
        this(name, ShareAccessArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ShareAccess(String name, ShareAccessArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ShareAccess(String name, ShareAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:sharedfilesystem/shareAccess:ShareAccess", name, args == null ? ShareAccessArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ShareAccess(String name, Output<String> id, @Nullable ShareAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("openstack:sharedfilesystem/shareAccess:ShareAccess", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "accessKey"
            ))
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
    public static ShareAccess get(String name, Output<String> id, @Nullable ShareAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ShareAccess(name, id, state, options);
    }
}
