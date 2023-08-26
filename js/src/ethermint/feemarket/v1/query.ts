/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import { Params } from "../../../ethermint/feemarket/v1/feemarket";
import { BrowserHeaders } from "browser-headers";

export const protobufPackage = "ethermint.feemarket.v1";

/** QueryParamsRequest defines the request type for querying x/evm parameters. */
export interface QueryParamsRequest {}

/** QueryParamsResponse defines the response type for querying x/evm parameters. */
export interface QueryParamsResponse {
  /** params define the evm module parameters. */
  params?: Params;
}

/**
 * QueryBaseFeeRequest defines the request type for querying the EIP1559 base
 * fee.
 */
export interface QueryBaseFeeRequest {}

/** BaseFeeResponse returns the EIP1559 base fee. */
export interface QueryBaseFeeResponse {
  baseFee: string;
}

/**
 * QueryBlockGasRequest defines the request type for querying the EIP1559 base
 * fee.
 */
export interface QueryBlockGasRequest {}

/** QueryBlockGasResponse returns block gas used for a given height. */
export interface QueryBlockGasResponse {
  gas: Long;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryBaseFeeRequest: object = {};

export const QueryBaseFeeRequest = {
  encode(_: QueryBaseFeeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBaseFeeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBaseFeeRequest } as QueryBaseFeeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryBaseFeeRequest {
    const message = { ...baseQueryBaseFeeRequest } as QueryBaseFeeRequest;
    return message;
  },

  toJSON(_: QueryBaseFeeRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryBaseFeeRequest>): QueryBaseFeeRequest {
    const message = { ...baseQueryBaseFeeRequest } as QueryBaseFeeRequest;
    return message;
  },
};

const baseQueryBaseFeeResponse: object = { baseFee: "" };

export const QueryBaseFeeResponse = {
  encode(message: QueryBaseFeeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseFee !== "") {
      writer.uint32(10).string(message.baseFee);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBaseFeeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBaseFeeResponse } as QueryBaseFeeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.baseFee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBaseFeeResponse {
    const message = { ...baseQueryBaseFeeResponse } as QueryBaseFeeResponse;
    if (object.baseFee !== undefined && object.baseFee !== null) {
      message.baseFee = String(object.baseFee);
    } else {
      message.baseFee = "";
    }
    return message;
  },

  toJSON(message: QueryBaseFeeResponse): unknown {
    const obj: any = {};
    message.baseFee !== undefined && (obj.baseFee = message.baseFee);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBaseFeeResponse>): QueryBaseFeeResponse {
    const message = { ...baseQueryBaseFeeResponse } as QueryBaseFeeResponse;
    if (object.baseFee !== undefined && object.baseFee !== null) {
      message.baseFee = object.baseFee;
    } else {
      message.baseFee = "";
    }
    return message;
  },
};

const baseQueryBlockGasRequest: object = {};

export const QueryBlockGasRequest = {
  encode(_: QueryBlockGasRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBlockGasRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBlockGasRequest } as QueryBlockGasRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryBlockGasRequest {
    const message = { ...baseQueryBlockGasRequest } as QueryBlockGasRequest;
    return message;
  },

  toJSON(_: QueryBlockGasRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryBlockGasRequest>): QueryBlockGasRequest {
    const message = { ...baseQueryBlockGasRequest } as QueryBlockGasRequest;
    return message;
  },
};

const baseQueryBlockGasResponse: object = { gas: Long.ZERO };

export const QueryBlockGasResponse = {
  encode(message: QueryBlockGasResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.gas.isZero()) {
      writer.uint32(8).int64(message.gas);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryBlockGasResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBlockGasResponse } as QueryBlockGasResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gas = reader.int64() as Long;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBlockGasResponse {
    const message = { ...baseQueryBlockGasResponse } as QueryBlockGasResponse;
    if (object.gas !== undefined && object.gas !== null) {
      message.gas = Long.fromString(object.gas);
    } else {
      message.gas = Long.ZERO;
    }
    return message;
  },

  toJSON(message: QueryBlockGasResponse): unknown {
    const obj: any = {};
    message.gas !== undefined && (obj.gas = (message.gas || Long.ZERO).toString());
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBlockGasResponse>): QueryBlockGasResponse {
    const message = { ...baseQueryBlockGasResponse } as QueryBlockGasResponse;
    if (object.gas !== undefined && object.gas !== null) {
      message.gas = object.gas as Long;
    } else {
      message.gas = Long.ZERO;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Params queries the parameters of x/feemarket module. */
  Params(request: DeepPartial<QueryParamsRequest>, metadata?: grpc.Metadata): Promise<QueryParamsResponse>;
  /** BaseFee queries the base fee of the parent block of the current block. */
  BaseFee(request: DeepPartial<QueryBaseFeeRequest>, metadata?: grpc.Metadata): Promise<QueryBaseFeeResponse>;
  /** BlockGas queries the gas used at a given block height */
  BlockGas(
    request: DeepPartial<QueryBlockGasRequest>,
    metadata?: grpc.Metadata,
  ): Promise<QueryBlockGasResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.BaseFee = this.BaseFee.bind(this);
    this.BlockGas = this.BlockGas.bind(this);
  }

  Params(request: DeepPartial<QueryParamsRequest>, metadata?: grpc.Metadata): Promise<QueryParamsResponse> {
    return this.rpc.unary(QueryParamsDesc, QueryParamsRequest.fromPartial(request), metadata);
  }

  BaseFee(
    request: DeepPartial<QueryBaseFeeRequest>,
    metadata?: grpc.Metadata,
  ): Promise<QueryBaseFeeResponse> {
    return this.rpc.unary(QueryBaseFeeDesc, QueryBaseFeeRequest.fromPartial(request), metadata);
  }

  BlockGas(
    request: DeepPartial<QueryBlockGasRequest>,
    metadata?: grpc.Metadata,
  ): Promise<QueryBlockGasResponse> {
    return this.rpc.unary(QueryBlockGasDesc, QueryBlockGasRequest.fromPartial(request), metadata);
  }
}

export const QueryDesc = {
  serviceName: "ethermint.feemarket.v1.Query",
};

export const QueryParamsDesc: UnaryMethodDefinitionish = {
  methodName: "Params",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QueryParamsRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...QueryParamsResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QueryBaseFeeDesc: UnaryMethodDefinitionish = {
  methodName: "BaseFee",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QueryBaseFeeRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...QueryBaseFeeResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QueryBlockGasDesc: UnaryMethodDefinitionish = {
  methodName: "BlockGas",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return QueryBlockGasRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...QueryBlockGasResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

interface UnaryMethodDefinitionishR extends grpc.UnaryMethodDefinition<any, any> {
  requestStream: any;
  responseStream: any;
}

type UnaryMethodDefinitionish = UnaryMethodDefinitionishR;

interface Rpc {
  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    request: any,
    metadata: grpc.Metadata | undefined,
  ): Promise<any>;
}

export class GrpcWebImpl {
  private host: string;
  private options: {
    transport?: grpc.TransportFactory;

    debug?: boolean;
    metadata?: grpc.Metadata;
  };

  constructor(
    host: string,
    options: {
      transport?: grpc.TransportFactory;

      debug?: boolean;
      metadata?: grpc.Metadata;
    },
  ) {
    this.host = host;
    this.options = options;
  }

  unary<T extends UnaryMethodDefinitionish>(
    methodDesc: T,
    _request: any,
    metadata: grpc.Metadata | undefined,
  ): Promise<any> {
    const request = { ..._request, ...methodDesc.requestType };
    const maybeCombinedMetadata =
      metadata && this.options.metadata
        ? new BrowserHeaders({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
        : metadata || this.options.metadata;
    return new Promise((resolve, reject) => {
      grpc.unary(methodDesc, {
        request,
        host: this.host,
        metadata: maybeCombinedMetadata,
        transport: this.options.transport,
        debug: this.options.debug,
        onEnd: function (response) {
          if (response.status === grpc.Code.OK) {
            resolve(response.message);
          } else {
            const err = new Error(response.statusMessage) as any;
            err.code = response.status;
            err.metadata = response.trailers;
            reject(err);
          }
        },
      });
    });
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined | Long;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}
