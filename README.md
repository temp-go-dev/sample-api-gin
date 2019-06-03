# sample-api-gin

[![Build Status](https://travis-ci.org/temp-go-dev/sample-api-gin.svg?branch=master)](https://travis-ci.org/temp-go-dev/sample-api-gin)

[![codebeat badge](https://codebeat.co/badges/42b2d504-40ed-48ba-bb81-6316072bf29c)](https://codebeat.co/projects/github-com-temp-go-dev-sample-api-gin-master)

[![CodeFactor](https://www.codefactor.io/repository/github/temp-go-dev/sample-api-gin/badge)](https://www.codefactor.io/repository/github/temp-go-dev/sample-api-gin)

## Build

__Build Package__

```bash
go build -a -installsuffix cgo -o sample-api-gin .
```

__Docker Build__

```bash
docker build -t sample-api-gin:0.0.1 .
docker images
REPOSITORY                           TAG                 IMAGE ID            CREATED             SIZE
sample-api-gin                       0.0.1               402c4f49de4d        16 minutes ago      18.5MB
```

## Run

__Docker Container Run__

```bash
docker run --rm -d -e MYSQL_ROOT_PASSWORD=password --name mysql -p 3306:3306 sample-db_mysql
docker run --rm -d --link mysql -name sample-api-gin -p 8080:8080 sample-api-gin:0.0.1
```

MySQLをあらかじめ起動およびマイグレーションしておく必要がある。  
なお、MySQLは[sample-db](https://github.com/temp-go-dev/sample-db)で作成したものを使っている。

## Kick API

CURLで打ち込んでみる。
あらかじめデータを登録しておかなければ404になる。

```bash
$ curl -v localhost:8080/users
* Expire in 0 ms for 6 (transfer 0xf030b0)
* Expire in 1 ms for 1 (transfer 0xf030b0)
* Expire in 1 ms for 1 (transfer 0xf030b0)
* Expire in 2 ms for 1 (transfer 0xf030b0)
*   Trying ::1...
* TCP_NODELAY set
* Expire in 149995 ms for 3 (transfer 0xf030b0)
* Expire in 200 ms for 4 (transfer 0xf030b0)
* Connected to localhost (::1) port 8080 (#0)
> GET /users HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.0
> Accept: */*
>
< HTTP/1.1 404 Not Found
< Date: Thu, 30 May 2019 08:25:05 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```