// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.OpenStack.SharedFileSystem
{
    /// <summary>
    /// Use this resource to configure a security service.
    /// 
    /// &gt; **Note:** All arguments including the security service password will be
    /// stored in the raw state as plain-text. [Read more about sensitive data in
    /// state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// A security service stores configuration information for clients for
    /// authentication and authorization (AuthN/AuthZ). For example, a share server
    /// will be the client for an existing service such as LDAP, Kerberos, or
    /// Microsoft Active Directory.
    /// 
    /// Minimum supported Manila microversion is 2.7.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using OpenStack = Pulumi.OpenStack;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var securityservice1 = new OpenStack.SharedFileSystem.SecurityService("securityservice1", new()
    ///     {
    ///         Description = "created by terraform",
    ///         DnsIp = "192.168.199.10",
    ///         Domain = "example.com",
    ///         Ou = "CN=Computers,DC=example,DC=com",
    ///         Password = "s8cret",
    ///         Server = "192.168.199.10",
    ///         Type = "active_directory",
    ///         User = "joinDomainUser",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported by specifying the ID of the security service:
    /// 
    /// ```sh
    /// $ pulumi import openstack:sharedfilesystem/securityService:SecurityService securityservice_1 id
    /// ```
    /// </summary>
    [OpenStackResourceType("openstack:sharedfilesystem/securityService:SecurityService")]
    public partial class SecurityService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The human-readable description for the security service.
        /// Changing this updates the description of the existing security service.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The security service DNS IP address that is used inside the
        /// tenant network.
        /// </summary>
        [Output("dnsIp")]
        public Output<string?> DnsIp { get; private set; } = null!;

        /// <summary>
        /// The security service domain.
        /// </summary>
        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        /// <summary>
        /// The name of the security service. Changing this updates the name
        /// of the existing security service.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The security service ou. An organizational unit can be added to
        /// specify where the share ends up. New in Manila microversion 2.44.
        /// </summary>
        [Output("ou")]
        public Output<string?> Ou { get; private set; } = null!;

        /// <summary>
        /// The user password, if you specify a user.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The owner of the Security Service.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a security service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security service.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The security service host name or IP address.
        /// </summary>
        [Output("server")]
        public Output<string?> Server { get; private set; } = null!;

        /// <summary>
        /// The security service type - can either be active\_directory,
        /// kerberos or ldap.  Changing this updates the existing security service.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The security service user or group name that is used by the
        /// tenant.
        /// </summary>
        [Output("user")]
        public Output<string?> User { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityService(string name, SecurityServiceArgs args, CustomResourceOptions? options = null)
            : base("openstack:sharedfilesystem/securityService:SecurityService", name, args ?? new SecurityServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityService(string name, Input<string> id, SecurityServiceState? state = null, CustomResourceOptions? options = null)
            : base("openstack:sharedfilesystem/securityService:SecurityService", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityService Get(string name, Input<string> id, SecurityServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityService(name, id, state, options);
        }
    }

    public sealed class SecurityServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description for the security service.
        /// Changing this updates the description of the existing security service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The security service DNS IP address that is used inside the
        /// tenant network.
        /// </summary>
        [Input("dnsIp")]
        public Input<string>? DnsIp { get; set; }

        /// <summary>
        /// The security service domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The name of the security service. Changing this updates the name
        /// of the existing security service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The security service ou. An organizational unit can be added to
        /// specify where the share ends up. New in Manila microversion 2.44.
        /// </summary>
        [Input("ou")]
        public Input<string>? Ou { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The user password, if you specify a user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a security service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security service.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The security service host name or IP address.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// The security service type - can either be active\_directory,
        /// kerberos or ldap.  Changing this updates the existing security service.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The security service user or group name that is used by the
        /// tenant.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public SecurityServiceArgs()
        {
        }
        public static new SecurityServiceArgs Empty => new SecurityServiceArgs();
    }

    public sealed class SecurityServiceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The human-readable description for the security service.
        /// Changing this updates the description of the existing security service.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The security service DNS IP address that is used inside the
        /// tenant network.
        /// </summary>
        [Input("dnsIp")]
        public Input<string>? DnsIp { get; set; }

        /// <summary>
        /// The security service domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The name of the security service. Changing this updates the name
        /// of the existing security service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The security service ou. An organizational unit can be added to
        /// specify where the share ends up. New in Manila microversion 2.44.
        /// </summary>
        [Input("ou")]
        public Input<string>? Ou { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The user password, if you specify a user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The owner of the Security Service.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The region in which to obtain the V2 Shared File System client.
        /// A Shared File System client is needed to create a security service. If omitted, the
        /// `region` argument of the provider is used. Changing this creates a new
        /// security service.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The security service host name or IP address.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// The security service type - can either be active\_directory,
        /// kerberos or ldap.  Changing this updates the existing security service.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The security service user or group name that is used by the
        /// tenant.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public SecurityServiceState()
        {
        }
        public static new SecurityServiceState Empty => new SecurityServiceState();
    }
}
