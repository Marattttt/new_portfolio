// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var fastrunner_pb = require('./fastrunner_pb.js');

function serialize_fastrunner_GoRequest(arg) {
  if (!(arg instanceof fastrunner_pb.GoRequest)) {
    throw new Error('Expected argument of type fastrunner.GoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_fastrunner_GoRequest(buffer_arg) {
  return fastrunner_pb.GoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_fastrunner_RunResponse(arg) {
  if (!(arg instanceof fastrunner_pb.RunResponse)) {
    throw new Error('Expected argument of type fastrunner.RunResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_fastrunner_RunResponse(buffer_arg) {
  return fastrunner_pb.RunResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var FastRunnerService = exports.FastRunnerService = {
  runGoLang: {
    path: '/fastrunner.FastRunner/RunGoLang',
    requestStream: false,
    responseStream: false,
    requestType: fastrunner_pb.GoRequest,
    responseType: fastrunner_pb.RunResponse,
    requestSerialize: serialize_fastrunner_GoRequest,
    requestDeserialize: deserialize_fastrunner_GoRequest,
    responseSerialize: serialize_fastrunner_RunResponse,
    responseDeserialize: deserialize_fastrunner_RunResponse,
  },
};

exports.FastRunnerClient = grpc.makeGenericClientConstructor(FastRunnerService);
