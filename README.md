# goserver

A basic HTTP server written in Go. Originally, part of the Boot.dev _Learning Docker_ course, it is being elevated as part of the Boot.dev _Learning HTTP Servers in Go_ course.

## build

```
$ go build
```

### build container image

```
$ docker build . --tag [TAG]
```

## run

```
$ GOSERVER_ROOT="./public" GOSERVER_PORT="8080" goserver
```

### run in a container

```
$ docker run --detach --publish="8080:8080" --volume="./public:/var/run/goserver/public" --name local-goserver [TAG|IMAGE ID]
```

