// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.firewall;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.firewall.inputs.GetPolicyArgs;
import com.pulumi.openstack.firewall.inputs.GetPolicyPlainArgs;
import com.pulumi.openstack.firewall.outputs.GetPolicyResult;
import java.util.concurrent.CompletableFuture;

public final class FirewallFunctions {
    /**
     * Use this data source to get firewall policy information of an available OpenStack firewall policy.
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
     * import com.pulumi.openstack.firewall.inputs.GetPolicyArgs;
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
     *         final var policy = FirewallFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPolicyResult> getPolicy() {
        return getPolicy(GetPolicyArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get firewall policy information of an available OpenStack firewall policy.
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
     * import com.pulumi.openstack.firewall.inputs.GetPolicyArgs;
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
     *         final var policy = FirewallFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPolicyResult> getPolicyPlain() {
        return getPolicyPlain(GetPolicyPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get firewall policy information of an available OpenStack firewall policy.
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
     * import com.pulumi.openstack.firewall.inputs.GetPolicyArgs;
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
     *         final var policy = FirewallFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPolicyResult> getPolicy(GetPolicyArgs args) {
        return getPolicy(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get firewall policy information of an available OpenStack firewall policy.
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
     * import com.pulumi.openstack.firewall.inputs.GetPolicyArgs;
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
     *         final var policy = FirewallFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPolicyResult> getPolicyPlain(GetPolicyPlainArgs args) {
        return getPolicyPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get firewall policy information of an available OpenStack firewall policy.
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
     * import com.pulumi.openstack.firewall.inputs.GetPolicyArgs;
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
     *         final var policy = FirewallFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPolicyResult> getPolicy(GetPolicyArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:firewall/getPolicy:getPolicy", TypeShape.of(GetPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get firewall policy information of an available OpenStack firewall policy.
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
     * import com.pulumi.openstack.firewall.inputs.GetPolicyArgs;
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
     *         final var policy = FirewallFunctions.getPolicy(GetPolicyArgs.builder()
     *             .name("tf_test_policy")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPolicyResult> getPolicyPlain(GetPolicyPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:firewall/getPolicy:getPolicy", TypeShape.of(GetPolicyResult.class), args, Utilities.withVersion(options));
    }
}
