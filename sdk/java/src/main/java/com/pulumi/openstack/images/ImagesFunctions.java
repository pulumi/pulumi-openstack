// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.openstack.images;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.openstack.Utilities;
import com.pulumi.openstack.images.inputs.GetImageArgs;
import com.pulumi.openstack.images.inputs.GetImageIdsArgs;
import com.pulumi.openstack.images.inputs.GetImageIdsPlainArgs;
import com.pulumi.openstack.images.inputs.GetImagePlainArgs;
import com.pulumi.openstack.images.outputs.GetImageIdsResult;
import com.pulumi.openstack.images.outputs.GetImageResult;
import java.util.concurrent.CompletableFuture;

public final class ImagesFunctions {
    /**
     * Use this data source to get the ID of an available OpenStack image.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageArgs;
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
     *         final var ubuntu = ImagesFunctions.getImage(GetImageArgs.builder()
     *             .mostRecent(true)
     *             .name(&#34;Ubuntu 16.04&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetImageResult> getImage() {
        return getImage(GetImageArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack image.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageArgs;
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
     *         final var ubuntu = ImagesFunctions.getImage(GetImageArgs.builder()
     *             .mostRecent(true)
     *             .name(&#34;Ubuntu 16.04&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetImageResult> getImagePlain() {
        return getImagePlain(GetImagePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack image.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageArgs;
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
     *         final var ubuntu = ImagesFunctions.getImage(GetImageArgs.builder()
     *             .mostRecent(true)
     *             .name(&#34;Ubuntu 16.04&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetImageResult> getImage(GetImageArgs args) {
        return getImage(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack image.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageArgs;
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
     *         final var ubuntu = ImagesFunctions.getImage(GetImageArgs.builder()
     *             .mostRecent(true)
     *             .name(&#34;Ubuntu 16.04&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetImageResult> getImagePlain(GetImagePlainArgs args) {
        return getImagePlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get the ID of an available OpenStack image.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageArgs;
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
     *         final var ubuntu = ImagesFunctions.getImage(GetImageArgs.builder()
     *             .mostRecent(true)
     *             .name(&#34;Ubuntu 16.04&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetImageResult> getImage(GetImageArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:images/getImage:getImage", TypeShape.of(GetImageResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get the ID of an available OpenStack image.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageArgs;
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
     *         final var ubuntu = ImagesFunctions.getImage(GetImageArgs.builder()
     *             .mostRecent(true)
     *             .name(&#34;Ubuntu 16.04&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetImageResult> getImagePlain(GetImagePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:images/getImage:getImage", TypeShape.of(GetImageResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get a list of Openstack Image IDs matching the
     * specified criteria.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageIdsArgs;
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
     *         final var images = ImagesFunctions.getImageIds(GetImageIdsArgs.builder()
     *             .nameRegex(&#34;^Ubuntu 16\\.04.*-amd64&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .sort(&#34;updated_at&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetImageIdsResult> getImageIds() {
        return getImageIds(GetImageIdsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Openstack Image IDs matching the
     * specified criteria.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageIdsArgs;
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
     *         final var images = ImagesFunctions.getImageIds(GetImageIdsArgs.builder()
     *             .nameRegex(&#34;^Ubuntu 16\\.04.*-amd64&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .sort(&#34;updated_at&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetImageIdsResult> getImageIdsPlain() {
        return getImageIdsPlain(GetImageIdsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Openstack Image IDs matching the
     * specified criteria.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageIdsArgs;
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
     *         final var images = ImagesFunctions.getImageIds(GetImageIdsArgs.builder()
     *             .nameRegex(&#34;^Ubuntu 16\\.04.*-amd64&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .sort(&#34;updated_at&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetImageIdsResult> getImageIds(GetImageIdsArgs args) {
        return getImageIds(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Openstack Image IDs matching the
     * specified criteria.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageIdsArgs;
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
     *         final var images = ImagesFunctions.getImageIds(GetImageIdsArgs.builder()
     *             .nameRegex(&#34;^Ubuntu 16\\.04.*-amd64&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .sort(&#34;updated_at&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetImageIdsResult> getImageIdsPlain(GetImageIdsPlainArgs args) {
        return getImageIdsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get a list of Openstack Image IDs matching the
     * specified criteria.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageIdsArgs;
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
     *         final var images = ImagesFunctions.getImageIds(GetImageIdsArgs.builder()
     *             .nameRegex(&#34;^Ubuntu 16\\.04.*-amd64&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .sort(&#34;updated_at&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetImageIdsResult> getImageIds(GetImageIdsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("openstack:images/getImageIds:getImageIds", TypeShape.of(GetImageIdsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get a list of Openstack Image IDs matching the
     * specified criteria.
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
     * import com.pulumi.openstack.images.ImagesFunctions;
     * import com.pulumi.openstack.images.inputs.GetImageIdsArgs;
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
     *         final var images = ImagesFunctions.getImageIds(GetImageIdsArgs.builder()
     *             .nameRegex(&#34;^Ubuntu 16\\.04.*-amd64&#34;)
     *             .properties(Map.of(&#34;key&#34;, &#34;value&#34;))
     *             .sort(&#34;updated_at&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetImageIdsResult> getImageIdsPlain(GetImageIdsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("openstack:images/getImageIds:getImageIds", TypeShape.of(GetImageIdsResult.class), args, Utilities.withVersion(options));
    }
}
