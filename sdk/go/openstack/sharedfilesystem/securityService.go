// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to configure a security service.
//
// > **Note:** All arguments including the security service password will be
// stored in the raw state as plain-text. [Read more about sensitive data in
// state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// A security service stores configuration information for clients for
// authentication and authorization (AuthN/AuthZ). For example, a share server
// will be the client for an existing service such as LDAP, Kerberos, or
// Microsoft Active Directory.
//
// Minimum supported Manila microversion is 2.7.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/sharedfilesystem"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := sharedfilesystem.NewSecurityService(ctx, "securityservice_1", &sharedfilesystem.SecurityServiceArgs{
//				Name:        pulumi.String("security"),
//				Description: pulumi.String("created by terraform"),
//				Type:        pulumi.String("active_directory"),
//				Server:      pulumi.String("192.168.199.10"),
//				DnsIp:       pulumi.String("192.168.199.10"),
//				Domain:      pulumi.String("example.com"),
//				Ou:          pulumi.String("CN=Computers,DC=example,DC=com"),
//				User:        pulumi.String("joinDomainUser"),
//				Password:    pulumi.String("s8cret"),
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
// ## Import
//
// This resource can be imported by specifying the ID of the security service:
//
// ```sh
// $ pulumi import openstack:sharedfilesystem/securityService:SecurityService securityservice_1 id
// ```
type SecurityService struct {
	pulumi.CustomResourceState

	// The human-readable description for the security service.
	// Changing this updates the description of the existing security service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The security service DNS IP address that is used inside the
	// tenant network.
	DnsIp pulumi.StringPtrOutput `pulumi:"dnsIp"`
	// The security service domain.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// The name of the security service. Changing this updates the name
	// of the existing security service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The security service ou. An organizational unit can be added to
	// specify where the share ends up. New in Manila microversion 2.44.
	Ou pulumi.StringPtrOutput `pulumi:"ou"`
	// The user password, if you specify a user.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The owner of the Security Service.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a security service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security service.
	Region pulumi.StringOutput `pulumi:"region"`
	// The security service host name or IP address.
	Server pulumi.StringPtrOutput `pulumi:"server"`
	// The security service type - can either be active\_directory,
	// kerberos or ldap.  Changing this updates the existing security service.
	Type pulumi.StringOutput `pulumi:"type"`
	// The security service user or group name that is used by the
	// tenant.
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewSecurityService registers a new resource with the given unique name, arguments, and options.
func NewSecurityService(ctx *pulumi.Context,
	name string, args *SecurityServiceArgs, opts ...pulumi.ResourceOption) (*SecurityService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityService
	err := ctx.RegisterResource("openstack:sharedfilesystem/securityService:SecurityService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityService gets an existing SecurityService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityServiceState, opts ...pulumi.ResourceOption) (*SecurityService, error) {
	var resource SecurityService
	err := ctx.ReadResource("openstack:sharedfilesystem/securityService:SecurityService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityService resources.
type securityServiceState struct {
	// The human-readable description for the security service.
	// Changing this updates the description of the existing security service.
	Description *string `pulumi:"description"`
	// The security service DNS IP address that is used inside the
	// tenant network.
	DnsIp *string `pulumi:"dnsIp"`
	// The security service domain.
	Domain *string `pulumi:"domain"`
	// The name of the security service. Changing this updates the name
	// of the existing security service.
	Name *string `pulumi:"name"`
	// The security service ou. An organizational unit can be added to
	// specify where the share ends up. New in Manila microversion 2.44.
	Ou *string `pulumi:"ou"`
	// The user password, if you specify a user.
	Password *string `pulumi:"password"`
	// The owner of the Security Service.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a security service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security service.
	Region *string `pulumi:"region"`
	// The security service host name or IP address.
	Server *string `pulumi:"server"`
	// The security service type - can either be active\_directory,
	// kerberos or ldap.  Changing this updates the existing security service.
	Type *string `pulumi:"type"`
	// The security service user or group name that is used by the
	// tenant.
	User *string `pulumi:"user"`
}

type SecurityServiceState struct {
	// The human-readable description for the security service.
	// Changing this updates the description of the existing security service.
	Description pulumi.StringPtrInput
	// The security service DNS IP address that is used inside the
	// tenant network.
	DnsIp pulumi.StringPtrInput
	// The security service domain.
	Domain pulumi.StringPtrInput
	// The name of the security service. Changing this updates the name
	// of the existing security service.
	Name pulumi.StringPtrInput
	// The security service ou. An organizational unit can be added to
	// specify where the share ends up. New in Manila microversion 2.44.
	Ou pulumi.StringPtrInput
	// The user password, if you specify a user.
	Password pulumi.StringPtrInput
	// The owner of the Security Service.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a security service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security service.
	Region pulumi.StringPtrInput
	// The security service host name or IP address.
	Server pulumi.StringPtrInput
	// The security service type - can either be active\_directory,
	// kerberos or ldap.  Changing this updates the existing security service.
	Type pulumi.StringPtrInput
	// The security service user or group name that is used by the
	// tenant.
	User pulumi.StringPtrInput
}

func (SecurityServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityServiceState)(nil)).Elem()
}

type securityServiceArgs struct {
	// The human-readable description for the security service.
	// Changing this updates the description of the existing security service.
	Description *string `pulumi:"description"`
	// The security service DNS IP address that is used inside the
	// tenant network.
	DnsIp *string `pulumi:"dnsIp"`
	// The security service domain.
	Domain *string `pulumi:"domain"`
	// The name of the security service. Changing this updates the name
	// of the existing security service.
	Name *string `pulumi:"name"`
	// The security service ou. An organizational unit can be added to
	// specify where the share ends up. New in Manila microversion 2.44.
	Ou *string `pulumi:"ou"`
	// The user password, if you specify a user.
	Password *string `pulumi:"password"`
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a security service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security service.
	Region *string `pulumi:"region"`
	// The security service host name or IP address.
	Server *string `pulumi:"server"`
	// The security service type - can either be active\_directory,
	// kerberos or ldap.  Changing this updates the existing security service.
	Type string `pulumi:"type"`
	// The security service user or group name that is used by the
	// tenant.
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a SecurityService resource.
type SecurityServiceArgs struct {
	// The human-readable description for the security service.
	// Changing this updates the description of the existing security service.
	Description pulumi.StringPtrInput
	// The security service DNS IP address that is used inside the
	// tenant network.
	DnsIp pulumi.StringPtrInput
	// The security service domain.
	Domain pulumi.StringPtrInput
	// The name of the security service. Changing this updates the name
	// of the existing security service.
	Name pulumi.StringPtrInput
	// The security service ou. An organizational unit can be added to
	// specify where the share ends up. New in Manila microversion 2.44.
	Ou pulumi.StringPtrInput
	// The user password, if you specify a user.
	Password pulumi.StringPtrInput
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a security service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security service.
	Region pulumi.StringPtrInput
	// The security service host name or IP address.
	Server pulumi.StringPtrInput
	// The security service type - can either be active\_directory,
	// kerberos or ldap.  Changing this updates the existing security service.
	Type pulumi.StringInput
	// The security service user or group name that is used by the
	// tenant.
	User pulumi.StringPtrInput
}

func (SecurityServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityServiceArgs)(nil)).Elem()
}

type SecurityServiceInput interface {
	pulumi.Input

	ToSecurityServiceOutput() SecurityServiceOutput
	ToSecurityServiceOutputWithContext(ctx context.Context) SecurityServiceOutput
}

func (*SecurityService) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityService)(nil)).Elem()
}

func (i *SecurityService) ToSecurityServiceOutput() SecurityServiceOutput {
	return i.ToSecurityServiceOutputWithContext(context.Background())
}

func (i *SecurityService) ToSecurityServiceOutputWithContext(ctx context.Context) SecurityServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityServiceOutput)
}

// SecurityServiceArrayInput is an input type that accepts SecurityServiceArray and SecurityServiceArrayOutput values.
// You can construct a concrete instance of `SecurityServiceArrayInput` via:
//
//	SecurityServiceArray{ SecurityServiceArgs{...} }
type SecurityServiceArrayInput interface {
	pulumi.Input

	ToSecurityServiceArrayOutput() SecurityServiceArrayOutput
	ToSecurityServiceArrayOutputWithContext(context.Context) SecurityServiceArrayOutput
}

type SecurityServiceArray []SecurityServiceInput

func (SecurityServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityService)(nil)).Elem()
}

func (i SecurityServiceArray) ToSecurityServiceArrayOutput() SecurityServiceArrayOutput {
	return i.ToSecurityServiceArrayOutputWithContext(context.Background())
}

func (i SecurityServiceArray) ToSecurityServiceArrayOutputWithContext(ctx context.Context) SecurityServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityServiceArrayOutput)
}

// SecurityServiceMapInput is an input type that accepts SecurityServiceMap and SecurityServiceMapOutput values.
// You can construct a concrete instance of `SecurityServiceMapInput` via:
//
//	SecurityServiceMap{ "key": SecurityServiceArgs{...} }
type SecurityServiceMapInput interface {
	pulumi.Input

	ToSecurityServiceMapOutput() SecurityServiceMapOutput
	ToSecurityServiceMapOutputWithContext(context.Context) SecurityServiceMapOutput
}

type SecurityServiceMap map[string]SecurityServiceInput

func (SecurityServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityService)(nil)).Elem()
}

func (i SecurityServiceMap) ToSecurityServiceMapOutput() SecurityServiceMapOutput {
	return i.ToSecurityServiceMapOutputWithContext(context.Background())
}

func (i SecurityServiceMap) ToSecurityServiceMapOutputWithContext(ctx context.Context) SecurityServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityServiceMapOutput)
}

type SecurityServiceOutput struct{ *pulumi.OutputState }

func (SecurityServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityService)(nil)).Elem()
}

func (o SecurityServiceOutput) ToSecurityServiceOutput() SecurityServiceOutput {
	return o
}

func (o SecurityServiceOutput) ToSecurityServiceOutputWithContext(ctx context.Context) SecurityServiceOutput {
	return o
}

// The human-readable description for the security service.
// Changing this updates the description of the existing security service.
func (o SecurityServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The security service DNS IP address that is used inside the
// tenant network.
func (o SecurityServiceOutput) DnsIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringPtrOutput { return v.DnsIp }).(pulumi.StringPtrOutput)
}

// The security service domain.
func (o SecurityServiceOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// The name of the security service. Changing this updates the name
// of the existing security service.
func (o SecurityServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The security service ou. An organizational unit can be added to
// specify where the share ends up. New in Manila microversion 2.44.
func (o SecurityServiceOutput) Ou() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringPtrOutput { return v.Ou }).(pulumi.StringPtrOutput)
}

// The user password, if you specify a user.
func (o SecurityServiceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The owner of the Security Service.
func (o SecurityServiceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V2 Shared File System client.
// A Shared File System client is needed to create a security service. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// security service.
func (o SecurityServiceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The security service host name or IP address.
func (o SecurityServiceOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringPtrOutput { return v.Server }).(pulumi.StringPtrOutput)
}

// The security service type - can either be active\_directory,
// kerberos or ldap.  Changing this updates the existing security service.
func (o SecurityServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The security service user or group name that is used by the
// tenant.
func (o SecurityServiceOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityService) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

type SecurityServiceArrayOutput struct{ *pulumi.OutputState }

func (SecurityServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityService)(nil)).Elem()
}

func (o SecurityServiceArrayOutput) ToSecurityServiceArrayOutput() SecurityServiceArrayOutput {
	return o
}

func (o SecurityServiceArrayOutput) ToSecurityServiceArrayOutputWithContext(ctx context.Context) SecurityServiceArrayOutput {
	return o
}

func (o SecurityServiceArrayOutput) Index(i pulumi.IntInput) SecurityServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityService {
		return vs[0].([]*SecurityService)[vs[1].(int)]
	}).(SecurityServiceOutput)
}

type SecurityServiceMapOutput struct{ *pulumi.OutputState }

func (SecurityServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityService)(nil)).Elem()
}

func (o SecurityServiceMapOutput) ToSecurityServiceMapOutput() SecurityServiceMapOutput {
	return o
}

func (o SecurityServiceMapOutput) ToSecurityServiceMapOutputWithContext(ctx context.Context) SecurityServiceMapOutput {
	return o
}

func (o SecurityServiceMapOutput) MapIndex(k pulumi.StringInput) SecurityServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityService {
		return vs[0].(map[string]*SecurityService)[vs[1].(string)]
	}).(SecurityServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityServiceInput)(nil)).Elem(), &SecurityService{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityServiceArrayInput)(nil)).Elem(), SecurityServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityServiceMapInput)(nil)).Elem(), SecurityServiceMap{})
	pulumi.RegisterOutputType(SecurityServiceOutput{})
	pulumi.RegisterOutputType(SecurityServiceArrayOutput{})
	pulumi.RegisterOutputType(SecurityServiceMapOutput{})
}
