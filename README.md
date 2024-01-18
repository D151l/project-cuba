![Go icon](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/512px-Go_Logo_Blue.svg.png)

---

# Project CUBA - Chat Utility Bot Assistant

## Description

...

## Getting Started

The discord bot uses the following environment variables:

| Variable Name | Description                        |
|---------------|------------------------------------|
| `token` | The discord authentication token   |
| `statusType` | The type of the discord bot status. See [Status Types](#status-types) |
| `statusContent` | The content of the status          |
| `streamingURL` | The URL of the stream when the status type is `STREAMING` |

### Status Types
- `PLAYING`
- `STREAMING` (requires environment variable `streamingURL`)
- `LISTENING`
- `WATCHING`

## Todo
- [ ] soon

## Installation

### Installations with Docker

The docker image is available on [GitHub Container Registry](https://github.com/d151l/discordbot/pkgs/container/discordbot)

#### Docker CLI

```bash
docker run -d \
    --name project-cuba \
    -e token=your_token \
    ghcr.io/d151l/project-cuba:master
```

#### Docker Compose

```yaml
version: "3.8"
services:
  project-cuba:
        container_name: "project-cuba"
        image: ghcr.io/d151l/project-cuba:master
        restart: always
        pull_policy: always
        environment:
            - token=your_token
```
