// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerinfra

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a V1 Magnum node group resource within OpenStack.
//
// ## Example Usage
//
// ### Create a Nodegroup
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-openstack/sdk/v5/go/openstack/containerinfra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := containerinfra.NewNodeGroup(ctx, "nodegroup_1", &containerinfra.NodeGroupArgs{
//				Name:      pulumi.String("nodegroup_1"),
//				ClusterId: pulumi.String("b9a45c5c-cd03-4958-82aa-b80bf93cb922"),
//				NodeCount: pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Attributes reference
//
// The following attributes are exported:
//
// * `region` - See Argument Reference above.
// * `name` - See Argument Reference above.
// * `projectId` - See Argument Reference above.
// * `createdAt` - The time at which node group was created.
// * `updatedAt` - The time at which node group was created.
// * `dockerVolumeSize` - See Argument Reference above.
// * `role` - See Argument Reference above.
// * `imageId` - See Argument Reference above.
// * `flavorId` - See Argument Reference above.
// * `labels` - See Argument Reference above.
// * `nodeCount` - See Argument Reference above.
// * `minNodeCount` - See Argument Reference above.
// * `maxNodeCount` - See Argument Reference above.
// * `role` - See Argument Reference above.
//
// ## Import
//
// Node groups can be imported using the `id` (cluster_id/nodegroup_id), e.g.
//
// ```sh
// $ pulumi import openstack:containerinfra/nodeGroup:NodeGroup nodegroup_1 b9a45c5c-cd03-4958-82aa-b80bf93cb922/ce0f9463-dd25-474b-9fe8-94de63e5e42b
// ```
type NodeGroup struct {
	pulumi.CustomResourceState

	// The UUID of the V1 Container Infra cluster.
	// Changing this creates a new node group.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The size (in GB) of the Docker volume.
	// Changing this creates a new node group.
	DockerVolumeSize pulumi.IntOutput `pulumi:"dockerVolumeSize"`
	// The flavor for the nodes of the node group. Can be set
	// via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// node group.
	FlavorId pulumi.StringOutput `pulumi:"flavorId"`
	// The reference to an image that is used for nodes of the
	// node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
	// Changing this updates the image attribute of the existing node group.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The list of key value pairs representing additional
	// properties of the node group. Changing this creates a new node group.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The maximum number of nodes for the node group.
	// Changing this update the maximum number of nodes of the node group.
	MaxNodeCount pulumi.IntPtrOutput `pulumi:"maxNodeCount"`
	// Indicates whether the provided labels should be
	// merged with cluster labels. Changing this creates a new nodegroup.
	MergeLabels pulumi.BoolPtrOutput `pulumi:"mergeLabels"`
	// The minimum number of nodes for the node group.
	// Changing this update the minimum number of nodes of the node group.
	MinNodeCount pulumi.IntOutput `pulumi:"minNodeCount"`
	// The name of the node group. Changing this creates a new
	// node group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of nodes for the node group. Changing
	// this update the number of nodes of the node group.
	NodeCount pulumi.IntPtrOutput `pulumi:"nodeCount"`
	// The project of the node group. Required if admin
	// wants to create a cluster in another project. Changing this creates a new
	// node group.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// node group.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role of nodes in the node group. Changing this
	// creates a new node group.
	Role      pulumi.StringOutput `pulumi:"role"`
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewNodeGroup registers a new resource with the given unique name, arguments, and options.
func NewNodeGroup(ctx *pulumi.Context,
	name string, args *NodeGroupArgs, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NodeGroup
	err := ctx.RegisterResource("openstack:containerinfra/nodeGroup:NodeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeGroup gets an existing NodeGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeGroupState, opts ...pulumi.ResourceOption) (*NodeGroup, error) {
	var resource NodeGroup
	err := ctx.ReadResource("openstack:containerinfra/nodeGroup:NodeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeGroup resources.
type nodeGroupState struct {
	// The UUID of the V1 Container Infra cluster.
	// Changing this creates a new node group.
	ClusterId *string `pulumi:"clusterId"`
	CreatedAt *string `pulumi:"createdAt"`
	// The size (in GB) of the Docker volume.
	// Changing this creates a new node group.
	DockerVolumeSize *int `pulumi:"dockerVolumeSize"`
	// The flavor for the nodes of the node group. Can be set
	// via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// node group.
	FlavorId *string `pulumi:"flavorId"`
	// The reference to an image that is used for nodes of the
	// node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
	// Changing this updates the image attribute of the existing node group.
	ImageId *string `pulumi:"imageId"`
	// The list of key value pairs representing additional
	// properties of the node group. Changing this creates a new node group.
	Labels map[string]string `pulumi:"labels"`
	// The maximum number of nodes for the node group.
	// Changing this update the maximum number of nodes of the node group.
	MaxNodeCount *int `pulumi:"maxNodeCount"`
	// Indicates whether the provided labels should be
	// merged with cluster labels. Changing this creates a new nodegroup.
	MergeLabels *bool `pulumi:"mergeLabels"`
	// The minimum number of nodes for the node group.
	// Changing this update the minimum number of nodes of the node group.
	MinNodeCount *int `pulumi:"minNodeCount"`
	// The name of the node group. Changing this creates a new
	// node group.
	Name *string `pulumi:"name"`
	// The number of nodes for the node group. Changing
	// this update the number of nodes of the node group.
	NodeCount *int `pulumi:"nodeCount"`
	// The project of the node group. Required if admin
	// wants to create a cluster in another project. Changing this creates a new
	// node group.
	ProjectId *string `pulumi:"projectId"`
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// node group.
	Region *string `pulumi:"region"`
	// The role of nodes in the node group. Changing this
	// creates a new node group.
	Role      *string `pulumi:"role"`
	UpdatedAt *string `pulumi:"updatedAt"`
}

type NodeGroupState struct {
	// The UUID of the V1 Container Infra cluster.
	// Changing this creates a new node group.
	ClusterId pulumi.StringPtrInput
	CreatedAt pulumi.StringPtrInput
	// The size (in GB) of the Docker volume.
	// Changing this creates a new node group.
	DockerVolumeSize pulumi.IntPtrInput
	// The flavor for the nodes of the node group. Can be set
	// via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// node group.
	FlavorId pulumi.StringPtrInput
	// The reference to an image that is used for nodes of the
	// node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
	// Changing this updates the image attribute of the existing node group.
	ImageId pulumi.StringPtrInput
	// The list of key value pairs representing additional
	// properties of the node group. Changing this creates a new node group.
	Labels pulumi.StringMapInput
	// The maximum number of nodes for the node group.
	// Changing this update the maximum number of nodes of the node group.
	MaxNodeCount pulumi.IntPtrInput
	// Indicates whether the provided labels should be
	// merged with cluster labels. Changing this creates a new nodegroup.
	MergeLabels pulumi.BoolPtrInput
	// The minimum number of nodes for the node group.
	// Changing this update the minimum number of nodes of the node group.
	MinNodeCount pulumi.IntPtrInput
	// The name of the node group. Changing this creates a new
	// node group.
	Name pulumi.StringPtrInput
	// The number of nodes for the node group. Changing
	// this update the number of nodes of the node group.
	NodeCount pulumi.IntPtrInput
	// The project of the node group. Required if admin
	// wants to create a cluster in another project. Changing this creates a new
	// node group.
	ProjectId pulumi.StringPtrInput
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// node group.
	Region pulumi.StringPtrInput
	// The role of nodes in the node group. Changing this
	// creates a new node group.
	Role      pulumi.StringPtrInput
	UpdatedAt pulumi.StringPtrInput
}

func (NodeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupState)(nil)).Elem()
}

type nodeGroupArgs struct {
	// The UUID of the V1 Container Infra cluster.
	// Changing this creates a new node group.
	ClusterId string `pulumi:"clusterId"`
	// The size (in GB) of the Docker volume.
	// Changing this creates a new node group.
	DockerVolumeSize *int `pulumi:"dockerVolumeSize"`
	// The flavor for the nodes of the node group. Can be set
	// via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// node group.
	FlavorId *string `pulumi:"flavorId"`
	// The reference to an image that is used for nodes of the
	// node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
	// Changing this updates the image attribute of the existing node group.
	ImageId *string `pulumi:"imageId"`
	// The list of key value pairs representing additional
	// properties of the node group. Changing this creates a new node group.
	Labels map[string]string `pulumi:"labels"`
	// The maximum number of nodes for the node group.
	// Changing this update the maximum number of nodes of the node group.
	MaxNodeCount *int `pulumi:"maxNodeCount"`
	// Indicates whether the provided labels should be
	// merged with cluster labels. Changing this creates a new nodegroup.
	MergeLabels *bool `pulumi:"mergeLabels"`
	// The minimum number of nodes for the node group.
	// Changing this update the minimum number of nodes of the node group.
	MinNodeCount *int `pulumi:"minNodeCount"`
	// The name of the node group. Changing this creates a new
	// node group.
	Name *string `pulumi:"name"`
	// The number of nodes for the node group. Changing
	// this update the number of nodes of the node group.
	NodeCount *int `pulumi:"nodeCount"`
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// node group.
	Region *string `pulumi:"region"`
	// The role of nodes in the node group. Changing this
	// creates a new node group.
	Role *string `pulumi:"role"`
}

// The set of arguments for constructing a NodeGroup resource.
type NodeGroupArgs struct {
	// The UUID of the V1 Container Infra cluster.
	// Changing this creates a new node group.
	ClusterId pulumi.StringInput
	// The size (in GB) of the Docker volume.
	// Changing this creates a new node group.
	DockerVolumeSize pulumi.IntPtrInput
	// The flavor for the nodes of the node group. Can be set
	// via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
	// node group.
	FlavorId pulumi.StringPtrInput
	// The reference to an image that is used for nodes of the
	// node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
	// Changing this updates the image attribute of the existing node group.
	ImageId pulumi.StringPtrInput
	// The list of key value pairs representing additional
	// properties of the node group. Changing this creates a new node group.
	Labels pulumi.StringMapInput
	// The maximum number of nodes for the node group.
	// Changing this update the maximum number of nodes of the node group.
	MaxNodeCount pulumi.IntPtrInput
	// Indicates whether the provided labels should be
	// merged with cluster labels. Changing this creates a new nodegroup.
	MergeLabels pulumi.BoolPtrInput
	// The minimum number of nodes for the node group.
	// Changing this update the minimum number of nodes of the node group.
	MinNodeCount pulumi.IntPtrInput
	// The name of the node group. Changing this creates a new
	// node group.
	Name pulumi.StringPtrInput
	// The number of nodes for the node group. Changing
	// this update the number of nodes of the node group.
	NodeCount pulumi.IntPtrInput
	// The region in which to obtain the V1 Container Infra
	// client. A Container Infra client is needed to create a cluster. If omitted,
	// the `region` argument of the provider is used. Changing this creates a new
	// node group.
	Region pulumi.StringPtrInput
	// The role of nodes in the node group. Changing this
	// creates a new node group.
	Role pulumi.StringPtrInput
}

func (NodeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeGroupArgs)(nil)).Elem()
}

type NodeGroupInput interface {
	pulumi.Input

	ToNodeGroupOutput() NodeGroupOutput
	ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput
}

func (*NodeGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroup)(nil)).Elem()
}

func (i *NodeGroup) ToNodeGroupOutput() NodeGroupOutput {
	return i.ToNodeGroupOutputWithContext(context.Background())
}

func (i *NodeGroup) ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupOutput)
}

// NodeGroupArrayInput is an input type that accepts NodeGroupArray and NodeGroupArrayOutput values.
// You can construct a concrete instance of `NodeGroupArrayInput` via:
//
//	NodeGroupArray{ NodeGroupArgs{...} }
type NodeGroupArrayInput interface {
	pulumi.Input

	ToNodeGroupArrayOutput() NodeGroupArrayOutput
	ToNodeGroupArrayOutputWithContext(context.Context) NodeGroupArrayOutput
}

type NodeGroupArray []NodeGroupInput

func (NodeGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeGroup)(nil)).Elem()
}

func (i NodeGroupArray) ToNodeGroupArrayOutput() NodeGroupArrayOutput {
	return i.ToNodeGroupArrayOutputWithContext(context.Background())
}

func (i NodeGroupArray) ToNodeGroupArrayOutputWithContext(ctx context.Context) NodeGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupArrayOutput)
}

// NodeGroupMapInput is an input type that accepts NodeGroupMap and NodeGroupMapOutput values.
// You can construct a concrete instance of `NodeGroupMapInput` via:
//
//	NodeGroupMap{ "key": NodeGroupArgs{...} }
type NodeGroupMapInput interface {
	pulumi.Input

	ToNodeGroupMapOutput() NodeGroupMapOutput
	ToNodeGroupMapOutputWithContext(context.Context) NodeGroupMapOutput
}

type NodeGroupMap map[string]NodeGroupInput

func (NodeGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeGroup)(nil)).Elem()
}

func (i NodeGroupMap) ToNodeGroupMapOutput() NodeGroupMapOutput {
	return i.ToNodeGroupMapOutputWithContext(context.Background())
}

func (i NodeGroupMap) ToNodeGroupMapOutputWithContext(ctx context.Context) NodeGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupMapOutput)
}

type NodeGroupOutput struct{ *pulumi.OutputState }

func (NodeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodeGroup)(nil)).Elem()
}

func (o NodeGroupOutput) ToNodeGroupOutput() NodeGroupOutput {
	return o
}

func (o NodeGroupOutput) ToNodeGroupOutputWithContext(ctx context.Context) NodeGroupOutput {
	return o
}

// The UUID of the V1 Container Infra cluster.
// Changing this creates a new node group.
func (o NodeGroupOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

func (o NodeGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The size (in GB) of the Docker volume.
// Changing this creates a new node group.
func (o NodeGroupOutput) DockerVolumeSize() pulumi.IntOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.IntOutput { return v.DockerVolumeSize }).(pulumi.IntOutput)
}

// The flavor for the nodes of the node group. Can be set
// via the `OS_MAGNUM_FLAVOR` environment variable. Changing this creates a new
// node group.
func (o NodeGroupOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.FlavorId }).(pulumi.StringOutput)
}

// The reference to an image that is used for nodes of the
// node group. Can be set via the `OS_MAGNUM_IMAGE` environment variable.
// Changing this updates the image attribute of the existing node group.
func (o NodeGroupOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The list of key value pairs representing additional
// properties of the node group. Changing this creates a new node group.
func (o NodeGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The maximum number of nodes for the node group.
// Changing this update the maximum number of nodes of the node group.
func (o NodeGroupOutput) MaxNodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.IntPtrOutput { return v.MaxNodeCount }).(pulumi.IntPtrOutput)
}

// Indicates whether the provided labels should be
// merged with cluster labels. Changing this creates a new nodegroup.
func (o NodeGroupOutput) MergeLabels() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.BoolPtrOutput { return v.MergeLabels }).(pulumi.BoolPtrOutput)
}

// The minimum number of nodes for the node group.
// Changing this update the minimum number of nodes of the node group.
func (o NodeGroupOutput) MinNodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.IntOutput { return v.MinNodeCount }).(pulumi.IntOutput)
}

// The name of the node group. Changing this creates a new
// node group.
func (o NodeGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The number of nodes for the node group. Changing
// this update the number of nodes of the node group.
func (o NodeGroupOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.IntPtrOutput { return v.NodeCount }).(pulumi.IntPtrOutput)
}

// The project of the node group. Required if admin
// wants to create a cluster in another project. Changing this creates a new
// node group.
func (o NodeGroupOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region in which to obtain the V1 Container Infra
// client. A Container Infra client is needed to create a cluster. If omitted,
// the `region` argument of the provider is used. Changing this creates a new
// node group.
func (o NodeGroupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The role of nodes in the node group. Changing this
// creates a new node group.
func (o NodeGroupOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o NodeGroupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *NodeGroup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type NodeGroupArrayOutput struct{ *pulumi.OutputState }

func (NodeGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodeGroup)(nil)).Elem()
}

func (o NodeGroupArrayOutput) ToNodeGroupArrayOutput() NodeGroupArrayOutput {
	return o
}

func (o NodeGroupArrayOutput) ToNodeGroupArrayOutputWithContext(ctx context.Context) NodeGroupArrayOutput {
	return o
}

func (o NodeGroupArrayOutput) Index(i pulumi.IntInput) NodeGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NodeGroup {
		return vs[0].([]*NodeGroup)[vs[1].(int)]
	}).(NodeGroupOutput)
}

type NodeGroupMapOutput struct{ *pulumi.OutputState }

func (NodeGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodeGroup)(nil)).Elem()
}

func (o NodeGroupMapOutput) ToNodeGroupMapOutput() NodeGroupMapOutput {
	return o
}

func (o NodeGroupMapOutput) ToNodeGroupMapOutputWithContext(ctx context.Context) NodeGroupMapOutput {
	return o
}

func (o NodeGroupMapOutput) MapIndex(k pulumi.StringInput) NodeGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NodeGroup {
		return vs[0].(map[string]*NodeGroup)[vs[1].(string)]
	}).(NodeGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupInput)(nil)).Elem(), &NodeGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupArrayInput)(nil)).Elem(), NodeGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeGroupMapInput)(nil)).Elem(), NodeGroupMap{})
	pulumi.RegisterOutputType(NodeGroupOutput{})
	pulumi.RegisterOutputType(NodeGroupArrayOutput{})
	pulumi.RegisterOutputType(NodeGroupMapOutput{})
}
