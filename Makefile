protoc-gen-helloworld:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./examples/helloworld/helloworld/helloworld.proto

protoc-gen-route:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./examples/route_guide/routeguide/route_guide.proto

clean-protoc:
	rm -f ./examples/**/**/*.pb.go
	rm -f ./examples/**/**/*_grpc.pb.go
