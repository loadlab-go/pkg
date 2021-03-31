pb:
	protoc --go_opt=paths=source_relative --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative auth/*.proto
	protoc --go_opt=paths=source_relative --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative user/*.proto

cleanpb:
	find . -name '*.pb.go' -delete
