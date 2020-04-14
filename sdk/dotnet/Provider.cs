// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack
{
    /// <summary>
    /// The provider type for the openstack package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("openstack", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to `true`, OpenStack authorization will be perfomed automatically, if the initial auth token get expired. This is
        /// useful, when the token TTL is low or the overall Terraform provider execution time expected to be greater than the
        /// initial token TTL.
        /// </summary>
        [Input("allowReauth", json: true)]
        public Input<bool>? AllowReauth { get; set; }

        /// <summary>
        /// Application Credential ID to login with.
        /// </summary>
        [Input("applicationCredentialId")]
        public Input<string>? ApplicationCredentialId { get; set; }

        /// <summary>
        /// Application Credential name to login with.
        /// </summary>
        [Input("applicationCredentialName")]
        public Input<string>? ApplicationCredentialName { get; set; }

        /// <summary>
        /// Application Credential secret to login with.
        /// </summary>
        [Input("applicationCredentialSecret")]
        public Input<string>? ApplicationCredentialSecret { get; set; }

        /// <summary>
        /// The Identity authentication URL.
        /// </summary>
        [Input("authUrl")]
        public Input<string>? AuthUrl { get; set; }

        /// <summary>
        /// A Custom CA certificate.
        /// </summary>
        [Input("cacertFile")]
        public Input<string>? CacertFile { get; set; }

        /// <summary>
        /// A client certificate to authenticate with.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// An entry in a `clouds.yaml` file to use.
        /// </summary>
        [Input("cloud")]
        public Input<string>? Cloud { get; set; }

        /// <summary>
        /// The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
        /// </summary>
        [Input("defaultDomain")]
        public Input<string>? DefaultDomain { get; set; }

        /// <summary>
        /// If set to `true`, OpenStack authorization will be perfomed, when the service provider client is called.
        /// </summary>
        [Input("delayedAuth", json: true)]
        public Input<bool>? DelayedAuth { get; set; }

        /// <summary>
        /// If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
        /// </summary>
        [Input("disableNoCacheHeader", json: true)]
        public Input<bool>? DisableNoCacheHeader { get; set; }

        /// <summary>
        /// The ID of the Domain to scope to (Identity v3).
        /// </summary>
        [Input("domainId")]
        public Input<string>? DomainId { get; set; }

        /// <summary>
        /// The name of the Domain to scope to (Identity v3).
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("endpointOverrides", json: true)]
        private InputMap<object>? _endpointOverrides;

        /// <summary>
        /// A map of services with an endpoint to override what was from the Keystone catalog
        /// </summary>
        public InputMap<object> EndpointOverrides
        {
            get => _endpointOverrides ?? (_endpointOverrides = new InputMap<object>());
            set => _endpointOverrides = value;
        }

        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// Trust self-signed certificates.
        /// </summary>
        [Input("insecure", json: true)]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// A client private key to authenticate with.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// How many times HTTP connection should be retried until giving up.
        /// </summary>
        [Input("maxRetries", json: true)]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// Password to login with.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The ID of the domain where the proejct resides (Identity v3).
        /// </summary>
        [Input("projectDomainId")]
        public Input<string>? ProjectDomainId { get; set; }

        /// <summary>
        /// The name of the domain where the project resides (Identity v3).
        /// </summary>
        [Input("projectDomainName")]
        public Input<string>? ProjectDomainName { get; set; }

        /// <summary>
        /// The OpenStack region to connect to.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
        /// </summary>
        [Input("swauth", json: true)]
        public Input<bool>? Swauth { get; set; }

        /// <summary>
        /// The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
        /// </summary>
        [Input("tenantName")]
        public Input<string>? TenantName { get; set; }

        /// <summary>
        /// Authentication token to use as an alternative to username/password.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
        /// </summary>
        [Input("useOctavia", json: true)]
        public Input<bool>? UseOctavia { get; set; }

        /// <summary>
        /// The ID of the domain where the user resides (Identity v3).
        /// </summary>
        [Input("userDomainId")]
        public Input<string>? UserDomainId { get; set; }

        /// <summary>
        /// The name of the domain where the user resides (Identity v3).
        /// </summary>
        [Input("userDomainName")]
        public Input<string>? UserDomainName { get; set; }

        /// <summary>
        /// Username to login with.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// Username to login with.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public ProviderArgs()
        {
            AllowReauth = Utilities.GetEnvBoolean("OS_ALLOW_REAUTH");
            ApplicationCredentialId = Utilities.GetEnv("OS_APPLICATION_CREDENTIAL_ID");
            ApplicationCredentialName = Utilities.GetEnv("OS_APPLICATION_CREDENTIAL_NAME");
            ApplicationCredentialSecret = Utilities.GetEnv("OS_APPLICATION_CREDENTIAL_SECRET");
            AuthUrl = Utilities.GetEnv("OS_AUTH_URL");
            CacertFile = Utilities.GetEnv("OS_CACERT");
            Cert = Utilities.GetEnv("OS_CERT");
            Cloud = Utilities.GetEnv("OS_CLOUD");
            DefaultDomain = Utilities.GetEnv("OS_DEFAULT_DOMAIN") ?? "default";
            DelayedAuth = Utilities.GetEnvBoolean("OS_DELAYED_AUTH");
            DomainId = Utilities.GetEnv("OS_DOMAIN_ID");
            DomainName = Utilities.GetEnv("OS_DOMAIN_NAME");
            EndpointType = Utilities.GetEnv("OS_ENDPOINT_TYPE");
            Insecure = Utilities.GetEnvBoolean("OS_INSECURE");
            Key = Utilities.GetEnv("OS_KEY");
            Password = Utilities.GetEnv("OS_PASSWORD");
            ProjectDomainId = Utilities.GetEnv("OS_PROJECT_DOMAIN_ID");
            ProjectDomainName = Utilities.GetEnv("OS_PROJECT_DOMAIN_NAME");
            Region = Utilities.GetEnv("OS_REGION_NAME");
            Swauth = Utilities.GetEnvBoolean("OS_SWAUTH");
            TenantId = Utilities.GetEnv("OS_TENANT_ID", "OS_PROJECT_ID");
            TenantName = Utilities.GetEnv("OS_TENANT_NAME", "OS_PROJECT_NAME");
            Token = Utilities.GetEnv("OS_TOKEN", "OS_AUTH_TOKEN");
            UseOctavia = Utilities.GetEnvBoolean("OS_USE_OCTAVIA");
            UserDomainId = Utilities.GetEnv("OS_USER_DOMAIN_ID");
            UserDomainName = Utilities.GetEnv("OS_USER_DOMAIN_NAME");
            UserId = Utilities.GetEnv("OS_USER_ID");
            UserName = Utilities.GetEnv("OS_USERNAME");
        }
    }
}
