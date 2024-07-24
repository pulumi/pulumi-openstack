// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openstack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-openstack/sdk/v4/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the openstack package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// Application Credential ID to login with.
	ApplicationCredentialId pulumi.StringPtrOutput `pulumi:"applicationCredentialId"`
	// Application Credential name to login with.
	ApplicationCredentialName pulumi.StringPtrOutput `pulumi:"applicationCredentialName"`
	// Application Credential secret to login with.
	ApplicationCredentialSecret pulumi.StringPtrOutput `pulumi:"applicationCredentialSecret"`
	// The Identity authentication URL.
	AuthUrl pulumi.StringPtrOutput `pulumi:"authUrl"`
	// A Custom CA certificate.
	CacertFile pulumi.StringPtrOutput `pulumi:"cacertFile"`
	// A client certificate to authenticate with.
	Cert pulumi.StringPtrOutput `pulumi:"cert"`
	// An entry in a `clouds.yaml` file to use.
	Cloud pulumi.StringPtrOutput `pulumi:"cloud"`
	// The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
	DefaultDomain pulumi.StringPtrOutput `pulumi:"defaultDomain"`
	// The ID of the Domain to scope to (Identity v3).
	DomainId pulumi.StringPtrOutput `pulumi:"domainId"`
	// The name of the Domain to scope to (Identity v3).
	DomainName   pulumi.StringPtrOutput `pulumi:"domainName"`
	EndpointType pulumi.StringPtrOutput `pulumi:"endpointType"`
	// A client private key to authenticate with.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// Password to login with.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The ID of the domain where the proejct resides (Identity v3).
	ProjectDomainId pulumi.StringPtrOutput `pulumi:"projectDomainId"`
	// The name of the domain where the project resides (Identity v3).
	ProjectDomainName pulumi.StringPtrOutput `pulumi:"projectDomainName"`
	// The OpenStack region to connect to.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
	TenantName pulumi.StringPtrOutput `pulumi:"tenantName"`
	// Authentication token to use as an alternative to username/password.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// The ID of the domain where the user resides (Identity v3).
	UserDomainId pulumi.StringPtrOutput `pulumi:"userDomainId"`
	// The name of the domain where the user resides (Identity v3).
	UserDomainName pulumi.StringPtrOutput `pulumi:"userDomainName"`
	// User ID to login with.
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
	// Username to login with.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.AllowReauth == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "OS_ALLOW_REAUTH"); d != nil {
			args.AllowReauth = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.Cloud == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OS_CLOUD"); d != nil {
			args.Cloud = pulumi.StringPtr(d.(string))
		}
	}
	if args.DelayedAuth == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "OS_DELAYED_AUTH"); d != nil {
			args.DelayedAuth = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.EndpointType == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OS_ENDPOINT_TYPE"); d != nil {
			args.EndpointType = pulumi.StringPtr(d.(string))
		}
	}
	if args.Insecure == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "OS_INSECURE"); d != nil {
			args.Insecure = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.Region == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OS_REGION_NAME"); d != nil {
			args.Region = pulumi.StringPtr(d.(string))
		}
	}
	if args.Swauth == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "OS_SWAUTH"); d != nil {
			args.Swauth = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:openstack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If set to `false`, OpenStack authorization won't be perfomed automatically, if the initial auth token get expired.
	// Defaults to `true`
	AllowReauth *bool `pulumi:"allowReauth"`
	// Application Credential ID to login with.
	ApplicationCredentialId *string `pulumi:"applicationCredentialId"`
	// Application Credential name to login with.
	ApplicationCredentialName *string `pulumi:"applicationCredentialName"`
	// Application Credential secret to login with.
	ApplicationCredentialSecret *string `pulumi:"applicationCredentialSecret"`
	// The Identity authentication URL.
	AuthUrl *string `pulumi:"authUrl"`
	// A Custom CA certificate.
	CacertFile *string `pulumi:"cacertFile"`
	// A client certificate to authenticate with.
	Cert *string `pulumi:"cert"`
	// An entry in a `clouds.yaml` file to use.
	Cloud *string `pulumi:"cloud"`
	// The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
	DefaultDomain *string `pulumi:"defaultDomain"`
	// If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
	// to `true`.
	DelayedAuth *bool `pulumi:"delayedAuth"`
	// If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
	DisableNoCacheHeader *bool `pulumi:"disableNoCacheHeader"`
	// The ID of the Domain to scope to (Identity v3).
	DomainId *string `pulumi:"domainId"`
	// The name of the Domain to scope to (Identity v3).
	DomainName *string `pulumi:"domainName"`
	// Outputs very verbose logs with all calls made to and responses from OpenStack
	EnableLogging *bool `pulumi:"enableLogging"`
	// A map of services with an endpoint to override what was from the Keystone catalog
	EndpointOverrides map[string]interface{} `pulumi:"endpointOverrides"`
	EndpointType      *string                `pulumi:"endpointType"`
	// Trust self-signed certificates.
	Insecure *bool `pulumi:"insecure"`
	// A client private key to authenticate with.
	Key *string `pulumi:"key"`
	// How many times HTTP connection should be retried until giving up.
	MaxRetries *int `pulumi:"maxRetries"`
	// Password to login with.
	Password *string `pulumi:"password"`
	// The ID of the domain where the proejct resides (Identity v3).
	ProjectDomainId *string `pulumi:"projectDomainId"`
	// The name of the domain where the project resides (Identity v3).
	ProjectDomainName *string `pulumi:"projectDomainName"`
	// The OpenStack region to connect to.
	Region *string `pulumi:"region"`
	// Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
	Swauth *bool `pulumi:"swauth"`
	// If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).
	SystemScope *bool `pulumi:"systemScope"`
	// The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
	TenantId *string `pulumi:"tenantId"`
	// The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
	TenantName *string `pulumi:"tenantName"`
	// Authentication token to use as an alternative to username/password.
	Token *string `pulumi:"token"`
	// If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
	//
	// Deprecated: Users not using loadbalancer resources can ignore this message. Support for neutron-lbaas will be removed on next major release. Octavia will be the only supported method for loadbalancer resources. Users using octavia will have to remove 'use_octavia' option from the provider configuration block. Users using neutron-lbaas will have to migrate/upgrade to octavia.
	UseOctavia *bool `pulumi:"useOctavia"`
	// The ID of the domain where the user resides (Identity v3).
	UserDomainId *string `pulumi:"userDomainId"`
	// The name of the domain where the user resides (Identity v3).
	UserDomainName *string `pulumi:"userDomainName"`
	// User ID to login with.
	UserId *string `pulumi:"userId"`
	// Username to login with.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If set to `false`, OpenStack authorization won't be perfomed automatically, if the initial auth token get expired.
	// Defaults to `true`
	AllowReauth pulumi.BoolPtrInput
	// Application Credential ID to login with.
	ApplicationCredentialId pulumi.StringPtrInput
	// Application Credential name to login with.
	ApplicationCredentialName pulumi.StringPtrInput
	// Application Credential secret to login with.
	ApplicationCredentialSecret pulumi.StringPtrInput
	// The Identity authentication URL.
	AuthUrl pulumi.StringPtrInput
	// A Custom CA certificate.
	CacertFile pulumi.StringPtrInput
	// A client certificate to authenticate with.
	Cert pulumi.StringPtrInput
	// An entry in a `clouds.yaml` file to use.
	Cloud pulumi.StringPtrInput
	// The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
	DefaultDomain pulumi.StringPtrInput
	// If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
	// to `true`.
	DelayedAuth pulumi.BoolPtrInput
	// If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
	DisableNoCacheHeader pulumi.BoolPtrInput
	// The ID of the Domain to scope to (Identity v3).
	DomainId pulumi.StringPtrInput
	// The name of the Domain to scope to (Identity v3).
	DomainName pulumi.StringPtrInput
	// Outputs very verbose logs with all calls made to and responses from OpenStack
	EnableLogging pulumi.BoolPtrInput
	// A map of services with an endpoint to override what was from the Keystone catalog
	EndpointOverrides pulumi.MapInput
	EndpointType      pulumi.StringPtrInput
	// Trust self-signed certificates.
	Insecure pulumi.BoolPtrInput
	// A client private key to authenticate with.
	Key pulumi.StringPtrInput
	// How many times HTTP connection should be retried until giving up.
	MaxRetries pulumi.IntPtrInput
	// Password to login with.
	Password pulumi.StringPtrInput
	// The ID of the domain where the proejct resides (Identity v3).
	ProjectDomainId pulumi.StringPtrInput
	// The name of the domain where the project resides (Identity v3).
	ProjectDomainName pulumi.StringPtrInput
	// The OpenStack region to connect to.
	Region pulumi.StringPtrInput
	// Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
	Swauth pulumi.BoolPtrInput
	// If set to `true`, system scoped authorization will be enabled. Defaults to `false` (Identity v3).
	SystemScope pulumi.BoolPtrInput
	// The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
	TenantId pulumi.StringPtrInput
	// The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
	TenantName pulumi.StringPtrInput
	// Authentication token to use as an alternative to username/password.
	Token pulumi.StringPtrInput
	// If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
	//
	// Deprecated: Users not using loadbalancer resources can ignore this message. Support for neutron-lbaas will be removed on next major release. Octavia will be the only supported method for loadbalancer resources. Users using octavia will have to remove 'use_octavia' option from the provider configuration block. Users using neutron-lbaas will have to migrate/upgrade to octavia.
	UseOctavia pulumi.BoolPtrInput
	// The ID of the domain where the user resides (Identity v3).
	UserDomainId pulumi.StringPtrInput
	// The name of the domain where the user resides (Identity v3).
	UserDomainName pulumi.StringPtrInput
	// User ID to login with.
	UserId pulumi.StringPtrInput
	// Username to login with.
	UserName pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// Application Credential ID to login with.
func (o ProviderOutput) ApplicationCredentialId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApplicationCredentialId }).(pulumi.StringPtrOutput)
}

// Application Credential name to login with.
func (o ProviderOutput) ApplicationCredentialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApplicationCredentialName }).(pulumi.StringPtrOutput)
}

// Application Credential secret to login with.
func (o ProviderOutput) ApplicationCredentialSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApplicationCredentialSecret }).(pulumi.StringPtrOutput)
}

// The Identity authentication URL.
func (o ProviderOutput) AuthUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AuthUrl }).(pulumi.StringPtrOutput)
}

// A Custom CA certificate.
func (o ProviderOutput) CacertFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CacertFile }).(pulumi.StringPtrOutput)
}

// A client certificate to authenticate with.
func (o ProviderOutput) Cert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Cert }).(pulumi.StringPtrOutput)
}

// An entry in a `clouds.yaml` file to use.
func (o ProviderOutput) Cloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Cloud }).(pulumi.StringPtrOutput)
}

// The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
func (o ProviderOutput) DefaultDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.DefaultDomain }).(pulumi.StringPtrOutput)
}

// The ID of the Domain to scope to (Identity v3).
func (o ProviderOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

// The name of the Domain to scope to (Identity v3).
func (o ProviderOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.EndpointType }).(pulumi.StringPtrOutput)
}

// A client private key to authenticate with.
func (o ProviderOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Key }).(pulumi.StringPtrOutput)
}

// Password to login with.
func (o ProviderOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The ID of the domain where the proejct resides (Identity v3).
func (o ProviderOutput) ProjectDomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ProjectDomainId }).(pulumi.StringPtrOutput)
}

// The name of the domain where the project resides (Identity v3).
func (o ProviderOutput) ProjectDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ProjectDomainName }).(pulumi.StringPtrOutput)
}

// The OpenStack region to connect to.
func (o ProviderOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
func (o ProviderOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
func (o ProviderOutput) TenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TenantName }).(pulumi.StringPtrOutput)
}

// Authentication token to use as an alternative to username/password.
func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

// The ID of the domain where the user resides (Identity v3).
func (o ProviderOutput) UserDomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.UserDomainId }).(pulumi.StringPtrOutput)
}

// The name of the domain where the user resides (Identity v3).
func (o ProviderOutput) UserDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.UserDomainName }).(pulumi.StringPtrOutput)
}

// User ID to login with.
func (o ProviderOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

// Username to login with.
func (o ProviderOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
