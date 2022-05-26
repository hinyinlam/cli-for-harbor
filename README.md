# Harbor Go CLI (Unofficial)
<p align="center"><img src="doc/cli-logo.gif?raw=true"/></p>

TLDR: How to install `harbor` command
---

Version:
---
This is Alpha version, NOT for production, no support provided, unofficial CLI.

Standalone binary `harbor`:
---

## Brew
For Mac (ARM64 / AMD64), just:

`brew install hinyinlam/tap/harbor-cli`

## Other OS - Windows/Linux
Download at
[Release Page](https://github.com/hinyinlam/cli-for-harbor/releases)
Then move `harbor` to your PATH

(I have not test if Windows / Linux version works, it likely works)

Tanzu Plugin:
---
Download `tanzu-harbor-plugin-allarch.tar.gz` from [release page](https://github.com/hinyinlam/cli-for-harbor/releases)

To install: 
```
tar xvf tanzu-harbor-plugin-allarch.tar.gz
tanzu-harbor-plugin/install.sh
```

To verify Tanzu plugin install successfully: 

Run `tanzu plugin list`

It should show similar output to:
```
tanzu plugin list
  NAME                DESCRIPTION                                                        SCOPE       DISCOVERY            VERSION  STATUS
  ...
  harbor              Unofficial Harbor CLI                                              Standalone  harbor-local-source  v0.0.4   update available
  ...
```

You should be able to run `tanzu harbor` in your terminal

Both `tanzu harbor` plugin and `harbor` standalone operate in the same way.

# How to use - Authentication:
Before you perform any operation, `harbor` command require your authentication for each request.

An example of login parameters has been included in `example-yamls/login-params.yaml` file.

To login:

Imperative style:

`harbor login --username username --password password --url https://myharbor.local`

Declarative Style

`harbor login -f login-params.yaml`


You can then run `harbor project list` or `tanzu harbor project list` command.

# How to use - Examples:
There are examples in `example-yamls` folder for various request parameter for different CLI call

When you need help for YAML example, there is `--show-request-yaml` for most command.

For example - You want to create a project declaratively:
`harbor project create -f myproject.yaml`

The fields and structure for `myproject.yaml` can be find out using `harbor project create --show-request-yaml`

Also, a visual way is to visit the OpenAPI spec from Harbor, for example [GoHarbor Public Demo env](https://demo.goharbor.io/devcenter-api-2.0)

# How to build:
Makefile are designed to run in Mac OS, tested in MacBook M1 (ARM64)

## Building binary:
`make clean && make all`
## Building the tanzu plugin release:
`make standalone_binary`

## Building the tanzu plugin release:
`make package_tanzu_plugin`

## Releasing to GitHub
Both Standalone and Tanzu Harbor Plugin

`make clean all && make release`

Question?
---
Please feel free to open issues or contact with Hin Lam <hinl@vmware.com>