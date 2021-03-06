kudos
===

[![Build Status](https://github.com/darwinz/kudos/workflows/Go/badge.svg)](https://github.com/darwinz/kudos/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/darwinz/kudos)](https://goreportcard.com/report/github.com/darwinz/kudos)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/darwinz/kudos/master/LICENSE)

Go API and Vue.js SPA frontend. GitHub repo keyword search with a feature allowing users to upvote (give kudos) to found repos.  Authenticated through Okta OIDC.

## Prerequisites

- Docker (https://docs.docker.com/get-docker)
- docker-compose (https://docs.docker.com/compose/install)
- yarn (https://classic.yarnpkg.com/en/docs/getting-started)

## Installation / Setup

#### Okta setup

Authentication for the app is provided through an integration with Okta.  Add authentication for the app by signing up for a [free developer account](https://developer.okta.com/signup/) and creating an OIDC application in Okta.

Once logged in, create a new application by clicking "Add Application", then select the "Single-Page App" platform option.  The default application settings should work.

#### Environment Variables

- MONGO_USER
- MONGO_PASS
- MONGO_HOST
- MONGO_DB

Your app is now ready.  Run the following commands to get going

```shell
make setup
make run_server
make run_client
```

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu), macOS, and Windows

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
