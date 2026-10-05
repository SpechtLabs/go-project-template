---
title: Install PROJECT_NAME
icon: mdi:download
createTime: 2026/10/05 12:00:00
permalink: /guides/install/
---

This guide installs the `PROJECT_NAME` binary on Linux or macOS.

## With Homebrew

```shell
brew install spechtlabs/tap/PROJECT_NAME
```

## From a release

Every [release](https://github.com/SpechtLabs/PROJECT_NAME/releases) has an archive per platform, a `checksums.txt`, and a cosign signature over the checksums. Download the archive for your platform and `checksums.txt`, then check the archive and unpack it:

```shell
sha256sum --check --ignore-missing checksums.txt
tar -xzf PROJECT_NAME_*_linux_amd64.tar.gz PROJECT_NAME
```

To check the signature too, download `checksums.txt.sigstore.json` and verify it with cosign:

```shell
cosign verify-blob \
  --bundle checksums.txt.sigstore.json \
  --certificate-identity-regexp '^https://github.com/SpechtLabs/PROJECT_NAME/' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  checksums.txt
```

## From source

With Go installed:

```shell
go install github.com/spechtlabs/PROJECT_NAME/cmd/PROJECT_NAME@latest
```
