module github.com/pulumi/pulumi-openstack

go 1.12

require (
	github.com/chzyer/logex v1.1.11-0.20160617073814-96a4d311aa9b // indirect
	github.com/coreos/bbolt v1.3.1-coreos.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.0 // indirect
	github.com/hashicorp/serf v0.8.2-0.20171022020050-c20a0b1b1ea9 // indirect
	github.com/hashicorp/terraform v0.12.2
	github.com/miekg/dns v1.0.14 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.23-0.20190715212628-02ffff88409f
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190716112909-08d502e9b427
	github.com/terraform-providers/terraform-provider-openstack v1.20.0
)

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)
