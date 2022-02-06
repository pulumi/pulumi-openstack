// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// If set to `false`, OpenStack authorization won't be perfomed automatically, if the initial auth token get expired.
// Defaults to `true`
func GetAllowReauth(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "openstack:allowReauth")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "OS_ALLOW_REAUTH").(bool)
}

// Application Credential ID to login with.
func GetApplicationCredentialId(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:applicationCredentialId")
}

// Application Credential name to login with.
func GetApplicationCredentialName(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:applicationCredentialName")
}

// Application Credential secret to login with.
func GetApplicationCredentialSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:applicationCredentialSecret")
}

// The Identity authentication URL.
func GetAuthUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:authUrl")
}

// A Custom CA certificate.
func GetCacertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:cacertFile")
}

// A client certificate to authenticate with.
func GetCert(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:cert")
}

// An entry in a `clouds.yaml` file to use.
func GetCloud(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "openstack:cloud")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "OS_CLOUD").(string)
}

// The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
func GetDefaultDomain(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:defaultDomain")
}

// If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
// to `true`.
func GetDelayedAuth(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "openstack:delayedAuth")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "OS_DELAYED_AUTH").(bool)
}

// If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
func GetDisableNoCacheHeader(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "openstack:disableNoCacheHeader")
}

// The ID of the Domain to scope to (Identity v3).
func GetDomainId(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:domainId")
}

// The name of the Domain to scope to (Identity v3).
func GetDomainName(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:domainName")
}

// Outputs very verbose logs with all calls made to and responses from OpenStack
func GetEnableLogging(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "openstack:enableLogging")
}

// A map of services with an endpoint to override what was from the Keystone catalog
func GetEndpointOverrides(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:endpointOverrides")
}
func GetEndpointType(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "openstack:endpointType")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "OS_ENDPOINT_TYPE").(string)
}

// Trust self-signed certificates.
func GetInsecure(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "openstack:insecure")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "OS_INSECURE").(bool)
}

// A client private key to authenticate with.
func GetKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:key")
}

// How many times HTTP connection should be retried until giving up.
func GetMaxRetries(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "openstack:maxRetries")
}

// Password to login with.
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:password")
}

// The ID of the domain where the proejct resides (Identity v3).
func GetProjectDomainId(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:projectDomainId")
}

// The name of the domain where the project resides (Identity v3).
func GetProjectDomainName(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:projectDomainName")
}

// The OpenStack region to connect to.
func GetRegion(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "openstack:region")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "OS_REGION_NAME").(string)
}

// Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
func GetSwauth(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "openstack:swauth")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "OS_SWAUTH").(bool)
}

// The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
func GetTenantId(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:tenantId")
}

// The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
func GetTenantName(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:tenantName")
}

// Authentication token to use as an alternative to username/password.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:token")
}

// If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
func GetUseOctavia(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "openstack:useOctavia")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "OS_USE_OCTAVIA").(bool)
}

// The ID of the domain where the user resides (Identity v3).
func GetUserDomainId(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:userDomainId")
}

// The name of the domain where the user resides (Identity v3).
func GetUserDomainName(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:userDomainName")
}

// Username to login with.
func GetUserId(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:userId")
}

// Username to login with.
func GetUserName(ctx *pulumi.Context) string {
	return config.Get(ctx, "openstack:userName")
}
