# CoolAdminGo

[![Go Reference](https://pkg.go.dev/badge/github.com/cool-team-official/cool-admin-go.svg)](https://pkg.go.dev/github.com/cool-team-official/cool-admin-go)
[![Build Release](https://github.com/cool-team-official/cool-admin-go/actions/workflows/release.yml/badge.svg)](https://github.com/cool-team-official/cool-admin-go/actions/workflows/release.yml)

CoolAdmin of Go Backend

Under development, not finished

- [docs](https://pkg.go.dev/github.com/cool-team-official/cool-admin-go) - development documentation
- [github](https://github.com/cool-team-official/cool-admin-go) - cool-admin-go 
- [release](https://github.com/cool-team-official/cool-admin-go/releases) - cool-tools compiled version download
- [issues](https://github.com/cool-team-official/cool-admin-go/issues) - Q & A
- [CoolAdmin](https://cool-js.com) - CoolAdmin Official Website
- [GoFrame](https://goframe.org) - GoFrame 

## TODO

- [x] Basic CURD structure construction
- [x] model layers are generated quickly
- [x] Fast generation of service layer
- [x] rapid generation of controller layer
- [x] Middleware error status and error code are unified with node version
- [x] Unification of global error codes
- [x] The controller layer interface input is unified with the node version
- [ ] Optimize development tool installation experience
- [x] service layer logic completion
- [ ] base module logic completion
- [ ] Other project requirements

## Instructions for use

First download and install the Go language environment. If it is not installed, please install it yourself. If it is already installed, please configure the environment variables yourself.

### 1. Install `cool-tools` development tools

install command

```bash
curl -sSL  https://raw.fastgit.org/cool-team-official/cool-admin-go/master/scripts/install-cool-tools.sh | sh
```

Verify

```bash
cool-tools version
```

### 2.Create new Project

```bash
cool-tools init hello-world
```

### 3.Install `gf` tool

```bash
curl -sSL  https://raw.fastgit.org/cool-team-official/cool-admin-go/master/scripts/install-gf.sh | sh
```

### 4.install dependencies

```bash
go mod tidy
```

### 5.start the development environment 

```bash
gf run main.go
```

enjoy it!
