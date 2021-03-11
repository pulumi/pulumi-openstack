CHANGELOG
=========

## HEAD (Unreleased)
* Fix typo for password configuration

---

## 2.16.0 (2021-03-09)
* Upgrade to v1.39.0 of the OpenStack Terraform Provider

## 2.15.0 (2021-02-24)
* Upgrade to v1.38.0 of the OpenStack Terraform Provider

## 2.14.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 2.14.0 (2021-02-16)
* Upgrade to v1.37.0 of the OpenStack Terraform Provider

## 2.13.0 (2021-02-03)
* Upgrade to v1.36.0 of the OpenStack Terraform Provider

## 2.12.0 (2021-01-29)
* Upgrade to pulumi-terraform-bridge v2.18.1

## 2.11.0 (2021-01-19)
* Upgrade to v1.35.0 of the OpenStack Terraform Provider

## 2.10.2 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 2.10.1 (2020-12-22)
* Upgrade to v1.34.1 of the OpenStack Terraform Provider

## 2.10.0 (2020-12-21)
* Upgrade to v1.34.0 of the OpenStack Terraform Provider
* Upgrade to v2.16.0 of pulumi-terraform-bridge which includes
  * Bug fix: Correcting an issue where replacements were not being correctly identified and leading to panics.
  * Bug fix: Corrects a panic caused by reading sets with partially-known elements.
  * Preserve unknowns during provider preview
  * This adds support for import specific examples in documentation  
  **PLEASE NOTE:**
  * `openstack.Orchestration.StackV1` has had an output value rename from `Outputs` to `StackOutputs`.  

## 2.9.0 (2020-11-17)
* Upgrade to v1.33.0 of the OpenStack Terraform Provider

## 2.8.1 (2020-11-06)
* Upgrade to pulumi-terraform-bridge v2.12.1

## 2.8.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.7.0 (2020-10-06)
* Upgrade to v1.32.0 of the OpenStack Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.8.0
* Upgrade to Pulumi v2.10.0

## 2.6.0 (2020-08-31)
* Upgrade to v1.31.0 of the OpenStack Terraform Provider

## 2.5.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.4.0 (2020-08-06)
* Upgrade to v1.30.0 of the OpenStack Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.6.0
* Upgrade to Pulumi v2.7.1

## 2.3.0 (2020-06-30)
* Upgrade to v1.29.0 of the OpenStack Terraform Provider

## 2.2.3 (2020-06-11)
* Switch to GitHub actions for build

## 2.2.2 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.2.1 (2020-05-11)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.2.0 (2020-05-04)
* Upgrade to v1.28.0 of the OpenStack Terraform Provider
  ** Please Note ** in `openstack.containerinfra.Cluster` `NodeAddresses` is now a List not a string
* Upgrade to pulumi-terraform-bridge v2.2.0

## 2.1.0 (2020-04-28)
* Regenerate datasource examples to be async
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-17)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.5.0 (2020-04-01)
- Upgrade to pulumi-terraform-bridge v1.8.4
- Upgrade to Pulumi v1.13.1
- Change layout to support new Go module structure

## 1.4.0 (2020-03-23)
* Upgrade to v1.26.0 of the OpenStack Terraform Provider
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.3.0 (2019-12-17)
* Namespace names in .NET SDK are adjusted to PascalCase
([#70](https://github.com/pulumi/pulumi-openstack/pull/70)).
* Upgrade to v1.25.0 of the OpenStack Terraform Provider
* Upgrade to pulumi-terraform-bridge v1.5.2

## 1.2.0 (2019-12-04)
* Upgrade to support go 1.13.x
* Upgrade to v1.4.3 of the pulumi-terraform-bridge

## 1.1.0 (2019-11-13)
* Add support for DotNet SDK Generation

## 1.0.0 (2019-10-23)
* Regenerate SDK against tf2pulumi 0.6.0
* Upgrade to v1.24.0 of the OpenStack Terraform Provider

## 0.17.13 (2019-09-24)
* Upgrade to v1.23.0 of the OpenStack Terraform Provider

## 0.17.12 (2019-09-05)
* Upgrade to v1.22.0 of the OpenStack Terraform Provider

## 0.17.11 (2019-09-05)
* Upgrade to Pulumi v1.0.0

## 0.17.10 (2019-08-30)
* Regenerate SDKs based on v1.21.1 of the OpenStack Terraform Provider

## 0.17.9 (2019-08-29)
* Upgrade pulumi-terraform to 3f206601e7

## 0.17.8 (2019-08-20)
* Depend on latest pulumi package

## 0.17.7 (2019-08-09)
* Update to pulumi-terraform@9db2fc93cd
* Upgrade to v1.21.1 of the OpenStack Terraform Provider

## 0.17.6 (2019-07-12)
* Upgrade to v1.20.0 of the OpenStack Terraform Provider

## 0.17.5 (2019-07-09)
* Fix detailed diffs with nested computed values.

## 0.17.4 (2019-07-08)
* Communicate detailed information about the difference between a resource's desired and actual state during a Pulumi update

## 0.17.3
* Add TypeScript type guards for each resource class ([7ace3e9b5f](https://github.com/pulumi/pulumi-terraform/commit/7ace3e9b5f2dcd4692b029ba4b80360582d7949a))
* Update to pulumi-terraform@3635bed3a5 which stops maps containing `.` being treated as nested maps.

## 0.17.2 (2019-05-31)
* Update to v1.19.0 of the OpenStack Terraform provider

## 0.17.1 (2019-04-22)
* Update to v1.17.0 of the OpenStack Terraform provider

## 0.17.0 (2019-03-06)
* Update to v1.16.0 of the OpenStack Terraform provider
* Depend on the latest version of the `pulumi` SDK

## 0.16.5 (2019-02-13)
* Update to v1.15.1 of the OpenStack Terraform provider
* Add support for the `deleteBeforeReplace` resource option and improved delete-before-replace behaviour introduced in Pulumi v0.16.14

## 0.16.4 (2019-01-19)
* Update to v1.14.0 of the OpenStack Terraform provider
* Add documentation comments to the Node.js SDK

## 0.16.3 (2018-12-05)
* Update to v1.12.0 of the OpenStack Terraform provider

## 0.16.2 (2018-11-13)
* Support Python 3.6 and higher, instead of only Python 3.7

## 0.16.1 (2018-10-22)
* Depend on the latest version of the `pulumi` SDK

## 0.16.0 (2018-10-16)
* Updated to the latest version of the `pulumi` SDK

## 0.15.0 (2018-08-14)
* Initial version of the OpenStack provoider
