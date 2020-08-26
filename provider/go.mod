module github.com/pulumi/pulumi-openstack/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.7.2
	github.com/pulumi/pulumi/sdk/v2 v2.9.1-0.20200821035132-629254334213
	github.com/terraform-providers/terraform-provider-openstack v1.30.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
