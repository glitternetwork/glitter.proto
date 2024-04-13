/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import { BrowserHeaders } from "browser-headers";
import { Argument, ResultSet } from "../index/sql_engine";

export const protobufPackage = "blockved.glitterchain.index";

export interface MsgSchema {
  ownerAddress: string;
  schemaName: string;
  body: string;
}

export interface MsgSchemaResponse {}

export interface MsgDoc {
  ownerAddress: string;
  schemaName: string;
  body: string;
}

export interface MsgDocResponse {}

export interface SQLExecRequest {
  uid: string;
  sql: string;
  arguments: Argument[];
}

export interface SQLExecResponse {
  results: ResultSet[];
  tookTimes: number;
}

export interface SQLGrantRequest {
  uid: string;
  onTable: string;
  toUID: string;
  role: string;
  onDatabase: string;
}

export interface SQLGrantResponse {}

export interface SQLAnalyzerRequest {
  uid: string;
  opType: string;
  isDelete: boolean;
  tokenFilter?: TokenFilter;
  tokenizer?: Tokenizer;
  analyzer?: Analyzer;
}

export interface TokenFilter {
  tokenFilterName: string;
  tokenFilterEnName: string;
  tokenFilterType: Long;
  dictFileCid: string;
  comment: string;
}

export interface Tokenizer {
  tokenizerName: string;
  tokenizerEnName: string;
  tokenizerType: Long;
  dictFileCid: string;
  comment: string;
}

export interface Analyzer {
  analyzerName: string;
  analyzerEnName: string;
  charFilters: string;
  tokenizer: string;
  tokenFilters: string;
  comment: string;
}

export interface SQLAnalyzerResponse {}

const baseMsgSchema: object = { ownerAddress: "", schemaName: "", body: "" };

export const MsgSchema = {
  encode(message: MsgSchema, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSchema {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSchema } as MsgSchema;
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

  fromJSON(object: any): MsgSchema {
    const message = { ...baseMsgSchema } as MsgSchema;
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

  toJSON(message: MsgSchema): unknown {
    const obj: any = {};
    message.ownerAddress !== undefined && (obj.ownerAddress = message.ownerAddress);
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    message.body !== undefined && (obj.body = message.body);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSchema>): MsgSchema {
    const message = { ...baseMsgSchema } as MsgSchema;
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

const baseMsgSchemaResponse: object = {};

export const MsgSchemaResponse = {
  encode(_: MsgSchemaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSchemaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSchemaResponse } as MsgSchemaResponse;
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

  fromJSON(_: any): MsgSchemaResponse {
    const message = { ...baseMsgSchemaResponse } as MsgSchemaResponse;
    return message;
  },

  toJSON(_: MsgSchemaResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSchemaResponse>): MsgSchemaResponse {
    const message = { ...baseMsgSchemaResponse } as MsgSchemaResponse;
    return message;
  },
};

const baseMsgDoc: object = { ownerAddress: "", schemaName: "", body: "" };

export const MsgDoc = {
  encode(message: MsgDoc, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDoc {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDoc } as MsgDoc;
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

  fromJSON(object: any): MsgDoc {
    const message = { ...baseMsgDoc } as MsgDoc;
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

  toJSON(message: MsgDoc): unknown {
    const obj: any = {};
    message.ownerAddress !== undefined && (obj.ownerAddress = message.ownerAddress);
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    message.body !== undefined && (obj.body = message.body);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDoc>): MsgDoc {
    const message = { ...baseMsgDoc } as MsgDoc;
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

const baseMsgDocResponse: object = {};

export const MsgDocResponse = {
  encode(_: MsgDocResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDocResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDocResponse } as MsgDocResponse;
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

  fromJSON(_: any): MsgDocResponse {
    const message = { ...baseMsgDocResponse } as MsgDocResponse;
    return message;
  },

  toJSON(_: MsgDocResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDocResponse>): MsgDocResponse {
    const message = { ...baseMsgDocResponse } as MsgDocResponse;
    return message;
  },
};

const baseSQLExecRequest: object = { uid: "", sql: "" };

export const SQLExecRequest = {
  encode(message: SQLExecRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.sql !== "") {
      writer.uint32(18).string(message.sql);
    }
    for (const v of message.arguments) {
      Argument.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLExecRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLExecRequest } as SQLExecRequest;
    message.arguments = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.sql = reader.string();
          break;
        case 3:
          message.arguments.push(Argument.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLExecRequest {
    const message = { ...baseSQLExecRequest } as SQLExecRequest;
    message.arguments = [];
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.sql !== undefined && object.sql !== null) {
      message.sql = String(object.sql);
    } else {
      message.sql = "";
    }
    if (object.arguments !== undefined && object.arguments !== null) {
      for (const e of object.arguments) {
        message.arguments.push(Argument.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: SQLExecRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.sql !== undefined && (obj.sql = message.sql);
    if (message.arguments) {
      obj.arguments = message.arguments.map((e) => (e ? Argument.toJSON(e) : undefined));
    } else {
      obj.arguments = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SQLExecRequest>): SQLExecRequest {
    const message = { ...baseSQLExecRequest } as SQLExecRequest;
    message.arguments = [];
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.sql !== undefined && object.sql !== null) {
      message.sql = object.sql;
    } else {
      message.sql = "";
    }
    if (object.arguments !== undefined && object.arguments !== null) {
      for (const e of object.arguments) {
        message.arguments.push(Argument.fromPartial(e));
      }
    }
    return message;
  },
};

const baseSQLExecResponse: object = { tookTimes: 0 };

export const SQLExecResponse = {
  encode(message: SQLExecResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.results) {
      ResultSet.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.tookTimes !== 0) {
      writer.uint32(21).float(message.tookTimes);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLExecResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLExecResponse } as SQLExecResponse;
    message.results = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.results.push(ResultSet.decode(reader, reader.uint32()));
          break;
        case 2:
          message.tookTimes = reader.float();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLExecResponse {
    const message = { ...baseSQLExecResponse } as SQLExecResponse;
    message.results = [];
    if (object.results !== undefined && object.results !== null) {
      for (const e of object.results) {
        message.results.push(ResultSet.fromJSON(e));
      }
    }
    if (object.tookTimes !== undefined && object.tookTimes !== null) {
      message.tookTimes = Number(object.tookTimes);
    } else {
      message.tookTimes = 0;
    }
    return message;
  },

  toJSON(message: SQLExecResponse): unknown {
    const obj: any = {};
    if (message.results) {
      obj.results = message.results.map((e) => (e ? ResultSet.toJSON(e) : undefined));
    } else {
      obj.results = [];
    }
    message.tookTimes !== undefined && (obj.tookTimes = message.tookTimes);
    return obj;
  },

  fromPartial(object: DeepPartial<SQLExecResponse>): SQLExecResponse {
    const message = { ...baseSQLExecResponse } as SQLExecResponse;
    message.results = [];
    if (object.results !== undefined && object.results !== null) {
      for (const e of object.results) {
        message.results.push(ResultSet.fromPartial(e));
      }
    }
    if (object.tookTimes !== undefined && object.tookTimes !== null) {
      message.tookTimes = object.tookTimes;
    } else {
      message.tookTimes = 0;
    }
    return message;
  },
};

const baseSQLGrantRequest: object = { uid: "", onTable: "", toUID: "", role: "", onDatabase: "" };

export const SQLGrantRequest = {
  encode(message: SQLGrantRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.onTable !== "") {
      writer.uint32(18).string(message.onTable);
    }
    if (message.toUID !== "") {
      writer.uint32(26).string(message.toUID);
    }
    if (message.role !== "") {
      writer.uint32(34).string(message.role);
    }
    if (message.onDatabase !== "") {
      writer.uint32(42).string(message.onDatabase);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLGrantRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLGrantRequest } as SQLGrantRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.onTable = reader.string();
          break;
        case 3:
          message.toUID = reader.string();
          break;
        case 4:
          message.role = reader.string();
          break;
        case 5:
          message.onDatabase = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLGrantRequest {
    const message = { ...baseSQLGrantRequest } as SQLGrantRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.onTable !== undefined && object.onTable !== null) {
      message.onTable = String(object.onTable);
    } else {
      message.onTable = "";
    }
    if (object.toUID !== undefined && object.toUID !== null) {
      message.toUID = String(object.toUID);
    } else {
      message.toUID = "";
    }
    if (object.role !== undefined && object.role !== null) {
      message.role = String(object.role);
    } else {
      message.role = "";
    }
    if (object.onDatabase !== undefined && object.onDatabase !== null) {
      message.onDatabase = String(object.onDatabase);
    } else {
      message.onDatabase = "";
    }
    return message;
  },

  toJSON(message: SQLGrantRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.onTable !== undefined && (obj.onTable = message.onTable);
    message.toUID !== undefined && (obj.toUID = message.toUID);
    message.role !== undefined && (obj.role = message.role);
    message.onDatabase !== undefined && (obj.onDatabase = message.onDatabase);
    return obj;
  },

  fromPartial(object: DeepPartial<SQLGrantRequest>): SQLGrantRequest {
    const message = { ...baseSQLGrantRequest } as SQLGrantRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.onTable !== undefined && object.onTable !== null) {
      message.onTable = object.onTable;
    } else {
      message.onTable = "";
    }
    if (object.toUID !== undefined && object.toUID !== null) {
      message.toUID = object.toUID;
    } else {
      message.toUID = "";
    }
    if (object.role !== undefined && object.role !== null) {
      message.role = object.role;
    } else {
      message.role = "";
    }
    if (object.onDatabase !== undefined && object.onDatabase !== null) {
      message.onDatabase = object.onDatabase;
    } else {
      message.onDatabase = "";
    }
    return message;
  },
};

const baseSQLGrantResponse: object = {};

export const SQLGrantResponse = {
  encode(_: SQLGrantResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLGrantResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLGrantResponse } as SQLGrantResponse;
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

  fromJSON(_: any): SQLGrantResponse {
    const message = { ...baseSQLGrantResponse } as SQLGrantResponse;
    return message;
  },

  toJSON(_: SQLGrantResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<SQLGrantResponse>): SQLGrantResponse {
    const message = { ...baseSQLGrantResponse } as SQLGrantResponse;
    return message;
  },
};

const baseSQLAnalyzerRequest: object = { uid: "", opType: "", isDelete: false };

export const SQLAnalyzerRequest = {
  encode(message: SQLAnalyzerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.opType !== "") {
      writer.uint32(18).string(message.opType);
    }
    if (message.isDelete === true) {
      writer.uint32(24).bool(message.isDelete);
    }
    if (message.tokenFilter !== undefined) {
      TokenFilter.encode(message.tokenFilter, writer.uint32(34).fork()).ldelim();
    }
    if (message.tokenizer !== undefined) {
      Tokenizer.encode(message.tokenizer, writer.uint32(42).fork()).ldelim();
    }
    if (message.analyzer !== undefined) {
      Analyzer.encode(message.analyzer, writer.uint32(50).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLAnalyzerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLAnalyzerRequest } as SQLAnalyzerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.opType = reader.string();
          break;
        case 3:
          message.isDelete = reader.bool();
          break;
        case 4:
          message.tokenFilter = TokenFilter.decode(reader, reader.uint32());
          break;
        case 5:
          message.tokenizer = Tokenizer.decode(reader, reader.uint32());
          break;
        case 6:
          message.analyzer = Analyzer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLAnalyzerRequest {
    const message = { ...baseSQLAnalyzerRequest } as SQLAnalyzerRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.opType !== undefined && object.opType !== null) {
      message.opType = String(object.opType);
    } else {
      message.opType = "";
    }
    if (object.isDelete !== undefined && object.isDelete !== null) {
      message.isDelete = Boolean(object.isDelete);
    } else {
      message.isDelete = false;
    }
    if (object.tokenFilter !== undefined && object.tokenFilter !== null) {
      message.tokenFilter = TokenFilter.fromJSON(object.tokenFilter);
    } else {
      message.tokenFilter = undefined;
    }
    if (object.tokenizer !== undefined && object.tokenizer !== null) {
      message.tokenizer = Tokenizer.fromJSON(object.tokenizer);
    } else {
      message.tokenizer = undefined;
    }
    if (object.analyzer !== undefined && object.analyzer !== null) {
      message.analyzer = Analyzer.fromJSON(object.analyzer);
    } else {
      message.analyzer = undefined;
    }
    return message;
  },

  toJSON(message: SQLAnalyzerRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.opType !== undefined && (obj.opType = message.opType);
    message.isDelete !== undefined && (obj.isDelete = message.isDelete);
    message.tokenFilter !== undefined &&
      (obj.tokenFilter = message.tokenFilter ? TokenFilter.toJSON(message.tokenFilter) : undefined);
    message.tokenizer !== undefined &&
      (obj.tokenizer = message.tokenizer ? Tokenizer.toJSON(message.tokenizer) : undefined);
    message.analyzer !== undefined &&
      (obj.analyzer = message.analyzer ? Analyzer.toJSON(message.analyzer) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<SQLAnalyzerRequest>): SQLAnalyzerRequest {
    const message = { ...baseSQLAnalyzerRequest } as SQLAnalyzerRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.opType !== undefined && object.opType !== null) {
      message.opType = object.opType;
    } else {
      message.opType = "";
    }
    if (object.isDelete !== undefined && object.isDelete !== null) {
      message.isDelete = object.isDelete;
    } else {
      message.isDelete = false;
    }
    if (object.tokenFilter !== undefined && object.tokenFilter !== null) {
      message.tokenFilter = TokenFilter.fromPartial(object.tokenFilter);
    } else {
      message.tokenFilter = undefined;
    }
    if (object.tokenizer !== undefined && object.tokenizer !== null) {
      message.tokenizer = Tokenizer.fromPartial(object.tokenizer);
    } else {
      message.tokenizer = undefined;
    }
    if (object.analyzer !== undefined && object.analyzer !== null) {
      message.analyzer = Analyzer.fromPartial(object.analyzer);
    } else {
      message.analyzer = undefined;
    }
    return message;
  },
};

const baseTokenFilter: object = {
  tokenFilterName: "",
  tokenFilterEnName: "",
  tokenFilterType: Long.ZERO,
  dictFileCid: "",
  comment: "",
};

export const TokenFilter = {
  encode(message: TokenFilter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tokenFilterName !== "") {
      writer.uint32(10).string(message.tokenFilterName);
    }
    if (message.tokenFilterEnName !== "") {
      writer.uint32(18).string(message.tokenFilterEnName);
    }
    if (!message.tokenFilterType.isZero()) {
      writer.uint32(24).int64(message.tokenFilterType);
    }
    if (message.dictFileCid !== "") {
      writer.uint32(34).string(message.dictFileCid);
    }
    if (message.comment !== "") {
      writer.uint32(42).string(message.comment);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TokenFilter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTokenFilter } as TokenFilter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenFilterName = reader.string();
          break;
        case 2:
          message.tokenFilterEnName = reader.string();
          break;
        case 3:
          message.tokenFilterType = reader.int64() as Long;
          break;
        case 4:
          message.dictFileCid = reader.string();
          break;
        case 5:
          message.comment = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenFilter {
    const message = { ...baseTokenFilter } as TokenFilter;
    if (object.tokenFilterName !== undefined && object.tokenFilterName !== null) {
      message.tokenFilterName = String(object.tokenFilterName);
    } else {
      message.tokenFilterName = "";
    }
    if (object.tokenFilterEnName !== undefined && object.tokenFilterEnName !== null) {
      message.tokenFilterEnName = String(object.tokenFilterEnName);
    } else {
      message.tokenFilterEnName = "";
    }
    if (object.tokenFilterType !== undefined && object.tokenFilterType !== null) {
      message.tokenFilterType = Long.fromString(object.tokenFilterType);
    } else {
      message.tokenFilterType = Long.ZERO;
    }
    if (object.dictFileCid !== undefined && object.dictFileCid !== null) {
      message.dictFileCid = String(object.dictFileCid);
    } else {
      message.dictFileCid = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = String(object.comment);
    } else {
      message.comment = "";
    }
    return message;
  },

  toJSON(message: TokenFilter): unknown {
    const obj: any = {};
    message.tokenFilterName !== undefined && (obj.tokenFilterName = message.tokenFilterName);
    message.tokenFilterEnName !== undefined && (obj.tokenFilterEnName = message.tokenFilterEnName);
    message.tokenFilterType !== undefined &&
      (obj.tokenFilterType = (message.tokenFilterType || Long.ZERO).toString());
    message.dictFileCid !== undefined && (obj.dictFileCid = message.dictFileCid);
    message.comment !== undefined && (obj.comment = message.comment);
    return obj;
  },

  fromPartial(object: DeepPartial<TokenFilter>): TokenFilter {
    const message = { ...baseTokenFilter } as TokenFilter;
    if (object.tokenFilterName !== undefined && object.tokenFilterName !== null) {
      message.tokenFilterName = object.tokenFilterName;
    } else {
      message.tokenFilterName = "";
    }
    if (object.tokenFilterEnName !== undefined && object.tokenFilterEnName !== null) {
      message.tokenFilterEnName = object.tokenFilterEnName;
    } else {
      message.tokenFilterEnName = "";
    }
    if (object.tokenFilterType !== undefined && object.tokenFilterType !== null) {
      message.tokenFilterType = object.tokenFilterType as Long;
    } else {
      message.tokenFilterType = Long.ZERO;
    }
    if (object.dictFileCid !== undefined && object.dictFileCid !== null) {
      message.dictFileCid = object.dictFileCid;
    } else {
      message.dictFileCid = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = object.comment;
    } else {
      message.comment = "";
    }
    return message;
  },
};

const baseTokenizer: object = {
  tokenizerName: "",
  tokenizerEnName: "",
  tokenizerType: Long.ZERO,
  dictFileCid: "",
  comment: "",
};

export const Tokenizer = {
  encode(message: Tokenizer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tokenizerName !== "") {
      writer.uint32(10).string(message.tokenizerName);
    }
    if (message.tokenizerEnName !== "") {
      writer.uint32(18).string(message.tokenizerEnName);
    }
    if (!message.tokenizerType.isZero()) {
      writer.uint32(24).int64(message.tokenizerType);
    }
    if (message.dictFileCid !== "") {
      writer.uint32(34).string(message.dictFileCid);
    }
    if (message.comment !== "") {
      writer.uint32(42).string(message.comment);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Tokenizer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTokenizer } as Tokenizer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenizerName = reader.string();
          break;
        case 2:
          message.tokenizerEnName = reader.string();
          break;
        case 3:
          message.tokenizerType = reader.int64() as Long;
          break;
        case 4:
          message.dictFileCid = reader.string();
          break;
        case 5:
          message.comment = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Tokenizer {
    const message = { ...baseTokenizer } as Tokenizer;
    if (object.tokenizerName !== undefined && object.tokenizerName !== null) {
      message.tokenizerName = String(object.tokenizerName);
    } else {
      message.tokenizerName = "";
    }
    if (object.tokenizerEnName !== undefined && object.tokenizerEnName !== null) {
      message.tokenizerEnName = String(object.tokenizerEnName);
    } else {
      message.tokenizerEnName = "";
    }
    if (object.tokenizerType !== undefined && object.tokenizerType !== null) {
      message.tokenizerType = Long.fromString(object.tokenizerType);
    } else {
      message.tokenizerType = Long.ZERO;
    }
    if (object.dictFileCid !== undefined && object.dictFileCid !== null) {
      message.dictFileCid = String(object.dictFileCid);
    } else {
      message.dictFileCid = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = String(object.comment);
    } else {
      message.comment = "";
    }
    return message;
  },

  toJSON(message: Tokenizer): unknown {
    const obj: any = {};
    message.tokenizerName !== undefined && (obj.tokenizerName = message.tokenizerName);
    message.tokenizerEnName !== undefined && (obj.tokenizerEnName = message.tokenizerEnName);
    message.tokenizerType !== undefined &&
      (obj.tokenizerType = (message.tokenizerType || Long.ZERO).toString());
    message.dictFileCid !== undefined && (obj.dictFileCid = message.dictFileCid);
    message.comment !== undefined && (obj.comment = message.comment);
    return obj;
  },

  fromPartial(object: DeepPartial<Tokenizer>): Tokenizer {
    const message = { ...baseTokenizer } as Tokenizer;
    if (object.tokenizerName !== undefined && object.tokenizerName !== null) {
      message.tokenizerName = object.tokenizerName;
    } else {
      message.tokenizerName = "";
    }
    if (object.tokenizerEnName !== undefined && object.tokenizerEnName !== null) {
      message.tokenizerEnName = object.tokenizerEnName;
    } else {
      message.tokenizerEnName = "";
    }
    if (object.tokenizerType !== undefined && object.tokenizerType !== null) {
      message.tokenizerType = object.tokenizerType as Long;
    } else {
      message.tokenizerType = Long.ZERO;
    }
    if (object.dictFileCid !== undefined && object.dictFileCid !== null) {
      message.dictFileCid = object.dictFileCid;
    } else {
      message.dictFileCid = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = object.comment;
    } else {
      message.comment = "";
    }
    return message;
  },
};

const baseAnalyzer: object = {
  analyzerName: "",
  analyzerEnName: "",
  charFilters: "",
  tokenizer: "",
  tokenFilters: "",
  comment: "",
};

export const Analyzer = {
  encode(message: Analyzer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.analyzerName !== "") {
      writer.uint32(10).string(message.analyzerName);
    }
    if (message.analyzerEnName !== "") {
      writer.uint32(18).string(message.analyzerEnName);
    }
    if (message.charFilters !== "") {
      writer.uint32(26).string(message.charFilters);
    }
    if (message.tokenizer !== "") {
      writer.uint32(34).string(message.tokenizer);
    }
    if (message.tokenFilters !== "") {
      writer.uint32(42).string(message.tokenFilters);
    }
    if (message.comment !== "") {
      writer.uint32(50).string(message.comment);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Analyzer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAnalyzer } as Analyzer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.analyzerName = reader.string();
          break;
        case 2:
          message.analyzerEnName = reader.string();
          break;
        case 3:
          message.charFilters = reader.string();
          break;
        case 4:
          message.tokenizer = reader.string();
          break;
        case 5:
          message.tokenFilters = reader.string();
          break;
        case 6:
          message.comment = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Analyzer {
    const message = { ...baseAnalyzer } as Analyzer;
    if (object.analyzerName !== undefined && object.analyzerName !== null) {
      message.analyzerName = String(object.analyzerName);
    } else {
      message.analyzerName = "";
    }
    if (object.analyzerEnName !== undefined && object.analyzerEnName !== null) {
      message.analyzerEnName = String(object.analyzerEnName);
    } else {
      message.analyzerEnName = "";
    }
    if (object.charFilters !== undefined && object.charFilters !== null) {
      message.charFilters = String(object.charFilters);
    } else {
      message.charFilters = "";
    }
    if (object.tokenizer !== undefined && object.tokenizer !== null) {
      message.tokenizer = String(object.tokenizer);
    } else {
      message.tokenizer = "";
    }
    if (object.tokenFilters !== undefined && object.tokenFilters !== null) {
      message.tokenFilters = String(object.tokenFilters);
    } else {
      message.tokenFilters = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = String(object.comment);
    } else {
      message.comment = "";
    }
    return message;
  },

  toJSON(message: Analyzer): unknown {
    const obj: any = {};
    message.analyzerName !== undefined && (obj.analyzerName = message.analyzerName);
    message.analyzerEnName !== undefined && (obj.analyzerEnName = message.analyzerEnName);
    message.charFilters !== undefined && (obj.charFilters = message.charFilters);
    message.tokenizer !== undefined && (obj.tokenizer = message.tokenizer);
    message.tokenFilters !== undefined && (obj.tokenFilters = message.tokenFilters);
    message.comment !== undefined && (obj.comment = message.comment);
    return obj;
  },

  fromPartial(object: DeepPartial<Analyzer>): Analyzer {
    const message = { ...baseAnalyzer } as Analyzer;
    if (object.analyzerName !== undefined && object.analyzerName !== null) {
      message.analyzerName = object.analyzerName;
    } else {
      message.analyzerName = "";
    }
    if (object.analyzerEnName !== undefined && object.analyzerEnName !== null) {
      message.analyzerEnName = object.analyzerEnName;
    } else {
      message.analyzerEnName = "";
    }
    if (object.charFilters !== undefined && object.charFilters !== null) {
      message.charFilters = object.charFilters;
    } else {
      message.charFilters = "";
    }
    if (object.tokenizer !== undefined && object.tokenizer !== null) {
      message.tokenizer = object.tokenizer;
    } else {
      message.tokenizer = "";
    }
    if (object.tokenFilters !== undefined && object.tokenFilters !== null) {
      message.tokenFilters = object.tokenFilters;
    } else {
      message.tokenFilters = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = object.comment;
    } else {
      message.comment = "";
    }
    return message;
  },
};

const baseSQLAnalyzerResponse: object = {};

export const SQLAnalyzerResponse = {
  encode(_: SQLAnalyzerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLAnalyzerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLAnalyzerResponse } as SQLAnalyzerResponse;
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

  fromJSON(_: any): SQLAnalyzerResponse {
    const message = { ...baseSQLAnalyzerResponse } as SQLAnalyzerResponse;
    return message;
  },

  toJSON(_: SQLAnalyzerResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<SQLAnalyzerResponse>): SQLAnalyzerResponse {
    const message = { ...baseSQLAnalyzerResponse } as SQLAnalyzerResponse;
    return message;
  },
};

export interface Msg {
  ManageSchema(request: DeepPartial<MsgSchema>, metadata?: grpc.Metadata): Promise<MsgSchemaResponse>;
  ManageDoc(request: DeepPartial<MsgDoc>, metadata?: grpc.Metadata): Promise<MsgDocResponse>;
  SQLExec(request: DeepPartial<SQLExecRequest>, metadata?: grpc.Metadata): Promise<SQLExecResponse>;
  SQLGrant(request: DeepPartial<SQLGrantRequest>, metadata?: grpc.Metadata): Promise<SQLGrantResponse>;
  SQLAnalyzer(
    request: DeepPartial<SQLAnalyzerRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SQLAnalyzerResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ManageSchema = this.ManageSchema.bind(this);
    this.ManageDoc = this.ManageDoc.bind(this);
    this.SQLExec = this.SQLExec.bind(this);
    this.SQLGrant = this.SQLGrant.bind(this);
    this.SQLAnalyzer = this.SQLAnalyzer.bind(this);
  }

  ManageSchema(request: DeepPartial<MsgSchema>, metadata?: grpc.Metadata): Promise<MsgSchemaResponse> {
    return this.rpc.unary(MsgManageSchemaDesc, MsgSchema.fromPartial(request), metadata);
  }

  ManageDoc(request: DeepPartial<MsgDoc>, metadata?: grpc.Metadata): Promise<MsgDocResponse> {
    return this.rpc.unary(MsgManageDocDesc, MsgDoc.fromPartial(request), metadata);
  }

  SQLExec(request: DeepPartial<SQLExecRequest>, metadata?: grpc.Metadata): Promise<SQLExecResponse> {
    return this.rpc.unary(MsgSQLExecDesc, SQLExecRequest.fromPartial(request), metadata);
  }

  SQLGrant(request: DeepPartial<SQLGrantRequest>, metadata?: grpc.Metadata): Promise<SQLGrantResponse> {
    return this.rpc.unary(MsgSQLGrantDesc, SQLGrantRequest.fromPartial(request), metadata);
  }

  SQLAnalyzer(
    request: DeepPartial<SQLAnalyzerRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SQLAnalyzerResponse> {
    return this.rpc.unary(MsgSQLAnalyzerDesc, SQLAnalyzerRequest.fromPartial(request), metadata);
  }
}

export const MsgDesc = {
  serviceName: "blockved.glitterchain.index.Msg",
};

export const MsgManageSchemaDesc: UnaryMethodDefinitionish = {
  methodName: "ManageSchema",
  service: MsgDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return MsgSchema.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...MsgSchemaResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const MsgManageDocDesc: UnaryMethodDefinitionish = {
  methodName: "ManageDoc",
  service: MsgDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return MsgDoc.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...MsgDocResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const MsgSQLExecDesc: UnaryMethodDefinitionish = {
  methodName: "SQLExec",
  service: MsgDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SQLExecRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SQLExecResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const MsgSQLGrantDesc: UnaryMethodDefinitionish = {
  methodName: "SQLGrant",
  service: MsgDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SQLGrantRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SQLGrantResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const MsgSQLAnalyzerDesc: UnaryMethodDefinitionish = {
  methodName: "SQLAnalyzer",
  service: MsgDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SQLAnalyzerRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SQLAnalyzerResponse.decode(data),
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
