// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package orchestration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 stack resource within OpenStack.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-openstack/sdk/v3/go/openstack/orchestration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := orchestration.NewStackV1(ctx, "stack1", &orchestration.StackV1Args{
// 			DisableRollback: pulumi.Bool(true),
// 			EnvironmentOpts: pulumi.AnyMap{
// 				"Bin": pulumi.Any(fmt.Sprintf("%v%v", "\n", "\n")),
// 			},
// 			Parameters: pulumi.AnyMap{
// 				"length": pulumi.Any(4),
// 			},
// 			TemplateOpts: pulumi.AnyMap{
// 				"Bin": pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "heat_template_version: 2013-05-23\n", "parameters:\n", "  length:\n", "    type: number\n", "resources:\n", "  test_res:\n", "    type: OS::Heat::TestResource\n", "  random:\n", "    type: OS::Heat::RandomString\n", "    properties:\n", "      length: {get_param: length}\n", "\n")),
// 			},
// 			Timeout: pulumi.Int(30),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// stacks can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import openstack:orchestration/stackV1:StackV1 stack_1 ea257959-eeb1-4c10-8d33-26f0409a755d
// ```
type StackV1 struct {
	pulumi.CustomResourceState

	// A list of stack outputs.
	StackOutputs StackV1StackOutputArrayOutput `pulumi:"StackOutputs"`
	// List of stack capabilities for stack.
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The description of the stack resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	DisableRollback pulumi.BoolOutput `pulumi:"disableRollback"`
	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	EnvironmentOpts pulumi.MapOutput `pulumi:"environmentOpts"`
	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of notification topics for stack.
	NotificationTopics pulumi.StringArrayOutput `pulumi:"notificationTopics"`
	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// The region in which to create the stack. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new stack.
	Region pulumi.StringOutput `pulumi:"region"`
	// The status of the stack.
	Status pulumi.StringOutput `pulumi:"status"`
	// The reason for the current status of the stack.
	StatusReason pulumi.StringOutput `pulumi:"statusReason"`
	// A list of tags to assosciate with the Stack
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The description of the stack template.
	TemplateDescription pulumi.StringOutput `pulumi:"templateDescription"`
	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	TemplateOpts pulumi.MapOutput `pulumi:"templateOpts"`
	// The timeout for stack action in minutes.
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	UpdatedTime pulumi.StringOutput `pulumi:"updatedTime"`
}

// NewStackV1 registers a new resource with the given unique name, arguments, and options.
func NewStackV1(ctx *pulumi.Context,
	name string, args *StackV1Args, opts ...pulumi.ResourceOption) (*StackV1, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TemplateOpts == nil {
		return nil, errors.New("invalid value for required argument 'TemplateOpts'")
	}
	var resource StackV1
	err := ctx.RegisterResource("openstack:orchestration/stackV1:StackV1", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStackV1 gets an existing StackV1 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStackV1(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackV1State, opts ...pulumi.ResourceOption) (*StackV1, error) {
	var resource StackV1
	err := ctx.ReadResource("openstack:orchestration/stackV1:StackV1", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StackV1 resources.
type stackV1State struct {
	// A list of stack outputs.
	StackOutputs []StackV1StackOutput `pulumi:"StackOutputs"`
	// List of stack capabilities for stack.
	Capabilities []string `pulumi:"capabilities"`
	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	CreationTime *string `pulumi:"creationTime"`
	// The description of the stack resource.
	Description *string `pulumi:"description"`
	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	DisableRollback *bool `pulumi:"disableRollback"`
	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	EnvironmentOpts map[string]interface{} `pulumi:"environmentOpts"`
	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	Name *string `pulumi:"name"`
	// List of notification topics for stack.
	NotificationTopics []string `pulumi:"notificationTopics"`
	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The region in which to create the stack. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new stack.
	Region *string `pulumi:"region"`
	// The status of the stack.
	Status *string `pulumi:"status"`
	// The reason for the current status of the stack.
	StatusReason *string `pulumi:"statusReason"`
	// A list of tags to assosciate with the Stack
	Tags []string `pulumi:"tags"`
	// The description of the stack template.
	TemplateDescription *string `pulumi:"templateDescription"`
	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	TemplateOpts map[string]interface{} `pulumi:"templateOpts"`
	// The timeout for stack action in minutes.
	Timeout *int `pulumi:"timeout"`
	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	UpdatedTime *string `pulumi:"updatedTime"`
}

type StackV1State struct {
	// A list of stack outputs.
	StackOutputs StackV1StackOutputArrayInput
	// List of stack capabilities for stack.
	Capabilities pulumi.StringArrayInput
	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	CreationTime pulumi.StringPtrInput
	// The description of the stack resource.
	Description pulumi.StringPtrInput
	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	DisableRollback pulumi.BoolPtrInput
	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	EnvironmentOpts pulumi.MapInput
	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	Name pulumi.StringPtrInput
	// List of notification topics for stack.
	NotificationTopics pulumi.StringArrayInput
	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	Parameters pulumi.MapInput
	// The region in which to create the stack. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new stack.
	Region pulumi.StringPtrInput
	// The status of the stack.
	Status pulumi.StringPtrInput
	// The reason for the current status of the stack.
	StatusReason pulumi.StringPtrInput
	// A list of tags to assosciate with the Stack
	Tags pulumi.StringArrayInput
	// The description of the stack template.
	TemplateDescription pulumi.StringPtrInput
	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	TemplateOpts pulumi.MapInput
	// The timeout for stack action in minutes.
	Timeout pulumi.IntPtrInput
	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	UpdatedTime pulumi.StringPtrInput
}

func (StackV1State) ElementType() reflect.Type {
	return reflect.TypeOf((*stackV1State)(nil)).Elem()
}

type stackV1Args struct {
	// A list of stack outputs.
	StackOutputs []StackV1StackOutput `pulumi:"StackOutputs"`
	// List of stack capabilities for stack.
	Capabilities []string `pulumi:"capabilities"`
	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	CreationTime *string `pulumi:"creationTime"`
	// The description of the stack resource.
	Description *string `pulumi:"description"`
	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	DisableRollback *bool `pulumi:"disableRollback"`
	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	EnvironmentOpts map[string]interface{} `pulumi:"environmentOpts"`
	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	Name *string `pulumi:"name"`
	// List of notification topics for stack.
	NotificationTopics []string `pulumi:"notificationTopics"`
	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The region in which to create the stack. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new stack.
	Region *string `pulumi:"region"`
	// The status of the stack.
	Status *string `pulumi:"status"`
	// The reason for the current status of the stack.
	StatusReason *string `pulumi:"statusReason"`
	// A list of tags to assosciate with the Stack
	Tags []string `pulumi:"tags"`
	// The description of the stack template.
	TemplateDescription *string `pulumi:"templateDescription"`
	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	TemplateOpts map[string]interface{} `pulumi:"templateOpts"`
	// The timeout for stack action in minutes.
	Timeout *int `pulumi:"timeout"`
	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	UpdatedTime *string `pulumi:"updatedTime"`
}

// The set of arguments for constructing a StackV1 resource.
type StackV1Args struct {
	// A list of stack outputs.
	StackOutputs StackV1StackOutputArrayInput
	// List of stack capabilities for stack.
	Capabilities pulumi.StringArrayInput
	// The date and time when the resource was created. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	CreationTime pulumi.StringPtrInput
	// The description of the stack resource.
	Description pulumi.StringPtrInput
	// Enables or disables deletion of all stack
	// resources when a stack creation fails. Default is true, meaning all
	// resources are not deleted when stack creation fails.
	DisableRollback pulumi.BoolPtrInput
	// Environment key/value pairs to associate with
	// the stack which contains details for the environment of the stack.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Environment Opts.
	EnvironmentOpts pulumi.MapInput
	// A unique name for the stack. It must start with an
	// alphabetic character. Changing this updates the stack's name.
	Name pulumi.StringPtrInput
	// List of notification topics for stack.
	NotificationTopics pulumi.StringArrayInput
	// User-defined key/value pairs as parameters to pass
	// to the template. Changing this updates the existing stack parameters.
	Parameters pulumi.MapInput
	// The region in which to create the stack. If
	// omitted, the `region` argument of the provider is used. Changing this
	// creates a new stack.
	Region pulumi.StringPtrInput
	// The status of the stack.
	Status pulumi.StringPtrInput
	// The reason for the current status of the stack.
	StatusReason pulumi.StringPtrInput
	// A list of tags to assosciate with the Stack
	Tags pulumi.StringArrayInput
	// The description of the stack template.
	TemplateDescription pulumi.StringPtrInput
	// Template key/value pairs to associate with the
	// stack which contains either the template file or url.
	// Allowed keys: Bin, URL, Files. Changing this updates the existing stack
	// Template Opts.
	TemplateOpts pulumi.MapInput
	// The timeout for stack action in minutes.
	Timeout pulumi.IntPtrInput
	// The date and time when the resource was updated. The date
	// and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm
	// For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included,
	// is the time zone as an offset from UTC.
	UpdatedTime pulumi.StringPtrInput
}

func (StackV1Args) ElementType() reflect.Type {
	return reflect.TypeOf((*stackV1Args)(nil)).Elem()
}

type StackV1Input interface {
	pulumi.Input

	ToStackV1Output() StackV1Output
	ToStackV1OutputWithContext(ctx context.Context) StackV1Output
}

func (*StackV1) ElementType() reflect.Type {
	return reflect.TypeOf((**StackV1)(nil)).Elem()
}

func (i *StackV1) ToStackV1Output() StackV1Output {
	return i.ToStackV1OutputWithContext(context.Background())
}

func (i *StackV1) ToStackV1OutputWithContext(ctx context.Context) StackV1Output {
	return pulumi.ToOutputWithContext(ctx, i).(StackV1Output)
}

// StackV1ArrayInput is an input type that accepts StackV1Array and StackV1ArrayOutput values.
// You can construct a concrete instance of `StackV1ArrayInput` via:
//
//          StackV1Array{ StackV1Args{...} }
type StackV1ArrayInput interface {
	pulumi.Input

	ToStackV1ArrayOutput() StackV1ArrayOutput
	ToStackV1ArrayOutputWithContext(context.Context) StackV1ArrayOutput
}

type StackV1Array []StackV1Input

func (StackV1Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackV1)(nil)).Elem()
}

func (i StackV1Array) ToStackV1ArrayOutput() StackV1ArrayOutput {
	return i.ToStackV1ArrayOutputWithContext(context.Background())
}

func (i StackV1Array) ToStackV1ArrayOutputWithContext(ctx context.Context) StackV1ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackV1ArrayOutput)
}

// StackV1MapInput is an input type that accepts StackV1Map and StackV1MapOutput values.
// You can construct a concrete instance of `StackV1MapInput` via:
//
//          StackV1Map{ "key": StackV1Args{...} }
type StackV1MapInput interface {
	pulumi.Input

	ToStackV1MapOutput() StackV1MapOutput
	ToStackV1MapOutputWithContext(context.Context) StackV1MapOutput
}

type StackV1Map map[string]StackV1Input

func (StackV1Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackV1)(nil)).Elem()
}

func (i StackV1Map) ToStackV1MapOutput() StackV1MapOutput {
	return i.ToStackV1MapOutputWithContext(context.Background())
}

func (i StackV1Map) ToStackV1MapOutputWithContext(ctx context.Context) StackV1MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StackV1MapOutput)
}

type StackV1Output struct{ *pulumi.OutputState }

func (StackV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((**StackV1)(nil)).Elem()
}

func (o StackV1Output) ToStackV1Output() StackV1Output {
	return o
}

func (o StackV1Output) ToStackV1OutputWithContext(ctx context.Context) StackV1Output {
	return o
}

type StackV1ArrayOutput struct{ *pulumi.OutputState }

func (StackV1ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StackV1)(nil)).Elem()
}

func (o StackV1ArrayOutput) ToStackV1ArrayOutput() StackV1ArrayOutput {
	return o
}

func (o StackV1ArrayOutput) ToStackV1ArrayOutputWithContext(ctx context.Context) StackV1ArrayOutput {
	return o
}

func (o StackV1ArrayOutput) Index(i pulumi.IntInput) StackV1Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StackV1 {
		return vs[0].([]*StackV1)[vs[1].(int)]
	}).(StackV1Output)
}

type StackV1MapOutput struct{ *pulumi.OutputState }

func (StackV1MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StackV1)(nil)).Elem()
}

func (o StackV1MapOutput) ToStackV1MapOutput() StackV1MapOutput {
	return o
}

func (o StackV1MapOutput) ToStackV1MapOutputWithContext(ctx context.Context) StackV1MapOutput {
	return o
}

func (o StackV1MapOutput) MapIndex(k pulumi.StringInput) StackV1Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StackV1 {
		return vs[0].(map[string]*StackV1)[vs[1].(string)]
	}).(StackV1Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StackV1Input)(nil)).Elem(), &StackV1{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackV1ArrayInput)(nil)).Elem(), StackV1Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*StackV1MapInput)(nil)).Elem(), StackV1Map{})
	pulumi.RegisterOutputType(StackV1Output{})
	pulumi.RegisterOutputType(StackV1ArrayOutput{})
	pulumi.RegisterOutputType(StackV1MapOutput{})
}
