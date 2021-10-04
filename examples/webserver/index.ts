import * as os from "@pulumi/openstack";

// flavorName and imageName may need to change based on what's availible in the
// OpenStack installation you're using.
const instance = new os.compute.Instance("test", {
	flavorName: "s1-2",
	imageName: "Ubuntu 18.04",
});

export let instanceIP = instance.accessIpV4;
