[![Actions Status](https://github.com/pulumi/pulumi-openstack/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-openstack/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fopenstack.svg)](https://www.npmjs.com/package/@pulumi/openstack)
[![Python version](https://badge.fury.io/py/pulumi-openstack.svg)](https://pypi.org/project/pulumi-openstack)
[![NuGet version](https://badge.fury.io/nu/pulumi.openstack.svg)](https://badge.fury.io/nu/pulumi.openstack)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-openstack/sdk/v2/go)](https://pkg.go.dev/github.com/pulumi/pulumi-openstack/sdk/v2/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-openstack/blob/master/LICENSE)

# OpenStack Resource Provider

The OpenStack resource provider for Pulumi lets you use OpenStack resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/openstack

or `yarn`:

    $ yarn add @pulumi/openstack

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_openstack

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-openstack/sdk/v2/go/...
    
### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Openstack

## Configuration

The following configuration points are available:

- `openstack:authUrl` - (Optional) The Identity authentication URL. If omitted, the `OS_AUTH_URL` environment variable is used.
- `openstack:cloud` - (Optional) An entry in a `clouds.yaml` file. See the OpenStack `openstacksdk`
  [documentation](https://docs.openstack.org/openstacksdk/latest/user/config/configuration.html) for more information about 
  `clouds.yaml` files. If omitted, the `OS_CLOUD` environment variable is used.
- `openstack:region` - (Optional) The region of the OpenStack cloud to use. If omitted, the `OS_REGION_NAME` environment 
  variable is used. If `OS_REGION_NAME` is not set, then no region will be used. It should be possible to omit the region 
  in single-region OpenStack environments, but this behavior may vary depending on the OpenStack environment being used.
- `openstack:userName` - (Optional) The Username to login with. If omitted, the `OS_USERNAME` environment variable is used.
- `openstack:userId` - (Optional) The User ID to login with. If omitted, the `OS_USER_ID` environment variable is used.
- `openstack:applicationCredentialId` - (Optional) (Identity v3 only) The ID of an application credential to authenticate with. An
  `applicationCredentialSecret` has to bet set along with this parameter. Can be set via the `OS_APPLICATION_CREDENTIAL_ID` 
  environment variable.
- `openstack:applicationCredentialName` - (Optional) (Identity v3 only) The name of an application credential to authenticate with. 
  Conflicts with the `applicationCredentialName`, requires `userId`, or `userName` and `userDomainName` (or `userDomainId`) to be set.
  Can be set via the `OS_APPLICATION_CREDENTIAL_NAME` environment variable.
- `openstack:applicationCredentialSecret` - (Optional) (Identity v3 only) The secret of an application credential to authenticate with. 
  Required by `applicationCredentialId` or `applicationCredentialName`. Can be set via the `OS_APPLICATION_CREDENTIAL_SECRET` 
  environment variable. 
- `openstack:tenantId` - (Optional) The ID of the Tenant (Identity v2) or Project (Identity v3) to login with. If omitted, the 
  `OS_TENANT_ID` or `OS_PROJECT_ID` environment variables are used.
- `openstack:tenantName` - (Optional) The Name of the Tenant (Identity v2) or Project (Identity v3) to login with. If omitted, 
  the `OS_TENANT_NAME` or `OS_PROJECT_NAME` environment variable are used.
- `openstack:tenantName` - (Optional) The Password to login with. If omitted, the
  `OS_PASSWORD` environment variable is used.
- `openstack:token` - (Optional) A token is an expiring, temporary means of access issued via the Keystone service. By specifying 
  a token, you do not have to specify a username/password combination, since the token was already created by a username/password 
  out of band of the provider. If omitted, the `OS_TOKEN` or `OS_AUTH_TOKEN` environment variables are used.
- `openstack:userDomainName` - (Optional) The domain name where the user is located. If omitted, the `OS_USER_DOMAIN_NAME` 
  environment variable is checked.
- `openstack:userDomainId` - (Optional) The domain ID where the user is located. If omitted, the `OS_USER_DOMAIN_ID` environment 
  variable is checked.
- `openstack:projectDomainName` - (Optional) The domain name where the project is located. If omitted, the `OS_PROJECT_DOMAIN_NAME` 
  environment variable is checked.
- `openstack:projectDomainId` - (Optional) The domain ID where the project is located. If omitted, the `OS_PROJECT_DOMAIN_ID` 
  environment variable is checked.
- `openstack:domainId` - (Optional) The ID of the Domain to scope to (Identity v3). If omitted, the `OS_DOMAIN_ID` environment 
  variable is checked.
- `openstack:domainName` - (Optional) The Name of the Domain to scope to (Identity v3). If omitted, the `OS_DOMAIN_NAME` environment 
  variable is checked.
- `openstack:defaultDomain` - (Optional) The ID of the Domain to scope to if no other domain is specified (Identity v3). If omitted, 
  the environment variable `OS_DEFAULT_DOMAIN` is checked or a default value of `default` will be used.
- `openstack:insecure` - (Optional) Trust self-signed SSL certificates. If omitted, the `OS_INSECURE` environment variable is used.
- `openstack:cacertFile` - (Optional) Specify a custom CA certificate when communicating over SSL. You can specify either a path 
  to the file or the contents of the certificate. If omitted, the `OS_CACERT` environment variable is used.
- `openstack:cert` - (Optional) Specify client certificate file for SSL client authentication. You can specify either a path to 
  the file or the contents of the certificate. If omitted the `OS_CERT` environment variable is used.
- `openstack:key` - (Optional) Specify client private key file for SSL client authentication. You can specify either a path 
  to the file or the contents of the key. If omitted the `OS_KEY` environment variable is used.
- `openstack:endpointType` - (Optional) Specify which type of endpoint to use from the service catalog. It can be set using the 
  `OS_ENDPOINT_TYPE` environment variable. If not set, public endpoints is used.
- `openstack:endpointOverrides` - (Optional) A set of key/value pairs that can override an endpoint for a specified OpenStack service. 
  Setting an override requires you to specify the full and complete endpoint URL. This might also invalidate any region you have set, 
  too. Please use this at your own risk.
- `openstack:swauth` - (Optional) Set to `true` to authenticate against Swauth, a Swift-native authentication system. If omitted, the 
  `OS_SWAUTH` environment variable is used. You must also set `username` to the Swauth/Swift username such as `username:project`. 
  Set the `password` to the Swauth/Swift key. Finally, set `auth_url` as the location of the Swift service. Note that this
  will only work when used with the OpenStack Object Storage resources.
- `openstack:userOctavia` - (Optional) If set to `true`, API requests will go the Load Balancer service (Octavia) instead of 
  the Networking service (Neutron).
- `openstack:disableNoCacheHeader` - (Optional) If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
  If omitted this header is added to all API requests to force HTTP caches (if any) to go upstream instead of serving cached responses.
- `openstack:delayedAuth` - (Optional) If set to `true`, OpenStack authorization will be perfomed, when the service provider client is called.
- `openstack:allowReauth` - (Optional) If set to `true`, OpenStack authorization will be perfomed automatically, if the initial auth token get 
  expired. This is useful, when the token TTL is low or the overall provider execution time expected to be greater than the initial token TTL.

## Reference

For further information, please visit [the OpenStack provider docs](https://www.pulumi.com/docs/intro/cloud-providers/openstack) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/openstack).
