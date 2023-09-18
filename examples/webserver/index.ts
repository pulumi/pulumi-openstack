import * as os from "@pulumi/openstack";

// flavorName and imageName may need to change based on what's availible in the
// OpenStack installation you're using.
const instance = new os.compute.Instance("test", {
    flavorName: "m1.tiny",
    imageName: "Ubuntu 22.04",

    // Networks are optional if there is only one network. This example assumes there is a network called "Ext-Net".
    //
    // See https://www.pulumi.com/registry/packages/openstack/api-docs/networking/network/ for a different example that
    // provisions a network as part of the Pulumi program.
    networks: [{
        name: "Ext-Net"
    }],
});

export let instanceIP = instance.accessIpV4;
