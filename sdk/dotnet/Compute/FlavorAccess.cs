// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.Compute
{
    /// <summary>
    /// Manages a project access for flavor V2 resource within OpenStack.
    /// 
    /// Note: You _must_ have admin privileges in your OpenStack cloud to use
    /// this resource.
    /// 
    /// ---
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var project1 = new OpenStack.Identity.Project("project1", new OpenStack.Identity.ProjectArgs
    ///         {
    ///         });
    ///         var flavor1 = new OpenStack.Compute.Flavor("flavor1", new OpenStack.Compute.FlavorArgs
    ///         {
    ///             Disk = "20",
    ///             IsPublic = false,
    ///             Ram = "8096",
    ///             Vcpus = "2",
    ///         });
    ///         var access1 = new OpenStack.Compute.FlavorAccess("access1", new OpenStack.Compute.FlavorAccessArgs
    ///         {
    ///             FlavorId = flavor1.Id,
    ///             TenantId = project1.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class FlavorAccess : Pulumi.CustomResource
    {
        /// <summary>
        /// The UUID of flavor to use. Changing this creates a new flavor access.
        /// </summary>
        [Output("flavorId")]
        public Output<string> FlavorId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new flavor access.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The UUID of tenant which is allowed to use the flavor.
        /// Changing this creates a new flavor access.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a FlavorAccess resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlavorAccess(string name, FlavorAccessArgs args, CustomResourceOptions? options = null)
            : base("openstack:compute/flavorAccess:FlavorAccess", name, args ?? new FlavorAccessArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlavorAccess(string name, Input<string> id, FlavorAccessState? state = null, CustomResourceOptions? options = null)
            : base("openstack:compute/flavorAccess:FlavorAccess", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FlavorAccess resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlavorAccess Get(string name, Input<string> id, FlavorAccessState? state = null, CustomResourceOptions? options = null)
        {
            return new FlavorAccess(name, id, state, options);
        }
    }

    public sealed class FlavorAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of flavor to use. Changing this creates a new flavor access.
        /// </summary>
        [Input("flavorId", required: true)]
        public Input<string> FlavorId { get; set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new flavor access.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of tenant which is allowed to use the flavor.
        /// Changing this creates a new flavor access.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public FlavorAccessArgs()
        {
        }
    }

    public sealed class FlavorAccessState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UUID of flavor to use. Changing this creates a new flavor access.
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Compute client.
        /// If omitted, the `region` argument of the provider is used.
        /// Changing this creates a new flavor access.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The UUID of tenant which is allowed to use the flavor.
        /// Changing this creates a new flavor access.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public FlavorAccessState()
        {
        }
    }
}
