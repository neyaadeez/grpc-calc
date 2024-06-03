proto:
	protoc -Iproto --go_out=. --go_opt=module=github.com/neyaadeez/grpc-calc --go-grpc_out=. --go-grpc_opt=module=github.com/neyaadeez/grpc-calc proto/*.proto

clean:
	rm -rf ./bin/*

gobuild: clean
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client