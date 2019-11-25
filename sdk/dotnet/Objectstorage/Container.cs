// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Openstack.Objectstorage
{
    /// <summary>
    /// Manages a V1 container resource within OpenStack.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/objectstorage_container_v1.html.markdown.
    /// </summary>
    public partial class Container : Pulumi.CustomResource
    {
        /// <summary>
        /// Sets an access control list (ACL) that grants
        /// read access. This header can contain a comma-delimited list of users that
        /// can read the container (allows the GET method for all objects in the
        /// container). Changing this updates the access control list read access.
        /// </summary>
        [Output("containerRead")]
        public Output<string?> ContainerRead { get; private set; } = null!;

        /// <summary>
        /// The secret key for container synchronization.
        /// Changing this updates container synchronization.
        /// </summary>
        [Output("containerSyncKey")]
        public Output<string?> ContainerSyncKey { get; private set; } = null!;

        /// <summary>
        /// The destination for container synchronization.
        /// Changing this updates container synchronization.
        /// </summary>
        [Output("containerSyncTo")]
        public Output<string?> ContainerSyncTo { get; private set; } = null!;

        /// <summary>
        /// Sets an ACL that grants write access.
        /// Changing this updates the access control list write access.
        /// </summary>
        [Output("containerWrite")]
        public Output<string?> ContainerWrite { get; private set; } = null!;

        /// <summary>
        /// The MIME type for the container. Changing this
        /// updates the MIME type.
        /// </summary>
        [Output("contentType")]
        public Output<string?> ContentType { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// Custom key/value pairs to associate with the container.
        /// Changing this updates the existing container metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// A unique name for the container. Changing this creates a
        /// new container.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the container. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new container.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Enable object versioning. The structure is described below.
        /// </summary>
        [Output("versioning")]
        public Output<Outputs.ContainerVersioning?> Versioning { get; private set; } = null!;


        /// <summary>
        /// Create a Container resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Container(string name, ContainerArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack:objectstorage/container:Container", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Container(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
            : base("openstack:objectstorage/container:Container", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Container resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Container Get(string name, Input<string> id, ContainerState? state = null, CustomResourceOptions? options = null)
        {
            return new Container(name, id, state, options);
        }
    }

    public sealed class ContainerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets an access control list (ACL) that grants
        /// read access. This header can contain a comma-delimited list of users that
        /// can read the container (allows the GET method for all objects in the
        /// container). Changing this updates the access control list read access.
        /// </summary>
        [Input("containerRead")]
        public Input<string>? ContainerRead { get; set; }

        /// <summary>
        /// The secret key for container synchronization.
        /// Changing this updates container synchronization.
        /// </summary>
        [Input("containerSyncKey")]
        public Input<string>? ContainerSyncKey { get; set; }

        /// <summary>
        /// The destination for container synchronization.
        /// Changing this updates container synchronization.
        /// </summary>
        [Input("containerSyncTo")]
        public Input<string>? ContainerSyncTo { get; set; }

        /// <summary>
        /// Sets an ACL that grants write access.
        /// Changing this updates the access control list write access.
        /// </summary>
        [Input("containerWrite")]
        public Input<string>? ContainerWrite { get; set; }

        /// <summary>
        /// The MIME type for the container. Changing this
        /// updates the MIME type.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Custom key/value pairs to associate with the container.
        /// Changing this updates the existing container metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// A unique name for the container. Changing this creates a
        /// new container.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the container. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new container.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Enable object versioning. The structure is described below.
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.ContainerVersioningArgs>? Versioning { get; set; }

        public ContainerArgs()
        {
        }
    }

    public sealed class ContainerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sets an access control list (ACL) that grants
        /// read access. This header can contain a comma-delimited list of users that
        /// can read the container (allows the GET method for all objects in the
        /// container). Changing this updates the access control list read access.
        /// </summary>
        [Input("containerRead")]
        public Input<string>? ContainerRead { get; set; }

        /// <summary>
        /// The secret key for container synchronization.
        /// Changing this updates container synchronization.
        /// </summary>
        [Input("containerSyncKey")]
        public Input<string>? ContainerSyncKey { get; set; }

        /// <summary>
        /// The destination for container synchronization.
        /// Changing this updates container synchronization.
        /// </summary>
        [Input("containerSyncTo")]
        public Input<string>? ContainerSyncTo { get; set; }

        /// <summary>
        /// Sets an ACL that grants write access.
        /// Changing this updates the access control list write access.
        /// </summary>
        [Input("containerWrite")]
        public Input<string>? ContainerWrite { get; set; }

        /// <summary>
        /// The MIME type for the container. Changing this
        /// updates the MIME type.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Custom key/value pairs to associate with the container.
        /// Changing this updates the existing container metadata.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// A unique name for the container. Changing this creates a
        /// new container.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The region in which to create the container. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new container.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Enable object versioning. The structure is described below.
        /// </summary>
        [Input("versioning")]
        public Input<Inputs.ContainerVersioningGetArgs>? Versioning { get; set; }

        public ContainerState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ContainerVersioningArgs : Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ContainerVersioningArgs()
        {
        }
    }

    public sealed class ContainerVersioningGetArgs : Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ContainerVersioningGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ContainerVersioning
    {
        public readonly string Location;
        public readonly string Type;

        [OutputConstructor]
        private ContainerVersioning(
            string location,
            string type)
        {
            Location = location;
            Type = type;
        }
    }
    }
}
