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

public struct GetUserResponse : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static void ValidateVersion() { FlatBufferConstants.FLATBUFFERS_25_2_10(); }
  public static GetUserResponse GetRootAsGetUserResponse(ByteBuffer _bb) { return GetRootAsGetUserResponse(_bb, new GetUserResponse()); }
  public static GetUserResponse GetRootAsGetUserResponse(ByteBuffer _bb, GetUserResponse obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p = new Table(_i, _bb); }
  public GetUserResponse __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public services.User? User { get { int o = __p.__offset(4); return o != 0 ? (services.User?)(new services.User()).__assign(__p.__indirect(o + __p.bb_pos), __p.bb) : null; } }
  public bool Found { get { int o = __p.__offset(6); return o != 0 ? 0!=__p.bb.Get(o + __p.bb_pos) : (bool)false; } }

  public static Offset<services.GetUserResponse> CreateGetUserResponse(FlatBufferBuilder builder,
      Offset<services.User> userOffset = default(Offset<services.User>),
      bool found = false) {
    builder.StartTable(2);
    GetUserResponse.AddUser(builder, userOffset);
    GetUserResponse.AddFound(builder, found);
    return GetUserResponse.EndGetUserResponse(builder);
  }

  public static void StartGetUserResponse(FlatBufferBuilder builder) { builder.StartTable(2); }
  public static void AddUser(FlatBufferBuilder builder, Offset<services.User> userOffset) { builder.AddOffset(0, userOffset.Value, 0); }
  public static void AddFound(FlatBufferBuilder builder, bool found) { builder.AddBool(1, found, false); }
  public static Offset<services.GetUserResponse> EndGetUserResponse(FlatBufferBuilder builder) {
    int o = builder.EndTable();
    return new Offset<services.GetUserResponse>(o);
  }
}


static public class GetUserResponseVerify
{
  static public bool Verify(Google.FlatBuffers.Verifier verifier, uint tablePos)
  {
    return verifier.VerifyTableStart(tablePos)
      && verifier.VerifyTable(tablePos, 4 /*User*/, services.UserVerify.Verify, false)
      && verifier.VerifyField(tablePos, 6 /*Found*/, 1 /*bool*/, 1, false)
      && verifier.VerifyTableEnd(tablePos);
  }
}

}
