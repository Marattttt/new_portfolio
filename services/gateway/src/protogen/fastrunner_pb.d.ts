// package: fastrunner
// file: fastrunner.proto

import * as jspb from 'google-protobuf';

export class GoRequest extends jspb.Message {
  getCode(): string;
  setCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GoRequest): GoRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GoRequest;
  static deserializeBinaryFromReader(message: GoRequest, reader: jspb.BinaryReader): GoRequest;
}

export namespace GoRequest {
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

