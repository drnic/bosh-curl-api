# bosh curl API

An BOSH Director API library that delegates all authentication and client API calls to the `bosh` CLI.

The library requires that you have `bosh` already installed and are using the appropriate `BOSH_ENVIRONMENT` environment variables.

This library assumes your `bosh` CLI is built with this pull request https://github.com/cloudfoundry/bosh-cli/pull/408 to provide the `bosh curl` sub-command.