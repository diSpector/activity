### Generate structs for grpc
```sh
protoc --go_out=pkg/activity/grpc --go-grpc_out=pkg/activity/grpc --go_opt=module=github.com/diSpector/activity.git/pkg/activity/grpc --go-grpc_opt=module=github.com/diSpector/activity.git/pkg/activity/grpc api/proto/activity.proto
```

### Run
Run server:
```sh
cd cmd/server
go run main.go
```

Run cli:
```sh
cd cmd/cli
go run main.go activity
go run main.go activity -s 
go run main.go activity add -p 5 -d "Learn Golang"
go run main.go activity search -d "Learn Golang"
go run main.go activity list
```

### Migrate
```sh
migrate -source file://migrations/server -database sqlite://storage/sqlite.db up
migrate -source file://migrations/server -database sqlite://storage/sqlite.db down
```

### Test Pyramid
![test pyramid](./assets/pyramid.png)

### Mockery
> https://github.com/vektra/mockery

