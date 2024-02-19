# go-module-template

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Actions Status](https://github.com/btnguyen2k/go-module-template/workflows/ci/badge.svg)](https://github.com/btnguyen2k/go-module-template/actions)
[![Release](https://img.shields.io/github/release/btnguyen2k/go-module-template.svg?style=flat-square)](RELEASE-NOTES.md)

Template to quickly spin up a Go module project.

## Features

- [Go modules](https://blog.golang.org/using-go-modules) enabled.
- Template for README, LICENSE, RELEASE-NOTES, and .gitignore files.
- GitHub actions included:
  - `dependabot.yaml`, `automerge-dependabot.yaml`: automatically update dependencies and merge PRs from dependabot.
  - `ci.yaml`: automatically run tests and generate code coverage report.
  - `release.yaml`: automatically create a new release.
  - `codeql.yaml`: automatically run CodeQL analysis.

## Usage

1. Create new project from this template:
   - Click on the `Use this template` button to create a new repository from this template.
   - Or use the `gh` command line tool: `gh repo create <your-repository> --template btnguyen2k/go-module-template`
   - (Less preferred method) Or simply clone/fork this repository.
2. Update `go.mod` and `module.go` file to reflect your module name and required Go version.
3. Review other source code and test files, either update them to reflect your module's name and functionality or remove them.
4. `LICENSE.tpl.md`, `README.tpl.md` and `RELEASE-NOTES.tpl.md` files are templates; update them to reflect your module's name and functionality; then rename them to `LICENSE.md`, `README.md` and `RELEASE-NOTES.md`.
5. Update other files to suit your needs.
6. Happy coding!

**Workflows**

Workflows implemented by this template are as the following:

- `dependabot.yaml` configures dependencies are checked and updated weekly. `dependabot` will create a PR for each dependency update. `automerge-dependabot.yaml` is triggered to automatically merged PRs to `main` branch.
- `codeql.yaml` is triggered on every push, pr and periodically to run CodeQL analysis.
- `ci.yaml` is triggered on every push to any branch to run tests and generate code coverage report.
- Once PR is approved and merged to `release` branch, `release.yaml` is triggered to create a new release. Then a new PR is created to merge `release` branch to `main` branch.

A suggested git workflow to use with this template is as the following:

- Work on your code in development/feature branches as usual.
- Once ready, create a PR to merge your development/feature branch to `release` branch.
  - Once the PR is merged, `release.yaml` is triggered to create a new release.
  - Then a new PR is created to merge `release` branch to `main` branch. Note: you have to review and approve the PR by yourself to finalize the merge.

> Remember to enable the setting "Allow GitHub Actions to create and approve pull requests" from project's `Settings -> Actions -> General`.

## License

This template is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Contributing & Support

Feel free to create [pull requests](https://github.com/btnguyen2k/go-module-template/pulls) or [issues](https://github.com/btnguyen2k/go-module-template/issues) to report bugs or suggest new features. If you find this project useful, please star it.
