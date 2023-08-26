/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "blockved.glitterchain.index";

export interface Schema {
  ownerAddress: string;
  schemaName: string;
  body: string;
}

export interface Doc {
  ownerAddress: string;
  schemaName: string;
  body: string;
}

const baseSchema: object = { ownerAddress: "", schemaName: "", body: "" };

export const Schema = {
  encode(message: Schema, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ownerAddress !== "") {
      writer.uint32(10).string(message.ownerAddress);
    }
    if (message.schemaName !== "") {
      writer.uint32(18).string(message.schemaName);
    }
    if (message.body !== "") {
      writer.uint32(26).string(message.body);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Schema {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSchema } as Schema;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ownerAddress = reader.string();
          break;
        case 2:
          message.schemaName = reader.string();
          break;
        case 3:
          message.body = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Schema {
    const message = { ...baseSchema } as Schema;
    if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
      message.ownerAddress = String(object.ownerAddress);
    } else {
      message.ownerAddress = "";
    }
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = String(object.schemaName);
    } else {
      message.schemaName = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = String(object.body);
    } else {
      message.body = "";
    }
    return message;
  },

  toJSON(message: Schema): unknown {
    const obj: any = {};
    message.ownerAddress !== undefined && (obj.ownerAddress = message.ownerAddress);
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    message.body !== undefined && (obj.body = message.body);
    return obj;
  },

  fromPartial(object: DeepPartial<Schema>): Schema {
    const message = { ...baseSchema } as Schema;
    if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
      message.ownerAddress = object.ownerAddress;
    } else {
      message.ownerAddress = "";
    }
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = object.schemaName;
    } else {
      message.schemaName = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = object.body;
    } else {
      message.body = "";
    }
    return message;
  },
};

const baseDoc: object = { ownerAddress: "", schemaName: "", body: "" };

export const Doc = {
  encode(message: Doc, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.ownerAddress !== "") {
      writer.uint32(10).string(message.ownerAddress);
    }
    if (message.schemaName !== "") {
      writer.uint32(18).string(message.schemaName);
    }
    if (message.body !== "") {
      writer.uint32(26).string(message.body);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Doc {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDoc } as Doc;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ownerAddress = reader.string();
          break;
        case 2:
          message.schemaName = reader.string();
          break;
        case 3:
          message.body = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Doc {
    const message = { ...baseDoc } as Doc;
    if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
      message.ownerAddress = String(object.ownerAddress);
    } else {
      message.ownerAddress = "";
    }
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = String(object.schemaName);
    } else {
      message.schemaName = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = String(object.body);
    } else {
      message.body = "";
    }
    return message;
  },

  toJSON(message: Doc): unknown {
    const obj: any = {};
    message.ownerAddress !== undefined && (obj.ownerAddress = message.ownerAddress);
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    message.body !== undefined && (obj.body = message.body);
    return obj;
  },

  fromPartial(object: DeepPartial<Doc>): Doc {
    const message = { ...baseDoc } as Doc;
    if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
      message.ownerAddress = object.ownerAddress;
    } else {
      message.ownerAddress = "";
    }
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = object.schemaName;
    } else {
      message.schemaName = "";
    }
    if (object.body !== undefined && object.body !== null) {
      message.body = object.body;
    } else {
      message.body = "";
    }
    return message;
  },
};

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
