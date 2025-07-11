// DO NOT EDIT!
// go generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10

package services

import (
	flatbuffers "github.com/google/flatbuffers/go"

	common "example.com/buffman/fb/common"
)

type SystemMetricsRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsSystemMetricsRequest(buf []byte, offset flatbuffers.UOffsetT) *SystemMetricsRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SystemMetricsRequest{}
	x.Init(buf, n+offset)
	return x
}

func FinishSystemMetricsRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsSystemMetricsRequest(buf []byte, offset flatbuffers.UOffsetT) *SystemMetricsRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SystemMetricsRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedSystemMetricsRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *SystemMetricsRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SystemMetricsRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SystemMetricsRequest) StartDate(obj *common.Timestamp) *common.Timestamp {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(common.Timestamp)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SystemMetricsRequest) EndDate(obj *common.Timestamp) *common.Timestamp {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(common.Timestamp)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func SystemMetricsRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SystemMetricsRequestAddStartDate(builder *flatbuffers.Builder, startDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(startDate), 0)
}
func SystemMetricsRequestAddEndDate(builder *flatbuffers.Builder, endDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(endDate), 0)
}
func SystemMetricsRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
