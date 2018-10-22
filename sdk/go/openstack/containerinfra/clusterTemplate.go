// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a V1 Magnum cluster template resource within OpenStack.
type ClusterTemplate struct {
	s *pulumi.ResourceState
}

// NewClusterTemplate registers a new resource with the given unique name, arguments, and options.
func NewClusterTemplate(ctx *pulumi.Context,
	name string, args *ClusterTemplateArgs, opts ...pulumi.ResourceOpt) (*ClusterTemplate, error) {
	if args == nil || args.Coe == nil {
		return nil, errors.New("missing required argument 'Coe'")
	}
	if args == nil || args.Image == nil {
		return nil, errors.New("missing required argument 'Image'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiserverPort"] = nil
		inputs["clusterDistro"] = nil
		inputs["coe"] = nil
		inputs["dnsNameserver"] = nil
		inputs["dockerStorageDriver"] = nil
		inputs["dockerVolumeSize"] = nil
		inputs["externalNetworkId"] = nil
		inputs["fixedNetwork"] = nil
		inputs["fixedSubnet"] = nil
		inputs["flavor"] = nil
		inputs["floatingIpEnabled"] = nil
		inputs["httpProxy"] = nil
		inputs["httpsProxy"] = nil
		inputs["image"] = nil
		inputs["insecureRegistry"] = nil
		inputs["keypairId"] = nil
		inputs["labels"] = nil
		inputs["masterFlavor"] = nil
		inputs["masterLbEnabled"] = nil
		inputs["name"] = nil
		inputs["networkDriver"] = nil
		inputs["noProxy"] = nil
		inputs["public"] = nil
		inputs["region"] = nil
		inputs["registryEnabled"] = nil
		inputs["serverType"] = nil
		inputs["tlsDisabled"] = nil
		inputs["volumeDriver"] = nil
	} else {
		inputs["apiserverPort"] = args.ApiserverPort
		inputs["clusterDistro"] = args.ClusterDistro
		inputs["coe"] = args.Coe
		inputs["dnsNameserver"] = args.DnsNameserver
		inputs["dockerStorageDriver"] = args.DockerStorageDriver
		inputs["dockerVolumeSize"] = args.DockerVolumeSize
		inputs["externalNetworkId"] = args.ExternalNetworkId
		inputs["fixedNetwork"] = args.FixedNetwork
		inputs["fixedSubnet"] = args.FixedSubnet
		inputs["flavor"] = args.Flavor
		inputs["floatingIpEnabled"] = args.FloatingIpEnabled
		inputs["httpProxy"] = args.HttpProxy
		inputs["httpsProxy"] = args.HttpsProxy
		inputs["image"] = args.Image
		inputs["insecureRegistry"] = args.InsecureRegistry
		inputs["keypairId"] = args.KeypairId
		inputs["labels"] = args.Labels
		inputs["masterFlavor"] = args.MasterFlavor
		inputs["masterLbEnabled"] = args.MasterLbEnabled
		inputs["name"] = args.Name
		inputs["networkDriver"] = args.NetworkDriver
		inputs["noProxy"] = args.NoProxy
		inputs["public"] = args.Public
		inputs["region"] = args.Region
		inputs["registryEnabled"] = args.RegistryEnabled
		inputs["serverType"] = args.ServerType
		inputs["tlsDisabled"] = args.TlsDisabled
		inputs["volumeDriver"] = args.VolumeDriver
	}
	inputs["createdAt"] = nil
	inputs["projectId"] = nil
	inputs["updatedAt"] = nil
	inputs["userId"] = nil
	s, err := ctx.RegisterResource("openstack:containerinfra/clusterTemplate:ClusterTemplate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterTemplate{s: s}, nil
}

// GetClusterTemplate gets an existing ClusterTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterTemplate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterTemplateState, opts ...pulumi.ResourceOpt) (*ClusterTemplate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiserverPort"] = state.ApiserverPort
		inputs["clusterDistro"] = state.ClusterDistro
		inputs["coe"] = state.Coe
		inputs["createdAt"] = state.CreatedAt
		inputs["dnsNameserver"] = state.DnsNameserver
		inputs["dockerStorageDriver"] = state.DockerStorageDriver
		inputs["dockerVolumeSize"] = state.DockerVolumeSize
		inputs["externalNetworkId"] = state.ExternalNetworkId
		inputs["fixedNetwork"] = state.FixedNetwork
		inputs["fixedSubnet"] = state.FixedSubnet
		inputs["flavor"] = state.Flavor
		inputs["floatingIpEnabled"] = state.FloatingIpEnabled
		inputs["httpProxy"] = state.HttpProxy
		inputs["httpsProxy"] = state.HttpsProxy
		inputs["image"] = state.Image
		inputs["insecureRegistry"] = state.InsecureRegistry
		inputs["keypairId"] = state.KeypairId
		inputs["labels"] = state.Labels
		inputs["masterFlavor"] = state.MasterFlavor
		inputs["masterLbEnabled"] = state.MasterLbEnabled
		inputs["name"] = state.Name
		inputs["networkDriver"] = state.NetworkDriver
		inputs["noProxy"] = state.NoProxy
		inputs["projectId"] = state.ProjectId
		inputs["public"] = state.Public
		inputs["region"] = state.Region
		inputs["registryEnabled"] = state.RegistryEnabled
		inputs["serverType"] = state.ServerType
		inputs["tlsDisabled"] = state.TlsDisabled
		inputs["updatedAt"] = state.UpdatedAt
		inputs["userId"] = state.UserId
		inputs["volumeDriver"] = state.VolumeDriver
	}
	s, err := ctx.ReadResource("openstack:containerinfra/clusterTemplate:ClusterTemplate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ClusterTemplate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ClusterTemplate) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ClusterTemplate) ID() *pulumi.IDOutput {
	return r.s.ID
}

func (r *ClusterTemplate) ApiserverPort() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["apiserverPort"])
}

func (r *ClusterTemplate) ClusterDistro() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["clusterDistro"])
}

func (r *ClusterTemplate) Coe() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["coe"])
}

func (r *ClusterTemplate) CreatedAt() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createdAt"])
}

func (r *ClusterTemplate) DnsNameserver() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsNameserver"])
}

func (r *ClusterTemplate) DockerStorageDriver() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dockerStorageDriver"])
}

func (r *ClusterTemplate) DockerVolumeSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["dockerVolumeSize"])
}

func (r *ClusterTemplate) ExternalNetworkId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["externalNetworkId"])
}

func (r *ClusterTemplate) FixedNetwork() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fixedNetwork"])
}

func (r *ClusterTemplate) FixedSubnet() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fixedSubnet"])
}

func (r *ClusterTemplate) Flavor() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["flavor"])
}

func (r *ClusterTemplate) FloatingIpEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["floatingIpEnabled"])
}

func (r *ClusterTemplate) HttpProxy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["httpProxy"])
}

func (r *ClusterTemplate) HttpsProxy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["httpsProxy"])
}

func (r *ClusterTemplate) Image() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["image"])
}

func (r *ClusterTemplate) InsecureRegistry() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["insecureRegistry"])
}

func (r *ClusterTemplate) KeypairId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keypairId"])
}

func (r *ClusterTemplate) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *ClusterTemplate) MasterFlavor() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterFlavor"])
}

func (r *ClusterTemplate) MasterLbEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["masterLbEnabled"])
}

func (r *ClusterTemplate) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *ClusterTemplate) NetworkDriver() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkDriver"])
}

func (r *ClusterTemplate) NoProxy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["noProxy"])
}

func (r *ClusterTemplate) ProjectId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["projectId"])
}

func (r *ClusterTemplate) Public() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["public"])
}

func (r *ClusterTemplate) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

func (r *ClusterTemplate) RegistryEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["registryEnabled"])
}

func (r *ClusterTemplate) ServerType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serverType"])
}

func (r *ClusterTemplate) TlsDisabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["tlsDisabled"])
}

func (r *ClusterTemplate) UpdatedAt() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["updatedAt"])
}

func (r *ClusterTemplate) UserId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["userId"])
}

func (r *ClusterTemplate) VolumeDriver() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["volumeDriver"])
}

// Input properties used for looking up and filtering ClusterTemplate resources.
type ClusterTemplateState struct {
	ApiserverPort interface{}
	ClusterDistro interface{}
	Coe interface{}
	CreatedAt interface{}
	DnsNameserver interface{}
	DockerStorageDriver interface{}
	DockerVolumeSize interface{}
	ExternalNetworkId interface{}
	FixedNetwork interface{}
	FixedSubnet interface{}
	Flavor interface{}
	FloatingIpEnabled interface{}
	HttpProxy interface{}
	HttpsProxy interface{}
	Image interface{}
	InsecureRegistry interface{}
	KeypairId interface{}
	Labels interface{}
	MasterFlavor interface{}
	MasterLbEnabled interface{}
	Name interface{}
	NetworkDriver interface{}
	NoProxy interface{}
	ProjectId interface{}
	Public interface{}
	Region interface{}
	RegistryEnabled interface{}
	ServerType interface{}
	TlsDisabled interface{}
	UpdatedAt interface{}
	UserId interface{}
	VolumeDriver interface{}
}

// The set of arguments for constructing a ClusterTemplate resource.
type ClusterTemplateArgs struct {
	ApiserverPort interface{}
	ClusterDistro interface{}
	Coe interface{}
	DnsNameserver interface{}
	DockerStorageDriver interface{}
	DockerVolumeSize interface{}
	ExternalNetworkId interface{}
	FixedNetwork interface{}
	FixedSubnet interface{}
	Flavor interface{}
	FloatingIpEnabled interface{}
	HttpProxy interface{}
	HttpsProxy interface{}
	Image interface{}
	InsecureRegistry interface{}
	KeypairId interface{}
	Labels interface{}
	MasterFlavor interface{}
	MasterLbEnabled interface{}
	Name interface{}
	NetworkDriver interface{}
	NoProxy interface{}
	Public interface{}
	Region interface{}
	RegistryEnabled interface{}
	ServerType interface{}
	TlsDisabled interface{}
	VolumeDriver interface{}
}
