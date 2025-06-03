# webconsole-proto

PoC for the Aether network function configuration system.

This PoC implements the AccessAndMobility endpoint which returns 
- PLMN ID list
- SNSSAI list
- TAC list

Using 2 types of services:
1. [gRPC](#grpc-service)
2. [REST HTTP](#rest-http-service)

The server starts running and response to the received requests.
The client sends a request every 5 seconds.

## gRPC Service

### Build

``` 
go build -o bin/grpc-server grpc/cmd/server/main.go
go build -o bin/grpc-client grpc/cmd/client/main.go
```

### Run Server

```
./bin/grpc-server
```

### Run Client

```
./bin/grpc-client
```

### Generate Code

```
protoc --go_out=. --go-grpc_out=. proto/config.proto
```

## REST HTTP Service

### Build

``` 
go build -o bin/rest-server rest/cmd/server/main.go
go build -o bin/rest-client rest/cmd/client/main.go
```

### Run Server

```
./bin/rest-server
```

### Run Client

```
./bin/rest-client
```

### Generate Code

```
npx openapi-generator-cli version

openapi-generator-cli generate \
  -i ./webconsole-api.yaml \
  -g go-gin-server \
  -o ./webconsole-server \
  --additional-properties=validateRequired=true

```
