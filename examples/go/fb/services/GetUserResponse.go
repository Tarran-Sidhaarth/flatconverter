// DO NOT EDIT!
// go generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10

package services

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GetUserResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsGetUserResponse(buf []byte, offset flatbuffers.UOffsetT) *GetUserResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetUserResponse{}
	x.Init(buf, n+offset)
	return x
}

func FinishGetUserResponseBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsGetUserResponse(buf []byte, offset flatbuffers.UOffsetT) *GetUserResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GetUserResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedGetUserResponseBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *GetUserResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GetUserResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GetUserResponse) User(obj *User) *User {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(User)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *GetUserResponse) Found() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *GetUserResponse) MutateFound(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func GetUserResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func GetUserResponseAddUser(builder *flatbuffers.Builder, user flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(user), 0)
}
func GetUserResponseAddFound(builder *flatbuffers.Builder, found bool) {
	builder.PrependBoolSlot(1, found, false)
}
func GetUserResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
