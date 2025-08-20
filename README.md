# proto

#golang proto gen

protoc --proto_path=proto  --go_out=. --go-grpc_out=require_unimplemented_servers=false:.    user/models/sign_in.proto <br />
protoc --proto_path=proto  --go_out=. --go-grpc_out=require_unimplemented_servers=false:.    user/service.proto <br />

#typescript proto gen
    
#window

npx protoc   --proto_path=proto   --ts_proto_out=pb/typescript   --ts_proto_opt=outputServices=grpc-js,esModuleInterop=true -I proto proto\user\models\sign_in.proto
npx protoc   --proto_path=proto   --ts_proto_out=pb/typescript   --ts_proto_opt=outputServices=grpc-js,esModuleInterop=true -I proto proto\user\service.proto

