// DO NOT EDIT!
// go generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10

package services

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CreateUserResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsCreateUserResponse(buf []byte, offset flatbuffers.UOffsetT) *CreateUserResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CreateUserResponse{}
	x.Init(buf, n+offset)
	return x
}

func FinishCreateUserResponseBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsCreateUserResponse(buf []byte, offset flatbuffers.UOffsetT) *CreateUserResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CreateUserResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedCreateUserResponseBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *CreateUserResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CreateUserResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CreateUserResponse) User(obj *User) *User {
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

func (rcv *CreateUserResponse) Success() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *CreateUserResponse) MutateSuccess(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *CreateUserResponse) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CreateUserResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CreateUserResponseAddUser(builder *flatbuffers.Builder, user flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(user), 0)
}
func CreateUserResponseAddSuccess(builder *flatbuffers.Builder, success bool) {
	builder.PrependBoolSlot(1, success, false)
}
func CreateUserResponseAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(message), 0)
}
func CreateUserResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
