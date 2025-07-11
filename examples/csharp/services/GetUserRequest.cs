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

public struct GetUserRequest : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static void ValidateVersion() { FlatBufferConstants.FLATBUFFERS_25_2_10(); }
  public static GetUserRequest GetRootAsGetUserRequest(ByteBuffer _bb) { return GetRootAsGetUserRequest(_bb, new GetUserRequest()); }
  public static GetUserRequest GetRootAsGetUserRequest(ByteBuffer _bb, GetUserRequest obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p = new Table(_i, _bb); }
  public GetUserRequest __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public string UserId { get { int o = __p.__offset(4); return o != 0 ? __p.__string(o + __p.bb_pos) : null; } }
#if ENABLE_SPAN_T
  public Span<byte> GetUserIdBytes() { return __p.__vector_as_span<byte>(4, 1); }
#else
  public ArraySegment<byte>? GetUserIdBytes() { return __p.__vector_as_arraysegment(4); }
#endif
  public byte[] GetUserIdArray() { return __p.__vector_as_array<byte>(4); }

  public static Offset<services.GetUserRequest> CreateGetUserRequest(FlatBufferBuilder builder,
      StringOffset user_idOffset = default(StringOffset)) {
    builder.StartTable(1);
    GetUserRequest.AddUserId(builder, user_idOffset);
    return GetUserRequest.EndGetUserRequest(builder);
  }

  public static void StartGetUserRequest(FlatBufferBuilder builder) { builder.StartTable(1); }
  public static void AddUserId(FlatBufferBuilder builder, StringOffset userIdOffset) { builder.AddOffset(0, userIdOffset.Value, 0); }
  public static Offset<services.GetUserRequest> EndGetUserRequest(FlatBufferBuilder builder) {
    int o = builder.EndTable();
    return new Offset<services.GetUserRequest>(o);
  }
}


static public class GetUserRequestVerify
{
  static public bool Verify(Google.FlatBuffers.Verifier verifier, uint tablePos)
  {
    return verifier.VerifyTableStart(tablePos)
      && verifier.VerifyString(tablePos, 4 /*UserId*/, false)
      && verifier.VerifyTableEnd(tablePos);
  }
}

}
