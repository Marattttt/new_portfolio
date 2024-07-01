// package: jsrunner
// file: jsrunner.proto

import * as jspb from 'google-protobuf';

export class JsRequest extends jspb.Message {
  getCode(): string;
  setCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: JsRequest): JsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JsRequest;
  static deserializeBinaryFromReader(message: JsRequest, reader: jspb.BinaryReader): JsRequest;
}

export namespace JsRequest {
  export type AsObject = {
    code: string,
  }
}

export class RunResponse extends jspb.Message {
  getError(): string;
  setError(value: string): void;

  getOutput(): string;
  setOutput(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RunResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RunResponse): RunResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RunResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RunResponse;
  static deserializeBinaryFromReader(message: RunResponse, reader: jspb.BinaryReader): RunResponse;
}

export namespace RunResponse {
  export type AsObject = {
    error: string,
    output: string,
  }
}

