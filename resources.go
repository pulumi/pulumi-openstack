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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-openstack/openstack"

	"github.com/pulumi/pulumi-terraform-bridge/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
)

// all of the OpenStack token components used below.
const (
	// packages:
	openstackPkg = "openstack"
	// modules:
	blockstorageMod     = "BlockStorage"     // Block Storage
	computeMod          = "Compute"          // Compute
	containerinfraMod   = "ContainerInfra"   // Container Infrastructure
	databaseMod         = "Database"         // Database
	dnsMod              = "Dns"              // DNS
	identityMod         = "Identity"         // Identity
	imagesMod           = "Images"           // Images
	keymanagerMod       = "KeyManager"       // Key Manager
	networkingMod       = "Networking"       // Networking
	lbMod               = "LoadBalancer"     // Load Balancer
	firewallMod         = "Firewall"         // Firewall
	osMod               = "ObjectStorage"    // Object Storage
	orchestrationMod    = "Orchestration"    // Orchestration
	sharedfilesystemMod = "SharedFileSystem" // Shared FileSystem
	vpnaasMod           = "VPNaaS"           // VPNaaS
)

var namespaceMap = map[string]string{
	"openstack": "OpenStack",
}

// openstackMember manufactures a type token for the OpenStack package and the given module and type.  It automatically
// uses the OpenStack package and names the file by simply lower casing the resource's first character.
func openstackMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(openstackPkg + ":" + token + ":" + mem)
}

// openstackType manufactures a type token for the OpenStack package and the given module and type.
func openstackType(mod string, typ string) tokens.Type {
	return tokens.Type(openstackMember(mod, typ))
}

// openstackDataSource manufactures a standard resource token given a module and resource name.
func openstackDataSource(mod string, res string) tokens.ModuleMember {
	return openstackMember(mod, res)
}

// openstackResource manufactures a standard resource token given a module and resource name.
func openstackResource(mod string, res string) tokens.Type {
	return openstackType(mod, res)
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
		License:     "Apache-2.0",
		Repository:  "https://github.com/pulumi/pulumi-openstack",
		Config: map[string]*tfbridge.SchemaInfo{
			"auth_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_AUTH_URL"},
				},
			},
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_REGION_NAME"},
				},
			},
			"user_name": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_USERNAME"},
				},
			},
			"user_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_USER_ID"},
				},
			},
			"tenant_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"OS_TENANT_ID",
						"OS_PROJECT_ID",
					},
				},
			},
			"tenant_name": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"OS_TENANT_NAME",
						"OS_PROJECT_NAME",
					},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_PASSWORD"},
				},
			},
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"OS_TOKEN",
						"OS_AUTH_TOKEN",
					},
				},
			},
			"user_domain_name": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_USER_DOMAIN_NAME"},
				},
			},
			"user_domain_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_USER_DOMAIN_ID"},
				},
			},
			"project_domain_name": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_PROJECT_DOMAIN_NAME"},
				},
			},
			"project_domain_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_PROJECT_DOMAIN_ID"},
				},
			},
			"domain_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_DOMAIN_ID"},
				},
			},
			"domain_name": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_DOMAIN_NAME"},
				},
			},
			"default_domain": {
				Default: &tfbridge.DefaultInfo{
					Value:   "default",
					EnvVars: []string{"OS_DEFAULT_DOMAIN"},
				},
			},
			"insecure": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_INSECURE"},
				},
			},
			"endpoint_type": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_ENDPOINT_TYPE"},
				},
			},
			"cacert_file": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_CACERT"},
				},
			},
			"cert": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_CERT"},
				},
			},
			"key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_KEY"},
				},
			},
			"swauth": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_SWAUTH"},
				},
			},
			"use_octavia": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_USE_OCTAVIA"},
				},
			},
			"cloud": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_CLOUD"},
				},
			},
			"application_credential_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_APPLICATION_CREDENTIAL_ID"},
				},
			},
			"application_credential_name": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_APPLICATION_CREDENTIAL_NAME"},
				},
			},
			"application_credential_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_APPLICATION_CREDENTIAL_SECRET"},
				},
			},
			"delayed_auth": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_DELAYED_AUTH"},
				},
			},
			"allow_reauth": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_ALLOW_REAUTH"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Block Storage
			"openstack_blockstorage_volume_v1":        {Tok: openstackResource(blockstorageMod, "VolumeV1")},
			"openstack_blockstorage_volume_v2":        {Tok: openstackResource(blockstorageMod, "VolumeV2")},
			"openstack_blockstorage_volume_attach_v2": {Tok: openstackResource(blockstorageMod, "VolumeAttachV2")},
			"openstack_blockstorage_volume_v3":        {Tok: openstackResource(blockstorageMod, "Volume")},
			"openstack_blockstorage_volume_attach_v3": {Tok: openstackResource(blockstorageMod, "VolumeAttach")},
			"openstack_blockstorage_quotaset_v2":      {Tok: openstackResource(blockstorageMod, "QuoteSetV2")},
			"openstack_blockstorage_quotaset_v3":      {Tok: openstackResource(blockstorageMod, "QuoteSetV3")},

			// Compute
			"openstack_compute_flavor_v2":               {Tok: openstackResource(computeMod, "Flavor")},
			"openstack_compute_flavor_access_v2":        {Tok: openstackResource(computeMod, "FlavorAccess")},
			"openstack_compute_floatingip_v2":           {Tok: openstackResource(computeMod, "FloatingIp")},
			"openstack_compute_floatingip_associate_v2": {Tok: openstackResource(computeMod, "FloatingIpAssociate")},
			"openstack_compute_instance_v2":             {Tok: openstackResource(computeMod, "Instance")},
			"openstack_compute_interface_attach_v2":     {Tok: openstackResource(computeMod, "InterfaceAttach")},
			"openstack_compute_keypair_v2":              {Tok: openstackResource(computeMod, "Keypair")},
			"openstack_compute_secgroup_v2":             {Tok: openstackResource(computeMod, "SecGroup")},
			"openstack_compute_servergroup_v2":          {Tok: openstackResource(computeMod, "ServerGroup")},
			"openstack_compute_volume_attach_v2":        {Tok: openstackResource(computeMod, "VolumeAttach")},
			"openstack_compute_quotaset_v2":             {Tok: openstackResource(computeMod, "QuotaSetV2")},

			// Container Infrastructure
			"openstack_containerinfra_cluster_v1":         {Tok: openstackResource(containerinfraMod, "Cluster")},
			"openstack_containerinfra_clustertemplate_v1": {Tok: openstackResource(containerinfraMod, "ClusterTemplate")},

			// Database
			"openstack_db_instance_v1":      {Tok: openstackResource(databaseMod, "Instance")},
			"openstack_db_database_v1":      {Tok: openstackResource(databaseMod, "Database")},
			"openstack_db_user_v1":          {Tok: openstackResource(databaseMod, "User")},
			"openstack_db_configuration_v1": {Tok: openstackResource(databaseMod, "Configuration")},

			// DNS
			"openstack_dns_recordset_v2": {Tok: openstackResource(dnsMod, "RecordSet")},
			"openstack_dns_zone_v2":      {Tok: openstackResource(dnsMod, "Zone")},

			// Identity
			"openstack_identity_application_credential_v3": {Tok: openstackResource(identityMod, "ApplicationCredential")},
			"openstack_identity_project_v3":                {Tok: openstackResource(identityMod, "Project")},
			"openstack_identity_role_v3":                   {Tok: openstackResource(identityMod, "Role")},
			"openstack_identity_role_assignment_v3":        {Tok: openstackResource(identityMod, "RoleAssignment")},
			"openstack_identity_user_v3":                   {Tok: openstackResource(identityMod, "User")},
			"openstack_identity_endpoint_v3":               {Tok: openstackResource(identityMod, "EndpointV3")},
			"openstack_identity_service_v3":                {Tok: openstackResource(identityMod, "ServiceV3")},

			// Images
			"openstack_images_image_v2":               {Tok: openstackResource(imagesMod, "Image")},
			"openstack_images_image_access_v2":        {Tok: openstackResource(imagesMod, "ImageAccess")},
			"openstack_images_image_access_accept_v2": {Tok: openstackResource(imagesMod, "ImageAccessAccept")},

			// Networking
			"openstack_networking_addressscope_v2":            {Tok: openstackResource(networkingMod, "AddressScope")},
			"openstack_networking_floatingip_v2":              {Tok: openstackResource(networkingMod, "FloatingIp")},
			"openstack_networking_floatingip_associate_v2":    {Tok: openstackResource(networkingMod, "FloatingIpAssociate")},
			"openstack_networking_network_v2":                 {Tok: openstackResource(networkingMod, "Network")},
			"openstack_networking_port_v2":                    {Tok: openstackResource(networkingMod, "Port")},
			"openstack_networking_port_secgroup_associate_v2": {Tok: openstackResource(networkingMod, "PortSecGroupAssociate")},
			"openstack_networking_router_interface_v2":        {Tok: openstackResource(networkingMod, "RouterInterface")},
			"openstack_networking_router_route_v2":            {Tok: openstackResource(networkingMod, "RouterRoute")},
			"openstack_networking_router_v2":                  {Tok: openstackResource(networkingMod, "Router")},
			"openstack_networking_subnet_v2": {
				Tok: openstackResource(networkingMod, "Subnet"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"allocation_pool": {
						Name: "allocationPools",
					},
					"allocation_pools": {
						Name: "allocationPoolsCollection",
					},
				},
			},
			"openstack_networking_subnet_route_v2":  {Tok: openstackResource(networkingMod, "SubnetRoute")},
			"openstack_networking_subnetpool_v2":    {Tok: openstackResource(networkingMod, "SubnetPool")},
			"openstack_networking_secgroup_v2":      {Tok: openstackResource(networkingMod, "SecGroup")},
			"openstack_networking_secgroup_rule_v2": {Tok: openstackResource(networkingMod, "SecGroupRule")},
			"openstack_networking_trunk_v2":         {Tok: openstackResource(networkingMod, "Trunk")},
			"openstack_networking_qos_bandwidth_limit_rule_v2": {
				Tok: openstackResource(networkingMod, "QosBandwidthLimitRule"),
			},
			"openstack_networking_qos_dscp_marking_rule_v2": {
				Tok: openstackResource(networkingMod, "QosDscpMarkingRule"),
			},
			"openstack_networking_qos_minimum_bandwidth_rule_v2": {
				Tok: openstackResource(networkingMod, "QosMinimumBandwidthRule"),
			},
			"openstack_networking_qos_policy_v2":  {Tok: openstackResource(networkingMod, "QosPolicy")},
			"openstack_networking_rbac_policy_v2": {Tok: openstackResource(networkingMod, "RbacPolicyV2")},
			"openstack_networking_quota_v2":       {Tok: openstackResource(networkingMod, "QuotaV2")},

			// Load Balancer
			"openstack_lb_member_v1":       {Tok: openstackResource(lbMod, "MemberV1")},
			"openstack_lb_monitor_v1":      {Tok: openstackResource(lbMod, "MonitorV1")},
			"openstack_lb_pool_v1":         {Tok: openstackResource(lbMod, "PoolV1")},
			"openstack_lb_vip_v1":          {Tok: openstackResource(lbMod, "Vip")},
			"openstack_lb_l7policy_v2":     {Tok: openstackResource(lbMod, "L7PolicyV2")},
			"openstack_lb_l7rule_v2":       {Tok: openstackResource(lbMod, "L7RuleV2")},
			"openstack_lb_loadbalancer_v2": {Tok: openstackResource(lbMod, "LoadBalancer")},
			"openstack_lb_listener_v2":     {Tok: openstackResource(lbMod, "Listener")},
			"openstack_lb_pool_v2":         {Tok: openstackResource(lbMod, "Pool")},
			"openstack_lb_member_v2":       {Tok: openstackResource(lbMod, "Member")},
			"openstack_lb_monitor_v2":      {Tok: openstackResource(lbMod, "Monitor")},

			// Firewall
			"openstack_fw_firewall_v1": {Tok: openstackResource(firewallMod, "Firewall")},
			"openstack_fw_policy_v1":   {Tok: openstackResource(firewallMod, "Policy")},
			"openstack_fw_rule_v1":     {Tok: openstackResource(firewallMod, "Rule")},

			// Object Storage
			"openstack_objectstorage_container_v1": {Tok: openstackResource(osMod, "Container")},
			"openstack_objectstorage_object_v1":    {Tok: openstackResource(osMod, "ContainerObject")},
			"openstack_objectstorage_tempurl_v1":   {Tok: openstackResource(osMod, "TempUrl")},

			// Shared Filesystem
			"openstack_sharedfilesystem_securityservice_v2": {Tok: openstackResource(sharedfilesystemMod, "SecurityService")},
			"openstack_sharedfilesystem_share_v2":           {Tok: openstackResource(sharedfilesystemMod, "Share")},
			"openstack_sharedfilesystem_share_access_v2":    {Tok: openstackResource(sharedfilesystemMod, "ShareAccess")},
			"openstack_sharedfilesystem_sharenetwork_v2":    {Tok: openstackResource(sharedfilesystemMod, "ShareNetwork")},

			// VPNaaS
			"openstack_vpnaas_ipsec_policy_v2":    {Tok: openstackResource(vpnaasMod, "IpSecPolicy")},
			"openstack_vpnaas_ike_policy_v2":      {Tok: openstackResource(vpnaasMod, "IkePolicy")},
			"openstack_vpnaas_service_v2":         {Tok: openstackResource(vpnaasMod, "Service")},
			"openstack_vpnaas_endpoint_group_v2":  {Tok: openstackResource(vpnaasMod, "EndpointGroup")},
			"openstack_vpnaas_site_connection_v2": {Tok: openstackResource(vpnaasMod, "SiteConnection")},

			// KeyManager
			"openstack_keymanager_container_v1": {Tok: openstackResource(keymanagerMod, "ContainerV1")},
			"openstack_keymanager_secret_v1":    {Tok: openstackResource(keymanagerMod, "SecretV1")},

			// Orchestration
			"openstack_orchestration_stack_v1": {Tok: openstackResource(orchestrationMod, "StackV1")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Block Storage
			"openstack_blockstorage_availability_zones_v3": {
				Tok: openstackDataSource(blockstorageMod, "getAvailabilityZonesV3"),
			},
			"openstack_blockstorage_snapshot_v2": {Tok: openstackDataSource(blockstorageMod, "getSnapshotV2")},
			"openstack_blockstorage_snapshot_v3": {Tok: openstackDataSource(blockstorageMod, "getSnapshotV3")},
			"openstack_blockstorage_volume_v2":   {Tok: openstackDataSource(blockstorageMod, "getVolumeV2")},
			"openstack_blockstorage_volume_v3":   {Tok: openstackDataSource(blockstorageMod, "getVolumeV3")},

			// Compute
			"openstack_compute_availability_zones_v2": {Tok: openstackDataSource(computeMod, "getAvailabilityZones")},
			"openstack_compute_flavor_v2":             {Tok: openstackDataSource(computeMod, "getFlavor")},
			"openstack_compute_keypair_v2":            {Tok: openstackDataSource(computeMod, "getKeypair")},

			// Container Infrastructure
			"openstack_containerinfra_cluster_v1":         {Tok: openstackDataSource(containerinfraMod, "getCluster")},
			"openstack_containerinfra_clustertemplate_v1": {Tok: openstackDataSource(containerinfraMod, "getClusterTemplate")},

			// DNS
			"openstack_dns_zone_v2": {Tok: openstackDataSource(dnsMod, "getDnsZone")},

			// Identity
			"openstack_identity_project_v3":    {Tok: openstackDataSource(identityMod, "getProject")},
			"openstack_identity_role_v3":       {Tok: openstackDataSource(identityMod, "getRole")},
			"openstack_identity_user_v3":       {Tok: openstackDataSource(identityMod, "getUser")},
			"openstack_identity_auth_scope_v3": {Tok: openstackDataSource(identityMod, "getAuthScope")},
			"openstack_identity_endpoint_v3":   {Tok: openstackDataSource(identityMod, "getEndpoint")},
			"openstack_identity_group_v3":      {Tok: openstackDataSource(identityMod, "getGroup")},
			"openstack_identity_service_v3":    {Tok: openstackDataSource(identityMod, "getService")},

			// Images
			"openstack_images_image_v2": {Tok: openstackDataSource(imagesMod, "getImage")},

			// Networking
			"openstack_networking_network_v2":      {Tok: openstackDataSource(networkingMod, "getNetwork")},
			"openstack_networking_port_v2":         {Tok: openstackDataSource(networkingMod, "getPort")},
			"openstack_networking_port_ids_v2":     {Tok: openstackDataSource(networkingMod, "getPortIds")},
			"openstack_networking_router_v2":       {Tok: openstackDataSource(networkingMod, "getRouter")},
			"openstack_networking_secgroup_v2":     {Tok: openstackDataSource(networkingMod, "getSecGroup")},
			"openstack_networking_subnet_v2":       {Tok: openstackDataSource(networkingMod, "getSubnet")},
			"openstack_networking_subnetpool_v2":   {Tok: openstackDataSource(networkingMod, "getSubnetPool")},
			"openstack_networking_trunk_v2":        {Tok: openstackDataSource(networkingMod, "getTrunk")},
			"openstack_networking_floatingip_v2":   {Tok: openstackDataSource(networkingMod, "getFloatingIp")},
			"openstack_networking_addressscope_v2": {Tok: openstackDataSource(networkingMod, "getAddressScope")},
			"openstack_networking_qos_bandwidth_limit_rule_v2": {
				Tok: openstackDataSource(networkingMod, "getQosBandwidthLimitRule"),
			},
			"openstack_networking_qos_dscp_marking_rule_v2": {
				Tok: openstackDataSource(networkingMod, "getQosDscpMarkingRule"),
			},
			"openstack_networking_qos_minimum_bandwidth_rule_v2": {
				Tok: openstackDataSource(networkingMod, "getQosMinimumBandwidthRule"),
			},
			"openstack_networking_qos_policy_v2": {Tok: openstackDataSource(networkingMod, "getQosPolicy")},

			// Firewall
			"openstack_fw_policy_v1": {Tok: openstackDataSource(firewallMod, "getPolicy")},

			// Shared Filesystem
			"openstack_sharedfilesystem_availability_zones_v2": {
				Tok: openstackDataSource(sharedfilesystemMod, "getAvailbilityZones"),
			},
			"openstack_sharedfilesystem_share_v2":        {Tok: openstackDataSource(sharedfilesystemMod, "getShare")},
			"openstack_sharedfilesystem_sharenetwork_v2": {Tok: openstackDataSource(sharedfilesystemMod, "getShareNetwork")},
			"openstack_sharedfilesystem_snapshot_v2":     {Tok: openstackDataSource(sharedfilesystemMod, "getSnapshot")},

			// Keymanager
			"openstack_keymanager_secret_v1":    {Tok: openstackDataSource(keymanagerMod, "getSecret")},
			"openstack_keymanager_container_v1": {Tok: openstackDataSource(keymanagerMod, "getContainer")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^1.0.0",
			},
			Overlay: &tfbridge.OverlayInfo{
				Files:   []string{},
				Modules: map[string]*tfbridge.OverlayInfo{},
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "1.7.0-preview",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: namespaceMap,
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const openstackName = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[openstackName]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[openstackName]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					// Use conservative options that apply broadly for OpenStack.
					res.Fields[openstackName] = tfbridge.AutoName(openstackName, 255)
				}
			}
		}
	}

	return prov
}
