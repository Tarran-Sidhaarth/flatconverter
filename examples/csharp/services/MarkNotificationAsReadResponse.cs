// DO NOT EDIT!
// csharp generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10
//  automatically generated by the FlatBuffers compiler, do not modify
// </auto-generated>

namespace services
{

using global::System;
using global::System.Collections.Generic;
using global::Google.FlatBuffers;

public struct MarkNotificationAsReadResponse : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static void ValidateVersion() { FlatBufferConstants.FLATBUFFERS_25_2_10(); }
  public static MarkNotificationAsReadResponse GetRootAsMarkNotificationAsReadResponse(ByteBuffer _bb) { return GetRootAsMarkNotificationAsReadResponse(_bb, new MarkNotificationAsReadResponse()); }
  public static MarkNotificationAsReadResponse GetRootAsMarkNotificationAsReadResponse(ByteBuffer _bb, MarkNotificationAsReadResponse obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p = new Table(_i, _bb); }
  public MarkNotificationAsReadResponse __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public bool Success { get { int o = __p.__offset(4); return o != 0 ? 0!=__p.bb.Get(o + __p.bb_pos) : (bool)false; } }
  public string Message { get { int o = __p.__offset(6); return o != 0 ? __p.__string(o + __p.bb_pos) : null; } }
#if ENABLE_SPAN_T
  public Span<byte> GetMessageBytes() { return __p.__vector_as_span<byte>(6, 1); }
#else
  public ArraySegment<byte>? GetMessageBytes() { return __p.__vector_as_arraysegment(6); }
#endif
  public byte[] GetMessageArray() { return __p.__vector_as_array<byte>(6); }

  public static Offset<services.MarkNotificationAsReadResponse> CreateMarkNotificationAsReadResponse(FlatBufferBuilder builder,
      bool success = false,
      StringOffset messageOffset = default(StringOffset)) {
    builder.StartTable(2);
    MarkNotificationAsReadResponse.AddMessage(builder, messageOffset);
    MarkNotificationAsReadResponse.AddSuccess(builder, success);
    return MarkNotificationAsReadResponse.EndMarkNotificationAsReadResponse(builder);
  }

  public static void StartMarkNotificationAsReadResponse(FlatBufferBuilder builder) { builder.StartTable(2); }
  public static void AddSuccess(FlatBufferBuilder builder, bool success) { builder.AddBool(0, success, false); }
  public static void AddMessage(FlatBufferBuilder builder, StringOffset messageOffset) { builder.AddOffset(1, messageOffset.Value, 0); }
  public static Offset<services.MarkNotificationAsReadResponse> EndMarkNotificationAsReadResponse(FlatBufferBuilder builder) {
    int o = builder.EndTable();
    return new Offset<services.MarkNotificationAsReadResponse>(o);
  }
}


static public class MarkNotificationAsReadResponseVerify
{
  static public bool Verify(Google.FlatBuffers.Verifier verifier, uint tablePos)
  {
    return verifier.VerifyTableStart(tablePos)
      && verifier.VerifyField(tablePos, 4 /*Success*/, 1 /*bool*/, 1, false)
      && verifier.VerifyString(tablePos, 6 /*Message*/, false)
      && verifier.VerifyTableEnd(tablePos);
  }
}

}
