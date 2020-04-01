module github.com/pulumi/pulumi-openstack/provider

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/pulumi/pulumi/sdk v0.0.0-20200325225746-80f1989600a3
	github.com/terraform-providers/terraform-provider-openstack v1.26.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
