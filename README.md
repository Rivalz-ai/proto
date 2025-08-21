# proto

#golang proto gen

protoc --proto_path=proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. user/models/sign_in.proto <br />
protoc --proto_path=proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. user/service.proto <br />

#typescript proto gen

```bash
yarn gen

yarn publish
```
