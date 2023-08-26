/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "blockved.glitterchain.index";

export enum ColumnValueType {
  InvalidColumn = 0,
  IntColumn = 1,
  UintColumn = 2,
  FloatColumn = 3,
  BoolColumn = 4,
  StringColumn = 5,
  BytesColumn = 6,
  UNRECOGNIZED = -1,
}

export function columnValueTypeFromJSON(object: any): ColumnValueType {
  switch (object) {
    case 0:
    case "InvalidColumn":
      return ColumnValueType.InvalidColumn;
    case 1:
    case "IntColumn":
      return ColumnValueType.IntColumn;
    case 2:
    case "UintColumn":
      return ColumnValueType.UintColumn;
    case 3:
    case "FloatColumn":
      return ColumnValueType.FloatColumn;
    case 4:
    case "BoolColumn":
      return ColumnValueType.BoolColumn;
    case 5:
    case "StringColumn":
      return ColumnValueType.StringColumn;
    case 6:
    case "BytesColumn":
      return ColumnValueType.BytesColumn;
    case -1:
    case "UNRECOGNIZED":
    default:
      return ColumnValueType.UNRECOGNIZED;
  }
}

export function columnValueTypeToJSON(object: ColumnValueType): string {
  switch (object) {
    case ColumnValueType.InvalidColumn:
      return "InvalidColumn";
    case ColumnValueType.IntColumn:
      return "IntColumn";
    case ColumnValueType.UintColumn:
      return "UintColumn";
    case ColumnValueType.FloatColumn:
      return "FloatColumn";
    case ColumnValueType.BoolColumn:
      return "BoolColumn";
    case ColumnValueType.StringColumn:
      return "StringColumn";
    case ColumnValueType.BytesColumn:
      return "BytesColumn";
    default:
      return "UNKNOWN";
  }
}

export interface Argument {
  /** variable type */
  type: Argument_VarType;
  /** variable value */
  value: string;
}

export enum Argument_VarType {
  INT = 0,
  UINT = 1,
  FLOAT = 2,
  BOOL = 3,
  STRING = 4,
  BYTES = 5,
  UNRECOGNIZED = -1,
}

export function argument_VarTypeFromJSON(object: any): Argument_VarType {
  switch (object) {
    case 0:
    case "INT":
      return Argument_VarType.INT;
    case 1:
    case "UINT":
      return Argument_VarType.UINT;
    case 2:
    case "FLOAT":
      return Argument_VarType.FLOAT;
    case 3:
    case "BOOL":
      return Argument_VarType.BOOL;
    case 4:
    case "STRING":
      return Argument_VarType.STRING;
    case 5:
    case "BYTES":
      return Argument_VarType.BYTES;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Argument_VarType.UNRECOGNIZED;
  }
}

export function argument_VarTypeToJSON(object: Argument_VarType): string {
  switch (object) {
    case Argument_VarType.INT:
      return "INT";
    case Argument_VarType.UINT:
      return "UINT";
    case Argument_VarType.FLOAT:
      return "FLOAT";
    case Argument_VarType.BOOL:
      return "BOOL";
    case Argument_VarType.STRING:
      return "STRING";
    case Argument_VarType.BYTES:
      return "BYTES";
    default:
      return "UNKNOWN";
  }
}

export interface RowData {
  columns: string[];
}

export interface ResultSet {
  id: string;
  columnDefs: ColumnDef[];
  rows: RowData[];
}

export interface ColumnDef {
  columnName: string;
  columnType: string;
  comment: string;
  columnValueType: ColumnValueType;
}

export interface IndexDef {
  indexName: string;
  indexType: string;
  columns: ColumnDef[];
  isPrimaryKey: boolean;
  parser: string;
  comment: string;
}

export interface TableInfo {
  tableName: string;
  dbName: string;
  engine: string;
  columns: ColumnDef[];
  indexes: IndexDef[];
  createStatement: string;
  comment: string;
  creator: string;
  members: TableMember[];
}

export interface TableMember {
  uid: string;
  role: string;
}

export interface DatabaseInfo {
  dbName: string;
  creator: string;
  members: TableMember[];
}

const baseArgument: object = { type: 0, value: "" };

export const Argument = {
  encode(message: Argument, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Argument {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseArgument } as Argument;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.int32() as any;
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Argument {
    const message = { ...baseArgument } as Argument;
    if (object.type !== undefined && object.type !== null) {
      message.type = argument_VarTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: Argument): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = argument_VarTypeToJSON(message.type));
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<Argument>): Argument {
    const message = { ...baseArgument } as Argument;
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseRowData: object = { columns: "" };

export const RowData = {
  encode(message: RowData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.columns) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RowData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRowData } as RowData;
    message.columns = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.columns.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RowData {
    const message = { ...baseRowData } as RowData;
    message.columns = [];
    if (object.columns !== undefined && object.columns !== null) {
      for (const e of object.columns) {
        message.columns.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: RowData): unknown {
    const obj: any = {};
    if (message.columns) {
      obj.columns = message.columns.map((e) => e);
    } else {
      obj.columns = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<RowData>): RowData {
    const message = { ...baseRowData } as RowData;
    message.columns = [];
    if (object.columns !== undefined && object.columns !== null) {
      for (const e of object.columns) {
        message.columns.push(e);
      }
    }
    return message;
  },
};

const baseResultSet: object = { id: "" };

export const ResultSet = {
  encode(message: ResultSet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    for (const v of message.columnDefs) {
      ColumnDef.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.rows) {
      RowData.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResultSet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseResultSet } as ResultSet;
    message.columnDefs = [];
    message.rows = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.columnDefs.push(ColumnDef.decode(reader, reader.uint32()));
          break;
        case 3:
          message.rows.push(RowData.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ResultSet {
    const message = { ...baseResultSet } as ResultSet;
    message.columnDefs = [];
    message.rows = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.columnDefs !== undefined && object.columnDefs !== null) {
      for (const e of object.columnDefs) {
        message.columnDefs.push(ColumnDef.fromJSON(e));
      }
    }
    if (object.rows !== undefined && object.rows !== null) {
      for (const e of object.rows) {
        message.rows.push(RowData.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ResultSet): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    if (message.columnDefs) {
      obj.columnDefs = message.columnDefs.map((e) => (e ? ColumnDef.toJSON(e) : undefined));
    } else {
      obj.columnDefs = [];
    }
    if (message.rows) {
      obj.rows = message.rows.map((e) => (e ? RowData.toJSON(e) : undefined));
    } else {
      obj.rows = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ResultSet>): ResultSet {
    const message = { ...baseResultSet } as ResultSet;
    message.columnDefs = [];
    message.rows = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.columnDefs !== undefined && object.columnDefs !== null) {
      for (const e of object.columnDefs) {
        message.columnDefs.push(ColumnDef.fromPartial(e));
      }
    }
    if (object.rows !== undefined && object.rows !== null) {
      for (const e of object.rows) {
        message.rows.push(RowData.fromPartial(e));
      }
    }
    return message;
  },
};

const baseColumnDef: object = { columnName: "", columnType: "", comment: "", columnValueType: 0 };

export const ColumnDef = {
  encode(message: ColumnDef, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.columnName !== "") {
      writer.uint32(10).string(message.columnName);
    }
    if (message.columnType !== "") {
      writer.uint32(18).string(message.columnType);
    }
    if (message.comment !== "") {
      writer.uint32(26).string(message.comment);
    }
    if (message.columnValueType !== 0) {
      writer.uint32(32).int32(message.columnValueType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ColumnDef {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseColumnDef } as ColumnDef;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.columnName = reader.string();
          break;
        case 2:
          message.columnType = reader.string();
          break;
        case 3:
          message.comment = reader.string();
          break;
        case 4:
          message.columnValueType = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ColumnDef {
    const message = { ...baseColumnDef } as ColumnDef;
    if (object.columnName !== undefined && object.columnName !== null) {
      message.columnName = String(object.columnName);
    } else {
      message.columnName = "";
    }
    if (object.columnType !== undefined && object.columnType !== null) {
      message.columnType = String(object.columnType);
    } else {
      message.columnType = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = String(object.comment);
    } else {
      message.comment = "";
    }
    if (object.columnValueType !== undefined && object.columnValueType !== null) {
      message.columnValueType = columnValueTypeFromJSON(object.columnValueType);
    } else {
      message.columnValueType = 0;
    }
    return message;
  },

  toJSON(message: ColumnDef): unknown {
    const obj: any = {};
    message.columnName !== undefined && (obj.columnName = message.columnName);
    message.columnType !== undefined && (obj.columnType = message.columnType);
    message.comment !== undefined && (obj.comment = message.comment);
    message.columnValueType !== undefined &&
      (obj.columnValueType = columnValueTypeToJSON(message.columnValueType));
    return obj;
  },

  fromPartial(object: DeepPartial<ColumnDef>): ColumnDef {
    const message = { ...baseColumnDef } as ColumnDef;
    if (object.columnName !== undefined && object.columnName !== null) {
      message.columnName = object.columnName;
    } else {
      message.columnName = "";
    }
    if (object.columnType !== undefined && object.columnType !== null) {
      message.columnType = object.columnType;
    } else {
      message.columnType = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = object.comment;
    } else {
      message.comment = "";
    }
    if (object.columnValueType !== undefined && object.columnValueType !== null) {
      message.columnValueType = object.columnValueType;
    } else {
      message.columnValueType = 0;
    }
    return message;
  },
};

const baseIndexDef: object = { indexName: "", indexType: "", isPrimaryKey: false, parser: "", comment: "" };

export const IndexDef = {
  encode(message: IndexDef, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.indexName !== "") {
      writer.uint32(10).string(message.indexName);
    }
    if (message.indexType !== "") {
      writer.uint32(18).string(message.indexType);
    }
    for (const v of message.columns) {
      ColumnDef.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.isPrimaryKey === true) {
      writer.uint32(40).bool(message.isPrimaryKey);
    }
    if (message.parser !== "") {
      writer.uint32(50).string(message.parser);
    }
    if (message.comment !== "") {
      writer.uint32(58).string(message.comment);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IndexDef {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseIndexDef } as IndexDef;
    message.columns = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.indexName = reader.string();
          break;
        case 2:
          message.indexType = reader.string();
          break;
        case 4:
          message.columns.push(ColumnDef.decode(reader, reader.uint32()));
          break;
        case 5:
          message.isPrimaryKey = reader.bool();
          break;
        case 6:
          message.parser = reader.string();
          break;
        case 7:
          message.comment = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IndexDef {
    const message = { ...baseIndexDef } as IndexDef;
    message.columns = [];
    if (object.indexName !== undefined && object.indexName !== null) {
      message.indexName = String(object.indexName);
    } else {
      message.indexName = "";
    }
    if (object.indexType !== undefined && object.indexType !== null) {
      message.indexType = String(object.indexType);
    } else {
      message.indexType = "";
    }
    if (object.columns !== undefined && object.columns !== null) {
      for (const e of object.columns) {
        message.columns.push(ColumnDef.fromJSON(e));
      }
    }
    if (object.isPrimaryKey !== undefined && object.isPrimaryKey !== null) {
      message.isPrimaryKey = Boolean(object.isPrimaryKey);
    } else {
      message.isPrimaryKey = false;
    }
    if (object.parser !== undefined && object.parser !== null) {
      message.parser = String(object.parser);
    } else {
      message.parser = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = String(object.comment);
    } else {
      message.comment = "";
    }
    return message;
  },

  toJSON(message: IndexDef): unknown {
    const obj: any = {};
    message.indexName !== undefined && (obj.indexName = message.indexName);
    message.indexType !== undefined && (obj.indexType = message.indexType);
    if (message.columns) {
      obj.columns = message.columns.map((e) => (e ? ColumnDef.toJSON(e) : undefined));
    } else {
      obj.columns = [];
    }
    message.isPrimaryKey !== undefined && (obj.isPrimaryKey = message.isPrimaryKey);
    message.parser !== undefined && (obj.parser = message.parser);
    message.comment !== undefined && (obj.comment = message.comment);
    return obj;
  },

  fromPartial(object: DeepPartial<IndexDef>): IndexDef {
    const message = { ...baseIndexDef } as IndexDef;
    message.columns = [];
    if (object.indexName !== undefined && object.indexName !== null) {
      message.indexName = object.indexName;
    } else {
      message.indexName = "";
    }
    if (object.indexType !== undefined && object.indexType !== null) {
      message.indexType = object.indexType;
    } else {
      message.indexType = "";
    }
    if (object.columns !== undefined && object.columns !== null) {
      for (const e of object.columns) {
        message.columns.push(ColumnDef.fromPartial(e));
      }
    }
    if (object.isPrimaryKey !== undefined && object.isPrimaryKey !== null) {
      message.isPrimaryKey = object.isPrimaryKey;
    } else {
      message.isPrimaryKey = false;
    }
    if (object.parser !== undefined && object.parser !== null) {
      message.parser = object.parser;
    } else {
      message.parser = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = object.comment;
    } else {
      message.comment = "";
    }
    return message;
  },
};

const baseTableInfo: object = {
  tableName: "",
  dbName: "",
  engine: "",
  createStatement: "",
  comment: "",
  creator: "",
};

export const TableInfo = {
  encode(message: TableInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tableName !== "") {
      writer.uint32(10).string(message.tableName);
    }
    if (message.dbName !== "") {
      writer.uint32(18).string(message.dbName);
    }
    if (message.engine !== "") {
      writer.uint32(26).string(message.engine);
    }
    for (const v of message.columns) {
      ColumnDef.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.indexes) {
      IndexDef.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.createStatement !== "") {
      writer.uint32(50).string(message.createStatement);
    }
    if (message.comment !== "") {
      writer.uint32(58).string(message.comment);
    }
    if (message.creator !== "") {
      writer.uint32(66).string(message.creator);
    }
    for (const v of message.members) {
      TableMember.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TableInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTableInfo } as TableInfo;
    message.columns = [];
    message.indexes = [];
    message.members = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tableName = reader.string();
          break;
        case 2:
          message.dbName = reader.string();
          break;
        case 3:
          message.engine = reader.string();
          break;
        case 4:
          message.columns.push(ColumnDef.decode(reader, reader.uint32()));
          break;
        case 5:
          message.indexes.push(IndexDef.decode(reader, reader.uint32()));
          break;
        case 6:
          message.createStatement = reader.string();
          break;
        case 7:
          message.comment = reader.string();
          break;
        case 8:
          message.creator = reader.string();
          break;
        case 9:
          message.members.push(TableMember.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TableInfo {
    const message = { ...baseTableInfo } as TableInfo;
    message.columns = [];
    message.indexes = [];
    message.members = [];
    if (object.tableName !== undefined && object.tableName !== null) {
      message.tableName = String(object.tableName);
    } else {
      message.tableName = "";
    }
    if (object.dbName !== undefined && object.dbName !== null) {
      message.dbName = String(object.dbName);
    } else {
      message.dbName = "";
    }
    if (object.engine !== undefined && object.engine !== null) {
      message.engine = String(object.engine);
    } else {
      message.engine = "";
    }
    if (object.columns !== undefined && object.columns !== null) {
      for (const e of object.columns) {
        message.columns.push(ColumnDef.fromJSON(e));
      }
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(IndexDef.fromJSON(e));
      }
    }
    if (object.createStatement !== undefined && object.createStatement !== null) {
      message.createStatement = String(object.createStatement);
    } else {
      message.createStatement = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = String(object.comment);
    } else {
      message.comment = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(TableMember.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: TableInfo): unknown {
    const obj: any = {};
    message.tableName !== undefined && (obj.tableName = message.tableName);
    message.dbName !== undefined && (obj.dbName = message.dbName);
    message.engine !== undefined && (obj.engine = message.engine);
    if (message.columns) {
      obj.columns = message.columns.map((e) => (e ? ColumnDef.toJSON(e) : undefined));
    } else {
      obj.columns = [];
    }
    if (message.indexes) {
      obj.indexes = message.indexes.map((e) => (e ? IndexDef.toJSON(e) : undefined));
    } else {
      obj.indexes = [];
    }
    message.createStatement !== undefined && (obj.createStatement = message.createStatement);
    message.comment !== undefined && (obj.comment = message.comment);
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.members) {
      obj.members = message.members.map((e) => (e ? TableMember.toJSON(e) : undefined));
    } else {
      obj.members = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<TableInfo>): TableInfo {
    const message = { ...baseTableInfo } as TableInfo;
    message.columns = [];
    message.indexes = [];
    message.members = [];
    if (object.tableName !== undefined && object.tableName !== null) {
      message.tableName = object.tableName;
    } else {
      message.tableName = "";
    }
    if (object.dbName !== undefined && object.dbName !== null) {
      message.dbName = object.dbName;
    } else {
      message.dbName = "";
    }
    if (object.engine !== undefined && object.engine !== null) {
      message.engine = object.engine;
    } else {
      message.engine = "";
    }
    if (object.columns !== undefined && object.columns !== null) {
      for (const e of object.columns) {
        message.columns.push(ColumnDef.fromPartial(e));
      }
    }
    if (object.indexes !== undefined && object.indexes !== null) {
      for (const e of object.indexes) {
        message.indexes.push(IndexDef.fromPartial(e));
      }
    }
    if (object.createStatement !== undefined && object.createStatement !== null) {
      message.createStatement = object.createStatement;
    } else {
      message.createStatement = "";
    }
    if (object.comment !== undefined && object.comment !== null) {
      message.comment = object.comment;
    } else {
      message.comment = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(TableMember.fromPartial(e));
      }
    }
    return message;
  },
};

const baseTableMember: object = { uid: "", role: "" };

export const TableMember = {
  encode(message: TableMember, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uid !== "") {
      writer.uint32(10).string(message.uid);
    }
    if (message.role !== "") {
      writer.uint32(18).string(message.role);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TableMember {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTableMember } as TableMember;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = reader.string();
          break;
        case 2:
          message.role = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TableMember {
    const message = { ...baseTableMember } as TableMember;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    if (object.role !== undefined && object.role !== null) {
      message.role = String(object.role);
    } else {
      message.role = "";
    }
    return message;
  },

  toJSON(message: TableMember): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.role !== undefined && (obj.role = message.role);
    return obj;
  },

  fromPartial(object: DeepPartial<TableMember>): TableMember {
    const message = { ...baseTableMember } as TableMember;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    if (object.role !== undefined && object.role !== null) {
      message.role = object.role;
    } else {
      message.role = "";
    }
    return message;
  },
};

const baseDatabaseInfo: object = { dbName: "", creator: "" };

export const DatabaseInfo = {
  encode(message: DatabaseInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dbName !== "") {
      writer.uint32(10).string(message.dbName);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    for (const v of message.members) {
      TableMember.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DatabaseInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDatabaseInfo } as DatabaseInfo;
    message.members = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dbName = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.members.push(TableMember.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DatabaseInfo {
    const message = { ...baseDatabaseInfo } as DatabaseInfo;
    message.members = [];
    if (object.dbName !== undefined && object.dbName !== null) {
      message.dbName = String(object.dbName);
    } else {
      message.dbName = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(TableMember.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: DatabaseInfo): unknown {
    const obj: any = {};
    message.dbName !== undefined && (obj.dbName = message.dbName);
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.members) {
      obj.members = message.members.map((e) => (e ? TableMember.toJSON(e) : undefined));
    } else {
      obj.members = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<DatabaseInfo>): DatabaseInfo {
    const message = { ...baseDatabaseInfo } as DatabaseInfo;
    message.members = [];
    if (object.dbName !== undefined && object.dbName !== null) {
      message.dbName = object.dbName;
    } else {
      message.dbName = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.members !== undefined && object.members !== null) {
      for (const e of object.members) {
        message.members.push(TableMember.fromPartial(e));
      }
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
