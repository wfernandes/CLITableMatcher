# CF CLI Table Gomega Matcher

The [Cloud Foundry CLI][cf_cli] contains the `terminal.Table` type in order to print
output in table format onto the terminal.

In order to test the terminal output easily, the `ContainsCLITable` [gomega][gomega] matcher
can be used to verify if the output contains the specified table format.

See the tests or [the sample](/sample/main.go) for example usages.

### Tests

```bash
go test ./... -race
```

[cf_cli]: https://github.com/cloudfoundry/cli
[gomega]: http://onsi.github.io/gomega/