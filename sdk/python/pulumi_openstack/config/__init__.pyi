# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

allowReauth: Optional[bool]
"""
If set to `false`, OpenStack authorization won't be perfomed automatically, if the initial auth token get expired.
Defaults to `true`
"""

applicationCredentialId: Optional[str]
"""
Application Credential ID to login with.
"""

applicationCredentialName: Optional[str]
"""
Application Credential name to login with.
"""

applicationCredentialSecret: Optional[str]
"""
Application Credential secret to login with.
"""

authUrl: Optional[str]
"""
The Identity authentication URL.
"""

cacertFile: Optional[str]
"""
A Custom CA certificate.
"""

cert: Optional[str]
"""
A client certificate to authenticate with.
"""

cloud: Optional[str]
"""
An entry in a `clouds.yaml` file to use.
"""

defaultDomain: Optional[str]
"""
The name of the Domain ID to scope to if no other domain is specified. Defaults to `default` (Identity v3).
"""

delayedAuth: Optional[bool]
"""
If set to `false`, OpenStack authorization will be perfomed, every time the service provider client is called. Defaults
to `true`.
"""

disableNoCacheHeader: Optional[bool]
"""
If set to `true`, the HTTP `Cache-Control: no-cache` header will not be added by default to all API requests.
"""

domainId: Optional[str]
"""
The ID of the Domain to scope to (Identity v3).
"""

domainName: Optional[str]
"""
The name of the Domain to scope to (Identity v3).
"""

endpointOverrides: Optional[str]
"""
A map of services with an endpoint to override what was from the Keystone catalog
"""

endpointType: Optional[str]

insecure: Optional[bool]
"""
Trust self-signed certificates.
"""

key: Optional[str]
"""
A client private key to authenticate with.
"""

maxRetries: Optional[int]
"""
How many times HTTP connection should be retried until giving up.
"""

password: Optional[str]
"""
Password to login with.
"""

projectDomainId: Optional[str]
"""
The ID of the domain where the proejct resides (Identity v3).
"""

projectDomainName: Optional[str]
"""
The name of the domain where the project resides (Identity v3).
"""

region: Optional[str]
"""
The OpenStack region to connect to.
"""

swauth: Optional[bool]
"""
Use Swift's authentication system instead of Keystone. Only used for interaction with Swift.
"""

tenantId: Optional[str]
"""
The ID of the Tenant (Identity v2) or Project (Identity v3) to login with.
"""

tenantName: Optional[str]
"""
The name of the Tenant (Identity v2) or Project (Identity v3) to login with.
"""

token: Optional[str]
"""
Authentication token to use as an alternative to username/password.
"""

useOctavia: Optional[bool]
"""
If set to `true`, API requests will go the Load Balancer service (Octavia) instead of the Networking service (Neutron).
"""

userDomainId: Optional[str]
"""
The ID of the domain where the user resides (Identity v3).
"""

userDomainName: Optional[str]
"""
The name of the domain where the user resides (Identity v3).
"""

userId: Optional[str]
"""
Username to login with.
"""

userName: Optional[str]
"""
Username to login with.
"""
