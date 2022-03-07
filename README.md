# [Kata](https://github.com/dbtedman/kata) // [Spectacle](https://github.com/dbtedman/kata-spectacle)

> ⚠️ WARNING: Not production ready code.

[![CI GitHub Pipeline](https://img.shields.io/github/workflow/status/dbtedman/kata-spectacle/ci?style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/kata-spectacle/actions/workflows/ci.yml)

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

-   **Specification** - An [OpenAPI](https://www.openapis.org) specification, `openapi.json` or `openapi.yaml`.
-   **Specification Rendering** - A Specification that has been rendered into html + js + css for display. A reference to
    a directory in the file system where the rendered files are contained will be maintained by this entity.
-   **Git Platform** (GitLab, GitHub, Bitbucket)
-   **Git Platform Group** - A collection of one or more Git Repositories contained within a Git Platform.
-   **Git Repository**
-   **Git Commit**
-   **Git Tag**
-   **Git Branch**
-   **User** (Unauthenticated, Standard, Administrator, Autonomous)
-   **Historical Fact** - Interactions with this application are recorded as a series of Historical Facts (Events) that
    can be replayed to populate each projection. This represents an event sourced pattern.
-   **Projection** - A view of the data populated from the log of Historical Facts.

### Domain Use Cases

-   **Discover New Specifications**
-   **Check for Updated Specifications**
-   **Render Specification**
-   **Search for Specifications**
-   **Browse Specifications**

## License

The [MIT License](./LICENSE.md) is used by this project.
