---
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Openstack Provider
meta_desc: Provides an overview on how to configure the Pulumi Openstack provider.
layout: package
---

## Installation

The Openstack provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/openstack`](https://www.npmjs.com/package/@pulumi/openstack)
* Python: [`pulumi-openstack`](https://pypi.org/project/pulumi-openstack/)
* Go: [`github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack`](https://github.com/pulumi/pulumi-openstack)
* .NET: [`Pulumi.Openstack`](https://www.nuget.org/packages/Pulumi.Openstack)
* Java: [`com.pulumi/openstack`](https://central.sonatype.com/artifact/com.pulumi/openstack)

## Overview

The OpenStack provider is used to interact with the
many resources supported by OpenStack. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.
## Example Usage

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    openstack:authUrl:
        value: http://myauthurl:5000/v3
    openstack:password:
        value: pwd
    openstack:region:
        value: RegionOne
    openstack:tenantName:
        value: admin
    openstack:userName:
        value: admin

```
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as openstack from "@pulumi/openstack";

// Create a web server
const test_server = new openstack.compute.Instance("test-server", {});
```
{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    openstack:authUrl:
        value: http://myauthurl:5000/v3
    openstack:password:
        value: pwd
    openstack:region:
        value: RegionOne
    openstack:tenantName:
        value: admin
    openstack:userName:
        value: admin

```
```python
import pulumi
import pulumi_openstack as openstack

# Create a web server
test_server = openstack.compute.Instance("test-server")
```
{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    openstack:authUrl:
        value: http://myauthurl:5000/v3
    openstack:password:
        value: pwd
    openstack:region:
        value: RegionOne
    openstack:tenantName:
        value: admin
    openstack:userName:
        value: admin

```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using OpenStack = Pulumi.OpenStack;

return await Deployment.RunAsync(() =>
{
    // Create a web server
    var test_server = new OpenStack.Compute.Instance("test-server");

});

```
{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    openstack:authUrl:
        value: http://myauthurl:5000/v3
    openstack:password:
        value: pwd
    openstack:region:
        value: RegionOne
    openstack:tenantName:
        value: admin
    openstack:userName:
        value: admin

```
```go
package main

import (
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a web server
		_, err := compute.NewInstance(ctx, "test-server", nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    openstack:authUrl:
        value: http://myauthurl:5000/v3
    openstack:password:
        value: pwd
    openstack:region:
        value: RegionOne
    openstack:tenantName:
        value: admin
    openstack:userName:
        value: admin

```
```yaml
resources:
  # Create a web server
  test-server:
    type: openstack:compute:Instance
```
{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    openstack:authUrl:
        value: http://myauthurl:5000/v3
    openstack:password:
        value: pwd
    openstack:region:
        value: RegionOne
    openstack:tenantName:
        value: admin
    openstack:userName:
        value: admin

```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.openstack.compute.Instance;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        // Create a web server
        var test_server = new Instance("test-server");

    }
}
```
{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

* `authUrl` - (Optional; required if `cloud` is not specified) The Identity
  authentication URL. If omitted, the `OS_AUTH_URL` environment variable is used.

* `cloud` - (Optional; required if `authUrl` is not specified) An entry in a
  `clouds.yaml` file. See the OpenStack `openstacksdk`
  [documentation](https://docs.openstack.org/openstacksdk/latest/user/config/configuration.html)
  for more information about `clouds.yaml` files. If omitted, the `OS_CLOUD`
  environment variable is used.

* `region` - (Optional) The region of the OpenStack cloud to use. If omitted,
  the `OS_REGION_NAME` environment variable is used. If `OS_REGION_NAME` is
  not set, then no region will be used. It should be possible to omit the
  region in single-region OpenStack environments, but this behavior may vary
  depending on the OpenStack environment being used.

* `userName` - (Optional) The Username to login with. If omitted, the
  `OS_USERNAME` environment variable is used.

* `userId` - (Optional) The User ID to login with. If omitted, the
  `OS_USER_ID` environment variable is used.

* `applicationCredentialId` - (Optional) (Identity v3 only) The ID of an
  application credential to authenticate with. An
  `applicationCredentialSecret` has to bet set along with this parameter.
  If omitted, the `OS_APPLICATION_CREDENTIAL_ID` environment variable is used.

* `applicationCredentialName` - (Optional) (Identity v3 only) The name of an
  application credential to authenticate with. Requires `userId`, or
  `userName` and `userDomainName` (or `userDomainId`) to be set.
  If omitted, the `OS_APPLICATION_CREDENTIAL_NAME` environment variable is used.

* `applicationCredentialSecret` - (Optional) (Identity v3 only) The secret of
  an application credential to authenticate with. Required by
  `applicationCredentialId` or `applicationCredentialName`.
  If omitted, the `OS_APPLICATION_CREDENTIAL_SECRET` environment variable is used.

* `tenantId` - (Optional) The ID of the Tenant (Identity v2) or Project
  (Identity v3) to login with. If omitted, the `OS_TENANT_ID` or
  `OS_PROJECT_ID` environment variables are used.

* `tenantName` - (Optional) The Name of the Tenant (Identity v2) or Project
  (Identity v3) to login with. If omitted, the `OS_TENANT_NAME` or
  `OS_PROJECT_NAME` environment variable are used.

* `password` - (Optional) The Password to login with. If omitted, the
  `OS_PASSWORD` environment variable is used.

* `token` - (Optional; Required if not using `userName` and `password`)
  A token is an expiring, temporary means of access issued via the Keystone
  service. By specifying a token, you do not have to specify a username/password
  combination, since the token was already created by a username/password out of
  band of Pulumi. If omitted, the `OS_TOKEN` or `OS_AUTH_TOKEN` environment
  variables are used.

* `userDomainName` - (Optional) The domain name where the user is located. If
  omitted, the `OS_USER_DOMAIN_NAME` environment variable is checked.

* `userDomainId` - (Optional) The domain ID where the user is located. If
  omitted, the `OS_USER_DOMAIN_ID` environment variable is checked.

* `projectDomainName` - (Optional) The domain name where the project is
  located. If omitted, the `OS_PROJECT_DOMAIN_NAME` environment variable is
  checked.

* `projectDomainId` - (Optional) The domain ID where the project is located.
  If omitted, the `OS_PROJECT_DOMAIN_ID` environment variable is checked.

* `domainId` - (Optional) The ID of the Domain to scope to (Identity v3). If
  omitted, the `OS_DOMAIN_ID` environment variable is checked.

* `domainName` - (Optional) The Name of the Domain to scope to (Identity v3).
  If omitted, the following environment variables are checked (in this order):
  `OS_DOMAIN_NAME`.

* `defaultDomain` - (Optional) The ID of the Domain to scope to if no other
  domain is specified (Identity v3). If omitted, the environment variable
  `OS_DEFAULT_DOMAIN` is checked or a default value of "default" will be
  used.

* `systemScope` - (Optional) Set to `true` to enable system scoped authorization. If omitted, the `OS_SYSTEM_SCOPE` environment variable is used.

* `insecure` - (Optional) Trust self-signed SSL certificates. If omitted, the
  `OS_INSECURE` environment variable is used.

* `cacertFile` - (Optional) Specify a custom CA certificate when communicating
  over SSL. You can specify either a path to the file or the contents of the
  certificate. If omitted, the `OS_CACERT` environment variable is used.

* `cert` - (Optional) Specify client certificate file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the certificate. If omitted the `OS_CERT` environment variable is used.

* `key` - (Optional) Specify client private key file for SSL client
  authentication. You can specify either a path to the file or the contents of
  the key. If omitted the `OS_KEY` environment variable is used.

* `endpointType` - (Optional) Specify which type of endpoint to use from the
  service catalog. It can be set using the `OS_ENDPOINT_TYPE` environment
  variable. If not set, public endpoints is used.

* `endpointOverrides` - (Optional) A set of key/value pairs that can
  override an endpoint for a specified OpenStack service. Setting an override
  requires you to specify the full and complete endpoint URL. This might
  also invalidate any region you have set, too. Please see below for more details.
  Please use this at your own risk.

* `swauth` - (Optional) Set to `true` to authenticate against Swauth, a
  Swift-native authentication system. If omitted, the `OS_SWAUTH` environment
  variable is used. You must also set `username` to the Swauth/Swift username
  such as `username:project`. Set the `password` to the Swauth/Swift key.
  Finally, set `authUrl` as the location of the Swift service. Note that this
  will only work when used with the OpenStack Object Storage resources.

* `disableNoCacheHeader` - (Optional) If set to `true`, the HTTP
  `Cache-Control: no-cache` header will not be added by default to all API requests.
  If omitted this header is added to all API requests to force HTTP caches (if any)
  to go upstream instead of serving cached responses.

* `delayedAuth` - (Optional) If set to `false`, OpenStack authorization will be perfomed,
  every time the service provider client is called. Defaults to `true`.
  If omitted, the `OS_DELAYED_AUTH` environment variable is checked.

* `allowReauth` - (Optional) If set to `false`, OpenStack authorization won't be
  perfomed automatically, if the initial auth token get expired. Defaults to `true`.
  If omitted, the `OS_ALLOW_REAUTH` environment variable is checked.

* `maxRetries` - (Optional) If set to a value greater than 0, the OpenStack
  client will retry failed HTTP connections and Too Many Requests (429 code)
  HTTP responses with a `Retry-After` header within the specified value.

* `enableLogging` - (Optional) When enabled, generates verbose logs containing
  all the calls made to and responses received from OpenStack.
## Overriding Service API Endpoints

There might be a situation in which you want or need to override an API endpoint
rather than use the endpoint which was returned to you in the service catalog.
You can do this by configuring the `endpointOverrides` argument in the provider
configuration:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    openstack:endpointOverrides:
        value:
            network: https://example.com:9696/v2.0/
            volumev2: https://volumes.example.com:8776/v2/3eb25ae78e7b42d68276e9bca66c8e44/

```

Note how each URL ends in a "/" and the `volumev2` service includes the
tenant/project UUID. You must make sure you specify the full and complete
endpoint URL for this to work.

The service keys are the standard service entries used in the OpenStack
Identity/Keystone service catalog. This provider supports:

* `compute`: Compute / Nova v2
* `container-infra`: Container Infra / Magnum v1
* `database`: Database / Trove v1
* `dns`: DNS / Designate v2
* `identity`: Identity / Keystone v3
* `image`: Image / Glance v2
* `network`: Networking / Neutron v2
* `object-store`: Object Storage / Swift v1
* `octavia`: Load Balancing as a Service / Octavia v2
* `sharev2`: Shared Filesystem / Manila v2
* `volume`: Block Storage / Cinder v1
* `volumev2`: Block Storage / Cinder v2
* `volumev3`: Block Storage / Cinder v3
* `workflowv2`: Workflow / Mistral v2

Please use this feature at your own risk. If you are unsure about needing
to override an endpoint, you most likely do not need to override one.
### Overriding service types

You can override service types if needed. For example, to use the
`block-storage` service type instead of `volumev3`, configure the provider as
follows:

```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime:
config:
    openstack:endpointOverrides:
        value:
            volumev3: block-storage

```

This ensures that requests intended for `volumev3` are directed to the
`block-storage` service.
## OpenStack Releases and Versions

This provider aims to support "vanilla" OpenStack. This means that we do all
testing and development using the official upstream OpenStack code. If your
OpenStack environment has patches or modifications, we do our best to
accommodate these modifications, but we can't guarantee this.

We try to support *all* releases of OpenStack when we can. If your OpenStack
cloud is running an older release, we still should be able to support it.
### Octavia api versioning

[Octavia api](https://docs.openstack.org/api-ref/load-balancer/v2/) is using
minor versions when adding new features and functionality. The required minor
version of each feature are documented on the resource page. When using such a
feature ensure that your Openstack cloud supports the required minor version.
A simple way of checking which minor versions are supported on your Openstack
cloud is the following:

```shell
export OCTAVIA_URL=`openstack endpoint list --interface public --service octavia -f value -c URL`
export OS_TOKEN=`openstack token issue -f value -c id`
curl -s -H "X-Auth-Token: $OS_TOKEN" "$OCTAVIA_URL" | jq -r '.versions[]|select(.status=="SUPPORTED")|.id' | tail -1
```
### Rackspace Compatibility

Using this OpenStack provider with Rackspace is not supported and not
guaranteed to work; however, users have reported success with the
following notes in mind:

* Interacting with instances and networks has been seen to work. Interacting
  with all other resources is either untested or known to not work.

* Use your *password* instead of your Rackspace API KEY.

* To use networks, override the endpoint in your provider configuration

```
provider "openstack" {
  user_name = "your-username"
  password  = "your-password"
  tenant_id = "your-tenant-id"
  region    = "IAD"
  auth_url  = "https://identity.api.rackspacecloud.com/v2.0/"
  endpoint_overrides = {
    "network" = "https://iad.networks.api.rackspacecloud.com/v2.0/"
  }
}
```

* Explicitly define the public and private networks in your
  instances as shown below:

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as openstack from "@pulumi/openstack";

const myInstance = new openstack.compute.Instance("my_instance", {
    name: "my_instance",
    region: "DFW",
    imageId: "fabe045f-43f8-4991-9e6c-5cabd617538c",
    flavorId: "general1-4",
    keyPair: "provisioning_key",
    networks: [
        {
            uuid: "00000000-0000-0000-0000-000000000000",
            name: "public",
        },
        {
            uuid: "11111111-1111-1111-1111-111111111111",
            name: "private",
        },
    ],
});
```
{{% /choosable %}}
{{% choosable language python %}}
```python
import pulumi
import pulumi_openstack as openstack

my_instance = openstack.compute.Instance("my_instance",
    name="my_instance",
    region="DFW",
    image_id="fabe045f-43f8-4991-9e6c-5cabd617538c",
    flavor_id="general1-4",
    key_pair="provisioning_key",
    networks=[
        {
            "uuid": "00000000-0000-0000-0000-000000000000",
            "name": "public",
        },
        {
            "uuid": "11111111-1111-1111-1111-111111111111",
            "name": "private",
        },
    ])
```
{{% /choosable %}}
{{% choosable language csharp %}}
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using OpenStack = Pulumi.OpenStack;

return await Deployment.RunAsync(() =>
{
    var myInstance = new OpenStack.Compute.Instance("my_instance", new()
    {
        Name = "my_instance",
        Region = "DFW",
        ImageId = "fabe045f-43f8-4991-9e6c-5cabd617538c",
        FlavorId = "general1-4",
        KeyPair = "provisioning_key",
        Networks = new[]
        {
            new OpenStack.Compute.Inputs.InstanceNetworkArgs
            {
                Uuid = "00000000-0000-0000-0000-000000000000",
                Name = "public",
            },
            new OpenStack.Compute.Inputs.InstanceNetworkArgs
            {
                Uuid = "11111111-1111-1111-1111-111111111111",
                Name = "private",
            },
        },
    });

});

```
{{% /choosable %}}
{{% choosable language go %}}
```go
package main

import (
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewInstance(ctx, "my_instance", &compute.InstanceArgs{
			Name:     pulumi.String("my_instance"),
			Region:   pulumi.String("DFW"),
			ImageId:  pulumi.String("fabe045f-43f8-4991-9e6c-5cabd617538c"),
			FlavorId: pulumi.String("general1-4"),
			KeyPair:  pulumi.String("provisioning_key"),
			Networks: compute.InstanceNetworkArray{
				&compute.InstanceNetworkArgs{
					Uuid: pulumi.String("00000000-0000-0000-0000-000000000000"),
					Name: pulumi.String("public"),
				},
				&compute.InstanceNetworkArgs{
					Uuid: pulumi.String("11111111-1111-1111-1111-111111111111"),
					Name: pulumi.String("private"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
resources:
  myInstance:
    type: openstack:compute:Instance
    name: my_instance
    properties:
      name: my_instance
      region: DFW
      imageId: fabe045f-43f8-4991-9e6c-5cabd617538c
      flavorId: general1-4
      keyPair: provisioning_key
      networks:
        - uuid: 00000000-0000-0000-0000-000000000000
          name: public
        - uuid: 11111111-1111-1111-1111-111111111111
          name: private
```
{{% /choosable %}}
{{% choosable language java %}}
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.openstack.compute.Instance;
import com.pulumi.openstack.compute.InstanceArgs;
import com.pulumi.openstack.compute.inputs.InstanceNetworkArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var myInstance = new Instance("myInstance", InstanceArgs.builder()
            .name("my_instance")
            .region("DFW")
            .imageId("fabe045f-43f8-4991-9e6c-5cabd617538c")
            .flavorId("general1-4")
            .keyPair("provisioning_key")
            .networks(
                InstanceNetworkArgs.builder()
                    .uuid("00000000-0000-0000-0000-000000000000")
                    .name("public")
                    .build(),
                InstanceNetworkArgs.builder()
                    .uuid("11111111-1111-1111-1111-111111111111")
                    .name("private")
                    .build())
            .build());

    }
}
```
{{% /choosable %}}
{{< /chooser >}}

If you try using this provider with Rackspace and run into bugs, you
are welcomed to open a bug report / issue on Github, but please keep
in mind that this is unsupported and the reported bug may not be
able to be fixed.

If you have successfully used this provider with Rackspace and can
add any additional comments, please let us know.