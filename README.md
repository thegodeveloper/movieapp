# Microservices movieapp

## Running Consult in Docker

```shell
docker pull consul:1.15.4
```

```shell
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=dev-consul \
    consul:1.15.4 agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

## Run Consul UI

- http://localhost:8500/

## Run the services

### Run Metadata Service

```shell
cd metadata/cmd
go run *.go
```

### Run Rating Service

```shell
cd rating/cmd
go run *.go
```

### Run Movie Service

```shell
cd movie/cmd
go run *.go
```

### How to start more services

```shell
go run *.go --port 8093
```

## References

- https://developer.hashicorp.com/consul/tutorials/day-0/docker-container-agents
