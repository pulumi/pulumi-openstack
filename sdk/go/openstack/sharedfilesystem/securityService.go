// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sharedfilesystem

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this resource to configure a security service.
// 
// A security service stores configuration information for clients for
// authentication and authorization (AuthN/AuthZ). For example, a share server
// will be the client for an existing service such as LDAP, Kerberos, or
// Microsoft Active Directory.
// 
// Minimum supported Manila microversion is 2.7.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-openstack/blob/master/website/docs/r/sharedfilesystem_securityservice_v2.html.markdown.
type SecurityService struct {
	s *pulumi.ResourceState
}

// NewSecurityService registers a new resource with the given unique name, arguments, and options.
func NewSecurityService(ctx *pulumi.Context,
	name string, args *SecurityServiceArgs, opts ...pulumi.ResourceOpt) (*SecurityService, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["dnsIp"] = nil
		inputs["domain"] = nil
		inputs["name"] = nil
		inputs["ou"] = nil
		inputs["password"] = nil
		inputs["region"] = nil
		inputs["server"] = nil
		inputs["type"] = nil
		inputs["user"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["dnsIp"] = args.DnsIp
		inputs["domain"] = args.Domain
		inputs["name"] = args.Name
		inputs["ou"] = args.Ou
		inputs["password"] = args.Password
		inputs["region"] = args.Region
		inputs["server"] = args.Server
		inputs["type"] = args.Type
		inputs["user"] = args.User
	}
	inputs["projectId"] = nil
	s, err := ctx.RegisterResource("openstack:sharedfilesystem/securityService:SecurityService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityService{s: s}, nil
}

// GetSecurityService gets an existing SecurityService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecurityServiceState, opts ...pulumi.ResourceOpt) (*SecurityService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["dnsIp"] = state.DnsIp
		inputs["domain"] = state.Domain
		inputs["name"] = state.Name
		inputs["ou"] = state.Ou
		inputs["password"] = state.Password
		inputs["projectId"] = state.ProjectId
		inputs["region"] = state.Region
		inputs["server"] = state.Server
		inputs["type"] = state.Type
		inputs["user"] = state.User
	}
	s, err := ctx.ReadResource("openstack:sharedfilesystem/securityService:SecurityService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecurityService) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecurityService) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The human-readable description for the security service.
// Changing this updates the description of the existing security service.
func (r *SecurityService) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The security service DNS IP address that is used inside the
// tenant network.
func (r *SecurityService) DnsIp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsIp"])
}

// The security service domain.
func (r *SecurityService) Domain() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domain"])
}

// The name of the security service. Changing this updates the name
// of the existing security service.
func (r *SecurityService) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The security service ou. An organizational unit can be added to
// specify where the share ends up. New in Manila microversion 2.44.
func (r *SecurityService) Ou() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ou"])
}

// The user password, if you specify a user.
func (r *SecurityService) Password() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["password"])
}

// The owner of the Security Service.
func (r *SecurityService) ProjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["projectId"])
}

// The region in which to obtain the V2 Shared File System client.
// A Shared File System client is needed to create a security service. If omitted, the
// `region` argument of the provider is used. Changing this creates a new
// security service.
func (r *SecurityService) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The security service host name or IP address.
func (r *SecurityService) Server() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["server"])
}

// The security service type - can either be active\_directory,
// kerberos or ldap.  Changing this updates the existing security service.
func (r *SecurityService) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// The security service user or group name that is used by the
// tenant.
func (r *SecurityService) User() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["user"])
}

// Input properties used for looking up and filtering SecurityService resources.
type SecurityServiceState struct {
	// The human-readable description for the security service.
	// Changing this updates the description of the existing security service.
	Description interface{}
	// The security service DNS IP address that is used inside the
	// tenant network.
	DnsIp interface{}
	// The security service domain.
	Domain interface{}
	// The name of the security service. Changing this updates the name
	// of the existing security service.
	Name interface{}
	// The security service ou. An organizational unit can be added to
	// specify where the share ends up. New in Manila microversion 2.44.
	Ou interface{}
	// The user password, if you specify a user.
	Password interface{}
	// The owner of the Security Service.
	ProjectId interface{}
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a security service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security service.
	Region interface{}
	// The security service host name or IP address.
	Server interface{}
	// The security service type - can either be active\_directory,
	// kerberos or ldap.  Changing this updates the existing security service.
	Type interface{}
	// The security service user or group name that is used by the
	// tenant.
	User interface{}
}

// The set of arguments for constructing a SecurityService resource.
type SecurityServiceArgs struct {
	// The human-readable description for the security service.
	// Changing this updates the description of the existing security service.
	Description interface{}
	// The security service DNS IP address that is used inside the
	// tenant network.
	DnsIp interface{}
	// The security service domain.
	Domain interface{}
	// The name of the security service. Changing this updates the name
	// of the existing security service.
	Name interface{}
	// The security service ou. An organizational unit can be added to
	// specify where the share ends up. New in Manila microversion 2.44.
	Ou interface{}
	// The user password, if you specify a user.
	Password interface{}
	// The region in which to obtain the V2 Shared File System client.
	// A Shared File System client is needed to create a security service. If omitted, the
	// `region` argument of the provider is used. Changing this creates a new
	// security service.
	Region interface{}
	// The security service host name or IP address.
	Server interface{}
	// The security service type - can either be active\_directory,
	// kerberos or ldap.  Changing this updates the existing security service.
	Type interface{}
	// The security service user or group name that is used by the
	// tenant.
	User interface{}
}
