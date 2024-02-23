### Generate structs for grpc
```sh
protoc --go_out=pkg/activity/grpc --go-grpc_out=pkg/activity/grpc --go_opt=module=github.com/diSpector/activity.git/pkg/activity/grpc --go-grpc_opt=module=github.com/diSpector/activity.git/pkg/activity/grpc  api/proto/activity.proto
```