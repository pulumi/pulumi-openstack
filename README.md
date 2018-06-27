[![Build Status](https://travis-ci.com/pulumi/pulumi-openstack.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-openstack)

# OpenStack Resource Provider

The OpenStack resource provider for Pulumi lets you use OpenStack resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/openstack

or `yarn`:

    $ yarn add @pulumi/openstack

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_openstack

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-openstack/sdk/go/...

## Concepts

The `@pulumi/openstack` package provides a strongly-typed means to build cloud applications that create
and interact closely with OpenStack resources.  Resources are exposed for the entire OpenStack surface area,
including (but not limited to), 'blockstorage', 'compute', 'identity', 'loadbalancer', and more.

## Reference

For detailed reference documentation, please visit [the API docs](
https://pulumi.io/reference/pkg/nodejs/@pulumi/openstack/index.html).
