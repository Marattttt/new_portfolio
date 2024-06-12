// package: fastrunner
// file: fastrunner.proto

import * as grpc from '@grpc/grpc-js';
import * as fastrunner_pb from './fastrunner_pb';

interface IFastRunnerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  runGoLang: IFastRunnerService_IRunGoLang;
}

interface IFastRunnerService_IRunGoLang extends grpc.MethodDefinition<fastrunner_pb.GoRequest, fastrunner_pb.RunResponse> {
  path: '/fastrunner.FastRunner/RunGoLang'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<fastrunner_pb.GoRequest>;
  requestDeserialize: grpc.deserialize<fastrunner_pb.GoRequest>;
  responseSerialize: grpc.serialize<fastrunner_pb.RunResponse>;
  responseDeserialize: grpc.deserialize<fastrunner_pb.RunResponse>;
}

export const FastRunnerService: IFastRunnerService;
export interface IFastRunnerServer extends grpc.UntypedServiceImplementation {
  runGoLang: grpc.handleUnaryCall<fastrunner_pb.GoRequest, fastrunner_pb.RunResponse>;
}

export interface IFastRunnerClient {
  runGoLang(request: fastrunner_pb.GoRequest, callback: (error: grpc.ServiceError | null, response: fastrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  runGoLang(request: fastrunner_pb.GoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: fastrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  runGoLang(request: fastrunner_pb.GoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: fastrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
}

export class FastRunnerClient extends grpc.Client implements IFastRunnerClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
  public runGoLang(request: fastrunner_pb.GoRequest, callback: (error: grpc.ServiceError | null, response: fastrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  public runGoLang(request: fastrunner_pb.GoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: fastrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  public runGoLang(request: fastrunner_pb.GoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: fastrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
}

