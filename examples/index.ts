import * as os from "@pulumi/openstack";

// flavorName and imageName will need changing based on what's availible in the
// OpenStack installation your using.
const instance = new os.compute.Instance("test", {
	flavorName: "s1-2",
	imageName: "Ubuntu 16.04",
});

export let instanceIP = instance.accessIpV4;