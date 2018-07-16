# Contributing to Pulumi

Do you want to hack on Pulumi?  Awesome!  We are so happy to have you.

Please refer to the [main Pulumi repo](https://github.com/pulumi/pulumi/)'s [CONTRIBUTING.md file](
https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md) for details on how to do so.

## Committing Generated Code

Code generated for Pulumi SDKs should be checked in as part of the pull request containing a
particular change. To generate code after making changes, run `make` from the root of this
repository.

If a large number of seemingly-unrelated diffs are produced by `make` (for example, lots of changes
to comments unrelated to the change you are making), ensure that the latest dependencies for the
provider are installed by running `make ensure` in the root of the repository.

## Running Integration Tests

The examples and integration tests in this repository will create and destroy real OpenStack
cloud resources while running. Before running these tests, make sure that you have
configured environment variables for deploying to your target OpenStack cloud provider.  For example, you can use the following as created by an "OpenStack configuration file" provided by your OpenStack provider:

* `OS_AUTH_URL`
* `OS_IDENTITY_API_VERSION`
* `OS_TENANT_ID`
* `OS_TENANT_NAME` 
* `OS_USERNAME`
* `OS_PASSWORD`
* `OS_REGION_NAME`

After you have these environment variables set, `make test_all` will run all integration tests.

Test runs in TravisCI currently deploy to [OVH](https://www.ovh.com/world/public-cloud/instances/).