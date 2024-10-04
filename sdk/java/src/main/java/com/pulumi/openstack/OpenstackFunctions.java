// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.inputs.GetFwGroupV2Args;
import com.pulumi.openstack.inputs.GetFwGroupV2PlainArgs;
import com.pulumi.openstack.inputs.GetFwPolicyV2Args;
import com.pulumi.openstack.inputs.GetFwPolicyV2PlainArgs;
import com.pulumi.openstack.inputs.GetFwRuleV2Args;
import com.pulumi.openstack.inputs.GetFwRuleV2PlainArgs;
import com.pulumi.openstack.outputs.GetFwGroupV2Result;
import com.pulumi.openstack.outputs.GetFwPolicyV2Result;
import com.pulumi.openstack.outputs.GetFwRuleV2Result;
import java.util.concurrent.CompletableFuture;

public final class OpenstackFunctions {
    /**
     * Use this data source to get information of an available OpenStack firewall group v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetGroupV2Args;
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
     *         final var group = FirewallFunctions.getGroupV2(GetGroupV2Args.builder()
     *             .name("tf_test_group")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2
     * 
     */
    @Deprecated /* openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
    public static Output<GetFwGroupV2Result> getFwGroupV2() {
        return getFwGroupV2(GetFwGroupV2Args.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall group v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetGroupV2Args;
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
     *         final var group = FirewallFunctions.getGroupV2(GetGroupV2Args.builder()
     *             .name("tf_test_group")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2
     * 
     */
    @Deprecated /* openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
    public static CompletableFuture<GetFwGroupV2Result> getFwGroupV2Plain() {
        return getFwGroupV2Plain(GetFwGroupV2PlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall group v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetGroupV2Args;
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
     *         final var group = FirewallFunctions.getGroupV2(GetGroupV2Args.builder()
     *             .name("tf_test_group")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2
     * 
     */
    @Deprecated /* openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
    public static Output<GetFwGroupV2Result> getFwGroupV2(GetFwGroupV2Args args) {
        return getFwGroupV2(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall group v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetGroupV2Args;
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
     *         final var group = FirewallFunctions.getGroupV2(GetGroupV2Args.builder()
     *             .name("tf_test_group")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2
     * 
     */
    @Deprecated /* openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
    public static CompletableFuture<GetFwGroupV2Result> getFwGroupV2Plain(GetFwGroupV2PlainArgs args) {
        return getFwGroupV2Plain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall group v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetGroupV2Args;
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
     *         final var group = FirewallFunctions.getGroupV2(GetGroupV2Args.builder()
     *             .name("tf_test_group")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2
     * 
     */
    @Deprecated /* openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
    public static Output<GetFwGroupV2Result> getFwGroupV2(GetFwGroupV2Args args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:index/getFwGroupV2:getFwGroupV2", TypeShape.of(GetFwGroupV2Result.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information of an available OpenStack firewall group v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetGroupV2Args;
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
     *         final var group = FirewallFunctions.getGroupV2(GetGroupV2Args.builder()
     *             .name("tf_test_group")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2
     * 
     */
    @Deprecated /* openstack.index/getfwgroupv2.getFwGroupV2 has been deprecated in favor of openstack.firewall/getgroupv2.getGroupV2 */
    public static CompletableFuture<GetFwGroupV2Result> getFwGroupV2Plain(GetFwGroupV2PlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:index/getFwGroupV2:getFwGroupV2", TypeShape.of(GetFwGroupV2Result.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information of an available OpenStack firewall policy v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetPolicyV2Args;
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
     *         final var policy = FirewallFunctions.getPolicyV2(GetPolicyV2Args.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2
     * 
     */
    @Deprecated /* openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2 */
    public static Output<GetFwPolicyV2Result> getFwPolicyV2() {
        return getFwPolicyV2(GetFwPolicyV2Args.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall policy v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetPolicyV2Args;
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
     *         final var policy = FirewallFunctions.getPolicyV2(GetPolicyV2Args.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2
     * 
     */
    @Deprecated /* openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2 */
    public static CompletableFuture<GetFwPolicyV2Result> getFwPolicyV2Plain() {
        return getFwPolicyV2Plain(GetFwPolicyV2PlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall policy v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetPolicyV2Args;
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
     *         final var policy = FirewallFunctions.getPolicyV2(GetPolicyV2Args.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2
     * 
     */
    @Deprecated /* openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2 */
    public static Output<GetFwPolicyV2Result> getFwPolicyV2(GetFwPolicyV2Args args) {
        return getFwPolicyV2(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall policy v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetPolicyV2Args;
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
     *         final var policy = FirewallFunctions.getPolicyV2(GetPolicyV2Args.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2
     * 
     */
    @Deprecated /* openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2 */
    public static CompletableFuture<GetFwPolicyV2Result> getFwPolicyV2Plain(GetFwPolicyV2PlainArgs args) {
        return getFwPolicyV2Plain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall policy v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetPolicyV2Args;
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
     *         final var policy = FirewallFunctions.getPolicyV2(GetPolicyV2Args.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2
     * 
     */
    @Deprecated /* openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2 */
    public static Output<GetFwPolicyV2Result> getFwPolicyV2(GetFwPolicyV2Args args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:index/getFwPolicyV2:getFwPolicyV2", TypeShape.of(GetFwPolicyV2Result.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information of an available OpenStack firewall policy v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetPolicyV2Args;
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
     *         final var policy = FirewallFunctions.getPolicyV2(GetPolicyV2Args.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2
     * 
     */
    @Deprecated /* openstack.index/getfwpolicyv2.getFwPolicyV2 has been deprecated in favor of openstack.firewall/getpolicyv2.getPolicyV2 */
    public static CompletableFuture<GetFwPolicyV2Result> getFwPolicyV2Plain(GetFwPolicyV2PlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:index/getFwPolicyV2:getFwPolicyV2", TypeShape.of(GetFwPolicyV2Result.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information of an available OpenStack firewall rule v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetRuleV2Args;
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
     *         final var rule = FirewallFunctions.getRuleV2(GetRuleV2Args.builder()
     *             .name("tf_test_rule")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2
     * 
     */
    @Deprecated /* openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2 */
    public static Output<GetFwRuleV2Result> getFwRuleV2() {
        return getFwRuleV2(GetFwRuleV2Args.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall rule v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetRuleV2Args;
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
     *         final var rule = FirewallFunctions.getRuleV2(GetRuleV2Args.builder()
     *             .name("tf_test_rule")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2
     * 
     */
    @Deprecated /* openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2 */
    public static CompletableFuture<GetFwRuleV2Result> getFwRuleV2Plain() {
        return getFwRuleV2Plain(GetFwRuleV2PlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall rule v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetRuleV2Args;
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
     *         final var rule = FirewallFunctions.getRuleV2(GetRuleV2Args.builder()
     *             .name("tf_test_rule")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2
     * 
     */
    @Deprecated /* openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2 */
    public static Output<GetFwRuleV2Result> getFwRuleV2(GetFwRuleV2Args args) {
        return getFwRuleV2(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall rule v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetRuleV2Args;
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
     *         final var rule = FirewallFunctions.getRuleV2(GetRuleV2Args.builder()
     *             .name("tf_test_rule")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2
     * 
     */
    @Deprecated /* openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2 */
    public static CompletableFuture<GetFwRuleV2Result> getFwRuleV2Plain(GetFwRuleV2PlainArgs args) {
        return getFwRuleV2Plain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information of an available OpenStack firewall rule v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetRuleV2Args;
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
     *         final var rule = FirewallFunctions.getRuleV2(GetRuleV2Args.builder()
     *             .name("tf_test_rule")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2
     * 
     */
    @Deprecated /* openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2 */
    public static Output<GetFwRuleV2Result> getFwRuleV2(GetFwRuleV2Args args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:index/getFwRuleV2:getFwRuleV2", TypeShape.of(GetFwRuleV2Result.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information of an available OpenStack firewall rule v2.
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
     * import com.pulumi.openstack.firewall.FirewallFunctions;
     * import com.pulumi.openstack.firewall.inputs.GetRuleV2Args;
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
     *         final var rule = FirewallFunctions.getRuleV2(GetRuleV2Args.builder()
     *             .name("tf_test_rule")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * @deprecated
     * openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2
     * 
     */
    @Deprecated /* openstack.index/getfwrulev2.getFwRuleV2 has been deprecated in favor of openstack.firewall/getrulev2.getRuleV2 */
    public static CompletableFuture<GetFwRuleV2Result> getFwRuleV2Plain(GetFwRuleV2PlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:index/getFwRuleV2:getFwRuleV2", TypeShape.of(GetFwRuleV2Result.class), args, Utilities.withVersion(options));
    }
}
