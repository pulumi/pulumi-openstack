// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.keymanager;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
import com.pulumi.openstack.keymanager.inputs.GetContainerPlainArgs;
import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
import com.pulumi.openstack.keymanager.inputs.GetSecretPlainArgs;
import com.pulumi.openstack.keymanager.outputs.GetContainerResult;
import com.pulumi.openstack.keymanager.outputs.GetSecretResult;
import java.util.concurrent.CompletableFuture;

public final class KeymanagerFunctions {
    /**
     * Use this data source to get the ID of an available Barbican container.
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
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
     *         final var example = KeymanagerFunctions.getContainer(GetContainerArgs.builder()
     *             .name("my_container")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetContainerResult> getContainer() {
        return getContainer(GetContainerArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Barbican container.
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
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
     *         final var example = KeymanagerFunctions.getContainer(GetContainerArgs.builder()
     *             .name("my_container")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetContainerResult> getContainerPlain() {
        return getContainerPlain(GetContainerPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Barbican container.
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
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
     *         final var example = KeymanagerFunctions.getContainer(GetContainerArgs.builder()
     *             .name("my_container")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetContainerResult> getContainer(GetContainerArgs args) {
        return getContainer(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Barbican container.
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
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
     *         final var example = KeymanagerFunctions.getContainer(GetContainerArgs.builder()
     *             .name("my_container")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetContainerResult> getContainerPlain(GetContainerPlainArgs args) {
        return getContainerPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available Barbican container.
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
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
     *         final var example = KeymanagerFunctions.getContainer(GetContainerArgs.builder()
     *             .name("my_container")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetContainerResult> getContainer(GetContainerArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:keymanager/getContainer:getContainer", TypeShape.of(GetContainerResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Barbican container.
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
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
     *         final var example = KeymanagerFunctions.getContainer(GetContainerArgs.builder()
     *             .name("my_container")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetContainerResult> getContainer(GetContainerArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("openstack:keymanager/getContainer:getContainer", TypeShape.of(GetContainerResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available Barbican container.
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetContainerArgs;
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
     *         final var example = KeymanagerFunctions.getContainer(GetContainerArgs.builder()
     *             .name("my_container")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetContainerResult> getContainerPlain(GetContainerPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:keymanager/getContainer:getContainer", TypeShape.of(GetContainerResult.class), args, Utilities.withVersion(options));
    }
    /**
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var example = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Date Filters
     * 
     * The values for the `expiration_filter`, `created_at_filter`, and
     * `updated_at_filter` parameters are comma-separated lists of time stamps in
     * RFC3339 format. The time stamps can be prefixed with any of these comparison
     * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
     * (less-than), *lte:* (less-than-or-equal).
     * 
     * For example, to get a passphrase a Secret with CBC moda, that will expire in
     * January of 2020:
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var dateFilterExample = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .expirationFilter("gt:2020-01-01T00:00:00Z")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSecretResult> getSecret() {
        return getSecret(GetSecretArgs.Empty, InvokeOptions.Empty);
    }
    /**
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var example = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Date Filters
     * 
     * The values for the `expiration_filter`, `created_at_filter`, and
     * `updated_at_filter` parameters are comma-separated lists of time stamps in
     * RFC3339 format. The time stamps can be prefixed with any of these comparison
     * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
     * (less-than), *lte:* (less-than-or-equal).
     * 
     * For example, to get a passphrase a Secret with CBC moda, that will expire in
     * January of 2020:
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var dateFilterExample = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .expirationFilter("gt:2020-01-01T00:00:00Z")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSecretResult> getSecretPlain() {
        return getSecretPlain(GetSecretPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var example = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Date Filters
     * 
     * The values for the `expiration_filter`, `created_at_filter`, and
     * `updated_at_filter` parameters are comma-separated lists of time stamps in
     * RFC3339 format. The time stamps can be prefixed with any of these comparison
     * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
     * (less-than), *lte:* (less-than-or-equal).
     * 
     * For example, to get a passphrase a Secret with CBC moda, that will expire in
     * January of 2020:
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var dateFilterExample = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .expirationFilter("gt:2020-01-01T00:00:00Z")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSecretResult> getSecret(GetSecretArgs args) {
        return getSecret(args, InvokeOptions.Empty);
    }
    /**
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var example = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Date Filters
     * 
     * The values for the `expiration_filter`, `created_at_filter`, and
     * `updated_at_filter` parameters are comma-separated lists of time stamps in
     * RFC3339 format. The time stamps can be prefixed with any of these comparison
     * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
     * (less-than), *lte:* (less-than-or-equal).
     * 
     * For example, to get a passphrase a Secret with CBC moda, that will expire in
     * January of 2020:
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var dateFilterExample = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .expirationFilter("gt:2020-01-01T00:00:00Z")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSecretResult> getSecretPlain(GetSecretPlainArgs args) {
        return getSecretPlain(args, InvokeOptions.Empty);
    }
    /**
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var example = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Date Filters
     * 
     * The values for the `expiration_filter`, `created_at_filter`, and
     * `updated_at_filter` parameters are comma-separated lists of time stamps in
     * RFC3339 format. The time stamps can be prefixed with any of these comparison
     * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
     * (less-than), *lte:* (less-than-or-equal).
     * 
     * For example, to get a passphrase a Secret with CBC moda, that will expire in
     * January of 2020:
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var dateFilterExample = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .expirationFilter("gt:2020-01-01T00:00:00Z")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSecretResult> getSecret(GetSecretArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:keymanager/getSecret:getSecret", TypeShape.of(GetSecretResult.class), args, Utilities.withVersion(options));
    }
    /**
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var example = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Date Filters
     * 
     * The values for the `expiration_filter`, `created_at_filter`, and
     * `updated_at_filter` parameters are comma-separated lists of time stamps in
     * RFC3339 format. The time stamps can be prefixed with any of these comparison
     * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
     * (less-than), *lte:* (less-than-or-equal).
     * 
     * For example, to get a passphrase a Secret with CBC moda, that will expire in
     * January of 2020:
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var dateFilterExample = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .expirationFilter("gt:2020-01-01T00:00:00Z")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSecretResult> getSecret(GetSecretArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("openstack:keymanager/getSecret:getSecret", TypeShape.of(GetSecretResult.class), args, Utilities.withVersion(options));
    }
    /**
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
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var example = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Date Filters
     * 
     * The values for the `expiration_filter`, `created_at_filter`, and
     * `updated_at_filter` parameters are comma-separated lists of time stamps in
     * RFC3339 format. The time stamps can be prefixed with any of these comparison
     * operators: *gt:* (greater-than), *gte:* (greater-than-or-equal), *lt:*
     * (less-than), *lte:* (less-than-or-equal).
     * 
     * For example, to get a passphrase a Secret with CBC moda, that will expire in
     * January of 2020:
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.openstack.keymanager.KeymanagerFunctions;
     * import com.pulumi.openstack.keymanager.inputs.GetSecretArgs;
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
     *         final var dateFilterExample = KeymanagerFunctions.getSecret(GetSecretArgs.builder()
     *             .mode("cbc")
     *             .secretType("passphrase")
     *             .expirationFilter("gt:2020-01-01T00:00:00Z")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSecretResult> getSecretPlain(GetSecretPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:keymanager/getSecret:getSecret", TypeShape.of(GetSecretResult.class), args, Utilities.withVersion(options));
    }
}
