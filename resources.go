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
	// modules; in general, we took naming inspiration from the OpenStack SDK for Go:
	// https://godoc.org/github.com/OpenStack/openstack-sdk-for-go
	openstackAppInsights         = "appinsights"         // AppInsights
	openstackAppService          = "appservice"          // App Service
	openstackAutomation          = "automation"          // Automatio
	openstackCDN                 = "cdn"                 // CDN
	openstackCompute             = "compute"             // Virtual Machinesn
	openstackContainerService    = "containerservice"    // OpenStack Container Service
	openstackCore                = "core"                // Base Resources
	openstackCosmosDB            = "cosmosdb"            // Cosmos DB
	openstackDatalake            = "datalake"            // Data Lake
	openstackDNS                 = "dns"                 // DNS
	openstackMessaging           = "eventhub"            // Event Hub
	openstackFunctions           = "functions"           // Functions
	openstackKeyVault            = "keyvault"            // Key Vault
	openstackLB                  = "lb"                  // Load Balancer
	openstackIot                 = "iot"                 // IoT resource
	openstackMgmtResource        = "managementresource"  // Management Resource
	openstackMonitoring          = "monitoring"          // Metrics/monitoring resources
	openstackMySQL               = "mysql"               // MySql
	openstackNetwork             = "network"             // Networking
	openstackNetworkWatcher      = "networkwatcher"      // Network Watcher
	openstackOperationalInsights = "operationalinsights" // Operational Insights
	openstackPostgresql          = "postgresql"          // Postgress SQL
	openstackPolicy              = "policy"              // Policy
	openstackRecoveryServices    = "recoveryservices"    // Recovery Services
	openstackRedis               = "redis"               // RedisCache
	openstackRelay               = "relay"               // Relay
	openstackScheduler           = "scheduler"           // Scheduler
	openstackRole                = "role"                // OpenStack Role
	openstackSearch              = "search"              // Search
	openstackSQL                 = "sql"                 // SQL
	openstackStorage             = "storage"             // Storage
	openstackTrafficManager      = "trafficmanager"      // Traffic Manager
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
	p := openstackrm.Provider().(*schema.Provider)

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "openstackrm",
		Description: "A Pulumi package for creating and managing Microsoft OpenStack cloud resources.",
		Keywords:    []string{"pulumi", "openstack"},
		Homepage:    "https://pulumi.io",
		License:     "Apache 2.0",
		Repository:  "https://github.com/pulumi/pulumi-openstack",
		Resources: map[string]*tfbridge.ResourceInfo{
			// AppInsights
			"openstackrm_application_insights": {Tok: openstackResource(openstackAppInsights, "Insights")},

			// App Service
			"openstackrm_app_service":                         {Tok: openstackResource(openstackAppService, "AppService")},
			"openstackrm_app_service_custom_hostname_binding": {Tok: openstackResource(openstackAppService, "CustomHostnameBinding")},
			"openstackrm_app_service_plan":                    {Tok: openstackResource(openstackAppService, "Plan")},
			"openstackrm_app_service_slot":                    {Tok: openstackResource(openstackAppService, "Slot")},
			"openstackrm_app_service_active_slot":             {Tok: openstackResource(openstackAppService, "ActiveSlot")},
			"openstackrm_function_app":                        {Tok: openstackResource(openstackAppService, "FunctionApp")},

			// Automation
			"openstackrm_automation_account":    {Tok: openstackResource(openstackAutomation, "Account")},
			"openstackrm_automation_credential": {Tok: openstackResource(openstackAutomation, "Credential")},
			"openstackrm_automation_runbook":    {Tok: openstackResource(openstackAutomation, "RunBook")},
			"openstackrm_automation_schedule":   {Tok: openstackResource(openstackAutomation, "Schedule")},

			// Authorization
			"openstackrm_role_assignment": {Tok: openstackResource(openstackRole, "assignment")},
			"openstackrm_role_definition": {Tok: openstackResource(openstackRole, "Definition")},

			// OpenStack Container Service
			"openstackrm_container_registry": {Tok: openstackResource(openstackContainerService, "Registry")},
			"openstackrm_container_service":  {Tok: openstackResource(openstackContainerService, "Service")},
			"openstackrm_container_group":    {Tok: openstackResource(openstackContainerService, "Group")},
			"openstackrm_kubernetes_cluster": {Tok: openstackResource(openstackContainerService, "KubernetesCluster")},

			// Core
			"openstackrm_resource_group":      {Tok: openstackResource(openstackCore, "ResourceGroup")},
			"openstackrm_template_deployment": {Tok: openstackResource(openstackCore, "TemplateDeployment")},

			// CDN
			"openstackrm_cdn_endpoint": {Tok: openstackResource(openstackCDN, "Endpoint")},
			"openstackrm_cdn_profile":  {Tok: openstackResource(openstackCDN, "Profile")},

			// Compute
			"openstackrm_availability_set":          {Tok: openstackResource(openstackCompute, "AvailabilitySet")},
			"openstackrm_virtual_machine_extension": {Tok: openstackResource(openstackCompute, "Extension")},
			"openstackrm_virtual_machine":           {Tok: openstackResource(openstackCompute, "VirtualMachine")},
			"openstackrm_virtual_machine_scale_set": {Tok: openstackResource(openstackCompute, "ScaleSet")},
			"openstackrm_managed_disk":              {Tok: openstackResource(openstackCompute, "ManagedDisk")},
			"openstackrm_snapshot":                  {Tok: openstackResource(openstackCompute, "Snapshot")},
			"openstackrm_image":                     {Tok: openstackResource(openstackCompute, "Image")},

			// Data Lake
			"openstackrm_data_lake_store": {Tok: openstackResource(openstackDatalake, "Store")},

			// DNS
			"openstackrm_dns_a_record":     {Tok: openstackResource(openstackDNS, "ARecord")},
			"openstackrm_dns_aaaa_record":  {Tok: openstackResource(openstackDNS, "AaaaRecord")},
			"openstackrm_dns_cname_record": {Tok: openstackResource(openstackDNS, "CNameRecord")},
			"openstackrm_dns_mx_record":    {Tok: openstackResource(openstackDNS, "MxRecord")},
			"openstackrm_dns_ns_record":    {Tok: openstackResource(openstackDNS, "NsRecord")},
			"openstackrm_dns_ptr_record":   {Tok: openstackResource(openstackDNS, "PtrRecord")},
			"openstackrm_dns_srv_record":   {Tok: openstackResource(openstackDNS, "SrvRecord")},
			"openstackrm_dns_txt_record":   {Tok: openstackResource(openstackDNS, "TxtRecord")},
			"openstackrm_dns_zone":         {Tok: openstackResource(openstackDNS, "Zone")},

			// Messaging
			"openstackrm_eventgrid_topic":                     {Tok: openstackResource(openstackMessaging, "EventGridTopic")},
			"openstackrm_eventhub":                            {Tok: openstackResource(openstackMessaging, "EventHub")},
			"openstackrm_eventhub_authorization_rule":         {Tok: openstackResource(openstackMessaging, "EventHubAuthorizationRule")},
			"openstackrm_eventhub_consumer_group":             {Tok: openstackResource(openstackMessaging, "EventHubConsumerGroup")},
			"openstackrm_eventhub_namespace":                  {Tok: openstackResource(openstackMessaging, "EventHubNamespace")},
			"openstackrm_servicebus_namespace":                {Tok: openstackResource(openstackMessaging, "Namespace")},
			"openstackrm_servicebus_queue":                    {Tok: openstackResource(openstackMessaging, "Queue")},
			"openstackrm_servicebus_subscription":             {Tok: openstackResource(openstackMessaging, "Subscription")},
			"openstackrm_servicebus_subscription_rule":        {Tok: openstackResource(openstackMessaging, "SubscriptionRule")},
			"openstackrm_servicebus_topic":                    {Tok: openstackResource(openstackMessaging, "Topic")},
			"openstackrm_servicebus_topic_authorization_rule": {Tok: openstackResource(openstackMessaging, "TopicAuthorizationRule")},

			// IoT Resources
			"openstackrm_iothub": {Tok: openstackResource(openstackIot, "IoTHub"),
				Docs: &tfbridge.DocInfo{
					Source: "iothub.html.markdown",
				},
			},

			// KeyVault
			"openstackrm_key_vault": {
				Tok: openstackResource(openstackKeyVault, "KeyVault"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Ensure "sku" is a singleton
					"sku": {Name: "sku", MaxItemsOne: boolRef(true)},
				},
			},
			"openstackrm_key_vault_certificate": {Tok: openstackResource(openstackKeyVault, "Certifiate")},
			"openstackrm_key_vault_key":         {Tok: openstackResource(openstackKeyVault, "Key")},
			"openstackrm_key_vault_secret":      {Tok: openstackResource(openstackKeyVault, "Secret")},

			// LoadBalancer
			"openstackrm_lb": {
				Tok: openstackResource(openstackLB, "LoadBalancer"),
				Docs: &tfbridge.DocInfo{
					Source: "loadbalancer.html.markdown",
				},
			},
			"openstackrm_lb_backend_address_pool": {Tok: openstackResource(openstackLB, "BackendAddressPool"),
				Docs: &tfbridge.DocInfo{
					Source: "loadbalancer_backend_address_pool.html.markdown",
				},
			},
			"openstackrm_lb_nat_rule": {Tok: openstackResource(openstackLB, "NatRule"),
				Docs: &tfbridge.DocInfo{
					Source: "loadbalancer_nat_rule.html.markdown",
				},
			},
			"openstackrm_lb_nat_pool": {Tok: openstackResource(openstackLB, "NatPool"),
				Docs: &tfbridge.DocInfo{
					Source: "loadbalancer_nat_pool.html.markdown",
				}},
			"openstackrm_lb_probe": {Tok: openstackResource(openstackLB, "Probe"),
				Docs: &tfbridge.DocInfo{
					Source: "loadbalancer_probe.html.markdown",
				},
			},
			"openstackrm_lb_rule": {Tok: openstackResource(openstackLB, "Rule"),
				Docs: &tfbridge.DocInfo{
					Source: "loadbalancer_rule.html.markdown",
				},
			},

			// Operational Insights
			"openstackrm_log_analytics_workspace": {Tok: openstackResource(openstackOperationalInsights, "AnalyticsWorkspace")},
			"openstackrm_log_analytics_solution": {Tok: openstackResource(openstackOperationalInsights, "AnalyticsSolution"),
				Docs: &tfbridge.DocInfo{
					Source: "log_analytics_solution.html.markdown",
				}},

			// CosmosDB
			"openstackrm_cosmosdb_account": {Tok: openstackResource(openstackCosmosDB, "Account")},

			// Management Resource
			"openstackrm_management_lock": {Tok: openstackResource(openstackMgmtResource, "ManangementLock")},

			// Monitoring resources
			"openstackrm_metric_alertrule": {Tok: openstackResource(openstackMonitoring, "AlertRule")},

			// MySQL
			"openstackrm_mysql_configuration": {Tok: openstackResource(openstackMySQL, "Configuration")},
			"openstackrm_mysql_database":      {Tok: openstackResource(openstackMySQL, "Database")},
			"openstackrm_mysql_firewall_rule": {Tok: openstackResource(openstackMySQL, "FirewallRule")},
			"openstackrm_mysql_server":        {Tok: openstackResource(openstackMySQL, "Server")},

			// Postgress SQL
			"openstackrm_postgresql_configuration": {Tok: openstackResource(openstackPostgresql, "Configuration")},
			"openstackrm_postgresql_database":      {Tok: openstackResource(openstackPostgresql, "Database")},
			"openstackrm_postgresql_firewall_rule": {Tok: openstackResource(openstackPostgresql, "FirewallRule")},
			"openstackrm_postgresql_server":        {Tok: openstackResource(openstackPostgresql, "Server")},

			// Policy
			"openstackrm_policy_assignment": {Tok: openstackResource(openstackPolicy, "Assignment")},
			"openstackrm_policy_definition": {Tok: openstackResource(openstackPolicy, "Definition")},

			// SQL
			"openstackrm_sql_elasticpool":   {Tok: openstackResource(openstackSQL, "ElasticPool")},
			"openstackrm_sql_database":      {Tok: openstackResource(openstackSQL, "Database")},
			"openstackrm_sql_firewall_rule": {Tok: openstackResource(openstackSQL, "FirewallRule")},
			"openstackrm_sql_server":        {Tok: openstackResource(openstackSQL, "SqlServer")},
			"openstackrm_sql_virtual_network_rule": {Tok: openstackResource(openstackSQL, "VirtualNetworkRule"),
				Docs: &tfbridge.DocInfo{
					Source: "sql_virtual_network_rule.html.markdown",
				},
			},
			"openstackrm_sql_active_directory_administrator": {Tok: openstackResource(openstackSQL, "ActiveDirectoryAdministrator")},

			// Network
			"openstackrm_virtual_network":                    {Tok: openstackResource(openstackNetwork, "VirtualNetwork")},
			"openstackrm_virtual_network_peering":            {Tok: openstackResource(openstackNetwork, "VirtualNetworkPeering")},
			"openstackrm_virtual_network_gateway":            {Tok: openstackResource(openstackNetwork, "VirtualNetworkGateway")},
			"openstackrm_virtual_network_gateway_connection": {Tok: openstackResource(openstackNetwork, "VirtualNetworkGatewayConnection")},
			"openstackrm_local_network_gateway":              {Tok: openstackResource(openstackNetwork, "LocalNetworkGateway")},
			"openstackrm_application_gateway":                {Tok: openstackResource(openstackNetwork, "ApplicationGateway")},
			"openstackrm_application_security_group":         {Tok: openstackResource(openstackNetwork, "ApplicationSecurityGroup")},
			"openstackrm_network_interface":                  {Tok: openstackResource(openstackNetwork, "NetworkInterface")},
			"openstackrm_network_security_group":             {Tok: openstackResource(openstackNetwork, "NetworkSecurityGroup")},
			"openstackrm_network_security_rule":              {Tok: openstackResource(openstackNetwork, "NetworkSecurityRule")},
			"openstackrm_network_watcher":                    {Tok: openstackResource(openstackNetwork, "NetworkWatcher")},
			"openstackrm_packet_capture":                     {Tok: openstackResource(openstackNetwork, "PacketCapture")},
			"openstackrm_public_ip":                          {Tok: openstackResource(openstackNetwork, "PublicIp")},
			"openstackrm_route":                              {Tok: openstackResource(openstackNetwork, "Route")},
			"openstackrm_route_table":                        {Tok: openstackResource(openstackNetwork, "RouteTable")},
			"openstackrm_subnet":                             {Tok: openstackResource(openstackNetwork, "Subnet")},
			"openstackrm_express_route_circuit":              {Tok: openstackResource(openstackNetwork, "ExpressRouteCircuit")},
			"openstackrm_express_route_circuit_authorization": {Tok: openstackResource(openstackNetwork, "ExpressRouteCircuitAuthorization"),
				Docs: &tfbridge.DocInfo{
					Source: "express_route_circuit_authorization.html.markdown",
				},
			},
			"openstackrm_express_route_circuit_peering": {Tok: openstackResource(openstackNetwork, "ExpressRouteCircuitPeering"),
				Docs: &tfbridge.DocInfo{
					Source: "express_route_circuit_authorization.html.markdown",
				},
			},

			// Traffic Manager
			"openstackrm_traffic_manager_endpoint": {Tok: openstackResource(openstackTrafficManager, "Endpoint")},
			"openstackrm_traffic_manager_profile":  {Tok: openstackResource(openstackTrafficManager, "Profile")},

			// Recovery Services
			"openstackrm_recovery_services_vault": {Tok: openstackResource(openstackRecoveryServices, "Vault")},

			// Redis
			"openstackrm_redis_cache":         {Tok: openstackResource(openstackRedis, "Cache")},
			"openstackrm_redis_firewall_rule": {Tok: openstackResource(openstackRedis, "FirewallRule")},

			// Relay
			"openstackrm_relay_namespace": {Tok: openstackResource(openstackRelay, "Namespace")},

			// Scheduler
			"openstackrm_scheduler_job_collection": {Tok: openstackResource(openstackScheduler, "JobCollection")},

			// Search
			"openstackrm_search_service": {Tok: openstackResource(openstackSearch, "Service")},

			// Storage
			"openstackrm_storage_account":   {Tok: openstackResource(openstackStorage, "Account")},
			"openstackrm_storage_blob":      {Tok: openstackResource(openstackStorage, "Blob")},
			"openstackrm_storage_container": {Tok: openstackResource(openstackStorage, "Container")},
			"openstackrm_storage_share":     {Tok: openstackResource(openstackStorage, "Share")},
			"openstackrm_storage_queue":     {Tok: openstackResource(openstackStorage, "Queue")},
			"openstackrm_storage_table":     {Tok: openstackResource(openstackStorage, "Table")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"openstackrm_app_service": {
				Tok: openstackDataSource(openstackAppService, "getAppService"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Ensure "sku" is a singleton
					"sku": {Name: "sku", MaxItemsOne: boolRef(true)},
				}},
			"openstackrm_app_service_plan": {
				Tok: openstackDataSource(openstackAppService, "getAppServicePlan"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Ensure "sku" is a singleton
					"sku": {Name: "sku", MaxItemsOne: boolRef(true)},
				},
			},
			"openstackrm_subscriptions":      {Tok: openstackDataSource(openstackCore, "getSubscriptions")},
			"openstackrm_cdn_profile":        {Tok: openstackDataSource(openstackCDN, "getProfile")},
			"openstackrm_client_config":      {Tok: openstackDataSource(openstackCore, "getClientConfig")},
			"openstackrm_cosmosdb_account":   {Tok: openstackDataSource(openstackCosmosDB, "getAccount")},
			"openstackrm_data_lake_store":    {Tok: openstackDataSource(openstackDatalake, "getStore")},
			"openstackrm_eventhub_namespace": {Tok: openstackDataSource(openstackMessaging, "getEventhubNamespace")},
			"openstackrm_image":              {Tok: openstackDataSource(openstackCompute, "getImage")},
			"openstackrm_dns_zone":           {Tok: openstackDataSource(openstackDNS, "getZone")},
			"openstackrm_key_vault": {
				Tok: openstackDataSource(openstackKeyVault, "getKeyVault"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Ensure "sku" is a singleton
					"sku": {Name: "sku", MaxItemsOne: boolRef(true)},
				},
			},
			"openstackrm_key_vault_access_policy":               {Tok: openstackDataSource(openstackKeyVault, "getAccessPolicy")},
			"openstackrm_key_vault_secret":                      {Tok: openstackDataSource(openstackKeyVault, "getSecret")},
			"openstackrm_kubernetes_cluster":                    {Tok: openstackDataSource(openstackContainerService, "getKubernetesCluster")},
			"openstackrm_virtual_network":                       {Tok: openstackDataSource(openstackNetwork, "getVirtualNetwork")},
			"openstackrm_virtual_network_gateway":               {Tok: openstackDataSource(openstackNetwork, "getVirtualNetworkGateway")},
			"openstackrm_network_security_group":                {Tok: openstackDataSource(openstackNetwork, "getNetworkSecurityGroup")},
			"openstackrm_network_interface":                     {Tok: openstackDataSource(openstackNetwork, "getNetworkInterface")},
			"openstackrm_public_ip":                             {Tok: openstackDataSource(openstackNetwork, "getPublicIP")},
			"openstackrm_public_ips":                            {Tok: openstackDataSource(openstackNetwork, "getPublicIPs")},
			"openstackrm_application_security_group":            {Tok: openstackDataSource(openstackNetwork, "getApplicationSecurityGroup")},
			"openstackrm_recovery_services_vault":               {Tok: openstackDataSource(openstackRecoveryServices, "getVault")},
			"openstackrm_resource_group":                        {Tok: openstackDataSource(openstackCore, "getResourceGroup")},
			"openstackrm_snapshot":                              {Tok: openstackDataSource(openstackCompute, "getSnapshot")},
			"openstackrm_subnet":                                {Tok: openstackDataSource(openstackNetwork, "getSubnet")},
			"openstackrm_route_table":                           {Tok: openstackDataSource(openstackNetwork, "getRouteTable")},
			"openstackrm_subscription":                          {Tok: openstackDataSource(openstackCore, "getSubscription")},
			"openstackrm_platform_image":                        {Tok: openstackDataSource(openstackCompute, "getPlatformImage")},
			"openstackrm_managed_disk":                          {Tok: openstackDataSource(openstackCompute, "getManagedDisk")},
			"openstackrm_role_definition":                       {Tok: openstackDataSource(openstackRole, "getRoleDefinition")},
			"openstackrm_builtin_role_definition":               {Tok: openstackDataSource(openstackRole, "getBuiltinRoleDefinition")},
			"openstackrm_scheduler_job_collection":              {Tok: openstackDataSource(openstackScheduler, "getJobCollection")},
			"openstackrm_storage_account":                       {Tok: openstackDataSource(openstackStorage, "getAccount")},
			"openstackrm_storage_account_sas":                   {Tok: openstackDataSource(openstackStorage, "getAccountSAS")},
			"openstackrm_traffic_manager_geographical_location": {Tok: openstackDataSource(openstackTrafficManager, "getGeographicalLocation")},
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

	// TODO[pulumi/pulumi#280]: Until we can pass an Archive as an Asset, create a resource type
	// specifically for uploading ZIP blobs to OpenStack storage.
	prov.P.ResourcesMap["openstackrm_storage_zipblob"] = prov.P.ResourcesMap["openstackrm_storage_blob"]
	prov.Resources["openstackrm_storage_zipblob"] = &tfbridge.ResourceInfo{
		Tok: openstackResource(openstackStorage, "ZipBlob"),
		Fields: map[string]*tfbridge.SchemaInfo{
			"source": {
				Name: "content",
				Asset: &tfbridge.AssetTranslation{
					Kind:   tfbridge.FileArchive,
					Format: resource.ZIPArchive,
				},
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
					// Use conservative options that apply broadly for OpenStack.  See
					// https://docs.microsoft.com/en-us/openstack/architecture/best-practices/naming-conventions for
					// details.
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
