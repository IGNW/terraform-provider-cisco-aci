# Cisco ACI Provider

This repo contains a terraform provider for deploying networks on Cisco hardware via [Cisco ACI](https://www.cisco.com/c/en/us/solutions/data-center-virtualization/application-centric-infrastructure/index.html).

## Background

In modern technology stacks developers and systems engineers find it easy to spin up new networks in the cloud using Terraform. This process is still challenging in more traditional data centers continues to be cumbersome. In recent years, Cisco has released the ACI technology to make the process of setting up and configuring networks simpler and faster. The ACI layer adds
an Application Programmer's Interface (API) and a GUI for network engineers. This module aims to leverage the ACI capability and allow engineers to provision Cisco networks using Terraform.

More information on ACI capabilities can be found [here](docs/CISCO-ACI.md).

## What's a Provider?

Terraform is used to create, manage, and update infrastructure resources such as physical machines, VMs, network switches, containers, and more. Almost any infrastructure type can be represented as a resource in [Terraform](https://www.terraform.io/).

A provider is responsible for understanding API interactions and exposing resources.

## Building the Provider

Clone repository to: `$GOPATH/src/github.com/ignw/terraform-provider-cisco-aci`

```
$ mkdir -p $GOPATH/src/github.com/ignw
$ cd $GOPATH/src/github.com/ignw
$ git clone git@github.com:IGNW/terraform-provider-cisco-aci.git
```

Enter the provider directory, install tools and project dependencies:

```
$ cd $GOPATH/src/github.com/ignw/terraform-provider-cisco-aci
$ make tools
$ make deps
```

Build the provider:

```
$ make build
```

If the build process ends successfully, the provider binary will be generated in the following path: `$GOPATH/bin/terraform-provider-aci`

## Testing the Provider

In order to test the provider, you can simply run `make test`.

In order to run the full suite of acceptance tests, run `make testacc`.

Note: Acceptance tests create real resources, and often cost money to run.

```
$ make testacc
``ACI_URL=https://host:port ACI_USER=admin ACI_PASS=password ACI_ALLOW_INSECURE make testacc
```

## Using the Provider

If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory, run `terraform init` in the folder where you have your `.tf` files to initialize it.

For example:

```
mkdir -p ~/.terraform.d/plugins/linux_amd64
cp $GOPATH/bin/terraform-provider-aci ~/.terraform.d/plugins/linux_amd64/
cd examples/example_1
terraform init -plugin-dir ~/.terraform.d/plugins/linux_amd64
terraform plan
terraform apply
```

## Authentication

The ACI provider offers a flexible way of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

### Static credentials

You can provide your credentials using the aci provider parameters in the terraform file.

```
provider "aci" {
  url            = "https://host:port"
  username       = "user"
  password       = "password"
  allow_insecure = true
  domain         = "mydomain.com"
}
```

Usage:

```
$ terraform plan
```

### Environment variables

You can provide your credentials via the `ACI_URL`, `ACI_USER`, `ACI_PASS`, `ACI_ALLOW_INSECURE` and `ACI_DOMAIN` (optional) environment variables.

```
provider "aci" {}
```

Usage:

```
$ export ACI_URL="https://host:port"
$ export ACI_USER="user"
$ export ACI_PASS="password"
$ export ACI_ALLOW_INSECURE=true
$ export ACI_DOMAIN="mydomain.com"
$ terraform plan
```

## Who maintains this Provider?

This Module is maintained by [IGNW](http://www.ignw.io/). If you're looking for help or commercial
support, send an email to [support@infogroupnw.com](mailto:support@infogroupnw.com?Subject=Cisco%20ACI%20Provider).
IGNW can help with:

* Setup, customization, and support for this Provider.
* Modules for other types of infrastructure, such as VPCs, Docker clusters, databases, and continuous integration.
* Modules that meet compliance requirements, such as FedRamp, HIPAA.
* Consulting & Training on AWS, Azure, GCP, Terraform, and DevOps.

## Code included in this Module:

* [Cisco Terraform Provider](https://github.com/ignw/terraform-provider-cisco-aci): The module includes Terraform code to interact with the Cisco ACI API.

## How is this Provider versioned?

This Module follows the principles of [Semantic Versioning](http://semver.org/). You can find each new release,
along with the changelog, in the [Releases Page](../../releases).

During initial development, the major version will be 0 (e.g., `0.x.y`), which indicates the code does not yet have a
stable API. Once we hit `1.0.0`, we will make every effort to maintain a backwards compatible API and use the MAJOR,
MINOR, and PATCH versions on each release to indicate any incompatibilities.

## License

This code is released under the Mozilla . Please see [LICENSE](https://github.com/ignw/terraform-provider-cisco-aci/tree/master/LICENSE) and [NOTICE](https://github.com/ignw/terraform-provider-cisco-aci/tree/master/NOTICE) for more
details.

Copyright &copy; 2018 InfogroupNW, Inc.
