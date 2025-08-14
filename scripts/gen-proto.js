const glob = require('glob');
const { execSync } = require('child_process');

const files = glob.sync('proto/**/*.proto'); // tìm tất cả file proto
const cmd = `npx protoc   --proto_path=proto   --ts_proto_out=pb/typescript   --ts_proto_opt=outputServices=grpc-js,esModuleInterop=true -I proto ${files.join(' ')}`;
//glob 'proto/**/*.proto' --proto_path=proto  --ts_proto_opt=outputServices=grpc-js,esModuleInterop=true --grpc_out=pb/typescript {file}
execSync(cmd, { stdio: 'inherit' });
