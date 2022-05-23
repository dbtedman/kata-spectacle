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

| Env                 |     |
| ------------------- | --- |
| `APIS_GITLAB_TOKEN` |     |
| `APIS_GITLAB_URL`   |     |

```shell
nvm use && make && ./spectacle --help
```

## Design

### Domain Entities

| Entity                   | Notes                                                                                                                                                                                            |
| :----------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `GitBranch`              |                                                                                                                                                                                                  |
| `GitCommit`              |                                                                                                                                                                                                  |
| `GitPlatformGroup`       | A collection of one or more Git Repositories contained within a Git Platform.                                                                                                                    |
| `GitPlatform`            | (GitLab, GitHub, Bitbucket)                                                                                                                                                                      |
| `GitRepository`          |                                                                                                                                                                                                  |
| `GitTag`                 |                                                                                                                                                                                                  |
| `HistoricalFact`         | Interactions with this application are recorded as a series of Historical Facts (Events) that can be replayed to populate each projection. This represents an event sourced pattern.             |
| `Projection`             | A view of the data populated from the log of Historical Facts.                                                                                                                                   |
| `SpecificationRendering` | A Specification that has been rendered into html + js + css for display. A reference to a directory in the file system where the rendered files are contained will be maintained by this entity. |
| `Specification`          | An [OpenAPI](https://www.openapis.org) specification, `openapi.json` or `openapi.yaml`.                                                                                                          |
| `User`                   | (Unauthenticated, Standard, Administrator, Autonomous)                                                                                                                                           |

### Domain Use Cases

| Use Case                        | Notes |
| :------------------------------ | :---- |
| `BrowseSpecifications`          |       |
| `CheckForUpdatedSpecifications` |       |
| `DiscoverNewSpecifications`     |       |
| `RenderSpecification`           |       |
| `SearchForSpecifications`       |       |

### Gateways

| Gateway | Notes |
| :------ | :---- |
| ` `     |       |

### Security Mitigations

> Initially based on the [OWASP Top 10 - 2021](https://owasp.org/www-project-top-ten/).

| Security Risk                                                                                                                       | Mitigation |
| :---------------------------------------------------------------------------------------------------------------------------------- | :--------- |
| [A01:2021-Broken Access Control](https://owasp.org/Top10/A01_2021-Broken_Access_Control/)                                           |            |
| [A02:2021-Cryptographic Failures](https://owasp.org/Top10/A02_2021-Cryptographic_Failures/)                                         |            |
| [A03:2021-Injection](https://owasp.org/Top10/A03_2021-Injection/)                                                                   |            |
| [A04:2021-Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/)                                                       |            |
| [A05:2021-Security Misconfiguration](https://owasp.org/Top10/A05_2021-Security_Misconfiguration/)                                   |            |
| [A06:2021-Vulnerable and Outdated Components](https://owasp.org/Top10/A06_2021-Vulnerable_and_Outdated_Components/)                 |            |
| [A07:2021-Identification and Authentication Failures](https://owasp.org/Top10/A07_2021-Identification_and_Authentication_Failures/) |            |
| [A08:2021-Software and Data Integrity Failures](https://owasp.org/Top10/A08_2021-Software_and_Data_Integrity_Failures/)             |            |
| [A09:2021-Security Logging and Monitoring Failures](https://owasp.org/Top10/A09_2021-Security_Logging_and_Monitoring_Failures/)     |            |
| [A10:2021-Server-Side Request Forgery](https://owasp.org/Top10/A10_2021-Server-Side_Request_Forgery_%28SSRF%29/)                    |            |

## License

The [MIT License](./LICENSE.md) is used by this project.
