// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package objectstorage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this resource to generate an OpenStack Object Storage temporary URL.
// 
// The temporary URL will be valid for as long as TTL is set to (in seconds).
// Once the URL has expired, it will no longer be valid, but the resource
// will remain in place. If you wish to automatically regenerate a URL, set
// the `regenerate` argument to `true`. This will create a new resource with
// a new ID and URL.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/objectstorage_tempurl_v1.html.markdown.
type TempUrl struct {
	pulumi.CustomResourceState

	// The container name the object belongs to.
	Container pulumi.StringOutput `pulumi:"container"`
	// The method allowed when accessing this URL.
	// Valid values are `GET`, and `POST`. Default is `GET`.
	Method pulumi.StringPtrOutput `pulumi:"method"`
	// The object name the tempurl is for.
	Object pulumi.StringOutput `pulumi:"object"`
	// Whether to automatically regenerate the URL when
	// it has expired. If set to true, this will create a new resource with a new
	// ID and new URL. Defaults to false.
	Regenerate pulumi.BoolPtrOutput `pulumi:"regenerate"`
	// The region the tempurl is located in.
	Region pulumi.StringOutput `pulumi:"region"`
	Split pulumi.StringPtrOutput `pulumi:"split"`
	// The TTL, in seconds, for the URL. For how long it should
	// be valid.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
	// The URL
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewTempUrl registers a new resource with the given unique name, arguments, and options.
func NewTempUrl(ctx *pulumi.Context,
	name string, args *TempUrlArgs, opts ...pulumi.ResourceOption) (*TempUrl, error) {
	if args == nil || args.Container == nil {
		return nil, errors.New("missing required argument 'Container'")
	}
	if args == nil || args.Object == nil {
		return nil, errors.New("missing required argument 'Object'")
	}
	if args == nil || args.Ttl == nil {
		return nil, errors.New("missing required argument 'Ttl'")
	}
	if args == nil {
		args = &TempUrlArgs{}
	}
	var resource TempUrl
	err := ctx.RegisterResource("openstack:objectstorage/tempUrl:TempUrl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTempUrl gets an existing TempUrl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTempUrl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TempUrlState, opts ...pulumi.ResourceOption) (*TempUrl, error) {
	var resource TempUrl
	err := ctx.ReadResource("openstack:objectstorage/tempUrl:TempUrl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TempUrl resources.
type tempUrlState struct {
	// The container name the object belongs to.
	Container *string `pulumi:"container"`
	// The method allowed when accessing this URL.
	// Valid values are `GET`, and `POST`. Default is `GET`.
	Method *string `pulumi:"method"`
	// The object name the tempurl is for.
	Object *string `pulumi:"object"`
	// Whether to automatically regenerate the URL when
	// it has expired. If set to true, this will create a new resource with a new
	// ID and new URL. Defaults to false.
	Regenerate *bool `pulumi:"regenerate"`
	// The region the tempurl is located in.
	Region *string `pulumi:"region"`
	Split *string `pulumi:"split"`
	// The TTL, in seconds, for the URL. For how long it should
	// be valid.
	Ttl *int `pulumi:"ttl"`
	// The URL
	Url *string `pulumi:"url"`
}

type TempUrlState struct {
	// The container name the object belongs to.
	Container pulumi.StringPtrInput
	// The method allowed when accessing this URL.
	// Valid values are `GET`, and `POST`. Default is `GET`.
	Method pulumi.StringPtrInput
	// The object name the tempurl is for.
	Object pulumi.StringPtrInput
	// Whether to automatically regenerate the URL when
	// it has expired. If set to true, this will create a new resource with a new
	// ID and new URL. Defaults to false.
	Regenerate pulumi.BoolPtrInput
	// The region the tempurl is located in.
	Region pulumi.StringPtrInput
	Split pulumi.StringPtrInput
	// The TTL, in seconds, for the URL. For how long it should
	// be valid.
	Ttl pulumi.IntPtrInput
	// The URL
	Url pulumi.StringPtrInput
}

func (TempUrlState) ElementType() reflect.Type {
	return reflect.TypeOf((*tempUrlState)(nil)).Elem()
}

type tempUrlArgs struct {
	// The container name the object belongs to.
	Container string `pulumi:"container"`
	// The method allowed when accessing this URL.
	// Valid values are `GET`, and `POST`. Default is `GET`.
	Method *string `pulumi:"method"`
	// The object name the tempurl is for.
	Object string `pulumi:"object"`
	// Whether to automatically regenerate the URL when
	// it has expired. If set to true, this will create a new resource with a new
	// ID and new URL. Defaults to false.
	Regenerate *bool `pulumi:"regenerate"`
	// The region the tempurl is located in.
	Region *string `pulumi:"region"`
	Split *string `pulumi:"split"`
	// The TTL, in seconds, for the URL. For how long it should
	// be valid.
	Ttl int `pulumi:"ttl"`
}

// The set of arguments for constructing a TempUrl resource.
type TempUrlArgs struct {
	// The container name the object belongs to.
	Container pulumi.StringInput
	// The method allowed when accessing this URL.
	// Valid values are `GET`, and `POST`. Default is `GET`.
	Method pulumi.StringPtrInput
	// The object name the tempurl is for.
	Object pulumi.StringInput
	// Whether to automatically regenerate the URL when
	// it has expired. If set to true, this will create a new resource with a new
	// ID and new URL. Defaults to false.
	Regenerate pulumi.BoolPtrInput
	// The region the tempurl is located in.
	Region pulumi.StringPtrInput
	Split pulumi.StringPtrInput
	// The TTL, in seconds, for the URL. For how long it should
	// be valid.
	Ttl pulumi.IntInput
}

func (TempUrlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tempUrlArgs)(nil)).Elem()
}

