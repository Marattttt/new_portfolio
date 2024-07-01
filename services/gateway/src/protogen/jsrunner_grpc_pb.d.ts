// package: jsrunner
// file: jsrunner.proto

import * as grpc from '@grpc/grpc-js';
import * as jsrunner_pb from './jsrunner_pb';

interface IJsRunnerService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  runJs: IJsRunnerService_IRunJs;
}

interface IJsRunnerService_IRunJs extends grpc.MethodDefinition<jsrunner_pb.JsRequest, jsrunner_pb.RunResponse> {
  path: '/jsrunner.JsRunner/RunJs'
  requestStream: false
  responseStream: false
  requestSerialize: grpc.serialize<jsrunner_pb.JsRequest>;
  requestDeserialize: grpc.deserialize<jsrunner_pb.JsRequest>;
  responseSerialize: grpc.serialize<jsrunner_pb.RunResponse>;
  responseDeserialize: grpc.deserialize<jsrunner_pb.RunResponse>;
}

export const JsRunnerService: IJsRunnerService;
export interface IJsRunnerServer extends grpc.UntypedServiceImplementation {
  runJs: grpc.handleUnaryCall<jsrunner_pb.JsRequest, jsrunner_pb.RunResponse>;
}

export interface IJsRunnerClient {
  runJs(request: jsrunner_pb.JsRequest, callback: (error: grpc.ServiceError | null, response: jsrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  runJs(request: jsrunner_pb.JsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jsrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  runJs(request: jsrunner_pb.JsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jsrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
}

export class JsRunnerClient extends grpc.Client implements IJsRunnerClient {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
  public runJs(request: jsrunner_pb.JsRequest, callback: (error: grpc.ServiceError | null, response: jsrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  public runJs(request: jsrunner_pb.JsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: jsrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
  public runJs(request: jsrunner_pb.JsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: jsrunner_pb.RunResponse) => void): grpc.ClientUnaryCall;
}

