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
$ goreleaser build --snapshot --clean
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

## How to cut a new release

As mentioned in the section above, this project relies on
[goreleaser](https://goreleaser.com/intro/) for cross-platform builds, packaging
and release.

The release process will be triggered **once a new git tag is pushed** to
GitHub, regardless of the branch it is associated with. The process goes as follows:

1. A new git tag is created following [semver](https://semver.org/) conventions
2. The git tag is pushed to `revealgo`'s repository
3. A GitHub action will take care of the release process:
   - Run linting
   - Run `goreleaser release --clean`
   - Publish new release as a [draft](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository)

Upon success, the maintainer will take care of updating the release information
and changing it out from draft mode.

To test the release process locally you can run any of the following commands:

```
# Run release process WITHOUT checking for unstagged changes and AVOIDING to plublish
$ goreleaser release --rm-dist --skip-validate --skip-publish

# Run FULL release process. Requires the GITHUB_TOKEN env. variable.
# Read more here: https://goreleaser.com/scm/github/
$ goreleaser release --rm-dist
```