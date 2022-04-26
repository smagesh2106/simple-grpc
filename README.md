# simple-grpc
grpc example
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative      proto/greeter.proto

protoc -I=$PWD --go_out=$PWD  --go-grpc_out=$PWD  $PWD/proto/*.proto
