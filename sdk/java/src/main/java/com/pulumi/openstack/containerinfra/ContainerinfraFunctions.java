// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.containerinfra;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.containerinfra.inputs.GetClusterArgs;
import com.pulumi.openstack.containerinfra.inputs.GetClusterPlainArgs;
import com.pulumi.openstack.containerinfra.inputs.GetClusterTemplateArgs;
import com.pulumi.openstack.containerinfra.inputs.GetClusterTemplatePlainArgs;
import com.pulumi.openstack.containerinfra.inputs.GetNodeGroupArgs;
import com.pulumi.openstack.containerinfra.inputs.GetNodeGroupPlainArgs;
import com.pulumi.openstack.containerinfra.outputs.GetClusterResult;
import com.pulumi.openstack.containerinfra.outputs.GetClusterTemplateResult;
import com.pulumi.openstack.containerinfra.outputs.GetNodeGroupResult;
import java.util.concurrent.CompletableFuture;

public final class ContainerinfraFunctions {
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterArgs;
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
     *         final var cluster1 = ContainerinfraFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;cluster_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClusterResult> getCluster(GetClusterArgs args) {
        return getCluster(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterArgs;
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
     *         final var cluster1 = ContainerinfraFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;cluster_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClusterResult> getClusterPlain(GetClusterPlainArgs args) {
        return getClusterPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterArgs;
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
     *         final var cluster1 = ContainerinfraFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;cluster_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClusterResult> getCluster(GetClusterArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:containerinfra/getCluster:getCluster", TypeShape.of(GetClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterArgs;
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
     *         final var cluster1 = ContainerinfraFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;cluster_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClusterResult> getClusterPlain(GetClusterPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:containerinfra/getCluster:getCluster", TypeShape.of(GetClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster
     * template.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterTemplateArgs;
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
     *         final var clustertemplate1 = ContainerinfraFunctions.getClusterTemplate(GetClusterTemplateArgs.builder()
     *             .name(&#34;clustertemplate_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClusterTemplateResult> getClusterTemplate(GetClusterTemplateArgs args) {
        return getClusterTemplate(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster
     * template.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterTemplateArgs;
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
     *         final var clustertemplate1 = ContainerinfraFunctions.getClusterTemplate(GetClusterTemplateArgs.builder()
     *             .name(&#34;clustertemplate_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClusterTemplateResult> getClusterTemplatePlain(GetClusterTemplatePlainArgs args) {
        return getClusterTemplatePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster
     * template.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterTemplateArgs;
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
     *         final var clustertemplate1 = ContainerinfraFunctions.getClusterTemplate(GetClusterTemplateArgs.builder()
     *             .name(&#34;clustertemplate_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetClusterTemplateResult> getClusterTemplate(GetClusterTemplateArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:containerinfra/getClusterTemplate:getClusterTemplate", TypeShape.of(GetClusterTemplateResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available OpenStack Magnum cluster
     * template.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetClusterTemplateArgs;
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
     *         final var clustertemplate1 = ContainerinfraFunctions.getClusterTemplate(GetClusterTemplateArgs.builder()
     *             .name(&#34;clustertemplate_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetClusterTemplateResult> getClusterTemplatePlain(GetClusterTemplatePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:containerinfra/getClusterTemplate:getClusterTemplate", TypeShape.of(GetClusterTemplateResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information of an available OpenStack Magnum node group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetNodeGroupArgs;
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
     *         final var nodegroup1 = ContainerinfraFunctions.getNodeGroup(GetNodeGroupArgs.builder()
     *             .clusterId(&#34;cluster_1&#34;)
     *             .name(&#34;nodegroup_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetNodeGroupResult> getNodeGroup(GetNodeGroupArgs args) {
        return getNodeGroup(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack Magnum node group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetNodeGroupArgs;
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
     *         final var nodegroup1 = ContainerinfraFunctions.getNodeGroup(GetNodeGroupArgs.builder()
     *             .clusterId(&#34;cluster_1&#34;)
     *             .name(&#34;nodegroup_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetNodeGroupResult> getNodeGroupPlain(GetNodeGroupPlainArgs args) {
        return getNodeGroupPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack Magnum node group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetNodeGroupArgs;
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
     *         final var nodegroup1 = ContainerinfraFunctions.getNodeGroup(GetNodeGroupArgs.builder()
     *             .clusterId(&#34;cluster_1&#34;)
     *             .name(&#34;nodegroup_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetNodeGroupResult> getNodeGroup(GetNodeGroupArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:containerinfra/getNodeGroup:getNodeGroup", TypeShape.of(GetNodeGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information of an available OpenStack Magnum node group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.containerinfra.ContainerinfraFunctions;
     * import com.pulumi.openstack.containerinfra.inputs.GetNodeGroupArgs;
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
     *         final var nodegroup1 = ContainerinfraFunctions.getNodeGroup(GetNodeGroupArgs.builder()
     *             .clusterId(&#34;cluster_1&#34;)
     *             .name(&#34;nodegroup_1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetNodeGroupResult> getNodeGroupPlain(GetNodeGroupPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:containerinfra/getNodeGroup:getNodeGroup", TypeShape.of(GetNodeGroupResult.class), args, Utilities.withVersion(options));
    }
}
