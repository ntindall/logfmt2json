# Publishing logfmt2json

The package follows semantic versioning and uses [`goreleaser`](https://goreleaser.com/)
to publish artifacts to github.

Breaking changes may occur with minor version bumps prior to 1.0.0. To release
the package, run the following:

```sh
export GITHUB_TOKEN="your_github_dot_com_access_token"
NEXT_VERSION=v0.0.X #TODO, fill this in with what version you want to publish
git commit --allow-empty -m "$NEXT_VERSION"
git tag "$NEXT_TAG_VERSION"
git push --follow-tags
goreleaser release
```

The version of the binary is available by running

```sh
./logfmt2json --version
```
