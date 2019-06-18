# sample-api-gin

[![Build Status](https://travis-ci.org/temp-go-dev/sample-api-gin.svg?branch=master)](https://travis-ci.org/temp-go-dev/sample-api-gin)

[![codebeat badge](https://codebeat.co/badges/42b2d504-40ed-48ba-bb81-6316072bf29c)](https://codebeat.co/projects/github-com-temp-go-dev-sample-api-gin-master)

[![CodeFactor](https://www.codefactor.io/repository/github/temp-go-dev/sample-api-gin/badge)](https://www.codefactor.io/repository/github/temp-go-dev/sample-api-gin)

## Init

__Update Submodule__

GitのSubmoduleを利用しているため取得します。

```bash
cd {project-root-dir}
git submodule update --init --recursive
```

## Build

__Build Package__

Goの実行バイナリをビルドするためには、以下のコマンドを実行します。

```bash
cd {project-root-dir}
go build -a -installsuffix cgo -o ./build/sample-api-gin .
```

__Docker Build__

Goの実行バイナリを含有したDockerImageをビルドするためには、以下のコマンドを実行します。

```bash
cd {project-root-dir}
docker build -t sample-api-gin:0.0.1 .
docker images
REPOSITORY                           TAG                 IMAGE ID            CREATED             SIZE
sample-api-gin                       0.0.1               402c4f49de4d        16 minutes ago      18.5MB
```

## Run

__Docker Container Run__

ローカル環境にてサンプルを実行するには、以下のコマンドを実行します。  
現在調査中ですが、`sample-api-gin` がDB起動を待つ対応ができていないため、  
`docker-compose up -d` を2回叩いていただく必要がある場合があります。

```bash
$ cd {project-root-dir}
$ docker-compose up -d
$ cd sample-db
$ docker run -v "%CD%\script:/migrations" --network host  migrate/migrate -path=/migrations/ -database mysql://user:password@tcp(localhost:3306)/sampledb up
```
