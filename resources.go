// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openstack

import (
	"strings"
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-openstack/openstack"
)

// all of the OpenStack token components used below.
const (
	// packages:
	openstackPkg = "openstack"
	// modules:
	blockstorageMod        = "blockstorage"  // Block Storage
	computeMod             = "compute"       // Compute
	databaseMod            = "database"      // Database
	dnsMod                 = "dns"           // DNS
	identityMod            = "identity"      // Identity
	imagesMod              = "images"        // Images
	networkingMod          = "networking"    // Networking
	lbMod                  = "loadbalancer"  // Load Balancer
	firewallMod            = "firewall"      // Firewall
	osMod                  = "objectstorage" // Object Storage
	vpnaasMod              = "vpnaas"        // VPNaaS
)

// openstackMember manufactures a type token for the OpenStack package and the given module and type.
func openstackMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(openstackPkg + ":" + mod + ":" + mem)
}

// openstackType manufactures a type token for the OpenStack package and the given module and type.
func openstackType(mod string, typ string) tokens.Type {
	return tokens.Type(openstackMember(mod, typ))
}

// openstackDataSource manufactures a standard resource token given a module and resource name.  It automatically uses the OpenStack
// package and names the file by simply lower casing the data source's first character.
func openstackDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return openstackMember(mod+"/"+fn, res)
}

// openstackResource manufactures a standard resource token given a module and resource name.  It automatically uses the OpenStack
// package and names the file by simply lower casing the resource's first character.
func openstackResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return openstackType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// Provider returns additional overlaid schema and metadata associated with the openstack package.
func Provider() tfbridge.ProviderInfo {
	p := openstack.Provider().(*schema.Provider)

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "openstack",
		Description: "A Pulumi package for creating and managing OpenStack cloud resources.",
		Keywords:    []string{"pulumi", "openstack"},
		Homepage:    "https://pulumi.io",
		License:     "Apache 2.0",
		Repository:  "https://github.com/pulumi/pulumi-openstack",
		Resources: map[string]*tfbridge.ResourceInfo{
			// Block Storage
			"openstack_blockstorage_volume_v1": {Tok: openstackResource(blockstorageMod, "VolumeV1")},
			"openstack_blockstorage_volume_v2": {Tok: openstackResource(blockstorageMod, "VolumeV2")},
			"openstack_blockstorage_volume_attach_v2": {Tok: openstackResource(blockstorageMod, "VolumeAttachV2")},
			"openstack_blockstorage_volume_v3": {Tok: openstackResource(blockstorageMod, "Volume")},
			"openstack_blockstorage_volume_attach_v3": {Tok: openstackResource(blockstorageMod, "VolumeAttach")},

			// Compute
			"openstack_compute_flavor_v2": {Tok: openstackResource(computeMod, "Flavor")},
			"openstack_compute_floatingip_v2": {Tok: openstackResource(computeMod, "FloatingIP")},
			"openstack_compute_floatingip_associate_v2": {Tok: openstackResource(computeMod, "FloatingIpAssociate")},
			"openstack_compute_instance_v2": {Tok: openstackResource(computeMod, "Instance")},
			"openstack_compute_keypair_v2": {Tok: openstackResource(computeMod, "Keypair")},
			"openstack_compute_secgroup_v2": {Tok: openstackResource(computeMod, "SecGroup")},
			"openstack_compute_servergroup_v2": {Tok: openstackResource(computeMod, "ServerGroup")},
			"openstack_compute_volume_attach_v2": {Tok: openstackResource(computeMod, "VolumeAttach")},

			// Database
			"openstack_db_instance_v1": {Tok: openstackResource(databaseMod, "Instance")},
			"openstack_db_database_v1": {Tok: openstackResource(databaseMod, "Database")},
			"openstack_db_user_v1": {Tok: openstackResource(databaseMod, "User")},
			"openstack_db_configuration_v1": {Tok: openstackResource(databaseMod, "Configuration")},

			// DNS
			"openstack_dns_recordset_v2": {Tok: openstackResource(dnsMod, "RecordSet")},
			"openstack_dns_zone_v2": {Tok: openstackResource(dnsMod, "Zone")},

			// Identity
			"openstack_identity_project_v3": {Tok: openstackResource(identityMod, "Project")},
			"openstack_identity_role_v3": {Tok: openstackResource(identityMod, "Role")},
			"openstack_identity_role_assignment_v3": {Tok: openstackResource(identityMod, "RoleAssignment")},
			"openstack_identity_user_v3": {Tok: openstackResource(identityMod, "User")},

			// Images
			"openstack_images_image_v2": {Tok: openstackResource(imagesMod, "Image")},

			// Networking
			"openstack_networking_floatingip_v2": {Tok: openstackResource(networkingMod, "FloatingIp")},
			"openstack_networking_floatingip_associate_v2": {Tok: openstackResource(networkingMod, "FloatingIpAssociate")},
			"openstack_networking_network_v2": {Tok: openstackResource(networkingMod, "Network")},
			"openstack_networking_port_v2": {Tok: openstackResource(networkingMod, "Port")},
			"openstack_networking_router_interface_v2": {Tok: openstackResource(networkingMod, "RouterInterface")},
			"openstack_networking_router_route_v2": {Tok: openstackResource(networkingMod, "RouterRoute")},
			"openstack_networking_router_v2": {Tok: openstackResource(networkingMod, "Router")},
			"openstack_networking_subnet_v2": {Tok: openstackResource(networkingMod, "Subnet")},
			"openstack_networking_subnet_route_v2": {Tok: openstackResource(networkingMod, "SubnetRoute")},
			"openstack_networking_subnetpool_v2": {Tok: openstackResource(networkingMod, "SubnetPool")},
			"openstack_networking_secgroup_v2": {Tok: openstackResource(networkingMod, "SecGroup")},
			"openstack_networking_secgroup_rule_v2": {Tok: openstackResource(networkingMod, "SecGroupRule")},

			// Load Balancer
			"openstack_lb_member_v1": {Tok: openstackResource(lbMod, "MemberV1")},
			"openstack_lb_monitor_v1": {Tok: openstackResource(lbMod, "MonitorV1")},
			"openstack_lb_pool_v1": {Tok: openstackResource(lbMod, "PoolV1")},
			"openstack_lb_vip_v1": {Tok: openstackResource(lbMod, "Vip")},
			"openstack_lb_loadbalancer_v2": {Tok: openstackResource(lbMod, "LoadBalancer")},
			"openstack_lb_listener_v2": {Tok: openstackResource(lbMod, "Listener")},
			"openstack_lb_pool_v2": {Tok: openstackResource(lbMod, "Pool")},
			"openstack_lb_member_v2": {Tok: openstackResource(lbMod, "Member")},
			"openstack_lb_monitor_v2": {Tok: openstackResource(lbMod, "Monitor")},

			// Firewall
			"openstack_fw_firewall_v1": {Tok: openstackResource(firewallMod, "Firewall")},
			"openstack_fw_policy_v1": {Tok: openstackResource(firewallMod, "Policy")},
			"openstack_fw_rule_v1": {Tok: openstackResource(firewallMod, "Rule")},

			// Object Storage
			"openstack_objectstorage_container_v1": {Tok: openstackResource(osMod, "Container")},
			"openstack_objectstorage_object_v1": {Tok: openstackResource(osMod, "ContainerObject")},

			// VPNaaS
			"openstack_vpnaas_ipsec_policy_v2": {Tok: openstackResource(vpnaasMod, "IpsecPolicy")},
			"openstack_vpnaas_ike_policy_v2": {Tok: openstackResource(vpnaasMod, "IkePolicy")},
			"openstack_vpnaas_service_v2": {Tok: openstackResource(vpnaasMod, "Service")},
			"openstack_vpnaas_endpoint_group_v2": {Tok: openstackResource(vpnaasMod, "EndpointGroup")},
			"openstack_vpnaas_site_connection_v2": {Tok: openstackResource(vpnaasMod, "SiteConnection")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Compute
			"openstack_compute_flavor_v2": {Tok: openstackDataSource(computeMod, "getFlavor")},
			"openstack_compute_keypair_v2": {Tok: openstackDataSource(computeMod, "getKeypair")},

			// DNS
			"openstack_dns_zone_v2": {Tok: openstackDataSource(dnsMod, "getDnsZone")},

			// Identity
			"openstack_identity_project_v3": {Tok: openstackDataSource(identityMod, "getProject")},
			"openstack_identity_role_v3": {Tok: openstackDataSource(identityMod, "getRole")},
			"openstack_identity_user_v3": {Tok: openstackDataSource(identityMod, "getUser")},
			"openstack_identity_auth_scope_v3": {Tok: openstackDataSource(identityMod, "getAuthScope")},

			// Images
			"openstack_images_image_v2": {Tok: openstackDataSource(imagesMod, "getImage")},

			// Networking
			"openstack_networking_network_v2": {Tok: openstackDataSource(networkingMod, "getNetwork")},
			"openstack_networking_secgroup_v2": {Tok: openstackDataSource(networkingMod, "getSecGroup")},
			"openstack_networking_subnet_v2": {Tok: openstackDataSource(networkingMod, "getSubnet")},
			"openstack_networking_subnetpool_v2": {Tok: openstackDataSource(networkingMod, "getSubnetPool")},

		},
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^0.14.0-rc1",
			},
			Overlay: &tfbridge.OverlayInfo{
				Files:   []string{},
				Modules: map[string]*tfbridge.OverlayInfo{},
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=0.14.0rc1,<0.15.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const openstackName = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			if _, has := schema.Schema[openstackName]; has {
				if _, hasfield := res.Fields[openstackName]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					// Use conservative options that apply broadly for OpenStack.
					res.Fields[openstackName] = AutoName(openstackName, AutoNameOptions{
						ForceLowercase: true,
						Separator:      "",
						Maxlen:         24,
						Randlen:        8,
					})
				}
			}
		}
	}

	return prov
}

// IDEA: Consider moving this refactoring of AutoName to allow more flexible configuration back into pulumi-terraform.

// AutoNameOptions provides parameters to AutoName to control how names will be generated
type AutoNameOptions struct {
	// A separator between name and random portions of the
	Separator string
	// Maximum length of the generated name
	Maxlen int
	// Number of characters of random hex digits to add to the name
	Randlen int
	// A transform to apply to the name prior to adding random characters
	Transform func(string) string
	// Force the name to be lowercase prior to adding random characters
	ForceLowercase bool
}

// AutoName creates custom schema for a Terraform name property which is automatically populated from the resource's URN
// name, and tranformed based on the provided options.
func AutoName(name string, options AutoNameOptions) *tfbridge.SchemaInfo {
	return &tfbridge.SchemaInfo{
		Name: name,
		Default: &tfbridge.DefaultInfo{
			From: FromName(options),
		},
	}
}

// FromName automatically propagates a resource's URN onto the resulting default info.
func FromName(options AutoNameOptions) func(res *tfbridge.PulumiResource) (interface{}, error) {
	return func(res *tfbridge.PulumiResource) (interface{}, error) {
		// Take the URN name part, transform it if required, and then append some unique characters if requested.
		vs := string(res.URN.Name())
		if options.Transform != nil {
			vs = options.Transform(vs)
		} else if options.ForceLowercase {
			vs = strings.ToLower(vs)
		}
		if options.Randlen > 0 {
			return resource.NewUniqueHex(vs+options.Separator, options.Randlen, options.Maxlen)
		}
		if len(vs) > options.Maxlen {
			return "", errors.Errorf("name '%s' is longer than maximum length %d", vs, options.Maxlen)
		}
		return vs, nil
	}
}
