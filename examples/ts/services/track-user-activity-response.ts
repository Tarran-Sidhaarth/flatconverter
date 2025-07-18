// DO NOT EDIT!
// ts generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10

/* eslint-disable @typescript-eslint/no-unused-vars, @typescript-eslint/no-explicit-any, @typescript-eslint/no-non-null-assertion */

import * as flatbuffers from 'flatbuffers';

export class TrackUserActivityResponse {
  bb: flatbuffers.ByteBuffer|null = null;
  bb_pos = 0;
  __init(i:number, bb:flatbuffers.ByteBuffer):TrackUserActivityResponse {
  this.bb_pos = i;
  this.bb = bb;
  return this;
}

static getRootAsTrackUserActivityResponse(bb:flatbuffers.ByteBuffer, obj?:TrackUserActivityResponse):TrackUserActivityResponse {
  return (obj || new TrackUserActivityResponse()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

static getSizePrefixedRootAsTrackUserActivityResponse(bb:flatbuffers.ByteBuffer, obj?:TrackUserActivityResponse):TrackUserActivityResponse {
  bb.setPosition(bb.position() + flatbuffers.SIZE_PREFIX_LENGTH);
  return (obj || new TrackUserActivityResponse()).__init(bb.readInt32(bb.position()) + bb.position(), bb);
}

success():boolean {
  const offset = this.bb!.__offset(this.bb_pos, 4);
  return offset ? !!this.bb!.readInt8(this.bb_pos + offset) : false;
}

message():string|null
message(optionalEncoding:flatbuffers.Encoding):string|Uint8Array|null
message(optionalEncoding?:any):string|Uint8Array|null {
  const offset = this.bb!.__offset(this.bb_pos, 6);
  return offset ? this.bb!.__string(this.bb_pos + offset, optionalEncoding) : null;
}

static startTrackUserActivityResponse(builder:flatbuffers.Builder) {
  builder.startObject(2);
}

static addSuccess(builder:flatbuffers.Builder, success:boolean) {
  builder.addFieldInt8(0, +success, +false);
}

static addMessage(builder:flatbuffers.Builder, messageOffset:flatbuffers.Offset) {
  builder.addFieldOffset(1, messageOffset, 0);
}

static endTrackUserActivityResponse(builder:flatbuffers.Builder):flatbuffers.Offset {
  const offset = builder.endObject();
  return offset;
}

static createTrackUserActivityResponse(builder:flatbuffers.Builder, success:boolean, messageOffset:flatbuffers.Offset):flatbuffers.Offset {
  TrackUserActivityResponse.startTrackUserActivityResponse(builder);
  TrackUserActivityResponse.addSuccess(builder, success);
  TrackUserActivityResponse.addMessage(builder, messageOffset);
  return TrackUserActivityResponse.endTrackUserActivityResponse(builder);
}
}
