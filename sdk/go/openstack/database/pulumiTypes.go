// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package database

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ConfigurationConfiguration struct {
	// Configuration parameter name. Changing this creates a new resource.
	Name string `pulumi:"name"`
	// Configuration parameter value. Changing this creates a new resource.
	Value string `pulumi:"value"`
}

type ConfigurationConfigurationInput interface {
	pulumi.Input

	ToConfigurationConfigurationOutput() ConfigurationConfigurationOutput
	ToConfigurationConfigurationOutputWithContext(context.Context) ConfigurationConfigurationOutput
}

type ConfigurationConfigurationArgs struct {
	// Configuration parameter name. Changing this creates a new resource.
	Name pulumi.StringInput `pulumi:"name"`
	// Configuration parameter value. Changing this creates a new resource.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ConfigurationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationConfiguration)(nil)).Elem()
}

func (i ConfigurationConfigurationArgs) ToConfigurationConfigurationOutput() ConfigurationConfigurationOutput {
	return i.ToConfigurationConfigurationOutputWithContext(context.Background())
}

func (i ConfigurationConfigurationArgs) ToConfigurationConfigurationOutputWithContext(ctx context.Context) ConfigurationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationConfigurationOutput)
}

type ConfigurationConfigurationArrayInput interface {
	pulumi.Input

	ToConfigurationConfigurationArrayOutput() ConfigurationConfigurationArrayOutput
	ToConfigurationConfigurationArrayOutputWithContext(context.Context) ConfigurationConfigurationArrayOutput
}

type ConfigurationConfigurationArray []ConfigurationConfigurationInput

func (ConfigurationConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationConfiguration)(nil)).Elem()
}

func (i ConfigurationConfigurationArray) ToConfigurationConfigurationArrayOutput() ConfigurationConfigurationArrayOutput {
	return i.ToConfigurationConfigurationArrayOutputWithContext(context.Background())
}

func (i ConfigurationConfigurationArray) ToConfigurationConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationConfigurationArrayOutput)
}

type ConfigurationConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationConfiguration)(nil)).Elem()
}

func (o ConfigurationConfigurationOutput) ToConfigurationConfigurationOutput() ConfigurationConfigurationOutput {
	return o
}

func (o ConfigurationConfigurationOutput) ToConfigurationConfigurationOutputWithContext(ctx context.Context) ConfigurationConfigurationOutput {
	return o
}

// Configuration parameter name. Changing this creates a new resource.
func (o ConfigurationConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

// Configuration parameter value. Changing this creates a new resource.
func (o ConfigurationConfigurationOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationConfiguration) string { return v.Value }).(pulumi.StringOutput)
}

type ConfigurationConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationConfiguration)(nil)).Elem()
}

func (o ConfigurationConfigurationArrayOutput) ToConfigurationConfigurationArrayOutput() ConfigurationConfigurationArrayOutput {
	return o
}

func (o ConfigurationConfigurationArrayOutput) ToConfigurationConfigurationArrayOutputWithContext(ctx context.Context) ConfigurationConfigurationArrayOutput {
	return o
}

func (o ConfigurationConfigurationArrayOutput) Index(i pulumi.IntInput) ConfigurationConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationConfiguration {
		return vs[0].([]ConfigurationConfiguration)[vs[1].(int)]
	}).(ConfigurationConfigurationOutput)
}

type ConfigurationDatastore struct {
	// Database engine type to be used with this configuration. Changing this creates a new resource.
	Type string `pulumi:"type"`
	// Version of database engine type to be used with this configuration. Changing this creates a new resource.
	Version string `pulumi:"version"`
}

type ConfigurationDatastoreInput interface {
	pulumi.Input

	ToConfigurationDatastoreOutput() ConfigurationDatastoreOutput
	ToConfigurationDatastoreOutputWithContext(context.Context) ConfigurationDatastoreOutput
}

type ConfigurationDatastoreArgs struct {
	// Database engine type to be used with this configuration. Changing this creates a new resource.
	Type pulumi.StringInput `pulumi:"type"`
	// Version of database engine type to be used with this configuration. Changing this creates a new resource.
	Version pulumi.StringInput `pulumi:"version"`
}

func (ConfigurationDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationDatastore)(nil)).Elem()
}

func (i ConfigurationDatastoreArgs) ToConfigurationDatastoreOutput() ConfigurationDatastoreOutput {
	return i.ToConfigurationDatastoreOutputWithContext(context.Background())
}

func (i ConfigurationDatastoreArgs) ToConfigurationDatastoreOutputWithContext(ctx context.Context) ConfigurationDatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationDatastoreOutput)
}

func (i ConfigurationDatastoreArgs) ToConfigurationDatastorePtrOutput() ConfigurationDatastorePtrOutput {
	return i.ToConfigurationDatastorePtrOutputWithContext(context.Background())
}

func (i ConfigurationDatastoreArgs) ToConfigurationDatastorePtrOutputWithContext(ctx context.Context) ConfigurationDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationDatastoreOutput).ToConfigurationDatastorePtrOutputWithContext(ctx)
}

type ConfigurationDatastorePtrInput interface {
	pulumi.Input

	ToConfigurationDatastorePtrOutput() ConfigurationDatastorePtrOutput
	ToConfigurationDatastorePtrOutputWithContext(context.Context) ConfigurationDatastorePtrOutput
}

type configurationDatastorePtrType ConfigurationDatastoreArgs

func ConfigurationDatastorePtr(v *ConfigurationDatastoreArgs) ConfigurationDatastorePtrInput {
	return (*configurationDatastorePtrType)(v)
}

func (*configurationDatastorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationDatastore)(nil)).Elem()
}

func (i *configurationDatastorePtrType) ToConfigurationDatastorePtrOutput() ConfigurationDatastorePtrOutput {
	return i.ToConfigurationDatastorePtrOutputWithContext(context.Background())
}

func (i *configurationDatastorePtrType) ToConfigurationDatastorePtrOutputWithContext(ctx context.Context) ConfigurationDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationDatastorePtrOutput)
}

type ConfigurationDatastoreOutput struct{ *pulumi.OutputState }

func (ConfigurationDatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationDatastore)(nil)).Elem()
}

func (o ConfigurationDatastoreOutput) ToConfigurationDatastoreOutput() ConfigurationDatastoreOutput {
	return o
}

func (o ConfigurationDatastoreOutput) ToConfigurationDatastoreOutputWithContext(ctx context.Context) ConfigurationDatastoreOutput {
	return o
}

func (o ConfigurationDatastoreOutput) ToConfigurationDatastorePtrOutput() ConfigurationDatastorePtrOutput {
	return o.ToConfigurationDatastorePtrOutputWithContext(context.Background())
}

func (o ConfigurationDatastoreOutput) ToConfigurationDatastorePtrOutputWithContext(ctx context.Context) ConfigurationDatastorePtrOutput {
	return o.ApplyT(func(v ConfigurationDatastore) *ConfigurationDatastore {
		return &v
	}).(ConfigurationDatastorePtrOutput)
}

// Database engine type to be used with this configuration. Changing this creates a new resource.
func (o ConfigurationDatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationDatastore) string { return v.Type }).(pulumi.StringOutput)
}

// Version of database engine type to be used with this configuration. Changing this creates a new resource.
func (o ConfigurationDatastoreOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationDatastore) string { return v.Version }).(pulumi.StringOutput)
}

type ConfigurationDatastorePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationDatastorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationDatastore)(nil)).Elem()
}

func (o ConfigurationDatastorePtrOutput) ToConfigurationDatastorePtrOutput() ConfigurationDatastorePtrOutput {
	return o
}

func (o ConfigurationDatastorePtrOutput) ToConfigurationDatastorePtrOutputWithContext(ctx context.Context) ConfigurationDatastorePtrOutput {
	return o
}

func (o ConfigurationDatastorePtrOutput) Elem() ConfigurationDatastoreOutput {
	return o.ApplyT(func(v *ConfigurationDatastore) ConfigurationDatastore { return *v }).(ConfigurationDatastoreOutput)
}

// Database engine type to be used with this configuration. Changing this creates a new resource.
func (o ConfigurationDatastorePtrOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationDatastore) string { return v.Type }).(pulumi.StringOutput)
}

// Version of database engine type to be used with this configuration. Changing this creates a new resource.
func (o ConfigurationDatastorePtrOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationDatastore) string { return v.Version }).(pulumi.StringOutput)
}

type InstanceDatabase struct {
	// Database character set. Changing this creates a
	// new instance.
	Charset *string `pulumi:"charset"`
	// Database collation. Changing this creates a new instance.
	Collate *string `pulumi:"collate"`
	// Database to be created on new instance. Changing this creates a
	// new instance.
	Name string `pulumi:"name"`
}

type InstanceDatabaseInput interface {
	pulumi.Input

	ToInstanceDatabaseOutput() InstanceDatabaseOutput
	ToInstanceDatabaseOutputWithContext(context.Context) InstanceDatabaseOutput
}

type InstanceDatabaseArgs struct {
	// Database character set. Changing this creates a
	// new instance.
	Charset pulumi.StringPtrInput `pulumi:"charset"`
	// Database collation. Changing this creates a new instance.
	Collate pulumi.StringPtrInput `pulumi:"collate"`
	// Database to be created on new instance. Changing this creates a
	// new instance.
	Name pulumi.StringInput `pulumi:"name"`
}

func (InstanceDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDatabase)(nil)).Elem()
}

func (i InstanceDatabaseArgs) ToInstanceDatabaseOutput() InstanceDatabaseOutput {
	return i.ToInstanceDatabaseOutputWithContext(context.Background())
}

func (i InstanceDatabaseArgs) ToInstanceDatabaseOutputWithContext(ctx context.Context) InstanceDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDatabaseOutput)
}

type InstanceDatabaseArrayInput interface {
	pulumi.Input

	ToInstanceDatabaseArrayOutput() InstanceDatabaseArrayOutput
	ToInstanceDatabaseArrayOutputWithContext(context.Context) InstanceDatabaseArrayOutput
}

type InstanceDatabaseArray []InstanceDatabaseInput

func (InstanceDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceDatabase)(nil)).Elem()
}

func (i InstanceDatabaseArray) ToInstanceDatabaseArrayOutput() InstanceDatabaseArrayOutput {
	return i.ToInstanceDatabaseArrayOutputWithContext(context.Background())
}

func (i InstanceDatabaseArray) ToInstanceDatabaseArrayOutputWithContext(ctx context.Context) InstanceDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDatabaseArrayOutput)
}

type InstanceDatabaseOutput struct{ *pulumi.OutputState }

func (InstanceDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDatabase)(nil)).Elem()
}

func (o InstanceDatabaseOutput) ToInstanceDatabaseOutput() InstanceDatabaseOutput {
	return o
}

func (o InstanceDatabaseOutput) ToInstanceDatabaseOutputWithContext(ctx context.Context) InstanceDatabaseOutput {
	return o
}

// Database character set. Changing this creates a
// new instance.
func (o InstanceDatabaseOutput) Charset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceDatabase) *string { return v.Charset }).(pulumi.StringPtrOutput)
}

// Database collation. Changing this creates a new instance.
func (o InstanceDatabaseOutput) Collate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceDatabase) *string { return v.Collate }).(pulumi.StringPtrOutput)
}

// Database to be created on new instance. Changing this creates a
// new instance.
func (o InstanceDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceDatabase) string { return v.Name }).(pulumi.StringOutput)
}

type InstanceDatabaseArrayOutput struct{ *pulumi.OutputState }

func (InstanceDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceDatabase)(nil)).Elem()
}

func (o InstanceDatabaseArrayOutput) ToInstanceDatabaseArrayOutput() InstanceDatabaseArrayOutput {
	return o
}

func (o InstanceDatabaseArrayOutput) ToInstanceDatabaseArrayOutputWithContext(ctx context.Context) InstanceDatabaseArrayOutput {
	return o
}

func (o InstanceDatabaseArrayOutput) Index(i pulumi.IntInput) InstanceDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceDatabase {
		return vs[0].([]InstanceDatabase)[vs[1].(int)]
	}).(InstanceDatabaseOutput)
}

type InstanceDatastore struct {
	// Database engine type to be used in new instance. Changing this
	// creates a new instance.
	Type string `pulumi:"type"`
	// Version of database engine type to be used in new instance.
	// Changing this creates a new instance.
	Version string `pulumi:"version"`
}

type InstanceDatastoreInput interface {
	pulumi.Input

	ToInstanceDatastoreOutput() InstanceDatastoreOutput
	ToInstanceDatastoreOutputWithContext(context.Context) InstanceDatastoreOutput
}

type InstanceDatastoreArgs struct {
	// Database engine type to be used in new instance. Changing this
	// creates a new instance.
	Type pulumi.StringInput `pulumi:"type"`
	// Version of database engine type to be used in new instance.
	// Changing this creates a new instance.
	Version pulumi.StringInput `pulumi:"version"`
}

func (InstanceDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDatastore)(nil)).Elem()
}

func (i InstanceDatastoreArgs) ToInstanceDatastoreOutput() InstanceDatastoreOutput {
	return i.ToInstanceDatastoreOutputWithContext(context.Background())
}

func (i InstanceDatastoreArgs) ToInstanceDatastoreOutputWithContext(ctx context.Context) InstanceDatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDatastoreOutput)
}

func (i InstanceDatastoreArgs) ToInstanceDatastorePtrOutput() InstanceDatastorePtrOutput {
	return i.ToInstanceDatastorePtrOutputWithContext(context.Background())
}

func (i InstanceDatastoreArgs) ToInstanceDatastorePtrOutputWithContext(ctx context.Context) InstanceDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDatastoreOutput).ToInstanceDatastorePtrOutputWithContext(ctx)
}

type InstanceDatastorePtrInput interface {
	pulumi.Input

	ToInstanceDatastorePtrOutput() InstanceDatastorePtrOutput
	ToInstanceDatastorePtrOutputWithContext(context.Context) InstanceDatastorePtrOutput
}

type instanceDatastorePtrType InstanceDatastoreArgs

func InstanceDatastorePtr(v *InstanceDatastoreArgs) InstanceDatastorePtrInput {
	return (*instanceDatastorePtrType)(v)
}

func (*instanceDatastorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceDatastore)(nil)).Elem()
}

func (i *instanceDatastorePtrType) ToInstanceDatastorePtrOutput() InstanceDatastorePtrOutput {
	return i.ToInstanceDatastorePtrOutputWithContext(context.Background())
}

func (i *instanceDatastorePtrType) ToInstanceDatastorePtrOutputWithContext(ctx context.Context) InstanceDatastorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDatastorePtrOutput)
}

type InstanceDatastoreOutput struct{ *pulumi.OutputState }

func (InstanceDatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDatastore)(nil)).Elem()
}

func (o InstanceDatastoreOutput) ToInstanceDatastoreOutput() InstanceDatastoreOutput {
	return o
}

func (o InstanceDatastoreOutput) ToInstanceDatastoreOutputWithContext(ctx context.Context) InstanceDatastoreOutput {
	return o
}

func (o InstanceDatastoreOutput) ToInstanceDatastorePtrOutput() InstanceDatastorePtrOutput {
	return o.ToInstanceDatastorePtrOutputWithContext(context.Background())
}

func (o InstanceDatastoreOutput) ToInstanceDatastorePtrOutputWithContext(ctx context.Context) InstanceDatastorePtrOutput {
	return o.ApplyT(func(v InstanceDatastore) *InstanceDatastore {
		return &v
	}).(InstanceDatastorePtrOutput)
}

// Database engine type to be used in new instance. Changing this
// creates a new instance.
func (o InstanceDatastoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceDatastore) string { return v.Type }).(pulumi.StringOutput)
}

// Version of database engine type to be used in new instance.
// Changing this creates a new instance.
func (o InstanceDatastoreOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceDatastore) string { return v.Version }).(pulumi.StringOutput)
}

type InstanceDatastorePtrOutput struct{ *pulumi.OutputState }

func (InstanceDatastorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceDatastore)(nil)).Elem()
}

func (o InstanceDatastorePtrOutput) ToInstanceDatastorePtrOutput() InstanceDatastorePtrOutput {
	return o
}

func (o InstanceDatastorePtrOutput) ToInstanceDatastorePtrOutputWithContext(ctx context.Context) InstanceDatastorePtrOutput {
	return o
}

func (o InstanceDatastorePtrOutput) Elem() InstanceDatastoreOutput {
	return o.ApplyT(func(v *InstanceDatastore) InstanceDatastore { return *v }).(InstanceDatastoreOutput)
}

// Database engine type to be used in new instance. Changing this
// creates a new instance.
func (o InstanceDatastorePtrOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceDatastore) string { return v.Type }).(pulumi.StringOutput)
}

// Version of database engine type to be used in new instance.
// Changing this creates a new instance.
func (o InstanceDatastorePtrOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceDatastore) string { return v.Version }).(pulumi.StringOutput)
}

type InstanceNetwork struct {
	// Specifies a fixed IPv4 address to be used on this
	// network. Changing this creates a new instance.
	FixedIpV4 *string `pulumi:"fixedIpV4"`
	// Specifies a fixed IPv6 address to be used on this
	// network. Changing this creates a new instance.
	FixedIpV6 *string `pulumi:"fixedIpV6"`
	// The port UUID of a
	// network to attach to the instance. Changing this creates a new instance.
	Port *string `pulumi:"port"`
	// The network UUID to
	// attach to the instance. Changing this creates a new instance.
	Uuid *string `pulumi:"uuid"`
}

type InstanceNetworkInput interface {
	pulumi.Input

	ToInstanceNetworkOutput() InstanceNetworkOutput
	ToInstanceNetworkOutputWithContext(context.Context) InstanceNetworkOutput
}

type InstanceNetworkArgs struct {
	// Specifies a fixed IPv4 address to be used on this
	// network. Changing this creates a new instance.
	FixedIpV4 pulumi.StringPtrInput `pulumi:"fixedIpV4"`
	// Specifies a fixed IPv6 address to be used on this
	// network. Changing this creates a new instance.
	FixedIpV6 pulumi.StringPtrInput `pulumi:"fixedIpV6"`
	// The port UUID of a
	// network to attach to the instance. Changing this creates a new instance.
	Port pulumi.StringPtrInput `pulumi:"port"`
	// The network UUID to
	// attach to the instance. Changing this creates a new instance.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
}

func (InstanceNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNetwork)(nil)).Elem()
}

func (i InstanceNetworkArgs) ToInstanceNetworkOutput() InstanceNetworkOutput {
	return i.ToInstanceNetworkOutputWithContext(context.Background())
}

func (i InstanceNetworkArgs) ToInstanceNetworkOutputWithContext(ctx context.Context) InstanceNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNetworkOutput)
}

type InstanceNetworkArrayInput interface {
	pulumi.Input

	ToInstanceNetworkArrayOutput() InstanceNetworkArrayOutput
	ToInstanceNetworkArrayOutputWithContext(context.Context) InstanceNetworkArrayOutput
}

type InstanceNetworkArray []InstanceNetworkInput

func (InstanceNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceNetwork)(nil)).Elem()
}

func (i InstanceNetworkArray) ToInstanceNetworkArrayOutput() InstanceNetworkArrayOutput {
	return i.ToInstanceNetworkArrayOutputWithContext(context.Background())
}

func (i InstanceNetworkArray) ToInstanceNetworkArrayOutputWithContext(ctx context.Context) InstanceNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNetworkArrayOutput)
}

type InstanceNetworkOutput struct{ *pulumi.OutputState }

func (InstanceNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNetwork)(nil)).Elem()
}

func (o InstanceNetworkOutput) ToInstanceNetworkOutput() InstanceNetworkOutput {
	return o
}

func (o InstanceNetworkOutput) ToInstanceNetworkOutputWithContext(ctx context.Context) InstanceNetworkOutput {
	return o
}

// Specifies a fixed IPv4 address to be used on this
// network. Changing this creates a new instance.
func (o InstanceNetworkOutput) FixedIpV4() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceNetwork) *string { return v.FixedIpV4 }).(pulumi.StringPtrOutput)
}

// Specifies a fixed IPv6 address to be used on this
// network. Changing this creates a new instance.
func (o InstanceNetworkOutput) FixedIpV6() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceNetwork) *string { return v.FixedIpV6 }).(pulumi.StringPtrOutput)
}

// The port UUID of a
// network to attach to the instance. Changing this creates a new instance.
func (o InstanceNetworkOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceNetwork) *string { return v.Port }).(pulumi.StringPtrOutput)
}

// The network UUID to
// attach to the instance. Changing this creates a new instance.
func (o InstanceNetworkOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceNetwork) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

type InstanceNetworkArrayOutput struct{ *pulumi.OutputState }

func (InstanceNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceNetwork)(nil)).Elem()
}

func (o InstanceNetworkArrayOutput) ToInstanceNetworkArrayOutput() InstanceNetworkArrayOutput {
	return o
}

func (o InstanceNetworkArrayOutput) ToInstanceNetworkArrayOutputWithContext(ctx context.Context) InstanceNetworkArrayOutput {
	return o
}

func (o InstanceNetworkArrayOutput) Index(i pulumi.IntInput) InstanceNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceNetwork {
		return vs[0].([]InstanceNetwork)[vs[1].(int)]
	}).(InstanceNetworkOutput)
}

type InstanceUser struct {
	// A list of databases that user will have access to. If not specified,
	// user has access to all databases on th einstance. Changing this creates a new instance.
	Databases []string `pulumi:"databases"`
	// An ip address or % sign indicating what ip addresses can connect with
	// this user credentials. Changing this creates a new instance.
	Host *string `pulumi:"host"`
	// Database to be created on new instance. Changing this creates a
	// new instance.
	Name string `pulumi:"name"`
	// User's password. Changing this creates a
	// new instance.
	Password *string `pulumi:"password"`
}

type InstanceUserInput interface {
	pulumi.Input

	ToInstanceUserOutput() InstanceUserOutput
	ToInstanceUserOutputWithContext(context.Context) InstanceUserOutput
}

type InstanceUserArgs struct {
	// A list of databases that user will have access to. If not specified,
	// user has access to all databases on th einstance. Changing this creates a new instance.
	Databases pulumi.StringArrayInput `pulumi:"databases"`
	// An ip address or % sign indicating what ip addresses can connect with
	// this user credentials. Changing this creates a new instance.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// Database to be created on new instance. Changing this creates a
	// new instance.
	Name pulumi.StringInput `pulumi:"name"`
	// User's password. Changing this creates a
	// new instance.
	Password pulumi.StringPtrInput `pulumi:"password"`
}

func (InstanceUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceUser)(nil)).Elem()
}

func (i InstanceUserArgs) ToInstanceUserOutput() InstanceUserOutput {
	return i.ToInstanceUserOutputWithContext(context.Background())
}

func (i InstanceUserArgs) ToInstanceUserOutputWithContext(ctx context.Context) InstanceUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceUserOutput)
}

type InstanceUserArrayInput interface {
	pulumi.Input

	ToInstanceUserArrayOutput() InstanceUserArrayOutput
	ToInstanceUserArrayOutputWithContext(context.Context) InstanceUserArrayOutput
}

type InstanceUserArray []InstanceUserInput

func (InstanceUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceUser)(nil)).Elem()
}

func (i InstanceUserArray) ToInstanceUserArrayOutput() InstanceUserArrayOutput {
	return i.ToInstanceUserArrayOutputWithContext(context.Background())
}

func (i InstanceUserArray) ToInstanceUserArrayOutputWithContext(ctx context.Context) InstanceUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceUserArrayOutput)
}

type InstanceUserOutput struct{ *pulumi.OutputState }

func (InstanceUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceUser)(nil)).Elem()
}

func (o InstanceUserOutput) ToInstanceUserOutput() InstanceUserOutput {
	return o
}

func (o InstanceUserOutput) ToInstanceUserOutputWithContext(ctx context.Context) InstanceUserOutput {
	return o
}

// A list of databases that user will have access to. If not specified,
// user has access to all databases on th einstance. Changing this creates a new instance.
func (o InstanceUserOutput) Databases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceUser) []string { return v.Databases }).(pulumi.StringArrayOutput)
}

// An ip address or % sign indicating what ip addresses can connect with
// this user credentials. Changing this creates a new instance.
func (o InstanceUserOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceUser) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// Database to be created on new instance. Changing this creates a
// new instance.
func (o InstanceUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceUser) string { return v.Name }).(pulumi.StringOutput)
}

// User's password. Changing this creates a
// new instance.
func (o InstanceUserOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceUser) *string { return v.Password }).(pulumi.StringPtrOutput)
}

type InstanceUserArrayOutput struct{ *pulumi.OutputState }

func (InstanceUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceUser)(nil)).Elem()
}

func (o InstanceUserArrayOutput) ToInstanceUserArrayOutput() InstanceUserArrayOutput {
	return o
}

func (o InstanceUserArrayOutput) ToInstanceUserArrayOutputWithContext(ctx context.Context) InstanceUserArrayOutput {
	return o
}

func (o InstanceUserArrayOutput) Index(i pulumi.IntInput) InstanceUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceUser {
		return vs[0].([]InstanceUser)[vs[1].(int)]
	}).(InstanceUserOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationConfigurationOutput{})
	pulumi.RegisterOutputType(ConfigurationConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationDatastoreOutput{})
	pulumi.RegisterOutputType(ConfigurationDatastorePtrOutput{})
	pulumi.RegisterOutputType(InstanceDatabaseOutput{})
	pulumi.RegisterOutputType(InstanceDatabaseArrayOutput{})
	pulumi.RegisterOutputType(InstanceDatastoreOutput{})
	pulumi.RegisterOutputType(InstanceDatastorePtrOutput{})
	pulumi.RegisterOutputType(InstanceNetworkOutput{})
	pulumi.RegisterOutputType(InstanceNetworkArrayOutput{})
	pulumi.RegisterOutputType(InstanceUserOutput{})
	pulumi.RegisterOutputType(InstanceUserArrayOutput{})
}
