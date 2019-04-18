// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the ID and public key of an OpenStack keypair.
func LookupKeypair(ctx *pulumi.Context, args *GetKeypairArgs) (*GetKeypairResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["region"] = args.Region
	}
	outputs, err := ctx.Invoke("openstack:compute/getKeypair:getKeypair", inputs)
	if err != nil {
		return nil, err
	}
	return &GetKeypairResult{
		Fingerprint: outputs["fingerprint"],
		Name: outputs["name"],
		PublicKey: outputs["publicKey"],
		Region: outputs["region"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getKeypair.
type GetKeypairArgs struct {
	// The unique name of the keypair.
	Name interface{}
	// The region in which to obtain the V2 Compute client.
	// If omitted, the `region` argument of the provider is used.
	Region interface{}
}

// A collection of values returned by getKeypair.
type GetKeypairResult struct {
	// The fingerprint of the OpenSSH key.
	Fingerprint interface{}
	// See Argument Reference above.
	Name interface{}
	// The OpenSSH-formatted public key of the keypair.
	PublicKey interface{}
	// See Argument Reference above.
	Region interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
