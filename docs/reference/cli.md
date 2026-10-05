---
title: Command line
icon: mdi:terminal
createTime: 2026/10/05 12:00:00
permalink: /reference/cli/
---

Every `PROJECT_NAME` command, its arguments and its flags. Every command also takes `--help`.

## PROJECT_NAME version

Shows the release version, plus the commit, commit time, Go version and platform the binary was built with.

```shell
PROJECT_NAME version
```

Binaries built with `go run` don't carry the commit details, so they report those fields as `unknown`.

## Exit status

| Status | Meaning |
| ------ | ------- |
| 0 | The command succeeded. |
| 1 | The command failed, after printing why. |
