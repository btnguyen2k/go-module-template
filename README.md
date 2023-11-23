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

**Default workflow**

- Dependencies are checked and updated weekly by `dependabot.yaml`, and then PRs are automatically merged to `main` branch by `automerge-dependabot.yaml`.
- `ci.yaml` is triggered on every push to any branch to run tests and generate code coverage report.
- Once PR is approved and merged to `release` branch, `release.yaml` is triggered to create a new release. Then a new PR is created to merge `release` branch to `main` branch.

## Usage

1. Create new project from this template:
   - Click on the `Use this template` button to create a new repository from this template.
   - Or use the `gh` command line tool: `gh repo create <your-repository> --template btnguyen2k/go-module-template`
   - (Less preferred method) Or simply clone/fork this repository.
2. Update `go.mod` and `module.go` file to reflect your module name and required Go version.
3. Review other source code and test files, either update them to reflect your module's name and functionality or remove them.
4. `LICENSE.tpl.md`, `README.tpl.md` and `RELEASE-NOTES.tpl.md` files are templates. Update them to reflect your module's name and functionality. Then rename them to `LICENSE.md`, `README.md` and `RELEASE-NOTES.md`.
5. Update other files to suit your needs.
6. Happy coding!

## License

This template is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Contributing & Support

Feel free to create [pull requests](https://github.com/btnguyen2k/go-module-template/pulls) or [issues](https://github.com/btnguyen2k/go-module-template/issues) to report bugs or suggest new features. If you find this project useful, please start it.
