// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 container object resource within OpenStack.
//
// ## Example Usage
//
// ### Example with simple content
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			container1, err := objectstorage.NewContainer(ctx, "container_1", &objectstorage.ContainerArgs{
//				Region: pulumi.String("RegionOne"),
//				Name:   pulumi.String("tf-test-container-1"),
//				Metadata: pulumi.Map{
//					pulumi.Any(map[string]interface{}{
//						"test": "true",
//					}),
//				},
//				ContentType: pulumi.String("application/json"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = objectstorage.NewContainerObject(ctx, "doc_1", &objectstorage.ContainerObjectArgs{
//				Region:        pulumi.String("RegionOne"),
//				ContainerName: container1.Name,
//				Name:          pulumi.String("test/default.json"),
//				Metadata: pulumi.Map{
//					pulumi.Any(map[string]interface{}{
//						"test": "true",
//					}),
//				},
//				ContentType: pulumi.String("application/json"),
//				Content:     pulumi.String("               {\n                 \"foo\" : \"bar\"\n               }\n"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Example with content from file
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/objectstorage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			container1, err := objectstorage.NewContainer(ctx, "container_1", &objectstorage.ContainerArgs{
//				Region: pulumi.String("RegionOne"),
//				Name:   pulumi.String("tf-test-container-1"),
//				Metadata: pulumi.Map{
//					pulumi.Any(map[string]interface{}{
//						"test": "true",
//					}),
//				},
//				ContentType: pulumi.String("application/json"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = objectstorage.NewContainerObject(ctx, "doc_1", &objectstorage.ContainerObjectArgs{
//				Region:        pulumi.String("RegionOne"),
//				ContainerName: container1.Name,
//				Name:          pulumi.String("test/default.json"),
//				Metadata: pulumi.Map{
//					pulumi.Any(map[string]interface{}{
//						"test": "true",
//					}),
//				},
//				ContentType: pulumi.String("application/json"),
//				Source:      pulumi.String("./default.json"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ContainerObject struct {
	pulumi.CustomResourceState

	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	ContainerName pulumi.StringOutput `pulumi:"containerName"`
	// A string representing the content of the object. Conflicts with
	// `source` and `copyFrom`.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	ContentDisposition pulumi.StringOutput `pulumi:"contentDisposition"`
	// A string representing the value of the Content-Encoding
	// metadata.
	ContentEncoding pulumi.StringOutput `pulumi:"contentEncoding"`
	// If the operation succeeds, this value is zero (0) or the
	// length of informational or error text in the response body.
	ContentLength pulumi.IntOutput `pulumi:"contentLength"`
	// A string which sets the MIME type for the object.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// A string representing the name of an object
	// used to create the new object by copying the `copyFrom` object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with `source` and
	// `content`.
	CopyFrom pulumi.StringPtrOutput `pulumi:"copyFrom"`
	// The date and time the system responded to the request, using the preferred
	// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
	// time is always in UTC.
	Date pulumi.StringOutput `pulumi:"date"`
	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	DeleteAfter pulumi.IntPtrOutput `pulumi:"deleteAfter"`
	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	DeleteAt pulumi.StringOutput `pulumi:"deleteAt"`
	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	DetectContentType pulumi.BoolPtrOutput `pulumi:"detectContentType"`
	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The date and time when the object was last modified. The date and time
	// stamp format is ISO 8601:
	// CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00.
	// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
	// example, the offset value is -05:00.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	Metadata     pulumi.MapOutput    `pulumi:"metadata"`
	// A unique name for the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	ObjectManifest pulumi.StringOutput `pulumi:"objectManifest"`
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region pulumi.StringOutput `pulumi:"region"`
	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with `source` and `copyFrom`.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// A unique transaction ID for this request. Your service provider might
	// need this value if you report a problem.
	TransId pulumi.StringOutput `pulumi:"transId"`
}

// NewContainerObject registers a new resource with the given unique name, arguments, and options.
func NewContainerObject(ctx *pulumi.Context,
	name string, args *ContainerObjectArgs, opts ...pulumi.ResourceOption) (*ContainerObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerObject
	err := ctx.RegisterResource("openstack:objectstorage/containerObject:ContainerObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerObject gets an existing ContainerObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerObjectState, opts ...pulumi.ResourceOption) (*ContainerObject, error) {
	var resource ContainerObject
	err := ctx.ReadResource("openstack:objectstorage/containerObject:ContainerObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerObject resources.
type containerObjectState struct {
	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	ContainerName *string `pulumi:"containerName"`
	// A string representing the content of the object. Conflicts with
	// `source` and `copyFrom`.
	Content *string `pulumi:"content"`
	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// A string representing the value of the Content-Encoding
	// metadata.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// If the operation succeeds, this value is zero (0) or the
	// length of informational or error text in the response body.
	ContentLength *int `pulumi:"contentLength"`
	// A string which sets the MIME type for the object.
	ContentType *string `pulumi:"contentType"`
	// A string representing the name of an object
	// used to create the new object by copying the `copyFrom` object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with `source` and
	// `content`.
	CopyFrom *string `pulumi:"copyFrom"`
	// The date and time the system responded to the request, using the preferred
	// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
	// time is always in UTC.
	Date *string `pulumi:"date"`
	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	DeleteAfter *int `pulumi:"deleteAfter"`
	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	DeleteAt *string `pulumi:"deleteAt"`
	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	DetectContentType *bool `pulumi:"detectContentType"`
	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	Etag *string `pulumi:"etag"`
	// The date and time when the object was last modified. The date and time
	// stamp format is ISO 8601:
	// CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00.
	// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
	// example, the offset value is -05:00.
	LastModified *string                `pulumi:"lastModified"`
	Metadata     map[string]interface{} `pulumi:"metadata"`
	// A unique name for the object.
	Name *string `pulumi:"name"`
	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	ObjectManifest *string `pulumi:"objectManifest"`
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region *string `pulumi:"region"`
	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with `source` and `copyFrom`.
	Source *string `pulumi:"source"`
	// A unique transaction ID for this request. Your service provider might
	// need this value if you report a problem.
	TransId *string `pulumi:"transId"`
}

type ContainerObjectState struct {
	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	ContainerName pulumi.StringPtrInput
	// A string representing the content of the object. Conflicts with
	// `source` and `copyFrom`.
	Content pulumi.StringPtrInput
	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	ContentDisposition pulumi.StringPtrInput
	// A string representing the value of the Content-Encoding
	// metadata.
	ContentEncoding pulumi.StringPtrInput
	// If the operation succeeds, this value is zero (0) or the
	// length of informational or error text in the response body.
	ContentLength pulumi.IntPtrInput
	// A string which sets the MIME type for the object.
	ContentType pulumi.StringPtrInput
	// A string representing the name of an object
	// used to create the new object by copying the `copyFrom` object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with `source` and
	// `content`.
	CopyFrom pulumi.StringPtrInput
	// The date and time the system responded to the request, using the preferred
	// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
	// time is always in UTC.
	Date pulumi.StringPtrInput
	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	DeleteAfter pulumi.IntPtrInput
	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	DeleteAt pulumi.StringPtrInput
	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	DetectContentType pulumi.BoolPtrInput
	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	Etag pulumi.StringPtrInput
	// The date and time when the object was last modified. The date and time
	// stamp format is ISO 8601:
	// CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00.
	// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
	// example, the offset value is -05:00.
	LastModified pulumi.StringPtrInput
	Metadata     pulumi.MapInput
	// A unique name for the object.
	Name pulumi.StringPtrInput
	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	ObjectManifest pulumi.StringPtrInput
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region pulumi.StringPtrInput
	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with `source` and `copyFrom`.
	Source pulumi.StringPtrInput
	// A unique transaction ID for this request. Your service provider might
	// need this value if you report a problem.
	TransId pulumi.StringPtrInput
}

func (ContainerObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerObjectState)(nil)).Elem()
}

type containerObjectArgs struct {
	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	ContainerName string `pulumi:"containerName"`
	// A string representing the content of the object. Conflicts with
	// `source` and `copyFrom`.
	Content *string `pulumi:"content"`
	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// A string representing the value of the Content-Encoding
	// metadata.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// A string which sets the MIME type for the object.
	ContentType *string `pulumi:"contentType"`
	// A string representing the name of an object
	// used to create the new object by copying the `copyFrom` object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with `source` and
	// `content`.
	CopyFrom *string `pulumi:"copyFrom"`
	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	DeleteAfter *int `pulumi:"deleteAfter"`
	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	DeleteAt *string `pulumi:"deleteAt"`
	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	DetectContentType *bool `pulumi:"detectContentType"`
	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	Etag     *string                `pulumi:"etag"`
	Metadata map[string]interface{} `pulumi:"metadata"`
	// A unique name for the object.
	Name *string `pulumi:"name"`
	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	ObjectManifest *string `pulumi:"objectManifest"`
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region *string `pulumi:"region"`
	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with `source` and `copyFrom`.
	Source *string `pulumi:"source"`
}

// The set of arguments for constructing a ContainerObject resource.
type ContainerObjectArgs struct {
	// A unique (within an account) name for the container.
	// The container name must be from 1 to 256 characters long and can start
	// with any character and contain any pattern. Character set must be UTF-8.
	// The container name cannot contain a slash (/) character because this
	// character delimits the container and object name. For example, the path
	// /v1/account/www/pages specifies the www container, not the www/pages container.
	ContainerName pulumi.StringInput
	// A string representing the content of the object. Conflicts with
	// `source` and `copyFrom`.
	Content pulumi.StringPtrInput
	// A string which specifies the override behavior for
	// the browser. For example, this header might specify that the browser use a download
	// program to save this file rather than show the file, which is the default.
	ContentDisposition pulumi.StringPtrInput
	// A string representing the value of the Content-Encoding
	// metadata.
	ContentEncoding pulumi.StringPtrInput
	// A string which sets the MIME type for the object.
	ContentType pulumi.StringPtrInput
	// A string representing the name of an object
	// used to create the new object by copying the `copyFrom` object. The value is in form
	// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
	// container and object before you include them in the header. Conflicts with `source` and
	// `content`.
	CopyFrom pulumi.StringPtrInput
	// An integer representing the number of seconds after which the
	// system removes the object. Internally, the Object Storage system stores this value in
	// the X-Delete-At metadata item.
	DeleteAfter pulumi.IntPtrInput
	// An string representing the date when the system removes the object.
	// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
	DeleteAt pulumi.StringPtrInput
	// If set to true, Object Storage guesses the content
	// type based on the file extension and ignores the value sent in the Content-Type
	// header, if present.
	DetectContentType pulumi.BoolPtrInput
	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	Etag     pulumi.StringPtrInput
	Metadata pulumi.MapInput
	// A unique name for the object.
	Name pulumi.StringPtrInput
	// A string set to specify that this is a dynamic large
	// object manifest object. The value is the container and object name prefix of the
	// segment objects in the form container/prefix. You must UTF-8-encode and then
	// URL-encode the names of the container and prefix before you include them in this
	// header.
	ObjectManifest pulumi.StringPtrInput
	// The region in which to create the container. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new container.
	Region pulumi.StringPtrInput
	// A string representing the local path of a file which will be used
	// as the object's content. Conflicts with `source` and `copyFrom`.
	Source pulumi.StringPtrInput
}

func (ContainerObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerObjectArgs)(nil)).Elem()
}

type ContainerObjectInput interface {
	pulumi.Input

	ToContainerObjectOutput() ContainerObjectOutput
	ToContainerObjectOutputWithContext(ctx context.Context) ContainerObjectOutput
}

func (*ContainerObject) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerObject)(nil)).Elem()
}

func (i *ContainerObject) ToContainerObjectOutput() ContainerObjectOutput {
	return i.ToContainerObjectOutputWithContext(context.Background())
}

func (i *ContainerObject) ToContainerObjectOutputWithContext(ctx context.Context) ContainerObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerObjectOutput)
}

// ContainerObjectArrayInput is an input type that accepts ContainerObjectArray and ContainerObjectArrayOutput values.
// You can construct a concrete instance of `ContainerObjectArrayInput` via:
//
//	ContainerObjectArray{ ContainerObjectArgs{...} }
type ContainerObjectArrayInput interface {
	pulumi.Input

	ToContainerObjectArrayOutput() ContainerObjectArrayOutput
	ToContainerObjectArrayOutputWithContext(context.Context) ContainerObjectArrayOutput
}

type ContainerObjectArray []ContainerObjectInput

func (ContainerObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerObject)(nil)).Elem()
}

func (i ContainerObjectArray) ToContainerObjectArrayOutput() ContainerObjectArrayOutput {
	return i.ToContainerObjectArrayOutputWithContext(context.Background())
}

func (i ContainerObjectArray) ToContainerObjectArrayOutputWithContext(ctx context.Context) ContainerObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerObjectArrayOutput)
}

// ContainerObjectMapInput is an input type that accepts ContainerObjectMap and ContainerObjectMapOutput values.
// You can construct a concrete instance of `ContainerObjectMapInput` via:
//
//	ContainerObjectMap{ "key": ContainerObjectArgs{...} }
type ContainerObjectMapInput interface {
	pulumi.Input

	ToContainerObjectMapOutput() ContainerObjectMapOutput
	ToContainerObjectMapOutputWithContext(context.Context) ContainerObjectMapOutput
}

type ContainerObjectMap map[string]ContainerObjectInput

func (ContainerObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerObject)(nil)).Elem()
}

func (i ContainerObjectMap) ToContainerObjectMapOutput() ContainerObjectMapOutput {
	return i.ToContainerObjectMapOutputWithContext(context.Background())
}

func (i ContainerObjectMap) ToContainerObjectMapOutputWithContext(ctx context.Context) ContainerObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerObjectMapOutput)
}

type ContainerObjectOutput struct{ *pulumi.OutputState }

func (ContainerObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerObject)(nil)).Elem()
}

func (o ContainerObjectOutput) ToContainerObjectOutput() ContainerObjectOutput {
	return o
}

func (o ContainerObjectOutput) ToContainerObjectOutputWithContext(ctx context.Context) ContainerObjectOutput {
	return o
}

// A unique (within an account) name for the container.
// The container name must be from 1 to 256 characters long and can start
// with any character and contain any pattern. Character set must be UTF-8.
// The container name cannot contain a slash (/) character because this
// character delimits the container and object name. For example, the path
// /v1/account/www/pages specifies the www container, not the www/pages container.
func (o ContainerObjectOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.ContainerName }).(pulumi.StringOutput)
}

// A string representing the content of the object. Conflicts with
// `source` and `copyFrom`.
func (o ContainerObjectOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// A string which specifies the override behavior for
// the browser. For example, this header might specify that the browser use a download
// program to save this file rather than show the file, which is the default.
func (o ContainerObjectOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.ContentDisposition }).(pulumi.StringOutput)
}

// A string representing the value of the Content-Encoding
// metadata.
func (o ContainerObjectOutput) ContentEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.ContentEncoding }).(pulumi.StringOutput)
}

// If the operation succeeds, this value is zero (0) or the
// length of informational or error text in the response body.
func (o ContainerObjectOutput) ContentLength() pulumi.IntOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.IntOutput { return v.ContentLength }).(pulumi.IntOutput)
}

// A string which sets the MIME type for the object.
func (o ContainerObjectOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// A string representing the name of an object
// used to create the new object by copying the `copyFrom` object. The value is in form
// {container}/{object}. You must UTF-8-encode and then URL-encode the names of the
// container and object before you include them in the header. Conflicts with `source` and
// `content`.
func (o ContainerObjectOutput) CopyFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringPtrOutput { return v.CopyFrom }).(pulumi.StringPtrOutput)
}

// The date and time the system responded to the request, using the preferred
// format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The
// time is always in UTC.
func (o ContainerObjectOutput) Date() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.Date }).(pulumi.StringOutput)
}

// An integer representing the number of seconds after which the
// system removes the object. Internally, the Object Storage system stores this value in
// the X-Delete-At metadata item.
func (o ContainerObjectOutput) DeleteAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.IntPtrOutput { return v.DeleteAfter }).(pulumi.IntPtrOutput)
}

// An string representing the date when the system removes the object.
// For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.
func (o ContainerObjectOutput) DeleteAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.DeleteAt }).(pulumi.StringOutput)
}

// If set to true, Object Storage guesses the content
// type based on the file extension and ignores the value sent in the Content-Type
// header, if present.
func (o ContainerObjectOutput) DetectContentType() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.BoolPtrOutput { return v.DetectContentType }).(pulumi.BoolPtrOutput)
}

// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
func (o ContainerObjectOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The date and time when the object was last modified. The date and time
// stamp format is ISO 8601:
// CCYY-MM-DDThh:mm:ss±hh:mm
// For example, 2015-08-27T09:49:58-05:00.
// The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous
// example, the offset value is -05:00.
func (o ContainerObjectOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o ContainerObjectOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.MapOutput { return v.Metadata }).(pulumi.MapOutput)
}

// A unique name for the object.
func (o ContainerObjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A string set to specify that this is a dynamic large
// object manifest object. The value is the container and object name prefix of the
// segment objects in the form container/prefix. You must UTF-8-encode and then
// URL-encode the names of the container and prefix before you include them in this
// header.
func (o ContainerObjectOutput) ObjectManifest() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.ObjectManifest }).(pulumi.StringOutput)
}

// The region in which to create the container. If
// omitted, the `region` argument of the provider is used. Changing this
// creates a new container.
func (o ContainerObjectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A string representing the local path of a file which will be used
// as the object's content. Conflicts with `source` and `copyFrom`.
func (o ContainerObjectOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// A unique transaction ID for this request. Your service provider might
// need this value if you report a problem.
func (o ContainerObjectOutput) TransId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerObject) pulumi.StringOutput { return v.TransId }).(pulumi.StringOutput)
}

type ContainerObjectArrayOutput struct{ *pulumi.OutputState }

func (ContainerObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerObject)(nil)).Elem()
}

func (o ContainerObjectArrayOutput) ToContainerObjectArrayOutput() ContainerObjectArrayOutput {
	return o
}

func (o ContainerObjectArrayOutput) ToContainerObjectArrayOutputWithContext(ctx context.Context) ContainerObjectArrayOutput {
	return o
}

func (o ContainerObjectArrayOutput) Index(i pulumi.IntInput) ContainerObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerObject {
		return vs[0].([]*ContainerObject)[vs[1].(int)]
	}).(ContainerObjectOutput)
}

type ContainerObjectMapOutput struct{ *pulumi.OutputState }

func (ContainerObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerObject)(nil)).Elem()
}

func (o ContainerObjectMapOutput) ToContainerObjectMapOutput() ContainerObjectMapOutput {
	return o
}

func (o ContainerObjectMapOutput) ToContainerObjectMapOutputWithContext(ctx context.Context) ContainerObjectMapOutput {
	return o
}

func (o ContainerObjectMapOutput) MapIndex(k pulumi.StringInput) ContainerObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerObject {
		return vs[0].(map[string]*ContainerObject)[vs[1].(string)]
	}).(ContainerObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerObjectInput)(nil)).Elem(), &ContainerObject{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerObjectArrayInput)(nil)).Elem(), ContainerObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerObjectMapInput)(nil)).Elem(), ContainerObjectMap{})
	pulumi.RegisterOutputType(ContainerObjectOutput{})
	pulumi.RegisterOutputType(ContainerObjectArrayOutput{})
	pulumi.RegisterOutputType(ContainerObjectMapOutput{})
}
