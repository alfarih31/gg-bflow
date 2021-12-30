# GG-BFlow (Go GRPC Buffer Flow)

[![Go Reference](https://pkg.go.dev/badge/github.com/alfarih31/gg-bflow.svg)](https://pkg.go.dev/github.com/alfarih31/gg-bflow)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/alfarih31/gg-bflow?style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/alfarih31/gg-bflow?style=flat-square)

`GG-BFlow` will behave `Hub-like`/`Messaging-like Protocol`/`Streaming-like Protocol` for streaming your `buffer` data to your client.

`GG-BFlow` utilizing some of the technologies, such as:
1. [gRPC](https://grpc.io/)
2. [memcached](https://memcached.org/)
3. [mongodb](https://www.mongodb.com/)

## Pre
1. gRPC `v1.4.3`
2. Memcached `v1.6.12`
3. MongoDB `v5.0.5`
4. make

## Setup

### Install Dependencies

If you are using unix based system, you can install this project dependencies by running `make` command in root of this project directory

```shell
make install
```

### Configure

#### Configure GG-BFlow

`GG-BFlow` load configuration from `.env` or `system wide Environment Variables`. See [.env.example](.env.example) for `.env` template.

Belows are list & description of needed configuration:

|          | Key                     | Description                                                                                                                                                                                                                                    | Required | Remarks                                                        |
|----------|-------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|----------------------------------------------------------------|
| LOG      | LOG_FORMAT              | Log formatting. Available formats: `console`, `json`                                                                                                                                                                                           |          | String. Default `json`                                         |
|          | LOG_LEVEL               | Log level filter. Available levels: `debug`, `error`, `info`, `warn`                                                                                                                                                                           |          | String. Default `info`                                         |
|          |                         |                                                                                                                                                                                                                                                |          |                                                                |
| BFlow    | BFLOW_HOSTNAME          | Hostname where `GG-BFlow` is running                                                                                                                                                                                                           |          | String. Empty for localhost                                    |
|          | BFLOW_PORT              | Port to access `GG-BFlow` gRPC connection                                                                                                                                                                                                      |          | Number. Default `50051`                                        |
|          | BFLOW_API_KEY           | Key to securing access to `GG-BFlow`                                                                                                                                                                                                           | **✓**    | String                                                         |
|          | BFLOW_AUTHORIZED_CLIENT | List of authorized client name which will access `GG-BFlow`. Client need to connect with `Authorization`  in `metadata` with `Basic Auth` format. Client name as `username` and `BFLOW_API_KEY` as `password`. Ex: `CLIENT_NAME:BFLOW_API_KEY` | **✓**    | String Array (comma `,` separated). Example: `client1,client2` |
|          | BFLOW_BUFFER_SIZE_LIMIT | Limit maximum size of `buffer`                                                                                                                                                                                                                 | **✓**    | Number. Use `0` for no limit                                   |
|          | BFLOW_BUFFER_EXP        | Buffer will remain exist in this seconds                                                                                                                                                                                                       | **✓**    | Number. Use `0` for no expiration                              |
|          |                         |                                                                                                                                                                                                                                                |          |                                                                |
| Mongo    | MONGO_DATABASE          | MongoDB database name                                                                                                                                                                                                                          | **✓**    | String                                                         |
|          | MONGO_HOST              | MongoDB host                                                                                                                                                                                                                                   | **✓**    | String                                                         |
|          | MONGO_PORT              | MongoDB port                                                                                                                                                                                                                                   | **✓**    | Number                                                         |
|          | MONGO_USER              | MongoDB username                                                                                                                                                                                                                               | **✓**    | String                                                         |
|          | MONGO_PASS              | MongoDB password                                                                                                                                                                                                                               | **✓**    | String                                                         |
|          |                         |                                                                                                                                                                                                                                                |          |                                                                |
| Memcache | MEMCACHE_HOST           | Memcache Host                                                                                                                                                                                                                                  | **✓**    | String                                                         |
|          | MEMCACHE_PORT           | Memcache port                                                                                                                                                                                                                                  | **✓**    | Number                                                         |
|          |                         |                                                                                                                                                                                                                                                |          |                                                                |

## Development

### Docker Build

You can build `GG-BFlow` using `docker` by build & run this project using `docker-compose` config in `deployments/docker-compose.yaml`. Belows are step to build this project using `docker-compose`

1. Create `docker bridge network` 

```shell
docker network create gg-bflow-network
```

2. Build this project

```shell
docker-compose build -f ./deployments/docker-compose.yaml
```

## Installation

To install this package, you need to install Go (**version 1.17+ is required**) & initiate your Go workspace first.

1. After you initiate your workspace then you can install this package with below command.

```shell
go get -u github.com/alfarih31/gg-bflow
```

2. Import it in your code

```go
import "github.com/alfarih31/gg-bflow"
```

## Contributors ##

- Alfarih Faza <alfarihfz@gmail.com>

## License

This project is licensed under the - see the [LICENSE.md](LICENSE.md) file for details