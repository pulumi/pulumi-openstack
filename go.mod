module github.com/pulumi/pulumi-openstack

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.6.5
	github.com/terraform-providers/terraform-provider-openstack v1.25.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
