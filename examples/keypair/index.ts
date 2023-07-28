import * as os from "@pulumi/openstack";

const keypair = new os.compute.Keypair("test", {});

export const privateKey = keypair.privateKey;
