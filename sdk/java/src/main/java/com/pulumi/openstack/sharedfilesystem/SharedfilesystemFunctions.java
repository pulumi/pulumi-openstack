// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.sharedfilesystem;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesPlainArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.GetShareArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkPlainArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.GetSharePlainArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotArgs;
import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotPlainArgs;
import com.pulumi.openstack.sharedfilesystem.outputs.GetAvailbilityZonesResult;
import com.pulumi.openstack.sharedfilesystem.outputs.GetShareNetworkResult;
import com.pulumi.openstack.sharedfilesystem.outputs.GetShareResult;
import com.pulumi.openstack.sharedfilesystem.outputs.GetSnapshotResult;
import java.util.concurrent.CompletableFuture;

public final class SharedfilesystemFunctions {
    /**
     * Use this data source to get a list of Shared File System availability zones
     * from OpenStack
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesArgs;
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
     *         final var zones = SharedfilesystemFunctions.getAvailbilityZones();
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAvailbilityZonesResult> getAvailbilityZones() {
        return getAvailbilityZones(GetAvailbilityZonesArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Shared File System availability zones
     * from OpenStack
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesArgs;
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
     *         final var zones = SharedfilesystemFunctions.getAvailbilityZones();
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAvailbilityZonesResult> getAvailbilityZonesPlain() {
        return getAvailbilityZonesPlain(GetAvailbilityZonesPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Shared File System availability zones
     * from OpenStack
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesArgs;
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
     *         final var zones = SharedfilesystemFunctions.getAvailbilityZones();
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAvailbilityZonesResult> getAvailbilityZones(GetAvailbilityZonesArgs args) {
        return getAvailbilityZones(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Shared File System availability zones
     * from OpenStack
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesArgs;
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
     *         final var zones = SharedfilesystemFunctions.getAvailbilityZones();
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAvailbilityZonesResult> getAvailbilityZonesPlain(GetAvailbilityZonesPlainArgs args) {
        return getAvailbilityZonesPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Shared File System availability zones
     * from OpenStack
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesArgs;
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
     *         final var zones = SharedfilesystemFunctions.getAvailbilityZones();
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAvailbilityZonesResult> getAvailbilityZones(GetAvailbilityZonesArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones", TypeShape.of(GetAvailbilityZonesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get a list of Shared File System availability zones
     * from OpenStack
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetAvailbilityZonesArgs;
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
     *         final var zones = SharedfilesystemFunctions.getAvailbilityZones();
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAvailbilityZonesResult> getAvailbilityZonesPlain(GetAvailbilityZonesPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:sharedfilesystem/getAvailbilityZones:getAvailbilityZones", TypeShape.of(GetAvailbilityZonesResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Shared File System share.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareArgs;
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
     *         final var share1 = SharedfilesystemFunctions.getShare(GetShareArgs.builder()
     *             .name("share_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetShareResult> getShare() {
        return getShare(GetShareArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareArgs;
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
     *         final var share1 = SharedfilesystemFunctions.getShare(GetShareArgs.builder()
     *             .name("share_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetShareResult> getSharePlain() {
        return getSharePlain(GetSharePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareArgs;
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
     *         final var share1 = SharedfilesystemFunctions.getShare(GetShareArgs.builder()
     *             .name("share_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetShareResult> getShare(GetShareArgs args) {
        return getShare(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareArgs;
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
     *         final var share1 = SharedfilesystemFunctions.getShare(GetShareArgs.builder()
     *             .name("share_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetShareResult> getSharePlain(GetSharePlainArgs args) {
        return getSharePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareArgs;
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
     *         final var share1 = SharedfilesystemFunctions.getShare(GetShareArgs.builder()
     *             .name("share_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetShareResult> getShare(GetShareArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:sharedfilesystem/getShare:getShare", TypeShape.of(GetShareResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Shared File System share.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareArgs;
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
     *         final var share1 = SharedfilesystemFunctions.getShare(GetShareArgs.builder()
     *             .name("share_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetShareResult> getSharePlain(GetSharePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:sharedfilesystem/getShare:getShare", TypeShape.of(GetShareResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Shared File System share network.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkArgs;
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
     *         final var sharenetwork1 = SharedfilesystemFunctions.getShareNetwork(GetShareNetworkArgs.builder()
     *             .name("sharenetwork_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetShareNetworkResult> getShareNetwork() {
        return getShareNetwork(GetShareNetworkArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share network.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkArgs;
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
     *         final var sharenetwork1 = SharedfilesystemFunctions.getShareNetwork(GetShareNetworkArgs.builder()
     *             .name("sharenetwork_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetShareNetworkResult> getShareNetworkPlain() {
        return getShareNetworkPlain(GetShareNetworkPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share network.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkArgs;
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
     *         final var sharenetwork1 = SharedfilesystemFunctions.getShareNetwork(GetShareNetworkArgs.builder()
     *             .name("sharenetwork_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetShareNetworkResult> getShareNetwork(GetShareNetworkArgs args) {
        return getShareNetwork(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share network.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkArgs;
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
     *         final var sharenetwork1 = SharedfilesystemFunctions.getShareNetwork(GetShareNetworkArgs.builder()
     *             .name("sharenetwork_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetShareNetworkResult> getShareNetworkPlain(GetShareNetworkPlainArgs args) {
        return getShareNetworkPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System share network.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkArgs;
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
     *         final var sharenetwork1 = SharedfilesystemFunctions.getShareNetwork(GetShareNetworkArgs.builder()
     *             .name("sharenetwork_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetShareNetworkResult> getShareNetwork(GetShareNetworkArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:sharedfilesystem/getShareNetwork:getShareNetwork", TypeShape.of(GetShareNetworkResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Shared File System share network.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetShareNetworkArgs;
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
     *         final var sharenetwork1 = SharedfilesystemFunctions.getShareNetwork(GetShareNetworkArgs.builder()
     *             .name("sharenetwork_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetShareNetworkResult> getShareNetworkPlain(GetShareNetworkPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:sharedfilesystem/getShareNetwork:getShareNetwork", TypeShape.of(GetShareNetworkResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Shared File System snapshot.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotArgs;
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
     *         final var snapshot1 = SharedfilesystemFunctions.getSnapshot(GetSnapshotArgs.builder()
     *             .name("snapshot_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSnapshotResult> getSnapshot() {
        return getSnapshot(GetSnapshotArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System snapshot.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotArgs;
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
     *         final var snapshot1 = SharedfilesystemFunctions.getSnapshot(GetSnapshotArgs.builder()
     *             .name("snapshot_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSnapshotResult> getSnapshotPlain() {
        return getSnapshotPlain(GetSnapshotPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System snapshot.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotArgs;
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
     *         final var snapshot1 = SharedfilesystemFunctions.getSnapshot(GetSnapshotArgs.builder()
     *             .name("snapshot_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSnapshotResult> getSnapshot(GetSnapshotArgs args) {
        return getSnapshot(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System snapshot.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotArgs;
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
     *         final var snapshot1 = SharedfilesystemFunctions.getSnapshot(GetSnapshotArgs.builder()
     *             .name("snapshot_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSnapshotResult> getSnapshotPlain(GetSnapshotPlainArgs args) {
        return getSnapshotPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Shared File System snapshot.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotArgs;
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
     *         final var snapshot1 = SharedfilesystemFunctions.getSnapshot(GetSnapshotArgs.builder()
     *             .name("snapshot_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSnapshotResult> getSnapshot(GetSnapshotArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:sharedfilesystem/getSnapshot:getSnapshot", TypeShape.of(GetSnapshotResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Shared File System snapshot.
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
     * import com.pulumi.openstack.sharedfilesystem.SharedfilesystemFunctions;
     * import com.pulumi.openstack.sharedfilesystem.inputs.GetSnapshotArgs;
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
     *         final var snapshot1 = SharedfilesystemFunctions.getSnapshot(GetSnapshotArgs.builder()
     *             .name("snapshot_1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSnapshotResult> getSnapshotPlain(GetSnapshotPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:sharedfilesystem/getSnapshot:getSnapshot", TypeShape.of(GetSnapshotResult.class), args, Utilities.withVersion(options));
    }
}
