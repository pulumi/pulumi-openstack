CHANGELOG
=========

## HEAD (Unreleased)
_(none)_

---

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
