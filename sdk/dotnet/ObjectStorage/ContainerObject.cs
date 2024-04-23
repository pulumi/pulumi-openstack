// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.ObjectStorage
{
    /// <summary>
    /// Manages a V1 container object resource within OpenStack.
    /// 
    /// ## Example Usage
    /// 
    /// ### Example with simple content
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var container1 = new OpenStack.ObjectStorage.Container("container_1", new()
    ///     {
    ///         Region = "RegionOne",
    ///         Name = "tf-test-container-1",
    ///         Metadata = new[]
    ///         {
    ///             new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["test"] = "true",
    ///             },
    ///         },
    ///         ContentType = "application/json",
    ///     });
    /// 
    ///     var doc1 = new OpenStack.ObjectStorage.ContainerObject("doc_1", new()
    ///     {
    ///         Region = "RegionOne",
    ///         ContainerName = container1.Name,
    ///         Name = "test/default.json",
    ///         Metadata = new[]
    ///         {
    ///             new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["test"] = "true",
    ///             },
    ///         },
    ///         ContentType = "application/json",
    ///         Content = @"               {
    ///                  ""foo"" : ""bar""
    ///                }
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example with content from file
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var container1 = new OpenStack.ObjectStorage.Container("container_1", new()
    ///     {
    ///         Region = "RegionOne",
    ///         Name = "tf-test-container-1",
    ///         Metadata = new[]
    ///         {
    ///             new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["test"] = "true",
    ///             },
    ///         },
    ///         ContentType = "application/json",
    ///     });
    /// 
    ///     var doc1 = new OpenStack.ObjectStorage.ContainerObject("doc_1", new()
    ///     {
    ///         Region = "RegionOne",
    ///         ContainerName = container1.Name,
    ///         Name = "test/default.json",
    ///         Metadata = new[]
    ///         {
    ///             new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["test"] = "true",
    ///             },
    ///         },
    ///         ContentType = "application/json",
    ///         Source = "./default.json",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:objectstorage/containerObject:ContainerObject")]
    public partial class ContainerObject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A unique (within an account) name for the container. 
        /// The container name must be from 1 to 256 characters long and can start
        /// with any character and contain any pattern. Character set must be UTF-8.
        /// The container name cannot contain a slash (/) character because this
        /// character delimits the container and object name. For example, the path
        /// /v1/account/www/pages specifies the www container, not the www/pages container.
        /// </summary>
        [Output("containerName")]
        public Output<string> ContainerName { get; private set; } = null!;

        /// <summary>
        /// A string representing the content of the object. Conflicts with
        /// `source` and `copy_from`.
        /// </summary>
        [Output("content")]
        public Output<string?> Content { get; private set; } = null!;

        /// <summary>
        /// A string which specifies the override behavior for 
        /// the browser. For example, this header might specify that the browser use a download
        /// program to save this file rather than show the file, which is the default.
        /// </summary>
        [Output("contentDisposition")]
        public Output<string> ContentDisposition { get; private set; } = null!;

        /// <summary>
        /// A string representing the value of the Content-Encoding
        /// metadata.
        /// </summary>
        [Output("contentEncoding")]
        public Output<string> ContentEncoding { get; private set; } = null!;

        /// <summary>
        /// If the operation succeeds, this value is zero (0) or the 
        /// length of informational or error text in the response body.
        /// </summary>
        [Output("contentLength")]
        public Output<int> ContentLength { get; private set; } = null!;

        /// <summary>
        /// A string which sets the MIME type for the object.
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// A string representing the name of an object 
        /// used to create the new object by copying the `copy_from` object. The value is in form
        /// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
        /// container and object before you include them in the header. Conflicts with `source` and
        /// `content`.
        /// </summary>
        [Output("copyFrom")]
        public Output<string?> CopyFrom { get; private set; } = null!;

        /// <summary>
        /// The date and time the system responded to the request, using the preferred 
        /// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
        /// time is always in UTC.
        /// </summary>
        [Output("date")]
        public Output<string> Date { get; private set; } = null!;

        /// <summary>
        /// An integer representing the number of seconds after which the
        /// system removes the object. Internally, the Object Storage system stores this value in
        /// the X-Delete-At metadata item.
        /// </summary>
        [Output("deleteAfter")]
        public Output<int?> DeleteAfter { get; private set; } = null!;

        /// <summary>
        /// An string representing the date when the system removes the object. 
        /// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
        /// </summary>
        [Output("deleteAt")]
        public Output<string> DeleteAt { get; private set; } = null!;

        /// <summary>
        /// If set to true, Object Storage guesses the content 
        /// type based on the file extension and ignores the value sent in the Content-Type
        /// header, if present.
        /// </summary>
        [Output("detectContentType")]
        public Output<bool?> DetectContentType { get; private set; } = null!;

        /// <summary>
        /// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The date and time when the object was last modified. The date and time 
        /// stamp format is ISO 8601:
        /// CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00.
        /// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
        /// example, the offset value is -05:00.
        /// </summary>
        [Output("lastModified")]
        public Output<string> LastModified { get; private set; } = null!;

        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// A unique name for the object.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A string set to specify that this is a dynamic large 
        /// object manifest object. The value is the container and object name prefix of the
        /// segment objects in the form container/prefix. You must UTF-8-encode and then
        /// URL-encode the names of the container and prefix before you include them in this
        /// header.
        /// </summary>
        [Output("objectManifest")]
        public Output<string> ObjectManifest { get; private set; } = null!;

        /// <summary>
        /// The region in which to create the container. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new container.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A string representing the local path of a file which will be used
        /// as the object's content. Conflicts with `source` and `copy_from`.
        /// </summary>
        [Output("source")]
        public Output<string?> Source { get; private set; } = null!;

        /// <summary>
        /// A unique transaction ID for this request. Your service provider might 
        /// need this value if you report a problem.
        /// </summary>
        [Output("transId")]
        public Output<string> TransId { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerObject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerObject(string name, ContainerObjectArgs args, CustomResourceOptions? options = null)
            : base("openstack:objectstorage/containerObject:ContainerObject", name, args ?? new ContainerObjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerObject(string name, Input<string> id, ContainerObjectState? state = null, CustomResourceOptions? options = null)
            : base("openstack:objectstorage/containerObject:ContainerObject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerObject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerObject Get(string name, Input<string> id, ContainerObjectState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerObject(name, id, state, options);
        }
    }

    public sealed class ContainerObjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique (within an account) name for the container. 
        /// The container name must be from 1 to 256 characters long and can start
        /// with any character and contain any pattern. Character set must be UTF-8.
        /// The container name cannot contain a slash (/) character because this
        /// character delimits the container and object name. For example, the path
        /// /v1/account/www/pages specifies the www container, not the www/pages container.
        /// </summary>
        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// A string representing the content of the object. Conflicts with
        /// `source` and `copy_from`.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// A string which specifies the override behavior for 
        /// the browser. For example, this header might specify that the browser use a download
        /// program to save this file rather than show the file, which is the default.
        /// </summary>
        [Input("contentDisposition")]
        public Input<string>? ContentDisposition { get; set; }

        /// <summary>
        /// A string representing the value of the Content-Encoding
        /// metadata.
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// A string which sets the MIME type for the object.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// A string representing the name of an object 
        /// used to create the new object by copying the `copy_from` object. The value is in form
        /// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
        /// container and object before you include them in the header. Conflicts with `source` and
        /// `content`.
        /// </summary>
        [Input("copyFrom")]
        public Input<string>? CopyFrom { get; set; }

        /// <summary>
        /// An integer representing the number of seconds after which the
        /// system removes the object. Internally, the Object Storage system stores this value in
        /// the X-Delete-At metadata item.
        /// </summary>
        [Input("deleteAfter")]
        public Input<int>? DeleteAfter { get; set; }

        /// <summary>
        /// An string representing the date when the system removes the object. 
        /// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
        /// </summary>
        [Input("deleteAt")]
        public Input<string>? DeleteAt { get; set; }

        /// <summary>
        /// If set to true, Object Storage guesses the content 
        /// type based on the file extension and ignores the value sent in the Content-Type
        /// header, if present.
        /// </summary>
        [Input("detectContentType")]
        public Input<bool>? DetectContentType { get; set; }

        /// <summary>
        /// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// A unique name for the object.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A string set to specify that this is a dynamic large 
        /// object manifest object. The value is the container and object name prefix of the
        /// segment objects in the form container/prefix. You must UTF-8-encode and then
        /// URL-encode the names of the container and prefix before you include them in this
        /// header.
        /// </summary>
        [Input("objectManifest")]
        public Input<string>? ObjectManifest { get; set; }

        /// <summary>
        /// The region in which to create the container. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new container.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// A string representing the local path of a file which will be used
        /// as the object's content. Conflicts with `source` and `copy_from`.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        public ContainerObjectArgs()
        {
        }
        public static new ContainerObjectArgs Empty => new ContainerObjectArgs();
    }

    public sealed class ContainerObjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique (within an account) name for the container. 
        /// The container name must be from 1 to 256 characters long and can start
        /// with any character and contain any pattern. Character set must be UTF-8.
        /// The container name cannot contain a slash (/) character because this
        /// character delimits the container and object name. For example, the path
        /// /v1/account/www/pages specifies the www container, not the www/pages container.
        /// </summary>
        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        /// <summary>
        /// A string representing the content of the object. Conflicts with
        /// `source` and `copy_from`.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// A string which specifies the override behavior for 
        /// the browser. For example, this header might specify that the browser use a download
        /// program to save this file rather than show the file, which is the default.
        /// </summary>
        [Input("contentDisposition")]
        public Input<string>? ContentDisposition { get; set; }

        /// <summary>
        /// A string representing the value of the Content-Encoding
        /// metadata.
        /// </summary>
        [Input("contentEncoding")]
        public Input<string>? ContentEncoding { get; set; }

        /// <summary>
        /// If the operation succeeds, this value is zero (0) or the 
        /// length of informational or error text in the response body.
        /// </summary>
        [Input("contentLength")]
        public Input<int>? ContentLength { get; set; }

        /// <summary>
        /// A string which sets the MIME type for the object.
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// A string representing the name of an object 
        /// used to create the new object by copying the `copy_from` object. The value is in form
        /// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
        /// container and object before you include them in the header. Conflicts with `source` and
        /// `content`.
        /// </summary>
        [Input("copyFrom")]
        public Input<string>? CopyFrom { get; set; }

        /// <summary>
        /// The date and time the system responded to the request, using the preferred 
        /// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
        /// time is always in UTC.
        /// </summary>
        [Input("date")]
        public Input<string>? Date { get; set; }

        /// <summary>
        /// An integer representing the number of seconds after which the
        /// system removes the object. Internally, the Object Storage system stores this value in
        /// the X-Delete-At metadata item.
        /// </summary>
        [Input("deleteAfter")]
        public Input<int>? DeleteAfter { get; set; }

        /// <summary>
        /// An string representing the date when the system removes the object. 
        /// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
        /// </summary>
        [Input("deleteAt")]
        public Input<string>? DeleteAt { get; set; }

        /// <summary>
        /// If set to true, Object Storage guesses the content 
        /// type based on the file extension and ignores the value sent in the Content-Type
        /// header, if present.
        /// </summary>
        [Input("detectContentType")]
        public Input<bool>? DetectContentType { get; set; }

        /// <summary>
        /// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The date and time when the object was last modified. The date and time 
        /// stamp format is ISO 8601:
        /// CCYY-MM-DDThh:mm:ss±hh:mm
        /// For example, 2015-08-27T09:49:58-05:00.
        /// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
        /// example, the offset value is -05:00.
        /// </summary>
        [Input("lastModified")]
        public Input<string>? LastModified { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// A unique name for the object.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A string set to specify that this is a dynamic large 
        /// object manifest object. The value is the container and object name prefix of the
        /// segment objects in the form container/prefix. You must UTF-8-encode and then
        /// URL-encode the names of the container and prefix before you include them in this
        /// header.
        /// </summary>
        [Input("objectManifest")]
        public Input<string>? ObjectManifest { get; set; }

        /// <summary>
        /// The region in which to create the container. If
        /// omitted, the `region` argument of the provider is used. Changing this
        /// creates a new container.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// A string representing the local path of a file which will be used
        /// as the object's content. Conflicts with `source` and `copy_from`.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// A unique transaction ID for this request. Your service provider might 
        /// need this value if you report a problem.
        /// </summary>
        [Input("transId")]
        public Input<string>? TransId { get; set; }

        public ContainerObjectState()
        {
        }
        public static new ContainerObjectState Empty => new ContainerObjectState();
    }
}
