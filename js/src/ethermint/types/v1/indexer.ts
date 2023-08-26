/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "ethermint.types.v1";

/** TxResult is the value stored in eth tx indexer */
export interface TxResult {
  /** the block height */
  height: Long;
  /** cosmos tx index */
  txIndex: number;
  /** the msg index in a batch tx */
  msgIndex: number;
  /**
   * eth tx index, the index in the list of valid eth tx in the block,
   * aka. the transaction list returned by eth_getBlock api.
   */
  ethTxIndex: number;
  /** if the eth tx is failed */
  failed: boolean;
  /**
   * gas used by tx, if exceeds block gas limit,
   * it's set to gas limit which is what's actually deducted by ante handler.
   */
  gasUsed: Long;
  /** the cumulative gas used within current batch tx */
  cumulativeGasUsed: Long;
}

const baseTxResult: object = {
  height: Long.ZERO,
  txIndex: 0,
  msgIndex: 0,
  ethTxIndex: 0,
  failed: false,
  gasUsed: Long.UZERO,
  cumulativeGasUsed: Long.UZERO,
};

export const TxResult = {
  encode(message: TxResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.height.isZero()) {
      writer.uint32(8).int64(message.height);
    }
    if (message.txIndex !== 0) {
      writer.uint32(16).uint32(message.txIndex);
    }
    if (message.msgIndex !== 0) {
      writer.uint32(24).uint32(message.msgIndex);
    }
    if (message.ethTxIndex !== 0) {
      writer.uint32(32).int32(message.ethTxIndex);
    }
    if (message.failed === true) {
      writer.uint32(40).bool(message.failed);
    }
    if (!message.gasUsed.isZero()) {
      writer.uint32(48).uint64(message.gasUsed);
    }
    if (!message.cumulativeGasUsed.isZero()) {
      writer.uint32(56).uint64(message.cumulativeGasUsed);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TxResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTxResult } as TxResult;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.height = reader.int64() as Long;
          break;
        case 2:
          message.txIndex = reader.uint32();
          break;
        case 3:
          message.msgIndex = reader.uint32();
          break;
        case 4:
          message.ethTxIndex = reader.int32();
          break;
        case 5:
          message.failed = reader.bool();
          break;
        case 6:
          message.gasUsed = reader.uint64() as Long;
          break;
        case 7:
          message.cumulativeGasUsed = reader.uint64() as Long;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TxResult {
    const message = { ...baseTxResult } as TxResult;
    if (object.height !== undefined && object.height !== null) {
      message.height = Long.fromString(object.height);
    } else {
      message.height = Long.ZERO;
    }
    if (object.txIndex !== undefined && object.txIndex !== null) {
      message.txIndex = Number(object.txIndex);
    } else {
      message.txIndex = 0;
    }
    if (object.msgIndex !== undefined && object.msgIndex !== null) {
      message.msgIndex = Number(object.msgIndex);
    } else {
      message.msgIndex = 0;
    }
    if (object.ethTxIndex !== undefined && object.ethTxIndex !== null) {
      message.ethTxIndex = Number(object.ethTxIndex);
    } else {
      message.ethTxIndex = 0;
    }
    if (object.failed !== undefined && object.failed !== null) {
      message.failed = Boolean(object.failed);
    } else {
      message.failed = false;
    }
    if (object.gasUsed !== undefined && object.gasUsed !== null) {
      message.gasUsed = Long.fromString(object.gasUsed);
    } else {
      message.gasUsed = Long.UZERO;
    }
    if (object.cumulativeGasUsed !== undefined && object.cumulativeGasUsed !== null) {
      message.cumulativeGasUsed = Long.fromString(object.cumulativeGasUsed);
    } else {
      message.cumulativeGasUsed = Long.UZERO;
    }
    return message;
  },

  toJSON(message: TxResult): unknown {
    const obj: any = {};
    message.height !== undefined && (obj.height = (message.height || Long.ZERO).toString());
    message.txIndex !== undefined && (obj.txIndex = message.txIndex);
    message.msgIndex !== undefined && (obj.msgIndex = message.msgIndex);
    message.ethTxIndex !== undefined && (obj.ethTxIndex = message.ethTxIndex);
    message.failed !== undefined && (obj.failed = message.failed);
    message.gasUsed !== undefined && (obj.gasUsed = (message.gasUsed || Long.UZERO).toString());
    message.cumulativeGasUsed !== undefined &&
      (obj.cumulativeGasUsed = (message.cumulativeGasUsed || Long.UZERO).toString());
    return obj;
  },

  fromPartial(object: DeepPartial<TxResult>): TxResult {
    const message = { ...baseTxResult } as TxResult;
    if (object.height !== undefined && object.height !== null) {
      message.height = object.height as Long;
    } else {
      message.height = Long.ZERO;
    }
    if (object.txIndex !== undefined && object.txIndex !== null) {
      message.txIndex = object.txIndex;
    } else {
      message.txIndex = 0;
    }
    if (object.msgIndex !== undefined && object.msgIndex !== null) {
      message.msgIndex = object.msgIndex;
    } else {
      message.msgIndex = 0;
    }
    if (object.ethTxIndex !== undefined && object.ethTxIndex !== null) {
      message.ethTxIndex = object.ethTxIndex;
    } else {
      message.ethTxIndex = 0;
    }
    if (object.failed !== undefined && object.failed !== null) {
      message.failed = object.failed;
    } else {
      message.failed = false;
    }
    if (object.gasUsed !== undefined && object.gasUsed !== null) {
      message.gasUsed = object.gasUsed as Long;
    } else {
      message.gasUsed = Long.UZERO;
    }
    if (object.cumulativeGasUsed !== undefined && object.cumulativeGasUsed !== null) {
      message.cumulativeGasUsed = object.cumulativeGasUsed as Long;
    } else {
      message.cumulativeGasUsed = Long.UZERO;
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
