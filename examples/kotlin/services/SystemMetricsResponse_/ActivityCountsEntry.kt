// DO NOT EDIT!
// kotlin generated by Buffman 💪
// Versions:
// 		Buffman: 1.0.0
// 		Flatc: v25.2.10

package services.SystemMetricsResponse_

import com.google.flatbuffers.BaseVector
import com.google.flatbuffers.BooleanVector
import com.google.flatbuffers.ByteVector
import com.google.flatbuffers.Constants
import com.google.flatbuffers.DoubleVector
import com.google.flatbuffers.FlatBufferBuilder
import com.google.flatbuffers.FloatVector
import com.google.flatbuffers.LongVector
import com.google.flatbuffers.StringVector
import com.google.flatbuffers.Struct
import com.google.flatbuffers.Table
import com.google.flatbuffers.UnionVector
import java.nio.ByteBuffer
import java.nio.ByteOrder
import kotlin.math.sign

@Suppress("unused")
class ActivityCountsEntry : Table() {

    fun __init(_i: Int, _bb: ByteBuffer)  {
        __reset(_i, _bb)
    }
    fun __assign(_i: Int, _bb: ByteBuffer) : ActivityCountsEntry {
        __init(_i, _bb)
        return this
    }
    val key : String?
        get() {
            val o = __offset(4)
            return if (o != 0) {
                __string(o + bb_pos)
            } else {
                null
            }
        }
    val keyAsByteBuffer : ByteBuffer get() = __vector_as_bytebuffer(4, 1)
    fun keyInByteBuffer(_bb: ByteBuffer) : ByteBuffer = __vector_in_bytebuffer(_bb, 4, 1)
    val value : Int
        get() {
            val o = __offset(6)
            return if(o != 0) bb.getInt(o + bb_pos) else 0
        }
    companion object {
        fun validateVersion() = Constants.FLATBUFFERS_25_2_10()
        fun getRootAsActivityCountsEntry(_bb: ByteBuffer): ActivityCountsEntry = getRootAsActivityCountsEntry(_bb, ActivityCountsEntry())
        fun getRootAsActivityCountsEntry(_bb: ByteBuffer, obj: ActivityCountsEntry): ActivityCountsEntry {
            _bb.order(ByteOrder.LITTLE_ENDIAN)
            return (obj.__assign(_bb.getInt(_bb.position()) + _bb.position(), _bb))
        }
        fun createActivityCountsEntry(builder: FlatBufferBuilder, keyOffset: Int, value: Int) : Int {
            builder.startTable(2)
            addValue(builder, value)
            addKey(builder, keyOffset)
            return endActivityCountsEntry(builder)
        }
        fun startActivityCountsEntry(builder: FlatBufferBuilder) = builder.startTable(2)
        fun addKey(builder: FlatBufferBuilder, key: Int) = builder.addOffset(0, key, 0)
        fun addValue(builder: FlatBufferBuilder, value: Int) = builder.addInt(1, value, 0)
        fun endActivityCountsEntry(builder: FlatBufferBuilder) : Int {
            val o = builder.endTable()
            return o
        }
    }
}
