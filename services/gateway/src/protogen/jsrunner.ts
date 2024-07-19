// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.181.0
//   protoc               v3.19.6
// source: jsrunner.proto

/* eslint-disable */
import {
  type CallOptions,
  ChannelCredentials,
  Client,
  type ClientOptions,
  type ClientUnaryCall,
  type handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  type ServiceError,
  type UntypedServiceImplementation,
} from "@grpc/grpc-js";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "jsrunner";

export interface JsRunRequest {
  code: string;
}

export interface JsRunResponse {
  error: string;
  output: string;
}

function createBaseJsRunRequest(): JsRunRequest {
  return { code: "" };
}

export const JsRunRequest = {
  encode(message: JsRunRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.code !== "") {
      writer.uint32(10).string(message.code);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JsRunRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJsRunRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.code = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): JsRunRequest {
    return { code: isSet(object.code) ? globalThis.String(object.code) : "" };
  },

  toJSON(message: JsRunRequest): unknown {
    const obj: any = {};
    if (message.code !== "") {
      obj.code = message.code;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<JsRunRequest>, I>>(base?: I): JsRunRequest {
    return JsRunRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<JsRunRequest>, I>>(object: I): JsRunRequest {
    const message = createBaseJsRunRequest();
    message.code = object.code ?? "";
    return message;
  },
};

function createBaseJsRunResponse(): JsRunResponse {
  return { error: "", output: "" };
}

export const JsRunResponse = {
  encode(message: JsRunResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.error !== "") {
      writer.uint32(10).string(message.error);
    }
    if (message.output !== "") {
      writer.uint32(18).string(message.output);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): JsRunResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJsRunResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.error = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.output = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): JsRunResponse {
    return {
      error: isSet(object.error) ? globalThis.String(object.error) : "",
      output: isSet(object.output) ? globalThis.String(object.output) : "",
    };
  },

  toJSON(message: JsRunResponse): unknown {
    const obj: any = {};
    if (message.error !== "") {
      obj.error = message.error;
    }
    if (message.output !== "") {
      obj.output = message.output;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<JsRunResponse>, I>>(base?: I): JsRunResponse {
    return JsRunResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<JsRunResponse>, I>>(object: I): JsRunResponse {
    const message = createBaseJsRunResponse();
    message.error = object.error ?? "";
    message.output = object.output ?? "";
    return message;
  },
};

export type JsRunnerService = typeof JsRunnerService;
export const JsRunnerService = {
  runJs: {
    path: "/jsrunner.JsRunner/RunJs",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: JsRunRequest) => Buffer.from(JsRunRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => JsRunRequest.decode(value),
    responseSerialize: (value: JsRunResponse) => Buffer.from(JsRunResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => JsRunResponse.decode(value),
  },
} as const;

export interface JsRunnerServer extends UntypedServiceImplementation {
  runJs: handleUnaryCall<JsRunRequest, JsRunResponse>;
}

export interface JsRunnerClient extends Client {
  runJs(
    request: JsRunRequest,
    callback: (error: ServiceError | null, response: JsRunResponse) => void,
  ): ClientUnaryCall;
  runJs(
    request: JsRunRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: JsRunResponse) => void,
  ): ClientUnaryCall;
  runJs(
    request: JsRunRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: JsRunResponse) => void,
  ): ClientUnaryCall;
}

export const JsRunnerClient = makeGenericClientConstructor(JsRunnerService, "jsrunner.JsRunner") as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): JsRunnerClient;
  service: typeof JsRunnerService;
  serviceName: string;
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}