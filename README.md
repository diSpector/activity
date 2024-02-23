### Generate structs for grpc
```sh
protoc --go_out=pkg/activity/grpc --go-grpc_out=pkg/activity/grpc --go_opt=module=github.com/diSpector/activity.git/pkg/activity/grpc --go-grpc_opt=module=github.com/diSpector/activity.git/pkg/activity/grpc  api/proto/activity.proto
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
```