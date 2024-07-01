// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var gorunner_pb = require('./gorunner_pb.js');

function serialize_multirunner_RunGoRequest(arg) {
  if (!(arg instanceof gorunner_pb.RunGoRequest)) {
    throw new Error('Expected argument of type multirunner.RunGoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_multirunner_RunGoRequest(buffer_arg) {
  return gorunner_pb.RunGoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_multirunner_RunResponse(arg) {
  if (!(arg instanceof gorunner_pb.RunResponse)) {
    throw new Error('Expected argument of type multirunner.RunResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_multirunner_RunResponse(buffer_arg) {
  return gorunner_pb.RunResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var GoRunnerService = exports.GoRunnerService = {
  runGo: {
    path: '/multirunner.GoRunner/RunGo',
    requestStream: false,
    responseStream: false,
    requestType: gorunner_pb.RunGoRequest,
    responseType: gorunner_pb.RunResponse,
    requestSerialize: serialize_multirunner_RunGoRequest,
    requestDeserialize: deserialize_multirunner_RunGoRequest,
    responseSerialize: serialize_multirunner_RunResponse,
    responseDeserialize: deserialize_multirunner_RunResponse,
  },
};

exports.GoRunnerClient = grpc.makeGenericClientConstructor(GoRunnerService);
