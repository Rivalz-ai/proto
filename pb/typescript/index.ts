// Export services as namespaces to avoid conflicts
export { ProtocolServiceClient, ProtocolServiceService } from './protocol/service';
export type { ProtocolServiceServer } from './protocol/service';

export { UserServiceClient, UserServiceService } from './user/service';
export type { UserServiceServer } from './user/service';

export { PriceServiceClient, PriceServiceService } from './price/service';
export type { PriceServiceServer } from './price/service';

// Export models as namespaces
export * as ProtocolModels from './protocol/models/protocol';
export * as PriceModels from './price/models/price';
export * as UserModels from './user/models/sign_in'; 