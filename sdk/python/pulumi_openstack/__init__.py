# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
from . import _utilities
import typing
# Export this package's modules as members:
from .bgpvpn_network_associate_v2 import *
from .bgpvpn_port_associate_v2 import *
from .bgpvpn_router_associate_v2 import *
from .bgpvpn_v2 import *
from .get_fw_group_v2 import *
from .get_fw_policy_v2 import *
from .get_fw_rule_v2 import *
from .lb_flavorprofile_v2 import *
from .lb_loadbalancer_v2 import *
from .provider import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_openstack.blockstorage as __blockstorage
    blockstorage = __blockstorage
    import pulumi_openstack.compute as __compute
    compute = __compute
    import pulumi_openstack.config as __config
    config = __config
    import pulumi_openstack.containerinfra as __containerinfra
    containerinfra = __containerinfra
    import pulumi_openstack.database as __database
    database = __database
    import pulumi_openstack.dns as __dns
    dns = __dns
    import pulumi_openstack.firewall as __firewall
    firewall = __firewall
    import pulumi_openstack.identity as __identity
    identity = __identity
    import pulumi_openstack.images as __images
    images = __images
    import pulumi_openstack.keymanager as __keymanager
    keymanager = __keymanager
    import pulumi_openstack.loadbalancer as __loadbalancer
    loadbalancer = __loadbalancer
    import pulumi_openstack.networking as __networking
    networking = __networking
    import pulumi_openstack.objectstorage as __objectstorage
    objectstorage = __objectstorage
    import pulumi_openstack.orchestration as __orchestration
    orchestration = __orchestration
    import pulumi_openstack.sharedfilesystem as __sharedfilesystem
    sharedfilesystem = __sharedfilesystem
    import pulumi_openstack.vpnaas as __vpnaas
    vpnaas = __vpnaas
else:
    blockstorage = _utilities.lazy_import('pulumi_openstack.blockstorage')
    compute = _utilities.lazy_import('pulumi_openstack.compute')
    config = _utilities.lazy_import('pulumi_openstack.config')
    containerinfra = _utilities.lazy_import('pulumi_openstack.containerinfra')
    database = _utilities.lazy_import('pulumi_openstack.database')
    dns = _utilities.lazy_import('pulumi_openstack.dns')
    firewall = _utilities.lazy_import('pulumi_openstack.firewall')
    identity = _utilities.lazy_import('pulumi_openstack.identity')
    images = _utilities.lazy_import('pulumi_openstack.images')
    keymanager = _utilities.lazy_import('pulumi_openstack.keymanager')
    loadbalancer = _utilities.lazy_import('pulumi_openstack.loadbalancer')
    networking = _utilities.lazy_import('pulumi_openstack.networking')
    objectstorage = _utilities.lazy_import('pulumi_openstack.objectstorage')
    orchestration = _utilities.lazy_import('pulumi_openstack.orchestration')
    sharedfilesystem = _utilities.lazy_import('pulumi_openstack.sharedfilesystem')
    vpnaas = _utilities.lazy_import('pulumi_openstack.vpnaas')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "openstack",
  "mod": "blockstorage/qosAssociationV3",
  "fqn": "pulumi_openstack.blockstorage",
  "classes": {
   "openstack:blockstorage/qosAssociationV3:QosAssociationV3": "QosAssociationV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "blockstorage/qosV3",
  "fqn": "pulumi_openstack.blockstorage",
  "classes": {
   "openstack:blockstorage/qosV3:QosV3": "QosV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "blockstorage/quoteSetV3",
  "fqn": "pulumi_openstack.blockstorage",
  "classes": {
   "openstack:blockstorage/quoteSetV3:QuoteSetV3": "QuoteSetV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "blockstorage/volume",
  "fqn": "pulumi_openstack.blockstorage",
  "classes": {
   "openstack:blockstorage/volume:Volume": "Volume"
  }
 },
 {
  "pkg": "openstack",
  "mod": "blockstorage/volumeAttach",
  "fqn": "pulumi_openstack.blockstorage",
  "classes": {
   "openstack:blockstorage/volumeAttach:VolumeAttach": "VolumeAttach"
  }
 },
 {
  "pkg": "openstack",
  "mod": "blockstorage/volumeTypeAccessV3",
  "fqn": "pulumi_openstack.blockstorage",
  "classes": {
   "openstack:blockstorage/volumeTypeAccessV3:VolumeTypeAccessV3": "VolumeTypeAccessV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "blockstorage/volumeTypeV3",
  "fqn": "pulumi_openstack.blockstorage",
  "classes": {
   "openstack:blockstorage/volumeTypeV3:VolumeTypeV3": "VolumeTypeV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/aggregateV2",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/aggregateV2:AggregateV2": "AggregateV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/flavor",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/flavor:Flavor": "Flavor"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/flavorAccess",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/flavorAccess:FlavorAccess": "FlavorAccess"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/instance",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/instance:Instance": "Instance"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/interfaceAttach",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/interfaceAttach:InterfaceAttach": "InterfaceAttach"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/keypair",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/keypair:Keypair": "Keypair"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/quotaSetV2",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/quotaSetV2:QuotaSetV2": "QuotaSetV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/serverGroup",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/serverGroup:ServerGroup": "ServerGroup"
  }
 },
 {
  "pkg": "openstack",
  "mod": "compute/volumeAttach",
  "fqn": "pulumi_openstack.compute",
  "classes": {
   "openstack:compute/volumeAttach:VolumeAttach": "VolumeAttach"
  }
 },
 {
  "pkg": "openstack",
  "mod": "containerinfra/cluster",
  "fqn": "pulumi_openstack.containerinfra",
  "classes": {
   "openstack:containerinfra/cluster:Cluster": "Cluster"
  }
 },
 {
  "pkg": "openstack",
  "mod": "containerinfra/clusterTemplate",
  "fqn": "pulumi_openstack.containerinfra",
  "classes": {
   "openstack:containerinfra/clusterTemplate:ClusterTemplate": "ClusterTemplate"
  }
 },
 {
  "pkg": "openstack",
  "mod": "containerinfra/nodeGroup",
  "fqn": "pulumi_openstack.containerinfra",
  "classes": {
   "openstack:containerinfra/nodeGroup:NodeGroup": "NodeGroup"
  }
 },
 {
  "pkg": "openstack",
  "mod": "database/configuration",
  "fqn": "pulumi_openstack.database",
  "classes": {
   "openstack:database/configuration:Configuration": "Configuration"
  }
 },
 {
  "pkg": "openstack",
  "mod": "database/database",
  "fqn": "pulumi_openstack.database",
  "classes": {
   "openstack:database/database:Database": "Database"
  }
 },
 {
  "pkg": "openstack",
  "mod": "database/instance",
  "fqn": "pulumi_openstack.database",
  "classes": {
   "openstack:database/instance:Instance": "Instance"
  }
 },
 {
  "pkg": "openstack",
  "mod": "database/user",
  "fqn": "pulumi_openstack.database",
  "classes": {
   "openstack:database/user:User": "User"
  }
 },
 {
  "pkg": "openstack",
  "mod": "dns/recordSet",
  "fqn": "pulumi_openstack.dns",
  "classes": {
   "openstack:dns/recordSet:RecordSet": "RecordSet"
  }
 },
 {
  "pkg": "openstack",
  "mod": "dns/transferAccept",
  "fqn": "pulumi_openstack.dns",
  "classes": {
   "openstack:dns/transferAccept:TransferAccept": "TransferAccept"
  }
 },
 {
  "pkg": "openstack",
  "mod": "dns/transferRequest",
  "fqn": "pulumi_openstack.dns",
  "classes": {
   "openstack:dns/transferRequest:TransferRequest": "TransferRequest"
  }
 },
 {
  "pkg": "openstack",
  "mod": "dns/zone",
  "fqn": "pulumi_openstack.dns",
  "classes": {
   "openstack:dns/zone:Zone": "Zone"
  }
 },
 {
  "pkg": "openstack",
  "mod": "firewall/groupV2",
  "fqn": "pulumi_openstack.firewall",
  "classes": {
   "openstack:firewall/groupV2:GroupV2": "GroupV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "firewall/policyV2",
  "fqn": "pulumi_openstack.firewall",
  "classes": {
   "openstack:firewall/policyV2:PolicyV2": "PolicyV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "firewall/ruleV2",
  "fqn": "pulumi_openstack.firewall",
  "classes": {
   "openstack:firewall/ruleV2:RuleV2": "RuleV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/applicationCredential",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/applicationCredential:ApplicationCredential": "ApplicationCredential"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/ec2CredentialV3",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/ec2CredentialV3:Ec2CredentialV3": "Ec2CredentialV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/endpointV3",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/endpointV3:EndpointV3": "EndpointV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/groupV3",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/groupV3:GroupV3": "GroupV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/inheritRoleAssignment",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/inheritRoleAssignment:InheritRoleAssignment": "InheritRoleAssignment"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/project",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/project:Project": "Project"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/role",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/role:Role": "Role"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/roleAssignment",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/roleAssignment:RoleAssignment": "RoleAssignment"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/serviceV3",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/serviceV3:ServiceV3": "ServiceV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/user",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/user:User": "User"
  }
 },
 {
  "pkg": "openstack",
  "mod": "identity/userMembershipV3",
  "fqn": "pulumi_openstack.identity",
  "classes": {
   "openstack:identity/userMembershipV3:UserMembershipV3": "UserMembershipV3"
  }
 },
 {
  "pkg": "openstack",
  "mod": "images/image",
  "fqn": "pulumi_openstack.images",
  "classes": {
   "openstack:images/image:Image": "Image"
  }
 },
 {
  "pkg": "openstack",
  "mod": "images/imageAccess",
  "fqn": "pulumi_openstack.images",
  "classes": {
   "openstack:images/imageAccess:ImageAccess": "ImageAccess"
  }
 },
 {
  "pkg": "openstack",
  "mod": "images/imageAccessAccept",
  "fqn": "pulumi_openstack.images",
  "classes": {
   "openstack:images/imageAccessAccept:ImageAccessAccept": "ImageAccessAccept"
  }
 },
 {
  "pkg": "openstack",
  "mod": "index/bgpvpnNetworkAssociateV2",
  "fqn": "pulumi_openstack",
  "classes": {
   "openstack:index/bgpvpnNetworkAssociateV2:BgpvpnNetworkAssociateV2": "BgpvpnNetworkAssociateV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "index/bgpvpnPortAssociateV2",
  "fqn": "pulumi_openstack",
  "classes": {
   "openstack:index/bgpvpnPortAssociateV2:BgpvpnPortAssociateV2": "BgpvpnPortAssociateV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "index/bgpvpnRouterAssociateV2",
  "fqn": "pulumi_openstack",
  "classes": {
   "openstack:index/bgpvpnRouterAssociateV2:BgpvpnRouterAssociateV2": "BgpvpnRouterAssociateV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "index/bgpvpnV2",
  "fqn": "pulumi_openstack",
  "classes": {
   "openstack:index/bgpvpnV2:BgpvpnV2": "BgpvpnV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "index/lbFlavorprofileV2",
  "fqn": "pulumi_openstack",
  "classes": {
   "openstack:index/lbFlavorprofileV2:LbFlavorprofileV2": "LbFlavorprofileV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "index/lbLoadbalancerV2",
  "fqn": "pulumi_openstack",
  "classes": {
   "openstack:index/lbLoadbalancerV2:LbLoadbalancerV2": "LbLoadbalancerV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "keymanager/containerV1",
  "fqn": "pulumi_openstack.keymanager",
  "classes": {
   "openstack:keymanager/containerV1:ContainerV1": "ContainerV1"
  }
 },
 {
  "pkg": "openstack",
  "mod": "keymanager/orderV1",
  "fqn": "pulumi_openstack.keymanager",
  "classes": {
   "openstack:keymanager/orderV1:OrderV1": "OrderV1"
  }
 },
 {
  "pkg": "openstack",
  "mod": "keymanager/secretV1",
  "fqn": "pulumi_openstack.keymanager",
  "classes": {
   "openstack:keymanager/secretV1:SecretV1": "SecretV1"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/l7PolicyV2",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/l7PolicyV2:L7PolicyV2": "L7PolicyV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/l7RuleV2",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/l7RuleV2:L7RuleV2": "L7RuleV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/listener",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/listener:Listener": "Listener"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/member",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/member:Member": "Member"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/members",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/members:Members": "Members"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/monitor",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/monitor:Monitor": "Monitor"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/pool",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/pool:Pool": "Pool"
  }
 },
 {
  "pkg": "openstack",
  "mod": "loadbalancer/quota",
  "fqn": "pulumi_openstack.loadbalancer",
  "classes": {
   "openstack:loadbalancer/quota:Quota": "Quota"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/addressScope",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/addressScope:AddressScope": "AddressScope"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/floatingIp",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/floatingIp:FloatingIp": "FloatingIp"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/floatingIpAssociate",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/floatingIpAssociate:FloatingIpAssociate": "FloatingIpAssociate"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/network",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/network:Network": "Network"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/port",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/port:Port": "Port"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/portForwardingV2",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/portForwardingV2:PortForwardingV2": "PortForwardingV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/portSecGroupAssociate",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/portSecGroupAssociate:PortSecGroupAssociate": "PortSecGroupAssociate"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/qosBandwidthLimitRule",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/qosBandwidthLimitRule:QosBandwidthLimitRule": "QosBandwidthLimitRule"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/qosDscpMarkingRule",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/qosDscpMarkingRule:QosDscpMarkingRule": "QosDscpMarkingRule"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/qosMinimumBandwidthRule",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/qosMinimumBandwidthRule:QosMinimumBandwidthRule": "QosMinimumBandwidthRule"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/qosPolicy",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/qosPolicy:QosPolicy": "QosPolicy"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/quotaV2",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/quotaV2:QuotaV2": "QuotaV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/rbacPolicyV2",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/rbacPolicyV2:RbacPolicyV2": "RbacPolicyV2"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/router",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/router:Router": "Router"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/routerInterface",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/routerInterface:RouterInterface": "RouterInterface"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/routerRoute",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/routerRoute:RouterRoute": "RouterRoute"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/secGroup",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/secGroup:SecGroup": "SecGroup"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/secGroupRule",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/secGroupRule:SecGroupRule": "SecGroupRule"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/subnet",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/subnet:Subnet": "Subnet"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/subnetPool",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/subnetPool:SubnetPool": "SubnetPool"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/subnetRoute",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/subnetRoute:SubnetRoute": "SubnetRoute"
  }
 },
 {
  "pkg": "openstack",
  "mod": "networking/trunk",
  "fqn": "pulumi_openstack.networking",
  "classes": {
   "openstack:networking/trunk:Trunk": "Trunk"
  }
 },
 {
  "pkg": "openstack",
  "mod": "objectstorage/accountV1",
  "fqn": "pulumi_openstack.objectstorage",
  "classes": {
   "openstack:objectstorage/accountV1:AccountV1": "AccountV1"
  }
 },
 {
  "pkg": "openstack",
  "mod": "objectstorage/container",
  "fqn": "pulumi_openstack.objectstorage",
  "classes": {
   "openstack:objectstorage/container:Container": "Container"
  }
 },
 {
  "pkg": "openstack",
  "mod": "objectstorage/containerObject",
  "fqn": "pulumi_openstack.objectstorage",
  "classes": {
   "openstack:objectstorage/containerObject:ContainerObject": "ContainerObject"
  }
 },
 {
  "pkg": "openstack",
  "mod": "objectstorage/tempUrl",
  "fqn": "pulumi_openstack.objectstorage",
  "classes": {
   "openstack:objectstorage/tempUrl:TempUrl": "TempUrl"
  }
 },
 {
  "pkg": "openstack",
  "mod": "orchestration/stackV1",
  "fqn": "pulumi_openstack.orchestration",
  "classes": {
   "openstack:orchestration/stackV1:StackV1": "StackV1"
  }
 },
 {
  "pkg": "openstack",
  "mod": "sharedfilesystem/securityService",
  "fqn": "pulumi_openstack.sharedfilesystem",
  "classes": {
   "openstack:sharedfilesystem/securityService:SecurityService": "SecurityService"
  }
 },
 {
  "pkg": "openstack",
  "mod": "sharedfilesystem/share",
  "fqn": "pulumi_openstack.sharedfilesystem",
  "classes": {
   "openstack:sharedfilesystem/share:Share": "Share"
  }
 },
 {
  "pkg": "openstack",
  "mod": "sharedfilesystem/shareAccess",
  "fqn": "pulumi_openstack.sharedfilesystem",
  "classes": {
   "openstack:sharedfilesystem/shareAccess:ShareAccess": "ShareAccess"
  }
 },
 {
  "pkg": "openstack",
  "mod": "sharedfilesystem/shareNetwork",
  "fqn": "pulumi_openstack.sharedfilesystem",
  "classes": {
   "openstack:sharedfilesystem/shareNetwork:ShareNetwork": "ShareNetwork"
  }
 },
 {
  "pkg": "openstack",
  "mod": "vpnaas/endpointGroup",
  "fqn": "pulumi_openstack.vpnaas",
  "classes": {
   "openstack:vpnaas/endpointGroup:EndpointGroup": "EndpointGroup"
  }
 },
 {
  "pkg": "openstack",
  "mod": "vpnaas/ikePolicy",
  "fqn": "pulumi_openstack.vpnaas",
  "classes": {
   "openstack:vpnaas/ikePolicy:IkePolicy": "IkePolicy"
  }
 },
 {
  "pkg": "openstack",
  "mod": "vpnaas/ipSecPolicy",
  "fqn": "pulumi_openstack.vpnaas",
  "classes": {
   "openstack:vpnaas/ipSecPolicy:IpSecPolicy": "IpSecPolicy"
  }
 },
 {
  "pkg": "openstack",
  "mod": "vpnaas/service",
  "fqn": "pulumi_openstack.vpnaas",
  "classes": {
   "openstack:vpnaas/service:Service": "Service"
  }
 },
 {
  "pkg": "openstack",
  "mod": "vpnaas/siteConnection",
  "fqn": "pulumi_openstack.vpnaas",
  "classes": {
   "openstack:vpnaas/siteConnection:SiteConnection": "SiteConnection"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "openstack",
  "token": "pulumi:providers:openstack",
  "fqn": "pulumi_openstack",
  "class": "Provider"
 }
]
"""
)
