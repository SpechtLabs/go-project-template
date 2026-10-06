# PROJECT_NAME

PROJECT_NAME does one thing, and does it well.

[![CI](https://github.com/SpechtLabs/PROJECT_NAME/actions/workflows/ci.yaml/badge.svg)](https://github.com/SpechtLabs/PROJECT_NAME/actions/workflows/ci.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/spechtlabs/PROJECT_NAME.svg)](https://pkg.go.dev/github.com/spechtlabs/PROJECT_NAME)
[![codecov](https://codecov.io/gh/SpechtLabs/PROJECT_NAME/graph/badge.svg)](https://codecov.io/gh/SpechtLabs/PROJECT_NAME)

<!-- template:begin -->

## Using this template

This is the template new SpechtLabs Go projects start from. It holds the tooling every project shares, as [sigil](https://github.com/SpechtLabs/sigil) runs it, around a minimal command line: a root command and `PROJECT_NAME version`. Everything a project needs is wired up and green from the first commit, so a new project starts with its own code rather than with CI.

### Start a project

1. Create the repository from the template and clone it:

   ```shell
   gh repo create SpechtLabs/my-tool --template SpechtLabs/go-project-template --public --clone
   cd my-tool
   ```

2. Rename the project. The script replaces the `PROJECT_NAME` placeholder in every tracked file, renames `cmd/PROJECT_NAME`, removes this section and deletes itself:

   ```shell
   ./scripts/init.sh my-tool
   ```

3. Replace the placeholder text the script can't write for you: the one-line description ("does one thing, and does it well") in this README, `internal/cli/cli.go`, `.goreleaser.yaml` and `docs/`, and the docs pages themselves.

4. Check that everything still passes, and commit:

   ```shell
   mise trust && mise install
   mise run check
   git commit -am "chore: start my-tool from go-project-template"
   ```

5. Set up the repository on GitHub:
   - **Merging:** allow squash merges only, with the pull request title and description as the commit message, and allow auto-merge (Renovate automerges patch and minor updates once CI passes).
   - **Secrets:** `CODECOV_TOKEN` (the coverage upload; optional for a public repository), `RELEASE_PLEASE_TOKEN` (a fine-grained token with read and write access to contents, pull requests and issues, so the release pull request's CI runs; release.yaml explains why) and `HOMEBREW_TAP_GITHUB_TOKEN` (write access to SpechtLabs/homebrew-tap, for the cask). The organization may already provide them.
   - **Renovate:** make sure the Renovate app covers the new repository.
   - **Docs:** set up the project's site on StaticPages (pages.specht-labs.de) before the first deploy, or delete `docs/` and `.github/workflows/docs-website.yaml` if the project has no website.
   - **Docs sections:** the site's shared components come from [docs-kit](https://github.com/SpechtLabs/docs-kit). To show the project's contributors or releases on the home page, add `SpechtLabs/my-tool` to the `repos` list of `docsKitPlugin` in `docs/.vuepress/config.ts` once the repository is public. The build fetches every listed repository from the GitHub API and fails in CI when it can't.

The release and docs deploy jobs skip in go-project-template itself, so the template never tags a version or publishes a cask. A repository created from it has another name, so they run there with nothing to change.

### What's in the box

| File | What it does |
| ---- | ------------ |
| `.mise.toml` | Every tool, pinned to an exact version, and the tasks that build and check the project. CI runs the same tasks. `mise tasks` lists them. |
| `.golangci.yaml`, `.custom-gcl.yml` | golangci-lint v2 with the [golint-sl](https://golint.specht-labs.de/) plugin, built into `./custom-gcl` by `mise run lint-build`. |
| `.github/workflows/ci.yaml` | What a pull request runs; it calls `go.yaml`. |
| `.github/workflows/go.yaml` | Lint, unit tests with the race detector and coverage for Codecov, and a GoReleaser snapshot build. |
| `.github/workflows/release.yaml` | release-please, gated on `go.yaml`, and GoReleaser once a release is created. A manual run with a tag repairs a release that failed to publish. |
| `.github/workflows/pr-hygiene.yaml` | Checks the pull request title is a Conventional Commit, checks release-please can parse the squash commit (`.github/commit-lint/`), and labels the pull request (`.github/labeler.yml`). |
| `.github/workflows/docs-website.yaml` | Lints and builds the VuePress site in `docs/`, and deploys it from main and once a day, so the GitHub data docs-kit fetches at build time stays current. |
| `.goreleaser.yaml` | Linux and macOS binaries for amd64 and arm64, a cosign-signed `checksums.txt`, and a Homebrew cask in SpechtLabs/homebrew-tap. |
| `.release-please-config.json` | Versions and changelog from Conventional Commits; `docs/` changes don't cut releases. |
| `renovate.json` | Keeps Go modules, tools, actions and docs dependencies current, automerging all but majors. |
| `codecov.yml` | Fails a pull request whose project or patch coverage drops below 85%. |
| `.pre-commit-config.yaml`, `.editorconfig`, `.yamllint.yaml`, `.markdownlint-cli2.yaml` | Editor and commit-time checks that match CI. |

### Leaving pieces out

- **A library with no binary:** delete `cmd/`, `.goreleaser.yaml`, the `build` job in `go.yaml`, the `goreleaser` job in `release.yaml`, and the GoReleaser tasks in `.mise.toml`. [go-otel-utils](https://github.com/SpechtLabs/go-otel-utils) shows the result, for several modules.
- **No website:** delete `docs/`, `.github/workflows/docs-website.yaml`, the docs tasks in `.mise.toml`, and the docs rules in `renovate.json` and `.github/labeler.yml`.

<!-- template:end -->

## Install

```shell
brew install spechtlabs/tap/PROJECT_NAME
```

Release archives for Linux and macOS are on the [releases page](https://github.com/SpechtLabs/PROJECT_NAME/releases). The [installation guide](docs/guides/install.md) covers checking their signatures and building from source.

## Usage

```shell
PROJECT_NAME version
```

`PROJECT_NAME --help` lists every command, and the [command line reference](docs/reference/cli.md) describes them.

## Development

The tools come from [mise](https://mise.jdx.dev/), pinned in `.mise.toml`:

```shell
mise trust && mise install
```

| Task | What it does |
| ---- | ------------ |
| `mise run build` | Build `bin/PROJECT_NAME`. |
| `mise run test` | Run the tests with the race detector, writing `coverage.txt`. |
| `mise run lint` | Lint Go with golangci-lint and golint-sl, YAML with yamllint, and the workflows with actionlint. |
| `mise run fmt` | Format go.mod, the Go sources and the Markdown. |
| `mise run check` | Everything CI checks. Run it before you push. |
| `mise run snapshot` | Build every release binary into `dist/` with GoReleaser. |
| `mise run docs-dev` | Serve the documentation website with hot reload. |
| `mise run update-gha` | Pin new GitHub Actions to commit SHAs. |

`pre-commit install` runs the commit-time checks on every commit.

## Contributing

Pull requests are squash-merged, with the pull request title as the commit subject and its description as the body. [release-please](https://github.com/googleapis/release-please) reads them to decide the next version and write the changelog, so the title must be a [Conventional Commit](https://www.conventionalcommits.org/en/v1.0.0/): `feat:` and `fix:` for changes users see, and `chore:`, `ci:`, `build:`, `docs:`, `refactor:` or `test:` for the rest, which don't cut a release.

## License

[Apache License 2.0](LICENSE)
