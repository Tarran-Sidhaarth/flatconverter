// DO NOT EDIT!
// csharp generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10
//  automatically generated by the FlatBuffers compiler, do not modify
// </auto-generated>

namespace common
{

using global::System;
using global::System.Collections.Generic;
using global::Google.FlatBuffers;

public struct Timestamp : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static void ValidateVersion() { FlatBufferConstants.FLATBUFFERS_25_2_10(); }
  public static Timestamp GetRootAsTimestamp(ByteBuffer _bb) { return GetRootAsTimestamp(_bb, new Timestamp()); }
  public static Timestamp GetRootAsTimestamp(ByteBuffer _bb, Timestamp obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p = new Table(_i, _bb); }
  public Timestamp __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public long Seconds { get { int o = __p.__offset(4); return o != 0 ? __p.bb.GetLong(o + __p.bb_pos) : (long)0; } }
  public int Nanos { get { int o = __p.__offset(6); return o != 0 ? __p.bb.GetInt(o + __p.bb_pos) : (int)0; } }

  public static Offset<common.Timestamp> CreateTimestamp(FlatBufferBuilder builder,
      long seconds = 0,
      int nanos = 0) {
    builder.StartTable(2);
    Timestamp.AddSeconds(builder, seconds);
    Timestamp.AddNanos(builder, nanos);
    return Timestamp.EndTimestamp(builder);
  }

  public static void StartTimestamp(FlatBufferBuilder builder) { builder.StartTable(2); }
  public static void AddSeconds(FlatBufferBuilder builder, long seconds) { builder.AddLong(0, seconds, 0); }
  public static void AddNanos(FlatBufferBuilder builder, int nanos) { builder.AddInt(1, nanos, 0); }
  public static Offset<common.Timestamp> EndTimestamp(FlatBufferBuilder builder) {
    int o = builder.EndTable();
    return new Offset<common.Timestamp>(o);
  }
}


static public class TimestampVerify
{
  static public bool Verify(Google.FlatBuffers.Verifier verifier, uint tablePos)
  {
    return verifier.VerifyTableStart(tablePos)
      && verifier.VerifyField(tablePos, 4 /*Seconds*/, 8 /*long*/, 8, false)
      && verifier.VerifyField(tablePos, 6 /*Nanos*/, 4 /*int*/, 4, false)
      && verifier.VerifyTableEnd(tablePos);
  }
}

}
