# Protobuf TypeScript Definitions

This package contains TypeScript definitions for protobuf services and models generated from `.proto` files.

## Installation

```bash
npm install rivalz-proto
```

## Usage

### Import Services

```typescript
import { 
  ProtocolServiceClient, 
  UserServiceClient 
} from 'rivalz-proto';

// Create a client
const protocolClient = new ProtocolServiceClient(
  'localhost:50051', 
  credentials.createInsecure()
);

const userClient = new UserServiceClient(
  'localhost:50052', 
  credentials.createInsecure()
);
```

### Import Models

```typescript
import { 
  ProtocolModels, 
  PriceModels, 
  UserModels 
} from 'rivalz-proto';

// Use models
const request = new ProtocolModels.GetInfoRequest();
const signInRequest = new UserModels.SignInRequest();
```

### Example: Using UserService

```typescript
import { UserServiceClient, UserModels } from 'rivalz-proto';
import { credentials } from '@grpc/grpc-js';

const client = new UserServiceClient(
  'localhost:50051',
  credentials.createInsecure()
);

const request = new UserModels.SignInRequest();
request.walletAddress = '0x123...';

client.signInWithWallet(request, (error, response) => {
  if (error) {
    console.error('Error:', error);
    return;
  }
  console.log('Response:', response);
});
```

## Available Services

- **ProtocolService**: Protocol-related operations
- **UserService**: User authentication and management

## Available Models

- **ProtocolModels**: Protocol-related data structures
- **PriceModels**: Price-related data structures  
- **UserModels**: User-related data structures

## Dependencies

This package requires:
- `@grpc/grpc-js`: For gRPC client/server functionality
- `protobufjs`: For protobuf encoding/decoding

## License

MIT 