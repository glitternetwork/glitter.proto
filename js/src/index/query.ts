/* eslint-disable */
import Long from "long";
import { grpc } from "@improbable-eng/grpc-web";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { Schema, Doc } from "../index/index";
import { Params } from "../index/params";
import {
  ColumnValueType,
  Argument,
  ResultSet,
  TableInfo,
  DatabaseInfo,
  columnValueTypeFromJSON,
  columnValueTypeToJSON,
} from "../index/sql_engine";
import { BrowserHeaders } from "browser-headers";

export const protobufPackage = "blockved.glitterchain.index";

/** ListSchemaRequest is request type for the Query RPC method. */
export interface ListSchemaRequest {
  pagination?: PageRequest;
}

/** ListSchemaResponse is response type for the Query RPC method. */
export interface ListSchemaResponse {
  schemas: Schema[];
  pagination?: PageResponse;
}

/** GetSchemaRequest is request type for the Query RPC method. */
export interface GetSchemaRequest {
  schemaName: string;
}

/** GetSchemaResponse is response type for the Query RPC method. */
export interface GetSchemaResponse {
  schema?: Schema;
}

/** GetDocsRequest is request type for the Query RPC method. */
export interface GetDocRequest {
  schemaName: string;
  docID: string;
}

/** GetDocsResponse is response type for the Query RPC method. */
export interface GetDocResponse {
  doc?: Doc;
}

/** GetDocsRequest is request type for the Query RPC method. */
export interface SearchDocRequest {
  schemaName: string;
  query: string;
  queryField: string[];
  aggsField: string[];
  filter?: SearchFilters;
  orderBy: string;
  limit: Long;
  page: Long;
  factor?: FieldValueFactor;
}

/** GetDocsResponse is response type for the Query RPC method. */
export interface SearchDocResponse {
  searchTime: Long;
  schemaName: string;
  meta?: Meta;
  items: DocItem[];
  facet: { [key: string]: SearchFilters };
}

export interface SearchDocResponse_FacetEntry {
  key: string;
  value?: SearchFilters;
}

export interface Meta {
  currentPage: Long;
  totalPages: Long;
  totalResults: Long;
  size: Long;
  sortedBy: string;
}

export interface DocItem {
  Highlight: { [key: string]: strings };
  data: string;
  score: number;
}

export interface DocItem_HighlightEntry {
  key: string;
  value?: strings;
}

export interface strings {
  strings: string[];
}

export interface SearchFilters {
  searchFilter: SearchFilter[];
}

export interface SearchFilter {
  type: string;
  field: string;
  value: string;
  from: number;
  to: number;
  docCount: Long;
}

export interface FieldValueFactor {
  field: string;
  factor: number;
  missing: number;
  weight: number;
  modifier: string;
}

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: Params;
}

export interface SQLQueryRequest {
  sql: string;
  arguments: Argument[];
}

export interface SQLQueryResponse {
  results: ResultSet[];
  tookTimes: number;
}

export interface SimpleSQLQueryResponse {
  result: SimpleSQLQueryResponse_ResultSet[];
}

export interface SimpleSQLQueryResponse_ResultSet {
  row: { [key: string]: SimpleSQLQueryResponse_RowValue };
}

export interface SimpleSQLQueryResponse_ResultSet_RowEntry {
  key: string;
  value?: SimpleSQLQueryResponse_RowValue;
}

export interface SimpleSQLQueryResponse_RowValue {
  value: string;
  columnValueType: ColumnValueType;
}

export interface SQLListTablesRequest {
  keyword: string;
  uid: string;
  page: number;
  pageSize: number;
  database: string;
}

export interface SQLListTablesResponse {
  tables: TableInfo[];
}

export interface SQLListDatabasesRequest {}

export interface SQLListDatabasesResponse {
  databases: DatabaseInfo[];
}

export interface ShowCreateTableRequest {
  databaseName: string;
  tableName: string;
}

export interface ShowCreateTableResponse {
  schema: string;
}

const baseListSchemaRequest: object = {};

export const ListSchemaRequest = {
  encode(message: ListSchemaRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSchemaRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseListSchemaRequest } as ListSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListSchemaRequest {
    const message = { ...baseListSchemaRequest } as ListSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: ListSchemaRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<ListSchemaRequest>): ListSchemaRequest {
    const message = { ...baseListSchemaRequest } as ListSchemaRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseListSchemaResponse: object = {};

export const ListSchemaResponse = {
  encode(message: ListSchemaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.schemas) {
      Schema.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListSchemaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseListSchemaResponse } as ListSchemaResponse;
    message.schemas = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemas.push(Schema.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ListSchemaResponse {
    const message = { ...baseListSchemaResponse } as ListSchemaResponse;
    message.schemas = [];
    if (object.schemas !== undefined && object.schemas !== null) {
      for (const e of object.schemas) {
        message.schemas.push(Schema.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: ListSchemaResponse): unknown {
    const obj: any = {};
    if (message.schemas) {
      obj.schemas = message.schemas.map((e) => (e ? Schema.toJSON(e) : undefined));
    } else {
      obj.schemas = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<ListSchemaResponse>): ListSchemaResponse {
    const message = { ...baseListSchemaResponse } as ListSchemaResponse;
    message.schemas = [];
    if (object.schemas !== undefined && object.schemas !== null) {
      for (const e of object.schemas) {
        message.schemas.push(Schema.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseGetSchemaRequest: object = { schemaName: "" };

export const GetSchemaRequest = {
  encode(message: GetSchemaRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.schemaName !== "") {
      writer.uint32(10).string(message.schemaName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSchemaRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGetSchemaRequest } as GetSchemaRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemaName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetSchemaRequest {
    const message = { ...baseGetSchemaRequest } as GetSchemaRequest;
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = String(object.schemaName);
    } else {
      message.schemaName = "";
    }
    return message;
  },

  toJSON(message: GetSchemaRequest): unknown {
    const obj: any = {};
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    return obj;
  },

  fromPartial(object: DeepPartial<GetSchemaRequest>): GetSchemaRequest {
    const message = { ...baseGetSchemaRequest } as GetSchemaRequest;
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = object.schemaName;
    } else {
      message.schemaName = "";
    }
    return message;
  },
};

const baseGetSchemaResponse: object = {};

export const GetSchemaResponse = {
  encode(message: GetSchemaResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.schema !== undefined) {
      Schema.encode(message.schema, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSchemaResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGetSchemaResponse } as GetSchemaResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schema = Schema.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetSchemaResponse {
    const message = { ...baseGetSchemaResponse } as GetSchemaResponse;
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = Schema.fromJSON(object.schema);
    } else {
      message.schema = undefined;
    }
    return message;
  },

  toJSON(message: GetSchemaResponse): unknown {
    const obj: any = {};
    message.schema !== undefined && (obj.schema = message.schema ? Schema.toJSON(message.schema) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GetSchemaResponse>): GetSchemaResponse {
    const message = { ...baseGetSchemaResponse } as GetSchemaResponse;
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = Schema.fromPartial(object.schema);
    } else {
      message.schema = undefined;
    }
    return message;
  },
};

const baseGetDocRequest: object = { schemaName: "", docID: "" };

export const GetDocRequest = {
  encode(message: GetDocRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.schemaName !== "") {
      writer.uint32(10).string(message.schemaName);
    }
    if (message.docID !== "") {
      writer.uint32(18).string(message.docID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetDocRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGetDocRequest } as GetDocRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemaName = reader.string();
          break;
        case 2:
          message.docID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetDocRequest {
    const message = { ...baseGetDocRequest } as GetDocRequest;
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = String(object.schemaName);
    } else {
      message.schemaName = "";
    }
    if (object.docID !== undefined && object.docID !== null) {
      message.docID = String(object.docID);
    } else {
      message.docID = "";
    }
    return message;
  },

  toJSON(message: GetDocRequest): unknown {
    const obj: any = {};
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    message.docID !== undefined && (obj.docID = message.docID);
    return obj;
  },

  fromPartial(object: DeepPartial<GetDocRequest>): GetDocRequest {
    const message = { ...baseGetDocRequest } as GetDocRequest;
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = object.schemaName;
    } else {
      message.schemaName = "";
    }
    if (object.docID !== undefined && object.docID !== null) {
      message.docID = object.docID;
    } else {
      message.docID = "";
    }
    return message;
  },
};

const baseGetDocResponse: object = {};

export const GetDocResponse = {
  encode(message: GetDocResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.doc !== undefined) {
      Doc.encode(message.doc, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetDocResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGetDocResponse } as GetDocResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.doc = Doc.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetDocResponse {
    const message = { ...baseGetDocResponse } as GetDocResponse;
    if (object.doc !== undefined && object.doc !== null) {
      message.doc = Doc.fromJSON(object.doc);
    } else {
      message.doc = undefined;
    }
    return message;
  },

  toJSON(message: GetDocResponse): unknown {
    const obj: any = {};
    message.doc !== undefined && (obj.doc = message.doc ? Doc.toJSON(message.doc) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GetDocResponse>): GetDocResponse {
    const message = { ...baseGetDocResponse } as GetDocResponse;
    if (object.doc !== undefined && object.doc !== null) {
      message.doc = Doc.fromPartial(object.doc);
    } else {
      message.doc = undefined;
    }
    return message;
  },
};

const baseSearchDocRequest: object = {
  schemaName: "",
  query: "",
  queryField: "",
  aggsField: "",
  orderBy: "",
  limit: Long.ZERO,
  page: Long.ZERO,
};

export const SearchDocRequest = {
  encode(message: SearchDocRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.schemaName !== "") {
      writer.uint32(10).string(message.schemaName);
    }
    if (message.query !== "") {
      writer.uint32(18).string(message.query);
    }
    for (const v of message.queryField) {
      writer.uint32(26).string(v!);
    }
    for (const v of message.aggsField) {
      writer.uint32(34).string(v!);
    }
    if (message.filter !== undefined) {
      SearchFilters.encode(message.filter, writer.uint32(42).fork()).ldelim();
    }
    if (message.orderBy !== "") {
      writer.uint32(50).string(message.orderBy);
    }
    if (!message.limit.isZero()) {
      writer.uint32(56).int64(message.limit);
    }
    if (!message.page.isZero()) {
      writer.uint32(64).int64(message.page);
    }
    if (message.factor !== undefined) {
      FieldValueFactor.encode(message.factor, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchDocRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSearchDocRequest } as SearchDocRequest;
    message.queryField = [];
    message.aggsField = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schemaName = reader.string();
          break;
        case 2:
          message.query = reader.string();
          break;
        case 3:
          message.queryField.push(reader.string());
          break;
        case 4:
          message.aggsField.push(reader.string());
          break;
        case 5:
          message.filter = SearchFilters.decode(reader, reader.uint32());
          break;
        case 6:
          message.orderBy = reader.string();
          break;
        case 7:
          message.limit = reader.int64() as Long;
          break;
        case 8:
          message.page = reader.int64() as Long;
          break;
        case 9:
          message.factor = FieldValueFactor.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchDocRequest {
    const message = { ...baseSearchDocRequest } as SearchDocRequest;
    message.queryField = [];
    message.aggsField = [];
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = String(object.schemaName);
    } else {
      message.schemaName = "";
    }
    if (object.query !== undefined && object.query !== null) {
      message.query = String(object.query);
    } else {
      message.query = "";
    }
    if (object.queryField !== undefined && object.queryField !== null) {
      for (const e of object.queryField) {
        message.queryField.push(String(e));
      }
    }
    if (object.aggsField !== undefined && object.aggsField !== null) {
      for (const e of object.aggsField) {
        message.aggsField.push(String(e));
      }
    }
    if (object.filter !== undefined && object.filter !== null) {
      message.filter = SearchFilters.fromJSON(object.filter);
    } else {
      message.filter = undefined;
    }
    if (object.orderBy !== undefined && object.orderBy !== null) {
      message.orderBy = String(object.orderBy);
    } else {
      message.orderBy = "";
    }
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = Long.fromString(object.limit);
    } else {
      message.limit = Long.ZERO;
    }
    if (object.page !== undefined && object.page !== null) {
      message.page = Long.fromString(object.page);
    } else {
      message.page = Long.ZERO;
    }
    if (object.factor !== undefined && object.factor !== null) {
      message.factor = FieldValueFactor.fromJSON(object.factor);
    } else {
      message.factor = undefined;
    }
    return message;
  },

  toJSON(message: SearchDocRequest): unknown {
    const obj: any = {};
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    message.query !== undefined && (obj.query = message.query);
    if (message.queryField) {
      obj.queryField = message.queryField.map((e) => e);
    } else {
      obj.queryField = [];
    }
    if (message.aggsField) {
      obj.aggsField = message.aggsField.map((e) => e);
    } else {
      obj.aggsField = [];
    }
    message.filter !== undefined &&
      (obj.filter = message.filter ? SearchFilters.toJSON(message.filter) : undefined);
    message.orderBy !== undefined && (obj.orderBy = message.orderBy);
    message.limit !== undefined && (obj.limit = (message.limit || Long.ZERO).toString());
    message.page !== undefined && (obj.page = (message.page || Long.ZERO).toString());
    message.factor !== undefined &&
      (obj.factor = message.factor ? FieldValueFactor.toJSON(message.factor) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<SearchDocRequest>): SearchDocRequest {
    const message = { ...baseSearchDocRequest } as SearchDocRequest;
    message.queryField = [];
    message.aggsField = [];
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = object.schemaName;
    } else {
      message.schemaName = "";
    }
    if (object.query !== undefined && object.query !== null) {
      message.query = object.query;
    } else {
      message.query = "";
    }
    if (object.queryField !== undefined && object.queryField !== null) {
      for (const e of object.queryField) {
        message.queryField.push(e);
      }
    }
    if (object.aggsField !== undefined && object.aggsField !== null) {
      for (const e of object.aggsField) {
        message.aggsField.push(e);
      }
    }
    if (object.filter !== undefined && object.filter !== null) {
      message.filter = SearchFilters.fromPartial(object.filter);
    } else {
      message.filter = undefined;
    }
    if (object.orderBy !== undefined && object.orderBy !== null) {
      message.orderBy = object.orderBy;
    } else {
      message.orderBy = "";
    }
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = object.limit as Long;
    } else {
      message.limit = Long.ZERO;
    }
    if (object.page !== undefined && object.page !== null) {
      message.page = object.page as Long;
    } else {
      message.page = Long.ZERO;
    }
    if (object.factor !== undefined && object.factor !== null) {
      message.factor = FieldValueFactor.fromPartial(object.factor);
    } else {
      message.factor = undefined;
    }
    return message;
  },
};

const baseSearchDocResponse: object = { searchTime: Long.ZERO, schemaName: "" };

export const SearchDocResponse = {
  encode(message: SearchDocResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.searchTime.isZero()) {
      writer.uint32(8).int64(message.searchTime);
    }
    if (message.schemaName !== "") {
      writer.uint32(18).string(message.schemaName);
    }
    if (message.meta !== undefined) {
      Meta.encode(message.meta, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.items) {
      DocItem.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    Object.entries(message.facet).forEach(([key, value]) => {
      SearchDocResponse_FacetEntry.encode({ key: key as any, value }, writer.uint32(42).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchDocResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSearchDocResponse } as SearchDocResponse;
    message.items = [];
    message.facet = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.searchTime = reader.int64() as Long;
          break;
        case 2:
          message.schemaName = reader.string();
          break;
        case 3:
          message.meta = Meta.decode(reader, reader.uint32());
          break;
        case 4:
          message.items.push(DocItem.decode(reader, reader.uint32()));
          break;
        case 5:
          const entry5 = SearchDocResponse_FacetEntry.decode(reader, reader.uint32());
          if (entry5.value !== undefined) {
            message.facet[entry5.key] = entry5.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchDocResponse {
    const message = { ...baseSearchDocResponse } as SearchDocResponse;
    message.items = [];
    message.facet = {};
    if (object.searchTime !== undefined && object.searchTime !== null) {
      message.searchTime = Long.fromString(object.searchTime);
    } else {
      message.searchTime = Long.ZERO;
    }
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = String(object.schemaName);
    } else {
      message.schemaName = "";
    }
    if (object.meta !== undefined && object.meta !== null) {
      message.meta = Meta.fromJSON(object.meta);
    } else {
      message.meta = undefined;
    }
    if (object.items !== undefined && object.items !== null) {
      for (const e of object.items) {
        message.items.push(DocItem.fromJSON(e));
      }
    }
    if (object.facet !== undefined && object.facet !== null) {
      Object.entries(object.facet).forEach(([key, value]) => {
        message.facet[key] = SearchFilters.fromJSON(value);
      });
    }
    return message;
  },

  toJSON(message: SearchDocResponse): unknown {
    const obj: any = {};
    message.searchTime !== undefined && (obj.searchTime = (message.searchTime || Long.ZERO).toString());
    message.schemaName !== undefined && (obj.schemaName = message.schemaName);
    message.meta !== undefined && (obj.meta = message.meta ? Meta.toJSON(message.meta) : undefined);
    if (message.items) {
      obj.items = message.items.map((e) => (e ? DocItem.toJSON(e) : undefined));
    } else {
      obj.items = [];
    }
    obj.facet = {};
    if (message.facet) {
      Object.entries(message.facet).forEach(([k, v]) => {
        obj.facet[k] = SearchFilters.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SearchDocResponse>): SearchDocResponse {
    const message = { ...baseSearchDocResponse } as SearchDocResponse;
    message.items = [];
    message.facet = {};
    if (object.searchTime !== undefined && object.searchTime !== null) {
      message.searchTime = object.searchTime as Long;
    } else {
      message.searchTime = Long.ZERO;
    }
    if (object.schemaName !== undefined && object.schemaName !== null) {
      message.schemaName = object.schemaName;
    } else {
      message.schemaName = "";
    }
    if (object.meta !== undefined && object.meta !== null) {
      message.meta = Meta.fromPartial(object.meta);
    } else {
      message.meta = undefined;
    }
    if (object.items !== undefined && object.items !== null) {
      for (const e of object.items) {
        message.items.push(DocItem.fromPartial(e));
      }
    }
    if (object.facet !== undefined && object.facet !== null) {
      Object.entries(object.facet).forEach(([key, value]) => {
        if (value !== undefined) {
          message.facet[key] = SearchFilters.fromPartial(value);
        }
      });
    }
    return message;
  },
};

const baseSearchDocResponse_FacetEntry: object = { key: "" };

export const SearchDocResponse_FacetEntry = {
  encode(message: SearchDocResponse_FacetEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      SearchFilters.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchDocResponse_FacetEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSearchDocResponse_FacetEntry } as SearchDocResponse_FacetEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = SearchFilters.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchDocResponse_FacetEntry {
    const message = { ...baseSearchDocResponse_FacetEntry } as SearchDocResponse_FacetEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = SearchFilters.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: SearchDocResponse_FacetEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? SearchFilters.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<SearchDocResponse_FacetEntry>): SearchDocResponse_FacetEntry {
    const message = { ...baseSearchDocResponse_FacetEntry } as SearchDocResponse_FacetEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = SearchFilters.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

const baseMeta: object = {
  currentPage: Long.ZERO,
  totalPages: Long.ZERO,
  totalResults: Long.ZERO,
  size: Long.ZERO,
  sortedBy: "",
};

export const Meta = {
  encode(message: Meta, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.currentPage.isZero()) {
      writer.uint32(8).int64(message.currentPage);
    }
    if (!message.totalPages.isZero()) {
      writer.uint32(16).int64(message.totalPages);
    }
    if (!message.totalResults.isZero()) {
      writer.uint32(24).int64(message.totalResults);
    }
    if (!message.size.isZero()) {
      writer.uint32(32).int64(message.size);
    }
    if (message.sortedBy !== "") {
      writer.uint32(42).string(message.sortedBy);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Meta {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMeta } as Meta;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.currentPage = reader.int64() as Long;
          break;
        case 2:
          message.totalPages = reader.int64() as Long;
          break;
        case 3:
          message.totalResults = reader.int64() as Long;
          break;
        case 4:
          message.size = reader.int64() as Long;
          break;
        case 5:
          message.sortedBy = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Meta {
    const message = { ...baseMeta } as Meta;
    if (object.currentPage !== undefined && object.currentPage !== null) {
      message.currentPage = Long.fromString(object.currentPage);
    } else {
      message.currentPage = Long.ZERO;
    }
    if (object.totalPages !== undefined && object.totalPages !== null) {
      message.totalPages = Long.fromString(object.totalPages);
    } else {
      message.totalPages = Long.ZERO;
    }
    if (object.totalResults !== undefined && object.totalResults !== null) {
      message.totalResults = Long.fromString(object.totalResults);
    } else {
      message.totalResults = Long.ZERO;
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = Long.fromString(object.size);
    } else {
      message.size = Long.ZERO;
    }
    if (object.sortedBy !== undefined && object.sortedBy !== null) {
      message.sortedBy = String(object.sortedBy);
    } else {
      message.sortedBy = "";
    }
    return message;
  },

  toJSON(message: Meta): unknown {
    const obj: any = {};
    message.currentPage !== undefined && (obj.currentPage = (message.currentPage || Long.ZERO).toString());
    message.totalPages !== undefined && (obj.totalPages = (message.totalPages || Long.ZERO).toString());
    message.totalResults !== undefined && (obj.totalResults = (message.totalResults || Long.ZERO).toString());
    message.size !== undefined && (obj.size = (message.size || Long.ZERO).toString());
    message.sortedBy !== undefined && (obj.sortedBy = message.sortedBy);
    return obj;
  },

  fromPartial(object: DeepPartial<Meta>): Meta {
    const message = { ...baseMeta } as Meta;
    if (object.currentPage !== undefined && object.currentPage !== null) {
      message.currentPage = object.currentPage as Long;
    } else {
      message.currentPage = Long.ZERO;
    }
    if (object.totalPages !== undefined && object.totalPages !== null) {
      message.totalPages = object.totalPages as Long;
    } else {
      message.totalPages = Long.ZERO;
    }
    if (object.totalResults !== undefined && object.totalResults !== null) {
      message.totalResults = object.totalResults as Long;
    } else {
      message.totalResults = Long.ZERO;
    }
    if (object.size !== undefined && object.size !== null) {
      message.size = object.size as Long;
    } else {
      message.size = Long.ZERO;
    }
    if (object.sortedBy !== undefined && object.sortedBy !== null) {
      message.sortedBy = object.sortedBy;
    } else {
      message.sortedBy = "";
    }
    return message;
  },
};

const baseDocItem: object = { data: "", score: 0 };

export const DocItem = {
  encode(message: DocItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.Highlight).forEach(([key, value]) => {
      DocItem_HighlightEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    if (message.data !== "") {
      writer.uint32(18).string(message.data);
    }
    if (message.score !== 0) {
      writer.uint32(25).double(message.score);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DocItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDocItem } as DocItem;
    message.Highlight = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = DocItem_HighlightEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.Highlight[entry1.key] = entry1.value;
          }
          break;
        case 2:
          message.data = reader.string();
          break;
        case 3:
          message.score = reader.double();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DocItem {
    const message = { ...baseDocItem } as DocItem;
    message.Highlight = {};
    if (object.Highlight !== undefined && object.Highlight !== null) {
      Object.entries(object.Highlight).forEach(([key, value]) => {
        message.Highlight[key] = strings.fromJSON(value);
      });
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = Number(object.score);
    } else {
      message.score = 0;
    }
    return message;
  },

  toJSON(message: DocItem): unknown {
    const obj: any = {};
    obj.Highlight = {};
    if (message.Highlight) {
      Object.entries(message.Highlight).forEach(([k, v]) => {
        obj.Highlight[k] = strings.toJSON(v);
      });
    }
    message.data !== undefined && (obj.data = message.data);
    message.score !== undefined && (obj.score = message.score);
    return obj;
  },

  fromPartial(object: DeepPartial<DocItem>): DocItem {
    const message = { ...baseDocItem } as DocItem;
    message.Highlight = {};
    if (object.Highlight !== undefined && object.Highlight !== null) {
      Object.entries(object.Highlight).forEach(([key, value]) => {
        if (value !== undefined) {
          message.Highlight[key] = strings.fromPartial(value);
        }
      });
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = 0;
    }
    return message;
  },
};

const baseDocItem_HighlightEntry: object = { key: "" };

export const DocItem_HighlightEntry = {
  encode(message: DocItem_HighlightEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      strings.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DocItem_HighlightEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDocItem_HighlightEntry } as DocItem_HighlightEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = strings.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DocItem_HighlightEntry {
    const message = { ...baseDocItem_HighlightEntry } as DocItem_HighlightEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = strings.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: DocItem_HighlightEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? strings.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<DocItem_HighlightEntry>): DocItem_HighlightEntry {
    const message = { ...baseDocItem_HighlightEntry } as DocItem_HighlightEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = strings.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

const basestrings: object = { strings: "" };

export const strings = {
  encode(message: strings, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.strings) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): strings {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basestrings } as strings;
    message.strings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.strings.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): strings {
    const message = { ...basestrings } as strings;
    message.strings = [];
    if (object.strings !== undefined && object.strings !== null) {
      for (const e of object.strings) {
        message.strings.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: strings): unknown {
    const obj: any = {};
    if (message.strings) {
      obj.strings = message.strings.map((e) => e);
    } else {
      obj.strings = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<strings>): strings {
    const message = { ...basestrings } as strings;
    message.strings = [];
    if (object.strings !== undefined && object.strings !== null) {
      for (const e of object.strings) {
        message.strings.push(e);
      }
    }
    return message;
  },
};

const baseSearchFilters: object = {};

export const SearchFilters = {
  encode(message: SearchFilters, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.searchFilter) {
      SearchFilter.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchFilters {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSearchFilters } as SearchFilters;
    message.searchFilter = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.searchFilter.push(SearchFilter.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchFilters {
    const message = { ...baseSearchFilters } as SearchFilters;
    message.searchFilter = [];
    if (object.searchFilter !== undefined && object.searchFilter !== null) {
      for (const e of object.searchFilter) {
        message.searchFilter.push(SearchFilter.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: SearchFilters): unknown {
    const obj: any = {};
    if (message.searchFilter) {
      obj.searchFilter = message.searchFilter.map((e) => (e ? SearchFilter.toJSON(e) : undefined));
    } else {
      obj.searchFilter = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SearchFilters>): SearchFilters {
    const message = { ...baseSearchFilters } as SearchFilters;
    message.searchFilter = [];
    if (object.searchFilter !== undefined && object.searchFilter !== null) {
      for (const e of object.searchFilter) {
        message.searchFilter.push(SearchFilter.fromPartial(e));
      }
    }
    return message;
  },
};

const baseSearchFilter: object = { type: "", field: "", value: "", from: 0, to: 0, docCount: Long.ZERO };

export const SearchFilter = {
  encode(message: SearchFilter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.field !== "") {
      writer.uint32(18).string(message.field);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    if (message.from !== 0) {
      writer.uint32(37).float(message.from);
    }
    if (message.to !== 0) {
      writer.uint32(45).float(message.to);
    }
    if (!message.docCount.isZero()) {
      writer.uint32(48).int64(message.docCount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SearchFilter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSearchFilter } as SearchFilter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.field = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        case 4:
          message.from = reader.float();
          break;
        case 5:
          message.to = reader.float();
          break;
        case 6:
          message.docCount = reader.int64() as Long;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SearchFilter {
    const message = { ...baseSearchFilter } as SearchFilter;
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (object.field !== undefined && object.field !== null) {
      message.field = String(object.field);
    } else {
      message.field = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = Number(object.from);
    } else {
      message.from = 0;
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = Number(object.to);
    } else {
      message.to = 0;
    }
    if (object.docCount !== undefined && object.docCount !== null) {
      message.docCount = Long.fromString(object.docCount);
    } else {
      message.docCount = Long.ZERO;
    }
    return message;
  },

  toJSON(message: SearchFilter): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.field !== undefined && (obj.field = message.field);
    message.value !== undefined && (obj.value = message.value);
    message.from !== undefined && (obj.from = message.from);
    message.to !== undefined && (obj.to = message.to);
    message.docCount !== undefined && (obj.docCount = (message.docCount || Long.ZERO).toString());
    return obj;
  },

  fromPartial(object: DeepPartial<SearchFilter>): SearchFilter {
    const message = { ...baseSearchFilter } as SearchFilter;
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (object.field !== undefined && object.field !== null) {
      message.field = object.field;
    } else {
      message.field = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.from !== undefined && object.from !== null) {
      message.from = object.from;
    } else {
      message.from = 0;
    }
    if (object.to !== undefined && object.to !== null) {
      message.to = object.to;
    } else {
      message.to = 0;
    }
    if (object.docCount !== undefined && object.docCount !== null) {
      message.docCount = object.docCount as Long;
    } else {
      message.docCount = Long.ZERO;
    }
    return message;
  },
};

const baseFieldValueFactor: object = { field: "", factor: 0, missing: 0, weight: 0, modifier: "" };

export const FieldValueFactor = {
  encode(message: FieldValueFactor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.field !== "") {
      writer.uint32(10).string(message.field);
    }
    if (message.factor !== 0) {
      writer.uint32(17).double(message.factor);
    }
    if (message.missing !== 0) {
      writer.uint32(25).double(message.missing);
    }
    if (message.weight !== 0) {
      writer.uint32(33).double(message.weight);
    }
    if (message.modifier !== "") {
      writer.uint32(42).string(message.modifier);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FieldValueFactor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFieldValueFactor } as FieldValueFactor;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.field = reader.string();
          break;
        case 2:
          message.factor = reader.double();
          break;
        case 3:
          message.missing = reader.double();
          break;
        case 4:
          message.weight = reader.double();
          break;
        case 5:
          message.modifier = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FieldValueFactor {
    const message = { ...baseFieldValueFactor } as FieldValueFactor;
    if (object.field !== undefined && object.field !== null) {
      message.field = String(object.field);
    } else {
      message.field = "";
    }
    if (object.factor !== undefined && object.factor !== null) {
      message.factor = Number(object.factor);
    } else {
      message.factor = 0;
    }
    if (object.missing !== undefined && object.missing !== null) {
      message.missing = Number(object.missing);
    } else {
      message.missing = 0;
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = Number(object.weight);
    } else {
      message.weight = 0;
    }
    if (object.modifier !== undefined && object.modifier !== null) {
      message.modifier = String(object.modifier);
    } else {
      message.modifier = "";
    }
    return message;
  },

  toJSON(message: FieldValueFactor): unknown {
    const obj: any = {};
    message.field !== undefined && (obj.field = message.field);
    message.factor !== undefined && (obj.factor = message.factor);
    message.missing !== undefined && (obj.missing = message.missing);
    message.weight !== undefined && (obj.weight = message.weight);
    message.modifier !== undefined && (obj.modifier = message.modifier);
    return obj;
  },

  fromPartial(object: DeepPartial<FieldValueFactor>): FieldValueFactor {
    const message = { ...baseFieldValueFactor } as FieldValueFactor;
    if (object.field !== undefined && object.field !== null) {
      message.field = object.field;
    } else {
      message.field = "";
    }
    if (object.factor !== undefined && object.factor !== null) {
      message.factor = object.factor;
    } else {
      message.factor = 0;
    }
    if (object.missing !== undefined && object.missing !== null) {
      message.missing = object.missing;
    } else {
      message.missing = 0;
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = object.weight;
    } else {
      message.weight = 0;
    }
    if (object.modifier !== undefined && object.modifier !== null) {
      message.modifier = object.modifier;
    } else {
      message.modifier = "";
    }
    return message;
  },
};

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

const baseSQLQueryRequest: object = { sql: "" };

export const SQLQueryRequest = {
  encode(message: SQLQueryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sql !== "") {
      writer.uint32(10).string(message.sql);
    }
    for (const v of message.arguments) {
      Argument.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLQueryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLQueryRequest } as SQLQueryRequest;
    message.arguments = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sql = reader.string();
          break;
        case 2:
          message.arguments.push(Argument.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLQueryRequest {
    const message = { ...baseSQLQueryRequest } as SQLQueryRequest;
    message.arguments = [];
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

  toJSON(message: SQLQueryRequest): unknown {
    const obj: any = {};
    message.sql !== undefined && (obj.sql = message.sql);
    if (message.arguments) {
      obj.arguments = message.arguments.map((e) => (e ? Argument.toJSON(e) : undefined));
    } else {
      obj.arguments = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SQLQueryRequest>): SQLQueryRequest {
    const message = { ...baseSQLQueryRequest } as SQLQueryRequest;
    message.arguments = [];
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

const baseSQLQueryResponse: object = { tookTimes: 0 };

export const SQLQueryResponse = {
  encode(message: SQLQueryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.results) {
      ResultSet.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.tookTimes !== 0) {
      writer.uint32(21).float(message.tookTimes);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLQueryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLQueryResponse } as SQLQueryResponse;
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

  fromJSON(object: any): SQLQueryResponse {
    const message = { ...baseSQLQueryResponse } as SQLQueryResponse;
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

  toJSON(message: SQLQueryResponse): unknown {
    const obj: any = {};
    if (message.results) {
      obj.results = message.results.map((e) => (e ? ResultSet.toJSON(e) : undefined));
    } else {
      obj.results = [];
    }
    message.tookTimes !== undefined && (obj.tookTimes = message.tookTimes);
    return obj;
  },

  fromPartial(object: DeepPartial<SQLQueryResponse>): SQLQueryResponse {
    const message = { ...baseSQLQueryResponse } as SQLQueryResponse;
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

const baseSimpleSQLQueryResponse: object = {};

export const SimpleSQLQueryResponse = {
  encode(message: SimpleSQLQueryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.result) {
      SimpleSQLQueryResponse_ResultSet.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SimpleSQLQueryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSimpleSQLQueryResponse } as SimpleSQLQueryResponse;
    message.result = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.result.push(SimpleSQLQueryResponse_ResultSet.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SimpleSQLQueryResponse {
    const message = { ...baseSimpleSQLQueryResponse } as SimpleSQLQueryResponse;
    message.result = [];
    if (object.result !== undefined && object.result !== null) {
      for (const e of object.result) {
        message.result.push(SimpleSQLQueryResponse_ResultSet.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: SimpleSQLQueryResponse): unknown {
    const obj: any = {};
    if (message.result) {
      obj.result = message.result.map((e) => (e ? SimpleSQLQueryResponse_ResultSet.toJSON(e) : undefined));
    } else {
      obj.result = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SimpleSQLQueryResponse>): SimpleSQLQueryResponse {
    const message = { ...baseSimpleSQLQueryResponse } as SimpleSQLQueryResponse;
    message.result = [];
    if (object.result !== undefined && object.result !== null) {
      for (const e of object.result) {
        message.result.push(SimpleSQLQueryResponse_ResultSet.fromPartial(e));
      }
    }
    return message;
  },
};

const baseSimpleSQLQueryResponse_ResultSet: object = {};

export const SimpleSQLQueryResponse_ResultSet = {
  encode(message: SimpleSQLQueryResponse_ResultSet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.row).forEach(([key, value]) => {
      SimpleSQLQueryResponse_ResultSet_RowEntry.encode(
        { key: key as any, value },
        writer.uint32(10).fork(),
      ).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SimpleSQLQueryResponse_ResultSet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSimpleSQLQueryResponse_ResultSet } as SimpleSQLQueryResponse_ResultSet;
    message.row = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = SimpleSQLQueryResponse_ResultSet_RowEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.row[entry1.key] = entry1.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SimpleSQLQueryResponse_ResultSet {
    const message = { ...baseSimpleSQLQueryResponse_ResultSet } as SimpleSQLQueryResponse_ResultSet;
    message.row = {};
    if (object.row !== undefined && object.row !== null) {
      Object.entries(object.row).forEach(([key, value]) => {
        message.row[key] = SimpleSQLQueryResponse_RowValue.fromJSON(value);
      });
    }
    return message;
  },

  toJSON(message: SimpleSQLQueryResponse_ResultSet): unknown {
    const obj: any = {};
    obj.row = {};
    if (message.row) {
      Object.entries(message.row).forEach(([k, v]) => {
        obj.row[k] = SimpleSQLQueryResponse_RowValue.toJSON(v);
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SimpleSQLQueryResponse_ResultSet>): SimpleSQLQueryResponse_ResultSet {
    const message = { ...baseSimpleSQLQueryResponse_ResultSet } as SimpleSQLQueryResponse_ResultSet;
    message.row = {};
    if (object.row !== undefined && object.row !== null) {
      Object.entries(object.row).forEach(([key, value]) => {
        if (value !== undefined) {
          message.row[key] = SimpleSQLQueryResponse_RowValue.fromPartial(value);
        }
      });
    }
    return message;
  },
};

const baseSimpleSQLQueryResponse_ResultSet_RowEntry: object = { key: "" };

export const SimpleSQLQueryResponse_ResultSet_RowEntry = {
  encode(
    message: SimpleSQLQueryResponse_ResultSet_RowEntry,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      SimpleSQLQueryResponse_RowValue.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SimpleSQLQueryResponse_ResultSet_RowEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseSimpleSQLQueryResponse_ResultSet_RowEntry,
    } as SimpleSQLQueryResponse_ResultSet_RowEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = SimpleSQLQueryResponse_RowValue.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SimpleSQLQueryResponse_ResultSet_RowEntry {
    const message = {
      ...baseSimpleSQLQueryResponse_ResultSet_RowEntry,
    } as SimpleSQLQueryResponse_ResultSet_RowEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = SimpleSQLQueryResponse_RowValue.fromJSON(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },

  toJSON(message: SimpleSQLQueryResponse_ResultSet_RowEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined &&
      (obj.value = message.value ? SimpleSQLQueryResponse_RowValue.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<SimpleSQLQueryResponse_ResultSet_RowEntry>,
  ): SimpleSQLQueryResponse_ResultSet_RowEntry {
    const message = {
      ...baseSimpleSQLQueryResponse_ResultSet_RowEntry,
    } as SimpleSQLQueryResponse_ResultSet_RowEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = SimpleSQLQueryResponse_RowValue.fromPartial(object.value);
    } else {
      message.value = undefined;
    }
    return message;
  },
};

const baseSimpleSQLQueryResponse_RowValue: object = { value: "", columnValueType: 0 };

export const SimpleSQLQueryResponse_RowValue = {
  encode(message: SimpleSQLQueryResponse_RowValue, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.value !== "") {
      writer.uint32(10).string(message.value);
    }
    if (message.columnValueType !== 0) {
      writer.uint32(16).int32(message.columnValueType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SimpleSQLQueryResponse_RowValue {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSimpleSQLQueryResponse_RowValue } as SimpleSQLQueryResponse_RowValue;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.value = reader.string();
          break;
        case 2:
          message.columnValueType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SimpleSQLQueryResponse_RowValue {
    const message = { ...baseSimpleSQLQueryResponse_RowValue } as SimpleSQLQueryResponse_RowValue;
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    if (object.columnValueType !== undefined && object.columnValueType !== null) {
      message.columnValueType = columnValueTypeFromJSON(object.columnValueType);
    } else {
      message.columnValueType = 0;
    }
    return message;
  },

  toJSON(message: SimpleSQLQueryResponse_RowValue): unknown {
    const obj: any = {};
    message.value !== undefined && (obj.value = message.value);
    message.columnValueType !== undefined &&
      (obj.columnValueType = columnValueTypeToJSON(message.columnValueType));
    return obj;
  },

  fromPartial(object: DeepPartial<SimpleSQLQueryResponse_RowValue>): SimpleSQLQueryResponse_RowValue {
    const message = { ...baseSimpleSQLQueryResponse_RowValue } as SimpleSQLQueryResponse_RowValue;
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    if (object.columnValueType !== undefined && object.columnValueType !== null) {
      message.columnValueType = object.columnValueType;
    } else {
      message.columnValueType = 0;
    }
    return message;
  },
};

const baseSQLListTablesRequest: object = { keyword: "", uid: "", page: 0, pageSize: 0, database: "" };

export const SQLListTablesRequest = {
  encode(message: SQLListTablesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.keyword !== "") {
      writer.uint32(10).string(message.keyword);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    if (message.page !== 0) {
      writer.uint32(24).int32(message.page);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int32(message.pageSize);
    }
    if (message.database !== "") {
      writer.uint32(42).string(message.database);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLListTablesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLListTablesRequest } as SQLListTablesRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.keyword = reader.string();
          break;
        case 2:
          message.uid = reader.string();
          break;
        case 3:
          message.page = reader.int32();
          break;
        case 4:
          message.pageSize = reader.int32();
          break;
        case 5:
          message.database = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLListTablesRequest {
    const message = { ...baseSQLListTablesRequest } as SQLListTablesRequest;
    if (object.keyword !== undefined && object.keyword !== null) {
      message.keyword = String(object.keyword);
    } else {
      message.keyword = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.page !== undefined && object.page !== null) {
      message.page = Number(object.page);
    } else {
      message.page = 0;
    }
    if (object.pageSize !== undefined && object.pageSize !== null) {
      message.pageSize = Number(object.pageSize);
    } else {
      message.pageSize = 0;
    }
    if (object.database !== undefined && object.database !== null) {
      message.database = String(object.database);
    } else {
      message.database = "";
    }
    return message;
  },

  toJSON(message: SQLListTablesRequest): unknown {
    const obj: any = {};
    message.keyword !== undefined && (obj.keyword = message.keyword);
    message.uid !== undefined && (obj.uid = message.uid);
    message.page !== undefined && (obj.page = message.page);
    message.pageSize !== undefined && (obj.pageSize = message.pageSize);
    message.database !== undefined && (obj.database = message.database);
    return obj;
  },

  fromPartial(object: DeepPartial<SQLListTablesRequest>): SQLListTablesRequest {
    const message = { ...baseSQLListTablesRequest } as SQLListTablesRequest;
    if (object.keyword !== undefined && object.keyword !== null) {
      message.keyword = object.keyword;
    } else {
      message.keyword = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.page !== undefined && object.page !== null) {
      message.page = object.page;
    } else {
      message.page = 0;
    }
    if (object.pageSize !== undefined && object.pageSize !== null) {
      message.pageSize = object.pageSize;
    } else {
      message.pageSize = 0;
    }
    if (object.database !== undefined && object.database !== null) {
      message.database = object.database;
    } else {
      message.database = "";
    }
    return message;
  },
};

const baseSQLListTablesResponse: object = {};

export const SQLListTablesResponse = {
  encode(message: SQLListTablesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tables) {
      TableInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLListTablesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLListTablesResponse } as SQLListTablesResponse;
    message.tables = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tables.push(TableInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLListTablesResponse {
    const message = { ...baseSQLListTablesResponse } as SQLListTablesResponse;
    message.tables = [];
    if (object.tables !== undefined && object.tables !== null) {
      for (const e of object.tables) {
        message.tables.push(TableInfo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: SQLListTablesResponse): unknown {
    const obj: any = {};
    if (message.tables) {
      obj.tables = message.tables.map((e) => (e ? TableInfo.toJSON(e) : undefined));
    } else {
      obj.tables = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SQLListTablesResponse>): SQLListTablesResponse {
    const message = { ...baseSQLListTablesResponse } as SQLListTablesResponse;
    message.tables = [];
    if (object.tables !== undefined && object.tables !== null) {
      for (const e of object.tables) {
        message.tables.push(TableInfo.fromPartial(e));
      }
    }
    return message;
  },
};

const baseSQLListDatabasesRequest: object = {};

export const SQLListDatabasesRequest = {
  encode(_: SQLListDatabasesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLListDatabasesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLListDatabasesRequest } as SQLListDatabasesRequest;
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

  fromJSON(_: any): SQLListDatabasesRequest {
    const message = { ...baseSQLListDatabasesRequest } as SQLListDatabasesRequest;
    return message;
  },

  toJSON(_: SQLListDatabasesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<SQLListDatabasesRequest>): SQLListDatabasesRequest {
    const message = { ...baseSQLListDatabasesRequest } as SQLListDatabasesRequest;
    return message;
  },
};

const baseSQLListDatabasesResponse: object = {};

export const SQLListDatabasesResponse = {
  encode(message: SQLListDatabasesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.databases) {
      DatabaseInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SQLListDatabasesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSQLListDatabasesResponse } as SQLListDatabasesResponse;
    message.databases = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.databases.push(DatabaseInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SQLListDatabasesResponse {
    const message = { ...baseSQLListDatabasesResponse } as SQLListDatabasesResponse;
    message.databases = [];
    if (object.databases !== undefined && object.databases !== null) {
      for (const e of object.databases) {
        message.databases.push(DatabaseInfo.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: SQLListDatabasesResponse): unknown {
    const obj: any = {};
    if (message.databases) {
      obj.databases = message.databases.map((e) => (e ? DatabaseInfo.toJSON(e) : undefined));
    } else {
      obj.databases = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<SQLListDatabasesResponse>): SQLListDatabasesResponse {
    const message = { ...baseSQLListDatabasesResponse } as SQLListDatabasesResponse;
    message.databases = [];
    if (object.databases !== undefined && object.databases !== null) {
      for (const e of object.databases) {
        message.databases.push(DatabaseInfo.fromPartial(e));
      }
    }
    return message;
  },
};

const baseShowCreateTableRequest: object = { databaseName: "", tableName: "" };

export const ShowCreateTableRequest = {
  encode(message: ShowCreateTableRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.databaseName !== "") {
      writer.uint32(10).string(message.databaseName);
    }
    if (message.tableName !== "") {
      writer.uint32(18).string(message.tableName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ShowCreateTableRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseShowCreateTableRequest } as ShowCreateTableRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.databaseName = reader.string();
          break;
        case 2:
          message.tableName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ShowCreateTableRequest {
    const message = { ...baseShowCreateTableRequest } as ShowCreateTableRequest;
    if (object.databaseName !== undefined && object.databaseName !== null) {
      message.databaseName = String(object.databaseName);
    } else {
      message.databaseName = "";
    }
    if (object.tableName !== undefined && object.tableName !== null) {
      message.tableName = String(object.tableName);
    } else {
      message.tableName = "";
    }
    return message;
  },

  toJSON(message: ShowCreateTableRequest): unknown {
    const obj: any = {};
    message.databaseName !== undefined && (obj.databaseName = message.databaseName);
    message.tableName !== undefined && (obj.tableName = message.tableName);
    return obj;
  },

  fromPartial(object: DeepPartial<ShowCreateTableRequest>): ShowCreateTableRequest {
    const message = { ...baseShowCreateTableRequest } as ShowCreateTableRequest;
    if (object.databaseName !== undefined && object.databaseName !== null) {
      message.databaseName = object.databaseName;
    } else {
      message.databaseName = "";
    }
    if (object.tableName !== undefined && object.tableName !== null) {
      message.tableName = object.tableName;
    } else {
      message.tableName = "";
    }
    return message;
  },
};

const baseShowCreateTableResponse: object = { schema: "" };

export const ShowCreateTableResponse = {
  encode(message: ShowCreateTableResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.schema !== "") {
      writer.uint32(10).string(message.schema);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ShowCreateTableResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseShowCreateTableResponse } as ShowCreateTableResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.schema = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ShowCreateTableResponse {
    const message = { ...baseShowCreateTableResponse } as ShowCreateTableResponse;
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = String(object.schema);
    } else {
      message.schema = "";
    }
    return message;
  },

  toJSON(message: ShowCreateTableResponse): unknown {
    const obj: any = {};
    message.schema !== undefined && (obj.schema = message.schema);
    return obj;
  },

  fromPartial(object: DeepPartial<ShowCreateTableResponse>): ShowCreateTableResponse {
    const message = { ...baseShowCreateTableResponse } as ShowCreateTableResponse;
    if (object.schema !== undefined && object.schema !== null) {
      message.schema = object.schema;
    } else {
      message.schema = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  ListSchema(request: DeepPartial<ListSchemaRequest>, metadata?: grpc.Metadata): Promise<ListSchemaResponse>;
  /** Parameters queries the parameters of the module. */
  GetSchema(request: DeepPartial<GetSchemaRequest>, metadata?: grpc.Metadata): Promise<GetSchemaResponse>;
  /** Parameters queries the parameters of the module. */
  GetDoc(request: DeepPartial<GetDocRequest>, metadata?: grpc.Metadata): Promise<GetDocResponse>;
  /** Parameters queries the parameters of the module. */
  SearchDoc(request: DeepPartial<SearchDocRequest>, metadata?: grpc.Metadata): Promise<SearchDocResponse>;
  /** Parameters queries the parameters of the module. */
  Params(request: DeepPartial<QueryParamsRequest>, metadata?: grpc.Metadata): Promise<QueryParamsResponse>;
  /** SQLQuery do queries */
  SQLQuery(request: DeepPartial<SQLQueryRequest>, metadata?: grpc.Metadata): Promise<SQLQueryResponse>;
  /** SimpleSQLQuery do queries */
  SimpleSQLQuery(
    request: DeepPartial<SQLQueryRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SimpleSQLQueryResponse>;
  /** Parameters queries the parameters of the module. */
  SQLListTables(
    request: DeepPartial<SQLListTablesRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SQLListTablesResponse>;
  SQLListDatabases(
    request: DeepPartial<SQLListDatabasesRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SQLListDatabasesResponse>;
  ShowCreateTable(
    request: DeepPartial<ShowCreateTableRequest>,
    metadata?: grpc.Metadata,
  ): Promise<ShowCreateTableResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;

  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ListSchema = this.ListSchema.bind(this);
    this.GetSchema = this.GetSchema.bind(this);
    this.GetDoc = this.GetDoc.bind(this);
    this.SearchDoc = this.SearchDoc.bind(this);
    this.Params = this.Params.bind(this);
    this.SQLQuery = this.SQLQuery.bind(this);
    this.SimpleSQLQuery = this.SimpleSQLQuery.bind(this);
    this.SQLListTables = this.SQLListTables.bind(this);
    this.SQLListDatabases = this.SQLListDatabases.bind(this);
    this.ShowCreateTable = this.ShowCreateTable.bind(this);
  }

  ListSchema(request: DeepPartial<ListSchemaRequest>, metadata?: grpc.Metadata): Promise<ListSchemaResponse> {
    return this.rpc.unary(QueryListSchemaDesc, ListSchemaRequest.fromPartial(request), metadata);
  }

  GetSchema(request: DeepPartial<GetSchemaRequest>, metadata?: grpc.Metadata): Promise<GetSchemaResponse> {
    return this.rpc.unary(QueryGetSchemaDesc, GetSchemaRequest.fromPartial(request), metadata);
  }

  GetDoc(request: DeepPartial<GetDocRequest>, metadata?: grpc.Metadata): Promise<GetDocResponse> {
    return this.rpc.unary(QueryGetDocDesc, GetDocRequest.fromPartial(request), metadata);
  }

  SearchDoc(request: DeepPartial<SearchDocRequest>, metadata?: grpc.Metadata): Promise<SearchDocResponse> {
    return this.rpc.unary(QuerySearchDocDesc, SearchDocRequest.fromPartial(request), metadata);
  }

  Params(request: DeepPartial<QueryParamsRequest>, metadata?: grpc.Metadata): Promise<QueryParamsResponse> {
    return this.rpc.unary(QueryParamsDesc, QueryParamsRequest.fromPartial(request), metadata);
  }

  SQLQuery(request: DeepPartial<SQLQueryRequest>, metadata?: grpc.Metadata): Promise<SQLQueryResponse> {
    return this.rpc.unary(QuerySQLQueryDesc, SQLQueryRequest.fromPartial(request), metadata);
  }

  SimpleSQLQuery(
    request: DeepPartial<SQLQueryRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SimpleSQLQueryResponse> {
    return this.rpc.unary(QuerySimpleSQLQueryDesc, SQLQueryRequest.fromPartial(request), metadata);
  }

  SQLListTables(
    request: DeepPartial<SQLListTablesRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SQLListTablesResponse> {
    return this.rpc.unary(QuerySQLListTablesDesc, SQLListTablesRequest.fromPartial(request), metadata);
  }

  SQLListDatabases(
    request: DeepPartial<SQLListDatabasesRequest>,
    metadata?: grpc.Metadata,
  ): Promise<SQLListDatabasesResponse> {
    return this.rpc.unary(QuerySQLListDatabasesDesc, SQLListDatabasesRequest.fromPartial(request), metadata);
  }

  ShowCreateTable(
    request: DeepPartial<ShowCreateTableRequest>,
    metadata?: grpc.Metadata,
  ): Promise<ShowCreateTableResponse> {
    return this.rpc.unary(QueryShowCreateTableDesc, ShowCreateTableRequest.fromPartial(request), metadata);
  }
}

export const QueryDesc = {
  serviceName: "blockved.glitterchain.index.Query",
};

export const QueryListSchemaDesc: UnaryMethodDefinitionish = {
  methodName: "ListSchema",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ListSchemaRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...ListSchemaResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QueryGetSchemaDesc: UnaryMethodDefinitionish = {
  methodName: "GetSchema",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetSchemaRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...GetSchemaResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QueryGetDocDesc: UnaryMethodDefinitionish = {
  methodName: "GetDoc",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return GetDocRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...GetDocResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QuerySearchDocDesc: UnaryMethodDefinitionish = {
  methodName: "SearchDoc",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SearchDocRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SearchDocResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
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

export const QuerySQLQueryDesc: UnaryMethodDefinitionish = {
  methodName: "SQLQuery",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SQLQueryRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SQLQueryResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QuerySimpleSQLQueryDesc: UnaryMethodDefinitionish = {
  methodName: "SimpleSQLQuery",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SQLQueryRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SimpleSQLQueryResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QuerySQLListTablesDesc: UnaryMethodDefinitionish = {
  methodName: "SQLListTables",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SQLListTablesRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SQLListTablesResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QuerySQLListDatabasesDesc: UnaryMethodDefinitionish = {
  methodName: "SQLListDatabases",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return SQLListDatabasesRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...SQLListDatabasesResponse.decode(data),
        toObject() {
          return this;
        },
      };
    },
  } as any,
};

export const QueryShowCreateTableDesc: UnaryMethodDefinitionish = {
  methodName: "ShowCreateTable",
  service: QueryDesc,
  requestStream: false,
  responseStream: false,
  requestType: {
    serializeBinary() {
      return ShowCreateTableRequest.encode(this).finish();
    },
  } as any,
  responseType: {
    deserializeBinary(data: Uint8Array) {
      return {
        ...ShowCreateTableResponse.decode(data),
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
