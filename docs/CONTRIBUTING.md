# Contributing to revealgo

:tada: First off, thanks for taking the time to contribute! :tada:

## How to build this project

As any other Golang written binary, `revealgo` can be build on different
manners.

### Local environment build

This one was already addressed on the [README.md](../README.md) file. Basically
what you need is to perform the followings steps to have your local `revealgo`
binary:

1. Clone this repository
2. Update git-submodule references
3. Either `go build [...]` or `go install [...]` this repository

```shell
$ git clone https://github.com/yusukebe/revealgo.git
$ cd revealgo
$ git submodule update --init --recursive
$ go install ./cmd/revealgo
```

### Cross-platform binaries build

For CI and Release builds another tool, called
[goreleaser](https://goreleaser.com/intro/), is used for orchestating the
cross-platform binaries management. Assuming you already have said tool in your
local environment, a snapshot build can be done by:

1. Clone this repository
2. Run `goreleaser` build for snapshots using the `--snapshot` flag

```shell
$ git clone https://github.com/yusukebe/revealgo.git
$ cd revealgo
$ goreleaser build --snapshot --rm-dist
$ tree dist
dist
├── config.yaml
├── revealgo_darwin_amd64
│   └── revealgo
├── revealgo_darwin_arm64
│   └── revealgo
├── revealgo_linux_386
│   └── revealgo
├── revealgo_linux_amd64
│   └── revealgo
├── revealgo_linux_arm_6
│   └── revealgo
├── revealgo_linux_arm64
│   └── revealgo
├── revealgo_windows_386
│   └── revealgo.exe
└── revealgo_windows_amd64
└── revealgo.exe
```

**NOTE:** `goreleaser` has an strong opinion about where the binary version
should come from, which is the git tag. Bear this in mind when you use the
`goreleaser release` subcommand.
