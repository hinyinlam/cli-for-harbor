# Harbor Go CLI (Unofficial)
<p align="center"><img src="doc/cli-logo.gif?raw=true"/></p>

TLDR: How to install `harbor` command
---
For Mac, just:

brew install hinyinlam/tap/harbor-cli

Other OS pls download at
[Release Page](https://github.com/hinyinlam/cli-for-harbor/releases)

(I have not test if Windows / Linux version works, it likely works)

Version:
---
This is Alpha version, NOT for production, no support provided, unofficial CLI.


Standalone form factor:
---
This standalone form factor is designed to behave the same as the Tanzu Plugin

You can download `cli-for-harbor_$version_$OS_$Arch.tar.gz` from release page

move `harbor` to your PATH for execution

Install as Tanzu Plugin:
---
Download [tanzu-framework-plugins-registry-allarch.tar.gz](https://github.com/hinyinlam/cli-for-harbor-go/releases/download/v0.0.1/tanzu-framework-plugins-registry-allarch.tar.gz) from the release page

To install: 
```
tar xvf tanzu-framework-plugins-registry-allarch.tar.gz

cp -rvf $(PWD)/registry-plugins/distribution/* ~/.config/tanzu-plugins/distribution/

tanzu plugin source add --name registry-local-source --type local --uri $(PWD)/registry-plugins/discovery/registry

tanzu plugin install harbor
```

To verify Tanzu plugin install successfully: 

Run `tanzu plugin list`

It should show similar output to:
```
NAME                DESCRIPTION                                                        SCOPE       DISCOVERY              VERSION  STATUS     
  login               Login to the platform                                              Standalone  default                v0.11.4  installed  
  management-cluster  Kubernetes management-cluster operations                           Standalone  default                v0.11.4  installed  
  package             Tanzu package management                                           Standalone  default                v0.11.4  installed  
  pinniped-auth       Pinniped authentication operations (usually not directly invoked)  Standalone  default                v0.11.4  installed  
  secret              Tanzu secret management                                            Standalone  default                v0.11.4  installed  
  harbor              Build Tanzu components                                             Standalone  registry-local-source  v0.0.1   installed  
  builder             Build Tanzu components                                             Standalone  test                   v0.11.4  installed 

```

Where you can see above output contain `harbor` as one of the installed plugin

You should be able to run `tanzu harbor` in your terminal


Authentication:
---
Before you perform any operation, `harbor` command require your authentication for each request.

An example of login parameters has been included in `login-params.yaml` file.

To login:

Imperative style:

`harbor login --username username --password password --url https://myharbor.local`

Declarative Style

`harbor login -f login-params.yaml`


You can then run `harbor project list` or `tanzu harbor project list` command.


How to build:
===
Makefile are designed to run in Mac OS

Building binary:
---
`make clean && make all`

Building the tanzu plugin release:
---
`make package_tanzu_plugin`


Question?
---
Please feel free to open issues or contact with Hin Lam <hinl@vmware.com>