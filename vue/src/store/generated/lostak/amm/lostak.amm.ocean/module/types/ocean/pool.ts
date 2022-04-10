/* eslint-disable */
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "lostak.amm.ocean";

export interface Pool {
  index: string;
  tokenA: Coin | undefined;
  tokenB: Coin | undefined;
  shares: Coin | undefined;
  swapFee: string;
}

const basePool: object = { index: "", swapFee: "" };

export const Pool = {
  encode(message: Pool, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.tokenA !== undefined) {
      Coin.encode(message.tokenA, writer.uint32(18).fork()).ldelim();
    }
    if (message.tokenB !== undefined) {
      Coin.encode(message.tokenB, writer.uint32(26).fork()).ldelim();
    }
    if (message.shares !== undefined) {
      Coin.encode(message.shares, writer.uint32(34).fork()).ldelim();
    }
    if (message.swapFee !== "") {
      writer.uint32(42).string(message.swapFee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePool } as Pool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.tokenA = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.tokenB = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.shares = Coin.decode(reader, reader.uint32());
          break;
        case 5:
          message.swapFee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pool {
    const message = { ...basePool } as Pool;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = Coin.fromJSON(object.tokenA);
    } else {
      message.tokenA = undefined;
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = Coin.fromJSON(object.tokenB);
    } else {
      message.tokenB = undefined;
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    if (object.swapFee !== undefined && object.swapFee !== null) {
      message.swapFee = String(object.swapFee);
    } else {
      message.swapFee = "";
    }
    return message;
  },

  toJSON(message: Pool): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.tokenA !== undefined &&
      (obj.tokenA = message.tokenA ? Coin.toJSON(message.tokenA) : undefined);
    message.tokenB !== undefined &&
      (obj.tokenB = message.tokenB ? Coin.toJSON(message.tokenB) : undefined);
    message.shares !== undefined &&
      (obj.shares = message.shares ? Coin.toJSON(message.shares) : undefined);
    message.swapFee !== undefined && (obj.swapFee = message.swapFee);
    return obj;
  },

  fromPartial(object: DeepPartial<Pool>): Pool {
    const message = { ...basePool } as Pool;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.tokenA !== undefined && object.tokenA !== null) {
      message.tokenA = Coin.fromPartial(object.tokenA);
    } else {
      message.tokenA = undefined;
    }
    if (object.tokenB !== undefined && object.tokenB !== null) {
      message.tokenB = Coin.fromPartial(object.tokenB);
    } else {
      message.tokenB = undefined;
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = Coin.fromPartial(object.shares);
    } else {
      message.shares = undefined;
    }
    if (object.swapFee !== undefined && object.swapFee !== null) {
      message.swapFee = object.swapFee;
    } else {
      message.swapFee = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
