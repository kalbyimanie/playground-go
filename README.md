# playground-go
### required environment variables
set the following environment variables for service runtime
| env var name  | value type |
| ------------- | ------------- |
| PORT_NUMBER  | string  |
<br /><br />

### how to debug the app
```shell
$ cd playground-go/; export PORT_NUMBER=80; go run main.go
```

### how to build the application
**1. using makefile**<br /> <br />
we need to build the image first:
```shell
$ make build-builder && make build-service
```
set the port number for container and host to be bound and exposed.
```shell
$ export PORT_NUMBER=80
```
run the application container.
```shell
$ make run-service
```
<br /><br />**2. docker compose**
```shell
$ docker-compose up --build
```
