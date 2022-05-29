# [Kata](https://github.com/dbtedman/kata) // [Spectacle](https://github.com/dbtedman/kata-spectacle)

> ⚠️ WARNING: Not production ready code.

[![ci workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-spectacle/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-spectacle/actions/workflows/ci.yml)
[![sast workflow status](https://img.shields.io/github/workflow/status/dbtedman/kata-spectacle/sast?style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/kata-spectacle/actions/workflows/sast.yml)
![language: go](https://img.shields.io/badge/language-go-blue.svg?style=for-the-badge)

Discover projects within a hosted git platform that contain an OpenAPI Specifications so that an index can be generated.

-   [Getting Started](#getting-started)
-   [Design](#design)
-   [License](#license)

## Getting Started

```shell
nvm use && make && ./spectacle --help
```

## Design

### Concepts

_Placeholder_

### Domain

#### [Git Branch]()

_Placeholder_

#### [Git Commit]()

_Placeholder_

#### [Git Platform Group]()

A collection of one or more Git Repositories contained within a Git Platform.

#### [Git Platform]()

(GitLab, GitHub, Bitbucket)

#### [Git Repository]()

_Placeholder_

#### [Git Tag]()

_Placeholder_

#### [Historical Fact]()

Interactions with this application are recorded as a series of Historical Facts (Events) that can be replayed to populate each projection. This represents an event sourced pattern.

#### [Projection]()

A view of the data populated from the log of Historical Facts.

#### [Specification Rendering]()

A Specification that has been rendered into html + js + css for display. A reference to a directory in the file system where the rendered files are contained will be maintained by this entity.

#### [Specification]()

An [OpenAPI](https://www.openapis.org) specification, `openapi.json` or `openapi.yaml`.

#### [User]()

(Unauthenticated, Standard, Administrator, Autonomous)

### Gateways

#### [GitLab]()

_Placeholder_

### Security Mitigations

> Initially based on the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/).

#### [A01:2021-Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)

[Github Security](https://github.com/features/security) detects secrets incorrectly committed into the repository.

#### [A02:2021-Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)

_Placeholder_

#### [A03:2021-Injection](https://owasp.org/Top10/A03_2021-Injection/)

_Placeholder_

#### [A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)

_Placeholder_

#### [A05:2021-Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)

_Placeholder_

#### [A06:2021-Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)

[Snyk](https://snyk.io) and [Github Security](https://github.com/features/security) scan Gradle and NPM dependencies for know vulnerabilities and create pull requests to resolve the vulnerabilities when available.

#### [A07:2021-Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/)

_Placeholder_

#### [A08:2021-Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)

_Placeholder_

#### [A09:2021-Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)

_Placeholder_

#### [A10:2021-Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)

_Placeholder_

## License

The [MIT License](./LICENSE.md) is used by this project.
